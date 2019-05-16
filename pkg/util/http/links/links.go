package links

import (
	goformat "fmt"
	goos "os"
	"regexp"
	godefaultruntime "runtime"
	gotime "time"
)

var linkRegex = regexp.MustCompile(`\<(.+?)\>\s*;\s*rel="(.+?)"(?:\s*,\s*)?`)

func ParseLinks(header string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
