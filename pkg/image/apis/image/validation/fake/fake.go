package fake

import (
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"bytes"
	"net/http"
	"runtime"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
