package controller

import (
 "encoding/binary"
 "encoding/json"
 "fmt"
 "hash/fnv"
 "sync"
 "sync/atomic"
 "time"
 apps "k8s.io/api/apps/v1"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 "k8s.io/apimachinery/pkg/api/meta"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/clock"
 "k8s.io/apimachinery/pkg/util/rand"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/strategicpatch"
 "k8s.io/apimachinery/pkg/util/wait"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/tools/record"
 "k8s.io/client-go/util/integer"
 clientretry "k8s.io/client-go/util/retry"
 podutil "k8s.io/kubernetes/pkg/api/v1/pod"
 _ "k8s.io/kubernetes/pkg/apis/core/install"
 "k8s.io/kubernetes/pkg/apis/core/validation"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 hashutil "k8s.io/kubernetes/pkg/util/hash"
 taintutils "k8s.io/kubernetes/pkg/util/taints"
 "k8s.io/klog"
)

const (
 ExpectationsTimeout       = 5 * time.Minute
 SlowStartInitialBatchSize = 1
)

var UpdateTaintBackoff = wait.Backoff{Steps: 5, Duration: 100 * time.Millisecond, Jitter: 1.0}
var ShutdownTaint = &v1.Taint{Key: schedulerapi.TaintNodeShutdown, Effect: v1.TaintEffectNoSchedule}
var (
 KeyFunc = cache.DeletionHandlingMetaNamespaceKeyFunc
)

type ResyncPeriodFunc func() time.Duration

func NoResyncPeriodFunc() time.Duration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return 0
}
func StaticResyncPeriodFunc(resyncPeriod time.Duration) ResyncPeriodFunc {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func() time.Duration {
  return resyncPeriod
 }
}

var ExpKeyFunc = func(obj interface{}) (string, error) {
 if e, ok := obj.(*ControlleeExpectations); ok {
  return e.key, nil
 }
 return "", fmt.Errorf("Could not find key for obj %#v", obj)
}

type ControllerExpectationsInterface interface {
 GetExpectations(controllerKey string) (*ControlleeExpectations, bool, error)
 SatisfiedExpectations(controllerKey string) bool
 DeleteExpectations(controllerKey string)
 SetExpectations(controllerKey string, add, del int) error
 ExpectCreations(controllerKey string, adds int) error
 ExpectDeletions(controllerKey string, dels int) error
 CreationObserved(controllerKey string)
 DeletionObserved(controllerKey string)
 RaiseExpectations(controllerKey string, add, del int)
 LowerExpectations(controllerKey string, add, del int)
}
type ControllerExpectations struct{ cache.Store }

func (r *ControllerExpectations) GetExpectations(controllerKey string) (*ControlleeExpectations, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if exp, exists, err := r.GetByKey(controllerKey); err == nil && exists {
  return exp.(*ControlleeExpectations), true, nil
 } else {
  return nil, false, err
 }
}
func (r *ControllerExpectations) DeleteExpectations(controllerKey string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if exp, exists, err := r.GetByKey(controllerKey); err == nil && exists {
  if err := r.Delete(exp); err != nil {
   klog.V(2).Infof("Error deleting expectations for controller %v: %v", controllerKey, err)
  }
 }
}
func (r *ControllerExpectations) SatisfiedExpectations(controllerKey string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if exp, exists, err := r.GetExpectations(controllerKey); exists {
  if exp.Fulfilled() {
   klog.V(4).Infof("Controller expectations fulfilled %#v", exp)
   return true
  } else if exp.isExpired() {
   klog.V(4).Infof("Controller expectations expired %#v", exp)
   return true
  } else {
   klog.V(4).Infof("Controller still waiting on expectations %#v", exp)
   return false
  }
 } else if err != nil {
  klog.V(2).Infof("Error encountered while checking expectations %#v, forcing sync", err)
 } else {
  klog.V(4).Infof("Controller %v either never recorded expectations, or the ttl expired.", controllerKey)
 }
 return true
}
func (exp *ControlleeExpectations) isExpired() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return clock.RealClock{}.Since(exp.timestamp) > ExpectationsTimeout
}
func (r *ControllerExpectations) SetExpectations(controllerKey string, add, del int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 exp := &ControlleeExpectations{add: int64(add), del: int64(del), key: controllerKey, timestamp: clock.RealClock{}.Now()}
 klog.V(4).Infof("Setting expectations %#v", exp)
 return r.Add(exp)
}
func (r *ControllerExpectations) ExpectCreations(controllerKey string, adds int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.SetExpectations(controllerKey, adds, 0)
}
func (r *ControllerExpectations) ExpectDeletions(controllerKey string, dels int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.SetExpectations(controllerKey, 0, dels)
}
func (r *ControllerExpectations) LowerExpectations(controllerKey string, add, del int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if exp, exists, err := r.GetExpectations(controllerKey); err == nil && exists {
  exp.Add(int64(-add), int64(-del))
  klog.V(4).Infof("Lowered expectations %#v", exp)
 }
}
func (r *ControllerExpectations) RaiseExpectations(controllerKey string, add, del int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if exp, exists, err := r.GetExpectations(controllerKey); err == nil && exists {
  exp.Add(int64(add), int64(del))
  klog.V(4).Infof("Raised expectations %#v", exp)
 }
}
func (r *ControllerExpectations) CreationObserved(controllerKey string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.LowerExpectations(controllerKey, 1, 0)
}
func (r *ControllerExpectations) DeletionObserved(controllerKey string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.LowerExpectations(controllerKey, 0, 1)
}

