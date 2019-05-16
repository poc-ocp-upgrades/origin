package controller

import (
	"fmt"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
	"sync"
)

type BaseControllerRefManager struct {
	Controller   metav1.Object
	Selector     labels.Selector
	canAdoptErr  error
	canAdoptOnce sync.Once
	CanAdoptFunc func() error
}

func (m *BaseControllerRefManager) CanAdopt() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.canAdoptOnce.Do(func() {
		if m.CanAdoptFunc != nil {
			m.canAdoptErr = m.CanAdoptFunc()
		}
	})
	return m.canAdoptErr
}
func (m *BaseControllerRefManager) ClaimObject(obj metav1.Object, match func(metav1.Object) bool, adopt, release func(metav1.Object) error) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerRef := metav1.GetControllerOf(obj)
	if controllerRef != nil {
		if controllerRef.UID != m.Controller.GetUID() {
			return false, nil
		}
		if match(obj) {
			return true, nil
		}
		if m.Controller.GetDeletionTimestamp() != nil {
			return false, nil
		}
		if err := release(obj); err != nil {
			if errors.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}
		return false, nil
	}
	if m.Controller.GetDeletionTimestamp() != nil || !match(obj) {
		return false, nil
	}
	if obj.GetDeletionTimestamp() != nil {
		return false, nil
	}
	if err := adopt(obj); err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

type PodControllerRefManager struct {
	BaseControllerRefManager
	controllerKind schema.GroupVersionKind
	podControl     PodControlInterface
}

func NewPodControllerRefManager(podControl PodControlInterface, controller metav1.Object, selector labels.Selector, controllerKind schema.GroupVersionKind, canAdopt func() error) *PodControllerRefManager {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PodControllerRefManager{BaseControllerRefManager: BaseControllerRefManager{Controller: controller, Selector: selector, CanAdoptFunc: canAdopt}, controllerKind: controllerKind, podControl: podControl}
}
func (m *PodControllerRefManager) ClaimPods(pods []*v1.Pod, filters ...func(*v1.Pod) bool) ([]*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var claimed []*v1.Pod
	var errlist []error
	match := func(obj metav1.Object) bool {
		pod := obj.(*v1.Pod)
		if !m.Selector.Matches(labels.Set(pod.Labels)) {
			return false
		}
		for _, filter := range filters {
			if !filter(pod) {
				return false
			}
		}
		return true
	}
	adopt := func(obj metav1.Object) error {
		return m.AdoptPod(obj.(*v1.Pod))
	}
	release := func(obj metav1.Object) error {
		return m.ReleasePod(obj.(*v1.Pod))
	}
	for _, pod := range pods {
		ok, err := m.ClaimObject(pod, match, adopt, release)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, pod)
		}
	}
	return claimed, utilerrors.NewAggregate(errlist)
}
func (m *PodControllerRefManager) AdoptPod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := m.CanAdopt(); err != nil {
		return fmt.Errorf("can't adopt Pod %v/%v (%v): %v", pod.Namespace, pod.Name, pod.UID, err)
	}
	addControllerPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName(), m.Controller.GetUID(), pod.UID)
	return m.podControl.PatchPod(pod.Namespace, pod.Name, []byte(addControllerPatch))
}
func (m *PodControllerRefManager) ReleasePod(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("patching pod %s_%s to remove its controllerRef to %s/%s:%s", pod.Namespace, pod.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.Controller.GetUID(), pod.UID)
	err := m.podControl.PatchPod(pod.Namespace, pod.Name, []byte(deleteOwnerRefPatch))
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		if errors.IsInvalid(err) {
			return nil
		}
	}
	return err
}

type ReplicaSetControllerRefManager struct {
	BaseControllerRefManager
	controllerKind schema.GroupVersionKind
	rsControl      RSControlInterface
}

