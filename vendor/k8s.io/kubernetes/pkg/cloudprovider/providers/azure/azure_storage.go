package azure

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
	"k8s.io/klog"
)

const (
	defaultStorageAccountType      = string(storage.StandardLRS)
	defaultStorageAccountKind      = storage.StorageV2
	fileShareAccountNamePrefix     = "f"
	sharedDiskAccountNamePrefix    = "ds"
	dedicatedDiskAccountNamePrefix = "dd"
)

func (az *Cloud) CreateFileShare(shareName, accountName, accountType, accountKind, resourceGroup, location string, requestGiB int) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if resourceGroup == "" {
		resourceGroup = az.resourceGroup
	}
	account, key, err := az.ensureStorageAccount(accountName, accountType, accountKind, resourceGroup, location, fileShareAccountNamePrefix)
	if err != nil {
		return "", "", fmt.Errorf("could not get storage key for storage account %s: %v", accountName, err)
	}
	if err := az.createFileShare(account, key, shareName, requestGiB); err != nil {
		return "", "", fmt.Errorf("failed to create share %s in account %s: %v", shareName, account, err)
	}
	klog.V(4).Infof("created share %s in account %s", shareName, account)
	return account, key, nil
}
func (az *Cloud) DeleteFileShare(accountName, accountKey, shareName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := az.deleteFileShare(accountName, accountKey, shareName); err != nil {
		return err
	}
	klog.V(4).Infof("share %s deleted", shareName)
	return nil
}
func (az *Cloud) ResizeFileShare(accountName, accountKey, name string, sizeGiB int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return az.resizeFileShare(accountName, accountKey, name, sizeGiB)
}