type Expectations interface{ Fulfilled() bool }
type ControlleeExpectations struct {
 add       int64
 del       int64
 key       string
 timestamp time.Time
}

func (e *ControlleeExpectations) Add(add, del int64) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 atomic.AddInt64(&e.add, add)
 atomic.AddInt64(&e.del, del)
}
func (e *ControlleeExpectations) Fulfilled() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return atomic.LoadInt64(&e.add) <= 0 && atomic.LoadInt64(&e.del) <= 0
}
func (e *ControlleeExpectations) GetExpectations() (int64, int64) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return atomic.LoadInt64(&e.add), atomic.LoadInt64(&e.del)
}
func NewControllerExpectations() *ControllerExpectations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &ControllerExpectations{cache.NewStore(ExpKeyFunc)}
}

var UIDSetKeyFunc = func(obj interface{}) (string, error) {
 if u, ok := obj.(*UIDSet); ok {
  return u.key, nil
 }
 return "", fmt.Errorf("Could not find key for obj %#v", obj)
}

type UIDSet struct {
 sets.String
 key string
}
type UIDTrackingControllerExpectations struct {
 ControllerExpectationsInterface
 uidStoreLock sync.Mutex
 uidStore     cache.Store
}

func (u *UIDTrackingControllerExpectations) GetUIDs(controllerKey string) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if uid, exists, err := u.uidStore.GetByKey(controllerKey); err == nil && exists {
  return uid.(*UIDSet).String
 }
 return nil
}
func (u *UIDTrackingControllerExpectations) ExpectDeletions(rcKey string, deletedKeys []string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 u.uidStoreLock.Lock()
 defer u.uidStoreLock.Unlock()
 if existing := u.GetUIDs(rcKey); existing != nil && existing.Len() != 0 {
  klog.Errorf("Clobbering existing delete keys: %+v", existing)
 }
 expectedUIDs := sets.NewString()
 for _, k := range deletedKeys {
  expectedUIDs.Insert(k)
 }
 klog.V(4).Infof("Controller %v waiting on deletions for: %+v", rcKey, deletedKeys)
 if err := u.uidStore.Add(&UIDSet{expectedUIDs, rcKey}); err != nil {
  return err
 }
 return u.ControllerExpectationsInterface.ExpectDeletions(rcKey, expectedUIDs.Len())
}
func (u *UIDTrackingControllerExpectations) DeletionObserved(rcKey, deleteKey string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 u.uidStoreLock.Lock()
 defer u.uidStoreLock.Unlock()
 uids := u.GetUIDs(rcKey)
 if uids != nil && uids.Has(deleteKey) {
  klog.V(4).Infof("Controller %v received delete for pod %v", rcKey, deleteKey)
  u.ControllerExpectationsInterface.DeletionObserved(rcKey)
  uids.Delete(deleteKey)
 }
}
func (u *UIDTrackingControllerExpectations) DeleteExpectations(rcKey string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 u.uidStoreLock.Lock()
 defer u.uidStoreLock.Unlock()
 u.ControllerExpectationsInterface.DeleteExpectations(rcKey)
 if uidExp, exists, err := u.uidStore.GetByKey(rcKey); err == nil && exists {
  if err := u.uidStore.Delete(uidExp); err != nil {
   klog.V(2).Infof("Error deleting uid expectations for controller %v: %v", rcKey, err)
  }
 }
}
func NewUIDTrackingControllerExpectations(ce ControllerExpectationsInterface) *UIDTrackingControllerExpectations {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &UIDTrackingControllerExpectations{ControllerExpectationsInterface: ce, uidStore: cache.NewStore(UIDSetKeyFunc)}
}

