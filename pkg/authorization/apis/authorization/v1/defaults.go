package v1

import (
	"github.com/openshift/api/authorization/v1"
	internal "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"k8s.io/apimachinery/pkg/api/equality"
)

var oldAllowAllPolicyRule = v1.PolicyRule{APIGroups: nil, Verbs: []string{internal.VerbAll}, Resources: []string{internal.ResourceAll}}

func SetDefaults_PolicyRule(obj *v1.PolicyRule) {
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
	if obj == nil {
		return
	}
	oldAllowAllRule := obj.APIGroups == nil && len(obj.Verbs) == 1 && obj.Verbs[0] == internal.VerbAll && len(obj.Resources) == 1 && obj.Resources[0] == internal.ResourceAll && len(obj.AttributeRestrictions.Raw) == 0 && len(obj.ResourceNames) == 0 && len(obj.NonResourceURLsSlice) == 0 && equality.Semantic.DeepEqual(oldAllowAllPolicyRule, *obj)
	if oldAllowAllRule {
		obj.APIGroups = []string{internal.APIGroupAll}
	}
	if len(obj.Resources) > 0 && len(obj.APIGroups) == 0 {
		obj.APIGroups = []string{""}
	}
}
