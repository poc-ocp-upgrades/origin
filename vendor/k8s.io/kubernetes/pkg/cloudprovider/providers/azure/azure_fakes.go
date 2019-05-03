package azure

import (
 "context"
 "fmt"
 "math/rand"
 "net/http"
 "strings"
 "sync"
 "time"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 cloudprovider "k8s.io/cloud-provider"
 "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
 "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
 "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-07-01/storage"
 "github.com/Azure/go-autorest/autorest"
 "github.com/Azure/go-autorest/autorest/to"
)

type fakeAzureLBClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.LoadBalancer
}

func newFakeAzureLBClient() *fakeAzureLBClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fLBC := &fakeAzureLBClient{}
 fLBC.FakeStore = make(map[string]map[string]network.LoadBalancer)
 fLBC.mutex = &sync.Mutex{}
 return fLBC
}
func (fLBC *fakeAzureLBClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fLBC.mutex.Lock()
 defer fLBC.mutex.Unlock()
 if _, ok := fLBC.FakeStore[resourceGroupName]; !ok {
  fLBC.FakeStore[resourceGroupName] = make(map[string]network.LoadBalancer)
 }
 if parameters.FrontendIPConfigurations != nil {
  for idx, config := range *parameters.FrontendIPConfigurations {
   if config.PrivateIPAllocationMethod == network.Dynamic {
    (*parameters.FrontendIPConfigurations)[idx].PrivateIPAddress = getRandomIPPtr()
   }
  }
 }
 fLBC.FakeStore[resourceGroupName][loadBalancerName] = parameters
 return nil, nil
}
func (fLBC *fakeAzureLBClient) Delete(ctx context.Context, resourceGroupName string, loadBalancerName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fLBC.mutex.Lock()
 defer fLBC.mutex.Unlock()
 if rgLBs, ok := fLBC.FakeStore[resourceGroupName]; ok {
  if _, ok := rgLBs[loadBalancerName]; ok {
   delete(rgLBs, loadBalancerName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusNotFound}, nil
}
func (fLBC *fakeAzureLBClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fLBC.mutex.Lock()
 defer fLBC.mutex.Unlock()
 if _, ok := fLBC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fLBC.FakeStore[resourceGroupName][loadBalancerName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such LB"}
}
func (fLBC *fakeAzureLBClient) List(ctx context.Context, resourceGroupName string) (result []network.LoadBalancer, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fLBC.mutex.Lock()
 defer fLBC.mutex.Unlock()
 var value []network.LoadBalancer
 if _, ok := fLBC.FakeStore[resourceGroupName]; ok {
  for _, v := range fLBC.FakeStore[resourceGroupName] {
   value = append(value, v)
  }
 }
 return value, nil
}

type fakeAzurePIPClient struct {
 mutex          *sync.Mutex
 FakeStore      map[string]map[string]network.PublicIPAddress
 SubscriptionID string
}

const publicIPAddressIDTemplate = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/publicIPAddresses/%s"

func getpublicIPAddressID(subscriptionID string, resourceGroupName, pipName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf(publicIPAddressIDTemplate, subscriptionID, resourceGroupName, pipName)
}
func newFakeAzurePIPClient(subscriptionID string) *fakeAzurePIPClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC := &fakeAzurePIPClient{}
 fAPC.FakeStore = make(map[string]map[string]network.PublicIPAddress)
 fAPC.SubscriptionID = subscriptionID
 fAPC.mutex = &sync.Mutex{}
 return fAPC
}
func (fAPC *fakeAzurePIPClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC.mutex.Lock()
 defer fAPC.mutex.Unlock()
 if _, ok := fAPC.FakeStore[resourceGroupName]; !ok {
  fAPC.FakeStore[resourceGroupName] = make(map[string]network.PublicIPAddress)
 }
 pipID := getpublicIPAddressID(fAPC.SubscriptionID, resourceGroupName, publicIPAddressName)
 parameters.ID = &pipID
 if parameters.PublicIPAddressPropertiesFormat != nil && parameters.PublicIPAddressPropertiesFormat.PublicIPAllocationMethod == network.Static {
  parameters.IPAddress = getRandomIPPtr()
 }
 fAPC.FakeStore[resourceGroupName][publicIPAddressName] = parameters
 return nil, nil
}
func (fAPC *fakeAzurePIPClient) Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC.mutex.Lock()
 defer fAPC.mutex.Unlock()
 if rgPIPs, ok := fAPC.FakeStore[resourceGroupName]; ok {
  if _, ok := rgPIPs[publicIPAddressName]; ok {
   delete(rgPIPs, publicIPAddressName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusNotFound}, nil
}
func (fAPC *fakeAzurePIPClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC.mutex.Lock()
 defer fAPC.mutex.Unlock()
 if _, ok := fAPC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fAPC.FakeStore[resourceGroupName][publicIPAddressName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such PIP"}
}
func (fAPC *fakeAzurePIPClient) List(ctx context.Context, resourceGroupName string) (result []network.PublicIPAddress, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC.mutex.Lock()
 defer fAPC.mutex.Unlock()
 var value []network.PublicIPAddress
 if _, ok := fAPC.FakeStore[resourceGroupName]; ok {
  for _, v := range fAPC.FakeStore[resourceGroupName] {
   value = append(value, v)
  }
 }
 return value, nil
}
func (fAPC *fakeAzurePIPClient) setFakeStore(store map[string]map[string]network.PublicIPAddress) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fAPC.mutex.Lock()
 defer fAPC.mutex.Unlock()
 fAPC.FakeStore = store
}

type fakeAzureInterfacesClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.Interface
}

func newFakeAzureInterfacesClient() *fakeAzureInterfacesClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fIC := &fakeAzureInterfacesClient{}
 fIC.FakeStore = make(map[string]map[string]network.Interface)
 fIC.mutex = &sync.Mutex{}
 return fIC
}
func (fIC *fakeAzureInterfacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters network.Interface) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fIC.mutex.Lock()
 defer fIC.mutex.Unlock()
 if _, ok := fIC.FakeStore[resourceGroupName]; !ok {
  fIC.FakeStore[resourceGroupName] = make(map[string]network.Interface)
 }
 fIC.FakeStore[resourceGroupName][networkInterfaceName] = parameters
 return nil, nil
}
func (fIC *fakeAzureInterfacesClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fIC.mutex.Lock()
 defer fIC.mutex.Unlock()
 if _, ok := fIC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fIC.FakeStore[resourceGroupName][networkInterfaceName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such Interface"}
}
func (fIC *fakeAzureInterfacesClient) GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fIC.mutex.Lock()
 defer fIC.mutex.Unlock()
 if _, ok := fIC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fIC.FakeStore[resourceGroupName][networkInterfaceName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such Interface"}
}
func (fIC *fakeAzureInterfacesClient) setFakeStore(store map[string]map[string]network.Interface) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fIC.mutex.Lock()
 defer fIC.mutex.Unlock()
 fIC.FakeStore = store
}

type fakeAzureVirtualMachinesClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]compute.VirtualMachine
}

func newFakeAzureVirtualMachinesClient() *fakeAzureVirtualMachinesClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC := &fakeAzureVirtualMachinesClient{}
 fVMC.FakeStore = make(map[string]map[string]compute.VirtualMachine)
 fVMC.mutex = &sync.Mutex{}
 return fVMC
}
func (fVMC *fakeAzureVirtualMachinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachine) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 if _, ok := fVMC.FakeStore[resourceGroupName]; !ok {
  fVMC.FakeStore[resourceGroupName] = make(map[string]compute.VirtualMachine)
 }
 fVMC.FakeStore[resourceGroupName][VMName] = parameters
 return nil, nil
}
func (fVMC *fakeAzureVirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 if _, ok := fVMC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fVMC.FakeStore[resourceGroupName][VMName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such VM"}
}
func (fVMC *fakeAzureVirtualMachinesClient) List(ctx context.Context, resourceGroupName string) (result []compute.VirtualMachine, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 result = []compute.VirtualMachine{}
 if _, ok := fVMC.FakeStore[resourceGroupName]; ok {
  for _, v := range fVMC.FakeStore[resourceGroupName] {
   result = append(result, v)
  }
 }
 return result, nil
}
func (fVMC *fakeAzureVirtualMachinesClient) setFakeStore(store map[string]map[string]compute.VirtualMachine) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 fVMC.FakeStore = store
}

type fakeAzureSubnetsClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.Subnet
}

func newFakeAzureSubnetsClient() *fakeAzureSubnetsClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fASC := &fakeAzureSubnetsClient{}
 fASC.FakeStore = make(map[string]map[string]network.Subnet)
 fASC.mutex = &sync.Mutex{}
 return fASC
}
func (fASC *fakeAzureSubnetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fASC.mutex.Lock()
 defer fASC.mutex.Unlock()
 rgVnet := strings.Join([]string{resourceGroupName, virtualNetworkName}, "AND")
 if _, ok := fASC.FakeStore[rgVnet]; !ok {
  fASC.FakeStore[rgVnet] = make(map[string]network.Subnet)
 }
 fASC.FakeStore[rgVnet][subnetName] = subnetParameters
 return nil, nil
}
func (fASC *fakeAzureSubnetsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fASC.mutex.Lock()
 defer fASC.mutex.Unlock()
 rgVnet := strings.Join([]string{resourceGroupName, virtualNetworkName}, "AND")
 if rgSubnets, ok := fASC.FakeStore[rgVnet]; ok {
  if _, ok := rgSubnets[subnetName]; ok {
   delete(rgSubnets, subnetName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusNotFound}, nil
}
func (fASC *fakeAzureSubnetsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fASC.mutex.Lock()
 defer fASC.mutex.Unlock()
 rgVnet := strings.Join([]string{resourceGroupName, virtualNetworkName}, "AND")
 if _, ok := fASC.FakeStore[rgVnet]; ok {
  if entity, ok := fASC.FakeStore[rgVnet][subnetName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such Subnet"}
}
func (fASC *fakeAzureSubnetsClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result []network.Subnet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fASC.mutex.Lock()
 defer fASC.mutex.Unlock()
 rgVnet := strings.Join([]string{resourceGroupName, virtualNetworkName}, "AND")
 var value []network.Subnet
 if _, ok := fASC.FakeStore[rgVnet]; ok {
  for _, v := range fASC.FakeStore[rgVnet] {
   value = append(value, v)
  }
 }
 return value, nil
}

type fakeAzureNSGClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.SecurityGroup
}

func newFakeAzureNSGClient() *fakeAzureNSGClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fNSG := &fakeAzureNSGClient{}
 fNSG.FakeStore = make(map[string]map[string]network.SecurityGroup)
 fNSG.mutex = &sync.Mutex{}
 return fNSG
}
func (fNSG *fakeAzureNSGClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fNSG.mutex.Lock()
 defer fNSG.mutex.Unlock()
 if _, ok := fNSG.FakeStore[resourceGroupName]; !ok {
  fNSG.FakeStore[resourceGroupName] = make(map[string]network.SecurityGroup)
 }
 fNSG.FakeStore[resourceGroupName][networkSecurityGroupName] = parameters
 return nil, nil
}
func (fNSG *fakeAzureNSGClient) Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fNSG.mutex.Lock()
 defer fNSG.mutex.Unlock()
 if rgSGs, ok := fNSG.FakeStore[resourceGroupName]; ok {
  if _, ok := rgSGs[networkSecurityGroupName]; ok {
   delete(rgSGs, networkSecurityGroupName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusNotFound}, nil
}
func (fNSG *fakeAzureNSGClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fNSG.mutex.Lock()
 defer fNSG.mutex.Unlock()
 if _, ok := fNSG.FakeStore[resourceGroupName]; ok {
  if entity, ok := fNSG.FakeStore[resourceGroupName][networkSecurityGroupName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such NSG"}
}
func (fNSG *fakeAzureNSGClient) List(ctx context.Context, resourceGroupName string) (result []network.SecurityGroup, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fNSG.mutex.Lock()
 defer fNSG.mutex.Unlock()
 var value []network.SecurityGroup
 if _, ok := fNSG.FakeStore[resourceGroupName]; ok {
  for _, v := range fNSG.FakeStore[resourceGroupName] {
   value = append(value, v)
  }
 }
 return value, nil
}
func getRandomIPPtr() *string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 rand.Seed(time.Now().UnixNano())
 return to.StringPtr(fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)))
}

type fakeVirtualMachineScaleSetVMsClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]compute.VirtualMachineScaleSetVM
}

