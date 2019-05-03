package gce

import (
 "context"
 "fmt"
 "net"
 "net/http"
 "strings"
 "time"
 "cloud.google.com/go/compute/metadata"
 computebeta "google.golang.org/api/compute/v0.beta"
 compute "google.golang.org/api/compute/v1"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/apimachinery/pkg/util/wait"
 cloudprovider "k8s.io/cloud-provider"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
 kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
)

const (
 defaultZone = ""
)

func newInstancesMetricContext(request, zone string) *metricContext {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return newGenericMetricContext("instances", request, unusedMetricLabel, zone, computeV1Version)
}
func splitNodesByZone(nodes []*v1.Node) map[string][]*v1.Node {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zones := make(map[string][]*v1.Node)
 for _, n := range nodes {
  z := getZone(n)
  if z != defaultZone {
   zones[z] = append(zones[z], n)
  }
 }
 return zones
}
func getZone(n *v1.Node) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 zone, ok := n.Labels[kubeletapis.LabelZoneFailureDomain]
 if !ok {
  return defaultZone
 }
 return zone
}
func makeHostURL(projectsAPIEndpoint, projectID, zone, host string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 host = canonicalizeInstanceName(host)
 return projectsAPIEndpoint + strings.Join([]string{projectID, "zones", zone, "instances", host}, "/")
}
func (g *Cloud) ToInstanceReferences(zone string, instanceNames []string) (refs []*compute.InstanceReference) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ins := range instanceNames {
  instanceLink := makeHostURL(g.service.BasePath, g.projectID, zone, ins)
  refs = append(refs, &compute.InstanceReference{Instance: instanceLink})
 }
 return refs
}
func (g *Cloud) NodeAddresses(_ context.Context, _ types.NodeName) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 internalIP, err := metadata.Get("instance/network-interfaces/0/ip")
 if err != nil {
  return nil, fmt.Errorf("couldn't get internal IP: %v", err)
 }
 externalIP, err := metadata.Get("instance/network-interfaces/0/access-configs/0/external-ip")
 if err != nil {
  return nil, fmt.Errorf("couldn't get external IP: %v", err)
 }
 addresses := []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: internalIP}, {Type: v1.NodeExternalIP, Address: externalIP}}
 if internalDNSFull, err := metadata.Get("instance/hostname"); err != nil {
  klog.Warningf("couldn't get full internal DNS name: %v", err)
 } else {
  addresses = append(addresses, v1.NodeAddress{Type: v1.NodeInternalDNS, Address: internalDNSFull}, v1.NodeAddress{Type: v1.NodeHostName, Address: internalDNSFull})
 }
 return addresses, nil
}
func (g *Cloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 _, zone, name, err := splitProviderID(providerID)
 if err != nil {
  return []v1.NodeAddress{}, err
 }
 instance, err := g.c.Instances().Get(ctx, meta.ZonalKey(canonicalizeInstanceName(name), zone))
 if err != nil {
  return []v1.NodeAddress{}, fmt.Errorf("error while querying for providerID %q: %v", providerID, err)
 }
 if len(instance.NetworkInterfaces) < 1 {
  return []v1.NodeAddress{}, fmt.Errorf("could not find network interfaces for providerID %q", providerID)
 }
 networkInterface := instance.NetworkInterfaces[0]
 nodeAddresses := []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: networkInterface.NetworkIP}}
 for _, config := range networkInterface.AccessConfigs {
  nodeAddresses = append(nodeAddresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: config.NatIP})
 }
 return nodeAddresses, nil
}
func (g *Cloud) instanceByProviderID(providerID string) (*gceInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 project, zone, name, err := splitProviderID(providerID)
 if err != nil {
  return nil, err
 }
 instance, err := g.getInstanceFromProjectInZoneByName(project, zone, name)
 if err != nil {
  if isHTTPErrorCode(err, http.StatusNotFound) {
   return nil, cloudprovider.InstanceNotFound
  }
  return nil, err
 }
 return instance, nil
}
func (g *Cloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false, cloudprovider.NotImplemented
}
func (g *Cloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instance, err := g.instanceByProviderID(providerID)
 if err != nil {
  return "", err
 }
 return instance.Type, nil
}
func (g *Cloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := g.instanceByProviderID(providerID)
 if err != nil {
  if err == cloudprovider.InstanceNotFound {
   return false, nil
  }
  return false, err
 }
 return true, nil
}
func (g *Cloud) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceName := mapNodeNameToInstanceName(nodeName)
 if g.useMetadataServer {
  if g.isCurrentInstance(instanceName) {
   projectID, zone, err := getProjectAndZone()
   if err == nil {
    return projectID + "/" + zone + "/" + canonicalizeInstanceName(instanceName), nil
   }
  }
 }
 instance, err := g.getInstanceByName(instanceName)
 if err != nil {
  return "", err
 }
 return g.projectID + "/" + instance.Zone + "/" + instance.Name, nil
}
func (g *Cloud) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceName := mapNodeNameToInstanceName(nodeName)
 if g.useMetadataServer {
  if g.isCurrentInstance(instanceName) {
   mType, err := getCurrentMachineTypeViaMetadata()
   if err == nil {
    return mType, nil
   }
  }
 }
 instance, err := g.getInstanceByName(instanceName)
 if err != nil {
  return "", err
 }
 return instance.Type, nil
}
func (g *Cloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 return wait.Poll(2*time.Second, 30*time.Second, func() (bool, error) {
  project, err := g.c.Projects().Get(ctx, g.projectID)
  if err != nil {
   klog.Errorf("Could not get project: %v", err)
   return false, nil
  }
  keyString := fmt.Sprintf("%s:%s %s@%s", user, strings.TrimSpace(string(keyData)), user, user)
  found := false
  for _, item := range project.CommonInstanceMetadata.Items {
   if item.Key == "sshKeys" {
    if strings.Contains(*item.Value, keyString) {
     klog.Info("SSHKey already in project metadata")
     return true, nil
    }
    value := *item.Value + "\n" + keyString
    item.Value = &value
    found = true
    break
   }
  }
  if !found {
   klog.Infof("Failed to find sshKeys metadata, creating a new item")
   project.CommonInstanceMetadata.Items = append(project.CommonInstanceMetadata.Items, &compute.MetadataItems{Key: "sshKeys", Value: &keyString})
  }
  mc := newInstancesMetricContext("add_ssh_key", "")
  err = g.c.Projects().SetCommonInstanceMetadata(ctx, g.projectID, project.CommonInstanceMetadata)
  mc.Observe(err)
  if err != nil {
   klog.Errorf("Could not Set Metadata: %v", err)
   return false, nil
  }
  klog.Infof("Successfully added sshKey to project metadata")
  return true, nil
 })
}
func (g *Cloud) GetAllCurrentZones() (sets.String, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if g.nodeInformerSynced == nil {
  klog.Warningf("Cloud object does not have informers set, should only happen in E2E binary.")
  return g.GetAllZonesFromCloudProvider()
 }
 g.nodeZonesLock.Lock()
 defer g.nodeZonesLock.Unlock()
 if !g.nodeInformerSynced() {
  return nil, fmt.Errorf("node informer is not synced when trying to GetAllCurrentZones")
 }
 zones := sets.NewString()
 for zone, nodes := range g.nodeZones {
  if len(nodes) > 0 {
   zones.Insert(zone)
  }
 }
 return zones, nil
}
func (g *Cloud) GetAllZonesFromCloudProvider() (sets.String, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 zones := sets.NewString()
 for _, zone := range g.managedZones {
  instances, err := g.c.Instances().List(ctx, zone, filter.None)
  if err != nil {
   return sets.NewString(), err
  }
  if len(instances) > 0 {
   zones.Insert(zone)
  }
 }
 return zones, nil
}
func (g *Cloud) InsertInstance(project string, zone string, i *compute.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 mc := newInstancesMetricContext("create", zone)
 return mc.Observe(g.c.Instances().Insert(ctx, meta.ZonalKey(i.Name, zone), i))
}
func (g *Cloud) ListInstanceNames(project, zone string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 l, err := g.c.Instances().List(ctx, zone, filter.None)
 if err != nil {
  return "", err
 }
 var names []string
 for _, i := range l {
  names = append(names, i.Name)
 }
 return strings.Join(names, " "), nil
}
func (g *Cloud) DeleteInstance(project, zone, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 return g.c.Instances().Delete(ctx, meta.ZonalKey(name, zone))
}
func (g *Cloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return types.NodeName(hostname), nil
}
func (g *Cloud) AliasRanges(nodeName types.NodeName) (cidrs []string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 var instance *gceInstance
 instance, err = g.getInstanceByName(mapNodeNameToInstanceName(nodeName))
 if err != nil {
  return
 }
 var res *computebeta.Instance
 res, err = g.c.BetaInstances().Get(ctx, meta.ZonalKey(instance.Name, lastComponent(instance.Zone)))
 if err != nil {
  return
 }
 for _, networkInterface := range res.NetworkInterfaces {
  for _, r := range networkInterface.AliasIpRanges {
   cidrs = append(cidrs, r.IpCidrRange)
  }
 }
 return
}
func (g *Cloud) AddAliasToInstance(nodeName types.NodeName, alias *net.IPNet) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 v1instance, err := g.getInstanceByName(mapNodeNameToInstanceName(nodeName))
 if err != nil {
  return err
 }
 instance, err := g.c.BetaInstances().Get(ctx, meta.ZonalKey(v1instance.Name, lastComponent(v1instance.Zone)))
 if err != nil {
  return err
 }
 switch len(instance.NetworkInterfaces) {
 case 0:
  return fmt.Errorf("instance %q has no network interfaces", nodeName)
 case 1:
 default:
  klog.Warningf("Instance %q has more than one network interface, using only the first (%v)", nodeName, instance.NetworkInterfaces)
 }
 iface := &computebeta.NetworkInterface{}
 iface.Name = instance.NetworkInterfaces[0].Name
 iface.Fingerprint = instance.NetworkInterfaces[0].Fingerprint
 iface.AliasIpRanges = append(iface.AliasIpRanges, &computebeta.AliasIpRange{IpCidrRange: alias.String(), SubnetworkRangeName: g.secondaryRangeName})
 mc := newInstancesMetricContext("add_alias", v1instance.Zone)
 err = g.c.BetaInstances().UpdateNetworkInterface(ctx, meta.ZonalKey(instance.Name, lastComponent(instance.Zone)), iface.Name, iface)
 return mc.Observe(err)
}
func (g *Cloud) getInstancesByNames(names []string) ([]*gceInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 found := map[string]*gceInstance{}
 remaining := len(names)
 nodeInstancePrefix := g.nodeInstancePrefix
 for _, name := range names {
  name = canonicalizeInstanceName(name)
  if !strings.HasPrefix(name, g.nodeInstancePrefix) {
   klog.Warningf("Instance %q does not conform to prefix %q, removing filter", name, g.nodeInstancePrefix)
   nodeInstancePrefix = ""
  }
  found[name] = nil
 }
 for _, zone := range g.managedZones {
  if remaining == 0 {
   break
  }
  instances, err := g.c.Instances().List(ctx, zone, filter.Regexp("name", nodeInstancePrefix+".*"))
  if err != nil {
   return nil, err
  }
  for _, inst := range instances {
   if remaining == 0 {
    break
   }
   if _, ok := found[inst.Name]; !ok {
    continue
   }
   if found[inst.Name] != nil {
    klog.Errorf("Instance name %q was duplicated (in zone %q and %q)", inst.Name, zone, found[inst.Name].Zone)
    continue
   }
   found[inst.Name] = &gceInstance{Zone: zone, Name: inst.Name, ID: inst.Id, Disks: inst.Disks, Type: lastComponent(inst.MachineType)}
   remaining--
  }
 }
 if remaining > 0 {
  var failed []string
  for k := range found {
   if found[k] == nil {
    failed = append(failed, k)
   }
  }
  klog.Errorf("Failed to retrieve instances: %v", failed)
  return nil, cloudprovider.InstanceNotFound
 }
 var ret []*gceInstance
 for _, instance := range found {
  ret = append(ret, instance)
 }
 return ret, nil
}
func (g *Cloud) getInstanceByName(name string) (*gceInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, zone := range g.managedZones {
  instance, err := g.getInstanceFromProjectInZoneByName(g.projectID, zone, name)
  if err != nil {
   if isHTTPErrorCode(err, http.StatusNotFound) {
    continue
   }
   klog.Errorf("getInstanceByName: failed to get instance %s in zone %s; err: %v", name, zone, err)
   return nil, err
  }
  return instance, nil
 }
 return nil, cloudprovider.InstanceNotFound
}
func (g *Cloud) getInstanceFromProjectInZoneByName(project, zone, name string) (*gceInstance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 name = canonicalizeInstanceName(name)
 mc := newInstancesMetricContext("get", zone)
 res, err := g.c.Instances().Get(ctx, meta.ZonalKey(name, zone))
 mc.Observe(err)
 if err != nil {
  return nil, err
 }
 return &gceInstance{Zone: lastComponent(res.Zone), Name: res.Name, ID: res.Id, Disks: res.Disks, Type: lastComponent(res.MachineType)}, nil
}
func getInstanceIDViaMetadata() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result, err := metadata.Get("instance/hostname")
 if err != nil {
  return "", err
 }
 parts := strings.Split(result, ".")
 if len(parts) == 0 {
  return "", fmt.Errorf("unexpected response: %s", result)
 }
 return parts[0], nil
}
func getCurrentMachineTypeViaMetadata() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mType, err := metadata.Get("instance/machine-type")
 if err != nil {
  return "", fmt.Errorf("couldn't get machine type: %v", err)
 }
 parts := strings.Split(mType, "/")
 if len(parts) != 4 {
  return "", fmt.Errorf("unexpected response for machine type: %s", mType)
 }
 return parts[3], nil
}
func (g *Cloud) isCurrentInstance(instanceID string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 currentInstanceID, err := getInstanceIDViaMetadata()
 if err != nil {
  klog.Errorf("Failed to fetch instanceID via Metadata: %v", err)
  return false
 }
 return currentInstanceID == canonicalizeInstanceName(instanceID)
}
func (g *Cloud) computeHostTags(hosts []*gceInstance) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ctx, cancel := cloud.ContextWithCallTimeout()
 defer cancel()
 hostNamesByZone := make(map[string]map[string]bool)
 nodeInstancePrefix := g.nodeInstancePrefix
 for _, host := range hosts {
  if !strings.HasPrefix(host.Name, g.nodeInstancePrefix) {
   klog.Warningf("instance %v does not conform to prefix '%s', ignoring filter", host, g.nodeInstancePrefix)
   nodeInstancePrefix = ""
  }
  z, ok := hostNamesByZone[host.Zone]
  if !ok {
   z = make(map[string]bool)
   hostNamesByZone[host.Zone] = z
  }
  z[host.Name] = true
 }
 tags := sets.NewString()
 filt := filter.None
 if nodeInstancePrefix != "" {
  filt = filter.Regexp("name", nodeInstancePrefix+".*")
 }
 for zone, hostNames := range hostNamesByZone {
  instances, err := g.c.Instances().List(ctx, zone, filt)
  if err != nil {
   return nil, err
  }
  for _, instance := range instances {
   if !hostNames[instance.Name] {
    continue
   }
   longestTag := ""
   for _, tag := range instance.Tags.Items {
    if strings.HasPrefix(instance.Name, tag) && len(tag) > len(longestTag) {
     longestTag = tag
    }
   }
   if len(longestTag) > 0 {
    tags.Insert(longestTag)
   } else {
    return nil, fmt.Errorf("could not find any tag that is a prefix of instance name for instance %s", instance.Name)
   }
  }
 }
 if len(tags) == 0 {
  return nil, fmt.Errorf("no instances found")
 }
 return tags.List(), nil
}
func (g *Cloud) GetNodeTags(nodeNames []string) ([]string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(g.nodeTags) > 0 {
  return g.nodeTags, nil
 }
 g.computeNodeTagLock.Lock()
 defer g.computeNodeTagLock.Unlock()
 hosts := sets.NewString(nodeNames...)
 if hosts.Equal(g.lastKnownNodeNames) {
  return g.lastComputedNodeTags, nil
 }
 instances, err := g.getInstancesByNames(nodeNames)
 if err != nil {
  return nil, err
 }
 tags, err := g.computeHostTags(instances)
 if err != nil {
  return nil, err
 }
 g.lastKnownNodeNames = hosts
 g.lastComputedNodeTags = tags
 return tags, nil
}
