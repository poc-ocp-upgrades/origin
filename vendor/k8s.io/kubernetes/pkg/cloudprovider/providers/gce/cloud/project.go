package cloud

import (
	"context"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

type ProjectRouter interface {
	ProjectID(ctx context.Context, version meta.Version, service string) string
}
type SingleProjectRouter struct{ ID string }

func (r *SingleProjectRouter) ProjectID(ctx context.Context, version meta.Version, service string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return r.ID
}
