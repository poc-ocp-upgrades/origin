package storage

import (
	"context"
	"errors"
	"fmt"
	goformat "fmt"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/storage"
	storeerr "k8s.io/apiserver/pkg/storage/errors"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/rangeallocation"
	"k8s.io/kubernetes/pkg/registry/core/service/allocator"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

var (
	errorUnableToAllocate = errors.New("unable to allocate")
)

type Etcd struct {
	lock     sync.Mutex
	alloc    allocator.Snapshottable
	storage  storage.Interface
	last     string
	baseKey  string
	resource schema.GroupResource
}

var _ allocator.Interface = &Etcd{}
var _ rangeallocation.RangeRegistry = &Etcd{}

func NewEtcd(alloc allocator.Snapshottable, baseKey string, resource schema.GroupResource, config *storagebackend.Config) *Etcd {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage, d := generic.NewRawStorage(config)
	registry.RegisterStorageCleanup(d)
	return &Etcd{alloc: alloc, storage: storage, baseKey: baseKey, resource: resource}
}
func (e *Etcd) Allocate(offset int) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	ok, err := e.alloc.Allocate(offset)
	if !ok || err != nil {
		return ok, err
	}
	err = e.tryUpdate(func() error {
		ok, err := e.alloc.Allocate(offset)
		if err != nil {
			return err
		}
		if !ok {
			return errorUnableToAllocate
		}
		return nil
	})
	if err != nil {
		if err == errorUnableToAllocate {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (e *Etcd) AllocateNext() (int, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	offset, ok, err := e.alloc.AllocateNext()
	if !ok || err != nil {
		return offset, ok, err
	}
	err = e.tryUpdate(func() error {
		ok, err := e.alloc.Allocate(offset)
		if err != nil {
			return err
		}
		if !ok {
			offset, ok, err = e.alloc.AllocateNext()
			if err != nil {
				return err
			}
			if !ok {
				return errorUnableToAllocate
			}
			return nil
		}
		return nil
	})
	return offset, ok, err
}
func (e *Etcd) Release(item int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	if err := e.alloc.Release(item); err != nil {
		return err
	}
	return e.tryUpdate(func() error {
		return e.alloc.Release(item)
	})
}
func (e *Etcd) ForEach(fn func(int)) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	e.alloc.ForEach(fn)
}
func (e *Etcd) tryUpdate(fn func() error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := e.storage.GuaranteedUpdate(context.TODO(), e.baseKey, &api.RangeAllocation{}, true, nil, storage.SimpleUpdate(func(input runtime.Object) (output runtime.Object, err error) {
		existing := input.(*api.RangeAllocation)
		if len(existing.ResourceVersion) == 0 {
			return nil, fmt.Errorf("cannot allocate resources of type %s at this time", e.resource.String())
		}
		if existing.ResourceVersion != e.last {
			if err := e.alloc.Restore(existing.Range, existing.Data); err != nil {
				return nil, err
			}
			if err := fn(); err != nil {
				return nil, err
			}
		}
		e.last = existing.ResourceVersion
		rangeSpec, data := e.alloc.Snapshot()
		existing.Range = rangeSpec
		existing.Data = data
		return existing, nil
	}))
	return storeerr.InterpretUpdateError(err, e.resource, "")
}
func (e *Etcd) Get() (*api.RangeAllocation, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	existing := &api.RangeAllocation{}
	if err := e.storage.Get(context.TODO(), e.baseKey, "", existing, true); err != nil {
		return nil, storeerr.InterpretGetError(err, e.resource, "")
	}
	return existing, nil
}
func (e *Etcd) CreateOrUpdate(snapshot *api.RangeAllocation) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	last := ""
	err := e.storage.GuaranteedUpdate(context.TODO(), e.baseKey, &api.RangeAllocation{}, true, nil, storage.SimpleUpdate(func(input runtime.Object) (output runtime.Object, err error) {
		existing := input.(*api.RangeAllocation)
		switch {
		case len(snapshot.ResourceVersion) != 0 && len(existing.ResourceVersion) != 0:
			if snapshot.ResourceVersion != existing.ResourceVersion {
				return nil, k8serr.NewConflict(e.resource, "", fmt.Errorf("the provided resource version does not match"))
			}
		case len(existing.ResourceVersion) != 0:
			return nil, k8serr.NewConflict(e.resource, "", fmt.Errorf("another caller has already initialized the resource"))
		}
		last = snapshot.ResourceVersion
		return snapshot, nil
	}))
	if err != nil {
		return storeerr.InterpretUpdateError(err, e.resource, "")
	}
	err = e.alloc.Restore(snapshot.Range, snapshot.Data)
	if err == nil {
		e.last = last
	}
	return err
}
func (e *Etcd) Has(item int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.alloc.Has(item)
}
func (e *Etcd) Free() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.alloc.Free()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
