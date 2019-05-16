package node

import (
	networkapi "github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network"
)

type singleTenantPlugin struct{}

func NewSingleTenantPlugin() osdnPolicy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &singleTenantPlugin{}
}
func (sp *singleTenantPlugin) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return network.SingleTenantPluginName
}
func (sp *singleTenantPlugin) SupportsVNIDs() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (sp *singleTenantPlugin) Start(node *OsdnNode) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	otx := node.oc.NewTransaction()
	otx.AddFlow("table=80, priority=200, actions=output:NXM_NX_REG2[]")
	return otx.Commit()
}
func (sp *singleTenantPlugin) AddNetNamespace(netns *networkapi.NetNamespace) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (sp *singleTenantPlugin) UpdateNetNamespace(netns *networkapi.NetNamespace, oldNetID uint32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (sp *singleTenantPlugin) DeleteNetNamespace(netns *networkapi.NetNamespace) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (sp *singleTenantPlugin) GetVNID(namespace string) (uint32, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return 0, nil
}
func (sp *singleTenantPlugin) GetNamespaces(vnid uint32) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (sp *singleTenantPlugin) GetMulticastEnabled(vnid uint32) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (sp *singleTenantPlugin) EnsureVNIDRules(vnid uint32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (sp *singleTenantPlugin) SyncVNIDRules() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
