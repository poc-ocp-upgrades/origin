package user

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

type User struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	FullName   string
	Identities []string
	Groups     []string
}
type UserList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []User
}
type Identity struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	ProviderName     string
	ProviderUserName string
	User             core.ObjectReference
	Extra            map[string]string
}
type IdentityList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Identity
}
type UserIdentityMapping struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Identity core.ObjectReference
	User     core.ObjectReference
}
type Group struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Users []string
}
type GroupList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Group
}
