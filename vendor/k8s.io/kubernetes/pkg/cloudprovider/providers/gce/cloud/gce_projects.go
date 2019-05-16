package cloud

import (
	"context"
	"fmt"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
	"net/http"
)

type ProjectsOps interface {
	Get(ctx context.Context, projectID string) (*compute.Project, error)
	SetCommonInstanceMetadata(ctx context.Context, projectID string, m *compute.Metadata) error
}
type MockProjectOpsState struct{ metadata map[string]*compute.Metadata }

func (m *MockProjects) Get(ctx context.Context, projectID string) (*compute.Project, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.Lock.Lock()
	defer m.Lock.Unlock()
	if p, ok := m.Objects[*meta.GlobalKey(projectID)]; ok {
		return p.ToGA(), nil
	}
	return nil, &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockProjects %v not found", projectID)}
}
func (g *GCEProjects) Get(ctx context.Context, projectID string) (*compute.Project, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Projects"}
	if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
		return nil, err
	}
	call := g.s.GA.Projects.Get(projectID)
	call.Context(ctx)
	return call.Do()
}
func (m *MockProjects) SetCommonInstanceMetadata(ctx context.Context, projectID string, meta *compute.Metadata) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if m.X == nil {
		m.X = &MockProjectOpsState{metadata: map[string]*compute.Metadata{}}
	}
	state := m.X.(*MockProjectOpsState)
	state.metadata[projectID] = meta
	return nil
}
func (g *GCEProjects) SetCommonInstanceMetadata(ctx context.Context, projectID string, m *compute.Metadata) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rk := &RateLimitKey{ProjectID: projectID, Operation: "SetCommonInstanceMetadata", Version: meta.Version("ga"), Service: "Projects"}
	if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
		return err
	}
	call := g.s.GA.Projects.SetCommonInstanceMetadata(projectID, m)
	call.Context(ctx)
	op, err := call.Do()
	if err != nil {
		return err
	}
	return g.s.WaitForCompletion(ctx, op)
}
