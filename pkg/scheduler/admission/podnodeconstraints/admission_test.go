package podnodeconstraints

import (
	"bytes"
	"fmt"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/batch"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/auth/nodeidentifier"
	"k8s.io/kubernetes/pkg/serviceaccount"
	oapps "github.com/openshift/api/apps"
	"github.com/openshift/api/security"
	_ "github.com/openshift/origin/pkg/api/install"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/scheduler/admission/apis/podnodeconstraints"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
)

func TestPodNodeConstraints(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := metav1.NamespaceDefault
	tests := []struct {
		config			*podnodeconstraints.PodNodeConstraintsConfig
		resource		runtime.Object
		kind			schema.GroupKind
		groupresource		schema.GroupResource
		userinfo		user.Info
		reviewResponse		*authorizationapi.SubjectAccessReviewResponse
		expectedResource	string
		expectedErrorMsg	string
	}{{config: emptyConfig(), resource: defaultPod(), userinfo: serviceaccount.UserInfo("", "", ""), reviewResponse: reviewResponse(false, ""), expectedResource: "pods/binding", expectedErrorMsg: ""}, {config: testConfig(), resource: nodeSelectorPod(), userinfo: serviceaccount.UserInfo("", "", ""), reviewResponse: reviewResponse(false, ""), expectedResource: "pods/binding", expectedErrorMsg: "node selection by label(s) [bogus] is prohibited by policy for your role"}, {config: testConfig(), resource: nodeNamePod(), userinfo: serviceaccount.UserInfo("herpy", "derpy", ""), reviewResponse: reviewResponse(false, ""), expectedResource: "pods/binding", expectedErrorMsg: "node selection by nodeName is prohibited by policy for your role"}, {config: testConfig(), resource: nodeNameNodeSelectorPod(), userinfo: serviceaccount.UserInfo("herpy", "derpy", ""), reviewResponse: reviewResponse(false, ""), expectedResource: "pods/binding", expectedErrorMsg: "node selection by nodeName and label(s) [bogus] is prohibited by policy for your role"}, {config: testConfig(), resource: nodeSelectorPod(), userinfo: serviceaccount.UserInfo("openshift-infra", "daemonset-controller", ""), reviewResponse: reviewResponse(true, ""), expectedResource: "pods/binding", expectedErrorMsg: ""}, {config: testConfig(), resource: nodeNamePod(), userinfo: serviceaccount.UserInfo("openshift-infra", "daemonset-controller", ""), reviewResponse: reviewResponse(true, ""), expectedResource: "pods/binding", expectedErrorMsg: ""}, {config: nil, resource: defaultPod(), userinfo: serviceaccount.UserInfo("", "", ""), reviewResponse: reviewResponse(false, ""), expectedResource: "pods/binding", expectedErrorMsg: ""}, {config: testConfig(), resource: nodeNameMirrorPod(), userinfo: &user.DefaultInfo{Name: "system:node:frank", Groups: []string{user.NodesGroup}}, expectedErrorMsg: ""}, {config: testConfig(), resource: nodeNamePod(), userinfo: &user.DefaultInfo{Name: "system:node:frank", Groups: []string{user.NodesGroup}}, expectedErrorMsg: "node selection by nodeName is prohibited by policy for your role"}, {config: testConfig(), resource: nodeNameMirrorPod(), userinfo: &user.DefaultInfo{Name: "system:node:bob", Groups: []string{user.NodesGroup}}, expectedErrorMsg: "node selection by nodeName is prohibited by policy for your role"}, {config: testConfig(), resource: nodeNamePod(), userinfo: &user.DefaultInfo{Name: "system:node:bob", Groups: []string{user.NodesGroup}}, expectedErrorMsg: "node selection by nodeName is prohibited by policy for your role"}}
	for i, tc := range tests {
		var expectedError error
		errPrefix := fmt.Sprintf("%d", i)
		prc := NewPodNodeConstraints(tc.config, nodeidentifier.NewDefaultNodeIdentifier())
		prc.(initializer.WantsAuthorizer).SetAuthorizer(fakeAuthorizer(t))
		err := prc.(admission.InitializationValidator).ValidateInitialization()
		if err != nil {
			checkAdmitError(t, err, expectedError, errPrefix)
			continue
		}
		attrs := admission.NewAttributesRecord(tc.resource, nil, kapi.Kind("Pod").WithVersion("version"), ns, "test", kapi.Resource("pods").WithVersion("version"), "", admission.Create, false, tc.userinfo)
		if tc.expectedErrorMsg != "" {
			expectedError = admission.NewForbidden(attrs, fmt.Errorf(tc.expectedErrorMsg))
		}
		err = prc.(admission.ValidationInterface).Validate(attrs)
		checkAdmitError(t, err, expectedError, errPrefix)
	}
}
func TestPodNodeConstraintsPodUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := metav1.NamespaceDefault
	var expectedError error
	errPrefix := "PodUpdate"
	prc := NewPodNodeConstraints(testConfig(), nodeidentifier.NewDefaultNodeIdentifier())
	prc.(initializer.WantsAuthorizer).SetAuthorizer(fakeAuthorizer(t))
	err := prc.(admission.InitializationValidator).ValidateInitialization()
	if err != nil {
		checkAdmitError(t, err, expectedError, errPrefix)
		return
	}
	attrs := admission.NewAttributesRecord(nodeNamePod(), nodeNamePod(), kapi.Kind("Pod").WithVersion("version"), ns, "test", kapi.Resource("pods").WithVersion("version"), "", admission.Update, false, serviceaccount.UserInfo("", "", ""))
	err = prc.(admission.ValidationInterface).Validate(attrs)
	checkAdmitError(t, err, expectedError, errPrefix)
}
func TestPodNodeConstraintsNonHandledResources(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := metav1.NamespaceDefault
	errPrefix := "ResourceQuotaTest"
	var expectedError error
	prc := NewPodNodeConstraints(testConfig(), nodeidentifier.NewDefaultNodeIdentifier())
	prc.(initializer.WantsAuthorizer).SetAuthorizer(fakeAuthorizer(t))
	err := prc.(admission.InitializationValidator).ValidateInitialization()
	if err != nil {
		checkAdmitError(t, err, expectedError, errPrefix)
		return
	}
	attrs := admission.NewAttributesRecord(resourceQuota(), nil, kapi.Kind("ResourceQuota").WithVersion("version"), ns, "test", kapi.Resource("resourcequotas").WithVersion("version"), "", admission.Create, false, serviceaccount.UserInfo("", "", ""))
	err = prc.(admission.ValidationInterface).Validate(attrs)
	checkAdmitError(t, err, expectedError, errPrefix)
}
func TestPodNodeConstraintsResources(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := metav1.NamespaceDefault
	testconfigs := []struct {
		config		*podnodeconstraints.PodNodeConstraintsConfig
		userinfo	user.Info
		reviewResponse	*authorizationapi.SubjectAccessReviewResponse
	}{{config: testConfig(), userinfo: serviceaccount.UserInfo("", "", ""), reviewResponse: reviewResponse(false, "")}}
	testresources := []struct {
		resource	func(bool) runtime.Object
		kind		schema.GroupKind
		groupresource	schema.GroupResource
		prefix		string
	}{{resource: replicationController, kind: kapi.Kind("ReplicationController"), groupresource: kapi.Resource("replicationcontrollers"), prefix: "ReplicationController"}, {resource: deployment, kind: extensions.Kind("Deployment"), groupresource: extensions.Resource("deployments"), prefix: "Deployment"}, {resource: replicaSet, kind: extensions.Kind("ReplicaSet"), groupresource: extensions.Resource("replicasets"), prefix: "ReplicaSet"}, {resource: job, kind: batch.Kind("Job"), groupresource: batch.Resource("jobs"), prefix: "Job"}, {resource: deploymentConfig, kind: oapps.Kind("DeploymentConfig"), groupresource: oapps.Resource("deploymentconfigs"), prefix: "DeploymentConfig"}, {resource: podTemplate, kind: kapi.Kind("PodTemplate"), groupresource: kapi.Resource("podtemplates"), prefix: "PodTemplate"}, {resource: podSecurityPolicySubjectReview, kind: security.Kind("PodSecurityPolicySubjectReview"), groupresource: security.Resource("podsecuritypolicysubjectreviews"), prefix: "PodSecurityPolicy"}, {resource: podSecurityPolicySelfSubjectReview, kind: security.Kind("PodSecurityPolicySelfSubjectReview"), groupresource: security.Resource("podsecuritypolicyselfsubjectreviews"), prefix: "PodSecurityPolicy"}, {resource: podSecurityPolicyReview, kind: security.Kind("PodSecurityPolicyReview"), groupresource: security.Resource("podsecuritypolicyreviews"), prefix: "PodSecurityPolicy"}}
	testparams := []struct {
		nodeselector		bool
		expectedErrorMsg	string
		prefix			string
	}{{nodeselector: true, expectedErrorMsg: "node selection by label(s) [bogus] is prohibited by policy for your role", prefix: "with nodeSelector"}, {nodeselector: false, expectedErrorMsg: "", prefix: "without nodeSelector"}}
	testops := []struct{ operation admission.Operation }{{operation: admission.Create}, {operation: admission.Update}}
	for _, tc := range testconfigs {
		for _, tr := range testresources {
			for _, tp := range testparams {
				for _, top := range testops {
					var expectedError error
					errPrefix := fmt.Sprintf("%s; %s; %s", tr.prefix, tp.prefix, top.operation)
					prc := NewPodNodeConstraints(tc.config, nodeidentifier.NewDefaultNodeIdentifier())
					prc.(initializer.WantsAuthorizer).SetAuthorizer(fakeAuthorizer(t))
					err := prc.(admission.InitializationValidator).ValidateInitialization()
					if err != nil {
						checkAdmitError(t, err, expectedError, errPrefix)
						continue
					}
					attrs := admission.NewAttributesRecord(tr.resource(tp.nodeselector), nil, tr.kind.WithVersion("version"), ns, "test", tr.groupresource.WithVersion("version"), "", top.operation, false, tc.userinfo)
					if tp.expectedErrorMsg != "" {
						expectedError = admission.NewForbidden(attrs, fmt.Errorf(tp.expectedErrorMsg))
					}
					err = prc.(admission.ValidationInterface).Validate(attrs)
					checkAdmitError(t, err, expectedError, errPrefix)
				}
			}
		}
	}
}
func emptyConfig() *podnodeconstraints.PodNodeConstraintsConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podnodeconstraints.PodNodeConstraintsConfig{}
}
func testConfig() *podnodeconstraints.PodNodeConstraintsConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podnodeconstraints.PodNodeConstraintsConfig{NodeSelectorLabelBlacklist: []string{"bogus"}}
}
func defaultPod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	return pod
}
func pod(ns bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	if ns {
		pod.Spec.NodeSelector = map[string]string{"bogus": "frank"}
	}
	return pod
}
func nodeNameNodeSelectorPod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	pod.Spec.NodeName = "frank"
	pod.Spec.NodeSelector = map[string]string{"bogus": "frank"}
	return pod
}
func nodeNamePod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	pod.Spec.NodeName = "frank"
	return pod
}
func nodeNameMirrorPod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	pod.Annotations = map[string]string{kapi.MirrorPodAnnotationKey: "true"}
	pod.Spec.NodeName = "frank"
	return pod
}
func nodeSelectorPod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	pod.Spec.NodeSelector = map[string]string{"bogus": "frank"}
	return pod
}
func emptyNodeSelectorPod() *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod := &kapi.Pod{}
	pod.Spec.NodeSelector = map[string]string{}
	return pod
}
func podSpec(setNodeSelector bool) *kapi.PodSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ps := &kapi.PodSpec{}
	if setNodeSelector {
		ps.NodeSelector = map[string]string{"bogus": "frank"}
	}
	return ps
}
func podTemplateSpec(setNodeSelector bool) *kapi.PodTemplateSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pts := &kapi.PodTemplateSpec{}
	if setNodeSelector {
		pts.Spec.NodeSelector = map[string]string{"bogus": "frank"}
	}
	return pts
}
func podTemplate(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pt := &kapi.PodTemplate{}
	pt.Template = *podTemplateSpec(setNodeSelector)
	return pt
}
func replicationController(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc := &kapi.ReplicationController{}
	rc.Spec.Template = podTemplateSpec(setNodeSelector)
	return rc
}
func deployment(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &apps.Deployment{}
	d.Spec.Template = *podTemplateSpec(setNodeSelector)
	return d
}
func replicaSet(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs := &apps.ReplicaSet{}
	rs.Spec.Template = *podTemplateSpec(setNodeSelector)
	return rs
}
func job(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	j := &batch.Job{}
	j.Spec.Template = *podTemplateSpec(setNodeSelector)
	return j
}
func resourceQuota() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rq := &kapi.ResourceQuota{}
	return rq
}
func deploymentConfig(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := &appsapi.DeploymentConfig{}
	dc.Spec.Template = podTemplateSpec(setNodeSelector)
	return dc
}
func podSecurityPolicySubjectReview(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pspsr := &securityapi.PodSecurityPolicySubjectReview{}
	pspsr.Spec.Template.Spec = *podSpec(setNodeSelector)
	return pspsr
}
func podSecurityPolicySelfSubjectReview(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pspssr := &securityapi.PodSecurityPolicySelfSubjectReview{}
	pspssr.Spec.Template.Spec = *podSpec(setNodeSelector)
	return pspssr
}
func podSecurityPolicyReview(setNodeSelector bool) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pspr := &securityapi.PodSecurityPolicyReview{}
	pspr.Spec.Template.Spec = *podSpec(setNodeSelector)
	return pspr
}
func checkAdmitError(t *testing.T, err error, expectedError error, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case expectedError == nil && err == nil:
	case expectedError != nil && err != nil && err.Error() != expectedError.Error():
		t.Errorf("%s: expected error %q, got: %q", prefix, expectedError.Error(), err.Error())
	case expectedError == nil && err != nil:
		t.Errorf("%s: expected no error, got: %q", prefix, err.Error())
	case expectedError != nil && err == nil:
		t.Errorf("%s: expected error %q, no error received", prefix, expectedError.Error())
	}
}

