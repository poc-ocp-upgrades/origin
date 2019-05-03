package registrytest

import (
 "context"
 "sync"
 "k8s.io/apimachinery/pkg/api/errors"
 metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/watch"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type NodeRegistry struct {
 Err   error
 Node  string
 Nodes api.NodeList
 sync.Mutex
}

func MakeNodeList(nodes []string, nodeResources api.NodeResources) *api.NodeList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 list := api.NodeList{Items: make([]api.Node, len(nodes))}
 for i := range nodes {
  list.Items[i].Name = nodes[i]
  list.Items[i].Status.Capacity = nodeResources.Capacity
 }
 return &list
}
func NewNodeRegistry(nodes []string, nodeResources api.NodeResources) *NodeRegistry {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &NodeRegistry{Nodes: *MakeNodeList(nodes, nodeResources)}
}
func (r *NodeRegistry) SetError(err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 r.Err = err
}
func (r *NodeRegistry) ListNodes(ctx context.Context, options *metainternalversion.ListOptions) (*api.NodeList, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 return &r.Nodes, r.Err
}
func (r *NodeRegistry) CreateNode(ctx context.Context, node *api.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 r.Node = node.Name
 r.Nodes.Items = append(r.Nodes.Items, *node)
 return r.Err
}
func (r *NodeRegistry) UpdateNode(ctx context.Context, node *api.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 for i, item := range r.Nodes.Items {
  if item.Name == node.Name {
   r.Nodes.Items[i] = *node
   return r.Err
  }
 }
 return r.Err
}
func (r *NodeRegistry) GetNode(ctx context.Context, nodeID string, options *metav1.GetOptions) (*api.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 if r.Err != nil {
  return nil, r.Err
 }
 for _, node := range r.Nodes.Items {
  if node.Name == nodeID {
   return &node, nil
  }
 }
 return nil, errors.NewNotFound(api.Resource("nodes"), nodeID)
}
func (r *NodeRegistry) DeleteNode(ctx context.Context, nodeID string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 r.Lock()
 defer r.Unlock()
 var newList []api.Node
 for _, node := range r.Nodes.Items {
  if node.Name != nodeID {
   newList = append(newList, api.Node{ObjectMeta: metav1.ObjectMeta{Name: node.Name}})
  }
 }
 r.Nodes.Items = newList
 return r.Err
}
func (r *NodeRegistry) WatchNodes(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, r.Err
}
