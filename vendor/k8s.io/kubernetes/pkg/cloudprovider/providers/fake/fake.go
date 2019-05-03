package fake

import (
 "context"
 "fmt"
 "net"
 "regexp"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 cloudprovider "k8s.io/cloud-provider"
)

const defaultProviderName = "fake"

type FakeBalancer struct {
 Name           string
 Region         string
 LoadBalancerIP string
 Ports          []v1.ServicePort
 Hosts          []*v1.Node
}
type FakeUpdateBalancerCall struct {
 Service *v1.Service
 Hosts   []*v1.Node
}
type FakeCloud struct {
 Exists                  bool
 Err                     error
 ExistsByProviderID      bool
 ErrByProviderID         error
 NodeShutdown            bool
 ErrShutdownByProviderID error
 Calls                   []string
 Addresses               []v1.NodeAddress
 addressesMux            sync.Mutex
 ExtID                   map[types.NodeName]string
 InstanceTypes           map[types.NodeName]string
 Machines                []types.NodeName
 NodeResources           *v1.NodeResources
 ClusterList             []string
 MasterName              string
 ExternalIP              net.IP
 Balancers               map[string]FakeBalancer
 UpdateCalls             []FakeUpdateBalancerCall
 RouteMap                map[string]*FakeRoute
 Lock                    sync.Mutex
 Provider                string
 addCallLock             sync.Mutex
 cloudprovider.Zone
 VolumeLabelMap map[string]map[string]string
 RequestDelay   time.Duration
}
type FakeRoute struct {
 ClusterName string
 Route       cloudprovider.Route
}

func (f *FakeCloud) addCall(desc string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCallLock.Lock()
 defer f.addCallLock.Unlock()
 time.Sleep(f.RequestDelay)
 f.Calls = append(f.Calls, desc)
}
func (f *FakeCloud) ClearCalls() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Calls = []string{}
}
func (f *FakeCloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (f *FakeCloud) ListClusters(ctx context.Context) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f.ClusterList, f.Err
}
func (f *FakeCloud) Master(ctx context.Context, name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f.MasterName, f.Err
}
func (f *FakeCloud) Clusters() (cloudprovider.Clusters, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, true
}
func (f *FakeCloud) ProviderName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if f.Provider == "" {
  return defaultProviderName
 }
 return f.Provider
}
func (f *FakeCloud) HasClusterID() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (f *FakeCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, true
}
func (f *FakeCloud) Instances() (cloudprovider.Instances, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, true
}
func (f *FakeCloud) Zones() (cloudprovider.Zones, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, true
}
func (f *FakeCloud) Routes() (cloudprovider.Routes, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return f, true
}
func (f *FakeCloud) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (*v1.LoadBalancerStatus, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 status := &v1.LoadBalancerStatus{}
 status.Ingress = []v1.LoadBalancerIngress{{IP: f.ExternalIP.String()}}
 return status, f.Exists, f.Err
}
func (f *FakeCloud) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.DefaultLoadBalancerName(service)
}
func (f *FakeCloud) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("create")
 if f.Balancers == nil {
  f.Balancers = make(map[string]FakeBalancer)
 }
 name := f.GetLoadBalancerName(ctx, clusterName, service)
 spec := service.Spec
 zone, err := f.GetZone(context.TODO())
 if err != nil {
  return nil, err
 }
 region := zone.Region
 f.Balancers[name] = FakeBalancer{name, region, spec.LoadBalancerIP, spec.Ports, nodes}
 status := &v1.LoadBalancerStatus{}
 status.Ingress = []v1.LoadBalancerIngress{{IP: f.ExternalIP.String()}}
 return status, f.Err
}
func (f *FakeCloud) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("update")
 f.UpdateCalls = append(f.UpdateCalls, FakeUpdateBalancerCall{service, nodes})
 return f.Err
}
func (f *FakeCloud) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("delete")
 return f.Err
}
func (f *FakeCloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.NotImplemented
}
func (f *FakeCloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return types.NodeName(hostname), nil
}
func (f *FakeCloud) NodeAddresses(ctx context.Context, instance types.NodeName) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("node-addresses")
 f.addressesMux.Lock()
 defer f.addressesMux.Unlock()
 return f.Addresses, f.Err
}
func (f *FakeCloud) SetNodeAddresses(nodeAddresses []v1.NodeAddress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addressesMux.Lock()
 defer f.addressesMux.Unlock()
 f.Addresses = nodeAddresses
}
func (f *FakeCloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("node-addresses-by-provider-id")
 return f.Addresses, f.Err
}
func (f *FakeCloud) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("instance-id")
 return f.ExtID[nodeName], nil
}
func (f *FakeCloud) InstanceType(ctx context.Context, instance types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("instance-type")
 return f.InstanceTypes[instance], nil
}
func (f *FakeCloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("instance-type-by-provider-id")
 return f.InstanceTypes[types.NodeName(providerID)], nil
}
func (f *FakeCloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("instance-exists-by-provider-id")
 return f.ExistsByProviderID, f.ErrByProviderID
}
func (f *FakeCloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("instance-shutdown-by-provider-id")
 return f.NodeShutdown, f.ErrShutdownByProviderID
}
func (f *FakeCloud) List(filter string) ([]types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("list")
 result := []types.NodeName{}
 for _, machine := range f.Machines {
  if match, _ := regexp.MatchString(filter, string(machine)); match {
   result = append(result, machine)
  }
 }
 return result, f.Err
}
func (f *FakeCloud) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("get-zone")
 return f.Zone, f.Err
}
func (f *FakeCloud) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("get-zone-by-provider-id")
 return f.Zone, f.Err
}
func (f *FakeCloud) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.addCall("get-zone-by-node-name")
 return f.Zone, f.Err
}
func (f *FakeCloud) ListRoutes(ctx context.Context, clusterName string) ([]*cloudprovider.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock.Lock()
 defer f.Lock.Unlock()
 f.addCall("list-routes")
 var routes []*cloudprovider.Route
 for _, fakeRoute := range f.RouteMap {
  if clusterName == fakeRoute.ClusterName {
   routeCopy := fakeRoute.Route
   routes = append(routes, &routeCopy)
  }
 }
 return routes, f.Err
}
func (f *FakeCloud) CreateRoute(ctx context.Context, clusterName string, nameHint string, route *cloudprovider.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock.Lock()
 defer f.Lock.Unlock()
 f.addCall("create-route")
 name := clusterName + "-" + nameHint
 if _, exists := f.RouteMap[name]; exists {
  f.Err = fmt.Errorf("route %q already exists", name)
  return f.Err
 }
 fakeRoute := FakeRoute{}
 fakeRoute.Route = *route
 fakeRoute.Route.Name = name
 fakeRoute.ClusterName = clusterName
 f.RouteMap[name] = &fakeRoute
 return nil
}
func (f *FakeCloud) DeleteRoute(ctx context.Context, clusterName string, route *cloudprovider.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 f.Lock.Lock()
 defer f.Lock.Unlock()
 f.addCall("delete-route")
 name := route.Name
 if _, exists := f.RouteMap[name]; !exists {
  f.Err = fmt.Errorf("no route found with name %q", name)
  return f.Err
 }
 delete(f.RouteMap, name)
 return nil
}
func (c *FakeCloud) GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if val, ok := c.VolumeLabelMap[pv.Name]; ok {
  return val, nil
 }
 return nil, fmt.Errorf("label not found for volume")
}