const (
 FailedCreatePodReason     = "FailedCreate"
 SuccessfulCreatePodReason = "SuccessfulCreate"
 FailedDeletePodReason     = "FailedDelete"
 SuccessfulDeletePodReason = "SuccessfulDelete"
)

type RSControlInterface interface {
 PatchReplicaSet(namespace, name string, data []byte) error
}
type RealRSControl struct {
 KubeClient clientset.Interface
 Recorder   record.EventRecorder
}

var _ RSControlInterface = &RealRSControl{}

func (r RealRSControl) PatchReplicaSet(namespace, name string, data []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := r.KubeClient.ExtensionsV1beta1().ReplicaSets(namespace).Patch(name, types.StrategicMergePatchType, data)
 return err
}

type ControllerRevisionControlInterface interface {
 PatchControllerRevision(namespace, name string, data []byte) error
}
type RealControllerRevisionControl struct{ KubeClient clientset.Interface }

var _ ControllerRevisionControlInterface = &RealControllerRevisionControl{}

func (r RealControllerRevisionControl) PatchControllerRevision(namespace, name string, data []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := r.KubeClient.AppsV1beta1().ControllerRevisions(namespace).Patch(name, types.StrategicMergePatchType, data)
 return err
}

type PodControlInterface interface {
 CreatePods(namespace string, template *v1.PodTemplateSpec, object runtime.Object) error
 CreatePodsOnNode(nodeName, namespace string, template *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error
 CreatePodsWithControllerRef(namespace string, template *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error
 DeletePod(namespace string, podID string, object runtime.Object) error
 PatchPod(namespace, name string, data []byte) error
}
type RealPodControl struct {
 KubeClient clientset.Interface
 Recorder   record.EventRecorder
}

var _ PodControlInterface = &RealPodControl{}

func getPodsLabelSet(template *v1.PodTemplateSpec) labels.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 desiredLabels := make(labels.Set)
 for k, v := range template.Labels {
  desiredLabels[k] = v
 }
 return desiredLabels
}
func getPodsFinalizers(template *v1.PodTemplateSpec) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 desiredFinalizers := make([]string, len(template.Finalizers))
 copy(desiredFinalizers, template.Finalizers)
 return desiredFinalizers
}
func getPodsAnnotationSet(template *v1.PodTemplateSpec) labels.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 desiredAnnotations := make(labels.Set)
 for k, v := range template.Annotations {
  desiredAnnotations[k] = v
 }
 return desiredAnnotations
}
func getPodsPrefix(controllerName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prefix := fmt.Sprintf("%s-", controllerName)
 if len(validation.ValidatePodName(prefix, true)) != 0 {
  prefix = controllerName
 }
 return prefix
}
func validateControllerRef(controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if controllerRef == nil {
  return fmt.Errorf("controllerRef is nil")
 }
 if len(controllerRef.APIVersion) == 0 {
  return fmt.Errorf("controllerRef has empty APIVersion")
 }
 if len(controllerRef.Kind) == 0 {
  return fmt.Errorf("controllerRef has empty Kind")
 }
 if controllerRef.Controller == nil || *controllerRef.Controller != true {
  return fmt.Errorf("controllerRef.Controller is not set to true")
 }
 if controllerRef.BlockOwnerDeletion == nil || *controllerRef.BlockOwnerDeletion != true {
  return fmt.Errorf("controllerRef.BlockOwnerDeletion is not set")
 }
 return nil
}
func (r RealPodControl) CreatePods(namespace string, template *v1.PodTemplateSpec, object runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.createPods("", namespace, template, object, nil)
}
func (r RealPodControl) CreatePodsWithControllerRef(namespace string, template *v1.PodTemplateSpec, controllerObject runtime.Object, controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := validateControllerRef(controllerRef); err != nil {
  return err
 }
 return r.createPods("", namespace, template, controllerObject, controllerRef)
}
func (r RealPodControl) CreatePodsOnNode(nodeName, namespace string, template *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := validateControllerRef(controllerRef); err != nil {
  return err
 }
 return r.createPods(nodeName, namespace, template, object, controllerRef)
}
func (r RealPodControl) PatchPod(namespace, name string, data []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := r.KubeClient.CoreV1().Pods(namespace).Patch(name, types.StrategicMergePatchType, data)
 return err
}
func GetPodFromTemplate(template *v1.PodTemplateSpec, parentObject runtime.Object, controllerRef *metav1.OwnerReference) (*v1.Pod, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 desiredLabels := getPodsLabelSet(template)
 desiredFinalizers := getPodsFinalizers(template)
 desiredAnnotations := getPodsAnnotationSet(template)
 accessor, err := meta.Accessor(parentObject)
 if err != nil {
  return nil, fmt.Errorf("parentObject does not have ObjectMeta, %v", err)
 }
 prefix := getPodsPrefix(accessor.GetName())
 pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Labels: desiredLabels, Annotations: desiredAnnotations, GenerateName: prefix, Finalizers: desiredFinalizers}}
 if controllerRef != nil {
  pod.OwnerReferences = append(pod.OwnerReferences, *controllerRef)
 }
 pod.Spec = *template.Spec.DeepCopy()
 return pod, nil
}
func (r RealPodControl) createPods(nodeName, namespace string, template *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pod, err := GetPodFromTemplate(template, object, controllerRef)
 if err != nil {
  return err
 }
 if len(nodeName) != 0 {
  pod.Spec.NodeName = nodeName
 }
 if labels.Set(pod.Labels).AsSelectorPreValidated().Empty() {
  return fmt.Errorf("unable to create pods, no labels")
 }
 if newPod, err := r.KubeClient.CoreV1().Pods(namespace).Create(pod); err != nil {
  r.Recorder.Eventf(object, v1.EventTypeWarning, FailedCreatePodReason, "Error creating: %v", err)
  return err
 } else {
  accessor, err := meta.Accessor(object)
  if err != nil {
   klog.Errorf("parentObject does not have ObjectMeta, %v", err)
   return nil
  }
  klog.V(4).Infof("Controller %v created pod %v", accessor.GetName(), newPod.Name)
  r.Recorder.Eventf(object, v1.EventTypeNormal, SuccessfulCreatePodReason, "Created pod: %v", newPod.Name)
 }
 return nil
}
func (r RealPodControl) DeletePod(namespace string, podID string, object runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 accessor, err := meta.Accessor(object)
 if err != nil {
  return fmt.Errorf("object does not have ObjectMeta, %v", err)
 }
 klog.V(2).Infof("Controller %v deleting pod %v/%v", accessor.GetName(), namespace, podID)
 if err := r.KubeClient.CoreV1().Pods(namespace).Delete(podID, nil); err != nil && !apierrors.IsNotFound(err) {
  r.Recorder.Eventf(object, v1.EventTypeWarning, FailedDeletePodReason, "Error deleting: %v", err)
  return fmt.Errorf("unable to delete pods: %v", err)
 } else {
  r.Recorder.Eventf(object, v1.EventTypeNormal, SuccessfulDeletePodReason, "Deleted pod: %v", podID)
 }
 return nil
}

