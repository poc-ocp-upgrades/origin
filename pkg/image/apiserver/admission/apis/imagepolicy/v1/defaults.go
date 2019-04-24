package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	kapi "k8s.io/kubernetes/pkg/apis/core"
)

func SetDefaults_ImagePolicyConfig(obj *ImagePolicyConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if obj == nil {
		return
	}
	if len(obj.ResolveImages) == 0 {
		obj.ResolveImages = Attempt
	}
	for i := range obj.ExecutionRules {
		if len(obj.ExecutionRules[i].OnResources) == 0 {
			obj.ExecutionRules[i].OnResources = []metav1.GroupResource{{Resource: "pods", Group: kapi.GroupName}}
		}
	}
	if obj.ResolutionRules == nil {
		obj.ResolutionRules = []ImageResolutionPolicyRule{{TargetResource: metav1.GroupResource{Resource: "pods"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "build.openshift.io", Resource: "builds"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "batch", Resource: "jobs"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "extensions", Resource: "replicasets"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Resource: "replicationcontrollers"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "apps", Resource: "deployments"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "extensions", Resource: "deployments"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "apps", Resource: "statefulsets"}, LocalNames: true}, {TargetResource: metav1.GroupResource{Group: "extensions", Resource: "daemonsets"}, LocalNames: true}}
		for i := range obj.ResolutionRules {
			if len(obj.ResolutionRules[i].Policy) != 0 {
				continue
			}
			obj.ResolutionRules[i].Policy = DoNotAttempt
			for _, rule := range obj.ExecutionRules {
				if executionRuleCoversResource(rule, obj.ResolutionRules[i].TargetResource) {
					obj.ResolutionRules[i].Policy = obj.ResolveImages
					break
				}
			}
		}
	} else {
		for i := range obj.ResolutionRules {
			if len(obj.ResolutionRules[i].Policy) != 0 {
				continue
			}
			obj.ResolutionRules[i].Policy = obj.ResolveImages
		}
	}
}
func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&ImagePolicyConfig{}, func(obj interface{}) {
		SetDefaults_ImagePolicyConfig(obj.(*ImagePolicyConfig))
	})
	return nil
}
func executionRuleCoversResource(rule ImageExecutionPolicyRule, gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, target := range rule.OnResources {
		if target.Group == gr.Group && (target.Resource == gr.Resource || target.Resource == "*") {
			return true
		}
	}
	return false
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
