package scheduler

import (
 "fmt"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/util/sets"
 clientset "k8s.io/client-go/kubernetes"
 corelisters "k8s.io/client-go/listers/core/v1"
 "k8s.io/kubernetes/pkg/scheduler/algorithm"
 schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
 "k8s.io/kubernetes/pkg/scheduler/factory"
 internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
 "k8s.io/kubernetes/pkg/scheduler/util"
)

type FakeConfigurator struct{ Config *factory.Config }

func (fc *FakeConfigurator) GetPredicateMetadataProducer() (algorithm.PredicateMetadataProducer, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("not implemented")
}
func (fc *FakeConfigurator) GetPredicates(predicateKeys sets.String) (map[string]algorithm.FitPredicate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("not implemented")
}
func (fc *FakeConfigurator) GetHardPodAffinitySymmetricWeight() int32 {
 _logClusterCodePath()
 defer _logClusterCodePath()
 panic("not implemented")
}
func (fc *FakeConfigurator) MakeDefaultErrorFunc(backoff *util.PodBackoff, podQueue internalqueue.SchedulingQueue) func(pod *v1.Pod, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fc *FakeConfigurator) GetNodeLister() corelisters.NodeLister {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fc *FakeConfigurator) GetClient() clientset.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fc *FakeConfigurator) GetScheduledPodLister() corelisters.PodLister {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fc *FakeConfigurator) Create() (*factory.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fc.Config, nil
}
func (fc *FakeConfigurator) CreateFromProvider(providerName string) (*factory.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fc.Config, nil
}
func (fc *FakeConfigurator) CreateFromConfig(policy schedulerapi.Policy) (*factory.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fc.Config, nil
}
func (fc *FakeConfigurator) CreateFromKeys(predicateKeys, priorityKeys sets.String, extenders []algorithm.SchedulerExtender) (*factory.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fc.Config, nil
}
