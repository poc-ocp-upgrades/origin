package integration

import (
	"strings"
	"testing"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	kapierror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	"github.com/openshift/origin/pkg/security/apis/security"
	securityclient "github.com/openshift/origin/pkg/security/generated/internalclientset/typed/security/internalversion"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestPodUpdateSCCEnforcement(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	clusterAdminKubeClientset, err := testutil.GetClusterAdminKubeClient(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	projectName := "hammer-project"
	if _, _, err := testserver.CreateNewProject(clusterAdminClientConfig, projectName, "harold"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	haroldKubeClient, haroldClientConfig, err := testutil.GetClientForUser(clusterAdminClientConfig, "harold")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := testserver.WaitForServiceAccounts(clusterAdminKubeClientset, projectName, []string{"default"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	privilegedPod := getPrivilegedPod("unsafe")
	if _, err := haroldKubeClient.CoreV1().Pods(projectName).Create(privilegedPod); !isForbiddenBySCC(err) {
		t.Fatalf("missing forbidden: %v", err)
	}
	actualPod, err := clusterAdminKubeClientset.CoreV1().Pods(projectName).Create(privilegedPod)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actualPod.Spec.Containers[0].Image = "something-nefarious"
	if _, err := haroldKubeClient.CoreV1().Pods(projectName).Update(actualPod); !isForbiddenBySCC(err) {
		t.Fatalf("missing forbidden: %v", err)
	}
	haroldCorev1Rest := corev1client.NewForConfigOrDie(haroldClientConfig).RESTClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := &metav1.Status{}
	err = haroldCorev1Rest.Post().Resource("pods").Namespace(projectName).Name(actualPod.Name).SubResource("exec").Param("container", "first").Do().Into(result)
	if !isForbiddenBySCCExecRestrictions(err) {
		t.Fatalf("missing forbidden by SCCExecRestrictions: %v", err)
	}
	actualPod.Spec.HostPID = false
	if _, err := haroldKubeClient.CoreV1().Pods(projectName).Update(actualPod); err == nil {
		t.Fatalf("missing error: %v", err)
	}
}
func TestAllowedSCCViaRBAC(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatal(err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	clusterAdminKubeClientset, err := testutil.GetClusterAdminKubeClient(clusterAdminKubeConfig)
	if err != nil {
		t.Fatal(err)
	}
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(clusterAdminKubeConfig)
	if err != nil {
		t.Fatal(err)
	}
	project1 := "project1"
	project2 := "project2"
	user1 := "user1"
	user2 := "user2"
	clusterRole := "all-scc"
	rule := rbacv1helpers.NewRule("use").Groups("security.openshift.io").Resources("securitycontextconstraints").RuleOrDie()
	if _, err := clusterAdminKubeClientset.RbacV1().ClusterRoles().Create(&rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: clusterRole}, Rules: []rbacv1.PolicyRule{rule}}); err != nil {
		t.Fatal(err)
	}
	user1Client, user1Config, err := testserver.CreateNewProject(clusterAdminClientConfig, project1, user1)
	if err != nil {
		t.Fatal(err)
	}
	user1SecurityClient := securityclient.NewForConfigOrDie(user1Config)
	user2Client, user2Config, err := testserver.CreateNewProject(clusterAdminClientConfig, project2, user2)
	if err != nil {
		t.Fatal(err)
	}
	user2SecurityClient := securityclient.NewForConfigOrDie(user2Config)
	if err := testserver.WaitForServiceAccounts(user1Client, project1, []string{"default"}); err != nil {
		t.Fatal(err)
	}
	if err := testserver.WaitForServiceAccounts(user2Client, project2, []string{"default"}); err != nil {
		t.Fatal(err)
	}
	if _, err := user1Client.CoreV1().Pods(project1).Create(getPrivilegedPod("test1")); !isForbiddenBySCC(err) {
		t.Fatalf("missing forbidden for user1: %v", err)
	}
	if _, err := user2Client.CoreV1().Pods(project2).Create(getPrivilegedPod("test2")); !isForbiddenBySCC(err) {
		t.Fatalf("missing forbidden for user2: %v", err)
	}
	rb := rbacv1helpers.NewRoleBindingForClusterRole(clusterRole, project1).Users(user1).BindingOrDie()
	if _, err := clusterAdminKubeClientset.RbacV1().RoleBindings(project1).Create(&rb); err != nil {
		t.Fatal(err)
	}
	rbEditUser1Project2 := rbacv1helpers.NewRoleBindingForClusterRole("edit", project2).Users(user1).BindingOrDie()
	if _, err := clusterAdminKubeClientset.RbacV1().RoleBindings(project2).Create(&rbEditUser1Project2); err != nil {
		t.Fatal(err)
	}
	rbEditUser2Project1 := rbacv1helpers.NewRoleBindingForClusterRole("edit", project1).Users(user2).BindingOrDie()
	if _, err := clusterAdminKubeClientset.RbacV1().RoleBindings(project1).Create(&rbEditUser2Project1); err != nil {
		t.Fatal(err)
	}
	crb := rbacv1helpers.NewClusterBinding(clusterRole).Users(user2).BindingOrDie()
	if _, err := clusterAdminKubeClientset.RbacV1().ClusterRoleBindings().Create(&crb); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(user1Client.AuthorizationV1(), project1, rule.Verbs[0], schema.GroupResource{Group: rule.APIGroups[0], Resource: rule.Resources[0]}, true); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(user1Client.AuthorizationV1(), project2, "create", schema.GroupResource{Resource: "pods"}, true); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForPolicyUpdate(user2Client.AuthorizationV1(), project1, "create", schema.GroupResource{Resource: "pods"}, true); err != nil {
		t.Fatal(err)
	}
	if err := testutil.WaitForClusterPolicyUpdate(user2Client.AuthorizationV1(), rule.Verbs[0], schema.GroupResource{Group: rule.APIGroups[0], Resource: rule.Resources[0]}, true); err != nil {
		t.Fatal(err)
	}
	if _, err := user1Client.CoreV1().Pods(project1).Create(getPrivilegedPod("test3")); err != nil {
		t.Fatalf("user1 failed to create pod in project1 via local binding: %v", err)
	}
	if _, err := user1Client.CoreV1().Pods(project2).Create(getPrivilegedPod("test4")); !isForbiddenBySCC(err) {
		t.Fatalf("missing forbidden for user1 in project2: %v", err)
	}
	if _, err := user2Client.CoreV1().Pods(project1).Create(getPrivilegedPod("test5")); err != nil {
		t.Fatalf("user2 failed to create pod in project1 via cluster binding: %v", err)
	}
	if _, err := user2Client.CoreV1().Pods(project2).Create(getPrivilegedPod("test6")); err != nil {
		t.Fatalf("user2 failed to create pod in project2 via cluster binding: %v", err)
	}
	user1PSPReview, err := user1SecurityClient.PodSecurityPolicySelfSubjectReviews(project1).Create(runAsRootPSPSSR())
	if err != nil {
		t.Fatal(err)
	}
	if allowedBy := user1PSPReview.Status.AllowedBy; allowedBy == nil || allowedBy.Name != "anyuid" {
		t.Fatalf("user1 failed PSP SSR in project1: %v", allowedBy)
	}
	user2PSPReview, err := user2SecurityClient.PodSecurityPolicySelfSubjectReviews(project2).Create(runAsRootPSPSSR())
	if err != nil {
		t.Fatal(err)
	}
	if allowedBy := user2PSPReview.Status.AllowedBy; allowedBy == nil || allowedBy.Name != "anyuid" {
		t.Fatalf("user2 failed PSP SSR in project2: %v", allowedBy)
	}
}
func isForbiddenBySCC(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kapierror.IsForbidden(err) && strings.Contains(err.Error(), "unable to validate against any security context constraint")
}
func isForbiddenBySCCExecRestrictions(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kapierror.IsForbidden(err) && strings.Contains(err.Error(), "pod's security context exceeds your permissions")
}
func getPrivilegedPod(name string) *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: corev1.PodSpec{Containers: []corev1.Container{{Name: "first", Image: "something-innocuous"}}, HostPID: true}}
}
func runAsRootPSPSSR() *security.PodSecurityPolicySelfSubjectReview {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &security.PodSecurityPolicySelfSubjectReview{Spec: security.PodSecurityPolicySelfSubjectReviewSpec{Template: kapi.PodTemplateSpec{Spec: kapi.PodSpec{Containers: []kapi.Container{{Name: "fake", Image: "fake", SecurityContext: &kapi.SecurityContext{RunAsUser: new(int64)}}}}}}}
}
