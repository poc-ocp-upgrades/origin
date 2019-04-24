package imagepolicy

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
	"k8s.io/client-go/informers"
	"k8s.io/apiserver/pkg/admission/initializer"
	"github.com/hashicorp/golang-lru"
	"k8s.io/klog"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/diff"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	corev1listers "k8s.io/client-go/listers/core/v1"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	internalimagereferencemutators "github.com/openshift/origin/pkg/api/imagereferencemutators/internalversion"
	oadmission "github.com/openshift/origin/pkg/cmd/server/admission"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imagepolicy "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
	"github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/validation"
	"github.com/openshift/origin/pkg/image/apiserver/admission/imagepolicy/rules"
	imageinternalclient "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	"k8s.io/client-go/rest"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(imagepolicy.PluginName, func(input io.Reader) (admission.Interface, error) {
		config := &imagepolicy.ImagePolicyConfig{}
		if input != nil {
			configContent, err := ioutil.ReadAll(input)
			if err != nil {
				return nil, err
			}
			scheme := runtime.NewScheme()
			utilruntime.Must(imagepolicy.Install(scheme))
			codecs := serializer.NewCodecFactory(scheme)
			err = runtime.DecodeInto(codecs.UniversalDecoder(imagepolicy.GroupVersion), configContent, config)
			if err != nil {
				return nil, err
			}
		}
		imagepolicy.SetDefaults_ImagePolicyConfig(config)
		if errs := validation.Validate(config); len(errs) > 0 {
			return nil, errs.ToAggregate()
		}
		klog.V(5).Infof("%s admission controller loaded with config: %#v", imagepolicy.PluginName, config)
		return newImagePolicyPlugin(config)
	})
}

type imagePolicyPlugin struct {
	*admission.Handler
	config				*imagepolicy.ImagePolicyConfig
	client				imageinternalclient.ImageInterface
	accepter			rules.Accepter
	integratedRegistryMatcher	integratedRegistryMatcher
	nsLister			corev1listers.NamespaceLister
	resolver			imageResolver
}

var _ = initializer.WantsExternalKubeInformerFactory(&imagePolicyPlugin{})
var _ = oadmission.WantsRESTClientConfig(&imagePolicyPlugin{})
var _ = oadmission.WantsDefaultRegistryFunc(&imagePolicyPlugin{})
var _ = admission.ValidationInterface(&imagePolicyPlugin{})
var _ = admission.MutationInterface(&imagePolicyPlugin{})

type integratedRegistryMatcher struct{ rules.RegistryMatcher }
type imageResolver interface {
	ResolveObjectReference(ref *kapi.ObjectReference, defaultNamespace string, forceResolveLocalNames bool) (*rules.ImagePolicyAttributes, error)
}
type imageResolutionPolicy interface {
	RequestsResolution(metav1.GroupResource) bool
	FailOnResolutionFailure(metav1.GroupResource) bool
	RewriteImagePullSpec(attr *rules.ImagePolicyAttributes, isUpdate bool, gr metav1.GroupResource) bool
}

