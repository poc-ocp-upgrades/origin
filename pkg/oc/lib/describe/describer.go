package describe

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
	"k8s.io/klog"
	units "github.com/docker/go-units"
	corev1 "k8s.io/api/core/v1"
	kerrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/kubectl/describe"
	"k8s.io/kubernetes/pkg/kubectl/describe/versioned"
	oapps "github.com/openshift/api/apps"
	"github.com/openshift/api/authorization"
	"github.com/openshift/api/build"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/api/image"
	"github.com/openshift/api/network"
	"github.com/openshift/api/oauth"
	"github.com/openshift/api/project"
	"github.com/openshift/api/quota"
	quotav1 "github.com/openshift/api/quota/v1"
	"github.com/openshift/api/route"
	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/api/security"
	"github.com/openshift/api/template"
	"github.com/openshift/api/user"
	appstypedclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	buildv1clienttyped "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	onetworktypedclient "github.com/openshift/client-go/network/clientset/versioned/typed/network/v1"
	quotaclient "github.com/openshift/client-go/quota/clientset/versioned/typed/quota/v1"
	oapi "github.com/openshift/origin/pkg/api"
	"github.com/openshift/origin/pkg/api/legacy"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	oauthorizationclient "github.com/openshift/origin/pkg/authorization/generated/internalclientset/typed/authorization/internalversion"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	buildutil "github.com/openshift/origin/pkg/build/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageclient "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	oauthclient "github.com/openshift/origin/pkg/oauth/generated/internalclientset/typed/oauth/internalversion"
	ocbuildapihelpers "github.com/openshift/origin/pkg/oc/lib/buildapihelpers"
	"github.com/openshift/origin/pkg/oc/lib/routedisplayhelpers"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectclient "github.com/openshift/origin/pkg/project/generated/internalclientset/typed/project/internalversion"
	quotaconvert "github.com/openshift/origin/pkg/quota/apis/quota"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	routev1conversions "github.com/openshift/origin/pkg/route/apis/route/v1"
	routeclient "github.com/openshift/origin/pkg/route/generated/internalclientset/typed/route/internalversion"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityclient "github.com/openshift/origin/pkg/security/generated/internalclientset/typed/security/internalversion"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	templateclient "github.com/openshift/origin/pkg/template/generated/internalclientset/typed/template/internalversion"
	userclient "github.com/openshift/origin/pkg/user/generated/internalclientset/typed/user/internalversion"
)

func describerMap(clientConfig *rest.Config, kclient kubernetes.Interface, host string) map[schema.GroupKind]describe.Describer {
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
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	oauthorizationClient, err := oauthorizationclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	onetworkClient, err := onetworktypedclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	userClient, err := userclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	quotaClient, err := quotaclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	imageClient, err := imageclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	appsClient, err := appstypedclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	buildClient, err := buildv1clienttyped.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	templateClient, err := templateclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	routeClient, err := routeclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	projectClient, err := projectclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	oauthClient, err := oauthclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	securityClient, err := securityclient.NewForConfig(clientConfig)
	if err != nil {
		klog.V(1).Info(err)
	}
	m := map[schema.GroupKind]describe.Describer{oapps.Kind("DeploymentConfig"): &DeploymentConfigDescriber{appsClient, kubeClient, nil}, build.Kind("Build"): &BuildDescriber{buildClient, kclient}, build.Kind("BuildConfig"): &BuildConfigDescriber{buildClient, kclient, host}, image.Kind("Image"): &ImageDescriber{imageClient}, image.Kind("ImageStream"): &ImageStreamDescriber{imageClient}, image.Kind("ImageStreamTag"): &ImageStreamTagDescriber{imageClient}, image.Kind("ImageStreamImage"): &ImageStreamImageDescriber{imageClient}, route.Kind("Route"): &RouteDescriber{routeClient, kclient}, project.Kind("Project"): &ProjectDescriber{projectClient, kclient}, template.Kind("Template"): &TemplateDescriber{templateClient, meta.NewAccessor(), legacyscheme.Scheme, nil}, template.Kind("TemplateInstance"): &TemplateInstanceDescriber{kclient, templateClient, nil}, authorization.Kind("RoleBinding"): &RoleBindingDescriber{oauthorizationClient}, authorization.Kind("Role"): &RoleDescriber{oauthorizationClient}, authorization.Kind("ClusterRoleBinding"): &ClusterRoleBindingDescriber{oauthorizationClient}, authorization.Kind("ClusterRole"): &ClusterRoleDescriber{oauthorizationClient}, authorization.Kind("RoleBindingRestriction"): &RoleBindingRestrictionDescriber{oauthorizationClient}, oauth.Kind("OAuthAccessToken"): &OAuthAccessTokenDescriber{oauthClient}, user.Kind("Identity"): &IdentityDescriber{userClient}, user.Kind("User"): &UserDescriber{userClient}, user.Kind("Group"): &GroupDescriber{userClient}, user.Kind("UserIdentityMapping"): &UserIdentityMappingDescriber{userClient}, quota.Kind("ClusterResourceQuota"): &ClusterQuotaDescriber{quotaClient}, quota.Kind("AppliedClusterResourceQuota"): &AppliedClusterQuotaDescriber{quotaClient}, network.Kind("ClusterNetwork"): &ClusterNetworkDescriber{onetworkClient}, network.Kind("HostSubnet"): &HostSubnetDescriber{onetworkClient}, network.Kind("NetNamespace"): &NetNamespaceDescriber{onetworkClient}, network.Kind("EgressNetworkPolicy"): &EgressNetworkPolicyDescriber{onetworkClient}, security.Kind("SecurityContextConstraints"): &SecurityContextConstraintsDescriber{securityClient}}
	for gk, d := range m {
		m[legacy.Kind(gk.Kind)] = d
	}
	return m
}
func DescriberFor(kind schema.GroupKind, clientConfig *rest.Config, kubeClient kubernetes.Interface, host string) (describe.Describer, bool) {
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
	f, ok := describerMap(clientConfig, kubeClient, host)[kind]
	if ok {
		return f, true
	}
	return nil, false
}

type BuildDescriber struct {
	buildClient	buildv1clienttyped.BuildV1Interface
	kubeClient	kubernetes.Interface
}

func (d *BuildDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.buildClient.Builds(namespace)
	buildObj, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	events, _ := d.kubeClient.CoreV1().Events(namespace).Search(legacyscheme.Scheme, buildObj)
	if events == nil {
		events = &corev1.EventList{}
	}
	if pod, err := d.kubeClient.CoreV1().Pods(namespace).Get(buildutil.GetBuildPodName(buildObj), metav1.GetOptions{}); err == nil {
		if podEvents, _ := d.kubeClient.CoreV1().Events(namespace).Search(legacyscheme.Scheme, pod); podEvents != nil {
			events.Items = append(events.Items, podEvents.Items...)
		}
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, buildObj.ObjectMeta)
		fmt.Fprintln(out, "")
		status := bold(buildObj.Status.Phase)
		if buildObj.Status.Message != "" {
			status += " (" + buildObj.Status.Message + ")"
		}
		formatString(out, "Status", status)
		if buildObj.Status.StartTimestamp != nil && !buildObj.Status.StartTimestamp.IsZero() {
			formatString(out, "Started", buildObj.Status.StartTimestamp.Time.Format(time.RFC1123))
		}
		formatString(out, "Duration", describeBuildDuration(buildObj))
		for _, stage := range buildObj.Status.Stages {
			duration := stage.StartTime.Time.Add(time.Duration(stage.DurationMilliseconds * int64(time.Millisecond))).Round(time.Second).Sub(stage.StartTime.Time.Round(time.Second))
			formatString(out, fmt.Sprintf("  %v", stage.Name), fmt.Sprintf("  %v", duration))
		}
		fmt.Fprintln(out, "")
		if buildObj.Status.Config != nil {
			formatString(out, "Build Config", buildObj.Status.Config.Name)
		}
		formatString(out, "Build Pod", buildutil.GetBuildPodName(buildObj))
		if buildObj.Status.Output.To != nil && len(buildObj.Status.Output.To.ImageDigest) > 0 {
			formatString(out, "Image Digest", buildObj.Status.Output.To.ImageDigest)
		}
		describeCommonSpec(buildObj.Spec.CommonSpec, out)
		describeBuildTriggerCauses(buildObj.Spec.TriggeredBy, out)
		if len(buildObj.Status.LogSnippet) != 0 {
			formatString(out, "Log Tail", buildObj.Status.LogSnippet)
		}
		if settings.ShowEvents {
			versioned.DescribeEvents(events, versioned.NewPrefixWriter(out))
		}
		return nil
	})
}
func describeBuildDuration(build *buildv1.Build) string {
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
	t := metav1.Now().Rfc3339Copy()
	if build.Status.StartTimestamp == nil && build.Status.CompletionTimestamp != nil && (build.Status.Phase == buildv1.BuildPhaseCancelled || build.Status.Phase == buildv1.BuildPhaseFailed || build.Status.Phase == buildv1.BuildPhaseError) {
		return fmt.Sprintf("waited for %s", build.Status.CompletionTimestamp.Rfc3339Copy().Time.Sub(build.CreationTimestamp.Rfc3339Copy().Time))
	} else if build.Status.StartTimestamp == nil && build.Status.Phase != buildv1.BuildPhaseCancelled {
		return fmt.Sprintf("waiting for %v", t.Sub(build.CreationTimestamp.Rfc3339Copy().Time))
	} else if build.Status.StartTimestamp != nil && build.Status.CompletionTimestamp == nil {
		duration := metav1.Now().Rfc3339Copy().Time.Sub(build.Status.StartTimestamp.Rfc3339Copy().Time)
		return fmt.Sprintf("running for %v", duration)
	} else if build.Status.CompletionTimestamp == nil && build.Status.StartTimestamp == nil && build.Status.Phase == buildv1.BuildPhaseCancelled {
		return "<none>"
	}
	duration := build.Status.CompletionTimestamp.Rfc3339Copy().Time.Sub(build.Status.StartTimestamp.Rfc3339Copy().Time)
	return fmt.Sprintf("%v", duration)
}

