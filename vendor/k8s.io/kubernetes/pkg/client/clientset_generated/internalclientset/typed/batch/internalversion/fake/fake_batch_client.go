package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/batch/internalversion"
)

type FakeBatch struct{ *testing.Fake }

func (c *FakeBatch) CronJobs(namespace string) internalversion.CronJobInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeCronJobs{c, namespace}
}
func (c *FakeBatch) Jobs(namespace string) internalversion.JobInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeJobs{c, namespace}
}
func (c *FakeBatch) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
