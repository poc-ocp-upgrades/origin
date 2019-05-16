package testing

import (
	goformat "fmt"
	securityv1 "github.com/openshift/api/security/v1"
	allocator "github.com/openshift/origin/pkg/security"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func CreateSAForTest() *corev1.ServiceAccount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "default"}}
}
func CreateNamespaceForTest() *corev1.Namespace {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default", Annotations: map[string]string{allocator.UIDRangeAnnotation: "1/3", allocator.MCSAnnotation: "s0:c1,c0", allocator.SupplementalGroupsAnnotation: "2/3"}}}
}
func UserScc(user string) *securityv1.SecurityContextConstraints {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var uid int64 = 9999
	fsGroup := int64(1)
	return &securityv1.SecurityContextConstraints{ObjectMeta: metav1.ObjectMeta{SelfLink: "/api/version/securitycontextconstraints/" + user, Name: user}, Users: []string{user}, SELinuxContext: securityv1.SELinuxContextStrategyOptions{Type: securityv1.SELinuxStrategyRunAsAny}, RunAsUser: securityv1.RunAsUserStrategyOptions{Type: securityv1.RunAsUserStrategyMustRunAs, UID: &uid}, FSGroup: securityv1.FSGroupStrategyOptions{Type: securityv1.FSGroupStrategyMustRunAs, Ranges: []securityv1.IDRange{{Min: fsGroup, Max: fsGroup}}}, SupplementalGroups: securityv1.SupplementalGroupsStrategyOptions{Type: securityv1.SupplementalGroupsStrategyRunAsAny}}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
