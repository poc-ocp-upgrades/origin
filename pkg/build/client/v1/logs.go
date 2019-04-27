package v1

import (
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	buildv1 "github.com/openshift/api/build/v1"
)

type BuildLogInterface interface {
	Logs(name string, options buildv1.BuildLogOptions) *rest.Request
}

func NewBuildLogClient(c rest.Interface, ns string) BuildLogInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &buildLogs{client: c, ns: ns}
}

type buildLogs struct {
	client	rest.Interface
	ns	string
}

func (c *buildLogs) Logs(name string, options buildv1.BuildLogOptions) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Get().Namespace(c.ns).Resource("builds").Name(name).SubResource("log").VersionedParams(&options, legacyscheme.ParameterCodec)
}