type BuildConfigDescriber struct {
	buildClient	buildv1clienttyped.BuildV1Interface
	kubeClient	kubernetes.Interface
	host		string
}

func nameAndNamespace(ns, name string) string {
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
	if len(ns) != 0 {
		return fmt.Sprintf("%s/%s", ns, name)
	}
	return name
}
func describeCommonSpec(p buildv1.CommonSpec, out *tabwriter.Writer) {
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
	formatString(out, "\nStrategy", buildapihelpers.StrategyType(p.Strategy))
	noneType := true
	if p.Source.Git != nil {
		noneType = false
		formatString(out, "URL", p.Source.Git.URI)
		if len(p.Source.Git.Ref) > 0 {
			formatString(out, "Ref", p.Source.Git.Ref)
		}
		if len(p.Source.ContextDir) > 0 {
			formatString(out, "ContextDir", p.Source.ContextDir)
		}
		if p.Source.SourceSecret != nil {
			formatString(out, "Source Secret", p.Source.SourceSecret.Name)
		}
		squashGitInfo(p.Revision, out)
	}
	if p.Source.Dockerfile != nil {
		if len(strings.TrimSpace(*p.Source.Dockerfile)) == 0 {
			formatString(out, "Dockerfile", "")
		} else {
			fmt.Fprintf(out, "Dockerfile:\n")
			for _, s := range strings.Split(*p.Source.Dockerfile, "\n") {
				fmt.Fprintf(out, "  %s\n", s)
			}
		}
	}
	switch {
	case p.Strategy.DockerStrategy != nil:
		describeDockerStrategy(p.Strategy.DockerStrategy, out)
	case p.Strategy.SourceStrategy != nil:
		describeSourceStrategy(p.Strategy.SourceStrategy, out)
	case p.Strategy.CustomStrategy != nil:
		describeCustomStrategy(p.Strategy.CustomStrategy, out)
	case p.Strategy.JenkinsPipelineStrategy != nil:
		describeJenkinsPipelineStrategy(p.Strategy.JenkinsPipelineStrategy, out)
	}
	if p.Output.To != nil {
		formatString(out, "Output to", fmt.Sprintf("%s %s", p.Output.To.Kind, nameAndNamespace(p.Output.To.Namespace, p.Output.To.Name)))
	}
	if p.Source.Binary != nil {
		noneType = false
		if len(p.Source.Binary.AsFile) > 0 {
			formatString(out, "Binary", fmt.Sprintf("provided as file %q on build", p.Source.Binary.AsFile))
		} else {
			formatString(out, "Binary", "provided on build")
		}
	}
	if len(p.Source.Secrets) > 0 {
		result := []string{}
		for _, s := range p.Source.Secrets {
			result = append(result, fmt.Sprintf("%s->%s", s.Secret.Name, filepath.Clean(s.DestinationDir)))
		}
		formatString(out, "Build Secrets", strings.Join(result, ","))
	}
	if len(p.Source.ConfigMaps) > 0 {
		result := []string{}
		for _, c := range p.Source.ConfigMaps {
			result = append(result, fmt.Sprintf("%s->%s", c.ConfigMap.Name, filepath.Clean(c.DestinationDir)))
		}
		formatString(out, "Build ConfigMaps", strings.Join(result, ","))
	}
	if len(p.Source.Images) == 1 && len(p.Source.Images[0].Paths) == 1 {
		noneType = false
		imageObj := p.Source.Images[0]
		path := imageObj.Paths[0]
		formatString(out, "Image Source", fmt.Sprintf("copies %s from %s to %s", path.SourcePath, nameAndNamespace(imageObj.From.Namespace, imageObj.From.Name), path.DestinationDir))
	} else {
		for _, image := range p.Source.Images {
			noneType = false
			formatString(out, "Image Source", fmt.Sprintf("%s", nameAndNamespace(image.From.Namespace, image.From.Name)))
			for _, path := range image.Paths {
				fmt.Fprintf(out, "\t- %s -> %s\n", path.SourcePath, path.DestinationDir)
			}
			for _, name := range image.As {
				fmt.Fprintf(out, "\t- as %s\n", name)
			}
		}
	}
	if noneType {
		formatString(out, "Empty Source", "no input source provided")
	}
	describePostCommitHook(p.PostCommit, out)
	if p.Output.PushSecret != nil {
		formatString(out, "Push Secret", p.Output.PushSecret.Name)
	}
	if p.CompletionDeadlineSeconds != nil {
		formatString(out, "Fail Build After", time.Duration(*p.CompletionDeadlineSeconds)*time.Second)
	}
}
func describePostCommitHook(hook buildv1.BuildPostCommitSpec, out *tabwriter.Writer) {
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
	command := hook.Command
	args := hook.Args
	script := hook.Script
	if len(command) == 0 && len(args) == 0 && len(script) == 0 {
		return
	}
	if len(script) != 0 {
		command = []string{"/bin/sh", "-ic"}
		if len(args) > 0 {
			args = append([]string{script, command[0]}, args...)
		} else {
			args = []string{script}
		}
	}
	if len(command) == 0 {
		command = []string{"<image-entrypoint>"}
	}
	all := append(command, args...)
	for i, v := range all {
		all[i] = fmt.Sprintf("%q", v)
	}
	formatString(out, "Post Commit Hook", fmt.Sprintf("[%s]", strings.Join(all, ", ")))
}
func describeSourceStrategy(s *buildv1.SourceBuildStrategy, out *tabwriter.Writer) {
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
	if len(s.From.Name) != 0 {
		formatString(out, "From Image", fmt.Sprintf("%s %s", s.From.Kind, nameAndNamespace(s.From.Namespace, s.From.Name)))
	}
	if len(s.Scripts) != 0 {
		formatString(out, "Scripts", s.Scripts)
	}
	if s.PullSecret != nil {
		formatString(out, "Pull Secret Name", s.PullSecret.Name)
	}
	if s.Incremental != nil && *s.Incremental {
		formatString(out, "Incremental Build", "yes")
	}
	if s.ForcePull {
		formatString(out, "Force Pull", "yes")
	}
}
func describeDockerStrategy(s *buildv1.DockerBuildStrategy, out *tabwriter.Writer) {
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
	if s.From != nil && len(s.From.Name) != 0 {
		formatString(out, "From Image", fmt.Sprintf("%s %s", s.From.Kind, nameAndNamespace(s.From.Namespace, s.From.Name)))
	}
	if len(s.DockerfilePath) != 0 {
		formatString(out, "Dockerfile Path", s.DockerfilePath)
	}
	if s.PullSecret != nil {
		formatString(out, "Pull Secret Name", s.PullSecret.Name)
	}
	if s.NoCache {
		formatString(out, "No Cache", "true")
	}
	if s.ForcePull {
		formatString(out, "Force Pull", "true")
	}
}
func describeCustomStrategy(s *buildv1.CustomBuildStrategy, out *tabwriter.Writer) {
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
	if len(s.From.Name) != 0 {
		formatString(out, "Image Reference", fmt.Sprintf("%s %s", s.From.Kind, nameAndNamespace(s.From.Namespace, s.From.Name)))
	}
	if s.ExposeDockerSocket {
		formatString(out, "Expose Docker Socket", "yes")
	}
	if s.ForcePull {
		formatString(out, "Force Pull", "yes")
	}
	if s.PullSecret != nil {
		formatString(out, "Pull Secret Name", s.PullSecret.Name)
	}
	for i, env := range s.Env {
		if i == 0 {
			formatString(out, "Environment", formatEnv(env))
		} else {
			formatString(out, "", formatEnv(env))
		}
	}
}
func describeJenkinsPipelineStrategy(s *buildv1.JenkinsPipelineBuildStrategy, out *tabwriter.Writer) {
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
	if len(s.JenkinsfilePath) != 0 {
		formatString(out, "Jenkinsfile path", s.JenkinsfilePath)
	}
	if len(s.Jenkinsfile) != 0 {
		fmt.Fprintf(out, "Jenkinsfile contents:\n")
		for _, s := range strings.Split(s.Jenkinsfile, "\n") {
			fmt.Fprintf(out, "  %s\n", s)
		}
	}
	if len(s.Jenkinsfile) == 0 && len(s.JenkinsfilePath) == 0 {
		formatString(out, "Jenkinsfile", "from source repository root")
	}
}
func (d *BuildConfigDescriber) DescribeTriggers(bc *buildv1.BuildConfig, out *tabwriter.Writer) {
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
	describeBuildTriggers(bc.Spec.Triggers, bc.Name, bc.Namespace, out, d)
}
func describeBuildTriggers(triggers []buildv1.BuildTriggerPolicy, name, namespace string, w *tabwriter.Writer, d *BuildConfigDescriber) {
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
	if len(triggers) == 0 {
		formatString(w, "Triggered by", "<none>")
		return
	}
	labels := []string{}
	for _, t := range triggers {
		switch t.Type {
		case buildv1.GitHubWebHookBuildTriggerType, buildv1.GenericWebHookBuildTriggerType, buildv1.GitLabWebHookBuildTriggerType, buildv1.BitbucketWebHookBuildTriggerType:
			continue
		case buildv1.ConfigChangeBuildTriggerType:
			labels = append(labels, "Config")
		case buildv1.ImageChangeBuildTriggerType:
			if t.ImageChange != nil && t.ImageChange.From != nil && len(t.ImageChange.From.Name) > 0 {
				labels = append(labels, fmt.Sprintf("Image(%s %s)", t.ImageChange.From.Kind, t.ImageChange.From.Name))
			} else {
				labels = append(labels, string(t.Type))
			}
		case "":
			labels = append(labels, "<unknown>")
		default:
			labels = append(labels, string(t.Type))
		}
	}
	desc := strings.Join(labels, ", ")
	formatString(w, "Triggered by", desc)
	webHooks := webHooksDescribe(triggers, name, namespace, d.buildClient.RESTClient())
	seenHookTypes := make(map[string]bool)
	for webHookType, webHookDesc := range webHooks {
		fmt.Fprintf(w, "Webhook %s:\n", strings.Title(webHookType))
		for _, trigger := range webHookDesc {
			_, seen := seenHookTypes[webHookType]
			if webHookType != string(buildapi.GenericWebHookBuildTriggerType) && seen {
				continue
			}
			seenHookTypes[webHookType] = true
			fmt.Fprintf(w, "\tURL:\t%s\n", trigger.URL)
			if webHookType == string(buildapi.GenericWebHookBuildTriggerType) && trigger.AllowEnv != nil {
				fmt.Fprintf(w, fmt.Sprintf("\t%s:\t%v\n", "AllowEnv", *trigger.AllowEnv))
			}
		}
	}
}
func (d *BuildConfigDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.buildClient.BuildConfigs(namespace)
	buildConfig, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	buildList, err := d.buildClient.Builds(namespace).List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	buildList.Items = ocbuildapihelpers.FilterBuilds(buildList.Items, ocbuildapihelpers.ByBuildConfigPredicate(name))
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, buildConfig.ObjectMeta)
		if buildConfig.Status.LastVersion == 0 {
			formatString(out, "Latest Version", "Never built")
		} else {
			formatString(out, "Latest Version", strconv.FormatInt(buildConfig.Status.LastVersion, 10))
		}
		describeCommonSpec(buildConfig.Spec.CommonSpec, out)
		formatString(out, "\nBuild Run Policy", string(buildConfig.Spec.RunPolicy))
		d.DescribeTriggers(buildConfig, out)
		if buildConfig.Spec.SuccessfulBuildsHistoryLimit != nil || buildConfig.Spec.FailedBuildsHistoryLimit != nil {
			fmt.Fprintf(out, "Builds History Limit:\n")
			if buildConfig.Spec.SuccessfulBuildsHistoryLimit != nil {
				fmt.Fprintf(out, "\tSuccessful:\t%s\n", strconv.Itoa(int(*buildConfig.Spec.SuccessfulBuildsHistoryLimit)))
			}
			if buildConfig.Spec.FailedBuildsHistoryLimit != nil {
				fmt.Fprintf(out, "\tFailed:\t%s\n", strconv.Itoa(int(*buildConfig.Spec.FailedBuildsHistoryLimit)))
			}
		}
		if len(buildList.Items) > 0 {
			fmt.Fprintf(out, "\nBuild\tStatus\tDuration\tCreation Time\n")
			builds := buildList.Items
			sort.Sort(sort.Reverse(buildapihelpers.BuildSliceByCreationTimestamp(builds)))
			for i, build := range builds {
				fmt.Fprintf(out, "%s \t%s \t%v \t%v\n", build.Name, strings.ToLower(string(build.Status.Phase)), describeBuildDuration(&build), build.CreationTimestamp.Rfc3339Copy().Time)
				if i == 9 {
					break
				}
			}
		}
		if settings.ShowEvents {
			events, _ := d.kubeClient.CoreV1().Events(namespace).Search(legacyscheme.Scheme, buildConfig)
			if events != nil {
				fmt.Fprint(out, "\n")
				versioned.DescribeEvents(events, versioned.NewPrefixWriter(out))
			}
		}
		return nil
	})
}

