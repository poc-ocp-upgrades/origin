package apidocs

import (
	"strings"
)

type byPathName []operation

func (o byPathName) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o byPathName) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o byPathName) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[i].PathName < o[j].PathName
}

type bySubresource []operation

func (o bySubresource) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o bySubresource) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o bySubresource) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[i].Subresource() < o[j].Subresource()
}

type byNamespaced []operation

func (o byNamespaced) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o byNamespaced) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o byNamespaced) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[j].Namespaced() && !o[i].Namespaced()
}

type byProxy []operation

func (o byProxy) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o byProxy) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o byProxy) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[j].IsProxy() && !o[i].IsProxy()
}

type byPlural []operation

func (o byPlural) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o byPlural) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o byPlural) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[j].Plural() && !o[i].Plural()
}

type byOperationVerb []operation

var byOperationVerbOrder = map[string]int{"Options": 1, "Create": 2, "Head": 3, "Get": 4, "Watch": 5, "Update": 6, "Patch": 7, "Delete": 8}

func (o byOperationVerb) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o byOperationVerb) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o byOperationVerb) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return byOperationVerbOrder[o[i].Verb()] < byOperationVerbOrder[o[j].Verb()]
}

type parentTopicsByRoot []Topic

func (o parentTopicsByRoot) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o parentTopicsByRoot) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o parentTopicsByRoot) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip, jp := strings.Split(strings.Trim(o[i].Name, "/"), "/"), strings.Split(strings.Trim(o[j].Name, "/"), "/")
	return ip[0] < jp[0]
}

type parentTopicsByGroup []Topic

func (o parentTopicsByGroup) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o parentTopicsByGroup) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o parentTopicsByGroup) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip, jp := strings.Split(strings.Trim(o[i].Name, "/"), "/"), strings.Split(strings.Trim(o[j].Name, "/"), "/")
	var ig, jg string
	if len(ip) == 3 {
		ig = strings.Join(ReverseStringSlice(strings.Split(ip[1], ".")), ".")
	}
	if len(jp) == 3 {
		jg = strings.Join(ReverseStringSlice(strings.Split(jp[1], ".")), ".")
	}
	if strings.Contains(ig, ".") == strings.Contains(jg, ".") {
		return ig < jg
	}
	return strings.Contains(jg, ".") && !strings.Contains(ig, ".")
}

type parentTopicsByVersion []Topic

func (o parentTopicsByVersion) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o parentTopicsByVersion) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o parentTopicsByVersion) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip, jp := strings.Split(strings.Trim(o[i].Name, "/"), "/"), strings.Split(strings.Trim(o[j].Name, "/"), "/")
	return ip[len(ip)-1] < jp[len(jp)-1]
}

type childTopicsByName []Topic

func (o childTopicsByName) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(o)
}
func (o childTopicsByName) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o[i], o[j] = o[j], o[i]
}
func (o childTopicsByName) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o[i].Name < o[j].Name
}
