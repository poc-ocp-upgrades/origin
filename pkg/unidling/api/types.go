package api

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	IdledAtAnnotation       = "idling.alpha.openshift.io/idled-at"
	UnidleTargetAnnotation  = "idling.alpha.openshift.io/unidle-targets"
	PreviousScaleAnnotation = "idling.alpha.openshift.io/previous-scale"
	NeedPodsReason          = "NeedPods"
)

type RecordedScaleReference struct {
	CrossGroupObjectReference `json:",inline" protobuf:"bytes,1,opt,name=crossVersionObjectReference"`
	Replicas                  int32 `json:"replicas" protobuf:"varint,2,opt,name=replicas"`
}
type CrossGroupObjectReference struct {
	Kind       string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	Name       string `json:"name" protobuf:"bytes,2,opt,name=name"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
	Group      string `json:"group,omitempty" protobuf:"bytes,3,opt,name=group"`
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