type OAuthAccessTokenDescriber struct{ client oauthclient.OauthInterface }

func (d *OAuthAccessTokenDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.client.OAuthAccessTokens()
	oAuthAccessToken, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	var timeCreated time.Time = oAuthAccessToken.ObjectMeta.CreationTimestamp.Time
	expires := "never"
	if oAuthAccessToken.ExpiresIn > 0 {
		var timeExpired time.Time = timeCreated.Add(time.Duration(oAuthAccessToken.ExpiresIn) * time.Second)
		expires = formatToHumanDuration(timeExpired.Sub(time.Now()))
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, oAuthAccessToken.ObjectMeta)
		formatString(out, "Scopes", oAuthAccessToken.Scopes)
		formatString(out, "Expires In", expires)
		formatString(out, "User Name", oAuthAccessToken.UserName)
		formatString(out, "User UID", oAuthAccessToken.UserUID)
		formatString(out, "Client Name", oAuthAccessToken.ClientName)
		return nil
	})
}

type ImageDescriber struct{ c imageclient.ImageInterface }

func (d *ImageDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.Images()
	image, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeImage(image, "")
}
func describeImageSignature(s imageapi.ImageSignature, out *tabwriter.Writer) error {
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
	formatString(out, "\tName", s.Name)
	formatString(out, "\tType", s.Type)
	if s.IssuedBy == nil {
		formatString(out, "\tStatus", "Unverified")
	} else {
		formatString(out, "\tStatus", "Verified")
		formatString(out, "\tIssued By", s.IssuedBy.CommonName)
		if len(s.Conditions) > 0 {
			for _, c := range s.Conditions {
				formatString(out, "\t", fmt.Sprintf("Signature is %s (%s on %s)", string(c.Type), c.Message, fmt.Sprintf("%s", c.LastProbeTime)))
			}
		}
	}
	return nil
}
func DescribeImage(image *imageapi.Image, imageName string) (string, error) {
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
	return tabbedString(func(out *tabwriter.Writer) error {
		if len(imageName) > 0 {
			formatString(out, "Image Name", imageName)
		}
		formatString(out, "Docker Image", image.DockerImageReference)
		formatString(out, "Name", image.Name)
		if !image.CreationTimestamp.IsZero() {
			formatTime(out, "Created", image.CreationTimestamp.Time)
		}
		if len(image.Labels) > 0 {
			formatMapStringString(out, "Labels", image.Labels)
		}
		if len(image.Annotations) > 0 {
			formatAnnotations(out, image.ObjectMeta, "")
		}
		switch l := len(image.DockerImageLayers); l {
		case 0:
			formatString(out, "Layer Size", units.HumanSize(float64(image.DockerImageMetadata.Size)))
		case 1:
			formatString(out, "Image Size", units.HumanSize(float64(image.DockerImageMetadata.Size)))
		default:
			formatString(out, "Image Size", fmt.Sprintf("%s in %d layers", units.HumanSize(float64(image.DockerImageMetadata.Size)), len(image.DockerImageLayers)))
			var layers []string
			for _, layer := range image.DockerImageLayers {
				layers = append(layers, fmt.Sprintf("%s\t%s", units.HumanSize(float64(layer.LayerSize)), layer.Name))
			}
			formatString(out, "Layers", strings.Join(layers, "\n"))
		}
		if len(image.Signatures) > 0 {
			for _, s := range image.Signatures {
				formatString(out, "Image Signatures", " ")
				if err := describeImageSignature(s, out); err != nil {
					return err
				}
			}
		}
		formatString(out, "Image Created", fmt.Sprintf("%s ago", formatRelativeTime(image.DockerImageMetadata.Created.Time)))
		formatString(out, "Author", image.DockerImageMetadata.Author)
		formatString(out, "Arch", image.DockerImageMetadata.Architecture)
		describeDockerImage(out, image.DockerImageMetadata.Config)
		return nil
	})
}
func describeDockerImage(out *tabwriter.Writer, image *imageapi.DockerConfig) {
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
	if image == nil {
		return
	}
	hasCommand := false
	if len(image.Entrypoint) > 0 {
		hasCommand = true
		formatString(out, "Entrypoint", strings.Join(image.Entrypoint, " "))
	}
	if len(image.Cmd) > 0 {
		hasCommand = true
		formatString(out, "Command", strings.Join(image.Cmd, " "))
	}
	if !hasCommand {
		formatString(out, "Command", "")
	}
	formatString(out, "Working Dir", image.WorkingDir)
	formatString(out, "User", image.User)
	ports := sets.NewString()
	for k := range image.ExposedPorts {
		ports.Insert(k)
	}
	formatString(out, "Exposes Ports", strings.Join(ports.List(), ", "))
	formatMapStringString(out, "Docker Labels", image.Labels)
	for i, env := range image.Env {
		if i == 0 {
			formatString(out, "Environment", env)
		} else {
			fmt.Fprintf(out, "\t%s\n", env)
		}
	}
	volumes := sets.NewString()
	for k := range image.Volumes {
		volumes.Insert(k)
	}
	for i, volume := range volumes.List() {
		if i == 0 {
			formatString(out, "Volumes", volume)
		} else {
			fmt.Fprintf(out, "\t%s\n", volume)
		}
	}
}

