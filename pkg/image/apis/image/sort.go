package image

import (
	"sort"
	"time"
)

type tag struct {
	Name	string
	Created	time.Time
}
type byCreationTimestamp []tag

func (t byCreationTimestamp) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(t)
}
func (t byCreationTimestamp) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	delta := t[i].Created.Sub(t[j].Created)
	if delta > 0 {
		return true
	}
	if delta < 0 {
		return false
	}
	return t[i].Name < t[j].Name
}
func (t byCreationTimestamp) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t[i], t[j] = t[j], t[i]
}
func SortStatusTags(tags map[string]TagEventList) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tagSlice := make([]tag, len(tags))
	index := 0
	for tag, list := range tags {
		tagSlice[index].Name = tag
		if len(list.Items) > 0 {
			tagSlice[index].Created = list.Items[0].Created.Round(time.Second)
		}
		index++
	}
	sort.Sort(byCreationTimestamp(tagSlice))
	actual := make([]string, len(tagSlice))
	for i, tag := range tagSlice {
		actual[i] = tag.Name
	}
	return actual
}
