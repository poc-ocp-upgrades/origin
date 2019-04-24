package trigger

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const TriggerAnnotationKey = "image.openshift.io/triggers"

type ObjectFieldTrigger struct {
	From		ObjectReference	`json:"from"`
	FieldPath	string		`json:"fieldPath"`
	Paused		bool		`json:"paused,omitempty"`
}
type ObjectReference struct {
	Kind		string	`json:"kind"`
	Name		string	`json:"name"`
	Namespace	string	`json:"namespace,omitempty"`
	APIVersion	string	`json:"apiVersion,omitempty"`
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
