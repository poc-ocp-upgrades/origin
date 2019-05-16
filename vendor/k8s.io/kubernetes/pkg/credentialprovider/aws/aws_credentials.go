package credentials

import (
	"encoding/base64"
	"fmt"
	goformat "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/credentialprovider"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const awsChinaRegionPrefix = "cn-"
const awsStandardDNSSuffix = "amazonaws.com"
const awsChinaDNSSuffix = "amazonaws.com.cn"
const registryURLTemplate = "*.dkr.ecr.%s.%s"

func awsHandlerLogger(req *request.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	service := req.ClientInfo.ServiceName
	region := req.Config.Region
	name := "?"
	if req.Operation != nil {
		name = req.Operation.Name
	}
	klog.V(3).Infof("AWS request: %s:%s in %s", service, name, *region)
}

type tokenGetter interface {
	GetAuthorizationToken(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)
}
type ecrTokenGetter struct{ svc *ecr.ECR }

func (p *ecrTokenGetter) GetAuthorizationToken(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return p.svc.GetAuthorizationToken(input)
}

type lazyEcrProvider struct {
	region         string
	regionURL      string
	actualProvider *credentialprovider.CachingDockerConfigProvider
}

var _ credentialprovider.DockerConfigProvider = &lazyEcrProvider{}

type ecrProvider struct {
	region    string
	regionURL string
	getter    tokenGetter
}

var _ credentialprovider.DockerConfigProvider = &ecrProvider{}

func registryURL(region string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dnsSuffix := awsStandardDNSSuffix
	if strings.HasPrefix(region, awsChinaRegionPrefix) {
		dnsSuffix = awsChinaDNSSuffix
	}
	return fmt.Sprintf(registryURLTemplate, region, dnsSuffix)
}
func RegisterCredentialsProvider(region string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("registering credentials provider for AWS region %q", region)
	credentialprovider.RegisterCredentialProvider("aws-ecr-"+region, &lazyEcrProvider{region: region, regionURL: registryURL(region)})
}
func (p *lazyEcrProvider) Enabled() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (p *lazyEcrProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.actualProvider == nil {
		klog.V(2).Infof("Creating ecrProvider for %s", p.region)
		p.actualProvider = &credentialprovider.CachingDockerConfigProvider{Provider: newEcrProvider(p.region, nil), Lifetime: 11*time.Hour + 55*time.Minute}
		if !p.actualProvider.Enabled() {
			return nil
		}
	}
	entry := p.actualProvider.Provide()[p.regionURL]
	return &entry
}
func (p *lazyEcrProvider) Provide() credentialprovider.DockerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	entry := credentialprovider.DockerConfigEntry{Provider: p}
	cfg := credentialprovider.DockerConfig{}
	cfg[p.regionURL] = entry
	return cfg
}
func newEcrProvider(region string, getter tokenGetter) *ecrProvider {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ecrProvider{region: region, regionURL: registryURL(region), getter: getter}
}
func (p *ecrProvider) Enabled() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.region == "" {
		klog.Errorf("Called ecrProvider.Enabled() with no region set")
		return false
	}
	getter := &ecrTokenGetter{svc: ecr.New(session.New(&aws.Config{Credentials: nil, Region: &p.region}))}
	getter.svc.Handlers.Sign.PushFrontNamed(request.NamedHandler{Name: "k8s/logger", Fn: awsHandlerLogger})
	p.getter = getter
	return true
}
func (p *ecrProvider) LazyProvide() *credentialprovider.DockerConfigEntry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (p *ecrProvider) Provide() credentialprovider.DockerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := credentialprovider.DockerConfig{}
	params := &ecr.GetAuthorizationTokenInput{}
	output, err := p.getter.GetAuthorizationToken(params)
	if err != nil {
		klog.Errorf("while requesting ECR authorization token %v", err)
		return cfg
	}
	if output == nil {
		klog.Errorf("Got back no ECR token")
		return cfg
	}
	for _, data := range output.AuthorizationData {
		if data.ProxyEndpoint != nil && data.AuthorizationToken != nil {
			decodedToken, err := base64.StdEncoding.DecodeString(aws.StringValue(data.AuthorizationToken))
			if err != nil {
				klog.Errorf("while decoding token for endpoint %v %v", data.ProxyEndpoint, err)
				return cfg
			}
			parts := strings.SplitN(string(decodedToken), ":", 2)
			user := parts[0]
			password := parts[1]
			entry := credentialprovider.DockerConfigEntry{Username: user, Password: password, Email: "not@val.id"}
			klog.V(3).Infof("Adding credentials for user %s in %s", user, p.region)
			cfg[p.regionURL] = entry
		}
	}
	return cfg
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
