package azure

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/rubiojr/go-vhd/vhd"
	kwait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	vhdContainerName         = "vhds"
	useHTTPSForBlobBasedDisk = true
	blobServiceName          = "blob"
)

type storageAccountState struct {
	name                    string
	saType                  storage.SkuName
	key                     string
	diskCount               int32
	isValidating            int32
	defaultContainerCreated bool
}
type BlobDiskController struct {
	common   *controllerCommon
	accounts map[string]*storageAccountState
}

var (
	accountsLock = &sync.Mutex{}
)

func (c *BlobDiskController) initStorageAccounts() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accountsLock.Lock()
	defer accountsLock.Unlock()
	if c.accounts == nil {
		accounts, err := c.getAllStorageAccounts()
		if err != nil {
			klog.Errorf("azureDisk - getAllStorageAccounts error: %v", err)
			c.accounts = make(map[string]*storageAccountState)
		}
		c.accounts = accounts
	}
}
func (c *BlobDiskController) CreateVolume(blobName, accountName, accountType, location string, requestGB int) (string, string, int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	account, key, err := c.common.cloud.ensureStorageAccount(accountName, accountType, string(defaultStorageAccountKind), c.common.resourceGroup, location, dedicatedDiskAccountNamePrefix)
	if err != nil {
		return "", "", 0, fmt.Errorf("could not get storage key for storage account %s: %v", accountName, err)
	}
	client, err := azstorage.NewBasicClientOnSovereignCloud(account, key, c.common.cloud.Environment)
	if err != nil {
		return "", "", 0, err
	}
	blobClient := client.GetBlobService()
	diskName, diskURI, err := c.createVHDBlobDisk(blobClient, account, blobName, vhdContainerName, int64(requestGB))
	if err != nil {
		return "", "", 0, err
	}
	klog.V(4).Infof("azureDisk - created vhd blob uri: %s", diskURI)
	return diskName, diskURI, requestGB, err
}
func (c *BlobDiskController) DeleteVolume(diskURI string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("azureDisk - begin to delete volume %s", diskURI)
	accountName, blob, err := c.common.cloud.getBlobNameAndAccountFromURI(diskURI)
	if err != nil {
		return fmt.Errorf("failed to parse vhd URI %v", err)
	}
	key, err := c.common.cloud.getStorageAccesskey(accountName, c.common.resourceGroup)
	if err != nil {
		return fmt.Errorf("no key for storage account %s, err %v", accountName, err)
	}
	err = c.common.cloud.deleteVhdBlob(accountName, key, blob)
	if err != nil {
		klog.Warningf("azureDisk - failed to delete blob %s err: %v", diskURI, err)
		detail := err.Error()
		if strings.Contains(detail, errLeaseIDMissing) {
			return volume.NewDeletedVolumeInUseError(fmt.Sprintf("disk %q is still in use while being deleted", diskURI))
		}
		return fmt.Errorf("failed to delete vhd %v, account %s, blob %s, err: %v", diskURI, accountName, blob, err)
	}
	klog.V(4).Infof("azureDisk - blob %s deleted", diskURI)
	return nil
}
func (c *BlobDiskController) getBlobNameAndAccountFromURI(diskURI string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme := "http"
	if useHTTPSForBlobBasedDisk {
		scheme = "https"
	}
	host := fmt.Sprintf("%s://(.*).%s.%s", scheme, blobServiceName, c.common.storageEndpointSuffix)
	reStr := fmt.Sprintf("%s/%s/(.*)", host, vhdContainerName)
	re := regexp.MustCompile(reStr)
	res := re.FindSubmatch([]byte(diskURI))
	if len(res) < 3 {
		return "", "", fmt.Errorf("invalid vhd URI for regex %s: %s", reStr, diskURI)
	}
	return string(res[1]), string(res[2]), nil
}
func (c *BlobDiskController) createVHDBlobDisk(blobClient azstorage.BlobStorageClient, accountName, vhdName, containerName string, sizeGB int64) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	container := blobClient.GetContainerReference(containerName)
	size := 1024 * 1024 * 1024 * sizeGB
	vhdSize := size + vhd.VHD_HEADER_SIZE
	vhdName = vhdName + ".vhd"
	tags := make(map[string]string)
	tags["createdby"] = "k8sAzureDataDisk"
	klog.V(4).Infof("azureDisk - creating page blob %s in container %s account %s", vhdName, containerName, accountName)
	blob := container.GetBlobReference(vhdName)
	blob.Properties.ContentLength = vhdSize
	blob.Metadata = tags
	err := blob.PutPageBlob(nil)
	if err != nil {
		detail := err.Error()
		if strings.Contains(detail, errContainerNotFound) {
			err = container.Create(&azstorage.CreateContainerOptions{Access: azstorage.ContainerAccessTypePrivate})
			if err == nil {
				err = blob.PutPageBlob(nil)
			}
		}
	}
	if err != nil {
		return "", "", fmt.Errorf("failed to put page blob %s in container %s: %v", vhdName, containerName, err)
	}
	h, err := createVHDHeader(uint64(size))
	if err != nil {
		blob.DeleteIfExists(nil)
		return "", "", fmt.Errorf("failed to create vhd header, err: %v", err)
	}
	blobRange := azstorage.BlobRange{Start: uint64(size), End: uint64(vhdSize - 1)}
	if err = blob.WriteRange(blobRange, bytes.NewBuffer(h[:vhd.VHD_HEADER_SIZE]), nil); err != nil {
		klog.Infof("azureDisk - failed to put header page for data disk %s in container %s account %s, error was %s\n", vhdName, containerName, accountName, err.Error())
		return "", "", err
	}
	scheme := "http"
	if useHTTPSForBlobBasedDisk {
		scheme = "https"
	}
	host := fmt.Sprintf("%s://%s.%s.%s", scheme, accountName, blobServiceName, c.common.storageEndpointSuffix)
	uri := fmt.Sprintf("%s/%s/%s", host, containerName, vhdName)
	return vhdName, uri, nil
}
func (c *BlobDiskController) deleteVhdBlob(accountName, accountKey, blobName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client, err := azstorage.NewBasicClientOnSovereignCloud(accountName, accountKey, c.common.cloud.Environment)
	if err != nil {
		return err
	}
	blobSvc := client.GetBlobService()
	container := blobSvc.GetContainerReference(vhdContainerName)
	blob := container.GetBlobReference(blobName)
	return blob.Delete(nil)
}
func (c *BlobDiskController) CreateBlobDisk(dataDiskName string, storageAccountType storage.SkuName, sizeGB int) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("azureDisk - creating blob data disk named:%s on StorageAccountType:%s", dataDiskName, storageAccountType)
	c.initStorageAccounts()
	storageAccountName, err := c.findSANameForDisk(storageAccountType)
	if err != nil {
		return "", err
	}
	blobClient, err := c.getBlobSvcClient(storageAccountName)
	if err != nil {
		return "", err
	}
	_, diskURI, err := c.createVHDBlobDisk(blobClient, storageAccountName, dataDiskName, vhdContainerName, int64(sizeGB))
	if err != nil {
		return "", err
	}
	atomic.AddInt32(&c.accounts[storageAccountName].diskCount, 1)
	return diskURI, nil
}
func (c *BlobDiskController) DeleteBlobDisk(diskURI string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storageAccountName, vhdName, err := diskNameandSANameFromURI(diskURI)
	if err != nil {
		return err
	}
	_, ok := c.accounts[storageAccountName]
	if !ok {
		klog.V(4).Infof("azureDisk - deleting volume %s", diskURI)
		return c.DeleteVolume(diskURI)
	}
	blobSvc, err := c.getBlobSvcClient(storageAccountName)
	if err != nil {
		return err
	}
	klog.V(4).Infof("azureDisk - About to delete vhd file %s on storage account %s container %s", vhdName, storageAccountName, vhdContainerName)
	container := blobSvc.GetContainerReference(vhdContainerName)
	blob := container.GetBlobReference(vhdName)
	_, err = blob.DeleteIfExists(nil)
	if c.accounts[storageAccountName].diskCount == -1 {
		if diskCount, err := c.getDiskCount(storageAccountName); err != nil {
			c.accounts[storageAccountName].diskCount = int32(diskCount)
		} else {
			klog.Warningf("azureDisk - failed to get disk count for %s however the delete disk operation was ok", storageAccountName)
			return nil
		}
	}
	atomic.AddInt32(&c.accounts[storageAccountName].diskCount, -1)
	return err
}
func (c *BlobDiskController) getStorageAccountKey(SAName string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if account, exists := c.accounts[SAName]; exists && account.key != "" {
		return c.accounts[SAName].key, nil
	}
	ctx, cancel := getContextWithCancel()
	defer cancel()
	listKeysResult, err := c.common.cloud.StorageAccountClient.ListKeys(ctx, c.common.resourceGroup, SAName)
	if err != nil {
		return "", err
	}
	if listKeysResult.Keys == nil {
		return "", fmt.Errorf("azureDisk - empty listKeysResult in storage account:%s keys", SAName)
	}
	for _, v := range *listKeysResult.Keys {
		if v.Value != nil && *v.Value == "key1" {
			if _, ok := c.accounts[SAName]; !ok {
				klog.Warningf("azureDisk - account %s was not cached while getting keys", SAName)
				return *v.Value, nil
			}
		}
		c.accounts[SAName].key = *v.Value
		return c.accounts[SAName].key, nil
	}
	return "", fmt.Errorf("couldn't find key named key1 in storage account:%s keys", SAName)
}
func (c *BlobDiskController) getBlobSvcClient(SAName string) (azstorage.BlobStorageClient, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := ""
	var client azstorage.Client
	var blobSvc azstorage.BlobStorageClient
	var err error
	if key, err = c.getStorageAccountKey(SAName); err != nil {
		return blobSvc, err
	}
	if client, err = azstorage.NewBasicClientOnSovereignCloud(SAName, key, c.common.cloud.Environment); err != nil {
		return blobSvc, err
	}
	blobSvc = client.GetBlobService()
	return blobSvc, nil
}
func (c *BlobDiskController) ensureDefaultContainer(storageAccountName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	var blobSvc azstorage.BlobStorageClient
	if v, ok := c.accounts[storageAccountName]; ok && v.defaultContainerCreated {
		return nil
	}
	bExist, provisionState, _ := c.getStorageAccountState(storageAccountName)
	if !bExist {
		return fmt.Errorf("azureDisk - account %s does not exist while trying to create/ensure default container", storageAccountName)
	}
	if provisionState != storage.Succeeded {
		counter := 1
		for swapped := atomic.CompareAndSwapInt32(&c.accounts[storageAccountName].isValidating, 0, 1); swapped != true; {
			time.Sleep(3 * time.Second)
			counter = counter + 1
			if counter >= 20 {
				return fmt.Errorf("azureDisk - timeout waiting to acquire lock to validate account:%s readiness", storageAccountName)
			}
		}
		defer func() {
			c.accounts[storageAccountName].isValidating = 0
		}()
		if v, ok := c.accounts[storageAccountName]; ok && v.defaultContainerCreated {
			return nil
		}
		err = kwait.ExponentialBackoff(defaultBackOff, func() (bool, error) {
			_, provisionState, err := c.getStorageAccountState(storageAccountName)
			if err != nil {
				klog.V(4).Infof("azureDisk - GetStorageAccount:%s err %s", storageAccountName, err.Error())
				return false, nil
			}
			if provisionState == storage.Succeeded {
				return true, nil
			}
			klog.V(4).Infof("azureDisk - GetStorageAccount:%s not ready yet (not flagged Succeeded by ARM)", storageAccountName)
			return false, nil
		})
		if err != nil {
			if err == kwait.ErrWaitTimeout {
				return fmt.Errorf("azureDisk - timed out waiting for storage account %s to become ready", storageAccountName)
			}
			return err
		}
	}
	if blobSvc, err = c.getBlobSvcClient(storageAccountName); err != nil {
		return err
	}
	container := blobSvc.GetContainerReference(vhdContainerName)
	bCreated, err := container.CreateIfNotExists(&azstorage.CreateContainerOptions{Access: azstorage.ContainerAccessTypePrivate})
	if err != nil {
		return err
	}
	if bCreated {
		klog.V(2).Infof("azureDisk - storage account:%s had no default container(%s) and it was created \n", storageAccountName, vhdContainerName)
	}
	c.accounts[storageAccountName].defaultContainerCreated = true
	return nil
}
func (c *BlobDiskController) getDiskCount(SAName string) (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if c.accounts[SAName].diskCount != -1 {
		return int(c.accounts[SAName].diskCount), nil
	}
	var err error
	var blobSvc azstorage.BlobStorageClient
	if err = c.ensureDefaultContainer(SAName); err != nil {
		return 0, err
	}
	if blobSvc, err = c.getBlobSvcClient(SAName); err != nil {
		return 0, err
	}
	params := azstorage.ListBlobsParameters{}
	container := blobSvc.GetContainerReference(vhdContainerName)
	response, err := container.ListBlobs(params)
	if err != nil {
		return 0, err
	}
	klog.V(4).Infof("azure-Disk -  refreshed data count for account %s and found %v", SAName, len(response.Blobs))
	c.accounts[SAName].diskCount = int32(len(response.Blobs))
	return int(c.accounts[SAName].diskCount), nil
}
func (c *BlobDiskController) getAllStorageAccounts() (map[string]*storageAccountState, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	accountListResult, err := c.common.cloud.StorageAccountClient.ListByResourceGroup(ctx, c.common.resourceGroup)
	if err != nil {
		return nil, err
	}
	if accountListResult.Value == nil {
		return nil, fmt.Errorf("azureDisk - empty accountListResult")
	}
	accounts := make(map[string]*storageAccountState)
	for _, v := range *accountListResult.Value {
		if v.Name == nil || v.Sku == nil {
			klog.Info("azureDisk - accountListResult Name or Sku is nil")
			continue
		}
		if !strings.HasPrefix(*v.Name, sharedDiskAccountNamePrefix) {
			continue
		}
		klog.Infof("azureDisk - identified account %s as part of shared PVC accounts", *v.Name)
		sastate := &storageAccountState{name: *v.Name, saType: (*v.Sku).Name, diskCount: -1}
		accounts[*v.Name] = sastate
	}
	return accounts, nil
}
func (c *BlobDiskController) createStorageAccount(storageAccountName string, storageAccountType storage.SkuName, location string, checkMaxAccounts bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bExist, _, _ := c.getStorageAccountState(storageAccountName)
	if bExist {
		newAccountState := &storageAccountState{diskCount: -1, saType: storageAccountType, name: storageAccountName}
		c.addAccountState(storageAccountName, newAccountState)
	}
	if !bExist {
		if len(c.accounts) == maxStorageAccounts && checkMaxAccounts {
			return fmt.Errorf("azureDisk - can not create new storage account, current storage accounts count:%v Max is:%v", len(c.accounts), maxStorageAccounts)
		}
		klog.V(2).Infof("azureDisk - Creating storage account %s type %s", storageAccountName, string(storageAccountType))
		cp := storage.AccountCreateParameters{Sku: &storage.Sku{Name: storageAccountType}, Kind: defaultStorageAccountKind, Tags: map[string]*string{"created-by": to.StringPtr("azure-dd")}, Location: &location}
		ctx, cancel := getContextWithCancel()
		defer cancel()
		_, err := c.common.cloud.StorageAccountClient.Create(ctx, c.common.resourceGroup, storageAccountName, cp)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("Create Storage Account: %s, error: %s", storageAccountName, err))
		}
		newAccountState := &storageAccountState{diskCount: -1, saType: storageAccountType, name: storageAccountName}
		c.addAccountState(storageAccountName, newAccountState)
	}
	return c.ensureDefaultContainer(storageAccountName)
}
func (c *BlobDiskController) findSANameForDisk(storageAccountType storage.SkuName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	maxDiskCount := maxDisksPerStorageAccounts
	SAName := ""
	totalDiskCounts := 0
	countAccounts := 0
	for _, v := range c.accounts {
		if !strings.HasPrefix(v.name, sharedDiskAccountNamePrefix) {
			continue
		}
		if v.saType == storageAccountType {
			dCount, err := c.getDiskCount(v.name)
			if err != nil {
				return "", err
			}
			totalDiskCounts = totalDiskCounts + dCount
			countAccounts = countAccounts + 1
			if dCount == 0 {
				klog.V(2).Infof("azureDisk - account %s identified for a new disk  is because it has 0 allocated disks", v.name)
				return v.name, nil
			}
			if dCount < maxDiskCount {
				maxDiskCount = dCount
				SAName = v.name
			}
		}
	}
	if SAName == "" {
		klog.V(2).Infof("azureDisk - failed to identify a suitable account for new disk and will attempt to create new account")
		SAName = generateStorageAccountName(sharedDiskAccountNamePrefix)
		err := c.createStorageAccount(SAName, storageAccountType, c.common.location, true)
		if err != nil {
			return "", err
		}
		return SAName, nil
	}
	disksAfter := totalDiskCounts + 1
	avgUtilization := float64(disksAfter) / float64(countAccounts*maxDisksPerStorageAccounts)
	aboveAvg := (avgUtilization > storageAccountUtilizationBeforeGrowing)
	if aboveAvg && countAccounts < maxStorageAccounts {
		klog.V(2).Infof("azureDisk - shared storageAccounts utilization(%v) >  grow-at-avg-utilization (%v). New storage account will be created", avgUtilization, storageAccountUtilizationBeforeGrowing)
		SAName = generateStorageAccountName(sharedDiskAccountNamePrefix)
		err := c.createStorageAccount(SAName, storageAccountType, c.common.location, true)
		if err != nil {
			return "", err
		}
		return SAName, nil
	}
	if aboveAvg && countAccounts == maxStorageAccounts {
		klog.Infof("azureDisk - shared storageAccounts utilization(%v) > grow-at-avg-utilization (%v). But k8s maxed on SAs for PVC(%v). k8s will now exceed grow-at-avg-utilization without adding accounts", avgUtilization, storageAccountUtilizationBeforeGrowing, maxStorageAccounts)
	}
	return SAName, nil
}
func (c *BlobDiskController) getStorageAccountState(storageAccountName string) (bool, storage.ProvisioningState, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := getContextWithCancel()
	defer cancel()
	account, err := c.common.cloud.StorageAccountClient.GetProperties(ctx, c.common.resourceGroup, storageAccountName)
	if err != nil {
		return false, "", err
	}
	return true, account.AccountProperties.ProvisioningState, nil
}
func (c *BlobDiskController) addAccountState(key string, state *storageAccountState) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	accountsLock.Lock()
	defer accountsLock.Unlock()
	if _, ok := c.accounts[key]; !ok {
		c.accounts[key] = state
	}
}
func createVHDHeader(size uint64) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	h := vhd.CreateFixedHeader(size, &vhd.VHDOptions{})
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, h)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func diskNameandSANameFromURI(diskURI string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	uri, err := url.Parse(diskURI)
	if err != nil {
		return "", "", err
	}
	hostName := uri.Host
	storageAccountName := strings.Split(hostName, ".")[0]
	segments := strings.Split(uri.Path, "/")
	diskNameVhd := segments[len(segments)-1]
	return storageAccountName, diskNameVhd, nil
}
