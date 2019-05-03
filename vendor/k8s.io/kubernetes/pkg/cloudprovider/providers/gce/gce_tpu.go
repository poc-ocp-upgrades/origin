package gce

import (
 "context"
 "encoding/json"
 "fmt"
 "net/http"
 "time"
 "google.golang.org/api/googleapi"
 tpuapi "google.golang.org/api/tpu/v1"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/util/wait"
)

func newTPUService(client *http.Client) (*tpuService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s, err := tpuapi.New(client)
 if err != nil {
  return nil, err
 }
 return &tpuService{projects: tpuapi.NewProjectsService(s)}, nil
}

type tpuService struct{ projects *tpuapi.ProjectsService }

func (g *Cloud) CreateTPU(ctx context.Context, name, zone string, node *tpuapi.Node) (*tpuapi.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 mc := newTPUMetricContext("create", zone)
 defer mc.Observe(err)
 var op *tpuapi.Operation
 parent := getTPUParentName(g.projectID, zone)
 op, err = g.tpuService.projects.Locations.Nodes.Create(parent, node).NodeId(name).Do()
 if err != nil {
  return nil, err
 }
 klog.V(2).Infof("Creating Cloud TPU %q in zone %q with operation %q", name, zone, op.Name)
 op, err = g.waitForTPUOp(ctx, op)
 if err != nil {
  return nil, err
 }
 err = getErrorFromTPUOp(op)
 if err != nil {
  return nil, err
 }
 output := new(tpuapi.Node)
 err = json.Unmarshal(op.Response, output)
 if err != nil {
  err = fmt.Errorf("failed to unmarshal response from operation %q: response = %v, err = %v", op.Name, op.Response, err)
  return nil, err
 }
 return output, nil
}
func (g *Cloud) DeleteTPU(ctx context.Context, name, zone string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 mc := newTPUMetricContext("delete", zone)
 defer mc.Observe(err)
 var op *tpuapi.Operation
 name = getTPUName(g.projectID, zone, name)
 op, err = g.tpuService.projects.Locations.Nodes.Delete(name).Do()
 if err != nil {
  return err
 }
 klog.V(2).Infof("Deleting Cloud TPU %q in zone %q with operation %q", name, zone, op.Name)
 op, err = g.waitForTPUOp(ctx, op)
 if err != nil {
  return err
 }
 err = getErrorFromTPUOp(op)
 if err != nil {
  return err
 }
 return nil
}
func (g *Cloud) GetTPU(ctx context.Context, name, zone string) (*tpuapi.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mc := newTPUMetricContext("get", zone)
 name = getTPUName(g.projectID, zone, name)
 node, err := g.tpuService.projects.Locations.Nodes.Get(name).Do()
 if err != nil {
  return nil, mc.Observe(err)
 }
 return node, mc.Observe(nil)
}
func (g *Cloud) ListTPUs(ctx context.Context, zone string) ([]*tpuapi.Node, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mc := newTPUMetricContext("list", zone)
 parent := getTPUParentName(g.projectID, zone)
 response, err := g.tpuService.projects.Locations.Nodes.List(parent).Do()
 if err != nil {
  return nil, mc.Observe(err)
 }
 return response.Nodes, mc.Observe(nil)
}
func (g *Cloud) ListLocations(ctx context.Context) ([]*tpuapi.Location, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mc := newTPUMetricContext("list_locations", "")
 parent := getTPUProjectURL(g.projectID)
 response, err := g.tpuService.projects.Locations.List(parent).Do()
 if err != nil {
  return nil, mc.Observe(err)
 }
 return response.Locations, mc.Observe(nil)
}
func (g *Cloud) waitForTPUOp(ctx context.Context, op *tpuapi.Operation) (*tpuapi.Operation, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err := wait.PollInfinite(30*time.Second, func() (bool, error) {
  select {
  case <-ctx.Done():
   klog.V(3).Infof("Context for operation %q has been cancelled: %s", op.Name, ctx.Err())
   return true, ctx.Err()
  default:
  }
  klog.V(3).Infof("Waiting for operation %q to complete...", op.Name)
  start := time.Now()
  g.operationPollRateLimiter.Accept()
  duration := time.Now().Sub(start)
  if duration > 5*time.Second {
   klog.V(2).Infof("Getting operation %q throttled for %v", op.Name, duration)
  }
  var err error
  op, err = g.tpuService.projects.Locations.Operations.Get(op.Name).Do()
  if err != nil {
   return true, err
  }
  if op.Done {
   klog.V(3).Infof("Operation %q has completed", op.Name)
   return true, nil
  }
  return false, nil
 }); err != nil {
  return nil, fmt.Errorf("failed to wait for operation %q: %s", op.Name, err)
 }
 return op, nil
}
func newTPUMetricContext(request, zone string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("tpus", request, unusedMetricLabel, zone, "v1")
}
func getErrorFromTPUOp(op *tpuapi.Operation) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if op != nil && op.Error != nil {
  return &googleapi.Error{Code: op.ServerResponse.HTTPStatusCode, Message: op.Error.Message}
 }
 return nil
}
func getTPUProjectURL(project string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("projects/%s", project)
}
func getTPUParentName(project, zone string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("projects/%s/locations/%s", project, zone)
}
func getTPUName(project, zone, name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("projects/%s/locations/%s/nodes/%s", project, zone, name)
}
