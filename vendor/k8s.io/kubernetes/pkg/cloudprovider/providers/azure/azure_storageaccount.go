package azure

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/klog"
	"strings"
)

type accountWithLocation struct{ Name, StorageType, Location string }

func (az *Cloud) getStorageAccounts(matchingAccountType, matchingAccountKind, resourceGroup, matchingLocation string) ([]accountWithLocation, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := getContextWithCancel()
	defer cancel()
	result, err := az.StorageAccountClient.ListByResourceGroup(ctx, resourceGroup)
	if err != nil {
		return nil, err
	}
	if result.Value == nil {
		return nil, fmt.Errorf("unexpected error when listing storage accounts from resource group %s", resourceGroup)
	}
	accounts := []accountWithLocation{}
	for _, acct := range *result.Value {
		if acct.Name != nil && acct.Location != nil && acct.Sku != nil {
			storageType := string((*acct.Sku).Name)
			if matchingAccountType != "" && !strings.EqualFold(matchingAccountType, storageType) {
				continue
			}
			if matchingAccountKind != "" && !strings.EqualFold(matchingAccountKind, string(acct.Kind)) {
				continue
			}
			location := *acct.Location
			if matchingLocation != "" && !strings.EqualFold(matchingLocation, location) {
				continue
			}
			accounts = append(accounts, accountWithLocation{Name: *acct.Name, StorageType: storageType, Location: location})
		}
	}
	return accounts, nil
}
func (az *Cloud) getStorageAccesskey(account, resourceGroup string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := getContextWithCancel()
	defer cancel()
	result, err := az.StorageAccountClient.ListKeys(ctx, resourceGroup, account)
	if err != nil {
		return "", err
	}
	if result.Keys == nil {
		return "", fmt.Errorf("empty keys")
	}
	for _, k := range *result.Keys {
		if k.Value != nil && *k.Value != "" {
			v := *k.Value
			if ind := strings.LastIndex(v, " "); ind >= 0 {
				v = v[(ind + 1):]
			}
			return v, nil
		}
	}
	return "", fmt.Errorf("no valid keys")
}
func (az *Cloud) ensureStorageAccount(accountName, accountType, accountKind, resourceGroup, location, genAccountNamePrefix string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(accountName) == 0 {
		accounts, err := az.getStorageAccounts(accountType, accountKind, resourceGroup, location)
		if err != nil {
			return "", "", fmt.Errorf("could not list storage accounts for account type %s: %v", accountType, err)
		}
		if len(accounts) > 0 {
			accountName = accounts[0].Name
			klog.V(4).Infof("found a matching account %s type %s location %s", accounts[0].Name, accounts[0].StorageType, accounts[0].Location)
		}
		if len(accountName) == 0 {
			accountName = generateStorageAccountName(genAccountNamePrefix)
			if location == "" {
				location = az.Location
			}
			if accountType == "" {
				accountType = defaultStorageAccountType
			}
			kind := defaultStorageAccountKind
			if accountKind != "" {
				kind = storage.Kind(accountKind)
			}
			klog.V(2).Infof("azure - no matching account found, begin to create a new account %s in resource group %s, location: %s, accountType: %s, accountKind: %s", accountName, resourceGroup, location, accountType, kind)
			cp := storage.AccountCreateParameters{Sku: &storage.Sku{Name: storage.SkuName(accountType)}, Kind: kind, AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{EnableHTTPSTrafficOnly: to.BoolPtr(true)}, Tags: map[string]*string{"created-by": to.StringPtr("azure")}, Location: &location}
			ctx, cancel := getContextWithCancel()
			defer cancel()
			_, err := az.StorageAccountClient.Create(ctx, resourceGroup, accountName, cp)
			if err != nil {
				return "", "", fmt.Errorf(fmt.Sprintf("Failed to create storage account %s, error: %s", accountName, err))
			}
		}
	}
	accountKey, err := az.getStorageAccesskey(accountName, resourceGroup)
	if err != nil {
		return "", "", fmt.Errorf("could not get storage key for storage account %s: %v", accountName, err)
	}
	return accountName, accountKey, nil
}
