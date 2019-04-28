package controller

import (
	"sync"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/flowcontrol"
)

type scheduler struct {
	handle		func(key, value interface{})
	position	int
	limiter		flowcontrol.RateLimiter
	mu		sync.Mutex
	buckets		[]bucket
}
type bucket map[interface{}]interface{}

func newScheduler(bucketCount int, bucketLimiter flowcontrol.RateLimiter, fn func(key, value interface{})) *scheduler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bucketCount++
	buckets := make([]bucket, bucketCount)
	for i := range buckets {
		buckets[i] = make(bucket)
	}
	return &scheduler{handle: fn, buckets: buckets, limiter: bucketLimiter}
}
func (s *scheduler) RunUntil(ch <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go utilwait.Until(s.RunOnce, 0, ch)
}
func (s *scheduler) RunOnce() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, value, last := s.next()
	if last {
		s.limiter.Accept()
		return
	}
	s.handle(key, value)
}
func (s *scheduler) at(inc int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (s.position + inc + len(s.buckets)) % len(s.buckets)
}
func (s *scheduler) next() (interface{}, interface{}, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	last := s.buckets[s.position]
	for k, v := range last {
		delete(last, k)
		s.buckets[s.at(-1)][k] = v
		return k, v, false
	}
	s.position = s.at(1)
	return nil, nil, true
}
func (s *scheduler) Add(key, value interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, bucket := range s.buckets {
		delete(bucket, key)
	}
	n := len(s.buckets)
	base := s.position + n
	target, least := 0, 0
	for i := n - 1; i > 0; i-- {
		position := (base + i) % n
		size := len(s.buckets[position])
		if size == 0 {
			target = position
			break
		}
		if size < least || least == 0 {
			target = position
			least = size
		}
	}
	s.buckets[target][key] = value
}
func (s *scheduler) Remove(key, value interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	match := true
	for _, bucket := range s.buckets {
		if value != nil {
			if old, ok := bucket[key]; ok && old != value {
				match = false
				continue
			}
		}
		delete(bucket, key)
	}
	return match
}
func (s *scheduler) Delay(key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	last := s.at(-1)
	for i, bucket := range s.buckets {
		if i == last {
			continue
		}
		if value, ok := bucket[key]; ok {
			delete(bucket, key)
			s.buckets[last][key] = value
		}
	}
}
func (s *scheduler) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for _, bucket := range s.buckets {
		count += len(bucket)
	}
	return count
}
func (s *scheduler) Map() map[interface{}]interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[interface{}]interface{})
	for _, bucket := range s.buckets {
		for k, v := range bucket {
			out[k] = v
		}
	}
	return out
}
