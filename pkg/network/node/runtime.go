package node

import (
	"fmt"
	"time"
	kwait "k8s.io/apimachinery/pkg/util/wait"
	kubeletapi "k8s.io/kubernetes/pkg/kubelet/apis/cri"
	kruntimeapi "k8s.io/kubernetes/pkg/kubelet/apis/cri/runtime/v1alpha2"
	kubeletremote "k8s.io/kubernetes/pkg/kubelet/remote"
)

func (node *OsdnNode) getRuntimeService() (kubeletapi.RuntimeService, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if node.runtimeService != nil {
		return node.runtimeService, nil
	}
	err := kwait.ExponentialBackoff(kwait.Backoff{Duration: 100 * time.Millisecond, Factor: 1.2, Steps: 24}, func() (bool, error) {
		runtimeService, err := kubeletremote.NewRemoteRuntimeService(node.runtimeEndpoint, node.runtimeRequestTimeout)
		if err != nil {
			return false, nil
		}
		if _, err := runtimeService.ListPodSandbox(&kruntimeapi.PodSandboxFilter{}); err != nil {
			return false, nil
		}
		node.runtimeService = runtimeService
		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch runtime service: %v", err)
	}
	return node.runtimeService, nil
}
func (node *OsdnNode) getPodSandboxID(filter *kruntimeapi.PodSandboxFilter) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runtimeService, err := node.getRuntimeService()
	if err != nil {
		return "", err
	}
	podSandboxList, err := runtimeService.ListPodSandbox(filter)
	if err != nil {
		return "", fmt.Errorf("failed to list pod sandboxes: %v", err)
	}
	if len(podSandboxList) == 0 {
		return "", fmt.Errorf("pod sandbox not found for filter: %v", filter)
	}
	return podSandboxList[0].Id, nil
}
func (node *OsdnNode) getPodSandboxes() (map[string]*kruntimeapi.PodSandbox, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runtimeService, err := node.getRuntimeService()
	if err != nil {
		return nil, err
	}
	podSandboxList, err := runtimeService.ListPodSandbox(&kruntimeapi.PodSandboxFilter{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pod sandboxes: %v", err)
	}
	podSandboxMap := make(map[string]*kruntimeapi.PodSandbox)
	for _, sandbox := range podSandboxList {
		podSandboxMap[sandbox.Id] = sandbox
	}
	return podSandboxMap, nil
}
