package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeproxyconfigv1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
)

const GroupName = "kubeproxy.config.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
var (
	localSchemeBuilder = &kubeproxyconfigv1alpha1.SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	localSchemeBuilder.Register(addDefaultingFuncs)
}
