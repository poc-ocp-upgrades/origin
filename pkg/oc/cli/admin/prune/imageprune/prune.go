package imageprune

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/registry/api/errcode"
	gonum "github.com/gonum/graph"
	"k8s.io/klog"
	kappsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kerrapi "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/retry"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/library-go/pkg/image/reference"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageutil "github.com/openshift/origin/pkg/image/util"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	buildgraph "github.com/openshift/origin/pkg/oc/lib/graph/buildgraph/nodes"
	"github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	imagegraph "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph/nodes"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	ReferencedImageEdgeKind		= "ReferencedImage"
	WeakReferencedImageEdgeKind	= "WeakReferencedImage"
	ReferencedImageConfigEdgeKind	= "ReferencedImageConfig"
	ReferencedImageLayerEdgeKind	= "ReferencedImageLayer"
	ReferencedImageManifestEdgeKind	= "ReferencedImageManifest"
	defaultPruneImageWorkerCount	= 5
)

type RegistryClientFactoryFunc func() (*http.Client, error)
type ImagePrunerFactoryFunc func() (ImageDeleter, error)

func FakeRegistryClientFactory() (*http.Client, error) {
	_logClusterCodePath()
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
	return nil, nil
}

type pruneAlgorithm struct {
	keepYoungerThan		time.Time
	keepTagRevisions	int
	pruneOverSizeLimit	bool
	namespace		string
	allImages		bool
	pruneRegistry		bool
}
type ImageDeleter interface {
	DeleteImage(image *imagev1.Image) error
}
type ImageStreamDeleter interface {
	GetImageStream(stream *imagev1.ImageStream) (*imagev1.ImageStream, error)
	UpdateImageStream(stream *imagev1.ImageStream) (*imagev1.ImageStream, error)
	NotifyImageStreamPrune(stream *imagev1.ImageStream, updatedTags []string, deletedTags []string)
}
type BlobDeleter interface {
	DeleteBlob(registryClient *http.Client, registryURL *url.URL, blob string) error
}
type LayerLinkDeleter interface {
	DeleteLayerLink(registryClient *http.Client, registryURL *url.URL, repo, linkName string) error
}
type ManifestDeleter interface {
	DeleteManifest(registryClient *http.Client, registryURL *url.URL, repo, manifest string) error
}
type PrunerOptions struct {
	KeepYoungerThan		*time.Duration
	KeepTagRevisions	*int
	PruneOverSizeLimit	*bool
	AllImages		*bool
	PruneRegistry		*bool
	Namespace		string
	Images			*imagev1.ImageList
	ImageWatcher		watch.Interface
	Streams			*imagev1.ImageStreamList
	StreamWatcher		watch.Interface
	Pods			*corev1.PodList
	RCs			*corev1.ReplicationControllerList
	BCs			*buildv1.BuildConfigList
	Builds			*buildv1.BuildList
	DSs			*kappsv1.DaemonSetList
	Deployments		*kappsv1.DeploymentList
	DCs			*appsv1.DeploymentConfigList
	RSs			*kappsv1.ReplicaSetList
	LimitRanges		map[string][]*corev1.LimitRange
	DryRun			bool
	RegistryClientFactory	RegistryClientFactoryFunc
	RegistryURL		*url.URL
	IgnoreInvalidRefs	bool
	NumWorkers		int
}
type Pruner interface {
	Prune(imagePrunerFactory ImagePrunerFactoryFunc, streamPruner ImageStreamDeleter, layerLinkPruner LayerLinkDeleter, blobPruner BlobDeleter, manifestPruner ManifestDeleter) (deletions []Deletion, failures []Failure)
}
type pruner struct {
	g			genericgraph.Graph
	algorithm		pruneAlgorithm
	ignoreInvalidRefs	bool
	registryClientFactory	RegistryClientFactoryFunc
	registryURL		*url.URL
	imageWatcher		watch.Interface
	imageStreamWatcher	watch.Interface
	imageStreamLimits	map[string][]*corev1.LimitRange
	queue			*nodeItem
	processedImages		map[*imagegraph.ImageNode]*Job
	numWorkers		int
}

var _ Pruner = &pruner{}

