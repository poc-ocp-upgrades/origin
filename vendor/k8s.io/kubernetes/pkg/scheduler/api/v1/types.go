package v1

import (
	gojson "encoding/json"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	restclient "k8s.io/client-go/rest"
	"time"
)

type Policy struct {
	metav1.TypeMeta                `json:",inline"`
	Predicates                     []PredicatePolicy `json:"predicates"`
	Priorities                     []PriorityPolicy  `json:"priorities"`
	ExtenderConfigs                []ExtenderConfig  `json:"extenders"`
	HardPodAffinitySymmetricWeight int               `json:"hardPodAffinitySymmetricWeight"`
	AlwaysCheckAllPredicates       bool              `json:"alwaysCheckAllPredicates"`
}
type PredicatePolicy struct {
	Name     string             `json:"name"`
	Argument *PredicateArgument `json:"argument"`
}
type PriorityPolicy struct {
	Name     string            `json:"name"`
	Weight   int               `json:"weight"`
	Argument *PriorityArgument `json:"argument"`
}
type PredicateArgument struct {
	ServiceAffinity *ServiceAffinity `json:"serviceAffinity"`
	LabelsPresence  *LabelsPresence  `json:"labelsPresence"`
}
type PriorityArgument struct {
	ServiceAntiAffinity               *ServiceAntiAffinity               `json:"serviceAntiAffinity"`
	LabelPreference                   *LabelPreference                   `json:"labelPreference"`
	RequestedToCapacityRatioArguments *RequestedToCapacityRatioArguments `json:"requestedToCapacityRatioArguments"`
}
type ServiceAffinity struct {
	Labels []string `json:"labels"`
}
type LabelsPresence struct {
	Labels   []string `json:"labels"`
	Presence bool     `json:"presence"`
}
type ServiceAntiAffinity struct {
	Label string `json:"label"`
}
type LabelPreference struct {
	Label    string `json:"label"`
	Presence bool   `json:"presence"`
}
type RequestedToCapacityRatioArguments struct {
	UtilizationShape []UtilizationShapePoint `json:"shape"`
}
type UtilizationShapePoint struct {
	Utilization int `json:"utilization"`
	Score       int `json:"score"`
}
type ExtenderManagedResource struct {
	Name               apiv1.ResourceName `json:"name,casttype=ResourceName"`
	IgnoredByScheduler bool               `json:"ignoredByScheduler,omitempty"`
}
type ExtenderConfig struct {
	URLPrefix        string `json:"urlPrefix"`
	FilterVerb       string `json:"filterVerb,omitempty"`
	PreemptVerb      string `json:"preemptVerb,omitempty"`
	PrioritizeVerb   string `json:"prioritizeVerb,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	BindVerb         string
	EnableHTTPS      bool                        `json:"enableHttps,omitempty"`
	TLSConfig        *restclient.TLSClientConfig `json:"tlsConfig,omitempty"`
	HTTPTimeout      time.Duration               `json:"httpTimeout,omitempty"`
	NodeCacheCapable bool                        `json:"nodeCacheCapable,omitempty"`
	ManagedResources []ExtenderManagedResource   `json:"managedResources,omitempty"`
	Ignorable        bool                        `json:"ignorable,omitempty"`
}
type caseInsensitiveExtenderConfig *ExtenderConfig

func (t *ExtenderConfig) UnmarshalJSON(b []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return gojson.Unmarshal(b, caseInsensitiveExtenderConfig(t))
}

type ExtenderArgs struct {
	Pod       *apiv1.Pod      `json:"pod"`
	Nodes     *apiv1.NodeList `json:"nodes,omitempty"`
	NodeNames *[]string       `json:"nodenames,omitempty"`
}
type ExtenderPreemptionResult struct {
	NodeNameToMetaVictims map[string]*MetaVictims `json:"nodeNameToMetaVictims,omitempty"`
}
type ExtenderPreemptionArgs struct {
	Pod                   *apiv1.Pod              `json:"pod"`
	NodeNameToVictims     map[string]*Victims     `json:"nodeToVictims,omitempty"`
	NodeNameToMetaVictims map[string]*MetaVictims `json:"nodeNameToMetaVictims,omitempty"`
}
type Victims struct {
	Pods             []*apiv1.Pod `json:"pods"`
	NumPDBViolations int          `json:"numPDBViolations"`
}
type MetaPod struct {
	UID string `json:"uid"`
}
type MetaVictims struct {
	Pods             []*MetaPod `json:"pods"`
	NumPDBViolations int        `json:"numPDBViolations"`
}
type FailedNodesMap map[string]string
type ExtenderFilterResult struct {
	Nodes       *apiv1.NodeList `json:"nodes,omitempty"`
	NodeNames   *[]string       `json:"nodenames,omitempty"`
	FailedNodes FailedNodesMap  `json:"failedNodes,omitempty"`
	Error       string          `json:"error,omitempty"`
}
type ExtenderBindingArgs struct {
	PodName      string
	PodNamespace string
	PodUID       types.UID
	Node         string
}
type ExtenderBindingResult struct{ Error string }
type HostPriority struct {
	Host  string `json:"host"`
	Score int    `json:"score"`
}
type HostPriorityList []HostPriority

func (h HostPriorityList) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(h)
}
func (h HostPriorityList) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if h[i].Score == h[j].Score {
		return h[i].Host < h[j].Host
	}
	return h[i].Score < h[j].Score
}
func (h HostPriorityList) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h[i], h[j] = h[j], h[i]
}
