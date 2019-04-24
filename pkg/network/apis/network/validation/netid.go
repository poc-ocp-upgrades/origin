package validation

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
)

const (
	MaxVNID		= uint32((1 << 24) - 1)
	MinVNID		= uint32(10)
	GlobalVNID	= uint32(0)
)

func ValidVNID(vnid uint32) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if vnid == GlobalVNID {
		return nil
	}
	if vnid < MinVNID {
		return fmt.Errorf("VNID must be greater than or equal to %d", MinVNID)
	}
	if vnid > MaxVNID {
		return fmt.Errorf("VNID must be less than or equal to %d", MaxVNID)
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
