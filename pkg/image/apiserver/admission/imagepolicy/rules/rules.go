package rules

import (
	"k8s.io/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imagepolicy "github.com/openshift/origin/pkg/image/apiserver/admission/apis/imagepolicy/v1"
)

type ImagePolicyAttributes struct {
	Resource		metav1.GroupResource
	Name			imageapi.DockerImageReference
	Image			*imageapi.Image
	ExcludedRules		sets.String
	IntegratedRegistry	bool
	LocalRewrite		bool
}
type RegistryMatcher interface{ Matches(name string) bool }
type RegistryNameMatcher func() (string, bool)

func (m RegistryNameMatcher) Matches(name string) bool {
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
	current, ok := m()
	if !ok {
		return false
	}
	return current == name
}

type nameSet []string

func (m nameSet) Matches(name string) bool {
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
	for _, s := range m {
		if s == name {
			return true
		}
	}
	return false
}
func NewRegistryMatcher(names []string) RegistryMatcher {
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
	return nameSet(names)
}

type resourceSet map[metav1.GroupResource]struct{}

func imageConditionInfo(rule *imagepolicy.ImageCondition) (covers resourceSet, selectors []labels.Selector, err error) {
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
	covers = make(resourceSet)
	for _, gr := range rule.OnResources {
		covers[gr] = struct{}{}
	}
	for i := range rule.MatchImageLabels {
		s, err := metav1.LabelSelectorAsSelector(&rule.MatchImageLabels[i])
		if err != nil {
			return nil, nil, err
		}
		selectors = append(selectors, s)
	}
	return covers, selectors, nil
}
func requiresImage(rule *imagepolicy.ImageCondition) bool {
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
	switch {
	case len(rule.MatchImageLabels) > 0, len(rule.MatchImageAnnotations) > 0, len(rule.MatchDockerImageLabels) > 0:
		return true
	}
	return false
}
func matchImageCondition(condition *imagepolicy.ImageCondition, integrated RegistryMatcher, attrs *ImagePolicyAttributes) bool {
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
	result := matchImageConditionValues(condition, integrated, attrs)
	klog.V(5).Infof("image matches conditions for %q: %t(invert=%t)", condition.Name, result, condition.InvertMatch)
	if condition.InvertMatch {
		result = !result
	}
	return result
}
func matchImageConditionValues(rule *imagepolicy.ImageCondition, integrated RegistryMatcher, attrs *ImagePolicyAttributes) bool {
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
	if rule.MatchIntegratedRegistry && !(attrs.IntegratedRegistry || integrated.Matches(attrs.Name.Registry)) {
		klog.V(5).Infof("image registry %v does not match integrated registry", attrs.Name.Registry)
		return false
	}
	if len(rule.MatchRegistries) > 0 && !hasAnyMatch(attrs.Name.Registry, rule.MatchRegistries) {
		klog.V(5).Infof("image registry %v does not match registries from rule: %#v", attrs.Name.Registry, rule.MatchRegistries)
		return false
	}
	image := attrs.Image
	if image == nil {
		if rule.SkipOnResolutionFailure {
			klog.V(5).Infof("rule does not match because image did not resolve and SkipOnResolutionFailure is true")
			return false
		}
		r := requiresImage(rule)
		klog.V(5).Infof("image did not resolve, rule requires image metadata for matching: %t", r)
		return !r
	}
	if len(rule.MatchDockerImageLabels) > 0 {
		if image.DockerImageMetadata.Config == nil {
			klog.V(5).Infof("image has no labels to match rule labels")
			return false
		}
		if !matchKeyValue(image.DockerImageMetadata.Config.Labels, rule.MatchDockerImageLabels) {
			klog.V(5).Infof("image labels %#v do not match rule labels %#v", image.DockerImageMetadata.Config.Labels, rule.MatchDockerImageLabels)
			return false
		}
	}
	if !matchKeyValue(image.Annotations, rule.MatchImageAnnotations) {
		klog.V(5).Infof("image annotations %#v do not match rule annotations %#v", image.Annotations, rule.MatchImageAnnotations)
		return false
	}
	for _, s := range rule.MatchImageLabelSelectors {
		if !s.Matches(labels.Set(image.Labels)) {
			klog.V(5).Infof("image label selectors %#v do not match rule label selectors %#v", image.Labels, s)
			return false
		}
	}
	return true
}
func matchKeyValue(all map[string]string, conditions []imagepolicy.ValueCondition) bool {
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
	for _, condition := range conditions {
		switch {
		case condition.Set:
			if _, ok := all[condition.Key]; !ok {
				return false
			}
		default:
			if all[condition.Key] != condition.Value {
				return false
			}
		}
	}
	return true
}
func hasAnyMatch(name string, all []string) bool {
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
	for _, s := range all {
		if name == s {
			return true
		}
	}
	return false
}
