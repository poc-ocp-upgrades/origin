package testing

import (
	"fmt"
	goformat "fmt"
	pflag "github.com/spf13/pflag"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/cmd/kube-apiserver/app"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	"net"
	"os"
	goos "os"
	"path"
	"runtime"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type TearDownFunc func()
type TestServerInstanceOptions struct{ DisableStorageCleanup bool }
type TestServer struct {
	ClientConfig *restclient.Config
	ServerOpts   *options.ServerRunOptions
	TearDownFn   TearDownFunc
	TmpDir       string
}
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})
}

func NewDefaultTestServerOptions() *TestServerInstanceOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &TestServerInstanceOptions{DisableStorageCleanup: false}
}
func StartTestServer(t Logger, instanceOptions *TestServerInstanceOptions, customFlags []string, storageConfig *storagebackend.Config) (result TestServer, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if instanceOptions == nil {
		instanceOptions = NewDefaultTestServerOptions()
	}
	if !instanceOptions.DisableStorageCleanup {
		registry.TrackStorageCleanup()
	}
	stopCh := make(chan struct{})
	tearDown := func() {
		if !instanceOptions.DisableStorageCleanup {
			registry.CleanupStorage()
		}
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
	result.TmpDir, err = ioutil.TempDir("", "kubernetes-kube-apiserver")
	if err != nil {
		return result, fmt.Errorf("failed to create temp dir: %v", err)
	}
	fs := pflag.NewFlagSet("test", pflag.PanicOnError)
	s := options.NewServerRunOptions()
	for _, f := range s.Flags().FlagSets {
		fs.AddFlagSet(f)
	}
	s.InsecureServing.BindPort = 0
	s.SecureServing.Listener, s.SecureServing.BindPort, err = createLocalhostListenerOnFreePort()
	if err != nil {
		return result, fmt.Errorf("failed to create listener: %v", err)
	}
	s.SecureServing.ServerCert.CertDirectory = result.TmpDir
	s.SecureServing.ExternalAddress = s.SecureServing.Listener.Addr().(*net.TCPAddr).IP
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return result, fmt.Errorf("failed to get current file")
	}
	s.SecureServing.ServerCert.FixtureDirectory = path.Join(path.Dir(thisFile), "testdata")
	s.ServiceClusterIPRange.IP = net.IPv4(10, 0, 0, 0)
	s.ServiceClusterIPRange.Mask = net.CIDRMask(16, 32)
	s.Etcd.StorageConfig = *storageConfig
	s.APIEnablement.RuntimeConfig.Set("api/all=true")
	fs.Parse(customFlags)
	completedOptions, err := app.Complete(s)
	if err != nil {
		return result, fmt.Errorf("failed to set default ServerRunOptions: %v", err)
	}
	t.Logf("runtime-config=%v", completedOptions.APIEnablement.RuntimeConfig)
	t.Logf("Starting kube-apiserver on port %d...", s.SecureServing.BindPort)
	server, err := app.CreateServerChain(completedOptions, stopCh)
	if err != nil {
		return result, fmt.Errorf("failed to create server chain: %v", err)
	}
	go func(stopCh <-chan struct{}) {
		if err := server.PrepareRun().Run(stopCh); err != nil {
			t.Errorf("kube-apiserver failed run: %v", err)
		}
	}(stopCh)
	t.Logf("Waiting for /healthz to be ok...")
	client, err := kubernetes.NewForConfig(server.LoopbackClientConfig)
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
	result.ClientConfig = server.LoopbackClientConfig
	result.ServerOpts = s
	result.TearDownFn = tearDown
	return result, nil
}
func StartTestServerOrDie(t Logger, instanceOptions *TestServerInstanceOptions, flags []string, storageConfig *storagebackend.Config) *TestServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := StartTestServer(t, instanceOptions, flags, storageConfig)
	if err == nil {
		return &result
	}
	t.Fatalf("failed to launch server: %v", err)
	return nil
}
func createLocalhostListenerOnFreePort() (net.Listener, int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ln, err := net.Listen("tcp", "127.0.0.1:0")
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
