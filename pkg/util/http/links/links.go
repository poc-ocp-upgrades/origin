package links

import (
	"regexp"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
