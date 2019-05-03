package gce

import (
 "errors"
 "fmt"
 "net"
 "net/http"
 "regexp"
 "sort"
 "strings"
 "sync"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/apimachinery/pkg/util/sets"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/mock"
 "cloud.google.com/go/compute/metadata"
 compute "google.golang.org/api/compute/v1"
 "google.golang.org/api/googleapi"
)

func fakeGCECloud(vals TestClusterValues) (*Cloud, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 gce := NewFakeGCECloud(vals)
 gce.AlphaFeatureGate = NewAlphaFeatureGate([]string{})
 gce.nodeInformerSynced = func() bool {
  return true
 }
 mockGCE := gce.c.(*cloud.MockGCE)
 mockGCE.MockTargetPools.AddInstanceHook = mock.AddInstanceHook
 mockGCE.MockTargetPools.RemoveInstanceHook = mock.RemoveInstanceHook
 mockGCE.MockForwardingRules.InsertHook = mock.InsertFwdRuleHook
 mockGCE.MockAddresses.InsertHook = mock.InsertAddressHook
 mockGCE.MockAlphaAddresses.InsertHook = mock.InsertAlphaAddressHook
 mockGCE.MockAlphaAddresses.X = mock.AddressAttributes{}
 mockGCE.MockAddresses.X = mock.AddressAttributes{}
 mockGCE.MockInstanceGroups.X = mock.InstanceGroupAttributes{InstanceMap: make(map[meta.Key]map[string]*compute.InstanceWithNamedPorts), Lock: &sync.Mutex{}}
 mockGCE.MockInstanceGroups.AddInstancesHook = mock.AddInstancesHook
 mockGCE.MockInstanceGroups.RemoveInstancesHook = mock.RemoveInstancesHook
 mockGCE.MockInstanceGroups.ListInstancesHook = mock.ListInstancesHook
 mockGCE.MockRegionBackendServices.UpdateHook = mock.UpdateRegionBackendServiceHook
 mockGCE.MockHealthChecks.UpdateHook = mock.UpdateHealthCheckHook
 mockGCE.MockFirewalls.UpdateHook = mock.UpdateFirewallHook
 keyGA := meta.GlobalKey("key-ga")
 mockGCE.MockZones.Objects[*keyGA] = &cloud.MockZonesObj{Obj: &compute.Zone{Name: vals.ZoneName, Region: gce.getRegionLink(vals.Region)}}
 return gce, nil
}

type gceInstance struct {
 Zone  string
 Name  string
 ID    uint64
 Disks []*compute.AttachedDisk
 Type  string
}

var (
 autoSubnetIPRange = &net.IPNet{IP: net.ParseIP("10.128.0.0"), Mask: net.CIDRMask(9, 32)}
)
var providerIDRE = regexp.MustCompile(`^` + ProviderName + `://([^/]+)/([^/]+)/([^/]+)$`)