func newImagePolicyPlugin(parsed *imagepolicy.ImagePolicyConfig) (*imagePolicyPlugin, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := integratedRegistryMatcher{RegistryMatcher: rules.NewRegistryMatcher(nil)}
	accepter, err := rules.NewExecutionRulesAccepter(parsed.ExecutionRules, m)
	if err != nil {
		return nil, err
	}
	return &imagePolicyPlugin{Handler: admission.NewHandler(admission.Create, admission.Update), config: parsed, accepter: accepter, integratedRegistryMatcher: m}, nil
}
func (a *imagePolicyPlugin) SetDefaultRegistryFunc(fn func() (string, bool)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.integratedRegistryMatcher.RegistryMatcher = rules.RegistryNameMatcher(fn)
}
func (a *imagePolicyPlugin) SetRESTClientConfig(restClientConfig rest.Config) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	a.client, err = imageinternalclient.NewForConfig(&restClientConfig)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
}
func (a *imagePolicyPlugin) SetExternalKubeInformerFactory(kubeInformers informers.SharedInformerFactory) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.nsLister = kubeInformers.Core().V1().Namespaces().Lister()
}
func (a *imagePolicyPlugin) ValidateInitialization() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.client == nil {
		return fmt.Errorf("%s needs an Openshift client", imagepolicy.PluginName)
	}
	if a.nsLister == nil {
		return fmt.Errorf("%s needs a namespace lister", imagepolicy.PluginName)
	}
	imageResolver, err := newImageResolutionCache(a.client, a.integratedRegistryMatcher)
	if err != nil {
		return fmt.Errorf("unable to create image policy controller: %v", err)
	}
	a.resolver = imageResolver
	return nil
}
func (a *imagePolicyPlugin) Admit(attr admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.admit(attr, true)
}
func (a *imagePolicyPlugin) Validate(attr admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.admit(attr, false)
}
func (a *imagePolicyPlugin) admit(attr admission.Attributes, mutationAllowed bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch attr.GetOperation() {
	case admission.Create, admission.Update:
		if len(attr.GetSubresource()) > 0 {
			return nil
		}
	default:
		return nil
	}
	policy := resolutionConfig{a.config}
	schemagr := attr.GetResource().GroupResource()
	apigr := metav1.GroupResource{Resource: schemagr.Resource, Group: schemagr.Group}
	if !a.accepter.Covers(apigr) && !policy.Covers(apigr) {
		return nil
	}
	klog.V(5).Infof("running image policy admission for %s:%s/%s", attr.GetKind(), attr.GetNamespace(), attr.GetName())
	m, err := internalimagereferencemutators.GetImageReferenceMutator(attr.GetObject(), attr.GetOldObject())
	if err != nil {
		return apierrs.NewForbidden(schemagr, attr.GetName(), fmt.Errorf("unable to apply image policy against objects of type %T: %v", attr.GetObject(), err))
	}
	if !mutationAllowed {
		m = &mutationPreventer{m}
	}
	annotations, _ := internalimagereferencemutators.GetAnnotationAccessor(attr.GetObject())
	var excluded sets.String
	if ns := attr.GetNamespace(); len(ns) > 0 {
		if ns, err := a.nsLister.Get(ns); err == nil {
			if value := ns.Annotations[imagepolicy.IgnorePolicyRulesAnnotation]; len(value) > 0 {
				excluded = sets.NewString(strings.Split(value, ",")...)
			}
		}
	}
	if err := accept(a.accepter, policy, a.resolver, m, annotations, attr, excluded); err != nil {
		return err
	}
	return nil
}

type mutationPreventer struct {
	m internalimagereferencemutators.ImageReferenceMutator
}

func (m *mutationPreventer) Mutate(fn internalimagereferencemutators.ImageReferenceMutateFunc) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.m.Mutate(func(ref *kapi.ObjectReference) error {
		original := ref.DeepCopy()
		if err := fn(ref); err != nil {
			return fmt.Errorf("error in image policy validation: %v", err)
		}
		if !reflect.DeepEqual(ref, original) {
			klog.V(2).Infof("disallowed mutation in image policy validation: %s", diff.ObjectGoPrintSideBySide(original, ref))
			return fmt.Errorf("this image is prohibited by policy (changed after admission)")
		}
		return nil
	})
}

type imageResolutionCache struct {
	imageClient	imageinternalclient.ImageInterface
	integrated	rules.RegistryMatcher
	expiration	time.Duration
	cache		*lru.Cache
}
type imageCacheEntry struct {
	expires	time.Time
	image	*imageapi.Image
}

func newImageResolutionCache(imageClient imageinternalclient.ImageInterface, integratedRegistry rules.RegistryMatcher) (*imageResolutionCache, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	imageCache, err := lru.New(128)
	if err != nil {
		return nil, err
	}
	return &imageResolutionCache{imageClient: imageClient, integrated: integratedRegistry, cache: imageCache, expiration: time.Minute}, nil
}

var now = time.Now