type ImageStreamTagDescriber struct{ c imageclient.ImageInterface }

func (d *ImageStreamTagDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.ImageStreamTags(namespace)
	repo, tag, err := imageapi.ParseImageStreamTagName(name)
	if err != nil {
		return "", err
	}
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	imageStreamTag, err := c.Get(repo+":"+tag, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeImage(&imageStreamTag.Image, imageStreamTag.Image.Name)
}

type ImageStreamImageDescriber struct{ c imageclient.ImageInterface }

func (d *ImageStreamImageDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.ImageStreamImages(namespace)
	imageStreamImage, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeImage(&imageStreamImage.Image, imageStreamImage.Image.Name)
}

type ImageStreamDescriber struct{ ImageClient imageclient.ImageInterface }

func (d *ImageStreamDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.ImageClient.ImageStreams(namespace)
	imageStream, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeImageStream(imageStream)
}
func DescribeImageStream(imageStream *imageapi.ImageStream) (string, error) {
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
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, imageStream.ObjectMeta)
		if len(imageStream.Status.PublicDockerImageRepository) > 0 {
			formatString(out, "Image Repository", imageStream.Status.PublicDockerImageRepository)
		} else {
			formatString(out, "Image Repository", imageStream.Status.DockerImageRepository)
		}
		formatString(out, "Image Lookup", fmt.Sprintf("local=%t", imageStream.Spec.LookupPolicy.Local))
		formatImageStreamTags(out, imageStream)
		return nil
	})
}

type RouteDescriber struct {
	routeClient	routeclient.RouteInterface
	kubeClient	kubernetes.Interface
}
type routeEndpointInfo struct {
	*corev1.Endpoints
	Err	error
}

func (d *RouteDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.routeClient.Routes(namespace)
	route, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	backends := append([]routeapi.RouteTargetReference{route.Spec.To}, route.Spec.AlternateBackends...)
	totalWeight := int32(0)
	endpoints := make(map[string]routeEndpointInfo)
	for _, backend := range backends {
		if backend.Weight != nil {
			totalWeight += *backend.Weight
		}
		ep, endpointsErr := d.kubeClient.CoreV1().Endpoints(namespace).Get(backend.Name, metav1.GetOptions{})
		endpoints[backend.Name] = routeEndpointInfo{ep, endpointsErr}
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		var hostName string
		formatMeta(out, route.ObjectMeta)
		if len(route.Spec.Host) > 0 {
			formatString(out, "Requested Host", route.Spec.Host)
			for _, ingress := range route.Status.Ingress {
				if route.Spec.Host != ingress.Host {
					continue
				}
				hostName = ""
				if len(ingress.RouterCanonicalHostname) > 0 {
					hostName = fmt.Sprintf(" (host %s)", ingress.RouterCanonicalHostname)
				}
				external := routev1.RouteIngress{}
				if err := routev1conversions.Convert_route_RouteIngress_To_v1_RouteIngress(&ingress, &external, nil); err != nil {
					return err
				}
				switch status, condition := routedisplayhelpers.IngressConditionStatus(&external, routev1.RouteAdmitted); status {
				case corev1.ConditionTrue:
					fmt.Fprintf(out, "\t  exposed on router %s%s %s ago\n", ingress.RouterName, hostName, strings.ToLower(formatRelativeTime(condition.LastTransitionTime.Time)))
				case corev1.ConditionFalse:
					fmt.Fprintf(out, "\t  rejected by router %s: %s%s (%s ago)\n", ingress.RouterName, hostName, condition.Reason, strings.ToLower(formatRelativeTime(condition.LastTransitionTime.Time)))
					if len(condition.Message) > 0 {
						fmt.Fprintf(out, "\t    %s\n", condition.Message)
					}
				}
			}
		} else {
			formatString(out, "Requested Host", "<auto>")
		}
		for _, ingress := range route.Status.Ingress {
			if route.Spec.Host == ingress.Host {
				continue
			}
			hostName = ""
			if len(ingress.RouterCanonicalHostname) > 0 {
				hostName = fmt.Sprintf(" (host %s)", ingress.RouterCanonicalHostname)
			}
			external := routev1.RouteIngress{}
			if err := routev1conversions.Convert_route_RouteIngress_To_v1_RouteIngress(&ingress, &external, nil); err != nil {
				return err
			}
			switch status, condition := routedisplayhelpers.IngressConditionStatus(&external, routev1.RouteAdmitted); status {
			case corev1.ConditionTrue:
				fmt.Fprintf(out, "\t%s exposed on router %s %s%s ago\n", ingress.Host, ingress.RouterName, hostName, strings.ToLower(formatRelativeTime(condition.LastTransitionTime.Time)))
			case corev1.ConditionFalse:
				fmt.Fprintf(out, "\trejected by router %s: %s%s (%s ago)\n", ingress.RouterName, hostName, condition.Reason, strings.ToLower(formatRelativeTime(condition.LastTransitionTime.Time)))
				if len(condition.Message) > 0 {
					fmt.Fprintf(out, "\t  %s\n", condition.Message)
				}
			}
		}
		formatString(out, "Path", route.Spec.Path)
		tlsTerm := ""
		insecurePolicy := ""
		if route.Spec.TLS != nil {
			tlsTerm = string(route.Spec.TLS.Termination)
			insecurePolicy = string(route.Spec.TLS.InsecureEdgeTerminationPolicy)
		}
		formatString(out, "TLS Termination", tlsTerm)
		formatString(out, "Insecure Policy", insecurePolicy)
		if route.Spec.Port != nil {
			formatString(out, "Endpoint Port", route.Spec.Port.TargetPort.String())
		} else {
			formatString(out, "Endpoint Port", "<all endpoint ports>")
		}
		for _, backend := range backends {
			fmt.Fprintln(out)
			formatString(out, "Service", backend.Name)
			weight := int32(0)
			if backend.Weight != nil {
				weight = *backend.Weight
			}
			if weight > 0 {
				fmt.Fprintf(out, "Weight:\t%d (%d%%)\n", weight, weight*100/totalWeight)
			} else {
				formatString(out, "Weight", "0")
			}
			info := endpoints[backend.Name]
			if info.Err != nil {
				formatString(out, "Endpoints", fmt.Sprintf("<error: %v>", info.Err))
				continue
			}
			endpoints := info.Endpoints
			if len(endpoints.Subsets) == 0 {
				formatString(out, "Endpoints", "<none>")
				continue
			}
			list := []string{}
			max := 3
			count := 0
			for i := range endpoints.Subsets {
				ss := &endpoints.Subsets[i]
				for p := range ss.Ports {
					for a := range ss.Addresses {
						if len(list) < max {
							list = append(list, fmt.Sprintf("%s:%d", ss.Addresses[a].IP, ss.Ports[p].Port))
						}
						count++
					}
				}
			}
			ends := strings.Join(list, ", ")
			if count > max {
				ends += fmt.Sprintf(" + %d more...", count-max)
			}
			formatString(out, "Endpoints", ends)
		}
		return nil
	})
}

