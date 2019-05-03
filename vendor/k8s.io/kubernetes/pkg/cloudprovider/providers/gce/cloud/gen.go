package cloud

import (
 "context"
 "fmt"
 "net/http"
 "sync"
 "google.golang.org/api/googleapi"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/filter"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
 alpha "google.golang.org/api/compute/v0.alpha"
 beta "google.golang.org/api/compute/v0.beta"
 ga "google.golang.org/api/compute/v1"
)

type Cloud interface {
 Addresses() Addresses
 AlphaAddresses() AlphaAddresses
 BetaAddresses() BetaAddresses
 GlobalAddresses() GlobalAddresses
 BackendServices() BackendServices
 BetaBackendServices() BetaBackendServices
 AlphaBackendServices() AlphaBackendServices
 RegionBackendServices() RegionBackendServices
 AlphaRegionBackendServices() AlphaRegionBackendServices
 Disks() Disks
 RegionDisks() RegionDisks
 Firewalls() Firewalls
 ForwardingRules() ForwardingRules
 AlphaForwardingRules() AlphaForwardingRules
 GlobalForwardingRules() GlobalForwardingRules
 HealthChecks() HealthChecks
 BetaHealthChecks() BetaHealthChecks
 AlphaHealthChecks() AlphaHealthChecks
 HttpHealthChecks() HttpHealthChecks
 HttpsHealthChecks() HttpsHealthChecks
 InstanceGroups() InstanceGroups
 Instances() Instances
 BetaInstances() BetaInstances
 AlphaInstances() AlphaInstances
 AlphaNetworkEndpointGroups() AlphaNetworkEndpointGroups
 BetaNetworkEndpointGroups() BetaNetworkEndpointGroups
 Projects() Projects
 Regions() Regions
 Routes() Routes
 BetaSecurityPolicies() BetaSecurityPolicies
 SslCertificates() SslCertificates
 TargetHttpProxies() TargetHttpProxies
 TargetHttpsProxies() TargetHttpsProxies
 TargetPools() TargetPools
 UrlMaps() UrlMaps
 Zones() Zones
}

func NewGCE(s *Service) *GCE {
 _logClusterCodePath()
 defer _logClusterCodePath()
 g := &GCE{gceAddresses: &GCEAddresses{s}, gceAlphaAddresses: &GCEAlphaAddresses{s}, gceBetaAddresses: &GCEBetaAddresses{s}, gceGlobalAddresses: &GCEGlobalAddresses{s}, gceBackendServices: &GCEBackendServices{s}, gceBetaBackendServices: &GCEBetaBackendServices{s}, gceAlphaBackendServices: &GCEAlphaBackendServices{s}, gceRegionBackendServices: &GCERegionBackendServices{s}, gceAlphaRegionBackendServices: &GCEAlphaRegionBackendServices{s}, gceDisks: &GCEDisks{s}, gceRegionDisks: &GCERegionDisks{s}, gceFirewalls: &GCEFirewalls{s}, gceForwardingRules: &GCEForwardingRules{s}, gceAlphaForwardingRules: &GCEAlphaForwardingRules{s}, gceGlobalForwardingRules: &GCEGlobalForwardingRules{s}, gceHealthChecks: &GCEHealthChecks{s}, gceBetaHealthChecks: &GCEBetaHealthChecks{s}, gceAlphaHealthChecks: &GCEAlphaHealthChecks{s}, gceHttpHealthChecks: &GCEHttpHealthChecks{s}, gceHttpsHealthChecks: &GCEHttpsHealthChecks{s}, gceInstanceGroups: &GCEInstanceGroups{s}, gceInstances: &GCEInstances{s}, gceBetaInstances: &GCEBetaInstances{s}, gceAlphaInstances: &GCEAlphaInstances{s}, gceAlphaNetworkEndpointGroups: &GCEAlphaNetworkEndpointGroups{s}, gceBetaNetworkEndpointGroups: &GCEBetaNetworkEndpointGroups{s}, gceProjects: &GCEProjects{s}, gceRegions: &GCERegions{s}, gceRoutes: &GCERoutes{s}, gceBetaSecurityPolicies: &GCEBetaSecurityPolicies{s}, gceSslCertificates: &GCESslCertificates{s}, gceTargetHttpProxies: &GCETargetHttpProxies{s}, gceTargetHttpsProxies: &GCETargetHttpsProxies{s}, gceTargetPools: &GCETargetPools{s}, gceUrlMaps: &GCEUrlMaps{s}, gceZones: &GCEZones{s}}
 return g
}

var _ Cloud = (*GCE)(nil)

type GCE struct {
 gceAddresses                  *GCEAddresses
 gceAlphaAddresses             *GCEAlphaAddresses
 gceBetaAddresses              *GCEBetaAddresses
 gceGlobalAddresses            *GCEGlobalAddresses
 gceBackendServices            *GCEBackendServices
 gceBetaBackendServices        *GCEBetaBackendServices
 gceAlphaBackendServices       *GCEAlphaBackendServices
 gceRegionBackendServices      *GCERegionBackendServices
 gceAlphaRegionBackendServices *GCEAlphaRegionBackendServices
 gceDisks                      *GCEDisks
 gceRegionDisks                *GCERegionDisks
 gceFirewalls                  *GCEFirewalls
 gceForwardingRules            *GCEForwardingRules
 gceAlphaForwardingRules       *GCEAlphaForwardingRules
 gceGlobalForwardingRules      *GCEGlobalForwardingRules
 gceHealthChecks               *GCEHealthChecks
 gceBetaHealthChecks           *GCEBetaHealthChecks
 gceAlphaHealthChecks          *GCEAlphaHealthChecks
 gceHttpHealthChecks           *GCEHttpHealthChecks
 gceHttpsHealthChecks          *GCEHttpsHealthChecks
 gceInstanceGroups             *GCEInstanceGroups
 gceInstances                  *GCEInstances
 gceBetaInstances              *GCEBetaInstances
 gceAlphaInstances             *GCEAlphaInstances
 gceAlphaNetworkEndpointGroups *GCEAlphaNetworkEndpointGroups
 gceBetaNetworkEndpointGroups  *GCEBetaNetworkEndpointGroups
 gceProjects                   *GCEProjects
 gceRegions                    *GCERegions
 gceRoutes                     *GCERoutes
 gceBetaSecurityPolicies       *GCEBetaSecurityPolicies
 gceSslCertificates            *GCESslCertificates
 gceTargetHttpProxies          *GCETargetHttpProxies
 gceTargetHttpsProxies         *GCETargetHttpsProxies
 gceTargetPools                *GCETargetPools
 gceUrlMaps                    *GCEUrlMaps
 gceZones                      *GCEZones
}

func (gce *GCE) Addresses() Addresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAddresses
}
func (gce *GCE) AlphaAddresses() AlphaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaAddresses
}
func (gce *GCE) BetaAddresses() BetaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaAddresses
}
func (gce *GCE) GlobalAddresses() GlobalAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceGlobalAddresses
}
func (gce *GCE) BackendServices() BackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBackendServices
}
func (gce *GCE) BetaBackendServices() BetaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaBackendServices
}
func (gce *GCE) AlphaBackendServices() AlphaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaBackendServices
}
func (gce *GCE) RegionBackendServices() RegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceRegionBackendServices
}
func (gce *GCE) AlphaRegionBackendServices() AlphaRegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaRegionBackendServices
}
func (gce *GCE) Disks() Disks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceDisks
}
func (gce *GCE) RegionDisks() RegionDisks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceRegionDisks
}
func (gce *GCE) Firewalls() Firewalls {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceFirewalls
}
func (gce *GCE) ForwardingRules() ForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceForwardingRules
}
func (gce *GCE) AlphaForwardingRules() AlphaForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaForwardingRules
}
func (gce *GCE) GlobalForwardingRules() GlobalForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceGlobalForwardingRules
}
func (gce *GCE) HealthChecks() HealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceHealthChecks
}
func (gce *GCE) BetaHealthChecks() BetaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaHealthChecks
}
func (gce *GCE) AlphaHealthChecks() AlphaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaHealthChecks
}
func (gce *GCE) HttpHealthChecks() HttpHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceHttpHealthChecks
}
func (gce *GCE) HttpsHealthChecks() HttpsHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceHttpsHealthChecks
}
func (gce *GCE) InstanceGroups() InstanceGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceInstanceGroups
}
func (gce *GCE) Instances() Instances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceInstances
}
func (gce *GCE) BetaInstances() BetaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaInstances
}
func (gce *GCE) AlphaInstances() AlphaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaInstances
}
func (gce *GCE) AlphaNetworkEndpointGroups() AlphaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceAlphaNetworkEndpointGroups
}
func (gce *GCE) BetaNetworkEndpointGroups() BetaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaNetworkEndpointGroups
}
func (gce *GCE) Projects() Projects {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceProjects
}
func (gce *GCE) Regions() Regions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceRegions
}
func (gce *GCE) Routes() Routes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceRoutes
}
func (gce *GCE) BetaSecurityPolicies() BetaSecurityPolicies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceBetaSecurityPolicies
}
func (gce *GCE) SslCertificates() SslCertificates {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceSslCertificates
}
func (gce *GCE) TargetHttpProxies() TargetHttpProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceTargetHttpProxies
}
func (gce *GCE) TargetHttpsProxies() TargetHttpsProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceTargetHttpsProxies
}
func (gce *GCE) TargetPools() TargetPools {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceTargetPools
}
func (gce *GCE) UrlMaps() UrlMaps {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceUrlMaps
}
func (gce *GCE) Zones() Zones {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return gce.gceZones
}
func NewMockGCE(projectRouter ProjectRouter) *MockGCE {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mockAddressesObjs := map[meta.Key]*MockAddressesObj{}
 mockBackendServicesObjs := map[meta.Key]*MockBackendServicesObj{}
 mockDisksObjs := map[meta.Key]*MockDisksObj{}
 mockFirewallsObjs := map[meta.Key]*MockFirewallsObj{}
 mockForwardingRulesObjs := map[meta.Key]*MockForwardingRulesObj{}
 mockGlobalAddressesObjs := map[meta.Key]*MockGlobalAddressesObj{}
 mockGlobalForwardingRulesObjs := map[meta.Key]*MockGlobalForwardingRulesObj{}
 mockHealthChecksObjs := map[meta.Key]*MockHealthChecksObj{}
 mockHttpHealthChecksObjs := map[meta.Key]*MockHttpHealthChecksObj{}
 mockHttpsHealthChecksObjs := map[meta.Key]*MockHttpsHealthChecksObj{}
 mockInstanceGroupsObjs := map[meta.Key]*MockInstanceGroupsObj{}
 mockInstancesObjs := map[meta.Key]*MockInstancesObj{}
 mockNetworkEndpointGroupsObjs := map[meta.Key]*MockNetworkEndpointGroupsObj{}
 mockProjectsObjs := map[meta.Key]*MockProjectsObj{}
 mockRegionBackendServicesObjs := map[meta.Key]*MockRegionBackendServicesObj{}
 mockRegionDisksObjs := map[meta.Key]*MockRegionDisksObj{}
 mockRegionsObjs := map[meta.Key]*MockRegionsObj{}
 mockRoutesObjs := map[meta.Key]*MockRoutesObj{}
 mockSecurityPoliciesObjs := map[meta.Key]*MockSecurityPoliciesObj{}
 mockSslCertificatesObjs := map[meta.Key]*MockSslCertificatesObj{}
 mockTargetHttpProxiesObjs := map[meta.Key]*MockTargetHttpProxiesObj{}
 mockTargetHttpsProxiesObjs := map[meta.Key]*MockTargetHttpsProxiesObj{}
 mockTargetPoolsObjs := map[meta.Key]*MockTargetPoolsObj{}
 mockUrlMapsObjs := map[meta.Key]*MockUrlMapsObj{}
 mockZonesObjs := map[meta.Key]*MockZonesObj{}
 mock := &MockGCE{MockAddresses: NewMockAddresses(projectRouter, mockAddressesObjs), MockAlphaAddresses: NewMockAlphaAddresses(projectRouter, mockAddressesObjs), MockBetaAddresses: NewMockBetaAddresses(projectRouter, mockAddressesObjs), MockGlobalAddresses: NewMockGlobalAddresses(projectRouter, mockGlobalAddressesObjs), MockBackendServices: NewMockBackendServices(projectRouter, mockBackendServicesObjs), MockBetaBackendServices: NewMockBetaBackendServices(projectRouter, mockBackendServicesObjs), MockAlphaBackendServices: NewMockAlphaBackendServices(projectRouter, mockBackendServicesObjs), MockRegionBackendServices: NewMockRegionBackendServices(projectRouter, mockRegionBackendServicesObjs), MockAlphaRegionBackendServices: NewMockAlphaRegionBackendServices(projectRouter, mockRegionBackendServicesObjs), MockDisks: NewMockDisks(projectRouter, mockDisksObjs), MockRegionDisks: NewMockRegionDisks(projectRouter, mockRegionDisksObjs), MockFirewalls: NewMockFirewalls(projectRouter, mockFirewallsObjs), MockForwardingRules: NewMockForwardingRules(projectRouter, mockForwardingRulesObjs), MockAlphaForwardingRules: NewMockAlphaForwardingRules(projectRouter, mockForwardingRulesObjs), MockGlobalForwardingRules: NewMockGlobalForwardingRules(projectRouter, mockGlobalForwardingRulesObjs), MockHealthChecks: NewMockHealthChecks(projectRouter, mockHealthChecksObjs), MockBetaHealthChecks: NewMockBetaHealthChecks(projectRouter, mockHealthChecksObjs), MockAlphaHealthChecks: NewMockAlphaHealthChecks(projectRouter, mockHealthChecksObjs), MockHttpHealthChecks: NewMockHttpHealthChecks(projectRouter, mockHttpHealthChecksObjs), MockHttpsHealthChecks: NewMockHttpsHealthChecks(projectRouter, mockHttpsHealthChecksObjs), MockInstanceGroups: NewMockInstanceGroups(projectRouter, mockInstanceGroupsObjs), MockInstances: NewMockInstances(projectRouter, mockInstancesObjs), MockBetaInstances: NewMockBetaInstances(projectRouter, mockInstancesObjs), MockAlphaInstances: NewMockAlphaInstances(projectRouter, mockInstancesObjs), MockAlphaNetworkEndpointGroups: NewMockAlphaNetworkEndpointGroups(projectRouter, mockNetworkEndpointGroupsObjs), MockBetaNetworkEndpointGroups: NewMockBetaNetworkEndpointGroups(projectRouter, mockNetworkEndpointGroupsObjs), MockProjects: NewMockProjects(projectRouter, mockProjectsObjs), MockRegions: NewMockRegions(projectRouter, mockRegionsObjs), MockRoutes: NewMockRoutes(projectRouter, mockRoutesObjs), MockBetaSecurityPolicies: NewMockBetaSecurityPolicies(projectRouter, mockSecurityPoliciesObjs), MockSslCertificates: NewMockSslCertificates(projectRouter, mockSslCertificatesObjs), MockTargetHttpProxies: NewMockTargetHttpProxies(projectRouter, mockTargetHttpProxiesObjs), MockTargetHttpsProxies: NewMockTargetHttpsProxies(projectRouter, mockTargetHttpsProxiesObjs), MockTargetPools: NewMockTargetPools(projectRouter, mockTargetPoolsObjs), MockUrlMaps: NewMockUrlMaps(projectRouter, mockUrlMapsObjs), MockZones: NewMockZones(projectRouter, mockZonesObjs)}
 return mock
}

var _ Cloud = (*MockGCE)(nil)

type MockGCE struct {
 MockAddresses                  *MockAddresses
 MockAlphaAddresses             *MockAlphaAddresses
 MockBetaAddresses              *MockBetaAddresses
 MockGlobalAddresses            *MockGlobalAddresses
 MockBackendServices            *MockBackendServices
 MockBetaBackendServices        *MockBetaBackendServices
 MockAlphaBackendServices       *MockAlphaBackendServices
 MockRegionBackendServices      *MockRegionBackendServices
 MockAlphaRegionBackendServices *MockAlphaRegionBackendServices
 MockDisks                      *MockDisks
 MockRegionDisks                *MockRegionDisks
 MockFirewalls                  *MockFirewalls
 MockForwardingRules            *MockForwardingRules
 MockAlphaForwardingRules       *MockAlphaForwardingRules
 MockGlobalForwardingRules      *MockGlobalForwardingRules
 MockHealthChecks               *MockHealthChecks
 MockBetaHealthChecks           *MockBetaHealthChecks
 MockAlphaHealthChecks          *MockAlphaHealthChecks
 MockHttpHealthChecks           *MockHttpHealthChecks
 MockHttpsHealthChecks          *MockHttpsHealthChecks
 MockInstanceGroups             *MockInstanceGroups
 MockInstances                  *MockInstances
 MockBetaInstances              *MockBetaInstances
 MockAlphaInstances             *MockAlphaInstances
 MockAlphaNetworkEndpointGroups *MockAlphaNetworkEndpointGroups
 MockBetaNetworkEndpointGroups  *MockBetaNetworkEndpointGroups
 MockProjects                   *MockProjects
 MockRegions                    *MockRegions
 MockRoutes                     *MockRoutes
 MockBetaSecurityPolicies       *MockBetaSecurityPolicies
 MockSslCertificates            *MockSslCertificates
 MockTargetHttpProxies          *MockTargetHttpProxies
 MockTargetHttpsProxies         *MockTargetHttpsProxies
 MockTargetPools                *MockTargetPools
 MockUrlMaps                    *MockUrlMaps
 MockZones                      *MockZones
}

func (mock *MockGCE) Addresses() Addresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAddresses
}
func (mock *MockGCE) AlphaAddresses() AlphaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaAddresses
}
func (mock *MockGCE) BetaAddresses() BetaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaAddresses
}
func (mock *MockGCE) GlobalAddresses() GlobalAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockGlobalAddresses
}
func (mock *MockGCE) BackendServices() BackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBackendServices
}
func (mock *MockGCE) BetaBackendServices() BetaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaBackendServices
}
func (mock *MockGCE) AlphaBackendServices() AlphaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaBackendServices
}
func (mock *MockGCE) RegionBackendServices() RegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockRegionBackendServices
}
func (mock *MockGCE) AlphaRegionBackendServices() AlphaRegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaRegionBackendServices
}
func (mock *MockGCE) Disks() Disks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockDisks
}
func (mock *MockGCE) RegionDisks() RegionDisks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockRegionDisks
}
func (mock *MockGCE) Firewalls() Firewalls {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockFirewalls
}
func (mock *MockGCE) ForwardingRules() ForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockForwardingRules
}
func (mock *MockGCE) AlphaForwardingRules() AlphaForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaForwardingRules
}
func (mock *MockGCE) GlobalForwardingRules() GlobalForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockGlobalForwardingRules
}
func (mock *MockGCE) HealthChecks() HealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockHealthChecks
}
func (mock *MockGCE) BetaHealthChecks() BetaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaHealthChecks
}
func (mock *MockGCE) AlphaHealthChecks() AlphaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaHealthChecks
}
func (mock *MockGCE) HttpHealthChecks() HttpHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockHttpHealthChecks
}
func (mock *MockGCE) HttpsHealthChecks() HttpsHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockHttpsHealthChecks
}
func (mock *MockGCE) InstanceGroups() InstanceGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockInstanceGroups
}
func (mock *MockGCE) Instances() Instances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockInstances
}
func (mock *MockGCE) BetaInstances() BetaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaInstances
}
func (mock *MockGCE) AlphaInstances() AlphaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaInstances
}
func (mock *MockGCE) AlphaNetworkEndpointGroups() AlphaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockAlphaNetworkEndpointGroups
}
func (mock *MockGCE) BetaNetworkEndpointGroups() BetaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaNetworkEndpointGroups
}
func (mock *MockGCE) Projects() Projects {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockProjects
}
func (mock *MockGCE) Regions() Regions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockRegions
}
func (mock *MockGCE) Routes() Routes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockRoutes
}
func (mock *MockGCE) BetaSecurityPolicies() BetaSecurityPolicies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockBetaSecurityPolicies
}
func (mock *MockGCE) SslCertificates() SslCertificates {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockSslCertificates
}
func (mock *MockGCE) TargetHttpProxies() TargetHttpProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockTargetHttpProxies
}
func (mock *MockGCE) TargetHttpsProxies() TargetHttpsProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockTargetHttpsProxies
}
func (mock *MockGCE) TargetPools() TargetPools {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockTargetPools
}
func (mock *MockGCE) UrlMaps() UrlMaps {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockUrlMaps
}
func (mock *MockGCE) Zones() Zones {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return mock.MockZones
}

type MockAddressesObj struct{ Obj interface{} }

