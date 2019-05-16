package markcontrolplane

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func MarkControlPlane(client clientset.Interface, controlPlaneName string, taints []v1.Taint) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[mark-control-plane] Marking the node %s as control-plane by adding the label \"%s=''\"\n", controlPlaneName, constants.LabelNodeRoleMaster)
	if taints != nil && len(taints) > 0 {
		taintStrs := []string{}
		for _, taint := range taints {
			taintStrs = append(taintStrs, taint.ToString())
		}
		fmt.Printf("[mark-control-plane] Marking the node %s as control-plane by adding the taints %v\n", controlPlaneName, taintStrs)
	}
	return apiclient.PatchNode(client, controlPlaneName, func(n *v1.Node) {
		markMasterNode(n, taints)
	})
}
func taintExists(taint v1.Taint, taints []v1.Taint) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, t := range taints {
		if t == taint {
			return true
		}
	}
	return false
}
func markMasterNode(n *v1.Node, taints []v1.Taint) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	n.ObjectMeta.Labels[constants.LabelNodeRoleMaster] = ""
	for _, nt := range n.Spec.Taints {
		if !taintExists(nt, taints) {
			taints = append(taints, nt)
		}
	}
	n.Spec.Taints = taints
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
