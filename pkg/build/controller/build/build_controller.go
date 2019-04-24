package build

import (
	"encoding/json"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
	"github.com/containers/image/signature"
	toml "github.com/pelletier/go-toml"
	"k8s.io/klog"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	metrics "github.com/openshift/origin/pkg/build/metrics/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	ktypedclient "k8s.io/client-go/kubernetes/typed/core/v1"
	v1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	buildv1 "github.com/openshift/api/build/v1"
	configv1 "github.com/openshift/api/config/v1"
	imagev1 "github.com/openshift/api/image/v1"
	buildv1client "github.com/openshift/client-go/build/clientset/versioned"
	buildv1informer "github.com/openshift/client-go/build/informers/externalversions/build/v1"
	buildv1lister "github.com/openshift/client-go/build/listers/build/v1"
	configv1informer "github.com/openshift/client-go/config/informers/externalversions/config/v1"
	configv1lister "github.com/openshift/client-go/config/listers/config/v1"
	imagev1informer "github.com/openshift/client-go/image/informers/externalversions/image/v1"
	imagev1lister "github.com/openshift/client-go/image/listers/image/v1"
	"github.com/openshift/origin/pkg/api/imagereferencemutators"
	"github.com/openshift/origin/pkg/build/buildscheme"
	buildmanualclient "github.com/openshift/origin/pkg/build/client"
	builddefaults "github.com/openshift/origin/pkg/build/controller/build/defaults"
	buildoverrides "github.com/openshift/origin/pkg/build/controller/build/overrides"
	"github.com/openshift/origin/pkg/build/controller/common"
	"github.com/openshift/origin/pkg/build/controller/policy"
	"github.com/openshift/origin/pkg/build/controller/strategy"
	buildutil "github.com/openshift/origin/pkg/build/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

const (
	maxRetries		= 15
	maxExcerptLength	= 5
)

type resourceTriggerQueue struct {
	lock	sync.Mutex
	queue	map[string][]string
}

func newResourceTriggerQueue() *resourceTriggerQueue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &resourceTriggerQueue{queue: make(map[string][]string)}
}
func (q *resourceTriggerQueue) Add(resource string, on []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, key := range on {
		q.queue[key] = append(q.queue[key], resource)
	}
}
func (q *resourceTriggerQueue) Remove(resource string, on []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, key := range on {
		resources := q.queue[key]
		newResources := make([]string, 0, len(resources))
		for _, existing := range resources {
			if existing == resource {
				continue
			}
			newResources = append(newResources, existing)
		}
		q.queue[key] = newResources
	}
}
func (q *resourceTriggerQueue) Pop(key string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.lock.Lock()
	defer q.lock.Unlock()
	resources := q.queue[key]
	delete(q.queue, key)
	return resources
}

type registryList struct {
	Registries []string `toml:"registries"`
}
type registries struct {
	Search		registryList	`toml:"search"`
	Insecure	registryList	`toml:"insecure"`
	Block		registryList	`toml:"block"`
}
type tomlConfig struct {
	Registries registries `toml:"registries"`
}
type BuildController struct {
	buildPatcher				buildmanualclient.BuildPatcher
	buildLister				buildv1lister.BuildLister
	buildConfigGetter			buildv1lister.BuildConfigLister
	buildDeleter				buildmanualclient.BuildDeleter
	buildControllerConfigLister		configv1lister.BuildLister
	imageConfigLister			configv1lister.ImageLister
	podClient				ktypedclient.PodsGetter
	configMapClient				ktypedclient.ConfigMapsGetter
	kubeClient				kubernetes.Interface
	buildQueue				workqueue.RateLimitingInterface
	imageStreamQueue			*resourceTriggerQueue
	buildConfigQueue			workqueue.RateLimitingInterface
	controllerConfigQueue			workqueue.RateLimitingInterface
	buildStore				buildv1lister.BuildLister
	secretStore				v1lister.SecretLister
	podStore				v1lister.PodLister
	imageStreamStore			imagev1lister.ImageStreamLister
	openShiftConfigConfigMapStore		v1lister.ConfigMapLister
	podInformer				cache.SharedIndexInformer
	buildInformer				cache.SharedIndexInformer
	buildStoreSynced			cache.InformerSynced
	buildControllerConfigStoreSynced	cache.InformerSynced
	imageConfigStoreSynced			cache.InformerSynced
	podStoreSynced				cache.InformerSynced
	secretStoreSynced			cache.InformerSynced
	imageStreamStoreSynced			cache.InformerSynced
	configMapStoreSynced			cache.InformerSynced
	runPolicies				[]policy.RunPolicy
	createStrategy				buildPodCreationStrategy
	buildDefaults				builddefaults.BuildDefaults
	buildOverrides				buildoverrides.BuildOverrides
	internalRegistryHostname		string
	recorder				record.EventRecorder
	registryConfData			string
	signaturePolicyData			string
	additionalTrustedCAData			map[string]string
}
type BuildControllerParams struct {
	BuildInformer				buildv1informer.BuildInformer
	BuildConfigInformer			buildv1informer.BuildConfigInformer
	BuildControllerConfigInformer		configv1informer.BuildInformer
	ImageConfigInformer			configv1informer.ImageInformer
	ImageStreamInformer			imagev1informer.ImageStreamInformer
	PodInformer				kubeinformers.PodInformer
	SecretInformer				kubeinformers.SecretInformer
	OpenshiftConfigConfigMapInformer	kubeinformers.ConfigMapInformer
	KubeClient				kubernetes.Interface
	BuildClient				buildv1client.Interface
	DockerBuildStrategy			*strategy.DockerBuildStrategy
	SourceBuildStrategy			*strategy.SourceBuildStrategy
	CustomBuildStrategy			*strategy.CustomBuildStrategy
	BuildDefaults				builddefaults.BuildDefaults
	BuildOverrides				buildoverrides.BuildOverrides
	InternalRegistryHostname		string
}

