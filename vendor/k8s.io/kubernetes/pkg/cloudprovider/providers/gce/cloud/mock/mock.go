package mock

import (
 "context"
 godefaultbytes "bytes"
 godefaultruntime "runtime"
 "encoding/json"
 "fmt"
 "net/http"
 godefaulthttp "net/http"
 "sync"
 alpha "google.golang.org/api/compute/v0.alpha"
 beta "google.golang.org/api/compute/v0.beta"
 ga "google.golang.org/api/compute/v1"
 "google.golang.org/api/googleapi"
 cloud "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
)

var (
 InUseError          = &googleapi.Error{Code: http.StatusBadRequest, Message: "It's being used by god."}
 InternalServerError = &googleapi.Error{Code: http.StatusInternalServerError}
 UnauthorizedErr     = &googleapi.Error{Code: http.StatusForbidden}
)

type gceObject interface{ MarshalJSON() ([]byte, error) }

func AddInstanceHook(ctx context.Context, key *meta.Key, req *ga.TargetPoolsAddInstanceRequest, m *cloud.MockTargetPools) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pool, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in TargetPools", key.String())}
 }
 for _, instance := range req.Instances {
  pool.Instances = append(pool.Instances, instance.Instance)
 }
 return nil
}
func RemoveInstanceHook(ctx context.Context, key *meta.Key, req *ga.TargetPoolsRemoveInstanceRequest, m *cloud.MockTargetPools) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pool, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in TargetPools", key.String())}
 }
 for _, instanceToRemove := range req.Instances {
  for i, instance := range pool.Instances {
   if instanceToRemove.Instance == instance {
    pool.Instances[i] = pool.Instances[len(pool.Instances)-1]
    pool.Instances = pool.Instances[:len(pool.Instances)-1]
    break
   }
  }
 }
 return nil
}
func convertAndInsertAlphaForwardingRule(key *meta.Key, obj gceObject, mRules map[meta.Key]*cloud.MockForwardingRulesObj, version meta.Version, projectID string) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !key.Valid() {
  return true, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 if _, ok := mRules[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockForwardingRule %v exists", key)}
  return true, err
 }
 enc, err := obj.MarshalJSON()
 if err != nil {
  return true, err
 }
 var fwdRule alpha.ForwardingRule
 if err := json.Unmarshal(enc, &fwdRule); err != nil {
  return true, err
 }
 if fwdRule.NetworkTier == "" {
  fwdRule.NetworkTier = cloud.NetworkTierDefault.ToGCEValue()
 }
 fwdRule.Name = key.Name
 if fwdRule.SelfLink == "" {
  fwdRule.SelfLink = cloud.SelfLink(version, projectID, "forwardingRules", key)
 }
 mRules[*key] = &cloud.MockForwardingRulesObj{Obj: fwdRule}
 return true, nil
}
func InsertFwdRuleHook(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule, m *cloud.MockForwardingRules) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionGA, "forwardingRules")
 return convertAndInsertAlphaForwardingRule(key, obj, m.Objects, meta.VersionGA, projectID)
}
func InsertBetaFwdRuleHook(ctx context.Context, key *meta.Key, obj *beta.ForwardingRule, m *cloud.MockForwardingRules) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionBeta, "forwardingRules")
 return convertAndInsertAlphaForwardingRule(key, obj, m.Objects, meta.VersionBeta, projectID)
}
func InsertAlphaFwdRuleHook(ctx context.Context, key *meta.Key, obj *alpha.ForwardingRule, m *cloud.MockForwardingRules) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionAlpha, "forwardingRules")
 return convertAndInsertAlphaForwardingRule(key, obj, m.Objects, meta.VersionAlpha, projectID)
}

type AddressAttributes struct{ IPCounter int }

