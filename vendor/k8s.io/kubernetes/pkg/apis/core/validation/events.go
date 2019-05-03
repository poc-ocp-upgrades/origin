package validation

import (
 "fmt"
 "time"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/util/validation"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/kubernetes/pkg/apis/core"
)

const (
 ReportingInstanceLengthLimit = 128
 ActionLengthLimit            = 128
 ReasonLengthLimit            = 128
 NoteLengthLimit              = 1024
)

func ValidateEvent(event *core.Event) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 zeroTime := time.Time{}
 if event.EventTime.Time == zeroTime {
  if len(event.InvolvedObject.Namespace) == 0 {
   if event.Namespace != metav1.NamespaceNone && event.Namespace != metav1.NamespaceDefault {
    allErrs = append(allErrs, field.Invalid(field.NewPath("involvedObject", "namespace"), event.InvolvedObject.Namespace, "does not match event.namespace"))
   }
  } else {
   if event.Namespace != event.InvolvedObject.Namespace {
    allErrs = append(allErrs, field.Invalid(field.NewPath("involvedObject", "namespace"), event.InvolvedObject.Namespace, "does not match event.namespace"))
   }
  }
 } else {
  if len(event.InvolvedObject.Namespace) == 0 && event.Namespace != metav1.NamespaceSystem {
   allErrs = append(allErrs, field.Invalid(field.NewPath("involvedObject", "namespace"), event.InvolvedObject.Namespace, "does not match event.namespace"))
  }
  if len(event.ReportingController) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("reportingController"), ""))
  }
  for _, msg := range validation.IsQualifiedName(event.ReportingController) {
   allErrs = append(allErrs, field.Invalid(field.NewPath("reportingController"), event.ReportingController, msg))
  }
  if len(event.ReportingInstance) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("reportingInstance"), ""))
  }
  if len(event.ReportingInstance) > ReportingInstanceLengthLimit {
   allErrs = append(allErrs, field.Invalid(field.NewPath("repotingIntance"), "", fmt.Sprintf("can have at most %v characters", ReportingInstanceLengthLimit)))
  }
  if len(event.Action) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("action"), ""))
  }
  if len(event.Action) > ActionLengthLimit {
   allErrs = append(allErrs, field.Invalid(field.NewPath("action"), "", fmt.Sprintf("can have at most %v characters", ActionLengthLimit)))
  }
  if len(event.Reason) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("reason"), ""))
  }
  if len(event.Reason) > ReasonLengthLimit {
   allErrs = append(allErrs, field.Invalid(field.NewPath("reason"), "", fmt.Sprintf("can have at most %v characters", ReasonLengthLimit)))
  }
  if len(event.Message) > NoteLengthLimit {
   allErrs = append(allErrs, field.Invalid(field.NewPath("message"), "", fmt.Sprintf("can have at most %v characters", NoteLengthLimit)))
  }
 }
 for _, msg := range validation.IsDNS1123Subdomain(event.Namespace) {
  allErrs = append(allErrs, field.Invalid(field.NewPath("namespace"), event.Namespace, msg))
 }
 return allErrs
}