func (c *imageResolutionCache) ResolveObjectReference(ref *kapi.ObjectReference, defaultNamespace string, forceResolveLocalNames bool) (*rules.ImagePolicyAttributes, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch ref.Kind {
	case "ImageStreamTag":
		ns := ref.Namespace
		if len(ns) == 0 {
			ns = defaultNamespace
		}
		name, tag, ok := imageapi.SplitImageStreamTag(ref.Name)
		if !ok {
			return &rules.ImagePolicyAttributes{IntegratedRegistry: true}, fmt.Errorf("references of kind ImageStreamTag must be of the form NAME:TAG")
		}
		return c.resolveImageStreamTag(ns, name, tag, false, false)
	case "ImageStreamImage":
		ns := ref.Namespace
		if len(ns) == 0 {
			ns = defaultNamespace
		}
		name, id, ok := imageapi.SplitImageStreamImage(ref.Name)
		if !ok {
			return &rules.ImagePolicyAttributes{IntegratedRegistry: true}, fmt.Errorf("references of kind ImageStreamImage must be of the form NAME@DIGEST")
		}
		return c.resolveImageStreamImage(ns, name, id)
	case "DockerImage":
		ref, err := imageapi.ParseDockerImageReference(ref.Name)
		if err != nil {
			return nil, err
		}
		return c.resolveImageReference(ref, defaultNamespace, forceResolveLocalNames)
	default:
		return nil, fmt.Errorf("image policy does not allow image references of kind %q", ref.Kind)
	}
}
func (c *imageResolutionCache) resolveImageReference(ref imageapi.DockerImageReference, defaultNamespace string, forceResolveLocalNames bool) (*rules.ImagePolicyAttributes, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ref.ID) > 0 {
		now := now()
		if value, ok := c.cache.Get(ref.ID); ok {
			cached := value.(imageCacheEntry)
			if now.Before(cached.expires) {
				return &rules.ImagePolicyAttributes{Name: ref, Image: cached.image}, nil
			}
		}
		image, err := c.imageClient.Images().Get(ref.ID, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		c.cache.Add(ref.ID, imageCacheEntry{expires: now.Add(c.expiration), image: image})
		return &rules.ImagePolicyAttributes{Name: ref, Image: image, IntegratedRegistry: c.integrated.Matches(ref.Registry)}, nil
	}
	fullReference := c.integrated.Matches(ref.Registry)
	partialReference := forceResolveLocalNames || (len(ref.Registry) == 0 && len(ref.Namespace) == 0 && len(ref.Name) > 0)
	if !fullReference && !partialReference {
		return nil, fmt.Errorf("(%s) could not be resolved to an exact image reference", ref.Exact())
	}
	tag := ref.Tag
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	if len(ref.Namespace) == 0 || forceResolveLocalNames {
		ref.Namespace = defaultNamespace
	}
	return c.resolveImageStreamTag(ref.Namespace, ref.Name, tag, partialReference, forceResolveLocalNames)
}
func (c *imageResolutionCache) resolveImageStreamTag(namespace, name, tag string, partial, forceResolveLocalNames bool) (*rules.ImagePolicyAttributes, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	attrs := &rules.ImagePolicyAttributes{IntegratedRegistry: true}
	resolved, err := c.imageClient.ImageStreamTags(namespace).Get(imageapi.JoinImageStreamTag(name, tag), metav1.GetOptions{})
	if err != nil {
		if partial {
			attrs.IntegratedRegistry = false
		}
		if isImageStreamTagNotFound(err) {
			if stream, err := c.imageClient.ImageStreams(namespace).Get(name, metav1.GetOptions{}); err == nil && (forceResolveLocalNames || stream.Spec.LookupPolicy.Local) && len(stream.Status.DockerImageRepository) > 0 {
				if ref, err := imageapi.ParseDockerImageReference(stream.Status.DockerImageRepository); err == nil {
					klog.V(4).Infof("%s/%s:%s points to a local name resolving stream, but the tag does not exist", namespace, name, tag)
					ref.Tag = tag
					attrs.Name = ref
					attrs.LocalRewrite = true
					return attrs, nil
				}
			}
		}
		return attrs, err
	}
	if partial {
		if !forceResolveLocalNames && !resolved.LookupPolicy.Local {
			attrs.IntegratedRegistry = false
			return attrs, fmt.Errorf("ImageStreamTag does not allow local references and the resource did not request image stream resolution")
		}
		attrs.LocalRewrite = true
	}
	ref, err := imageapi.ParseDockerImageReference(resolved.Image.DockerImageReference)
	if err != nil {
		return attrs, fmt.Errorf("image reference %s could not be parsed: %v", resolved.Image.DockerImageReference, err)
	}
	ref.Tag = ""
	ref.ID = resolved.Image.Name
	now := now()
	c.cache.Add(resolved.Image.Name, imageCacheEntry{expires: now.Add(c.expiration), image: &resolved.Image})
	attrs.Name = ref
	attrs.Image = &resolved.Image
	return attrs, nil
}
func (c *imageResolutionCache) resolveImageStreamImage(namespace, name, id string) (*rules.ImagePolicyAttributes, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	attrs := &rules.ImagePolicyAttributes{IntegratedRegistry: true}
	resolved, err := c.imageClient.ImageStreamImages(namespace).Get(imageapi.JoinImageStreamImage(name, id), metav1.GetOptions{})
	if err != nil {
		return attrs, err
	}
	ref, err := imageapi.ParseDockerImageReference(resolved.Image.DockerImageReference)
	if err != nil {
		return attrs, fmt.Errorf("ImageStreamTag could not be resolved: %v", err)
	}
	now := now()
	c.cache.Add(resolved.Image.Name, imageCacheEntry{expires: now.Add(c.expiration), image: &resolved.Image})
	attrs.Name = ref
	attrs.Image = &resolved.Image
	return attrs, nil
}
func isImageStreamTagNotFound(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil || !apierrs.IsNotFound(err) {
		return false
	}
	status, ok := err.(apierrs.APIStatus)
	if !ok {
		return false
	}
	details := status.Status().Details
	if details == nil {
		return false
	}
	return details.Kind == "imagestreamtags" && details.Group == "image.openshift.io"
}

