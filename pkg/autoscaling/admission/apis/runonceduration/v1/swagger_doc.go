package v1

var map_RunOnceDurationConfig = map[string]string{"": "RunOnceDurationConfig is the configuration for the RunOnceDuration plugin. It specifies a maximum value for ActiveDeadlineSeconds for a run-once pod. The project that contains the pod may specify a different setting. That setting will take precedence over the one configured for the plugin here.", "activeDeadlineSecondsOverride": "ActiveDeadlineSecondsOverride is the maximum value to set on containers of run-once pods Only a positive value is valid. Absence of a value means that the plugin won't make any changes to the pod It is kept this way for compatibility. Only change it in a new version of the API."}

func (RunOnceDurationConfig) SwaggerDoc() map[string]string {
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
	return map_RunOnceDurationConfig
}
