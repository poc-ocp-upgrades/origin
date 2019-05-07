package api

import (
	api "k8s.io/kubernetes/pkg/apis/core"
)

const (
	TaintNodeNotReady              = "node.kubernetes.io/not-ready"
	DeprecatedTaintNodeNotReady    = "node.alpha.kubernetes.io/notReady"
	TaintNodeUnreachable           = "node.kubernetes.io/unreachable"
	DeprecatedTaintNodeUnreachable = "node.alpha.kubernetes.io/unreachable"
	TaintNodeUnschedulable         = "node.kubernetes.io/unschedulable"
	TaintNodeOutOfDisk             = "node.kubernetes.io/out-of-disk"
	TaintNodeMemoryPressure        = "node.kubernetes.io/memory-pressure"
	TaintNodeDiskPressure          = "node.kubernetes.io/disk-pressure"
	TaintNodeNetworkUnavailable    = "node.kubernetes.io/network-unavailable"
	TaintNodePIDPressure           = "node.kubernetes.io/pid-pressure"
	TaintExternalCloudProvider     = "node.cloudprovider.kubernetes.io/uninitialized"
	TaintNodeShutdown              = "node.cloudprovider.kubernetes.io/shutdown"
	NodeFieldSelectorKeyNodeName   = api.ObjectNameField
)
