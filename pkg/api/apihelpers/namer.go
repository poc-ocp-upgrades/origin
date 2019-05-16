package apihelpers

import (
	"fmt"
	"hash/fnv"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
)

func GetName(base, suffix string, maxLength int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if maxLength <= 0 {
		return ""
	}
	name := fmt.Sprintf("%s-%s", base, suffix)
	if len(name) <= maxLength {
		return name
	}
	baseLength := maxLength - 10 - len(suffix)
	if baseLength < 0 {
		prefix := base[0:min(len(base), max(0, maxLength-9))]
		shortName := fmt.Sprintf("%s-%s", prefix, hash(name))
		return shortName[:min(maxLength, len(shortName))]
	}
	prefix := base[0:baseLength]
	return fmt.Sprintf("%s-%s-%s", prefix, hash(base), suffix)
}
func GetPodName(base, suffix string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GetName(base, suffix, kvalidation.DNS1123SubdomainMaxLength)
}
func GetConfigMapName(base, suffix string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GetName(base, suffix, kvalidation.DNS1123SubdomainMaxLength)
}
func max(a, b int) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if b > a {
		return b
	}
	return a
}
func min(a, b int) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if b < a {
		return b
	}
	return a
}
func hash(s string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hash := fnv.New32a()
	hash.Write([]byte(s))
	intHash := hash.Sum32()
	result := fmt.Sprintf("%08x", intHash)
	return result
}
