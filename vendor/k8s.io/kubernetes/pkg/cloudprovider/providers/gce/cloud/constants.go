package cloud

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.ToUpper(string(n))
}
func NetworkTierGCEValueToType(s string) NetworkTier {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch s {
	case NetworkTierStandard.ToGCEValue():
		return NetworkTierStandard
	case NetworkTierPremium.ToGCEValue():
		return NetworkTierPremium
	default:
		return NetworkTier(s)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
