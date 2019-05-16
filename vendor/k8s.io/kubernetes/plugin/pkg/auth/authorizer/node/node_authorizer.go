package node

import (
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	csiv1alpha1 "k8s.io/csi-api/pkg/apis/csi/v1alpha1"
	"k8s.io/klog"
	coordapi "k8s.io/kubernetes/pkg/apis/coordination"
	api "k8s.io/kubernetes/pkg/apis/core"
	storageapi "k8s.io/kubernetes/pkg/apis/storage"
	"k8s.io/kubernetes/pkg/auth/nodeidentifier"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
	"k8s.io/kubernetes/third_party/forked/gonum/graph"
	"k8s.io/kubernetes/third_party/forked/gonum/graph/traverse"
)

type NodeAuthorizer struct {
	graph      *Graph
	identifier nodeidentifier.NodeIdentifier
	nodeRules  []rbacv1.PolicyRule
	features   utilfeature.FeatureGate
}

func NewAuthorizer(graph *Graph, identifier nodeidentifier.NodeIdentifier, rules []rbacv1.PolicyRule) authorizer.Authorizer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &NodeAuthorizer{graph: graph, identifier: identifier, nodeRules: rules, features: utilfeature.DefaultFeatureGate}
}

var (
	configMapResource   = api.Resource("configmaps")
	secretResource      = api.Resource("secrets")
	pvcResource         = api.Resource("persistentvolumeclaims")
	pvResource          = api.Resource("persistentvolumes")
	vaResource          = storageapi.Resource("volumeattachments")
	svcAcctResource     = api.Resource("serviceaccounts")
	leaseResource       = coordapi.Resource("leases")
	csiNodeInfoResource = csiv1alpha1.Resource("csinodeinfos")
)