func getProjectAndZone() (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 result, err := metadata.Get("instance/zone")
 if err != nil {
  return "", "", err
 }
 parts := strings.Split(result, "/")
 if len(parts) != 4 {
  return "", "", fmt.Errorf("unexpected response: %s", result)
 }
 zone := parts[3]
 projectID, err := metadata.ProjectID()
 if err != nil {
  return "", "", err
 }
 return projectID, zone, nil
}
func (g *Cloud) raiseFirewallChangeNeededEvent(svc *v1.Service, cmd string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 msg := fmt.Sprintf("Firewall change required by network admin: `%v`", cmd)
 if g.eventRecorder != nil && svc != nil {
  g.eventRecorder.Event(svc, v1.EventTypeNormal, "LoadBalancerManualChange", msg)
 }
}
func FirewallToGCloudCreateCmd(fw *compute.Firewall, projectID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 args := firewallToGcloudArgs(fw, projectID)
 return fmt.Sprintf("gcloud compute firewall-rules create %v --network %v %v", fw.Name, getNameFromLink(fw.Network), args)
}
func FirewallToGCloudUpdateCmd(fw *compute.Firewall, projectID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 args := firewallToGcloudArgs(fw, projectID)
 return fmt.Sprintf("gcloud compute firewall-rules update %v %v", fw.Name, args)
}
func FirewallToGCloudDeleteCmd(fwName, projectID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("gcloud compute firewall-rules delete %v --project %v", fwName, projectID)
}
func firewallToGcloudArgs(fw *compute.Firewall, projectID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var allPorts []string
 for _, a := range fw.Allowed {
  for _, p := range a.Ports {
   allPorts = append(allPorts, fmt.Sprintf("%v:%v", a.IPProtocol, p))
  }
 }
 sort.Strings(allPorts)
 allow := strings.Join(allPorts, ",")
 sort.Strings(fw.SourceRanges)
 srcRngs := strings.Join(fw.SourceRanges, ",")
 sort.Strings(fw.TargetTags)
 targets := strings.Join(fw.TargetTags, ",")
 return fmt.Sprintf("--description %q --allow %v --source-ranges %v --target-tags %v --project %v", fw.Description, allow, srcRngs, targets, projectID)
}
func canonicalizeInstanceName(name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ix := strings.Index(name, ".")
 if ix != -1 {
  name = name[:ix]
 }
 return name
}
func lastComponent(s string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 lastSlash := strings.LastIndex(s, "/")
 if lastSlash != -1 {
  s = s[lastSlash+1:]
 }
 return s
}
func mapNodeNameToInstanceName(nodeName types.NodeName) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return string(nodeName)
}
func GetGCERegion(zone string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ix := strings.LastIndex(zone, "-")
 if ix == -1 {
  return "", fmt.Errorf("unexpected zone: %s", zone)
 }
 return zone[:ix], nil
}
func isHTTPErrorCode(err error, code int) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiErr, ok := err.(*googleapi.Error)
 return ok && apiErr.Code == code
}
func isInUsedByError(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 apiErr, ok := err.(*googleapi.Error)
 if !ok || apiErr.Code != http.StatusBadRequest {
  return false
 }
 return strings.Contains(apiErr.Message, "being used by")
}
func splitProviderID(providerID string) (project, zone, instance string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 matches := providerIDRE.FindStringSubmatch(providerID)
 if len(matches) != 4 {
  return "", "", "", errors.New("error splitting providerID")
 }
 return matches[1], matches[2], matches[3], nil
}
func equalStringSets(x, y []string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(x) != len(y) {
  return false
 }
 xString := sets.NewString(x...)
 yString := sets.NewString(y...)
 return xString.Equal(yString)
}
func isNotFound(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return isHTTPErrorCode(err, http.StatusNotFound)
}
func ignoreNotFound(err error) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if err == nil || isNotFound(err) {
  return nil
 }
 return err
}
func isNotFoundOrInUse(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return isNotFound(err) || isInUsedByError(err)
}
func isForbidden(err error) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return isHTTPErrorCode(err, http.StatusForbidden)
}
func makeGoogleAPINotFoundError(message string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &googleapi.Error{Code: http.StatusNotFound, Message: message}
}
func makeGoogleAPIError(code int, message string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &googleapi.Error{Code: code, Message: message}
}
func handleAlphaNetworkTierGetError(err error) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if isForbidden(err) {
  return cloud.NetworkTierDefault.ToGCEValue(), nil
 }
 return "", err
}
func containsCIDR(outer, inner *net.IPNet) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return outer.Contains(firstIPInRange(inner)) && outer.Contains(lastIPInRange(inner))
}
func firstIPInRange(ipNet *net.IPNet) net.IP {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ipNet.IP.Mask(ipNet.Mask)
}
func lastIPInRange(cidr *net.IPNet) net.IP {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ip := append([]byte{}, cidr.IP...)
 for i, b := range cidr.Mask {
  ip[i] |= ^b
 }
 return ip
}
func subnetsInCIDR(subnets []*compute.Subnetwork, cidr *net.IPNet) ([]*compute.Subnetwork, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var res []*compute.Subnetwork
 for _, subnet := range subnets {
  _, subnetRange, err := net.ParseCIDR(subnet.IpCidrRange)
  if err != nil {
   return nil, fmt.Errorf("unable to parse CIDR %q for subnet %q: %v", subnet.IpCidrRange, subnet.Name, err)
  }
  if containsCIDR(cidr, subnetRange) {
   res = append(res, subnet)
  }
 }
 return res, nil
}

type netType string

const (
 netTypeLegacy netType = "LEGACY"
 netTypeAuto   netType = "AUTO"
 netTypeCustom netType = "CUSTOM"
)

func typeOfNetwork(network *compute.Network) netType {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if network.IPv4Range != "" {
  return netTypeLegacy
 }
 if network.AutoCreateSubnetworks {
  return netTypeAuto
 }
 return netTypeCustom
}
func getLocationName(project, zoneOrRegion string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("projects/%s/locations/%s", project, zoneOrRegion)
}
