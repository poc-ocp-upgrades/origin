package scheduling

import (
	"fmt"
	goformat "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var systemPriorityClasses = []*PriorityClass{{ObjectMeta: metav1.ObjectMeta{Name: SystemNodeCritical}, Value: SystemCriticalPriority + 1000, Description: "Used for system critical pods that must not be moved from their current node."}, {ObjectMeta: metav1.ObjectMeta{Name: SystemClusterCritical}, Value: SystemCriticalPriority, Description: "Used for system critical pods that must run in the cluster, but can be moved to another node if necessary."}}

func SystemPriorityClasses() []*PriorityClass {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return systemPriorityClasses
}
func IsKnownSystemPriorityClass(pc *PriorityClass) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, spc := range systemPriorityClasses {
		if spc.Name == pc.Name {
			if spc.Value != pc.Value {
				return false, fmt.Errorf("value of %v PriorityClass must be %v", spc.Name, spc.Value)
			}
			if spc.GlobalDefault != pc.GlobalDefault {
				return false, fmt.Errorf("globalDefault of %v PriorityClass must be %v", spc.Name, spc.GlobalDefault)
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("%v is not a known system priority class", pc.Name)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
