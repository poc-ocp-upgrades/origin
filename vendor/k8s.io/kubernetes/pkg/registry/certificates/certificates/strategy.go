package certificates

import (
 "context"
 "fmt"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
 "k8s.io/apiserver/pkg/storage/names"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 "k8s.io/kubernetes/pkg/apis/certificates"
 "k8s.io/kubernetes/pkg/apis/certificates/validation"
)

type csrStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = csrStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (csrStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (csrStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (csrStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 csr := obj.(*certificates.CertificateSigningRequest)
 csr.Spec.Username = ""
 csr.Spec.UID = ""
 csr.Spec.Groups = nil
 csr.Spec.Extra = nil
 if user, ok := genericapirequest.UserFrom(ctx); ok {
  csr.Spec.Username = user.GetName()
  csr.Spec.UID = user.GetUID()
  csr.Spec.Groups = user.GetGroups()
  if extra := user.GetExtra(); len(extra) > 0 {
   csr.Spec.Extra = map[string]certificates.ExtraValue{}
   for k, v := range extra {
    csr.Spec.Extra[k] = certificates.ExtraValue(v)
   }
  }
 }
 csr.Status = certificates.CertificateSigningRequestStatus{}
 csr.Status.Conditions = []certificates.CertificateSigningRequestCondition{}
}
func (csrStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newCSR := obj.(*certificates.CertificateSigningRequest)
 oldCSR := old.(*certificates.CertificateSigningRequest)
 newCSR.Spec = oldCSR.Spec
 newCSR.Status = oldCSR.Status
}
func (csrStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 csr := obj.(*certificates.CertificateSigningRequest)
 return validation.ValidateCertificateSigningRequest(csr)
}
func (csrStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (csrStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldCSR := old.(*certificates.CertificateSigningRequest)
 newCSR := obj.(*certificates.CertificateSigningRequest)
 return validation.ValidateCertificateSigningRequestUpdate(newCSR, oldCSR)
}
func (csrStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (s csrStrategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 csr, ok := obj.(*certificates.CertificateSigningRequest)
 if !ok {
  return fmt.Errorf("unexpected object: %v", obj)
 }
 s.PrepareForCreate(ctx, obj)
 if exact {
  return nil
 }
 csr.Status = certificates.CertificateSigningRequestStatus{}
 return nil
}

type csrStatusStrategy struct{ csrStrategy }

var StatusStrategy = csrStatusStrategy{Strategy}

func (csrStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newCSR := obj.(*certificates.CertificateSigningRequest)
 oldCSR := old.(*certificates.CertificateSigningRequest)
 newCSR.Spec = oldCSR.Spec
 newCSR.Status.Conditions = oldCSR.Status.Conditions
 for i := range newCSR.Status.Conditions {
  if newCSR.Status.Conditions[i].LastUpdateTime.IsZero() {
   newCSR.Status.Conditions[i].LastUpdateTime = metav1.Now()
  }
 }
}
func (csrStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateCertificateSigningRequestUpdate(obj.(*certificates.CertificateSigningRequest), old.(*certificates.CertificateSigningRequest))
}
func (csrStatusStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}

type csrApprovalStrategy struct{ csrStrategy }

var ApprovalStrategy = csrApprovalStrategy{Strategy}

func (csrApprovalStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newCSR := obj.(*certificates.CertificateSigningRequest)
 oldCSR := old.(*certificates.CertificateSigningRequest)
 newCSR.Spec = oldCSR.Spec
 oldCSR.Status.Conditions = newCSR.Status.Conditions
 for i := range newCSR.Status.Conditions {
  if newCSR.Status.Conditions[i].LastUpdateTime.IsZero() {
   newCSR.Status.Conditions[i].LastUpdateTime = metav1.Now()
  }
 }
 newCSR.Status = oldCSR.Status
}
func (csrApprovalStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateCertificateSigningRequestUpdate(obj.(*certificates.CertificateSigningRequest), old.(*certificates.CertificateSigningRequest))
}
