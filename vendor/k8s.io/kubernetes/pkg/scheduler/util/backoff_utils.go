package util

import (
	"fmt"
	goformat "fmt"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"sync/atomic"
	"time"
	gotime "time"
)

type clock interface{ Now() time.Time }
type realClock struct{}

func (realClock) Now() time.Time {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return time.Now()
}

type BackoffEntry struct {
	backoff     time.Duration
	lastUpdate  time.Time
	reqInFlight int32
}

func (b *BackoffEntry) tryLock() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return atomic.CompareAndSwapInt32(&b.reqInFlight, 0, 1)
}
func (b *BackoffEntry) unlock() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !atomic.CompareAndSwapInt32(&b.reqInFlight, 1, 0) {
		panic(fmt.Sprintf("unexpected state on unlocking: %+v", b))
	}
}
func (b *BackoffEntry) TryWait(maxDuration time.Duration) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !b.tryLock() {
		return false
	}
	defer b.unlock()
	b.wait(maxDuration)
	return true
}
func (b *BackoffEntry) getBackoff(maxDuration time.Duration) time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	duration := b.backoff
	newDuration := time.Duration(duration) * 2
	if newDuration > maxDuration {
		newDuration = maxDuration
	}
	b.backoff = newDuration
	klog.V(4).Infof("Backing off %s", duration.String())
	return duration
}
func (b *BackoffEntry) wait(maxDuration time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	time.Sleep(b.getBackoff(maxDuration))
}

type PodBackoff struct {
	perPodBackoff   map[ktypes.NamespacedName]*BackoffEntry
	lock            sync.Mutex
	clock           clock
	defaultDuration time.Duration
	maxDuration     time.Duration
}

func (p *PodBackoff) MaxDuration() time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return p.maxDuration
}
func CreateDefaultPodBackoff() *PodBackoff {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return CreatePodBackoff(1*time.Second, 60*time.Second)
}
func CreatePodBackoff(defaultDuration, maxDuration time.Duration) *PodBackoff {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return CreatePodBackoffWithClock(defaultDuration, maxDuration, realClock{})
}
func CreatePodBackoffWithClock(defaultDuration, maxDuration time.Duration, clock clock) *PodBackoff {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PodBackoff{perPodBackoff: map[ktypes.NamespacedName]*BackoffEntry{}, clock: clock, defaultDuration: defaultDuration, maxDuration: maxDuration}
}
func (p *PodBackoff) GetEntry(podID ktypes.NamespacedName) *BackoffEntry {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.lock.Lock()
	defer p.lock.Unlock()
	entry, ok := p.perPodBackoff[podID]
	if !ok {
		entry = &BackoffEntry{backoff: p.defaultDuration}
		p.perPodBackoff[podID] = entry
	}
	entry.lastUpdate = p.clock.Now()
	return entry
}
func (p *PodBackoff) Gc() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	p.lock.Lock()
	defer p.lock.Unlock()
	now := p.clock.Now()
	for podID, entry := range p.perPodBackoff {
		if now.Sub(entry.lastUpdate) > p.maxDuration {
			delete(p.perPodBackoff, podID)
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