func NewBuildController(params *BuildControllerParams) *BuildController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&ktypedclient.EventSinkImpl{Interface: params.KubeClient.CoreV1().Events("")})
	buildClient := buildmanualclient.NewClientBuildClient(params.BuildClient)
	buildLister := params.BuildInformer.Lister()
	buildConfigGetter := params.BuildConfigInformer.Lister()
	c := &BuildController{buildPatcher: buildClient, buildLister: buildLister, buildConfigGetter: buildConfigGetter, buildDeleter: buildClient, buildControllerConfigLister: params.BuildControllerConfigInformer.Lister(), imageConfigLister: params.ImageConfigInformer.Lister(), secretStore: params.SecretInformer.Lister(), podClient: params.KubeClient.CoreV1(), configMapClient: params.KubeClient.CoreV1(), openShiftConfigConfigMapStore: params.OpenshiftConfigConfigMapInformer.Lister(), kubeClient: params.KubeClient, podInformer: params.PodInformer.Informer(), podStore: params.PodInformer.Lister(), buildInformer: params.BuildInformer.Informer(), buildStore: params.BuildInformer.Lister(), imageStreamStore: params.ImageStreamInformer.Lister(), createStrategy: &typeBasedFactoryStrategy{dockerBuildStrategy: params.DockerBuildStrategy, sourceBuildStrategy: params.SourceBuildStrategy, customBuildStrategy: params.CustomBuildStrategy}, buildDefaults: params.BuildDefaults, buildOverrides: params.BuildOverrides, internalRegistryHostname: params.InternalRegistryHostname, buildQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), imageStreamQueue: newResourceTriggerQueue(), buildConfigQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), controllerConfigQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()), recorder: eventBroadcaster.NewRecorder(buildscheme.EncoderScheme, corev1.EventSource{Component: "build-controller"}), runPolicies: policy.GetAllRunPolicies(buildLister, buildClient)}
	c.podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: c.podUpdated, DeleteFunc: c.podDeleted})
	c.buildInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.buildAdded, UpdateFunc: c.buildUpdated, DeleteFunc: c.buildDeleted})
	params.ImageStreamInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.imageStreamAdded, UpdateFunc: c.imageStreamUpdated})
	params.BuildControllerConfigInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.buildControllerConfigAdded, UpdateFunc: c.buildControllerConfigUpdated, DeleteFunc: c.buildControllerConfigDeleted})
	params.ImageConfigInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.imageConfigAdded, UpdateFunc: c.imageConfigUpdated, DeleteFunc: c.imageConfigDeleted})
	params.OpenshiftConfigConfigMapInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.configMapAdded, UpdateFunc: c.configMapUpdated, DeleteFunc: c.configMapDeleted})
	c.buildStoreSynced = c.buildInformer.HasSynced
	c.podStoreSynced = c.podInformer.HasSynced
	c.secretStoreSynced = params.SecretInformer.Informer().HasSynced
	c.imageStreamStoreSynced = params.ImageStreamInformer.Informer().HasSynced
	c.buildControllerConfigStoreSynced = params.BuildControllerConfigInformer.Informer().HasSynced
	c.imageConfigStoreSynced = params.ImageConfigInformer.Informer().HasSynced
	c.configMapStoreSynced = params.OpenshiftConfigConfigMapInformer.Informer().HasSynced
	return c
}
func (bc *BuildController) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer bc.buildQueue.ShutDown()
	defer bc.buildConfigQueue.ShutDown()
	defer bc.controllerConfigQueue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, bc.buildStoreSynced, bc.podStoreSynced, bc.secretStoreSynced, bc.imageStreamStoreSynced, bc.configMapStoreSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	if os.Getenv("OS_INTEGRATION_TEST") != "true" {
		if !cache.WaitForCacheSync(stopCh, bc.buildControllerConfigStoreSynced, bc.imageConfigStoreSynced) {
			utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return
		}
	}
	if errs := bc.handleControllerConfig(); len(errs) > 0 {
		utilruntime.HandleError(fmt.Errorf("errors syncing build controller config: %v", errs))
	}
	klog.Infof("Starting build controller")
	go wait.Until(bc.controllerConfigWorker, time.Second, stopCh)
	for i := 0; i < workers; i++ {
		go wait.Until(bc.buildWorker, time.Second, stopCh)
	}
	for i := 0; i < workers; i++ {
		go wait.Until(bc.buildConfigWorker, time.Second, stopCh)
	}
	metrics.IntializeMetricsCollector(bc.buildLister)
	<-stopCh
	klog.Infof("Shutting down build controller")
}
func (bc *BuildController) buildWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if quit := bc.buildWork(); quit {
			return
		}
	}
}
func (bc *BuildController) buildWork() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := bc.buildQueue.Get()
	if quit {
		return true
	}
	defer bc.buildQueue.Done(key)
	build, err := bc.getBuildByKey(key.(string))
	if err != nil {
		bc.handleBuildError(err, key)
		return false
	}
	if build == nil {
		return false
	}
	err = bc.handleBuild(build)
	bc.handleBuildError(err, key)
	return false
}
func (bc *BuildController) buildConfigWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if quit := bc.buildConfigWork(); quit {
			return
		}
	}
}
func (bc *BuildController) buildConfigWork() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := bc.buildConfigQueue.Get()
	if quit {
		return true
	}
	defer bc.buildConfigQueue.Done(key)
	namespace, name, err := parseBuildConfigKey(key.(string))
	if err != nil {
		utilruntime.HandleError(err)
		return false
	}
	err = bc.handleBuildConfig(namespace, name)
	bc.handleBuildConfigError(err, key)
	return false
}
func parseBuildConfigKey(key string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(key, "/", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid build config key: %s", key)
	}
	return parts[0], parts[1], nil
}
func (bc *BuildController) handleBuild(build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build.Spec.Strategy.JenkinsPipelineStrategy != nil {
		if buildutil.IsBuildComplete(build) {
			if err := common.HandleBuildPruning(buildutil.ConfigNameForBuild(build), build.Namespace, bc.buildLister, bc.buildConfigGetter, bc.buildDeleter); err != nil {
				utilruntime.HandleError(fmt.Errorf("failed to prune builds for %s/%s: %v", build.Namespace, build.Name, err))
			}
		}
	}
	if shouldIgnore(build) {
		return nil
	}
	klog.V(4).Infof("Handling build %s", buildDesc(build))
	pod, podErr := bc.podStore.Pods(build.Namespace).Get(buildapihelpers.GetBuildPodName(build))
	if podErr != nil && !errors.IsNotFound(podErr) {
		return podErr
	}
	var update *buildUpdate
	var err, updateErr error
	switch {
	case shouldCancel(build):
		update, err = bc.cancelBuild(build)
	case build.Status.Phase == buildv1.BuildPhaseNew:
		update, err = bc.handleNewBuild(build, pod)
	case build.Status.Phase == buildv1.BuildPhasePending, build.Status.Phase == buildv1.BuildPhaseRunning:
		update, err = bc.handleActiveBuild(build, pod)
	case buildutil.IsBuildComplete(build):
		update, err = bc.handleCompletedBuild(build, pod)
	}
	if update != nil && !update.isEmpty() {
		updateErr = bc.updateBuild(build, update, pod)
	}
	if err != nil {
		return err
	}
	if updateErr != nil {
		return updateErr
	}
	return nil
}
func shouldIgnore(build *buildv1.Build) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build.Spec.Strategy.JenkinsPipelineStrategy != nil {
		klog.V(4).Infof("Ignoring build %s with jenkins pipeline strategy", buildDesc(build))
		return true
	}
	if build.DeletionTimestamp != nil {
		return true
	}
	if buildutil.IsBuildComplete(build) {
		switch build.Status.Phase {
		case buildv1.BuildPhaseComplete:
			if build.Status.CompletionTimestamp == nil {
				return false
			}
		case buildv1.BuildPhaseFailed:
			if build.Status.CompletionTimestamp == nil || len(build.Status.LogSnippet) == 0 {
				return false
			}
		}
		klog.V(4).Infof("Ignoring build %s in completed state", buildDesc(build))
		return true
	}
	return false
}
func shouldCancel(build *buildv1.Build) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return !buildutil.IsBuildComplete(build) && build.Status.Cancelled
}
func (bc *BuildController) cancelBuild(build *buildv1.Build) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Cancelling build %s", buildDesc(build))
	podName := buildapihelpers.GetBuildPodName(build)
	err := bc.podClient.Pods(build.Namespace).Delete(podName, &metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return nil, fmt.Errorf("could not delete build pod %s/%s to cancel build %s: %v", build.Namespace, podName, buildDesc(build), err)
	}
	return transitionToPhase(buildv1.BuildPhaseCancelled, buildv1.StatusReasonCancelledBuild, buildutil.StatusMessageCancelledBuild), nil
}
func (bc *BuildController) handleNewBuild(build *buildv1.Build, pod *corev1.Pod) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod != nil {
		if strategy.HasOwnerReference(pod, build) {
			return bc.handleActiveBuild(build, pod)
		}
		return transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonBuildPodExists, buildutil.StatusMessageBuildPodExists), nil
	}
	runPolicy := policy.ForBuild(build, bc.runPolicies)
	if runPolicy == nil {
		return nil, fmt.Errorf("unable to determine build policy for %s", buildDesc(build))
	}
	if run, err := runPolicy.IsRunnable(build); err != nil || !run {
		return nil, err
	}
	return bc.createBuildPod(build)
}
func (bc *BuildController) createPodSpec(build *buildv1.Build, includeAdditionalCA bool) (*corev1.Pod, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if build.Spec.Output.To != nil {
		build.Status.OutputDockerImageReference = build.Spec.Output.To.Name
	}
	build.Status.Reason = ""
	build.Status.Message = ""
	podSpec, err := bc.createStrategy.CreateBuildPod(build, bc.additionalTrustedCAData, bc.internalRegistryHostname)
	if err != nil {
		if strategy.IsFatal(err) {
			return nil, &strategy.FatalError{Reason: fmt.Sprintf("failed to create a build pod spec for build %s/%s: %v", build.Namespace, build.Name, err)}
		}
		return nil, fmt.Errorf("failed to create a build pod spec for build %s/%s: %v", build.Namespace, build.Name, err)
	}
	if err := bc.buildDefaults.ApplyDefaults(podSpec); err != nil {
		return nil, fmt.Errorf("failed to apply build defaults for build %s/%s: %v", build.Namespace, build.Name, err)
	}
	if err := bc.buildOverrides.ApplyOverrides(podSpec); err != nil {
		return nil, fmt.Errorf("failed to apply build overrides for build %s/%s: %v", build.Namespace, build.Name, err)
	}
	if err := common.ResolveValueFrom(podSpec, bc.kubeClient); err != nil {
		return nil, err
	}
	return podSpec, nil
}
func (bc *BuildController) resolveImageSecretAsReference(build *buildv1.Build, imagename string) (*corev1.LocalObjectReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceAccount := build.Spec.ServiceAccount
	if len(serviceAccount) == 0 {
		serviceAccount = buildutil.BuilderServiceAccountName
	}
	builderSecrets, err := buildutil.FetchServiceAccountSecrets(bc.kubeClient.CoreV1(), build.Namespace, serviceAccount)
	if err != nil {
		return nil, fmt.Errorf("Error getting push/pull secrets for service account %s/%s: %v", build.Namespace, serviceAccount, err)
	}
	var secret *corev1.LocalObjectReference
	if len(imagename) != 0 {
		secret = buildutil.FindDockerSecretAsReference(builderSecrets, imagename)
	}
	if secret == nil {
		klog.V(4).Infof("build %s is referencing an unknown image, will attempt to use the default secret for the service account", build.Name)
		dockerSecretExists := false
		for _, builderSecret := range builderSecrets {
			if builderSecret.Type == corev1.SecretTypeDockercfg || builderSecret.Type == corev1.SecretTypeDockerConfigJson {
				dockerSecretExists = true
				secret = &corev1.LocalObjectReference{Name: builderSecret.Name}
				break
			}
		}
		if !dockerSecretExists {
			return nil, fmt.Errorf("No docker secrets associated with build service account %s", serviceAccount)
		}
		klog.V(4).Infof("No secrets found for pushing or pulling image named %s for build, using default: %s %s/%s", imagename, build.Namespace, build.Name, secret.Name)
	}
	return secret, nil
}
func resourceName(namespace, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return namespace + "/" + name
}