type FakePodControl struct {
 sync.Mutex
 Templates       []v1.PodTemplateSpec
 ControllerRefs  []metav1.OwnerReference
 DeletePodName   []string
 Patches         [][]byte
 Err             error
 CreateLimit     int
 CreateCallCount int
}

var _ PodControlInterface = &FakePodControl{}

func (f *FakePodControl) PatchPod(namespace, name string, data []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.Patches = append(f.Patches, data)
 if f.Err != nil {
  return f.Err
 }
 return nil
}
func (f *FakePodControl) CreatePods(namespace string, spec *v1.PodTemplateSpec, object runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.CreateCallCount++
 if f.CreateLimit != 0 && f.CreateCallCount > f.CreateLimit {
  return fmt.Errorf("Not creating pod, limit %d already reached (create call %d)", f.CreateLimit, f.CreateCallCount)
 }
 f.Templates = append(f.Templates, *spec)
 if f.Err != nil {
  return f.Err
 }
 return nil
}
func (f *FakePodControl) CreatePodsWithControllerRef(namespace string, spec *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.CreateCallCount++
 if f.CreateLimit != 0 && f.CreateCallCount > f.CreateLimit {
  return fmt.Errorf("Not creating pod, limit %d already reached (create call %d)", f.CreateLimit, f.CreateCallCount)
 }
 f.Templates = append(f.Templates, *spec)
 f.ControllerRefs = append(f.ControllerRefs, *controllerRef)
 if f.Err != nil {
  return f.Err
 }
 return nil
}
func (f *FakePodControl) CreatePodsOnNode(nodeName, namespace string, template *v1.PodTemplateSpec, object runtime.Object, controllerRef *metav1.OwnerReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.CreateCallCount++
 if f.CreateLimit != 0 && f.CreateCallCount > f.CreateLimit {
  return fmt.Errorf("Not creating pod, limit %d already reached (create call %d)", f.CreateLimit, f.CreateCallCount)
 }
 f.Templates = append(f.Templates, *template)
 f.ControllerRefs = append(f.ControllerRefs, *controllerRef)
 if f.Err != nil {
  return f.Err
 }
 return nil
}
func (f *FakePodControl) DeletePod(namespace string, podID string, object runtime.Object) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.DeletePodName = append(f.DeletePodName, podID)
 if f.Err != nil {
  return f.Err
 }
 return nil
}
func (f *FakePodControl) Clear() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock()
 defer f.Unlock()
 f.DeletePodName = []string{}
 f.Templates = []v1.PodTemplateSpec{}
 f.ControllerRefs = []metav1.OwnerReference{}
 f.Patches = [][]byte{}
 f.CreateLimit = 0
 f.CreateCallCount = 0
}

