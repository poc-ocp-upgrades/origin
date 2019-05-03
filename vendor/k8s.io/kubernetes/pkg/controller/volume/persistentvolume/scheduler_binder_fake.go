package persistentvolume

import "k8s.io/api/core/v1"

type FakeVolumeBinderConfig struct {
 AllBound             bool
 FindUnboundSatsified bool
 FindBoundSatsified   bool
 FindErr              error
 AssumeErr            error
 BindErr              error
}

func NewFakeVolumeBinder(config *FakeVolumeBinderConfig) *FakeVolumeBinder {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &FakeVolumeBinder{config: config}
}

type FakeVolumeBinder struct {
 config       *FakeVolumeBinderConfig
 AssumeCalled bool
 BindCalled   bool
}

func (b *FakeVolumeBinder) FindPodVolumes(pod *v1.Pod, node *v1.Node) (unboundVolumesSatisfied, boundVolumesSatsified bool, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return b.config.FindUnboundSatsified, b.config.FindBoundSatsified, b.config.FindErr
}
func (b *FakeVolumeBinder) AssumePodVolumes(assumedPod *v1.Pod, nodeName string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 b.AssumeCalled = true
 return b.config.AllBound, b.config.AssumeErr
}
func (b *FakeVolumeBinder) BindPodVolumes(assumedPod *v1.Pod) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 b.BindCalled = true
 return b.config.BindErr
}
func (b *FakeVolumeBinder) GetBindingsCache() PodBindingCache {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