func NewPruner(options PrunerOptions) (Pruner, kerrors.Aggregate) {
	_logClusterCodePath()
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
	klog.V(1).Infof("Creating image pruner with keepYoungerThan=%v, keepTagRevisions=%s, pruneOverSizeLimit=%s, allImages=%s", options.KeepYoungerThan, getValue(options.KeepTagRevisions), getValue(options.PruneOverSizeLimit), getValue(options.AllImages))
	algorithm := pruneAlgorithm{}
	if options.KeepYoungerThan != nil {
		algorithm.keepYoungerThan = metav1.Now().Add(-*options.KeepYoungerThan)
	}
	if options.KeepTagRevisions != nil {
		algorithm.keepTagRevisions = *options.KeepTagRevisions
	}
	if options.PruneOverSizeLimit != nil {
		algorithm.pruneOverSizeLimit = *options.PruneOverSizeLimit
	}
	algorithm.allImages = true
	if options.AllImages != nil {
		algorithm.allImages = *options.AllImages
	}
	algorithm.pruneRegistry = true
	if options.PruneRegistry != nil {
		algorithm.pruneRegistry = *options.PruneRegistry
	}
	algorithm.namespace = options.Namespace
	p := &pruner{algorithm: algorithm, ignoreInvalidRefs: options.IgnoreInvalidRefs, registryClientFactory: options.RegistryClientFactory, registryURL: options.RegistryURL, processedImages: make(map[*imagegraph.ImageNode]*Job), imageWatcher: options.ImageWatcher, imageStreamWatcher: options.StreamWatcher, imageStreamLimits: options.LimitRanges, numWorkers: options.NumWorkers}
	if p.numWorkers < 1 {
		p.numWorkers = defaultPruneImageWorkerCount
	}
	if err := p.buildGraph(options); err != nil {
		return nil, err
	}
	return p, nil
}
func (p *pruner) buildGraph(options PrunerOptions) kerrors.Aggregate {
	_logClusterCodePath()
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
	p.g = genericgraph.New()
	var errs []error
	errs = append(errs, p.addImagesToGraph(options.Images)...)
	errs = append(errs, p.addImageStreamsToGraph(options.Streams, options.LimitRanges)...)
	errs = append(errs, p.addPodsToGraph(options.Pods)...)
	errs = append(errs, p.addReplicationControllersToGraph(options.RCs)...)
	errs = append(errs, p.addBuildConfigsToGraph(options.BCs)...)
	errs = append(errs, p.addBuildsToGraph(options.Builds)...)
	errs = append(errs, p.addDaemonSetsToGraph(options.DSs)...)
	errs = append(errs, p.addDeploymentsToGraph(options.Deployments)...)
	errs = append(errs, p.addDeploymentConfigsToGraph(options.DCs)...)
	errs = append(errs, p.addReplicaSetsToGraph(options.RSs)...)
	return kerrors.NewAggregate(errs)
}
func getValue(option interface{}) string {
	_logClusterCodePath()
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
	if v := reflect.ValueOf(option); !v.IsNil() {
		return fmt.Sprintf("%v", v.Elem())
	}
	return "<nil>"
}
func (p *pruner) addImagesToGraph(images *imagev1.ImageList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range images.Items {
		image := &images.Items[i]
		klog.V(4).Infof("Adding image %q to graph", image.Name)
		imageNode := imagegraph.EnsureImageNode(p.g, image)
		if err := imageutil.ImageWithMetadata(image); err != nil {
			klog.V(1).Infof("Failed to read image metadata for image %s: %v", image.Name, err)
			errs = append(errs, err)
			continue
		}
		dockerImage, ok := image.DockerImageMetadata.Object.(*dockerv10.DockerImage)
		if !ok {
			klog.V(1).Infof("Failed to read image metadata for image %s", image.Name)
			errs = append(errs, fmt.Errorf("Failed to read image metadata for image %s", image.Name))
			continue
		}
		if image.DockerImageManifestMediaType == schema2.MediaTypeManifest && len(dockerImage.ID) > 0 {
			configName := dockerImage.ID
			klog.V(4).Infof("Adding image config %q to graph", configName)
			configNode := imagegraph.EnsureImageComponentConfigNode(p.g, configName)
			p.g.AddEdge(imageNode, configNode, ReferencedImageConfigEdgeKind)
		}
		for _, layer := range image.DockerImageLayers {
			klog.V(4).Infof("Adding image layer %q to graph", layer.Name)
			layerNode := imagegraph.EnsureImageComponentLayerNode(p.g, layer.Name)
			p.g.AddEdge(imageNode, layerNode, ReferencedImageLayerEdgeKind)
		}
		klog.V(4).Infof("Adding image manifest %q to graph", image.Name)
		manifestNode := imagegraph.EnsureImageComponentManifestNode(p.g, image.Name)
		p.g.AddEdge(imageNode, manifestNode, ReferencedImageManifestEdgeKind)
	}
	return errs
}
func (p *pruner) addImageStreamsToGraph(streams *imagev1.ImageStreamList, limits map[string][]*corev1.LimitRange) []error {
	_logClusterCodePath()
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
	for i := range streams.Items {
		stream := &streams.Items[i]
		klog.V(4).Infof("Examining ImageStream %s", getName(stream))
		oldImageRevisionReferenceKind := WeakReferencedImageEdgeKind
		if !p.algorithm.pruneOverSizeLimit && stream.CreationTimestamp.Time.After(p.algorithm.keepYoungerThan) {
			oldImageRevisionReferenceKind = ReferencedImageEdgeKind
		}
		klog.V(4).Infof("Adding ImageStream %s to graph", getName(stream))
		isNode := imagegraph.EnsureImageStreamNode(p.g, stream)
		imageStreamNode := isNode.(*imagegraph.ImageStreamNode)
		for _, tag := range stream.Status.Tags {
			istNode := imagegraph.EnsureImageStreamTagNode(p.g, makeISTagWithStream(stream, tag.Tag))
			for i, tagEvent := range tag.Items {
				imageNode := imagegraph.FindImage(p.g, tag.Items[i].Image)
				if imageNode == nil {
					klog.V(2).Infof("Unable to find image %q in graph (from tag=%q, revision=%d, dockerImageReference=%s) - skipping", tag.Items[i].Image, tag.Tag, tagEvent.Generation, tag.Items[i].DockerImageReference)
					continue
				}
				kind := oldImageRevisionReferenceKind
				if p.algorithm.pruneOverSizeLimit {
					if exceedsLimits(stream, imageNode.Image, limits) {
						kind = WeakReferencedImageEdgeKind
					} else {
						kind = ReferencedImageEdgeKind
					}
				} else {
					if i < p.algorithm.keepTagRevisions {
						kind = ReferencedImageEdgeKind
					}
				}
				if i == 0 {
					klog.V(4).Infof("Adding edge (kind=%s) from %q to %q", kind, istNode.UniqueName(), imageNode.UniqueName())
					p.g.AddEdge(istNode, imageNode, kind)
				}
				klog.V(4).Infof("Checking for existing strong reference from stream %s to image %s", getName(stream), imageNode.Image.Name)
				if edge := p.g.Edge(imageStreamNode, imageNode); edge != nil && p.g.EdgeKinds(edge).Has(ReferencedImageEdgeKind) {
					klog.V(4).Infof("Strong reference found")
					continue
				}
				klog.V(4).Infof("Adding edge (kind=%s) from %q to %q", kind, imageStreamNode.UniqueName(), imageNode.UniqueName())
				p.g.AddEdge(imageStreamNode, imageNode, kind)
				klog.V(4).Infof("Adding stream->(layer|config) references")
				for _, s := range p.g.From(imageNode) {
					cn, ok := s.(*imagegraph.ImageComponentNode)
					if !ok {
						continue
					}
					klog.V(4).Infof("Adding reference from stream %s to %s", getName(stream), cn.Describe())
					switch cn.Type {
					case imagegraph.ImageComponentTypeConfig:
						p.g.AddEdge(imageStreamNode, s, ReferencedImageConfigEdgeKind)
					case imagegraph.ImageComponentTypeLayer:
						p.g.AddEdge(imageStreamNode, s, ReferencedImageLayerEdgeKind)
					case imagegraph.ImageComponentTypeManifest:
						p.g.AddEdge(imageStreamNode, s, ReferencedImageManifestEdgeKind)
					default:
						utilruntime.HandleError(fmt.Errorf("internal error: unhandled image component type %q", cn.Type))
					}
				}
			}
		}
	}
	return nil
}
func exceedsLimits(is *imagev1.ImageStream, image *imagev1.Image, limits map[string][]*corev1.LimitRange) bool {
	_logClusterCodePath()
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
	limitRanges, ok := limits[is.Namespace]
	if !ok || len(limitRanges) == 0 {
		return false
	}
	if err := imageutil.ImageWithMetadata(image); err != nil {
		return false
	}
	dockerImage, ok := image.DockerImageMetadata.Object.(*dockerv10.DockerImage)
	if !ok {
		return false
	}
	imageSize := resource.NewQuantity(dockerImage.Size, resource.BinarySI)
	for _, limitRange := range limitRanges {
		if limitRange == nil {
			continue
		}
		for _, limit := range limitRange.Spec.Limits {
			if limit.Type != imagev1.LimitTypeImage {
				continue
			}
			limitQuantity, ok := limit.Max[corev1.ResourceStorage]
			if !ok {
				continue
			}
			if limitQuantity.Cmp(*imageSize) < 0 {
				klog.V(4).Infof("Image %s in stream %s exceeds limit %s: %v vs %v", image.Name, getName(is), limitRange.Name, *imageSize, limitQuantity)
				return true
			}
		}
	}
	return false
}
func (p *pruner) addPodsToGraph(pods *corev1.PodList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range pods.Items {
		pod := &pods.Items[i]
		desc := fmt.Sprintf("Pod %s", getName(pod))
		klog.V(4).Infof("Examining %s", desc)
		if pod.Status.Phase != corev1.PodRunning && pod.Status.Phase != corev1.PodPending {
			if !pod.CreationTimestamp.Time.After(p.algorithm.keepYoungerThan) {
				klog.V(4).Infof("Ignoring %s for image reference counting because it's not running/pending and is too old", desc)
				continue
			}
		}
		klog.V(4).Infof("Adding %s to graph", desc)
		podNode := kubegraph.EnsurePodNode(p.g, pod)
		errs = append(errs, p.addPodSpecToGraph(getRef(pod), &pod.Spec, podNode)...)
	}
	return errs
}
func (p *pruner) addPodSpecToGraph(referrer *corev1.ObjectReference, spec *corev1.PodSpec, predecessor gonum.Node) []error {
	_logClusterCodePath()
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
	var errs []error
	for j := range spec.Containers {
		container := spec.Containers[j]
		if len(strings.TrimSpace(container.Image)) == 0 {
			klog.V(4).Infof("Ignoring edge from %s because container has no reference to image", getKindName(referrer))
			continue
		}
		klog.V(4).Infof("Examining container image %q", container.Image)
		ref, err := reference.Parse(container.Image)
		if err != nil {
			klog.Warningf("Unable to parse DockerImageReference %q of %s: %v - skipping", container.Image, getKindName(referrer), err)
			if !p.ignoreInvalidRefs {
				errs = append(errs, newErrBadReferenceToImage(container.Image, referrer, err.Error()))
			}
			continue
		}
		if len(ref.ID) == 0 {
			ref = ref.DockerClientDefaults()
			klog.V(4).Infof("%q has no image ID", container.Image)
			node := p.g.Find(imagegraph.ImageStreamTagNodeName(makeISTag(ref.Namespace, ref.Name, ref.Tag)))
			if node == nil {
				klog.V(4).Infof("No image stream tag found for %q - skipping", container.Image)
				continue
			}
			for _, n := range p.g.From(node) {
				imgNode, ok := n.(*imagegraph.ImageNode)
				if !ok {
					continue
				}
				klog.V(4).Infof("Adding edge from pod to image %q referenced by %s:%s", imgNode.Image.Name, ref.RepositoryName(), ref.Tag)
				p.g.AddEdge(predecessor, imgNode, ReferencedImageEdgeKind)
			}
			continue
		}
		imageNode := imagegraph.FindImage(p.g, ref.ID)
		if imageNode == nil {
			klog.V(2).Infof("Unable to find image %q referenced by %s in the graph - skipping", ref.ID, getKindName(referrer))
			continue
		}
		klog.V(4).Infof("Adding edge from %s to image %v", getKindName(referrer), imageNode)
		p.g.AddEdge(predecessor, imageNode, ReferencedImageEdgeKind)
	}
	return errs
}
func (p *pruner) addReplicationControllersToGraph(rcs *corev1.ReplicationControllerList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range rcs.Items {
		rc := &rcs.Items[i]
		desc := fmt.Sprintf("ReplicationController %s", getName(rc))
		klog.V(4).Infof("Examining %s", desc)
		rcNode := kubegraph.EnsureReplicationControllerNode(p.g, rc)
		errs = append(errs, p.addPodSpecToGraph(getRef(rc), &rc.Spec.Template.Spec, rcNode)...)
	}
	return errs
}
func (p *pruner) addDaemonSetsToGraph(dss *kappsv1.DaemonSetList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range dss.Items {
		ds := &dss.Items[i]
		desc := fmt.Sprintf("DaemonSet %s", getName(ds))
		klog.V(4).Infof("Examining %s", desc)
		dsNode := kubegraph.EnsureDaemonSetNode(p.g, ds)
		errs = append(errs, p.addPodSpecToGraph(getRef(ds), &ds.Spec.Template.Spec, dsNode)...)
	}
	return errs
}
func (p *pruner) addDeploymentsToGraph(dmnts *kappsv1.DeploymentList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range dmnts.Items {
		d := &dmnts.Items[i]
		ref := getRef(d)
		klog.V(4).Infof("Examining %s", getKindName(ref))
		dNode := kubegraph.EnsureDeploymentNode(p.g, d)
		errs = append(errs, p.addPodSpecToGraph(ref, &d.Spec.Template.Spec, dNode)...)
	}
	return errs
}
func (p *pruner) addDeploymentConfigsToGraph(dcs *appsv1.DeploymentConfigList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range dcs.Items {
		dc := &dcs.Items[i]
		ref := getRef(dc)
		klog.V(4).Infof("Examining %s", getKindName(ref))
		dcNode := appsgraph.EnsureDeploymentConfigNode(p.g, dc)
		errs = append(errs, p.addPodSpecToGraph(getRef(dc), &dc.Spec.Template.Spec, dcNode)...)
	}
	return errs
}
func (p *pruner) addReplicaSetsToGraph(rss *kappsv1.ReplicaSetList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range rss.Items {
		rs := &rss.Items[i]
		ref := getRef(rs)
		klog.V(4).Infof("Examining %s", getKindName(ref))
		rsNode := kubegraph.EnsureReplicaSetNode(p.g, rs)
		errs = append(errs, p.addPodSpecToGraph(ref, &rs.Spec.Template.Spec, rsNode)...)
	}
	return errs
}
func (p *pruner) addBuildConfigsToGraph(bcs *buildv1.BuildConfigList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range bcs.Items {
		bc := &bcs.Items[i]
		ref := getRef(bc)
		klog.V(4).Infof("Examining %s", getKindName(ref))
		bcNode := buildgraph.EnsureBuildConfigNode(p.g, bc)
		errs = append(errs, p.addBuildStrategyImageReferencesToGraph(ref, bc.Spec.Strategy, bcNode)...)
	}
	return errs
}
func (p *pruner) addBuildsToGraph(builds *buildv1.BuildList) []error {
	_logClusterCodePath()
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
	var errs []error
	for i := range builds.Items {
		build := &builds.Items[i]
		ref := getRef(build)
		klog.V(4).Infof("Examining %s", getKindName(ref))
		buildNode := buildgraph.EnsureBuildNode(p.g, build)
		errs = append(errs, p.addBuildStrategyImageReferencesToGraph(ref, build.Spec.Strategy, buildNode)...)
	}
	return errs
}
func (p *pruner) resolveISTagName(g genericgraph.Graph, referrer *corev1.ObjectReference, istagName string) (*imagegraph.ImageStreamTagNode, error) {
	_logClusterCodePath()
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
	name, tag, err := imageapi.ParseImageStreamTagName(istagName)
	if err != nil {
		if p.ignoreInvalidRefs {
			klog.Warningf("Failed to parse ImageStreamTag name %q: %v", istagName, err)
			return nil, nil
		}
		return nil, newErrBadReferenceTo("ImageStreamTag", istagName, referrer, err.Error())
	}
	node := g.Find(imagegraph.ImageStreamTagNodeName(makeISTag(referrer.Namespace, name, tag)))
	if istNode, ok := node.(*imagegraph.ImageStreamTagNode); ok {
		return istNode, nil
	}
	return nil, nil
}
func (p *pruner) addBuildStrategyImageReferencesToGraph(referrer *corev1.ObjectReference, strategy buildv1.BuildStrategy, predecessor gonum.Node) []error {
	_logClusterCodePath()
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
	from := buildapihelpers.GetInputReference(strategy)
	if from == nil {
		klog.V(4).Infof("Unable to determine 'from' reference - skipping")
		return nil
	}
	klog.V(4).Infof("Examining build strategy with from: %#v", from)
	var imageID string
	switch from.Kind {
	case "DockerImage":
		if len(strings.TrimSpace(from.Name)) == 0 {
			klog.V(4).Infof("Ignoring edge from %s because build strategy has no reference to image", getKindName(referrer))
			return nil
		}
		ref, err := reference.Parse(from.Name)
		if err != nil {
			klog.Warningf("Failed to parse DockerImage name %q of %s: %v", from.Name, getKindName(referrer), err)
			if !p.ignoreInvalidRefs {
				return []error{newErrBadReferenceToImage(from.Name, referrer, err.Error())}
			}
			return nil
		}
		imageID = ref.ID
	case "ImageStreamImage":
		_, id, err := imageapi.ParseImageStreamImageName(from.Name)
		if err != nil {
			klog.Warningf("Failed to parse ImageStreamImage name %q of %s: %v", from.Name, getKindName(referrer), err)
			if !p.ignoreInvalidRefs {
				return []error{newErrBadReferenceTo("ImageStreamImage", from.Name, referrer, err.Error())}
			}
			return nil
		}
		imageID = id
	case "ImageStreamTag":
		istNode, err := p.resolveISTagName(p.g, referrer, from.Name)
		if err != nil {
			klog.V(4).Infof(err.Error())
			return []error{err}
		}
		if istNode == nil {
			klog.V(2).Infof("%s referenced by %s could not be found", getKindName(from), getKindName(referrer))
			return nil
		}
		for _, n := range p.g.From(istNode) {
			imgNode, ok := n.(*imagegraph.ImageNode)
			if !ok {
				continue
			}
			imageID = imgNode.Image.Name
			break
		}
		if len(imageID) == 0 {
			klog.V(4).Infof("No image referenced by %s found", getKindName(from))
			return nil
		}
	default:
		klog.V(4).Infof("Ignoring unrecognized source location %q in %s", getKindName(from), getKindName(referrer))
		return nil
	}
	klog.V(4).Infof("Looking for image %q in graph", imageID)
	imageNode := imagegraph.FindImage(p.g, imageID)
	if imageNode == nil {
		klog.V(2).Infof("Unable to find image %q in graph referenced by %s - skipping", imageID, getKindName(referrer))
		return nil
	}
	klog.V(4).Infof("Adding edge from %s to image %s", predecessor, imageNode.Image.Name)
	p.g.AddEdge(predecessor, imageNode, ReferencedImageEdgeKind)
	return nil
}
func (p *pruner) handleImageStreamEvent(event watch.Event) {
	_logClusterCodePath()
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
	getIsNode := func() (*imagev1.ImageStream, *imagegraph.ImageStreamNode) {
		is, ok := event.Object.(*imagev1.ImageStream)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("internal error: expected ImageStream object in %s event, not %T", event.Type, event.Object))
			return nil, nil
		}
		n := p.g.Find(imagegraph.ImageStreamNodeName(is))
		if isNode, ok := n.(*imagegraph.ImageStreamNode); ok {
			return is, isNode
		}
		return is, nil
	}
	switch event.Type {
	case watch.Added:
		is, isNode := getIsNode()
		if is == nil {
			return
		}
		if isNode != nil {
			klog.V(4).Infof("Ignoring added ImageStream %s that is already present in the graph", getName(is))
			return
		}
		klog.V(4).Infof("Adding ImageStream %s to the graph", getName(is))
		p.addImageStreamsToGraph(&imagev1.ImageStreamList{Items: []imagev1.ImageStream{*is}}, p.imageStreamLimits)
	case watch.Modified:
		is, isNode := getIsNode()
		if is == nil {
			return
		}
		if isNode != nil {
			klog.V(4).Infof("Removing updated ImageStream %s from the graph", getName(is))
			p.g.RemoveNode(isNode)
		}
		klog.V(4).Infof("Adding updated ImageStream %s back to the graph", getName(is))
		p.addImageStreamsToGraph(&imagev1.ImageStreamList{Items: []imagev1.ImageStream{*is}}, p.imageStreamLimits)
	}
}
func (p *pruner) handleImageEvent(event watch.Event) {
	_logClusterCodePath()
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
	getImageNode := func() (*imagev1.Image, *imagegraph.ImageNode) {
		img, ok := event.Object.(*imagev1.Image)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("internal error: expected Image object in %s event, not %T", event.Type, event.Object))
			return nil, nil
		}
		return img, imagegraph.FindImage(p.g, img.Name)
	}
	switch event.Type {
	case watch.Added:
		img, imgNode := getImageNode()
		if img == nil {
			return
		}
		if imgNode != nil {
			klog.V(4).Infof("Ignoring added Image %s that is already present in the graph", img.Name)
			return
		}
		klog.V(4).Infof("Adding new Image %s to the graph", img.Name)
		p.addImagesToGraph(&imagev1.ImageList{Items: []imagev1.Image{*img}})
	case watch.Deleted:
		img, imgNode := getImageNode()
		if imgNode == nil {
			klog.V(4).Infof("Ignoring event for deleted Image %s that is not present in the graph", img.Name)
			return
		}
		klog.V(4).Infof("Removing deleted image %s from the graph", img.Name)
		p.g.RemoveNode(imgNode)
	}
}
func getImageNodes(nodes []gonum.Node) map[string]*imagegraph.ImageNode {
	_logClusterCodePath()
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
	ret := make(map[string]*imagegraph.ImageNode)
	for i := range nodes {
		if node, ok := nodes[i].(*imagegraph.ImageNode); ok {
			ret[node.Image.Name] = node
		}
	}
	return ret
}
func edgeKind(g genericgraph.Graph, from, to gonum.Node, desiredKind string) bool {
	_logClusterCodePath()
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
	edge := g.Edge(from, to)
	kinds := g.EdgeKinds(edge)
	return kinds.Has(desiredKind)
}
func imageIsPrunable(g genericgraph.Graph, imageNode *imagegraph.ImageNode, algorithm pruneAlgorithm) bool {
	_logClusterCodePath()
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
	if !algorithm.allImages {
		if imageNode.Image.Annotations[imageapi.ManagedByOpenShiftAnnotation] != "true" {
			klog.V(4).Infof("Image %q with DockerImageReference %q belongs to an external registry - skipping", imageNode.Image.Name, imageNode.Image.DockerImageReference)
			return false
		}
	}
	if !algorithm.pruneOverSizeLimit && imageNode.Image.CreationTimestamp.Time.After(algorithm.keepYoungerThan) {
		klog.V(4).Infof("Image %q is younger than minimum pruning age", imageNode.Image.Name)
		return false
	}
	for _, n := range g.To(imageNode) {
		klog.V(4).Infof("Examining predecessor %#v", n)
		if edgeKind(g, n, imageNode, ReferencedImageEdgeKind) {
			klog.V(4).Infof("Strong reference detected")
			return false
		}
	}
	return true
}
func calculatePrunableImages(g genericgraph.Graph, imageNodes map[string]*imagegraph.ImageNode, algorithm pruneAlgorithm) []*imagegraph.ImageNode {
	_logClusterCodePath()
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
	prunable := []*imagegraph.ImageNode{}
	for _, imageNode := range imageNodes {
		klog.V(4).Infof("Examining image %q", imageNode.Image.Name)
		if imageIsPrunable(g, imageNode, algorithm) {
			klog.V(4).Infof("Image %q is prunable", imageNode.Image.Name)
			prunable = append(prunable, imageNode)
		}
	}
	return prunable
}
func pruneStreams(g genericgraph.Graph, prunableImageNodes []*imagegraph.ImageNode, streamPruner ImageStreamDeleter, keepYoungerThan time.Time) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
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
	imageNameToNode := map[string]*imagegraph.ImageNode{}
	for _, node := range prunableImageNodes {
		imageNameToNode[node.Image.Name] = node
	}
	noChangeErr := errors.New("nothing changed")
	klog.V(4).Infof("Removing pruned image references from streams")
	for _, node := range g.Nodes() {
		streamNode, ok := node.(*imagegraph.ImageStreamNode)
		if !ok {
			continue
		}
		streamName := getName(streamNode.ImageStream)
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			stream, err := streamPruner.GetImageStream(streamNode.ImageStream)
			if err != nil {
				if kerrapi.IsNotFound(err) {
					klog.V(4).Infof("Unable to get image stream %s: removed during prune", streamName)
					return noChangeErr
				}
				return err
			}
			updatedTags := sets.NewString()
			deletedTags := sets.NewString()
			for _, tag := range stream.Status.Tags {
				if updated, deleted := pruneISTagHistory(g, imageNameToNode, keepYoungerThan, streamName, stream, tag.Tag); deleted {
					deletedTags.Insert(tag.Tag)
				} else if updated {
					updatedTags.Insert(tag.Tag)
				}
			}
			if updatedTags.Len() == 0 && deletedTags.Len() == 0 {
				return noChangeErr
			}
			updatedStream, err := streamPruner.UpdateImageStream(stream)
			if err == nil {
				streamPruner.NotifyImageStreamPrune(stream, updatedTags.List(), deletedTags.List())
				streamNode.ImageStream = updatedStream
			}
			if kerrapi.IsNotFound(err) {
				klog.V(4).Infof("Unable to update image stream %s: removed during prune", streamName)
				return nil
			}
			return err
		})
		if err == noChangeErr {
			continue
		}
		if err != nil {
			failures = append(failures, Failure{Node: streamNode, Err: err})
		} else {
			deletions = append(deletions, Deletion{Node: streamNode})
		}
	}
	klog.V(4).Infof("Done removing pruned image references from streams")
	return
}
func strengthenReferencesFromFailedImageStreams(g genericgraph.Graph, failures []Failure) {
	_logClusterCodePath()
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
	for _, f := range failures {
		for _, n := range g.From(f.Node) {
			imageNode, ok := n.(*imagegraph.ImageNode)
			if !ok {
				continue
			}
			edge := g.Edge(f.Node, imageNode)
			if edge == nil {
				continue
			}
			kinds := g.EdgeKinds(edge)
			if kinds.Has(ReferencedImageEdgeKind) {
				continue
			}
			g.RemoveEdge(edge)
			g.AddEdge(f.Node, imageNode, ReferencedImageEdgeKind)
		}
	}
}
func pruneISTagHistory(g genericgraph.Graph, prunableImageNodes map[string]*imagegraph.ImageNode, keepYoungerThan time.Time, streamName string, imageStream *imagev1.ImageStream, tag string) (tagUpdated, tagDeleted bool) {
	_logClusterCodePath()
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
	history, _ := imageutil.StatusHasTag(imageStream, tag)
	newHistory := imagev1.NamedTagEventList{Tag: tag}
	for _, tagEvent := range history.Items {
		klog.V(4).Infof("Checking image stream tag %s:%s generation %d with image %q", streamName, tag, tagEvent.Generation, tagEvent.Image)
		if ok, reason := tagEventIsPrunable(tagEvent, g, prunableImageNodes, keepYoungerThan); ok {
			klog.V(4).Infof("Image stream tag %s:%s generation %d - removing because %s", streamName, tag, tagEvent.Generation, reason)
			tagUpdated = true
		} else {
			klog.V(4).Infof("Image stream tag %s:%s generation %d - keeping because %s", streamName, tag, tagEvent.Generation, reason)
			newHistory.Items = append(newHistory.Items, tagEvent)
		}
	}
	if len(newHistory.Items) == 0 {
		klog.V(4).Infof("Image stream tag %s:%s - removing empty tag", streamName, tag)
		tags := []imagev1.NamedTagEventList{}
		for i := range imageStream.Status.Tags {
			t := imageStream.Status.Tags[i]
			if t.Tag != tag {
				tags = append(tags, t)
			}
		}
		imageStream.Status.Tags = tags
		tagDeleted = true
		tagUpdated = false
	} else if tagUpdated {
		for i := range imageStream.Status.Tags {
			t := imageStream.Status.Tags[i]
			if t.Tag == tag {
				imageStream.Status.Tags[i] = newHistory
				break
			}
		}
	}
	return
}
func tagEventIsPrunable(tagEvent imagev1.TagEvent, g genericgraph.Graph, prunableImageNodes map[string]*imagegraph.ImageNode, keepYoungerThan time.Time) (ok bool, reason string) {
	_logClusterCodePath()
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
	if _, ok := prunableImageNodes[tagEvent.Image]; ok {
		return true, fmt.Sprintf("image %q matches deleted image", tagEvent.Image)
	}
	n := imagegraph.FindImage(g, tagEvent.Image)
	if n != nil {
		return false, fmt.Sprintf("image %q is not deleted", tagEvent.Image)
	}
	if n == nil && !tagEvent.Created.After(keepYoungerThan) {
		return true, fmt.Sprintf("image %q is absent", tagEvent.Image)
	}
	return false, "the tag event is younger than threshold"
}