func newFakeVirtualMachineScaleSetVMsClient() *fakeVirtualMachineScaleSetVMsClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC := &fakeVirtualMachineScaleSetVMsClient{}
 fVMC.FakeStore = make(map[string]map[string]compute.VirtualMachineScaleSetVM)
 fVMC.mutex = &sync.Mutex{}
 return fVMC
}
func (fVMC *fakeVirtualMachineScaleSetVMsClient) setFakeStore(store map[string]map[string]compute.VirtualMachineScaleSetVM) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 fVMC.FakeStore = store
}
func (fVMC *fakeVirtualMachineScaleSetVMsClient) List(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result []compute.VirtualMachineScaleSetVM, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 result = []compute.VirtualMachineScaleSetVM{}
 if _, ok := fVMC.FakeStore[resourceGroupName]; ok {
  for _, v := range fVMC.FakeStore[resourceGroupName] {
   result = append(result, v)
  }
 }
 return result, nil
}
func (fVMC *fakeVirtualMachineScaleSetVMsClient) Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVM, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 vmKey := fmt.Sprintf("%s_%s", VMScaleSetName, instanceID)
 if scaleSetMap, ok := fVMC.FakeStore[resourceGroupName]; ok {
  if entity, ok := scaleSetMap[vmKey]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "No such VirtualMachineScaleSetVM"}
}
func (fVMC *fakeVirtualMachineScaleSetVMsClient) GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 _, err = fVMC.Get(ctx, resourceGroupName, VMScaleSetName, instanceID)
 if err != nil {
  return result, err
 }
 return result, nil
}
func (fVMC *fakeVirtualMachineScaleSetVMsClient) Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.VirtualMachineScaleSetVM) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMC.mutex.Lock()
 defer fVMC.mutex.Unlock()
 vmKey := fmt.Sprintf("%s_%s", VMScaleSetName, instanceID)
 if scaleSetMap, ok := fVMC.FakeStore[resourceGroupName]; ok {
  if _, ok := scaleSetMap[vmKey]; ok {
   scaleSetMap[vmKey] = parameters
  }
 }
 return nil, nil
}

type fakeVirtualMachineScaleSetsClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]compute.VirtualMachineScaleSet
}

func newFakeVirtualMachineScaleSetsClient() *fakeVirtualMachineScaleSetsClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMSSC := &fakeVirtualMachineScaleSetsClient{}
 fVMSSC.FakeStore = make(map[string]map[string]compute.VirtualMachineScaleSet)
 fVMSSC.mutex = &sync.Mutex{}
 return fVMSSC
}
func (fVMSSC *fakeVirtualMachineScaleSetsClient) setFakeStore(store map[string]map[string]compute.VirtualMachineScaleSet) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMSSC.mutex.Lock()
 defer fVMSSC.mutex.Unlock()
 fVMSSC.FakeStore = store
}
func (fVMSSC *fakeVirtualMachineScaleSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSet) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMSSC.mutex.Lock()
 defer fVMSSC.mutex.Unlock()
 if _, ok := fVMSSC.FakeStore[resourceGroupName]; !ok {
  fVMSSC.FakeStore[resourceGroupName] = make(map[string]compute.VirtualMachineScaleSet)
 }
 fVMSSC.FakeStore[resourceGroupName][VMScaleSetName] = parameters
 return nil, nil
}
func (fVMSSC *fakeVirtualMachineScaleSetsClient) Get(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMSSC.mutex.Lock()
 defer fVMSSC.mutex.Unlock()
 if scaleSetMap, ok := fVMSSC.FakeStore[resourceGroupName]; ok {
  if entity, ok := scaleSetMap[VMScaleSetName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "No such ScaleSet"}
}
func (fVMSSC *fakeVirtualMachineScaleSetsClient) List(ctx context.Context, resourceGroupName string) (result []compute.VirtualMachineScaleSet, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fVMSSC.mutex.Lock()
 defer fVMSSC.mutex.Unlock()
 result = []compute.VirtualMachineScaleSet{}
 if _, ok := fVMSSC.FakeStore[resourceGroupName]; ok {
  for _, v := range fVMSSC.FakeStore[resourceGroupName] {
   result = append(result, v)
  }
 }
 return result, nil
}
func (fVMSSC *fakeVirtualMachineScaleSetsClient) UpdateInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, nil
}

type fakeRoutesClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.Route
}

func newFakeRoutesClient() *fakeRoutesClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRC := &fakeRoutesClient{}
 fRC.FakeStore = make(map[string]map[string]network.Route)
 fRC.mutex = &sync.Mutex{}
 return fRC
}
func (fRC *fakeRoutesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRC.mutex.Lock()
 defer fRC.mutex.Unlock()
 if _, ok := fRC.FakeStore[routeTableName]; !ok {
  fRC.FakeStore[routeTableName] = make(map[string]network.Route)
 }
 fRC.FakeStore[routeTableName][routeName] = routeParameters
 return nil, nil
}
func (fRC *fakeRoutesClient) Delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRC.mutex.Lock()
 defer fRC.mutex.Unlock()
 if routes, ok := fRC.FakeStore[routeTableName]; ok {
  if _, ok := routes[routeName]; ok {
   delete(routes, routeName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusNotFound}, nil
}

type fakeRouteTablesClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]network.RouteTable
 Calls     []string
}

func newFakeRouteTablesClient() *fakeRouteTablesClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRTC := &fakeRouteTablesClient{}
 fRTC.FakeStore = make(map[string]map[string]network.RouteTable)
 fRTC.mutex = &sync.Mutex{}
 return fRTC
}
func (fRTC *fakeRouteTablesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters network.RouteTable) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRTC.mutex.Lock()
 defer fRTC.mutex.Unlock()
 fRTC.Calls = append(fRTC.Calls, "CreateOrUpdate")
 if _, ok := fRTC.FakeStore[resourceGroupName]; !ok {
  fRTC.FakeStore[resourceGroupName] = make(map[string]network.RouteTable)
 }
 fRTC.FakeStore[resourceGroupName][routeTableName] = parameters
 return nil, nil
}
func (fRTC *fakeRouteTablesClient) Get(ctx context.Context, resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fRTC.mutex.Lock()
 defer fRTC.mutex.Unlock()
 fRTC.Calls = append(fRTC.Calls, "Get")
 if _, ok := fRTC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fRTC.FakeStore[resourceGroupName][routeTableName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such RouteTable"}
}

type fakeFileClient struct{}

func (fFC *fakeFileClient) createFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fFC *fakeFileClient) deleteFileShare(accountName, accountKey, name string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}
func (fFC *fakeFileClient) resizeFileShare(accountName, accountKey, name string, sizeGiB int) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil
}

type fakeStorageAccountClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]storage.Account
 Keys      storage.AccountListKeysResult
 Accounts  storage.AccountListResult
 Err       error
}

func newFakeStorageAccountClient() *fakeStorageAccountClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fSAC := &fakeStorageAccountClient{}
 fSAC.FakeStore = make(map[string]map[string]storage.Account)
 fSAC.mutex = &sync.Mutex{}
 return fSAC
}
func (fSAC *fakeStorageAccountClient) Create(ctx context.Context, resourceGroupName string, accountName string, parameters storage.AccountCreateParameters) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fSAC.mutex.Lock()
 defer fSAC.mutex.Unlock()
 if _, ok := fSAC.FakeStore[resourceGroupName]; !ok {
  fSAC.FakeStore[resourceGroupName] = make(map[string]storage.Account)
 }
 fSAC.FakeStore[resourceGroupName][accountName] = storage.Account{Name: &accountName, Sku: parameters.Sku, Kind: parameters.Kind, Location: parameters.Location, Identity: parameters.Identity, Tags: parameters.Tags, AccountProperties: &storage.AccountProperties{}}
 return nil, nil
}
func (fSAC *fakeStorageAccountClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fSAC.mutex.Lock()
 defer fSAC.mutex.Unlock()
 if rgAccounts, ok := fSAC.FakeStore[resourceGroupName]; ok {
  if _, ok := rgAccounts[accountName]; ok {
   delete(rgAccounts, accountName)
   result.Response = &http.Response{StatusCode: http.StatusAccepted}
   return result, nil
  }
 }
 result.Response = &http.Response{StatusCode: http.StatusNotFound}
 err = autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such StorageAccount"}
 return result, err
}
func (fSAC *fakeStorageAccountClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fSAC.Keys, fSAC.Err
}
func (fSAC *fakeStorageAccountClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storage.AccountListResult, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fSAC.Accounts, fSAC.Err
}
func (fSAC *fakeStorageAccountClient) GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fSAC.mutex.Lock()
 defer fSAC.mutex.Unlock()
 if _, ok := fSAC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fSAC.FakeStore[resourceGroupName][accountName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such StorageAccount"}
}

