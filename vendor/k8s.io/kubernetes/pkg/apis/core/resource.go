package core

import (
 "k8s.io/apimachinery/pkg/api/resource"
)

func (self ResourceName) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return string(self)
}
func (self *ResourceList) Cpu() *resource.Quantity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, ok := (*self)[ResourceCPU]; ok {
  return &val
 }
 return &resource.Quantity{Format: resource.DecimalSI}
}
func (self *ResourceList) Memory() *resource.Quantity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, ok := (*self)[ResourceMemory]; ok {
  return &val
 }
 return &resource.Quantity{Format: resource.BinarySI}
}
func (self *ResourceList) Pods() *resource.Quantity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, ok := (*self)[ResourcePods]; ok {
  return &val
 }
 return &resource.Quantity{}
}
func (self *ResourceList) StorageEphemeral() *resource.Quantity {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, ok := (*self)[ResourceEphemeralStorage]; ok {
  return &val
 }
 return &resource.Quantity{}
}
