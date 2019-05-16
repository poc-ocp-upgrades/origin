package v1

import (
	goformat "fmt"
	buildv1 "github.com/openshift/api/build/v1"
	"io"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type BuildInstantiateBinaryInterface interface {
	InstantiateBinary(name string, options *buildv1.BinaryBuildRequestOptions, r io.Reader) (*buildv1.Build, error)
}

func NewBuildInstantiateBinaryClient(c rest.Interface, ns string) BuildInstantiateBinaryInterface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &buildInstatiateBinary{client: c, ns: ns}
}

type buildInstatiateBinary struct {
	client rest.Interface
	ns     string
}

func (c *buildInstatiateBinary) InstantiateBinary(name string, options *buildv1.BinaryBuildRequestOptions, r io.Reader) (*buildv1.Build, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := &buildv1.Build{}
	err := c.client.Post().Namespace(c.ns).Resource("buildconfigs").Name(name).SubResource("instantiatebinary").Body(r).VersionedParams(options, legacyscheme.ParameterCodec).Do().Into(result)
	return result, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
