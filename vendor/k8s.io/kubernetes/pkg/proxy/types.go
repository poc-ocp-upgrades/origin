package proxy

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type ProxyProvider interface {
	Sync()
	SyncLoop()
}
type ServicePortName struct {
	types.NamespacedName
	Port string
}

func (spn ServicePortName) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s:%s", spn.NamespacedName.String(), spn.Port)
}

type ServicePort interface {
	String() string
	ClusterIPString() string
	ExternalIPStrings() []string
	GetProtocol() v1.Protocol
	GetHealthCheckNodePort() int
	GetNodePort() int
}
type Endpoint interface {
	String() string
	GetIsLocal() bool
	IP() string
	Port() (int, error)
	Equal(Endpoint) bool
}
type ServiceEndpoint struct {
	Endpoint        string
	ServicePortName ServicePortName
}
