package app

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/disruption"
	"net/http"
)

func startDisruptionController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var group = "policy"
	var version = "v1beta1"
	var resource = "poddisruptionbudgets"
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: group, Version: version, Resource: resource}] {
		klog.Infof("Refusing to start disruption because resource %q in group %q is not available.", resource, group+"/"+version)
		return nil, false, nil
	}
	go disruption.NewDisruptionController(ctx.InformerFactory.Core().V1().Pods(), ctx.InformerFactory.Policy().V1beta1().PodDisruptionBudgets(), ctx.InformerFactory.Core().V1().ReplicationControllers(), ctx.InformerFactory.Extensions().V1beta1().ReplicaSets(), ctx.InformerFactory.Extensions().V1beta1().Deployments(), ctx.InformerFactory.Apps().V1beta1().StatefulSets(), ctx.ClientBuilder.ClientOrDie("disruption-controller")).Run(ctx.Stop)
	return nil, true, nil
}
