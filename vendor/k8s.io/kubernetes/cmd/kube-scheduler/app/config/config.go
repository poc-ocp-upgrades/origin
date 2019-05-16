package config

import (
	goformat "fmt"
	apiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/record"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Config struct {
	ComponentConfig        kubeschedulerconfig.KubeSchedulerConfiguration
	LoopbackClientConfig   *restclient.Config
	InsecureServing        *apiserver.DeprecatedInsecureServingInfo
	InsecureMetricsServing *apiserver.DeprecatedInsecureServingInfo
	Authentication         apiserver.AuthenticationInfo
	Authorization          apiserver.AuthorizationInfo
	SecureServing          *apiserver.SecureServingInfo
	Client                 clientset.Interface
	InformerFactory        informers.SharedInformerFactory
	PodInformer            coreinformers.PodInformer
	EventClient            v1core.EventsGetter
	Recorder               record.EventRecorder
	Broadcaster            record.EventBroadcaster
	LeaderElection         *leaderelection.LeaderElectionConfig
}
type completedConfig struct{ *Config }
type CompletedConfig struct{ *completedConfig }

func (c *Config) Complete() CompletedConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cc := completedConfig{c}
	if c.InsecureServing != nil {
		c.InsecureServing.Name = "healthz"
	}
	if c.InsecureMetricsServing != nil {
		c.InsecureMetricsServing.Name = "metrics"
	}
	apiserver.AuthorizeClientBearerToken(c.LoopbackClientConfig, &c.Authentication, &c.Authorization)
	return CompletedConfig{&cc}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
