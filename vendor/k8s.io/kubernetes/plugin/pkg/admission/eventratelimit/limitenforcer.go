package eventratelimit

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/util/flowcontrol"
	api "k8s.io/kubernetes/pkg/apis/core"
	eventratelimitapi "k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit"
	"strings"
)

const (
	defaultCacheSize = 4096
)

type limitEnforcer struct {
	limitType eventratelimitapi.LimitType
	cache     cache
	keyFunc   func(admission.Attributes) string
}

func newLimitEnforcer(config eventratelimitapi.Limit, clock flowcontrol.Clock) (*limitEnforcer, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rateLimiterFactory := func() flowcontrol.RateLimiter {
		return flowcontrol.NewTokenBucketRateLimiterWithClock(float32(config.QPS), int(config.Burst), clock)
	}
	if config.Type == eventratelimitapi.ServerLimitType {
		return &limitEnforcer{limitType: config.Type, cache: &singleCache{rateLimiter: rateLimiterFactory()}, keyFunc: getServerKey}, nil
	}
	cacheSize := int(config.CacheSize)
	if cacheSize == 0 {
		cacheSize = defaultCacheSize
	}
	underlyingCache, err := lru.New(cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create lru cache: %v", err)
	}
	cache := &lruCache{rateLimiterFactory: rateLimiterFactory, cache: underlyingCache}
	var keyFunc func(admission.Attributes) string
	switch t := config.Type; t {
	case eventratelimitapi.NamespaceLimitType:
		keyFunc = getNamespaceKey
	case eventratelimitapi.UserLimitType:
		keyFunc = getUserKey
	case eventratelimitapi.SourceAndObjectLimitType:
		keyFunc = getSourceAndObjectKey
	default:
		return nil, fmt.Errorf("unknown event rate limit type: %v", t)
	}
	return &limitEnforcer{limitType: config.Type, cache: cache, keyFunc: keyFunc}, nil
}
func (enforcer *limitEnforcer) accept(attr admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := enforcer.keyFunc(attr)
	rateLimiter := enforcer.cache.get(key)
	allow := rateLimiter.TryAccept()
	if !allow {
		return fmt.Errorf("limit reached on type %v for key %v", enforcer.limitType, key)
	}
	return nil
}
func getServerKey(attr admission.Attributes) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ""
}
func getNamespaceKey(attr admission.Attributes) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return attr.GetNamespace()
}
func getUserKey(attr admission.Attributes) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	userInfo := attr.GetUserInfo()
	if userInfo == nil {
		return ""
	}
	return userInfo.GetName()
}
func getSourceAndObjectKey(attr admission.Attributes) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	object := attr.GetObject()
	if object == nil {
		return ""
	}
	event, ok := object.(*api.Event)
	if !ok {
		return ""
	}
	return strings.Join([]string{event.Source.Component, event.Source.Host, event.InvolvedObject.Kind, event.InvolvedObject.Namespace, event.InvolvedObject.Name, string(event.InvolvedObject.UID), event.InvolvedObject.APIVersion}, "")
}
