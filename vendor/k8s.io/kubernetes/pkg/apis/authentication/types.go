package authentication

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ImpersonateUserHeader            = "Impersonate-User"
	ImpersonateGroupHeader           = "Impersonate-Group"
	ImpersonateUserExtraHeaderPrefix = "Impersonate-Extra-"
)

type TokenReview struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   TokenReviewSpec
	Status TokenReviewStatus
}
type TokenReviewSpec struct {
	Token     string
	Audiences []string
}
type TokenReviewStatus struct {
	Authenticated bool
	User          UserInfo
	Audiences     []string
	Error         string
}
type UserInfo struct {
	Username string
	UID      string
	Groups   []string
	Extra    map[string]ExtraValue
}
type ExtraValue []string
type TokenRequest struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   TokenRequestSpec
	Status TokenRequestStatus
}
type TokenRequestSpec struct {
	Audiences         []string
	ExpirationSeconds int64
	BoundObjectRef    *BoundObjectReference
}
type TokenRequestStatus struct {
	Token               string
	ExpirationTimestamp metav1.Time
}
type BoundObjectReference struct {
	Kind       string
	APIVersion string
	Name       string
	UID        types.UID
}