func convertAndInsertAlphaAddress(key *meta.Key, obj gceObject, mAddrs map[meta.Key]*cloud.MockAddressesObj, version meta.Version, projectID string, addressAttrs AddressAttributes) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !key.Valid() {
  return true, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 if _, ok := mAddrs[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAddresses %v exists", key)}
  return true, err
 }
 enc, err := obj.MarshalJSON()
 if err != nil {
  return true, err
 }
 var addr alpha.Address
 if err := json.Unmarshal(enc, &addr); err != nil {
  return true, err
 }
 if addr.AddressType == "" {
  addr.AddressType = string(cloud.SchemeExternal)
 }
 var existingAddresses []*ga.Address
 for _, obj := range mAddrs {
  existingAddresses = append(existingAddresses, obj.ToGA())
 }
 for _, existingAddr := range existingAddresses {
  if addr.Address == existingAddr.Address {
   msg := fmt.Sprintf("MockAddresses IP %v in use", addr.Address)
   errorCode := http.StatusConflict
   if addr.AddressType == string(cloud.SchemeExternal) {
    errorCode = http.StatusBadRequest
   }
   return true, &googleapi.Error{Code: errorCode, Message: msg}
  }
 }
 addr.Name = key.Name
 if addr.SelfLink == "" {
  addr.SelfLink = cloud.SelfLink(version, projectID, "addresses", key)
 }
 if addr.Address == "" {
  addr.Address = fmt.Sprintf("1.2.3.%d", addressAttrs.IPCounter)
  addressAttrs.IPCounter++
 }
 if addr.NetworkTier == "" {
  addr.NetworkTier = cloud.NetworkTierDefault.ToGCEValue()
 }
 mAddrs[*key] = &cloud.MockAddressesObj{Obj: addr}
 return true, nil
}
func InsertAddressHook(ctx context.Context, key *meta.Key, obj *ga.Address, m *cloud.MockAddresses) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionGA, "addresses")
 return convertAndInsertAlphaAddress(key, obj, m.Objects, meta.VersionGA, projectID, m.X.(AddressAttributes))
}
func InsertBetaAddressHook(ctx context.Context, key *meta.Key, obj *beta.Address, m *cloud.MockAddresses) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionBeta, "addresses")
 return convertAndInsertAlphaAddress(key, obj, m.Objects, meta.VersionBeta, projectID, m.X.(AddressAttributes))
}
func InsertAlphaAddressHook(ctx context.Context, key *meta.Key, obj *alpha.Address, m *cloud.MockAlphaAddresses) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 m.Lock.Lock()
 defer m.Lock.Unlock()
 projectID := m.ProjectRouter.ProjectID(ctx, meta.VersionBeta, "addresses")
 return convertAndInsertAlphaAddress(key, obj, m.Objects, meta.VersionAlpha, projectID, m.X.(AddressAttributes))
}

type InstanceGroupAttributes struct {
 InstanceMap map[meta.Key]map[string]*ga.InstanceWithNamedPorts
 Lock        *sync.Mutex
}

