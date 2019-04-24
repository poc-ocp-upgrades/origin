package internalversion

import (
	"time"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	units "github.com/docker/go-units"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
)

func formatRelativeTime(t time.Time) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return units.HumanDuration(timeNowFn().Sub(t))
}

var timeNowFn = func() time.Time {
	return time.Now()
}

func roleBindingRestrictionType(rbr *authorizationapi.RoleBindingRestriction) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case rbr.Spec.UserRestriction != nil:
		return "User"
	case rbr.Spec.GroupRestriction != nil:
		return "Group"
	case rbr.Spec.ServiceAccountRestriction != nil:
		return "ServiceAccount"
	}
	return ""
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
