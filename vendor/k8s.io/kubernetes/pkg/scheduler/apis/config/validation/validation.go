package validation

import (
	goformat "fmt"
	apimachinery "k8s.io/apimachinery/pkg/apis/config/validation"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apiserver "k8s.io/apiserver/pkg/apis/config/validation"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ValidateKubeSchedulerConfiguration(cc *config.KubeSchedulerConfiguration) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apimachinery.ValidateClientConnectionConfiguration(&cc.ClientConnection, field.NewPath("clientConnection"))...)
	allErrs = append(allErrs, ValidateKubeSchedulerLeaderElectionConfiguration(&cc.LeaderElection, field.NewPath("leaderElection"))...)
	if len(cc.SchedulerName) == 0 {
		allErrs = append(allErrs, field.Required(field.NewPath("schedulerName"), ""))
	}
	for _, msg := range validation.IsValidSocketAddr(cc.HealthzBindAddress) {
		allErrs = append(allErrs, field.Invalid(field.NewPath("healthzBindAddress"), cc.HealthzBindAddress, msg))
	}
	for _, msg := range validation.IsValidSocketAddr(cc.MetricsBindAddress) {
		allErrs = append(allErrs, field.Invalid(field.NewPath("metricsBindAddress"), cc.MetricsBindAddress, msg))
	}
	if cc.HardPodAffinitySymmetricWeight < 0 || cc.HardPodAffinitySymmetricWeight > 100 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("hardPodAffinitySymmetricWeight"), cc.HardPodAffinitySymmetricWeight, "not in valid range 0-100"))
	}
	if cc.BindTimeoutSeconds == nil {
		allErrs = append(allErrs, field.Required(field.NewPath("bindTimeoutSeconds"), ""))
	}
	if cc.PercentageOfNodesToScore < 0 || cc.PercentageOfNodesToScore > 100 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("percentageOfNodesToScore"), cc.PercentageOfNodesToScore, "not in valid range 0-100"))
	}
	return allErrs
}
func ValidateKubeSchedulerLeaderElectionConfiguration(cc *config.KubeSchedulerLeaderElectionConfiguration, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allErrs := field.ErrorList{}
	if !cc.LeaderElectionConfiguration.LeaderElect {
		return allErrs
	}
	allErrs = append(allErrs, apiserver.ValidateLeaderElectionConfiguration(&cc.LeaderElectionConfiguration, field.NewPath("leaderElectionConfiguration"))...)
	if len(cc.LockObjectNamespace) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("lockObjectNamespace"), ""))
	}
	if len(cc.LockObjectName) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("lockObjectName"), ""))
	}
	return allErrs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
