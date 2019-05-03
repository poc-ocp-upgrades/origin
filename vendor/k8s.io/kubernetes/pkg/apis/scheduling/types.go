package scheduling

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
 DefaultPriorityWhenNoDefaultClassExists = 0
 HighestUserDefinablePriority            = int32(1000000000)
 SystemCriticalPriority                  = 2 * HighestUserDefinablePriority
 SystemPriorityClassPrefix               = "system-"
 SystemClusterCritical                   = SystemPriorityClassPrefix + "cluster-critical"
 SystemNodeCritical                      = SystemPriorityClassPrefix + "node-critical"
)

type PriorityClass struct {
 metav1.TypeMeta
 metav1.ObjectMeta
 Value         int32
 GlobalDefault bool
 Description   string
}
type PriorityClassList struct {
 metav1.TypeMeta
 metav1.ListMeta
 Items []PriorityClass
}
