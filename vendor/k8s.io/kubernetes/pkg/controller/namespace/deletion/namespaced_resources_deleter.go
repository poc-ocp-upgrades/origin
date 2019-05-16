package deletion

import (
	"fmt"
	goformat "fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	v1clientset "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/klog"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

type NamespacedResourcesDeleterInterface interface{ Delete(nsName string) error }

func NewNamespacedResourcesDeleter(nsClient v1clientset.NamespaceInterface, dynamicClient dynamic.Interface, podsGetter v1clientset.PodsGetter, discoverResourcesFn func() ([]*metav1.APIResourceList, error), finalizerToken v1.FinalizerName, deleteNamespaceWhenDone bool) NamespacedResourcesDeleterInterface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	d := &namespacedResourcesDeleter{nsClient: nsClient, dynamicClient: dynamicClient, podsGetter: podsGetter, opCache: &operationNotSupportedCache{m: make(map[operationKey]bool)}, discoverResourcesFn: discoverResourcesFn, finalizerToken: finalizerToken, deleteNamespaceWhenDone: deleteNamespaceWhenDone}
	d.initOpCache()
	return d
}

var _ NamespacedResourcesDeleterInterface = &namespacedResourcesDeleter{}

type namespacedResourcesDeleter struct {
	nsClient                v1clientset.NamespaceInterface
	dynamicClient           dynamic.Interface
	podsGetter              v1clientset.PodsGetter
	opCache                 *operationNotSupportedCache
	discoverResourcesFn     func() ([]*metav1.APIResourceList, error)
	finalizerToken          v1.FinalizerName
	deleteNamespaceWhenDone bool
}

func (d *namespacedResourcesDeleter) Delete(nsName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespace, err := d.nsClient.Get(nsName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	if namespace.DeletionTimestamp == nil {
		return nil
	}
	klog.V(5).Infof("namespace controller - syncNamespace - namespace: %s, finalizerToken: %s", namespace.Name, d.finalizerToken)
	namespace, err = d.retryOnConflictError(namespace, d.updateNamespaceStatusFunc)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	if namespace.DeletionTimestamp.IsZero() {
		return nil
	}
	if d.deleteNamespaceWhenDone && finalized(namespace) {
		return d.deleteNamespace(namespace)
	}
	estimate, err := d.deleteAllContent(namespace.Name, *namespace.DeletionTimestamp)
	_, _ = d.retryOnConflictError(namespace, d.updateSetDeletionFailureFunc(err))
	if err != nil {
		return err
	}
	if estimate > 0 {
		return &ResourcesRemainingError{estimate}
	}
	namespace, err = d.retryOnConflictError(namespace, d.finalizeNamespace)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	if d.deleteNamespaceWhenDone && finalized(namespace) {
		return d.deleteNamespace(namespace)
	}
	return nil
}
func (d *namespacedResourcesDeleter) initOpCache() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resources, err := d.discoverResourcesFn()
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to get all supported resources from server: %v", err))
	}
	if len(resources) == 0 {
		klog.Fatalf("Unable to get any supported resources from server: %v", err)
	}
	deletableGroupVersionResources := []schema.GroupVersionResource{}
	for _, rl := range resources {
		gv, err := schema.ParseGroupVersion(rl.GroupVersion)
		if err != nil {
			klog.Errorf("Failed to parse GroupVersion %q, skipping: %v", rl.GroupVersion, err)
			continue
		}
		for _, r := range rl.APIResources {
			gvr := schema.GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: r.Name}
			verbs := sets.NewString([]string(r.Verbs)...)
			if !verbs.Has("delete") {
				klog.V(6).Infof("Skipping resource %v because it cannot be deleted.", gvr)
			}
			for _, op := range []operation{operationList, operationDeleteCollection} {
				if !verbs.Has(string(op)) {
					d.opCache.setNotSupported(operationKey{operation: op, gvr: gvr})
				}
			}
			deletableGroupVersionResources = append(deletableGroupVersionResources, gvr)
		}
	}
}
func (d *namespacedResourcesDeleter) deleteNamespace(namespace *v1.Namespace) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var opts *metav1.DeleteOptions
	uid := namespace.UID
	if len(uid) > 0 {
		opts = &metav1.DeleteOptions{Preconditions: &metav1.Preconditions{UID: &uid}}
	}
	err := d.nsClient.Delete(namespace.Name, opts)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	return nil
}

