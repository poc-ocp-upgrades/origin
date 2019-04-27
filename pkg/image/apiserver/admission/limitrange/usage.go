package limitrange

import (
	"fmt"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	imagev1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

type InternalImageReferenceHandler func(imageReference string, inSpec, inStatus bool)

func GetImageStreamUsage(is *imageapi.ImageStream) corev1.ResourceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	specRefs := resource.NewQuantity(0, resource.DecimalSI)
	statusRefs := resource.NewQuantity(0, resource.DecimalSI)
	processImageStreamImages(is, false, func(ref string, inSpec, inStatus bool) {
		if inSpec {
			specRefs.Set(specRefs.Value() + 1)
		}
		if inStatus {
			statusRefs.Set(statusRefs.Value() + 1)
		}
	})
	return corev1.ResourceList{imagev1.ResourceImageStreamTags: *specRefs, imagev1.ResourceImageStreamImages: *statusRefs}
}
func processImageStreamImages(is *imageapi.ImageStream, specOnly bool, handler InternalImageReferenceHandler) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	type sources struct{ inSpec, inStatus bool }
	var statusReferences sets.String
	imageReferences := make(map[string]*sources)
	specReferences := gatherImagesFromImageStreamSpec(is)
	for ref := range specReferences {
		imageReferences[ref] = &sources{inSpec: true}
	}
	if !specOnly {
		statusReferences = gatherImagesFromImageStreamStatus(is)
		for ref := range statusReferences {
			if s, exists := imageReferences[ref]; exists {
				s.inStatus = true
			} else {
				imageReferences[ref] = &sources{inStatus: true}
			}
		}
	}
	for ref, s := range imageReferences {
		handler(ref, s.inSpec, s.inStatus)
	}
}
func gatherImagesFromImageStreamStatus(is *imageapi.ImageStream) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := sets.NewString()
	for _, history := range is.Status.Tags {
		for i := range history.Items {
			ref := history.Items[i].Image
			if len(ref) == 0 {
				continue
			}
			res.Insert(ref)
		}
	}
	return res
}
func gatherImagesFromImageStreamSpec(is *imageapi.ImageStream) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := sets.NewString()
	for _, tagRef := range is.Spec.Tags {
		if tagRef.From == nil {
			continue
		}
		ref, err := getImageReferenceForObjectReference(is.Namespace, tagRef.From)
		if err != nil {
			klog.V(4).Infof("could not process object reference: %v", err)
			continue
		}
		res.Insert(ref)
	}
	return res
}
func getImageReferenceForObjectReference(namespace string, objRef *kapi.ObjectReference) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch objRef.Kind {
	case "ImageStreamImage", "DockerImage":
		res, err := imageapi.ParseDockerImageReference(objRef.Name)
		if err != nil {
			return "", err
		}
		if objRef.Kind == "ImageStreamImage" {
			if res.Namespace == "" {
				res.Namespace = objRef.Namespace
			}
			if res.Namespace == "" {
				res.Namespace = namespace
			}
			if len(res.ID) == 0 {
				return "", fmt.Errorf("missing id in ImageStreamImage reference %q", objRef.Name)
			}
		} else {
			res = res.DockerClientDefaults()
		}
		return res.DaemonMinimal().Exact(), nil
	case "ImageStreamTag":
		isName, tag, err := imageapi.ParseImageStreamTagName(objRef.Name)
		if err != nil {
			return "", err
		}
		ns := namespace
		if len(objRef.Namespace) > 0 {
			ns = objRef.Namespace
		}
		return cache.MetaNamespaceKeyFunc(&metav1.ObjectMeta{Namespace: ns, Name: imageapi.JoinImageStreamTag(isName, tag)})
	}
	return "", fmt.Errorf("unsupported object reference kind %s", objRef.Kind)
}