type byLayerCountAndAge []*imagegraph.ImageNode

func (b byLayerCountAndAge) Len() int {
	_logClusterCodePath()
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
	return len(b)
}
func (b byLayerCountAndAge) Swap(i, j int) {
	_logClusterCodePath()
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
	b[i], b[j] = b[j], b[i]
}
func (b byLayerCountAndAge) Less(i, j int) bool {
	_logClusterCodePath()
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
	fst, snd := b[i].Image, b[j].Image
	if len(fst.DockerImageLayers) > len(snd.DockerImageLayers) {
		return true
	}
	if len(fst.DockerImageLayers) < len(snd.DockerImageLayers) {
		return false
	}
	return fst.CreationTimestamp.Before(&snd.CreationTimestamp) || (!snd.CreationTimestamp.Before(&fst.CreationTimestamp) && fst.Name < snd.Name)
}

type nodeItem struct {
	node		*imagegraph.ImageNode
	prev, next	*nodeItem
}

func (i *nodeItem) pop() (node *imagegraph.ImageNode, next *nodeItem) {
	_logClusterCodePath()
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
	n, p := i.next, i.prev
	if p != nil {
		p.next = n
	}
	if n != nil {
		n.prev = p
	}
	return i.node, n
}
func insertAfter(item *nodeItem, node *imagegraph.ImageNode) *nodeItem {
	_logClusterCodePath()
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
	newItem := &nodeItem{node: node, prev: item}
	if item != nil {
		if item.next != nil {
			item.next.prev = newItem
			newItem.next = item.next
		}
		item.next = newItem
	}
	return newItem
}
func makeQueue(nodes []*imagegraph.ImageNode) *nodeItem {
	_logClusterCodePath()
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
	var head, tail *nodeItem
	for i, n := range nodes {
		tail = insertAfter(tail, n)
		if i == 0 {
			head = tail
		}
	}
	return head
}
func (p *pruner) Prune(imagePrunerFactory ImagePrunerFactoryFunc, streamPruner ImageStreamDeleter, layerLinkPruner LayerLinkDeleter, blobPruner BlobDeleter, manifestPruner ManifestDeleter) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
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
	allNodes := p.g.Nodes()
	imageNodes := getImageNodes(allNodes)
	prunable := calculatePrunableImages(p.g, imageNodes, p.algorithm)
	deletions, failures = pruneStreams(p.g, prunable, streamPruner, p.algorithm.keepYoungerThan)
	if len(p.algorithm.namespace) > 0 || len(prunable) == 0 {
		return deletions, failures
	}
	strengthenReferencesFromFailedImageStreams(p.g, failures)
	sort.Sort(byLayerCountAndAge(prunable))
	p.queue = makeQueue(prunable)
	var (
		jobChan		= make(chan *Job)
		resultChan	= make(chan JobResult)
	)
	defer close(jobChan)
	for i := 0; i < p.numWorkers; i++ {
		worker, err := NewWorker(p.algorithm, p.registryClientFactory, p.registryURL, imagePrunerFactory, streamPruner, layerLinkPruner, blobPruner, manifestPruner)
		if err != nil {
			failures = append(failures, Failure{Err: fmt.Errorf("failed to initialize worker: %v", err)})
			return
		}
		go worker.Run(jobChan, resultChan)
	}
	ds, fs := p.runLoop(jobChan, resultChan)
	deletions = append(deletions, ds...)
	failures = append(failures, fs...)
	return
}
func (p *pruner) runLoop(jobChan chan<- *Job, resultChan <-chan JobResult) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
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
	imgUpdateChan := p.imageWatcher.ResultChan()
	isUpdateChan := p.imageStreamWatcher.ResultChan()
	for {
		for len(p.processedImages) < p.numWorkers {
			job, blocked := p.getNextJob()
			if blocked {
				break
			}
			if job == nil {
				if len(p.processedImages) == 0 {
					return
				}
				break
			}
			jobChan <- job
			p.processedImages[job.Image] = job
		}
		select {
		case res := <-resultChan:
			p.updateGraphWithResult(&res)
			for _, deletion := range res.Deletions {
				deletions = append(deletions, deletion)
			}
			for _, failure := range res.Failures {
				failures = append(failures, failure)
			}
			delete(p.processedImages, res.Job.Image)
		case <-isUpdateChan:
		case <-imgUpdateChan:
		}
	}
}
func (p *pruner) getNextJob() (job *Job, blocked bool) {
	_logClusterCodePath()
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
	if p.queue == nil {
		return
	}
	pop := func(item *nodeItem) (*imagegraph.ImageNode, *nodeItem) {
		node, next := item.pop()
		if item == p.queue {
			p.queue = next
		}
		return node, next
	}
	for item := p.queue; item != nil; {
		if !imageIsPrunable(p.g, item.node, p.algorithm) {
			_, item = pop(item)
			continue
		}
		if components, blocked := getImageComponents(p.g, p.processedImages, item.node); !blocked {
			job = &Job{Image: item.node, Components: components}
			_, item = pop(item)
			break
		}
		item = item.next
	}
	blocked = job == nil && p.queue != nil
	return
}
func (p *pruner) updateGraphWithResult(res *JobResult) {
	_logClusterCodePath()
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
	imageDeleted := false
	for _, d := range res.Deletions {
		switch d.Node.(type) {
		case *imagegraph.ImageNode:
			imageDeleted = true
			p.g.RemoveNode(d.Node)
		case *imagegraph.ImageComponentNode:
			if d.Parent == nil {
				p.g.RemoveNode(d.Node)
				continue
			}
			isn, ok := d.Parent.(*imagegraph.ImageStreamNode)
			if !ok {
				continue
			}
			edge := p.g.Edge(isn, d.Node)
			if edge == nil {
				continue
			}
			p.g.RemoveEdge(edge)
		case *imagegraph.ImageStreamNode:
		default:
			utilruntime.HandleError(fmt.Errorf("internal error: unhandled graph node %t", d.Node))
		}
	}
	if imageDeleted {
		return
	}
}
func getImageComponents(g genericgraph.Graph, processedImages map[*imagegraph.ImageNode]*Job, image *imagegraph.ImageNode) (components ComponentRetentions, blocked bool) {
	_logClusterCodePath()
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
	components = make(ComponentRetentions)
	for _, node := range g.From(image) {
		kinds := g.EdgeKinds(g.Edge(image, node))
		if len(kinds.Intersection(sets.NewString(ReferencedImageLayerEdgeKind, ReferencedImageConfigEdgeKind, ReferencedImageManifestEdgeKind))) == 0 {
			continue
		}
		imageStrongRefCounter := 0
		imageMarkedForDeletionCounter := 0
		referencingStreams := map[*imagegraph.ImageStreamNode]struct{}{}
		referencingImages := map[*imagegraph.ImageNode]struct{}{}
		comp, ok := node.(*imagegraph.ImageComponentNode)
		if !ok {
			continue
		}
		for _, ref := range g.To(comp) {
			switch t := ref.(type) {
			case (*imagegraph.ImageNode):
				imageStrongRefCounter++
				if _, processed := processedImages[t]; processed {
					imageMarkedForDeletionCounter++
				}
				referencingImages[t] = struct{}{}
			case *imagegraph.ImageStreamNode:
				referencingStreams[t] = struct{}{}
			default:
				continue
			}
		}
		switch {
		case imageStrongRefCounter < 2:
			components.Add(comp, true)
		case imageStrongRefCounter-imageMarkedForDeletionCounter < 2:
			return nil, true
		default:
			components.Add(comp, false)
		}
		if addComponentReferencingStreams(g, components, referencingImages, referencingStreams, processedImages, comp) {
			return nil, true
		}
	}
	return
}
func addComponentReferencingStreams(g genericgraph.Graph, components ComponentRetentions, referencingImages map[*imagegraph.ImageNode]struct{}, referencingStreams map[*imagegraph.ImageStreamNode]struct{}, processedImages map[*imagegraph.ImageNode]*Job, comp *imagegraph.ImageComponentNode) (blocked bool) {
	_logClusterCodePath()
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
streamLoop:
	for stream := range referencingStreams {
		refCounter := 0
		markedForDeletionCounter := 0
		for image := range referencingImages {
			edge := g.Edge(stream, image)
			if edge == nil {
				continue
			}
			kinds := g.EdgeKinds(edge)
			if kinds.Has(ReferencedImageEdgeKind) {
				components.AddReferencingStreams(comp, false, stream)
				continue streamLoop
			}
			if !kinds.Has(WeakReferencedImageEdgeKind) {
				continue
			}
			refCounter++
			if _, processed := processedImages[image]; processed {
				markedForDeletionCounter++
			}
			if refCounter-markedForDeletionCounter > 1 {
				components.AddReferencingStreams(comp, false, stream)
				continue streamLoop
			}
		}
		switch {
		case refCounter < 2:
			components.AddReferencingStreams(comp, true, stream)
		case refCounter-markedForDeletionCounter < 2:
			return true
		default:
			components.AddReferencingStreams(comp, false, stream)
		}
	}
	return false
}
func imageComponentIsPrunable(g genericgraph.Graph, cn *imagegraph.ImageComponentNode) bool {
	_logClusterCodePath()
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
	for _, predecessor := range g.To(cn) {
		klog.V(4).Infof("Examining predecessor %#v of image config %v", predecessor, cn)
		if g.Kind(predecessor) == imagegraph.ImageNodeKind {
			klog.V(4).Infof("Config %v has an image predecessor", cn)
			return false
		}
	}
	return true
}
func streamsReferencingImageComponent(g genericgraph.Graph, cn *imagegraph.ImageComponentNode) []*imagegraph.ImageStreamNode {
	_logClusterCodePath()
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
	ret := []*imagegraph.ImageStreamNode{}
	for _, predecessor := range g.To(cn) {
		if g.Kind(predecessor) != imagegraph.ImageStreamNodeKind {
			continue
		}
		ret = append(ret, predecessor.(*imagegraph.ImageStreamNode))
	}
	return ret
}

