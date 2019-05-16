package podsecuritypolicy

import (
	"fmt"
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/util/errors"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/apparmor"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/capabilities"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/group"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/seccomp"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/selinux"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/sysctl"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/user"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type simpleStrategyFactory struct{}

var _ StrategyFactory = &simpleStrategyFactory{}

func NewSimpleStrategyFactory() StrategyFactory {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &simpleStrategyFactory{}
}
func (f *simpleStrategyFactory) CreateStrategies(psp *policy.PodSecurityPolicy, namespace string) (*ProviderStrategies, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	userStrat, err := createUserStrategy(&psp.Spec.RunAsUser)
	if err != nil {
		errs = append(errs, err)
	}
	var groupStrat group.GroupStrategy
	if utilfeature.DefaultFeatureGate.Enabled(features.RunAsGroup) {
		groupStrat, err = createRunAsGroupStrategy(psp.Spec.RunAsGroup)
		if err != nil {
			errs = append(errs, err)
		}
	}
	seLinuxStrat, err := createSELinuxStrategy(&psp.Spec.SELinux)
	if err != nil {
		errs = append(errs, err)
	}
	appArmorStrat, err := createAppArmorStrategy(psp)
	if err != nil {
		errs = append(errs, err)
	}
	seccompStrat, err := createSeccompStrategy(psp)
	if err != nil {
		errs = append(errs, err)
	}
	fsGroupStrat, err := createFSGroupStrategy(&psp.Spec.FSGroup)
	if err != nil {
		errs = append(errs, err)
	}
	supGroupStrat, err := createSupplementalGroupStrategy(&psp.Spec.SupplementalGroups)
	if err != nil {
		errs = append(errs, err)
	}
	capStrat, err := createCapabilitiesStrategy(psp.Spec.DefaultAddCapabilities, psp.Spec.RequiredDropCapabilities, psp.Spec.AllowedCapabilities)
	if err != nil {
		errs = append(errs, err)
	}
	sysctlsStrat := createSysctlsStrategy(sysctl.SafeSysctlWhitelist(), psp.Spec.AllowedUnsafeSysctls, psp.Spec.ForbiddenSysctls)
	if len(errs) > 0 {
		return nil, errors.NewAggregate(errs)
	}
	strategies := &ProviderStrategies{RunAsUserStrategy: userStrat, RunAsGroupStrategy: groupStrat, SELinuxStrategy: seLinuxStrat, AppArmorStrategy: appArmorStrat, FSGroupStrategy: fsGroupStrat, SupplementalGroupStrategy: supGroupStrat, CapabilitiesStrategy: capStrat, SeccompStrategy: seccompStrat, SysctlsStrategy: sysctlsStrat}
	return strategies, nil
}
func createUserStrategy(opts *policy.RunAsUserStrategyOptions) (user.RunAsUserStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch opts.Rule {
	case policy.RunAsUserStrategyMustRunAs:
		return user.NewMustRunAs(opts)
	case policy.RunAsUserStrategyMustRunAsNonRoot:
		return user.NewRunAsNonRoot(opts)
	case policy.RunAsUserStrategyRunAsAny:
		return user.NewRunAsAny(opts)
	default:
		return nil, fmt.Errorf("Unrecognized RunAsUser strategy type %s", opts.Rule)
	}
}
func createRunAsGroupStrategy(opts *policy.RunAsGroupStrategyOptions) (group.GroupStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if opts == nil {
		return group.NewRunAsAny()
	}
	switch opts.Rule {
	case policy.RunAsGroupStrategyMustRunAs:
		return group.NewMustRunAs(opts.Ranges)
	case policy.RunAsGroupStrategyRunAsAny:
		return group.NewRunAsAny()
	case policy.RunAsGroupStrategyMayRunAs:
		return group.NewMayRunAs(opts.Ranges)
	default:
		return nil, fmt.Errorf("Unrecognized RunAsGroup strategy type %s", opts.Rule)
	}
}
func createSELinuxStrategy(opts *policy.SELinuxStrategyOptions) (selinux.SELinuxStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch opts.Rule {
	case policy.SELinuxStrategyMustRunAs:
		return selinux.NewMustRunAs(opts)
	case policy.SELinuxStrategyRunAsAny:
		return selinux.NewRunAsAny(opts)
	default:
		return nil, fmt.Errorf("Unrecognized SELinuxContext strategy type %s", opts.Rule)
	}
}
func createAppArmorStrategy(psp *policy.PodSecurityPolicy) (apparmor.Strategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apparmor.NewStrategy(psp.Annotations), nil
}
func createSeccompStrategy(psp *policy.PodSecurityPolicy) (seccomp.Strategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return seccomp.NewStrategy(psp.Annotations), nil
}
func createFSGroupStrategy(opts *policy.FSGroupStrategyOptions) (group.GroupStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch opts.Rule {
	case policy.FSGroupStrategyRunAsAny:
		return group.NewRunAsAny()
	case policy.FSGroupStrategyMayRunAs:
		return group.NewMayRunAs(opts.Ranges)
	case policy.FSGroupStrategyMustRunAs:
		return group.NewMustRunAs(opts.Ranges)
	default:
		return nil, fmt.Errorf("Unrecognized FSGroup strategy type %s", opts.Rule)
	}
}
func createSupplementalGroupStrategy(opts *policy.SupplementalGroupsStrategyOptions) (group.GroupStrategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch opts.Rule {
	case policy.SupplementalGroupsStrategyRunAsAny:
		return group.NewRunAsAny()
	case policy.SupplementalGroupsStrategyMayRunAs:
		return group.NewMayRunAs(opts.Ranges)
	case policy.SupplementalGroupsStrategyMustRunAs:
		return group.NewMustRunAs(opts.Ranges)
	default:
		return nil, fmt.Errorf("Unrecognized SupplementalGroups strategy type %s", opts.Rule)
	}
}
func createCapabilitiesStrategy(defaultAddCaps, requiredDropCaps, allowedCaps []corev1.Capability) (capabilities.Strategy, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return capabilities.NewDefaultCapabilities(defaultAddCaps, requiredDropCaps, allowedCaps)
}
func createSysctlsStrategy(safeWhitelist, allowedUnsafeSysctls, forbiddenSysctls []string) sysctl.SysctlsStrategy {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return sysctl.NewMustMatchPatterns(safeWhitelist, allowedUnsafeSysctls, forbiddenSysctls)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
