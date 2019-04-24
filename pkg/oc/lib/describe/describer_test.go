package describe

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"text/tabwriter"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/api/image/docker10"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	_ "github.com/openshift/origin/pkg/api/install"
	"github.com/openshift/origin/pkg/api/legacy"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
)

type describeClient struct {
	T		*testing.T
	Namespace	string
	Err		error
}

var DescriberCoverageExceptions = []reflect.Type{reflect.TypeOf(&buildapi.BuildLog{}), reflect.TypeOf(&buildapi.BuildLogOptions{}), reflect.TypeOf(&buildapi.BinaryBuildRequestOptions{}), reflect.TypeOf(&buildapi.BuildRequest{}), reflect.TypeOf(&appsapi.DeploymentConfigRollback{}), reflect.TypeOf(&appsapi.DeploymentLog{}), reflect.TypeOf(&appsapi.DeploymentLogOptions{}), reflect.TypeOf(&appsapi.DeploymentRequest{}), reflect.TypeOf(&imageapi.DockerImage{}), reflect.TypeOf(&imageapi.ImageStreamImport{}), reflect.TypeOf(&oauthapi.OAuthAccessToken{}), reflect.TypeOf(&oauthapi.OAuthAuthorizeToken{}), reflect.TypeOf(&oauthapi.OAuthClientAuthorization{}), reflect.TypeOf(&projectapi.ProjectRequest{}), reflect.TypeOf(&templateapi.TemplateInstance{}), reflect.TypeOf(&templateapi.BrokerTemplateInstance{}), reflect.TypeOf(&authorizationapi.IsPersonalSubjectAccessReview{}), reflect.TypeOf(&imageapi.ImageSignature{}), reflect.TypeOf(&authorizationapi.SubjectAccessReviewResponse{}), reflect.TypeOf(&authorizationapi.ResourceAccessReviewResponse{}), reflect.TypeOf(&authorizationapi.SubjectAccessReview{}), reflect.TypeOf(&authorizationapi.ResourceAccessReview{}), reflect.TypeOf(&authorizationapi.LocalSubjectAccessReview{}), reflect.TypeOf(&authorizationapi.LocalResourceAccessReview{}), reflect.TypeOf(&authorizationapi.SelfSubjectRulesReview{}), reflect.TypeOf(&authorizationapi.SubjectRulesReview{}), reflect.TypeOf(&securityapi.PodSecurityPolicySubjectReview{}), reflect.TypeOf(&securityapi.PodSecurityPolicySelfSubjectReview{}), reflect.TypeOf(&securityapi.PodSecurityPolicyReview{}), reflect.TypeOf(&oauthapi.OAuthRedirectReference{})}
var MissingDescriberCoverageExceptions = []reflect.Type{reflect.TypeOf(&imageapi.ImageStreamMapping{}), reflect.TypeOf(&oauthapi.OAuthClient{})}

