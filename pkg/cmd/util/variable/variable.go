package variable

import (
	"fmt"
	"os"
	"strings"
	"github.com/openshift/origin/pkg/version"
)

type KeyFunc func(key string) (string, bool)

func Expand(s string, fns ...KeyFunc) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	val, _ := ExpandStrict(s, append(fns, Empty)...)
	return val
}
func ExpandStrict(s string, fns ...KeyFunc) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	unmatched := []string{}
	result := os.Expand(s, func(key string) string {
		for _, fn := range fns {
			val, ok := fn(key)
			if !ok {
				continue
			}
			return val
		}
		unmatched = append(unmatched, key)
		return ""
	})
	switch len(unmatched) {
	case 0:
		return result, nil
	case 1:
		return "", fmt.Errorf("the key %q in %q is not recognized", unmatched[0], s)
	default:
		return "", fmt.Errorf("multiple keys in %q were not recognized: %s", s, strings.Join(unmatched, ", "))
	}
}
func Empty(s string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", true
}
func Identity(key string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("${%s}", key), true
}
func Versions(key string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch key {
	case "shortcommit":
		s := overrideVersion.GitCommit
		if len(s) > 7 {
			s = s[:7]
		}
		return s, true
	case "version":
		s := lastSemanticVersionWithoutModifiers(overrideVersion.GitVersion)
		return s, true
	default:
		return "", false
	}
}
func Env(key string) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return os.Getenv(key), true
}

var overrideVersion = version.Get()

func lastSemanticVersionWithoutModifiers(version string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.Split(version, "-")
	return strings.Split(parts[0], "+")[0]
}
