package azure

import (
 "fmt"
 azs "github.com/Azure/azure-sdk-for-go/storage"
 "github.com/Azure/go-autorest/autorest/azure"
 "k8s.io/klog"
)

const (
 useHTTPS = true
)

type FileClient interface {
 createFileShare(accountName, accountKey, name string, sizeGiB int) error
 deleteFileShare(accountName, accountKey, name string) error
 resizeFileShare(accountName, accountKey, name string, sizeGiB int) error
}

func (az *Cloud) createFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return az.FileClient.createFileShare(accountName, accountKey, name, sizeGiB)
}
func (az *Cloud) deleteFileShare(accountName, accountKey, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return az.FileClient.deleteFileShare(accountName, accountKey, name)
}
func (az *Cloud) resizeFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return az.FileClient.resizeFileShare(accountName, accountKey, name, sizeGiB)
}

type azureFileClient struct{ env azure.Environment }

func (f *azureFileClient) createFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fileClient, err := f.getFileSvcClient(accountName, accountKey)
 if err != nil {
  return err
 }
 share := fileClient.GetShareReference(name)
 share.Properties.Quota = sizeGiB
 if err = share.Create(nil); err != nil {
  return fmt.Errorf("failed to create file share, err: %v", err)
 }
 return nil
}
func (f *azureFileClient) deleteFileShare(accountName, accountKey, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fileClient, err := f.getFileSvcClient(accountName, accountKey)
 if err != nil {
  return err
 }
 return fileClient.GetShareReference(name).Delete(nil)
}
func (f *azureFileClient) resizeFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fileClient, err := f.getFileSvcClient(accountName, accountKey)
 if err != nil {
  return err
 }
 share := fileClient.GetShareReference(name)
 if share.Properties.Quota >= sizeGiB {
  klog.Warningf("file share size(%dGi) is already greater or equal than requested size(%dGi), accountName: %s, shareName: %s", share.Properties.Quota, sizeGiB, accountName, name)
  return nil
 }
 share.Properties.Quota = sizeGiB
 if err = share.SetProperties(nil); err != nil {
  return fmt.Errorf("failed to set quota on file share %s, err: %v", name, err)
 }
 klog.V(4).Infof("resize file share completed, accountName: %s, shareName: %s, sizeGiB: %d", accountName, name, sizeGiB)
 return nil
}
func (f *azureFileClient) getFileSvcClient(accountName, accountKey string) (*azs.FileServiceClient, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fileClient, err := azs.NewClient(accountName, accountKey, f.env.StorageEndpointSuffix, azs.DefaultAPIVersion, useHTTPS)
 if err != nil {
  return nil, fmt.Errorf("error creating azure client: %v", err)
 }
 fc := fileClient.GetFileService()
 return &fc, nil
}