type imageDeleter struct{ images imagev1client.ImagesGetter }

var _ ImageDeleter = &imageDeleter{}

func NewImageDeleter(images imagev1client.ImagesGetter) ImageDeleter {
	_logClusterCodePath()
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
	return &imageDeleter{images: images}
}
func (p *imageDeleter) DeleteImage(image *imagev1.Image) error {
	_logClusterCodePath()
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
	klog.V(4).Infof("Deleting image %q", image.Name)
	return p.images.Images().Delete(image.Name, metav1.NewDeleteOptions(0))
}

type imageStreamDeleter struct {
	streams imagev1client.ImageStreamsGetter
}

var _ ImageStreamDeleter = &imageStreamDeleter{}

func NewImageStreamDeleter(streams imagev1client.ImageStreamsGetter) ImageStreamDeleter {
	_logClusterCodePath()
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
	return &imageStreamDeleter{streams: streams}
}
func (p *imageStreamDeleter) GetImageStream(stream *imagev1.ImageStream) (*imagev1.ImageStream, error) {
	_logClusterCodePath()
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
	return p.streams.ImageStreams(stream.Namespace).Get(stream.Name, metav1.GetOptions{})
}
func (p *imageStreamDeleter) UpdateImageStream(stream *imagev1.ImageStream) (*imagev1.ImageStream, error) {
	_logClusterCodePath()
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
	klog.V(4).Infof("Updating ImageStream %s", getName(stream))
	is, err := p.streams.ImageStreams(stream.Namespace).UpdateStatus(stream)
	if err == nil {
		klog.V(5).Infof("Updated ImageStream: %#v", is)
	}
	return is, err
}
func (p *imageStreamDeleter) NotifyImageStreamPrune(stream *imagev1.ImageStream, updatedTags []string, deletedTags []string) {
	_logClusterCodePath()
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
	return
}
func deleteFromRegistry(registryClient *http.Client, url string) error {
	_logClusterCodePath()
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
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	klog.V(5).Infof(`Sending request "%s %s" to the registry`, req.Method, req.URL.String())
	resp, err := registryClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		klog.Warningf("Unable to prune layer %s, returned %v", url, resp.Status)
		return nil
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(resp.Status)
	}
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusAccepted {
		klog.V(1).Infof("Unexpected status code in response: %d", resp.StatusCode)
		var response errcode.Errors
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&response); err != nil {
			return err
		}
		klog.V(1).Infof("Response: %#v", response)
		return &response
	}
	return err
}

