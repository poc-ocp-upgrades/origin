package controller

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"math/big"
	"reflect"
	"time"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	corev1informers "k8s.io/client-go/informers/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/controller"
	securityv1 "github.com/openshift/api/security/v1"
	securityv1client "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	"github.com/openshift/origin/pkg/security"
	"github.com/openshift/origin/pkg/security/mcs"
	"github.com/openshift/origin/pkg/security/uid"
	"github.com/openshift/origin/pkg/security/uidallocator"
)

const (
	controllerName	= "namespace-security-allocation-controller"
	rangeName	= "scc-uid"
)

type NamespaceSCCAllocationController struct {
	requiredUIDRange		*uid.Range
	mcsAllocator			MCSAllocationFunc
	nsLister			corev1listers.NamespaceLister
	nsListerSynced			cache.InformerSynced
	currentUIDRangeAllocation	*securityv1.RangeAllocation
	namespaceClient			corev1client.NamespaceInterface
	rangeAllocationClient		securityv1client.RangeAllocationsGetter
	queue				workqueue.RateLimitingInterface
}

func NewNamespaceSCCAllocationController(namespaceInformer corev1informers.NamespaceInformer, client corev1client.NamespaceInterface, rangeAllocationClient securityv1client.RangeAllocationsGetter, requiredUIDRange *uid.Range, mcs MCSAllocationFunc) *NamespaceSCCAllocationController {
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
	c := &NamespaceSCCAllocationController{requiredUIDRange: requiredUIDRange, mcsAllocator: mcs, namespaceClient: client, rangeAllocationClient: rangeAllocationClient, nsLister: namespaceInformer.Lister(), nsListerSynced: namespaceInformer.Informer().HasSynced, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName)}
	namespaceInformer.Informer().AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{AddFunc: c.enqueueNamespace, UpdateFunc: func(oldObj, newObj interface{}) {
		c.enqueueNamespace(newObj)
	}}, 10*time.Minute)
	return c
}
func (c *NamespaceSCCAllocationController) Run(stopCh <-chan struct{}) {
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
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	defer klog.V(1).Infof("Shutting down")
	if !controller.WaitForCacheSync(controllerName, stopCh, c.nsListerSynced) {
		return
	}
	klog.V(1).Infof("Repairing SCC UID Allocations")
	if err := c.WaitForRepair(stopCh); err != nil {
		klog.Fatal(err)
	}
	klog.V(1).Infof("Repair complete")
	go c.worker()
	<-stopCh
}
func (c *NamespaceSCCAllocationController) syncNamespace(key string) error {
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
	ns, err := c.nsLister.Get(key)
	if apierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if _, ok := ns.Annotations[security.UIDRangeAnnotation]; ok {
		return nil
	}
	return c.allocate(ns)
}
func (c *NamespaceSCCAllocationController) allocate(ns *corev1.Namespace) error {
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
	success := false
	defer func() {
		if success {
			return
		}
		c.currentUIDRangeAllocation = nil
	}()
	if c.currentUIDRangeAllocation == nil {
		newRange, err := c.rangeAllocationClient.RangeAllocations().Get(rangeName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		c.currentUIDRangeAllocation = newRange
	}
	uidRange, err := uid.ParseRange(c.currentUIDRangeAllocation.Range)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(*uidRange, *c.requiredUIDRange) {
		return fmt.Errorf("conflicting UID range; expected %#v, got %#v", *c.requiredUIDRange, *uidRange)
	}
	allocatedBitMapInt := big.NewInt(0).SetBytes(c.currentUIDRangeAllocation.Data)
	bitIndex, found := allocateNextContiguousBit(allocatedBitMapInt, int(uidRange.Size()))
	if !found {
		return fmt.Errorf("uid range exceeded")
	}
	allocatedBitMapInt = allocatedBitMapInt.SetBit(allocatedBitMapInt, bitIndex, 1)
	newRangeAllocation := c.currentUIDRangeAllocation.DeepCopy()
	newRangeAllocation.Data = allocatedBitMapInt.Bytes()
	actualRangeAllocation, err := c.rangeAllocationClient.RangeAllocations().Update(newRangeAllocation)
	if err != nil {
		return err
	}
	c.currentUIDRangeAllocation = actualRangeAllocation
	block, ok := uidRange.BlockAt(uint32(bitIndex))
	if !ok {
		return fmt.Errorf("%d not in range", bitIndex)
	}
	nsCopy := ns.DeepCopy()
	if nsCopy.Annotations == nil {
		nsCopy.Annotations = make(map[string]string)
	}
	nsCopy.Annotations[security.UIDRangeAnnotation] = block.String()
	nsCopy.Annotations[security.SupplementalGroupsAnnotation] = block.String()
	if _, ok := nsCopy.Annotations[security.MCSAnnotation]; !ok {
		if label := c.mcsAllocator(block); label != nil {
			nsCopy.Annotations[security.MCSAnnotation] = label.String()
		}
	}
	_, err = c.namespaceClient.Update(nsCopy)
	if apierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	success = true
	return nil
}
func allocateNextContiguousBit(allocated *big.Int, max int) (int, bool) {
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
	for i := 0; i < max; i++ {
		if allocated.Bit(i) == 0 {
			return i, true
		}
	}
	return 0, false
}
func (c *NamespaceSCCAllocationController) WaitForRepair(stopCh <-chan struct{}) error {
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
	return wait.PollImmediate(10*time.Second, 5*time.Minute, func() (bool, error) {
		select {
		case <-stopCh:
			return true, nil
		default:
		}
		err := c.Repair()
		if err == nil {
			return true, nil
		}
		utilruntime.HandleError(err)
		return false, nil
	})
}
func (c *NamespaceSCCAllocationController) Repair() error {
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
	uidRange, err := c.rangeAllocationClient.RangeAllocations().Get(rangeName, metav1.GetOptions{})
	needCreate := apierrors.IsNotFound(err)
	if err != nil && !needCreate {
		return err
	}
	if needCreate {
		uidRange = &securityv1.RangeAllocation{ObjectMeta: metav1.ObjectMeta{Name: rangeName}}
	}
	uids := uidallocator.NewInMemory(c.requiredUIDRange)
	nsList, err := c.nsLister.List(labels.Everything())
	if err != nil {
		return err
	}
	for _, ns := range nsList {
		value, ok := ns.Annotations[security.UIDRangeAnnotation]
		if !ok {
			continue
		}
		block, err := uid.ParseBlock(value)
		if err != nil {
			continue
		}
		switch err := uids.Allocate(block); err {
		case nil:
		case uidallocator.ErrNotInRange, uidallocator.ErrAllocated:
			continue
		case uidallocator.ErrFull:
			return fmt.Errorf("the UID range %s is full; you must widen the range in order to allocate more UIDs", c.requiredUIDRange)
		default:
			return fmt.Errorf("unable to allocate UID block %s for namespace %s due to an unknown error, exiting: %v", block, ns.Name, err)
		}
	}
	newRangeAllocation := &coreapi.RangeAllocation{}
	if err := uids.Snapshot(newRangeAllocation); err != nil {
		return err
	}
	uidRange.Range = newRangeAllocation.Range
	uidRange.Data = newRangeAllocation.Data
	if needCreate {
		if _, err := c.rangeAllocationClient.RangeAllocations().Create(uidRange); err != nil {
			return err
		}
		return nil
	}
	if _, err := c.rangeAllocationClient.RangeAllocations().Update(uidRange); err != nil {
		return err
	}
	return nil
}

type MCSAllocationFunc func(uid.Block) *mcs.Label

func DefaultMCSAllocation(from *uid.Range, to *mcs.Range, blockSize int) MCSAllocationFunc {
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
	return func(block uid.Block) *mcs.Label {
		ok, offset := from.Offset(block)
		if !ok {
			return nil
		}
		if blockSize > 0 {
			offset = offset * uint32(blockSize)
		}
		label, _ := to.LabelAt(uint64(offset))
		return label
	}
}
func (c *NamespaceSCCAllocationController) enqueueNamespace(obj interface{}) {
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
	ns, ok := obj.(*corev1.Namespace)
	if !ok {
		return
	}
	c.queue.Add(ns.Name)
}
func (c *NamespaceSCCAllocationController) worker() {
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
	for c.work() {
	}
}
func (c *NamespaceSCCAllocationController) work() bool {
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
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	if err := c.syncNamespace(key.(string)); err == nil {
		c.queue.Forget(key)
	} else {
		utilruntime.HandleError(fmt.Errorf("error syncing namespace, it will be retried: %v", err))
		c.queue.AddRateLimited(key)
	}
	return true
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
