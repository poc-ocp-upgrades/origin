package admission

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/apis/authentication"
)

type AdmissionReview struct {
	metav1.TypeMeta
	Request  *AdmissionRequest
	Response *AdmissionResponse
}
type AdmissionRequest struct {
	UID         types.UID
	Kind        metav1.GroupVersionKind
	Resource    metav1.GroupVersionResource
	SubResource string
	Name        string
	Namespace   string
	Operation   Operation
	UserInfo    authentication.UserInfo
	Object      runtime.Object
	OldObject   runtime.Object
	DryRun      *bool
}
type AdmissionResponse struct {
	UID              types.UID
	Allowed          bool
	Result           *metav1.Status
	Patch            []byte
	PatchType        *PatchType
	AuditAnnotations map[string]string
}
type PatchType string

const (
	PatchTypeJSONPatch PatchType = "JSONPatch"
)

type Operation string

const (
	Create  Operation = "CREATE"
	Update  Operation = "UPDATE"
	Delete  Operation = "DELETE"
	Connect Operation = "CONNECT"
)
