package credentialprovider

import (
	"k8s.io/klog"
	"os"
	"reflect"
	"sync"
	"time"
)

type DockerConfigProvider interface {
	Enabled() bool
	Provide() DockerConfig
	LazyProvide() *DockerConfigEntry
}

func LazyProvide(creds LazyAuthConfiguration) AuthConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if creds.Provider != nil {
		entry := *creds.Provider.LazyProvide()
		return DockerConfigEntryToLazyAuthConfiguration(entry).AuthConfig
	}
	return creds.AuthConfig
}

type defaultDockerConfigProvider struct{}

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	RegisterCredentialProvider(".dockercfg", &CachingDockerConfigProvider{Provider: &defaultDockerConfigProvider{}, Lifetime: 5 * time.Minute})
}

type CachingDockerConfigProvider struct {
	Provider          DockerConfigProvider
	Lifetime          time.Duration
	cacheDockerConfig DockerConfig
	expiration        time.Time
	mu                sync.Mutex
}

func (d *defaultDockerConfigProvider) Enabled() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (d *defaultDockerConfigProvider) Provide() DockerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg, err := ReadDockerConfigFile(); err == nil {
		return cfg
	} else if !os.IsNotExist(err) {
		klog.V(4).Infof("Unable to parse Docker config file: %v", err)
	}
	return DockerConfig{}
}
func (d *defaultDockerConfigProvider) LazyProvide() *DockerConfigEntry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (d *CachingDockerConfigProvider) Enabled() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.Provider.Enabled()
}
func (d *CachingDockerConfigProvider) LazyProvide() *DockerConfigEntry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (d *CachingDockerConfigProvider) Provide() DockerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	d.mu.Lock()
	defer d.mu.Unlock()
	if time.Now().Before(d.expiration) {
		return d.cacheDockerConfig
	}
	klog.V(2).Infof("Refreshing cache for provider: %v", reflect.TypeOf(d.Provider).String())
	d.cacheDockerConfig = d.Provider.Provide()
	d.expiration = time.Now().Add(d.Lifetime)
	return d.cacheDockerConfig
}
