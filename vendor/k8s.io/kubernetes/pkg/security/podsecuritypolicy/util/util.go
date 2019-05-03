package util

import (
 "fmt"
 "strings"
 policy "k8s.io/api/policy/v1beta1"
 "k8s.io/apimachinery/pkg/util/sets"
 api "k8s.io/kubernetes/pkg/apis/core"
)

const (
 ValidatedPSPAnnotation = "kubernetes.io/psp"
)

func GetAllFSTypesExcept(exceptions ...string) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fstypes := GetAllFSTypesAsSet()
 for _, e := range exceptions {
  fstypes.Delete(e)
 }
 return fstypes
}
func GetAllFSTypesAsSet() sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fstypes := sets.NewString()
 fstypes.Insert(string(policy.HostPath), string(policy.AzureFile), string(policy.Flocker), string(policy.FlexVolume), string(policy.EmptyDir), string(policy.GCEPersistentDisk), string(policy.AWSElasticBlockStore), string(policy.GitRepo), string(policy.Secret), string(policy.NFS), string(policy.ISCSI), string(policy.Glusterfs), string(policy.PersistentVolumeClaim), string(policy.RBD), string(policy.Cinder), string(policy.CephFS), string(policy.DownwardAPI), string(policy.FC), string(policy.ConfigMap), string(policy.VsphereVolume), string(policy.Quobyte), string(policy.AzureDisk), string(policy.PhotonPersistentDisk), string(policy.StorageOS), string(policy.Projected), string(policy.PortworxVolume), string(policy.ScaleIO), string(policy.CSI))
 return fstypes
}
func GetVolumeFSType(v api.Volume) (policy.FSType, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch {
 case v.HostPath != nil:
  return policy.HostPath, nil
 case v.EmptyDir != nil:
  return policy.EmptyDir, nil
 case v.GCEPersistentDisk != nil:
  return policy.GCEPersistentDisk, nil
 case v.AWSElasticBlockStore != nil:
  return policy.AWSElasticBlockStore, nil
 case v.GitRepo != nil:
  return policy.GitRepo, nil
 case v.Secret != nil:
  return policy.Secret, nil
 case v.NFS != nil:
  return policy.NFS, nil
 case v.ISCSI != nil:
  return policy.ISCSI, nil
 case v.Glusterfs != nil:
  return policy.Glusterfs, nil
 case v.PersistentVolumeClaim != nil:
  return policy.PersistentVolumeClaim, nil
 case v.RBD != nil:
  return policy.RBD, nil
 case v.FlexVolume != nil:
  return policy.FlexVolume, nil
 case v.Cinder != nil:
  return policy.Cinder, nil
 case v.CephFS != nil:
  return policy.CephFS, nil
 case v.Flocker != nil:
  return policy.Flocker, nil
 case v.DownwardAPI != nil:
  return policy.DownwardAPI, nil
 case v.FC != nil:
  return policy.FC, nil
 case v.AzureFile != nil:
  return policy.AzureFile, nil
 case v.ConfigMap != nil:
  return policy.ConfigMap, nil
 case v.VsphereVolume != nil:
  return policy.VsphereVolume, nil
 case v.Quobyte != nil:
  return policy.Quobyte, nil
 case v.AzureDisk != nil:
  return policy.AzureDisk, nil
 case v.PhotonPersistentDisk != nil:
  return policy.PhotonPersistentDisk, nil
 case v.StorageOS != nil:
  return policy.StorageOS, nil
 case v.Projected != nil:
  return policy.Projected, nil
 case v.PortworxVolume != nil:
  return policy.PortworxVolume, nil
 case v.ScaleIO != nil:
  return policy.ScaleIO, nil
 }
 return "", fmt.Errorf("unknown volume type for volume: %#v", v)
}
func FSTypeToStringSet(fsTypes []policy.FSType) sets.String {
 _logClusterCodePath()
 defer _logClusterCodePath()
 set := sets.NewString()
 for _, v := range fsTypes {
  set.Insert(string(v))
 }
 return set
}
func PSPAllowsAllVolumes(psp *policy.PodSecurityPolicy) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return PSPAllowsFSType(psp, policy.All)
}
func PSPAllowsFSType(psp *policy.PodSecurityPolicy, fsType policy.FSType) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if psp == nil {
  return false
 }
 for _, v := range psp.Spec.Volumes {
  if v == fsType || v == policy.All {
   return true
  }
 }
 return false
}
func UserFallsInRange(id int64, rng policy.IDRange) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return id >= rng.Min && id <= rng.Max
}
func GroupFallsInRange(id int64, rng policy.IDRange) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return id >= rng.Min && id <= rng.Max
}
func AllowsHostVolumePath(psp *policy.PodSecurityPolicy, hostPath string) (pathIsAllowed, mustBeReadOnly bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if psp == nil {
  return false, false
 }
 if len(psp.Spec.AllowedHostPaths) == 0 {
  return true, false
 }
 for _, allowedPath := range psp.Spec.AllowedHostPaths {
  if hasPathPrefix(hostPath, allowedPath.PathPrefix) {
   if !allowedPath.ReadOnly {
    return true, allowedPath.ReadOnly
   }
   pathIsAllowed = true
   mustBeReadOnly = true
  }
 }
 return pathIsAllowed, mustBeReadOnly
}
func hasPathPrefix(s, pathPrefix string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s = strings.TrimSuffix(s, "/")
 pathPrefix = strings.TrimSuffix(pathPrefix, "/")
 if !strings.HasPrefix(s, pathPrefix) {
  return false
 }
 pathPrefixLength := len(pathPrefix)
 if len(s) == pathPrefixLength {
  return true
 }
 if s[pathPrefixLength:pathPrefixLength+1] == "/" {
  return true
 }
 return false
}
func EqualStringSlices(a, b []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(a) != len(b) {
  return false
 }
 for i := 0; i < len(a); i++ {
  if a[i] != b[i] {
   return false
  }
 }
 return true
}
