package azure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	metadataCacheTTL = time.Minute
	metadataCacheKey = "InstanceMetadata"
	metadataURL      = "http://169.254.169.254/metadata/instance"
)

type NetworkMetadata struct {
	Interface []NetworkInterface `json:"interface"`
}
type NetworkInterface struct {
	IPV4 NetworkData `json:"ipv4"`
	IPV6 NetworkData `json:"ipv6"`
	MAC  string      `json:"macAddress"`
}
type NetworkData struct {
	IPAddress []IPAddress `json:"ipAddress"`
	Subnet    []Subnet    `json:"subnet"`
}
type IPAddress struct {
	PrivateIP string `json:"privateIPAddress"`
	PublicIP  string `json:"publicIPAddress"`
}
type Subnet struct {
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
}
type ComputeMetadata struct {
	SKU            string `json:"sku,omitempty"`
	Name           string `json:"name,omitempty"`
	Zone           string `json:"zone,omitempty"`
	VMSize         string `json:"vmSize,omitempty"`
	OSType         string `json:"osType,omitempty"`
	Location       string `json:"location,omitempty"`
	FaultDomain    string `json:"platformFaultDomain,omitempty"`
	UpdateDomain   string `json:"platformUpdateDomain,omitempty"`
	ResourceGroup  string `json:"resourceGroupName,omitempty"`
	VMScaleSetName string `json:"vmScaleSetName,omitempty"`
}
type InstanceMetadata struct {
	Compute *ComputeMetadata `json:"compute,omitempty"`
	Network *NetworkMetadata `json:"network,omitempty"`
}
type InstanceMetadataService struct {
	metadataURL string
	imsCache    *timedCache
}

func NewInstanceMetadataService(metadataURL string) (*InstanceMetadataService, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ims := &InstanceMetadataService{metadataURL: metadataURL}
	imsCache, err := newTimedcache(metadataCacheTTL, ims.getInstanceMetadata)
	if err != nil {
		return nil, err
	}
	ims.imsCache = imsCache
	return ims, nil
}
func (ims *InstanceMetadataService) getInstanceMetadata(key string) (interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", ims.metadataURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Metadata", "True")
	req.Header.Add("User-Agent", "golang/kubernetes-cloud-provider")
	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2017-12-01")
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failure of getting instance metadata with response %q", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := InstanceMetadata{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
func (ims *InstanceMetadataService) GetMetadata() (*InstanceMetadata, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cache, err := ims.imsCache.Get(metadataCacheKey)
	if err != nil {
		return nil, err
	}
	if cache == nil {
		return nil, fmt.Errorf("failure of getting instance metadata")
	}
	if metadata, ok := cache.(*InstanceMetadata); ok {
		return metadata, nil
	}
	return nil, fmt.Errorf("failure of getting instance metadata")
}
