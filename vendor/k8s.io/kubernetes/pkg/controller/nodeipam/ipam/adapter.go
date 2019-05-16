package ipam

import (
	"context"
	"encoding/json"
	goformat "fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"k8s.io/metrics/pkg/client/clientset/versioned/scheme"
	"net"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type adapter struct {
	k8s      clientset.Interface
	cloud    *gce.Cloud
	recorder record.EventRecorder
}

func newAdapter(k8s clientset.Interface, cloud *gce.Cloud) *adapter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := &adapter{k8s: k8s, cloud: cloud}
	broadcaster := record.NewBroadcaster()
	broadcaster.StartLogging(klog.Infof)
	ret.recorder = broadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "cloudCIDRAllocator"})
	klog.V(0).Infof("Sending events to api server.")
	broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: k8s.CoreV1().Events("")})
	return ret
}
func (a *adapter) Alias(ctx context.Context, nodeName string) (*net.IPNet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cidrs, err := a.cloud.AliasRanges(types.NodeName(nodeName))
	if err != nil {
		return nil, err
	}
	switch len(cidrs) {
	case 0:
		return nil, nil
	case 1:
		break
	default:
		klog.Warningf("Node %q has more than one alias assigned (%v), defaulting to the first", nodeName, cidrs)
	}
	_, cidrRange, err := net.ParseCIDR(cidrs[0])
	if err != nil {
		return nil, err
	}
	return cidrRange, nil
}
func (a *adapter) AddAlias(ctx context.Context, nodeName string, cidrRange *net.IPNet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return a.cloud.AddAliasToInstance(types.NodeName(nodeName), cidrRange)
}
func (a *adapter) Node(ctx context.Context, name string) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return a.k8s.CoreV1().Nodes().Get(name, metav1.GetOptions{})
}
func (a *adapter) UpdateNodePodCIDR(ctx context.Context, node *v1.Node, cidrRange *net.IPNet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	patch := map[string]interface{}{"apiVersion": node.APIVersion, "kind": node.Kind, "metadata": map[string]interface{}{"name": node.Name}, "spec": map[string]interface{}{"podCIDR": cidrRange.String()}}
	bytes, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	_, err = a.k8s.CoreV1().Nodes().Patch(node.Name, types.StrategicMergePatchType, bytes)
	return err
}
func (a *adapter) UpdateNodeNetworkUnavailable(nodeName string, unavailable bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	condition := v1.ConditionFalse
	if unavailable {
		condition = v1.ConditionTrue
	}
	return nodeutil.SetNodeCondition(a.k8s, types.NodeName(nodeName), v1.NodeCondition{Type: v1.NodeNetworkUnavailable, Status: condition, Reason: "RouteCreated", Message: "NodeController created an implicit route", LastTransitionTime: metav1.Now()})
}
func (a *adapter) EmitNodeWarningEvent(nodeName, reason, fmt string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ref := &v1.ObjectReference{Kind: "Node", Name: nodeName}
	a.recorder.Eventf(ref, v1.EventTypeNormal, reason, fmt, args...)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
