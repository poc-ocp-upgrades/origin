package features

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strconv"
	"strings"
	gotime "time"
)

const (
	CoreDNS              = "CoreDNS"
	DynamicKubeletConfig = "DynamicKubeletConfig"
)

var coreDNSMessage = "featureGates:CoreDNS has been removed in v1.13\n" + "\tUse kubeadm-config to select which DNS addon to install."
var InitFeatureGates = FeatureList{CoreDNS: {FeatureSpec: utilfeature.FeatureSpec{Default: true, PreRelease: utilfeature.Deprecated}, HiddenInHelpText: true, DeprecationMessage: coreDNSMessage}}

type Feature struct {
	utilfeature.FeatureSpec
	MinimumVersion     *version.Version
	HiddenInHelpText   bool
	DeprecationMessage string
}
type FeatureList map[string]Feature

func ValidateVersion(allFeatures FeatureList, requestedFeatures map[string]bool, requestedVersion string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if requestedVersion == "" {
		return nil
	}
	parsedExpVersion, err := version.ParseSemantic(requestedVersion)
	if err != nil {
		return errors.Wrapf(err, "error parsing version %s", requestedVersion)
	}
	for k := range requestedFeatures {
		if minVersion := allFeatures[k].MinimumVersion; minVersion != nil {
			if !parsedExpVersion.AtLeast(minVersion) {
				return errors.Errorf("the requested Kubernetes version (%s) is incompatible with the %s feature gate, which needs %s as a minimum", requestedVersion, k, minVersion)
			}
		}
	}
	return nil
}
func Enabled(featureList map[string]bool, featureName string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if enabled, ok := featureList[string(featureName)]; ok {
		return enabled
	}
	return InitFeatureGates[string(featureName)].Default
}
func Supports(featureList FeatureList, featureName string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for k, v := range featureList {
		if featureName == string(k) {
			return v.PreRelease != utilfeature.Deprecated
		}
	}
	return false
}
func Keys(featureList FeatureList) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var list []string
	for k := range featureList {
		list = append(list, string(k))
	}
	return list
}
func KnownFeatures(f *FeatureList) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var known []string
	for k, v := range *f {
		if v.HiddenInHelpText {
			continue
		}
		pre := ""
		if v.PreRelease != utilfeature.GA {
			pre = fmt.Sprintf("%s - ", v.PreRelease)
		}
		known = append(known, fmt.Sprintf("%s=true|false (%sdefault=%t)", k, pre, v.Default))
	}
	sort.Strings(known)
	return known
}
func NewFeatureGate(f *FeatureList, value string) (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	featureGate := map[string]bool{}
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) != 2 {
			return nil, errors.Errorf("missing bool value for feature-gate key:%s", s)
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		featureSpec, ok := (*f)[k]
		if !ok {
			return nil, errors.Errorf("unrecognized feature-gate key: %s", k)
		}
		if featureSpec.PreRelease == utilfeature.Deprecated {
			return nil, errors.Errorf("feature-gate key is deprecated: %s", k)
		}
		boolValue, err := strconv.ParseBool(v)
		if err != nil {
			return nil, errors.Errorf("invalid value %v for feature-gate key: %s, use true|false instead", v, k)
		}
		featureGate[k] = boolValue
	}
	return featureGate, nil
}
func CheckDeprecatedFlags(f *FeatureList, features map[string]bool) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	deprecatedMsg := map[string]string{}
	for k := range features {
		featureSpec, ok := (*f)[k]
		if !ok {
			deprecatedMsg[k] = fmt.Sprintf("Unknown feature gate flag: %s", k)
		}
		if featureSpec.PreRelease == utilfeature.Deprecated {
			if _, ok := deprecatedMsg[k]; !ok {
				deprecatedMsg[k] = featureSpec.DeprecationMessage
			}
		}
	}
	return deprecatedMsg
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
