package config

import (
	goformat "fmt"
	apiserver "k8s.io/apiserver/pkg/server"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Config struct {
	OpenShiftContext     OpenShiftContext
	ComponentConfig      kubectrlmgrconfig.KubeControllerManagerConfiguration
	SecureServing        *apiserver.SecureServingInfo
	LoopbackClientConfig *restclient.Config
	InsecureServing      *apiserver.DeprecatedInsecureServingInfo
	Authentication       apiserver.AuthenticationInfo
	Authorization        apiserver.AuthorizationInfo
	Client               *clientset.Clientset
	LeaderElectionClient *clientset.Clientset
	Kubeconfig           *restclient.Config
	EventRecorder        record.EventRecorder
}
type completedConfig struct{ *Config }
type CompletedConfig struct{ *completedConfig }

func (c *Config) Complete() *CompletedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cc := completedConfig{c}
	apiserver.AuthorizeClientBearerToken(c.LoopbackClientConfig, &c.Authentication, &c.Authorization)
	return &CompletedConfig{&cc}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
