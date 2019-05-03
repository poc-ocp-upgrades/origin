package fake

import (
 rest "k8s.io/client-go/rest"
 testing "k8s.io/client-go/testing"
 internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/storage/internalversion"
)

type FakeStorage struct{ *testing.Fake }

func (c *FakeStorage) StorageClasses() internalversion.StorageClassInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeStorageClasses{c}
}
func (c *FakeStorage) VolumeAttachments() internalversion.VolumeAttachmentInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeVolumeAttachments{c}
}
func (c *FakeStorage) RESTClient() rest.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret *rest.RESTClient
 return ret
}