type ByLogging []*v1.Pod

func (s ByLogging) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(s)
}
func (s ByLogging) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s[i], s[j] = s[j], s[i]
}
func (s ByLogging) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if s[i].Spec.NodeName != s[j].Spec.NodeName && (len(s[i].Spec.NodeName) == 0 || len(s[j].Spec.NodeName) == 0) {
  return len(s[i].Spec.NodeName) > 0
 }
 m := map[v1.PodPhase]int{v1.PodRunning: 0, v1.PodUnknown: 1, v1.PodPending: 2}
 if m[s[i].Status.Phase] != m[s[j].Status.Phase] {
  return m[s[i].Status.Phase] < m[s[j].Status.Phase]
 }
 if podutil.IsPodReady(s[i]) != podutil.IsPodReady(s[j]) {
  return podutil.IsPodReady(s[i])
 }
 if podutil.IsPodReady(s[i]) && podutil.IsPodReady(s[j]) && !podReadyTime(s[i]).Equal(podReadyTime(s[j])) {
  return afterOrZero(podReadyTime(s[j]), podReadyTime(s[i]))
 }
 if maxContainerRestarts(s[i]) != maxContainerRestarts(s[j]) {
  return maxContainerRestarts(s[i]) > maxContainerRestarts(s[j])
 }
 if !s[i].CreationTimestamp.Equal(&s[j].CreationTimestamp) {
  return afterOrZero(&s[j].CreationTimestamp, &s[i].CreationTimestamp)
 }
 return false
}

type ActivePods []*v1.Pod

