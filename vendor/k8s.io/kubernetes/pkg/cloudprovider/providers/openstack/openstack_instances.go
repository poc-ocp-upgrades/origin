package openstack

import (
 "context"
 "fmt"
 "regexp"
 "github.com/gophercloud/gophercloud"
 "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
 "k8s.io/klog"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 cloudprovider "k8s.io/cloud-provider"
)

type Instances struct {
 compute *gophercloud.ServiceClient
 opts    MetadataOpts
}

const (
 instanceShutoff = "SHUTOFF"
)

func (os *OpenStack) Instances() (cloudprovider.Instances, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Info("openstack.Instances() called")
 compute, err := os.NewComputeV2()
 if err != nil {
  klog.Errorf("unable to access compute v2 API : %v", err)
  return nil, false
 }
 klog.V(4).Info("Claiming to support Instances")
 return &Instances{compute: compute, opts: os.metadataOpts}, true
}
func (i *Instances) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 md, err := getMetadata(i.opts.SearchOrder)
 if err != nil {
  return "", err
 }
 return types.NodeName(md.Name), nil
}
func (i *Instances) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.NotImplemented
}
func (i *Instances) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(4).Infof("NodeAddresses(%v) called", name)
 addrs, err := getAddressesByName(i.compute, name)
 if err != nil {
  return nil, err
 }
 klog.V(4).Infof("NodeAddresses(%v) => %v", name, addrs)
 return addrs, nil
}
func (i *Instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceID, err := instanceIDFromProviderID(providerID)
 if err != nil {
  return []v1.NodeAddress{}, err
 }
 server, err := servers.Get(i.compute, instanceID).Extract()
 if err != nil {
  return []v1.NodeAddress{}, err
 }
 addresses, err := nodeAddresses(server)
 if err != nil {
  return []v1.NodeAddress{}, err
 }
 return addresses, nil
}
func (i *Instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceID, err := instanceIDFromProviderID(providerID)
 if err != nil {
  return false, err
 }
 _, err = servers.Get(i.compute, instanceID).Extract()
 if err != nil {
  if isNotFound(err) {
   return false, nil
  }
  return false, err
 }
 return true, nil
}
func (i *Instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceID, err := instanceIDFromProviderID(providerID)
 if err != nil {
  return false, err
 }
 server, err := servers.Get(i.compute, instanceID).Extract()
 if err != nil {
  return false, err
 }
 if server.Status == instanceShutoff {
  return true, nil
 }
 return false, nil
}
func (os *OpenStack) InstanceID() (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(os.localInstanceID) == 0 {
  id, err := readInstanceID(os.metadataOpts.SearchOrder)
  if err != nil {
   return "", err
  }
  os.localInstanceID = id
 }
 return os.localInstanceID, nil
}
func (i *Instances) InstanceID(ctx context.Context, name types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 srv, err := getServerByName(i.compute, name)
 if err != nil {
  if err == ErrNotFound {
   return "", cloudprovider.InstanceNotFound
  }
  return "", err
 }
 return "/" + srv.ID, nil
}
func (i *Instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 instanceID, err := instanceIDFromProviderID(providerID)
 if err != nil {
  return "", err
 }
 server, err := servers.Get(i.compute, instanceID).Extract()
 if err != nil {
  return "", err
 }
 return srvInstanceType(server)
}
func (i *Instances) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 srv, err := getServerByName(i.compute, name)
 if err != nil {
  return "", err
 }
 return srvInstanceType(srv)
}
func srvInstanceType(srv *servers.Server) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 keys := []string{"name", "id", "original_name"}
 for _, key := range keys {
  val, found := srv.Flavor[key]
  if found {
   flavor, ok := val.(string)
   if ok {
    return flavor, nil
   }
  }
 }
 return "", fmt.Errorf("flavor name/id not found")
}
func instanceIDFromProviderID(providerID string) (instanceID string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var providerIDRegexp = regexp.MustCompile(`^` + ProviderName + `:///([^/]+)$`)
 matches := providerIDRegexp.FindStringSubmatch(providerID)
 if len(matches) != 2 {
  return "", fmt.Errorf("ProviderID \"%s\" didn't match expected format \"openstack:///InstanceID\"", providerID)
 }
 return matches[1], nil
}