func NewReplicaSetControllerRefManager(rsControl RSControlInterface, controller metav1.Object, selector labels.Selector, controllerKind schema.GroupVersionKind, canAdopt func() error) *ReplicaSetControllerRefManager {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ReplicaSetControllerRefManager{BaseControllerRefManager: BaseControllerRefManager{Controller: controller, Selector: selector, CanAdoptFunc: canAdopt}, controllerKind: controllerKind, rsControl: rsControl}
}
func (m *ReplicaSetControllerRefManager) ClaimReplicaSets(sets []*apps.ReplicaSet) ([]*apps.ReplicaSet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var claimed []*apps.ReplicaSet
	var errlist []error
	match := func(obj metav1.Object) bool {
		return m.Selector.Matches(labels.Set(obj.GetLabels()))
	}
	adopt := func(obj metav1.Object) error {
		return m.AdoptReplicaSet(obj.(*apps.ReplicaSet))
	}
	release := func(obj metav1.Object) error {
		return m.ReleaseReplicaSet(obj.(*apps.ReplicaSet))
	}
	for _, rs := range sets {
		ok, err := m.ClaimObject(rs, match, adopt, release)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, rs)
		}
	}
	return claimed, utilerrors.NewAggregate(errlist)
}
func (m *ReplicaSetControllerRefManager) AdoptReplicaSet(rs *apps.ReplicaSet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := m.CanAdopt(); err != nil {
		return fmt.Errorf("can't adopt ReplicaSet %v/%v (%v): %v", rs.Namespace, rs.Name, rs.UID, err)
	}
	addControllerPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName(), m.Controller.GetUID(), rs.UID)
	return m.rsControl.PatchReplicaSet(rs.Namespace, rs.Name, []byte(addControllerPatch))
}
func (m *ReplicaSetControllerRefManager) ReleaseReplicaSet(replicaSet *apps.ReplicaSet) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("patching ReplicaSet %s_%s to remove its controllerRef to %s/%s:%s", replicaSet.Namespace, replicaSet.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.Controller.GetUID(), replicaSet.UID)
	err := m.rsControl.PatchReplicaSet(replicaSet.Namespace, replicaSet.Name, []byte(deleteOwnerRefPatch))
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		if errors.IsInvalid(err) {
			return nil
		}
	}
	return err
}
func RecheckDeletionTimestamp(getObject func() (metav1.Object, error)) func() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() error {
		obj, err := getObject()
		if err != nil {
			return fmt.Errorf("can't recheck DeletionTimestamp: %v", err)
		}
		if obj.GetDeletionTimestamp() != nil {
			return fmt.Errorf("%v/%v has just been deleted at %v", obj.GetNamespace(), obj.GetName(), obj.GetDeletionTimestamp())
		}
		return nil
	}
}

type ControllerRevisionControllerRefManager struct {
	BaseControllerRefManager
	controllerKind schema.GroupVersionKind
	crControl      ControllerRevisionControlInterface
}

func NewControllerRevisionControllerRefManager(crControl ControllerRevisionControlInterface, controller metav1.Object, selector labels.Selector, controllerKind schema.GroupVersionKind, canAdopt func() error) *ControllerRevisionControllerRefManager {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ControllerRevisionControllerRefManager{BaseControllerRefManager: BaseControllerRefManager{Controller: controller, Selector: selector, CanAdoptFunc: canAdopt}, controllerKind: controllerKind, crControl: crControl}
}
func (m *ControllerRevisionControllerRefManager) ClaimControllerRevisions(histories []*apps.ControllerRevision) ([]*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var claimed []*apps.ControllerRevision
	var errlist []error
	match := func(obj metav1.Object) bool {
		return m.Selector.Matches(labels.Set(obj.GetLabels()))
	}
	adopt := func(obj metav1.Object) error {
		return m.AdoptControllerRevision(obj.(*apps.ControllerRevision))
	}
	release := func(obj metav1.Object) error {
		return m.ReleaseControllerRevision(obj.(*apps.ControllerRevision))
	}
	for _, h := range histories {
		ok, err := m.ClaimObject(h, match, adopt, release)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, h)
		}
	}
	return claimed, utilerrors.NewAggregate(errlist)
}
func (m *ControllerRevisionControllerRefManager) AdoptControllerRevision(history *apps.ControllerRevision) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := m.CanAdopt(); err != nil {
		return fmt.Errorf("can't adopt ControllerRevision %v/%v (%v): %v", history.Namespace, history.Name, history.UID, err)
	}
	addControllerPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName(), m.Controller.GetUID(), history.UID)
	return m.crControl.PatchControllerRevision(history.Namespace, history.Name, []byte(addControllerPatch))
}
func (m *ControllerRevisionControllerRefManager) ReleaseControllerRevision(history *apps.ControllerRevision) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("patching ControllerRevision %s_%s to remove its controllerRef to %s/%s:%s", history.Namespace, history.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.Controller.GetUID(), history.UID)
	err := m.crControl.PatchControllerRevision(history.Namespace, history.Name, []byte(deleteOwnerRefPatch))
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		if errors.IsInvalid(err) {
			return nil
		}
	}
	return err
}
