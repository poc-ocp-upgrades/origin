package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/autoscaling/internalversion"
)

type FakeAutoscaling struct{ *testing.Fake }

func (c *FakeAutoscaling) HorizontalPodAutoscalers(namespace string) internalversion.HorizontalPodAutoscalerInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeHorizontalPodAutoscalers{c, namespace}
}
func (c *FakeAutoscaling) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
