package vclib

import (
 "context"
 "fmt"
 "github.com/vmware/govmomi/object"
 "github.com/vmware/govmomi/property"
 "github.com/vmware/govmomi/vim25/mo"
 "github.com/vmware/govmomi/vim25/soap"
 "github.com/vmware/govmomi/vim25/types"
 "k8s.io/klog"
)

type Datastore struct {
 *object.Datastore
 Datacenter *Datacenter
}
type DatastoreInfo struct {
 *Datastore
 Info *types.DatastoreInfo
}

func (di DatastoreInfo) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("Datastore: %+v, datastore URL: %s", di.Datastore, di.Info.Url)
}
func (ds *Datastore) CreateDirectory(ctx context.Context, directoryPath string, createParents bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fileManager := object.NewFileManager(ds.Client())
 err := fileManager.MakeDirectory(ctx, directoryPath, ds.Datacenter.Datacenter, createParents)
 if err != nil {
  if soap.IsSoapFault(err) {
   soapFault := soap.ToSoapFault(err)
   if _, ok := soapFault.VimFault().(types.FileAlreadyExists); ok {
    return ErrFileAlreadyExist
   }
  }
  return err
 }
 klog.V(LogLevel).Infof("Created dir with path as %+q", directoryPath)
 return nil
}
func (ds *Datastore) GetType(ctx context.Context) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var dsMo mo.Datastore
 pc := property.DefaultCollector(ds.Client())
 err := pc.RetrieveOne(ctx, ds.Datastore.Reference(), []string{"summary"}, &dsMo)
 if err != nil {
  klog.Errorf("Failed to retrieve datastore summary property. err: %v", err)
  return "", err
 }
 return dsMo.Summary.Type, nil
}
func (ds *Datastore) IsCompatibleWithStoragePolicy(ctx context.Context, storagePolicyID string) (bool, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 pbmClient, err := NewPbmClient(ctx, ds.Client())
 if err != nil {
  klog.Errorf("Failed to get new PbmClient Object. err: %v", err)
  return false, "", err
 }
 return pbmClient.IsDatastoreCompatible(ctx, storagePolicyID, ds)
}