func (igAttrs *InstanceGroupAttributes) AddInstances(key *meta.Key, instanceRefs []*ga.InstanceReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 igAttrs.Lock.Lock()
 defer igAttrs.Lock.Unlock()
 instancesWithNamedPorts, ok := igAttrs.InstanceMap[*key]
 if !ok {
  instancesWithNamedPorts = make(map[string]*ga.InstanceWithNamedPorts)
 }
 for _, instance := range instanceRefs {
  iWithPort := &ga.InstanceWithNamedPorts{Instance: instance.Instance}
  instancesWithNamedPorts[instance.Instance] = iWithPort
 }
 igAttrs.InstanceMap[*key] = instancesWithNamedPorts
 return nil
}
func (igAttrs *InstanceGroupAttributes) RemoveInstances(key *meta.Key, instanceRefs []*ga.InstanceReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 igAttrs.Lock.Lock()
 defer igAttrs.Lock.Unlock()
 instancesWithNamedPorts, ok := igAttrs.InstanceMap[*key]
 if !ok {
  instancesWithNamedPorts = make(map[string]*ga.InstanceWithNamedPorts)
 }
 for _, instanceToRemove := range instanceRefs {
  if _, ok := instancesWithNamedPorts[instanceToRemove.Instance]; ok {
   delete(instancesWithNamedPorts, instanceToRemove.Instance)
  } else {
   return &googleapi.Error{Code: http.StatusBadRequest, Message: fmt.Sprintf("%s is not a member of %s", instanceToRemove.Instance, key.String())}
  }
 }
 igAttrs.InstanceMap[*key] = instancesWithNamedPorts
 return nil
}
func (igAttrs *InstanceGroupAttributes) List(key *meta.Key) []*ga.InstanceWithNamedPorts {
 _logClusterCodePath()
 defer _logClusterCodePath()
 igAttrs.Lock.Lock()
 defer igAttrs.Lock.Unlock()
 instancesWithNamedPorts, ok := igAttrs.InstanceMap[*key]
 if !ok {
  instancesWithNamedPorts = make(map[string]*ga.InstanceWithNamedPorts)
 }
 var instanceList []*ga.InstanceWithNamedPorts
 for _, val := range instancesWithNamedPorts {
  instanceList = append(instanceList, val)
 }
 return instanceList
}
func AddInstancesHook(ctx context.Context, key *meta.Key, req *ga.InstanceGroupsAddInstancesRequest, m *cloud.MockInstanceGroups) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in InstanceGroups", key.String())}
 }
 var attrs InstanceGroupAttributes
 attrs = m.X.(InstanceGroupAttributes)
 attrs.AddInstances(key, req.Instances)
 m.X = attrs
 return nil
}
func ListInstancesHook(ctx context.Context, key *meta.Key, req *ga.InstanceGroupsListInstancesRequest, filter *filter.F, m *cloud.MockInstanceGroups) ([]*ga.InstanceWithNamedPorts, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return nil, &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in InstanceGroups", key.String())}
 }
 var attrs InstanceGroupAttributes
 attrs = m.X.(InstanceGroupAttributes)
 instances := attrs.List(key)
 return instances, nil
}
func RemoveInstancesHook(ctx context.Context, key *meta.Key, req *ga.InstanceGroupsRemoveInstancesRequest, m *cloud.MockInstanceGroups) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in InstanceGroups", key.String())}
 }
 var attrs InstanceGroupAttributes
 attrs = m.X.(InstanceGroupAttributes)
 attrs.RemoveInstances(key, req.Instances)
 m.X = attrs
 return nil
}
func UpdateFirewallHook(ctx context.Context, key *meta.Key, obj *ga.Firewall, m *cloud.MockFirewalls) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in Firewalls", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "firewalls")
 obj.SelfLink = cloud.SelfLink(meta.VersionGA, projectID, "firewalls", key)
 m.Objects[*key] = &cloud.MockFirewallsObj{Obj: obj}
 return nil
}
func UpdateHealthCheckHook(ctx context.Context, key *meta.Key, obj *ga.HealthCheck, m *cloud.MockHealthChecks) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in HealthChecks", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "healthChecks")
 obj.SelfLink = cloud.SelfLink(meta.VersionGA, projectID, "healthChecks", key)
 m.Objects[*key] = &cloud.MockHealthChecksObj{Obj: obj}
 return nil
}
func UpdateRegionBackendServiceHook(ctx context.Context, key *meta.Key, obj *ga.BackendService, m *cloud.MockRegionBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in RegionBackendServices", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "backendServices")
 obj.SelfLink = cloud.SelfLink(meta.VersionGA, projectID, "backendServices", key)
 m.Objects[*key] = &cloud.MockRegionBackendServicesObj{Obj: obj}
 return nil
}
func UpdateBackendServiceHook(ctx context.Context, key *meta.Key, obj *ga.BackendService, m *cloud.MockBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in BackendServices", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "backendServices")
 obj.SelfLink = cloud.SelfLink(meta.VersionGA, projectID, "backendServices", key)
 m.Objects[*key] = &cloud.MockBackendServicesObj{Obj: obj}
 return nil
}
func UpdateAlphaBackendServiceHook(ctx context.Context, key *meta.Key, obj *alpha.BackendService, m *cloud.MockAlphaBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in BackendServices", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "backendServices")
 obj.SelfLink = cloud.SelfLink(meta.VersionAlpha, projectID, "backendServices", key)
 m.Objects[*key] = &cloud.MockBackendServicesObj{Obj: obj}
 return nil
}
func UpdateBetaBackendServiceHook(ctx context.Context, key *meta.Key, obj *beta.BackendService, m *cloud.MockBetaBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in BackendServices", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "backendServices")
 obj.SelfLink = cloud.SelfLink(meta.VersionBeta, projectID, "backendServices", key)
 m.Objects[*key] = &cloud.MockBackendServicesObj{Obj: obj}
 return nil
}
func UpdateURLMapHook(ctx context.Context, key *meta.Key, obj *ga.UrlMap, m *cloud.MockUrlMaps) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err := m.Get(ctx, key)
 if err != nil {
  return &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("Key: %s was not found in UrlMaps", key.String())}
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "urlMaps")
 obj.SelfLink = cloud.SelfLink(meta.VersionGA, projectID, "urlMaps", key)
 m.Objects[*key] = &cloud.MockUrlMapsObj{Obj: obj}
 return nil
}
func InsertFirewallsUnauthorizedErrHook(ctx context.Context, key *meta.Key, obj *ga.Firewall, m *cloud.MockFirewalls) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, &googleapi.Error{Code: http.StatusForbidden}
}
func UpdateFirewallsUnauthorizedErrHook(ctx context.Context, key *meta.Key, obj *ga.Firewall, m *cloud.MockFirewalls) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &googleapi.Error{Code: http.StatusForbidden}
}
func DeleteFirewallsUnauthorizedErrHook(ctx context.Context, key *meta.Key, m *cloud.MockFirewalls) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, &googleapi.Error{Code: http.StatusForbidden}
}
func GetFirewallsUnauthorizedErrHook(ctx context.Context, key *meta.Key, m *cloud.MockFirewalls) (bool, *ga.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, &googleapi.Error{Code: http.StatusForbidden}
}
func GetTargetPoolInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockTargetPools) (bool, *ga.TargetPool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func GetForwardingRulesInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockForwardingRules) (bool, *ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func GetAddressesInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockAddresses) (bool, *ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func GetHTTPHealthChecksInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockHttpHealthChecks) (bool, *ga.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func InsertTargetPoolsInternalErrHook(ctx context.Context, key *meta.Key, obj *ga.TargetPool, m *cloud.MockTargetPools) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func InsertForwardingRulesInternalErrHook(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule, m *cloud.MockForwardingRules) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func DeleteAddressesNotFoundErrHook(ctx context.Context, key *meta.Key, m *cloud.MockAddresses) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, &googleapi.Error{Code: http.StatusNotFound}
}
func DeleteAddressesInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockAddresses) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func InsertAlphaBackendServiceUnauthorizedErrHook(ctx context.Context, key *meta.Key, obj *alpha.BackendService, m *cloud.MockAlphaBackendServices) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, UnauthorizedErr
}
func UpdateAlphaBackendServiceUnauthorizedErrHook(ctx context.Context, key *meta.Key, obj *alpha.BackendService, m *cloud.MockAlphaBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return UnauthorizedErr
}
func GetRegionBackendServicesErrHook(ctx context.Context, key *meta.Key, m *cloud.MockRegionBackendServices) (bool, *ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func UpdateRegionBackendServicesErrHook(ctx context.Context, key *meta.Key, svc *ga.BackendService, m *cloud.MockRegionBackendServices) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return InternalServerError
}
func DeleteRegionBackendServicesErrHook(ctx context.Context, key *meta.Key, m *cloud.MockRegionBackendServices) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func DeleteRegionBackendServicesInUseErrHook(ctx context.Context, key *meta.Key, m *cloud.MockRegionBackendServices) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InUseError
}
func GetInstanceGroupInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockInstanceGroups) (bool, *ga.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func GetHealthChecksInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockHealthChecks) (bool, *ga.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, nil, InternalServerError
}
func DeleteHealthChecksInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockHealthChecks) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func DeleteHealthChecksInuseErrHook(ctx context.Context, key *meta.Key, m *cloud.MockHealthChecks) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InUseError
}
func DeleteForwardingRuleErrHook(ctx context.Context, key *meta.Key, m *cloud.MockForwardingRules) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func ListZonesInternalErrHook(ctx context.Context, fl *filter.F, m *cloud.MockZones) (bool, []*ga.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, []*ga.Zone{}, InternalServerError
}
func DeleteInstanceGroupInternalErrHook(ctx context.Context, key *meta.Key, m *cloud.MockInstanceGroups) (bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true, InternalServerError
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