type ProjectDescriber struct {
	projectClient	projectclient.ProjectInterface
	kubeClient	kubernetes.Interface
}

func (d *ProjectDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	projectsClient := d.projectClient.Projects()
	project, err := projectsClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	resourceQuotasClient := d.kubeClient.CoreV1().ResourceQuotas(name)
	resourceQuotaList, err := resourceQuotasClient.List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	limitRangesClient := d.kubeClient.CoreV1().LimitRanges(name)
	limitRangeList, err := limitRangesClient.List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	nodeSelector := ""
	if len(project.ObjectMeta.Annotations) > 0 {
		if ns, ok := project.ObjectMeta.Annotations[projectapi.ProjectNodeSelector]; ok {
			nodeSelector = ns
		}
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, project.ObjectMeta)
		formatString(out, "Display Name", project.Annotations[oapi.OpenShiftDisplayName])
		formatString(out, "Description", project.Annotations[oapi.OpenShiftDescription])
		formatString(out, "Status", project.Status.Phase)
		formatString(out, "Node Selector", nodeSelector)
		if len(resourceQuotaList.Items) == 0 {
			formatString(out, "Quota", "")
		} else {
			fmt.Fprintf(out, "Quota:\n")
			for i := range resourceQuotaList.Items {
				resourceQuota := &resourceQuotaList.Items[i]
				fmt.Fprintf(out, "\tName:\t%s\n", resourceQuota.Name)
				fmt.Fprintf(out, "\tResource\tUsed\tHard\n")
				fmt.Fprintf(out, "\t--------\t----\t----\n")
				resources := []corev1.ResourceName{}
				for resource := range resourceQuota.Status.Hard {
					resources = append(resources, resource)
				}
				sort.Sort(versioned.SortableResourceNames(resources))
				msg := "\t%v\t%v\t%v\n"
				for i := range resources {
					resource := resources[i]
					hardQuantity := resourceQuota.Status.Hard[resource]
					usedQuantity := resourceQuota.Status.Used[resource]
					fmt.Fprintf(out, msg, resource, usedQuantity.String(), hardQuantity.String())
				}
			}
		}
		if len(limitRangeList.Items) == 0 {
			formatString(out, "Resource limits", "")
		} else {
			fmt.Fprintf(out, "Resource limits:\n")
			for i := range limitRangeList.Items {
				limitRange := &limitRangeList.Items[i]
				fmt.Fprintf(out, "\tName:\t%s\n", limitRange.Name)
				fmt.Fprintf(out, "\tType\tResource\tMin\tMax\tDefault Request\tDefault Limit\tMax Limit/Request Ratio\n")
				fmt.Fprintf(out, "\t----\t--------\t---\t---\t---------------\t-------------\t-----------------------\n")
				for i := range limitRange.Spec.Limits {
					item := limitRange.Spec.Limits[i]
					maxResources := item.Max
					minResources := item.Min
					defaultLimitResources := item.Default
					defaultRequestResources := item.DefaultRequest
					ratio := item.MaxLimitRequestRatio
					set := map[corev1.ResourceName]bool{}
					for k := range maxResources {
						set[k] = true
					}
					for k := range minResources {
						set[k] = true
					}
					for k := range defaultLimitResources {
						set[k] = true
					}
					for k := range defaultRequestResources {
						set[k] = true
					}
					for k := range ratio {
						set[k] = true
					}
					for k := range set {
						maxValue := "-"
						minValue := "-"
						defaultLimitValue := "-"
						defaultRequestValue := "-"
						ratioValue := "-"
						maxQuantity, maxQuantityFound := maxResources[k]
						if maxQuantityFound {
							maxValue = maxQuantity.String()
						}
						minQuantity, minQuantityFound := minResources[k]
						if minQuantityFound {
							minValue = minQuantity.String()
						}
						defaultLimitQuantity, defaultLimitQuantityFound := defaultLimitResources[k]
						if defaultLimitQuantityFound {
							defaultLimitValue = defaultLimitQuantity.String()
						}
						defaultRequestQuantity, defaultRequestQuantityFound := defaultRequestResources[k]
						if defaultRequestQuantityFound {
							defaultRequestValue = defaultRequestQuantity.String()
						}
						ratioQuantity, ratioQuantityFound := ratio[k]
						if ratioQuantityFound {
							ratioValue = ratioQuantity.String()
						}
						msg := "\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n"
						fmt.Fprintf(out, msg, item.Type, k, minValue, maxValue, defaultRequestValue, defaultLimitValue, ratioValue)
					}
				}
			}
		}
		return nil
	})
}

type TemplateDescriber struct {
	templateClient	templateclient.TemplateInterface
	meta.MetadataAccessor
	runtime.ObjectTyper
	describe.ObjectDescriber
}

func (d *TemplateDescriber) DescribeMessage(msg string, out *tabwriter.Writer) {
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
	if len(msg) == 0 {
		msg = "<none>"
	}
	formatString(out, "Message", msg)
}
func (d *TemplateDescriber) DescribeParameters(params []templateapi.Parameter, out *tabwriter.Writer) {
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
	formatString(out, "Parameters", " ")
	indent := "    "
	for _, p := range params {
		formatString(out, indent+"Name", p.Name)
		if len(p.DisplayName) > 0 {
			formatString(out, indent+"Display Name", p.DisplayName)
		}
		if len(p.Description) > 0 {
			formatString(out, indent+"Description", p.Description)
		}
		formatString(out, indent+"Required", p.Required)
		if len(p.Generate) == 0 {
			formatString(out, indent+"Value", p.Value)
			out.Write([]byte("\n"))
			continue
		}
		if len(p.Value) > 0 {
			formatString(out, indent+"Value", p.Value)
			formatString(out, indent+"Generated (ignored)", p.Generate)
			formatString(out, indent+"From", p.From)
		} else {
			formatString(out, indent+"Generated", p.Generate)
			formatString(out, indent+"From", p.From)
		}
		out.Write([]byte("\n"))
	}
}
func (d *TemplateDescriber) describeObjects(objects []runtime.Object, out *tabwriter.Writer) {
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
	formatString(out, "Objects", " ")
	indent := "    "
	for _, obj := range objects {
		if d.ObjectDescriber != nil {
			output, err := d.DescribeObject(obj)
			if err != nil {
				fmt.Fprintf(out, "error: %v\n", err)
				continue
			}
			fmt.Fprint(out, output)
			fmt.Fprint(out, "\n")
			continue
		}
		name, _ := d.MetadataAccessor.Name(obj)
		groupKind := "<unknown>"
		if gvk, _, err := d.ObjectTyper.ObjectKinds(obj); err == nil {
			gk := gvk[0].GroupKind()
			groupKind = gk.String()
		} else {
			if unstructured, ok := obj.(*unstructured.Unstructured); ok {
				gvk := unstructured.GroupVersionKind()
				gk := gvk.GroupKind()
				groupKind = gk.String()
			}
		}
		fmt.Fprintf(out, fmt.Sprintf("%s%s\t%s\n", indent, groupKind, name))
	}
}
func (d *TemplateDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.templateClient.Templates(namespace)
	template, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return d.DescribeTemplate(template)
}
func (d *TemplateDescriber) DescribeTemplate(template *templateapi.Template) (string, error) {
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
	_ = runtime.DecodeList(template.Objects, unstructured.UnstructuredJSONScheme)
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, template.ObjectMeta)
		out.Write([]byte("\n"))
		out.Flush()
		d.DescribeParameters(template.Parameters, out)
		out.Write([]byte("\n"))
		formatString(out, "Object Labels", formatLabels(template.ObjectLabels))
		out.Write([]byte("\n"))
		d.DescribeMessage(template.Message, out)
		out.Write([]byte("\n"))
		out.Flush()
		d.describeObjects(template.Objects, out)
		return nil
	})
}

type TemplateInstanceDescriber struct {
	kubeClient	kubernetes.Interface
	templateClient	templateclient.TemplateInterface
	describe.ObjectDescriber
}

