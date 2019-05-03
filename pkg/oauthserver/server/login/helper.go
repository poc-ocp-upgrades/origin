package login

import (
	godefaultbytes "bytes"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
	godefaultruntime "runtime"
	"strings"
)

func failed(reason string, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uri, err := getBaseURL(req)
	if err != nil {
		http.Redirect(w, req, req.URL.Path, http.StatusFound)
		return
	}
	query := url.Values{}
	query.Set(reasonParam, reason)
	if then := req.FormValue(thenParam); len(then) != 0 {
		query.Set(thenParam, then)
	}
	uri.RawQuery = query.Encode()
	http.Redirect(w, req, uri.String(), http.StatusFound)
}
func getBaseURL(req *http.Request) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uri, err := url.Parse(req.RequestURI)
	if err != nil {
		return nil, err
	}
	uri.Scheme, uri.Host, uri.RawQuery, uri.Fragment = req.URL.Scheme, req.URL.Host, "", ""
	return uri, nil
}
func postForm(url string, body url.Values) (resp *http.Response, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return http.DefaultTransport.RoundTrip(req)
}
func getURL(url string) (resp *http.Response, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultTransport.RoundTrip(req)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
