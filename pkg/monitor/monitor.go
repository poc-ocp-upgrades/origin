package monitor

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

type Monitor struct {
	interval	time.Duration
	samplers	[]SamplerFunc
	lock		sync.Mutex
	events		[]*Event
	samples		[]*sample
}

func NewMonitor() *Monitor {
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
	return &Monitor{interval: 15 * time.Second}
}

var _ Interface = &Monitor{}

func (m *Monitor) StartSampling(ctx context.Context) {
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
	if m.interval == 0 {
		return
	}
	go func() {
		ticker := time.NewTicker(m.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				m.sample()
				return
			}
			m.sample()
		}
	}()
}
func (m *Monitor) AddSampler(fn SamplerFunc) {
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
	m.lock.Lock()
	defer m.lock.Unlock()
	m.samplers = append(m.samplers, fn)
}
func (m *Monitor) Record(conditions ...Condition) {
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
	if len(conditions) == 0 {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	t := time.Now().UTC()
	for _, condition := range conditions {
		m.events = append(m.events, &Event{At: t, Condition: condition})
	}
}
func (m *Monitor) sample() {
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
	m.lock.Lock()
	samplers := m.samplers
	m.lock.Unlock()
	now := time.Now().UTC()
	var conditions []*Condition
	for _, fn := range samplers {
		conditions = append(conditions, fn(now)...)
	}
	if len(conditions) == 0 {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	t := time.Now().UTC()
	m.samples = append(m.samples, &sample{at: t, conditions: conditions})
}
func (m *Monitor) snapshot() ([]*sample, []*Event) {
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
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.samples, m.events
}
func (m *Monitor) Conditions(from, to time.Time) EventIntervals {
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
	samples, _ := m.snapshot()
	return filterSamples(samples, from, to)
}
func (m *Monitor) Events(from, to time.Time) EventIntervals {
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
	samples, events := m.snapshot()
	intervals := filterSamples(samples, from, to)
	events = filterEvents(events, from, to)
	mustSort := len(intervals) > 0
	for i := range events {
		if i > 0 && events[i-1].At.After(events[i].At) {
			fmt.Printf("ERROR: event %d out of order\n  %#v\n  %#v\n", i, events[i-1], events[i])
		}
		at := events[i].At
		condition := &events[i].Condition
		intervals = append(intervals, &EventInterval{From: at, To: at, Condition: condition})
	}
	if mustSort {
		sort.Sort(intervals)
	}
	return intervals
}
func filterSamples(samples []*sample, from, to time.Time) EventIntervals {
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
	if len(samples) == 0 {
		return nil
	}
	if !from.IsZero() {
		first := sort.Search(len(samples), func(i int) bool {
			return samples[i].at.After(from)
		})
		if first == -1 {
			return nil
		}
		samples = samples[first:]
	}
	if !to.IsZero() {
		for i, sample := range samples {
			if sample.at.After(to) {
				samples = samples[:i]
				break
			}
		}
	}
	if len(samples) == 0 {
		return nil
	}
	intervals := make(EventIntervals, 0, len(samples)*2)
	last, next := make(map[Condition]*EventInterval), make(map[Condition]*EventInterval)
	for _, sample := range samples {
		for _, condition := range sample.conditions {
			interval, ok := last[*condition]
			if ok {
				interval.To = sample.at
				next[*condition] = interval
				continue
			}
			interval = &EventInterval{Condition: condition, From: sample.at, To: sample.at}
			next[*condition] = interval
			intervals = append(intervals, interval)
		}
		for k := range last {
			delete(last, k)
		}
		last, next = next, last
	}
	return intervals
}
func filterEvents(events []*Event, from, to time.Time) []*Event {
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
	if from.IsZero() && to.IsZero() {
		return events
	}
	first := sort.Search(len(events), func(i int) bool {
		return events[i].At.After(from)
	})
	if first == -1 {
		return nil
	}
	if to.IsZero() {
		return events[first:]
	}
	for i := first; i < len(events); i++ {
		if events[i].At.After(to) {
			return events[first:i]
		}
	}
	return events[first:]
}
