package internalversion

import (
 "k8s.io/apimachinery/pkg/types"
 api "k8s.io/kubernetes/pkg/apis/core"
)

type NodeExpansion interface {
 PatchStatus(nodeName string, data []byte) (*api.Node, error)
}

func (c *nodes) PatchStatus(nodeName string, data []byte) (*api.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result := &api.Node{}
 err := c.client.Patch(types.StrategicMergePatchType).Resource("nodes").Name(nodeName).SubResource("status").Body(data).Do().Into(result)
 return result, err
}
