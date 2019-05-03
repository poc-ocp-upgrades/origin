package simple

import (
	"fmt"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/klog"
	"strings"
)

const defaultDNSSuffix = "router.default.svc.cluster.local"

type SimpleAllocationPlugin struct{ DNSSuffix string }

func NewSimpleAllocationPlugin(suffix string) (*SimpleAllocationPlugin, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(suffix) == 0 {
		suffix = defaultDNSSuffix
	}
	klog.V(4).Infof("Route plugin initialized with suffix=%s", suffix)
	if len(kvalidation.IsDNS1123Subdomain(suffix)) != 0 {
		return nil, fmt.Errorf("invalid DNS suffix: %s", suffix)
	}
	return &SimpleAllocationPlugin{DNSSuffix: suffix}, nil
}
func (p *SimpleAllocationPlugin) Allocate(route *routeapi.Route) (*routeapi.RouterShard, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Allocating global shard *.%s to Route: %s", p.DNSSuffix, route.Name)
	return &routeapi.RouterShard{ShardName: "global", DNSSuffix: p.DNSSuffix}, nil
}
func (p *SimpleAllocationPlugin) GenerateHostname(route *routeapi.Route, shard *routeapi.RouterShard) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(route.Name) == 0 || len(route.Namespace) == 0 {
		return ""
	}
	return fmt.Sprintf("%s-%s.%s", strings.Replace(route.Name, ".", "-", -1), route.Namespace, shard.DNSSuffix)
}
