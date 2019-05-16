package statefulset

import (
	"fmt"
	apps "k8s.io/api/apps/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientset "k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/util/retry"
)

type StatefulSetStatusUpdaterInterface interface {
	UpdateStatefulSetStatus(set *apps.StatefulSet, status *apps.StatefulSetStatus) error
}

func NewRealStatefulSetStatusUpdater(client clientset.Interface, setLister appslisters.StatefulSetLister) StatefulSetStatusUpdaterInterface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &realStatefulSetStatusUpdater{client, setLister}
}

type realStatefulSetStatusUpdater struct {
	client    clientset.Interface
	setLister appslisters.StatefulSetLister
}

func (ssu *realStatefulSetStatusUpdater) UpdateStatefulSetStatus(set *apps.StatefulSet, status *apps.StatefulSetStatus) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		set.Status = *status
		_, updateErr := ssu.client.AppsV1().StatefulSets(set.Namespace).UpdateStatus(set)
		if updateErr == nil {
			return nil
		}
		if updated, err := ssu.setLister.StatefulSets(set.Namespace).Get(set.Name); err == nil {
			set = updated.DeepCopy()
		} else {
			utilruntime.HandleError(fmt.Errorf("error getting updated StatefulSet %s/%s from lister: %v", set.Namespace, set.Name, err))
		}
		return updateErr
	})
}

var _ StatefulSetStatusUpdaterInterface = &realStatefulSetStatusUpdater{}
