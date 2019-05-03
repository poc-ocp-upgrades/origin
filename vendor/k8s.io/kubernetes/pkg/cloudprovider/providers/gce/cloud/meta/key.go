package meta

import (
 "fmt"
 "regexp"
)

type Key struct {
 Name   string
 Zone   string
 Region string
}
type KeyType string

const (
 Zonal    = "zonal"
 Regional = "regional"
 Global   = "global"
)

var (
 locationRegexp = regexp.MustCompile("^[a-z](?:[-a-z0-9]+)?$")
)

func ZonalKey(name, zone string) *Key {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &Key{name, zone, ""}
}
func RegionalKey(name, region string) *Key {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &Key{name, "", region}
}
func GlobalKey(name string) *Key {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &Key{name, "", ""}
}
func (k *Key) Type() KeyType {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch {
 case k.Zone != "":
  return Zonal
 case k.Region != "":
  return Regional
 default:
  return Global
 }
}
func (k Key) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch k.Type() {
 case Zonal:
  return fmt.Sprintf("Key{%q, zone: %q}", k.Name, k.Zone)
 case Regional:
  return fmt.Sprintf("Key{%q, region: %q}", k.Name, k.Region)
 default:
  return fmt.Sprintf("Key{%q}", k.Name)
 }
}
func (k *Key) Valid() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if k.Zone != "" && k.Region != "" {
  return false
 }
 switch {
 case k.Region != "":
  return locationRegexp.Match([]byte(k.Region))
 case k.Zone != "":
  return locationRegexp.Match([]byte(k.Zone))
 }
 return true
}
func KeysToMap(keys ...Key) map[Key]bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := map[Key]bool{}
 for _, k := range keys {
  ret[k] = true
 }
 return ret
}
