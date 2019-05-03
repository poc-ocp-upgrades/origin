package authorization

import (
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SubjectAccessReview struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   SubjectAccessReviewSpec
 Status SubjectAccessReviewStatus
}
type SelfSubjectAccessReview struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   SelfSubjectAccessReviewSpec
 Status SubjectAccessReviewStatus
}
type LocalSubjectAccessReview struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   SubjectAccessReviewSpec
 Status SubjectAccessReviewStatus
}
type ResourceAttributes struct {
 Namespace   string
 Verb        string
 Group       string
 Version     string
 Resource    string
 Subresource string
 Name        string
}
type NonResourceAttributes struct {
 Path string
 Verb string
}
type SubjectAccessReviewSpec struct {
 ResourceAttributes    *ResourceAttributes
 NonResourceAttributes *NonResourceAttributes
 User                  string
 Groups                []string
 Extra                 map[string]ExtraValue
 UID                   string
}
type ExtraValue []string
type SelfSubjectAccessReviewSpec struct {
 ResourceAttributes    *ResourceAttributes
 NonResourceAttributes *NonResourceAttributes
}
type SubjectAccessReviewStatus struct {
 Allowed         bool
 Denied          bool
 Reason          string
 EvaluationError string
}
type SelfSubjectRulesReview struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Spec   SelfSubjectRulesReviewSpec
 Status SubjectRulesReviewStatus
}
type SelfSubjectRulesReviewSpec struct{ Namespace string }
type SubjectRulesReviewStatus struct {
 ResourceRules    []ResourceRule
 NonResourceRules []NonResourceRule
 Incomplete       bool
 EvaluationError  string
}
type ResourceRule struct {
 Verbs         []string
 APIGroups     []string
 Resources     []string
 ResourceNames []string
}
type NonResourceRule struct {
 Verbs           []string
 NonResourceURLs []string
}
