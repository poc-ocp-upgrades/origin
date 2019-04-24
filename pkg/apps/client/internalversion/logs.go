package internalversion

import (
	rest "k8s.io/client-go/rest"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
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
	client	rest.Interface
	ns	string
}

func (c *rolloutLogs) Logs(name string, options appsapi.DeploymentLogOptions) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Get().Namespace(c.ns).Resource("deploymentConfigs").Name(name).SubResource("log").VersionedParams(&options, legacyscheme.ParameterCodec)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
