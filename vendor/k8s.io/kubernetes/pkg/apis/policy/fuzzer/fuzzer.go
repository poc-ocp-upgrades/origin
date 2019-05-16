package fuzzer

import (
	goformat "fmt"
	fuzz "github.com/google/gofuzz"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kubernetes/pkg/apis/policy"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{func(s *policy.PodDisruptionBudgetStatus, c fuzz.Continue) {
		c.FuzzNoCustom(s)
		s.PodDisruptionsAllowed = int32(c.Rand.Intn(2))
	}, func(psp *policy.PodSecurityPolicySpec, c fuzz.Continue) {
		c.FuzzNoCustom(psp)
		runAsUserRules := []policy.RunAsUserStrategy{policy.RunAsUserStrategyMustRunAsNonRoot, policy.RunAsUserStrategyMustRunAs, policy.RunAsUserStrategyRunAsAny}
		psp.RunAsUser.Rule = runAsUserRules[c.Rand.Intn(len(runAsUserRules))]
		runAsGroupRules := []policy.RunAsGroupStrategy{policy.RunAsGroupStrategyMustRunAs, policy.RunAsGroupStrategyRunAsAny, policy.RunAsGroupStrategyMayRunAs}
		psp.RunAsGroup = &policy.RunAsGroupStrategyOptions{}
		psp.RunAsGroup.Rule = runAsGroupRules[c.Rand.Intn(len(runAsGroupRules))]
		seLinuxRules := []policy.SELinuxStrategy{policy.SELinuxStrategyMustRunAs, policy.SELinuxStrategyRunAsAny}
		psp.SELinux.Rule = seLinuxRules[c.Rand.Intn(len(seLinuxRules))]
		supplementalGroupsRules := []policy.SupplementalGroupsStrategyType{policy.SupplementalGroupsStrategyRunAsAny, policy.SupplementalGroupsStrategyMayRunAs, policy.SupplementalGroupsStrategyMustRunAs}
		psp.SupplementalGroups.Rule = supplementalGroupsRules[c.Rand.Intn(len(supplementalGroupsRules))]
		fsGroupRules := []policy.FSGroupStrategyType{policy.FSGroupStrategyMustRunAs, policy.FSGroupStrategyMayRunAs, policy.FSGroupStrategyRunAsAny}
		psp.FSGroup.Rule = fsGroupRules[c.Rand.Intn(len(fsGroupRules))]
	}}
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