func (d *TemplateInstanceDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.templateClient.TemplateInstances(namespace)
	templateInstance, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return d.DescribeTemplateInstance(templateInstance, namespace, settings)
}
func (d *TemplateInstanceDescriber) DescribeTemplateInstance(templateInstance *templateapi.TemplateInstance, namespace string, settings describe.DescriberSettings) (string, error) {
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
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, templateInstance.ObjectMeta)
		out.Write([]byte("\n"))
		out.Flush()
		d.DescribeConditions(templateInstance.Status.Conditions, out)
		out.Write([]byte("\n"))
		out.Flush()
		d.DescribeObjects(templateInstance.Status.Objects, out)
		out.Write([]byte("\n"))
		out.Flush()
		d.DescribeParameters(templateInstance.Spec.Template, namespace, templateInstance.Spec.Secret.Name, out)
		out.Write([]byte("\n"))
		out.Flush()
		return nil
	})
}
func (d *TemplateInstanceDescriber) DescribeConditions(conditions []templateapi.TemplateInstanceCondition, out *tabwriter.Writer) {
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
	formatString(out, "Conditions", " ")
	indent := "    "
	for _, c := range conditions {
		formatString(out, indent+"Type", c.Type)
		formatString(out, indent+"Status", c.Status)
		formatString(out, indent+"LastTransitionTime", c.LastTransitionTime)
		formatString(out, indent+"Reason", c.Reason)
		formatString(out, indent+"Message", c.Message)
		out.Write([]byte("\n"))
	}
}
func (d *TemplateInstanceDescriber) DescribeObjects(objects []templateapi.TemplateInstanceObject, out *tabwriter.Writer) {
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
	formatString(out, "Objects", " ")
	indent := "    "
	for _, o := range objects {
		formatString(out, indent+o.Ref.Kind, fmt.Sprintf("%s/%s", o.Ref.Namespace, o.Ref.Name))
	}
}
func (d *TemplateInstanceDescriber) DescribeParameters(template templateapi.Template, namespace, name string, out *tabwriter.Writer) {
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
	secret, err := d.kubeClient.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})
	formatString(out, "Parameters", " ")
	if kerrs.IsForbidden(err) || kerrs.IsUnauthorized(err) {
		fmt.Fprintf(out, "Unable to access parameters, insufficient permissions.")
		return
	} else if kerrs.IsNotFound(err) {
		fmt.Fprintf(out, "Unable to access parameters, secret not found: %s", secret.Name)
		return
	} else if err != nil {
		fmt.Fprintf(out, "Unknown error occurred, please rerun with loglevel > 4 for more information")
		klog.V(4).Infof("%v", err)
		return
	}
	indent := "    "
	if len(template.Parameters) == 0 {
		fmt.Fprintf(out, indent+"No parameters found.")
	} else {
		for _, p := range template.Parameters {
			if val, ok := secret.Data[p.Name]; ok {
				formatString(out, indent+p.Name, fmt.Sprintf("%d bytes", len(val)))
			}
		}
	}
}

type IdentityDescriber struct{ c userclient.UserInterface }

func (d *IdentityDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	userClient := d.c.Users()
	identityClient := d.c.Identities()
	identity, err := identityClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, identity.ObjectMeta)
		if len(identity.User.Name) == 0 {
			formatString(out, "User Name", identity.User.Name)
			formatString(out, "User UID", identity.User.UID)
		} else {
			resolvedUser, err := userClient.Get(identity.User.Name, metav1.GetOptions{})
			nameValue := identity.User.Name
			uidValue := string(identity.User.UID)
			if kerrs.IsNotFound(err) {
				nameValue += fmt.Sprintf(" (Error: User does not exist)")
			} else if err != nil {
				nameValue += fmt.Sprintf(" (Error: User lookup failed)")
			} else {
				if !sets.NewString(resolvedUser.Identities...).Has(name) {
					nameValue += fmt.Sprintf(" (Error: User identities do not include %s)", name)
				}
				if resolvedUser.UID != identity.User.UID {
					uidValue += fmt.Sprintf(" (Error: Actual user UID is %s)", string(resolvedUser.UID))
				}
			}
			formatString(out, "User Name", nameValue)
			formatString(out, "User UID", uidValue)
		}
		return nil
	})
}

type UserIdentityMappingDescriber struct{ c userclient.UserInterface }

func (d *UserIdentityMappingDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.UserIdentityMappings()
	mapping, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, mapping.ObjectMeta)
		formatString(out, "Identity", mapping.Identity.Name)
		formatString(out, "User Name", mapping.User.Name)
		formatString(out, "User UID", mapping.User.UID)
		return nil
	})
}

type UserDescriber struct{ c userclient.UserInterface }

func (d *UserDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	userClient := d.c.Users()
	identityClient := d.c.Identities()
	user, err := userClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, user.ObjectMeta)
		if len(user.FullName) > 0 {
			formatString(out, "Full Name", user.FullName)
		}
		if len(user.Identities) == 0 {
			formatString(out, "Identities", "<none>")
		} else {
			for i, identity := range user.Identities {
				resolvedIdentity, err := identityClient.Get(identity, metav1.GetOptions{})
				value := identity
				if kerrs.IsNotFound(err) {
					value += fmt.Sprintf(" (Error: Identity does not exist)")
				} else if err != nil {
					value += fmt.Sprintf(" (Error: Identity lookup failed)")
				} else if resolvedIdentity.User.Name != name {
					value += fmt.Sprintf(" (Error: Identity maps to user name '%s')", resolvedIdentity.User.Name)
				} else if resolvedIdentity.User.UID != user.UID {
					value += fmt.Sprintf(" (Error: Identity maps to user UID '%s')", resolvedIdentity.User.UID)
				}
				if i == 0 {
					formatString(out, "Identities", value)
				} else {
					fmt.Fprintf(out, "           \t%s\n", value)
				}
			}
		}
		return nil
	})
}

type GroupDescriber struct{ c userclient.UserInterface }

func (d *GroupDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	group, err := d.c.Groups().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, group.ObjectMeta)
		if len(group.Users) == 0 {
			formatString(out, "Users", "<none>")
		} else {
			for i, user := range group.Users {
				if i == 0 {
					formatString(out, "Users", user)
				} else {
					fmt.Fprintf(out, "           \t%s\n", user)
				}
			}
		}
		return nil
	})
}

const PolicyRuleHeadings = "Verbs\tNon-Resource URLs\tResource Names\tAPI Groups\tResources"

func DescribePolicyRule(out *tabwriter.Writer, rule authorizationapi.PolicyRule, indent string) {
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
	if rule.AttributeRestrictions != nil {
		return
	}
	fmt.Fprintf(out, indent+"%v\t%v\t%v\t%v\t%v\n", rule.Verbs.List(), rule.NonResourceURLs.List(), rule.ResourceNames.List(), rule.APIGroups, rule.Resources.List())
}

type RoleDescriber struct {
	c oauthorizationclient.AuthorizationInterface
}

func (d *RoleDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.Roles(namespace)
	role, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeRole(role)
}
func DescribeRole(role *authorizationapi.Role) (string, error) {
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
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, role.ObjectMeta)
		fmt.Fprint(out, PolicyRuleHeadings+"\n")
		for _, rule := range role.Rules {
			DescribePolicyRule(out, rule, "")
		}
		return nil
	})
}

type RoleBindingDescriber struct {
	c oauthorizationclient.AuthorizationInterface
}

func (d *RoleBindingDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.RoleBindings(namespace)
	roleBinding, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	var role *authorizationapi.Role
	if len(roleBinding.RoleRef.Namespace) == 0 {
		var clusterRole *authorizationapi.ClusterRole
		clusterRole, err = d.c.ClusterRoles().Get(roleBinding.RoleRef.Name, metav1.GetOptions{})
		role = authorizationapi.ToRole(clusterRole)
	} else {
		role, err = d.c.Roles(roleBinding.RoleRef.Namespace).Get(roleBinding.RoleRef.Name, metav1.GetOptions{})
	}
	return DescribeRoleBinding(roleBinding, role, err)
}
func DescribeRoleBinding(roleBinding *authorizationapi.RoleBinding, role *authorizationapi.Role, err error) (string, error) {
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
	users, groups, sas, others := authorizationapi.SubjectsStrings(roleBinding.Namespace, roleBinding.Subjects)
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, roleBinding.ObjectMeta)
		formatString(out, "Role", roleBinding.RoleRef.Namespace+"/"+roleBinding.RoleRef.Name)
		formatString(out, "Users", strings.Join(users, ", "))
		formatString(out, "Groups", strings.Join(groups, ", "))
		formatString(out, "ServiceAccounts", strings.Join(sas, ", "))
		formatString(out, "Subjects", strings.Join(others, ", "))
		switch {
		case err != nil:
			formatString(out, "Policy Rules", fmt.Sprintf("error: %v", err))
		case role != nil:
			fmt.Fprint(out, PolicyRuleHeadings+"\n")
			for _, rule := range role.Rules {
				DescribePolicyRule(out, rule, "")
			}
		default:
			formatString(out, "Policy Rules", "<none>")
		}
		return nil
	})
}

