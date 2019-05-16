package scope

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

func Add(has []string, new []string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sorted := sortAndCopy(has)
	for _, s := range new {
		i := sort.SearchStrings(sorted, s)
		if i == len(sorted) {
			sorted = append(sorted, s)
		} else if sorted[i] != s {
			sorted = append(sorted, "")
			copy(sorted[i+1:], sorted[i:])
			sorted[i] = s
		}
	}
	return sorted
}
func Split(scope string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scope = strings.TrimSpace(scope)
	if scope == "" {
		return []string{}
	}
	return strings.Split(scope, " ")
}
func Join(scopes []string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.Join(scopes, " ")
}
func Covers(has, requested []string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(requested) == 0 && len(has) > 0 {
		return false
	}
	has, requested = sortAndCopy(has), sortAndCopy(requested)
NextRequested:
	for i := range requested {
		for j := range has {
			if has[j] == requested[i] {
				continue NextRequested
			}
		}
		return false
	}
	return true
}
func sortAndCopy(arr []string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	sort.Sort(sort.StringSlice(newArr))
	return newArr
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
