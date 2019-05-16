package https

import (
	goformat "fmt"
	"io/ioutil"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/kubernetes/cmd/kubeadm/app/discovery/file"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func RetrieveValidatedConfigInfo(httpsURL, clustername string) (*clientcmdapi.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client := &http.Client{Transport: netutil.SetOldTransportDefaults(&http.Transport{})}
	response, err := client.Get(httpsURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	kubeconfig, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.Load(kubeconfig)
	if err != nil {
		return nil, err
	}
	return file.ValidateConfigInfo(config, clustername)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
