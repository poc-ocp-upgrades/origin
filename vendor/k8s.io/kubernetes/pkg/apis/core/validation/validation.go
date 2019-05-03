package validation

import (
 "encoding/json"
 "fmt"
 "math"
 "net"
 "path"
 "path/filepath"
 "reflect"
 "regexp"
 "strings"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 apiequality "k8s.io/apimachinery/pkg/api/equality"
 "k8s.io/apimachinery/pkg/api/resource"
 apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 unversionedvalidation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime/schema"
 "k8s.io/apimachinery/pkg/util/diff"
 "k8s.io/apimachinery/pkg/util/intstr"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/validation"
 "k8s.io/apimachinery/pkg/util/validation/field"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 apiservice "k8s.io/kubernetes/pkg/api/service"
 "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/helper"
 podshelper "k8s.io/kubernetes/pkg/apis/core/pods"
 corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
 v1helper "k8s.io/kubernetes/pkg/apis/core/v1/helper"
 "k8s.io/kubernetes/pkg/capabilities"
 "k8s.io/kubernetes/pkg/features"
 "k8s.io/kubernetes/pkg/fieldpath"
 "k8s.io/kubernetes/pkg/master/ports"
 "k8s.io/kubernetes/pkg/security/apparmor"
)

const isNegativeErrorMsg string = apimachineryvalidation.IsNegativeErrorMsg
const isInvalidQuotaResource string = `must be a standard resource for quota`
const fieldImmutableErrorMsg string = apimachineryvalidation.FieldImmutableErrorMsg
const isNotIntegerErrorMsg string = `must be an integer`
const isNotPositiveErrorMsg string = `must be greater than zero`
const csiDriverNameRexpErrMsg string = "must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character"
const csiDriverNameRexpFmt string = `^[a-zA-Z0-9][-a-zA-Z0-9_.]{0,61}[a-zA-Z-0-9]$`

var pdPartitionErrorMsg string = validation.InclusiveRangeError(1, 255)
var fileModeErrorMsg string = "must be a number between 0 and 0777 (octal), both inclusive"
var BannedOwners = apimachineryvalidation.BannedOwners
var iscsiInitiatorIqnRegex = regexp.MustCompile(`iqn\.\d{4}-\d{2}\.([[:alnum:]-.]+)(:[^,;*&$|\s]+)$`)
var iscsiInitiatorEuiRegex = regexp.MustCompile(`^eui.[[:alnum:]]{16}$`)
var iscsiInitiatorNaaRegex = regexp.MustCompile(`^naa.[[:alnum:]]{32}$`)
var csiDriverNameRexp = regexp.MustCompile(csiDriverNameRexpFmt)

func ValidateHasLabel(meta metav1.ObjectMeta, fldPath *field.Path, key, expectedValue string) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 actualValue, found := meta.Labels[key]
 if !found {
  allErrs = append(allErrs, field.Required(fldPath.Child("labels").Key(key), fmt.Sprintf("must be '%s'", expectedValue)))
  return allErrs
 }
 if actualValue != expectedValue {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("labels").Key(key), meta.Labels, fmt.Sprintf("must be '%s'", expectedValue)))
 }
 return allErrs
}
func ValidateAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apimachineryvalidation.ValidateAnnotations(annotations, fldPath)
}
func ValidateDNS1123Label(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, msg := range validation.IsDNS1123Label(value) {
  allErrs = append(allErrs, field.Invalid(fldPath, value, msg))
 }
 return allErrs
}
func ValidateDNS1123Subdomain(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, msg := range validation.IsDNS1123Subdomain(value) {
  allErrs = append(allErrs, field.Invalid(fldPath, value, msg))
 }
 return allErrs
}
func ValidatePodSpecificAnnotations(annotations map[string]string, spec *core.PodSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if value, isMirror := annotations[core.MirrorPodAnnotationKey]; isMirror {
  if len(spec.NodeName) == 0 {
   allErrs = append(allErrs, field.Invalid(fldPath.Key(core.MirrorPodAnnotationKey), value, "must set spec.nodeName if mirror pod annotation is set"))
  }
 }
 if annotations[core.TolerationsAnnotationKey] != "" {
  allErrs = append(allErrs, ValidateTolerationsInPodAnnotations(annotations, fldPath)...)
 }
 allErrs = append(allErrs, ValidateSeccompPodAnnotations(annotations, fldPath)...)
 allErrs = append(allErrs, ValidateAppArmorPodAnnotations(annotations, spec, fldPath)...)
 return allErrs
}
func ValidateTolerationsInPodAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 tolerations, err := helper.GetTolerationsFromPodAnnotations(annotations)
 if err != nil {
  allErrs = append(allErrs, field.Invalid(fldPath, core.TolerationsAnnotationKey, err.Error()))
  return allErrs
 }
 if len(tolerations) > 0 {
  allErrs = append(allErrs, ValidateTolerations(tolerations, fldPath.Child(core.TolerationsAnnotationKey))...)
 }
 return allErrs
}
func ValidatePodSpecificAnnotationUpdates(newPod, oldPod *core.Pod, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 newAnnotations := newPod.Annotations
 oldAnnotations := oldPod.Annotations
 for k, oldVal := range oldAnnotations {
  if newVal, exists := newAnnotations[k]; exists && newVal == oldVal {
   continue
  }
  if strings.HasPrefix(k, apparmor.ContainerAnnotationKeyPrefix) {
   allErrs = append(allErrs, field.Forbidden(fldPath.Key(k), "may not remove or update AppArmor annotations"))
  }
  if k == core.MirrorPodAnnotationKey {
   allErrs = append(allErrs, field.Forbidden(fldPath.Key(k), "may not remove or update mirror pod annotation"))
  }
 }
 for k := range newAnnotations {
  if _, ok := oldAnnotations[k]; ok {
   continue
  }
  if strings.HasPrefix(k, apparmor.ContainerAnnotationKeyPrefix) {
   allErrs = append(allErrs, field.Forbidden(fldPath.Key(k), "may not add AppArmor annotations"))
  }
  if k == core.MirrorPodAnnotationKey {
   allErrs = append(allErrs, field.Forbidden(fldPath.Key(k), "may not add mirror pod annotation"))
  }
 }
 allErrs = append(allErrs, ValidatePodSpecificAnnotations(newAnnotations, &newPod.Spec, fldPath)...)
 return allErrs
}
func ValidateEndpointsSpecificAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 return allErrs
}

type ValidateNameFunc apimachineryvalidation.ValidateNameFunc

var ValidatePodName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateReplicationControllerName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateServiceName = apimachineryvalidation.NameIsDNS1035Label
var ValidateNodeName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateNamespaceName = apimachineryvalidation.ValidateNamespaceName
var ValidateLimitRangeName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateResourceQuotaName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateSecretName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateServiceAccountName = apimachineryvalidation.ValidateServiceAccountName
var ValidateEndpointsName = apimachineryvalidation.NameIsDNSSubdomain
var ValidateClusterName = apimachineryvalidation.ValidateClusterName
var ValidateClassName = apimachineryvalidation.NameIsDNSSubdomain
var ValidatePriorityClassName = apimachineryvalidation.NameIsDNSSubdomain

func ValidateRuntimeClassName(name string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var allErrs field.ErrorList
 for _, msg := range apimachineryvalidation.NameIsDNSSubdomain(name, false) {
  allErrs = append(allErrs, field.Invalid(fldPath, name, msg))
 }
 return allErrs
}
func ValidateNonnegativeField(value int64, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apimachineryvalidation.ValidateNonnegativeField(value, fldPath)
}
func ValidateNonnegativeQuantity(value resource.Quantity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if value.Cmp(resource.Quantity{}) < 0 {
  allErrs = append(allErrs, field.Invalid(fldPath, value.String(), isNegativeErrorMsg))
 }
 return allErrs
}
func ValidatePositiveQuantityValue(value resource.Quantity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if value.Cmp(resource.Quantity{}) <= 0 {
  allErrs = append(allErrs, field.Invalid(fldPath, value.String(), isNotPositiveErrorMsg))
 }
 return allErrs
}
func ValidateImmutableField(newVal, oldVal interface{}, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return apimachineryvalidation.ValidateImmutableField(newVal, oldVal, fldPath)
}
func ValidateImmutableAnnotation(newVal string, oldVal string, annotation string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if oldVal != newVal {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("annotations", annotation), newVal, fieldImmutableErrorMsg))
 }
 return allErrs
}
func ValidateObjectMeta(meta *metav1.ObjectMeta, requiresNamespace bool, nameFn ValidateNameFunc, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apimachineryvalidation.ValidateObjectMeta(meta, requiresNamespace, apimachineryvalidation.ValidateNameFunc(nameFn), fldPath)
 for i := range meta.Finalizers {
  allErrs = append(allErrs, validateKubeFinalizerName(string(meta.Finalizers[i]), fldPath.Child("finalizers").Index(i))...)
 }
 return allErrs
}
func ValidateObjectMetaUpdate(newMeta, oldMeta *metav1.ObjectMeta, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apimachineryvalidation.ValidateObjectMetaUpdate(newMeta, oldMeta, fldPath)
 for i := range newMeta.Finalizers {
  allErrs = append(allErrs, validateKubeFinalizerName(string(newMeta.Finalizers[i]), fldPath.Child("finalizers").Index(i))...)
 }
 return allErrs
}
func ValidateVolumes(volumes []core.Volume, fldPath *field.Path) (map[string]core.VolumeSource, field.ErrorList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allNames := sets.String{}
 vols := make(map[string]core.VolumeSource)
 for i, vol := range volumes {
  idxPath := fldPath.Index(i)
  namePath := idxPath.Child("name")
  el := validateVolumeSource(&vol.VolumeSource, idxPath, vol.Name)
  if len(vol.Name) == 0 {
   el = append(el, field.Required(namePath, ""))
  } else {
   el = append(el, ValidateDNS1123Label(vol.Name, namePath)...)
  }
  if allNames.Has(vol.Name) {
   el = append(el, field.Duplicate(namePath, vol.Name))
  }
  if len(el) == 0 {
   allNames.Insert(vol.Name)
   vols[vol.Name] = vol.VolumeSource
  } else {
   allErrs = append(allErrs, el...)
  }
 }
 return vols, allErrs
}
func IsMatchedVolume(name string, volumes map[string]core.VolumeSource) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, ok := volumes[name]; ok {
  return true
 }
 return false
}
func isMatchedDevice(name string, volumes map[string]core.VolumeSource) (bool, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if source, ok := volumes[name]; ok {
  if source.PersistentVolumeClaim != nil {
   return true, true
  }
  return true, false
 }
 return false, false
}
func mountNameAlreadyExists(name string, devices map[string]string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, ok := devices[name]; ok {
  return true
 }
 return false
}
func mountPathAlreadyExists(mountPath string, devices map[string]string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, devPath := range devices {
  if mountPath == devPath {
   return true
  }
 }
 return false
}
func deviceNameAlreadyExists(name string, mounts map[string]string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if _, ok := mounts[name]; ok {
  return true
 }
 return false
}
func devicePathAlreadyExists(devicePath string, mounts map[string]string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, mountPath := range mounts {
  if mountPath == devicePath {
   return true
  }
 }
 return false
}
func validateVolumeSource(source *core.VolumeSource, fldPath *field.Path, volName string) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 numVolumes := 0
 allErrs := field.ErrorList{}
 if source.EmptyDir != nil {
  numVolumes++
  if !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
   if source.EmptyDir.SizeLimit != nil && source.EmptyDir.SizeLimit.Cmp(resource.Quantity{}) != 0 {
    allErrs = append(allErrs, field.Forbidden(fldPath.Child("emptyDir").Child("sizeLimit"), "SizeLimit field disabled by feature-gate for EmptyDir volumes"))
   }
  } else {
   if source.EmptyDir.SizeLimit != nil && source.EmptyDir.SizeLimit.Cmp(resource.Quantity{}) < 0 {
    allErrs = append(allErrs, field.Forbidden(fldPath.Child("emptyDir").Child("sizeLimit"), "SizeLimit field must be a valid resource quantity"))
   }
  }
  if !utilfeature.DefaultFeatureGate.Enabled(features.HugePages) && source.EmptyDir.Medium == core.StorageMediumHugePages {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("emptyDir").Child("medium"), "HugePages medium is disabled by feature-gate for EmptyDir volumes"))
  }
 }
 if source.HostPath != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("hostPath"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateHostPathVolumeSource(source.HostPath, fldPath.Child("hostPath"))...)
  }
 }
 if source.GitRepo != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("gitRepo"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateGitRepoVolumeSource(source.GitRepo, fldPath.Child("gitRepo"))...)
  }
 }
 if source.GCEPersistentDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("gcePersistentDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateGCEPersistentDiskVolumeSource(source.GCEPersistentDisk, fldPath.Child("persistentDisk"))...)
  }
 }
 if source.AWSElasticBlockStore != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("awsElasticBlockStore"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAWSElasticBlockStoreVolumeSource(source.AWSElasticBlockStore, fldPath.Child("awsElasticBlockStore"))...)
  }
 }
 if source.Secret != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("secret"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateSecretVolumeSource(source.Secret, fldPath.Child("secret"))...)
  }
 }
 if source.NFS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("nfs"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateNFSVolumeSource(source.NFS, fldPath.Child("nfs"))...)
  }
 }
 if source.ISCSI != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("iscsi"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateISCSIVolumeSource(source.ISCSI, fldPath.Child("iscsi"))...)
  }
  if source.ISCSI.InitiatorName != nil && len(volName+":"+source.ISCSI.TargetPortal) > 64 {
   tooLongErr := "Total length of <volume name>:<iscsi.targetPortal> must be under 64 characters if iscsi.initiatorName is specified."
   allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), volName, tooLongErr))
  }
 }
 if source.Glusterfs != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("glusterfs"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateGlusterfsVolumeSource(source.Glusterfs, fldPath.Child("glusterfs"))...)
  }
 }
 if source.Flocker != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("flocker"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateFlockerVolumeSource(source.Flocker, fldPath.Child("flocker"))...)
  }
 }
 if source.PersistentVolumeClaim != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("persistentVolumeClaim"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validatePersistentClaimVolumeSource(source.PersistentVolumeClaim, fldPath.Child("persistentVolumeClaim"))...)
  }
 }
 if source.RBD != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("rbd"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateRBDVolumeSource(source.RBD, fldPath.Child("rbd"))...)
  }
 }
 if source.Cinder != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("cinder"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateCinderVolumeSource(source.Cinder, fldPath.Child("cinder"))...)
  }
 }
 if source.CephFS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("cephFS"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateCephFSVolumeSource(source.CephFS, fldPath.Child("cephfs"))...)
  }
 }
 if source.Quobyte != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("quobyte"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateQuobyteVolumeSource(source.Quobyte, fldPath.Child("quobyte"))...)
  }
 }
 if source.DownwardAPI != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("downwarAPI"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateDownwardAPIVolumeSource(source.DownwardAPI, fldPath.Child("downwardAPI"))...)
  }
 }
 if source.FC != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("fc"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateFCVolumeSource(source.FC, fldPath.Child("fc"))...)
  }
 }
 if source.FlexVolume != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("flexVolume"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateFlexVolumeSource(source.FlexVolume, fldPath.Child("flexVolume"))...)
  }
 }
 if source.ConfigMap != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("configMap"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateConfigMapVolumeSource(source.ConfigMap, fldPath.Child("configMap"))...)
  }
 }
 if source.AzureFile != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("azureFile"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAzureFile(source.AzureFile, fldPath.Child("azureFile"))...)
  }
 }
 if source.VsphereVolume != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("vsphereVolume"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateVsphereVolumeSource(source.VsphereVolume, fldPath.Child("vsphereVolume"))...)
  }
 }
 if source.PhotonPersistentDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("photonPersistentDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validatePhotonPersistentDiskVolumeSource(source.PhotonPersistentDisk, fldPath.Child("photonPersistentDisk"))...)
  }
 }
 if source.PortworxVolume != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("portworxVolume"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validatePortworxVolumeSource(source.PortworxVolume, fldPath.Child("portworxVolume"))...)
  }
 }
 if source.AzureDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("azureDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAzureDisk(source.AzureDisk, fldPath.Child("azureDisk"))...)
  }
 }
 if source.StorageOS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("storageos"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateStorageOSVolumeSource(source.StorageOS, fldPath.Child("storageos"))...)
  }
 }
 if source.Projected != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("projected"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateProjectedVolumeSource(source.Projected, fldPath.Child("projected"))...)
  }
 }
 if source.ScaleIO != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("scaleIO"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateScaleIOVolumeSource(source.ScaleIO, fldPath.Child("scaleIO"))...)
  }
 }
 if numVolumes == 0 {
  allErrs = append(allErrs, field.Required(fldPath, "must specify a volume type"))
 }
 return allErrs
}
func validateHostPathVolumeSource(hostPath *core.HostPathVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(hostPath.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
  return allErrs
 }
 allErrs = append(allErrs, validatePathNoBacksteps(hostPath.Path, fldPath.Child("path"))...)
 allErrs = append(allErrs, validateHostPathType(hostPath.Type, fldPath.Child("type"))...)
 return allErrs
}
func validateGitRepoVolumeSource(gitRepo *core.GitRepoVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(gitRepo.Repository) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("repository"), ""))
 }
 pathErrs := validateLocalDescendingPath(gitRepo.Directory, fldPath.Child("directory"))
 allErrs = append(allErrs, pathErrs...)
 return allErrs
}
func validateISCSIVolumeSource(iscsi *core.ISCSIVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(iscsi.TargetPortal) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("targetPortal"), ""))
 }
 if len(iscsi.IQN) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("iqn"), ""))
 } else {
  if !strings.HasPrefix(iscsi.IQN, "iqn") && !strings.HasPrefix(iscsi.IQN, "eui") && !strings.HasPrefix(iscsi.IQN, "naa") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format starting with iqn, eui, or naa"))
  } else if strings.HasPrefix(iscsi.IQN, "iqn") && !iscsiInitiatorIqnRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  } else if strings.HasPrefix(iscsi.IQN, "eui") && !iscsiInitiatorEuiRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  } else if strings.HasPrefix(iscsi.IQN, "naa") && !iscsiInitiatorNaaRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  }
 }
 if iscsi.Lun < 0 || iscsi.Lun > 255 {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("lun"), iscsi.Lun, validation.InclusiveRangeError(0, 255)))
 }
 if (iscsi.DiscoveryCHAPAuth || iscsi.SessionCHAPAuth) && iscsi.SecretRef == nil {
  allErrs = append(allErrs, field.Required(fldPath.Child("secretRef"), ""))
 }
 if iscsi.InitiatorName != nil {
  initiator := *iscsi.InitiatorName
  if !strings.HasPrefix(initiator, "iqn") && !strings.HasPrefix(initiator, "eui") && !strings.HasPrefix(initiator, "naa") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format starting with iqn, eui, or naa"))
  }
  if strings.HasPrefix(initiator, "iqn") && !iscsiInitiatorIqnRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  } else if strings.HasPrefix(initiator, "eui") && !iscsiInitiatorEuiRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  } else if strings.HasPrefix(initiator, "naa") && !iscsiInitiatorNaaRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  }
 }
 return allErrs
}
func validateISCSIPersistentVolumeSource(iscsi *core.ISCSIPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(iscsi.TargetPortal) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("targetPortal"), ""))
 }
 if len(iscsi.IQN) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("iqn"), ""))
 } else {
  if !strings.HasPrefix(iscsi.IQN, "iqn") && !strings.HasPrefix(iscsi.IQN, "eui") && !strings.HasPrefix(iscsi.IQN, "naa") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  } else if strings.HasPrefix(iscsi.IQN, "iqn") && !iscsiInitiatorIqnRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  } else if strings.HasPrefix(iscsi.IQN, "eui") && !iscsiInitiatorEuiRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  } else if strings.HasPrefix(iscsi.IQN, "naa") && !iscsiInitiatorNaaRegex.MatchString(iscsi.IQN) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("iqn"), iscsi.IQN, "must be valid format"))
  }
 }
 if iscsi.Lun < 0 || iscsi.Lun > 255 {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("lun"), iscsi.Lun, validation.InclusiveRangeError(0, 255)))
 }
 if (iscsi.DiscoveryCHAPAuth || iscsi.SessionCHAPAuth) && iscsi.SecretRef == nil {
  allErrs = append(allErrs, field.Required(fldPath.Child("secretRef"), ""))
 }
 if iscsi.SecretRef != nil {
  if len(iscsi.SecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "name"), ""))
  }
 }
 if iscsi.InitiatorName != nil {
  initiator := *iscsi.InitiatorName
  if !strings.HasPrefix(initiator, "iqn") && !strings.HasPrefix(initiator, "eui") && !strings.HasPrefix(initiator, "naa") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  }
  if strings.HasPrefix(initiator, "iqn") && !iscsiInitiatorIqnRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  } else if strings.HasPrefix(initiator, "eui") && !iscsiInitiatorEuiRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  } else if strings.HasPrefix(initiator, "naa") && !iscsiInitiatorNaaRegex.MatchString(initiator) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("initiatorname"), initiator, "must be valid format"))
  }
 }
 return allErrs
}
func validateFCVolumeSource(fc *core.FCVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(fc.TargetWWNs) < 1 && len(fc.WWIDs) < 1 {
  allErrs = append(allErrs, field.Required(fldPath.Child("targetWWNs"), "must specify either targetWWNs or wwids, but not both"))
 }
 if len(fc.TargetWWNs) != 0 && len(fc.WWIDs) != 0 {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("targetWWNs"), fc.TargetWWNs, "targetWWNs and wwids can not be specified simultaneously"))
 }
 if len(fc.TargetWWNs) != 0 {
  if fc.Lun == nil {
   allErrs = append(allErrs, field.Required(fldPath.Child("lun"), "lun is required if targetWWNs is specified"))
  } else {
   if *fc.Lun < 0 || *fc.Lun > 255 {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("lun"), fc.Lun, validation.InclusiveRangeError(0, 255)))
   }
  }
 }
 return allErrs
}
func validateGCEPersistentDiskVolumeSource(pd *core.GCEPersistentDiskVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(pd.PDName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("pdName"), ""))
 }
 if pd.Partition < 0 || pd.Partition > 255 {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("partition"), pd.Partition, pdPartitionErrorMsg))
 }
 return allErrs
}
func validateAWSElasticBlockStoreVolumeSource(PD *core.AWSElasticBlockStoreVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(PD.VolumeID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeID"), ""))
 }
 if PD.Partition < 0 || PD.Partition > 255 {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("partition"), PD.Partition, pdPartitionErrorMsg))
 }
 return allErrs
}
func validateSecretVolumeSource(secretSource *core.SecretVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(secretSource.SecretName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("secretName"), ""))
 }
 secretMode := secretSource.DefaultMode
 if secretMode != nil && (*secretMode > 0777 || *secretMode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("defaultMode"), *secretMode, fileModeErrorMsg))
 }
 itemsPath := fldPath.Child("items")
 for i, kp := range secretSource.Items {
  itemPath := itemsPath.Index(i)
  allErrs = append(allErrs, validateKeyToPath(&kp, itemPath)...)
 }
 return allErrs
}
func validateConfigMapVolumeSource(configMapSource *core.ConfigMapVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(configMapSource.Name) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
 }
 configMapMode := configMapSource.DefaultMode
 if configMapMode != nil && (*configMapMode > 0777 || *configMapMode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("defaultMode"), *configMapMode, fileModeErrorMsg))
 }
 itemsPath := fldPath.Child("items")
 for i, kp := range configMapSource.Items {
  itemPath := itemsPath.Index(i)
  allErrs = append(allErrs, validateKeyToPath(&kp, itemPath)...)
 }
 return allErrs
}
func validateKeyToPath(kp *core.KeyToPath, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(kp.Key) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("key"), ""))
 }
 if len(kp.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
 }
 allErrs = append(allErrs, validateLocalNonReservedPath(kp.Path, fldPath.Child("path"))...)
 if kp.Mode != nil && (*kp.Mode > 0777 || *kp.Mode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("mode"), *kp.Mode, fileModeErrorMsg))
 }
 return allErrs
}
func validatePersistentClaimVolumeSource(claim *core.PersistentVolumeClaimVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(claim.ClaimName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("claimName"), ""))
 }
 return allErrs
}
func validateNFSVolumeSource(nfs *core.NFSVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(nfs.Server) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("server"), ""))
 }
 if len(nfs.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
 }
 if !path.IsAbs(nfs.Path) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("path"), nfs.Path, "must be an absolute path"))
 }
 return allErrs
}
func validateQuobyteVolumeSource(quobyte *core.QuobyteVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(quobyte.Registry) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("registry"), "must be a host:port pair or multiple pairs separated by commas"))
 } else {
  for _, hostPortPair := range strings.Split(quobyte.Registry, ",") {
   if _, _, err := net.SplitHostPort(hostPortPair); err != nil {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("registry"), quobyte.Registry, "must be a host:port pair or multiple pairs separated by commas"))
   }
  }
 }
 if len(quobyte.Volume) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volume"), ""))
 }
 return allErrs
}
func validateGlusterfsVolumeSource(glusterfs *core.GlusterfsVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(glusterfs.EndpointsName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("endpoints"), ""))
 }
 if len(glusterfs.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
 }
 return allErrs
}
func validateGlusterfsPersistentVolumeSource(glusterfs *core.GlusterfsPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(glusterfs.EndpointsName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("endpoints"), ""))
 }
 if len(glusterfs.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
 }
 if glusterfs.EndpointsNamespace != nil {
  endpointNs := glusterfs.EndpointsNamespace
  if *endpointNs == "" {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("endpointsNamespace"), *endpointNs, "if the endpointnamespace is set, it must be a valid namespace name"))
  } else {
   for _, msg := range ValidateNamespaceName(*endpointNs, false) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("endpointsNamespace"), *endpointNs, msg))
   }
  }
 }
 return allErrs
}
func validateFlockerVolumeSource(flocker *core.FlockerVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(flocker.DatasetName) == 0 && len(flocker.DatasetUUID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath, "one of datasetName and datasetUUID is required"))
 }
 if len(flocker.DatasetName) != 0 && len(flocker.DatasetUUID) != 0 {
  allErrs = append(allErrs, field.Invalid(fldPath, "resource", "datasetName and datasetUUID can not be specified simultaneously"))
 }
 if strings.Contains(flocker.DatasetName, "/") {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("datasetName"), flocker.DatasetName, "must not contain '/'"))
 }
 return allErrs
}