func (s ActivePods) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(s)
}
func (s ActivePods) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s[i], s[j] = s[j], s[i]
}
func (s ActivePods) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if s[i].Spec.NodeName != s[j].Spec.NodeName && (len(s[i].Spec.NodeName) == 0 || len(s[j].Spec.NodeName) == 0) {
  return len(s[i].Spec.NodeName) == 0
 }
 m := map[v1.PodPhase]int{v1.PodPending: 0, v1.PodUnknown: 1, v1.PodRunning: 2}
 if m[s[i].Status.Phase] != m[s[j].Status.Phase] {
  return m[s[i].Status.Phase] < m[s[j].Status.Phase]
 }
 if podutil.IsPodReady(s[i]) != podutil.IsPodReady(s[j]) {
  return !podutil.IsPodReady(s[i])
 }
 if podutil.IsPodReady(s[i]) && podutil.IsPodReady(s[j]) && !podReadyTime(s[i]).Equal(podReadyTime(s[j])) {
  return afterOrZero(podReadyTime(s[i]), podReadyTime(s[j]))
 }
 if maxContainerRestarts(s[i]) != maxContainerRestarts(s[j]) {
  return maxContainerRestarts(s[i]) > maxContainerRestarts(s[j])
 }
 if !s[i].CreationTimestamp.Equal(&s[j].CreationTimestamp) {
  return afterOrZero(&s[i].CreationTimestamp, &s[j].CreationTimestamp)
 }
 return false
}
func afterOrZero(t1, t2 *metav1.Time) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if t1.Time.IsZero() || t2.Time.IsZero() {
  return t1.Time.IsZero()
 }
 return t1.After(t2.Time)
}
func podReadyTime(pod *v1.Pod) *metav1.Time {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if podutil.IsPodReady(pod) {
  for _, c := range pod.Status.Conditions {
   if c.Type == v1.PodReady && c.Status == v1.ConditionTrue {
    return &c.LastTransitionTime
   }
  }
 }
 return &metav1.Time{}
}
func maxContainerRestarts(pod *v1.Pod) int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 maxRestarts := 0
 for _, c := range pod.Status.ContainerStatuses {
  maxRestarts = integer.IntMax(maxRestarts, int(c.RestartCount))
 }
 return maxRestarts
}
func FilterActivePods(pods []*v1.Pod) []*v1.Pod {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var result []*v1.Pod
 for _, p := range pods {
  if IsPodActive(p) {
   result = append(result, p)
  } else {
   klog.V(4).Infof("Ignoring inactive pod %v/%v in state %v, deletion time %v", p.Namespace, p.Name, p.Status.Phase, p.DeletionTimestamp)
  }
 }
 return result
}
func IsPodActive(p *v1.Pod) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return v1.PodSucceeded != p.Status.Phase && v1.PodFailed != p.Status.Phase && p.DeletionTimestamp == nil
}
func FilterActiveReplicaSets(replicaSets []*apps.ReplicaSet) []*apps.ReplicaSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 activeFilter := func(rs *apps.ReplicaSet) bool {
  return rs != nil && *(rs.Spec.Replicas) > 0
 }
 return FilterReplicaSets(replicaSets, activeFilter)
}

type filterRS func(rs *apps.ReplicaSet) bool

func FilterReplicaSets(RSes []*apps.ReplicaSet, filterFn filterRS) []*apps.ReplicaSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var filtered []*apps.ReplicaSet
 for i := range RSes {
  if filterFn(RSes[i]) {
   filtered = append(filtered, RSes[i])
  }
 }
 return filtered
}
func PodKey(pod *v1.Pod) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%v/%v", pod.Namespace, pod.Name)
}

type ControllersByCreationTimestamp []*v1.ReplicationController

func (o ControllersByCreationTimestamp) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o ControllersByCreationTimestamp) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o ControllersByCreationTimestamp) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
  return o[i].Name < o[j].Name
 }
 return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}

type ReplicaSetsByCreationTimestamp []*apps.ReplicaSet

func (o ReplicaSetsByCreationTimestamp) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o ReplicaSetsByCreationTimestamp) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o ReplicaSetsByCreationTimestamp) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
  return o[i].Name < o[j].Name
 }
 return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}

type ReplicaSetsBySizeOlder []*apps.ReplicaSet

func (o ReplicaSetsBySizeOlder) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o ReplicaSetsBySizeOlder) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o ReplicaSetsBySizeOlder) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if *(o[i].Spec.Replicas) == *(o[j].Spec.Replicas) {
  return ReplicaSetsByCreationTimestamp(o).Less(i, j)
 }
 return *(o[i].Spec.Replicas) > *(o[j].Spec.Replicas)
}

type ReplicaSetsBySizeNewer []*apps.ReplicaSet

