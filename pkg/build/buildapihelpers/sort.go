package buildapihelpers

import buildv1 "github.com/openshift/api/build/v1"

type BuildSliceByCreationTimestamp []buildv1.Build

func (b BuildSliceByCreationTimestamp) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b)
}
func (b BuildSliceByCreationTimestamp) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b[i].CreationTimestamp.Before(&b[j].CreationTimestamp)
}
func (b BuildSliceByCreationTimestamp) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b[i], b[j] = b[j], b[i]
}
