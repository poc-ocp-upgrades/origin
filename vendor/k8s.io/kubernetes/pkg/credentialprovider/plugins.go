package credentialprovider

import (
	"k8s.io/klog"
	"reflect"
	"sort"
	"sync"
)

var providersMutex sync.Mutex
var providers = make(map[string]DockerConfigProvider)

func RegisterCredentialProvider(name string, provider DockerConfigProvider) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	providersMutex.Lock()
	defer providersMutex.Unlock()
	_, found := providers[name]
	if found {
		klog.Fatalf("Credential provider %q was registered twice", name)
	}
	klog.V(4).Infof("Registered credential provider %q", name)
	providers[name] = provider
}
func NewDockerKeyring() DockerKeyring {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	keyring := &lazyDockerKeyring{Providers: make([]DockerConfigProvider, 0)}
	keys := reflect.ValueOf(providers).MapKeys()
	stringKeys := make([]string, len(keys))
	for ix := range keys {
		stringKeys[ix] = keys[ix].String()
	}
	sort.Strings(stringKeys)
	for _, key := range stringKeys {
		provider := providers[key]
		if provider.Enabled() {
			klog.V(4).Infof("Registering credential provider: %v", key)
			keyring.Providers = append(keyring.Providers, provider)
		}
	}
	return keyring
}