var validVolumeDownwardAPIFieldPathExpressions = sets.NewString("metadata.name", "metadata.namespace", "metadata.labels", "metadata.annotations", "metadata.uid")

func validateDownwardAPIVolumeFile(file *core.DownwardAPIVolumeFile, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(file.Path) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
 }
 allErrs = append(allErrs, validateLocalNonReservedPath(file.Path, fldPath.Child("path"))...)
 if file.FieldRef != nil {
  allErrs = append(allErrs, validateObjectFieldSelector(file.FieldRef, &validVolumeDownwardAPIFieldPathExpressions, fldPath.Child("fieldRef"))...)
  if file.ResourceFieldRef != nil {
   allErrs = append(allErrs, field.Invalid(fldPath, "resource", "fieldRef and resourceFieldRef can not be specified simultaneously"))
  }
 } else if file.ResourceFieldRef != nil {
  allErrs = append(allErrs, validateContainerResourceFieldSelector(file.ResourceFieldRef, &validContainerResourceFieldPathExpressions, fldPath.Child("resourceFieldRef"), true)...)
 } else {
  allErrs = append(allErrs, field.Required(fldPath, "one of fieldRef and resourceFieldRef is required"))
 }
 if file.Mode != nil && (*file.Mode > 0777 || *file.Mode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("mode"), *file.Mode, fileModeErrorMsg))
 }
 return allErrs
}
func validateDownwardAPIVolumeSource(downwardAPIVolume *core.DownwardAPIVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 downwardAPIMode := downwardAPIVolume.DefaultMode
 if downwardAPIMode != nil && (*downwardAPIMode > 0777 || *downwardAPIMode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("defaultMode"), *downwardAPIMode, fileModeErrorMsg))
 }
 for _, file := range downwardAPIVolume.Items {
  allErrs = append(allErrs, validateDownwardAPIVolumeFile(&file, fldPath)...)
 }
 return allErrs
}
func validateProjectionSources(projection *core.ProjectedVolumeSource, projectionMode *int32, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allPaths := sets.String{}
 for i, source := range projection.Sources {
  numSources := 0
  srcPath := fldPath.Child("sources").Index(i)
  if projPath := srcPath.Child("secret"); source.Secret != nil {
   numSources++
   if len(source.Secret.Name) == 0 {
    allErrs = append(allErrs, field.Required(projPath.Child("name"), ""))
   }
   itemsPath := projPath.Child("items")
   for i, kp := range source.Secret.Items {
    itemPath := itemsPath.Index(i)
    allErrs = append(allErrs, validateKeyToPath(&kp, itemPath)...)
    if len(kp.Path) > 0 {
     curPath := kp.Path
     if !allPaths.Has(curPath) {
      allPaths.Insert(curPath)
     } else {
      allErrs = append(allErrs, field.Invalid(fldPath, source.Secret.Name, "conflicting duplicate paths"))
     }
    }
   }
  }
  if projPath := srcPath.Child("configMap"); source.ConfigMap != nil {
   numSources++
   if len(source.ConfigMap.Name) == 0 {
    allErrs = append(allErrs, field.Required(projPath.Child("name"), ""))
   }
   itemsPath := projPath.Child("items")
   for i, kp := range source.ConfigMap.Items {
    itemPath := itemsPath.Index(i)
    allErrs = append(allErrs, validateKeyToPath(&kp, itemPath)...)
    if len(kp.Path) > 0 {
     curPath := kp.Path
     if !allPaths.Has(curPath) {
      allPaths.Insert(curPath)
     } else {
      allErrs = append(allErrs, field.Invalid(fldPath, source.ConfigMap.Name, "conflicting duplicate paths"))
     }
    }
   }
  }
  if projPath := srcPath.Child("downwardAPI"); source.DownwardAPI != nil {
   numSources++
   for _, file := range source.DownwardAPI.Items {
    allErrs = append(allErrs, validateDownwardAPIVolumeFile(&file, projPath)...)
    if len(file.Path) > 0 {
     curPath := file.Path
     if !allPaths.Has(curPath) {
      allPaths.Insert(curPath)
     } else {
      allErrs = append(allErrs, field.Invalid(fldPath, curPath, "conflicting duplicate paths"))
     }
    }
   }
  }
  if projPath := srcPath.Child("serviceAccountToken"); source.ServiceAccountToken != nil {
   numSources++
   if !utilfeature.DefaultFeatureGate.Enabled(features.TokenRequestProjection) {
    allErrs = append(allErrs, field.Forbidden(projPath, "TokenRequestProjection feature is not enabled"))
   }
   if source.ServiceAccountToken.ExpirationSeconds < 10*60 {
    allErrs = append(allErrs, field.Invalid(projPath.Child("expirationSeconds"), source.ServiceAccountToken.ExpirationSeconds, "may not specify a duration less than 10 minutes"))
   }
   if source.ServiceAccountToken.ExpirationSeconds > 1<<32 {
    allErrs = append(allErrs, field.Invalid(projPath.Child("expirationSeconds"), source.ServiceAccountToken.ExpirationSeconds, "may not specify a duration larger than 2^32 seconds"))
   }
   if source.ServiceAccountToken.Path == "" {
    allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
   }
  }
  if numSources > 1 {
   allErrs = append(allErrs, field.Forbidden(srcPath, "may not specify more than 1 volume type"))
  }
 }
 return allErrs
}
func validateProjectedVolumeSource(projection *core.ProjectedVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 projectionMode := projection.DefaultMode
 if projectionMode != nil && (*projectionMode > 0777 || *projectionMode < 0) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("defaultMode"), *projectionMode, fileModeErrorMsg))
 }
 allErrs = append(allErrs, validateProjectionSources(projection, projectionMode, fldPath)...)
 return allErrs
}

var supportedHostPathTypes = sets.NewString(string(core.HostPathUnset), string(core.HostPathDirectoryOrCreate), string(core.HostPathDirectory), string(core.HostPathFileOrCreate), string(core.HostPathFile), string(core.HostPathSocket), string(core.HostPathCharDev), string(core.HostPathBlockDev))

func validateHostPathType(hostPathType *core.HostPathType, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if hostPathType != nil && !supportedHostPathTypes.Has(string(*hostPathType)) {
  allErrs = append(allErrs, field.NotSupported(fldPath, hostPathType, supportedHostPathTypes.List()))
 }
 return allErrs
}
func validateLocalDescendingPath(targetPath string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if path.IsAbs(targetPath) {
  allErrs = append(allErrs, field.Invalid(fldPath, targetPath, "must be a relative path"))
 }
 allErrs = append(allErrs, validatePathNoBacksteps(targetPath, fldPath)...)
 return allErrs
}
func validatePathNoBacksteps(targetPath string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 parts := strings.Split(filepath.ToSlash(targetPath), "/")
 for _, item := range parts {
  if item == ".." {
   allErrs = append(allErrs, field.Invalid(fldPath, targetPath, "must not contain '..'"))
   break
  }
 }
 return allErrs
}
func validateMountPropagation(mountPropagation *core.MountPropagationMode, container *core.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if mountPropagation == nil {
  return allErrs
 }
 supportedMountPropagations := sets.NewString(string(core.MountPropagationBidirectional), string(core.MountPropagationHostToContainer), string(core.MountPropagationNone))
 if !supportedMountPropagations.Has(string(*mountPropagation)) {
  allErrs = append(allErrs, field.NotSupported(fldPath, *mountPropagation, supportedMountPropagations.List()))
 }
 if container == nil {
  return allErrs
 }
 privileged := container.SecurityContext != nil && container.SecurityContext.Privileged != nil && *container.SecurityContext.Privileged
 if *mountPropagation == core.MountPropagationBidirectional && !privileged {
  allErrs = append(allErrs, field.Forbidden(fldPath, "Bidirectional mount propagation is available only to privileged containers"))
 }
 return allErrs
}
func validateLocalNonReservedPath(targetPath string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, validateLocalDescendingPath(targetPath, fldPath)...)
 if strings.HasPrefix(targetPath, "..") && !strings.HasPrefix(targetPath, "../") {
  allErrs = append(allErrs, field.Invalid(fldPath, targetPath, "must not start with '..'"))
 }
 return allErrs
}
func validateRBDVolumeSource(rbd *core.RBDVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(rbd.CephMonitors) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("monitors"), ""))
 }
 if len(rbd.RBDImage) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("image"), ""))
 }
 return allErrs
}
func validateRBDPersistentVolumeSource(rbd *core.RBDPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(rbd.CephMonitors) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("monitors"), ""))
 }
 if len(rbd.RBDImage) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("image"), ""))
 }
 return allErrs
}
func validateCinderVolumeSource(cd *core.CinderVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cd.VolumeID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeID"), ""))
 }
 if cd.SecretRef != nil {
  if len(cd.SecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "name"), ""))
  }
 }
 return allErrs
}
func validateCinderPersistentVolumeSource(cd *core.CinderPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cd.VolumeID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeID"), ""))
 }
 if cd.SecretRef != nil {
  if len(cd.SecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "name"), ""))
  }
  if len(cd.SecretRef.Namespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "namespace"), ""))
  }
 }
 return allErrs
}
func validateCephFSVolumeSource(cephfs *core.CephFSVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cephfs.Monitors) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("monitors"), ""))
 }
 return allErrs
}
func validateCephFSPersistentVolumeSource(cephfs *core.CephFSPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cephfs.Monitors) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("monitors"), ""))
 }
 return allErrs
}
func validateFlexVolumeSource(fv *core.FlexVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(fv.Driver) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("driver"), ""))
 }
 for k := range fv.Options {
  namespace := k
  if parts := strings.SplitN(k, "/", 2); len(parts) == 2 {
   namespace = parts[0]
  }
  normalized := "." + strings.ToLower(namespace)
  if strings.HasSuffix(normalized, ".kubernetes.io") || strings.HasSuffix(normalized, ".k8s.io") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("options").Key(k), k, "kubernetes.io and k8s.io namespaces are reserved"))
  }
 }
 return allErrs
}
func validateFlexPersistentVolumeSource(fv *core.FlexPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(fv.Driver) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("driver"), ""))
 }
 for k := range fv.Options {
  namespace := k
  if parts := strings.SplitN(k, "/", 2); len(parts) == 2 {
   namespace = parts[0]
  }
  normalized := "." + strings.ToLower(namespace)
  if strings.HasSuffix(normalized, ".kubernetes.io") || strings.HasSuffix(normalized, ".k8s.io") {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("options").Key(k), k, "kubernetes.io and k8s.io namespaces are reserved"))
  }
 }
 return allErrs
}
func validateAzureFile(azure *core.AzureFileVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if azure.SecretName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("secretName"), ""))
 }
 if azure.ShareName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("shareName"), ""))
 }
 return allErrs
}
func validateAzureFilePV(azure *core.AzureFilePersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if azure.SecretName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("secretName"), ""))
 }
 if azure.ShareName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("shareName"), ""))
 }
 if azure.SecretNamespace != nil {
  if len(*azure.SecretNamespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretNamespace"), ""))
  }
 }
 return allErrs
}
func validateAzureDisk(azure *core.AzureDiskVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var supportedCachingModes = sets.NewString(string(core.AzureDataDiskCachingNone), string(core.AzureDataDiskCachingReadOnly), string(core.AzureDataDiskCachingReadWrite))
 var supportedDiskKinds = sets.NewString(string(core.AzureSharedBlobDisk), string(core.AzureDedicatedBlobDisk), string(core.AzureManagedDisk))
 diskUriSupportedManaged := []string{"/subscriptions/{sub-id}/resourcegroups/{group-name}/providers/microsoft.compute/disks/{disk-id}"}
 diskUriSupportedblob := []string{"https://{account-name}.blob.core.windows.net/{container-name}/{disk-name}.vhd"}
 allErrs := field.ErrorList{}
 if azure.DiskName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("diskName"), ""))
 }
 if azure.DataDiskURI == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("diskURI"), ""))
 }
 if azure.CachingMode != nil && !supportedCachingModes.Has(string(*azure.CachingMode)) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("cachingMode"), *azure.CachingMode, supportedCachingModes.List()))
 }
 if azure.Kind != nil && !supportedDiskKinds.Has(string(*azure.Kind)) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("kind"), *azure.Kind, supportedDiskKinds.List()))
 }
 if azure.Kind != nil && *azure.Kind == core.AzureManagedDisk && strings.Index(azure.DataDiskURI, "/subscriptions/") != 0 {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("diskURI"), azure.DataDiskURI, diskUriSupportedManaged))
 }
 if azure.Kind != nil && *azure.Kind != core.AzureManagedDisk && strings.Index(azure.DataDiskURI, "https://") != 0 {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("diskURI"), azure.DataDiskURI, diskUriSupportedblob))
 }
 return allErrs
}
func validateVsphereVolumeSource(cd *core.VsphereVirtualDiskVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cd.VolumePath) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumePath"), ""))
 }
 return allErrs
}
func validatePhotonPersistentDiskVolumeSource(cd *core.PhotonPersistentDiskVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(cd.PdID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("pdID"), ""))
 }
 return allErrs
}
func validatePortworxVolumeSource(pwx *core.PortworxVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(pwx.VolumeID) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeID"), ""))
 }
 return allErrs
}
func validateScaleIOVolumeSource(sio *core.ScaleIOVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if sio.Gateway == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("gateway"), ""))
 }
 if sio.System == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("system"), ""))
 }
 if sio.VolumeName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeName"), ""))
 }
 return allErrs
}
func validateScaleIOPersistentVolumeSource(sio *core.ScaleIOPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if sio.Gateway == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("gateway"), ""))
 }
 if sio.System == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("system"), ""))
 }
 if sio.VolumeName == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeName"), ""))
 }
 return allErrs
}
func validateLocalVolumeSource(ls *core.LocalVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if ls.Path == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("path"), ""))
  return allErrs
 }
 allErrs = append(allErrs, validatePathNoBacksteps(ls.Path, fldPath.Child("path"))...)
 return allErrs
}
func validateStorageOSVolumeSource(storageos *core.StorageOSVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(storageos.VolumeName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeName"), ""))
 } else {
  allErrs = append(allErrs, ValidateDNS1123Label(storageos.VolumeName, fldPath.Child("volumeName"))...)
 }
 if len(storageos.VolumeNamespace) > 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(storageos.VolumeNamespace, fldPath.Child("volumeNamespace"))...)
 }
 if storageos.SecretRef != nil {
  if len(storageos.SecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "name"), ""))
  }
 }
 return allErrs
}
func validateStorageOSPersistentVolumeSource(storageos *core.StorageOSPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(storageos.VolumeName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeName"), ""))
 } else {
  allErrs = append(allErrs, ValidateDNS1123Label(storageos.VolumeName, fldPath.Child("volumeName"))...)
 }
 if len(storageos.VolumeNamespace) > 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(storageos.VolumeNamespace, fldPath.Child("volumeNamespace"))...)
 }
 if storageos.SecretRef != nil {
  if len(storageos.SecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "name"), ""))
  }
  if len(storageos.SecretRef.Namespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("secretRef", "namespace"), ""))
  }
 }
 return allErrs
}
func ValidateCSIDriverName(driverName string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(driverName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath, ""))
 }
 if len(driverName) > 63 {
  allErrs = append(allErrs, field.TooLong(fldPath, driverName, 63))
 }
 if !csiDriverNameRexp.MatchString(driverName) {
  allErrs = append(allErrs, field.Invalid(fldPath, driverName, validation.RegexError(csiDriverNameRexpErrMsg, csiDriverNameRexpFmt, "csi-hostpath")))
 }
 return allErrs
}
func validateCSIPersistentVolumeSource(csi *core.CSIPersistentVolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateCSIDriverName(csi.Driver, fldPath.Child("driver"))...)
 if len(csi.VolumeHandle) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("volumeHandle"), ""))
 }
 if csi.ControllerPublishSecretRef != nil {
  if len(csi.ControllerPublishSecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("controllerPublishSecretRef", "name"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.ControllerPublishSecretRef.Name, fldPath.Child("name"))...)
  }
  if len(csi.ControllerPublishSecretRef.Namespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("controllerPublishSecretRef", "namespace"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.ControllerPublishSecretRef.Namespace, fldPath.Child("namespace"))...)
  }
 }
 if csi.NodePublishSecretRef != nil {
  if len(csi.NodePublishSecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("nodePublishSecretRef ", "name"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.NodePublishSecretRef.Name, fldPath.Child("name"))...)
  }
  if len(csi.NodePublishSecretRef.Namespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("nodePublishSecretRef ", "namespace"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.NodePublishSecretRef.Namespace, fldPath.Child("namespace"))...)
  }
 }
 if csi.NodeStageSecretRef != nil {
  if len(csi.NodeStageSecretRef.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("nodeStageSecretRef", "name"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.NodeStageSecretRef.Name, fldPath.Child("name"))...)
  }
  if len(csi.NodeStageSecretRef.Namespace) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("nodeStageSecretRef", "namespace"), ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(csi.NodeStageSecretRef.Namespace, fldPath.Child("namespace"))...)
  }
 }
 return allErrs
}

var ValidatePersistentVolumeName = apimachineryvalidation.NameIsDNSSubdomain
var supportedAccessModes = sets.NewString(string(core.ReadWriteOnce), string(core.ReadOnlyMany), string(core.ReadWriteMany))
var supportedReclaimPolicy = sets.NewString(string(core.PersistentVolumeReclaimDelete), string(core.PersistentVolumeReclaimRecycle), string(core.PersistentVolumeReclaimRetain))
var supportedVolumeModes = sets.NewString(string(core.PersistentVolumeBlock), string(core.PersistentVolumeFilesystem))
var supportedDataSourceAPIGroupKinds = map[schema.GroupKind]bool{{Group: "snapshot.storage.k8s.io", Kind: "VolumeSnapshot"}: true}

