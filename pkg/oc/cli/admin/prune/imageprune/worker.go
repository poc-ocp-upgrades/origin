package imageprune

import (
	"fmt"
	"net/http"
	"net/url"
	gonum "github.com/gonum/graph"
	"k8s.io/klog"
	kerrapi "k8s.io/apimachinery/pkg/api/errors"
	imagegraph "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph/nodes"
)

type ComponentRetention struct {
	ReferencingStreams	map[*imagegraph.ImageStreamNode]bool
	PrunableGlobally	bool
}
type ComponentRetentions map[*imagegraph.ImageComponentNode]*ComponentRetention

func (cr ComponentRetentions) add(comp *imagegraph.ImageComponentNode) *ComponentRetention {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := cr[comp]; ok {
		return cr[comp]
	}
	cr[comp] = &ComponentRetention{ReferencingStreams: make(map[*imagegraph.ImageStreamNode]bool)}
	return cr[comp]
}
func (cr ComponentRetentions) Add(comp *imagegraph.ImageComponentNode, globallyPrunable bool) *ComponentRetention {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := cr.add(comp)
	r.PrunableGlobally = globallyPrunable
	return r
}
func (cr ComponentRetentions) AddReferencingStreams(comp *imagegraph.ImageComponentNode, prunable bool, streams ...*imagegraph.ImageStreamNode) *ComponentRetention {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := cr.add(comp)
	for _, n := range streams {
		r.ReferencingStreams[n] = prunable
	}
	return r
}

type Job struct {
	Image		*imagegraph.ImageNode
	Components	ComponentRetentions
}

func enumerateImageComponents(crs ComponentRetentions, compType *imagegraph.ImageComponentType, withPreserved bool, handler func(comp *imagegraph.ImageComponentNode, prunable bool)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for c, retention := range crs {
		if !withPreserved && !retention.PrunableGlobally {
			continue
		}
		if compType != nil && c.Type != *compType {
			continue
		}
		handler(c, retention.PrunableGlobally)
	}
}
func enumerateImageStreamComponents(crs ComponentRetentions, compType *imagegraph.ImageComponentType, withPreserved bool, handler func(comp *imagegraph.ImageComponentNode, stream *imagegraph.ImageStreamNode, prunable bool)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for c, cr := range crs {
		if compType != nil && c.Type != *compType {
			continue
		}
		for s, prunable := range cr.ReferencingStreams {
			if withPreserved || prunable {
				handler(c, s, prunable)
			}
		}
	}
}

type Deletion struct {
	Node	gonum.Node
	Parent	gonum.Node
}
type Failure struct {
	Node	gonum.Node
	Parent	gonum.Node
	Err	error
}

var _ error = &Failure{}

func (pf *Failure) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pf.String()
}
func (pf *Failure) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pf.Node == nil {
		return fmt.Sprintf("failed to prune blob: %v", pf.Err)
	}
	switch t := pf.Node.(type) {
	case *imagegraph.ImageStreamNode:
		return fmt.Sprintf("failed to update ImageStream %s: %v", getName(t.ImageStream), pf.Err)
	case *imagegraph.ImageNode:
		return fmt.Sprintf("failed to delete Image %s: %v", t.Image.DockerImageReference, pf.Err)
	case *imagegraph.ImageComponentNode:
		detail := ""
		if isn, ok := pf.Parent.(*imagegraph.ImageStreamNode); ok {
			detail = " in repository " + getName(isn.ImageStream)
		}
		switch t.Type {
		case imagegraph.ImageComponentTypeConfig:
			return fmt.Sprintf("failed to delete image config link %s%s: %v", t.Component, detail, pf.Err)
		case imagegraph.ImageComponentTypeLayer:
			return fmt.Sprintf("failed to delete image layer link %s%s: %v", t.Component, detail, pf.Err)
		case imagegraph.ImageComponentTypeManifest:
			return fmt.Sprintf("failed to delete image manifest link %s%s: %v", t.Component, detail, pf.Err)
		default:
			return fmt.Sprintf("failed to delete %s%s: %v", t.String(), detail, pf.Err)
		}
	default:
		return fmt.Sprintf("failed to delete %v: %v", t, pf.Err)
	}
}

type JobResult struct {
	Job		*Job
	Deletions	[]Deletion
	Failures	[]Failure
}

func (jr *JobResult) update(deletions []Deletion, failures []Failure) *JobResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	jr.Deletions = append(jr.Deletions, deletions...)
	jr.Failures = append(jr.Failures, failures...)
	return jr
}

type Worker interface {
	Run(in <-chan *Job, out chan<- JobResult)
}
type worker struct {
	algorithm	pruneAlgorithm
	registryClient	*http.Client
	registryURL	*url.URL
	imagePruner	ImageDeleter
	streamPruner	ImageStreamDeleter
	layerLinkPruner	LayerLinkDeleter
	blobPruner	BlobDeleter
	manifestPruner	ManifestDeleter
}

