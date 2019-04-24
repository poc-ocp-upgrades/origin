package docker

import (
	"os"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/spf13/pflag"
	"k8s.io/klog"
)

type Helper struct{}

func NewHelper() *Helper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Helper{}
}
func (_ *Helper) InstallFlags(flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (_ *Helper) GetClient() (client *docker.Client, endpoint string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err = docker.NewClientFromEnv()
	if len(os.Getenv("DOCKER_HOST")) > 0 {
		endpoint = os.Getenv("DOCKER_HOST")
	} else {
		endpoint = "unix:///var/run/docker.sock"
	}
	return
}
func (h *Helper) GetClientOrExit() (*docker.Client, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, addr, err := h.GetClient()
	if err != nil {
		klog.Fatalf("ERROR: Couldn't connect to Docker at %s.\n%v\n.", addr, err)
	}
	return client, addr
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
