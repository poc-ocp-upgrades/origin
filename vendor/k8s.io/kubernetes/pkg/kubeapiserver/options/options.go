package options

import (
 "net"
 utilnet "k8s.io/apimachinery/pkg/util/net"
)

var DefaultServiceNodePortRange = utilnet.PortRange{Base: 30000, Size: 2768}
var DefaultServiceIPCIDR net.IPNet = net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)}