type ResourcesRemainingError struct{ Estimate int64 }

func (e *ResourcesRemainingError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("some content remains in the namespace, estimate %d seconds before it is removed", e.Estimate)
}

type operation string

const (
	operationDeleteCollection operation = "deletecollection"
	operationList             operation = "list"
	finalizerEstimateSeconds  int64     = int64(15)
)

type operationKey struct {
	operation operation
	gvr       schema.GroupVersionResource
}
type operationNotSupportedCache struct {
	lock sync.RWMutex
	m    map[operationKey]bool
}

func (o *operationNotSupportedCache) isSupported(key operationKey) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.lock.RLock()
	defer o.lock.RUnlock()
	return !o.m[key]
}
func (o *operationNotSupportedCache) setNotSupported(key operationKey) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o.lock.Lock()
	defer o.lock.Unlock()
	o.m[key] = true
}

type updateNamespaceFunc func(namespace *v1.Namespace) (*v1.Namespace, error)

func (d *namespacedResourcesDeleter) retryOnConflictError(namespace *v1.Namespace, fn updateNamespaceFunc) (result *v1.Namespace, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	latestNamespace := namespace
	for {
		result, err = fn(latestNamespace)
		if err == nil {
			return result, nil
		}
		if !errors.IsConflict(err) {
			return nil, err
		}
		prevNamespace := latestNamespace
		latestNamespace, err = d.nsClient.Get(latestNamespace.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if prevNamespace.UID != latestNamespace.UID {
			return nil, fmt.Errorf("namespace uid has changed across retries")
		}
	}
}
func (d *namespacedResourcesDeleter) updateNamespaceStatusFunc(namespace *v1.Namespace) (*v1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if namespace.DeletionTimestamp.IsZero() || namespace.Status.Phase == v1.NamespaceTerminating {
		return namespace, nil
	}
	newNamespace := v1.Namespace{}
	newNamespace.ObjectMeta = namespace.ObjectMeta
	newNamespace.Status = namespace.Status
	newNamespace.Status.Phase = v1.NamespaceTerminating
	return d.nsClient.UpdateStatus(&newNamespace)
}
func (d *namespacedResourcesDeleter) updateSetDeletionFailureFunc(err error) func(namespace *v1.Namespace) (*v1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errorAnnotationName := "namespace-controller.kcm.openshift.io/deletion-error"
	return func(namespace *v1.Namespace) (*v1.Namespace, error) {
		newNamespace := namespace.DeepCopy()
		if err == nil || len(err.Error()) == 0 {
			delete(newNamespace.Annotations, errorAnnotationName)
		} else {
			newNamespace.Annotations[errorAnnotationName] = err.Error()
		}
		return d.nsClient.Update(newNamespace)
	}
}
func finalized(namespace *v1.Namespace) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(namespace.Spec.Finalizers) == 0
}
func (d *namespacedResourcesDeleter) finalizeNamespace(namespace *v1.Namespace) (*v1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	namespaceFinalize := v1.Namespace{}
	namespaceFinalize.ObjectMeta = namespace.ObjectMeta
	namespaceFinalize.Spec = namespace.Spec
	finalizerSet := sets.NewString()
	for i := range namespace.Spec.Finalizers {
		if namespace.Spec.Finalizers[i] != d.finalizerToken {
			finalizerSet.Insert(string(namespace.Spec.Finalizers[i]))
		}
	}
	namespaceFinalize.Spec.Finalizers = make([]v1.FinalizerName, 0, len(finalizerSet))
	for _, value := range finalizerSet.List() {
		namespaceFinalize.Spec.Finalizers = append(namespaceFinalize.Spec.Finalizers, v1.FinalizerName(value))
	}
	namespace, err := d.nsClient.Finalize(&namespaceFinalize)
	if err != nil {
		if errors.IsNotFound(err) {
			return namespace, nil
		}
	}
	return namespace, err
}
func (d *namespacedResourcesDeleter) deleteCollection(gvr schema.GroupVersionResource, namespace string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("namespace controller - deleteCollection - namespace: %s, gvr: %v", namespace, gvr)
	key := operationKey{operation: operationDeleteCollection, gvr: gvr}
	if !d.opCache.isSupported(key) {
		klog.V(5).Infof("namespace controller - deleteCollection ignored since not supported - namespace: %s, gvr: %v", namespace, gvr)
		return false, nil
	}
	background := metav1.DeletePropagationBackground
	opts := &metav1.DeleteOptions{PropagationPolicy: &background}
	err := d.dynamicClient.Resource(gvr).Namespace(namespace).DeleteCollection(opts, metav1.ListOptions{})
	if err == nil {
		return true, nil
	}
	if errors.IsMethodNotSupported(err) || errors.IsNotFound(err) {
		klog.V(5).Infof("namespace controller - deleteCollection not supported - namespace: %s, gvr: %v", namespace, gvr)
		d.opCache.setNotSupported(key)
		return false, nil
	}
	klog.V(5).Infof("namespace controller - deleteCollection unexpected error - namespace: %s, gvr: %v, error: %v", namespace, gvr, err)
	return true, err
}
func (d *namespacedResourcesDeleter) listCollection(gvr schema.GroupVersionResource, namespace string) (*unstructured.UnstructuredList, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("namespace controller - listCollection - namespace: %s, gvr: %v", namespace, gvr)
	key := operationKey{operation: operationList, gvr: gvr}
	if !d.opCache.isSupported(key) {
		klog.V(5).Infof("namespace controller - listCollection ignored since not supported - namespace: %s, gvr: %v", namespace, gvr)
		return nil, false, nil
	}
	unstructuredList, err := d.dynamicClient.Resource(gvr).Namespace(namespace).List(metav1.ListOptions{IncludeUninitialized: true})
	if err == nil {
		return unstructuredList, true, nil
	}
	if errors.IsMethodNotSupported(err) || errors.IsNotFound(err) {
		klog.V(5).Infof("namespace controller - listCollection not supported - namespace: %s, gvr: %v", namespace, gvr)
		d.opCache.setNotSupported(key)
		return nil, false, nil
	}
	return nil, true, err
}
func (d *namespacedResourcesDeleter) deleteEachItem(gvr schema.GroupVersionResource, namespace string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("namespace controller - deleteEachItem - namespace: %s, gvr: %v", namespace, gvr)
	unstructuredList, listSupported, err := d.listCollection(gvr, namespace)
	if err != nil {
		return err
	}
	if !listSupported {
		return nil
	}
	for _, item := range unstructuredList.Items {
		background := metav1.DeletePropagationBackground
		opts := &metav1.DeleteOptions{PropagationPolicy: &background}
		if err = d.dynamicClient.Resource(gvr).Namespace(namespace).Delete(item.GetName(), opts); err != nil && !errors.IsNotFound(err) && !errors.IsMethodNotSupported(err) {
			return err
		}
	}
	return nil
}
func (d *namespacedResourcesDeleter) deleteAllContentForGroupVersionResource(gvr schema.GroupVersionResource, namespace string, namespaceDeletedAt metav1.Time) (int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - namespace: %s, gvr: %v", namespace, gvr)
	estimate, err := d.estimateGracefulTermination(gvr, namespace, namespaceDeletedAt)
	if err != nil {
		klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - unable to estimate - namespace: %s, gvr: %v, err: %v", namespace, gvr, err)
		return estimate, err
	}
	klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - estimate - namespace: %s, gvr: %v, estimate: %v", namespace, gvr, estimate)
	deleteCollectionSupported, err := d.deleteCollection(gvr, namespace)
	if err != nil {
		return estimate, err
	}
	if !deleteCollectionSupported {
		err = d.deleteEachItem(gvr, namespace)
		if err != nil {
			return estimate, err
		}
	}
	klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - checking for no more items in namespace: %s, gvr: %v", namespace, gvr)
	unstructuredList, listSupported, err := d.listCollection(gvr, namespace)
	if err != nil {
		klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - error verifying no items in namespace: %s, gvr: %v, err: %v", namespace, gvr, err)
		return estimate, err
	}
	if !listSupported {
		return estimate, nil
	}
	klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - items remaining - namespace: %s, gvr: %v, items: %v", namespace, gvr, len(unstructuredList.Items))
	if len(unstructuredList.Items) != 0 && estimate == int64(0) {
		for _, item := range unstructuredList.Items {
			if len(item.GetFinalizers()) > 0 {
				klog.V(5).Infof("namespace controller - deleteAllContentForGroupVersionResource - items remaining with finalizers - namespace: %s, gvr: %v, finalizers: %v", namespace, gvr, item.GetFinalizers())
				return finalizerEstimateSeconds, nil
			}
		}
		return estimate, fmt.Errorf("unexpected items still remain in namespace: %s for gvr: %v", namespace, gvr)
	}
	return estimate, nil
}
func (d *namespacedResourcesDeleter) deleteAllContent(namespace string, namespaceDeletedAt metav1.Time) (int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	estimate := int64(0)
	klog.V(4).Infof("namespace controller - deleteAllContent - namespace: %s", namespace)
	resources, err := d.discoverResourcesFn()
	if err != nil {
		errs = append(errs, err)
	}
	deletableResources := discovery.FilteredBy(discovery.SupportsAllVerbs{Verbs: []string{"delete"}}, resources)
	groupVersionResources, err := discovery.GroupVersionResources(deletableResources)
	if err != nil {
		errs = append(errs, err)
	}
	for gvr := range groupVersionResources {
		gvrEstimate, err := d.deleteAllContentForGroupVersionResource(gvr, namespace, namespaceDeletedAt)
		if err != nil {
			errs = append(errs, err)
		}
		if gvrEstimate > estimate {
			estimate = gvrEstimate
		}
	}
	if len(errs) > 0 {
		return estimate, utilerrors.NewAggregate(errs)
	}
	klog.V(4).Infof("namespace controller - deleteAllContent - namespace: %s, estimate: %v", namespace, estimate)
	return estimate, nil
}
func (d *namespacedResourcesDeleter) estimateGracefulTermination(gvr schema.GroupVersionResource, ns string, namespaceDeletedAt metav1.Time) (int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	groupResource := gvr.GroupResource()
	klog.V(5).Infof("namespace controller - estimateGracefulTermination - group %s, resource: %s", groupResource.Group, groupResource.Resource)
	estimate := int64(0)
	var err error
	switch groupResource {
	case schema.GroupResource{Group: "", Resource: "pods"}:
		estimate, err = d.estimateGracefulTerminationForPods(ns)
	}
	if err != nil {
		return estimate, err
	}
	duration := time.Since(namespaceDeletedAt.Time)
	allowedEstimate := time.Duration(estimate) * time.Second
	if duration >= allowedEstimate {
		estimate = int64(0)
	}
	return estimate, nil
}
func (d *namespacedResourcesDeleter) estimateGracefulTerminationForPods(ns string) (int64, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(5).Infof("namespace controller - estimateGracefulTerminationForPods - namespace %s", ns)
	estimate := int64(0)
	podsGetter := d.podsGetter
	if podsGetter == nil || reflect.ValueOf(podsGetter).IsNil() {
		return estimate, fmt.Errorf("unexpected: podsGetter is nil. Cannot estimate grace period seconds for pods")
	}
	items, err := podsGetter.Pods(ns).List(metav1.ListOptions{IncludeUninitialized: true})
	if err != nil {
		return estimate, err
	}
	for i := range items.Items {
		pod := items.Items[i]
		phase := pod.Status.Phase
		if v1.PodSucceeded == phase || v1.PodFailed == phase {
			continue
		}
		if pod.Spec.TerminationGracePeriodSeconds != nil {
			grace := *pod.Spec.TerminationGracePeriodSeconds
			if grace > estimate {
				estimate = grace
			}
		}
	}
	return estimate, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
