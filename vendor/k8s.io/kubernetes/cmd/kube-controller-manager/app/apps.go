package app

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/kubernetes/pkg/controller/daemon"
	"k8s.io/kubernetes/pkg/controller/deployment"
	"k8s.io/kubernetes/pkg/controller/replicaset"
	"k8s.io/kubernetes/pkg/controller/statefulset"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func startDaemonSetController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "daemonsets"}] {
		return nil, false, nil
	}
	dsc, err := daemon.NewNodeSelectorAwareDaemonSetsController(ctx.OpenShiftContext.OpenShiftDefaultProjectNodeSelector, ctx.OpenShiftContext.KubeDefaultProjectNodeSelector, ctx.InformerFactory.Core().V1().Namespaces(), ctx.InformerFactory.Apps().V1().DaemonSets(), ctx.InformerFactory.Apps().V1().ControllerRevisions(), ctx.InformerFactory.Core().V1().Pods(), ctx.InformerFactory.Core().V1().Nodes(), ctx.ClientBuilder.ClientOrDie("daemon-set-controller"), flowcontrol.NewBackOff(1*time.Second, 15*time.Minute))
	if err != nil {
		return nil, true, fmt.Errorf("error creating DaemonSets controller: %v", err)
	}
	go dsc.Run(int(ctx.ComponentConfig.DaemonSetController.ConcurrentDaemonSetSyncs), ctx.Stop)
	return nil, true, nil
}
func startStatefulSetController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"}] {
		return nil, false, nil
	}
	go statefulset.NewStatefulSetController(ctx.InformerFactory.Core().V1().Pods(), ctx.InformerFactory.Apps().V1().StatefulSets(), ctx.InformerFactory.Core().V1().PersistentVolumeClaims(), ctx.InformerFactory.Apps().V1().ControllerRevisions(), ctx.ClientBuilder.ClientOrDie("statefulset-controller")).Run(1, ctx.Stop)
	return nil, true, nil
}
func startReplicaSetController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "replicasets"}] {
		return nil, false, nil
	}
	go replicaset.NewReplicaSetController(ctx.InformerFactory.Apps().V1().ReplicaSets(), ctx.InformerFactory.Core().V1().Pods(), ctx.ClientBuilder.ClientOrDie("replicaset-controller"), replicaset.BurstReplicas).Run(int(ctx.ComponentConfig.ReplicaSetController.ConcurrentRSSyncs), ctx.Stop)
	return nil, true, nil
}
func startDeploymentController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}] {
		return nil, false, nil
	}
	dc, err := deployment.NewDeploymentController(ctx.InformerFactory.Apps().V1().Deployments(), ctx.InformerFactory.Apps().V1().ReplicaSets(), ctx.InformerFactory.Core().V1().Pods(), ctx.ClientBuilder.ClientOrDie("deployment-controller"))
	if err != nil {
		return nil, true, fmt.Errorf("error creating Deployment controller: %v", err)
	}
	go dc.Run(int(ctx.ComponentConfig.DeploymentController.ConcurrentDeploymentSyncs), ctx.Stop)
	return nil, true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
