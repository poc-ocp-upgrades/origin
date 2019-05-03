package gcp_credentials

import (
 "encoding/json"
 "io/ioutil"
 "net/http"
 "strings"
 "time"
 utilnet "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/credentialprovider"
)

const (
 metadataUrl              = "http://metadata.google.internal./computeMetadata/v1/"
 metadataAttributes       = metadataUrl + "instance/attributes/"
 dockerConfigKey          = metadataAttributes + "google-dockercfg"
 dockerConfigUrlKey       = metadataAttributes + "google-dockercfg-url"
 serviceAccounts          = metadataUrl + "instance/service-accounts/"
 metadataScopes           = metadataUrl + "instance/service-accounts/default/scopes"
 metadataToken            = metadataUrl + "instance/service-accounts/default/token"
 metadataEmail            = metadataUrl + "instance/service-accounts/default/email"
 storageScopePrefix       = "https://www.googleapis.com/auth/devstorage"
 cloudPlatformScopePrefix = "https://www.googleapis.com/auth/cloud-platform"
 defaultServiceAccount    = "default/"
)

var gceProductNameFile = "/sys/class/dmi/id/product_name"
var containerRegistryUrls = []string{"container.cloud.google.com", "gcr.io", "*.gcr.io"}
var metadataHeader = &http.Header{"Metadata-Flavor": []string{"Google"}}

type metadataProvider struct{ Client *http.Client }
type dockerConfigKeyProvider struct{ metadataProvider }
type dockerConfigUrlKeyProvider struct{ metadataProvider }
type containerRegistryProvider struct{ metadataProvider }

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 tr := utilnet.SetTransportDefaults(&http.Transport{})
 metadataHTTPClientTimeout := time.Second * 10
 httpClient := &http.Client{Transport: tr, Timeout: metadataHTTPClientTimeout}
 credentialprovider.RegisterCredentialProvider("google-dockercfg", &credentialprovider.CachingDockerConfigProvider{Provider: &dockerConfigKeyProvider{metadataProvider{Client: httpClient}}, Lifetime: 60 * time.Second})
 credentialprovider.RegisterCredentialProvider("google-dockercfg-url", &credentialprovider.CachingDockerConfigProvider{Provider: &dockerConfigUrlKeyProvider{metadataProvider{Client: httpClient}}, Lifetime: 60 * time.Second})
 credentialprovider.RegisterCredentialProvider("google-container-registry", &containerRegistryProvider{metadataProvider{Client: httpClient}})
}
func onGCEVM() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 data, err := ioutil.ReadFile(gceProductNameFile)
 if err != nil {
  klog.V(2).Infof("Error while reading product_name: %v", err)
  return false
 }
 name := strings.TrimSpace(string(data))
 return name == "Google" || name == "Google Compute Engine"
}
func (g *metadataProvider) Enabled() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return onGCEVM()
}
func (g *dockerConfigKeyProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (g *dockerConfigKeyProvider) Provide() credentialprovider.DockerConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if cfg, err := credentialprovider.ReadDockerConfigFileFromUrl(dockerConfigKey, g.Client, metadataHeader); err != nil {
  klog.Errorf("while reading 'google-dockercfg' metadata: %v", err)
 } else {
  return cfg
 }
 return credentialprovider.DockerConfig{}
}
func (g *dockerConfigUrlKeyProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (g *dockerConfigUrlKeyProvider) Provide() credentialprovider.DockerConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if url, err := credentialprovider.ReadUrl(dockerConfigUrlKey, g.Client, metadataHeader); err != nil {
  klog.Errorf("while reading 'google-dockercfg-url' metadata: %v", err)
 } else {
  if strings.HasPrefix(string(url), "http") {
   if cfg, err := credentialprovider.ReadDockerConfigFileFromUrl(string(url), g.Client, nil); err != nil {
    klog.Errorf("while reading 'google-dockercfg-url'-specified url: %s, %v", string(url), err)
   } else {
    return cfg
   }
  } else {
   klog.Errorf("Unsupported URL scheme: %s", string(url))
  }
 }
 return credentialprovider.DockerConfig{}
}
func runWithBackoff(f func() ([]byte, error)) []byte {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var backoff = 100 * time.Millisecond
 const maxBackoff = time.Minute
 for {
  value, err := f()
  if err == nil {
   return value
  }
  time.Sleep(backoff)
  backoff = backoff * 2
  if backoff > maxBackoff {
   backoff = maxBackoff
  }
 }
}
func (g *containerRegistryProvider) Enabled() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !onGCEVM() {
  return false
 }
 value := runWithBackoff(func() ([]byte, error) {
  value, err := credentialprovider.ReadUrl(serviceAccounts, g.Client, metadataHeader)
  if err != nil {
   klog.V(2).Infof("Failed to Get service accounts from gce metadata server: %v", err)
  }
  return value, err
 })
 defaultServiceAccountExists := false
 for _, sa := range strings.Split(string(value), "\n") {
  if strings.TrimSpace(sa) == defaultServiceAccount {
   defaultServiceAccountExists = true
   break
  }
 }
 if !defaultServiceAccountExists {
  klog.V(2).Infof("'default' service account does not exist. Found following service accounts: %q", string(value))
  return false
 }
 url := metadataScopes + "?alt=json"
 value = runWithBackoff(func() ([]byte, error) {
  value, err := credentialprovider.ReadUrl(url, g.Client, metadataHeader)
  if err != nil {
   klog.V(2).Infof("Failed to Get scopes in default service account from gce metadata server: %v", err)
  }
  return value, err
 })
 var scopes []string
 if err := json.Unmarshal(value, &scopes); err != nil {
  klog.Errorf("Failed to unmarshal scopes: %v", err)
  return false
 }
 for _, v := range scopes {
  if strings.HasPrefix(v, storageScopePrefix) || strings.HasPrefix(v, cloudPlatformScopePrefix) {
   return true
  }
 }
 klog.Warningf("Google container registry is disabled, no storage scope is available: %s", value)
 return false
}

type tokenBlob struct {
 AccessToken string `json:"access_token"`
}

func (g *containerRegistryProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (g *containerRegistryProvider) Provide() credentialprovider.DockerConfig {
 _logClusterCodePath()
 defer _logClusterCodePath()
 cfg := credentialprovider.DockerConfig{}
 tokenJsonBlob, err := credentialprovider.ReadUrl(metadataToken, g.Client, metadataHeader)
 if err != nil {
  klog.Errorf("while reading access token endpoint: %v", err)
  return cfg
 }
 email, err := credentialprovider.ReadUrl(metadataEmail, g.Client, metadataHeader)
 if err != nil {
  klog.Errorf("while reading email endpoint: %v", err)
  return cfg
 }
 var parsedBlob tokenBlob
 if err := json.Unmarshal([]byte(tokenJsonBlob), &parsedBlob); err != nil {
  klog.Errorf("while parsing json blob %s: %v", tokenJsonBlob, err)
  return cfg
 }
 entry := credentialprovider.DockerConfigEntry{Username: "_token", Password: parsedBlob.AccessToken, Email: string(email)}
 for _, k := range containerRegistryUrls {
  cfg[k] = entry
 }
 return cfg
}
