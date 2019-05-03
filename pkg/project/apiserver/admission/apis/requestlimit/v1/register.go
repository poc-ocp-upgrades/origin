package v1

import (
	"github.com/openshift/origin/pkg/project/apiserver/admission/apis/requestlimit"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "project.openshift.io"
)

var (
	GroupVersion  = schema.GroupVersion{Group: GroupName, Version: "v1"}
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, requestlimit.Install)
	Install       = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(GroupVersion, &ProjectRequestLimitConfig{})
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
