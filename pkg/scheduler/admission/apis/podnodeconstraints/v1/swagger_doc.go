package v1

var map_PodNodeConstraintsConfig = map[string]string{"": "PodNodeConstraintsConfig is the configuration for the pod node name and node selector constraint plug-in. For accounts, serviceaccounts and groups which lack the \"pods/binding\" permission, Loading this plugin will prevent setting NodeName on pod specs and will prevent setting NodeSelectors whose labels appear in the blacklist field \"NodeSelectorLabelBlacklist\"", "nodeSelectorLabelBlacklist": "NodeSelectorLabelBlacklist specifies a list of labels which cannot be set by entities without the \"pods/binding\" permission"}

func (PodNodeConstraintsConfig) SwaggerDoc() map[string]string {
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
	return map_PodNodeConstraintsConfig
}
