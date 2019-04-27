package analysis

import (
	"fmt"
	"github.com/gonum/graph"
	"github.com/openshift/origin/pkg/oc/lib/graph/appsgraph"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
	kubeedges "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph"
	kubegraph "github.com/openshift/origin/pkg/oc/lib/graph/kubegraph/nodes"
)

const (
	UnmountableSecretWarning	= "UnmountableSecret"
	MissingSecretWarning		= "MissingSecret"
	MissingLivenessProbeWarning	= "MissingLivenessProbe"
)

func FindUnmountableSecrets(g osgraph.Graph, f osgraph.Namer) []osgraph.Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	markers := []osgraph.Marker{}
	for _, uncastPodSpecNode := range g.NodesByKind(kubegraph.PodSpecNodeKind) {
		podSpecNode := uncastPodSpecNode.(*kubegraph.PodSpecNode)
		unmountableSecrets := CheckForUnmountableSecrets(g, podSpecNode)
		topLevelNode := osgraph.GetTopLevelContainerNode(g, podSpecNode)
		topLevelString := f.ResourceName(topLevelNode)
		saString := "MISSING_SA"
		saNodes := g.SuccessorNodesByEdgeKind(podSpecNode, kubeedges.ReferencedServiceAccountEdgeKind)
		if len(saNodes) > 0 {
			saString = f.ResourceName(saNodes[0])
		}
		for _, unmountableSecret := range unmountableSecrets {
			markers = append(markers, osgraph.Marker{Node: podSpecNode, RelatedNodes: []graph.Node{unmountableSecret}, Severity: osgraph.WarningSeverity, Key: UnmountableSecretWarning, Message: fmt.Sprintf("%s is attempting to mount a secret %s disallowed by %s", topLevelString, f.ResourceName(unmountableSecret), saString)})
		}
	}
	return markers
}
func FindMissingSecrets(g osgraph.Graph, f osgraph.Namer) []osgraph.Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	markers := []osgraph.Marker{}
	for _, uncastPodSpecNode := range g.NodesByKind(kubegraph.PodSpecNodeKind) {
		podSpecNode := uncastPodSpecNode.(*kubegraph.PodSpecNode)
		missingSecrets := CheckMissingMountedSecrets(g, podSpecNode)
		topLevelNode := osgraph.GetTopLevelContainerNode(g, podSpecNode)
		topLevelString := f.ResourceName(topLevelNode)
		for _, missingSecret := range missingSecrets {
			markers = append(markers, osgraph.Marker{Node: podSpecNode, RelatedNodes: []graph.Node{missingSecret}, Severity: osgraph.WarningSeverity, Key: UnmountableSecretWarning, Message: fmt.Sprintf("%s is attempting to mount a missing secret %s", topLevelString, f.ResourceName(missingSecret))})
		}
	}
	return markers
}
func FindMissingLivenessProbes(g osgraph.Graph, f osgraph.Namer, setProbeCommand string) []osgraph.Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	markers := []osgraph.Marker{}
	for _, uncastPodSpecNode := range g.NodesByKind(kubegraph.PodSpecNodeKind) {
		podSpecNode := uncastPodSpecNode.(*kubegraph.PodSpecNode)
		if hasLivenessProbe(podSpecNode) {
			continue
		}
		topLevelNode := osgraph.GetTopLevelContainerNode(g, podSpecNode)
		if hasControllerRefEdge(g, topLevelNode) {
			continue
		}
		if hasControllerOwnerReference(topLevelNode) {
			continue
		}
		topLevelString := f.ResourceName(topLevelNode)
		markers = append(markers, osgraph.Marker{Node: podSpecNode, RelatedNodes: []graph.Node{topLevelNode}, Severity: osgraph.InfoSeverity, Key: MissingLivenessProbeWarning, Message: fmt.Sprintf("%s has no liveness probe to verify pods are still running.", topLevelString), Suggestion: osgraph.Suggestion(fmt.Sprintf("%s %s --liveness ...", setProbeCommand, topLevelString))})
	}
	return markers
}
func hasLivenessProbe(podSpecNode *kubegraph.PodSpecNode) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, container := range podSpecNode.PodSpec.Containers {
		if container.LivenessProbe != nil {
			return true
		}
	}
	return false
}
func hasControllerOwnerReference(node graph.Node) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, ok := node.(*kubegraph.PodNode)
	if !ok {
		return false
	}
	for _, ref := range pod.OwnerReferences {
		if ref.Controller != nil && *ref.Controller == true {
			return true
		}
	}
	return false
}
func hasControllerRefEdge(g osgraph.Graph, node graph.Node) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	managedEdges := g.OutboundEdges(node, appsgraph.ManagedByControllerEdgeKind)
	return len(managedEdges) > 0
}
func CheckForUnmountableSecrets(g osgraph.Graph, podSpecNode *kubegraph.PodSpecNode) []*kubegraph.SecretNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	saNodes := g.SuccessorNodesByNodeAndEdgeKind(podSpecNode, kubegraph.ServiceAccountNodeKind, kubeedges.ReferencedServiceAccountEdgeKind)
	saMountableSecrets := []*kubegraph.SecretNode{}
	if len(saNodes) > 0 {
		saNode := saNodes[0].(*kubegraph.ServiceAccountNode)
		for _, secretNode := range g.SuccessorNodesByNodeAndEdgeKind(saNode, kubegraph.SecretNodeKind, kubeedges.MountableSecretEdgeKind) {
			saMountableSecrets = append(saMountableSecrets, secretNode.(*kubegraph.SecretNode))
		}
	}
	unmountableSecrets := []*kubegraph.SecretNode{}
	for _, uncastMountedSecretNode := range g.SuccessorNodesByNodeAndEdgeKind(podSpecNode, kubegraph.SecretNodeKind, kubeedges.MountedSecretEdgeKind) {
		mountedSecretNode := uncastMountedSecretNode.(*kubegraph.SecretNode)
		mountable := false
		for _, mountableSecretNode := range saMountableSecrets {
			if mountableSecretNode == mountedSecretNode {
				mountable = true
				break
			}
		}
		if !mountable {
			unmountableSecrets = append(unmountableSecrets, mountedSecretNode)
			continue
		}
	}
	return unmountableSecrets
}
func CheckMissingMountedSecrets(g osgraph.Graph, podSpecNode *kubegraph.PodSpecNode) []*kubegraph.SecretNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	missingSecrets := []*kubegraph.SecretNode{}
	for _, uncastMountedSecretNode := range g.SuccessorNodesByNodeAndEdgeKind(podSpecNode, kubegraph.SecretNodeKind, kubeedges.MountedSecretEdgeKind) {
		mountedSecretNode := uncastMountedSecretNode.(*kubegraph.SecretNode)
		if !mountedSecretNode.Found() {
			missingSecrets = append(missingSecrets, mountedSecretNode)
		}
	}
	return missingSecrets
}