func ValidatePersistentVolume(pv *core.PersistentVolume) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 metaPath := field.NewPath("metadata")
 allErrs := ValidateObjectMeta(&pv.ObjectMeta, false, ValidatePersistentVolumeName, metaPath)
 specPath := field.NewPath("spec")
 if len(pv.Spec.AccessModes) == 0 {
  allErrs = append(allErrs, field.Required(specPath.Child("accessModes"), ""))
 }
 for _, mode := range pv.Spec.AccessModes {
  if !supportedAccessModes.Has(string(mode)) {
   allErrs = append(allErrs, field.NotSupported(specPath.Child("accessModes"), mode, supportedAccessModes.List()))
  }
 }
 if len(pv.Spec.Capacity) == 0 {
  allErrs = append(allErrs, field.Required(specPath.Child("capacity"), ""))
 }
 if _, ok := pv.Spec.Capacity[core.ResourceStorage]; !ok || len(pv.Spec.Capacity) > 1 {
  allErrs = append(allErrs, field.NotSupported(specPath.Child("capacity"), pv.Spec.Capacity, []string{string(core.ResourceStorage)}))
 }
 capPath := specPath.Child("capacity")
 for r, qty := range pv.Spec.Capacity {
  allErrs = append(allErrs, validateBasicResource(qty, capPath.Key(string(r)))...)
  allErrs = append(allErrs, ValidatePositiveQuantityValue(qty, capPath.Key(string(r)))...)
 }
 if len(string(pv.Spec.PersistentVolumeReclaimPolicy)) > 0 {
  if !supportedReclaimPolicy.Has(string(pv.Spec.PersistentVolumeReclaimPolicy)) {
   allErrs = append(allErrs, field.NotSupported(specPath.Child("persistentVolumeReclaimPolicy"), pv.Spec.PersistentVolumeReclaimPolicy, supportedReclaimPolicy.List()))
  }
 }
 nodeAffinitySpecified, errs := validateVolumeNodeAffinity(pv.Spec.NodeAffinity, specPath.Child("nodeAffinity"))
 allErrs = append(allErrs, errs...)
 numVolumes := 0
 if pv.Spec.HostPath != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("hostPath"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateHostPathVolumeSource(pv.Spec.HostPath, specPath.Child("hostPath"))...)
  }
 }
 if pv.Spec.GCEPersistentDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("gcePersistentDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateGCEPersistentDiskVolumeSource(pv.Spec.GCEPersistentDisk, specPath.Child("persistentDisk"))...)
  }
 }
 if pv.Spec.AWSElasticBlockStore != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("awsElasticBlockStore"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAWSElasticBlockStoreVolumeSource(pv.Spec.AWSElasticBlockStore, specPath.Child("awsElasticBlockStore"))...)
  }
 }
 if pv.Spec.Glusterfs != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("glusterfs"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateGlusterfsPersistentVolumeSource(pv.Spec.Glusterfs, specPath.Child("glusterfs"))...)
  }
 }
 if pv.Spec.Flocker != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("flocker"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateFlockerVolumeSource(pv.Spec.Flocker, specPath.Child("flocker"))...)
  }
 }
 if pv.Spec.NFS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("nfs"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateNFSVolumeSource(pv.Spec.NFS, specPath.Child("nfs"))...)
  }
 }
 if pv.Spec.RBD != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("rbd"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateRBDPersistentVolumeSource(pv.Spec.RBD, specPath.Child("rbd"))...)
  }
 }
 if pv.Spec.Quobyte != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("quobyte"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateQuobyteVolumeSource(pv.Spec.Quobyte, specPath.Child("quobyte"))...)
  }
 }
 if pv.Spec.CephFS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("cephFS"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateCephFSPersistentVolumeSource(pv.Spec.CephFS, specPath.Child("cephfs"))...)
  }
 }
 if pv.Spec.ISCSI != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("iscsi"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateISCSIPersistentVolumeSource(pv.Spec.ISCSI, specPath.Child("iscsi"))...)
  }
  if pv.Spec.ISCSI.InitiatorName != nil && len(pv.ObjectMeta.Name+":"+pv.Spec.ISCSI.TargetPortal) > 64 {
   tooLongErr := "Total length of <volume name>:<iscsi.targetPortal> must be under 64 characters if iscsi.initiatorName is specified."
   allErrs = append(allErrs, field.Invalid(metaPath.Child("name"), pv.ObjectMeta.Name, tooLongErr))
  }
 }
 if pv.Spec.Cinder != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("cinder"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateCinderPersistentVolumeSource(pv.Spec.Cinder, specPath.Child("cinder"))...)
  }
 }
 if pv.Spec.FC != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("fc"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateFCVolumeSource(pv.Spec.FC, specPath.Child("fc"))...)
  }
 }
 if pv.Spec.FlexVolume != nil {
  numVolumes++
  allErrs = append(allErrs, validateFlexPersistentVolumeSource(pv.Spec.FlexVolume, specPath.Child("flexVolume"))...)
 }
 if pv.Spec.AzureFile != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("azureFile"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAzureFilePV(pv.Spec.AzureFile, specPath.Child("azureFile"))...)
  }
 }
 if pv.Spec.VsphereVolume != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("vsphereVolume"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateVsphereVolumeSource(pv.Spec.VsphereVolume, specPath.Child("vsphereVolume"))...)
  }
 }
 if pv.Spec.PhotonPersistentDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("photonPersistentDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validatePhotonPersistentDiskVolumeSource(pv.Spec.PhotonPersistentDisk, specPath.Child("photonPersistentDisk"))...)
  }
 }
 if pv.Spec.PortworxVolume != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("portworxVolume"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validatePortworxVolumeSource(pv.Spec.PortworxVolume, specPath.Child("portworxVolume"))...)
  }
 }
 if pv.Spec.AzureDisk != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("azureDisk"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateAzureDisk(pv.Spec.AzureDisk, specPath.Child("azureDisk"))...)
  }
 }
 if pv.Spec.ScaleIO != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("scaleIO"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateScaleIOPersistentVolumeSource(pv.Spec.ScaleIO, specPath.Child("scaleIO"))...)
  }
 }
 if pv.Spec.Local != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("local"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   if !utilfeature.DefaultFeatureGate.Enabled(features.PersistentLocalVolumes) {
    allErrs = append(allErrs, field.Forbidden(specPath.Child("local"), "Local volumes are disabled by feature-gate"))
   }
   allErrs = append(allErrs, validateLocalVolumeSource(pv.Spec.Local, specPath.Child("local"))...)
   if !nodeAffinitySpecified {
    allErrs = append(allErrs, field.Required(metaPath.Child("annotations"), "Local volume requires node affinity"))
   }
  }
 }
 if pv.Spec.StorageOS != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("storageos"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateStorageOSPersistentVolumeSource(pv.Spec.StorageOS, specPath.Child("storageos"))...)
  }
 }
 if pv.Spec.CSI != nil {
  if numVolumes > 0 {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("csi"), "may not specify more than 1 volume type"))
  } else {
   numVolumes++
   allErrs = append(allErrs, validateCSIPersistentVolumeSource(pv.Spec.CSI, specPath.Child("csi"))...)
  }
 }
 if numVolumes == 0 {
  allErrs = append(allErrs, field.Required(specPath, "must specify a volume type"))
 }
 if pv.Spec.HostPath != nil && path.Clean(pv.Spec.HostPath.Path) == "/" && pv.Spec.PersistentVolumeReclaimPolicy == core.PersistentVolumeReclaimRecycle {
  allErrs = append(allErrs, field.Forbidden(specPath.Child("persistentVolumeReclaimPolicy"), "may not be 'recycle' for a hostPath mount of '/'"))
 }
 if len(pv.Spec.StorageClassName) > 0 {
  for _, msg := range ValidateClassName(pv.Spec.StorageClassName, false) {
   allErrs = append(allErrs, field.Invalid(specPath.Child("storageClassName"), pv.Spec.StorageClassName, msg))
  }
 }
 if pv.Spec.VolumeMode != nil && !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  allErrs = append(allErrs, field.Forbidden(specPath.Child("volumeMode"), "PersistentVolume volumeMode is disabled by feature-gate"))
 } else if pv.Spec.VolumeMode != nil && !supportedVolumeModes.Has(string(*pv.Spec.VolumeMode)) {
  allErrs = append(allErrs, field.NotSupported(specPath.Child("volumeMode"), *pv.Spec.VolumeMode, supportedVolumeModes.List()))
 }
 return allErrs
}
func ValidatePersistentVolumeUpdate(newPv, oldPv *core.PersistentVolume) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = ValidatePersistentVolume(newPv)
 if !apiequality.Semantic.DeepEqual(newPv.Spec.PersistentVolumeSource, oldPv.Spec.PersistentVolumeSource) {
  allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "persistentvolumesource"), "is immutable after creation"))
 }
 newPv.Status = oldPv.Status
 if utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  allErrs = append(allErrs, ValidateImmutableField(newPv.Spec.VolumeMode, oldPv.Spec.VolumeMode, field.NewPath("volumeMode"))...)
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
  if oldPv.Spec.NodeAffinity != nil {
   allErrs = append(allErrs, ValidateImmutableField(newPv.Spec.NodeAffinity, oldPv.Spec.NodeAffinity, field.NewPath("nodeAffinity"))...)
  }
 }
 return allErrs
}
func ValidatePersistentVolumeStatusUpdate(newPv, oldPv *core.PersistentVolume) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newPv.ObjectMeta, &oldPv.ObjectMeta, field.NewPath("metadata"))
 if len(newPv.ResourceVersion) == 0 {
  allErrs = append(allErrs, field.Required(field.NewPath("resourceVersion"), ""))
 }
 newPv.Spec = oldPv.Spec
 return allErrs
}
func ValidatePersistentVolumeClaim(pvc *core.PersistentVolumeClaim) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&pvc.ObjectMeta, true, ValidatePersistentVolumeName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidatePersistentVolumeClaimSpec(&pvc.Spec, field.NewPath("spec"))...)
 return allErrs
}
func ValidatePersistentVolumeClaimSpec(spec *core.PersistentVolumeClaimSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(spec.AccessModes) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("accessModes"), "at least 1 access mode is required"))
 }
 if spec.Selector != nil {
  allErrs = append(allErrs, unversionedvalidation.ValidateLabelSelector(spec.Selector, fldPath.Child("selector"))...)
 }
 for _, mode := range spec.AccessModes {
  if mode != core.ReadWriteOnce && mode != core.ReadOnlyMany && mode != core.ReadWriteMany {
   allErrs = append(allErrs, field.NotSupported(fldPath.Child("accessModes"), mode, supportedAccessModes.List()))
  }
 }
 storageValue, ok := spec.Resources.Requests[core.ResourceStorage]
 if !ok {
  allErrs = append(allErrs, field.Required(fldPath.Child("resources").Key(string(core.ResourceStorage)), ""))
 } else {
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(core.ResourceStorage), storageValue, fldPath.Child("resources").Key(string(core.ResourceStorage)))...)
  allErrs = append(allErrs, ValidatePositiveQuantityValue(storageValue, fldPath.Child("resources").Key(string(core.ResourceStorage)))...)
 }
 if spec.StorageClassName != nil && len(*spec.StorageClassName) > 0 {
  for _, msg := range ValidateClassName(*spec.StorageClassName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("storageClassName"), *spec.StorageClassName, msg))
  }
 }
 if spec.VolumeMode != nil && !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("volumeMode"), "PersistentVolumeClaim volumeMode is disabled by feature-gate"))
 } else if spec.VolumeMode != nil && !supportedVolumeModes.Has(string(*spec.VolumeMode)) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("volumeMode"), *spec.VolumeMode, supportedVolumeModes.List()))
 }
 if spec.DataSource != nil && !utilfeature.DefaultFeatureGate.Enabled(features.VolumeSnapshotDataSource) {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("dataSource"), "VolumeSnapshotDataSource is disabled by feature-gate"))
 } else if spec.DataSource != nil {
  if len(spec.DataSource.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("dataSource", "name"), ""))
  }
  groupKind := schema.GroupKind{Group: "", Kind: spec.DataSource.Kind}
  if spec.DataSource.APIGroup != nil {
   groupKind.Group = string(*spec.DataSource.APIGroup)
  }
  groupKindList := make([]string, 0, len(supportedDataSourceAPIGroupKinds))
  for grp := range supportedDataSourceAPIGroupKinds {
   groupKindList = append(groupKindList, grp.String())
  }
  if !supportedDataSourceAPIGroupKinds[groupKind] {
   allErrs = append(allErrs, field.NotSupported(fldPath.Child("dataSource"), groupKind.String(), groupKindList))
  }
 }
 return allErrs
}
func ValidatePersistentVolumeClaimUpdate(newPvc, oldPvc *core.PersistentVolumeClaim) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newPvc.ObjectMeta, &oldPvc.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidatePersistentVolumeClaim(newPvc)...)
 newPvcClone := newPvc.DeepCopy()
 oldPvcClone := oldPvc.DeepCopy()
 if len(oldPvc.Spec.VolumeName) == 0 {
  oldPvcClone.Spec.VolumeName = newPvcClone.Spec.VolumeName
 }
 if validateStorageClassUpgrade(oldPvcClone.Annotations, newPvcClone.Annotations, oldPvcClone.Spec.StorageClassName, newPvcClone.Spec.StorageClassName) {
  newPvcClone.Spec.StorageClassName = nil
  metav1.SetMetaDataAnnotation(&newPvcClone.ObjectMeta, core.BetaStorageClassAnnotation, oldPvcClone.Annotations[core.BetaStorageClassAnnotation])
 } else {
  allErrs = append(allErrs, ValidateImmutableAnnotation(newPvc.ObjectMeta.Annotations[v1.BetaStorageClassAnnotation], oldPvc.ObjectMeta.Annotations[v1.BetaStorageClassAnnotation], v1.BetaStorageClassAnnotation, field.NewPath("metadata"))...)
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) {
  if newPvc.Status.Phase == core.ClaimBound && newPvcClone.Spec.Resources.Requests != nil {
   newPvcClone.Spec.Resources.Requests["storage"] = oldPvc.Spec.Resources.Requests["storage"]
  }
  oldSize := oldPvc.Spec.Resources.Requests["storage"]
  newSize := newPvc.Spec.Resources.Requests["storage"]
  if !apiequality.Semantic.DeepEqual(newPvcClone.Spec, oldPvcClone.Spec) {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec"), "is immutable after creation except resources.requests for bound claims"))
  }
  if newSize.Cmp(oldSize) < 0 {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "resources", "requests", "storage"), "field can not be less than previous value"))
  }
 } else {
  if !apiequality.Semantic.DeepEqual(newPvcClone.Spec, oldPvcClone.Spec) {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec"), "field is immutable after creation"))
  }
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  allErrs = append(allErrs, ValidateImmutableField(newPvc.Spec.VolumeMode, oldPvc.Spec.VolumeMode, field.NewPath("volumeMode"))...)
 }
 return allErrs
}
func validateStorageClassUpgrade(oldAnnotations, newAnnotations map[string]string, oldScName, newScName *string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 oldSc, oldAnnotationExist := oldAnnotations[core.BetaStorageClassAnnotation]
 newScInAnnotation, newAnnotationExist := newAnnotations[core.BetaStorageClassAnnotation]
 return oldAnnotationExist && oldScName == nil && (newScName != nil && *newScName == oldSc) && (!newAnnotationExist || newScInAnnotation == oldSc)
}
func ValidatePersistentVolumeClaimStatusUpdate(newPvc, oldPvc *core.PersistentVolumeClaim) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newPvc.ObjectMeta, &oldPvc.ObjectMeta, field.NewPath("metadata"))
 if len(newPvc.ResourceVersion) == 0 {
  allErrs = append(allErrs, field.Required(field.NewPath("resourceVersion"), ""))
 }
 if len(newPvc.Spec.AccessModes) == 0 {
  allErrs = append(allErrs, field.Required(field.NewPath("Spec", "accessModes"), ""))
 }
 if !utilfeature.DefaultFeatureGate.Enabled(features.ExpandPersistentVolumes) && len(newPvc.Status.Conditions) > 0 {
  conditionPath := field.NewPath("status", "conditions")
  allErrs = append(allErrs, field.Forbidden(conditionPath, "invalid field"))
 }
 capPath := field.NewPath("status", "capacity")
 for r, qty := range newPvc.Status.Capacity {
  allErrs = append(allErrs, validateBasicResource(qty, capPath.Key(string(r)))...)
 }
 newPvc.Spec = oldPvc.Spec
 return allErrs
}

var supportedPortProtocols = sets.NewString(string(core.ProtocolTCP), string(core.ProtocolUDP), string(core.ProtocolSCTP))

func validateContainerPorts(ports []core.ContainerPort, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allNames := sets.String{}
 for i, port := range ports {
  idxPath := fldPath.Index(i)
  if len(port.Name) > 0 {
   if msgs := validation.IsValidPortName(port.Name); len(msgs) != 0 {
    for i = range msgs {
     allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), port.Name, msgs[i]))
    }
   } else if allNames.Has(port.Name) {
    allErrs = append(allErrs, field.Duplicate(idxPath.Child("name"), port.Name))
   } else {
    allNames.Insert(port.Name)
   }
  }
  if port.ContainerPort == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("containerPort"), ""))
  } else {
   for _, msg := range validation.IsValidPortNum(int(port.ContainerPort)) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("containerPort"), port.ContainerPort, msg))
   }
  }
  if port.HostPort != 0 {
   for _, msg := range validation.IsValidPortNum(int(port.HostPort)) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("hostPort"), port.HostPort, msg))
   }
  }
  if len(port.Protocol) == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("protocol"), ""))
  } else if !utilfeature.DefaultFeatureGate.Enabled(features.SCTPSupport) && port.Protocol == core.ProtocolSCTP {
   allErrs = append(allErrs, field.NotSupported(idxPath.Child("protocol"), port.Protocol, []string{string(core.ProtocolTCP), string(core.ProtocolUDP)}))
  } else if !supportedPortProtocols.Has(string(port.Protocol)) {
   allErrs = append(allErrs, field.NotSupported(idxPath.Child("protocol"), port.Protocol, supportedPortProtocols.List()))
  }
 }
 return allErrs
}
func ValidateEnv(vars []core.EnvVar, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, ev := range vars {
  idxPath := fldPath.Index(i)
  if len(ev.Name) == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("name"), ""))
  } else {
   for _, msg := range validation.IsEnvVarName(ev.Name) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), ev.Name, msg))
   }
  }
  allErrs = append(allErrs, validateEnvVarValueFrom(ev, idxPath.Child("valueFrom"))...)
 }
 return allErrs
}

var validEnvDownwardAPIFieldPathExpressions = sets.NewString("metadata.name", "metadata.namespace", "metadata.uid", "spec.nodeName", "spec.serviceAccountName", "status.hostIP", "status.podIP")
var validContainerResourceFieldPathExpressions = sets.NewString("limits.cpu", "limits.memory", "limits.ephemeral-storage", "requests.cpu", "requests.memory", "requests.ephemeral-storage")

func validateEnvVarValueFrom(ev core.EnvVar, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if ev.ValueFrom == nil {
  return allErrs
 }
 numSources := 0
 if ev.ValueFrom.FieldRef != nil {
  numSources++
  allErrs = append(allErrs, validateObjectFieldSelector(ev.ValueFrom.FieldRef, &validEnvDownwardAPIFieldPathExpressions, fldPath.Child("fieldRef"))...)
 }
 if ev.ValueFrom.ResourceFieldRef != nil {
  numSources++
  allErrs = append(allErrs, validateContainerResourceFieldSelector(ev.ValueFrom.ResourceFieldRef, &validContainerResourceFieldPathExpressions, fldPath.Child("resourceFieldRef"), false)...)
 }
 if ev.ValueFrom.ConfigMapKeyRef != nil {
  numSources++
  allErrs = append(allErrs, validateConfigMapKeySelector(ev.ValueFrom.ConfigMapKeyRef, fldPath.Child("configMapKeyRef"))...)
 }
 if ev.ValueFrom.SecretKeyRef != nil {
  numSources++
  allErrs = append(allErrs, validateSecretKeySelector(ev.ValueFrom.SecretKeyRef, fldPath.Child("secretKeyRef"))...)
 }
 if numSources == 0 {
  allErrs = append(allErrs, field.Invalid(fldPath, "", "must specify one of: `fieldRef`, `resourceFieldRef`, `configMapKeyRef` or `secretKeyRef`"))
 } else if len(ev.Value) != 0 {
  if numSources != 0 {
   allErrs = append(allErrs, field.Invalid(fldPath, "", "may not be specified when `value` is not empty"))
  }
 } else if numSources > 1 {
  allErrs = append(allErrs, field.Invalid(fldPath, "", "may not have more than one field specified at a time"))
 }
 return allErrs
}
func validateObjectFieldSelector(fs *core.ObjectFieldSelector, expressions *sets.String, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(fs.APIVersion) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("apiVersion"), ""))
  return allErrs
 }
 if len(fs.FieldPath) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("fieldPath"), ""))
  return allErrs
 }
 internalFieldPath, _, err := podshelper.ConvertDownwardAPIFieldLabel(fs.APIVersion, fs.FieldPath, "")
 if err != nil {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("fieldPath"), fs.FieldPath, fmt.Sprintf("error converting fieldPath: %v", err)))
  return allErrs
 }
 if path, subscript, ok := fieldpath.SplitMaybeSubscriptedPath(internalFieldPath); ok {
  switch path {
  case "metadata.annotations":
   for _, msg := range validation.IsQualifiedName(strings.ToLower(subscript)) {
    allErrs = append(allErrs, field.Invalid(fldPath, subscript, msg))
   }
  case "metadata.labels":
   for _, msg := range validation.IsQualifiedName(subscript) {
    allErrs = append(allErrs, field.Invalid(fldPath, subscript, msg))
   }
  default:
   allErrs = append(allErrs, field.Invalid(fldPath, path, "does not support subscript"))
  }
 } else if !expressions.Has(path) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("fieldPath"), path, expressions.List()))
  return allErrs
 }
 return allErrs
}
func fsResourceIsEphemeralStorage(resource string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if resource == "limits.ephemeral-storage" || resource == "requests.ephemeral-storage" {
  return true
 }
 return false
}
func validateContainerResourceFieldSelector(fs *core.ResourceFieldSelector, expressions *sets.String, fldPath *field.Path, volume bool) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if volume && len(fs.ContainerName) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("containerName"), ""))
 } else if len(fs.Resource) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("resource"), ""))
 } else if !expressions.Has(fs.Resource) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("resource"), fs.Resource, expressions.List()))
 } else if fsResourceIsEphemeralStorage(fs.Resource) && !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
  allErrs = append(allErrs, field.Forbidden(fldPath, "Containers' ephemeral storage requests/limits disabled by feature-gate for Downward API"))
 }
 allErrs = append(allErrs, validateContainerResourceDivisor(fs.Resource, fs.Divisor, fldPath)...)
 return allErrs
}
func ValidateEnvFrom(vars []core.EnvFromSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, ev := range vars {
  idxPath := fldPath.Index(i)
  if len(ev.Prefix) > 0 {
   for _, msg := range validation.IsEnvVarName(ev.Prefix) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("prefix"), ev.Prefix, msg))
   }
  }
  numSources := 0
  if ev.ConfigMapRef != nil {
   numSources++
   allErrs = append(allErrs, validateConfigMapEnvSource(ev.ConfigMapRef, idxPath.Child("configMapRef"))...)
  }
  if ev.SecretRef != nil {
   numSources++
   allErrs = append(allErrs, validateSecretEnvSource(ev.SecretRef, idxPath.Child("secretRef"))...)
  }
  if numSources == 0 {
   allErrs = append(allErrs, field.Invalid(fldPath, "", "must specify one of: `configMapRef` or `secretRef`"))
  } else if numSources > 1 {
   allErrs = append(allErrs, field.Invalid(fldPath, "", "may not have more than one field specified at a time"))
  }
 }
 return allErrs
}
func validateConfigMapEnvSource(configMapSource *core.ConfigMapEnvSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(configMapSource.Name) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
 } else {
  for _, msg := range ValidateConfigMapName(configMapSource.Name, true) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), configMapSource.Name, msg))
  }
 }
 return allErrs
}
func validateSecretEnvSource(secretSource *core.SecretEnvSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(secretSource.Name) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
 } else {
  for _, msg := range ValidateSecretName(secretSource.Name, true) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), secretSource.Name, msg))
  }
 }
 return allErrs
}

var validContainerResourceDivisorForCPU = sets.NewString("1m", "1")
var validContainerResourceDivisorForMemory = sets.NewString("1", "1k", "1M", "1G", "1T", "1P", "1E", "1Ki", "1Mi", "1Gi", "1Ti", "1Pi", "1Ei")
var validContainerResourceDivisorForEphemeralStorage = sets.NewString("1", "1k", "1M", "1G", "1T", "1P", "1E", "1Ki", "1Mi", "1Gi", "1Ti", "1Pi", "1Ei")

