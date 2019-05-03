package aws

import (
 "encoding/json"
 "fmt"
 "github.com/aws/aws-sdk-go/service/ec2"
)

type IPPermissionSet map[string]*ec2.IpPermission

func NewIPPermissionSet(items ...*ec2.IpPermission) IPPermissionSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s := make(IPPermissionSet)
 s.Insert(items...)
 return s
}
func (s IPPermissionSet) Ungroup() IPPermissionSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 l := []*ec2.IpPermission{}
 for _, p := range s.List() {
  if len(p.IpRanges) <= 1 {
   l = append(l, p)
   continue
  }
  for _, ipRange := range p.IpRanges {
   c := &ec2.IpPermission{}
   *c = *p
   c.IpRanges = []*ec2.IpRange{ipRange}
   l = append(l, c)
  }
 }
 l2 := []*ec2.IpPermission{}
 for _, p := range l {
  if len(p.UserIdGroupPairs) <= 1 {
   l2 = append(l2, p)
   continue
  }
  for _, u := range p.UserIdGroupPairs {
   c := &ec2.IpPermission{}
   *c = *p
   c.UserIdGroupPairs = []*ec2.UserIdGroupPair{u}
   l2 = append(l, c)
  }
 }
 l3 := []*ec2.IpPermission{}
 for _, p := range l2 {
  if len(p.PrefixListIds) <= 1 {
   l3 = append(l3, p)
   continue
  }
  for _, v := range p.PrefixListIds {
   c := &ec2.IpPermission{}
   *c = *p
   c.PrefixListIds = []*ec2.PrefixListId{v}
   l3 = append(l3, c)
  }
 }
 return NewIPPermissionSet(l3...)
}
func (s IPPermissionSet) Insert(items ...*ec2.IpPermission) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, p := range items {
  k := keyForIPPermission(p)
  s[k] = p
 }
}
func (s IPPermissionSet) List() []*ec2.IpPermission {
 _logClusterCodePath()
 defer _logClusterCodePath()
 res := make([]*ec2.IpPermission, 0, len(s))
 for _, v := range s {
  res = append(res, v)
 }
 return res
}
func (s IPPermissionSet) IsSuperset(s2 IPPermissionSet) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for k := range s2 {
  _, found := s[k]
  if !found {
   return false
  }
 }
 return true
}
func (s IPPermissionSet) Equal(s2 IPPermissionSet) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(s) == len(s2) && s.IsSuperset(s2)
}
func (s IPPermissionSet) Difference(s2 IPPermissionSet) IPPermissionSet {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := NewIPPermissionSet()
 for k, v := range s {
  _, found := s2[k]
  if !found {
   result[k] = v
  }
 }
 return result
}
func (s IPPermissionSet) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(s)
}
func keyForIPPermission(p *ec2.IpPermission) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 v, err := json.Marshal(p)
 if err != nil {
  panic(fmt.Sprintf("error building JSON representation of ec2.IpPermission: %v", err))
 }
 return string(v)
}
