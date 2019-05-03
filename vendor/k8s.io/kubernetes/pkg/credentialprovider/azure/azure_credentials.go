package azure

import (
 "context"
 "io"
 "io/ioutil"
 "os"
 "time"
 "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2017-10-01/containerregistry"
 "github.com/Azure/go-autorest/autorest"
 "github.com/Azure/go-autorest/autorest/adal"
 "github.com/Azure/go-autorest/autorest/azure"
 "github.com/spf13/pflag"
 "k8s.io/klog"
 "sigs.k8s.io/yaml"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/azure/auth"
 "k8s.io/kubernetes/pkg/credentialprovider"
)

var flagConfigFile = pflag.String("azure-container-registry-config", "", "Path to the file containing Azure container registry configuration information.")

const dummyRegistryEmail = "name@contoso.com"

var containerRegistryUrls = []string{"*.azurecr.io", "*.azurecr.cn", "*.azurecr.de", "*.azurecr.us"}

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 credentialprovider.RegisterCredentialProvider("azure", &credentialprovider.CachingDockerConfigProvider{Provider: NewACRProvider(flagConfigFile), Lifetime: 1 * time.Minute})
}
func getContextWithCancel() (context.Context, context.CancelFunc) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return context.WithCancel(context.Background())
}

type RegistriesClient interface {
 List(ctx context.Context) ([]containerregistry.Registry, error)
}
type azRegistriesClient struct {
 client containerregistry.RegistriesClient
}

func newAzRegistriesClient(subscriptionID, endpoint string, token *adal.ServicePrincipalToken) *azRegistriesClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 registryClient := containerregistry.NewRegistriesClient(subscriptionID)
 registryClient.BaseURI = endpoint
 registryClient.Authorizer = autorest.NewBearerAuthorizer(token)
 return &azRegistriesClient{client: registryClient}
}
func (az *azRegistriesClient) List(ctx context.Context) ([]containerregistry.Registry, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 iterator, err := az.client.ListComplete(ctx)
 if err != nil {
  return nil, err
 }
 result := make([]containerregistry.Registry, 0)
 for ; iterator.NotDone(); err = iterator.Next() {
  if err != nil {
   return nil, err
  }
  result = append(result, iterator.Value())
 }
 return result, nil
}
func NewACRProvider(configFile *string) credentialprovider.DockerConfigProvider {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &acrProvider{file: configFile}
}

type acrProvider struct {
 file                  *string
 config                *auth.AzureAuthConfig
 environment           *azure.Environment
 registryClient        RegistriesClient
 servicePrincipalToken *adal.ServicePrincipalToken
}

func parseConfig(configReader io.Reader) (*auth.AzureAuthConfig, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var config auth.AzureAuthConfig
 if configReader == nil {
  return &config, nil
 }
 configContents, err := ioutil.ReadAll(configReader)
 if err != nil {
  return nil, err
 }
 err = yaml.Unmarshal(configContents, &config)
 if err != nil {
  return nil, err
 }
 return &config, nil
}
func (a *acrProvider) loadConfig(rdr io.Reader) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 a.config, err = parseConfig(rdr)
 if err != nil {
  klog.Errorf("Failed to load azure credential file: %v", err)
 }
 a.environment, err = auth.ParseAzureEnvironment(a.config.Cloud)
 if err != nil {
  return err
 }
 return nil
}
func (a *acrProvider) Enabled() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a.file == nil || len(*a.file) == 0 {
  klog.V(5).Infof("Azure config unspecified, disabling")
  return false
 }
 f, err := os.Open(*a.file)
 if err != nil {
  klog.Errorf("Failed to load config from file: %s", *a.file)
  return false
 }
 defer f.Close()
 err = a.loadConfig(f)
 if err != nil {
  klog.Errorf("Failed to load config from file: %s", *a.file)
  return false
 }
 a.servicePrincipalToken, err = auth.GetServicePrincipalToken(a.config, a.environment)
 if err != nil {
  klog.Errorf("Failed to create service principal token: %v", err)
  return false
 }
 a.registryClient = newAzRegistriesClient(a.config.SubscriptionID, a.environment.ResourceManagerEndpoint, a.servicePrincipalToken)
 return true
}
func (a *acrProvider) Provide() credentialprovider.DockerConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cfg := credentialprovider.DockerConfig{}
 ctx, cancel := getContextWithCancel()
 defer cancel()
 if a.config.UseManagedIdentityExtension {
  klog.V(4).Infof("listing registries")
  result, err := a.registryClient.List(ctx)
  if err != nil {
   klog.Errorf("Failed to list registries: %v", err)
   return cfg
  }
  for ix := range result {
   loginServer := getLoginServer(result[ix])
   klog.V(2).Infof("loginServer: %s", loginServer)
   cred, err := getACRDockerEntryFromARMToken(a, loginServer)
   if err != nil {
    continue
   }
   cfg[loginServer] = *cred
  }
 } else {
  for _, url := range containerRegistryUrls {
   cred := &credentialprovider.DockerConfigEntry{Username: a.config.AADClientID, Password: a.config.AADClientSecret, Email: dummyRegistryEmail}
   cfg[url] = *cred
  }
 }
 return cfg
}
func getLoginServer(registry containerregistry.Registry) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return *(*registry.RegistryProperties).LoginServer
}
func getACRDockerEntryFromARMToken(a *acrProvider, loginServer string) (*credentialprovider.DockerConfigEntry, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 armAccessToken := a.servicePrincipalToken.OAuthToken()
 klog.V(4).Infof("discovering auth redirects for: %s", loginServer)
 directive, err := receiveChallengeFromLoginServer(loginServer)
 if err != nil {
  klog.Errorf("failed to receive challenge: %s", err)
  return nil, err
 }
 klog.V(4).Infof("exchanging an acr refresh_token")
 registryRefreshToken, err := performTokenExchange(loginServer, directive, a.config.TenantID, armAccessToken)
 if err != nil {
  klog.Errorf("failed to perform token exchange: %s", err)
  return nil, err
 }
 klog.V(4).Infof("adding ACR docker config entry for: %s", loginServer)
 return &credentialprovider.DockerConfigEntry{Username: dockerTokenLoginUsernameGUID, Password: registryRefreshToken, Email: dummyRegistryEmail}, nil
}
func (a *acrProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