type layerLinkDeleter struct{}

var _ LayerLinkDeleter = &layerLinkDeleter{}

func NewLayerLinkDeleter() LayerLinkDeleter {
	_logClusterCodePath()
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
	return &layerLinkDeleter{}
}
func (p *layerLinkDeleter) DeleteLayerLink(registryClient *http.Client, registryURL *url.URL, repoName, linkName string) error {
	_logClusterCodePath()
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
	klog.V(4).Infof("Deleting layer link %s from repository %s/%s", linkName, registryURL.Host, repoName)
	return deleteFromRegistry(registryClient, fmt.Sprintf("%s/v2/%s/blobs/%s", registryURL.String(), repoName, linkName))
}

type blobDeleter struct{}

var _ BlobDeleter = &blobDeleter{}

func NewBlobDeleter() BlobDeleter {
	_logClusterCodePath()
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
	return &blobDeleter{}
}
func (p *blobDeleter) DeleteBlob(registryClient *http.Client, registryURL *url.URL, blob string) error {
	_logClusterCodePath()
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
	klog.V(4).Infof("Deleting blob %s from registry %s", blob, registryURL.Host)
	return deleteFromRegistry(registryClient, fmt.Sprintf("%s/admin/blobs/%s", registryURL.String(), blob))
}

type manifestDeleter struct{}

var _ ManifestDeleter = &manifestDeleter{}

func NewManifestDeleter() ManifestDeleter {
	_logClusterCodePath()
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
	return &manifestDeleter{}
}
func (p *manifestDeleter) DeleteManifest(registryClient *http.Client, registryURL *url.URL, repoName, manifest string) error {
	_logClusterCodePath()
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
	klog.V(4).Infof("Deleting manifest %s from repository %s/%s", manifest, registryURL.Host, repoName)
	return deleteFromRegistry(registryClient, fmt.Sprintf("%s/v2/%s/manifests/%s", registryURL.String(), repoName, manifest))
}
func makeISTag(namespace, name, tag string) *imagev1.ImageStreamTag {
	_logClusterCodePath()
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
	return &imagev1.ImageStreamTag{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: imageapi.JoinImageStreamTag(name, tag)}}
}
func makeISTagWithStream(is *imagev1.ImageStream, tag string) *imagev1.ImageStreamTag {
	_logClusterCodePath()
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
	return makeISTag(is.Namespace, is.Name, tag)
}
