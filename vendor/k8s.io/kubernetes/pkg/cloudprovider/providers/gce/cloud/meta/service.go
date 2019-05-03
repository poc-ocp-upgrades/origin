package meta

import (
 "errors"
 "fmt"
 "reflect"
 "sort"
)

type ServiceInfo struct {
 Object              string
 Service             string
 Resource            string
 version             Version
 keyType             KeyType
 serviceType         reflect.Type
 additionalMethods   []string
 options             int
 aggregatedListField string
}

func (i *ServiceInfo) Version() Version {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if i.version == "" {
  return VersionGA
 }
 return i.version
}
func (i *ServiceInfo) VersionTitle() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch i.Version() {
 case VersionGA:
  return "GA"
 case VersionAlpha:
  return "Alpha"
 case VersionBeta:
  return "Beta"
 }
 panic(fmt.Errorf("invalid version %q", i.Version()))
}
func (i *ServiceInfo) WrapType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch i.Version() {
 case VersionGA:
  return i.Service
 case VersionAlpha:
  return "Alpha" + i.Service
 case VersionBeta:
  return "Beta" + i.Service
 }
 return "Invalid"
}
func (i *ServiceInfo) WrapTypeOps() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.WrapType() + "Ops"
}
func (i *ServiceInfo) FQObjectType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%v.%v", i.Version(), i.Object)
}
func (i *ServiceInfo) ObjectListType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%v.%vList", i.Version(), i.Object)
}
func (i *ServiceInfo) ObjectAggregatedListType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("%v.%vAggregatedList", i.Version(), i.Object)
}
func (i *ServiceInfo) MockWrapType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "Mock" + i.WrapType()
}
func (i *ServiceInfo) MockField() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "Mock" + i.WrapType()
}
func (i *ServiceInfo) GCEWrapType() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "GCE" + i.WrapType()
}
func (i *ServiceInfo) Field() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "gce" + i.WrapType()
}
func (i *ServiceInfo) Methods() []*Method {
 _logClusterCodePath()
 defer _logClusterCodePath()
 methods := map[string]bool{}
 for _, m := range i.additionalMethods {
  methods[m] = true
 }
 var ret []*Method
 for j := 0; j < i.serviceType.NumMethod(); j++ {
  m := i.serviceType.Method(j)
  if _, ok := methods[m.Name]; !ok {
   continue
  }
  ret = append(ret, newMethod(i, m))
  methods[m.Name] = false
 }
 for k, b := range methods {
  if b {
   panic(fmt.Errorf("method %q was not found in service %q", k, i.Service))
  }
 }
 return ret
}
func (i *ServiceInfo) KeyIsGlobal() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.keyType == Global
}
func (i *ServiceInfo) KeyIsRegional() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.keyType == Regional
}
func (i *ServiceInfo) KeyIsZonal() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.keyType == Zonal
}
func (i *ServiceInfo) KeyIsProject() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.Service == "Projects"
}
func (i *ServiceInfo) MakeKey(name, location string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch i.keyType {
 case Global:
  return fmt.Sprintf("GlobalKey(%q)", name)
 case Regional:
  return fmt.Sprintf("RegionalKey(%q, %q)", name, location)
 case Zonal:
  return fmt.Sprintf("ZonalKey(%q, %q)", name, location)
 }
 return "Invalid"
}
func (i *ServiceInfo) GenerateGet() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&NoGet == 0
}
func (i *ServiceInfo) GenerateList() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&NoList == 0
}
func (i *ServiceInfo) GenerateDelete() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&NoDelete == 0
}
func (i *ServiceInfo) GenerateInsert() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&NoInsert == 0
}
func (i *ServiceInfo) GenerateCustomOps() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&CustomOps != 0
}
func (i *ServiceInfo) AggregatedList() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return i.options&AggregatedList != 0
}
func (i *ServiceInfo) AggregatedListField() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if i.aggregatedListField == "" {
  return i.Service
 }
 return i.aggregatedListField
}

type ServiceGroup struct {
 Alpha *ServiceInfo
 Beta  *ServiceInfo
 GA    *ServiceInfo
}

func (sg *ServiceGroup) Service() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sg.ServiceInfo().Service
}
func (sg *ServiceGroup) ServiceInfo() *ServiceInfo {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch {
 case sg.GA != nil:
  return sg.GA
 case sg.Alpha != nil:
  return sg.Alpha
 case sg.Beta != nil:
  return sg.Beta
 default:
  panic(errors.New("service group is empty"))
 }
}
func (sg *ServiceGroup) HasGA() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sg.GA != nil
}
func (sg *ServiceGroup) HasAlpha() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sg.Alpha != nil
}
func (sg *ServiceGroup) HasBeta() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return sg.Beta != nil
}
func groupServices(services []*ServiceInfo) map[string]*ServiceGroup {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := map[string]*ServiceGroup{}
 for _, si := range services {
  if _, ok := ret[si.Service]; !ok {
   ret[si.Service] = &ServiceGroup{}
  }
  group := ret[si.Service]
  switch si.Version() {
  case VersionAlpha:
   group.Alpha = si
  case VersionBeta:
   group.Beta = si
  case VersionGA:
   group.GA = si
  }
 }
 return ret
}

var AllServicesByGroup map[string]*ServiceGroup
var SortedServicesGroups []*ServiceGroup

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 AllServicesByGroup = groupServices(AllServices)
 for _, sg := range AllServicesByGroup {
  SortedServicesGroups = append(SortedServicesGroups, sg)
 }
 sort.Slice(SortedServicesGroups, func(i, j int) bool {
  return SortedServicesGroups[i].Service() < SortedServicesGroups[j].Service()
 })
}