var _ Worker = &worker{}

func NewWorker(algorithm pruneAlgorithm, registryClientFactory RegistryClientFactoryFunc, registryURL *url.URL, imagePrunerFactory ImagePrunerFactoryFunc, streamPruner ImageStreamDeleter, layerLinkPruner LayerLinkDeleter, blobPruner BlobDeleter, manifestPruner ManifestDeleter) (Worker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := registryClientFactory()
	if err != nil {
		return nil, err
	}
	imagePruner, err := imagePrunerFactory()
	if err != nil {
		return nil, err
	}
	return &worker{algorithm: algorithm, registryClient: client, registryURL: registryURL, imagePruner: imagePruner, streamPruner: streamPruner, layerLinkPruner: layerLinkPruner, blobPruner: blobPruner, manifestPruner: manifestPruner}, nil
}
func (w *worker) Run(in <-chan *Job, out chan<- JobResult) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		job, more := <-in
		if !more {
			return
		}
		out <- *w.prune(job)
	}
}
func (w *worker) prune(job *Job) *JobResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := &JobResult{Job: job}
	blobDeletions, blobFailures := []Deletion{}, []Failure{}
	if w.algorithm.pruneRegistry {
		res.update(pruneImageComponents(w.registryClient, w.registryURL, job.Components, w.layerLinkPruner))
		blobDeletions, blobFailures = pruneBlobs(w.registryClient, w.registryURL, job.Components, w.blobPruner)
		res.update(blobDeletions, blobFailures)
		res.update(pruneManifests(w.registryClient, w.registryURL, job.Components, w.manifestPruner))
	}
	if len(blobDeletions) > 0 || len(blobFailures) == 0 {
		res.update(pruneImages(job.Image, w.imagePruner))
	}
	return res
}
func pruneImages(imageNode *imagegraph.ImageNode, imagePruner ImageDeleter) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := imagePruner.DeleteImage(imageNode.Image)
	if err != nil {
		if kerrapi.IsNotFound(err) {
			klog.V(2).Infof("Skipping image %s that no longer exists", imageNode.Image.Name)
		} else {
			failures = append(failures, Failure{Node: imageNode, Err: err})
		}
	} else {
		deletions = append(deletions, Deletion{Node: imageNode})
	}
	return
}
func pruneImageComponents(registryClient *http.Client, registryURL *url.URL, crs ComponentRetentions, layerLinkDeleter LayerLinkDeleter) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	enumerateImageStreamComponents(crs, nil, false, func(comp *imagegraph.ImageComponentNode, stream *imagegraph.ImageStreamNode, _ bool) {
		if comp.Type == imagegraph.ImageComponentTypeManifest {
			return
		}
		streamName := getName(stream.ImageStream)
		klog.V(4).Infof("Pruning repository %s/%s: %s", registryURL.Host, streamName, comp.Describe())
		err := layerLinkDeleter.DeleteLayerLink(registryClient, registryURL, streamName, comp.Component)
		if err != nil {
			failures = append(failures, Failure{Node: comp, Parent: stream, Err: err})
		} else {
			deletions = append(deletions, Deletion{Node: comp, Parent: stream})
		}
	})
	return
}
func pruneBlobs(registryClient *http.Client, registryURL *url.URL, crs ComponentRetentions, blobPruner BlobDeleter) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	enumerateImageComponents(crs, nil, false, func(comp *imagegraph.ImageComponentNode, prunable bool) {
		err := blobPruner.DeleteBlob(registryClient, registryURL, comp.Component)
		if err != nil {
			failures = append(failures, Failure{Node: comp, Err: err})
		} else {
			deletions = append(deletions, Deletion{Node: comp})
		}
	})
	return
}
func pruneManifests(registryClient *http.Client, registryURL *url.URL, crs ComponentRetentions, manifestPruner ManifestDeleter) (deletions []Deletion, failures []Failure) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	manifestType := imagegraph.ImageComponentTypeManifest
	enumerateImageStreamComponents(crs, &manifestType, false, func(manifestNode *imagegraph.ImageComponentNode, stream *imagegraph.ImageStreamNode, _ bool) {
		repoName := getName(stream.ImageStream)
		klog.V(4).Infof("Pruning manifest %s in the repository %s/%s", manifestNode.Component, registryURL.Host, repoName)
		err := manifestPruner.DeleteManifest(registryClient, registryURL, repoName, manifestNode.Component)
		if err != nil {
			failures = append(failures, Failure{Node: manifestNode, Parent: stream, Err: err})
		} else {
			deletions = append(deletions, Deletion{Node: manifestNode, Parent: stream})
		}
	})
	return
}
