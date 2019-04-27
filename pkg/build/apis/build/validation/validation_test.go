package validation

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	_ "github.com/openshift/origin/pkg/build/apis/build/install"
)

func TestBuildValidationSuccess(t *testing.T) {
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
	build := &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "buildid", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseNew}}
	if result := ValidateBuild(build); len(result) > 0 {
		t.Errorf("Unexpected validation error returned %v", result)
	}
}
func checkDockerStrategyEmptySourceError(result field.ErrorList) bool {
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
	for _, err := range result {
		if err.Type == field.ErrorTypeInvalid && strings.Contains(err.Field, "spec.source") && strings.Contains(err.Detail, "must provide a value for at least one source input(git, binary, dockerfile, images).") {
			return true
		}
	}
	return false
}
func TestBuildEmptySource(t *testing.T) {
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
	builds := []buildapi.Build{{ObjectMeta: metav1.ObjectMeta{Name: "buildid", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseNew}}, {ObjectMeta: metav1.ObjectMeta{Name: "buildid", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseNew}}}
	for _, build := range builds {
		if result := ValidateBuild(&build); len(result) > 0 {
			t.Errorf("Unexpected validation error returned %v", result)
		}
	}
	badBuild := &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "buildid", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseNew}}
	if result := ValidateBuild(badBuild); len(result) == 0 {
		t.Error("An error should have occurred with a DockerStrategy / no source combo")
	} else {
		if !checkDockerStrategyEmptySourceError(result) {
			t.Errorf("The correct error was not found: %v", result)
		}
	}
}
func TestBuildConfigEmptySource(t *testing.T) {
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
	buildConfigs := []buildapi.BuildConfig{{ObjectMeta: metav1.ObjectMeta{Name: "config-id", Namespace: "namespace"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}}, {ObjectMeta: metav1.ObjectMeta{Name: "config-id", Namespace: "namespace"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}}}
	for _, buildConfig := range buildConfigs {
		if result := ValidateBuildConfig(&buildConfig); len(result) > 0 {
			t.Errorf("Unexpected validation error returned %v", result)
		}
	}
	badBuildConfig := buildapi.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: "config-id", Namespace: "namespace"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}}
	if result := ValidateBuildConfig(&badBuildConfig); len(result) == 0 {
		t.Error("An error should have occurred with a DockerStrategy / no source combo")
	} else {
		if !checkDockerStrategyEmptySourceError(result) {
			t.Errorf("The correct error was not found: %v", result)
		}
	}
}
func TestBuildValidationFailure(t *testing.T) {
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
	build := &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "", Namespace: ""}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseNew}}
	if result := ValidateBuild(build); len(result) != 2 {
		t.Errorf("Unexpected validation result: %v", result)
	}
}
func newDefaultParameters() buildapi.BuildSpec {
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
	return buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}
}
func newNonDefaultParameters() buildapi.BuildSpec {
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
	o := newDefaultParameters()
	o.Source.Git.URI = "changed"
	return o
}
func TestValidateBuildUpdate(t *testing.T) {
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
	old := &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseRunning}}
	errs := ValidateBuildUpdate(&buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseComplete}}, old)
	if len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}
	errorCases := map[string]struct {
		Old	*buildapi.Build
		Update	*buildapi.Build
		T	field.ErrorType
		F	string
	}{"changed spec": {Old: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters()}, Update: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newNonDefaultParameters()}, T: field.ErrorTypeInvalid, F: "spec"}, "update from terminal1": {Old: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseComplete}}, Update: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseRunning}}, T: field.ErrorTypeInvalid, F: "status.phase"}, "update from terminal2": {Old: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseCancelled}}, Update: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseRunning}}, T: field.ErrorTypeInvalid, F: "status.phase"}, "update from terminal3": {Old: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseError}}, Update: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseRunning}}, T: field.ErrorTypeInvalid, F: "status.phase"}, "update from terminal4": {Old: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseFailed}}, Update: &buildapi.Build{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault, Name: "my-build", ResourceVersion: "1"}, Spec: newDefaultParameters(), Status: buildapi.BuildStatus{Phase: buildapi.BuildPhaseRunning}}, T: field.ErrorTypeInvalid, F: "status.phase"}}
	for k, v := range errorCases {
		errs := ValidateBuildUpdate(v.Update, v.Old)
		if len(errs) == 0 {
			t.Errorf("expected failure %s for %v", k, v.Update)
			continue
		}
		for i := range errs {
			if errs[i].Type != v.T {
				t.Errorf("%s: expected errors to have type %s: %v", k, v.T, errs[i])
			}
			if errs[i].Field != v.F {
				t.Errorf("%s: expected errors to have field %s: %v", k, v.F, errs[i])
			}
		}
	}
}
func TestBuildConfigDockerStrategyImageChangeTrigger(t *testing.T) {
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
	buildConfig := &buildapi.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: "config-id", Namespace: "namespace"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}, Triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}}}
	errors := ValidateBuildConfig(buildConfig)
	switch len(errors) {
	case 0:
		t.Errorf("Expected validation error, got nothing")
	case 1:
		err := errors[0]
		if err.Type != field.ErrorTypeInvalid {
			t.Errorf("Expected error to be '%v', got '%v'", field.ErrorTypeInvalid, err.Type)
		}
	default:
		t.Errorf("Expected a single validation error, got %v", errors)
	}
}
func TestBuildConfigValidationFailureRequiredName(t *testing.T) {
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
	buildConfig := &buildapi.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: "", Namespace: "foo"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}}
	errors := ValidateBuildConfig(buildConfig)
	if len(errors) != 1 {
		t.Fatalf("Unexpected validation errors %v", errors)
	}
	err := errors[0]
	if err.Type != field.ErrorTypeRequired {
		t.Errorf("Unexpected error type, expected %s, got %s", field.ErrorTypeRequired, err.Type)
	}
	if err.Field != "metadata.name" {
		t.Errorf("Unexpected field name expected metadata.name, got %s", err.Field)
	}
}
func TestBuildConfigImageChangeTriggers(t *testing.T) {
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
	tests := []struct {
		name		string
		triggers	[]buildapi.BuildTriggerPolicy
		fromKind	string
		expectError	bool
		errorType	field.ErrorType
	}{{name: "valid default trigger with imagestreamtag", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}, fromKind: "ImageStreamTag", expectError: false}, {name: "invalid default trigger (imagestreamimage)", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}, fromKind: "ImageStreamImage", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "invalid default trigger (dockerimage)", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}, fromKind: "DockerImage", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "more than one default trigger", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "missing image change struct", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeRequired}, {name: "only one default image change trigger", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag"}}}}, fromKind: "ImageStreamTag", expectError: false}, {name: "invalid reference kind for trigger", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "empty reference kind for trigger", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Name: "myimage:tag"}}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "duplicate imagestreamtag references", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag"}}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag"}}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "duplicate imagestreamtag - same as strategy ref", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "builderimage:latest"}}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}, {name: "imagestreamtag references with same name, different ns", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag", Namespace: "ns1"}}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag", Namespace: "ns2"}}}}, fromKind: "ImageStreamTag", expectError: false}, {name: "imagestreamtag references with same name, same ns", triggers: []buildapi.BuildTriggerPolicy{{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag", Namespace: "ns"}}}, {Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimage:tag", Namespace: "ns"}}}}, fromKind: "ImageStreamTag", expectError: true, errorType: field.ErrorTypeInvalid}}
	for _, tc := range tests {
		buildConfig := &buildapi.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: "bar", Namespace: "foo"}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: tc.fromKind, Name: "builderimage:latest"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}, Triggers: tc.triggers}}
		errors := ValidateBuildConfig(buildConfig)
		if hasError := len(errors) > 0; hasError != tc.expectError {
			t.Errorf("%s: did not get expected result: %#v", tc.name, errors)
		}
		if len(errors) > 0 && tc.expectError && tc.errorType != "" {
			verr := errors[0]
			if verr.Type != tc.errorType {
				t.Errorf("%s: unexpected error type. Expected: %s. Got: %s", tc.name, tc.errorType, verr.Type)
			}
		}
	}
}
func TestBuildConfigValidationOutputFailure(t *testing.T) {
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
	buildConfig := &buildapi.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: ""}, Spec: buildapi.BuildConfigSpec{RunPolicy: buildapi.BuildRunPolicySerial, CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Name: "other"}}}}}
	if result := ValidateBuildConfig(buildConfig); len(result) != 3 {
		for _, e := range result {
			t.Errorf("Unexpected validation result %v", e)
		}
	}
}
func TestValidateBuildRequest(t *testing.T) {
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
	testCases := map[string]*buildapi.BuildRequest{string(field.ErrorTypeRequired) + "metadata.namespace": {ObjectMeta: metav1.ObjectMeta{Name: "requestName"}}, string(field.ErrorTypeRequired) + "metadata.name": {ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceDefault}}}
	for desc, tc := range testCases {
		errors := ValidateBuildRequest(tc)
		if len(desc) == 0 && len(errors) > 0 {
			t.Errorf("%s: Unexpected validation result: %v", desc, errors)
		}
		if len(desc) > 0 && len(errors) != 1 {
			t.Errorf("%s: Unexpected validation result: %v", desc, errors)
		}
		if len(desc) > 0 {
			err := errors[0]
			errDesc := string(err.Type) + err.Field
			if desc != errDesc {
				t.Errorf("Unexpected validation result for %s: expected %s, got %s", err.Field, desc, errDesc)
			}
		}
	}
}
func TestValidateSource(t *testing.T) {
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
	dockerfile := "FROM something"
	invalidProxyAddress := "some!@#$%^&*()url"
	errorCases := []struct {
		t		field.ErrorType
		path		string
		source		*buildapi.BuildSource
		ok		bool
		multiple	bool
		customStrategy	bool
		dockerStrategy	bool
		jenkinsStrategy	bool
	}{{t: field.ErrorTypeRequired, path: "git.uri", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: ""}}}, {t: field.ErrorTypeInvalid, path: "git.uri", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://%"}}}, {t: field.ErrorTypeInvalid, path: "contextDir", source: &buildapi.BuildSource{Dockerfile: &dockerfile, ContextDir: "../file"}}, {t: field.ErrorTypeInvalid, path: "git", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "https://example.com/repo.git"}, Binary: &buildapi.BinaryBuildSource{}}, multiple: true}, {t: field.ErrorTypeInvalid, path: "binary.asFile", source: &buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{AsFile: "/a/path"}}}, {t: field.ErrorTypeInvalid, path: "binary.asFile", source: &buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{AsFile: "/"}}}, {t: field.ErrorTypeInvalid, path: "binary.asFile", source: &buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{AsFile: "a\\b"}}}, {source: &buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{AsFile: "/././file"}}, ok: true}, {source: &buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{AsFile: "/././file"}, Dockerfile: &dockerfile}, ok: true}, {source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "https://example.com/repo.git"}, Dockerfile: &dockerfile}, ok: true}, {source: &buildapi.BuildSource{Dockerfile: &dockerfile}, ok: true}, {source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "https://example.com/repo.git"}, ContextDir: "contextDir"}, ok: true}, {t: field.ErrorTypeRequired, path: "sourceSecret.name", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://example.com/repo.git"}, SourceSecret: &kapi.LocalObjectReference{}, ContextDir: "contextDir/../somedir"}}, {t: field.ErrorTypeInvalid, path: "git.httpproxy", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "https://example.com/repo.git", ProxyConfig: buildapi.ProxyConfig{HTTPProxy: &invalidProxyAddress}}, ContextDir: "contextDir"}}, {t: field.ErrorTypeInvalid, path: "git.httpsproxy", source: &buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "https://example.com/repo.git", ProxyConfig: buildapi.ProxyConfig{HTTPSProxy: &invalidProxyAddress}}, ContextDir: "contextDir"}}, {ok: true, source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}, {From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}}}}, {t: field.ErrorTypeRequired, path: "images[0].paths", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}}}}, multiple: true}, {t: field.ErrorTypeInvalid, path: "images[0].paths[0].destinationDir", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "/test/dir"}}}}}}, {t: field.ErrorTypeInvalid, path: "images[0].paths[0].sourcePath", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "some/path", DestinationDir: "test/dir"}}}}}}, {t: field.ErrorTypeInvalid, path: "images[0].paths[0].destinationDir", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/../../dir"}}}}}}, {t: field.ErrorTypeInvalid, path: "images[0].from.kind", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "InvalidKind", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}}}}, {t: field.ErrorTypeRequired, path: "images[0].pullSecret.name", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "DockerImage", Name: "my-image:latest"}, PullSecret: &kapi.LocalObjectReference{Name: ""}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}}}}, {t: field.ErrorTypeInvalid, path: "images[0].as[0]", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, As: []string{""}}}}}, {t: field.ErrorTypeDuplicate, path: "images[1].as[1]", source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, As: []string{"a", "b"}}, {From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:v2"}, As: []string{"c", "a"}}}}}, {ok: true, source: &buildapi.BuildSource{Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "my-image:latest"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}, As: []string{"a"}}}}}, {t: field.ErrorTypeInvalid, path: "configMaps[0].configMap", source: &buildapi.BuildSource{ConfigMaps: []buildapi.ConfigMapBuildSource{{ConfigMap: kapi.LocalObjectReference{Name: "A@ba!dn#me"}, DestinationDir: "./some/relative/path"}}}}, {t: field.ErrorTypeInvalid, path: "configMaps[0].destinationDir", source: &buildapi.BuildSource{ConfigMaps: []buildapi.ConfigMapBuildSource{{ConfigMap: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "../bad/parent/path"}}}}, {t: field.ErrorTypeInvalid, path: "configMaps[0].destinationDir", dockerStrategy: true, source: &buildapi.BuildSource{ConfigMaps: []buildapi.ConfigMapBuildSource{{ConfigMap: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "/var/log/something"}}}}, {ok: true, source: &buildapi.BuildSource{ConfigMaps: []buildapi.ConfigMapBuildSource{{ConfigMap: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "/var/log/something"}}}}, {t: field.ErrorTypeInvalid, path: "secrets[0].secret", source: &buildapi.BuildSource{Secrets: []buildapi.SecretBuildSource{{Secret: kapi.LocalObjectReference{Name: "A@ba!dn#me"}, DestinationDir: "./some/relative/path"}}}}, {t: field.ErrorTypeInvalid, path: "secrets[0].destinationDir", source: &buildapi.BuildSource{Secrets: []buildapi.SecretBuildSource{{Secret: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "../bad/parent/path"}}}}, {t: field.ErrorTypeInvalid, path: "secrets[0].destinationDir", dockerStrategy: true, source: &buildapi.BuildSource{Secrets: []buildapi.SecretBuildSource{{Secret: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "/var/log/something"}}}}, {ok: true, source: &buildapi.BuildSource{Secrets: []buildapi.SecretBuildSource{{Secret: kapi.LocalObjectReference{Name: "good-secret-name"}, DestinationDir: "/var/log/something"}}}}}
	for i, tc := range errorCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			errors := validateSource(tc.source, tc.customStrategy, tc.dockerStrategy, tc.jenkinsStrategy, nil)
			switch len(errors) {
			case 0:
				if !tc.ok {
					t.Fatalf("Unexpected validation result: %v", errors)
				}
				return
			case 1:
				if tc.ok || tc.multiple {
					t.Fatalf("Unexpected validation result: %v", errors)
				}
			default:
				if tc.ok || !tc.multiple {
					t.Fatalf("Unexpected validation result: %v", errors)
				}
			}
			err := errors[0]
			if err.Type != tc.t {
				t.Fatalf("Expected error type %s, got %s", tc.t, err.Type)
			}
			if err.Field != tc.path {
				t.Fatalf("Expected error path %s, got %s", tc.path, err.Field)
			}
		})
	}
	errorCases[11].source.ContextDir = "."
	validateSource(errorCases[11].source, false, false, false, nil)
	if len(errorCases[11].source.ContextDir) != 0 {
		t.Errorf("ContextDir was not cleaned: %s", errorCases[11].source.ContextDir)
	}
}
func TestValidateStrategy(t *testing.T) {
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
	badPolicy := buildapi.ImageOptimizationPolicy("Unknown")
	goodPolicy := buildapi.ImageOptimizationNone
	errorCases := []struct {
		t		field.ErrorType
		path		string
		strategy	*buildapi.BuildStrategy
		ok		bool
		multiple	bool
	}{{t: field.ErrorTypeInvalid, path: "", strategy: &buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{}, DockerStrategy: &buildapi.DockerBuildStrategy{}, CustomStrategy: &buildapi.CustomBuildStrategy{}, JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{}}}, {t: field.ErrorTypeInvalid, path: "dockerStrategy.imageOptimizationPolicy", strategy: &buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{ImageOptimizationPolicy: &badPolicy}}}, {ok: true, strategy: &buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{ImageOptimizationPolicy: &goodPolicy}}}}
	for i, tc := range errorCases {
		errors := validateStrategy(tc.strategy, nil)
		switch len(errors) {
		case 0:
			if !tc.ok {
				t.Errorf("%d: Unexpected validation result: %v", i, errors)
			}
			continue
		case 1:
			if tc.ok || tc.multiple {
				t.Errorf("%d: Unexpected validation result: %v", i, errors)
				continue
			}
		default:
			if tc.ok || !tc.multiple {
				t.Errorf("%d: Unexpected validation result: %v", i, errors)
				continue
			}
		}
		err := errors[0]
		if err.Type != tc.t {
			t.Errorf("%d: Unexpected error type: %s", i, err.Type)
		}
		if err.Field != tc.path {
			t.Errorf("%d: Unexpected error path: %s", i, err.Field)
		}
	}
}
func TestValidateCommonSpec(t *testing.T) {
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
	zero := int64(0)
	longString := strings.Repeat("1234567890", 100*61)
	errorCases := []struct {
		err	string
		buildapi.CommonSpec
	}{{string(field.ErrorTypeInvalid) + "output.to.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "///some/long/value/with/no/meaning"}}}}, {string(field.ErrorTypeInvalid) + "output.to.kind", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStream", Name: "///some/long/value/with/no/meaning"}}}}, {string(field.ErrorTypeInvalid) + "output.to.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "///some/long/value/with/no/meaning"}}}}, {string(field.ErrorTypeInvalid) + "output.to.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "///some/long/value/with/no/meaning:latest"}}}}, {string(field.ErrorTypeInvalid) + "output.to.kind", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "Foo", Name: "test"}}}}, {string(field.ErrorTypeRequired) + "output.to.kind", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{}}}}, {string(field.ErrorTypeRequired) + "output.to.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStreamTag"}}}}, {string(field.ErrorTypeInvalid) + "output.to.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "missingtag", Namespace: "subdomain"}}}}, {string(field.ErrorTypeInvalid) + "output.to.namespace", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "test:tag", Namespace: "not_a_valid_subdomain"}}}}, {string(field.ErrorTypeRequired) + "strategy.sourceStrategy.from.kind", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeRequired) + "strategy.sourceStrategy.from.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.sourceStrategy.from.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "bad format"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.sourceStrategy.from.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "badformat"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.sourceStrategy.from.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "bad/format:latest"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeRequired) + "strategy.customStrategy.from.kind", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.customStrategy.from.name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "bad format"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "source.dockerfile", buildapi.CommonSpec{Source: buildapi.BuildSource{Dockerfile: &longString}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}}}, {string(field.ErrorTypeInvalid) + "source.dockerfile", buildapi.CommonSpec{Source: buildapi.BuildSource{Dockerfile: &longString, Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}}}, {string(field.ErrorTypeInvalid) + "completionDeadlineSeconds", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}, CompletionDeadlineSeconds: &zero}}, {string(field.ErrorTypeInvalid) + "source", buildapi.CommonSpec{Source: buildapi.BuildSource{}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.dockerStrategy.dockerfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{DockerfilePath: "/myDockerfile"}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.dockerStrategy.dockerfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{DockerfilePath: "../someDockerfile"}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.dockerStrategy.dockerfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{DockerfilePath: "someDockerfile/../../.."}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "strategy.dockerStrategy.dockerfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{DockerfilePath: ".."}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {string(field.ErrorTypeInvalid) + "postCommit", buildapi.CommonSpec{PostCommit: buildapi.BuildPostCommitSpec{Command: []string{"rake", "test"}, Script: "rake test"}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}}}, {string(field.ErrorTypeInvalid) + "source.git", buildapi.CommonSpec{Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{}}}}, {string(field.ErrorTypeInvalid) + "source.git", buildapi.CommonSpec{Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "b"}}}}, {string(field.ErrorTypeInvalid) + "strategy.jenkinsPipelineStrategy.jenkinsfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "/myJenkinsfile"}}}}, {string(field.ErrorTypeInvalid) + "strategy.jenkinsPipelineStrategy.jenkinsfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "../someJenkinsfile"}}}}, {string(field.ErrorTypeInvalid) + "strategy.jenkinsPipelineStrategy.jenkinsfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "someJenkinsfile/../../../"}}}}, {string(field.ErrorTypeInvalid) + "strategy.jenkinsPipelineStrategy.jenkinsfilePath", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: ".."}}}}, {string(field.ErrorTypeInvalid) + "strategy.jenkinsPipelineStrategy.jenkinsfile", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{Jenkinsfile: longString + longString}}}}, {string(field.ErrorTypeRequired) + "output.imageLabels[0].name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "", Value: ""}}}}}, {string(field.ErrorTypeInvalid) + "output.imageLabels[0].name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "%$#@!", Value: ""}}}}}, {string(field.ErrorTypeInvalid) + "output.imageLabels[1].name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "really", Value: "yes"}, {Name: "really", Value: "no"}}}}}, {string(field.ErrorTypeInvalid) + "output.imageLabels[3].name", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "a", Value: "1"}, {Name: "really", Value: "yes"}, {Name: "b", Value: "2"}, {Name: "really", Value: "no"}, {Name: "c", Value: "3"}}}}}, {string(field.ErrorTypeInvalid) + "nodeSelector[A@B!]", buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, NodeSelector: map[string]string{"A@B!": "C"}}}}
	for count, config := range errorCases {
		errors := validateCommonSpec(&config.CommonSpec, nil)
		if len(errors) != 1 {
			t.Errorf("Test[%d] %s: Unexpected validation result: %v", count, config.err, errors)
			continue
		}
		err := errors[0]
		errDesc := string(err.Type) + err.Field
		if config.err != errDesc {
			t.Errorf("Test[%d] Unexpected validation result for %s: expected %s, got %s", count, err.Field, config.err, errDesc)
		}
	}
}
func TestValidateCommonSpecSuccess(t *testing.T) {
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
	shortString := "FROM foo"
	testCases := []struct{ buildapi.CommonSpec }{{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "reponame"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: "imagestreamname:tag"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "ImageStreamImage", Name: "imagestreamimage"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Dockerfile: &shortString, Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "ImageStreamImage", Name: "imagestreamimage"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "ImageStreamImage", Name: "imagestreamimage"}, DockerfilePath: "dockerfiles/firstDockerfile"}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "reponame"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "registry/project/repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "registry/project/repository/data"}}}, Output: buildapi.BuildOutput{To: &kapi.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: nil}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "key"}}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, Output: buildapi.BuildOutput{ImageLabels: []buildapi.ImageLabel{{Name: "key", Value: "value )(*&"}}}}}, {CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{}}, NodeSelector: map[string]string{"A": "B", "C": "D"}}}, {buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{URI: "http://github.com/my/repository"}}, Strategy: buildapi.BuildStrategy{JenkinsPipelineStrategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "myJenkinsfile", Env: []kapi.EnvVar{{Name: "key", Value: "value"}}}}}}}
	for count, config := range testCases {
		errors := validateCommonSpec(&config.CommonSpec, nil)
		if len(errors) != 0 {
			t.Errorf("Test[%d] Unexpected validation error: %v", count, errors)
		}
	}
}
func TestValidateDockerfilePath(t *testing.T) {
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
	tests := []struct {
		strategy		*buildapi.DockerBuildStrategy
		expectedDockerfilePath	string
	}{{strategy: &buildapi.DockerBuildStrategy{DockerfilePath: "."}, expectedDockerfilePath: ""}, {strategy: &buildapi.DockerBuildStrategy{DockerfilePath: "somedir/.."}, expectedDockerfilePath: ""}, {strategy: &buildapi.DockerBuildStrategy{DockerfilePath: "somedir/../somedockerfile"}, expectedDockerfilePath: "somedockerfile"}, {strategy: &buildapi.DockerBuildStrategy{DockerfilePath: "somedir/somedockerfile"}, expectedDockerfilePath: "somedir/somedockerfile"}}
	for count, test := range tests {
		errors := validateDockerStrategy(test.strategy, nil)
		if len(errors) != 0 {
			t.Errorf("Test[%d] Unexpected validation error: %v", count, errors)
		}
		if test.strategy.DockerfilePath != test.expectedDockerfilePath {
			t.Errorf("Test[%d] Unexpected DockerfilePath: %v (expected: %s)", count, test.strategy.DockerfilePath, test.expectedDockerfilePath)
		}
	}
}
func TestValidateJenkinsfilePath(t *testing.T) {
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
	tests := []struct {
		strategy		*buildapi.JenkinsPipelineBuildStrategy
		expectedJenkinsfilePath	string
	}{{strategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "."}, expectedJenkinsfilePath: ""}, {strategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "somedir/.."}, expectedJenkinsfilePath: ""}, {strategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "somedir/../somedockerfile"}, expectedJenkinsfilePath: "somedockerfile"}, {strategy: &buildapi.JenkinsPipelineBuildStrategy{JenkinsfilePath: "somedir/somedockerfile"}, expectedJenkinsfilePath: "somedir/somedockerfile"}}
	for count, test := range tests {
		errors := validateJenkinsPipelineStrategy(test.strategy, nil)
		if len(errors) != 0 {
			t.Errorf("Test[%d] Unexpected validation error: %v", count, errors)
		}
		if test.strategy.JenkinsfilePath != test.expectedJenkinsfilePath {
			t.Errorf("Test[%d] Unexpected JenkinsfilePath: %v (expected: %s)", count, test.strategy.JenkinsfilePath, test.expectedJenkinsfilePath)
		}
	}
}
func TestValidateTrigger(t *testing.T) {
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
	tests := map[string]struct {
		trigger		buildapi.BuildTriggerPolicy
		expected	[]*field.Error
	}{"trigger without type": {trigger: buildapi.BuildTriggerPolicy{}, expected: []*field.Error{field.Required(field.NewPath("type"), "")}}, "trigger with unknown type": {trigger: buildapi.BuildTriggerPolicy{Type: "UnknownTriggerType"}, expected: []*field.Error{field.Invalid(field.NewPath("type"), "", "")}}, "GitHub type with no github webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType}, expected: []*field.Error{field.Required(field.NewPath("github"), "")}}, "GitHub trigger with no secret": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType, GitHubWebHook: &buildapi.WebHookTrigger{}}, expected: []*field.Error{field.Invalid(field.NewPath("github"), buildapi.WebHookTrigger{}, "must provide a value for at least one of secret or secretReference")}}, "GitHub trigger with generic webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{Secret: "secret101"}}, expected: []*field.Error{field.Required(field.NewPath("github"), "")}}, "GitHub trigger with allow env": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType, GitHubWebHook: &buildapi.WebHookTrigger{Secret: "secret101", AllowEnv: true}}, expected: []*field.Error{field.Invalid(field.NewPath("github", "allowEnv"), "", "")}}, "GitLab type with no gitlab webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitLabWebHookBuildTriggerType}, expected: []*field.Error{field.Required(field.NewPath("gitlab"), "")}}, "GitLab trigger with no secret": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitLabWebHookBuildTriggerType, GitLabWebHook: &buildapi.WebHookTrigger{}}, expected: []*field.Error{field.Invalid(field.NewPath("gitlab"), buildapi.WebHookTrigger{}, "must provide a value for at least one of secret or secretReference")}}, "GitLab trigger with generic webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitLabWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{Secret: "secret101"}}, expected: []*field.Error{field.Required(field.NewPath("gitlab"), "")}}, "GitLab trigger with allow env": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitLabWebHookBuildTriggerType, GitLabWebHook: &buildapi.WebHookTrigger{Secret: "secret101", AllowEnv: true}}, expected: []*field.Error{field.Invalid(field.NewPath("gitlab", "allowEnv"), "", "")}}, "Bitbucket type with no Bitbucket webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.BitbucketWebHookBuildTriggerType}, expected: []*field.Error{field.Required(field.NewPath("bitbucket"), "")}}, "Bitbucket trigger with no secret": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.BitbucketWebHookBuildTriggerType, BitbucketWebHook: &buildapi.WebHookTrigger{}}, expected: []*field.Error{field.Invalid(field.NewPath("bitbucket"), buildapi.WebHookTrigger{}, "must provide a value for at least one of secret or secretReference")}}, "Bitbucket trigger with generic webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.BitbucketWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{Secret: "secret101"}}, expected: []*field.Error{field.Required(field.NewPath("bitbucket"), "")}}, "Bitbucket trigger with allow env": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.BitbucketWebHookBuildTriggerType, BitbucketWebHook: &buildapi.WebHookTrigger{Secret: "secret101", AllowEnv: true}}, expected: []*field.Error{field.Invalid(field.NewPath("bitbucket", "allowEnv"), "", "")}}, "Generic trigger with no generic webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GenericWebHookBuildTriggerType}, expected: []*field.Error{field.Required(field.NewPath("generic"), "")}}, "Generic trigger with no secret": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GenericWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{}}, expected: []*field.Error{field.Invalid(field.NewPath("generic"), buildapi.WebHookTrigger{}, "must provide a value for at least one of secret or secretReference")}}, "Generic trigger with github webhook": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GenericWebHookBuildTriggerType, GitHubWebHook: &buildapi.WebHookTrigger{Secret: "secret101"}}, expected: []*field.Error{field.Required(field.NewPath("generic"), "")}}, "Webhook trigger with no secretref name": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GenericWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{}}}, expected: []*field.Error{field.Required(field.NewPath("generic.secretReference.name"), "")}}, "ImageChange trigger without params": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.ImageChangeBuildTriggerType}, expected: []*field.Error{field.Required(field.NewPath("imageChange"), "")}}, "valid GitHub trigger": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType, GitHubWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{Name: "mysecret"}}}}, "valid GitHub trigger with secretref": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitHubWebHookBuildTriggerType, GitHubWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{Name: "mysecret"}}}}, "valid GitLab trigger with secretref": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GitLabWebHookBuildTriggerType, GitLabWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{Name: "mysecret"}}}}, "valid Bitbucket trigger with secretref": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.BitbucketWebHookBuildTriggerType, BitbucketWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{Name: "mysecret"}}}}, "valid Generic trigger with secretref": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.GenericWebHookBuildTriggerType, GenericWebHook: &buildapi.WebHookTrigger{SecretReference: &buildapi.SecretLocalReference{Name: "mysecret"}}}}, "valid ImageChange trigger": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{LastTriggeredImageID: "asdf1234"}}}, "valid ImageChange trigger with empty fields": {trigger: buildapi.BuildTriggerPolicy{Type: buildapi.ImageChangeBuildTriggerType, ImageChange: &buildapi.ImageChangeTrigger{}}}}
	for desc, test := range tests {
		errors := validateTrigger(&test.trigger, &kapi.ObjectReference{Kind: "ImageStreamTag"}, nil)
		if len(test.expected) == 0 {
			if len(errors) != 0 {
				t.Errorf("%s: Got unexpected validation errors: %#v", desc, errors)
			}
			continue
		}
		if len(errors) != 1 {
			t.Errorf("%s: Expected one validation error, got %d", desc, len(errors))
			for i, err := range errors {
				validationError := err
				t.Errorf("  %d. %v", i+1, validationError)
			}
			continue
		}
		err := errors[0]
		validationError := err
		if validationError.Type != test.expected[0].Type {
			t.Errorf("%s: Expected error type %s, got %s", desc, test.expected[0].Type, validationError.Type)
		}
		if validationError.Field != test.expected[0].Field {
			t.Errorf("%s: Expected error field %s, got %s", desc, test.expected[0].Field, validationError.Field)
		}
	}
}
func TestValidateToImageReference(t *testing.T) {
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
	o := &kapi.ObjectReference{Name: "somename", Namespace: "somenamespace", Kind: "DockerImage"}
	errs := validateToImageReference(o, nil)
	if len(errs) != 1 {
		t.Errorf("Wrong number of errors: %v", errs)
	}
	err := errs[0]
	if err.Type != field.ErrorTypeInvalid {
		t.Errorf("Wrong error type, expected %v, got %v", field.ErrorTypeInvalid, err.Type)
	}
	if err.Field != "namespace" {
		t.Errorf("Error on wrong field, expected %s, got %s", "namespace", err.Field)
	}
}
func TestValidateStrategyEnvVars(t *testing.T) {
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
	tests := []struct {
		env		[]kapi.EnvVar
		errExpected	bool
		errField	string
		errType		field.ErrorType
	}{{env: []kapi.EnvVar{{Name: "", Value: "test"}}, errExpected: true, errField: "env[0].name", errType: field.ErrorTypeRequired}, {env: []kapi.EnvVar{{Name: " invalid,name", Value: "test"}}, errExpected: true, errField: "env[0].name", errType: field.ErrorTypeInvalid}, {env: []kapi.EnvVar{{Name: "VAR1", Value: "value1"}, {Name: "VAR2", Value: "value2"}}}}
	for i, tc := range tests {
		errs := ValidateStrategyEnv(tc.env, field.NewPath("env"))
		if !tc.errExpected {
			if len(errs) > 0 {
				t.Errorf("%d: unexpected error: %v", i, errs.ToAggregate())
			}
			continue
		}
		if tc.errExpected && len(errs) == 0 {
			t.Errorf("%d: expected error. Got none.", i)
			continue
		}
		err := errs[0]
		if err.Field != tc.errField {
			t.Errorf("%d: unexpected error field: %s", i, err.Field)
		}
		if err.Type != tc.errType {
			t.Errorf("%d: unexpected error type: %s", i, err.Type)
		}
	}
}
func TestValidatePostCommit(t *testing.T) {
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
	path := field.NewPath("postCommit")
	invalidSpec := buildapi.BuildPostCommitSpec{Command: []string{"rake", "test"}, Script: "rake test"}
	tests := []struct {
		spec	buildapi.BuildPostCommitSpec
		want	field.ErrorList
	}{{spec: buildapi.BuildPostCommitSpec{}, want: field.ErrorList{}}, {spec: buildapi.BuildPostCommitSpec{Script: "rake test"}, want: field.ErrorList{}}, {spec: buildapi.BuildPostCommitSpec{Command: []string{"rake", "test"}}, want: field.ErrorList{}}, {spec: buildapi.BuildPostCommitSpec{Command: []string{"/bin/bash", "-c"}, Args: []string{"rake test"}}, want: field.ErrorList{}}, {spec: invalidSpec, want: field.ErrorList{field.Invalid(path, invalidSpec, "cannot use command and script together")}}}
	for _, tt := range tests {
		if got := validatePostCommit(tt.spec, path); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("validatePostCommitSpec(%+v) = %v, want %v", tt.spec, got, tt.want)
		}
	}
}
func TestDiffBuildSpec(t *testing.T) {
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
	tests := []struct {
		name		string
		older, newer	buildapi.BuildSpec
		expected	string
	}{{name: "context dir", older: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{}}}, newer: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{ContextDir: "context-dir"}}}, expected: `{"spec":{"source":{"contextDir":"context-dir"}}}`}, {name: "same git build source", older: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{Ref: "https://github.com/openshift/origin.git"}}}}, newer: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{Ref: "https://github.com/openshift/origin.git"}}}}, expected: "{}"}, {name: "different git build source", older: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{Ref: "https://github.com/openshift/origin.git"}}}}, newer: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Git: &buildapi.GitBuildSource{Ref: "https://github.com/ose/origin.git"}}}}, expected: `{"spec":{"source":{"git":{"ref":"https://github.com/ose/origin.git"}}}}`}}
	for _, test := range tests {
		diff, err := diffBuildSpec(test.newer, test.older)
		if err != nil {
			t.Errorf("%s: unexpected: %v", test.name, err)
			continue
		}
		if diff != test.expected {
			t.Errorf("%s: expected: %s, got: %s", test.name, test.expected, diff)
		}
	}
}
func TestValidateBuildImageRefs(t *testing.T) {
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
	tests := []struct {
		name		string
		build		buildapi.Build
		expectedError	string
	}{{name: "valid docker build", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}}}}, expectedError: ""}, {name: "invalid docker image ref", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "!!!myimage:tag"}}}}}}, expectedError: "not a valid Docker pull specification: invalid reference format"}, {name: "docker build with ImageStreamTag in from", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "ImageStreamTag", Name: "myimagestream"}}}}}}, expectedError: "must be <name>:<tag>"}, {name: "s2i build with valid source image references", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}, Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimagestream:tag"}}}}}}, expectedError: ""}, {name: "image with sources uses ImageStreamTag", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}, Images: []buildapi.ImageSource{{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}, {From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: ""}, Paths: []buildapi.ImageSourcePath{{SourcePath: "/some/path", DestinationDir: "test/dir"}}}}}, Strategy: buildapi.BuildStrategy{SourceStrategy: &buildapi.SourceBuildStrategy{From: kapi.ObjectReference{Kind: "DockerImage", Name: "myimagestream:tag"}}}}}}, expectedError: "Required value"}, {name: "custom build with ImageStreamTag in from", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{CustomStrategy: &buildapi.CustomBuildStrategy{From: kapi.ObjectReference{Kind: "ImageStreamTag", Name: ":tag"}}}}}}, expectedError: "invalid"}}
	for _, tc := range tests {
		errs := ValidateBuild(&tc.build)
		if tc.expectedError == "" && len(errs) > 0 {
			t.Errorf("%s: Unexpected validation result: %v", tc.name, errs)
		}
		if tc.expectedError != "" {
			found := false
			for _, err := range errs {
				if strings.Contains(err.Error(), tc.expectedError) {
					found = true
				}
			}
			if !found {
				t.Errorf("%s: Expected to fail with %q, result: %v", tc.name, tc.expectedError, errs)
			}
		}
	}
}
func TestValidateBuildUpdateImageReferences(t *testing.T) {
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
	tests := []struct {
		name		string
		old		buildapi.Build
		build		buildapi.Build
		expectedError	string
	}{{name: "invalid docker image reference is ignored if it doesn't change", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default", ResourceVersion: "10"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "!!!myimage:tag"}}}}}}, old: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "!!!myimage:tag"}}}}}}, expectedError: ""}, {name: "docker image reference is immutable", build: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default", ResourceVersion: "10"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag2"}}}}}}, old: buildapi.Build{ObjectMeta: metav1.ObjectMeta{Name: "build", Namespace: "default"}, Spec: buildapi.BuildSpec{CommonSpec: buildapi.CommonSpec{Source: buildapi.BuildSource{Binary: &buildapi.BinaryBuildSource{}}, Strategy: buildapi.BuildStrategy{DockerStrategy: &buildapi.DockerBuildStrategy{From: &kapi.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}}}}, expectedError: "spec is immutable"}}
	for _, tc := range tests {
		errs := ValidateBuildUpdate(&tc.build, &tc.old)
		if tc.expectedError == "" && len(errs) > 0 {
			t.Errorf("%s: Unexpected validation result: %v", tc.name, errs)
		}
		if tc.expectedError != "" {
			found := false
			for _, err := range errs {
				if strings.Contains(err.Error(), tc.expectedError) {
					found = true
				}
			}
			if !found {
				t.Errorf("%s: Expected to fail with %q, result: %v", tc.name, tc.expectedError, errs)
			}
		}
	}
}
