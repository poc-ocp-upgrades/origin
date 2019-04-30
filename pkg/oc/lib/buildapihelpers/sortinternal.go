package buildapihelpers

import (
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
)

type BuildSliceByCreationTimestampInternal []buildapi.Build

func (b BuildSliceByCreationTimestampInternal) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b)
}
func (b BuildSliceByCreationTimestampInternal) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b[i].CreationTimestamp.Before(&b[j].CreationTimestamp)
}
func (b BuildSliceByCreationTimestampInternal) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b[i], b[j] = b[j], b[i]
}

type BuildPtrSliceByCreationTimestampInternal []*buildapi.Build

func (b BuildPtrSliceByCreationTimestampInternal) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b)
}
func (b BuildPtrSliceByCreationTimestampInternal) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b[i].CreationTimestamp.Before(&b[j].CreationTimestamp)
}
func (b BuildPtrSliceByCreationTimestampInternal) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b[i], b[j] = b[j], b[i]
}