func validateContainerResourceDivisor(rName string, divisor resource.Quantity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 unsetDivisor := resource.Quantity{}
 if unsetDivisor.Cmp(divisor) == 0 {
  return allErrs
 }
 switch rName {
 case "limits.cpu", "requests.cpu":
  if !validContainerResourceDivisorForCPU.Has(divisor.String()) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("divisor"), rName, "only divisor's values 1m and 1 are supported with the cpu resource"))
  }
 case "limits.memory", "requests.memory":
  if !validContainerResourceDivisorForMemory.Has(divisor.String()) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("divisor"), rName, "only divisor's values 1, 1k, 1M, 1G, 1T, 1P, 1E, 1Ki, 1Mi, 1Gi, 1Ti, 1Pi, 1Ei are supported with the memory resource"))
  }
 case "limits.ephemeral-storage", "requests.ephemeral-storage":
  if !validContainerResourceDivisorForEphemeralStorage.Has(divisor.String()) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("divisor"), rName, "only divisor's values 1, 1k, 1M, 1G, 1T, 1P, 1E, 1Ki, 1Mi, 1Gi, 1Ti, 1Pi, 1Ei are supported with the local ephemeral storage resource"))
  }
 }
 return allErrs
}
func validateConfigMapKeySelector(s *core.ConfigMapKeySelector, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 nameFn := ValidateNameFunc(ValidateSecretName)
 for _, msg := range nameFn(s.Name, false) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), s.Name, msg))
 }
 if len(s.Key) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("key"), ""))
 } else {
  for _, msg := range validation.IsConfigMapKey(s.Key) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("key"), s.Key, msg))
  }
 }
 return allErrs
}
func validateSecretKeySelector(s *core.SecretKeySelector, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 nameFn := ValidateNameFunc(ValidateSecretName)
 for _, msg := range nameFn(s.Name, false) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), s.Name, msg))
 }
 if len(s.Key) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("key"), ""))
 } else {
  for _, msg := range validation.IsConfigMapKey(s.Key) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("key"), s.Key, msg))
  }
 }
 return allErrs
}
func GetVolumeMountMap(mounts []core.VolumeMount) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 volmounts := make(map[string]string)
 for _, mnt := range mounts {
  volmounts[mnt.Name] = mnt.MountPath
 }
 return volmounts
}
func GetVolumeDeviceMap(devices []core.VolumeDevice) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 voldevices := make(map[string]string)
 for _, dev := range devices {
  voldevices[dev.Name] = dev.DevicePath
 }
 return voldevices
}
func ValidateVolumeMounts(mounts []core.VolumeMount, voldevices map[string]string, volumes map[string]core.VolumeSource, container *core.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 mountpoints := sets.NewString()
 for i, mnt := range mounts {
  idxPath := fldPath.Index(i)
  if len(mnt.Name) == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("name"), ""))
  }
  if !IsMatchedVolume(mnt.Name, volumes) {
   allErrs = append(allErrs, field.NotFound(idxPath.Child("name"), mnt.Name))
  }
  if len(mnt.MountPath) == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("mountPath"), ""))
  }
  if mountpoints.Has(mnt.MountPath) {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("mountPath"), mnt.MountPath, "must be unique"))
  }
  mountpoints.Insert(mnt.MountPath)
  if mountNameAlreadyExists(mnt.Name, voldevices) {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), mnt.Name, "must not already exist in volumeDevices"))
  }
  if mountPathAlreadyExists(mnt.MountPath, voldevices) {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("mountPath"), mnt.MountPath, "must not already exist as a path in volumeDevices"))
  }
  if len(mnt.SubPath) > 0 {
   if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeSubpath) {
    allErrs = append(allErrs, field.Forbidden(fldPath.Child("subPath"), "subPath is disabled by feature-gate"))
   } else {
    allErrs = append(allErrs, validateLocalDescendingPath(mnt.SubPath, fldPath.Child("subPath"))...)
   }
  }
  if mnt.MountPropagation != nil {
   allErrs = append(allErrs, validateMountPropagation(mnt.MountPropagation, container, fldPath.Child("mountPropagation"))...)
  }
 }
 return allErrs
}
func ValidateVolumeDevices(devices []core.VolumeDevice, volmounts map[string]string, volumes map[string]core.VolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 devicepath := sets.NewString()
 devicename := sets.NewString()
 if devices != nil && !utilfeature.DefaultFeatureGate.Enabled(features.BlockVolume) {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("volumeDevices"), "Container volumeDevices is disabled by feature-gate"))
  return allErrs
 }
 if devices != nil {
  for i, dev := range devices {
   idxPath := fldPath.Index(i)
   devName := dev.Name
   devPath := dev.DevicePath
   didMatch, isPVC := isMatchedDevice(devName, volumes)
   if len(devName) == 0 {
    allErrs = append(allErrs, field.Required(idxPath.Child("name"), ""))
   }
   if devicename.Has(devName) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), devName, "must be unique"))
   }
   if didMatch && !isPVC {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), devName, "can only use volume source type of PersistentVolumeClaim for block mode"))
   }
   if !didMatch {
    allErrs = append(allErrs, field.NotFound(idxPath.Child("name"), devName))
   }
   if len(devPath) == 0 {
    allErrs = append(allErrs, field.Required(idxPath.Child("devicePath"), ""))
   }
   if devicepath.Has(devPath) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("devicePath"), devPath, "must be unique"))
   }
   if len(devPath) > 0 && len(validatePathNoBacksteps(devPath, fldPath.Child("devicePath"))) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("devicePath"), devPath, "can not contain backsteps ('..')"))
   } else {
    devicepath.Insert(devPath)
   }
   if deviceNameAlreadyExists(devName, volmounts) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("name"), devName, "must not already exist in volumeMounts"))
   }
   if devicePathAlreadyExists(devPath, volmounts) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("devicePath"), devPath, "must not already exist as a path in volumeMounts"))
   }
   if len(devName) > 0 {
    devicename.Insert(devName)
   }
  }
 }
 return allErrs
}
func validateProbe(probe *core.Probe, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if probe == nil {
  return allErrs
 }
 allErrs = append(allErrs, validateHandler(&probe.Handler, fldPath)...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(probe.InitialDelaySeconds), fldPath.Child("initialDelaySeconds"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(probe.TimeoutSeconds), fldPath.Child("timeoutSeconds"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(probe.PeriodSeconds), fldPath.Child("periodSeconds"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(probe.SuccessThreshold), fldPath.Child("successThreshold"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(probe.FailureThreshold), fldPath.Child("failureThreshold"))...)
 return allErrs
}
func validateClientIPAffinityConfig(config *core.SessionAffinityConfig, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if config == nil {
  allErrs = append(allErrs, field.Required(fldPath, fmt.Sprintf("when session affinity type is %s", core.ServiceAffinityClientIP)))
  return allErrs
 }
 if config.ClientIP == nil {
  allErrs = append(allErrs, field.Required(fldPath.Child("clientIP"), fmt.Sprintf("when session affinity type is %s", core.ServiceAffinityClientIP)))
  return allErrs
 }
 if config.ClientIP.TimeoutSeconds == nil {
  allErrs = append(allErrs, field.Required(fldPath.Child("clientIP").Child("timeoutSeconds"), fmt.Sprintf("when session affinity type is %s", core.ServiceAffinityClientIP)))
  return allErrs
 }
 allErrs = append(allErrs, validateAffinityTimeout(config.ClientIP.TimeoutSeconds, fldPath.Child("clientIP").Child("timeoutSeconds"))...)
 return allErrs
}
func validateAffinityTimeout(timeout *int32, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if *timeout <= 0 || *timeout > core.MaxClientIPServiceAffinitySeconds {
  allErrs = append(allErrs, field.Invalid(fldPath, timeout, fmt.Sprintf("must be greater than 0 and less than %d", core.MaxClientIPServiceAffinitySeconds)))
 }
 return allErrs
}
func AccumulateUniqueHostPorts(containers []core.Container, accumulator *sets.String, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for ci, ctr := range containers {
  idxPath := fldPath.Index(ci)
  portsPath := idxPath.Child("ports")
  for pi := range ctr.Ports {
   idxPath := portsPath.Index(pi)
   port := ctr.Ports[pi].HostPort
   if port == 0 {
    continue
   }
   str := fmt.Sprintf("%s/%s/%d", ctr.Ports[pi].Protocol, ctr.Ports[pi].HostIP, port)
   if accumulator.Has(str) {
    allErrs = append(allErrs, field.Duplicate(idxPath.Child("hostPort"), str))
   } else {
    accumulator.Insert(str)
   }
  }
 }
 return allErrs
}
func checkHostPortConflicts(containers []core.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allPorts := sets.String{}
 return AccumulateUniqueHostPorts(containers, &allPorts, fldPath)
}
func validateExecAction(exec *core.ExecAction, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 if len(exec.Command) == 0 {
  allErrors = append(allErrors, field.Required(fldPath.Child("command"), ""))
 }
 return allErrors
}

var supportedHTTPSchemes = sets.NewString(string(core.URISchemeHTTP), string(core.URISchemeHTTPS))

func validateHTTPGetAction(http *core.HTTPGetAction, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 if len(http.Path) == 0 {
  allErrors = append(allErrors, field.Required(fldPath.Child("path"), ""))
 }
 allErrors = append(allErrors, ValidatePortNumOrName(http.Port, fldPath.Child("port"))...)
 if !supportedHTTPSchemes.Has(string(http.Scheme)) {
  allErrors = append(allErrors, field.NotSupported(fldPath.Child("scheme"), http.Scheme, supportedHTTPSchemes.List()))
 }
 for _, header := range http.HTTPHeaders {
  for _, msg := range validation.IsHTTPHeaderName(header.Name) {
   allErrors = append(allErrors, field.Invalid(fldPath.Child("httpHeaders"), header.Name, msg))
  }
 }
 return allErrors
}
func ValidatePortNumOrName(port intstr.IntOrString, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if port.Type == intstr.Int {
  for _, msg := range validation.IsValidPortNum(port.IntValue()) {
   allErrs = append(allErrs, field.Invalid(fldPath, port.IntValue(), msg))
  }
 } else if port.Type == intstr.String {
  for _, msg := range validation.IsValidPortName(port.StrVal) {
   allErrs = append(allErrs, field.Invalid(fldPath, port.StrVal, msg))
  }
 } else {
  allErrs = append(allErrs, field.InternalError(fldPath, fmt.Errorf("unknown type: %v", port.Type)))
 }
 return allErrs
}
func validateTCPSocketAction(tcp *core.TCPSocketAction, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ValidatePortNumOrName(tcp.Port, fldPath.Child("port"))
}
func validateHandler(handler *core.Handler, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 numHandlers := 0
 allErrors := field.ErrorList{}
 if handler.Exec != nil {
  if numHandlers > 0 {
   allErrors = append(allErrors, field.Forbidden(fldPath.Child("exec"), "may not specify more than 1 handler type"))
  } else {
   numHandlers++
   allErrors = append(allErrors, validateExecAction(handler.Exec, fldPath.Child("exec"))...)
  }
 }
 if handler.HTTPGet != nil {
  if numHandlers > 0 {
   allErrors = append(allErrors, field.Forbidden(fldPath.Child("httpGet"), "may not specify more than 1 handler type"))
  } else {
   numHandlers++
   allErrors = append(allErrors, validateHTTPGetAction(handler.HTTPGet, fldPath.Child("httpGet"))...)
  }
 }
 if handler.TCPSocket != nil {
  if numHandlers > 0 {
   allErrors = append(allErrors, field.Forbidden(fldPath.Child("tcpSocket"), "may not specify more than 1 handler type"))
  } else {
   numHandlers++
   allErrors = append(allErrors, validateTCPSocketAction(handler.TCPSocket, fldPath.Child("tcpSocket"))...)
  }
 }
 if numHandlers == 0 {
  allErrors = append(allErrors, field.Required(fldPath, "must specify a handler type"))
 }
 return allErrors
}
func validateLifecycle(lifecycle *core.Lifecycle, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if lifecycle.PostStart != nil {
  allErrs = append(allErrs, validateHandler(lifecycle.PostStart, fldPath.Child("postStart"))...)
 }
 if lifecycle.PreStop != nil {
  allErrs = append(allErrs, validateHandler(lifecycle.PreStop, fldPath.Child("preStop"))...)
 }
 return allErrs
}

var supportedPullPolicies = sets.NewString(string(core.PullAlways), string(core.PullIfNotPresent), string(core.PullNever))

func validatePullPolicy(policy core.PullPolicy, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 switch policy {
 case core.PullAlways, core.PullIfNotPresent, core.PullNever:
  break
 case "":
  allErrors = append(allErrors, field.Required(fldPath, ""))
 default:
  allErrors = append(allErrors, field.NotSupported(fldPath, policy, supportedPullPolicies.List()))
 }
 return allErrors
}
func validateInitContainers(containers, otherContainers []core.Container, deviceVolumes map[string]core.VolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var allErrs field.ErrorList
 if len(containers) > 0 {
  allErrs = append(allErrs, validateContainers(containers, true, deviceVolumes, fldPath)...)
 }
 allNames := sets.String{}
 for _, ctr := range otherContainers {
  allNames.Insert(ctr.Name)
 }
 for i, ctr := range containers {
  idxPath := fldPath.Index(i)
  if allNames.Has(ctr.Name) {
   allErrs = append(allErrs, field.Duplicate(idxPath.Child("name"), ctr.Name))
  }
  if len(ctr.Name) > 0 {
   allNames.Insert(ctr.Name)
  }
  if ctr.Lifecycle != nil {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("lifecycle"), ctr.Lifecycle, "must not be set for init containers"))
  }
  if ctr.LivenessProbe != nil {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("livenessProbe"), ctr.LivenessProbe, "must not be set for init containers"))
  }
  if ctr.ReadinessProbe != nil {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("readinessProbe"), ctr.ReadinessProbe, "must not be set for init containers"))
  }
 }
 return allErrs
}
func validateContainers(containers []core.Container, isInitContainers bool, volumes map[string]core.VolumeSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(containers) == 0 {
  return append(allErrs, field.Required(fldPath, ""))
 }
 allNames := sets.String{}
 for i, ctr := range containers {
  idxPath := fldPath.Index(i)
  namePath := idxPath.Child("name")
  volMounts := GetVolumeMountMap(ctr.VolumeMounts)
  volDevices := GetVolumeDeviceMap(ctr.VolumeDevices)
  if len(ctr.Name) == 0 {
   allErrs = append(allErrs, field.Required(namePath, ""))
  } else {
   allErrs = append(allErrs, ValidateDNS1123Label(ctr.Name, namePath)...)
  }
  if allNames.Has(ctr.Name) {
   allErrs = append(allErrs, field.Duplicate(namePath, ctr.Name))
  } else {
   allNames.Insert(ctr.Name)
  }
  if len(ctr.Image) == 0 {
   allErrs = append(allErrs, field.Required(idxPath.Child("image"), ""))
  }
  if ctr.Lifecycle != nil {
   allErrs = append(allErrs, validateLifecycle(ctr.Lifecycle, idxPath.Child("lifecycle"))...)
  }
  allErrs = append(allErrs, validateProbe(ctr.LivenessProbe, idxPath.Child("livenessProbe"))...)
  if ctr.LivenessProbe != nil && ctr.LivenessProbe.SuccessThreshold != 1 {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("livenessProbe", "successThreshold"), ctr.LivenessProbe.SuccessThreshold, "must be 1"))
  }
  switch ctr.TerminationMessagePolicy {
  case core.TerminationMessageReadFile, core.TerminationMessageFallbackToLogsOnError:
  case "":
   allErrs = append(allErrs, field.Required(idxPath.Child("terminationMessagePolicy"), "must be 'File' or 'FallbackToLogsOnError'"))
  default:
   allErrs = append(allErrs, field.Invalid(idxPath.Child("terminationMessagePolicy"), ctr.TerminationMessagePolicy, "must be 'File' or 'FallbackToLogsOnError'"))
  }
  allErrs = append(allErrs, validateProbe(ctr.ReadinessProbe, idxPath.Child("readinessProbe"))...)
  allErrs = append(allErrs, validateContainerPorts(ctr.Ports, idxPath.Child("ports"))...)
  allErrs = append(allErrs, ValidateEnv(ctr.Env, idxPath.Child("env"))...)
  allErrs = append(allErrs, ValidateEnvFrom(ctr.EnvFrom, idxPath.Child("envFrom"))...)
  allErrs = append(allErrs, ValidateVolumeMounts(ctr.VolumeMounts, volDevices, volumes, &ctr, idxPath.Child("volumeMounts"))...)
  allErrs = append(allErrs, ValidateVolumeDevices(ctr.VolumeDevices, volMounts, volumes, idxPath.Child("volumeDevices"))...)
  allErrs = append(allErrs, validatePullPolicy(ctr.ImagePullPolicy, idxPath.Child("imagePullPolicy"))...)
  allErrs = append(allErrs, ValidateResourceRequirements(&ctr.Resources, idxPath.Child("resources"))...)
  allErrs = append(allErrs, ValidateSecurityContext(ctr.SecurityContext, idxPath.Child("securityContext"))...)
 }
 if isInitContainers {
  for _, initContainer := range containers {
   allErrs = append(allErrs, checkHostPortConflicts([]core.Container{initContainer}, fldPath)...)
  }
 } else {
  allErrs = append(allErrs, checkHostPortConflicts(containers, fldPath)...)
 }
 return allErrs
}
func validateRestartPolicy(restartPolicy *core.RestartPolicy, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 switch *restartPolicy {
 case core.RestartPolicyAlways, core.RestartPolicyOnFailure, core.RestartPolicyNever:
  break
 case "":
  allErrors = append(allErrors, field.Required(fldPath, ""))
 default:
  validValues := []string{string(core.RestartPolicyAlways), string(core.RestartPolicyOnFailure), string(core.RestartPolicyNever)}
  allErrors = append(allErrors, field.NotSupported(fldPath, *restartPolicy, validValues))
 }
 return allErrors
}
func validateDNSPolicy(dnsPolicy *core.DNSPolicy, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 switch *dnsPolicy {
 case core.DNSClusterFirstWithHostNet, core.DNSClusterFirst, core.DNSDefault:
 case core.DNSNone:
  if !utilfeature.DefaultFeatureGate.Enabled(features.CustomPodDNS) {
   allErrors = append(allErrors, field.Invalid(fldPath, dnsPolicy, "DNSPolicy: can not use 'None', custom pod DNS is disabled by feature gate"))
  }
 case "":
  allErrors = append(allErrors, field.Required(fldPath, ""))
 default:
  validValues := []string{string(core.DNSClusterFirstWithHostNet), string(core.DNSClusterFirst), string(core.DNSDefault)}
  if utilfeature.DefaultFeatureGate.Enabled(features.CustomPodDNS) {
   validValues = append(validValues, string(core.DNSNone))
  }
  allErrors = append(allErrors, field.NotSupported(fldPath, dnsPolicy, validValues))
 }
 return allErrors
}

const (
 MaxDNSNameservers     = 3
 MaxDNSSearchPaths     = 6
 MaxDNSSearchListChars = 256
)

