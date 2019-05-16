package app

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/controller/cronjob"
	"k8s.io/kubernetes/pkg/controller/job"
	"net/http"
)

func startJobController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "jobs"}] {
		return nil, false, nil
	}
	go job.NewJobController(ctx.InformerFactory.Core().V1().Pods(), ctx.InformerFactory.Batch().V1().Jobs(), ctx.ClientBuilder.ClientOrDie("job-controller")).Run(int(ctx.ComponentConfig.JobController.ConcurrentJobSyncs), ctx.Stop)
	return nil, true, nil
}
func startCronJobController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "batch", Version: "v1beta1", Resource: "cronjobs"}] {
		return nil, false, nil
	}
	cjc, err := cronjob.NewCronJobController(ctx.ClientBuilder.ClientOrDie("cronjob-controller"))
	if err != nil {
		return nil, true, fmt.Errorf("error creating CronJob controller: %v", err)
	}
	go cjc.Run(ctx.Stop)
	return nil, true, nil
}
