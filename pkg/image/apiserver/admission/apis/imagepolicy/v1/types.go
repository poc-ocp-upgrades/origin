package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	PluginName                  = "image.openshift.io/ImagePolicy"
	IgnorePolicyRulesAnnotation = "alpha.image.policy.openshift.io/ignore-rules"
	ResolveNamesAnnotation      = "alpha.image.policy.openshift.io/resolve-names"
)

type ImagePolicyConfig struct {
	metav1.TypeMeta `json:",inline"`
	ResolveImages   ImageResolutionType         `json:"resolveImages"`
	ResolutionRules []ImageResolutionPolicyRule `json:"resolutionRules"`
	ExecutionRules  []ImageExecutionPolicyRule  `json:"executionRules"`
}
type ImageResolutionType string

var (
	RequiredRewrite ImageResolutionType = "RequiredRewrite"
	Required        ImageResolutionType = "Required"
	AttemptRewrite  ImageResolutionType = "AttemptRewrite"
	Attempt         ImageResolutionType = "Attempt"
	DoNotAttempt    ImageResolutionType = "DoNotAttempt"
)

type ImageResolutionPolicyRule struct {
	Policy         ImageResolutionType  `json:"policy"`
	TargetResource metav1.GroupResource `json:"targetResource"`
	LocalNames     bool                 `json:"localNames"`
}
type ImageExecutionPolicyRule struct {
	ImageCondition `json:",inline"`
	Reject         bool `json:"reject"`
}
type ImageCondition struct {
	Name                     string                 `json:"name"`
	IgnoreNamespaceOverride  bool                   `json:"ignoreNamespaceOverride"`
	OnResources              []metav1.GroupResource `json:"onResources"`
	InvertMatch              bool                   `json:"invertMatch"`
	MatchIntegratedRegistry  bool                   `json:"matchIntegratedRegistry"`
	MatchRegistries          []string               `json:"matchRegistries"`
	SkipOnResolutionFailure  bool                   `json:"skipOnResolutionFailure"`
	MatchDockerImageLabels   []ValueCondition       `json:"matchDockerImageLabels"`
	MatchImageLabels         []metav1.LabelSelector `json:"matchImageLabels"`
	MatchImageLabelSelectors []labels.Selector      `json:"-"`
	MatchImageAnnotations    []ValueCondition       `json:"matchImageAnnotations"`
}
type ValueCondition struct {
	Key   string `json:"key"`
	Set   bool   `json:"set"`
	Value string `json:"value"`
}
