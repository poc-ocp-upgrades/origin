package rest

import (
	goformat "fmt"
	appsapiv1 "k8s.io/api/apps/v1"
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
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apps.GroupName
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