func validateReadinessGates(readinessGates []core.PodReadinessGate, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if !utilfeature.DefaultFeatureGate.Enabled(features.PodReadinessGates) && len(readinessGates) > 0 {
  return append(allErrs, field.Forbidden(fldPath, "PodReadinessGates is disabled by feature gate"))
 }
 for i, value := range readinessGates {
  for _, msg := range validation.IsQualifiedName(string(value.ConditionType)) {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("conditionType"), string(value.ConditionType), msg))
  }
 }
 return allErrs
}
func validatePodDNSConfig(dnsConfig *core.PodDNSConfig, dnsPolicy *core.DNSPolicy, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if utilfeature.DefaultFeatureGate.Enabled(features.CustomPodDNS) && dnsPolicy != nil && *dnsPolicy == core.DNSNone {
  if dnsConfig == nil {
   return append(allErrs, field.Required(fldPath, fmt.Sprintf("must provide `dnsConfig` when `dnsPolicy` is %s", core.DNSNone)))
  }
  if len(dnsConfig.Nameservers) == 0 {
   return append(allErrs, field.Required(fldPath.Child("nameservers"), fmt.Sprintf("must provide at least one DNS nameserver when `dnsPolicy` is %s", core.DNSNone)))
  }
 }
 if dnsConfig != nil {
  if !utilfeature.DefaultFeatureGate.Enabled(features.CustomPodDNS) {
   return append(allErrs, field.Forbidden(fldPath, "DNSConfig: custom pod DNS is disabled by feature gate"))
  }
  if len(dnsConfig.Nameservers) > MaxDNSNameservers {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("nameservers"), dnsConfig.Nameservers, fmt.Sprintf("must not have more than %v nameservers", MaxDNSNameservers)))
  }
  for i, ns := range dnsConfig.Nameservers {
   if ip := net.ParseIP(ns); ip == nil {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("nameservers").Index(i), ns, "must be valid IP address"))
   }
  }
  if len(dnsConfig.Searches) > MaxDNSSearchPaths {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("searches"), dnsConfig.Searches, fmt.Sprintf("must not have more than %v search paths", MaxDNSSearchPaths)))
  }
  if len(strings.Join(dnsConfig.Searches, " ")) > MaxDNSSearchListChars {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("searches"), dnsConfig.Searches, "must not have more than 256 characters (including spaces) in the search list"))
  }
  for i, search := range dnsConfig.Searches {
   allErrs = append(allErrs, ValidateDNS1123Subdomain(search, fldPath.Child("searches").Index(i))...)
  }
  for i, option := range dnsConfig.Options {
   if len(option.Name) == 0 {
    allErrs = append(allErrs, field.Required(fldPath.Child("options").Index(i), "must not be empty"))
   }
  }
 }
 return allErrs
}
func validateHostNetwork(hostNetwork bool, containers []core.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 if hostNetwork {
  for i, container := range containers {
   portsPath := fldPath.Index(i).Child("ports")
   for i, port := range container.Ports {
    idxPath := portsPath.Index(i)
    if port.HostPort != port.ContainerPort {
     allErrors = append(allErrors, field.Invalid(idxPath.Child("containerPort"), port.ContainerPort, "must match `hostPort` when `hostNetwork` is true"))
    }
   }
  }
 }
 return allErrors
}
func validateImagePullSecrets(imagePullSecrets []core.LocalObjectReference, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 for i, currPullSecret := range imagePullSecrets {
  idxPath := fldPath.Index(i)
  strippedRef := core.LocalObjectReference{Name: currPullSecret.Name}
  if !reflect.DeepEqual(strippedRef, currPullSecret) {
   allErrors = append(allErrors, field.Invalid(idxPath, currPullSecret, "only name may be set"))
  }
 }
 return allErrors
}
func validateAffinity(affinity *core.Affinity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if affinity != nil {
  if affinity.NodeAffinity != nil {
   allErrs = append(allErrs, validateNodeAffinity(affinity.NodeAffinity, fldPath.Child("nodeAffinity"))...)
  }
  if affinity.PodAffinity != nil {
   allErrs = append(allErrs, validatePodAffinity(affinity.PodAffinity, fldPath.Child("podAffinity"))...)
  }
  if affinity.PodAntiAffinity != nil {
   allErrs = append(allErrs, validatePodAntiAffinity(affinity.PodAntiAffinity, fldPath.Child("podAntiAffinity"))...)
  }
 }
 return allErrs
}
func validateTaintEffect(effect *core.TaintEffect, allowEmpty bool, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !allowEmpty && len(*effect) == 0 {
  return field.ErrorList{field.Required(fldPath, "")}
 }
 allErrors := field.ErrorList{}
 switch *effect {
 case core.TaintEffectNoSchedule, core.TaintEffectPreferNoSchedule, core.TaintEffectNoExecute:
 default:
  validValues := []string{string(core.TaintEffectNoSchedule), string(core.TaintEffectPreferNoSchedule), string(core.TaintEffectNoExecute)}
  allErrors = append(allErrors, field.NotSupported(fldPath, *effect, validValues))
 }
 return allErrors
}
func validateOnlyAddedTolerations(newTolerations []core.Toleration, oldTolerations []core.Toleration, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, old := range oldTolerations {
  found := false
  old.TolerationSeconds = nil
  for _, new := range newTolerations {
   new.TolerationSeconds = nil
   if reflect.DeepEqual(old, new) {
    found = true
    break
   }
  }
  if !found {
   allErrs = append(allErrs, field.Forbidden(fldPath, "existing toleration can not be modified except its tolerationSeconds"))
   return allErrs
  }
 }
 allErrs = append(allErrs, ValidateTolerations(newTolerations, fldPath)...)
 return allErrs
}
func ValidateHostAliases(hostAliases []core.HostAlias, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, hostAlias := range hostAliases {
  if ip := net.ParseIP(hostAlias.IP); ip == nil {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("ip"), hostAlias.IP, "must be valid IP address"))
  }
  for _, hostname := range hostAlias.Hostnames {
   allErrs = append(allErrs, ValidateDNS1123Subdomain(hostname, fldPath.Child("hostnames"))...)
  }
 }
 return allErrs
}
func ValidateTolerations(tolerations []core.Toleration, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 for i, toleration := range tolerations {
  idxPath := fldPath.Index(i)
  if len(toleration.Key) > 0 {
   allErrors = append(allErrors, unversionedvalidation.ValidateLabelName(toleration.Key, idxPath.Child("key"))...)
  }
  if len(toleration.Key) == 0 && toleration.Operator != core.TolerationOpExists {
   allErrors = append(allErrors, field.Invalid(idxPath.Child("operator"), toleration.Operator, "operator must be Exists when `key` is empty, which means \"match all values and all keys\""))
  }
  if toleration.TolerationSeconds != nil && toleration.Effect != core.TaintEffectNoExecute {
   allErrors = append(allErrors, field.Invalid(idxPath.Child("effect"), toleration.Effect, "effect must be 'NoExecute' when `tolerationSeconds` is set"))
  }
  switch toleration.Operator {
  case core.TolerationOpEqual, "":
   if errs := validation.IsValidLabelValue(toleration.Value); len(errs) != 0 {
    allErrors = append(allErrors, field.Invalid(idxPath.Child("operator"), toleration.Value, strings.Join(errs, ";")))
   }
  case core.TolerationOpExists:
   if len(toleration.Value) > 0 {
    allErrors = append(allErrors, field.Invalid(idxPath.Child("operator"), toleration, "value must be empty when `operator` is 'Exists'"))
   }
  default:
   validValues := []string{string(core.TolerationOpEqual), string(core.TolerationOpExists)}
   allErrors = append(allErrors, field.NotSupported(idxPath.Child("operator"), toleration.Operator, validValues))
  }
  if len(toleration.Effect) > 0 {
   allErrors = append(allErrors, validateTaintEffect(&toleration.Effect, true, idxPath.Child("effect"))...)
  }
 }
 return allErrors
}
func toResourceNames(resources core.ResourceList) []core.ResourceName {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := []core.ResourceName{}
 for resourceName := range resources {
  result = append(result, resourceName)
 }
 return result
}
func toSet(resourceNames []core.ResourceName) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := sets.NewString()
 for _, resourceName := range resourceNames {
  result.Insert(string(resourceName))
 }
 return result
}
func toContainerResourcesSet(ctr *core.Container) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 resourceNames := toResourceNames(ctr.Resources.Requests)
 resourceNames = append(resourceNames, toResourceNames(ctr.Resources.Limits)...)
 return toSet(resourceNames)
}
func validateContainersOnlyForPod(containers []core.Container, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, ctr := range containers {
  idxPath := fldPath.Index(i)
  if len(ctr.Image) != len(strings.TrimSpace(ctr.Image)) {
   allErrs = append(allErrs, field.Invalid(idxPath.Child("image"), ctr.Image, "must not have leading or trailing whitespace"))
  }
 }
 return allErrs
}
func ValidatePod(pod *core.Pod) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fldPath := field.NewPath("metadata")
 allErrs := ValidateObjectMeta(&pod.ObjectMeta, true, ValidatePodName, fldPath)
 allErrs = append(allErrs, ValidatePodSpecificAnnotations(pod.ObjectMeta.Annotations, &pod.Spec, fldPath.Child("annotations"))...)
 allErrs = append(allErrs, ValidatePodSpec(&pod.Spec, field.NewPath("spec"))...)
 specPath := field.NewPath("spec")
 if pod.Spec.ServiceAccountName == "" {
  for vi, volume := range pod.Spec.Volumes {
   path := specPath.Child("volumes").Index(vi).Child("projected")
   if volume.Projected != nil {
    for si, source := range volume.Projected.Sources {
     saPath := path.Child("sources").Index(si).Child("serviceAccountToken")
     if source.ServiceAccountToken != nil {
      allErrs = append(allErrs, field.Forbidden(saPath, "must not be specified when serviceAccountName is not set"))
     }
    }
   }
  }
 }
 allErrs = append(allErrs, validateContainersOnlyForPod(pod.Spec.Containers, specPath.Child("containers"))...)
 allErrs = append(allErrs, validateContainersOnlyForPod(pod.Spec.InitContainers, specPath.Child("initContainers"))...)
 if utilfeature.DefaultFeatureGate.Enabled(features.HugePages) {
  hugePageResources := sets.NewString()
  for i := range pod.Spec.Containers {
   resourceSet := toContainerResourcesSet(&pod.Spec.Containers[i])
   for resourceStr := range resourceSet {
    if v1helper.IsHugePageResourceName(v1.ResourceName(resourceStr)) {
     hugePageResources.Insert(resourceStr)
    }
   }
  }
  if len(hugePageResources) > 1 {
   allErrs = append(allErrs, field.Invalid(specPath, hugePageResources, "must use a single hugepage size in a pod spec"))
  }
 }
 return allErrs
}
func ValidatePodSpec(spec *core.PodSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 vols, vErrs := ValidateVolumes(spec.Volumes, fldPath.Child("volumes"))
 allErrs = append(allErrs, vErrs...)
 allErrs = append(allErrs, validateContainers(spec.Containers, false, vols, fldPath.Child("containers"))...)
 allErrs = append(allErrs, validateInitContainers(spec.InitContainers, spec.Containers, vols, fldPath.Child("initContainers"))...)
 allErrs = append(allErrs, validateRestartPolicy(&spec.RestartPolicy, fldPath.Child("restartPolicy"))...)
 allErrs = append(allErrs, validateDNSPolicy(&spec.DNSPolicy, fldPath.Child("dnsPolicy"))...)
 allErrs = append(allErrs, unversionedvalidation.ValidateLabels(spec.NodeSelector, fldPath.Child("nodeSelector"))...)
 allErrs = append(allErrs, ValidatePodSecurityContext(spec.SecurityContext, spec, fldPath, fldPath.Child("securityContext"))...)
 allErrs = append(allErrs, validateImagePullSecrets(spec.ImagePullSecrets, fldPath.Child("imagePullSecrets"))...)
 allErrs = append(allErrs, validateAffinity(spec.Affinity, fldPath.Child("affinity"))...)
 allErrs = append(allErrs, validatePodDNSConfig(spec.DNSConfig, &spec.DNSPolicy, fldPath.Child("dnsConfig"))...)
 allErrs = append(allErrs, validateReadinessGates(spec.ReadinessGates, fldPath.Child("readinessGates"))...)
 if len(spec.ServiceAccountName) > 0 {
  for _, msg := range ValidateServiceAccountName(spec.ServiceAccountName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("serviceAccountName"), spec.ServiceAccountName, msg))
  }
 }
 if len(spec.NodeName) > 0 {
  for _, msg := range ValidateNodeName(spec.NodeName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("nodeName"), spec.NodeName, msg))
  }
 }
 if spec.ActiveDeadlineSeconds != nil {
  value := *spec.ActiveDeadlineSeconds
  if value < 1 || value > math.MaxInt32 {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("activeDeadlineSeconds"), value, validation.InclusiveRangeError(1, math.MaxInt32)))
  }
 }
 if len(spec.Hostname) > 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(spec.Hostname, fldPath.Child("hostname"))...)
 }
 if len(spec.Subdomain) > 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(spec.Subdomain, fldPath.Child("subdomain"))...)
 }
 if len(spec.Tolerations) > 0 {
  allErrs = append(allErrs, ValidateTolerations(spec.Tolerations, fldPath.Child("tolerations"))...)
 }
 if len(spec.HostAliases) > 0 {
  allErrs = append(allErrs, ValidateHostAliases(spec.HostAliases, fldPath.Child("hostAliases"))...)
 }
 if len(spec.PriorityClassName) > 0 {
  if utilfeature.DefaultFeatureGate.Enabled(features.PodPriority) {
   for _, msg := range ValidatePriorityClassName(spec.PriorityClassName, false) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("priorityClassName"), spec.PriorityClassName, msg))
   }
  }
 }
 if spec.RuntimeClassName != nil && utilfeature.DefaultFeatureGate.Enabled(features.RuntimeClass) {
  allErrs = append(allErrs, ValidateRuntimeClassName(*spec.RuntimeClassName, fldPath.Child("runtimeClassName"))...)
 }
 return allErrs
}
func ValidateNodeSelectorRequirement(rq core.NodeSelectorRequirement, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 switch rq.Operator {
 case core.NodeSelectorOpIn, core.NodeSelectorOpNotIn:
  if len(rq.Values) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Child("values"), "must be specified when `operator` is 'In' or 'NotIn'"))
  }
 case core.NodeSelectorOpExists, core.NodeSelectorOpDoesNotExist:
  if len(rq.Values) > 0 {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("values"), "may not be specified when `operator` is 'Exists' or 'DoesNotExist'"))
  }
 case core.NodeSelectorOpGt, core.NodeSelectorOpLt:
  if len(rq.Values) != 1 {
   allErrs = append(allErrs, field.Required(fldPath.Child("values"), "must be specified single value when `operator` is 'Lt' or 'Gt'"))
  }
 default:
  allErrs = append(allErrs, field.Invalid(fldPath.Child("operator"), rq.Operator, "not a valid selector operator"))
 }
 allErrs = append(allErrs, unversionedvalidation.ValidateLabelName(rq.Key, fldPath.Child("key"))...)
 return allErrs
}

var nodeFieldSelectorValidators = map[string]func(string, bool) []string{core.ObjectNameField: ValidateNodeName}

func ValidateNodeFieldSelectorRequirement(req core.NodeSelectorRequirement, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 switch req.Operator {
 case core.NodeSelectorOpIn, core.NodeSelectorOpNotIn:
  if len(req.Values) != 1 {
   allErrs = append(allErrs, field.Required(fldPath.Child("values"), "must be only one value when `operator` is 'In' or 'NotIn' for node field selector"))
  }
 default:
  allErrs = append(allErrs, field.Invalid(fldPath.Child("operator"), req.Operator, "not a valid selector operator"))
 }
 if vf, found := nodeFieldSelectorValidators[req.Key]; !found {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("key"), req.Key, "not a valid field selector key"))
 } else {
  for i, v := range req.Values {
   for _, msg := range vf(v, false) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("values").Index(i), v, msg))
   }
  }
 }
 return allErrs
}
func ValidateNodeSelectorTerm(term core.NodeSelectorTerm, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for j, req := range term.MatchExpressions {
  allErrs = append(allErrs, ValidateNodeSelectorRequirement(req, fldPath.Child("matchExpressions").Index(j))...)
 }
 for j, req := range term.MatchFields {
  allErrs = append(allErrs, ValidateNodeFieldSelectorRequirement(req, fldPath.Child("matchFields").Index(j))...)
 }
 return allErrs
}
func ValidateNodeSelector(nodeSelector *core.NodeSelector, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 termFldPath := fldPath.Child("nodeSelectorTerms")
 if len(nodeSelector.NodeSelectorTerms) == 0 {
  return append(allErrs, field.Required(termFldPath, "must have at least one node selector term"))
 }
 for i, term := range nodeSelector.NodeSelectorTerms {
  allErrs = append(allErrs, ValidateNodeSelectorTerm(term, termFldPath.Index(i))...)
 }
 return allErrs
}
func validateTopologySelectorLabelRequirement(rq core.TopologySelectorLabelRequirement, fldPath *field.Path) (sets.String, field.ErrorList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 valueSet := make(sets.String)
 valuesPath := fldPath.Child("values")
 if len(rq.Values) == 0 {
  allErrs = append(allErrs, field.Required(valuesPath, ""))
 }
 for i, value := range rq.Values {
  if valueSet.Has(value) {
   allErrs = append(allErrs, field.Duplicate(valuesPath.Index(i), value))
  }
  valueSet.Insert(value)
 }
 allErrs = append(allErrs, unversionedvalidation.ValidateLabelName(rq.Key, fldPath.Child("key"))...)
 return valueSet, allErrs
}
func ValidateTopologySelectorTerm(term core.TopologySelectorTerm, fldPath *field.Path) (map[string]sets.String, field.ErrorList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 exprMap := make(map[string]sets.String)
 exprPath := fldPath.Child("matchLabelExpressions")
 if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
  for i, req := range term.MatchLabelExpressions {
   idxPath := exprPath.Index(i)
   valueSet, exprErrs := validateTopologySelectorLabelRequirement(req, idxPath)
   allErrs = append(allErrs, exprErrs...)
   if _, exists := exprMap[req.Key]; exists {
    allErrs = append(allErrs, field.Duplicate(idxPath.Child("key"), req.Key))
   }
   exprMap[req.Key] = valueSet
  }
 } else if len(term.MatchLabelExpressions) != 0 {
  allErrs = append(allErrs, field.Forbidden(fldPath, "field is disabled by feature-gate VolumeScheduling"))
 }
 return exprMap, allErrs
}
func ValidateAvoidPodsInNodeAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 v1Avoids, err := v1helper.GetAvoidPodsFromNodeAnnotations(annotations)
 if err != nil {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("AvoidPods"), core.PreferAvoidPodsAnnotationKey, err.Error()))
  return allErrs
 }
 var avoids core.AvoidPods
 if err := corev1.Convert_v1_AvoidPods_To_core_AvoidPods(&v1Avoids, &avoids, nil); err != nil {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("AvoidPods"), core.PreferAvoidPodsAnnotationKey, err.Error()))
  return allErrs
 }
 if len(avoids.PreferAvoidPods) != 0 {
  for i, pa := range avoids.PreferAvoidPods {
   idxPath := fldPath.Child(core.PreferAvoidPodsAnnotationKey).Index(i)
   allErrs = append(allErrs, validatePreferAvoidPodsEntry(pa, idxPath)...)
  }
 }
 return allErrs
}
func validatePreferAvoidPodsEntry(avoidPodEntry core.PreferAvoidPodsEntry, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 if avoidPodEntry.PodSignature.PodController == nil {
  allErrors = append(allErrors, field.Required(fldPath.Child("PodSignature"), ""))
 } else {
  if *(avoidPodEntry.PodSignature.PodController.Controller) != true {
   allErrors = append(allErrors, field.Invalid(fldPath.Child("PodSignature").Child("PodController").Child("Controller"), *(avoidPodEntry.PodSignature.PodController.Controller), "must point to a controller"))
  }
 }
 return allErrors
}
func ValidatePreferredSchedulingTerms(terms []core.PreferredSchedulingTerm, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, term := range terms {
  if term.Weight <= 0 || term.Weight > 100 {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("weight"), term.Weight, "must be in the range 1-100"))
  }
  allErrs = append(allErrs, ValidateNodeSelectorTerm(term.Preference, fldPath.Index(i).Child("preference"))...)
 }
 return allErrs
}
func validatePodAffinityTerm(podAffinityTerm core.PodAffinityTerm, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, unversionedvalidation.ValidateLabelSelector(podAffinityTerm.LabelSelector, fldPath.Child("matchExpressions"))...)
 for _, name := range podAffinityTerm.Namespaces {
  for _, msg := range ValidateNamespaceName(name, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("namespace"), name, msg))
  }
 }
 if len(podAffinityTerm.TopologyKey) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("topologyKey"), "can not be empty"))
 }
 return append(allErrs, unversionedvalidation.ValidateLabelName(podAffinityTerm.TopologyKey, fldPath.Child("topologyKey"))...)
}
func validatePodAffinityTerms(podAffinityTerms []core.PodAffinityTerm, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, podAffinityTerm := range podAffinityTerms {
  allErrs = append(allErrs, validatePodAffinityTerm(podAffinityTerm, fldPath.Index(i))...)
 }
 return allErrs
}
func validateWeightedPodAffinityTerms(weightedPodAffinityTerms []core.WeightedPodAffinityTerm, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for j, weightedTerm := range weightedPodAffinityTerms {
  if weightedTerm.Weight <= 0 || weightedTerm.Weight > 100 {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(j).Child("weight"), weightedTerm.Weight, "must be in the range 1-100"))
  }
  allErrs = append(allErrs, validatePodAffinityTerm(weightedTerm.PodAffinityTerm, fldPath.Index(j).Child("podAffinityTerm"))...)
 }
 return allErrs
}
func validatePodAntiAffinity(podAntiAffinity *core.PodAntiAffinity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  allErrs = append(allErrs, validatePodAffinityTerms(podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution, fldPath.Child("requiredDuringSchedulingIgnoredDuringExecution"))...)
 }
 if podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
  allErrs = append(allErrs, validateWeightedPodAffinityTerms(podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution, fldPath.Child("preferredDuringSchedulingIgnoredDuringExecution"))...)
 }
 return allErrs
}
func validateNodeAffinity(na *core.NodeAffinity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if na.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  allErrs = append(allErrs, ValidateNodeSelector(na.RequiredDuringSchedulingIgnoredDuringExecution, fldPath.Child("requiredDuringSchedulingIgnoredDuringExecution"))...)
 }
 if len(na.PreferredDuringSchedulingIgnoredDuringExecution) > 0 {
  allErrs = append(allErrs, ValidatePreferredSchedulingTerms(na.PreferredDuringSchedulingIgnoredDuringExecution, fldPath.Child("preferredDuringSchedulingIgnoredDuringExecution"))...)
 }
 return allErrs
}
func validatePodAffinity(podAffinity *core.PodAffinity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if podAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
  allErrs = append(allErrs, validatePodAffinityTerms(podAffinity.RequiredDuringSchedulingIgnoredDuringExecution, fldPath.Child("requiredDuringSchedulingIgnoredDuringExecution"))...)
 }
 if podAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
  allErrs = append(allErrs, validateWeightedPodAffinityTerms(podAffinity.PreferredDuringSchedulingIgnoredDuringExecution, fldPath.Child("preferredDuringSchedulingIgnoredDuringExecution"))...)
 }
 return allErrs
}
func ValidateSeccompProfile(p string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if p == core.SeccompProfileRuntimeDefault || p == core.DeprecatedSeccompProfileDockerDefault {
  return nil
 }
 if p == "unconfined" {
  return nil
 }
 if strings.HasPrefix(p, "localhost/") {
  return validateLocalDescendingPath(strings.TrimPrefix(p, "localhost/"), fldPath)
 }
 return field.ErrorList{field.Invalid(fldPath, p, "must be a valid seccomp profile")}
}
func ValidateSeccompPodAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if p, exists := annotations[core.SeccompPodAnnotationKey]; exists {
  allErrs = append(allErrs, ValidateSeccompProfile(p, fldPath.Child(core.SeccompPodAnnotationKey))...)
 }
 for k, p := range annotations {
  if strings.HasPrefix(k, core.SeccompContainerAnnotationKeyPrefix) {
   allErrs = append(allErrs, ValidateSeccompProfile(p, fldPath.Child(k))...)
  }
 }
 return allErrs
}
func ValidateAppArmorPodAnnotations(annotations map[string]string, spec *core.PodSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for k, p := range annotations {
  if !strings.HasPrefix(k, apparmor.ContainerAnnotationKeyPrefix) {
   continue
  }
  if !utilfeature.DefaultFeatureGate.Enabled(features.AppArmor) {
   allErrs = append(allErrs, field.Forbidden(fldPath.Key(k), "AppArmor is disabled by feature-gate"))
   continue
  }
  containerName := strings.TrimPrefix(k, apparmor.ContainerAnnotationKeyPrefix)
  if !podSpecHasContainer(spec, containerName) {
   allErrs = append(allErrs, field.Invalid(fldPath.Key(k), containerName, "container not found"))
  }
  if err := apparmor.ValidateProfileFormat(p); err != nil {
   allErrs = append(allErrs, field.Invalid(fldPath.Key(k), p, err.Error()))
  }
 }
 return allErrs
}
func podSpecHasContainer(spec *core.PodSpec, containerName string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range spec.InitContainers {
  if c.Name == containerName {
   return true
  }
 }
 for _, c := range spec.Containers {
  if c.Name == containerName {
   return true
  }
 }
 return false
}

const (
 SysctlSegmentFmt string = "[a-z0-9]([-_a-z0-9]*[a-z0-9])?"
 SysctlFmt        string = "(" + SysctlSegmentFmt + "\\.)*" + SysctlSegmentFmt
 SysctlMaxLength  int    = 253
)

var sysctlRegexp = regexp.MustCompile("^" + SysctlFmt + "$")

