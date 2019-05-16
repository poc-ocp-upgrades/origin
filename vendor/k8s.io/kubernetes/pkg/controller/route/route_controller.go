package route

import (
	"context"
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	clientretry "k8s.io/client-go/util/retry"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	v1node "k8s.io/kubernetes/pkg/api/v1/node"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/util/metrics"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"net"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

const (
	maxConcurrentRouteCreations int = 200
)

var updateNetworkConditionBackoff = wait.Backoff{Steps: 5, Duration: 100 * time.Millisecond, Jitter: 1.0}

type RouteController struct {
	routes           cloudprovider.Routes
	kubeClient       clientset.Interface
	clusterName      string
	clusterCIDR      *net.IPNet
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced
	broadcaster      record.EventBroadcaster
	recorder         record.EventRecorder
}

func New(routes cloudprovider.Routes, kubeClient clientset.Interface, nodeInformer coreinformers.NodeInformer, clusterName string, clusterCIDR *net.IPNet) *RouteController {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		metrics.RegisterMetricAndTrackRateLimiterUsage("route_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
	}
	if clusterCIDR == nil {
		klog.Fatal("RouteController: Must specify clusterCIDR.")
	}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "route_controller"})
	rc := &RouteController{routes: routes, kubeClient: kubeClient, clusterName: clusterName, clusterCIDR: clusterCIDR, nodeLister: nodeInformer.Lister(), nodeListerSynced: nodeInformer.Informer().HasSynced, broadcaster: eventBroadcaster, recorder: recorder}
	return rc
}
func (rc *RouteController) Run(stopCh <-chan struct{}, syncPeriod time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer utilruntime.HandleCrash()
	klog.Info("Starting route controller")
	defer klog.Info("Shutting down route controller")
	if !controller.WaitForCacheSync("route", stopCh, rc.nodeListerSynced) {
		return
	}
	if rc.broadcaster != nil {
		rc.broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: rc.kubeClient.CoreV1().Events("")})
	}
	go wait.NonSlidingUntil(func() {
		if err := rc.reconcileNodeRoutes(); err != nil {
			klog.Errorf("Couldn't reconcile node routes: %v", err)
		}
	}, syncPeriod, stopCh)
	<-stopCh
}
func (rc *RouteController) reconcileNodeRoutes() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	routeList, err := rc.routes.ListRoutes(context.TODO(), rc.clusterName)
	if err != nil {
		return fmt.Errorf("error listing routes: %v", err)
	}
	nodes, err := rc.nodeLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error listing nodes: %v", err)
	}
	return rc.reconcile(nodes, routeList)
}
func (rc *RouteController) reconcile(nodes []*v1.Node, routes []*cloudprovider.Route) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeCIDRs := make(map[types.NodeName]string)
	routeMap := make(map[types.NodeName]*cloudprovider.Route)
	for _, route := range routes {
		if route.TargetNode != "" {
			routeMap[route.TargetNode] = route
		}
	}
	wg := sync.WaitGroup{}
	rateLimiter := make(chan struct{}, maxConcurrentRouteCreations)
	for _, node := range nodes {
		if node.Spec.PodCIDR == "" {
			continue
		}
		nodeName := types.NodeName(node.Name)
		r := routeMap[nodeName]
		if r == nil || r.DestinationCIDR != node.Spec.PodCIDR {
			route := &cloudprovider.Route{TargetNode: nodeName, DestinationCIDR: node.Spec.PodCIDR}
			nameHint := string(node.UID)
			wg.Add(1)
			go func(nodeName types.NodeName, nameHint string, route *cloudprovider.Route) {
				defer wg.Done()
				err := clientretry.RetryOnConflict(updateNetworkConditionBackoff, func() error {
					startTime := time.Now()
					rateLimiter <- struct{}{}
					klog.Infof("Creating route for node %s %s with hint %s, throttled %v", nodeName, route.DestinationCIDR, nameHint, time.Since(startTime))
					err := rc.routes.CreateRoute(context.TODO(), rc.clusterName, nameHint, route)
					<-rateLimiter
					rc.updateNetworkingCondition(nodeName, err == nil)
					if err != nil {
						msg := fmt.Sprintf("Could not create route %s %s for node %s after %v: %v", nameHint, route.DestinationCIDR, nodeName, time.Since(startTime), err)
						if rc.recorder != nil {
							rc.recorder.Eventf(&v1.ObjectReference{Kind: "Node", Name: string(nodeName), UID: types.UID(nodeName), Namespace: ""}, v1.EventTypeWarning, "FailedToCreateRoute", msg)
						}
						klog.V(4).Infof(msg)
						return err
					}
					klog.Infof("Created route for node %s %s with hint %s after %v", nodeName, route.DestinationCIDR, nameHint, time.Now().Sub(startTime))
					return nil
				})
				if err != nil {
					klog.Errorf("Could not create route %s %s for node %s: %v", nameHint, route.DestinationCIDR, nodeName, err)
				}
			}(nodeName, nameHint, route)
		} else {
			_, condition := v1node.GetNodeCondition(&node.Status, v1.NodeNetworkUnavailable)
			if condition == nil || condition.Status != v1.ConditionFalse {
				rc.updateNetworkingCondition(types.NodeName(node.Name), true)
			}
		}
		nodeCIDRs[nodeName] = node.Spec.PodCIDR
	}
	for _, route := range routes {
		if rc.isResponsibleForRoute(route) {
			if route.Blackhole || (nodeCIDRs[route.TargetNode] != route.DestinationCIDR) {
				wg.Add(1)
				go func(route *cloudprovider.Route, startTime time.Time) {
					defer wg.Done()
					klog.Infof("Deleting route %s %s", route.Name, route.DestinationCIDR)
					if err := rc.routes.DeleteRoute(context.TODO(), rc.clusterName, route); err != nil {
						klog.Errorf("Could not delete route %s %s after %v: %v", route.Name, route.DestinationCIDR, time.Since(startTime), err)
					} else {
						klog.Infof("Deleted route %s %s after %v", route.Name, route.DestinationCIDR, time.Since(startTime))
					}
				}(route, time.Now())
			}
		}
	}
	wg.Wait()
	return nil
}
func (rc *RouteController) updateNetworkingCondition(nodeName types.NodeName, routeCreated bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := clientretry.RetryOnConflict(updateNetworkConditionBackoff, func() error {
		var err error
		currentTime := metav1.Now()
		if routeCreated {
			err = nodeutil.SetNodeCondition(rc.kubeClient, nodeName, v1.NodeCondition{Type: v1.NodeNetworkUnavailable, Status: v1.ConditionFalse, Reason: "RouteCreated", Message: "RouteController created a route", LastTransitionTime: currentTime})
		} else {
			err = nodeutil.SetNodeCondition(rc.kubeClient, nodeName, v1.NodeCondition{Type: v1.NodeNetworkUnavailable, Status: v1.ConditionTrue, Reason: "NoRouteCreated", Message: "RouteController failed to create a route", LastTransitionTime: currentTime})
		}
		if err != nil {
			klog.V(4).Infof("Error updating node %s, retrying: %v", nodeName, err)
		}
		return err
	})
	if err != nil {
		klog.Errorf("Error updating node %s: %v", nodeName, err)
	}
	return err
}
func (rc *RouteController) isResponsibleForRoute(route *cloudprovider.Route) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, cidr, err := net.ParseCIDR(route.DestinationCIDR)
	if err != nil {
		klog.Errorf("Ignoring route %s, unparsable CIDR: %v", route.Name, err)
		return false
	}
	lastIP := make([]byte, len(cidr.IP))
	for i := range lastIP {
		lastIP[i] = cidr.IP[i] | ^cidr.Mask[i]
	}
	if !rc.clusterCIDR.Contains(cidr.IP) || !rc.clusterCIDR.Contains(lastIP) {
		return false
	}
	return true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
