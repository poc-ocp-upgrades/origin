package template

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	IconClassAnnotation		= "iconClass"
	ServiceBrokerRoot		= "/brokers/template.openshift.io"
	ServiceMetadataIconClass	= "console.openshift.io/iconClass"
	TemplateUIDIndex		= "templateuid"
	ExposeAnnotationPrefix		= "template.openshift.io/expose-"
	Base64ExposeAnnotationPrefix	= "template.openshift.io/base64-expose-"
	WaitForReadyAnnotation		= "template.alpha.openshift.io/wait-for-ready"
	BindableAnnotation		= "template.openshift.io/bindable"
	TemplateInstanceFinalizer	= "template.openshift.io/finalizer"
	TemplateInstanceOwner		= "template.openshift.io/template-instance-owner"
)

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
