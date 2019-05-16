package buildapihelpers

import buildv1 "github.com/openshift/api/build/v1"

type BuildSliceByCreationTimestamp []buildv1.Build

func (b BuildSliceByCreationTimestamp) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(b)
}
func (b BuildSliceByCreationTimestamp) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return b[i].CreationTimestamp.Before(&b[j].CreationTimestamp)
}
func (b BuildSliceByCreationTimestamp) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b[i], b[j] = b[j], b[i]
}
