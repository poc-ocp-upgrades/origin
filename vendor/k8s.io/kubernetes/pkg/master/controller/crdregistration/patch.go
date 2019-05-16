package crdregistration

func getGroupPriorityMin(group string) int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch group {
	case "config.openshift.io":
		return 1100
	case "operator.openshift.io":
		return 1080
	default:
		return 1000
	}
}
