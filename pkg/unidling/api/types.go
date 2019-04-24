package api

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	IdledAtAnnotation	= "idling.alpha.openshift.io/idled-at"
	UnidleTargetAnnotation	= "idling.alpha.openshift.io/unidle-targets"
	PreviousScaleAnnotation	= "idling.alpha.openshift.io/previous-scale"
	NeedPodsReason		= "NeedPods"
)

type RecordedScaleReference struct {
	CrossGroupObjectReference	`json:",inline" protobuf:"bytes,1,opt,name=crossVersionObjectReference"`
	Replicas			int32	`json:"replicas" protobuf:"varint,2,opt,name=replicas"`
}
type CrossGroupObjectReference struct {
	Kind		string	`json:"kind" protobuf:"bytes,1,opt,name=kind"`
	Name		string	`json:"name" protobuf:"bytes,2,opt,name=name"`
	APIVersion	string	`json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
	Group		string	`json:"group,omitempty" protobuf:"bytes,3,opt,name=group"`
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
