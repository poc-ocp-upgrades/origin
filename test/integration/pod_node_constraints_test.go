package integration

import (
	"testing"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	"k8s.io/client-go/kubernetes"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	oappsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	policy "github.com/openshift/origin/pkg/oc/cli/admin/policy"
	pluginapi "github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints"
	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestPodNodeConstraintsAdmissionPluginSetNodeNameClusterAdmin(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oclient, kclientset, fn := setupClusterAdminPodNodeConstraintsTest(t, &pluginapi.PodNodeConstraintsConfig{})
	defer fn()
	testPodNodeConstraintsObjectCreationWithPodTemplate(t, "set node name, cluster admin", kclientset, oclient, "nodename.example.com", nil, false)
}
func TestPodNodeConstraintsAdmissionPluginSetNodeNameNonAdmin(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &pluginapi.PodNodeConstraintsConfig{}
	oclient, kclientset, fn := setupUserPodNodeConstraintsTest(t, config, "derples")
	defer fn()
	testPodNodeConstraintsObjectCreationWithPodTemplate(t, "set node name, regular user", kclientset, oclient, "nodename.example.com", nil, true)
}
func TestPodNodeConstraintsAdmissionPluginSetNodeSelectorClusterAdmin(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &pluginapi.PodNodeConstraintsConfig{NodeSelectorLabelBlacklist: []string{"hostname"}}
	oclient, kclientset, fn := setupClusterAdminPodNodeConstraintsTest(t, config)
	defer fn()
	testPodNodeConstraintsObjectCreationWithPodTemplate(t, "set node selector, cluster admin", kclientset, oclient, "", map[string]string{"hostname": "foo"}, false)
}
func TestPodNodeConstraintsAdmissionPluginSetNodeSelectorNonAdmin(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &pluginapi.PodNodeConstraintsConfig{NodeSelectorLabelBlacklist: []string{"hostname"}}
	oclient, kclientset, fn := setupUserPodNodeConstraintsTest(t, config, "derples")
	defer fn()
	testPodNodeConstraintsObjectCreationWithPodTemplate(t, "set node selector, regular user", kclientset, oclient, "", map[string]string{"hostname": "foo"}, true)
}
func setupClusterAdminPodNodeConstraintsTest(t *testing.T, pluginConfig *pluginapi.PodNodeConstraintsConfig) (appsclient.Interface, kubernetes.Interface, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("error creating config: %v", err)
	}
	cfg := map[string]*configapi.AdmissionPluginConfig{"scheduling.openshift.io/PodNodeConstraints": {Configuration: pluginConfig}}
	masterConfig.AdmissionConfig.PluginConfig = cfg
	kubeConfigFile, err := testserver.StartConfiguredMaster(masterConfig)
	if err != nil {
		t.Fatalf("error starting server: %v", err)
	}
	kubeClientset, err := testutil.GetClusterAdminKubeClient(kubeConfigFile)
	if err != nil {
		t.Fatalf("error getting client: %v", err)
	}
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(kubeConfigFile)
	if err != nil {
		t.Fatalf("error getting client: %v", err)
	}
	ns := &corev1.Namespace{}
	ns.Name = testutil.Namespace()
	_, err = kubeClientset.CoreV1().Namespaces().Create(ns)
	if err != nil {
		t.Fatalf("error creating namespace: %v", err)
	}
	if err := testserver.WaitForPodCreationServiceAccounts(kubeClientset, testutil.Namespace()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return appsclient.NewForConfigOrDie(clusterAdminClientConfig), kubeClientset, func() {
		testserver.CleanupMasterEtcd(t, masterConfig)
	}
}
func setupUserPodNodeConstraintsTest(t *testing.T, pluginConfig *pluginapi.PodNodeConstraintsConfig, user string) (appsclient.Interface, kubernetes.Interface, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("error creating config: %v", err)
	}
	cfg := map[string]*configapi.AdmissionPluginConfig{"scheduling.openshift.io/PodNodeConstraints": {Configuration: pluginConfig}}
	masterConfig.AdmissionConfig.PluginConfig = cfg
	kubeConfigFile, err := testserver.StartConfiguredMaster(masterConfig)
	if err != nil {
		t.Fatalf("error starting server: %v", err)
	}
	clusterAdminClientConfig, err := testutil.GetClusterAdminClientConfig(kubeConfigFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	userkubeClientset, userClientConfig, err := testutil.GetClientForUser(clusterAdminClientConfig, user)
	if err != nil {
		t.Fatalf("error getting user/kube client: %v", err)
	}
	kubeClientset, err := testutil.GetClusterAdminKubeClient(kubeConfigFile)
	if err != nil {
		t.Fatalf("error getting kube client: %v", err)
	}
	ns := &corev1.Namespace{}
	ns.Name = testutil.Namespace()
	_, err = kubeClientset.CoreV1().Namespaces().Create(ns)
	if err != nil {
		t.Fatalf("error creating namespace: %v", err)
	}
	if err := testserver.WaitForServiceAccounts(kubeClientset, testutil.Namespace(), []string{bootstrappolicy.DefaultServiceAccountName}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	addUser := &policy.RoleModificationOptions{RoleName: bootstrappolicy.AdminRoleName, RoleKind: "ClusterRole", RbacClient: rbacv1client.NewForConfigOrDie(clusterAdminClientConfig), Users: []string{user}, PrintFlags: genericclioptions.NewPrintFlags(""), ToPrinter: func(string) (printers.ResourcePrinter, error) {
		return printers.NewDiscardingPrinter(), nil
	}}
	if err := addUser.AddRole(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return appsclient.NewForConfigOrDie(userClientConfig), userkubeClientset, func() {
		testserver.CleanupMasterEtcd(t, masterConfig)
	}
}
func testPodNodeConstraintsPodSpec(nodeName string, nodeSelector map[string]string) corev1.PodSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	spec := corev1.PodSpec{}
	spec.RestartPolicy = corev1.RestartPolicyAlways
	spec.NodeName = nodeName
	spec.NodeSelector = nodeSelector
	spec.Containers = []corev1.Container{{Name: "container", Image: "test/image"}}
	return spec
}
func testPodNodeConstraintsPod(nodeName string, nodeSelector map[string]string) *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &corev1.Pod{}
	pod.Name = "testpod"
	pod.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	return pod
}
func testPodNodeConstraintsReplicationController(nodeName string, nodeSelector map[string]string) *corev1.ReplicationController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc := &corev1.ReplicationController{}
	rc.Name = "testrc"
	rc.Spec.Replicas = int32Ptr(1)
	rc.Spec.Selector = map[string]string{"foo": "bar"}
	rc.Spec.Template = &corev1.PodTemplateSpec{}
	rc.Spec.Template.Labels = map[string]string{"foo": "bar"}
	rc.Spec.Template.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	return rc
}
func int32Ptr(in int32) *int32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &in
}
func testPodNodeConstraintsDeployment(nodeName string, nodeSelector map[string]string) *appsv1.Deployment {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &appsv1.Deployment{}
	d.Name = "testdeployment"
	d.Spec.Replicas = int32Ptr(1)
	d.Spec.Template.Labels = map[string]string{"foo": "bar"}
	d.Spec.Template.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	d.Spec.Selector = &metav1.LabelSelector{MatchLabels: map[string]string{"foo": "bar"}}
	return d
}
func testPodNodeConstraintsReplicaSet(nodeName string, nodeSelector map[string]string) *appsv1.ReplicaSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs := &appsv1.ReplicaSet{}
	rs.Name = "testrs"
	rs.Spec.Replicas = int32Ptr(1)
	rs.Spec.Template = corev1.PodTemplateSpec{}
	rs.Spec.Template.Labels = map[string]string{"foo": "bar"}
	rs.Spec.Template.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	rs.Spec.Selector = &metav1.LabelSelector{MatchLabels: map[string]string{"foo": "bar"}}
	return rs
}
func testPodNodeConstraintsJob(nodeName string, nodeSelector map[string]string) *batchv1.Job {
	_logClusterCodePath()
	defer _logClusterCodePath()
	job := &batchv1.Job{}
	job.Name = "testjob"
	job.Spec.Template.Labels = map[string]string{"foo": "bar"}
	job.Spec.Template.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	job.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyNever
	return job
}
func testPodNodeConstraintsDeploymentConfig(nodeName string, nodeSelector map[string]string) *oappsv1.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := &oappsv1.DeploymentConfig{}
	dc.Name = "testdc"
	dc.Spec.Replicas = 1
	dc.Spec.Template = &corev1.PodTemplateSpec{}
	dc.Spec.Template.Labels = map[string]string{"foo": "bar"}
	dc.Spec.Template.Spec = testPodNodeConstraintsPodSpec(nodeName, nodeSelector)
	dc.Spec.Selector = map[string]string{"foo": "bar"}
	return dc
}
func testPodNodeConstraintsObjectCreationWithPodTemplate(t *testing.T, name string, kclientset kubernetes.Interface, appsClient appsclient.Interface, nodeName string, nodeSelector map[string]string, expectError bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkForbiddenErr := func(objType string, err error) {
		if err == nil && expectError {
			t.Errorf("%s (%s): expected forbidden error but did not receive one", name, objType)
			return
		}
		if err != nil && !expectError {
			t.Errorf("%s (%s): got error but did not expect one: %v", name, objType, err)
			return
		}
		if err != nil && expectError && !kapierrors.IsForbidden(err) {
			t.Errorf("%s (%s): did not get an expected forbidden error: %v", name, objType, err)
			return
		}
	}
	pod := testPodNodeConstraintsPod(nodeName, nodeSelector)
	_, err := kclientset.CoreV1().Pods(testutil.Namespace()).Create(pod)
	checkForbiddenErr("pod", err)
	rc := testPodNodeConstraintsReplicationController(nodeName, nodeSelector)
	_, err = kclientset.CoreV1().ReplicationControllers(testutil.Namespace()).Create(rc)
	checkForbiddenErr("rc", err)
	rs := testPodNodeConstraintsReplicaSet(nodeName, nodeSelector)
	_, err = kclientset.AppsV1().ReplicaSets(testutil.Namespace()).Create(rs)
	checkForbiddenErr("replicaset", err)
	job := testPodNodeConstraintsJob(nodeName, nodeSelector)
	_, err = kclientset.BatchV1().Jobs(testutil.Namespace()).Create(job)
	checkForbiddenErr("job", err)
	dc := testPodNodeConstraintsDeploymentConfig(nodeName, nodeSelector)
	_, err = appsClient.AppsV1().DeploymentConfigs(testutil.Namespace()).Create(dc)
	checkForbiddenErr("dc", err)
}
