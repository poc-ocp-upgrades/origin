package vclib

import (
 "context"
 "fmt"
 "github.com/vmware/govmomi/pbm"
 "k8s.io/klog"
 pbmtypes "github.com/vmware/govmomi/pbm/types"
 "github.com/vmware/govmomi/vim25"
)

type PbmClient struct{ *pbm.Client }

func NewPbmClient(ctx context.Context, client *vim25.Client) (*PbmClient, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pbmClient, err := pbm.NewClient(ctx, client)
 if err != nil {
  klog.Errorf("Failed to create new Pbm Client. err: %+v", err)
  return nil, err
 }
 return &PbmClient{pbmClient}, nil
}
func (pbmClient *PbmClient) IsDatastoreCompatible(ctx context.Context, storagePolicyID string, datastore *Datastore) (bool, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 faultMessage := ""
 placementHub := pbmtypes.PbmPlacementHub{HubType: datastore.Reference().Type, HubId: datastore.Reference().Value}
 hubs := []pbmtypes.PbmPlacementHub{placementHub}
 req := []pbmtypes.BasePbmPlacementRequirement{&pbmtypes.PbmPlacementCapabilityProfileRequirement{ProfileId: pbmtypes.PbmProfileId{UniqueId: storagePolicyID}}}
 compatibilityResult, err := pbmClient.CheckRequirements(ctx, hubs, nil, req)
 if err != nil {
  klog.Errorf("Error occurred for CheckRequirements call. err %+v", err)
  return false, "", err
 }
 if compatibilityResult != nil && len(compatibilityResult) > 0 {
  compatibleHubs := compatibilityResult.CompatibleDatastores()
  if compatibleHubs != nil && len(compatibleHubs) > 0 {
   return true, "", nil
  }
  dsName, err := datastore.ObjectName(ctx)
  if err != nil {
   klog.Errorf("Failed to get datastore ObjectName")
   return false, "", err
  }
  if compatibilityResult[0].Error[0].LocalizedMessage == "" {
   faultMessage = "Datastore: " + dsName + " is not compatible with the storage policy."
  } else {
   faultMessage = "Datastore: " + dsName + " is not compatible with the storage policy. LocalizedMessage: " + compatibilityResult[0].Error[0].LocalizedMessage + "\n"
  }
  return false, faultMessage, nil
 }
 return false, "", fmt.Errorf("compatibilityResult is nil or empty")
}
func (pbmClient *PbmClient) GetCompatibleDatastores(ctx context.Context, dc *Datacenter, storagePolicyID string, datastores []*DatastoreInfo) ([]*DatastoreInfo, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var (
  dsMorNameMap                                = getDsMorNameMap(ctx, datastores)
  localizedMessagesForNotCompatibleDatastores = ""
 )
 compatibilityResult, err := pbmClient.GetPlacementCompatibilityResult(ctx, storagePolicyID, datastores)
 if err != nil {
  klog.Errorf("Error occurred while retrieving placement compatibility result for datastores: %+v with storagePolicyID: %s. err: %+v", datastores, storagePolicyID, err)
  return nil, "", err
 }
 compatibleHubs := compatibilityResult.CompatibleDatastores()
 var compatibleDatastoreList []*DatastoreInfo
 for _, hub := range compatibleHubs {
  compatibleDatastoreList = append(compatibleDatastoreList, getDatastoreFromPlacementHub(datastores, hub))
 }
 for _, res := range compatibilityResult {
  for _, err := range res.Error {
   dsName := dsMorNameMap[res.Hub.HubId]
   localizedMessage := ""
   if err.LocalizedMessage != "" {
    localizedMessage = "Datastore: " + dsName + " not compatible with the storage policy. LocalizedMessage: " + err.LocalizedMessage + "\n"
   } else {
    localizedMessage = "Datastore: " + dsName + " not compatible with the storage policy. \n"
   }
   localizedMessagesForNotCompatibleDatastores += localizedMessage
  }
 }
 if len(compatibleHubs) < 1 {
  klog.Errorf("No compatible datastores found that satisfy the storage policy requirements: %s", storagePolicyID)
  return nil, localizedMessagesForNotCompatibleDatastores, fmt.Errorf("No compatible datastores found that satisfy the storage policy requirements")
 }
 return compatibleDatastoreList, localizedMessagesForNotCompatibleDatastores, nil
}
func (pbmClient *PbmClient) GetPlacementCompatibilityResult(ctx context.Context, storagePolicyID string, datastore []*DatastoreInfo) (pbm.PlacementCompatibilityResult, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var hubs []pbmtypes.PbmPlacementHub
 for _, ds := range datastore {
  hubs = append(hubs, pbmtypes.PbmPlacementHub{HubType: ds.Reference().Type, HubId: ds.Reference().Value})
 }
 req := []pbmtypes.BasePbmPlacementRequirement{&pbmtypes.PbmPlacementCapabilityProfileRequirement{ProfileId: pbmtypes.PbmProfileId{UniqueId: storagePolicyID}}}
 res, err := pbmClient.CheckRequirements(ctx, hubs, nil, req)
 if err != nil {
  klog.Errorf("Error occurred for CheckRequirements call. err: %+v", err)
  return nil, err
 }
 return res, nil
}
func getDatastoreFromPlacementHub(datastore []*DatastoreInfo, pbmPlacementHub pbmtypes.PbmPlacementHub) *DatastoreInfo {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, ds := range datastore {
  if ds.Reference().Type == pbmPlacementHub.HubType && ds.Reference().Value == pbmPlacementHub.HubId {
   return ds
  }
 }
 return nil
}
func getDsMorNameMap(ctx context.Context, datastores []*DatastoreInfo) map[string]string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 dsMorNameMap := make(map[string]string)
 for _, ds := range datastores {
  dsObjectName, err := ds.ObjectName(ctx)
  if err == nil {
   dsMorNameMap[ds.Reference().Value] = dsObjectName
  } else {
   klog.Errorf("Error occurred while getting datastore object name. err: %+v", err)
  }
 }
 return dsMorNameMap
}
