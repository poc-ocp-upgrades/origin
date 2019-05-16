package options

import (
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"net"
)

var DefaultServiceNodePortRange = utilnet.PortRange{Base: 30000, Size: 2768}
var DefaultServiceIPCIDR net.IPNet = net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}
