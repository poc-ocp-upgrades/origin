package api

import (
 "time"
 "k8s.io/api/core/v1"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/types"
 restclient "k8s.io/client-go/rest"
)

const (
 MaxUint                         = ^uint(0)
 MaxInt                          = int(MaxUint >> 1)
 MaxTotalPriority                = MaxInt
 MaxPriority                     = 10
 MaxWeight                       = MaxInt / MaxPriority
 DefaultPercentageOfNodesToScore = 50
)

type Policy struct {
 metav1.TypeMeta
 Predicates                     []PredicatePolicy
 Priorities                     []PriorityPolicy
 ExtenderConfigs                []ExtenderConfig
 HardPodAffinitySymmetricWeight int32
 AlwaysCheckAllPredicates       bool
}
type PredicatePolicy struct {
 Name     string
 Argument *PredicateArgument
}
type PriorityPolicy struct {
 Name     string
 Weight   int
 Argument *PriorityArgument
}
type PredicateArgument struct {
 ServiceAffinity *ServiceAffinity
 LabelsPresence  *LabelsPresence
}
type PriorityArgument struct {
 ServiceAntiAffinity               *ServiceAntiAffinity
 LabelPreference                   *LabelPreference
 RequestedToCapacityRatioArguments *RequestedToCapacityRatioArguments
}
type ServiceAffinity struct{ Labels []string }
type LabelsPresence struct {
 Labels   []string
 Presence bool
}
type ServiceAntiAffinity struct{ Label string }
type LabelPreference struct {
 Label    string
 Presence bool
}
type RequestedToCapacityRatioArguments struct{ UtilizationShape []UtilizationShapePoint }
type UtilizationShapePoint struct {
 Utilization int
 Score       int
}
type ExtenderManagedResource struct {
 Name               v1.ResourceName
 IgnoredByScheduler bool
}
type ExtenderConfig struct {
 URLPrefix        string
 FilterVerb       string
 PreemptVerb      string
 PrioritizeVerb   string
 Weight           int
 BindVerb         string
 EnableHTTPS      bool
 TLSConfig        *restclient.TLSClientConfig
 HTTPTimeout      time.Duration
 NodeCacheCapable bool
 ManagedResources []ExtenderManagedResource
 Ignorable        bool
}
type ExtenderPreemptionResult struct{ NodeNameToMetaVictims map[string]*MetaVictims }
type ExtenderPreemptionArgs struct {
 Pod                   *v1.Pod
 NodeNameToVictims     map[string]*Victims
 NodeNameToMetaVictims map[string]*MetaVictims
}
type Victims struct {
 Pods             []*v1.Pod
 NumPDBViolations int
}
type MetaPod struct{ UID string }
type MetaVictims struct {
 Pods             []*MetaPod
 NumPDBViolations int
}
type ExtenderArgs struct {
 Pod       *v1.Pod
 Nodes     *v1.NodeList
 NodeNames *[]string
}
type FailedNodesMap map[string]string
type ExtenderFilterResult struct {
 Nodes       *v1.NodeList
 NodeNames   *[]string
 FailedNodes FailedNodesMap
 Error       string
}
type ExtenderBindingArgs struct {
 PodName      string
 PodNamespace string
 PodUID       types.UID
 Node         string
}
type ExtenderBindingResult struct{ Error string }
type HostPriority struct {
 Host  string
 Score int
}
type HostPriorityList []HostPriority

func (h HostPriorityList) Len() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return len(h)
}
func (h HostPriorityList) Less(i, j int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if h[i].Score == h[j].Score {
  return h[i].Host < h[j].Host
 }
 return h[i].Score < h[j].Score
}
func (h HostPriorityList) Swap(i, j int) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 h[i], h[j] = h[j], h[i]
}
