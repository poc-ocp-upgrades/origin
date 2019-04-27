package rsync

import (
	"io"
	"net/http"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

type portForwarder struct {
	Namespace	string
	PodName		string
	Client		kubernetes.Interface
	Config		*restclient.Config
	Out		io.Writer
	ErrOut		io.Writer
}

var _ forwarder = &portForwarder{}

func (f *portForwarder) ForwardPorts(ports []string, stopChan <-chan struct{}) error {
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
	req := f.Client.CoreV1().RESTClient().Post().Resource("pods").Namespace(f.Namespace).Name(f.PodName).SubResource("portforward")
	transport, upgrader, err := spdy.RoundTripperFor(f.Config)
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())
	readyChan := make(chan struct{})
	fw, err := portforward.New(dialer, ports, stopChan, readyChan, f.Out, f.ErrOut)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	go func() {
		errChan <- fw.ForwardPorts()
	}()
	select {
	case <-readyChan:
		return nil
	case err = <-errChan:
		return err
	}
}
func newPortForwarder(o *RsyncOptions) forwarder {
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
	return &portForwarder{Namespace: o.Namespace, PodName: o.PodName(), Client: o.Client, Config: o.Config, Out: o.Out, ErrOut: o.ErrOut}
}
