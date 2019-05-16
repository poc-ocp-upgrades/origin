package ovirt

import (
	"context"
	"encoding/xml"
	"fmt"
	goformat "fmt"
	"gopkg.in/gcfg.v1"
	"io"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"net"
	"net/http"
	"net/url"
	goos "os"
	"path"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

const ProviderName = "ovirt"

type OVirtInstance struct {
	UUID      string
	Name      string
	IPAddress string
}
type OVirtInstanceMap map[string]OVirtInstance
type OVirtCloud struct {
	VmsRequest   *url.URL
	HostsRequest *url.URL
}
type OVirtApiConfig struct {
	Connection struct {
		ApiEntry string `gcfg:"uri"`
		Username string `gcfg:"username"`
		Password string `gcfg:"password"`
	}
	Filters struct {
		VmsQuery string `gcfg:"vms"`
	}
}
type XmlVmAddress struct {
	Address string `xml:"address,attr"`
}
type XmlVmInfo struct {
	UUID      string         `xml:"id,attr"`
	Name      string         `xml:"name"`
	Hostname  string         `xml:"guest_info>fqdn"`
	Addresses []XmlVmAddress `xml:"guest_info>ips>ip"`
	State     string         `xml:"status>state"`
}
type XmlVmsList struct {
	XMLName xml.Name    `xml:"vms"`
	Vm      []XmlVmInfo `xml:"vm"`
}

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newOVirtCloud(config)
	})
}
func newOVirtCloud(config io.Reader) (*OVirtCloud, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if config == nil {
		return nil, fmt.Errorf("missing configuration file for ovirt cloud provider")
	}
	oVirtConfig := OVirtApiConfig{}
	oVirtConfig.Connection.Username = "admin@internal"
	if err := gcfg.ReadInto(&oVirtConfig, config); err != nil {
		return nil, err
	}
	if oVirtConfig.Connection.ApiEntry == "" {
		return nil, fmt.Errorf("missing ovirt uri in cloud provider configuration")
	}
	request, err := url.Parse(oVirtConfig.Connection.ApiEntry)
	if err != nil {
		return nil, err
	}
	request.Path = path.Join(request.Path, "vms")
	request.User = url.UserPassword(oVirtConfig.Connection.Username, oVirtConfig.Connection.Password)
	request.RawQuery = url.Values{"search": {oVirtConfig.Filters.VmsQuery}}.Encode()
	return &OVirtCloud{VmsRequest: request}, nil
}
func (v *OVirtCloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (v *OVirtCloud) Clusters() (cloudprovider.Clusters, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (v *OVirtCloud) ProviderName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ProviderName
}
func (v *OVirtCloud) HasClusterID() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (v *OVirtCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (v *OVirtCloud) Instances() (cloudprovider.Instances, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v, true
}
func (v *OVirtCloud) Zones() (cloudprovider.Zones, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (v *OVirtCloud) Routes() (cloudprovider.Routes, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, false
}
func (v *OVirtCloud) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name := mapNodeNameToInstanceName(nodeName)
	instance, err := v.fetchInstance(name)
	if err != nil {
		return nil, err
	}
	var address net.IP
	if instance.IPAddress != "" {
		address = net.ParseIP(instance.IPAddress)
		if address == nil {
			return nil, fmt.Errorf("couldn't parse address: %s", instance.IPAddress)
		}
	} else {
		resolved, err := net.LookupIP(name)
		if err != nil || len(resolved) < 1 {
			return nil, fmt.Errorf("couldn't lookup address: %s", name)
		}
		address = resolved[0]
	}
	return []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: address.String()}, {Type: v1.NodeExternalIP, Address: address.String()}}, nil
}
func (v *OVirtCloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []v1.NodeAddress{}, cloudprovider.NotImplemented
}
func mapNodeNameToInstanceName(nodeName types.NodeName) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return string(nodeName)
}
func (v *OVirtCloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false, cloudprovider.NotImplemented
}
func (v *OVirtCloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false, cloudprovider.NotImplemented
}
func (v *OVirtCloud) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name := mapNodeNameToInstanceName(nodeName)
	instance, err := v.fetchInstance(name)
	if err != nil {
		return "", err
	}
	return "/" + instance.UUID, err
}
func (v *OVirtCloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "", cloudprovider.NotImplemented
}
func (v *OVirtCloud) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "", nil
}
func getInstancesFromXml(body io.Reader) (OVirtInstanceMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if body == nil {
		return nil, fmt.Errorf("ovirt rest-api response body is missing")
	}
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	vmlist := XmlVmsList{}
	if err := xml.Unmarshal(content, &vmlist); err != nil {
		return nil, err
	}
	instances := make(OVirtInstanceMap)
	for _, vm := range vmlist.Vm {
		if vm.Hostname != "" && strings.ToLower(vm.State) == "up" {
			address := ""
			if len(vm.Addresses) > 0 {
				address = vm.Addresses[0].Address
			}
			instances[vm.Hostname] = OVirtInstance{UUID: vm.UUID, Name: vm.Name, IPAddress: address}
		}
	}
	return instances, nil
}
func (v *OVirtCloud) fetchAllInstances() (OVirtInstanceMap, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	response, err := http.Get(v.VmsRequest.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return getInstancesFromXml(response.Body)
}
func (v *OVirtCloud) fetchInstance(name string) (*OVirtInstance, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allInstances, err := v.fetchAllInstances()
	if err != nil {
		return nil, err
	}
	instance, found := allInstances[name]
	if !found {
		return nil, fmt.Errorf("cannot find instance: %s", name)
	}
	return &instance, nil
}
func (m *OVirtInstanceMap) ListSortedNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var names []string
	for k := range *m {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
func (v *OVirtCloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return types.NodeName(hostname), nil
}
func (v *OVirtCloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cloudprovider.NotImplemented
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
