package common

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	buildv1 "github.com/openshift/api/build/v1"
	buildfake "github.com/openshift/client-go/build/clientset/versioned/fake"
	buildclient "github.com/openshift/origin/pkg/build/client"
	buildutil "github.com/openshift/origin/pkg/build/util"
)

func mockBuildConfig(name string) buildv1.BuildConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	appName := strings.Split(name, "-")
	successfulBuildsToKeep := int32(2)
	failedBuildsToKeep := int32(3)
	return buildv1.BuildConfig{ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("%s-build", appName[0]), Namespace: "namespace", Labels: map[string]string{"app": appName[0]}}, Spec: buildv1.BuildConfigSpec{SuccessfulBuildsHistoryLimit: &successfulBuildsToKeep, FailedBuildsHistoryLimit: &failedBuildsToKeep}}
}
func mockBuild(name string, phase buildv1.BuildPhase, stamp *metav1.Time) buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	appName := strings.Split(name, "-")
	return buildv1.Build{ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(fmt.Sprintf("uid%v", appName[1])), Namespace: "namespace", CreationTimestamp: *stamp, Labels: map[string]string{"app": appName[0], buildutil.BuildConfigLabel: fmt.Sprintf("%v-build", appName[0]), "buildconfig": fmt.Sprintf("%v-build", appName[0])}, Annotations: map[string]string{buildutil.BuildConfigLabel: fmt.Sprintf("%v-build", appName[0])}}, Status: buildv1.BuildStatus{Phase: phase, StartTimestamp: stamp, Config: &corev1.ObjectReference{Name: fmt.Sprintf("%v-build", appName[0]), Namespace: "namespace"}}}
}
func mockBuildsList(length int) (buildv1.BuildConfig, []buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var builds []buildv1.Build
	buildPhaseList := []buildv1.BuildPhase{buildv1.BuildPhaseComplete, buildv1.BuildPhaseFailed, buildv1.BuildPhaseError, buildv1.BuildPhaseCancelled}
	addOrSubtract := []string{"+", "-"}
	j := 0
	for i := 0; i < length; i++ {
		duration, _ := time.ParseDuration(fmt.Sprintf("%v%vh", addOrSubtract[i%2], i))
		startTime := metav1.NewTime(time.Now().Add(duration))
		build := mockBuild(fmt.Sprintf("myapp-%v", i), buildPhaseList[j], &startTime)
		builds = append(builds, build)
		j++
		if j == 4 {
			j = 0
		}
	}
	return mockBuildConfig("myapp"), builds
}
func TestHandleBuildPruning(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var objects []runtime.Object
	buildconfig, builds := mockBuildsList(16)
	objects = append(objects, &buildconfig)
	for index := range builds {
		objects = append(objects, &builds[index])
	}
	buildClient := buildfake.NewSimpleClientset(objects...)
	build, err := buildClient.BuildV1().Builds("namespace").Get("myapp-0", metav1.GetOptions{})
	if err != nil {
		t.Errorf("%v", err)
	}
	buildLister := buildclient.NewClientBuildLister(buildClient.BuildV1())
	buildConfigGetter := buildclient.NewClientBuildConfigLister(buildClient.BuildV1())
	buildDeleter := buildclient.NewClientBuildClient(buildClient)
	bcName := buildutil.ConfigNameForBuild(build)
	successfulStartingBuilds, err := buildutil.BuildConfigBuilds(buildLister, build.Namespace, bcName, func(build *buildv1.Build) bool {
		return build.Status.Phase == buildv1.BuildPhaseComplete
	})
	sort.Sort(ByCreationTimestamp(successfulStartingBuilds))
	failedStartingBuilds, err := buildutil.BuildConfigBuilds(buildLister, build.Namespace, bcName, func(build *buildv1.Build) bool {
		return build.Status.Phase == buildv1.BuildPhaseFailed || build.Status.Phase == buildv1.BuildPhaseError || build.Status.Phase == buildv1.BuildPhaseCancelled
	})
	sort.Sort(ByCreationTimestamp(failedStartingBuilds))
	if len(successfulStartingBuilds)+len(failedStartingBuilds) != 16 {
		t.Errorf("should start with 16 builds, but started with %v instead", len(successfulStartingBuilds)+len(failedStartingBuilds))
	}
	if err := HandleBuildPruning(bcName, build.Namespace, buildLister, buildConfigGetter, buildDeleter); err != nil {
		t.Errorf("error pruning builds: %v", err)
	}
	successfulRemainingBuilds, err := buildutil.BuildConfigBuilds(buildLister, build.Namespace, bcName, func(build *buildv1.Build) bool {
		return build.Status.Phase == buildv1.BuildPhaseComplete
	})
	sort.Sort(ByCreationTimestamp(successfulRemainingBuilds))
	failedRemainingBuilds, err := buildutil.BuildConfigBuilds(buildLister, build.Namespace, bcName, func(build *buildv1.Build) bool {
		return build.Status.Phase == buildv1.BuildPhaseFailed || build.Status.Phase == buildv1.BuildPhaseError || build.Status.Phase == buildv1.BuildPhaseCancelled
	})
	sort.Sort(ByCreationTimestamp(failedRemainingBuilds))
	if len(successfulRemainingBuilds)+len(failedRemainingBuilds) != 5 {
		t.Errorf("there should only be 5 builds left, but instead there are %v", len(successfulRemainingBuilds)+len(failedRemainingBuilds))
	}
	if !reflect.DeepEqual(successfulStartingBuilds[:2], successfulRemainingBuilds) {
		t.Errorf("expected the two most recent successful builds should be left, but instead there were %v: %v", len(successfulRemainingBuilds), successfulRemainingBuilds)
	}
	if !reflect.DeepEqual(failedStartingBuilds[:3], failedRemainingBuilds) {
		t.Errorf("expected the three most recent failed builds to be left, but instead there were %v: %v", len(failedRemainingBuilds), failedRemainingBuilds)
	}
}
