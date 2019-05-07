package cache

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

func CreateNodeNameToInfoMap(pods []*v1.Pod, nodes []*v1.Node) map[string]*NodeInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeNameToInfo := make(map[string]*NodeInfo)
	for _, pod := range pods {
		nodeName := pod.Spec.NodeName
		if _, ok := nodeNameToInfo[nodeName]; !ok {
			nodeNameToInfo[nodeName] = NewNodeInfo()
		}
		nodeNameToInfo[nodeName].AddPod(pod)
	}
	imageExistenceMap := createImageExistenceMap(nodes)
	for _, node := range nodes {
		if _, ok := nodeNameToInfo[node.Name]; !ok {
			nodeNameToInfo[node.Name] = NewNodeInfo()
		}
		nodeInfo := nodeNameToInfo[node.Name]
		nodeInfo.SetNode(node)
		nodeInfo.imageStates = getNodeImageStates(node, imageExistenceMap)
	}
	return nodeNameToInfo
}
func getNodeImageStates(node *v1.Node, imageExistenceMap map[string]sets.String) map[string]*ImageStateSummary {
	_logClusterCodePath()
	defer _logClusterCodePath()
	imageStates := make(map[string]*ImageStateSummary)
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			imageStates[name] = &ImageStateSummary{Size: image.SizeBytes, NumNodes: len(imageExistenceMap[name])}
		}
	}
	return imageStates
}
func createImageExistenceMap(nodes []*v1.Node) map[string]sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	imageExistenceMap := make(map[string]sets.String)
	for _, node := range nodes {
		for _, image := range node.Status.Images {
			for _, name := range image.Names {
				if _, ok := imageExistenceMap[name]; !ok {
					imageExistenceMap[name] = sets.NewString(node.Name)
				} else {
					imageExistenceMap[name].Insert(node.Name)
				}
			}
		}
	}
	return imageExistenceMap
}
