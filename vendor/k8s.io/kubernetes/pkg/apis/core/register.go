package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = ""

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

func Kind(kind string) schema.GroupKind {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := scheme.AddIgnoredConversionType(&metav1.TypeMeta{}, &metav1.TypeMeta{}); err != nil {
		return err
	}
	scheme.AddKnownTypes(SchemeGroupVersion, &Pod{}, &PodList{}, &PodStatusResult{}, &PodTemplate{}, &PodTemplateList{}, &ReplicationControllerList{}, &ReplicationController{}, &ServiceList{}, &Service{}, &ServiceProxyOptions{}, &NodeList{}, &Node{}, &NodeProxyOptions{}, &Endpoints{}, &EndpointsList{}, &Binding{}, &Event{}, &EventList{}, &List{}, &LimitRange{}, &LimitRangeList{}, &ResourceQuota{}, &ResourceQuotaList{}, &Namespace{}, &NamespaceList{}, &ServiceAccount{}, &ServiceAccountList{}, &Secret{}, &SecretList{}, &PersistentVolume{}, &PersistentVolumeList{}, &PersistentVolumeClaim{}, &PersistentVolumeClaimList{}, &PodAttachOptions{}, &PodLogOptions{}, &PodExecOptions{}, &PodPortForwardOptions{}, &PodProxyOptions{}, &ComponentStatus{}, &ComponentStatusList{}, &SerializedReference{}, &RangeAllocation{}, &ConfigMap{}, &ConfigMapList{})
	return nil
}
