package template

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
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

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