func (m *MockAddressesObj) ToAlpha() *alpha.Address {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.Address); ok {
  return ret
 }
 ret := &alpha.Address{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.Address via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockAddressesObj) ToBeta() *beta.Address {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.Address); ok {
  return ret
 }
 ret := &beta.Address{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.Address via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockAddressesObj) ToGA() *ga.Address {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Address); ok {
  return ret
 }
 ret := &ga.Address{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Address via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockBackendServicesObj struct{ Obj interface{} }

func (m *MockBackendServicesObj) ToAlpha() *alpha.BackendService {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.BackendService); ok {
  return ret
 }
 ret := &alpha.BackendService{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.BackendService via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockBackendServicesObj) ToBeta() *beta.BackendService {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.BackendService); ok {
  return ret
 }
 ret := &beta.BackendService{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.BackendService via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockBackendServicesObj) ToGA() *ga.BackendService {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.BackendService); ok {
  return ret
 }
 ret := &ga.BackendService{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.BackendService via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockDisksObj struct{ Obj interface{} }

func (m *MockDisksObj) ToGA() *ga.Disk {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Disk); ok {
  return ret
 }
 ret := &ga.Disk{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Disk via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockFirewallsObj struct{ Obj interface{} }

func (m *MockFirewallsObj) ToGA() *ga.Firewall {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Firewall); ok {
  return ret
 }
 ret := &ga.Firewall{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Firewall via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockForwardingRulesObj struct{ Obj interface{} }

func (m *MockForwardingRulesObj) ToAlpha() *alpha.ForwardingRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.ForwardingRule); ok {
  return ret
 }
 ret := &alpha.ForwardingRule{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.ForwardingRule via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockForwardingRulesObj) ToGA() *ga.ForwardingRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.ForwardingRule); ok {
  return ret
 }
 ret := &ga.ForwardingRule{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.ForwardingRule via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockGlobalAddressesObj struct{ Obj interface{} }

func (m *MockGlobalAddressesObj) ToGA() *ga.Address {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Address); ok {
  return ret
 }
 ret := &ga.Address{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Address via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockGlobalForwardingRulesObj struct{ Obj interface{} }

func (m *MockGlobalForwardingRulesObj) ToGA() *ga.ForwardingRule {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.ForwardingRule); ok {
  return ret
 }
 ret := &ga.ForwardingRule{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.ForwardingRule via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockHealthChecksObj struct{ Obj interface{} }

func (m *MockHealthChecksObj) ToAlpha() *alpha.HealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.HealthCheck); ok {
  return ret
 }
 ret := &alpha.HealthCheck{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.HealthCheck via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockHealthChecksObj) ToBeta() *beta.HealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.HealthCheck); ok {
  return ret
 }
 ret := &beta.HealthCheck{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.HealthCheck via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockHealthChecksObj) ToGA() *ga.HealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.HealthCheck); ok {
  return ret
 }
 ret := &ga.HealthCheck{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.HealthCheck via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockHttpHealthChecksObj struct{ Obj interface{} }

func (m *MockHttpHealthChecksObj) ToGA() *ga.HttpHealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.HttpHealthCheck); ok {
  return ret
 }
 ret := &ga.HttpHealthCheck{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.HttpHealthCheck via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockHttpsHealthChecksObj struct{ Obj interface{} }

func (m *MockHttpsHealthChecksObj) ToGA() *ga.HttpsHealthCheck {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.HttpsHealthCheck); ok {
  return ret
 }
 ret := &ga.HttpsHealthCheck{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.HttpsHealthCheck via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockInstanceGroupsObj struct{ Obj interface{} }

func (m *MockInstanceGroupsObj) ToGA() *ga.InstanceGroup {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.InstanceGroup); ok {
  return ret
 }
 ret := &ga.InstanceGroup{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.InstanceGroup via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockInstancesObj struct{ Obj interface{} }

func (m *MockInstancesObj) ToAlpha() *alpha.Instance {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.Instance); ok {
  return ret
 }
 ret := &alpha.Instance{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.Instance via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockInstancesObj) ToBeta() *beta.Instance {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.Instance); ok {
  return ret
 }
 ret := &beta.Instance{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.Instance via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockInstancesObj) ToGA() *ga.Instance {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Instance); ok {
  return ret
 }
 ret := &ga.Instance{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Instance via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockNetworkEndpointGroupsObj struct{ Obj interface{} }

func (m *MockNetworkEndpointGroupsObj) ToAlpha() *alpha.NetworkEndpointGroup {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.NetworkEndpointGroup); ok {
  return ret
 }
 ret := &alpha.NetworkEndpointGroup{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.NetworkEndpointGroup via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockNetworkEndpointGroupsObj) ToBeta() *beta.NetworkEndpointGroup {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.NetworkEndpointGroup); ok {
  return ret
 }
 ret := &beta.NetworkEndpointGroup{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.NetworkEndpointGroup via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockProjectsObj struct{ Obj interface{} }

func (m *MockProjectsObj) ToGA() *ga.Project {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Project); ok {
  return ret
 }
 ret := &ga.Project{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Project via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockRegionBackendServicesObj struct{ Obj interface{} }

func (m *MockRegionBackendServicesObj) ToAlpha() *alpha.BackendService {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*alpha.BackendService); ok {
  return ret
 }
 ret := &alpha.BackendService{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *alpha.BackendService via JSON: %v", m.Obj, err)
 }
 return ret
}
func (m *MockRegionBackendServicesObj) ToGA() *ga.BackendService {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.BackendService); ok {
  return ret
 }
 ret := &ga.BackendService{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.BackendService via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockRegionDisksObj struct{ Obj interface{} }

func (m *MockRegionDisksObj) ToGA() *ga.Disk {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Disk); ok {
  return ret
 }
 ret := &ga.Disk{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Disk via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockRegionsObj struct{ Obj interface{} }

func (m *MockRegionsObj) ToGA() *ga.Region {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Region); ok {
  return ret
 }
 ret := &ga.Region{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Region via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockRoutesObj struct{ Obj interface{} }

func (m *MockRoutesObj) ToGA() *ga.Route {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Route); ok {
  return ret
 }
 ret := &ga.Route{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Route via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockSecurityPoliciesObj struct{ Obj interface{} }

func (m *MockSecurityPoliciesObj) ToBeta() *beta.SecurityPolicy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*beta.SecurityPolicy); ok {
  return ret
 }
 ret := &beta.SecurityPolicy{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *beta.SecurityPolicy via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockSslCertificatesObj struct{ Obj interface{} }

func (m *MockSslCertificatesObj) ToGA() *ga.SslCertificate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.SslCertificate); ok {
  return ret
 }
 ret := &ga.SslCertificate{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.SslCertificate via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockTargetHttpProxiesObj struct{ Obj interface{} }

func (m *MockTargetHttpProxiesObj) ToGA() *ga.TargetHttpProxy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.TargetHttpProxy); ok {
  return ret
 }
 ret := &ga.TargetHttpProxy{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.TargetHttpProxy via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockTargetHttpsProxiesObj struct{ Obj interface{} }

func (m *MockTargetHttpsProxiesObj) ToGA() *ga.TargetHttpsProxy {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.TargetHttpsProxy); ok {
  return ret
 }
 ret := &ga.TargetHttpsProxy{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.TargetHttpsProxy via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockTargetPoolsObj struct{ Obj interface{} }

func (m *MockTargetPoolsObj) ToGA() *ga.TargetPool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.TargetPool); ok {
  return ret
 }
 ret := &ga.TargetPool{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.TargetPool via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockUrlMapsObj struct{ Obj interface{} }

func (m *MockUrlMapsObj) ToGA() *ga.UrlMap {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.UrlMap); ok {
  return ret
 }
 ret := &ga.UrlMap{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.UrlMap via JSON: %v", m.Obj, err)
 }
 return ret
}

type MockZonesObj struct{ Obj interface{} }

func (m *MockZonesObj) ToGA() *ga.Zone {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if ret, ok := m.Obj.(*ga.Zone); ok {
  return ret
 }
 ret := &ga.Zone{}
 if err := copyViaJSON(ret, m.Obj); err != nil {
  klog.Errorf("Could not convert %T to *ga.Zone via JSON: %v", m.Obj, err)
 }
 return ret
}

type Addresses interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Address, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*ga.Address, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockAddresses(pr ProjectRouter, objs map[meta.Key]*MockAddressesObj) *MockAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAddresses{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAddresses struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockAddressesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockAddresses) (bool, *ga.Address, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockAddresses) (bool, []*ga.Address, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Address, m *MockAddresses) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockAddresses) (bool, error)
 X             interface{}
}

func (m *MockAddresses) Get(ctx context.Context, key *meta.Key) (*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAddresses.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockAddresses.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAddresses %v not found", key)}
 klog.V(5).Infof("MockAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockAddresses.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAddresses.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Address
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockAddresses.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockAddresses) Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAddresses %v exists", key)}
  klog.V(5).Infof("MockAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "addresses")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "addresses", key)
 m.Objects[*key] = &MockAddressesObj{obj}
 klog.V(5).Infof("MockAddresses.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAddresses.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAddresses %v not found", key)}
  klog.V(5).Infof("MockAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAddresses.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAddresses) Obj(o *ga.Address) *MockAddressesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockAddressesObj{o}
}

type GCEAddresses struct{ s *Service }

func (g *GCEAddresses) Get(ctx context.Context, key *meta.Key) (*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAddresses.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAddresses.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Addresses"}
 klog.V(5).Infof("GCEAddresses.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAddresses.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Addresses.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAddresses.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAddresses.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Addresses"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAddresses.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.GA.Addresses.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Address
 f := func(l *ga.AddressList) error {
  klog.V(5).Infof("GCEAddresses.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAddresses.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAddresses) Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAddresses.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAddresses.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "Addresses"}
 klog.V(5).Infof("GCEAddresses.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAddresses.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.Addresses.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAddresses.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAddresses.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAddresses.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAddresses.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "Addresses"}
 klog.V(5).Infof("GCEAddresses.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAddresses.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Addresses.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAddresses.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type AlphaAddresses interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.Address, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*alpha.Address, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.Address) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockAlphaAddresses(pr ProjectRouter, objs map[meta.Key]*MockAddressesObj) *MockAlphaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaAddresses{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaAddresses struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockAddressesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockAlphaAddresses) (bool, *alpha.Address, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockAlphaAddresses) (bool, []*alpha.Address, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *alpha.Address, m *MockAlphaAddresses) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockAlphaAddresses) (bool, error)
 X             interface{}
}

func (m *MockAlphaAddresses) Get(ctx context.Context, key *meta.Key) (*alpha.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaAddresses.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaAddresses.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaAddresses %v not found", key)}
 klog.V(5).Infof("MockAlphaAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockAlphaAddresses.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaAddresses.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.Address
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaAddresses.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaAddresses) Insert(ctx context.Context, key *meta.Key, obj *alpha.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaAddresses %v exists", key)}
  klog.V(5).Infof("MockAlphaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "addresses")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "addresses", key)
 m.Objects[*key] = &MockAddressesObj{obj}
 klog.V(5).Infof("MockAlphaAddresses.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaAddresses.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaAddresses %v not found", key)}
  klog.V(5).Infof("MockAlphaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaAddresses.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaAddresses) Obj(o *alpha.Address) *MockAddressesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockAddressesObj{o}
}

type GCEAlphaAddresses struct{ s *Service }

func (g *GCEAlphaAddresses) Get(ctx context.Context, key *meta.Key) (*alpha.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaAddresses.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaAddresses.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "Addresses"}
 klog.V(5).Infof("GCEAlphaAddresses.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.Addresses.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaAddresses.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaAddresses.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "Addresses"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaAddresses.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.Alpha.Addresses.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.Address
 f := func(l *alpha.AddressList) error {
  klog.V(5).Infof("GCEAlphaAddresses.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaAddresses.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaAddresses) Insert(ctx context.Context, key *meta.Key, obj *alpha.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaAddresses.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaAddresses.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "Addresses"}
 klog.V(5).Infof("GCEAlphaAddresses.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.Addresses.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaAddresses.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaAddresses.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaAddresses.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "Addresses"}
 klog.V(5).Infof("GCEAlphaAddresses.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.Addresses.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaAddresses.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type BetaAddresses interface {
 Get(ctx context.Context, key *meta.Key) (*beta.Address, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*beta.Address, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.Address) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockBetaAddresses(pr ProjectRouter, objs map[meta.Key]*MockAddressesObj) *MockBetaAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaAddresses{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaAddresses struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockAddressesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockBetaAddresses) (bool, *beta.Address, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockBetaAddresses) (bool, []*beta.Address, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *beta.Address, m *MockBetaAddresses) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockBetaAddresses) (bool, error)
 X             interface{}
}

func (m *MockBetaAddresses) Get(ctx context.Context, key *meta.Key) (*beta.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaAddresses.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaAddresses.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaAddresses %v not found", key)}
 klog.V(5).Infof("MockBetaAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*beta.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockBetaAddresses.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaAddresses.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.Address
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaAddresses.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockBetaAddresses) Insert(ctx context.Context, key *meta.Key, obj *beta.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaAddresses %v exists", key)}
  klog.V(5).Infof("MockBetaAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "addresses")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "addresses", key)
 m.Objects[*key] = &MockAddressesObj{obj}
 klog.V(5).Infof("MockBetaAddresses.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaAddresses.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaAddresses %v not found", key)}
  klog.V(5).Infof("MockBetaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaAddresses.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaAddresses) Obj(o *beta.Address) *MockAddressesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockAddressesObj{o}
}

type GCEBetaAddresses struct{ s *Service }

func (g *GCEBetaAddresses) Get(ctx context.Context, key *meta.Key) (*beta.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaAddresses.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaAddresses.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "Addresses"}
 klog.V(5).Infof("GCEBetaAddresses.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaAddresses.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.Addresses.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaAddresses.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaAddresses) List(ctx context.Context, region string, fl *filter.F) ([]*beta.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaAddresses.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "Addresses"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaAddresses.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.Beta.Addresses.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.Address
 f := func(l *beta.AddressList) error {
  klog.V(5).Infof("GCEBetaAddresses.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaAddresses.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaAddresses) Insert(ctx context.Context, key *meta.Key, obj *beta.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaAddresses.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaAddresses.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "Addresses"}
 klog.V(5).Infof("GCEBetaAddresses.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaAddresses.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.Addresses.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaAddresses.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaAddresses.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaAddresses.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaAddresses.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Addresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "Addresses"}
 klog.V(5).Infof("GCEBetaAddresses.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaAddresses.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.Addresses.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaAddresses.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type GlobalAddresses interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Address, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.Address, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockGlobalAddresses(pr ProjectRouter, objs map[meta.Key]*MockGlobalAddressesObj) *MockGlobalAddresses {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockGlobalAddresses{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockGlobalAddresses struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockGlobalAddressesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockGlobalAddresses) (bool, *ga.Address, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockGlobalAddresses) (bool, []*ga.Address, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Address, m *MockGlobalAddresses) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockGlobalAddresses) (bool, error)
 X             interface{}
}

func (m *MockGlobalAddresses) Get(ctx context.Context, key *meta.Key) (*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockGlobalAddresses.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockGlobalAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockGlobalAddresses.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockGlobalAddresses %v not found", key)}
 klog.V(5).Infof("MockGlobalAddresses.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockGlobalAddresses) List(ctx context.Context, fl *filter.F) ([]*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockGlobalAddresses.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockGlobalAddresses.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Address
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockGlobalAddresses.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockGlobalAddresses) Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockGlobalAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockGlobalAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockGlobalAddresses %v exists", key)}
  klog.V(5).Infof("MockGlobalAddresses.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "addresses")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "addresses", key)
 m.Objects[*key] = &MockGlobalAddressesObj{obj}
 klog.V(5).Infof("MockGlobalAddresses.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockGlobalAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockGlobalAddresses.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockGlobalAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockGlobalAddresses %v not found", key)}
  klog.V(5).Infof("MockGlobalAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockGlobalAddresses.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockGlobalAddresses) Obj(o *ga.Address) *MockGlobalAddressesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockGlobalAddressesObj{o}
}

type GCEGlobalAddresses struct{ s *Service }

func (g *GCEGlobalAddresses) Get(ctx context.Context, key *meta.Key) (*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalAddresses.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalAddresses.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalAddresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "GlobalAddresses"}
 klog.V(5).Infof("GCEGlobalAddresses.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.GlobalAddresses.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEGlobalAddresses.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEGlobalAddresses) List(ctx context.Context, fl *filter.F) ([]*ga.Address, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalAddresses.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalAddresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "GlobalAddresses"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEGlobalAddresses.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.GlobalAddresses.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Address
 f := func(l *ga.AddressList) error {
  klog.V(5).Infof("GCEGlobalAddresses.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEGlobalAddresses.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEGlobalAddresses.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEGlobalAddresses) Insert(ctx context.Context, key *meta.Key, obj *ga.Address) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalAddresses.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalAddresses.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalAddresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "GlobalAddresses"}
 klog.V(5).Infof("GCEGlobalAddresses.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.GlobalAddresses.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEGlobalAddresses.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEGlobalAddresses) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalAddresses.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalAddresses.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalAddresses")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "GlobalAddresses"}
 klog.V(5).Infof("GCEGlobalAddresses.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.GlobalAddresses.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEGlobalAddresses.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEGlobalAddresses.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type BackendServices interface {
 Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.BackendService, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error
 Delete(ctx context.Context, key *meta.Key) error
 GetHealth(context.Context, *meta.Key, *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error)
 Patch(context.Context, *meta.Key, *ga.BackendService) error
 Update(context.Context, *meta.Key, *ga.BackendService) error
}

func NewMockBackendServices(pr ProjectRouter, objs map[meta.Key]*MockBackendServicesObj) *MockBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBackendServices{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBackendServices struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockBackendServicesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockBackendServices) (bool, *ga.BackendService, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockBackendServices) (bool, []*ga.BackendService, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.BackendService, m *MockBackendServices) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockBackendServices) (bool, error)
 GetHealthHook func(context.Context, *meta.Key, *ga.ResourceGroupReference, *MockBackendServices) (*ga.BackendServiceGroupHealth, error)
 PatchHook     func(context.Context, *meta.Key, *ga.BackendService, *MockBackendServices) error
 UpdateHook    func(context.Context, *meta.Key, *ga.BackendService, *MockBackendServices) error
 X             interface{}
}

func (m *MockBackendServices) Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBackendServices.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockBackendServices.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBackendServices %v not found", key)}
 klog.V(5).Infof("MockBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBackendServices) List(ctx context.Context, fl *filter.F) ([]*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockBackendServices.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBackendServices.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.BackendService
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockBackendServices.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockBackendServices) Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBackendServices %v exists", key)}
  klog.V(5).Infof("MockBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "backendServices")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "backendServices", key)
 m.Objects[*key] = &MockBackendServicesObj{obj}
 klog.V(5).Infof("MockBackendServices.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBackendServices.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBackendServices %v not found", key)}
  klog.V(5).Infof("MockBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBackendServices.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBackendServices) Obj(o *ga.BackendService) *MockBackendServicesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockBackendServicesObj{o}
}
func (m *MockBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHealthHook != nil {
  return m.GetHealthHook(ctx, key, arg0, m)
 }
 return nil, fmt.Errorf("GetHealthHook must be set")
}
func (m *MockBackendServices) Patch(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.PatchHook != nil {
  return m.PatchHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEBackendServices struct{ s *Service }

func (g *GCEBackendServices) Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.BackendServices.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBackendServices.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBackendServices) List(ctx context.Context, fl *filter.F) ([]*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "BackendServices"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBackendServices.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.BackendServices.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.BackendService
 f := func(l *ga.BackendServiceList) error {
  klog.V(5).Infof("GCEBackendServices.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBackendServices.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBackendServices) Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.BackendServices.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBackendServices.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBackendServices.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.BackendServices.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBackendServices.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.GetHealth(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.GetHealth(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "GetHealth", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.GetHealth(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.GetHealth(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.BackendServices.GetHealth(projectID, key.Name, arg0)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBackendServices.GetHealth(%v, %v, ...) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBackendServices) Patch(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.Patch(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.Patch(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Patch", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.Patch(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.Patch(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.BackendServices.Patch(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBackendServices.Patch(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBackendServices.Patch(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBackendServices.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBackendServices.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBackendServices.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBackendServices.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.BackendServices.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type BetaBackendServices interface {
 Get(ctx context.Context, key *meta.Key) (*beta.BackendService, error)
 List(ctx context.Context, fl *filter.F) ([]*beta.BackendService, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.BackendService) error
 Delete(ctx context.Context, key *meta.Key) error
 SetSecurityPolicy(context.Context, *meta.Key, *beta.SecurityPolicyReference) error
 Update(context.Context, *meta.Key, *beta.BackendService) error
}

func NewMockBetaBackendServices(pr ProjectRouter, objs map[meta.Key]*MockBackendServicesObj) *MockBetaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaBackendServices{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaBackendServices struct {
 Lock                  sync.Mutex
 ProjectRouter         ProjectRouter
 Objects               map[meta.Key]*MockBackendServicesObj
 GetError              map[meta.Key]error
 ListError             *error
 InsertError           map[meta.Key]error
 DeleteError           map[meta.Key]error
 GetHook               func(ctx context.Context, key *meta.Key, m *MockBetaBackendServices) (bool, *beta.BackendService, error)
 ListHook              func(ctx context.Context, fl *filter.F, m *MockBetaBackendServices) (bool, []*beta.BackendService, error)
 InsertHook            func(ctx context.Context, key *meta.Key, obj *beta.BackendService, m *MockBetaBackendServices) (bool, error)
 DeleteHook            func(ctx context.Context, key *meta.Key, m *MockBetaBackendServices) (bool, error)
 SetSecurityPolicyHook func(context.Context, *meta.Key, *beta.SecurityPolicyReference, *MockBetaBackendServices) error
 UpdateHook            func(context.Context, *meta.Key, *beta.BackendService, *MockBetaBackendServices) error
 X                     interface{}
}

func (m *MockBetaBackendServices) Get(ctx context.Context, key *meta.Key) (*beta.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaBackendServices.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaBackendServices.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaBackendServices %v not found", key)}
 klog.V(5).Infof("MockBetaBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaBackendServices) List(ctx context.Context, fl *filter.F) ([]*beta.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockBetaBackendServices.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaBackendServices.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.BackendService
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaBackendServices.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockBetaBackendServices) Insert(ctx context.Context, key *meta.Key, obj *beta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaBackendServices %v exists", key)}
  klog.V(5).Infof("MockBetaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "backendServices")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "backendServices", key)
 m.Objects[*key] = &MockBackendServicesObj{obj}
 klog.V(5).Infof("MockBetaBackendServices.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaBackendServices %v not found", key)}
  klog.V(5).Infof("MockBetaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaBackendServices.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaBackendServices) Obj(o *beta.BackendService) *MockBackendServicesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockBackendServicesObj{o}
}
func (m *MockBetaBackendServices) SetSecurityPolicy(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetSecurityPolicyHook != nil {
  return m.SetSecurityPolicyHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *beta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEBetaBackendServices struct{ s *Service }

func (g *GCEBetaBackendServices) Get(ctx context.Context, key *meta.Key) (*beta.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaBackendServices.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBetaBackendServices.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.BackendServices.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaBackendServices.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaBackendServices) List(ctx context.Context, fl *filter.F) ([]*beta.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "BackendServices"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaBackendServices.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.Beta.BackendServices.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.BackendService
 f := func(l *beta.BackendServiceList) error {
  klog.V(5).Infof("GCEBetaBackendServices.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaBackendServices.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaBackendServices) Insert(ctx context.Context, key *meta.Key, obj *beta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaBackendServices.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBetaBackendServices.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.BackendServices.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaBackendServices.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaBackendServices.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBetaBackendServices.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.BackendServices.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBetaBackendServices) SetSecurityPolicy(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetSecurityPolicy", Version: meta.Version("beta"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.BackendServices.SetSecurityPolicy(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaBackendServices.SetSecurityPolicy(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *beta.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaBackendServices.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaBackendServices.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("beta"), Service: "BackendServices"}
 klog.V(5).Infof("GCEBetaBackendServices.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.BackendServices.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type AlphaBackendServices interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error)
 List(ctx context.Context, fl *filter.F) ([]*alpha.BackendService, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error
 Delete(ctx context.Context, key *meta.Key) error
 SetSecurityPolicy(context.Context, *meta.Key, *alpha.SecurityPolicyReference) error
 Update(context.Context, *meta.Key, *alpha.BackendService) error
}

func NewMockAlphaBackendServices(pr ProjectRouter, objs map[meta.Key]*MockBackendServicesObj) *MockAlphaBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaBackendServices{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaBackendServices struct {
 Lock                  sync.Mutex
 ProjectRouter         ProjectRouter
 Objects               map[meta.Key]*MockBackendServicesObj
 GetError              map[meta.Key]error
 ListError             *error
 InsertError           map[meta.Key]error
 DeleteError           map[meta.Key]error
 GetHook               func(ctx context.Context, key *meta.Key, m *MockAlphaBackendServices) (bool, *alpha.BackendService, error)
 ListHook              func(ctx context.Context, fl *filter.F, m *MockAlphaBackendServices) (bool, []*alpha.BackendService, error)
 InsertHook            func(ctx context.Context, key *meta.Key, obj *alpha.BackendService, m *MockAlphaBackendServices) (bool, error)
 DeleteHook            func(ctx context.Context, key *meta.Key, m *MockAlphaBackendServices) (bool, error)
 SetSecurityPolicyHook func(context.Context, *meta.Key, *alpha.SecurityPolicyReference, *MockAlphaBackendServices) error
 UpdateHook            func(context.Context, *meta.Key, *alpha.BackendService, *MockAlphaBackendServices) error
 X                     interface{}
}

func (m *MockAlphaBackendServices) Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaBackendServices.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaBackendServices.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaBackendServices %v not found", key)}
 klog.V(5).Infof("MockAlphaBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaBackendServices) List(ctx context.Context, fl *filter.F) ([]*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockAlphaBackendServices.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaBackendServices.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.BackendService
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaBackendServices.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaBackendServices) Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaBackendServices %v exists", key)}
  klog.V(5).Infof("MockAlphaBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "backendServices")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "backendServices", key)
 m.Objects[*key] = &MockBackendServicesObj{obj}
 klog.V(5).Infof("MockAlphaBackendServices.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaBackendServices %v not found", key)}
  klog.V(5).Infof("MockAlphaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaBackendServices.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaBackendServices) Obj(o *alpha.BackendService) *MockBackendServicesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockBackendServicesObj{o}
}
func (m *MockAlphaBackendServices) SetSecurityPolicy(ctx context.Context, key *meta.Key, arg0 *alpha.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetSecurityPolicyHook != nil {
  return m.SetSecurityPolicyHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockAlphaBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEAlphaBackendServices struct{ s *Service }

func (g *GCEAlphaBackendServices) Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaBackendServices.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "BackendServices"}
 klog.V(5).Infof("GCEAlphaBackendServices.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.BackendServices.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaBackendServices.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaBackendServices) List(ctx context.Context, fl *filter.F) ([]*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "BackendServices"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaBackendServices.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.Alpha.BackendServices.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.BackendService
 f := func(l *alpha.BackendServiceList) error {
  klog.V(5).Infof("GCEAlphaBackendServices.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaBackendServices.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaBackendServices) Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaBackendServices.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "BackendServices"}
 klog.V(5).Infof("GCEAlphaBackendServices.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.BackendServices.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaBackendServices.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaBackendServices.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "BackendServices"}
 klog.V(5).Infof("GCEAlphaBackendServices.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.BackendServices.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaBackendServices.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEAlphaBackendServices) SetSecurityPolicy(ctx context.Context, key *meta.Key, arg0 *alpha.SecurityPolicyReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetSecurityPolicy", Version: meta.Version("alpha"), Service: "BackendServices"}
 klog.V(5).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.BackendServices.SetSecurityPolicy(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaBackendServices.SetSecurityPolicy(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEAlphaBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaBackendServices.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaBackendServices.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "BackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("alpha"), Service: "BackendServices"}
 klog.V(5).Infof("GCEAlphaBackendServices.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.BackendServices.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type RegionBackendServices interface {
 Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*ga.BackendService, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error
 Delete(ctx context.Context, key *meta.Key) error
 GetHealth(context.Context, *meta.Key, *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error)
 Update(context.Context, *meta.Key, *ga.BackendService) error
}

func NewMockRegionBackendServices(pr ProjectRouter, objs map[meta.Key]*MockRegionBackendServicesObj) *MockRegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockRegionBackendServices{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockRegionBackendServices struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockRegionBackendServicesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockRegionBackendServices) (bool, *ga.BackendService, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockRegionBackendServices) (bool, []*ga.BackendService, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.BackendService, m *MockRegionBackendServices) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockRegionBackendServices) (bool, error)
 GetHealthHook func(context.Context, *meta.Key, *ga.ResourceGroupReference, *MockRegionBackendServices) (*ga.BackendServiceGroupHealth, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.BackendService, *MockRegionBackendServices) error
 X             interface{}
}

func (m *MockRegionBackendServices) Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRegionBackendServices.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockRegionBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockRegionBackendServices.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRegionBackendServices %v not found", key)}
 klog.V(5).Infof("MockRegionBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockRegionBackendServices) List(ctx context.Context, region string, fl *filter.F) ([]*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockRegionBackendServices.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockRegionBackendServices.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.BackendService
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockRegionBackendServices.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockRegionBackendServices) Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockRegionBackendServices %v exists", key)}
  klog.V(5).Infof("MockRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "backendServices")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "backendServices", key)
 m.Objects[*key] = &MockRegionBackendServicesObj{obj}
 klog.V(5).Infof("MockRegionBackendServices.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockRegionBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRegionBackendServices %v not found", key)}
  klog.V(5).Infof("MockRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockRegionBackendServices.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockRegionBackendServices) Obj(o *ga.BackendService) *MockRegionBackendServicesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockRegionBackendServicesObj{o}
}
func (m *MockRegionBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHealthHook != nil {
  return m.GetHealthHook(ctx, key, arg0, m)
 }
 return nil, fmt.Errorf("GetHealthHook must be set")
}
func (m *MockRegionBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCERegionBackendServices struct{ s *Service }

func (g *GCERegionBackendServices) Get(ctx context.Context, key *meta.Key) (*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionBackendServices.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCERegionBackendServices.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.RegionBackendServices.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCERegionBackendServices.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCERegionBackendServices) List(ctx context.Context, region string, fl *filter.F) ([]*ga.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCERegionBackendServices.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.GA.RegionBackendServices.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.BackendService
 f := func(l *ga.BackendServiceList) error {
  klog.V(5).Infof("GCERegionBackendServices.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCERegionBackendServices.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCERegionBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCERegionBackendServices) Insert(ctx context.Context, key *meta.Key, obj *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionBackendServices.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCERegionBackendServices.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.RegionBackendServices.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionBackendServices.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCERegionBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionBackendServices.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCERegionBackendServices.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.RegionBackendServices.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCERegionBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *ga.ResourceGroupReference) (*ga.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.GetHealth(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionBackendServices.GetHealth(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "GetHealth", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCERegionBackendServices.GetHealth(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.GetHealth(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.RegionBackendServices.GetHealth(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCERegionBackendServices.GetHealth(%v, %v, ...) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCERegionBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *ga.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionBackendServices.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionBackendServices.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCERegionBackendServices.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.RegionBackendServices.Update(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type AlphaRegionBackendServices interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*alpha.BackendService, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error
 Delete(ctx context.Context, key *meta.Key) error
 GetHealth(context.Context, *meta.Key, *alpha.ResourceGroupReference) (*alpha.BackendServiceGroupHealth, error)
 Update(context.Context, *meta.Key, *alpha.BackendService) error
}

func NewMockAlphaRegionBackendServices(pr ProjectRouter, objs map[meta.Key]*MockRegionBackendServicesObj) *MockAlphaRegionBackendServices {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaRegionBackendServices{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaRegionBackendServices struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockRegionBackendServicesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockAlphaRegionBackendServices) (bool, *alpha.BackendService, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockAlphaRegionBackendServices) (bool, []*alpha.BackendService, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *alpha.BackendService, m *MockAlphaRegionBackendServices) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockAlphaRegionBackendServices) (bool, error)
 GetHealthHook func(context.Context, *meta.Key, *alpha.ResourceGroupReference, *MockAlphaRegionBackendServices) (*alpha.BackendServiceGroupHealth, error)
 UpdateHook    func(context.Context, *meta.Key, *alpha.BackendService, *MockAlphaRegionBackendServices) error
 X             interface{}
}

func (m *MockAlphaRegionBackendServices) Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaRegionBackendServices.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaRegionBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaRegionBackendServices.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaRegionBackendServices %v not found", key)}
 klog.V(5).Infof("MockAlphaRegionBackendServices.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaRegionBackendServices) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockAlphaRegionBackendServices.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaRegionBackendServices.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.BackendService
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaRegionBackendServices.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaRegionBackendServices) Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaRegionBackendServices %v exists", key)}
  klog.V(5).Infof("MockAlphaRegionBackendServices.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "backendServices")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "backendServices", key)
 m.Objects[*key] = &MockRegionBackendServicesObj{obj}
 klog.V(5).Infof("MockAlphaRegionBackendServices.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaRegionBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaRegionBackendServices %v not found", key)}
  klog.V(5).Infof("MockAlphaRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaRegionBackendServices.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaRegionBackendServices) Obj(o *alpha.BackendService) *MockRegionBackendServicesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockRegionBackendServicesObj{o}
}
func (m *MockAlphaRegionBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *alpha.ResourceGroupReference) (*alpha.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHealthHook != nil {
  return m.GetHealthHook(ctx, key, arg0, m)
 }
 return nil, fmt.Errorf("GetHealthHook must be set")
}
func (m *MockAlphaRegionBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEAlphaRegionBackendServices struct{ s *Service }

func (g *GCEAlphaRegionBackendServices) Get(ctx context.Context, key *meta.Key) (*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaRegionBackendServices.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.RegionBackendServices.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaRegionBackendServices.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaRegionBackendServices) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.BackendService, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaRegionBackendServices.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.Alpha.RegionBackendServices.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.BackendService
 f := func(l *alpha.BackendServiceList) error {
  klog.V(5).Infof("GCEAlphaRegionBackendServices.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaRegionBackendServices.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaRegionBackendServices) Insert(ctx context.Context, key *meta.Key, obj *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.RegionBackendServices.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaRegionBackendServices.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaRegionBackendServices) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.RegionBackendServices.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaRegionBackendServices.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEAlphaRegionBackendServices) GetHealth(ctx context.Context, key *meta.Key, arg0 *alpha.ResourceGroupReference) (*alpha.BackendServiceGroupHealth, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.GetHealth(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaRegionBackendServices.GetHealth(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "GetHealth", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCEAlphaRegionBackendServices.GetHealth(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.GetHealth(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.RegionBackendServices.GetHealth(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaRegionBackendServices.GetHealth(%v, %v, ...) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaRegionBackendServices) Update(ctx context.Context, key *meta.Key, arg0 *alpha.BackendService) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "RegionBackendServices")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("alpha"), Service: "RegionBackendServices"}
 klog.V(5).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.RegionBackendServices.Update(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaRegionBackendServices.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type Disks interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Disk, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Disk, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error
 Delete(ctx context.Context, key *meta.Key) error
 Resize(context.Context, *meta.Key, *ga.DisksResizeRequest) error
}

func NewMockDisks(pr ProjectRouter, objs map[meta.Key]*MockDisksObj) *MockDisks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockDisks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockDisks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockDisksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockDisks) (bool, *ga.Disk, error)
 ListHook      func(ctx context.Context, zone string, fl *filter.F, m *MockDisks) (bool, []*ga.Disk, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Disk, m *MockDisks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockDisks) (bool, error)
 ResizeHook    func(context.Context, *meta.Key, *ga.DisksResizeRequest, *MockDisks) error
 X             interface{}
}

func (m *MockDisks) Get(ctx context.Context, key *meta.Key) (*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockDisks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockDisks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockDisks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockDisks %v not found", key)}
 klog.V(5).Infof("MockDisks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockDisks) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockDisks.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockDisks.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Disk
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockDisks.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockDisks) Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockDisks %v exists", key)}
  klog.V(5).Infof("MockDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "disks")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "disks", key)
 m.Objects[*key] = &MockDisksObj{obj}
 klog.V(5).Infof("MockDisks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockDisks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockDisks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockDisks %v not found", key)}
  klog.V(5).Infof("MockDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockDisks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockDisks) Obj(o *ga.Disk) *MockDisksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockDisksObj{o}
}
func (m *MockDisks) Resize(ctx context.Context, key *meta.Key, arg0 *ga.DisksResizeRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ResizeHook != nil {
  return m.ResizeHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEDisks struct{ s *Service }

func (g *GCEDisks) Get(ctx context.Context, key *meta.Key) (*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEDisks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEDisks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Disks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Disks"}
 klog.V(5).Infof("GCEDisks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEDisks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Disks.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEDisks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEDisks) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEDisks.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Disks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Disks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEDisks.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.GA.Disks.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Disk
 f := func(l *ga.DiskList) error {
  klog.V(5).Infof("GCEDisks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEDisks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEDisks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEDisks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEDisks) Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEDisks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEDisks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Disks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "Disks"}
 klog.V(5).Infof("GCEDisks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEDisks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.Disks.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEDisks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEDisks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEDisks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEDisks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEDisks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Disks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "Disks"}
 klog.V(5).Infof("GCEDisks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEDisks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Disks.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEDisks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEDisks) Resize(ctx context.Context, key *meta.Key, arg0 *ga.DisksResizeRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEDisks.Resize(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEDisks.Resize(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Disks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Resize", Version: meta.Version("ga"), Service: "Disks"}
 klog.V(5).Infof("GCEDisks.Resize(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEDisks.Resize(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Disks.Resize(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEDisks.Resize(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEDisks.Resize(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type RegionDisks interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Disk, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*ga.Disk, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error
 Delete(ctx context.Context, key *meta.Key) error
 Resize(context.Context, *meta.Key, *ga.RegionDisksResizeRequest) error
}

func NewMockRegionDisks(pr ProjectRouter, objs map[meta.Key]*MockRegionDisksObj) *MockRegionDisks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockRegionDisks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockRegionDisks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockRegionDisksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockRegionDisks) (bool, *ga.Disk, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockRegionDisks) (bool, []*ga.Disk, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Disk, m *MockRegionDisks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockRegionDisks) (bool, error)
 ResizeHook    func(context.Context, *meta.Key, *ga.RegionDisksResizeRequest, *MockRegionDisks) error
 X             interface{}
}

func (m *MockRegionDisks) Get(ctx context.Context, key *meta.Key) (*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRegionDisks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockRegionDisks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockRegionDisks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRegionDisks %v not found", key)}
 klog.V(5).Infof("MockRegionDisks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockRegionDisks) List(ctx context.Context, region string, fl *filter.F) ([]*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockRegionDisks.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockRegionDisks.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Disk
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockRegionDisks.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockRegionDisks) Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockRegionDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockRegionDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockRegionDisks %v exists", key)}
  klog.V(5).Infof("MockRegionDisks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "disks")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "disks", key)
 m.Objects[*key] = &MockRegionDisksObj{obj}
 klog.V(5).Infof("MockRegionDisks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockRegionDisks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRegionDisks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockRegionDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRegionDisks %v not found", key)}
  klog.V(5).Infof("MockRegionDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockRegionDisks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockRegionDisks) Obj(o *ga.Disk) *MockRegionDisksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockRegionDisksObj{o}
}
func (m *MockRegionDisks) Resize(ctx context.Context, key *meta.Key, arg0 *ga.RegionDisksResizeRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ResizeHook != nil {
  return m.ResizeHook(ctx, key, arg0, m)
 }
 return nil
}

type GCERegionDisks struct{ s *Service }

func (g *GCERegionDisks) Get(ctx context.Context, key *meta.Key) (*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionDisks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionDisks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionDisks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "RegionDisks"}
 klog.V(5).Infof("GCERegionDisks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionDisks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.RegionDisks.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCERegionDisks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCERegionDisks) List(ctx context.Context, region string, fl *filter.F) ([]*ga.Disk, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionDisks.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionDisks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "RegionDisks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCERegionDisks.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.GA.RegionDisks.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Disk
 f := func(l *ga.DiskList) error {
  klog.V(5).Infof("GCERegionDisks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCERegionDisks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCERegionDisks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCERegionDisks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCERegionDisks) Insert(ctx context.Context, key *meta.Key, obj *ga.Disk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionDisks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionDisks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionDisks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "RegionDisks"}
 klog.V(5).Infof("GCERegionDisks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionDisks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.RegionDisks.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionDisks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionDisks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCERegionDisks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionDisks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionDisks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionDisks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "RegionDisks"}
 klog.V(5).Infof("GCERegionDisks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionDisks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.RegionDisks.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionDisks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionDisks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCERegionDisks) Resize(ctx context.Context, key *meta.Key, arg0 *ga.RegionDisksResizeRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegionDisks.Resize(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegionDisks.Resize(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "RegionDisks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Resize", Version: meta.Version("ga"), Service: "RegionDisks"}
 klog.V(5).Infof("GCERegionDisks.Resize(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegionDisks.Resize(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.RegionDisks.Resize(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERegionDisks.Resize(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERegionDisks.Resize(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type Firewalls interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Firewall, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.Firewall, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Firewall) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *ga.Firewall) error
}

func NewMockFirewalls(pr ProjectRouter, objs map[meta.Key]*MockFirewallsObj) *MockFirewalls {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockFirewalls{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockFirewalls struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockFirewallsObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockFirewalls) (bool, *ga.Firewall, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockFirewalls) (bool, []*ga.Firewall, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Firewall, m *MockFirewalls) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockFirewalls) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.Firewall, *MockFirewalls) error
 X             interface{}
}

func (m *MockFirewalls) Get(ctx context.Context, key *meta.Key) (*ga.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockFirewalls.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockFirewalls.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockFirewalls.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockFirewalls %v not found", key)}
 klog.V(5).Infof("MockFirewalls.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockFirewalls) List(ctx context.Context, fl *filter.F) ([]*ga.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockFirewalls.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockFirewalls.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Firewall
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockFirewalls.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockFirewalls) Insert(ctx context.Context, key *meta.Key, obj *ga.Firewall) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockFirewalls.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockFirewalls.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockFirewalls %v exists", key)}
  klog.V(5).Infof("MockFirewalls.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "firewalls")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "firewalls", key)
 m.Objects[*key] = &MockFirewallsObj{obj}
 klog.V(5).Infof("MockFirewalls.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockFirewalls) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockFirewalls.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockFirewalls.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockFirewalls %v not found", key)}
  klog.V(5).Infof("MockFirewalls.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockFirewalls.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockFirewalls) Obj(o *ga.Firewall) *MockFirewallsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockFirewallsObj{o}
}
func (m *MockFirewalls) Update(ctx context.Context, key *meta.Key, arg0 *ga.Firewall) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEFirewalls struct{ s *Service }

func (g *GCEFirewalls) Get(ctx context.Context, key *meta.Key) (*ga.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEFirewalls.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEFirewalls.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Firewalls")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Firewalls"}
 klog.V(5).Infof("GCEFirewalls.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEFirewalls.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Firewalls.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEFirewalls.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEFirewalls) List(ctx context.Context, fl *filter.F) ([]*ga.Firewall, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEFirewalls.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Firewalls")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Firewalls"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEFirewalls.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.Firewalls.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Firewall
 f := func(l *ga.FirewallList) error {
  klog.V(5).Infof("GCEFirewalls.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEFirewalls.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEFirewalls.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEFirewalls.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEFirewalls) Insert(ctx context.Context, key *meta.Key, obj *ga.Firewall) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEFirewalls.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEFirewalls.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Firewalls")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "Firewalls"}
 klog.V(5).Infof("GCEFirewalls.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEFirewalls.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.Firewalls.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEFirewalls.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEFirewalls.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEFirewalls) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEFirewalls.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEFirewalls.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Firewalls")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "Firewalls"}
 klog.V(5).Infof("GCEFirewalls.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEFirewalls.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Firewalls.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEFirewalls.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEFirewalls.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEFirewalls) Update(ctx context.Context, key *meta.Key, arg0 *ga.Firewall) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEFirewalls.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEFirewalls.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Firewalls")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "Firewalls"}
 klog.V(5).Infof("GCEFirewalls.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEFirewalls.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Firewalls.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEFirewalls.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEFirewalls.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type ForwardingRules interface {
 Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*ga.ForwardingRule, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockForwardingRules(pr ProjectRouter, objs map[meta.Key]*MockForwardingRulesObj) *MockForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockForwardingRules{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockForwardingRules struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockForwardingRulesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockForwardingRules) (bool, *ga.ForwardingRule, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockForwardingRules) (bool, []*ga.ForwardingRule, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule, m *MockForwardingRules) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockForwardingRules) (bool, error)
 X             interface{}
}

func (m *MockForwardingRules) Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockForwardingRules.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockForwardingRules.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockForwardingRules %v not found", key)}
 klog.V(5).Infof("MockForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockForwardingRules) List(ctx context.Context, region string, fl *filter.F) ([]*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockForwardingRules.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockForwardingRules.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.ForwardingRule
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockForwardingRules.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockForwardingRules %v exists", key)}
  klog.V(5).Infof("MockForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "forwardingRules")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "forwardingRules", key)
 m.Objects[*key] = &MockForwardingRulesObj{obj}
 klog.V(5).Infof("MockForwardingRules.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockForwardingRules %v not found", key)}
  klog.V(5).Infof("MockForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockForwardingRules.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockForwardingRules) Obj(o *ga.ForwardingRule) *MockForwardingRulesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockForwardingRulesObj{o}
}

type GCEForwardingRules struct{ s *Service }

func (g *GCEForwardingRules) Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEForwardingRules.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEForwardingRules.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEForwardingRules.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEForwardingRules.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.ForwardingRules.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEForwardingRules.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEForwardingRules) List(ctx context.Context, region string, fl *filter.F) ([]*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEForwardingRules.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "ForwardingRules"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEForwardingRules.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.GA.ForwardingRules.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.ForwardingRule
 f := func(l *ga.ForwardingRuleList) error {
  klog.V(5).Infof("GCEForwardingRules.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEForwardingRules.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEForwardingRules.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEForwardingRules.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEForwardingRules.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEForwardingRules.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.ForwardingRules.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEForwardingRules.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEForwardingRules.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEForwardingRules.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEForwardingRules.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEForwardingRules.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEForwardingRules.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.ForwardingRules.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type AlphaForwardingRules interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.ForwardingRule, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*alpha.ForwardingRule, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.ForwardingRule) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockAlphaForwardingRules(pr ProjectRouter, objs map[meta.Key]*MockForwardingRulesObj) *MockAlphaForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaForwardingRules{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaForwardingRules struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockForwardingRulesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockAlphaForwardingRules) (bool, *alpha.ForwardingRule, error)
 ListHook      func(ctx context.Context, region string, fl *filter.F, m *MockAlphaForwardingRules) (bool, []*alpha.ForwardingRule, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *alpha.ForwardingRule, m *MockAlphaForwardingRules) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockAlphaForwardingRules) (bool, error)
 X             interface{}
}

func (m *MockAlphaForwardingRules) Get(ctx context.Context, key *meta.Key) (*alpha.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaForwardingRules.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaForwardingRules.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaForwardingRules %v not found", key)}
 klog.V(5).Infof("MockAlphaForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaForwardingRules) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockAlphaForwardingRules.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaForwardingRules.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.ForwardingRule
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaForwardingRules.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *alpha.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaForwardingRules %v exists", key)}
  klog.V(5).Infof("MockAlphaForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "forwardingRules")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "forwardingRules", key)
 m.Objects[*key] = &MockForwardingRulesObj{obj}
 klog.V(5).Infof("MockAlphaForwardingRules.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaForwardingRules %v not found", key)}
  klog.V(5).Infof("MockAlphaForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaForwardingRules.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaForwardingRules) Obj(o *alpha.ForwardingRule) *MockForwardingRulesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockForwardingRulesObj{o}
}

type GCEAlphaForwardingRules struct{ s *Service }

func (g *GCEAlphaForwardingRules) Get(ctx context.Context, key *meta.Key) (*alpha.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaForwardingRules.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaForwardingRules.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEAlphaForwardingRules.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.ForwardingRules.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaForwardingRules.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaForwardingRules) List(ctx context.Context, region string, fl *filter.F) ([]*alpha.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaForwardingRules.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "ForwardingRules"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaForwardingRules.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.Alpha.ForwardingRules.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.ForwardingRule
 f := func(l *alpha.ForwardingRuleList) error {
  klog.V(5).Infof("GCEAlphaForwardingRules.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaForwardingRules.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *alpha.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaForwardingRules.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaForwardingRules.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEAlphaForwardingRules.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.ForwardingRules.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaForwardingRules.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaForwardingRules.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaForwardingRules.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "ForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "ForwardingRules"}
 klog.V(5).Infof("GCEAlphaForwardingRules.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.ForwardingRules.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type GlobalForwardingRules interface {
 Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.ForwardingRule, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error
 Delete(ctx context.Context, key *meta.Key) error
 SetTarget(context.Context, *meta.Key, *ga.TargetReference) error
}

func NewMockGlobalForwardingRules(pr ProjectRouter, objs map[meta.Key]*MockGlobalForwardingRulesObj) *MockGlobalForwardingRules {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockGlobalForwardingRules{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockGlobalForwardingRules struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockGlobalForwardingRulesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockGlobalForwardingRules) (bool, *ga.ForwardingRule, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockGlobalForwardingRules) (bool, []*ga.ForwardingRule, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule, m *MockGlobalForwardingRules) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockGlobalForwardingRules) (bool, error)
 SetTargetHook func(context.Context, *meta.Key, *ga.TargetReference, *MockGlobalForwardingRules) error
 X             interface{}
}

func (m *MockGlobalForwardingRules) Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockGlobalForwardingRules.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockGlobalForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockGlobalForwardingRules.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockGlobalForwardingRules %v not found", key)}
 klog.V(5).Infof("MockGlobalForwardingRules.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockGlobalForwardingRules) List(ctx context.Context, fl *filter.F) ([]*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockGlobalForwardingRules.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockGlobalForwardingRules.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.ForwardingRule
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockGlobalForwardingRules.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockGlobalForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockGlobalForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockGlobalForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockGlobalForwardingRules %v exists", key)}
  klog.V(5).Infof("MockGlobalForwardingRules.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "forwardingRules")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "forwardingRules", key)
 m.Objects[*key] = &MockGlobalForwardingRulesObj{obj}
 klog.V(5).Infof("MockGlobalForwardingRules.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockGlobalForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockGlobalForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockGlobalForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockGlobalForwardingRules %v not found", key)}
  klog.V(5).Infof("MockGlobalForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockGlobalForwardingRules.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockGlobalForwardingRules) Obj(o *ga.ForwardingRule) *MockGlobalForwardingRulesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockGlobalForwardingRulesObj{o}
}
func (m *MockGlobalForwardingRules) SetTarget(ctx context.Context, key *meta.Key, arg0 *ga.TargetReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetTargetHook != nil {
  return m.SetTargetHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEGlobalForwardingRules struct{ s *Service }

func (g *GCEGlobalForwardingRules) Get(ctx context.Context, key *meta.Key) (*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalForwardingRules.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalForwardingRules.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "GlobalForwardingRules"}
 klog.V(5).Infof("GCEGlobalForwardingRules.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.GlobalForwardingRules.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEGlobalForwardingRules.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEGlobalForwardingRules) List(ctx context.Context, fl *filter.F) ([]*ga.ForwardingRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalForwardingRules.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "GlobalForwardingRules"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEGlobalForwardingRules.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.GlobalForwardingRules.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.ForwardingRule
 f := func(l *ga.ForwardingRuleList) error {
  klog.V(5).Infof("GCEGlobalForwardingRules.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEGlobalForwardingRules.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEGlobalForwardingRules.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEGlobalForwardingRules) Insert(ctx context.Context, key *meta.Key, obj *ga.ForwardingRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalForwardingRules.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalForwardingRules.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "GlobalForwardingRules"}
 klog.V(5).Infof("GCEGlobalForwardingRules.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.GlobalForwardingRules.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEGlobalForwardingRules.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEGlobalForwardingRules) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalForwardingRules.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalForwardingRules.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "GlobalForwardingRules"}
 klog.V(5).Infof("GCEGlobalForwardingRules.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.GlobalForwardingRules.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEGlobalForwardingRules.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEGlobalForwardingRules) SetTarget(ctx context.Context, key *meta.Key, arg0 *ga.TargetReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "GlobalForwardingRules")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetTarget", Version: meta.Version("ga"), Service: "GlobalForwardingRules"}
 klog.V(5).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.GlobalForwardingRules.SetTarget(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEGlobalForwardingRules.SetTarget(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type HealthChecks interface {
 Get(ctx context.Context, key *meta.Key) (*ga.HealthCheck, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.HealthCheck, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.HealthCheck) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *ga.HealthCheck) error
}

func NewMockHealthChecks(pr ProjectRouter, objs map[meta.Key]*MockHealthChecksObj) *MockHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockHealthChecks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockHealthChecks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockHealthChecksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockHealthChecks) (bool, *ga.HealthCheck, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockHealthChecks) (bool, []*ga.HealthCheck, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.HealthCheck, m *MockHealthChecks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockHealthChecks) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.HealthCheck, *MockHealthChecks) error
 X             interface{}
}

func (m *MockHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHealthChecks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockHealthChecks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHealthChecks %v not found", key)}
 klog.V(5).Infof("MockHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockHealthChecks.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockHealthChecks.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.HealthCheck
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockHealthChecks.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockHealthChecks %v exists", key)}
  klog.V(5).Infof("MockHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "healthChecks")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "healthChecks", key)
 m.Objects[*key] = &MockHealthChecksObj{obj}
 klog.V(5).Infof("MockHealthChecks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHealthChecks %v not found", key)}
  klog.V(5).Infof("MockHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockHealthChecks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockHealthChecks) Obj(o *ga.HealthCheck) *MockHealthChecksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockHealthChecksObj{o}
}
func (m *MockHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEHealthChecks struct{ s *Service }

func (g *GCEHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHealthChecks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHealthChecks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEHealthChecks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHealthChecks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.HealthChecks.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEHealthChecks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHealthChecks.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "HealthChecks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEHealthChecks.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.HealthChecks.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.HealthCheck
 f := func(l *ga.HealthCheckList) error {
  klog.V(5).Infof("GCEHealthChecks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEHealthChecks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHealthChecks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEHealthChecks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEHealthChecks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHealthChecks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.HealthChecks.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHealthChecks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHealthChecks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHealthChecks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHealthChecks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEHealthChecks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHealthChecks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HealthChecks.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHealthChecks.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHealthChecks.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEHealthChecks.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHealthChecks.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HealthChecks.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type BetaHealthChecks interface {
 Get(ctx context.Context, key *meta.Key) (*beta.HealthCheck, error)
 List(ctx context.Context, fl *filter.F) ([]*beta.HealthCheck, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.HealthCheck) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *beta.HealthCheck) error
}

func NewMockBetaHealthChecks(pr ProjectRouter, objs map[meta.Key]*MockHealthChecksObj) *MockBetaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaHealthChecks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaHealthChecks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockHealthChecksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockBetaHealthChecks) (bool, *beta.HealthCheck, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockBetaHealthChecks) (bool, []*beta.HealthCheck, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *beta.HealthCheck, m *MockBetaHealthChecks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockBetaHealthChecks) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *beta.HealthCheck, *MockBetaHealthChecks) error
 X             interface{}
}

func (m *MockBetaHealthChecks) Get(ctx context.Context, key *meta.Key) (*beta.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaHealthChecks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaHealthChecks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaHealthChecks %v not found", key)}
 klog.V(5).Infof("MockBetaHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaHealthChecks) List(ctx context.Context, fl *filter.F) ([]*beta.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockBetaHealthChecks.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaHealthChecks.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.HealthCheck
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaHealthChecks.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockBetaHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *beta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaHealthChecks %v exists", key)}
  klog.V(5).Infof("MockBetaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "healthChecks")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "healthChecks", key)
 m.Objects[*key] = &MockHealthChecksObj{obj}
 klog.V(5).Infof("MockBetaHealthChecks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaHealthChecks %v not found", key)}
  klog.V(5).Infof("MockBetaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaHealthChecks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaHealthChecks) Obj(o *beta.HealthCheck) *MockHealthChecksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockHealthChecksObj{o}
}
func (m *MockBetaHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *beta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEBetaHealthChecks struct{ s *Service }

func (g *GCEBetaHealthChecks) Get(ctx context.Context, key *meta.Key) (*beta.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaHealthChecks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaHealthChecks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEBetaHealthChecks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.HealthChecks.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaHealthChecks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaHealthChecks) List(ctx context.Context, fl *filter.F) ([]*beta.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaHealthChecks.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "HealthChecks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaHealthChecks.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.Beta.HealthChecks.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.HealthCheck
 f := func(l *beta.HealthCheckList) error {
  klog.V(5).Infof("GCEBetaHealthChecks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaHealthChecks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *beta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaHealthChecks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaHealthChecks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEBetaHealthChecks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.HealthChecks.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaHealthChecks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaHealthChecks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaHealthChecks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEBetaHealthChecks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.HealthChecks.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBetaHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *beta.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaHealthChecks.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaHealthChecks.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("beta"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEBetaHealthChecks.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.HealthChecks.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type AlphaHealthChecks interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.HealthCheck, error)
 List(ctx context.Context, fl *filter.F) ([]*alpha.HealthCheck, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.HealthCheck) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *alpha.HealthCheck) error
}

func NewMockAlphaHealthChecks(pr ProjectRouter, objs map[meta.Key]*MockHealthChecksObj) *MockAlphaHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaHealthChecks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaHealthChecks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockHealthChecksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockAlphaHealthChecks) (bool, *alpha.HealthCheck, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockAlphaHealthChecks) (bool, []*alpha.HealthCheck, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *alpha.HealthCheck, m *MockAlphaHealthChecks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockAlphaHealthChecks) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *alpha.HealthCheck, *MockAlphaHealthChecks) error
 X             interface{}
}

func (m *MockAlphaHealthChecks) Get(ctx context.Context, key *meta.Key) (*alpha.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaHealthChecks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaHealthChecks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaHealthChecks %v not found", key)}
 klog.V(5).Infof("MockAlphaHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaHealthChecks) List(ctx context.Context, fl *filter.F) ([]*alpha.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockAlphaHealthChecks.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaHealthChecks.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.HealthCheck
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaHealthChecks.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *alpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaHealthChecks %v exists", key)}
  klog.V(5).Infof("MockAlphaHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "healthChecks")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "healthChecks", key)
 m.Objects[*key] = &MockHealthChecksObj{obj}
 klog.V(5).Infof("MockAlphaHealthChecks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaHealthChecks %v not found", key)}
  klog.V(5).Infof("MockAlphaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaHealthChecks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaHealthChecks) Obj(o *alpha.HealthCheck) *MockHealthChecksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockHealthChecksObj{o}
}
func (m *MockAlphaHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *alpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEAlphaHealthChecks struct{ s *Service }

func (g *GCEAlphaHealthChecks) Get(ctx context.Context, key *meta.Key) (*alpha.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaHealthChecks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaHealthChecks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEAlphaHealthChecks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.HealthChecks.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaHealthChecks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaHealthChecks) List(ctx context.Context, fl *filter.F) ([]*alpha.HealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaHealthChecks.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "HealthChecks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaHealthChecks.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.Alpha.HealthChecks.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.HealthCheck
 f := func(l *alpha.HealthCheckList) error {
  klog.V(5).Infof("GCEAlphaHealthChecks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaHealthChecks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *alpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaHealthChecks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaHealthChecks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEAlphaHealthChecks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.HealthChecks.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaHealthChecks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaHealthChecks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaHealthChecks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEAlphaHealthChecks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.HealthChecks.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEAlphaHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *alpha.HealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "HealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("alpha"), Service: "HealthChecks"}
 klog.V(5).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.HealthChecks.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type HttpHealthChecks interface {
 Get(ctx context.Context, key *meta.Key) (*ga.HttpHealthCheck, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.HttpHealthCheck, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.HttpHealthCheck) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *ga.HttpHealthCheck) error
}

func NewMockHttpHealthChecks(pr ProjectRouter, objs map[meta.Key]*MockHttpHealthChecksObj) *MockHttpHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockHttpHealthChecks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockHttpHealthChecks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockHttpHealthChecksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockHttpHealthChecks) (bool, *ga.HttpHealthCheck, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockHttpHealthChecks) (bool, []*ga.HttpHealthCheck, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.HttpHealthCheck, m *MockHttpHealthChecks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockHttpHealthChecks) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.HttpHealthCheck, *MockHttpHealthChecks) error
 X             interface{}
}

func (m *MockHttpHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHttpHealthChecks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockHttpHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockHttpHealthChecks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHttpHealthChecks %v not found", key)}
 klog.V(5).Infof("MockHttpHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockHttpHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockHttpHealthChecks.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockHttpHealthChecks.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.HttpHealthCheck
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockHttpHealthChecks.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockHttpHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockHttpHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockHttpHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockHttpHealthChecks %v exists", key)}
  klog.V(5).Infof("MockHttpHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "httpHealthChecks")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "httpHealthChecks", key)
 m.Objects[*key] = &MockHttpHealthChecksObj{obj}
 klog.V(5).Infof("MockHttpHealthChecks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockHttpHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHttpHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockHttpHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHttpHealthChecks %v not found", key)}
  klog.V(5).Infof("MockHttpHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockHttpHealthChecks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockHttpHealthChecks) Obj(o *ga.HttpHealthCheck) *MockHttpHealthChecksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockHttpHealthChecksObj{o}
}
func (m *MockHttpHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEHttpHealthChecks struct{ s *Service }

func (g *GCEHttpHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpHealthChecks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpHealthChecks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "HttpHealthChecks"}
 klog.V(5).Infof("GCEHttpHealthChecks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.HttpHealthChecks.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEHttpHealthChecks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEHttpHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HttpHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpHealthChecks.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "HttpHealthChecks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEHttpHealthChecks.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.HttpHealthChecks.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.HttpHealthCheck
 f := func(l *ga.HttpHealthCheckList) error {
  klog.V(5).Infof("GCEHttpHealthChecks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEHttpHealthChecks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEHttpHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEHttpHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpHealthChecks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpHealthChecks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "HttpHealthChecks"}
 klog.V(5).Infof("GCEHttpHealthChecks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.HttpHealthChecks.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpHealthChecks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEHttpHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpHealthChecks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpHealthChecks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "HttpHealthChecks"}
 klog.V(5).Infof("GCEHttpHealthChecks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HttpHealthChecks.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEHttpHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HttpHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpHealthChecks.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpHealthChecks.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "HttpHealthChecks"}
 klog.V(5).Infof("GCEHttpHealthChecks.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HttpHealthChecks.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type HttpsHealthChecks interface {
 Get(ctx context.Context, key *meta.Key) (*ga.HttpsHealthCheck, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.HttpsHealthCheck, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.HttpsHealthCheck) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *ga.HttpsHealthCheck) error
}

func NewMockHttpsHealthChecks(pr ProjectRouter, objs map[meta.Key]*MockHttpsHealthChecksObj) *MockHttpsHealthChecks {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockHttpsHealthChecks{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockHttpsHealthChecks struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockHttpsHealthChecksObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockHttpsHealthChecks) (bool, *ga.HttpsHealthCheck, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockHttpsHealthChecks) (bool, []*ga.HttpsHealthCheck, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.HttpsHealthCheck, m *MockHttpsHealthChecks) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockHttpsHealthChecks) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.HttpsHealthCheck, *MockHttpsHealthChecks) error
 X             interface{}
}

func (m *MockHttpsHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHttpsHealthChecks.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockHttpsHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockHttpsHealthChecks.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHttpsHealthChecks %v not found", key)}
 klog.V(5).Infof("MockHttpsHealthChecks.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockHttpsHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockHttpsHealthChecks.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockHttpsHealthChecks.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.HttpsHealthCheck
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockHttpsHealthChecks.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockHttpsHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockHttpsHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockHttpsHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockHttpsHealthChecks %v exists", key)}
  klog.V(5).Infof("MockHttpsHealthChecks.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "httpsHealthChecks")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "httpsHealthChecks", key)
 m.Objects[*key] = &MockHttpsHealthChecksObj{obj}
 klog.V(5).Infof("MockHttpsHealthChecks.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockHttpsHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockHttpsHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockHttpsHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockHttpsHealthChecks %v not found", key)}
  klog.V(5).Infof("MockHttpsHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockHttpsHealthChecks.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockHttpsHealthChecks) Obj(o *ga.HttpsHealthCheck) *MockHttpsHealthChecksObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockHttpsHealthChecksObj{o}
}
func (m *MockHttpsHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEHttpsHealthChecks struct{ s *Service }

func (g *GCEHttpsHealthChecks) Get(ctx context.Context, key *meta.Key) (*ga.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpsHealthChecks.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpsHealthChecks.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpsHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "HttpsHealthChecks"}
 klog.V(5).Infof("GCEHttpsHealthChecks.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.HttpsHealthChecks.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEHttpsHealthChecks.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEHttpsHealthChecks) List(ctx context.Context, fl *filter.F) ([]*ga.HttpsHealthCheck, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpsHealthChecks.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpsHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "HttpsHealthChecks"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEHttpsHealthChecks.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.HttpsHealthChecks.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.HttpsHealthCheck
 f := func(l *ga.HttpsHealthCheckList) error {
  klog.V(5).Infof("GCEHttpsHealthChecks.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEHttpsHealthChecks.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEHttpsHealthChecks.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEHttpsHealthChecks) Insert(ctx context.Context, key *meta.Key, obj *ga.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpsHealthChecks.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpsHealthChecks.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpsHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "HttpsHealthChecks"}
 klog.V(5).Infof("GCEHttpsHealthChecks.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.HttpsHealthChecks.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpsHealthChecks.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEHttpsHealthChecks) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpsHealthChecks.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpsHealthChecks.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpsHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "HttpsHealthChecks"}
 klog.V(5).Infof("GCEHttpsHealthChecks.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HttpsHealthChecks.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpsHealthChecks.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEHttpsHealthChecks) Update(ctx context.Context, key *meta.Key, arg0 *ga.HttpsHealthCheck) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "HttpsHealthChecks")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "HttpsHealthChecks"}
 klog.V(5).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.HttpsHealthChecks.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEHttpsHealthChecks.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type InstanceGroups interface {
 Get(ctx context.Context, key *meta.Key) (*ga.InstanceGroup, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*ga.InstanceGroup, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.InstanceGroup) error
 Delete(ctx context.Context, key *meta.Key) error
 AddInstances(context.Context, *meta.Key, *ga.InstanceGroupsAddInstancesRequest) error
 ListInstances(context.Context, *meta.Key, *ga.InstanceGroupsListInstancesRequest, *filter.F) ([]*ga.InstanceWithNamedPorts, error)
 RemoveInstances(context.Context, *meta.Key, *ga.InstanceGroupsRemoveInstancesRequest) error
 SetNamedPorts(context.Context, *meta.Key, *ga.InstanceGroupsSetNamedPortsRequest) error
}

func NewMockInstanceGroups(pr ProjectRouter, objs map[meta.Key]*MockInstanceGroupsObj) *MockInstanceGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockInstanceGroups{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockInstanceGroups struct {
 Lock                sync.Mutex
 ProjectRouter       ProjectRouter
 Objects             map[meta.Key]*MockInstanceGroupsObj
 GetError            map[meta.Key]error
 ListError           *error
 InsertError         map[meta.Key]error
 DeleteError         map[meta.Key]error
 GetHook             func(ctx context.Context, key *meta.Key, m *MockInstanceGroups) (bool, *ga.InstanceGroup, error)
 ListHook            func(ctx context.Context, zone string, fl *filter.F, m *MockInstanceGroups) (bool, []*ga.InstanceGroup, error)
 InsertHook          func(ctx context.Context, key *meta.Key, obj *ga.InstanceGroup, m *MockInstanceGroups) (bool, error)
 DeleteHook          func(ctx context.Context, key *meta.Key, m *MockInstanceGroups) (bool, error)
 AddInstancesHook    func(context.Context, *meta.Key, *ga.InstanceGroupsAddInstancesRequest, *MockInstanceGroups) error
 ListInstancesHook   func(context.Context, *meta.Key, *ga.InstanceGroupsListInstancesRequest, *filter.F, *MockInstanceGroups) ([]*ga.InstanceWithNamedPorts, error)
 RemoveInstancesHook func(context.Context, *meta.Key, *ga.InstanceGroupsRemoveInstancesRequest, *MockInstanceGroups) error
 SetNamedPortsHook   func(context.Context, *meta.Key, *ga.InstanceGroupsSetNamedPortsRequest, *MockInstanceGroups) error
 X                   interface{}
}

func (m *MockInstanceGroups) Get(ctx context.Context, key *meta.Key) (*ga.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockInstanceGroups.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockInstanceGroups.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockInstanceGroups.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockInstanceGroups %v not found", key)}
 klog.V(5).Infof("MockInstanceGroups.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockInstanceGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockInstanceGroups.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockInstanceGroups.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.InstanceGroup
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockInstanceGroups.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockInstanceGroups) Insert(ctx context.Context, key *meta.Key, obj *ga.InstanceGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockInstanceGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockInstanceGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockInstanceGroups %v exists", key)}
  klog.V(5).Infof("MockInstanceGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "instanceGroups")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "instanceGroups", key)
 m.Objects[*key] = &MockInstanceGroupsObj{obj}
 klog.V(5).Infof("MockInstanceGroups.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockInstanceGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockInstanceGroups.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockInstanceGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockInstanceGroups %v not found", key)}
  klog.V(5).Infof("MockInstanceGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockInstanceGroups.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockInstanceGroups) Obj(o *ga.InstanceGroup) *MockInstanceGroupsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockInstanceGroupsObj{o}
}
func (m *MockInstanceGroups) AddInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsAddInstancesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AddInstancesHook != nil {
  return m.AddInstancesHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockInstanceGroups) ListInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsListInstancesRequest, fl *filter.F) ([]*ga.InstanceWithNamedPorts, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListInstancesHook != nil {
  return m.ListInstancesHook(ctx, key, arg0, fl, m)
 }
 return nil, nil
}
func (m *MockInstanceGroups) RemoveInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsRemoveInstancesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.RemoveInstancesHook != nil {
  return m.RemoveInstancesHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockInstanceGroups) SetNamedPorts(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsSetNamedPortsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetNamedPortsHook != nil {
  return m.SetNamedPortsHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEInstanceGroups struct{ s *Service }

func (g *GCEInstanceGroups) Get(ctx context.Context, key *meta.Key) (*ga.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.InstanceGroups.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEInstanceGroups.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEInstanceGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.InstanceGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "InstanceGroups"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEInstanceGroups.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.GA.InstanceGroups.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.InstanceGroup
 f := func(l *ga.InstanceGroupList) error {
  klog.V(5).Infof("GCEInstanceGroups.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEInstanceGroups.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEInstanceGroups.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEInstanceGroups) Insert(ctx context.Context, key *meta.Key, obj *ga.InstanceGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.InstanceGroups.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstanceGroups.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstanceGroups.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEInstanceGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.InstanceGroups.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstanceGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstanceGroups.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEInstanceGroups) AddInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsAddInstancesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AddInstances", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.InstanceGroups.AddInstances(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstanceGroups.AddInstances(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEInstanceGroups) ListInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsListInstancesRequest, fl *filter.F) ([]*ga.InstanceWithNamedPorts, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "ListInstances", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.InstanceGroups.ListInstances(projectID, key.Zone, key.Name, arg0)
 var all []*ga.InstanceWithNamedPorts
 f := func(l *ga.InstanceGroupsListInstances) error {
  klog.V(5).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...): page %+v", ctx, key, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...) = %v, %v", ctx, key, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...) = [%v items], %v", ctx, key, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEInstanceGroups.ListInstances(%v, %v, ...) = %v, %v", ctx, key, asStr, nil)
 }
 return all, nil
}
func (g *GCEInstanceGroups) RemoveInstances(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsRemoveInstancesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "RemoveInstances", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.InstanceGroups.RemoveInstances(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstanceGroups.RemoveInstances(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEInstanceGroups) SetNamedPorts(ctx context.Context, key *meta.Key, arg0 *ga.InstanceGroupsSetNamedPortsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "InstanceGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetNamedPorts", Version: meta.Version("ga"), Service: "InstanceGroups"}
 klog.V(5).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.InstanceGroups.SetNamedPorts(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstanceGroups.SetNamedPorts(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type Instances interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Instance, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Instance, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Instance) error
 Delete(ctx context.Context, key *meta.Key) error
 AttachDisk(context.Context, *meta.Key, *ga.AttachedDisk) error
 DetachDisk(context.Context, *meta.Key, string) error
}

func NewMockInstances(pr ProjectRouter, objs map[meta.Key]*MockInstancesObj) *MockInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockInstances{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockInstances struct {
 Lock           sync.Mutex
 ProjectRouter  ProjectRouter
 Objects        map[meta.Key]*MockInstancesObj
 GetError       map[meta.Key]error
 ListError      *error
 InsertError    map[meta.Key]error
 DeleteError    map[meta.Key]error
 GetHook        func(ctx context.Context, key *meta.Key, m *MockInstances) (bool, *ga.Instance, error)
 ListHook       func(ctx context.Context, zone string, fl *filter.F, m *MockInstances) (bool, []*ga.Instance, error)
 InsertHook     func(ctx context.Context, key *meta.Key, obj *ga.Instance, m *MockInstances) (bool, error)
 DeleteHook     func(ctx context.Context, key *meta.Key, m *MockInstances) (bool, error)
 AttachDiskHook func(context.Context, *meta.Key, *ga.AttachedDisk, *MockInstances) error
 DetachDiskHook func(context.Context, *meta.Key, string, *MockInstances) error
 X              interface{}
}

func (m *MockInstances) Get(ctx context.Context, key *meta.Key) (*ga.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockInstances.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockInstances.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockInstances.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockInstances %v not found", key)}
 klog.V(5).Infof("MockInstances.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockInstances.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockInstances.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Instance
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockInstances.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockInstances) Insert(ctx context.Context, key *meta.Key, obj *ga.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockInstances %v exists", key)}
  klog.V(5).Infof("MockInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "instances")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "instances", key)
 m.Objects[*key] = &MockInstancesObj{obj}
 klog.V(5).Infof("MockInstances.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockInstances.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockInstances %v not found", key)}
  klog.V(5).Infof("MockInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockInstances.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockInstances) Obj(o *ga.Instance) *MockInstancesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockInstancesObj{o}
}
func (m *MockInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *ga.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AttachDiskHook != nil {
  return m.AttachDiskHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DetachDiskHook != nil {
  return m.DetachDiskHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEInstances struct{ s *Service }

func (g *GCEInstances) Get(ctx context.Context, key *meta.Key) (*ga.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstances.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Instances"}
 klog.V(5).Infof("GCEInstances.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstances.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Instances.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEInstances.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*ga.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Instances"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEInstances.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.GA.Instances.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Instance
 f := func(l *ga.InstanceList) error {
  klog.V(5).Infof("GCEInstances.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEInstances.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEInstances.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEInstances.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEInstances) Insert(ctx context.Context, key *meta.Key, obj *ga.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstances.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "Instances"}
 klog.V(5).Infof("GCEInstances.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstances.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.Instances.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstances.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstances.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstances.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "Instances"}
 klog.V(5).Infof("GCEInstances.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstances.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Instances.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstances.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *ga.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.AttachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstances.AttachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AttachDisk", Version: meta.Version("ga"), Service: "Instances"}
 klog.V(5).Infof("GCEInstances.AttachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstances.AttachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Instances.AttachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEInstances.DetachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEInstances.DetachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "DetachDisk", Version: meta.Version("ga"), Service: "Instances"}
 klog.V(5).Infof("GCEInstances.DetachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEInstances.DetachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Instances.DetachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type BetaInstances interface {
 Get(ctx context.Context, key *meta.Key) (*beta.Instance, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*beta.Instance, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.Instance) error
 Delete(ctx context.Context, key *meta.Key) error
 AttachDisk(context.Context, *meta.Key, *beta.AttachedDisk) error
 DetachDisk(context.Context, *meta.Key, string) error
 UpdateNetworkInterface(context.Context, *meta.Key, string, *beta.NetworkInterface) error
}

func NewMockBetaInstances(pr ProjectRouter, objs map[meta.Key]*MockInstancesObj) *MockBetaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaInstances{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaInstances struct {
 Lock                       sync.Mutex
 ProjectRouter              ProjectRouter
 Objects                    map[meta.Key]*MockInstancesObj
 GetError                   map[meta.Key]error
 ListError                  *error
 InsertError                map[meta.Key]error
 DeleteError                map[meta.Key]error
 GetHook                    func(ctx context.Context, key *meta.Key, m *MockBetaInstances) (bool, *beta.Instance, error)
 ListHook                   func(ctx context.Context, zone string, fl *filter.F, m *MockBetaInstances) (bool, []*beta.Instance, error)
 InsertHook                 func(ctx context.Context, key *meta.Key, obj *beta.Instance, m *MockBetaInstances) (bool, error)
 DeleteHook                 func(ctx context.Context, key *meta.Key, m *MockBetaInstances) (bool, error)
 AttachDiskHook             func(context.Context, *meta.Key, *beta.AttachedDisk, *MockBetaInstances) error
 DetachDiskHook             func(context.Context, *meta.Key, string, *MockBetaInstances) error
 UpdateNetworkInterfaceHook func(context.Context, *meta.Key, string, *beta.NetworkInterface, *MockBetaInstances) error
 X                          interface{}
}

func (m *MockBetaInstances) Get(ctx context.Context, key *meta.Key) (*beta.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaInstances.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaInstances.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaInstances.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaInstances %v not found", key)}
 klog.V(5).Infof("MockBetaInstances.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*beta.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockBetaInstances.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaInstances.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.Instance
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaInstances.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockBetaInstances) Insert(ctx context.Context, key *meta.Key, obj *beta.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaInstances %v exists", key)}
  klog.V(5).Infof("MockBetaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "instances")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "instances", key)
 m.Objects[*key] = &MockInstancesObj{obj}
 klog.V(5).Infof("MockBetaInstances.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaInstances.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaInstances %v not found", key)}
  klog.V(5).Infof("MockBetaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaInstances.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaInstances) Obj(o *beta.Instance) *MockInstancesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockInstancesObj{o}
}
func (m *MockBetaInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *beta.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AttachDiskHook != nil {
  return m.AttachDiskHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DetachDiskHook != nil {
  return m.DetachDiskHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaInstances) UpdateNetworkInterface(ctx context.Context, key *meta.Key, arg0 string, arg1 *beta.NetworkInterface) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateNetworkInterfaceHook != nil {
  return m.UpdateNetworkInterfaceHook(ctx, key, arg0, arg1, m)
 }
 return nil
}

type GCEBetaInstances struct{ s *Service }

func (g *GCEBetaInstances) Get(ctx context.Context, key *meta.Key) (*beta.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.Instances.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaInstances.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*beta.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "Instances"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaInstances.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.Beta.Instances.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.Instance
 f := func(l *beta.InstanceList) error {
  klog.V(5).Infof("GCEBetaInstances.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaInstances.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaInstances.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaInstances.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaInstances) Insert(ctx context.Context, key *meta.Key, obj *beta.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.Instances.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaInstances.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaInstances.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.Instances.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaInstances.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBetaInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *beta.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AttachDisk", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.Instances.AttachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "DetachDisk", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.Instances.DetachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaInstances) UpdateNetworkInterface(ctx context.Context, key *meta.Key, arg0 string, arg1 *beta.NetworkInterface) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "UpdateNetworkInterface", Version: meta.Version("beta"), Service: "Instances"}
 klog.V(5).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.Instances.UpdateNetworkInterface(projectID, key.Zone, key.Name, arg0, arg1)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaInstances.UpdateNetworkInterface(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type AlphaInstances interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.Instance, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.Instance, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.Instance) error
 Delete(ctx context.Context, key *meta.Key) error
 AttachDisk(context.Context, *meta.Key, *alpha.AttachedDisk) error
 DetachDisk(context.Context, *meta.Key, string) error
 UpdateNetworkInterface(context.Context, *meta.Key, string, *alpha.NetworkInterface) error
}

func NewMockAlphaInstances(pr ProjectRouter, objs map[meta.Key]*MockInstancesObj) *MockAlphaInstances {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaInstances{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaInstances struct {
 Lock                       sync.Mutex
 ProjectRouter              ProjectRouter
 Objects                    map[meta.Key]*MockInstancesObj
 GetError                   map[meta.Key]error
 ListError                  *error
 InsertError                map[meta.Key]error
 DeleteError                map[meta.Key]error
 GetHook                    func(ctx context.Context, key *meta.Key, m *MockAlphaInstances) (bool, *alpha.Instance, error)
 ListHook                   func(ctx context.Context, zone string, fl *filter.F, m *MockAlphaInstances) (bool, []*alpha.Instance, error)
 InsertHook                 func(ctx context.Context, key *meta.Key, obj *alpha.Instance, m *MockAlphaInstances) (bool, error)
 DeleteHook                 func(ctx context.Context, key *meta.Key, m *MockAlphaInstances) (bool, error)
 AttachDiskHook             func(context.Context, *meta.Key, *alpha.AttachedDisk, *MockAlphaInstances) error
 DetachDiskHook             func(context.Context, *meta.Key, string, *MockAlphaInstances) error
 UpdateNetworkInterfaceHook func(context.Context, *meta.Key, string, *alpha.NetworkInterface, *MockAlphaInstances) error
 X                          interface{}
}

func (m *MockAlphaInstances) Get(ctx context.Context, key *meta.Key) (*alpha.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaInstances.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaInstances.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaInstances.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaInstances %v not found", key)}
 klog.V(5).Infof("MockAlphaInstances.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockAlphaInstances.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaInstances.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.Instance
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaInstances.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaInstances) Insert(ctx context.Context, key *meta.Key, obj *alpha.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaInstances %v exists", key)}
  klog.V(5).Infof("MockAlphaInstances.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "instances")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "instances", key)
 m.Objects[*key] = &MockInstancesObj{obj}
 klog.V(5).Infof("MockAlphaInstances.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaInstances.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaInstances %v not found", key)}
  klog.V(5).Infof("MockAlphaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaInstances.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaInstances) Obj(o *alpha.Instance) *MockInstancesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockInstancesObj{o}
}
func (m *MockAlphaInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *alpha.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AttachDiskHook != nil {
  return m.AttachDiskHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockAlphaInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DetachDiskHook != nil {
  return m.DetachDiskHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockAlphaInstances) UpdateNetworkInterface(ctx context.Context, key *meta.Key, arg0 string, arg1 *alpha.NetworkInterface) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateNetworkInterfaceHook != nil {
  return m.UpdateNetworkInterfaceHook(ctx, key, arg0, arg1, m)
 }
 return nil
}

type GCEAlphaInstances struct{ s *Service }

func (g *GCEAlphaInstances) Get(ctx context.Context, key *meta.Key) (*alpha.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.Instances.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaInstances.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaInstances) List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.Instance, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "Instances"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaInstances.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.Alpha.Instances.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.Instance
 f := func(l *alpha.InstanceList) error {
  klog.V(5).Infof("GCEAlphaInstances.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaInstances.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaInstances.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaInstances) Insert(ctx context.Context, key *meta.Key, obj *alpha.Instance) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.Instances.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaInstances.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaInstances.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaInstances) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.Instances.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaInstances.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaInstances.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEAlphaInstances) AttachDisk(ctx context.Context, key *meta.Key, arg0 *alpha.AttachedDisk) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AttachDisk", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.Instances.AttachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaInstances.AttachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEAlphaInstances) DetachDisk(ctx context.Context, key *meta.Key, arg0 string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "DetachDisk", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.Instances.DetachDisk(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaInstances.DetachDisk(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEAlphaInstances) UpdateNetworkInterface(ctx context.Context, key *meta.Key, arg0 string, arg1 *alpha.NetworkInterface) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "Instances")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "UpdateNetworkInterface", Version: meta.Version("alpha"), Service: "Instances"}
 klog.V(5).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.Instances.UpdateNetworkInterface(projectID, key.Zone, key.Name, arg0, arg1)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaInstances.UpdateNetworkInterface(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type AlphaNetworkEndpointGroups interface {
 Get(ctx context.Context, key *meta.Key) (*alpha.NetworkEndpointGroup, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.NetworkEndpointGroup, error)
 Insert(ctx context.Context, key *meta.Key, obj *alpha.NetworkEndpointGroup) error
 Delete(ctx context.Context, key *meta.Key) error
 AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*alpha.NetworkEndpointGroup, error)
 AttachNetworkEndpoints(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsAttachEndpointsRequest) error
 DetachNetworkEndpoints(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsDetachEndpointsRequest) error
 ListNetworkEndpoints(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsListEndpointsRequest, *filter.F) ([]*alpha.NetworkEndpointWithHealthStatus, error)
}

func NewMockAlphaNetworkEndpointGroups(pr ProjectRouter, objs map[meta.Key]*MockNetworkEndpointGroupsObj) *MockAlphaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockAlphaNetworkEndpointGroups{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockAlphaNetworkEndpointGroups struct {
 Lock                       sync.Mutex
 ProjectRouter              ProjectRouter
 Objects                    map[meta.Key]*MockNetworkEndpointGroupsObj
 GetError                   map[meta.Key]error
 ListError                  *error
 InsertError                map[meta.Key]error
 DeleteError                map[meta.Key]error
 AggregatedListError        *error
 GetHook                    func(ctx context.Context, key *meta.Key, m *MockAlphaNetworkEndpointGroups) (bool, *alpha.NetworkEndpointGroup, error)
 ListHook                   func(ctx context.Context, zone string, fl *filter.F, m *MockAlphaNetworkEndpointGroups) (bool, []*alpha.NetworkEndpointGroup, error)
 InsertHook                 func(ctx context.Context, key *meta.Key, obj *alpha.NetworkEndpointGroup, m *MockAlphaNetworkEndpointGroups) (bool, error)
 DeleteHook                 func(ctx context.Context, key *meta.Key, m *MockAlphaNetworkEndpointGroups) (bool, error)
 AggregatedListHook         func(ctx context.Context, fl *filter.F, m *MockAlphaNetworkEndpointGroups) (bool, map[string][]*alpha.NetworkEndpointGroup, error)
 AttachNetworkEndpointsHook func(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsAttachEndpointsRequest, *MockAlphaNetworkEndpointGroups) error
 DetachNetworkEndpointsHook func(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsDetachEndpointsRequest, *MockAlphaNetworkEndpointGroups) error
 ListNetworkEndpointsHook   func(context.Context, *meta.Key, *alpha.NetworkEndpointGroupsListEndpointsRequest, *filter.F, *MockAlphaNetworkEndpointGroups) ([]*alpha.NetworkEndpointWithHealthStatus, error)
 X                          interface{}
}

func (m *MockAlphaNetworkEndpointGroups) Get(ctx context.Context, key *meta.Key) (*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToAlpha()
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaNetworkEndpointGroups %v not found", key)}
 klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockAlphaNetworkEndpointGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*alpha.NetworkEndpointGroup
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs = append(objs, obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaNetworkEndpointGroups.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaNetworkEndpointGroups) Insert(ctx context.Context, key *meta.Key, obj *alpha.NetworkEndpointGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockAlphaNetworkEndpointGroups %v exists", key)}
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "alpha", "networkEndpointGroups")
 obj.SelfLink = SelfLink(meta.VersionAlpha, projectID, "networkEndpointGroups", key)
 m.Objects[*key] = &MockNetworkEndpointGroupsObj{obj}
 klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockAlphaNetworkEndpointGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockAlphaNetworkEndpointGroups %v not found", key)}
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockAlphaNetworkEndpointGroups.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockAlphaNetworkEndpointGroups) AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AggregatedListHook != nil {
  if intercept, objs, err := m.AggregatedListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.AggregatedListError != nil {
  err := *m.AggregatedListError
  klog.V(5).Infof("MockAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = nil, %v", ctx, fl, err)
  return nil, err
 }
 objs := map[string][]*alpha.NetworkEndpointGroup{}
 for _, obj := range m.Objects {
  res, err := ParseResourceURL(obj.ToAlpha().SelfLink)
  location := res.Key.Zone
  if err != nil {
   klog.V(5).Infof("MockAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = nil, %v", ctx, fl, err)
   return nil, err
  }
  if !fl.Match(obj.ToAlpha()) {
   continue
  }
  objs[location] = append(objs[location], obj.ToAlpha())
 }
 klog.V(5).Infof("MockAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockAlphaNetworkEndpointGroups) Obj(o *alpha.NetworkEndpointGroup) *MockNetworkEndpointGroupsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockNetworkEndpointGroupsObj{o}
}
func (m *MockAlphaNetworkEndpointGroups) AttachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsAttachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AttachNetworkEndpointsHook != nil {
  return m.AttachNetworkEndpointsHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockAlphaNetworkEndpointGroups) DetachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsDetachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DetachNetworkEndpointsHook != nil {
  return m.DetachNetworkEndpointsHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockAlphaNetworkEndpointGroups) ListNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsListEndpointsRequest, fl *filter.F) ([]*alpha.NetworkEndpointWithHealthStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListNetworkEndpointsHook != nil {
  return m.ListNetworkEndpointsHook(ctx, key, arg0, fl, m)
 }
 return nil, nil
}

type GCEAlphaNetworkEndpointGroups struct{ s *Service }

func (g *GCEAlphaNetworkEndpointGroups) Get(ctx context.Context, key *meta.Key) (*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.NetworkEndpointGroups.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEAlphaNetworkEndpointGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.Alpha.NetworkEndpointGroups.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*alpha.NetworkEndpointGroup
 f := func(l *alpha.NetworkEndpointGroupList) error {
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaNetworkEndpointGroups) Insert(ctx context.Context, key *meta.Key, obj *alpha.NetworkEndpointGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Alpha.NetworkEndpointGroups.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEAlphaNetworkEndpointGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.NetworkEndpointGroups.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEAlphaNetworkEndpointGroups) AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*alpha.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AggregatedList", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v): RateLimiter error: %v", ctx, fl, err)
  return nil, err
 }
 call := g.s.Alpha.NetworkEndpointGroups.AggregatedList(projectID)
 call.Context(ctx)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 all := map[string][]*alpha.NetworkEndpointGroup{}
 f := func(l *alpha.NetworkEndpointGroupAggregatedList) error {
  for k, v := range l.Items {
   klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v): page[%v]%+v", ctx, fl, k, v)
   all[k] = append(all[k], v.NetworkEndpointGroups...)
  }
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AggregatedList(%v, %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEAlphaNetworkEndpointGroups) AttachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsAttachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AttachNetworkEndpoints", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.NetworkEndpointGroups.AttachNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEAlphaNetworkEndpointGroups) DetachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsDetachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "DetachNetworkEndpoints", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Alpha.NetworkEndpointGroups.DetachNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEAlphaNetworkEndpointGroups) ListNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *alpha.NetworkEndpointGroupsListEndpointsRequest, fl *filter.F) ([]*alpha.NetworkEndpointWithHealthStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "alpha", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "ListNetworkEndpoints", Version: meta.Version("alpha"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Alpha.NetworkEndpointGroups.ListNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 var all []*alpha.NetworkEndpointWithHealthStatus
 f := func(l *alpha.NetworkEndpointGroupsListNetworkEndpoints) error {
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): page %+v", ctx, key, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = %v, %v", ctx, key, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = [%v items], %v", ctx, key, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEAlphaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = %v, %v", ctx, key, asStr, nil)
 }
 return all, nil
}

type BetaNetworkEndpointGroups interface {
 Get(ctx context.Context, key *meta.Key) (*beta.NetworkEndpointGroup, error)
 List(ctx context.Context, zone string, fl *filter.F) ([]*beta.NetworkEndpointGroup, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.NetworkEndpointGroup) error
 Delete(ctx context.Context, key *meta.Key) error
 AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*beta.NetworkEndpointGroup, error)
 AttachNetworkEndpoints(context.Context, *meta.Key, *beta.NetworkEndpointGroupsAttachEndpointsRequest) error
 DetachNetworkEndpoints(context.Context, *meta.Key, *beta.NetworkEndpointGroupsDetachEndpointsRequest) error
 ListNetworkEndpoints(context.Context, *meta.Key, *beta.NetworkEndpointGroupsListEndpointsRequest, *filter.F) ([]*beta.NetworkEndpointWithHealthStatus, error)
}

func NewMockBetaNetworkEndpointGroups(pr ProjectRouter, objs map[meta.Key]*MockNetworkEndpointGroupsObj) *MockBetaNetworkEndpointGroups {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaNetworkEndpointGroups{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaNetworkEndpointGroups struct {
 Lock                       sync.Mutex
 ProjectRouter              ProjectRouter
 Objects                    map[meta.Key]*MockNetworkEndpointGroupsObj
 GetError                   map[meta.Key]error
 ListError                  *error
 InsertError                map[meta.Key]error
 DeleteError                map[meta.Key]error
 AggregatedListError        *error
 GetHook                    func(ctx context.Context, key *meta.Key, m *MockBetaNetworkEndpointGroups) (bool, *beta.NetworkEndpointGroup, error)
 ListHook                   func(ctx context.Context, zone string, fl *filter.F, m *MockBetaNetworkEndpointGroups) (bool, []*beta.NetworkEndpointGroup, error)
 InsertHook                 func(ctx context.Context, key *meta.Key, obj *beta.NetworkEndpointGroup, m *MockBetaNetworkEndpointGroups) (bool, error)
 DeleteHook                 func(ctx context.Context, key *meta.Key, m *MockBetaNetworkEndpointGroups) (bool, error)
 AggregatedListHook         func(ctx context.Context, fl *filter.F, m *MockBetaNetworkEndpointGroups) (bool, map[string][]*beta.NetworkEndpointGroup, error)
 AttachNetworkEndpointsHook func(context.Context, *meta.Key, *beta.NetworkEndpointGroupsAttachEndpointsRequest, *MockBetaNetworkEndpointGroups) error
 DetachNetworkEndpointsHook func(context.Context, *meta.Key, *beta.NetworkEndpointGroupsDetachEndpointsRequest, *MockBetaNetworkEndpointGroups) error
 ListNetworkEndpointsHook   func(context.Context, *meta.Key, *beta.NetworkEndpointGroupsListEndpointsRequest, *filter.F, *MockBetaNetworkEndpointGroups) ([]*beta.NetworkEndpointWithHealthStatus, error)
 X                          interface{}
}

func (m *MockBetaNetworkEndpointGroups) Get(ctx context.Context, key *meta.Key) (*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaNetworkEndpointGroups %v not found", key)}
 klog.V(5).Infof("MockBetaNetworkEndpointGroups.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaNetworkEndpointGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, zone, fl, m); intercept {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.List(%v, %q, %v) = [%v items], %v", ctx, zone, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.List(%v, %q, %v) = nil, %v", ctx, zone, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.NetworkEndpointGroup
 for key, obj := range m.Objects {
  if key.Zone != zone {
   continue
  }
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaNetworkEndpointGroups.List(%v, %q, %v) = [%v items], nil", ctx, zone, fl, len(objs))
 return objs, nil
}
func (m *MockBetaNetworkEndpointGroups) Insert(ctx context.Context, key *meta.Key, obj *beta.NetworkEndpointGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaNetworkEndpointGroups %v exists", key)}
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "networkEndpointGroups")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "networkEndpointGroups", key)
 m.Objects[*key] = &MockNetworkEndpointGroupsObj{obj}
 klog.V(5).Infof("MockBetaNetworkEndpointGroups.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaNetworkEndpointGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaNetworkEndpointGroups %v not found", key)}
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaNetworkEndpointGroups.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaNetworkEndpointGroups) AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AggregatedListHook != nil {
  if intercept, objs, err := m.AggregatedListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.AggregatedListError != nil {
  err := *m.AggregatedListError
  klog.V(5).Infof("MockBetaNetworkEndpointGroups.AggregatedList(%v, %v) = nil, %v", ctx, fl, err)
  return nil, err
 }
 objs := map[string][]*beta.NetworkEndpointGroup{}
 for _, obj := range m.Objects {
  res, err := ParseResourceURL(obj.ToBeta().SelfLink)
  location := res.Key.Zone
  if err != nil {
   klog.V(5).Infof("MockBetaNetworkEndpointGroups.AggregatedList(%v, %v) = nil, %v", ctx, fl, err)
   return nil, err
  }
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs[location] = append(objs[location], obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockBetaNetworkEndpointGroups) Obj(o *beta.NetworkEndpointGroup) *MockNetworkEndpointGroupsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockNetworkEndpointGroupsObj{o}
}
func (m *MockBetaNetworkEndpointGroups) AttachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsAttachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AttachNetworkEndpointsHook != nil {
  return m.AttachNetworkEndpointsHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaNetworkEndpointGroups) DetachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsDetachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DetachNetworkEndpointsHook != nil {
  return m.DetachNetworkEndpointsHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaNetworkEndpointGroups) ListNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsListEndpointsRequest, fl *filter.F) ([]*beta.NetworkEndpointWithHealthStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListNetworkEndpointsHook != nil {
  return m.ListNetworkEndpointsHook(ctx, key, arg0, fl, m)
 }
 return nil, nil
}

type GCEBetaNetworkEndpointGroups struct{ s *Service }

func (g *GCEBetaNetworkEndpointGroups) Get(ctx context.Context, key *meta.Key) (*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.NetworkEndpointGroups.Get(projectID, key.Zone, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaNetworkEndpointGroups) List(ctx context.Context, zone string, fl *filter.F) ([]*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.List(%v, %v, %v) called", ctx, zone, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, zone, fl, projectID, rk)
 call := g.s.Beta.NetworkEndpointGroups.List(projectID, zone)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.NetworkEndpointGroup
 f := func(l *beta.NetworkEndpointGroupList) error {
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaNetworkEndpointGroups) Insert(ctx context.Context, key *meta.Key, obj *beta.NetworkEndpointGroup) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.NetworkEndpointGroups.Insert(projectID, key.Zone, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaNetworkEndpointGroups) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.NetworkEndpointGroups.Delete(projectID, key.Zone, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaNetworkEndpointGroups.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBetaNetworkEndpointGroups) AggregatedList(ctx context.Context, fl *filter.F) (map[string][]*beta.NetworkEndpointGroup, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AggregatedList", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v): RateLimiter error: %v", ctx, fl, err)
  return nil, err
 }
 call := g.s.Beta.NetworkEndpointGroups.AggregatedList(projectID)
 call.Context(ctx)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 all := map[string][]*beta.NetworkEndpointGroup{}
 f := func(l *beta.NetworkEndpointGroupAggregatedList) error {
  for k, v := range l.Items {
   klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v): page[%v]%+v", ctx, fl, k, v)
   all[k] = append(all[k], v.NetworkEndpointGroups...)
  }
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AggregatedList(%v, %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaNetworkEndpointGroups) AttachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsAttachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AttachNetworkEndpoints", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.NetworkEndpointGroups.AttachNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaNetworkEndpointGroups.AttachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaNetworkEndpointGroups) DetachNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsDetachEndpointsRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "DetachNetworkEndpoints", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.NetworkEndpointGroups.DetachNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaNetworkEndpointGroups.DetachNetworkEndpoints(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaNetworkEndpointGroups) ListNetworkEndpoints(ctx context.Context, key *meta.Key, arg0 *beta.NetworkEndpointGroupsListEndpointsRequest, fl *filter.F) ([]*beta.NetworkEndpointWithHealthStatus, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "NetworkEndpointGroups")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "ListNetworkEndpoints", Version: meta.Version("beta"), Service: "NetworkEndpointGroups"}
 klog.V(5).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.NetworkEndpointGroups.ListNetworkEndpoints(projectID, key.Zone, key.Name, arg0)
 var all []*beta.NetworkEndpointWithHealthStatus
 f := func(l *beta.NetworkEndpointGroupsListNetworkEndpoints) error {
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...): page %+v", ctx, key, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = %v, %v", ctx, key, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = [%v items], %v", ctx, key, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaNetworkEndpointGroups.ListNetworkEndpoints(%v, %v, ...) = %v, %v", ctx, key, asStr, nil)
 }
 return all, nil
}

type Projects interface{ ProjectsOps }

func NewMockProjects(pr ProjectRouter, objs map[meta.Key]*MockProjectsObj) *MockProjects {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockProjects{ProjectRouter: pr, Objects: objs}
 return mock
}

type MockProjects struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockProjectsObj
 X             interface{}
}

func (m *MockProjects) Obj(o *ga.Project) *MockProjectsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockProjectsObj{o}
}

type GCEProjects struct{ s *Service }
type Regions interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Region, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.Region, error)
}

func NewMockRegions(pr ProjectRouter, objs map[meta.Key]*MockRegionsObj) *MockRegions {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockRegions{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}}
 return mock
}

type MockRegions struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockRegionsObj
 GetError      map[meta.Key]error
 ListError     *error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockRegions) (bool, *ga.Region, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockRegions) (bool, []*ga.Region, error)
 X             interface{}
}

func (m *MockRegions) Get(ctx context.Context, key *meta.Key) (*ga.Region, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRegions.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockRegions.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockRegions.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRegions %v not found", key)}
 klog.V(5).Infof("MockRegions.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockRegions) List(ctx context.Context, fl *filter.F) ([]*ga.Region, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockRegions.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockRegions.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Region
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockRegions.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockRegions) Obj(o *ga.Region) *MockRegionsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockRegionsObj{o}
}

type GCERegions struct{ s *Service }

func (g *GCERegions) Get(ctx context.Context, key *meta.Key) (*ga.Region, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegions.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERegions.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Regions")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Regions"}
 klog.V(5).Infof("GCERegions.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERegions.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Regions.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCERegions.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCERegions) List(ctx context.Context, fl *filter.F) ([]*ga.Region, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERegions.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Regions")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Regions"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCERegions.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.Regions.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Region
 f := func(l *ga.RegionList) error {
  klog.V(5).Infof("GCERegions.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCERegions.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCERegions.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCERegions.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}

type Routes interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Route, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.Route, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.Route) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockRoutes(pr ProjectRouter, objs map[meta.Key]*MockRoutesObj) *MockRoutes {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockRoutes{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockRoutes struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockRoutesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockRoutes) (bool, *ga.Route, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockRoutes) (bool, []*ga.Route, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.Route, m *MockRoutes) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockRoutes) (bool, error)
 X             interface{}
}

func (m *MockRoutes) Get(ctx context.Context, key *meta.Key) (*ga.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRoutes.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockRoutes.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockRoutes.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRoutes %v not found", key)}
 klog.V(5).Infof("MockRoutes.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockRoutes) List(ctx context.Context, fl *filter.F) ([]*ga.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockRoutes.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockRoutes.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Route
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockRoutes.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockRoutes) Insert(ctx context.Context, key *meta.Key, obj *ga.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockRoutes.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockRoutes.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockRoutes %v exists", key)}
  klog.V(5).Infof("MockRoutes.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "routes")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "routes", key)
 m.Objects[*key] = &MockRoutesObj{obj}
 klog.V(5).Infof("MockRoutes.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockRoutes) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockRoutes.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockRoutes.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockRoutes %v not found", key)}
  klog.V(5).Infof("MockRoutes.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockRoutes.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockRoutes) Obj(o *ga.Route) *MockRoutesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockRoutesObj{o}
}

type GCERoutes struct{ s *Service }

func (g *GCERoutes) Get(ctx context.Context, key *meta.Key) (*ga.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERoutes.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERoutes.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Routes")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Routes"}
 klog.V(5).Infof("GCERoutes.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERoutes.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Routes.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCERoutes.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCERoutes) List(ctx context.Context, fl *filter.F) ([]*ga.Route, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERoutes.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Routes")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Routes"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCERoutes.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.Routes.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Route
 f := func(l *ga.RouteList) error {
  klog.V(5).Infof("GCERoutes.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCERoutes.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCERoutes.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCERoutes.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCERoutes) Insert(ctx context.Context, key *meta.Key, obj *ga.Route) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERoutes.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCERoutes.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Routes")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "Routes"}
 klog.V(5).Infof("GCERoutes.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERoutes.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.Routes.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERoutes.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERoutes.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCERoutes) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCERoutes.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCERoutes.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Routes")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "Routes"}
 klog.V(5).Infof("GCERoutes.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCERoutes.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.Routes.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCERoutes.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCERoutes.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type BetaSecurityPolicies interface {
 Get(ctx context.Context, key *meta.Key) (*beta.SecurityPolicy, error)
 List(ctx context.Context, fl *filter.F) ([]*beta.SecurityPolicy, error)
 Insert(ctx context.Context, key *meta.Key, obj *beta.SecurityPolicy) error
 Delete(ctx context.Context, key *meta.Key) error
 AddRule(context.Context, *meta.Key, *beta.SecurityPolicyRule) error
 GetRule(context.Context, *meta.Key) (*beta.SecurityPolicyRule, error)
 Patch(context.Context, *meta.Key, *beta.SecurityPolicy) error
 PatchRule(context.Context, *meta.Key, *beta.SecurityPolicyRule) error
 RemoveRule(context.Context, *meta.Key) error
}

func NewMockBetaSecurityPolicies(pr ProjectRouter, objs map[meta.Key]*MockSecurityPoliciesObj) *MockBetaSecurityPolicies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockBetaSecurityPolicies{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockBetaSecurityPolicies struct {
 Lock           sync.Mutex
 ProjectRouter  ProjectRouter
 Objects        map[meta.Key]*MockSecurityPoliciesObj
 GetError       map[meta.Key]error
 ListError      *error
 InsertError    map[meta.Key]error
 DeleteError    map[meta.Key]error
 GetHook        func(ctx context.Context, key *meta.Key, m *MockBetaSecurityPolicies) (bool, *beta.SecurityPolicy, error)
 ListHook       func(ctx context.Context, fl *filter.F, m *MockBetaSecurityPolicies) (bool, []*beta.SecurityPolicy, error)
 InsertHook     func(ctx context.Context, key *meta.Key, obj *beta.SecurityPolicy, m *MockBetaSecurityPolicies) (bool, error)
 DeleteHook     func(ctx context.Context, key *meta.Key, m *MockBetaSecurityPolicies) (bool, error)
 AddRuleHook    func(context.Context, *meta.Key, *beta.SecurityPolicyRule, *MockBetaSecurityPolicies) error
 GetRuleHook    func(context.Context, *meta.Key, *MockBetaSecurityPolicies) (*beta.SecurityPolicyRule, error)
 PatchHook      func(context.Context, *meta.Key, *beta.SecurityPolicy, *MockBetaSecurityPolicies) error
 PatchRuleHook  func(context.Context, *meta.Key, *beta.SecurityPolicyRule, *MockBetaSecurityPolicies) error
 RemoveRuleHook func(context.Context, *meta.Key, *MockBetaSecurityPolicies) error
 X              interface{}
}

func (m *MockBetaSecurityPolicies) Get(ctx context.Context, key *meta.Key) (*beta.SecurityPolicy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaSecurityPolicies.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockBetaSecurityPolicies.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToBeta()
  klog.V(5).Infof("MockBetaSecurityPolicies.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaSecurityPolicies %v not found", key)}
 klog.V(5).Infof("MockBetaSecurityPolicies.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockBetaSecurityPolicies) List(ctx context.Context, fl *filter.F) ([]*beta.SecurityPolicy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockBetaSecurityPolicies.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockBetaSecurityPolicies.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*beta.SecurityPolicy
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToBeta()) {
   continue
  }
  objs = append(objs, obj.ToBeta())
 }
 klog.V(5).Infof("MockBetaSecurityPolicies.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockBetaSecurityPolicies) Insert(ctx context.Context, key *meta.Key, obj *beta.SecurityPolicy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockBetaSecurityPolicies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockBetaSecurityPolicies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockBetaSecurityPolicies %v exists", key)}
  klog.V(5).Infof("MockBetaSecurityPolicies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "beta", "securityPolicies")
 obj.SelfLink = SelfLink(meta.VersionBeta, projectID, "securityPolicies", key)
 m.Objects[*key] = &MockSecurityPoliciesObj{obj}
 klog.V(5).Infof("MockBetaSecurityPolicies.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockBetaSecurityPolicies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockBetaSecurityPolicies.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockBetaSecurityPolicies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockBetaSecurityPolicies %v not found", key)}
  klog.V(5).Infof("MockBetaSecurityPolicies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockBetaSecurityPolicies.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockBetaSecurityPolicies) Obj(o *beta.SecurityPolicy) *MockSecurityPoliciesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockSecurityPoliciesObj{o}
}
func (m *MockBetaSecurityPolicies) AddRule(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AddRuleHook != nil {
  return m.AddRuleHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaSecurityPolicies) GetRule(ctx context.Context, key *meta.Key) (*beta.SecurityPolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetRuleHook != nil {
  return m.GetRuleHook(ctx, key, m)
 }
 return nil, fmt.Errorf("GetRuleHook must be set")
}
func (m *MockBetaSecurityPolicies) Patch(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.PatchHook != nil {
  return m.PatchHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaSecurityPolicies) PatchRule(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.PatchRuleHook != nil {
  return m.PatchRuleHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockBetaSecurityPolicies) RemoveRule(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.RemoveRuleHook != nil {
  return m.RemoveRuleHook(ctx, key, m)
 }
 return nil
}

type GCEBetaSecurityPolicies struct{ s *Service }

func (g *GCEBetaSecurityPolicies) Get(ctx context.Context, key *meta.Key) (*beta.SecurityPolicy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.SecurityPolicies.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaSecurityPolicies.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaSecurityPolicies) List(ctx context.Context, fl *filter.F) ([]*beta.SecurityPolicy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEBetaSecurityPolicies.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.Beta.SecurityPolicies.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*beta.SecurityPolicy
 f := func(l *beta.SecurityPolicyList) error {
  klog.V(5).Infof("GCEBetaSecurityPolicies.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEBetaSecurityPolicies.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEBetaSecurityPolicies.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEBetaSecurityPolicies) Insert(ctx context.Context, key *meta.Key, obj *beta.SecurityPolicy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.Beta.SecurityPolicies.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEBetaSecurityPolicies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.SecurityPolicies.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEBetaSecurityPolicies) AddRule(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AddRule", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.SecurityPolicies.AddRule(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.AddRule(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaSecurityPolicies) GetRule(ctx context.Context, key *meta.Key) (*beta.SecurityPolicyRule, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.GetRule(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.GetRule(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "GetRule", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.GetRule(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.GetRule(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.Beta.SecurityPolicies.GetRule(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEBetaSecurityPolicies.GetRule(%v, %v, ...) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEBetaSecurityPolicies) Patch(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Patch", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.SecurityPolicies.Patch(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.Patch(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaSecurityPolicies) PatchRule(ctx context.Context, key *meta.Key, arg0 *beta.SecurityPolicyRule) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "PatchRule", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.SecurityPolicies.PatchRule(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.PatchRule(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCEBetaSecurityPolicies) RemoveRule(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "beta", "SecurityPolicies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "RemoveRule", Version: meta.Version("beta"), Service: "SecurityPolicies"}
 klog.V(5).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.Beta.SecurityPolicies.RemoveRule(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEBetaSecurityPolicies.RemoveRule(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type SslCertificates interface {
 Get(ctx context.Context, key *meta.Key) (*ga.SslCertificate, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.SslCertificate, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.SslCertificate) error
 Delete(ctx context.Context, key *meta.Key) error
}

func NewMockSslCertificates(pr ProjectRouter, objs map[meta.Key]*MockSslCertificatesObj) *MockSslCertificates {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockSslCertificates{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockSslCertificates struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockSslCertificatesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockSslCertificates) (bool, *ga.SslCertificate, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockSslCertificates) (bool, []*ga.SslCertificate, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.SslCertificate, m *MockSslCertificates) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockSslCertificates) (bool, error)
 X             interface{}
}

func (m *MockSslCertificates) Get(ctx context.Context, key *meta.Key) (*ga.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockSslCertificates.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockSslCertificates.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockSslCertificates.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockSslCertificates %v not found", key)}
 klog.V(5).Infof("MockSslCertificates.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockSslCertificates) List(ctx context.Context, fl *filter.F) ([]*ga.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockSslCertificates.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockSslCertificates.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.SslCertificate
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockSslCertificates.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockSslCertificates) Insert(ctx context.Context, key *meta.Key, obj *ga.SslCertificate) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockSslCertificates.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockSslCertificates.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockSslCertificates %v exists", key)}
  klog.V(5).Infof("MockSslCertificates.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "sslCertificates")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "sslCertificates", key)
 m.Objects[*key] = &MockSslCertificatesObj{obj}
 klog.V(5).Infof("MockSslCertificates.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockSslCertificates) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockSslCertificates.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockSslCertificates.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockSslCertificates %v not found", key)}
  klog.V(5).Infof("MockSslCertificates.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockSslCertificates.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockSslCertificates) Obj(o *ga.SslCertificate) *MockSslCertificatesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockSslCertificatesObj{o}
}

type GCESslCertificates struct{ s *Service }

func (g *GCESslCertificates) Get(ctx context.Context, key *meta.Key) (*ga.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCESslCertificates.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCESslCertificates.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "SslCertificates")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "SslCertificates"}
 klog.V(5).Infof("GCESslCertificates.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCESslCertificates.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.SslCertificates.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCESslCertificates.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCESslCertificates) List(ctx context.Context, fl *filter.F) ([]*ga.SslCertificate, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCESslCertificates.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "SslCertificates")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "SslCertificates"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCESslCertificates.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.SslCertificates.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.SslCertificate
 f := func(l *ga.SslCertificateList) error {
  klog.V(5).Infof("GCESslCertificates.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCESslCertificates.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCESslCertificates.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCESslCertificates.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCESslCertificates) Insert(ctx context.Context, key *meta.Key, obj *ga.SslCertificate) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCESslCertificates.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCESslCertificates.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "SslCertificates")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "SslCertificates"}
 klog.V(5).Infof("GCESslCertificates.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCESslCertificates.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.SslCertificates.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCESslCertificates.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCESslCertificates.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCESslCertificates) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCESslCertificates.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCESslCertificates.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "SslCertificates")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "SslCertificates"}
 klog.V(5).Infof("GCESslCertificates.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCESslCertificates.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.SslCertificates.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCESslCertificates.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCESslCertificates.Delete(%v, %v) = %v", ctx, key, err)
 return err
}

type TargetHttpProxies interface {
 Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpProxy, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpProxy, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpProxy) error
 Delete(ctx context.Context, key *meta.Key) error
 SetUrlMap(context.Context, *meta.Key, *ga.UrlMapReference) error
}

func NewMockTargetHttpProxies(pr ProjectRouter, objs map[meta.Key]*MockTargetHttpProxiesObj) *MockTargetHttpProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockTargetHttpProxies{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockTargetHttpProxies struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockTargetHttpProxiesObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockTargetHttpProxies) (bool, *ga.TargetHttpProxy, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockTargetHttpProxies) (bool, []*ga.TargetHttpProxy, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.TargetHttpProxy, m *MockTargetHttpProxies) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockTargetHttpProxies) (bool, error)
 SetUrlMapHook func(context.Context, *meta.Key, *ga.UrlMapReference, *MockTargetHttpProxies) error
 X             interface{}
}

func (m *MockTargetHttpProxies) Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetHttpProxies.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpProxies.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockTargetHttpProxies.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetHttpProxies %v not found", key)}
 klog.V(5).Infof("MockTargetHttpProxies.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockTargetHttpProxies) List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockTargetHttpProxies.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockTargetHttpProxies.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.TargetHttpProxy
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockTargetHttpProxies.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockTargetHttpProxies) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockTargetHttpProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockTargetHttpProxies %v exists", key)}
  klog.V(5).Infof("MockTargetHttpProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "targetHttpProxies")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "targetHttpProxies", key)
 m.Objects[*key] = &MockTargetHttpProxiesObj{obj}
 klog.V(5).Infof("MockTargetHttpProxies.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockTargetHttpProxies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetHttpProxies.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetHttpProxies %v not found", key)}
  klog.V(5).Infof("MockTargetHttpProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockTargetHttpProxies.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockTargetHttpProxies) Obj(o *ga.TargetHttpProxy) *MockTargetHttpProxiesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockTargetHttpProxiesObj{o}
}
func (m *MockTargetHttpProxies) SetUrlMap(ctx context.Context, key *meta.Key, arg0 *ga.UrlMapReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetUrlMapHook != nil {
  return m.SetUrlMapHook(ctx, key, arg0, m)
 }
 return nil
}

type GCETargetHttpProxies struct{ s *Service }

func (g *GCETargetHttpProxies) Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpProxies.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpProxies.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "TargetHttpProxies"}
 klog.V(5).Infof("GCETargetHttpProxies.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.TargetHttpProxies.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCETargetHttpProxies.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCETargetHttpProxies) List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpProxies.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "TargetHttpProxies"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCETargetHttpProxies.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.TargetHttpProxies.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.TargetHttpProxy
 f := func(l *ga.TargetHttpProxyList) error {
  klog.V(5).Infof("GCETargetHttpProxies.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCETargetHttpProxies.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCETargetHttpProxies.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCETargetHttpProxies) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpProxies.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpProxies.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "TargetHttpProxies"}
 klog.V(5).Infof("GCETargetHttpProxies.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.TargetHttpProxies.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpProxies.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCETargetHttpProxies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpProxies.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpProxies.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "TargetHttpProxies"}
 klog.V(5).Infof("GCETargetHttpProxies.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetHttpProxies.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpProxies.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCETargetHttpProxies) SetUrlMap(ctx context.Context, key *meta.Key, arg0 *ga.UrlMapReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetUrlMap", Version: meta.Version("ga"), Service: "TargetHttpProxies"}
 klog.V(5).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetHttpProxies.SetUrlMap(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpProxies.SetUrlMap(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type TargetHttpsProxies interface {
 Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpsProxy, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpsProxy, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpsProxy) error
 Delete(ctx context.Context, key *meta.Key) error
 SetSslCertificates(context.Context, *meta.Key, *ga.TargetHttpsProxiesSetSslCertificatesRequest) error
 SetUrlMap(context.Context, *meta.Key, *ga.UrlMapReference) error
}

func NewMockTargetHttpsProxies(pr ProjectRouter, objs map[meta.Key]*MockTargetHttpsProxiesObj) *MockTargetHttpsProxies {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockTargetHttpsProxies{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockTargetHttpsProxies struct {
 Lock                   sync.Mutex
 ProjectRouter          ProjectRouter
 Objects                map[meta.Key]*MockTargetHttpsProxiesObj
 GetError               map[meta.Key]error
 ListError              *error
 InsertError            map[meta.Key]error
 DeleteError            map[meta.Key]error
 GetHook                func(ctx context.Context, key *meta.Key, m *MockTargetHttpsProxies) (bool, *ga.TargetHttpsProxy, error)
 ListHook               func(ctx context.Context, fl *filter.F, m *MockTargetHttpsProxies) (bool, []*ga.TargetHttpsProxy, error)
 InsertHook             func(ctx context.Context, key *meta.Key, obj *ga.TargetHttpsProxy, m *MockTargetHttpsProxies) (bool, error)
 DeleteHook             func(ctx context.Context, key *meta.Key, m *MockTargetHttpsProxies) (bool, error)
 SetSslCertificatesHook func(context.Context, *meta.Key, *ga.TargetHttpsProxiesSetSslCertificatesRequest, *MockTargetHttpsProxies) error
 SetUrlMapHook          func(context.Context, *meta.Key, *ga.UrlMapReference, *MockTargetHttpsProxies) error
 X                      interface{}
}

func (m *MockTargetHttpsProxies) Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetHttpsProxies.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpsProxies.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockTargetHttpsProxies.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetHttpsProxies %v not found", key)}
 klog.V(5).Infof("MockTargetHttpsProxies.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockTargetHttpsProxies) List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockTargetHttpsProxies.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockTargetHttpsProxies.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.TargetHttpsProxy
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockTargetHttpsProxies.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockTargetHttpsProxies) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpsProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockTargetHttpsProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpsProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockTargetHttpsProxies %v exists", key)}
  klog.V(5).Infof("MockTargetHttpsProxies.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "targetHttpsProxies")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "targetHttpsProxies", key)
 m.Objects[*key] = &MockTargetHttpsProxiesObj{obj}
 klog.V(5).Infof("MockTargetHttpsProxies.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockTargetHttpsProxies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetHttpsProxies.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockTargetHttpsProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetHttpsProxies %v not found", key)}
  klog.V(5).Infof("MockTargetHttpsProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockTargetHttpsProxies.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockTargetHttpsProxies) Obj(o *ga.TargetHttpsProxy) *MockTargetHttpsProxiesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockTargetHttpsProxiesObj{o}
}
func (m *MockTargetHttpsProxies) SetSslCertificates(ctx context.Context, key *meta.Key, arg0 *ga.TargetHttpsProxiesSetSslCertificatesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetSslCertificatesHook != nil {
  return m.SetSslCertificatesHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockTargetHttpsProxies) SetUrlMap(ctx context.Context, key *meta.Key, arg0 *ga.UrlMapReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.SetUrlMapHook != nil {
  return m.SetUrlMapHook(ctx, key, arg0, m)
 }
 return nil
}

type GCETargetHttpsProxies struct{ s *Service }

func (g *GCETargetHttpsProxies) Get(ctx context.Context, key *meta.Key) (*ga.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpsProxies.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 klog.V(5).Infof("GCETargetHttpsProxies.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.TargetHttpsProxies.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCETargetHttpsProxies.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCETargetHttpsProxies) List(ctx context.Context, fl *filter.F) ([]*ga.TargetHttpsProxy, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCETargetHttpsProxies.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.TargetHttpsProxies.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.TargetHttpsProxy
 f := func(l *ga.TargetHttpsProxyList) error {
  klog.V(5).Infof("GCETargetHttpsProxies.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCETargetHttpsProxies.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCETargetHttpsProxies.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCETargetHttpsProxies) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetHttpsProxy) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpsProxies.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 klog.V(5).Infof("GCETargetHttpsProxies.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.TargetHttpsProxies.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpsProxies.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCETargetHttpsProxies) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpsProxies.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 klog.V(5).Infof("GCETargetHttpsProxies.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetHttpsProxies.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpsProxies.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCETargetHttpsProxies) SetSslCertificates(ctx context.Context, key *meta.Key, arg0 *ga.TargetHttpsProxiesSetSslCertificatesRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetSslCertificates", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 klog.V(5).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetHttpsProxies.SetSslCertificates(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpsProxies.SetSslCertificates(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCETargetHttpsProxies) SetUrlMap(ctx context.Context, key *meta.Key, arg0 *ga.UrlMapReference) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetHttpsProxies")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "SetUrlMap", Version: meta.Version("ga"), Service: "TargetHttpsProxies"}
 klog.V(5).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetHttpsProxies.SetUrlMap(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetHttpsProxies.SetUrlMap(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type TargetPools interface {
 Get(ctx context.Context, key *meta.Key) (*ga.TargetPool, error)
 List(ctx context.Context, region string, fl *filter.F) ([]*ga.TargetPool, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.TargetPool) error
 Delete(ctx context.Context, key *meta.Key) error
 AddInstance(context.Context, *meta.Key, *ga.TargetPoolsAddInstanceRequest) error
 RemoveInstance(context.Context, *meta.Key, *ga.TargetPoolsRemoveInstanceRequest) error
}

func NewMockTargetPools(pr ProjectRouter, objs map[meta.Key]*MockTargetPoolsObj) *MockTargetPools {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockTargetPools{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockTargetPools struct {
 Lock               sync.Mutex
 ProjectRouter      ProjectRouter
 Objects            map[meta.Key]*MockTargetPoolsObj
 GetError           map[meta.Key]error
 ListError          *error
 InsertError        map[meta.Key]error
 DeleteError        map[meta.Key]error
 GetHook            func(ctx context.Context, key *meta.Key, m *MockTargetPools) (bool, *ga.TargetPool, error)
 ListHook           func(ctx context.Context, region string, fl *filter.F, m *MockTargetPools) (bool, []*ga.TargetPool, error)
 InsertHook         func(ctx context.Context, key *meta.Key, obj *ga.TargetPool, m *MockTargetPools) (bool, error)
 DeleteHook         func(ctx context.Context, key *meta.Key, m *MockTargetPools) (bool, error)
 AddInstanceHook    func(context.Context, *meta.Key, *ga.TargetPoolsAddInstanceRequest, *MockTargetPools) error
 RemoveInstanceHook func(context.Context, *meta.Key, *ga.TargetPoolsRemoveInstanceRequest, *MockTargetPools) error
 X                  interface{}
}

func (m *MockTargetPools) Get(ctx context.Context, key *meta.Key) (*ga.TargetPool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetPools.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockTargetPools.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockTargetPools.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetPools %v not found", key)}
 klog.V(5).Infof("MockTargetPools.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockTargetPools) List(ctx context.Context, region string, fl *filter.F) ([]*ga.TargetPool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, region, fl, m); intercept {
   klog.V(5).Infof("MockTargetPools.List(%v, %q, %v) = [%v items], %v", ctx, region, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockTargetPools.List(%v, %q, %v) = nil, %v", ctx, region, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.TargetPool
 for key, obj := range m.Objects {
  if key.Region != region {
   continue
  }
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockTargetPools.List(%v, %q, %v) = [%v items], nil", ctx, region, fl, len(objs))
 return objs, nil
}
func (m *MockTargetPools) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetPool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockTargetPools.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockTargetPools.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockTargetPools %v exists", key)}
  klog.V(5).Infof("MockTargetPools.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "targetPools")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "targetPools", key)
 m.Objects[*key] = &MockTargetPoolsObj{obj}
 klog.V(5).Infof("MockTargetPools.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockTargetPools) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockTargetPools.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockTargetPools.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockTargetPools %v not found", key)}
  klog.V(5).Infof("MockTargetPools.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockTargetPools.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockTargetPools) Obj(o *ga.TargetPool) *MockTargetPoolsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockTargetPoolsObj{o}
}
func (m *MockTargetPools) AddInstance(ctx context.Context, key *meta.Key, arg0 *ga.TargetPoolsAddInstanceRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.AddInstanceHook != nil {
  return m.AddInstanceHook(ctx, key, arg0, m)
 }
 return nil
}
func (m *MockTargetPools) RemoveInstance(ctx context.Context, key *meta.Key, arg0 *ga.TargetPoolsRemoveInstanceRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.RemoveInstanceHook != nil {
  return m.RemoveInstanceHook(ctx, key, arg0, m)
 }
 return nil
}

type GCETargetPools struct{ s *Service }

func (g *GCETargetPools) Get(ctx context.Context, key *meta.Key) (*ga.TargetPool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetPools.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "TargetPools"}
 klog.V(5).Infof("GCETargetPools.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetPools.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.TargetPools.Get(projectID, key.Region, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCETargetPools.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCETargetPools) List(ctx context.Context, region string, fl *filter.F) ([]*ga.TargetPool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.List(%v, %v, %v) called", ctx, region, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "TargetPools"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCETargetPools.List(%v, %v, %v): projectID = %v, rk = %+v", ctx, region, fl, projectID, rk)
 call := g.s.GA.TargetPools.List(projectID, region)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.TargetPool
 f := func(l *ga.TargetPoolList) error {
  klog.V(5).Infof("GCETargetPools.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCETargetPools.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCETargetPools.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCETargetPools.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCETargetPools) Insert(ctx context.Context, key *meta.Key, obj *ga.TargetPool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetPools.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "TargetPools"}
 klog.V(5).Infof("GCETargetPools.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetPools.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.TargetPools.Insert(projectID, key.Region, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetPools.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetPools.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCETargetPools) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetPools.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "TargetPools"}
 klog.V(5).Infof("GCETargetPools.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetPools.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetPools.Delete(projectID, key.Region, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetPools.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetPools.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCETargetPools) AddInstance(ctx context.Context, key *meta.Key, arg0 *ga.TargetPoolsAddInstanceRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.AddInstance(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetPools.AddInstance(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "AddInstance", Version: meta.Version("ga"), Service: "TargetPools"}
 klog.V(5).Infof("GCETargetPools.AddInstance(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetPools.AddInstance(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetPools.AddInstance(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetPools.AddInstance(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetPools.AddInstance(%v, %v, ...) = %+v", ctx, key, err)
 return err
}
func (g *GCETargetPools) RemoveInstance(ctx context.Context, key *meta.Key, arg0 *ga.TargetPoolsRemoveInstanceRequest) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCETargetPools.RemoveInstance(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCETargetPools.RemoveInstance(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "TargetPools")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "RemoveInstance", Version: meta.Version("ga"), Service: "TargetPools"}
 klog.V(5).Infof("GCETargetPools.RemoveInstance(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCETargetPools.RemoveInstance(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.TargetPools.RemoveInstance(projectID, key.Region, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCETargetPools.RemoveInstance(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCETargetPools.RemoveInstance(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type UrlMaps interface {
 Get(ctx context.Context, key *meta.Key) (*ga.UrlMap, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.UrlMap, error)
 Insert(ctx context.Context, key *meta.Key, obj *ga.UrlMap) error
 Delete(ctx context.Context, key *meta.Key) error
 Update(context.Context, *meta.Key, *ga.UrlMap) error
}

func NewMockUrlMaps(pr ProjectRouter, objs map[meta.Key]*MockUrlMapsObj) *MockUrlMaps {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockUrlMaps{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}, InsertError: map[meta.Key]error{}, DeleteError: map[meta.Key]error{}}
 return mock
}

type MockUrlMaps struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockUrlMapsObj
 GetError      map[meta.Key]error
 ListError     *error
 InsertError   map[meta.Key]error
 DeleteError   map[meta.Key]error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockUrlMaps) (bool, *ga.UrlMap, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockUrlMaps) (bool, []*ga.UrlMap, error)
 InsertHook    func(ctx context.Context, key *meta.Key, obj *ga.UrlMap, m *MockUrlMaps) (bool, error)
 DeleteHook    func(ctx context.Context, key *meta.Key, m *MockUrlMaps) (bool, error)
 UpdateHook    func(context.Context, *meta.Key, *ga.UrlMap, *MockUrlMaps) error
 X             interface{}
}

func (m *MockUrlMaps) Get(ctx context.Context, key *meta.Key) (*ga.UrlMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockUrlMaps.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockUrlMaps.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockUrlMaps.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockUrlMaps %v not found", key)}
 klog.V(5).Infof("MockUrlMaps.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockUrlMaps) List(ctx context.Context, fl *filter.F) ([]*ga.UrlMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockUrlMaps.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockUrlMaps.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.UrlMap
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockUrlMaps.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockUrlMaps) Insert(ctx context.Context, key *meta.Key, obj *ga.UrlMap) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.InsertHook != nil {
  if intercept, err := m.InsertHook(ctx, key, obj, m); intercept {
   klog.V(5).Infof("MockUrlMaps.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.InsertError[*key]; ok {
  klog.V(5).Infof("MockUrlMaps.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 if _, ok := m.Objects[*key]; ok {
  err := &googleapi.Error{Code: http.StatusConflict, Message: fmt.Sprintf("MockUrlMaps %v exists", key)}
  klog.V(5).Infof("MockUrlMaps.Insert(%v, %v, %+v) = %v", ctx, key, obj, err)
  return err
 }
 obj.Name = key.Name
 projectID := m.ProjectRouter.ProjectID(ctx, "ga", "urlMaps")
 obj.SelfLink = SelfLink(meta.VersionGA, projectID, "urlMaps", key)
 m.Objects[*key] = &MockUrlMapsObj{obj}
 klog.V(5).Infof("MockUrlMaps.Insert(%v, %v, %+v) = nil", ctx, key, obj)
 return nil
}
func (m *MockUrlMaps) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.DeleteHook != nil {
  if intercept, err := m.DeleteHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockUrlMaps.Delete(%v, %v) = %v", ctx, key, err)
   return err
  }
 }
 if !key.Valid() {
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.DeleteError[*key]; ok {
  klog.V(5).Infof("MockUrlMaps.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 if _, ok := m.Objects[*key]; !ok {
  err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockUrlMaps %v not found", key)}
  klog.V(5).Infof("MockUrlMaps.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 delete(m.Objects, *key)
 klog.V(5).Infof("MockUrlMaps.Delete(%v, %v) = nil", ctx, key)
 return nil
}
func (m *MockUrlMaps) Obj(o *ga.UrlMap) *MockUrlMapsObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockUrlMapsObj{o}
}
func (m *MockUrlMaps) Update(ctx context.Context, key *meta.Key, arg0 *ga.UrlMap) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.UpdateHook != nil {
  return m.UpdateHook(ctx, key, arg0, m)
 }
 return nil
}

type GCEUrlMaps struct{ s *Service }

func (g *GCEUrlMaps) Get(ctx context.Context, key *meta.Key) (*ga.UrlMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEUrlMaps.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEUrlMaps.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "UrlMaps")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "UrlMaps"}
 klog.V(5).Infof("GCEUrlMaps.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEUrlMaps.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.UrlMaps.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEUrlMaps.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEUrlMaps) List(ctx context.Context, fl *filter.F) ([]*ga.UrlMap, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEUrlMaps.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "UrlMaps")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "UrlMaps"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEUrlMaps.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.UrlMaps.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.UrlMap
 f := func(l *ga.UrlMapList) error {
  klog.V(5).Infof("GCEUrlMaps.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEUrlMaps.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEUrlMaps.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEUrlMaps.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func (g *GCEUrlMaps) Insert(ctx context.Context, key *meta.Key, obj *ga.UrlMap) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEUrlMaps.Insert(%v, %v, %+v): called", ctx, key, obj)
 if !key.Valid() {
  klog.V(2).Infof("GCEUrlMaps.Insert(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "UrlMaps")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Insert", Version: meta.Version("ga"), Service: "UrlMaps"}
 klog.V(5).Infof("GCEUrlMaps.Insert(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEUrlMaps.Insert(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 obj.Name = key.Name
 call := g.s.GA.UrlMaps.Insert(projectID, obj)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEUrlMaps.Insert(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEUrlMaps.Insert(%v, %v, %+v) = %+v", ctx, key, obj, err)
 return err
}
func (g *GCEUrlMaps) Delete(ctx context.Context, key *meta.Key) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEUrlMaps.Delete(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEUrlMaps.Delete(%v, %v): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "UrlMaps")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Delete", Version: meta.Version("ga"), Service: "UrlMaps"}
 klog.V(5).Infof("GCEUrlMaps.Delete(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEUrlMaps.Delete(%v, %v): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.UrlMaps.Delete(projectID, key.Name)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEUrlMaps.Delete(%v, %v) = %v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEUrlMaps.Delete(%v, %v) = %v", ctx, key, err)
 return err
}
func (g *GCEUrlMaps) Update(ctx context.Context, key *meta.Key, arg0 *ga.UrlMap) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEUrlMaps.Update(%v, %v, ...): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEUrlMaps.Update(%v, %v, ...): key is invalid (%#v)", ctx, key, key)
  return fmt.Errorf("invalid GCE key (%+v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "UrlMaps")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Update", Version: meta.Version("ga"), Service: "UrlMaps"}
 klog.V(5).Infof("GCEUrlMaps.Update(%v, %v, ...): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEUrlMaps.Update(%v, %v, ...): RateLimiter error: %v", ctx, key, err)
  return err
 }
 call := g.s.GA.UrlMaps.Update(projectID, key.Name, arg0)
 call.Context(ctx)
 op, err := call.Do()
 if err != nil {
  klog.V(4).Infof("GCEUrlMaps.Update(%v, %v, ...) = %+v", ctx, key, err)
  return err
 }
 err = g.s.WaitForCompletion(ctx, op)
 klog.V(4).Infof("GCEUrlMaps.Update(%v, %v, ...) = %+v", ctx, key, err)
 return err
}

type Zones interface {
 Get(ctx context.Context, key *meta.Key) (*ga.Zone, error)
 List(ctx context.Context, fl *filter.F) ([]*ga.Zone, error)
}

func NewMockZones(pr ProjectRouter, objs map[meta.Key]*MockZonesObj) *MockZones {
 _logClusterCodePath()
 defer _logClusterCodePath()
 mock := &MockZones{ProjectRouter: pr, Objects: objs, GetError: map[meta.Key]error{}}
 return mock
}

type MockZones struct {
 Lock          sync.Mutex
 ProjectRouter ProjectRouter
 Objects       map[meta.Key]*MockZonesObj
 GetError      map[meta.Key]error
 ListError     *error
 GetHook       func(ctx context.Context, key *meta.Key, m *MockZones) (bool, *ga.Zone, error)
 ListHook      func(ctx context.Context, fl *filter.F, m *MockZones) (bool, []*ga.Zone, error)
 X             interface{}
}

func (m *MockZones) Get(ctx context.Context, key *meta.Key) (*ga.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.GetHook != nil {
  if intercept, obj, err := m.GetHook(ctx, key, m); intercept {
   klog.V(5).Infof("MockZones.Get(%v, %s) = %+v, %v", ctx, key, obj, err)
   return obj, err
  }
 }
 if !key.Valid() {
  return nil, fmt.Errorf("invalid GCE key (%+v)", key)
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if err, ok := m.GetError[*key]; ok {
  klog.V(5).Infof("MockZones.Get(%v, %s) = nil, %v", ctx, key, err)
  return nil, err
 }
 if obj, ok := m.Objects[*key]; ok {
  typedObj := obj.ToGA()
  klog.V(5).Infof("MockZones.Get(%v, %s) = %+v, nil", ctx, key, typedObj)
  return typedObj, nil
 }
 err := &googleapi.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("MockZones %v not found", key)}
 klog.V(5).Infof("MockZones.Get(%v, %s) = nil, %v", ctx, key, err)
 return nil, err
}
func (m *MockZones) List(ctx context.Context, fl *filter.F) ([]*ga.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if m.ListHook != nil {
  if intercept, objs, err := m.ListHook(ctx, fl, m); intercept {
   klog.V(5).Infof("MockZones.List(%v, %v) = [%v items], %v", ctx, fl, len(objs), err)
   return objs, err
  }
 }
 m.Lock.Lock()
 defer m.Lock.Unlock()
 if m.ListError != nil {
  err := *m.ListError
  klog.V(5).Infof("MockZones.List(%v, %v) = nil, %v", ctx, fl, err)
  return nil, *m.ListError
 }
 var objs []*ga.Zone
 for _, obj := range m.Objects {
  if !fl.Match(obj.ToGA()) {
   continue
  }
  objs = append(objs, obj.ToGA())
 }
 klog.V(5).Infof("MockZones.List(%v, %v) = [%v items], nil", ctx, fl, len(objs))
 return objs, nil
}
func (m *MockZones) Obj(o *ga.Zone) *MockZonesObj {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &MockZonesObj{o}
}

type GCEZones struct{ s *Service }

func (g *GCEZones) Get(ctx context.Context, key *meta.Key) (*ga.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEZones.Get(%v, %v): called", ctx, key)
 if !key.Valid() {
  klog.V(2).Infof("GCEZones.Get(%v, %v): key is invalid (%#v)", ctx, key, key)
  return nil, fmt.Errorf("invalid GCE key (%#v)", key)
 }
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Zones")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "Get", Version: meta.Version("ga"), Service: "Zones"}
 klog.V(5).Infof("GCEZones.Get(%v, %v): projectID = %v, rk = %+v", ctx, key, projectID, rk)
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  klog.V(4).Infof("GCEZones.Get(%v, %v): RateLimiter error: %v", ctx, key, err)
  return nil, err
 }
 call := g.s.GA.Zones.Get(projectID, key.Name)
 call.Context(ctx)
 v, err := call.Do()
 klog.V(4).Infof("GCEZones.Get(%v, %v) = %+v, %v", ctx, key, v, err)
 return v, err
}
func (g *GCEZones) List(ctx context.Context, fl *filter.F) ([]*ga.Zone, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 klog.V(5).Infof("GCEZones.List(%v, %v) called", ctx, fl)
 projectID := g.s.ProjectRouter.ProjectID(ctx, "ga", "Zones")
 rk := &RateLimitKey{ProjectID: projectID, Operation: "List", Version: meta.Version("ga"), Service: "Zones"}
 if err := g.s.RateLimiter.Accept(ctx, rk); err != nil {
  return nil, err
 }
 klog.V(5).Infof("GCEZones.List(%v, %v): projectID = %v, rk = %+v", ctx, fl, projectID, rk)
 call := g.s.GA.Zones.List(projectID)
 if fl != filter.None {
  call.Filter(fl.String())
 }
 var all []*ga.Zone
 f := func(l *ga.ZoneList) error {
  klog.V(5).Infof("GCEZones.List(%v, ..., %v): page %+v", ctx, fl, l)
  all = append(all, l.Items...)
  return nil
 }
 if err := call.Pages(ctx, f); err != nil {
  klog.V(4).Infof("GCEZones.List(%v, ..., %v) = %v, %v", ctx, fl, nil, err)
  return nil, err
 }
 if klog.V(4) {
  klog.V(4).Infof("GCEZones.List(%v, ..., %v) = [%v items], %v", ctx, fl, len(all), nil)
 } else if klog.V(5) {
  var asStr []string
  for _, o := range all {
   asStr = append(asStr, fmt.Sprintf("%+v", o))
  }
  klog.V(5).Infof("GCEZones.List(%v, ..., %v) = %v, %v", ctx, fl, asStr, nil)
 }
 return all, nil
}
func NewAddressesResourceID(project, region, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.RegionalKey(name, region)
 return &ResourceID{project, "addresses", key}
}
func NewBackendServicesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "backendServices", key}
}
func NewDisksResourceID(project, zone, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.ZonalKey(name, zone)
 return &ResourceID{project, "disks", key}
}
func NewFirewallsResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "firewalls", key}
}
func NewForwardingRulesResourceID(project, region, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.RegionalKey(name, region)
 return &ResourceID{project, "forwardingRules", key}
}
func NewGlobalAddressesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "addresses", key}
}
func NewGlobalForwardingRulesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "forwardingRules", key}
}
func NewHealthChecksResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "healthChecks", key}
}
func NewHttpHealthChecksResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "httpHealthChecks", key}
}
func NewHttpsHealthChecksResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "httpsHealthChecks", key}
}
func NewInstanceGroupsResourceID(project, zone, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.ZonalKey(name, zone)
 return &ResourceID{project, "instanceGroups", key}
}
func NewInstancesResourceID(project, zone, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.ZonalKey(name, zone)
 return &ResourceID{project, "instances", key}
}
func NewNetworkEndpointGroupsResourceID(project, zone, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.ZonalKey(name, zone)
 return &ResourceID{project, "networkEndpointGroups", key}
}
func NewProjectsResourceID(project string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var key *meta.Key
 return &ResourceID{project, "projects", key}
}
func NewRegionBackendServicesResourceID(project, region, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.RegionalKey(name, region)
 return &ResourceID{project, "backendServices", key}
}
func NewRegionDisksResourceID(project, region, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.RegionalKey(name, region)
 return &ResourceID{project, "disks", key}
}
func NewRegionsResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "regions", key}
}
func NewRoutesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "routes", key}
}
func NewSecurityPoliciesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "securityPolicies", key}
}
func NewSslCertificatesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "sslCertificates", key}
}
func NewTargetHttpProxiesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "targetHttpProxies", key}
}
func NewTargetHttpsProxiesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "targetHttpsProxies", key}
}
func NewTargetPoolsResourceID(project, region, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.RegionalKey(name, region)
 return &ResourceID{project, "targetPools", key}
}
func NewUrlMapsResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "urlMaps", key}
}
func NewZonesResourceID(project, name string) *ResourceID {
 _logClusterCodePath()
 defer _logClusterCodePath()
 key := meta.GlobalKey(name)
 return &ResourceID{project, "zones", key}
}