func IsValidSysctlName(name string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(name) > SysctlMaxLength {
  return false
 }
 return sysctlRegexp.MatchString(name)
}
func validateSysctls(sysctls []core.Sysctl, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 names := make(map[string]struct{})
 for i, s := range sysctls {
  if len(s.Name) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Index(i).Child("name"), ""))
  } else if !IsValidSysctlName(s.Name) {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("name"), s.Name, fmt.Sprintf("must have at most %d characters and match regex %s", SysctlMaxLength, SysctlFmt)))
  } else if _, ok := names[s.Name]; ok {
   allErrs = append(allErrs, field.Duplicate(fldPath.Index(i).Child("name"), s.Name))
  }
  names[s.Name] = struct{}{}
 }
 return allErrs
}
func ValidatePodSecurityContext(securityContext *core.PodSecurityContext, spec *core.PodSpec, specPath, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if securityContext != nil {
  allErrs = append(allErrs, validateHostNetwork(securityContext.HostNetwork, spec.Containers, specPath.Child("containers"))...)
  if securityContext.FSGroup != nil {
   for _, msg := range validation.IsValidGroupID(*securityContext.FSGroup) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("fsGroup"), *(securityContext.FSGroup), msg))
   }
  }
  if securityContext.RunAsUser != nil {
   for _, msg := range validation.IsValidUserID(*securityContext.RunAsUser) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsUser"), *(securityContext.RunAsUser), msg))
   }
  }
  if securityContext.RunAsGroup != nil {
   for _, msg := range validation.IsValidGroupID(*securityContext.RunAsGroup) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsGroup"), *(securityContext.RunAsGroup), msg))
   }
  }
  for g, gid := range securityContext.SupplementalGroups {
   for _, msg := range validation.IsValidGroupID(gid) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("supplementalGroups").Index(g), gid, msg))
   }
  }
  if securityContext.ShareProcessNamespace != nil {
   if !utilfeature.DefaultFeatureGate.Enabled(features.PodShareProcessNamespace) {
    allErrs = append(allErrs, field.Forbidden(fldPath.Child("shareProcessNamespace"), "Process Namespace Sharing is disabled by PodShareProcessNamespace feature-gate"))
   } else if securityContext.HostPID && *securityContext.ShareProcessNamespace {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("shareProcessNamespace"), *securityContext.ShareProcessNamespace, "ShareProcessNamespace and HostPID cannot both be enabled"))
   }
  }
  if len(securityContext.Sysctls) != 0 {
   if utilfeature.DefaultFeatureGate.Enabled(features.Sysctls) {
    allErrs = append(allErrs, validateSysctls(securityContext.Sysctls, fldPath.Child("sysctls"))...)
   } else {
    allErrs = append(allErrs, field.Forbidden(fldPath.Child("sysctls"), "Sysctls are disabled by Sysctls feature-gate"))
   }
  }
 }
 return allErrs
}
func ValidateContainerUpdates(newContainers, oldContainers []core.Container, fldPath *field.Path) (allErrs field.ErrorList, stop bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs = field.ErrorList{}
 if len(newContainers) != len(oldContainers) {
  allErrs = append(allErrs, field.Forbidden(fldPath, "pod updates may not add or remove containers"))
  return allErrs, true
 }
 for i, ctr := range newContainers {
  if len(ctr.Image) == 0 {
   allErrs = append(allErrs, field.Required(fldPath.Index(i).Child("image"), ""))
  }
  if len(strings.TrimSpace(ctr.Image)) != len(ctr.Image) {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("image"), ctr.Image, "must not have leading or trailing whitespace"))
  }
 }
 return allErrs, false
}
func ValidatePodUpdate(newPod, oldPod *core.Pod) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fldPath := field.NewPath("metadata")
 allErrs := ValidateObjectMetaUpdate(&newPod.ObjectMeta, &oldPod.ObjectMeta, fldPath)
 allErrs = append(allErrs, ValidatePodSpecificAnnotationUpdates(newPod, oldPod, fldPath.Child("annotations"))...)
 specPath := field.NewPath("spec")
 containerErrs, stop := ValidateContainerUpdates(newPod.Spec.Containers, oldPod.Spec.Containers, specPath.Child("containers"))
 allErrs = append(allErrs, containerErrs...)
 if stop {
  return allErrs
 }
 containerErrs, stop = ValidateContainerUpdates(newPod.Spec.InitContainers, oldPod.Spec.InitContainers, specPath.Child("initContainers"))
 allErrs = append(allErrs, containerErrs...)
 if stop {
  return allErrs
 }
 if newPod.Spec.ActiveDeadlineSeconds != nil {
  newActiveDeadlineSeconds := *newPod.Spec.ActiveDeadlineSeconds
  if newActiveDeadlineSeconds < 0 || newActiveDeadlineSeconds > math.MaxInt32 {
   allErrs = append(allErrs, field.Invalid(specPath.Child("activeDeadlineSeconds"), newActiveDeadlineSeconds, validation.InclusiveRangeError(0, math.MaxInt32)))
   return allErrs
  }
  if oldPod.Spec.ActiveDeadlineSeconds != nil {
   oldActiveDeadlineSeconds := *oldPod.Spec.ActiveDeadlineSeconds
   if oldActiveDeadlineSeconds < newActiveDeadlineSeconds {
    allErrs = append(allErrs, field.Invalid(specPath.Child("activeDeadlineSeconds"), newActiveDeadlineSeconds, "must be less than or equal to previous value"))
    return allErrs
   }
  }
 } else if oldPod.Spec.ActiveDeadlineSeconds != nil {
  allErrs = append(allErrs, field.Invalid(specPath.Child("activeDeadlineSeconds"), newPod.Spec.ActiveDeadlineSeconds, "must not update from a positive integer to nil value"))
 }
 mungedPod := *newPod
 var newContainers []core.Container
 for ix, container := range mungedPod.Spec.Containers {
  container.Image = oldPod.Spec.Containers[ix].Image
  newContainers = append(newContainers, container)
 }
 mungedPod.Spec.Containers = newContainers
 var newInitContainers []core.Container
 for ix, container := range mungedPod.Spec.InitContainers {
  container.Image = oldPod.Spec.InitContainers[ix].Image
  newInitContainers = append(newInitContainers, container)
 }
 mungedPod.Spec.InitContainers = newInitContainers
 mungedPod.Spec.ActiveDeadlineSeconds = nil
 if oldPod.Spec.ActiveDeadlineSeconds != nil {
  activeDeadlineSeconds := *oldPod.Spec.ActiveDeadlineSeconds
  mungedPod.Spec.ActiveDeadlineSeconds = &activeDeadlineSeconds
 }
 mungedPod.Spec.Tolerations = oldPod.Spec.Tolerations
 allErrs = append(allErrs, validateOnlyAddedTolerations(newPod.Spec.Tolerations, oldPod.Spec.Tolerations, specPath.Child("tolerations"))...)
 if !apiequality.Semantic.DeepEqual(mungedPod.Spec, oldPod.Spec) {
  specDiff := diff.ObjectDiff(mungedPod.Spec, oldPod.Spec)
  allErrs = append(allErrs, field.Forbidden(specPath, fmt.Sprintf("pod updates may not change fields other than `spec.containers[*].image`, `spec.initContainers[*].image`, `spec.activeDeadlineSeconds` or `spec.tolerations` (only additions to existing tolerations)\n%v", specDiff)))
 }
 return allErrs
}
func ValidateContainerStateTransition(newStatuses, oldStatuses []core.ContainerStatus, fldpath *field.Path, restartPolicy core.RestartPolicy) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if restartPolicy == core.RestartPolicyAlways {
  return allErrs
 }
 for i, oldStatus := range oldStatuses {
  if oldStatus.State.Terminated == nil {
   continue
  }
  if oldStatus.State.Terminated.ExitCode != 0 && restartPolicy == core.RestartPolicyOnFailure {
   continue
  }
  for _, newStatus := range newStatuses {
   if oldStatus.Name == newStatus.Name && newStatus.State.Terminated == nil {
    allErrs = append(allErrs, field.Forbidden(fldpath.Index(i).Child("state"), "may not be transitioned to non-terminated state"))
   }
  }
 }
 return allErrs
}
func ValidatePodStatusUpdate(newPod, oldPod *core.Pod) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fldPath := field.NewPath("metadata")
 allErrs := ValidateObjectMetaUpdate(&newPod.ObjectMeta, &oldPod.ObjectMeta, fldPath)
 allErrs = append(allErrs, ValidatePodSpecificAnnotationUpdates(newPod, oldPod, fldPath.Child("annotations"))...)
 allErrs = append(allErrs, validatePodConditions(newPod.Status.Conditions, fldPath.Child("conditions"))...)
 fldPath = field.NewPath("status")
 if newPod.Spec.NodeName != oldPod.Spec.NodeName {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("nodeName"), "may not be changed directly"))
 }
 if newPod.Status.NominatedNodeName != oldPod.Status.NominatedNodeName && len(newPod.Status.NominatedNodeName) > 0 {
  for _, msg := range ValidateNodeName(newPod.Status.NominatedNodeName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("nominatedNodeName"), newPod.Status.NominatedNodeName, msg))
  }
 }
 allErrs = append(allErrs, ValidateContainerStateTransition(newPod.Status.ContainerStatuses, oldPod.Status.ContainerStatuses, fldPath.Child("containerStatuses"), oldPod.Spec.RestartPolicy)...)
 allErrs = append(allErrs, ValidateContainerStateTransition(newPod.Status.InitContainerStatuses, oldPod.Status.InitContainerStatuses, fldPath.Child("initContainerStatuses"), oldPod.Spec.RestartPolicy)...)
 newPod.Spec = oldPod.Spec
 return allErrs
}
func validatePodConditions(conditions []core.PodCondition, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 systemConditions := sets.NewString(string(core.PodScheduled), string(core.PodReady), string(core.PodInitialized))
 for i, condition := range conditions {
  if systemConditions.Has(string(condition.Type)) {
   continue
  }
  for _, msg := range validation.IsQualifiedName(string(condition.Type)) {
   allErrs = append(allErrs, field.Invalid(fldPath.Index(i).Child("Type"), string(condition.Type), msg))
  }
 }
 return allErrs
}
func ValidatePodBinding(binding *core.Binding) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(binding.Target.Kind) != 0 && binding.Target.Kind != "Node" {
  allErrs = append(allErrs, field.NotSupported(field.NewPath("target", "kind"), binding.Target.Kind, []string{"Node", "<empty>"}))
 }
 if len(binding.Target.Name) == 0 {
  allErrs = append(allErrs, field.Required(field.NewPath("target", "name"), ""))
 }
 return allErrs
}
func ValidatePodTemplate(pod *core.PodTemplate) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&pod.ObjectMeta, true, ValidatePodName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidatePodTemplateSpec(&pod.Template, field.NewPath("template"))...)
 return allErrs
}
func ValidatePodTemplateUpdate(newPod, oldPod *core.PodTemplate) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newPod.ObjectMeta, &oldPod.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidatePodTemplateSpec(&newPod.Template, field.NewPath("template"))...)
 return allErrs
}

var supportedSessionAffinityType = sets.NewString(string(core.ServiceAffinityClientIP), string(core.ServiceAffinityNone))
var supportedServiceType = sets.NewString(string(core.ServiceTypeClusterIP), string(core.ServiceTypeNodePort), string(core.ServiceTypeLoadBalancer), string(core.ServiceTypeExternalName))

