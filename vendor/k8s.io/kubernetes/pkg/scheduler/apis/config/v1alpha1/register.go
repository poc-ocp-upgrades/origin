package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeschedulerconfigv1alpha1 "k8s.io/kube-scheduler/config/v1alpha1"
)

const GroupName = "kubescheduler.config.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
var (
	localSchemeBuilder = &kubeschedulerconfigv1alpha1.SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(addDefaultingFuncs)
}