func TestDescriberCoverage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
main:
	for _, apiType := range legacyscheme.Scheme.KnownTypes(legacy.InternalGroupVersion) {
		if !strings.HasPrefix(apiType.PkgPath(), "github.com/openshift/origin") || strings.HasPrefix(apiType.PkgPath(), "github.com/openshift/origin/vendor/") {
			continue
		}
		if strings.HasSuffix(apiType.Name(), "List") {
			continue
		}
		ptrType := reflect.PtrTo(apiType)
		for _, exception := range DescriberCoverageExceptions {
			if ptrType == exception {
				continue main
			}
		}
		for _, exception := range MissingDescriberCoverageExceptions {
			if ptrType == exception {
				continue main
			}
		}
		gk := legacy.InternalGroupVersion.WithKind(apiType.Name()).GroupKind()
		_, ok := DescriberFor(gk, &rest.Config{}, fake.NewSimpleClientset(), "")
		if !ok {
			t.Errorf("missing describer for %v.  Check pkg/cmd/cli/describe/describer.go", apiType)
		}
	}
}
func TestDescribeBuildDuration(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type testBuild struct {
		build	*buildv1.Build
		output	string
	}
	now := metav1.Now()
	minuteAgo := metav1.Unix(now.Rfc3339Copy().Time.Unix()-60, 0)
	twoMinutesAgo := metav1.Unix(now.Rfc3339Copy().Time.Unix()-120, 0)
	threeMinutesAgo := metav1.Unix(now.Rfc3339Copy().Time.Unix()-180, 0)
	tests := []testBuild{{&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: minuteAgo}, Status: buildv1.BuildStatus{Phase: buildv1.BuildPhaseNew}}, "waiting for 1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: minuteAgo}, Status: buildv1.BuildStatus{Phase: buildv1.BuildPhasePending}}, "waiting for 1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: twoMinutesAgo}, Status: buildv1.BuildStatus{StartTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseRunning}}, "running for 1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{StartTimestamp: &twoMinutesAgo, CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseComplete}}, "1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{StartTimestamp: &twoMinutesAgo, CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseFailed}}, "1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{StartTimestamp: &twoMinutesAgo, CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseError}}, "1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseCancelled}}, "waited for 2m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{StartTimestamp: &twoMinutesAgo, CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseCancelled}}, "1m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseFailed}}, "waited for 2m"}, {&buildv1.Build{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: threeMinutesAgo}, Status: buildv1.BuildStatus{CompletionTimestamp: &minuteAgo, Phase: buildv1.BuildPhaseError}}, "waited for 2m"}}
	for i, tc := range tests {
		if actual, expected := describeBuildDuration(tc.build), tc.output; !strings.Contains(actual, expected) {
			t.Errorf("(%d) expected duration output %s, got %s", i, expected, actual)
		}
	}
}
func mkV1Pod(status corev1.PodPhase, exitCode int) *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "PodName"}, Status: corev1.PodStatus{Phase: status, ContainerStatuses: []corev1.ContainerStatus{{State: corev1.ContainerState{Terminated: &corev1.ContainerStateTerminated{ExitCode: int32(exitCode)}}}}}}
}
func mkPod(status kapi.PodPhase, exitCode int) *kapi.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kapi.Pod{ObjectMeta: metav1.ObjectMeta{Name: "PodName"}, Status: kapi.PodStatus{Phase: status, ContainerStatuses: []kapi.ContainerStatus{{State: kapi.ContainerState{Terminated: &kapi.ContainerStateTerminated{ExitCode: int32(exitCode)}}}}}}
}
func TestDescribePostCommitHook(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		hook	buildv1.BuildPostCommitSpec
		want	string
	}{{hook: buildv1.BuildPostCommitSpec{}, want: ""}, {hook: buildv1.BuildPostCommitSpec{Script: "go test"}, want: `"/bin/sh", "-ic", "go test"`}, {hook: buildv1.BuildPostCommitSpec{Command: []string{"go", "test"}}, want: `"go", "test"`}, {hook: buildv1.BuildPostCommitSpec{Args: []string{"go", "test"}}, want: `"<image-entrypoint>", "go", "test"`}, {hook: buildv1.BuildPostCommitSpec{Script: `go test "$@"`, Args: []string{"-v", "-timeout", "2s"}}, want: `"/bin/sh", "-ic", "go test \"$@\"", "/bin/sh", "-v", "-timeout", "2s"`}, {hook: buildv1.BuildPostCommitSpec{Command: []string{"go", "test"}, Args: []string{"-v", "-timeout", "2s"}}, want: `"go", "test", "-v", "-timeout", "2s"`}, {hook: buildv1.BuildPostCommitSpec{Script: "go test -v", Command: []string{"go", "test"}}, want: `"/bin/sh", "-ic", "go test -v"`}}
	for _, tt := range tests {
		var b bytes.Buffer
		out := tabwriter.NewWriter(&b, 0, 8, 0, '\t', 0)
		describePostCommitHook(tt.hook, out)
		if err := out.Flush(); err != nil {
			t.Fatalf("%+v: flush error: %v", tt.hook, err)
		}
		var want string
		if tt.want != "" {
			want = fmt.Sprintf("Post Commit Hook:\t[%s]\n", tt.want)
		}
		if got := b.String(); got != want {
			t.Errorf("describePostCommitHook(%+v, out) = %q, want %q", tt.hook, got, want)
		}
	}
}
func TestDescribeBuildSpec(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		spec	buildv1.BuildSpec
		want	string
	}{{spec: buildv1.BuildSpec{CommonSpec: buildv1.CommonSpec{Source: buildv1.BuildSource{Git: &buildv1.GitBuildSource{URI: "http://github.com/my/repository"}, ContextDir: "context"}, Strategy: buildv1.BuildStrategy{DockerStrategy: &buildv1.DockerBuildStrategy{}}, Output: buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, want: "URL"}, {spec: buildv1.BuildSpec{CommonSpec: buildv1.CommonSpec{Source: buildv1.BuildSource{}, Strategy: buildv1.BuildStrategy{SourceStrategy: &buildv1.SourceBuildStrategy{From: corev1.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, want: "Empty Source"}, {spec: buildv1.BuildSpec{CommonSpec: buildv1.CommonSpec{Source: buildv1.BuildSource{}, Strategy: buildv1.BuildStrategy{CustomStrategy: &buildv1.CustomBuildStrategy{From: corev1.ObjectReference{Kind: "DockerImage", Name: "myimage:tag"}}}, Output: buildv1.BuildOutput{To: &corev1.ObjectReference{Kind: "DockerImage", Name: "repository/data"}}}}, want: "Empty Source"}, {spec: buildv1.BuildSpec{CommonSpec: buildv1.CommonSpec{Source: buildv1.BuildSource{}, Strategy: buildv1.BuildStrategy{JenkinsPipelineStrategy: &buildv1.JenkinsPipelineBuildStrategy{Jenkinsfile: "openshiftBuild"}}}}, want: "openshiftBuild"}}
	for _, tt := range tests {
		var b bytes.Buffer
		out := tabwriter.NewWriter(&b, 0, 8, 0, '\t', 0)
		describeCommonSpec(tt.spec.CommonSpec, out)
		if err := out.Flush(); err != nil {
			t.Fatalf("%+v: flush error: %v", tt.spec, err)
		}
		if got := b.String(); !strings.Contains(got, tt.want) {
			t.Errorf("describeBuildSpec(%+v, out) = %q, should contain %q", tt.spec, got, tt.want)
		}
	}
}
func TestDescribeImage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		image	imageapi.Image
		want	[]string
	}{{image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: "test"}}, want: []string{"Name:.+test"}}, {image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: "test"}, DockerImageLayers: []imageapi.ImageLayer{{Name: "sha256:1234", LayerSize: 3409}, {Name: "sha256:5678", LayerSize: 1024}}}, want: []string{"Layers:.+3.409kB\\ssha256:1234", "1.024kB\\ssha256:5678", "Image Size:.+0B in 2 layers"}}, {image: imageapi.Image{ObjectMeta: metav1.ObjectMeta{Name: "test"}, DockerImageMetadata: docker10.DockerImage{Size: 4430}, DockerImageLayers: []imageapi.ImageLayer{{Name: "sha256:1234", LayerSize: 3409}, {Name: "sha256:5678", LayerSize: 1024}}}, want: []string{"Layers:.+3.409kB\\ssha256:1234", "1.024kB\\ssha256:5678", "Image Size:.+4.43kB in 2 layers"}}}
	for i := range tests {
		tt := tests[i]
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			out, err := DescribeImage(&tt.image, tt.image.Name)
			if err != nil {
				t.Fatal(err)
			}
			for _, match := range tt.want {
				if got := out; !regexp.MustCompile(match).MatchString(got) {
					t.Errorf("%s\nshould contain %q", got, match)
				}
			}
		})
	}
}