func ValidateService(service *core.Service) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&service.ObjectMeta, true, ValidateServiceName, field.NewPath("metadata"))
 specPath := field.NewPath("spec")
 isHeadlessService := service.Spec.ClusterIP == core.ClusterIPNone
 if len(service.Spec.Ports) == 0 && !isHeadlessService && service.Spec.Type != core.ServiceTypeExternalName {
  allErrs = append(allErrs, field.Required(specPath.Child("ports"), ""))
 }
 switch service.Spec.Type {
 case core.ServiceTypeLoadBalancer:
  for ix := range service.Spec.Ports {
   port := &service.Spec.Ports[ix]
   if port.Port == ports.KubeletPort {
    portPath := specPath.Child("ports").Index(ix)
    allErrs = append(allErrs, field.Invalid(portPath, port.Port, fmt.Sprintf("may not expose port %v externally since it is used by kubelet", ports.KubeletPort)))
   }
  }
  if service.Spec.ClusterIP == "None" {
   allErrs = append(allErrs, field.Invalid(specPath.Child("clusterIP"), service.Spec.ClusterIP, "may not be set to 'None' for LoadBalancer services"))
  }
 case core.ServiceTypeNodePort:
  if service.Spec.ClusterIP == "None" {
   allErrs = append(allErrs, field.Invalid(specPath.Child("clusterIP"), service.Spec.ClusterIP, "may not be set to 'None' for NodePort services"))
  }
 case core.ServiceTypeExternalName:
  if service.Spec.ClusterIP != "" {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("clusterIP"), "must be empty for ExternalName services"))
  }
  if len(service.Spec.ExternalName) > 0 {
   allErrs = append(allErrs, ValidateDNS1123Subdomain(service.Spec.ExternalName, specPath.Child("externalName"))...)
  } else {
   allErrs = append(allErrs, field.Required(specPath.Child("externalName"), ""))
  }
 }
 allPortNames := sets.String{}
 portsPath := specPath.Child("ports")
 for i := range service.Spec.Ports {
  portPath := portsPath.Index(i)
  allErrs = append(allErrs, validateServicePort(&service.Spec.Ports[i], len(service.Spec.Ports) > 1, isHeadlessService, &allPortNames, portPath)...)
 }
 if service.Spec.Selector != nil {
  allErrs = append(allErrs, unversionedvalidation.ValidateLabels(service.Spec.Selector, specPath.Child("selector"))...)
 }
 if len(service.Spec.SessionAffinity) == 0 {
  allErrs = append(allErrs, field.Required(specPath.Child("sessionAffinity"), ""))
 } else if !supportedSessionAffinityType.Has(string(service.Spec.SessionAffinity)) {
  allErrs = append(allErrs, field.NotSupported(specPath.Child("sessionAffinity"), service.Spec.SessionAffinity, supportedSessionAffinityType.List()))
 }
 if service.Spec.SessionAffinity == core.ServiceAffinityClientIP {
  allErrs = append(allErrs, validateClientIPAffinityConfig(service.Spec.SessionAffinityConfig, specPath.Child("sessionAffinityConfig"))...)
 } else if service.Spec.SessionAffinity == core.ServiceAffinityNone {
  if service.Spec.SessionAffinityConfig != nil {
   allErrs = append(allErrs, field.Forbidden(specPath.Child("sessionAffinityConfig"), fmt.Sprintf("must not be set when session affinity is %s", string(core.ServiceAffinityNone))))
  }
 }
 if helper.IsServiceIPSet(service) {
  if ip := net.ParseIP(service.Spec.ClusterIP); ip == nil {
   allErrs = append(allErrs, field.Invalid(specPath.Child("clusterIP"), service.Spec.ClusterIP, "must be empty, 'None', or a valid IP address"))
  }
 }
 ipPath := specPath.Child("externalIPs")
 for i, ip := range service.Spec.ExternalIPs {
  idxPath := ipPath.Index(i)
  if msgs := validation.IsValidIP(ip); len(msgs) != 0 {
   for i := range msgs {
    allErrs = append(allErrs, field.Invalid(idxPath, ip, msgs[i]))
   }
  } else {
   allErrs = append(allErrs, validateNonSpecialIP(ip, idxPath)...)
  }
 }
 if len(service.Spec.Type) == 0 {
  allErrs = append(allErrs, field.Required(specPath.Child("type"), ""))
 } else if !supportedServiceType.Has(string(service.Spec.Type)) {
  allErrs = append(allErrs, field.NotSupported(specPath.Child("type"), service.Spec.Type, supportedServiceType.List()))
 }
 if service.Spec.Type == core.ServiceTypeLoadBalancer {
  portsPath := specPath.Child("ports")
  includeProtocols := sets.NewString()
  for i := range service.Spec.Ports {
   portPath := portsPath.Index(i)
   if !utilfeature.DefaultFeatureGate.Enabled(features.SCTPSupport) && service.Spec.Ports[i].Protocol == core.ProtocolSCTP {
    allErrs = append(allErrs, field.NotSupported(portPath.Child("protocol"), service.Spec.Ports[i].Protocol, []string{string(core.ProtocolTCP), string(core.ProtocolUDP)}))
   } else if !supportedPortProtocols.Has(string(service.Spec.Ports[i].Protocol)) {
    allErrs = append(allErrs, field.Invalid(portPath.Child("protocol"), service.Spec.Ports[i].Protocol, "cannot create an external load balancer with non-TCP/UDP/SCTP ports"))
   } else {
    includeProtocols.Insert(string(service.Spec.Ports[i].Protocol))
   }
  }
  if includeProtocols.Len() > 1 {
   allErrs = append(allErrs, field.Invalid(portsPath, service.Spec.Ports, "cannot create an external load balancer with mix protocols"))
  }
 }
 if service.Spec.Type == core.ServiceTypeClusterIP {
  portsPath := specPath.Child("ports")
  for i := range service.Spec.Ports {
   portPath := portsPath.Index(i)
   if service.Spec.Ports[i].NodePort != 0 {
    allErrs = append(allErrs, field.Forbidden(portPath.Child("nodePort"), "may not be used when `type` is 'ClusterIP'"))
   }
  }
 }
 portsPath = specPath.Child("ports")
 nodePorts := make(map[core.ServicePort]bool)
 for i := range service.Spec.Ports {
  port := &service.Spec.Ports[i]
  if port.NodePort == 0 {
   continue
  }
  portPath := portsPath.Index(i)
  var key core.ServicePort
  key.Protocol = port.Protocol
  key.NodePort = port.NodePort
  _, found := nodePorts[key]
  if found {
   allErrs = append(allErrs, field.Duplicate(portPath.Child("nodePort"), port.NodePort))
  }
  nodePorts[key] = true
 }
 portsPath = specPath.Child("ports")
 ports := make(map[core.ServicePort]bool)
 for i, port := range service.Spec.Ports {
  portPath := portsPath.Index(i)
  key := core.ServicePort{Protocol: port.Protocol, Port: port.Port}
  _, found := ports[key]
  if found {
   allErrs = append(allErrs, field.Duplicate(portPath, key))
  }
  ports[key] = true
 }
 _, ok := service.Annotations[core.AnnotationLoadBalancerSourceRangesKey]
 if len(service.Spec.LoadBalancerSourceRanges) > 0 || ok {
  var fieldPath *field.Path
  var val string
  if len(service.Spec.LoadBalancerSourceRanges) > 0 {
   fieldPath = specPath.Child("LoadBalancerSourceRanges")
   val = fmt.Sprintf("%v", service.Spec.LoadBalancerSourceRanges)
  } else {
   fieldPath = field.NewPath("metadata", "annotations").Key(core.AnnotationLoadBalancerSourceRangesKey)
   val = service.Annotations[core.AnnotationLoadBalancerSourceRangesKey]
  }
  if service.Spec.Type != core.ServiceTypeLoadBalancer {
   allErrs = append(allErrs, field.Forbidden(fieldPath, "may only be used when `type` is 'LoadBalancer'"))
  }
  _, err := apiservice.GetLoadBalancerSourceRanges(service)
  if err != nil {
   allErrs = append(allErrs, field.Invalid(fieldPath, val, "must be a list of IP ranges. For example, 10.240.0.0/24,10.250.0.0/24 "))
  }
 }
 allErrs = append(allErrs, validateServiceExternalTrafficFieldsValue(service)...)
 return allErrs
}
func validateServicePort(sp *core.ServicePort, requireName, isHeadlessService bool, allNames *sets.String, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if requireName && len(sp.Name) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
 } else if len(sp.Name) != 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(sp.Name, fldPath.Child("name"))...)
  if allNames.Has(sp.Name) {
   allErrs = append(allErrs, field.Duplicate(fldPath.Child("name"), sp.Name))
  } else {
   allNames.Insert(sp.Name)
  }
 }
 for _, msg := range validation.IsValidPortNum(int(sp.Port)) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("port"), sp.Port, msg))
 }
 if len(sp.Protocol) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("protocol"), ""))
 } else if !utilfeature.DefaultFeatureGate.Enabled(features.SCTPSupport) && sp.Protocol == core.ProtocolSCTP {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("protocol"), sp.Protocol, []string{string(core.ProtocolTCP), string(core.ProtocolUDP)}))
 } else if !supportedPortProtocols.Has(string(sp.Protocol)) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("protocol"), sp.Protocol, supportedPortProtocols.List()))
 }
 allErrs = append(allErrs, ValidatePortNumOrName(sp.TargetPort, fldPath.Child("targetPort"))...)
 return allErrs
}
func validateServiceExternalTrafficFieldsValue(service *core.Service) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if service.Spec.ExternalTrafficPolicy != "" && service.Spec.ExternalTrafficPolicy != core.ServiceExternalTrafficPolicyTypeCluster && service.Spec.ExternalTrafficPolicy != core.ServiceExternalTrafficPolicyTypeLocal {
  allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("externalTrafficPolicy"), service.Spec.ExternalTrafficPolicy, fmt.Sprintf("ExternalTrafficPolicy must be empty, %v or %v", core.ServiceExternalTrafficPolicyTypeCluster, core.ServiceExternalTrafficPolicyTypeLocal)))
 }
 if service.Spec.HealthCheckNodePort < 0 {
  allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("healthCheckNodePort"), service.Spec.HealthCheckNodePort, "HealthCheckNodePort must be not less than 0"))
 }
 return allErrs
}
func ValidateServiceExternalTrafficFieldsCombination(service *core.Service) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if service.Spec.Type != core.ServiceTypeLoadBalancer && service.Spec.Type != core.ServiceTypeNodePort && service.Spec.ExternalTrafficPolicy != "" {
  allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "externalTrafficPolicy"), service.Spec.ExternalTrafficPolicy, "ExternalTrafficPolicy can only be set on NodePort and LoadBalancer service"))
 }
 if !apiservice.NeedsHealthCheck(service) && service.Spec.HealthCheckNodePort != 0 {
  allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "healthCheckNodePort"), service.Spec.HealthCheckNodePort, "HealthCheckNodePort can only be set on LoadBalancer service with ExternalTrafficPolicy=Local"))
 }
 return allErrs
}
func ValidateServiceUpdate(service, oldService *core.Service) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&service.ObjectMeta, &oldService.ObjectMeta, field.NewPath("metadata"))
 if service.Spec.Type != core.ServiceTypeExternalName {
  if oldService.Spec.Type != core.ServiceTypeExternalName && oldService.Spec.ClusterIP != "" {
   allErrs = append(allErrs, ValidateImmutableField(service.Spec.ClusterIP, oldService.Spec.ClusterIP, field.NewPath("spec", "clusterIP"))...)
  }
 }
 allErrs = append(allErrs, ValidateService(service)...)
 return allErrs
}
func ValidateServiceStatusUpdate(service, oldService *core.Service) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&service.ObjectMeta, &oldService.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateLoadBalancerStatus(&service.Status.LoadBalancer, field.NewPath("status", "loadBalancer"))...)
 return allErrs
}
func ValidateReplicationController(controller *core.ReplicationController) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&controller.ObjectMeta, true, ValidateReplicationControllerName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateReplicationControllerSpec(&controller.Spec, field.NewPath("spec"))...)
 return allErrs
}
func ValidateReplicationControllerUpdate(controller, oldController *core.ReplicationController) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&controller.ObjectMeta, &oldController.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateReplicationControllerSpec(&controller.Spec, field.NewPath("spec"))...)
 return allErrs
}
func ValidateReplicationControllerStatusUpdate(controller, oldController *core.ReplicationController) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&controller.ObjectMeta, &oldController.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateReplicationControllerStatus(controller.Status, field.NewPath("status"))...)
 return allErrs
}
func ValidateReplicationControllerStatus(status core.ReplicationControllerStatus, statusPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateNonnegativeField(int64(status.Replicas), statusPath.Child("replicas"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(status.FullyLabeledReplicas), statusPath.Child("fullyLabeledReplicas"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(status.ReadyReplicas), statusPath.Child("readyReplicas"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(status.AvailableReplicas), statusPath.Child("availableReplicas"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(status.ObservedGeneration), statusPath.Child("observedGeneration"))...)
 msg := "cannot be greater than status.replicas"
 if status.FullyLabeledReplicas > status.Replicas {
  allErrs = append(allErrs, field.Invalid(statusPath.Child("fullyLabeledReplicas"), status.FullyLabeledReplicas, msg))
 }
 if status.ReadyReplicas > status.Replicas {
  allErrs = append(allErrs, field.Invalid(statusPath.Child("readyReplicas"), status.ReadyReplicas, msg))
 }
 if status.AvailableReplicas > status.Replicas {
  allErrs = append(allErrs, field.Invalid(statusPath.Child("availableReplicas"), status.AvailableReplicas, msg))
 }
 if status.AvailableReplicas > status.ReadyReplicas {
  allErrs = append(allErrs, field.Invalid(statusPath.Child("availableReplicas"), status.AvailableReplicas, "cannot be greater than readyReplicas"))
 }
 return allErrs
}
func ValidateNonEmptySelector(selectorMap map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 selector := labels.Set(selectorMap).AsSelector()
 if selector.Empty() {
  allErrs = append(allErrs, field.Required(fldPath, ""))
 }
 return allErrs
}
func ValidatePodTemplateSpecForRC(template *core.PodTemplateSpec, selectorMap map[string]string, replicas int32, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if template == nil {
  allErrs = append(allErrs, field.Required(fldPath, ""))
 } else {
  selector := labels.Set(selectorMap).AsSelector()
  if !selector.Empty() {
   labels := labels.Set(template.Labels)
   if !selector.Matches(labels) {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("metadata", "labels"), template.Labels, "`selector` does not match template `labels`"))
   }
  }
  allErrs = append(allErrs, ValidatePodTemplateSpec(template, fldPath)...)
  if replicas > 1 {
   allErrs = append(allErrs, ValidateReadOnlyPersistentDisks(template.Spec.Volumes, fldPath.Child("spec", "volumes"))...)
  }
  if template.Spec.RestartPolicy != core.RestartPolicyAlways {
   allErrs = append(allErrs, field.NotSupported(fldPath.Child("spec", "restartPolicy"), template.Spec.RestartPolicy, []string{string(core.RestartPolicyAlways)}))
  }
  if template.Spec.ActiveDeadlineSeconds != nil {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("spec", "activeDeadlineSeconds"), template.Spec.ActiveDeadlineSeconds, "must not be specified"))
  }
 }
 return allErrs
}
func ValidateReplicationControllerSpec(spec *core.ReplicationControllerSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateNonnegativeField(int64(spec.MinReadySeconds), fldPath.Child("minReadySeconds"))...)
 allErrs = append(allErrs, ValidateNonEmptySelector(spec.Selector, fldPath.Child("selector"))...)
 allErrs = append(allErrs, ValidateNonnegativeField(int64(spec.Replicas), fldPath.Child("replicas"))...)
 allErrs = append(allErrs, ValidatePodTemplateSpecForRC(spec.Template, spec.Selector, spec.Replicas, fldPath.Child("template"))...)
 return allErrs
}
func ValidatePodTemplateSpec(spec *core.PodTemplateSpec, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, unversionedvalidation.ValidateLabels(spec.Labels, fldPath.Child("labels"))...)
 allErrs = append(allErrs, ValidateAnnotations(spec.Annotations, fldPath.Child("annotations"))...)
 allErrs = append(allErrs, ValidatePodSpecificAnnotations(spec.Annotations, &spec.Spec, fldPath.Child("annotations"))...)
 allErrs = append(allErrs, ValidatePodSpec(&spec.Spec, fldPath.Child("spec"))...)
 return allErrs
}
func ValidateReadOnlyPersistentDisks(volumes []core.Volume, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i := range volumes {
  vol := &volumes[i]
  idxPath := fldPath.Index(i)
  if vol.GCEPersistentDisk != nil {
   if vol.GCEPersistentDisk.ReadOnly == false {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("gcePersistentDisk", "readOnly"), false, "must be true for replicated pods > 1; GCE PD can only be mounted on multiple machines if it is read-only"))
   }
  }
 }
 return allErrs
}
func ValidateTaintsInNodeAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 taints, err := helper.GetTaintsFromNodeAnnotations(annotations)
 if err != nil {
  allErrs = append(allErrs, field.Invalid(fldPath, core.TaintsAnnotationKey, err.Error()))
  return allErrs
 }
 if len(taints) > 0 {
  allErrs = append(allErrs, validateNodeTaints(taints, fldPath.Child(core.TaintsAnnotationKey))...)
 }
 return allErrs
}
func validateNodeTaints(taints []core.Taint, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrors := field.ErrorList{}
 uniqueTaints := map[core.TaintEffect]sets.String{}
 for i, currTaint := range taints {
  idxPath := fldPath.Index(i)
  allErrors = append(allErrors, unversionedvalidation.ValidateLabelName(currTaint.Key, idxPath.Child("key"))...)
  if errs := validation.IsValidLabelValue(currTaint.Value); len(errs) != 0 {
   allErrors = append(allErrors, field.Invalid(idxPath.Child("value"), currTaint.Value, strings.Join(errs, ";")))
  }
  allErrors = append(allErrors, validateTaintEffect(&currTaint.Effect, false, idxPath.Child("effect"))...)
  if len(uniqueTaints[currTaint.Effect]) > 0 && uniqueTaints[currTaint.Effect].Has(currTaint.Key) {
   duplicatedError := field.Duplicate(idxPath, currTaint)
   duplicatedError.Detail = "taints must be unique by key and effect pair"
   allErrors = append(allErrors, duplicatedError)
   continue
  }
  if len(uniqueTaints[currTaint.Effect]) == 0 {
   uniqueTaints[currTaint.Effect] = sets.String{}
  }
  uniqueTaints[currTaint.Effect].Insert(currTaint.Key)
 }
 return allErrors
}
func ValidateNodeSpecificAnnotations(annotations map[string]string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if annotations[core.TaintsAnnotationKey] != "" {
  allErrs = append(allErrs, ValidateTaintsInNodeAnnotations(annotations, fldPath)...)
 }
 if annotations[core.PreferAvoidPodsAnnotationKey] != "" {
  allErrs = append(allErrs, ValidateAvoidPodsInNodeAnnotations(annotations, fldPath)...)
 }
 return allErrs
}
func ValidateNode(node *core.Node) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fldPath := field.NewPath("metadata")
 allErrs := ValidateObjectMeta(&node.ObjectMeta, false, ValidateNodeName, fldPath)
 allErrs = append(allErrs, ValidateNodeSpecificAnnotations(node.ObjectMeta.Annotations, fldPath.Child("annotations"))...)
 if len(node.Spec.Taints) > 0 {
  allErrs = append(allErrs, validateNodeTaints(node.Spec.Taints, fldPath.Child("taints"))...)
 }
 allErrs = append(allErrs, ValidateNodeResources(node)...)
 if !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  if node.Spec.ConfigSource != nil {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "configSource"), "configSource may only be set if the DynamicKubeletConfig feature gate is enabled)"))
  }
  if node.Status.Config != nil {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("status", "config"), "config may only be set if the DynamicKubeletConfig feature gate is enabled)"))
  }
 }
 if len(node.Spec.PodCIDR) != 0 {
  _, err := ValidateCIDR(node.Spec.PodCIDR)
  if err != nil {
   allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "podCIDR"), node.Spec.PodCIDR, "not a valid CIDR"))
  }
 }
 return allErrs
}
func ValidateNodeResources(node *core.Node) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 hugePageSizes := sets.NewString()
 for k, v := range node.Status.Capacity {
  resPath := field.NewPath("status", "capacity", string(k))
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
  if helper.IsHugePageResourceName(k) && v.Value() > int64(0) {
   hugePageSizes.Insert(string(k))
  }
  if len(hugePageSizes) > 1 {
   allErrs = append(allErrs, field.Invalid(resPath, v, "may not have pre-allocated hugepages for multiple page sizes"))
  }
 }
 hugePageSizes = sets.NewString()
 for k, v := range node.Status.Allocatable {
  resPath := field.NewPath("status", "allocatable", string(k))
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
  if helper.IsHugePageResourceName(k) && v.Value() > int64(0) {
   hugePageSizes.Insert(string(k))
  }
  if len(hugePageSizes) > 1 {
   allErrs = append(allErrs, field.Invalid(resPath, v, "may not have pre-allocated hugepages for multiple page sizes"))
  }
 }
 return allErrs
}
func ValidateNodeUpdate(node, oldNode *core.Node) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fldPath := field.NewPath("metadata")
 allErrs := ValidateObjectMetaUpdate(&node.ObjectMeta, &oldNode.ObjectMeta, fldPath)
 allErrs = append(allErrs, ValidateNodeSpecificAnnotations(node.ObjectMeta.Annotations, fldPath.Child("annotations"))...)
 allErrs = append(allErrs, ValidateNodeResources(node)...)
 addresses := make(map[core.NodeAddress]bool)
 for i, address := range node.Status.Addresses {
  if _, ok := addresses[address]; ok {
   allErrs = append(allErrs, field.Duplicate(field.NewPath("status", "addresses").Index(i), address))
  }
  addresses[address] = true
 }
 if len(oldNode.Spec.PodCIDR) == 0 {
  oldNode.Spec.PodCIDR = node.Spec.PodCIDR
 } else {
  if oldNode.Spec.PodCIDR != node.Spec.PodCIDR {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "podCIDR"), "node updates may not change podCIDR except from \"\" to valid"))
  }
 }
 if len(oldNode.Spec.ProviderID) == 0 {
  oldNode.Spec.ProviderID = node.Spec.ProviderID
 } else {
  if oldNode.Spec.ProviderID != node.Spec.ProviderID {
   allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "providerID"), "node updates may not change providerID except from \"\" to valid"))
  }
 }
 if utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  if node.Spec.ConfigSource != nil {
   allErrs = append(allErrs, validateNodeConfigSourceSpec(node.Spec.ConfigSource, field.NewPath("spec", "configSource"))...)
  }
  oldNode.Spec.ConfigSource = node.Spec.ConfigSource
  if node.Status.Config != nil {
   allErrs = append(allErrs, validateNodeConfigStatus(node.Status.Config, field.NewPath("status", "config"))...)
  }
  oldNode.Status.Config = node.Status.Config
 }
 oldNode.ObjectMeta = node.ObjectMeta
 oldNode.Status.Capacity = node.Status.Capacity
 oldNode.Spec.Unschedulable = node.Spec.Unschedulable
 oldNode.Status = node.Status
 if len(node.Spec.Taints) > 0 {
  allErrs = append(allErrs, validateNodeTaints(node.Spec.Taints, fldPath.Child("taints"))...)
 }
 oldNode.Spec.Taints = node.Spec.Taints
 if !apiequality.Semantic.DeepEqual(oldNode, node) {
  klog.V(4).Infof("Update failed validation %#v vs %#v", oldNode, node)
  allErrs = append(allErrs, field.Forbidden(field.NewPath(""), "node updates may only change labels, taints, or capacity (or configSource, if the DynamicKubeletConfig feature gate is enabled)"))
 }
 return allErrs
}
func validateNodeConfigSourceSpec(source *core.NodeConfigSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 count := int(0)
 if source.ConfigMap != nil {
  count++
  allErrs = append(allErrs, validateConfigMapNodeConfigSourceSpec(source.ConfigMap, fldPath.Child("configMap"))...)
 }
 if count != 1 {
  allErrs = append(allErrs, field.Invalid(fldPath, source, "exactly one reference subfield must be non-nil"))
 }
 return allErrs
}
func validateConfigMapNodeConfigSourceSpec(source *core.ConfigMapNodeConfigSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if string(source.UID) != "" {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("uid"), "uid must not be set in spec"))
 }
 if source.ResourceVersion != "" {
  allErrs = append(allErrs, field.Forbidden(fldPath.Child("resourceVersion"), "resourceVersion must not be set in spec"))
 }
 return append(allErrs, validateConfigMapNodeConfigSource(source, fldPath)...)
}
func validateNodeConfigStatus(status *core.NodeConfigStatus, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if status.Assigned != nil {
  allErrs = append(allErrs, validateNodeConfigSourceStatus(status.Assigned, fldPath.Child("assigned"))...)
 }
 if status.Active != nil {
  allErrs = append(allErrs, validateNodeConfigSourceStatus(status.Active, fldPath.Child("active"))...)
 }
 if status.LastKnownGood != nil {
  allErrs = append(allErrs, validateNodeConfigSourceStatus(status.LastKnownGood, fldPath.Child("lastKnownGood"))...)
 }
 return allErrs
}
func validateNodeConfigSourceStatus(source *core.NodeConfigSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 count := int(0)
 if source.ConfigMap != nil {
  count++
  allErrs = append(allErrs, validateConfigMapNodeConfigSourceStatus(source.ConfigMap, fldPath.Child("configMap"))...)
 }
 if count != 1 {
  allErrs = append(allErrs, field.Invalid(fldPath, source, "exactly one reference subfield must be non-nil"))
 }
 return allErrs
}
func validateConfigMapNodeConfigSourceStatus(source *core.ConfigMapNodeConfigSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if string(source.UID) == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("uid"), "uid must be set in status"))
 }
 if source.ResourceVersion == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("resourceVersion"), "resourceVersion must be set in status"))
 }
 return append(allErrs, validateConfigMapNodeConfigSource(source, fldPath)...)
}
func validateConfigMapNodeConfigSource(source *core.ConfigMapNodeConfigSource, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if source.Namespace == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("namespace"), "namespace must be set"))
 } else {
  for _, msg := range ValidateNameFunc(ValidateNamespaceName)(source.Namespace, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("namespace"), source.Namespace, msg))
  }
 }
 if source.Name == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), "name must be set"))
 } else {
  for _, msg := range ValidateNameFunc(ValidateConfigMapName)(source.Name, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), source.Name, msg))
  }
 }
 if source.KubeletConfigKey == "" {
  allErrs = append(allErrs, field.Required(fldPath.Child("kubeletConfigKey"), "kubeletConfigKey must be set"))
 } else {
  for _, msg := range validation.IsConfigMapKey(source.KubeletConfigKey) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("kubeletConfigKey"), source.KubeletConfigKey, msg))
  }
 }
 return allErrs
}
func validateResourceName(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, msg := range validation.IsQualifiedName(value) {
  allErrs = append(allErrs, field.Invalid(fldPath, value, msg))
 }
 if len(allErrs) != 0 {
  return allErrs
 }
 if len(strings.Split(value, "/")) == 1 {
  if !helper.IsStandardResourceName(value) {
   return append(allErrs, field.Invalid(fldPath, value, "must be a standard resource type or fully qualified"))
  }
 }
 return allErrs
}
func validateContainerResourceName(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := validateResourceName(value, fldPath)
 if len(strings.Split(value, "/")) == 1 {
  if !helper.IsStandardContainerResourceName(value) {
   return append(allErrs, field.Invalid(fldPath, value, "must be a standard resource for containers"))
  }
 } else if !helper.IsNativeResource(core.ResourceName(value)) {
  if !helper.IsExtendedResourceName(core.ResourceName(value)) {
   return append(allErrs, field.Invalid(fldPath, value, "doesn't follow extended resource name standard"))
  }
 }
 return allErrs
}
func isLocalStorageResource(name string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if name == string(core.ResourceEphemeralStorage) || name == string(core.ResourceRequestsEphemeralStorage) || name == string(core.ResourceLimitsEphemeralStorage) {
  return true
 }
 return false
}
func ValidateResourceQuotaResourceName(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := validateResourceName(value, fldPath)
 if isLocalStorageResource(value) && !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
  return append(allErrs, field.Forbidden(fldPath, "ResourceEphemeralStorage field disabled by feature-gate for ResourceQuota"))
 }
 if len(strings.Split(value, "/")) == 1 {
  if !helper.IsStandardQuotaResourceName(value) {
   return append(allErrs, field.Invalid(fldPath, value, isInvalidQuotaResource))
  }
 }
 return allErrs
}
func validateLimitRangeTypeName(value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, msg := range validation.IsQualifiedName(value) {
  allErrs = append(allErrs, field.Invalid(fldPath, value, msg))
 }
 if len(allErrs) != 0 {
  return allErrs
 }
 if len(strings.Split(value, "/")) == 1 {
  if !helper.IsStandardLimitRangeType(value) {
   return append(allErrs, field.Invalid(fldPath, value, "must be a standard limit type or fully qualified"))
  }
 }
 return allErrs
}
func validateLimitRangeResourceName(limitType core.LimitType, value string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if value == string(core.ResourceEphemeralStorage) && !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
  return append(allErrs, field.Forbidden(fldPath, "ResourceEphemeralStorage field disabled by feature-gate for Resource LimitRange"))
 }
 switch limitType {
 case core.LimitTypePod, core.LimitTypeContainer:
  return validateContainerResourceName(value, fldPath)
 default:
  return validateResourceName(value, fldPath)
 }
}
func ValidateLimitRange(limitRange *core.LimitRange) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&limitRange.ObjectMeta, true, ValidateLimitRangeName, field.NewPath("metadata"))
 limitTypeSet := map[core.LimitType]bool{}
 fldPath := field.NewPath("spec", "limits")
 for i := range limitRange.Spec.Limits {
  idxPath := fldPath.Index(i)
  limit := &limitRange.Spec.Limits[i]
  allErrs = append(allErrs, validateLimitRangeTypeName(string(limit.Type), idxPath.Child("type"))...)
  _, found := limitTypeSet[limit.Type]
  if found {
   allErrs = append(allErrs, field.Duplicate(idxPath.Child("type"), limit.Type))
  }
  limitTypeSet[limit.Type] = true
  keys := sets.String{}
  min := map[string]resource.Quantity{}
  max := map[string]resource.Quantity{}
  defaults := map[string]resource.Quantity{}
  defaultRequests := map[string]resource.Quantity{}
  maxLimitRequestRatios := map[string]resource.Quantity{}
  for k, q := range limit.Max {
   allErrs = append(allErrs, validateLimitRangeResourceName(limit.Type, string(k), idxPath.Child("max").Key(string(k)))...)
   keys.Insert(string(k))
   max[string(k)] = q
  }
  for k, q := range limit.Min {
   allErrs = append(allErrs, validateLimitRangeResourceName(limit.Type, string(k), idxPath.Child("min").Key(string(k)))...)
   keys.Insert(string(k))
   min[string(k)] = q
  }
  if limit.Type == core.LimitTypePod {
   if len(limit.Default) > 0 {
    allErrs = append(allErrs, field.Forbidden(idxPath.Child("default"), "may not be specified when `type` is 'Pod'"))
   }
   if len(limit.DefaultRequest) > 0 {
    allErrs = append(allErrs, field.Forbidden(idxPath.Child("defaultRequest"), "may not be specified when `type` is 'Pod'"))
   }
  } else {
   for k, q := range limit.Default {
    allErrs = append(allErrs, validateLimitRangeResourceName(limit.Type, string(k), idxPath.Child("default").Key(string(k)))...)
    keys.Insert(string(k))
    defaults[string(k)] = q
   }
   for k, q := range limit.DefaultRequest {
    allErrs = append(allErrs, validateLimitRangeResourceName(limit.Type, string(k), idxPath.Child("defaultRequest").Key(string(k)))...)
    keys.Insert(string(k))
    defaultRequests[string(k)] = q
   }
  }
  if limit.Type == core.LimitTypePersistentVolumeClaim {
   _, minQuantityFound := limit.Min[core.ResourceStorage]
   _, maxQuantityFound := limit.Max[core.ResourceStorage]
   if !minQuantityFound && !maxQuantityFound {
    allErrs = append(allErrs, field.Required(idxPath.Child("limits"), "either minimum or maximum storage value is required, but neither was provided"))
   }
  }
  for k, q := range limit.MaxLimitRequestRatio {
   allErrs = append(allErrs, validateLimitRangeResourceName(limit.Type, string(k), idxPath.Child("maxLimitRequestRatio").Key(string(k)))...)
   keys.Insert(string(k))
   maxLimitRequestRatios[string(k)] = q
  }
  for k := range keys {
   minQuantity, minQuantityFound := min[k]
   maxQuantity, maxQuantityFound := max[k]
   defaultQuantity, defaultQuantityFound := defaults[k]
   defaultRequestQuantity, defaultRequestQuantityFound := defaultRequests[k]
   maxRatio, maxRatioFound := maxLimitRequestRatios[k]
   if minQuantityFound && maxQuantityFound && minQuantity.Cmp(maxQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("min").Key(string(k)), minQuantity, fmt.Sprintf("min value %s is greater than max value %s", minQuantity.String(), maxQuantity.String())))
   }
   if defaultRequestQuantityFound && minQuantityFound && minQuantity.Cmp(defaultRequestQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("defaultRequest").Key(string(k)), defaultRequestQuantity, fmt.Sprintf("min value %s is greater than default request value %s", minQuantity.String(), defaultRequestQuantity.String())))
   }
   if defaultRequestQuantityFound && maxQuantityFound && defaultRequestQuantity.Cmp(maxQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("defaultRequest").Key(string(k)), defaultRequestQuantity, fmt.Sprintf("default request value %s is greater than max value %s", defaultRequestQuantity.String(), maxQuantity.String())))
   }
   if defaultRequestQuantityFound && defaultQuantityFound && defaultRequestQuantity.Cmp(defaultQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("defaultRequest").Key(string(k)), defaultRequestQuantity, fmt.Sprintf("default request value %s is greater than default limit value %s", defaultRequestQuantity.String(), defaultQuantity.String())))
   }
   if defaultQuantityFound && minQuantityFound && minQuantity.Cmp(defaultQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("default").Key(string(k)), minQuantity, fmt.Sprintf("min value %s is greater than default value %s", minQuantity.String(), defaultQuantity.String())))
   }
   if defaultQuantityFound && maxQuantityFound && defaultQuantity.Cmp(maxQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("default").Key(string(k)), maxQuantity, fmt.Sprintf("default value %s is greater than max value %s", defaultQuantity.String(), maxQuantity.String())))
   }
   if maxRatioFound && maxRatio.Cmp(*resource.NewQuantity(1, resource.DecimalSI)) < 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("maxLimitRequestRatio").Key(string(k)), maxRatio, fmt.Sprintf("ratio %s is less than 1", maxRatio.String())))
   }
   if maxRatioFound && minQuantityFound && maxQuantityFound {
    maxRatioValue := float64(maxRatio.Value())
    minQuantityValue := minQuantity.Value()
    maxQuantityValue := maxQuantity.Value()
    if maxRatio.Value() < resource.MaxMilliValue && minQuantityValue < resource.MaxMilliValue && maxQuantityValue < resource.MaxMilliValue {
     maxRatioValue = float64(maxRatio.MilliValue()) / 1000
     minQuantityValue = minQuantity.MilliValue()
     maxQuantityValue = maxQuantity.MilliValue()
    }
    maxRatioLimit := float64(maxQuantityValue) / float64(minQuantityValue)
    if maxRatioValue > maxRatioLimit {
     allErrs = append(allErrs, field.Invalid(idxPath.Child("maxLimitRequestRatio").Key(string(k)), maxRatio, fmt.Sprintf("ratio %s is greater than max/min = %f", maxRatio.String(), maxRatioLimit)))
    }
   }
   if !helper.IsOvercommitAllowed(core.ResourceName(k)) && defaultQuantityFound && defaultRequestQuantityFound && defaultQuantity.Cmp(defaultRequestQuantity) != 0 {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("defaultRequest").Key(string(k)), defaultRequestQuantity, fmt.Sprintf("default value %s must equal to defaultRequest value %s in %s", defaultQuantity.String(), defaultRequestQuantity.String(), k)))
   }
  }
 }
 return allErrs
}
func ValidateServiceAccount(serviceAccount *core.ServiceAccount) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&serviceAccount.ObjectMeta, true, ValidateServiceAccountName, field.NewPath("metadata"))
 return allErrs
}
func ValidateServiceAccountUpdate(newServiceAccount, oldServiceAccount *core.ServiceAccount) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newServiceAccount.ObjectMeta, &oldServiceAccount.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateServiceAccount(newServiceAccount)...)
 return allErrs
}
func ValidateSecret(secret *core.Secret) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&secret.ObjectMeta, true, ValidateSecretName, field.NewPath("metadata"))
 dataPath := field.NewPath("data")
 totalSize := 0
 for key, value := range secret.Data {
  for _, msg := range validation.IsConfigMapKey(key) {
   allErrs = append(allErrs, field.Invalid(dataPath.Key(key), key, msg))
  }
  totalSize += len(value)
 }
 if totalSize > core.MaxSecretSize {
  allErrs = append(allErrs, field.TooLong(dataPath, "", core.MaxSecretSize))
 }
 switch secret.Type {
 case core.SecretTypeServiceAccountToken:
  if value := secret.Annotations[core.ServiceAccountNameKey]; len(value) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("metadata", "annotations").Key(core.ServiceAccountNameKey), ""))
  }
 case core.SecretTypeOpaque, "":
 case core.SecretTypeDockercfg:
  dockercfgBytes, exists := secret.Data[core.DockerConfigKey]
  if !exists {
   allErrs = append(allErrs, field.Required(dataPath.Key(core.DockerConfigKey), ""))
   break
  }
  if err := json.Unmarshal(dockercfgBytes, &map[string]interface{}{}); err != nil {
   allErrs = append(allErrs, field.Invalid(dataPath.Key(core.DockerConfigKey), "<secret contents redacted>", err.Error()))
  }
 case core.SecretTypeDockerConfigJson:
  dockerConfigJsonBytes, exists := secret.Data[core.DockerConfigJsonKey]
  if !exists {
   allErrs = append(allErrs, field.Required(dataPath.Key(core.DockerConfigJsonKey), ""))
   break
  }
  if err := json.Unmarshal(dockerConfigJsonBytes, &map[string]interface{}{}); err != nil {
   allErrs = append(allErrs, field.Invalid(dataPath.Key(core.DockerConfigJsonKey), "<secret contents redacted>", err.Error()))
  }
 case core.SecretTypeBasicAuth:
  _, usernameFieldExists := secret.Data[core.BasicAuthUsernameKey]
  _, passwordFieldExists := secret.Data[core.BasicAuthPasswordKey]
  if !usernameFieldExists && !passwordFieldExists {
   allErrs = append(allErrs, field.Required(field.NewPath("data[%s]").Key(core.BasicAuthUsernameKey), ""))
   allErrs = append(allErrs, field.Required(field.NewPath("data[%s]").Key(core.BasicAuthPasswordKey), ""))
   break
  }
 case core.SecretTypeSSHAuth:
  if len(secret.Data[core.SSHAuthPrivateKey]) == 0 {
   allErrs = append(allErrs, field.Required(field.NewPath("data[%s]").Key(core.SSHAuthPrivateKey), ""))
   break
  }
 case core.SecretTypeTLS:
  if _, exists := secret.Data[core.TLSCertKey]; !exists {
   allErrs = append(allErrs, field.Required(dataPath.Key(core.TLSCertKey), ""))
  }
  if _, exists := secret.Data[core.TLSPrivateKeyKey]; !exists {
   allErrs = append(allErrs, field.Required(dataPath.Key(core.TLSPrivateKeyKey), ""))
  }
 default:
 }
 return allErrs
}
func ValidateSecretUpdate(newSecret, oldSecret *core.Secret) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newSecret.ObjectMeta, &oldSecret.ObjectMeta, field.NewPath("metadata"))
 if len(newSecret.Type) == 0 {
  newSecret.Type = oldSecret.Type
 }
 allErrs = append(allErrs, ValidateImmutableField(newSecret.Type, oldSecret.Type, field.NewPath("type"))...)
 allErrs = append(allErrs, ValidateSecret(newSecret)...)
 return allErrs
}

var ValidateConfigMapName = apimachineryvalidation.NameIsDNSSubdomain

