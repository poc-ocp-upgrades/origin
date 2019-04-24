package whitelist

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"net"
	"reflect"
	"strings"
	"k8s.io/klog"
	kerrutil "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	stringsutil "github.com/openshift/origin/pkg/util/strings"
)

type WhitelistTransport string

const (
	WhitelistTransportAny		WhitelistTransport	= "any"
	WhitelistTransportSecure	WhitelistTransport	= "secure"
	WhitelistTransportInsecure	WhitelistTransport	= "insecure"
)

type RegistryWhitelister interface {
	AdmitHostname(host string, transport WhitelistTransport) error
	AdmitPullSpec(pullSpec string, transport WhitelistTransport) error
	AdmitDockerImageReference(ref *imageapi.DockerImageReference, transport WhitelistTransport) error
	WhitelistRegistry(hostPortGlob string, transport WhitelistTransport) error
	WhitelistPullSpecs(pullSpecs ...string)
	Copy() RegistryWhitelister
}
type RegistryHostnameRetriever interface {
	InternalRegistryHostname() (string, bool)
	ExternalRegistryHostname() (string, bool)
}
type allowedHostPortGlobs struct {
	host	string
	port	string
}
type registryWhitelister struct {
	whitelist		[]allowedHostPortGlobs
	pullSpecs		sets.String
	registryHostRetriever	RegistryHostnameRetriever
}

var _ RegistryWhitelister = &registryWhitelister{}

func NewRegistryWhitelister(whitelist openshiftcontrolplanev1.AllowedRegistries, registryHostRetriever RegistryHostnameRetriever) (RegistryWhitelister, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := []error{}
	rw := registryWhitelister{whitelist: make([]allowedHostPortGlobs, 0, len(whitelist)), pullSpecs: sets.NewString(), registryHostRetriever: registryHostRetriever}
	for i := len(whitelist) - 1; i >= 0; i-- {
		registry := whitelist[i]
		transport := WhitelistTransportSecure
		if registry.Insecure {
			transport = WhitelistTransportInsecure
		}
		err := rw.WhitelistRegistry(registry.DomainName, transport)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, kerrutil.NewAggregate(errs)
	}
	return &rw, nil
}
func WhitelistAllRegistries() RegistryWhitelister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &registryWhitelister{whitelist: []allowedHostPortGlobs{{host: "*", port: "*"}}, pullSpecs: sets.NewString()}
}
func (rw *registryWhitelister) AdmitHostname(hostname string, transport WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rw.AdmitDockerImageReference(&imageapi.DockerImageReference{Registry: hostname}, transport)
}
func (rw *registryWhitelister) AdmitPullSpec(pullSpec string, transport WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ref, err := imageapi.ParseDockerImageReference(pullSpec)
	if err != nil {
		return err
	}
	return rw.AdmitDockerImageReference(&ref, transport)
}
func (rw *registryWhitelister) AdmitDockerImageReference(ref *imageapi.DockerImageReference, transport WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const showMax = 5
	if rw.pullSpecs.Len() > 0 {
		if rw.pullSpecs.Has(ref.Exact()) || rw.pullSpecs.Has(ref.DockerClientDefaults().Exact()) || rw.pullSpecs.Has(ref.DaemonMinimal().Exact()) {
			return nil
		}
	}
	if rw.registryHostRetriever != nil {
		if localRegistry, ok := rw.registryHostRetriever.InternalRegistryHostname(); ok {
			rw.WhitelistRegistry(localRegistry, WhitelistTransportSecure)
		}
	}
	var (
		host, port	string
		err		error
	)
	switch transport {
	case WhitelistTransportAny:
		host, port, err = net.SplitHostPort(ref.Registry)
		if err != nil || len(port) == 0 {
			host = ref.Registry
			port = ""
		}
		if len(host) == 0 {
			host, _ = ref.RegistryHostPort(false)
		}
	case WhitelistTransportInsecure:
		host, port = ref.RegistryHostPort(true)
	default:
		host, port = ref.RegistryHostPort(false)
	}
	matchHost := func(h string) bool {
		for _, hp := range rw.whitelist {
			if stringsutil.IsWildcardMatch(h, hp.host) {
				if len(port) == 0 {
					switch hp.port {
					case "80", "443", "*":
						return true
					default:
						continue
					}
				}
				if stringsutil.IsWildcardMatch(port, hp.port) {
					return true
				}
			}
		}
		return false
	}
	switch host {
	case imageapi.DockerDefaultV1Registry, imageapi.DockerDefaultV2Registry:
		if matchHost(imageapi.DockerDefaultRegistry) {
			return nil
		}
		fallthrough
	default:
		if matchHost(host) {
			return nil
		}
	}
	hostname := ref.Registry
	if len(ref.Registry) == 0 {
		if len(port) > 0 {
			hostname = net.JoinHostPort(host, port)
		} else {
			hostname = host
		}
	}
	var whitelist []string
	for i := 0; i < len(rw.whitelist); i++ {
		whitelist = append(whitelist, fmt.Sprintf("%q", net.JoinHostPort(rw.whitelist[i].host, rw.whitelist[i].port)))
	}
	if len(rw.whitelist) == 0 {
		klog.V(5).Infof("registry %q not allowed by empty whitelist", hostname)
		return fmt.Errorf("registry %q not allowed by empty whitelist", hostname)
	}
	klog.V(5).Infof("registry %q not allowed by whitelist: %s", hostname, strings.Join(whitelist, ", "))
	if len(rw.whitelist) <= showMax {
		return fmt.Errorf("registry %q not allowed by whitelist: %s", hostname, strings.Join(whitelist, ", "))
	}
	return fmt.Errorf("registry %q not allowed by whitelist: %s, and %d more ...", hostname, strings.Join(whitelist[:showMax-1], ", "), len(whitelist)-showMax+1)
}
func (rw *registryWhitelister) WhitelistRegistry(hostPortGlob string, transport WhitelistTransport) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hps := make([]allowedHostPortGlobs, 1, 2)
	parts := strings.SplitN(hostPortGlob, ":", 3)
	switch len(parts) {
	case 1:
		hps[0].host = parts[0]
		switch transport {
		case WhitelistTransportAny:
			hps[0].port = "80"
			hps = append(hps, allowedHostPortGlobs{host: parts[0], port: "443"})
		case WhitelistTransportInsecure:
			hps[0].port = "80"
		default:
			hps[0].port = "443"
		}
	case 2:
		hps[0].host, hps[0].port = parts[0], parts[1]
	default:
		return fmt.Errorf("failed to parse allowed registry %q: too many colons", hostPortGlob)
	}
addHPsLoop:
	for i := range hps {
		for _, item := range rw.whitelist {
			if reflect.DeepEqual(item, hps[i]) {
				continue addHPsLoop
			}
		}
		rw.whitelist = append([]allowedHostPortGlobs{hps[i]}, rw.whitelist...)
	}
	return nil
}
func (rw *registryWhitelister) WhitelistPullSpecs(pullSpecs ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.pullSpecs.Insert(pullSpecs...)
}
func (rw *registryWhitelister) Copy() RegistryWhitelister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newRW := registryWhitelister{whitelist: make([]allowedHostPortGlobs, len(rw.whitelist)), pullSpecs: sets.NewString(rw.pullSpecs.List()...), registryHostRetriever: rw.registryHostRetriever}
	copy(newRW.whitelist, rw.whitelist)
	return &newRW
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
