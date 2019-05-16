package api

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	OpenShiftDisplayName                   = "openshift.io/display-name"
	OpenShiftProviderDisplayNameAnnotation = "openshift.io/provider-display-name"
	OpenShiftDocumentationURLAnnotation    = "openshift.io/documentation-url"
	OpenShiftSupportURLAnnotation          = "openshift.io/support-url"
	OpenShiftDescription                   = "openshift.io/description"
	OpenShiftLongDescriptionAnnotation     = "openshift.io/long-description"
)

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
