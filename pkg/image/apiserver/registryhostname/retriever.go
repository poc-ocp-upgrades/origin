package registryhostname

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

type RegistryHostnameRetriever interface {
	InternalRegistryHostname() (string, bool)
	ExternalRegistryHostname() (string, bool)
}

func TestingRegistryHostnameRetriever(deprecatedDefaultRegistryEnvFn func() (string, bool), external, internal string) RegistryHostnameRetriever {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &defaultRegistryHostnameRetriever{deprecatedDefaultFn: deprecatedDefaultRegistryEnvFn, externalHostname: external, internalHostname: internal}
}
func DefaultRegistryHostnameRetriever(clientConfig *rest.Config, external, internal string) (RegistryHostnameRetriever, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	defaultRegistry := env("OPENSHIFT_DEFAULT_REGISTRY", "${DOCKER_REGISTRY_SERVICE_HOST}:${DOCKER_REGISTRY_SERVICE_PORT}")
	svcCache := newServiceResolverCache(kubeClient.CoreV1().Services(metav1.NamespaceDefault).Get)
	defaultRegistryFunc, err := svcCache.Defer(defaultRegistry)
	if err != nil {
		return nil, fmt.Errorf("OPENSHIFT_DEFAULT_REGISTRY variable is invalid %q: %v", defaultRegistry, err)
	}
	return &defaultRegistryHostnameRetriever{deprecatedDefaultFn: defaultRegistryFunc, externalHostname: external, internalHostname: internal}, nil
}
func env(key string, defaultValue string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

type defaultRegistryHostnameRetriever struct {
	deprecatedDefaultFn func() (string, bool)
	internalHostname    string
	externalHostname    string
}

func (r *defaultRegistryHostnameRetriever) InternalRegistryHostname() (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(r.internalHostname) > 0 {
		return r.internalHostname, true
	}
	if r.deprecatedDefaultFn != nil {
		return r.deprecatedDefaultFn()
	}
	return "", false
}
func (r *defaultRegistryHostnameRetriever) ExternalRegistryHostname() (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.externalHostname, len(r.externalHostname) > 0
}
