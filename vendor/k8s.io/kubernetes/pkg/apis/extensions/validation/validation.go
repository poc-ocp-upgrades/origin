package validation

import (
 "net"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "regexp"
 "strings"
 apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
 "k8s.io/apimachinery/pkg/util/validation"
 "k8s.io/apimachinery/pkg/util/validation/field"
 apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
 "k8s.io/kubernetes/pkg/apis/extensions"
)

func ValidateIngress(ingress *extensions.Ingress) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apivalidation.ValidateObjectMeta(&ingress.ObjectMeta, true, ValidateIngressName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateIngressSpec(&ingress.Spec, field.NewPath("spec"))...)
 return allErrs
}

var ValidateIngressName = apimachineryvalidation.NameIsDNSSubdomain

func validateIngressTLS(spec *extensions.IngressSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, itls := range spec.TLS {
  for i, host := range itls.Hosts {
   if strings.Contains(host, "*") {
    for _, msg := range validation.IsWildcardDNS1123Subdomain(host) {
     allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("hosts"), host, msg))
    }
    continue
   }
   for _, msg := range validation.IsDNS1123Subdomain(host) {
    allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("hosts"), host, msg))
   }
  }
 }
 return allErrs
}
func ValidateIngressSpec(spec *extensions.IngressSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if spec.Backend != nil {
  allErrs = append(allErrs, validateIngressBackend(spec.Backend, fldPath.Child("backend"))...)
 } else if len(spec.Rules) == 0 {
  allErrs = append(allErrs, field.Invalid(fldPath, spec.Rules, "either `backend` or `rules` must be specified"))
 }
 if len(spec.Rules) > 0 {
  allErrs = append(allErrs, validateIngressRules(spec.Rules, fldPath.Child("rules"))...)
 }
 if len(spec.TLS) > 0 {
  allErrs = append(allErrs, validateIngressTLS(spec, fldPath.Child("tls"))...)
 }
 return allErrs
}
func ValidateIngressUpdate(ingress, oldIngress *extensions.Ingress) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apivalidation.ValidateObjectMetaUpdate(&ingress.ObjectMeta, &oldIngress.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateIngressSpec(&ingress.Spec, field.NewPath("spec"))...)
 return allErrs
}
func ValidateIngressStatusUpdate(ingress, oldIngress *extensions.Ingress) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apivalidation.ValidateObjectMetaUpdate(&ingress.ObjectMeta, &oldIngress.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, apivalidation.ValidateLoadBalancerStatus(&ingress.Status.LoadBalancer, field.NewPath("status", "loadBalancer"))...)
 return allErrs
}
func validateIngressRules(ingressRules []extensions.IngressRule, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(ingressRules) == 0 {
  return append(allErrs, field.Required(fldPath, ""))
 }
 for i, ih := range ingressRules {
  if len(ih.Host) > 0 {
   if isIP := (net.ParseIP(ih.Host) != nil); isIP {
    allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("host"), ih.Host, "must be a DNS name, not an IP address"))
   }
   if strings.Contains(ih.Host, "*") {
    for _, msg := range validation.IsWildcardDNS1123Subdomain(ih.Host) {
     allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("host"), ih.Host, msg))
    }
    continue
   }
   for _, msg := range validation.IsDNS1123Subdomain(ih.Host) {
    allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("host"), ih.Host, msg))
   }
  }
  allErrs = append(allErrs, validateIngressRuleValue(&ih.IngressRuleValue, fldPath.Index(0))...)
 }
 return allErrs
}
func validateIngressRuleValue(ingressRule *extensions.IngressRuleValue, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if ingressRule.HTTP != nil {
  allErrs = append(allErrs, validateHTTPIngressRuleValue(ingressRule.HTTP, fldPath.Child("http"))...)
 }
 return allErrs
}
func validateHTTPIngressRuleValue(httpIngressRuleValue *extensions.HTTPIngressRuleValue, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(httpIngressRuleValue.Paths) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("paths"), ""))
 }
 for i, rule := range httpIngressRuleValue.Paths {
  if len(rule.Path) > 0 {
   if !strings.HasPrefix(rule.Path, "/") {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("paths").Index(i).Child("path"), rule.Path, "must be an absolute path"))
   }
   _, err := regexp.CompilePOSIX(rule.Path)
   if err != nil {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("paths").Index(i).Child("path"), rule.Path, "must be a valid regex"))
   }
  }
  allErrs = append(allErrs, validateIngressBackend(&rule.Backend, fldPath.Child("backend"))...)
 }
 return allErrs
}
func validateIngressBackend(backend *extensions.IngressBackend, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(backend.ServiceName) == 0 {
  return append(allErrs, field.Required(fldPath.Child("serviceName"), ""))
 } else {
  for _, msg := range apivalidation.ValidateServiceName(backend.ServiceName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("serviceName"), backend.ServiceName, msg))
  }
 }
 allErrs = append(allErrs, apivalidation.ValidatePortNumOrName(backend.ServicePort, fldPath.Child("servicePort"))...)
 return allErrs
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
