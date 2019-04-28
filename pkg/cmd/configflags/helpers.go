package configflags

import (
	"fmt"
	"sort"
	"strings"
)

func ArgsWithPrefix(args map[string][]string, prefix string) map[string][]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	filtered := map[string][]string{}
	for key, slice := range args {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		for _, val := range slice {
			filtered[key] = append(filtered[key], val)
		}
	}
	return filtered
}
func SetIfUnset(cmdLineArgs map[string][]string, key string, value ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := cmdLineArgs[key]; !ok {
		cmdLineArgs[key] = value
	}
}
func ToFlagSlice(args map[string][]string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var keys []string
	for key := range args {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var flags []string
	for _, key := range keys {
		for _, token := range args[key] {
			flags = append(flags, fmt.Sprintf("--%s=%v", key, token))
		}
	}
	return flags
}
