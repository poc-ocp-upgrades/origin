package rest

import (
	goformat "fmt"
	extensionsapiv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/extensions"
	daemonstore "k8s.io/kubernetes/pkg/registry/apps/daemonset/storage"
	deploymentstore "k8s.io/kubernetes/pkg/registry/apps/deployment/storage"
	replicasetstore "k8s.io/kubernetes/pkg/registry/apps/replicaset/storage"
	expcontrollerstore "k8s.io/kubernetes/pkg/registry/extensions/controller/storage"
	ingressstore "k8s.io/kubernetes/pkg/registry/extensions/ingress/storage"
	networkpolicystore "k8s.io/kubernetes/pkg/registry/networking/networkpolicy/storage"
	pspstore "k8s.io/kubernetes/pkg/registry/policy/podsecuritypolicy/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(extensions.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	if apiResourceConfigSource.VersionEnabled(extensionsapiv1beta1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[extensionsapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
	}
	return apiGroupInfo, true
}
func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storage := map[string]rest.Storage{}
	controllerStorage := expcontrollerstore.NewStorage(restOptionsGetter)
	storage["replicationcontrollers"] = controllerStorage.ReplicationController
	storage["replicationcontrollers/scale"] = controllerStorage.Scale
	daemonSetStorage, daemonSetStatusStorage := daemonstore.NewREST(restOptionsGetter)
	storage["daemonsets"] = daemonSetStorage.WithCategories(nil)
	storage["daemonsets/status"] = daemonSetStatusStorage
	deploymentStorage := deploymentstore.NewStorage(restOptionsGetter)
	storage["deployments"] = deploymentStorage.Deployment.WithCategories(nil)
	storage["deployments/status"] = deploymentStorage.Status
	storage["deployments/rollback"] = deploymentStorage.Rollback
	storage["deployments/scale"] = deploymentStorage.Scale
	ingressStorage, ingressStatusStorage := ingressstore.NewREST(restOptionsGetter)
	storage["ingresses"] = ingressStorage
	storage["ingresses/status"] = ingressStatusStorage
	podSecurityPolicyStorage := pspstore.NewREST(restOptionsGetter)
	storage["podSecurityPolicies"] = podSecurityPolicyStorage
	replicaSetStorage := replicasetstore.NewStorage(restOptionsGetter)
	storage["replicasets"] = replicaSetStorage.ReplicaSet.WithCategories(nil)
	storage["replicasets/status"] = replicaSetStorage.Status
	storage["replicasets/scale"] = replicaSetStorage.Scale
	networkExtensionsStorage := networkpolicystore.NewREST(restOptionsGetter)
	storage["networkpolicies"] = networkExtensionsStorage
	return storage
}
func (p RESTStorageProvider) GroupName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return extensions.GroupName
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
