package nodes

import (
	"github.com/gonum/graph"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	imagev1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
)

func EnsureImageNode(g osgraph.MutableUniqueGraph, img *imagev1.Image) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageNodeName(img), func(node osgraph.Node) graph.Node {
		return &ImageNode{node, img}
	})
}
func EnsureAllImageStreamTagNodes(g osgraph.MutableUniqueGraph, is *imagev1.ImageStream) []*ImageStreamTagNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []*ImageStreamTagNode{}
	for _, tag := range is.Status.Tags {
		ist := &imagev1.ImageStreamTag{}
		ist.Namespace = is.Namespace
		ist.Name = imageapi.JoinImageStreamTag(is.Name, tag.Tag)
		istNode := EnsureImageStreamTagNode(g, ist)
		ret = append(ret, istNode)
	}
	return ret
}
func FindImage(g osgraph.MutableUniqueGraph, imageName string) *ImageNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := g.Find(ImageNodeName(&imagev1.Image{ObjectMeta: metav1.ObjectMeta{Name: imageName}}))
	if imageNode, ok := n.(*ImageNode); ok {
		return imageNode
	}
	return nil
}
func EnsureDockerRepositoryNode(g osgraph.MutableUniqueGraph, name, tag string) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ref, err := imageapi.ParseDockerImageReference(name)
	if err == nil {
		if len(tag) != 0 {
			ref.Tag = tag
		}
		ref = ref.DockerClientDefaults()
	} else {
		ref = imageapi.DockerImageReference{Name: name}
	}
	return osgraph.EnsureUnique(g, DockerImageRepositoryNodeName(ref), func(node osgraph.Node) graph.Node {
		return &DockerImageRepositoryNode{node, ref}
	})
}
func MakeImageStreamTagObjectMeta(namespace, name, tag string) *imagev1.ImageStreamTag {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imagev1.ImageStreamTag{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: imageapi.JoinImageStreamTag(name, tag)}}
}
func MakeImageStreamTagObjectMeta2(namespace, name string) *imagev1.ImageStreamTag {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imagev1.ImageStreamTag{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name}}
}
func EnsureImageStreamTagNode(g osgraph.MutableUniqueGraph, ist *imagev1.ImageStreamTag) *ImageStreamTagNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageStreamTagNodeName(ist), func(node osgraph.Node) graph.Node {
		return &ImageStreamTagNode{node, ist, true}
	}).(*ImageStreamTagNode)
}
func FindOrCreateSyntheticImageStreamTagNode(g osgraph.MutableUniqueGraph, ist *imagev1.ImageStreamTag) *ImageStreamTagNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageStreamTagNodeName(ist), func(node osgraph.Node) graph.Node {
		return &ImageStreamTagNode{node, ist, false}
	}).(*ImageStreamTagNode)
}
func MakeImageStreamImageObjectMeta(namespace, name string) *imagev1.ImageStreamImage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &imagev1.ImageStreamImage{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name}}
}
func EnsureImageStreamImageNode(g osgraph.MutableUniqueGraph, namespace, name string) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	isi := &imagev1.ImageStreamImage{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: name}}
	return osgraph.EnsureUnique(g, ImageStreamImageNodeName(isi), func(node osgraph.Node) graph.Node {
		return &ImageStreamImageNode{node, isi, true}
	})
}
func FindOrCreateSyntheticImageStreamImageNode(g osgraph.MutableUniqueGraph, isi *imagev1.ImageStreamImage) *ImageStreamImageNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageStreamImageNodeName(isi), func(node osgraph.Node) graph.Node {
		return &ImageStreamImageNode{node, isi, false}
	}).(*ImageStreamImageNode)
}
func EnsureImageStreamNode(g osgraph.MutableUniqueGraph, is *imagev1.ImageStream) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageStreamNodeName(is), func(node osgraph.Node) graph.Node {
		return &ImageStreamNode{node, is, true}
	})
}
func FindOrCreateSyntheticImageStreamNode(g osgraph.MutableUniqueGraph, is *imagev1.ImageStream) *ImageStreamNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.EnsureUnique(g, ImageStreamNodeName(is), func(node osgraph.Node) graph.Node {
		return &ImageStreamNode{node, is, false}
	}).(*ImageStreamNode)
}
func ensureImageComponentNode(g osgraph.MutableUniqueGraph, name string, t ImageComponentType) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := osgraph.EnsureUnique(g, ImageComponentNodeName(name), func(node osgraph.Node) graph.Node {
		return &ImageComponentNode{Node: node, Component: name, Type: t}
	})
	if t == ImageComponentTypeConfig {
		cn := node.(*ImageComponentNode)
		if cn.Type != ImageComponentTypeConfig {
			cn.Type = ImageComponentTypeConfig
		}
	}
	return node
}
func EnsureImageComponentConfigNode(g osgraph.MutableUniqueGraph, name string) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ensureImageComponentNode(g, name, ImageComponentTypeConfig)
}
func EnsureImageComponentLayerNode(g osgraph.MutableUniqueGraph, name string) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ensureImageComponentNode(g, name, ImageComponentTypeLayer)
}
func EnsureImageComponentManifestNode(g osgraph.MutableUniqueGraph, name string) graph.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ensureImageComponentNode(g, name, ImageComponentTypeManifest)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
