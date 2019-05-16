package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"k8s.io/apimachinery/pkg/util/sets"
)

func stringSetToPointers(in sets.String) []*string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	out := make([]*string, 0, len(in))
	for k := range in {
		out = append(out, aws.String(k))
	}
	return out
}
func stringSetFromPointers(in []*string) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	out := sets.NewString()
	for i := range in {
		out.Insert(aws.StringValue(in[i]))
	}
	return out
}