type fakeDisksClient struct {
 mutex     *sync.Mutex
 FakeStore map[string]map[string]compute.Disk
}

func newFakeDisksClient() *fakeDisksClient {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fDC := &fakeDisksClient{}
 fDC.FakeStore = make(map[string]map[string]compute.Disk)
 fDC.mutex = &sync.Mutex{}
 return fDC
}
func (fDC *fakeDisksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, diskParameter compute.Disk) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fDC.mutex.Lock()
 defer fDC.mutex.Unlock()
 if _, ok := fDC.FakeStore[resourceGroupName]; !ok {
  fDC.FakeStore[resourceGroupName] = make(map[string]compute.Disk)
 }
 fDC.FakeStore[resourceGroupName][diskName] = diskParameter
 return nil, nil
}
func (fDC *fakeDisksClient) Delete(ctx context.Context, resourceGroupName string, diskName string) (resp *http.Response, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fDC.mutex.Lock()
 defer fDC.mutex.Unlock()
 if rgDisks, ok := fDC.FakeStore[resourceGroupName]; ok {
  if _, ok := rgDisks[diskName]; ok {
   delete(rgDisks, diskName)
   return nil, nil
  }
 }
 return &http.Response{StatusCode: http.StatusAccepted}, nil
}
func (fDC *fakeDisksClient) Get(ctx context.Context, resourceGroupName string, diskName string) (result compute.Disk, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fDC.mutex.Lock()
 defer fDC.mutex.Unlock()
 if _, ok := fDC.FakeStore[resourceGroupName]; ok {
  if entity, ok := fDC.FakeStore[resourceGroupName][diskName]; ok {
   return entity, nil
  }
 }
 return result, autorest.DetailedError{StatusCode: http.StatusNotFound, Message: "Not such Disk"}
}

type fakeVMSet struct {
 NodeToIP map[string]string
 Err      error
}

func (f *fakeVMSet) GetInstanceIDByNodeName(name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "", fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetInstanceTypeByNodeName(name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "", fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetIPByNodeName(name string) (string, string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ip, found := f.NodeToIP[name]
 if !found {
  return "", "", fmt.Errorf("not found")
 }
 return ip, "", nil
}
func (f *fakeVMSet) GetPrimaryInterface(nodeName string) (network.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return network.Interface{}, fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetNodeNameByProviderID(providerID string) (types.NodeName, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return types.NodeName(""), fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetZoneByNodeName(name string) (cloudprovider.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return cloudprovider.Zone{}, fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetPrimaryVMSetName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return ""
}
func (f *fakeVMSet) GetVMSetNames(service *v1.Service, nodes []*v1.Node) (availabilitySetNames *[]string, err error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) EnsureHostsInPool(service *v1.Service, nodes []*v1.Node, backendPoolID string, vmSetName string, isInternal bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) EnsureBackendPoolDeleted(service *v1.Service, poolID, vmSetName string, backendAddressPools *[]network.BackendAddressPool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) AttachDisk(isManagedDisk bool, diskName, diskURI string, nodeName types.NodeName, lun int32, cachingMode compute.CachingTypes) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) DetachDiskByName(diskName, diskURI string, nodeName types.NodeName) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetDataDisks(nodeName types.NodeName) ([]compute.DataDisk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return nil, fmt.Errorf("unimplemented")
}
func (f *fakeVMSet) GetPowerStatusByNodeName(name string) (string, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return "", fmt.Errorf("unimplemented")
}