type ClusterRoleDescriber struct {
	c oauthorizationclient.AuthorizationInterface
}

func (d *ClusterRoleDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.ClusterRoles()
	role, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeRole(authorizationapi.ToRole(role))
}

type ClusterRoleBindingDescriber struct {
	c oauthorizationclient.AuthorizationInterface
}

func (d *ClusterRoleBindingDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.ClusterRoleBindings()
	roleBinding, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	role, err := d.c.ClusterRoles().Get(roleBinding.RoleRef.Name, metav1.GetOptions{})
	return DescribeRoleBinding(authorizationapi.ToRoleBinding(roleBinding), authorizationapi.ToRole(role), err)
}
func describeBuildTriggerCauses(causes []buildv1.BuildTriggerCause, out *tabwriter.Writer) {
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
	if causes == nil {
		formatString(out, "\nBuild trigger cause", "<unknown>")
	}
	for _, cause := range causes {
		formatString(out, "\nBuild trigger cause", cause.Message)
		switch {
		case cause.GitHubWebHook != nil:
			squashGitInfo(cause.GitHubWebHook.Revision, out)
			formatString(out, "Secret", cause.GitHubWebHook.Secret)
		case cause.GitLabWebHook != nil:
			squashGitInfo(cause.GitLabWebHook.Revision, out)
			formatString(out, "Secret", cause.GitLabWebHook.Secret)
		case cause.BitbucketWebHook != nil:
			squashGitInfo(cause.BitbucketWebHook.Revision, out)
			formatString(out, "Secret", cause.BitbucketWebHook.Secret)
		case cause.GenericWebHook != nil:
			squashGitInfo(cause.GenericWebHook.Revision, out)
			formatString(out, "Secret", cause.GenericWebHook.Secret)
		case cause.ImageChangeBuild != nil:
			formatString(out, "Image ID", cause.ImageChangeBuild.ImageID)
			formatString(out, "Image Name/Kind", fmt.Sprintf("%s / %s", cause.ImageChangeBuild.FromRef.Name, cause.ImageChangeBuild.FromRef.Kind))
		}
	}
	fmt.Fprintf(out, "\n")
}
func squashGitInfo(sourceRevision *buildv1.SourceRevision, out *tabwriter.Writer) {
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
	if sourceRevision != nil && sourceRevision.Git != nil {
		rev := sourceRevision.Git
		var commit string
		if len(rev.Commit) > 7 {
			commit = rev.Commit[:7]
		} else {
			commit = rev.Commit
		}
		formatString(out, "Commit", fmt.Sprintf("%s (%s)", commit, rev.Message))
		hasAuthor := len(rev.Author.Name) != 0
		hasCommitter := len(rev.Committer.Name) != 0
		if hasAuthor && hasCommitter {
			if rev.Author.Name == rev.Committer.Name {
				formatString(out, "Author/Committer", rev.Author.Name)
			} else {
				formatString(out, "Author/Committer", fmt.Sprintf("%s / %s", rev.Author.Name, rev.Committer.Name))
			}
		} else if hasAuthor {
			formatString(out, "Author", rev.Author.Name)
		} else if hasCommitter {
			formatString(out, "Committer", rev.Committer.Name)
		}
	}
}

type ClusterQuotaDescriber struct{ c quotaclient.QuotaV1Interface }

func (d *ClusterQuotaDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	quota, err := d.c.ClusterResourceQuotas().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeClusterQuota(quota)
}
func DescribeClusterQuota(quota *quotav1.ClusterResourceQuota) (string, error) {
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
	labelSelector, err := metav1.LabelSelectorAsSelector(quota.Spec.Selector.LabelSelector)
	if err != nil {
		return "", err
	}
	nsSelector := make([]interface{}, 0, len(quota.Status.Namespaces))
	for _, nsQuota := range quota.Status.Namespaces {
		nsSelector = append(nsSelector, nsQuota.Namespace)
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, quota.ObjectMeta)
		fmt.Fprintf(out, "Namespace Selector: %q\n", nsSelector)
		fmt.Fprintf(out, "Label Selector: %s\n", labelSelector)
		fmt.Fprintf(out, "AnnotationSelector: %s\n", quota.Spec.Selector.AnnotationSelector)
		if len(quota.Spec.Quota.Scopes) > 0 {
			scopes := []string{}
			for _, scope := range quota.Spec.Quota.Scopes {
				scopes = append(scopes, string(scope))
			}
			sort.Strings(scopes)
			fmt.Fprintf(out, "Scopes:\t%s\n", strings.Join(scopes, ", "))
		}
		fmt.Fprintf(out, "Resource\tUsed\tHard\n")
		fmt.Fprintf(out, "--------\t----\t----\n")
		resources := []corev1.ResourceName{}
		for resource := range quota.Status.Total.Hard {
			resources = append(resources, resource)
		}
		sort.Sort(versioned.SortableResourceNames(resources))
		msg := "%v\t%v\t%v\n"
		for i := range resources {
			resource := resources[i]
			hardQuantity := quota.Status.Total.Hard[resource]
			usedQuantity := quota.Status.Total.Used[resource]
			fmt.Fprintf(out, msg, resource, usedQuantity.String(), hardQuantity.String())
		}
		return nil
	})
}

type AppliedClusterQuotaDescriber struct{ c quotaclient.QuotaV1Interface }

func (d *AppliedClusterQuotaDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	quota, err := d.c.AppliedClusterResourceQuotas(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return DescribeClusterQuota(quotaconvert.ConvertV1AppliedClusterResourceQuotaToV1ClusterResourceQuota(quota))
}

type ClusterNetworkDescriber struct {
	c onetworktypedclient.NetworkV1Interface
}

func (d *ClusterNetworkDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	cn, err := d.c.ClusterNetworks().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, cn.ObjectMeta)
		formatString(out, "Service Network", cn.ServiceNetwork)
		formatString(out, "Plugin Name", cn.PluginName)
		fmt.Fprintf(out, "ClusterNetworks:\n")
		fmt.Fprintf(out, "CIDR\tHost Subnet Length\n")
		fmt.Fprintf(out, "----\t------------------\n")
		for _, clusterNetwork := range cn.ClusterNetworks {
			fmt.Fprintf(out, "%s\t%d\n", clusterNetwork.CIDR, clusterNetwork.HostSubnetLength)
		}
		return nil
	})
}

type HostSubnetDescriber struct {
	c onetworktypedclient.NetworkV1Interface
}

func (d *HostSubnetDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	hs, err := d.c.HostSubnets().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, hs.ObjectMeta)
		formatString(out, "Node", hs.Host)
		formatString(out, "Node IP", hs.HostIP)
		formatString(out, "Pod Subnet", hs.Subnet)
		formatString(out, "Egress CIDRs", strings.Join(hs.EgressCIDRs, ", "))
		formatString(out, "Egress IPs", strings.Join(hs.EgressIPs, ", "))
		return nil
	})
}

type NetNamespaceDescriber struct {
	c onetworktypedclient.NetworkV1Interface
}

func (d *NetNamespaceDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	netns, err := d.c.NetNamespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, netns.ObjectMeta)
		formatString(out, "Name", netns.NetName)
		formatString(out, "ID", netns.NetID)
		formatString(out, "Egress IPs", strings.Join(netns.EgressIPs, ", "))
		return nil
	})
}

type EgressNetworkPolicyDescriber struct {
	c onetworktypedclient.NetworkV1Interface
}

func (d *EgressNetworkPolicyDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	c := d.c.EgressNetworkPolicies(namespace)
	policy, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, policy.ObjectMeta)
		for _, rule := range policy.Spec.Egress {
			if len(rule.To.CIDRSelector) > 0 {
				fmt.Fprintf(out, "Rule:\t%s to %s\n", rule.Type, rule.To.CIDRSelector)
			} else {
				fmt.Fprintf(out, "Rule:\t%s to %s\n", rule.Type, rule.To.DNSName)
			}
		}
		return nil
	})
}

type RoleBindingRestrictionDescriber struct {
	c oauthorizationclient.AuthorizationInterface
}

