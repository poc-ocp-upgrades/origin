package trigger

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
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

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
