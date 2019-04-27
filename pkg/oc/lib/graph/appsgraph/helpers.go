package appsgraph

import (
	"fmt"
	"sort"
	corev1 "k8s.io/api/core/v1"
	appsv1 "github.com/openshift/api/apps/v1"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	appsgraph "github.com/openshift/origin/pkg/oc/lib/graph/appsgraph/nodes"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

func RelevantDeployments(g osgraph.Graph, dcNode *appsgraph.DeploymentConfigNode) (*kubegraph.ReplicationControllerNode, []*kubegraph.ReplicationControllerNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allDeployments := []*kubegraph.ReplicationControllerNode{}
	uncastDeployments := g.SuccessorNodesByEdgeKind(dcNode, DeploymentEdgeKind)
	if len(uncastDeployments) == 0 {
		return nil, []*kubegraph.ReplicationControllerNode{}
	}
	for i := range uncastDeployments {
		allDeployments = append(allDeployments, uncastDeployments[i].(*kubegraph.ReplicationControllerNode))
	}
	sort.Sort(RecentDeploymentReferences(allDeployments))
	if dcNode.DeploymentConfig.Status.LatestVersion == appsutil.DeploymentVersionFor(allDeployments[0].ReplicationController) {
		return allDeployments[0], allDeployments[1:]
	}
	return nil, allDeployments
}
func BelongsToDeploymentConfig(config *appsv1.DeploymentConfig, b *corev1.ReplicationController) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.Annotations != nil {
		return config.Name == appsutil.DeploymentConfigNameFor(b)
	}
	return false
}

type RecentDeploymentReferences []*kubegraph.ReplicationControllerNode

func (m RecentDeploymentReferences) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m RecentDeploymentReferences) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m RecentDeploymentReferences) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return appsutil.DeploymentVersionFor(m[i].ReplicationController) > appsutil.DeploymentVersionFor(m[j].ReplicationController)
}

type TemplateImage struct {
	Image		string
	Ref		*imageapi.DockerImageReference
	From		*corev1.ObjectReference
	Container	*corev1.Container
}

func templateImageForContainer(container *corev1.Container, triggerFn TriggeredByFunc) (TemplateImage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ref imageapi.DockerImageReference
	if trigger, ok := triggerFn(container); ok {
		trigger.Image = container.Image
		trigger.Container = container
		return trigger, nil
	}
	ref, err := imageapi.ParseDockerImageReference(container.Image)
	if err != nil {
		return TemplateImage{Image: container.Image, Container: container}, err
	}
	return TemplateImage{Image: container.Image, Ref: &ref, Container: container}, nil
}
func TemplateImageForContainer(pod *corev1.PodSpec, triggerFn TriggeredByFunc, containerName string) (TemplateImage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range pod.Containers {
		container := &pod.Containers[i]
		if container.Name != containerName {
			continue
		}
		return templateImageForContainer(container, triggerFn)
	}
	for i := range pod.InitContainers {
		container := &pod.InitContainers[i]
		if container.Name != containerName {
			continue
		}
		return templateImageForContainer(container, triggerFn)
	}
	return TemplateImage{}, fmt.Errorf("no container %q found", containerName)
}
func eachTemplateImage(container *corev1.Container, triggerFn TriggeredByFunc, fn func(TemplateImage, error)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	image, err := templateImageForContainer(container, triggerFn)
	fn(image, err)
}
func EachTemplateImage(pod *corev1.PodSpec, triggerFn TriggeredByFunc, fn func(TemplateImage, error)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range pod.Containers {
		eachTemplateImage(&pod.Containers[i], triggerFn, fn)
	}
	for i := range pod.InitContainers {
		eachTemplateImage(&pod.InitContainers[i], triggerFn, fn)
	}
}

type TriggeredByFunc func(container *corev1.Container) (TemplateImage, bool)

func IgnoreTriggers(container *corev1.Container) (TemplateImage, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return TemplateImage{}, false
}
func DeploymentConfigHasTrigger(config *appsv1.DeploymentConfig) TriggeredByFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(container *corev1.Container) (TemplateImage, bool) {
		for _, trigger := range config.Spec.Triggers {
			params := trigger.ImageChangeParams
			if params == nil {
				continue
			}
			for _, name := range params.ContainerNames {
				if container.Name == name {
					if len(params.From.Name) == 0 {
						continue
					}
					from := params.From
					if len(from.Namespace) == 0 {
						from.Namespace = config.Namespace
					}
					return TemplateImage{Image: container.Image, From: &from}, true
				}
			}
		}
		return TemplateImage{}, false
	}
}