func (d *RoleBindingRestrictionDescriber) Describe(namespace, name string, settings describe.DescriberSettings) (string, error) {
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
	rbr, err := d.c.RoleBindingRestrictions(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return tabbedString(func(out *tabwriter.Writer) error {
		formatMeta(out, rbr.ObjectMeta)
		subjectType := roleBindingRestrictionType(rbr)
		if subjectType == "" {
			subjectType = "<none>"
		}
		formatString(out, "Subject type", subjectType)
		var labelSelectors []metav1.LabelSelector
		switch {
		case rbr.Spec.UserRestriction != nil:
			formatString(out, "Users", strings.Join(rbr.Spec.UserRestriction.Users, ", "))
			formatString(out, "Users in groups", strings.Join(rbr.Spec.UserRestriction.Groups, ", "))
			labelSelectors = rbr.Spec.UserRestriction.Selectors
		case rbr.Spec.GroupRestriction != nil:
			formatString(out, "Groups", strings.Join(rbr.Spec.GroupRestriction.Groups, ", "))
			labelSelectors = rbr.Spec.GroupRestriction.Selectors
		case rbr.Spec.ServiceAccountRestriction != nil:
			serviceaccounts := []string{}
			for _, sa := range rbr.Spec.ServiceAccountRestriction.ServiceAccounts {
				serviceaccounts = append(serviceaccounts, sa.Name)
			}
			formatString(out, "ServiceAccounts", strings.Join(serviceaccounts, ", "))
			formatString(out, "Namespaces", strings.Join(rbr.Spec.ServiceAccountRestriction.Namespaces, ", "))
		}
		if rbr.Spec.UserRestriction != nil || rbr.Spec.GroupRestriction != nil {
			if len(labelSelectors) == 0 {
				formatString(out, "Label selectors", "")
			} else {
				fmt.Fprintf(out, "Label selectors:\n")
				for _, labelSelector := range labelSelectors {
					selector, err := metav1.LabelSelectorAsSelector(&labelSelector)
					if err != nil {
						return err
					}
					fmt.Fprintf(out, "\t%s\n", selector)
				}
			}
		}
		return nil
	})
}

type SecurityContextConstraintsDescriber struct {
	c securityclient.SecurityContextConstraintsGetter
}

func (d *SecurityContextConstraintsDescriber) Describe(namespace, name string, s describe.DescriberSettings) (string, error) {
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
	scc, err := d.c.SecurityContextConstraints().Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return describeSecurityContextConstraints(scc)
}
func describeSecurityContextConstraints(scc *securityapi.SecurityContextConstraints) (string, error) {
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
	return tabbedString(func(out *tabwriter.Writer) error {
		fmt.Fprintf(out, "Name:\t%s\n", scc.Name)
		priority := ""
		if scc.Priority != nil {
			priority = fmt.Sprintf("%d", *scc.Priority)
		}
		fmt.Fprintf(out, "Priority:\t%s\n", stringOrNone(priority))
		fmt.Fprintf(out, "Access:\t\n")
		fmt.Fprintf(out, "  Users:\t%s\n", stringOrNone(strings.Join(scc.Users, ",")))
		fmt.Fprintf(out, "  Groups:\t%s\n", stringOrNone(strings.Join(scc.Groups, ",")))
		fmt.Fprintf(out, "Settings:\t\n")
		fmt.Fprintf(out, "  Allow Privileged:\t%t\n", scc.AllowPrivilegedContainer)
		fmt.Fprintf(out, "  Allow Privilege Escalation:\t%v\n", scc.AllowPrivilegeEscalation)
		fmt.Fprintf(out, "  Default Add Capabilities:\t%s\n", capsToString(scc.DefaultAddCapabilities))
		fmt.Fprintf(out, "  Required Drop Capabilities:\t%s\n", capsToString(scc.RequiredDropCapabilities))
		fmt.Fprintf(out, "  Allowed Capabilities:\t%s\n", capsToString(scc.AllowedCapabilities))
		fmt.Fprintf(out, "  Allowed Seccomp Profiles:\t%s\n", stringOrNone(strings.Join(scc.SeccompProfiles, ",")))
		fmt.Fprintf(out, "  Allowed Volume Types:\t%s\n", fsTypeToString(scc.Volumes))
		fmt.Fprintf(out, "  Allowed Flexvolumes:\t%s\n", flexVolumesToString(scc.AllowedFlexVolumes))
		fmt.Fprintf(out, "  Allowed Unsafe Sysctls:\t%s\n", sysctlsToString(scc.AllowedUnsafeSysctls))
		fmt.Fprintf(out, "  Forbidden Sysctls:\t%s\n", sysctlsToString(scc.ForbiddenSysctls))
		fmt.Fprintf(out, "  Allow Host Network:\t%t\n", scc.AllowHostNetwork)
		fmt.Fprintf(out, "  Allow Host Ports:\t%t\n", scc.AllowHostPorts)
		fmt.Fprintf(out, "  Allow Host PID:\t%t\n", scc.AllowHostPID)
		fmt.Fprintf(out, "  Allow Host IPC:\t%t\n", scc.AllowHostIPC)
		fmt.Fprintf(out, "  Read Only Root Filesystem:\t%t\n", scc.ReadOnlyRootFilesystem)
		fmt.Fprintf(out, "  Run As User Strategy: %s\t\n", string(scc.RunAsUser.Type))
		uid := ""
		if scc.RunAsUser.UID != nil {
			uid = strconv.FormatInt(*scc.RunAsUser.UID, 10)
		}
		fmt.Fprintf(out, "    UID:\t%s\n", stringOrNone(uid))
		uidRangeMin := ""
		if scc.RunAsUser.UIDRangeMin != nil {
			uidRangeMin = strconv.FormatInt(*scc.RunAsUser.UIDRangeMin, 10)
		}
		fmt.Fprintf(out, "    UID Range Min:\t%s\n", stringOrNone(uidRangeMin))
		uidRangeMax := ""
		if scc.RunAsUser.UIDRangeMax != nil {
			uidRangeMax = strconv.FormatInt(*scc.RunAsUser.UIDRangeMax, 10)
		}
		fmt.Fprintf(out, "    UID Range Max:\t%s\n", stringOrNone(uidRangeMax))
		fmt.Fprintf(out, "  SELinux Context Strategy: %s\t\n", string(scc.SELinuxContext.Type))
		var user, role, seLinuxType, level string
		if scc.SELinuxContext.SELinuxOptions != nil {
			user = scc.SELinuxContext.SELinuxOptions.User
			role = scc.SELinuxContext.SELinuxOptions.Role
			seLinuxType = scc.SELinuxContext.SELinuxOptions.Type
			level = scc.SELinuxContext.SELinuxOptions.Level
		}
		fmt.Fprintf(out, "    User:\t%s\n", stringOrNone(user))
		fmt.Fprintf(out, "    Role:\t%s\n", stringOrNone(role))
		fmt.Fprintf(out, "    Type:\t%s\n", stringOrNone(seLinuxType))
		fmt.Fprintf(out, "    Level:\t%s\n", stringOrNone(level))
		fmt.Fprintf(out, "  FSGroup Strategy: %s\t\n", string(scc.FSGroup.Type))
		fmt.Fprintf(out, "    Ranges:\t%s\n", idRangeToString(scc.FSGroup.Ranges))
		fmt.Fprintf(out, "  Supplemental Groups Strategy: %s\t\n", string(scc.SupplementalGroups.Type))
		fmt.Fprintf(out, "    Ranges:\t%s\n", idRangeToString(scc.SupplementalGroups.Ranges))
		return nil
	})
}
func stringOrNone(s string) string {
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
	return stringOrDefaultValue(s, "<none>")
}
func stringOrDefaultValue(s, defaultValue string) string {
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
	if len(s) > 0 {
		return s
	}
	return defaultValue
}
func fsTypeToString(volumes []securityapi.FSType) string {
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
	strVolumes := []string{}
	for _, v := range volumes {
		strVolumes = append(strVolumes, string(v))
	}
	return stringOrNone(strings.Join(strVolumes, ","))
}
func flexVolumesToString(flexVolumes []securityapi.AllowedFlexVolume) string {
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
	volumes := []string{}
	for _, flexVolume := range flexVolumes {
		volumes = append(volumes, "driver="+flexVolume.Driver)
	}
	return stringOrDefaultValue(strings.Join(volumes, ","), "<all>")
}
func sysctlsToString(sysctls []string) string {
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
	return stringOrNone(strings.Join(sysctls, ","))
}
func idRangeToString(ranges []securityapi.IDRange) string {
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
	formattedString := ""
	if ranges != nil {
		strRanges := []string{}
		for _, r := range ranges {
			strRanges = append(strRanges, fmt.Sprintf("%d-%d", r.Min, r.Max))
		}
		formattedString = strings.Join(strRanges, ",")
	}
	return stringOrNone(formattedString)
}
func capsToString(caps []kapi.Capability) string {
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
	formattedString := ""
	if caps != nil {
		strCaps := []string{}
		for _, c := range caps {
			strCaps = append(strCaps, string(c))
		}
		formattedString = strings.Join(strCaps, ",")
	}
	return stringOrNone(formattedString)
}
