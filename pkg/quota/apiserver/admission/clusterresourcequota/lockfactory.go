package clusterresourcequota

import (
	"sync"
)

type LockFactory interface{ GetLock(string) sync.Locker }
type DefaultLockFactory struct {
	lock	sync.RWMutex
	locks	map[string]sync.Locker
}

func NewDefaultLockFactory() *DefaultLockFactory {
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
	return &DefaultLockFactory{locks: map[string]sync.Locker{}}
}
func (f *DefaultLockFactory) GetLock(key string) sync.Locker {
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
	lock, exists := f.getExistingLock(key)
	if exists {
		return lock
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	lock = &sync.Mutex{}
	f.locks[key] = lock
	return lock
}
func (f *DefaultLockFactory) getExistingLock(key string) (sync.Locker, bool) {
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
	f.lock.RLock()
	defer f.lock.RUnlock()
	lock, exists := f.locks[key]
	return lock, exists
}