type fakeTestAuthorizer struct{ t *testing.T }

func fakeAuthorizer(t *testing.T) authorizer.Authorizer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &fakeTestAuthorizer{t: t}
}
func (a *fakeTestAuthorizer) Authorize(attributes authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ui := attributes.GetUser()
	if ui == nil {
		return authorizer.DecisionNoOpinion, "", fmt.Errorf("No valid UserInfo for Context")
	}
	if ui.GetName() == "system:serviceaccount:openshift-infra:daemonset-controller" {
		return authorizer.DecisionAllow, "", nil
	}
	return authorizer.DecisionNoOpinion, "", nil
}
func reviewResponse(allowed bool, msg string) *authorizationapi.SubjectAccessReviewResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authorizationapi.SubjectAccessReviewResponse{Allowed: allowed, Reason: msg}
}
func TestReadConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configStr := `apiVersion: scheduling.openshift.io/v1
kind: PodNodeConstraintsConfig
nodeSelectorLabelBlacklist:
  - bogus
  - foo
`
	buf := bytes.NewBufferString(configStr)
	config, err := readConfig(buf)
	if err != nil {
		t.Fatalf("unexpected error reading config: %v", err)
	}
	if len(config.NodeSelectorLabelBlacklist) == 0 {
		t.Fatalf("NodeSelectorLabelBlacklist didn't take specified value")
	}
}
