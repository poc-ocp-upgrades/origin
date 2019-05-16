package podautoscaler

import (
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func overrideMappingsForOapiDeploymentConfig(mappings []*apimeta.RESTMapping, err error, targetGK schema.GroupKind) ([]*apimeta.RESTMapping, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if (targetGK == schema.GroupKind{Kind: "DeploymentConfig"}) {
		err = nil
		mappings = []*apimeta.RESTMapping{{Resource: schema.GroupVersionResource{Group: "apps.openshift.io", Version: "v1", Resource: "deploymentconfigs"}, GroupVersionKind: schema.GroupVersionKind{Group: "apps.openshift.io", Version: "v1", Kind: "DeploymentConfig"}}}
	}
	return mappings, err
}