type resolutionConfig struct {
	config *imagepolicy.ImagePolicyConfig
}

func (config resolutionConfig) Covers(gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, rule := range config.config.ResolutionRules {
		if resolutionRuleCoversResource(rule.TargetResource, gr) {
			return true
		}
	}
	return false
}
func (config resolutionConfig) RequestsResolution(gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if RequestsResolution(config.config.ResolveImages) {
		return true
	}
	for _, rule := range config.config.ResolutionRules {
		if resolutionRuleCoversResource(rule.TargetResource, gr) {
			return true
		}
	}
	return false
}
func (config resolutionConfig) FailOnResolutionFailure(gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return FailOnResolutionFailure(config.config.ResolveImages)
}

var skipImageRewriteOnUpdate = map[metav1.GroupResource]struct{}{{Group: "batch", Resource: "jobs"}: {}, {Group: "build.openshift.io", Resource: "builds"}: {}, {Group: "apps", Resource: "statefulsets"}: {}}

func (config resolutionConfig) RewriteImagePullSpec(attr *rules.ImagePolicyAttributes, isUpdate bool, gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isUpdate {
		if _, ok := skipImageRewriteOnUpdate[gr]; ok {
			return false
		}
	}
	hasMatchingRule := false
	for _, rule := range config.config.ResolutionRules {
		if !resolutionRuleCoversResource(rule.TargetResource, gr) {
			continue
		}
		if rule.LocalNames && attr.LocalRewrite {
			return true
		}
		if RewriteImagePullSpec(rule.Policy) {
			return true
		}
		hasMatchingRule = true
	}
	if hasMatchingRule {
		return false
	}
	return RewriteImagePullSpec(config.config.ResolveImages)
}
func resolutionRuleCoversResource(rule metav1.GroupResource, gr metav1.GroupResource) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rule.Group == gr.Group && (rule.Resource == gr.Resource || rule.Resource == "*")
}
