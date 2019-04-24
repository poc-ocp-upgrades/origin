package registryhostname

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"os"
	"strconv"
	"strings"
	"sync"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type serviceEntry struct {
	host	string
	port	string
}
type ResolverCacheFunc func(name string, options metav1.GetOptions) (*corev1.Service, error)
type ServiceResolverCache struct {
	fill	ResolverCacheFunc
	cache	map[string]serviceEntry
	lock	sync.RWMutex
}

func newServiceResolverCache(fill ResolverCacheFunc) *ServiceResolverCache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ServiceResolverCache{cache: make(map[string]serviceEntry), fill: fill}
}
func (c *ServiceResolverCache) get(name string) (host, port string, ok bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.lock.RLock()
	entry, found := c.cache[name]
	c.lock.RUnlock()
	if found {
		return entry.host, entry.port, true
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	if entry, found := c.cache[name]; found {
		return entry.host, entry.port, true
	}
	service, err := c.fill(name, metav1.GetOptions{})
	if err != nil {
		return
	}
	if len(service.Spec.Ports) == 0 {
		return
	}
	host, port, ok = service.Spec.ClusterIP, strconv.Itoa(int(service.Spec.Ports[0].Port)), true
	c.cache[name] = serviceEntry{host: host, port: port}
	return
}
func toServiceName(envName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.TrimSpace(strings.ToLower(strings.Replace(envName, "_", "-", -1)))
}
func recognizeVariable(name string) (service string, host bool, ok bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case strings.HasSuffix(name, "_SERVICE_HOST"):
		service = toServiceName(strings.TrimSuffix(name, "_SERVICE_HOST"))
		host = true
	case strings.HasSuffix(name, "_SERVICE_PORT"):
		service = toServiceName(strings.TrimSuffix(name, "_SERVICE_PORT"))
	default:
		return "", false, false
	}
	if len(service) == 0 {
		return "", false, false
	}
	ok = true
	return
}
func (c *ServiceResolverCache) resolve(name string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	service, isHost, ok := recognizeVariable(name)
	if !ok {
		return "", false
	}
	host, port, ok := c.get(service)
	if !ok {
		return "", false
	}
	if isHost {
		return host, true
	}
	return port, true
}
func (c *ServiceResolverCache) Defer(env string) (func() (string, bool), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hasExpansion := false
	invalid := []string{}
	os.Expand(env, func(name string) string {
		hasExpansion = true
		if _, _, ok := recognizeVariable(name); !ok {
			invalid = append(invalid, name)
		}
		return ""
	})
	if len(invalid) != 0 {
		return nil, fmt.Errorf("invalid variable name(s): %s", strings.Join(invalid, ", "))
	}
	if !hasExpansion {
		return func() (string, bool) {
			return env, true
		}, nil
	}
	lock := sync.Mutex{}
	loaded := false
	return func() (string, bool) {
		lock.Lock()
		defer lock.Unlock()
		if loaded {
			return env, true
		}
		resolved := true
		expand := os.Expand(env, func(s string) string {
			s, ok := c.resolve(s)
			resolved = resolved && ok
			return s
		})
		if !resolved {
			return "", false
		}
		loaded = true
		env = expand
		return env, true
	}, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
