package eventratelimit

import (
	"github.com/hashicorp/golang-lru"
	"k8s.io/client-go/util/flowcontrol"
)

type cache interface {
	get(key interface{}) flowcontrol.RateLimiter
}
type singleCache struct{ rateLimiter flowcontrol.RateLimiter }

func (c *singleCache) get(key interface{}) flowcontrol.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.rateLimiter
}

type lruCache struct {
	rateLimiterFactory func() flowcontrol.RateLimiter
	cache              *lru.Cache
}

func (c *lruCache) get(key interface{}) flowcontrol.RateLimiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value, found := c.cache.Get(key)
	if !found {
		rateLimter := c.rateLimiterFactory()
		c.cache.Add(key, rateLimter)
		return rateLimter
	}
	return value.(flowcontrol.RateLimiter)
}
