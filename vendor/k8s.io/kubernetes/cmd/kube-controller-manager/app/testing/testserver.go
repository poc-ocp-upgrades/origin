package testing

import (
	"fmt"
	goformat "fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kubecontrollerconfig "k8s.io/kubernetes/cmd/kube-controller-manager/app/config"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	"net"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type TearDownFunc func()
type TestServer struct {
	LoopbackClientConfig *restclient.Config
	Options              *options.KubeControllerManagerOptions
	Config               *kubecontrollerconfig.Config
	TearDownFn           TearDownFunc
	TmpDir               string
}
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})
}

func StartTestServer(t Logger, customFlags []string) (result TestServer, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	stopCh := make(chan struct{})
	tearDown := func() {
		close(stopCh)
		if len(result.TmpDir) != 0 {
			os.RemoveAll(result.TmpDir)
		}
	}
	defer func() {
		if result.TearDownFn == nil {
			tearDown()
		}
	}()
	result.TmpDir, err = ioutil.TempDir("", "kube-controller-manager")
	if err != nil {
		return result, fmt.Errorf("failed to create temp dir: %v", err)
	}
	fs := pflag.NewFlagSet("test", pflag.PanicOnError)
	s, err := options.NewKubeControllerManagerOptions()
	if err != nil {
		return TestServer{}, err
	}
	all, disabled := app.KnownControllers(), app.ControllersDisabledByDefault.List()
	namedFlagSets := s.Flags(all, disabled)
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}
	fs.Parse(customFlags)
	if s.SecureServing.BindPort != 0 {
		s.SecureServing.Listener, s.SecureServing.BindPort, err = createListenerOnFreePort()
		if err != nil {
			return result, fmt.Errorf("failed to create listener: %v", err)
		}
		s.SecureServing.ServerCert.CertDirectory = result.TmpDir
		t.Logf("kube-controller-manager will listen securely on port %d...", s.SecureServing.BindPort)
	}
	if s.InsecureServing.BindPort != 0 {
		s.InsecureServing.Listener, s.InsecureServing.BindPort, err = createListenerOnFreePort()
		if err != nil {
			return result, fmt.Errorf("failed to create listener: %v", err)
		}
		t.Logf("kube-controller-manager will listen insecurely on port %d...", s.InsecureServing.BindPort)
	}
	config, err := s.Config(all, disabled)
	if err != nil {
		return result, fmt.Errorf("failed to create config from options: %v", err)
	}
	go func(stopCh <-chan struct{}) {
		if err := app.Run(config.Complete(), stopCh); err != nil {
			t.Errorf("kube-apiserver failed run: %v", err)
		}
	}(stopCh)
	t.Logf("Waiting for /healthz to be ok...")
	client, err := kubernetes.NewForConfig(config.LoopbackClientConfig)
	if err != nil {
		return result, fmt.Errorf("failed to create a client: %v", err)
	}
	err = wait.Poll(100*time.Millisecond, 30*time.Second, func() (bool, error) {
		result := client.CoreV1().RESTClient().Get().AbsPath("/healthz").Do()
		status := 0
		result.StatusCode(&status)
		if status == 200 {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return result, fmt.Errorf("failed to wait for /healthz to return ok: %v", err)
	}
	result.LoopbackClientConfig = config.LoopbackClientConfig
	result.Options = s
	result.Config = config
	result.TearDownFn = tearDown
	return result, nil
}
func StartTestServerOrDie(t Logger, flags []string) *TestServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := StartTestServer(t, flags)
	if err == nil {
		return &result
	}
	t.Fatalf("failed to launch server: %v", err)
	return nil
}
func createListenerOnFreePort() (net.Listener, int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, 0, err
	}
	tcpAddr, ok := ln.Addr().(*net.TCPAddr)
	if !ok {
		ln.Close()
		return nil, 0, fmt.Errorf("invalid listen address: %q", ln.Addr().String())
	}
	return ln, tcpAddr.Port, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
