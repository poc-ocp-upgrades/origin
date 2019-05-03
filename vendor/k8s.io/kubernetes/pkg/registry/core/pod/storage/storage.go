package storage

import (
 "context"
 "fmt"
 "net/http"
 "net/url"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/registry/generic"
 genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
 "k8s.io/apiserver/pkg/registry/rest"
 "k8s.io/apiserver/pkg/storage"
 storeerr "k8s.io/apiserver/pkg/storage/errors"
 "k8s.io/apiserver/pkg/util/dryrun"
 podutil "k8s.io/kubernetes/pkg/api/pod"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
 policyclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
 "k8s.io/kubernetes/pkg/kubelet/client"
 "k8s.io/kubernetes/pkg/printers"
 printersinternal "k8s.io/kubernetes/pkg/printers/internalversion"
 printerstorage "k8s.io/kubernetes/pkg/printers/storage"
 "k8s.io/kubernetes/pkg/registry/core/pod"
 podrest "k8s.io/kubernetes/pkg/registry/core/pod/rest"
)

type PodStorage struct {
 Pod         *REST
 Binding     *BindingREST
 Eviction    *EvictionREST
 Status      *StatusREST
 Log         *podrest.LogREST
 Proxy       *podrest.ProxyREST
 Exec        *podrest.ExecREST
 Attach      *podrest.AttachREST
 PortForward *podrest.PortForwardREST
}
type REST struct {
 *genericregistry.Store
 proxyTransport http.RoundTripper
}

func NewStorage(optsGetter generic.RESTOptionsGetter, k client.ConnectionInfoGetter, proxyTransport http.RoundTripper, podDisruptionBudgetClient policyclient.PodDisruptionBudgetsGetter) PodStorage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 store := &genericregistry.Store{NewFunc: func() runtime.Object {
  return &api.Pod{}
 }, NewListFunc: func() runtime.Object {
  return &api.PodList{}
 }, PredicateFunc: pod.MatchPod, DefaultQualifiedResource: api.Resource("pods"), CreateStrategy: pod.Strategy, UpdateStrategy: pod.Strategy, DeleteStrategy: pod.Strategy, ReturnDeletedObject: true, TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
 options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: pod.GetAttrs, TriggerFunc: pod.NodeNameTriggerFunc}
 if err := store.CompleteWithOptions(options); err != nil {
  panic(err)
 }
 statusStore := *store
 statusStore.UpdateStrategy = pod.StatusStrategy
 return PodStorage{Pod: &REST{store, proxyTransport}, Binding: &BindingREST{store: store}, Eviction: newEvictionStorage(store, podDisruptionBudgetClient), Status: &StatusREST{store: &statusStore}, Log: &podrest.LogREST{Store: store, KubeletConn: k}, Proxy: &podrest.ProxyREST{Store: store, ProxyTransport: proxyTransport}, Exec: &podrest.ExecREST{Store: store, KubeletConn: k}, Attach: &podrest.AttachREST{Store: store, KubeletConn: k}, PortForward: &podrest.PortForwardREST{Store: store, KubeletConn: k}}
}

var _ = rest.Redirector(&REST{})

func (r *REST) ResourceLocation(ctx context.Context, name string) (*url.URL, http.RoundTripper, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return pod.ResourceLocation(r, r.proxyTransport, ctx, name)
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"po"}
}

var _ rest.CategoriesProvider = &REST{}

func (r *REST) Categories() []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"all"}
}

type BindingREST struct{ store *genericregistry.Store }

func (r *BindingREST) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.store.NamespaceScoped()
}
func (r *BindingREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.Binding{}
}

var _ = rest.Creater(&BindingREST{})

func (r *BindingREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (out runtime.Object, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 binding := obj.(*api.Binding)
 if errs := validation.ValidatePodBinding(binding); len(errs) != 0 {
  return nil, errs.ToAggregate()
 }
 err = r.assignPod(ctx, binding.Name, binding.Target.Name, binding.Annotations, dryrun.IsDryRun(options.DryRun))
 out = &metav1.Status{Status: metav1.StatusSuccess}
 return
}
func (r *BindingREST) setPodHostAndAnnotations(ctx context.Context, podID, oldMachine, machine string, annotations map[string]string, dryRun bool) (finalPod *api.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 podKey, err := r.store.KeyFunc(ctx, podID)
 if err != nil {
  return nil, err
 }
 err = r.store.Storage.GuaranteedUpdate(ctx, podKey, &api.Pod{}, false, nil, storage.SimpleUpdate(func(obj runtime.Object) (runtime.Object, error) {
  pod, ok := obj.(*api.Pod)
  if !ok {
   return nil, fmt.Errorf("unexpected object: %#v", obj)
  }
  if pod.DeletionTimestamp != nil {
   return nil, fmt.Errorf("pod %s is being deleted, cannot be assigned to a host", pod.Name)
  }
  if pod.Spec.NodeName != oldMachine {
   return nil, fmt.Errorf("pod %v is already assigned to node %q", pod.Name, pod.Spec.NodeName)
  }
  pod.Spec.NodeName = machine
  if pod.Annotations == nil {
   pod.Annotations = make(map[string]string)
  }
  for k, v := range annotations {
   pod.Annotations[k] = v
  }
  podutil.UpdatePodCondition(&pod.Status, &api.PodCondition{Type: api.PodScheduled, Status: api.ConditionTrue})
  finalPod = pod
  return pod, nil
 }), dryRun)
 return finalPod, err
}
func (r *BindingREST) assignPod(ctx context.Context, podID string, machine string, annotations map[string]string, dryRun bool) (err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, err = r.setPodHostAndAnnotations(ctx, podID, "", machine, annotations, dryRun); err != nil {
  err = storeerr.InterpretGetError(err, api.Resource("pods"), podID)
  err = storeerr.InterpretUpdateError(err, api.Resource("pods"), podID)
  if _, ok := err.(*errors.StatusError); !ok {
   err = errors.NewConflict(api.Resource("pods/binding"), podID, err)
  }
 }
 return
}

type StatusREST struct{ store *genericregistry.Store }

func (r *StatusREST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.Pod{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}