func (r *NodeAuthorizer) Authorize(attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodeName, isNode := r.identifier.NodeIdentity(attrs.GetUser())
	if !isNode {
		return authorizer.DecisionNoOpinion, "", nil
	}
	if len(nodeName) == 0 {
		klog.V(2).Infof("NODE DENY: unknown node for user %q", attrs.GetUser().GetName())
		return authorizer.DecisionNoOpinion, fmt.Sprintf("unknown node for user %q", attrs.GetUser().GetName()), nil
	}
	if attrs.IsResourceRequest() {
		requestResource := schema.GroupResource{Group: attrs.GetAPIGroup(), Resource: attrs.GetResource()}
		switch requestResource {
		case secretResource:
			return r.authorizeReadNamespacedObject(nodeName, secretVertexType, attrs)
		case configMapResource:
			return r.authorizeReadNamespacedObject(nodeName, configMapVertexType, attrs)
		case pvcResource:
			if r.features.Enabled(features.ExpandPersistentVolumes) {
				if attrs.GetSubresource() == "status" {
					return r.authorizeStatusUpdate(nodeName, pvcVertexType, attrs)
				}
			}
			return r.authorizeGet(nodeName, pvcVertexType, attrs)
		case pvResource:
			return r.authorizeGet(nodeName, pvVertexType, attrs)
		case vaResource:
			if r.features.Enabled(features.CSIPersistentVolume) {
				return r.authorizeGet(nodeName, vaVertexType, attrs)
			}
			return authorizer.DecisionNoOpinion, fmt.Sprintf("disabled by feature gate %s", features.CSIPersistentVolume), nil
		case svcAcctResource:
			if r.features.Enabled(features.TokenRequest) {
				return r.authorizeCreateToken(nodeName, serviceAccountVertexType, attrs)
			}
			return authorizer.DecisionNoOpinion, fmt.Sprintf("disabled by feature gate %s", features.TokenRequest), nil
		case leaseResource:
			if r.features.Enabled(features.NodeLease) {
				return r.authorizeLease(nodeName, attrs)
			}
			return authorizer.DecisionNoOpinion, fmt.Sprintf("disabled by feature gate %s", features.NodeLease), nil
		case csiNodeInfoResource:
			if r.features.Enabled(features.KubeletPluginsWatcher) && r.features.Enabled(features.CSINodeInfo) {
				return r.authorizeCSINodeInfo(nodeName, attrs)
			}
			return authorizer.DecisionNoOpinion, fmt.Sprintf("disabled by feature gates %s and %s", features.KubeletPluginsWatcher, features.CSINodeInfo), nil
		}
	}
	if rbac.RulesAllow(attrs, r.nodeRules...) {
		return authorizer.DecisionAllow, "", nil
	}
	return authorizer.DecisionNoOpinion, "", nil
}
func (r *NodeAuthorizer) authorizeStatusUpdate(nodeName string, startingType vertexType, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch attrs.GetVerb() {
	case "update", "patch":
	default:
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only get/update/patch this type", nil
	}
	if attrs.GetSubresource() != "status" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only update status subresource", nil
	}
	return r.authorize(nodeName, startingType, attrs)
}
func (r *NodeAuthorizer) authorizeGet(nodeName string, startingType vertexType, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attrs.GetVerb() != "get" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only get individual resources of this type", nil
	}
	if len(attrs.GetSubresource()) > 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "cannot get subresource", nil
	}
	return r.authorize(nodeName, startingType, attrs)
}
func (r *NodeAuthorizer) authorizeReadNamespacedObject(nodeName string, startingType vertexType, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attrs.GetVerb() != "get" && attrs.GetVerb() != "list" && attrs.GetVerb() != "watch" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only read resources of this type", nil
	}
	if len(attrs.GetSubresource()) > 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "cannot read subresource", nil
	}
	if len(attrs.GetNamespace()) == 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only read namespaced object of this type", nil
	}
	return r.authorize(nodeName, startingType, attrs)
}
func (r *NodeAuthorizer) authorize(nodeName string, startingType vertexType, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(attrs.GetName()) == 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "No Object name found", nil
	}
	ok, err := r.hasPathFrom(nodeName, startingType, attrs.GetNamespace(), attrs.GetName())
	if err != nil {
		klog.V(2).Infof("NODE DENY: %v", err)
		return authorizer.DecisionNoOpinion, "no path found to object", nil
	}
	if !ok {
		klog.V(2).Infof("NODE DENY: %q %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "no path found to object", nil
	}
	return authorizer.DecisionAllow, "", nil
}
func (r *NodeAuthorizer) authorizeCreateToken(nodeName string, startingType vertexType, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attrs.GetVerb() != "create" || len(attrs.GetName()) == 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only create tokens for individual service accounts", nil
	}
	if attrs.GetSubresource() != "token" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only create token subresource of serviceaccount", nil
	}
	ok, err := r.hasPathFrom(nodeName, startingType, attrs.GetNamespace(), attrs.GetName())
	if err != nil {
		klog.V(2).Infof("NODE DENY: %v", err)
		return authorizer.DecisionNoOpinion, "no path found to object", nil
	}
	if !ok {
		klog.V(2).Infof("NODE DENY: %q %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "no path found to object", nil
	}
	return authorizer.DecisionAllow, "", nil
}
func (r *NodeAuthorizer) authorizeLease(nodeName string, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	verb := attrs.GetVerb()
	if verb != "get" && verb != "create" && verb != "update" && verb != "patch" && verb != "delete" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only get, create, update, patch, or delete a node lease", nil
	}
	if attrs.GetNamespace() != api.NamespaceNodeLease {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, fmt.Sprintf("can only access leases in the %q system namespace", api.NamespaceNodeLease), nil
	}
	if verb != "create" && attrs.GetName() != nodeName {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only access node lease with the same name as the requesting node", nil
	}
	return authorizer.DecisionAllow, "", nil
}
func (r *NodeAuthorizer) authorizeCSINodeInfo(nodeName string, attrs authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	verb := attrs.GetVerb()
	if verb != "get" && verb != "create" && verb != "update" && verb != "patch" && verb != "delete" {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only get, create, update, patch, or delete a CSINodeInfo", nil
	}
	if len(attrs.GetSubresource()) > 0 {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "cannot authorize CSINodeInfo subresources", nil
	}
	if verb != "create" && attrs.GetName() != nodeName {
		klog.V(2).Infof("NODE DENY: %s %#v", nodeName, attrs)
		return authorizer.DecisionNoOpinion, "can only access CSINodeInfo with the same name as the requesting node", nil
	}
	return authorizer.DecisionAllow, "", nil
}
func (r *NodeAuthorizer) hasPathFrom(nodeName string, startingType vertexType, startingNamespace, startingName string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.graph.lock.RLock()
	defer r.graph.lock.RUnlock()
	nodeVertex, exists := r.graph.getVertex_rlocked(nodeVertexType, "", nodeName)
	if !exists {
		return false, fmt.Errorf("unknown node %q cannot get %s %s/%s", nodeName, vertexTypes[startingType], startingNamespace, startingName)
	}
	startingVertex, exists := r.graph.getVertex_rlocked(startingType, startingNamespace, startingName)
	if !exists {
		return false, fmt.Errorf("node %q cannot get unknown %s %s/%s", nodeName, vertexTypes[startingType], startingNamespace, startingName)
	}
	if r.graph.destinationEdgeIndex[startingVertex.ID()].has(nodeVertex.ID()) {
		return true, nil
	}
	found := false
	traversal := &traverse.VisitingDepthFirst{EdgeFilter: func(edge graph.Edge) bool {
		if destinationEdge, ok := edge.(*destinationEdge); ok {
			if destinationEdge.DestinationID() != nodeVertex.ID() {
				return false
			}
			found = true
		}
		return true
	}}
	traversal.Walk(r.graph.graph, startingVertex, func(n graph.Node) bool {
		if n.ID() == nodeVertex.ID() {
			found = true
		}
		return found
	})
	if !found {
		return false, fmt.Errorf("node %q cannot get %s %s/%s, no path was found", nodeName, vertexTypes[startingType], startingNamespace, startingName)
	}
	return true, nil
}
