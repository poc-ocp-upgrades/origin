package deployconfig

import (
	"reflect"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	appstest "github.com/openshift/origin/pkg/apps/apis/apps/internaltest"
)

var (
	nonDefaultRevisionHistoryLimit = appsapi.DefaultRevisionHistoryLimit + 42
)

func int32ptr(v int32) *int32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v
}
func TestDeploymentConfigStrategy(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := apirequest.NewDefaultContext()
	if !CommonStrategy.NamespaceScoped() {
		t.Errorf("DeploymentConfig is namespace scoped")
	}
	if CommonStrategy.AllowCreateOnUpdate() {
		t.Errorf("DeploymentConfig should not allow create on update")
	}
	deploymentConfig := &appsapi.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"}, Spec: appstest.OkDeploymentConfigSpec()}
	CommonStrategy.PrepareForCreate(ctx, deploymentConfig)
	errs := CommonStrategy.Validate(ctx, deploymentConfig)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}
	updatedDeploymentConfig := &appsapi.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Name: "bar", Namespace: "default", Generation: 1}, Spec: appstest.OkDeploymentConfigSpec()}
	errs = CommonStrategy.ValidateUpdate(ctx, updatedDeploymentConfig, deploymentConfig)
	if len(errs) == 0 {
		t.Errorf("Expected error validating")
	}
	updatedDeploymentConfig.Name = "foo"
	updatedDeploymentConfig.ResourceVersion = "1"
	errs = CommonStrategy.ValidateUpdate(ctx, updatedDeploymentConfig, deploymentConfig)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}
	invalidDeploymentConfig := &appsapi.DeploymentConfig{}
	errs = CommonStrategy.Validate(ctx, invalidDeploymentConfig)
	if len(errs) == 0 {
		t.Errorf("Expected error validating")
	}
}
func TestPrepareForUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := apirequest.NewDefaultContext()
	tests := []struct {
		name		string
		prev		runtime.Object
		after		runtime.Object
		expected	runtime.Object
	}{{name: "latestVersion bump", prev: prevDeployment(), after: afterDeploymentVersionBump(), expected: expectedAfterVersionBump()}, {name: "spec change", prev: prevDeployment(), after: afterDeployment(), expected: expectedAfterDeployment()}}
	for _, test := range tests {
		strategy{}.PrepareForUpdate(ctx, test.after, test.prev)
		if !reflect.DeepEqual(test.expected, test.after) {
			t.Errorf("%s: unexpected object mismatch! Expected:\n%#v\ngot:\n%#v", test.name, test.expected, test.after)
		}
	}
}
func prevDeployment() *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default", Generation: 4, Annotations: make(map[string]string)}, Spec: appstest.OkDeploymentConfigSpec(), Status: appstest.OkDeploymentConfigStatus(1)}
}
func afterDeployment() *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := prevDeployment()
	dc.Spec.Replicas++
	return dc
}
func expectedAfterDeployment() *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := afterDeployment()
	dc.Generation++
	return dc
}
func afterDeploymentVersionBump() *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := prevDeployment()
	dc.Status.LatestVersion++
	return dc
}
func expectedAfterVersionBump() *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := afterDeploymentVersionBump()
	dc.Generation++
	return dc
}
func setRevisionHistoryLimit(v *int32, dc *appsapi.DeploymentConfig) *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc.Spec.RevisionHistoryLimit = v
	return dc
}
func okDeploymentConfig(generation int64) *appsapi.DeploymentConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := appstest.OkDeploymentConfig(0)
	dc.ObjectMeta.Generation = generation
	return dc
}
func TestLegacyStrategy_PrepareForCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nonDefaultRevisionHistoryLimit := appsapi.DefaultRevisionHistoryLimit + 42
	tt := []struct {
		obj		*appsapi.DeploymentConfig
		expected	*appsapi.DeploymentConfig
	}{{obj: setRevisionHistoryLimit(nil, okDeploymentConfig(0)), expected: setRevisionHistoryLimit(nil, okDeploymentConfig(1))}, {obj: setRevisionHistoryLimit(&nonDefaultRevisionHistoryLimit, okDeploymentConfig(0)), expected: setRevisionHistoryLimit(&nonDefaultRevisionHistoryLimit, okDeploymentConfig(1))}}
	ctx := apirequest.NewDefaultContext()
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			LegacyStrategy.PrepareForCreate(ctx, tc.obj)
			if !reflect.DeepEqual(tc.obj, tc.expected) {
				t.Fatalf("LegacyStrategy.PrepareForCreate failed:%s", diff.ObjectReflectDiff(tc.obj, tc.expected))
			}
			errs := LegacyStrategy.Validate(ctx, tc.obj)
			if len(errs) != 0 {
				t.Errorf("Unexpected error validating DeploymentConfig: %v", errs)
			}
		})
	}
}
func TestLegacyStrategy_DefaultGarbageCollectionPolicy(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	expected := rest.OrphanDependents
	got := LegacyStrategy.DefaultGarbageCollectionPolicy(nil)
	if got != expected {
		t.Fatalf("Default garbage collection policy for DeploymentConfigs should be %q (not %q)", expected, got)
	}
}
func TestGroupStrategy_PrepareForCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := []struct {
		obj		*appsapi.DeploymentConfig
		expected	*appsapi.DeploymentConfig
	}{{obj: setRevisionHistoryLimit(nil, okDeploymentConfig(0)), expected: setRevisionHistoryLimit(int32ptr(appsapi.DefaultRevisionHistoryLimit), okDeploymentConfig(1))}, {obj: setRevisionHistoryLimit(&nonDefaultRevisionHistoryLimit, okDeploymentConfig(0)), expected: setRevisionHistoryLimit(&nonDefaultRevisionHistoryLimit, okDeploymentConfig(1))}}
	ctx := apirequest.NewDefaultContext()
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			GroupStrategy.PrepareForCreate(ctx, tc.obj)
			if !reflect.DeepEqual(tc.obj, tc.expected) {
				t.Fatalf("GroupStrategy.PrepareForCreate failed:%s", diff.ObjectReflectDiff(tc.obj, tc.expected))
			}
			errs := GroupStrategy.Validate(ctx, tc.obj)
			if len(errs) != 0 {
				t.Errorf("Unexpected error validating DeploymentConfig: %v", errs)
			}
		})
	}
}
