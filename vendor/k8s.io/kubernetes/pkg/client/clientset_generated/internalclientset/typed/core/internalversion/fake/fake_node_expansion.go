package fake

import (
 types "k8s.io/apimachinery/pkg/types"
 core "k8s.io/client-go/testing"
 api "k8s.io/kubernetes/pkg/apis/core"
)

func (c *FakeNodes) PatchStatus(nodeName string, data []byte) (*api.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pt := types.StrategicMergePatchType
 obj, err := c.Fake.Invokes(core.NewRootPatchSubresourceAction(nodesResource, nodeName, pt, data, "status"), &api.Node{})
 if obj == nil {
  return nil, err
 }
 return obj.(*api.Node), err
}
