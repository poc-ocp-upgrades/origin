package controller

import (
	"container/heap"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

func TestScheduler(t *testing.T) {
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
	keys := []string{}
	s := newScheduler(2, flowcontrol.NewFakeAlwaysRateLimiter(), func(key, value interface{}) {
		keys = append(keys, key.(string))
	})
	for i := 0; i < 6; i++ {
		s.RunOnce()
		if len(keys) > 0 {
			t.Fatal(keys)
		}
		if s.position != (i+1)%3 {
			t.Fatal(s.position)
		}
	}
	s.Add("first", "test")
	found := false
	for i, buckets := range s.buckets {
		if _, ok := buckets["first"]; ok {
			found = true
		} else {
			continue
		}
		if i == s.position {
			t.Fatal("should not insert into current bucket")
		}
	}
	if !found {
		t.Fatal("expected to find key in a bucket")
	}
	s.Delay("first")
	for i := 0; i < 2; i++ {
		s.RunOnce()
	}
	if len(keys) != 0 {
		t.Fatal(keys)
	}
	s.RunOnce()
	if !reflect.DeepEqual(keys, []string{"first"}) {
		t.Fatal(keys)
	}
}
func TestSchedulerAddAndDelay(t *testing.T) {
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
	s := newScheduler(3, flowcontrol.NewFakeAlwaysRateLimiter(), func(key, value interface{}) {
	})
	s.Add("first", "other")
	if s.buckets[3]["first"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Add("second", "other")
	if s.buckets[2]["second"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Add("third", "other")
	if s.buckets[1]["third"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Add("fourth", "other")
	if s.buckets[3]["fourth"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Add("fifth", "other")
	if s.buckets[2]["fifth"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Remove("third", "other")
	s.Add("sixth", "other")
	if s.buckets[1]["sixth"] != "other" {
		t.Fatalf("placed key in wrong bucket: %#v", s.buckets)
	}
	s.Delay("second")
	if s.buckets[3]["second"] != "other" {
		t.Fatalf("delay placed key in wrong bucket: %#v", s.buckets)
	}
	s.Delay("third")
	if _, ok := s.buckets[3]["third"]; ok {
		t.Fatalf("delay placed key in wrong bucket: %#v", s.buckets)
	}
	s.Delay("fourth")
	if s.buckets[3]["fourth"] != "other" {
		t.Fatalf("delay placed key in wrong bucket: %#v", s.buckets)
	}
}
func TestSchedulerRemove(t *testing.T) {
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
	s := newScheduler(2, flowcontrol.NewFakeAlwaysRateLimiter(), func(key, value interface{}) {
	})
	s.Add("test", "other")
	if s.Remove("test", "value") {
		t.Fatal(s)
	}
	if !s.Remove("test", "other") {
		t.Fatal(s)
	}
	if s.Len() != 0 {
		t.Fatal(s)
	}
	s.Add("test", "other")
	s.Add("test", "new")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	if s.Remove("test", "other") {
		t.Fatal(s)
	}
	if !s.Remove("test", "new") {
		t.Fatal(s)
	}
}

type int64Heap []int64

var _ heap.Interface = &int64Heap{}

func (h int64Heap) Len() int {
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
	return len(h)
}
func (h int64Heap) Less(i, j int) bool {
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
	return h[i] < h[j]
}
func (h int64Heap) Swap(i, j int) {
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
	h[i], h[j] = h[j], h[i]
}
func (h *int64Heap) Push(x interface{}) {
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
	*h = append(*h, x.(int64))
}
func (h *int64Heap) Pop() interface{} {
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
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type wallClock struct{}

var _ flowcontrol.Clock = &wallClock{}

func (c wallClock) Now() time.Time {
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
	return time.Now()
}
func (c wallClock) Sleep(d time.Duration) {
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
	time.Sleep(d)
}

type fakeClock struct {
	c	sync.Cond
	threads	int
	now	int64
	wake	int64Heap
}

var _ flowcontrol.Clock = &fakeClock{}

func newFakeClock(threads int) flowcontrol.Clock {
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
	return &fakeClock{threads: threads, c: sync.Cond{L: &sync.Mutex{}}}
}
func (c *fakeClock) Now() time.Time {
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
	c.c.L.Lock()
	defer c.c.L.Unlock()
	return time.Unix(0, c.now)
}
func (c *fakeClock) Sleep(d time.Duration) {
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
	c.c.L.Lock()
	defer c.c.L.Unlock()
	wake := c.now + int64(d)
	heap.Push(&c.wake, wake)
	if len(c.wake) == c.threads {
		c.now = heap.Pop(&c.wake).(int64)
		for len(c.wake) > 0 && c.wake[0] == c.now {
			heap.Pop(&c.wake)
		}
		c.c.Broadcast()
	}
	for c.now < wake {
		c.c.Wait()
	}
}
func TestSchedulerSanity(t *testing.T) {
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
	const (
		buckets	= 4
		items	= 10
	)
	clock := newFakeClock(2)
	limiter := flowcontrol.NewTokenBucketRateLimiterWithClock(1, 1, clock)
	m := map[int]int{}
	s := newScheduler(buckets, limiter, func(key, value interface{}) {
		fmt.Printf("%v: %v\n", clock.Now().UTC(), key)
		m[key.(int)]++
	})
	for i := 0; i < items; i++ {
		s.Add(i, nil)
	}
	go s.RunUntil(make(chan struct{}))
	clock.Sleep((2*buckets-1)*time.Second + 1)
	expected := map[int]int{}
	for i := 0; i < items; i++ {
		expected[i] = 2
	}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("m did not match expected: %#v\n", m)
	}
}