func (o ReplicaSetsBySizeNewer) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(o)
}
func (o ReplicaSetsBySizeNewer) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o[i], o[j] = o[j], o[i]
}
func (o ReplicaSetsBySizeNewer) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if *(o[i].Spec.Replicas) == *(o[j].Spec.Replicas) {
  return ReplicaSetsByCreationTimestamp(o).Less(j, i)
 }
 return *(o[i].Spec.Replicas) > *(o[j].Spec.Replicas)
}
func AddOrUpdateTaintOnNode(c clientset.Interface, nodeName string, taints ...*v1.Taint) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(taints) == 0 {
  return nil
 }
 firstTry := true
 return clientretry.RetryOnConflict(UpdateTaintBackoff, func() error {
  var err error
  var oldNode *v1.Node
  if firstTry {
   oldNode, err = c.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{ResourceVersion: "0"})
   firstTry = false
  } else {
   oldNode, err = c.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
  }
  if err != nil {
   return err
  }
  var newNode *v1.Node
  oldNodeCopy := oldNode
  updated := false
  for _, taint := range taints {
   curNewNode, ok, err := taintutils.AddOrUpdateTaint(oldNodeCopy, taint)
   if err != nil {
    return fmt.Errorf("Failed to update taint of node!")
   }
   updated = updated || ok
   newNode = curNewNode
   oldNodeCopy = curNewNode
  }
  if !updated {
   return nil
  }
  return PatchNodeTaints(c, nodeName, oldNode, newNode)
 })
}
func RemoveTaintOffNode(c clientset.Interface, nodeName string, node *v1.Node, taints ...*v1.Taint) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(taints) == 0 {
  return nil
 }
 if node != nil {
  match := false
  for _, taint := range taints {
   if taintutils.TaintExists(node.Spec.Taints, taint) {
    match = true
    break
   }
  }
  if !match {
   return nil
  }
 }
 firstTry := true
 return clientretry.RetryOnConflict(UpdateTaintBackoff, func() error {
  var err error
  var oldNode *v1.Node
  if firstTry {
   oldNode, err = c.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{ResourceVersion: "0"})
   firstTry = false
  } else {
   oldNode, err = c.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
  }
  if err != nil {
   return err
  }
  var newNode *v1.Node
  oldNodeCopy := oldNode
  updated := false
  for _, taint := range taints {
   curNewNode, ok, err := taintutils.RemoveTaint(oldNodeCopy, taint)
   if err != nil {
    return fmt.Errorf("Failed to remove taint of node!")
   }
   updated = updated || ok
   newNode = curNewNode
   oldNodeCopy = curNewNode
  }
  if !updated {
   return nil
  }
  return PatchNodeTaints(c, nodeName, oldNode, newNode)
 })
}
func PatchNodeTaints(c clientset.Interface, nodeName string, oldNode *v1.Node, newNode *v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldData, err := json.Marshal(oldNode)
 if err != nil {
  return fmt.Errorf("failed to marshal old node %#v for node %q: %v", oldNode, nodeName, err)
 }
 newTaints := newNode.Spec.Taints
 newNodeClone := oldNode.DeepCopy()
 newNodeClone.Spec.Taints = newTaints
 newData, err := json.Marshal(newNodeClone)
 if err != nil {
  return fmt.Errorf("failed to marshal new node %#v for node %q: %v", newNodeClone, nodeName, err)
 }
 patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, v1.Node{})
 if err != nil {
  return fmt.Errorf("failed to create patch for node %q: %v", nodeName, err)
 }
 _, err = c.CoreV1().Nodes().Patch(nodeName, types.StrategicMergePatchType, patchBytes)
 return err
}
func WaitForCacheSync(controllerName string, stopCh <-chan struct{}, cacheSyncs ...cache.InformerSynced) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.Infof("Waiting for caches to sync for %s controller", controllerName)
 if !cache.WaitForCacheSync(stopCh, cacheSyncs...) {
  utilruntime.HandleError(fmt.Errorf("Unable to sync caches for %s controller", controllerName))
  return false
 }
 klog.Infof("Caches are synced for %s controller", controllerName)
 return true
}
func ComputeHash(template *v1.PodTemplateSpec, collisionCount *int32) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podTemplateSpecHasher := fnv.New32a()
 hashutil.DeepHashObject(podTemplateSpecHasher, *template)
 if collisionCount != nil {
  collisionCountBytes := make([]byte, 8)
  binary.LittleEndian.PutUint32(collisionCountBytes, uint32(*collisionCount))
  podTemplateSpecHasher.Write(collisionCountBytes)
 }
 return rand.SafeEncodeString(fmt.Sprint(podTemplateSpecHasher.Sum32()))
}
