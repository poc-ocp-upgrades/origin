package graphview

import (
	"sort"
	"k8s.io/apimachinery/pkg/util/sets"
)

type IntSet map[int]sets.Empty

func NewIntSet(items ...int) IntSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss := IntSet{}
	ss.Insert(items...)
	return ss
}
func (s IntSet) Insert(items ...int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, item := range items {
		s[item] = sets.Empty{}
	}
}
func (s IntSet) Delete(items ...int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, item := range items {
		delete(s, item)
	}
}
func (s IntSet) Has(item int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, contained := s[item]
	return contained
}
func (s IntSet) List() []int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := make([]int, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.IntSlice(res).Sort()
	return res
}
