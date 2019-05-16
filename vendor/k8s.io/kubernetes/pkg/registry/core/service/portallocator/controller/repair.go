package controller

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/registry/core/rangeallocation"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type Repair struct {
	interval      time.Duration
	serviceClient corev1client.ServicesGetter
	portRange     net.PortRange
	alloc         rangeallocation.RangeRegistry
	leaks         map[int]int
	recorder      record.EventRecorder
}

const numRepairsBeforeLeakCleanup = 3

func NewRepair(interval time.Duration, serviceClient corev1client.ServicesGetter, eventClient corev1client.EventsGetter, portRange net.PortRange, alloc rangeallocation.RangeRegistry) *Repair {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&corev1client.EventSinkImpl{Interface: eventClient.Events("")})
	recorder := eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: "portallocator-repair-controller"})
	return &Repair{interval: interval, serviceClient: serviceClient, portRange: portRange, alloc: alloc, leaks: map[int]int{}, recorder: recorder}
}
func (c *Repair) RunUntil(ch chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(func() {
		if err := c.RunOnce(); err != nil {
			runtime.HandleError(err)
		}
	}, c.interval, ch)
}
func (c *Repair) RunOnce() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return retry.RetryOnConflict(retry.DefaultBackoff, c.runOnce)
}
func (c *Repair) runOnce() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var snapshot *api.RangeAllocation
	err := wait.PollImmediate(time.Second, 10*time.Second, func() (bool, error) {
		var err error
		snapshot, err = c.alloc.Get()
		return err == nil, err
	})
	if err != nil {
		return fmt.Errorf("unable to refresh the port allocations: %v", err)
	}
	if snapshot.Range == "" {
		snapshot.Range = c.portRange.String()
	}
	stored, err := portallocator.NewFromSnapshot(snapshot)
	if err != nil {
		return fmt.Errorf("unable to rebuild allocator from snapshot: %v", err)
	}
	list, err := c.serviceClient.Services(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("unable to refresh the port block: %v", err)
	}
	rebuilt := portallocator.NewPortAllocator(c.portRange)
	for i := range list.Items {
		svc := &list.Items[i]
		ports := collectServiceNodePorts(svc)
		if len(ports) == 0 {
			continue
		}
		for _, port := range ports {
			switch err := rebuilt.Allocate(port); err {
			case nil:
				if stored.Has(port) {
					stored.Release(port)
				} else {
					c.recorder.Eventf(svc, v1.EventTypeWarning, "PortNotAllocated", "Port %d is not allocated; repairing", port)
					runtime.HandleError(fmt.Errorf("the node port %d for service %s/%s is not allocated; repairing", port, svc.Name, svc.Namespace))
				}
				delete(c.leaks, port)
			case portallocator.ErrAllocated:
				c.recorder.Eventf(svc, v1.EventTypeWarning, "PortAlreadyAllocated", "Port %d was assigned to multiple services; please recreate service", port)
				runtime.HandleError(fmt.Errorf("the node port %d for service %s/%s was assigned to multiple services; please recreate", port, svc.Name, svc.Namespace))
			case err.(*portallocator.ErrNotInRange):
				c.recorder.Eventf(svc, v1.EventTypeWarning, "PortOutOfRange", "Port %d is not within the port range %s; please recreate service", port, c.portRange)
				runtime.HandleError(fmt.Errorf("the port %d for service %s/%s is not within the port range %s; please recreate", port, svc.Name, svc.Namespace, c.portRange))
			case portallocator.ErrFull:
				c.recorder.Eventf(svc, v1.EventTypeWarning, "PortRangeFull", "Port range %s is full; you must widen the port range in order to create new services", c.portRange)
				return fmt.Errorf("the port range %s is full; you must widen the port range in order to create new services", c.portRange)
			default:
				c.recorder.Eventf(svc, v1.EventTypeWarning, "UnknownError", "Unable to allocate port %d due to an unknown error", port)
				return fmt.Errorf("unable to allocate port %d for service %s/%s due to an unknown error, exiting: %v", port, svc.Name, svc.Namespace, err)
			}
		}
	}
	stored.ForEach(func(port int) {
		count, found := c.leaks[port]
		switch {
		case !found:
			runtime.HandleError(fmt.Errorf("the node port %d may have leaked: flagging for later clean up", port))
			count = numRepairsBeforeLeakCleanup - 1
			fallthrough
		case count > 0:
			c.leaks[port] = count - 1
			if err := rebuilt.Allocate(port); err != nil {
				runtime.HandleError(fmt.Errorf("the node port %d may have leaked, but can not be allocated: %v", port, err))
			}
		default:
			runtime.HandleError(fmt.Errorf("the node port %d appears to have leaked: cleaning up", port))
		}
	})
	if err := rebuilt.Snapshot(snapshot); err != nil {
		return fmt.Errorf("unable to snapshot the updated port allocations: %v", err)
	}
	if err := c.alloc.CreateOrUpdate(snapshot); err != nil {
		if errors.IsConflict(err) {
			return err
		}
		return fmt.Errorf("unable to persist the updated port allocations: %v", err)
	}
	return nil
}
func collectServiceNodePorts(service *corev1.Service) []int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	servicePorts := []int{}
	for i := range service.Spec.Ports {
		servicePort := &service.Spec.Ports[i]
		if servicePort.NodePort != 0 {
			servicePorts = append(servicePorts, int(servicePort.NodePort))
		}
	}
	if service.Spec.HealthCheckNodePort != 0 {
		servicePorts = append(servicePorts, int(service.Spec.HealthCheckNodePort))
	}
	return servicePorts
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