func ValidateConfigMap(cfg *core.ConfigMap) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateObjectMeta(&cfg.ObjectMeta, true, ValidateConfigMapName, field.NewPath("metadata"))...)
 totalSize := 0
 for key, value := range cfg.Data {
  for _, msg := range validation.IsConfigMapKey(key) {
   allErrs = append(allErrs, field.Invalid(field.NewPath("data").Key(key), key, msg))
  }
  if _, isValue := cfg.BinaryData[key]; isValue {
   msg := "duplicate of key present in binaryData"
   allErrs = append(allErrs, field.Invalid(field.NewPath("data").Key(key), key, msg))
  }
  totalSize += len(value)
 }
 for key, value := range cfg.BinaryData {
  for _, msg := range validation.IsConfigMapKey(key) {
   allErrs = append(allErrs, field.Invalid(field.NewPath("binaryData").Key(key), key, msg))
  }
  totalSize += len(value)
 }
 if totalSize > core.MaxSecretSize {
  allErrs = append(allErrs, field.TooLong(field.NewPath(""), cfg, core.MaxSecretSize))
 }
 return allErrs
}
func ValidateConfigMapUpdate(newCfg, oldCfg *core.ConfigMap) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateObjectMetaUpdate(&newCfg.ObjectMeta, &oldCfg.ObjectMeta, field.NewPath("metadata"))...)
 allErrs = append(allErrs, ValidateConfigMap(newCfg)...)
 return allErrs
}
func validateBasicResource(quantity resource.Quantity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if quantity.Value() < 0 {
  return field.ErrorList{field.Invalid(fldPath, quantity.Value(), "must be a valid resource quantity")}
 }
 return field.ErrorList{}
}
func ValidateResourceRequirements(requirements *core.ResourceRequirements, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 limPath := fldPath.Child("limits")
 reqPath := fldPath.Child("requests")
 limContainsCpuOrMemory := false
 reqContainsCpuOrMemory := false
 limContainsHugePages := false
 reqContainsHugePages := false
 supportedQoSComputeResources := sets.NewString(string(core.ResourceCPU), string(core.ResourceMemory))
 for resourceName, quantity := range requirements.Limits {
  fldPath := limPath.Key(string(resourceName))
  allErrs = append(allErrs, validateContainerResourceName(string(resourceName), fldPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(resourceName), quantity, fldPath)...)
  if resourceName == core.ResourceEphemeralStorage && !utilfeature.DefaultFeatureGate.Enabled(features.LocalStorageCapacityIsolation) {
   allErrs = append(allErrs, field.Forbidden(limPath, "ResourceEphemeralStorage field disabled by feature-gate for ResourceRequirements"))
  }
  if helper.IsHugePageResourceName(resourceName) {
   if !utilfeature.DefaultFeatureGate.Enabled(features.HugePages) {
    allErrs = append(allErrs, field.Forbidden(limPath, fmt.Sprintf("%s field disabled by feature-gate for ResourceRequirements", resourceName)))
   } else {
    limContainsHugePages = true
   }
  }
  if supportedQoSComputeResources.Has(string(resourceName)) {
   limContainsCpuOrMemory = true
  }
 }
 for resourceName, quantity := range requirements.Requests {
  fldPath := reqPath.Key(string(resourceName))
  allErrs = append(allErrs, validateContainerResourceName(string(resourceName), fldPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(resourceName), quantity, fldPath)...)
  limitQuantity, exists := requirements.Limits[resourceName]
  if exists {
   if quantity.Cmp(limitQuantity) != 0 && !helper.IsOvercommitAllowed(resourceName) {
    allErrs = append(allErrs, field.Invalid(reqPath, quantity.String(), fmt.Sprintf("must be equal to %s limit", resourceName)))
   } else if quantity.Cmp(limitQuantity) > 0 {
    allErrs = append(allErrs, field.Invalid(reqPath, quantity.String(), fmt.Sprintf("must be less than or equal to %s limit", resourceName)))
   }
  } else if !helper.IsOvercommitAllowed(resourceName) {
   allErrs = append(allErrs, field.Required(limPath, "Limit must be set for non overcommitable resources"))
  }
  if helper.IsHugePageResourceName(resourceName) {
   reqContainsHugePages = true
  }
  if supportedQoSComputeResources.Has(string(resourceName)) {
   reqContainsCpuOrMemory = true
  }
 }
 if !limContainsCpuOrMemory && !reqContainsCpuOrMemory && (reqContainsHugePages || limContainsHugePages) {
  allErrs = append(allErrs, field.Forbidden(fldPath, fmt.Sprintf("HugePages require cpu or memory")))
 }
 return allErrs
}
func validateResourceQuotaScopes(resourceQuotaSpec *core.ResourceQuotaSpec, fld *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(resourceQuotaSpec.Scopes) == 0 {
  return allErrs
 }
 hardLimits := sets.NewString()
 for k := range resourceQuotaSpec.Hard {
  hardLimits.Insert(string(k))
 }
 fldPath := fld.Child("scopes")
 scopeSet := sets.NewString()
 for _, scope := range resourceQuotaSpec.Scopes {
  if !helper.IsStandardResourceQuotaScope(string(scope)) {
   allErrs = append(allErrs, field.Invalid(fldPath, resourceQuotaSpec.Scopes, "unsupported scope"))
  }
  for _, k := range hardLimits.List() {
   if helper.IsStandardQuotaResourceName(k) && !helper.IsResourceQuotaScopeValidForResource(scope, k) {
    allErrs = append(allErrs, field.Invalid(fldPath, resourceQuotaSpec.Scopes, "unsupported scope applied to resource"))
   }
  }
  scopeSet.Insert(string(scope))
 }
 invalidScopePairs := []sets.String{sets.NewString(string(core.ResourceQuotaScopeBestEffort), string(core.ResourceQuotaScopeNotBestEffort)), sets.NewString(string(core.ResourceQuotaScopeTerminating), string(core.ResourceQuotaScopeNotTerminating))}
 for _, invalidScopePair := range invalidScopePairs {
  if scopeSet.HasAll(invalidScopePair.List()...) {
   allErrs = append(allErrs, field.Invalid(fldPath, resourceQuotaSpec.Scopes, "conflicting scopes"))
  }
 }
 return allErrs
}
func validateScopedResourceSelectorRequirement(resourceQuotaSpec *core.ResourceQuotaSpec, fld *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 hardLimits := sets.NewString()
 for k := range resourceQuotaSpec.Hard {
  hardLimits.Insert(string(k))
 }
 fldPath := fld.Child("matchExpressions")
 scopeSet := sets.NewString()
 for _, req := range resourceQuotaSpec.ScopeSelector.MatchExpressions {
  if !helper.IsStandardResourceQuotaScope(string(req.ScopeName)) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("scopeName"), req.ScopeName, "unsupported scope"))
  }
  for _, k := range hardLimits.List() {
   if helper.IsStandardQuotaResourceName(k) && !helper.IsResourceQuotaScopeValidForResource(req.ScopeName, k) {
    allErrs = append(allErrs, field.Invalid(fldPath, resourceQuotaSpec.ScopeSelector, "unsupported scope applied to resource"))
   }
  }
  switch req.ScopeName {
  case core.ResourceQuotaScopeBestEffort, core.ResourceQuotaScopeNotBestEffort, core.ResourceQuotaScopeTerminating, core.ResourceQuotaScopeNotTerminating:
   if req.Operator != core.ScopeSelectorOpExists {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("operator"), req.Operator, "must be 'Exist' only operator when scope is any of ResourceQuotaScopeTerminating, ResourceQuotaScopeNotTerminating, ResourceQuotaScopeBestEffort and ResourceQuotaScopeNotBestEffort"))
   }
  }
  switch req.Operator {
  case core.ScopeSelectorOpIn, core.ScopeSelectorOpNotIn:
   if len(req.Values) == 0 {
    allErrs = append(allErrs, field.Required(fldPath.Child("values"), "must be at least one value when `operator` is 'In' or 'NotIn' for scope selector"))
   }
  case core.ScopeSelectorOpExists, core.ScopeSelectorOpDoesNotExist:
   if len(req.Values) != 0 {
    allErrs = append(allErrs, field.Invalid(fldPath.Child("values"), req.Values, "must be no value when `operator` is 'Exist' or 'DoesNotExist' for scope selector"))
   }
  default:
   allErrs = append(allErrs, field.Invalid(fldPath.Child("operator"), req.Operator, "not a valid selector operator"))
  }
  scopeSet.Insert(string(req.ScopeName))
 }
 invalidScopePairs := []sets.String{sets.NewString(string(core.ResourceQuotaScopeBestEffort), string(core.ResourceQuotaScopeNotBestEffort)), sets.NewString(string(core.ResourceQuotaScopeTerminating), string(core.ResourceQuotaScopeNotTerminating))}
 for _, invalidScopePair := range invalidScopePairs {
  if scopeSet.HasAll(invalidScopePair.List()...) {
   allErrs = append(allErrs, field.Invalid(fldPath, resourceQuotaSpec.Scopes, "conflicting scopes"))
  }
 }
 return allErrs
}
func validateScopeSelector(resourceQuotaSpec *core.ResourceQuotaSpec, fld *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if resourceQuotaSpec.ScopeSelector == nil {
  return allErrs
 }
 if !utilfeature.DefaultFeatureGate.Enabled(features.ResourceQuotaScopeSelectors) && resourceQuotaSpec.ScopeSelector != nil {
  allErrs = append(allErrs, field.Forbidden(fld.Child("scopeSelector"), "ResourceQuotaScopeSelectors feature-gate is disabled"))
 }
 allErrs = append(allErrs, validateScopedResourceSelectorRequirement(resourceQuotaSpec, fld.Child("scopeSelector"))...)
 return allErrs
}
func ValidateResourceQuota(resourceQuota *core.ResourceQuota) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&resourceQuota.ObjectMeta, true, ValidateResourceQuotaName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateResourceQuotaSpec(&resourceQuota.Spec, field.NewPath("spec"))...)
 allErrs = append(allErrs, ValidateResourceQuotaStatus(&resourceQuota.Status, field.NewPath("status"))...)
 return allErrs
}
func ValidateResourceQuotaStatus(status *core.ResourceQuotaStatus, fld *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 fldPath := fld.Child("hard")
 for k, v := range status.Hard {
  resPath := fldPath.Key(string(k))
  allErrs = append(allErrs, ValidateResourceQuotaResourceName(string(k), resPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
 }
 fldPath = fld.Child("used")
 for k, v := range status.Used {
  resPath := fldPath.Key(string(k))
  allErrs = append(allErrs, ValidateResourceQuotaResourceName(string(k), resPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
 }
 return allErrs
}
func ValidateResourceQuotaSpec(resourceQuotaSpec *core.ResourceQuotaSpec, fld *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 fldPath := fld.Child("hard")
 for k, v := range resourceQuotaSpec.Hard {
  resPath := fldPath.Key(string(k))
  allErrs = append(allErrs, ValidateResourceQuotaResourceName(string(k), resPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
 }
 allErrs = append(allErrs, validateResourceQuotaScopes(resourceQuotaSpec, fld)...)
 allErrs = append(allErrs, validateScopeSelector(resourceQuotaSpec, fld)...)
 return allErrs
}
func ValidateResourceQuantityValue(resource string, value resource.Quantity, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 allErrs = append(allErrs, ValidateNonnegativeQuantity(value, fldPath)...)
 if helper.IsIntegerResourceName(resource) {
  if value.MilliValue()%int64(1000) != int64(0) {
   allErrs = append(allErrs, field.Invalid(fldPath, value, isNotIntegerErrorMsg))
  }
 }
 return allErrs
}
func ValidateResourceQuotaUpdate(newResourceQuota, oldResourceQuota *core.ResourceQuota) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newResourceQuota.ObjectMeta, &oldResourceQuota.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateResourceQuotaSpec(&newResourceQuota.Spec, field.NewPath("spec"))...)
 fldPath := field.NewPath("spec", "scopes")
 oldScopes := sets.NewString()
 newScopes := sets.NewString()
 for _, scope := range newResourceQuota.Spec.Scopes {
  newScopes.Insert(string(scope))
 }
 for _, scope := range oldResourceQuota.Spec.Scopes {
  oldScopes.Insert(string(scope))
 }
 if !oldScopes.Equal(newScopes) {
  allErrs = append(allErrs, field.Invalid(fldPath, newResourceQuota.Spec.Scopes, fieldImmutableErrorMsg))
 }
 newResourceQuota.Status = oldResourceQuota.Status
 return allErrs
}
func ValidateResourceQuotaStatusUpdate(newResourceQuota, oldResourceQuota *core.ResourceQuota) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newResourceQuota.ObjectMeta, &oldResourceQuota.ObjectMeta, field.NewPath("metadata"))
 if len(newResourceQuota.ResourceVersion) == 0 {
  allErrs = append(allErrs, field.Required(field.NewPath("resourceVersion"), ""))
 }
 fldPath := field.NewPath("status", "hard")
 for k, v := range newResourceQuota.Status.Hard {
  resPath := fldPath.Key(string(k))
  allErrs = append(allErrs, ValidateResourceQuotaResourceName(string(k), resPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
 }
 fldPath = field.NewPath("status", "used")
 for k, v := range newResourceQuota.Status.Used {
  resPath := fldPath.Key(string(k))
  allErrs = append(allErrs, ValidateResourceQuotaResourceName(string(k), resPath)...)
  allErrs = append(allErrs, ValidateResourceQuantityValue(string(k), v, resPath)...)
 }
 newResourceQuota.Spec = oldResourceQuota.Spec
 return allErrs
}
func ValidateNamespace(namespace *core.Namespace) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&namespace.ObjectMeta, false, ValidateNamespaceName, field.NewPath("metadata"))
 for i := range namespace.Spec.Finalizers {
  allErrs = append(allErrs, validateFinalizerName(string(namespace.Spec.Finalizers[i]), field.NewPath("spec", "finalizers"))...)
 }
 return allErrs
}
func validateFinalizerName(stringValue string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := apimachineryvalidation.ValidateFinalizerName(stringValue, fldPath)
 for _, err := range validateKubeFinalizerName(stringValue, fldPath) {
  allErrs = append(allErrs, err)
 }
 return allErrs
}
func validateKubeFinalizerName(stringValue string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if len(strings.Split(stringValue, "/")) == 1 {
  if !helper.IsStandardFinalizerName(stringValue) {
   return append(allErrs, field.Invalid(fldPath, stringValue, "name is neither a standard finalizer name nor is it fully qualified"))
  }
 }
 return allErrs
}
func ValidateNamespaceUpdate(newNamespace *core.Namespace, oldNamespace *core.Namespace) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newNamespace.ObjectMeta, &oldNamespace.ObjectMeta, field.NewPath("metadata"))
 newNamespace.Spec.Finalizers = oldNamespace.Spec.Finalizers
 newNamespace.Status = oldNamespace.Status
 return allErrs
}
func ValidateNamespaceStatusUpdate(newNamespace, oldNamespace *core.Namespace) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newNamespace.ObjectMeta, &oldNamespace.ObjectMeta, field.NewPath("metadata"))
 newNamespace.Spec = oldNamespace.Spec
 if newNamespace.DeletionTimestamp.IsZero() {
  if newNamespace.Status.Phase != core.NamespaceActive {
   allErrs = append(allErrs, field.Invalid(field.NewPath("status", "Phase"), newNamespace.Status.Phase, "may only be 'Active' if `deletionTimestamp` is empty"))
  }
 } else {
  if newNamespace.Status.Phase != core.NamespaceTerminating {
   allErrs = append(allErrs, field.Invalid(field.NewPath("status", "Phase"), newNamespace.Status.Phase, "may only be 'Terminating' if `deletionTimestamp` is not empty"))
  }
 }
 return allErrs
}
func ValidateNamespaceFinalizeUpdate(newNamespace, oldNamespace *core.Namespace) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newNamespace.ObjectMeta, &oldNamespace.ObjectMeta, field.NewPath("metadata"))
 fldPath := field.NewPath("spec", "finalizers")
 for i := range newNamespace.Spec.Finalizers {
  idxPath := fldPath.Index(i)
  allErrs = append(allErrs, validateFinalizerName(string(newNamespace.Spec.Finalizers[i]), idxPath)...)
 }
 newNamespace.Status = oldNamespace.Status
 return allErrs
}
func ValidateEndpoints(endpoints *core.Endpoints) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMeta(&endpoints.ObjectMeta, true, ValidateEndpointsName, field.NewPath("metadata"))
 allErrs = append(allErrs, ValidateEndpointsSpecificAnnotations(endpoints.Annotations, field.NewPath("annotations"))...)
 allErrs = append(allErrs, validateEndpointSubsets(endpoints.Subsets, field.NewPath("subsets"))...)
 return allErrs
}
func validateEndpointSubsets(subsets []core.EndpointSubset, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i := range subsets {
  ss := &subsets[i]
  idxPath := fldPath.Index(i)
  if len(ss.Addresses) == 0 && len(ss.NotReadyAddresses) == 0 {
   allErrs = append(allErrs, field.Required(idxPath, "must specify `addresses` or `notReadyAddresses`"))
  }
  for addr := range ss.Addresses {
   allErrs = append(allErrs, validateEndpointAddress(&ss.Addresses[addr], idxPath.Child("addresses").Index(addr))...)
  }
  for addr := range ss.NotReadyAddresses {
   allErrs = append(allErrs, validateEndpointAddress(&ss.NotReadyAddresses[addr], idxPath.Child("notReadyAddresses").Index(addr))...)
  }
  for port := range ss.Ports {
   allErrs = append(allErrs, validateEndpointPort(&ss.Ports[port], len(ss.Ports) > 1, idxPath.Child("ports").Index(port))...)
  }
 }
 return allErrs
}
func validateEndpointAddress(address *core.EndpointAddress, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for _, msg := range validation.IsValidIP(address.IP) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("ip"), address.IP, msg))
 }
 if len(address.Hostname) > 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(address.Hostname, fldPath.Child("hostname"))...)
 }
 if address.NodeName != nil {
  for _, msg := range ValidateNodeName(*address.NodeName, false) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("nodeName"), *address.NodeName, msg))
  }
 }
 allErrs = append(allErrs, validateNonSpecialIP(address.IP, fldPath.Child("ip"))...)
 return allErrs
}
func validateNonSpecialIP(ipAddress string, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 ip := net.ParseIP(ipAddress)
 if ip == nil {
  allErrs = append(allErrs, field.Invalid(fldPath, ipAddress, "must be a valid IP address"))
  return allErrs
 }
 if ip.IsUnspecified() {
  allErrs = append(allErrs, field.Invalid(fldPath, ipAddress, "may not be unspecified (0.0.0.0)"))
 }
 if ip.IsLoopback() {
  allErrs = append(allErrs, field.Invalid(fldPath, ipAddress, "may not be in the loopback range (127.0.0.0/8)"))
 }
 if ip.IsLinkLocalUnicast() {
  allErrs = append(allErrs, field.Invalid(fldPath, ipAddress, "may not be in the link-local range (169.254.0.0/16)"))
 }
 if ip.IsLinkLocalMulticast() {
  allErrs = append(allErrs, field.Invalid(fldPath, ipAddress, "may not be in the link-local multicast range (224.0.0.0/24)"))
 }
 return allErrs
}
func validateEndpointPort(port *core.EndpointPort, requireName bool, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if requireName && len(port.Name) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
 } else if len(port.Name) != 0 {
  allErrs = append(allErrs, ValidateDNS1123Label(port.Name, fldPath.Child("name"))...)
 }
 for _, msg := range validation.IsValidPortNum(int(port.Port)) {
  allErrs = append(allErrs, field.Invalid(fldPath.Child("port"), port.Port, msg))
 }
 if len(port.Protocol) == 0 {
  allErrs = append(allErrs, field.Required(fldPath.Child("protocol"), ""))
 } else if !utilfeature.DefaultFeatureGate.Enabled(features.SCTPSupport) && port.Protocol == core.ProtocolSCTP {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("protocol"), port.Protocol, []string{string(core.ProtocolTCP), string(core.ProtocolUDP)}))
 } else if !supportedPortProtocols.Has(string(port.Protocol)) {
  allErrs = append(allErrs, field.NotSupported(fldPath.Child("protocol"), port.Protocol, supportedPortProtocols.List()))
 }
 return allErrs
}
func ValidateEndpointsUpdate(newEndpoints, oldEndpoints *core.Endpoints) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := ValidateObjectMetaUpdate(&newEndpoints.ObjectMeta, &oldEndpoints.ObjectMeta, field.NewPath("metadata"))
 allErrs = append(allErrs, validateEndpointSubsets(newEndpoints.Subsets, field.NewPath("subsets"))...)
 allErrs = append(allErrs, ValidateEndpointsSpecificAnnotations(newEndpoints.Annotations, field.NewPath("annotations"))...)
 return allErrs
}
func ValidateSecurityContext(sc *core.SecurityContext, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if sc == nil {
  return allErrs
 }
 if sc.Privileged != nil {
  if *sc.Privileged && !capabilities.Get().AllowPrivileged {
   allErrs = append(allErrs, field.Forbidden(fldPath.Child("privileged"), "disallowed by cluster policy"))
  }
 }
 if sc.RunAsUser != nil {
  for _, msg := range validation.IsValidUserID(*sc.RunAsUser) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsUser"), *sc.RunAsUser, msg))
  }
 }
 if sc.RunAsGroup != nil {
  for _, msg := range validation.IsValidGroupID(*sc.RunAsGroup) {
   allErrs = append(allErrs, field.Invalid(fldPath.Child("runAsGroup"), *sc.RunAsGroup, msg))
  }
 }
 if sc.AllowPrivilegeEscalation != nil && !*sc.AllowPrivilegeEscalation {
  if sc.Privileged != nil && *sc.Privileged {
   allErrs = append(allErrs, field.Invalid(fldPath, sc, "cannot set `allowPrivilegeEscalation` to false and `privileged` to true"))
  }
  if sc.Capabilities != nil {
   for _, cap := range sc.Capabilities.Add {
    if string(cap) == "CAP_SYS_ADMIN" {
     allErrs = append(allErrs, field.Invalid(fldPath, sc, "cannot set `allowPrivilegeEscalation` to false and `capabilities.Add` CAP_SYS_ADMIN"))
    }
   }
  }
 }
 return allErrs
}
func ValidatePodLogOptions(opts *core.PodLogOptions) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if opts.TailLines != nil && *opts.TailLines < 0 {
  allErrs = append(allErrs, field.Invalid(field.NewPath("tailLines"), *opts.TailLines, isNegativeErrorMsg))
 }
 if opts.LimitBytes != nil && *opts.LimitBytes < 1 {
  allErrs = append(allErrs, field.Invalid(field.NewPath("limitBytes"), *opts.LimitBytes, "must be greater than 0"))
 }
 switch {
 case opts.SinceSeconds != nil && opts.SinceTime != nil:
  allErrs = append(allErrs, field.Forbidden(field.NewPath(""), "at most one of `sinceTime` or `sinceSeconds` may be specified"))
 case opts.SinceSeconds != nil:
  if *opts.SinceSeconds < 1 {
   allErrs = append(allErrs, field.Invalid(field.NewPath("sinceSeconds"), *opts.SinceSeconds, "must be greater than 0"))
  }
 }
 return allErrs
}
func ValidateLoadBalancerStatus(status *core.LoadBalancerStatus, fldPath *field.Path) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 for i, ingress := range status.Ingress {
  idxPath := fldPath.Child("ingress").Index(i)
  if len(ingress.IP) > 0 {
   if isIP := (net.ParseIP(ingress.IP) != nil); !isIP {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("ip"), ingress.IP, "must be a valid IP address"))
   }
  }
  if len(ingress.Hostname) > 0 {
   for _, msg := range validation.IsDNS1123Subdomain(ingress.Hostname) {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("hostname"), ingress.Hostname, msg))
   }
   if isIP := (net.ParseIP(ingress.Hostname) != nil); isIP {
    allErrs = append(allErrs, field.Invalid(idxPath.Child("hostname"), ingress.Hostname, "must be a DNS name, not an IP address"))
   }
  }
 }
 return allErrs
}
func validateVolumeNodeAffinity(nodeAffinity *core.VolumeNodeAffinity, fldPath *field.Path) (bool, field.ErrorList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 allErrs := field.ErrorList{}
 if nodeAffinity == nil {
  return false, allErrs
 }
 if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
  allErrs = append(allErrs, field.Forbidden(fldPath, "Volume node affinity is disabled by feature-gate"))
 }
 if nodeAffinity.Required != nil {
  allErrs = append(allErrs, ValidateNodeSelector(nodeAffinity.Required, fldPath.Child("required"))...)
 } else {
  allErrs = append(allErrs, field.Required(fldPath.Child("required"), "must specify required node constraints"))
 }
 return true, allErrs
}
func ValidateCIDR(cidr string) (*net.IPNet, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, net, err := net.ParseCIDR(cidr)
 if err != nil {
  return nil, err
 }
 return net, nil
}
func IsDecremented(update, old *int32) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if update == nil && old != nil {
  return true
 }
 if update == nil || old == nil {
  return false
 }
 return *update < *old
}
