package bootstrappolicy

import (
	"reflect"
	"strings"
	"testing"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
)

const osClusterRoleAggregationPrefix = "system:openshift:"

var expectedAggregationMap = map[string]sets.String{"cluster-reader": sets.NewString("registry-viewer", "system:openshift:aggregate-to-view", "system:openshift:aggregate-to-cluster-reader")}

func TestPolicyAggregation(t *testing.T) {
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
	policyData := Policy()
	clusterRoles := policyData.ClusterRoles
	clusterRolesToAggregate := policyData.ClusterRolesToAggregate
	if len(clusterRoles) == 0 || len(clusterRolesToAggregate) == 0 {
		t.Fatalf("invalid policy data:\n%#v\n%#v", clusterRoles, clusterRolesToAggregate)
	}
	shouldHaveAggregationRuleSet := sets.NewString()
	newNameOfClusterRoleSet := sets.NewString()
	for oldName, newName := range clusterRolesToAggregate {
		if newNameOfClusterRoleSet.Has(newName) {
			t.Errorf("duplicate value %s for key %s", newName, oldName)
		}
		shouldHaveAggregationRuleSet.Insert(oldName)
		newNameOfClusterRoleSet.Insert(newName)
	}
	hasAggregationRuleSet := sets.NewString()
	aggregationMap := map[string]sets.String{}
	for i := range clusterRoles {
		cr := clusterRoles[i]
		if cr.AggregationRule == nil {
			continue
		}
		hasAggregationRuleSet.Insert(cr.Name)
		for j := range cr.AggregationRule.ClusterRoleSelectors {
			labelSelector := cr.AggregationRule.ClusterRoleSelectors[j]
			selector, err := v1.LabelSelectorAsSelector(&labelSelector)
			if err != nil {
				t.Errorf("invalid label selector %#v   at index %d for cluster role %s: %v", labelSelector, j, cr.Name, err)
				continue
			}
			for k := range clusterRoles {
				cr2 := clusterRoles[k]
				if selector.Matches(labels.Set(cr2.Labels)) {
					if cr.Name == cr2.Name {
						t.Errorf("invalid self match %s", cr.Name)
						continue
					}
					if aggregationMap[cr.Name] == nil {
						aggregationMap[cr.Name] = sets.NewString()
					}
					if aggregationMap[cr.Name].Has(cr2.Name) {
						t.Errorf("invalid duplicate entry %s for %s -> %s", cr2.Name, cr.Name, aggregationMap[cr.Name].List())
						continue
					}
					aggregationMap[cr.Name].Insert(cr2.Name)
				}
			}
		}
	}
	if !shouldHaveAggregationRuleSet.Equal(hasAggregationRuleSet) {
		missingClusterRoles := shouldHaveAggregationRuleSet.Difference(hasAggregationRuleSet).List()
		extraClusterRoles := hasAggregationRuleSet.Difference(shouldHaveAggregationRuleSet).List()
		t.Errorf("missing aggregation cluster roles = %s\nextra aggregation cluster roles = %s", missingClusterRoles, extraClusterRoles)
	}
	for parentClusterRole, childClusterRoles := range aggregationMap {
		newNameOfClusterRole, ok := clusterRolesToAggregate[parentClusterRole]
		if !ok {
			t.Errorf("cluster role %s in missing from %#v", parentClusterRole, clusterRolesToAggregate)
			continue
		}
		if !childClusterRoles.Has(newNameOfClusterRole) {
			t.Errorf("cluster role %s -> %s is missing the new name cluster role %s", parentClusterRole, childClusterRoles.List(), newNameOfClusterRole)
		}
		if !strings.HasPrefix(newNameOfClusterRole, osClusterRoleAggregationPrefix) {
			t.Errorf("invalid new name %s for old cluster role %s -> %s", newNameOfClusterRole, parentClusterRole, childClusterRoles.List())
		}
	}
	if !reflect.DeepEqual(expectedAggregationMap, aggregationMap) {
		t.Errorf("unexpected data in aggregationMap:\n%s", diff.ObjectDiff(expectedAggregationMap, aggregationMap))
	}
}