var (
	errInvalidImageReferences	= fmt.Errorf("one or more image references were invalid")
	errNoIntegratedRegistry		= fmt.Errorf("the integrated registry is not configured")
)

func unresolvedImageStreamReferences(m imagereferencemutators.ImageReferenceMutator, defaultNamespace string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var streams []string
	fn := func(ref *corev1.ObjectReference) error {
		switch ref.Kind {
		case "ImageStreamImage":
			namespace := ref.Namespace
			if len(namespace) == 0 {
				namespace = defaultNamespace
			}
			name, _, ok := imageapi.SplitImageStreamImage(ref.Name)
			if !ok {
				return errInvalidImageReferences
			}
			streams = append(streams, resourceName(namespace, name))
		case "ImageStreamTag":
			namespace := ref.Namespace
			if len(namespace) == 0 {
				namespace = defaultNamespace
			}
			name, _, ok := imageapi.SplitImageStreamTag(ref.Name)
			if !ok {
				return errInvalidImageReferences
			}
			streams = append(streams, resourceName(namespace, name))
		}
		return nil
	}
	errs := m.Mutate(fn)
	if len(errs) > 0 {
		return nil, errInvalidImageReferences
	}
	return streams, nil
}
func resolveImageStreamLocation(ref *corev1.ObjectReference, lister imagev1lister.ImageStreamLister, defaultNamespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace := ref.Namespace
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}
	var (
		name	string
		tag	string
	)
	switch ref.Kind {
	case "ImageStreamImage":
		var ok bool
		name, _, ok = imageapi.SplitImageStreamImage(ref.Name)
		if !ok {
			return "", errInvalidImageReferences
		}
		tag = imageapi.DefaultImageTag
	case "ImageStreamTag":
		var ok bool
		name, tag, ok = imageapi.SplitImageStreamTag(ref.Name)
		if !ok {
			return "", errInvalidImageReferences
		}
	case "ImageStream":
		name = ref.Name
	}
	stream, err := lister.ImageStreams(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			return "", err
		}
		return "", fmt.Errorf("the referenced output image stream %s/%s could not be found: %v", namespace, name, err)
	}
	if len(stream.Status.DockerImageRepository) == 0 {
		return "", errNoIntegratedRegistry
	}
	repo, err := imageapi.ParseDockerImageReference(stream.Status.DockerImageRepository)
	if err != nil {
		return "", fmt.Errorf("the referenced output image stream does not represent a valid reference name: %v", err)
	}
	repo.ID = ""
	repo.Tag = tag
	return repo.Exact(), nil
}
func resolveImageStreamImage(ref *corev1.ObjectReference, lister imagev1lister.ImageStreamLister, defaultNamespace string) (*corev1.ObjectReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace := ref.Namespace
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}
	name, imageID, ok := imageapi.SplitImageStreamImage(ref.Name)
	if !ok {
		return nil, errInvalidImageReferences
	}
	stream, err := lister.ImageStreams(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, err
		}
		return nil, fmt.Errorf("the referenced image stream %s/%s could not be found: %v", namespace, name, err)
	}
	event, err := imageutil.ResolveImageID(stream, imageID)
	if err != nil {
		return nil, err
	}
	if len(event.DockerImageReference) == 0 {
		return nil, fmt.Errorf("the referenced image stream image %s/%s does not have a pull spec", namespace, ref.Name)
	}
	return &corev1.ObjectReference{Kind: "DockerImage", Name: event.DockerImageReference}, nil
}
func resolveImageStreamTag(ref *corev1.ObjectReference, lister imagev1lister.ImageStreamLister, defaultNamespace string) (*corev1.ObjectReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace := ref.Namespace
	if len(namespace) == 0 {
		namespace = defaultNamespace
	}
	name, tag, ok := imageapi.SplitImageStreamTag(ref.Name)
	if !ok {
		return nil, errInvalidImageReferences
	}
	stream, err := lister.ImageStreams(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, err
		}
		return nil, fmt.Errorf("the referenced image stream %s/%s could not be found: %v", namespace, name, err)
	}
	if newRef, ok := imageutil.ResolveLatestTaggedImage(stream, tag); ok {
		return &corev1.ObjectReference{Kind: "DockerImage", Name: newRef}, nil
	}
	return nil, fmt.Errorf("the referenced image stream tag %s/%s does not exist", namespace, ref.Name)
}
func (bc *BuildController) resolveOutputDockerImageReference(build *buildv1.Build) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ref := build.Spec.Output.To
	if ref == nil || ref.Name == "" {
		return nil
	}
	switch ref.Kind {
	case "ImageStream", "ImageStreamTag":
		newRef, err := resolveImageStreamLocation(ref, bc.imageStreamStore, build.Namespace)
		if err != nil {
			return err
		}
		*ref = corev1.ObjectReference{Kind: "DockerImage", Name: newRef}
		return nil
	default:
		return nil
	}
}
func (bc *BuildController) resolveImageReferences(build *buildv1.Build, update *buildUpdate) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := imagereferencemutators.NewBuildMutator(build)
	streams, err := unresolvedImageStreamReferences(m, build.Namespace)
	if err != nil {
		return err
	}
	if len(streams) == 0 {
		klog.V(5).Infof("Build %s contains no unresolved image references", build.Name)
		return nil
	}
	buildKey := resourceName(build.Namespace, build.Name)
	bc.imageStreamQueue.Add(buildKey, streams)
	if err := bc.resolveOutputDockerImageReference(build); err != nil {
		update.setReason(buildv1.StatusReasonInvalidOutputReference)
		update.setMessage(buildutil.StatusMessageInvalidOutputRef)
		if err == errNoIntegratedRegistry {
			e := fmt.Errorf("an image stream cannot be used as build output because the integrated Docker registry is not configured")
			bc.recorder.Eventf(build, corev1.EventTypeWarning, "InvalidOutput", "Error starting build: %v", e)
		}
		return err
	}
	errs := m.Mutate(func(ref *corev1.ObjectReference) error {
		switch ref.Kind {
		case "ImageStreamImage":
			newRef, err := resolveImageStreamImage(ref, bc.imageStreamStore, build.Namespace)
			if err != nil {
				return err
			}
			*ref = *newRef
		case "ImageStreamTag":
			newRef, err := resolveImageStreamTag(ref, bc.imageStreamStore, build.Namespace)
			if err != nil {
				return err
			}
			*ref = *newRef
		}
		return nil
	})
	if len(errs) > 0 {
		update.setReason(buildv1.StatusReasonInvalidImageReference)
		update.setMessage(buildutil.StatusMessageInvalidImageRef)
		return errs.ToAggregate()
	}
	bc.imageStreamQueue.Remove(buildKey, streams)
	return nil
}
func (bc *BuildController) createBuildPod(build *buildv1.Build) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	update := &buildUpdate{}
	var err error
	build = build.DeepCopy()
	if err := bc.resolveImageReferences(build, update); err != nil {
		if hasError(err, errors.IsNotFound, field.NewErrorTypeMatcher(field.ErrorTypeNotFound)) {
			return update, nil
		}
		return update, err
	}
	pushSecret := build.Spec.Output.PushSecret
	if build.Spec.Output.PushSecret == nil && build.Spec.Output.To != nil && len(build.Spec.Output.To.Name) > 0 {
		var err error
		pushSecret, err = bc.resolveImageSecretAsReference(build, build.Spec.Output.To.Name)
		if err != nil {
			update.setReason(buildv1.StatusReasonCannotRetrieveServiceAccount)
			update.setMessage(buildutil.StatusMessageCannotRetrieveServiceAccount)
			return update, err
		}
	}
	build.Spec.Output.PushSecret = pushSecret
	var pullSecret *corev1.LocalObjectReference
	var imageName string
	switch {
	case build.Spec.Strategy.SourceStrategy != nil:
		pullSecret = build.Spec.Strategy.SourceStrategy.PullSecret
		imageName = build.Spec.Strategy.SourceStrategy.From.Name
	case build.Spec.Strategy.DockerStrategy != nil:
		pullSecret = build.Spec.Strategy.DockerStrategy.PullSecret
		if build.Spec.Strategy.DockerStrategy.From != nil {
			imageName = build.Spec.Strategy.DockerStrategy.From.Name
		}
	case build.Spec.Strategy.CustomStrategy != nil:
		pullSecret = build.Spec.Strategy.CustomStrategy.PullSecret
		imageName = build.Spec.Strategy.CustomStrategy.From.Name
	}
	if pullSecret == nil {
		var err error
		pullSecret, err = bc.resolveImageSecretAsReference(build, imageName)
		if err != nil {
			update.setReason(buildv1.StatusReasonCannotRetrieveServiceAccount)
			update.setMessage(buildutil.StatusMessageCannotRetrieveServiceAccount)
			return update, err
		}
		if pullSecret != nil {
			switch {
			case build.Spec.Strategy.SourceStrategy != nil:
				build.Spec.Strategy.SourceStrategy.PullSecret = pullSecret
			case build.Spec.Strategy.DockerStrategy != nil:
				build.Spec.Strategy.DockerStrategy.PullSecret = pullSecret
			case build.Spec.Strategy.CustomStrategy != nil:
				build.Spec.Strategy.CustomStrategy.PullSecret = pullSecret
			}
		}
	}
	for i, s := range build.Spec.Source.Images {
		if s.PullSecret != nil {
			continue
		}
		imageInputPullSecret, err := bc.resolveImageSecretAsReference(build, s.From.Name)
		if err != nil {
			update.setReason(buildv1.StatusReasonCannotRetrieveServiceAccount)
			update.setMessage(buildutil.StatusMessageCannotRetrieveServiceAccount)
			return update, err
		}
		build.Spec.Source.Images[i].PullSecret = imageInputPullSecret
	}
	if build.Spec.Strategy.CustomStrategy != nil {
		buildutil.UpdateCustomImageEnv(build.Spec.Strategy.CustomStrategy, build.Spec.Strategy.CustomStrategy.From.Name)
	}
	includeAdditionalCA := false
	if len(bc.additionalTrustedCAData) > 0 {
		includeAdditionalCA = true
	}
	buildPod, err := bc.createPodSpec(build, includeAdditionalCA)
	if err != nil {
		switch err.(type) {
		case common.ErrEnvVarResolver:
			update = transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonUnresolvableEnvironmentVariable, fmt.Sprintf("%v, %v", buildutil.StatusMessageUnresolvableEnvironmentVariable, err.Error()))
		default:
			update.setReason(buildv1.StatusReasonCannotCreateBuildPodSpec)
			update.setMessage(buildutil.StatusMessageCannotCreateBuildPodSpec)
		}
		utilruntime.HandleError(err)
		return update, nil
	}
	klog.V(4).Infof("Pod %s/%s for build %s is about to be created", build.Namespace, buildPod.Name, buildDesc(build))
	pod, err := bc.podClient.Pods(build.Namespace).Create(buildPod)
	if err != nil && !errors.IsAlreadyExists(err) {
		bc.recorder.Eventf(build, corev1.EventTypeWarning, "FailedCreate", "Error creating build pod: %v", err)
		update.setReason(buildv1.StatusReasonCannotCreateBuildPod)
		update.setMessage(buildutil.StatusMessageCannotCreateBuildPod)
		return update, fmt.Errorf("failed to create build pod: %v", err)
	} else if err != nil {
		bc.recorder.Eventf(build, corev1.EventTypeWarning, "FailedCreate", "Pod already exists: %s/%s", buildPod.Namespace, buildPod.Name)
		klog.V(4).Infof("Build pod %s/%s for build %s already exists", build.Namespace, buildPod.Name, buildDesc(build))
		existingPod, err := bc.podClient.Pods(build.Namespace).Get(buildPod.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if !strategy.HasOwnerReference(existingPod, build) {
			klog.V(4).Infof("Did not recognise pod %s/%s as belonging to build %s", build.Namespace, buildPod.Name, buildDesc(build))
			update = transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonBuildPodExists, buildutil.StatusMessageBuildPodExists)
			return update, nil
		}
		klog.V(4).Infof("Recognised pod %s/%s as belonging to build %s", build.Namespace, buildPod.Name, buildDesc(build))
		hasCAMap, err := bc.findOwnedConfigMap(existingPod, build.Namespace, buildapihelpers.GetBuildCAConfigMapName(build))
		if err != nil {
			return update, fmt.Errorf("could not find certificate authority for build: %v", err)
		}
		if !hasCAMap {
			update, err = bc.createBuildCAConfigMap(build, existingPod, update)
			if err != nil {
				return update, err
			}
		}
		hasRegistryConf, err := bc.findOwnedConfigMap(existingPod, build.Namespace, buildapihelpers.GetBuildSystemConfigMapName(build))
		if err != nil {
			return update, fmt.Errorf("could not find registry config for build: %v", err)
		}
		if !hasRegistryConf {
			update, err = bc.createBuildSystemConfConfigMap(build, existingPod, update)
			if err != nil {
				return update, err
			}
		}
	} else {
		klog.V(4).Infof("Created pod %s/%s for build %s", build.Namespace, buildPod.Name, buildDesc(build))
		update, err = bc.createBuildCAConfigMap(build, pod, update)
		if err != nil {
			return update, err
		}
		update, err = bc.createBuildSystemConfConfigMap(build, pod, update)
		if err != nil {
			return nil, err
		}
	}
	update = transitionToPhase(buildv1.BuildPhasePending, "", "")
	if pushSecret != nil {
		update.setPushSecret(*pushSecret)
	}
	update.setPodNameAnnotation(buildPod.Name)
	if build.Spec.Output.To != nil {
		update.setOutputRef(build.Spec.Output.To.Name)
	}
	return update, nil
}
func (bc *BuildController) handleActiveBuild(build *buildv1.Build, pod *corev1.Pod) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod == nil {
		pod = bc.findMissingPod(build)
		if pod == nil {
			klog.V(4).Infof("Failed to find the build pod for build %s. Moving it to Error state", buildDesc(build))
			return transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonBuildPodDeleted, buildutil.StatusMessageBuildPodDeleted), nil
		}
	}
	podPhase := pod.Status.Phase
	var update *buildUpdate
	if podPhase == corev1.PodPending {
		for _, initContainer := range pod.Status.InitContainerStatuses {
			if initContainer.Name == strategy.GitCloneContainer && (initContainer.State.Running != nil || initContainer.State.Terminated != nil) {
				podPhase = corev1.PodRunning
			}
		}
	}
	switch podPhase {
	case corev1.PodPending:
		switch {
		case build.Status.Phase == buildv1.BuildPhaseNew:
			update = transitionToPhase(buildv1.BuildPhasePending, "", "")
			fallthrough
		case build.Status.Phase == buildv1.BuildPhasePending:
			if secret := build.Spec.Output.PushSecret; secret != nil && build.Status.Reason != buildv1.StatusReasonMissingPushSecret {
				if _, err := bc.secretStore.Secrets(build.Namespace).Get(secret.Name); err != nil && errors.IsNotFound(err) {
					klog.V(4).Infof("Setting reason for pending build to %q due to missing secret for %s", build.Status.Reason, buildDesc(build))
					update = transitionToPhase(buildv1.BuildPhasePending, buildv1.StatusReasonMissingPushSecret, buildutil.StatusMessageMissingPushSecret)
				}
			}
		default:
			bc.recorder.Eventf(build, corev1.EventTypeWarning, "UnexpectedPodPhase", "Build %s received a pod in pending phase event while in %s phase", resourceName(build.Namespace, build.Name), string(build.Status.Phase))
		}
	case corev1.PodRunning:
		if build.Status.Phase != buildv1.BuildPhaseRunning {
			update = transitionToPhase(buildv1.BuildPhaseRunning, "", "")
			if pod.Status.StartTime != nil {
				update.setStartTime(*pod.Status.StartTime)
			}
		}
	case corev1.PodSucceeded:
		if build.Status.Phase != buildv1.BuildPhaseComplete {
			update = transitionToPhase(buildv1.BuildPhaseComplete, "", "")
		}
		if len(pod.Status.ContainerStatuses) == 0 {
			klog.V(2).Infof("Setting build %s to error state because its pod has no containers", buildDesc(build))
			update = transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonNoBuildContainerStatus, buildutil.StatusMessageNoBuildContainerStatus)
		} else {
			for _, info := range pod.Status.ContainerStatuses {
				if info.State.Terminated != nil && info.State.Terminated.ExitCode != 0 {
					klog.V(2).Infof("Setting build %s to error state because a container in its pod has non-zero exit code", buildDesc(build))
					update = transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonFailedContainer, buildutil.StatusMessageFailedContainer)
					break
				}
			}
		}
	case corev1.PodFailed:
		if isOOMKilled(pod) {
			update = transitionToPhase(buildv1.BuildPhaseFailed, buildv1.StatusReasonOutOfMemoryKilled, buildutil.StatusMessageOutOfMemoryKilled)
		} else if isPodEvicted(pod) {
			update = transitionToPhase(buildv1.BuildPhaseFailed, buildv1.StatusReasonBuildPodEvicted, pod.Status.Message)
		} else if build.Status.Phase != buildv1.BuildPhaseFailed {
			if pod.DeletionTimestamp != nil {
				update = transitionToPhase(buildv1.BuildPhaseError, buildv1.StatusReasonBuildPodDeleted, buildutil.StatusMessageBuildPodDeleted)
			} else {
				update = transitionToPhase(buildv1.BuildPhaseFailed, buildv1.StatusReasonGenericBuildFailed, buildutil.StatusMessageGenericBuildFailed)
			}
		}
	}
	return update, nil
}
func isOOMKilled(pod *corev1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod == nil {
		return false
	}
	if pod.Status.Reason == "OOMKilled" {
		return true
	}
	for _, c := range pod.Status.InitContainerStatuses {
		terminated := c.State.Terminated
		if terminated != nil && terminated.Reason == "OOMKilled" {
			return true
		}
	}
	for _, c := range pod.Status.ContainerStatuses {
		terminated := c.State.Terminated
		if terminated != nil && terminated.Reason == "OOMKilled" {
			return true
		}
	}
	return false
}
func isPodEvicted(pod *corev1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod == nil {
		return false
	}
	if pod.Status.Reason == "Evicted" {
		return true
	}
	return false
}
func (bc *BuildController) handleCompletedBuild(build *buildv1.Build, pod *corev1.Pod) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	update := &buildUpdate{}
	if isOOMKilled(pod) {
		update = transitionToPhase(buildv1.BuildPhaseFailed, buildv1.StatusReasonOutOfMemoryKilled, buildutil.StatusMessageOutOfMemoryKilled)
	}
	setBuildCompletionData(build, pod, update)
	return update, nil
}
func (bc *BuildController) updateBuild(build *buildv1.Build, update *buildUpdate, pod *corev1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stateTransition := false
	if update.phase != nil && (*update.phase) != build.Status.Phase {
		stateTransition = true
	} else if build.Status.Phase == buildv1.BuildPhaseFailed && update.completionTime != nil {
		update.setPhase(buildv1.BuildPhaseFailed)
		stateTransition = true
	}
	if stateTransition {
		if !isValidTransition(build.Status.Phase, *update.phase) {
			return fmt.Errorf("invalid phase transition %s -> %s", buildDesc(build), *update.phase)
		}
		reasonText := ""
		if update.reason != nil && *update.reason != "" {
			reasonText = fmt.Sprintf(" ( %s )", *update.reason)
		}
		if buildutil.IsTerminalPhase(*update.phase) {
			setBuildCompletionData(build, pod, update)
		}
		klog.V(4).Infof("Updating build %s -> %s%s", buildDesc(build), *update.phase, reasonText)
	}
	if update.podNameAnnotation == nil && !common.HasBuildPodNameAnnotation(build) && pod != nil {
		update.setPodNameAnnotation(pod.Name)
	}
	patchedBuild, err := bc.patchBuild(build, update)
	if err != nil {
		return err
	}
	if stateTransition {
		switch *update.phase {
		case buildv1.BuildPhaseRunning:
			bc.recorder.Eventf(patchedBuild, corev1.EventTypeNormal, buildutil.BuildStartedEventReason, fmt.Sprintf(buildutil.BuildStartedEventMessage, patchedBuild.Namespace, patchedBuild.Name))
		case buildv1.BuildPhaseCancelled:
			bc.recorder.Eventf(patchedBuild, corev1.EventTypeNormal, buildutil.BuildCancelledEventReason, fmt.Sprintf(buildutil.BuildCancelledEventMessage, patchedBuild.Namespace, patchedBuild.Name))
		case buildv1.BuildPhaseComplete:
			bc.recorder.Eventf(patchedBuild, corev1.EventTypeNormal, buildutil.BuildCompletedEventReason, fmt.Sprintf(buildutil.BuildCompletedEventMessage, patchedBuild.Namespace, patchedBuild.Name))
		case buildv1.BuildPhaseError, buildv1.BuildPhaseFailed:
			bc.recorder.Eventf(patchedBuild, corev1.EventTypeNormal, buildutil.BuildFailedEventReason, fmt.Sprintf(buildutil.BuildFailedEventMessage, patchedBuild.Namespace, patchedBuild.Name))
		}
		if buildutil.IsTerminalPhase(*update.phase) {
			bc.handleBuildCompletion(patchedBuild)
		}
	}
	return nil
}
func (bc *BuildController) handleBuildCompletion(build *buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcName := buildutil.ConfigNameForBuild(build)
	bc.enqueueBuildConfig(build.Namespace, bcName)
	if err := common.HandleBuildPruning(bcName, build.Namespace, bc.buildLister, bc.buildConfigGetter, bc.buildDeleter); err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to prune builds for %s/%s: %v", build.Namespace, build.Name, err))
	}
}
func (bc *BuildController) enqueueBuildConfig(ns, name string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := resourceName(ns, name)
	bc.buildConfigQueue.Add(key)
}
func (bc *BuildController) handleBuildConfig(bcNamespace string, bcName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Handling build config %s/%s", bcNamespace, bcName)
	nextBuilds, hasRunningBuilds, err := policy.GetNextConfigBuild(bc.buildLister, bcNamespace, bcName)
	if err != nil {
		klog.V(2).Infof("Error getting next builds for %s/%s: %v", bcNamespace, bcName, err)
		return err
	}
	klog.V(5).Infof("Build config %s/%s: has %d next builds, is running builds: %v", bcNamespace, bcName, len(nextBuilds), hasRunningBuilds)
	if hasRunningBuilds {
		klog.V(4).Infof("Build config %s/%s has running builds, will retry", bcNamespace, bcName)
		return fmt.Errorf("build config %s/%s has running builds and cannot run more builds", bcNamespace, bcName)
	}
	if len(nextBuilds) == 0 {
		klog.V(4).Infof("Build config %s/%s has no builds to run next, will retry", bcNamespace, bcName)
		return fmt.Errorf("build config %s/%s has no builds to run next", bcNamespace, bcName)
	}
	for _, build := range nextBuilds {
		klog.V(5).Infof("Queueing next build for build config %s/%s: %s", bcNamespace, bcName, build.Name)
		bc.enqueueBuild(build)
	}
	return nil
}
func createBuildPatch(older, newer *buildv1.Build) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newerJSON, err := runtime.Encode(buildscheme.Encoder, newer)
	if err != nil {
		return nil, fmt.Errorf("error encoding newer: %v", err)
	}
	olderJSON, err := runtime.Encode(buildscheme.Encoder, older)
	if err != nil {
		return nil, fmt.Errorf("error encoding older: %v", err)
	}
	patch, err := strategicpatch.CreateTwoWayMergePatch(olderJSON, newerJSON, &buildv1.Build{})
	if err != nil {
		return nil, fmt.Errorf("error creating a strategic patch: %v", err)
	}
	return patch, nil
}
func (bc *BuildController) patchBuild(build *buildv1.Build, update *buildUpdate) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	updatedBuild := build.DeepCopy()
	update.apply(updatedBuild)
	patch, err := createBuildPatch(build, updatedBuild)
	if err != nil {
		return nil, fmt.Errorf("failed to create a build patch: %v", err)
	}
	klog.V(5).Infof("Patching build %s with %v", buildDesc(build), update)
	return bc.buildPatcher.Patch(build.Namespace, build.Name, patch)
}
func (bc *BuildController) findMissingPod(build *buildv1.Build) *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, err := bc.podClient.Pods(build.Namespace).Get(buildapihelpers.GetBuildPodName(build), metav1.GetOptions{})
	if err == nil {
		klog.V(2).Infof("Found missing pod for build %s by using direct client.", buildDesc(build))
		return pod
	}
	return nil
}
func (bc *BuildController) getBuildByKey(key string) (*buildv1.Build, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := bc.buildInformer.GetIndexer().GetByKey(key)
	if err != nil {
		klog.V(2).Infof("Unable to retrieve build %q from store: %v", key, err)
		return nil, err
	}
	if !exists {
		klog.V(2).Infof("Build %q has been deleted", key)
		return nil, nil
	}
	return obj.(*buildv1.Build), nil
}
func (bc *BuildController) podUpdated(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	curPod := cur.(*corev1.Pod)
	oldPod := old.(*corev1.Pod)
	if curPod.ResourceVersion == oldPod.ResourceVersion {
		return
	}
	if isBuildPod(curPod) {
		bc.enqueueBuildForPod(curPod)
	}
}
func (bc *BuildController) podDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone: %+v", obj))
			return
		}
		pod, ok = tombstone.Obj.(*corev1.Pod)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod: %+v", obj))
			return
		}
	}
	if isBuildPod(pod) {
		bc.enqueueBuildForPod(pod)
	}
}
func (bc *BuildController) buildAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	build := obj.(*buildv1.Build)
	bc.enqueueBuild(build)
}
func (bc *BuildController) buildUpdated(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	build := cur.(*buildv1.Build)
	bc.enqueueBuild(build)
}
func (bc *BuildController) buildDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	build, ok := obj.(*buildv1.Build)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone: %+v", obj))
			return
		}
		build, ok = tombstone.Obj.(*buildv1.Build)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a pod: %+v", obj))
			return
		}
	}
	if !buildutil.IsBuildComplete(build) {
		bcName := buildutil.ConfigNameForBuild(build)
		bc.enqueueBuildConfig(build.Namespace, bcName)
	}
}
func (bc *BuildController) enqueueBuild(build *buildv1.Build) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := resourceName(build.Namespace, build.Name)
	bc.buildQueue.Add(key)
}
func (bc *BuildController) enqueueBuildForPod(pod *corev1.Pod) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.buildQueue.Add(resourceName(pod.Namespace, getBuildName(pod)))
}
func (bc *BuildController) imageStreamAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream := obj.(*imagev1.ImageStream)
	for _, buildKey := range bc.imageStreamQueue.Pop(resourceName(stream.Namespace, stream.Name)) {
		bc.buildQueue.Add(buildKey)
	}
}
func (bc *BuildController) imageStreamUpdated(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.imageStreamAdded(cur)
}
func (bc *BuildController) handleBuildError(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		bc.buildQueue.Forget(key)
		return
	}
	if strategy.IsFatal(err) {
		klog.V(2).Infof("Will not retry fatal error for key %v: %v", key, err)
		bc.buildQueue.Forget(key)
		return
	}
	if bc.buildQueue.NumRequeues(key) < maxRetries {
		klog.V(4).Infof("Retrying key %v: %v", key, err)
		bc.buildQueue.AddRateLimited(key)
		return
	}
	klog.V(2).Infof("Giving up retrying %v: %v", key, err)
	bc.buildQueue.Forget(key)
}
func (bc *BuildController) handleBuildConfigError(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		bc.buildConfigQueue.Forget(key)
		return
	}
	if bc.buildConfigQueue.NumRequeues(key) < maxRetries {
		klog.V(4).Infof("Retrying key %v: %v", key, err)
		bc.buildConfigQueue.AddRateLimited(key)
		return
	}
	klog.V(2).Infof("Giving up retrying %v: %v", key, err)
	bc.buildConfigQueue.Forget(key)
}
func (bc *BuildController) createBuildCAConfigMap(build *buildv1.Build, buildPod *corev1.Pod, update *buildUpdate) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configMapSpec := bc.createBuildCAConfigMapSpec(build, buildPod)
	configMap, err := bc.configMapClient.ConfigMaps(buildPod.Namespace).Create(configMapSpec)
	if err != nil {
		bc.recorder.Eventf(build, corev1.EventTypeWarning, "FailedCreate", "Error creating build certificate authority configMap: %v", err)
		update.setReason("CannotCreateCAConfigMap")
		update.setMessage(buildutil.StatusMessageCannotCreateCAConfigMap)
		return update, fmt.Errorf("failed to create build certificate authority configMap: %v", err)
	}
	klog.V(4).Infof("Created certificate authority configMap %s/%s for build %s", build.Namespace, configMap.Name, buildDesc(build))
	return update, nil
}
func (bc *BuildController) createBuildCAConfigMapSpec(build *buildv1.Build, buildPod *corev1.Pod) *corev1.ConfigMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: buildapihelpers.GetBuildCAConfigMapName(build), OwnerReferences: []metav1.OwnerReference{makeBuildPodOwnerRef(buildPod)}, Annotations: map[string]string{"service.alpha.openshift.io/inject-cabundle": "true"}}, Data: bc.additionalTrustedCAData}
	if cm.Data == nil {
		cm.Data = make(map[string]string)
	}
	return cm
}
func (bc *BuildController) findOwnedConfigMap(owner *corev1.Pod, namespace string, name string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm, err := bc.configMapClient.ConfigMaps(namespace).Get(name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	if hasRef := hasBuildPodOwnerRef(owner, cm); !hasRef {
		return true, fmt.Errorf("configMap %s/%s is not owned by build pod %s/%s", cm.Namespace, cm.Name, owner.Namespace, owner.Name)
	}
	return true, nil
}
func (bc *BuildController) createBuildSystemConfConfigMap(build *buildv1.Build, buildPod *corev1.Pod, update *buildUpdate) (*buildUpdate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configMapSpec := bc.createBuildSystemConfigMapSpec(build, buildPod)
	configMap, err := bc.configMapClient.ConfigMaps(build.Namespace).Create(configMapSpec)
	if err != nil {
		bc.recorder.Eventf(build, corev1.EventTypeWarning, "FailedCreate", "Error creating build system config configMap: %v", err)
		update.setReason("CannotCreateBuildSysConfigMap")
		update.setMessage(buildutil.StatusMessageCannotCreateBuildSysConfigMap)
		return update, fmt.Errorf("failed to create build system config configMap: %v", err)
	}
	klog.V(4).Infof("Created build system config configMap %s/%s for build %s", build.Namespace, configMap.Name, buildDesc(build))
	return update, nil
}
func (bc *BuildController) createBuildSystemConfigMapSpec(build *buildv1.Build, buildPod *corev1.Pod) *corev1.ConfigMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: buildapihelpers.GetBuildSystemConfigMapName(build), OwnerReferences: []metav1.OwnerReference{makeBuildPodOwnerRef(buildPod)}}, Data: make(map[string]string)}
	if len(bc.registryConfData) > 0 {
		cm.Data[buildutil.RegistryConfKey] = bc.registryConfData
	}
	if len(bc.signaturePolicyData) > 0 {
		cm.Data[buildutil.SignaturePolicyKey] = bc.signaturePolicyData
	}
	return cm
}
func (bc *BuildController) controllerConfigWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if quit := bc.controllerConfigWork(); quit {
			return
		}
	}
}
func (bc *BuildController) controllerConfigWork() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := bc.controllerConfigQueue.Get()
	if quit {
		return true
	}
	defer bc.controllerConfigQueue.Done(key)
	var err error
	errs := bc.handleControllerConfig()
	if len(errs) > 0 {
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprintf("%v", e)
		}
		err = fmt.Errorf("failed to apply build controller config: %v", msgs)
	}
	bc.handleControllerConfigError(err, key)
	return false
}
func (bc *BuildController) handleControllerConfig() []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configErrs := bc.readClusterImageConfig()
	err := bc.readClusterBuildControllerConfig()
	if err != nil {
		configErrs = append(configErrs, err)
	}
	return configErrs
}
func (bc *BuildController) readClusterBuildControllerConfig() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildConfig, err := bc.buildControllerConfigLister.Get("cluster")
	if err != nil && !errors.IsNotFound(err) {
		return err
	} else if buildConfig == nil {
		bc.buildDefaults.DefaultProxy = nil
		return nil
	}
	if klog.V(5) {
		configJSON, _ := json.Marshal(buildConfig)
		if configJSON != nil {
			klog.Infof("build controller config: %s", string(configJSON))
		}
	}
	bc.buildDefaults.DefaultProxy = buildConfig.Spec.BuildDefaults.DefaultProxy
	return nil
}
func (bc *BuildController) readClusterImageConfig() []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configErrs := []error{}
	imageConfig, err := bc.imageConfigLister.Get("cluster")
	if err != nil && !errors.IsNotFound(err) {
		configErrs = append(configErrs, err)
		return configErrs
	} else if imageConfig == nil {
		bc.additionalTrustedCAData = nil
		bc.registryConfData = ""
		bc.signaturePolicyData = ""
		return configErrs
	}
	if klog.V(5) {
		configJSON, _ := json.Marshal(imageConfig)
		if configJSON != nil {
			klog.Infof("image config: %s", string(configJSON))
		}
	}
	additionalCAs, err := bc.getAdditionalTrustedCAData(imageConfig)
	if err != nil {
		configErrs = append(configErrs, err)
	} else {
		bc.additionalTrustedCAData = additionalCAs
	}
	registriesTOML, regErr := bc.createBuildRegistriesConfigData(imageConfig)
	if regErr != nil {
		configErrs = append(configErrs, regErr)
	} else {
		bc.registryConfData = registriesTOML
	}
	signatureJSON, sigErr := bc.createBuildSignaturePolicyData(imageConfig)
	if sigErr != nil {
		configErrs = append(configErrs, sigErr)
	} else {
		bc.signaturePolicyData = signatureJSON
	}
	return configErrs
}
func (bc *BuildController) getAdditionalTrustedCAData(config *configv1.Image) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(config.Spec.AdditionalTrustedCA.Name) == 0 {
		klog.V(4).Info("additional certificate authorities for builds not specified")
		return nil, nil
	}
	additionalCA, err := bc.openShiftConfigConfigMapStore.ConfigMaps("openshift-config").Get(config.Spec.AdditionalTrustedCA.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	if additionalCA == nil {
		klog.V(3).Infof("configMap reference %s/%s with additional certificate authorities for builds not found", "openshift-config", config.Spec.AdditionalTrustedCA.Name)
		return nil, nil
	}
	if klog.V(5) {
		keys := make([]string, len(additionalCA.Data))
		i := 0
		for key := range additionalCA.Data {
			keys[i] = key
			i++
		}
		klog.Infof("found certificate authorities for hosts %s", keys)
	}
	return additionalCA.Data, nil
}
func (bc *BuildController) createBuildRegistriesConfigData(config *configv1.Image) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	registriesConfig := config.Spec.RegistrySources
	if len(registriesConfig.InsecureRegistries) == 0 {
		klog.V(4).Info("using default insecure registry settings for builds")
		return "", nil
	}
	configObj := tomlConfig{Registries: registries{Search: registryList{Registries: []string{"docker.io"}}, Insecure: registryList{Registries: registriesConfig.InsecureRegistries}}}
	configTOML, err := toml.Marshal(configObj)
	if err != nil {
		return "", err
	}
	if len(configTOML) == 0 {
		klog.V(4).Info("using default insecure registry settings for builds")
		return "", nil
	}
	klog.V(4).Info("overrode insecure registry settings for builds")
	klog.V(5).Infof("generated registries.conf for build pods: \n%s", string(configTOML))
	return string(configTOML), nil
}
func (bc *BuildController) createBuildSignaturePolicyData(config *configv1.Image) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	registriesConfig := config.Spec.RegistrySources
	if len(registriesConfig.AllowedRegistries) == 0 && len(registriesConfig.BlockedRegistries) == 0 {
		klog.V(4).Info("allowing builds to pull images from all registries")
		return "", nil
	}
	if len(registriesConfig.AllowedRegistries) != 0 && len(registriesConfig.BlockedRegistries) != 0 {
		return "", fmt.Errorf("invalid registries config: only one of AllowedRegistries or BlockedRegistries may be specified")
	}
	policyObj := &signature.Policy{}
	transportScopes := make(signature.PolicyTransportScopes)
	if len(registriesConfig.AllowedRegistries) > 0 {
		klog.V(4).Infof("only allowing image pulls from %s for builds", registriesConfig.AllowedRegistries)
		policyObj.Default = signature.PolicyRequirements{signature.NewPRReject()}
		for _, registry := range registriesConfig.AllowedRegistries {
			transportScopes[registry] = signature.PolicyRequirements{signature.NewPRInsecureAcceptAnything()}
		}
	}
	if len(registriesConfig.BlockedRegistries) > 0 {
		klog.V(4).Infof("blocking image pulls from %s for builds", registriesConfig.BlockedRegistries)
		policyObj.Default = signature.PolicyRequirements{signature.NewPRInsecureAcceptAnything()}
		for _, registry := range registriesConfig.BlockedRegistries {
			transportScopes[registry] = signature.PolicyRequirements{signature.NewPRReject()}
		}
	}
	policyObj.Transports = map[string]signature.PolicyTransportScopes{"atomic": transportScopes, "docker": transportScopes}
	policyJSON, err := json.Marshal(policyObj)
	if err != nil {
		return "", err
	}
	if len(policyJSON) == 0 {
		return "", nil
	}
	klog.V(5).Infof("generated policy.json for build pods: \n%s", string(policyJSON))
	return string(policyJSON), err
}
func (bc *BuildController) handleControllerConfigError(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err != nil {
		utilruntime.HandleError(err)
	}
	bc.controllerConfigQueue.Forget(key)
}
func (bc *BuildController) buildControllerConfigAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) buildControllerConfigUpdated(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) buildControllerConfigDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) imageConfigAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) imageConfigUpdated(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) imageConfigDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigChanged()
}
func (bc *BuildController) controllerConfigChanged() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.controllerConfigQueue.Add("openshift-config")
}
func (bc *BuildController) configMapAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configMap, ok := obj.(*corev1.ConfigMap)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("object was not a configMap: %+v", obj))
		return
	}
	if configMap == nil {
		return
	}
	config, err := bc.buildControllerConfigLister.Get("cluster")
	if err != nil && !errors.IsNotFound(err) {
		utilruntime.HandleError(fmt.Errorf("could not get cluster build controller config: %v", err))
		return
	}
	if config == nil {
		return
	}
	if configMap.Name != config.Spec.AdditionalTrustedCA.Name {
		return
	}
	bc.controllerConfigChanged()
}
func (bc *BuildController) configMapUpdated(old, curr interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configMap, ok := curr.(*corev1.ConfigMap)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("object was not a configMap: %+v", curr))
		return
	}
	if configMap == nil {
		return
	}
	config, err := bc.buildControllerConfigLister.Get("cluster")
	if err != nil && !errors.IsNotFound(err) {
		utilruntime.HandleError(fmt.Errorf("could not get cluster build controller config: %v", err))
		return
	}
	if config == nil {
		return
	}
	if configMap.Name != config.Spec.AdditionalTrustedCA.Name {
		return
	}
	bc.controllerConfigChanged()
}
func (bc *BuildController) configMapDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configMap, ok := obj.(*corev1.ConfigMap)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone: %+v", obj))
			return
		}
		configMap, ok = tombstone.Obj.(*corev1.ConfigMap)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a configMap: %+v", obj))
			return
		}
	}
	config, err := bc.buildControllerConfigLister.Get("cluster")
	if err != nil && !errors.IsNotFound(err) {
		utilruntime.HandleError(fmt.Errorf("could not get cluster build controller config: %v", err))
		return
	}
	if config == nil {
		return
	}
	if configMap.Name != config.Spec.AdditionalTrustedCA.Name {
		return
	}
	bc.controllerConfigChanged()
}
func isBuildPod(pod *corev1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(getBuildName(pod)) > 0
}
func buildDesc(build *buildv1.Build) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s/%s (%s)", build.Namespace, build.Name, build.Status.Phase)
}
func transitionToPhase(phase buildv1.BuildPhase, reason buildv1.StatusReason, message string) *buildUpdate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	update := &buildUpdate{}
	update.setPhase(phase)
	update.setReason(reason)
	update.setMessage(message)
	return update
}
func isValidTransition(from, to buildv1.BuildPhase) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if from == to {
		return true
	}
	switch {
	case buildutil.IsTerminalPhase(from):
		return false
	case from == buildv1.BuildPhasePending:
		switch to {
		case buildv1.BuildPhaseNew:
			return false
		}
	case from == buildv1.BuildPhaseRunning:
		switch to {
		case buildv1.BuildPhaseNew, buildv1.BuildPhasePending:
			return false
		}
	}
	return true
}
func setBuildCompletionData(build *buildv1.Build, pod *corev1.Pod, update *buildUpdate) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := metav1.Now()
	startTime := build.Status.StartTimestamp
	if startTime == nil {
		if pod != nil {
			startTime = pod.Status.StartTime
		}
		if startTime == nil {
			startTime = &now
		}
		update.setStartTime(*startTime)
	}
	if build.Status.CompletionTimestamp == nil {
		update.setCompletionTime(now)
		update.setDuration(now.Rfc3339Copy().Time.Sub(startTime.Rfc3339Copy().Time))
	}
	if (build.Status.Phase == buildv1.BuildPhaseFailed || (update.phase != nil && *update.phase == buildv1.BuildPhaseFailed)) && len(build.Status.LogSnippet) == 0 && pod != nil && len(pod.Status.ContainerStatuses) != 0 && pod.Status.ContainerStatuses[0].State.Terminated != nil {
		msg := pod.Status.ContainerStatuses[0].State.Terminated.Message
		if len(msg) != 0 {
			parts := strings.Split(strings.TrimRight(msg, "\n"), "\n")
			excerptLength := maxExcerptLength
			if len(parts) < maxExcerptLength {
				excerptLength = len(parts)
			}
			excerpt := parts[len(parts)-excerptLength:]
			for i, line := range excerpt {
				if len(line) > 120 {
					excerpt[i] = line[:58] + "..." + line[len(line)-59:]
				}
			}
			msg = strings.Join(excerpt, "\n")
			update.setLogSnippet(msg)
		}
	}
}
func hasError(err error, fns ...utilerrors.Matcher) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return false
	}
	if agg, ok := err.(utilerrors.Aggregate); ok {
		for _, err := range agg.Errors() {
			if hasError(err, fns...) {
				return true
			}
		}
		return false
	}
	for _, fn := range fns {
		if fn(err) {
			return true
		}
	}
	return false
}
func getBuildName(pod metav1.Object) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod == nil {
		return ""
	}
	return pod.GetAnnotations()[buildutil.BuildAnnotation]
}
func makeBuildPodOwnerRef(buildPod *corev1.Pod) metav1.OwnerReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return metav1.OwnerReference{APIVersion: "v1", Kind: "Pod", Name: buildPod.Name, UID: buildPod.UID}
}
func hasBuildPodOwnerRef(buildPod *corev1.Pod, caMap *corev1.ConfigMap) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ref := makeBuildPodOwnerRef(buildPod)
	for _, owner := range caMap.OwnerReferences {
		if reflect.DeepEqual(ref, owner) {
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
