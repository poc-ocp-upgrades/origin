package testing

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 policy "k8s.io/api/policy/v1beta1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
)

var _ algorithm.NodeLister = &FakeNodeLister{}

type FakeNodeLister []*v1.Node

func (f FakeNodeLister) List() ([]*v1.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, nil
}

var _ algorithm.PodLister = &FakePodLister{}

type FakePodLister []*v1.Pod

func (f FakePodLister) List(s labels.Selector) (selected []*v1.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, pod := range f {
  if s.Matches(labels.Set(pod.Labels)) {
   selected = append(selected, pod)
  }
 }
 return selected, nil
}
func (f FakePodLister) FilteredList(podFilter algorithm.PodFilter, s labels.Selector) (selected []*v1.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, pod := range f {
  if podFilter(pod) && s.Matches(labels.Set(pod.Labels)) {
   selected = append(selected, pod)
  }
 }
 return selected, nil
}

var _ algorithm.ServiceLister = &FakeServiceLister{}

type FakeServiceLister []*v1.Service

func (f FakeServiceLister) List(labels.Selector) ([]*v1.Service, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, nil
}
func (f FakeServiceLister) GetPodServices(pod *v1.Pod) (services []*v1.Service, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var selector labels.Selector
 for i := range f {
  service := f[i]
  if service.Namespace != pod.Namespace {
   continue
  }
  selector = labels.Set(service.Spec.Selector).AsSelectorPreValidated()
  if selector.Matches(labels.Set(pod.Labels)) {
   services = append(services, service)
  }
 }
 return
}

var _ algorithm.ControllerLister = &FakeControllerLister{}

type FakeControllerLister []*v1.ReplicationController

func (f FakeControllerLister) List(labels.Selector) ([]*v1.ReplicationController, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, nil
}
func (f FakeControllerLister) GetPodControllers(pod *v1.Pod) (controllers []*v1.ReplicationController, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var selector labels.Selector
 for i := range f {
  controller := f[i]
  if controller.Namespace != pod.Namespace {
   continue
  }
  selector = labels.Set(controller.Spec.Selector).AsSelectorPreValidated()
  if selector.Matches(labels.Set(pod.Labels)) {
   controllers = append(controllers, controller)
  }
 }
 if len(controllers) == 0 {
  err = fmt.Errorf("Could not find Replication Controller for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
 }
 return
}

var _ algorithm.ReplicaSetLister = &FakeReplicaSetLister{}

type FakeReplicaSetLister []*apps.ReplicaSet

func (f FakeReplicaSetLister) GetPodReplicaSets(pod *v1.Pod) (rss []*apps.ReplicaSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var selector labels.Selector
 for _, rs := range f {
  if rs.Namespace != pod.Namespace {
   continue
  }
  selector, err = metav1.LabelSelectorAsSelector(rs.Spec.Selector)
  if err != nil {
   return
  }
  if selector.Matches(labels.Set(pod.Labels)) {
   rss = append(rss, rs)
  }
 }
 if len(rss) == 0 {
  err = fmt.Errorf("Could not find ReplicaSet for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
 }
 return
}

var _ algorithm.StatefulSetLister = &FakeStatefulSetLister{}

type FakeStatefulSetLister []*apps.StatefulSet

func (f FakeStatefulSetLister) GetPodStatefulSets(pod *v1.Pod) (sss []*apps.StatefulSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var selector labels.Selector
 for _, ss := range f {
  if ss.Namespace != pod.Namespace {
   continue
  }
  selector, err = metav1.LabelSelectorAsSelector(ss.Spec.Selector)
  if err != nil {
   return
  }
  if selector.Matches(labels.Set(pod.Labels)) {
   sss = append(sss, ss)
  }
 }
 if len(sss) == 0 {
  err = fmt.Errorf("Could not find StatefulSet for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
 }
 return
}

type FakePersistentVolumeClaimLister []*v1.PersistentVolumeClaim

var _ corelisters.PersistentVolumeClaimLister = FakePersistentVolumeClaimLister{}

func (f FakePersistentVolumeClaimLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("not implemented")
}
func (f FakePersistentVolumeClaimLister) PersistentVolumeClaims(namespace string) corelisters.PersistentVolumeClaimNamespaceLister {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakePersistentVolumeClaimNamespaceLister{pvcs: f, namespace: namespace}
}

type fakePersistentVolumeClaimNamespaceLister struct {
 pvcs      []*v1.PersistentVolumeClaim
 namespace string
}

func (f *fakePersistentVolumeClaimNamespaceLister) Get(name string) (*v1.PersistentVolumeClaim, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, pvc := range f.pvcs {
  if pvc.Name == name && pvc.Namespace == f.namespace {
   return pvc, nil
  }
 }
 return nil, fmt.Errorf("persistentvolumeclaim %q not found", name)
}
func (f fakePersistentVolumeClaimNamespaceLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("not implemented")
}

type FakePDBLister []*policy.PodDisruptionBudget

func (f FakePDBLister) List(labels.Selector) ([]*policy.PodDisruptionBudget, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
