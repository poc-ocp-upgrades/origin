package generator

import (
	"context"
	"fmt"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	buildv1clienttyped "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1clienttyped "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	buildutil "github.com/openshift/origin/pkg/build/util"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	imageutil "github.com/openshift/origin/pkg/image/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/credentialprovider"
	credentialprovidersecrets "k8s.io/kubernetes/pkg/credentialprovider/secrets"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const conflictRetries = 3

type BuildGenerator struct {
	Client          GeneratorClient
	ServiceAccounts corev1client.ServiceAccountsGetter
	Secrets         corev1client.SecretsGetter
}
type GeneratorClient interface {
	GetBuildConfig(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.BuildConfig, error)
	UpdateBuildConfig(ctx context.Context, buildConfig *buildv1.BuildConfig) error
	GetBuild(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.Build, error)
	CreateBuild(ctx context.Context, build *buildv1.Build) error
	UpdateBuild(ctx context.Context, build *buildv1.Build) error
	GetImageStream(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStream, error)
	GetImageStreamImage(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamImage, error)
	GetImageStreamTag(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamTag, error)
}
type Client struct {
	BuildConfigs      buildv1clienttyped.BuildConfigsGetter
	Builds            buildv1clienttyped.BuildsGetter
	ImageStreams      imagev1clienttyped.ImageStreamsGetter
	ImageStreamImages imagev1clienttyped.ImageStreamImagesGetter
	ImageStreamTags   imagev1clienttyped.ImageStreamTagsGetter
}

func (c Client) GetBuildConfig(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.BuildConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.BuildConfigs.BuildConfigs(apirequest.NamespaceValue(ctx)).Get(name, *options)
}
func (c Client) UpdateBuildConfig(ctx context.Context, buildConfig *buildv1.BuildConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.BuildConfigs.BuildConfigs(apirequest.NamespaceValue(ctx)).Update(buildConfig)
	return err
}
func (c Client) GetBuild(ctx context.Context, name string, options *metav1.GetOptions) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Builds.Builds(apirequest.NamespaceValue(ctx)).Get(name, *options)
}
func (c Client) CreateBuild(ctx context.Context, build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Builds.Builds(apirequest.NamespaceValue(ctx)).Create(build)
	return err
}
func (c Client) UpdateBuild(ctx context.Context, build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Builds.Builds(apirequest.NamespaceValue(ctx)).Update(build)
	return err
}
func (c Client) GetImageStream(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ImageStreams.ImageStreams(apirequest.NamespaceValue(ctx)).Get(name, *options)
}
func (c Client) GetImageStreamImage(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamImage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ImageStreamImages.ImageStreamImages(apirequest.NamespaceValue(ctx)).Get(name, *options)
}
func (c Client) GetImageStreamTag(ctx context.Context, name string, options *metav1.GetOptions) (*imagev1.ImageStreamTag, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ImageStreamTags.ImageStreamTags(apirequest.NamespaceValue(ctx)).Get(name, *options)
}
func fetchServiceAccountSecrets(secrets corev1client.SecretsGetter, serviceAccounts corev1client.ServiceAccountsGetter, namespace, serviceAccount string) ([]corev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result []corev1.Secret
	sa, err := serviceAccounts.ServiceAccounts(namespace).Get(serviceAccount, metav1.GetOptions{})
	if err != nil {
		return result, fmt.Errorf("error getting push/pull secrets for service account %s/%s: %v", namespace, serviceAccount, err)
	}
	for _, ref := range sa.Secrets {
		secret, err := secrets.Secrets(namespace).Get(ref.Name, metav1.GetOptions{})
		if err != nil {
			continue
		}
		result = append(result, *secret)
	}
	return result, nil
}
func findImageChangeTrigger(bc *buildv1.BuildConfig, ref *corev1.ObjectReference) *buildv1.ImageChangeTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ref == nil {
		return nil
	}
	for _, trigger := range bc.Spec.Triggers {
		if trigger.Type != buildv1.ImageChangeBuildTriggerType {
			continue
		}
		imageChange := trigger.ImageChange
		triggerRef := imageChange.From
		if triggerRef == nil {
			triggerRef = buildutil.GetInputReference(bc.Spec.Strategy)
			if triggerRef == nil || triggerRef.Kind != "ImageStreamTag" {
				continue
			}
		}
		triggerNs := triggerRef.Namespace
		if triggerNs == "" {
			triggerNs = bc.Namespace
		}
		refNs := ref.Namespace
		if refNs == "" {
			refNs = bc.Namespace
		}
		if triggerRef.Name == ref.Name && triggerNs == refNs {
			return imageChange
		}
	}
	return nil
}
func describeBuildRequest(request *buildv1.BuildRequest) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	desc := fmt.Sprintf("BuildConfig: %s/%s", request.Namespace, request.Name)
	if request.Revision != nil {
		desc += fmt.Sprintf(", Revision: %#v", request.Revision.Git)
	}
	if request.TriggeredByImage != nil {
		desc += fmt.Sprintf(", TriggeredBy: %s/%s with stream: %s/%s", request.TriggeredByImage.Kind, request.TriggeredByImage.Name, request.From.Kind, request.From.Name)
	}
	if request.LastVersion != nil {
		desc += fmt.Sprintf(", LastVersion: %d", *request.LastVersion)
	}
	return desc
}
func updateBuildArgs(oldArgs *[]corev1.EnvVar, newArgs []corev1.EnvVar) []corev1.EnvVar {
	_logClusterCodePath()
	defer _logClusterCodePath()
	combined := make(map[string]string)
	for _, o := range *oldArgs {
		combined[o.Name] = o.Value
	}
	for _, n := range newArgs {
		combined[n.Name] = n.Value
	}
	var result []corev1.EnvVar
	for k, v := range combined {
		result = append(result, corev1.EnvVar{Name: k, Value: v})
	}
	return result
}
func (g *BuildGenerator) InstantiateInternal(ctx context.Context, request *buildapi.BuildRequest) (*buildapi.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	versionedRequest := &buildv1.BuildRequest{}
	if err := legacyscheme.Scheme.Convert(request, versionedRequest, nil); err != nil {
		return nil, fmt.Errorf("failed to convert internal BuildRequest to external: %v", err)
	}
	build, err := g.Instantiate(ctx, versionedRequest)
	if err != nil {
		return nil, err
	}
	internalBuild := &buildapi.Build{}
	if err := legacyscheme.Scheme.Convert(build, internalBuild, nil); err != nil {
		return nil, fmt.Errorf("failed to convert external Build to internal: %v", err)
	}
	return internalBuild, nil
}
func (g *BuildGenerator) Instantiate(ctx context.Context, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var build *buildv1.Build
	var err error
	for i := 0; i < conflictRetries; i++ {
		build, err = g.instantiate(ctx, request)
		if errors.IsConflict(err) {
			klog.V(4).Infof("instantiate returned conflict, try %d/%d", i+1, conflictRetries)
			continue
		}
		if err != nil {
			return nil, err
		}
		if err == nil {
			break
		}
	}
	return build, err
}
func (g *BuildGenerator) instantiate(ctx context.Context, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Generating Build from %s", describeBuildRequest(request))
	bc, err := g.Client.GetBuildConfig(ctx, request.Name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if isPaused(bc) {
		return nil, errors.NewBadRequest(fmt.Sprintf("can't instantiate from BuildConfig %s/%s: BuildConfig is paused", bc.Namespace, bc.Name))
	}
	if err := g.checkLastVersion(bc, request.LastVersion); err != nil {
		return nil, errors.NewBadRequest(err.Error())
	}
	if err := g.updateImageTriggers(ctx, bc, request.From, request.TriggeredByImage); err != nil {
		if _, ok := err.(errors.APIStatus); ok {
			return nil, err
		}
		return nil, errors.NewInternalError(err)
	}
	newBuild, err := g.generateBuildFromConfig(ctx, bc, request.Revision, request.Binary)
	if err != nil {
		if _, ok := err.(errors.APIStatus); ok {
			return nil, err
		}
		return nil, errors.NewInternalError(err)
	}
	newBuild.Annotations = mergeMaps(request.Annotations, newBuild.Annotations)
	newBuild.Labels = mergeMaps(request.Labels, newBuild.Labels)
	newBuild.Spec.TriggeredBy = request.TriggeredBy
	if len(request.Env) > 0 {
		buildutil.UpdateBuildEnv(newBuild, request.Env)
	}
	if request.DockerStrategyOptions != nil {
		dockerOpts := request.DockerStrategyOptions
		if dockerOpts.BuildArgs != nil && len(dockerOpts.BuildArgs) > 0 {
			if newBuild.Spec.Strategy.DockerStrategy == nil {
				return nil, errors.NewBadRequest(fmt.Sprintf("Cannot specify Docker build specific options on %s/%s, not a Docker build.", bc.Namespace, bc.ObjectMeta.Name))
			}
			newBuild.Spec.Strategy.DockerStrategy.BuildArgs = updateBuildArgs(&newBuild.Spec.Strategy.DockerStrategy.BuildArgs, dockerOpts.BuildArgs)
		}
		if dockerOpts.NoCache != nil {
			if newBuild.Spec.Strategy.DockerStrategy == nil {
				return nil, errors.NewBadRequest(fmt.Sprintf("Cannot specify Docker build specific options on %s/%s, not a Docker build.", bc.Namespace, bc.ObjectMeta.Name))
			}
			newBuild.Spec.Strategy.DockerStrategy.NoCache = *dockerOpts.NoCache
		}
	}
	if request.SourceStrategyOptions != nil {
		sourceOpts := request.SourceStrategyOptions
		if sourceOpts.Incremental != nil {
			if newBuild.Spec.Strategy.SourceStrategy == nil {
				return nil, errors.NewBadRequest(fmt.Sprintf("Cannot specify Source build specific options on %s/%s, not a Source build.", bc.Namespace, bc.ObjectMeta.Name))
			}
			newBuild.Spec.Strategy.SourceStrategy.Incremental = sourceOpts.Incremental
		}
	}
	klog.V(4).Infof("Build %s/%s has been generated from %s/%s BuildConfig", newBuild.Namespace, newBuild.ObjectMeta.Name, bc.Namespace, bc.ObjectMeta.Name)
	if err := g.Client.UpdateBuildConfig(ctx, bc); err != nil {
		klog.V(4).Infof("Failed to update BuildConfig %s/%s so no Build will be created", bc.Namespace, bc.Name)
		return nil, err
	}
	return g.createBuild(ctx, newBuild)
}
func (g *BuildGenerator) checkLastVersion(bc *buildv1.BuildConfig, lastVersion *int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lastVersion != nil && bc.Status.LastVersion != *lastVersion {
		klog.V(2).Infof("Aborting version triggered build for BuildConfig %s/%s because the BuildConfig LastVersion (%d) does not match the requested LastVersion (%d)", bc.Namespace, bc.Name, bc.Status.LastVersion, *lastVersion)
		return fmt.Errorf("the LastVersion(%v) on build config %s/%s does not match the build request LastVersion(%d)", bc.Status.LastVersion, bc.Namespace, bc.Name, *lastVersion)
	}
	return nil
}
func (g *BuildGenerator) updateImageTriggers(ctx context.Context, bc *buildv1.BuildConfig, from, triggeredBy *corev1.ObjectReference) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var requestTrigger *buildv1.ImageChangeTrigger
	if from != nil {
		requestTrigger = findImageChangeTrigger(bc, from)
	}
	if requestTrigger != nil && triggeredBy != nil && requestTrigger.LastTriggeredImageID == triggeredBy.Name {
		klog.V(2).Infof("Aborting imageid triggered build for BuildConfig %s/%s with imageid %s because the BuildConfig already matches this imageid", bc.Namespace, bc.Name, triggeredBy.Name)
		return fmt.Errorf("build config %s/%s has already instantiated a build for imageid %s", bc.Namespace, bc.Name, triggeredBy.Name)
	}
	for _, trigger := range bc.Spec.Triggers {
		if trigger.Type != buildv1.ImageChangeBuildTriggerType {
			continue
		}
		if triggeredBy != nil && trigger.ImageChange == requestTrigger {
			trigger.ImageChange.LastTriggeredImageID = triggeredBy.Name
			continue
		}
		triggerImageRef := trigger.ImageChange.From
		if triggerImageRef == nil {
			triggerImageRef = buildutil.GetInputReference(bc.Spec.Strategy)
		}
		if triggerImageRef == nil {
			klog.Warningf("Could not get ImageStream reference for default ImageChangeTrigger on BuildConfig %s/%s", bc.Namespace, bc.Name)
			continue
		}
		image, err := g.resolveImageStreamReference(ctx, *triggerImageRef, bc.Namespace)
		if err != nil {
			if trigger.ImageChange.From == nil {
				return err
			}
			klog.Warningf("Could not resolve trigger reference for build config %s/%s: %#v", bc.Namespace, bc.Name, triggerImageRef)
		}
		trigger.ImageChange.LastTriggeredImageID = image
	}
	return nil
}
func (g *BuildGenerator) CloneInternal(ctx context.Context, request *buildapi.BuildRequest) (*buildapi.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	versionedRequest := &buildv1.BuildRequest{}
	if err := legacyscheme.Scheme.Convert(request, versionedRequest, nil); err != nil {
		return nil, err
	}
	build, err := g.Clone(ctx, versionedRequest)
	if err != nil {
		return nil, err
	}
	internalBuild := &buildapi.Build{}
	if err := legacyscheme.Scheme.Convert(build, internalBuild, nil); err != nil {
		return nil, err
	}
	return internalBuild, nil
}
func (g *BuildGenerator) Clone(ctx context.Context, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var build *buildv1.Build
	var err error
	for i := 0; i < conflictRetries; i++ {
		build, err = g.clone(ctx, request)
		if err == nil || !errors.IsConflict(err) {
			break
		}
		klog.V(4).Infof("clone returned conflict, try %d/%d", i+1, conflictRetries)
	}
	return build, err
}
func (g *BuildGenerator) clone(ctx context.Context, request *buildv1.BuildRequest) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Generating build from build %s/%s", request.Namespace, request.Name)
	build, err := g.Client.GetBuild(ctx, request.Name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var buildConfig *buildv1.BuildConfig
	if build.Status.Config != nil {
		buildConfig, err = g.Client.GetBuildConfig(ctx, build.Status.Config.Name, &metav1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			return nil, err
		}
		if isPaused(buildConfig) {
			return nil, errors.NewInternalError(&buildutil.GeneratorFatalError{Reason: fmt.Sprintf("can't instantiate from BuildConfig %s/%s: BuildConfig is paused", buildConfig.Namespace, buildConfig.Name)})
		}
	}
	newBuild := generateBuildFromBuild(build, buildConfig)
	klog.V(4).Infof("Build %s/%s has been generated from Build %s/%s", newBuild.Namespace, newBuild.ObjectMeta.Name, build.Namespace, build.ObjectMeta.Name)
	newBuild.Spec.TriggeredBy = request.TriggeredBy
	if len(request.Env) > 0 {
		buildutil.UpdateBuildEnv(newBuild, request.Env)
	}
	if request.DockerStrategyOptions != nil {
		dockerOpts := request.DockerStrategyOptions
		if dockerOpts.BuildArgs != nil && len(dockerOpts.BuildArgs) > 0 {
			if newBuild.Spec.Strategy.DockerStrategy == nil {
				return nil, errors.NewBadRequest(fmt.Sprintf("Cannot specify build args on %s/%s, not a Docker build.", buildConfig.Namespace, buildConfig.ObjectMeta.Name))
			}
			newBuild.Spec.Strategy.DockerStrategy.BuildArgs = updateBuildArgs(&newBuild.Spec.Strategy.DockerStrategy.BuildArgs, dockerOpts.BuildArgs)
		}
	}
	if buildConfig != nil {
		if err := g.Client.UpdateBuildConfig(ctx, buildConfig); err != nil {
			klog.V(4).Infof("Failed to update BuildConfig %s/%s so no Build will be created", buildConfig.Namespace, buildConfig.Name)
			return nil, err
		}
	}
	return g.createBuild(ctx, newBuild)
}
func (g *BuildGenerator) createBuild(ctx context.Context, build *buildv1.Build) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !rest.ValidNamespace(ctx, &build.ObjectMeta) {
		return nil, errors.NewConflict(buildv1.Resource("build"), build.Namespace, fmt.Errorf("Build.Namespace does not match the provided context"))
	}
	rest.FillObjectMetaSystemFields(&build.ObjectMeta)
	err := g.Client.CreateBuild(ctx, build)
	if err != nil {
		return nil, err
	}
	return g.Client.GetBuild(ctx, build.Name, &metav1.GetOptions{})
}
func (g *BuildGenerator) generateBuildFromConfig(ctx context.Context, bc *buildv1.BuildConfig, revision *buildv1.SourceRevision, binary *buildv1.BinaryBuildSource) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildName := getNextBuildName(bc)
	bcCopy := bc.DeepCopy()
	serviceAccount := bcCopy.Spec.ServiceAccount
	if len(serviceAccount) == 0 {
		serviceAccount = bootstrappolicy.BuilderServiceAccountName
	}
	t := true
	build := &buildv1.Build{Spec: buildv1.BuildSpec{CommonSpec: buildv1.CommonSpec{ServiceAccount: serviceAccount, Source: bcCopy.Spec.Source, Strategy: bcCopy.Spec.Strategy, Output: bcCopy.Spec.Output, Revision: revision, Resources: bcCopy.Spec.Resources, PostCommit: bcCopy.Spec.PostCommit, CompletionDeadlineSeconds: bcCopy.Spec.CompletionDeadlineSeconds, NodeSelector: bcCopy.Spec.NodeSelector}}, ObjectMeta: metav1.ObjectMeta{Name: buildName, Labels: bcCopy.Labels, OwnerReferences: []metav1.OwnerReference{{APIVersion: buildv1.GroupVersion.String(), Kind: "BuildConfig", Name: bcCopy.Name, UID: bcCopy.UID, Controller: &t}}}, Status: buildv1.BuildStatus{Phase: buildv1.BuildPhaseNew, Config: &corev1.ObjectReference{Kind: "BuildConfig", Name: bcCopy.Name, Namespace: bcCopy.Namespace}}}
	setBuildSource(binary, build)
	setBuildAnnotationAndLabel(bcCopy, build)
	var builderSecrets []corev1.Secret
	var err error
	if builderSecrets, err = fetchServiceAccountSecrets(g.Secrets, g.ServiceAccounts, bcCopy.Namespace, serviceAccount); err != nil {
		return nil, err
	}
	if err = g.setBuildSourceImage(ctx, builderSecrets, bcCopy, &build.Spec.Source); err != nil {
		return nil, err
	}
	if err = g.setBaseImageAndPullSecretForBuildStrategy(ctx, builderSecrets, bcCopy, &build.Spec.Strategy); err != nil {
		return nil, err
	}
	return build, nil
}
func (g *BuildGenerator) setBuildSourceImage(ctx context.Context, builderSecrets []corev1.Secret, bcCopy *buildv1.BuildConfig, Source *buildv1.BuildSource) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	strategyImageChangeTrigger := getStrategyImageChangeTrigger(bcCopy)
	for i, sourceImage := range Source.Images {
		if sourceImage.PullSecret == nil {
			sourceImage.PullSecret = g.resolveImageSecret(ctx, builderSecrets, &sourceImage.From, bcCopy.Namespace)
		}
		var sourceImageSpec string
		if strategyFrom := buildutil.GetInputReference(bcCopy.Spec.Strategy); strategyFrom != nil && reflect.DeepEqual(sourceImage.From, *strategyFrom) && strategyImageChangeTrigger != nil {
			sourceImageSpec = strategyImageChangeTrigger.LastTriggeredImageID
		} else {
			refImageChangeTrigger := getImageChangeTriggerForRef(bcCopy, &sourceImage.From)
			if refImageChangeTrigger == nil {
				sourceImageSpec, err = g.resolveImageStreamReference(ctx, sourceImage.From, bcCopy.Namespace)
				if err != nil {
					return err
				}
			} else {
				sourceImageSpec = refImageChangeTrigger.LastTriggeredImageID
			}
		}
		sourceImage.From.Kind = "DockerImage"
		sourceImage.From.Name = sourceImageSpec
		sourceImage.From.Namespace = ""
		Source.Images[i] = sourceImage
	}
	return nil
}
func (g *BuildGenerator) setBaseImageAndPullSecretForBuildStrategy(ctx context.Context, builderSecrets []corev1.Secret, bcCopy *buildv1.BuildConfig, strategy *buildv1.BuildStrategy) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	var image string
	if strategyImageChangeTrigger := getStrategyImageChangeTrigger(bcCopy); strategyImageChangeTrigger != nil {
		image = strategyImageChangeTrigger.LastTriggeredImageID
	}
	switch {
	case strategy.SourceStrategy != nil:
		if image == "" {
			image, err = g.resolveImageStreamReference(ctx, strategy.SourceStrategy.From, bcCopy.Namespace)
			if err != nil {
				return err
			}
		}
		strategy.SourceStrategy.From = corev1.ObjectReference{Kind: "DockerImage", Name: image}
		if strategy.SourceStrategy.PullSecret == nil {
			strategy.SourceStrategy.PullSecret = g.resolveImageSecret(ctx, builderSecrets, &strategy.SourceStrategy.From, bcCopy.Namespace)
		}
	case strategy.DockerStrategy != nil && strategy.DockerStrategy.From != nil:
		if image == "" {
			image, err = g.resolveImageStreamReference(ctx, *strategy.DockerStrategy.From, bcCopy.Namespace)
			if err != nil {
				return err
			}
		}
		strategy.DockerStrategy.From = &corev1.ObjectReference{Kind: "DockerImage", Name: image}
		if strategy.DockerStrategy.PullSecret == nil {
			strategy.DockerStrategy.PullSecret = g.resolveImageSecret(ctx, builderSecrets, strategy.DockerStrategy.From, bcCopy.Namespace)
		}
	case strategy.CustomStrategy != nil:
		if image == "" {
			image, err = g.resolveImageStreamReference(ctx, strategy.CustomStrategy.From, bcCopy.Namespace)
			if err != nil {
				return err
			}
		}
		strategy.CustomStrategy.From = corev1.ObjectReference{Kind: "DockerImage", Name: image}
		if strategy.CustomStrategy.PullSecret == nil {
			strategy.CustomStrategy.PullSecret = g.resolveImageSecret(ctx, builderSecrets, &strategy.CustomStrategy.From, bcCopy.Namespace)
		}
		updateCustomImageEnv(strategy.CustomStrategy, image)
	}
	return nil
}
func (g *BuildGenerator) resolveImageStreamReference(ctx context.Context, from corev1.ObjectReference, defaultNamespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var namespace string
	if len(from.Namespace) != 0 {
		namespace = from.Namespace
	} else {
		namespace = defaultNamespace
	}
	klog.V(4).Infof("Resolving ImageStreamReference %s of Kind %s in namespace %s", from.Name, from.Kind, namespace)
	switch from.Kind {
	case "ImageStreamImage":
		name, id, err := imageutil.ParseImageStreamImageName(from.Name)
		if err != nil {
			err = resolveError(from.Kind, namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		stream, err := g.Client.GetImageStream(apirequest.WithNamespace(ctx, namespace), name, &metav1.GetOptions{})
		if err != nil {
			err = resolveError(from.Kind, namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		reference, ok := imageutil.DockerImageReferenceForImage(stream, id)
		if !ok {
			err = resolveError(from.Kind, namespace, from.Name, fmt.Errorf("unable to find corresponding tag for image %q", id))
			klog.V(2).Info(err)
			return "", err
		}
		klog.V(4).Infof("Resolved ImageStreamImage %s to image %q", from.Name, reference)
		return reference, nil
	case "ImageStreamTag":
		name, tag, err := imageutil.ParseImageStreamTagName(from.Name)
		if err != nil {
			err = resolveError(from.Kind, namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		stream, err := g.Client.GetImageStream(apirequest.WithNamespace(ctx, namespace), name, &metav1.GetOptions{})
		if err != nil {
			err = resolveError(from.Kind, namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		reference, ok := imageutil.ResolveLatestTaggedImage(stream, tag)
		if !ok {
			err = resolveError(from.Kind, namespace, from.Name, fmt.Errorf("unable to find latest tagged image"))
			klog.V(2).Info(err)
			return "", err
		}
		klog.V(4).Infof("Resolved ImageStreamTag %s to image %q", from.Name, reference)
		return reference, nil
	case "DockerImage":
		return from.Name, nil
	default:
		return "", fmt.Errorf("unknown From Kind %s", from.Kind)
	}
}
func (g *BuildGenerator) resolveImageStreamDockerRepository(ctx context.Context, from corev1.ObjectReference, defaultNamespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace := defaultNamespace
	if len(from.Namespace) > 0 {
		namespace = from.Namespace
	}
	klog.V(4).Infof("Resolving ImageStreamReference %s of Kind %s in namespace %s", from.Name, from.Kind, namespace)
	switch from.Kind {
	case "ImageStreamImage":
		imageStreamImage, err := g.Client.GetImageStreamImage(apirequest.WithNamespace(ctx, namespace), from.Name, &metav1.GetOptions{})
		if err != nil {
			err = resolveError(from.Kind, namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		image := imageStreamImage.Image
		klog.V(4).Infof("Resolved ImageStreamReference %s to image %s with reference %s in namespace %s", from.Name, image.Name, image.DockerImageReference, namespace)
		return image.DockerImageReference, nil
	case "ImageStreamTag":
		name := strings.Split(from.Name, ":")[0]
		is, err := g.Client.GetImageStream(apirequest.WithNamespace(ctx, namespace), name, &metav1.GetOptions{})
		if err != nil {
			err = resolveError("ImageStream", namespace, from.Name, err)
			klog.V(2).Info(err)
			return "", err
		}
		image, err := imageutil.DockerImageReferenceForStream(is)
		if err != nil {
			klog.V(2).Infof("Error resolving Docker image reference for %s/%s: %v", namespace, name, err)
			return "", err
		}
		klog.V(4).Infof("Resolved ImageStreamTag %s/%s to repository %s", namespace, from.Name, image)
		return image.String(), nil
	case "DockerImage":
		return from.Name, nil
	default:
		return "", fmt.Errorf("unknown From Kind %s", from.Kind)
	}
}
func (g *BuildGenerator) resolveImageSecret(ctx context.Context, secrets []corev1.Secret, imageRef *corev1.ObjectReference, buildNamespace string) *corev1.LocalObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(secrets) == 0 || imageRef == nil {
		return nil
	}
	imageSpec, err := g.resolveImageStreamDockerRepository(ctx, *imageRef, buildNamespace)
	if err != nil {
		klog.V(2).Infof("Unable to resolve the image name for %s/%s: %v", buildNamespace, imageRef, err)
		return nil
	}
	s := findDockerSecretAsInternalReference(secrets, imageSpec)
	if s == nil {
		klog.V(4).Infof("No secrets found for pushing or pulling the %s  %s/%s", imageRef.Kind, buildNamespace, imageRef.Name)
	}
	return s
}
func findDockerSecretAsInternalReference(secrets []corev1.Secret, image string) *corev1.LocalObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	emptyKeyring := credentialprovider.BasicDockerKeyring{}
	for _, secret := range secrets {
		externalSecret := corev1.Secret{}
		if err := legacyscheme.Scheme.Convert(&secret, &externalSecret, nil); err != nil {
			panic(err)
		}
		secretList := []corev1.Secret{externalSecret}
		keyring, err := credentialprovidersecrets.MakeDockerKeyring(secretList, &emptyKeyring)
		if err != nil {
			klog.V(2).Infof("Unable to make the Docker keyring for %s/%s secret: %v", secret.Name, secret.Namespace, err)
			continue
		}
		if _, found := keyring.Lookup(image); found {
			return &corev1.LocalObjectReference{Name: secret.Name}
		}
	}
	return nil
}
func resolveError(kind string, namespace string, name string, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := fmt.Sprintf("Error resolving %s %s in namespace %s: %v", kind, name, namespace, err)
	return &errors.StatusError{ErrStatus: metav1.Status{Status: metav1.StatusFailure, Code: http.StatusUnprocessableEntity, Reason: metav1.StatusReasonInvalid, Message: msg, Details: &metav1.StatusDetails{Kind: kind, Name: name, Causes: []metav1.StatusCause{{Field: "from", Message: msg}}}}}
}
func getNextBuildName(buildConfig *buildv1.BuildConfig) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildConfig.Status.LastVersion++
	return apihelpers.GetName(buildConfig.Name, strconv.FormatInt(buildConfig.Status.LastVersion, 10), kvalidation.DNS1123SubdomainMaxLength)
}
func updateCustomImageEnv(strategy *buildv1.CustomBuildStrategy, newImage string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strategy.Env == nil {
		strategy.Env = make([]corev1.EnvVar, 1)
		strategy.Env[0] = corev1.EnvVar{Name: buildutil.CustomBuildStrategyBaseImageKey, Value: newImage}
	} else {
		found := false
		for i := range strategy.Env {
			klog.V(4).Infof("Checking env variable %s %s", strategy.Env[i].Name, strategy.Env[i].Value)
			if strategy.Env[i].Name == buildutil.CustomBuildStrategyBaseImageKey {
				found = true
				strategy.Env[i].Value = newImage
				klog.V(4).Infof("Updated env variable %s to %s", strategy.Env[i].Name, strategy.Env[i].Value)
				break
			}
		}
		if !found {
			strategy.Env = append(strategy.Env, corev1.EnvVar{Name: buildutil.CustomBuildStrategyBaseImageKey, Value: newImage})
		}
	}
}
func generateBuildFromBuild(build *buildv1.Build, buildConfig *buildv1.BuildConfig) *buildv1.Build {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildCopy := build.DeepCopy()
	newBuild := &buildv1.Build{Spec: buildCopy.Spec, ObjectMeta: metav1.ObjectMeta{Name: getNextBuildNameFromBuild(buildCopy, buildConfig), Labels: buildCopy.ObjectMeta.Labels, Annotations: buildCopy.ObjectMeta.Annotations, OwnerReferences: buildCopy.ObjectMeta.OwnerReferences}, Status: buildv1.BuildStatus{Phase: buildv1.BuildPhaseNew, Config: buildCopy.Status.Config}}
	newBuild.Spec.Source.Type = ""
	newBuild.Spec.Source.Binary = nil
	if newBuild.Annotations == nil {
		newBuild.Annotations = make(map[string]string)
	}
	newBuild.Annotations[buildutil.BuildCloneAnnotation] = build.Name
	if buildConfig != nil {
		newBuild.Annotations[buildutil.BuildNumberAnnotation] = strconv.FormatInt(buildConfig.Status.LastVersion, 10)
	} else {
		delete(newBuild.Annotations, buildutil.BuildNumberAnnotation)
	}
	delete(newBuild.Annotations, buildutil.BuildJenkinsStatusJSONAnnotation)
	delete(newBuild.Annotations, buildutil.BuildJenkinsLogURLAnnotation)
	delete(newBuild.Annotations, buildutil.BuildJenkinsConsoleLogURLAnnotation)
	delete(newBuild.Annotations, buildutil.BuildJenkinsBlueOceanLogURLAnnotation)
	delete(newBuild.Annotations, buildutil.BuildJenkinsBuildURIAnnotation)
	delete(newBuild.Annotations, buildutil.BuildPodNameAnnotation)
	return newBuild
}
func getNextBuildNameFromBuild(build *buildv1.Build, buildConfig *buildv1.BuildConfig) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buildName string
	if buildConfig != nil {
		return getNextBuildName(buildConfig)
	}
	buildName = build.Name
	if matched, _ := regexp.MatchString(`^.+-\d{10}$`, buildName); matched {
		nameElems := strings.Split(buildName, "-")
		buildName = strings.Join(nameElems[:len(nameElems)-1], "-")
	}
	suffix := fmt.Sprintf("%v", metav1.Now().UnixNano())
	if len(suffix) > 10 {
		suffix = suffix[len(suffix)-10:]
	}
	return apihelpers.GetName(buildName, suffix, kvalidation.DNS1123SubdomainMaxLength)
}
func getStrategyImageChangeTrigger(bc *buildv1.BuildConfig) *buildv1.ImageChangeTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, trigger := range bc.Spec.Triggers {
		if trigger.Type == buildv1.ImageChangeBuildTriggerType && trigger.ImageChange.From == nil {
			return trigger.ImageChange
		}
	}
	return nil
}
func getImageChangeTriggerForRef(bc *buildv1.BuildConfig, ref *corev1.ObjectReference) *buildv1.ImageChangeTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ref == nil || ref.Kind != "ImageStreamTag" {
		return nil
	}
	for _, trigger := range bc.Spec.Triggers {
		if trigger.Type == buildv1.ImageChangeBuildTriggerType && trigger.ImageChange.From != nil && trigger.ImageChange.From.Name == ref.Name && trigger.ImageChange.From.Namespace == ref.Namespace {
			return trigger.ImageChange
		}
	}
	return nil
}
func setBuildSource(binary *buildv1.BinaryBuildSource, build *buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if binary != nil {
		build.Spec.Source.Git = nil
		build.Spec.Source.Binary = binary
		if build.Spec.Source.Dockerfile != nil && binary.AsFile == "Dockerfile" {
			build.Spec.Source.Dockerfile = nil
		}
	} else {
		build.Spec.Source.Type = ""
		build.Spec.Source.Binary = nil
	}
}
func setBuildAnnotationAndLabel(bcCopy *buildv1.BuildConfig, build *buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build.Annotations == nil {
		build.Annotations = make(map[string]string)
	}
	build.Annotations[buildutil.BuildNumberAnnotation] = strconv.FormatInt(bcCopy.Status.LastVersion, 10)
	build.Annotations[buildutil.BuildConfigAnnotation] = bcCopy.Name
	if build.Labels == nil {
		build.Labels = make(map[string]string)
	}
	build.Labels[buildutil.BuildConfigLabelDeprecated] = buildapihelpers.LabelValue(bcCopy.Name)
	build.Labels[buildutil.BuildConfigLabel] = buildapihelpers.LabelValue(bcCopy.Name)
	build.Labels[buildutil.BuildRunPolicyLabel] = string(bcCopy.Spec.RunPolicy)
}
func mergeMaps(a, b map[string]string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a == nil && b == nil {
		return nil
	}
	res := make(map[string]string)
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}
func isPaused(bc *buildv1.BuildConfig) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.ToLower(bc.Annotations[buildutil.BuildConfigPausedAnnotation]) == "true"
}
