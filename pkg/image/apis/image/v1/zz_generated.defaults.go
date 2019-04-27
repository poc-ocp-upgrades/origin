package v1

import (
	v1 "github.com/openshift/api/image/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.ImageStream{}, func(obj interface{}) {
		SetObjectDefaults_ImageStream(obj.(*v1.ImageStream))
	})
	scheme.AddTypeDefaultingFunc(&v1.ImageStreamImport{}, func(obj interface{}) {
		SetObjectDefaults_ImageStreamImport(obj.(*v1.ImageStreamImport))
	})
	scheme.AddTypeDefaultingFunc(&v1.ImageStreamList{}, func(obj interface{}) {
		SetObjectDefaults_ImageStreamList(obj.(*v1.ImageStreamList))
	})
	scheme.AddTypeDefaultingFunc(&v1.ImageStreamTag{}, func(obj interface{}) {
		SetObjectDefaults_ImageStreamTag(obj.(*v1.ImageStreamTag))
	})
	scheme.AddTypeDefaultingFunc(&v1.ImageStreamTagList{}, func(obj interface{}) {
		SetObjectDefaults_ImageStreamTagList(obj.(*v1.ImageStreamTagList))
	})
	return nil
}
func SetObjectDefaults_ImageStream(in *v1.ImageStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Spec.Tags {
		a := &in.Spec.Tags[i]
		SetDefaults_TagReferencePolicy(&a.ReferencePolicy)
	}
}
func SetObjectDefaults_ImageStreamImport(in *v1.ImageStreamImport) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Spec.Repository != nil {
		SetDefaults_TagReferencePolicy(&in.Spec.Repository.ReferencePolicy)
	}
	for i := range in.Spec.Images {
		a := &in.Spec.Images[i]
		SetDefaults_ImageImportSpec(a)
		SetDefaults_TagReferencePolicy(&a.ReferencePolicy)
	}
	if in.Status.Import != nil {
		SetObjectDefaults_ImageStream(in.Status.Import)
	}
}
func SetObjectDefaults_ImageStreamList(in *v1.ImageStreamList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ImageStream(a)
	}
}
func SetObjectDefaults_ImageStreamTag(in *v1.ImageStreamTag) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.Tag != nil {
		SetDefaults_TagReferencePolicy(&in.Tag.ReferencePolicy)
	}
}
func SetObjectDefaults_ImageStreamTagList(in *v1.ImageStreamTagList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ImageStreamTag(a)
	}
}
