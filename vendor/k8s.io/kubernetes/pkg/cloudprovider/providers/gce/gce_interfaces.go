package gce

import (
 computealpha "google.golang.org/api/compute/v0.alpha"
 computebeta "google.golang.org/api/compute/v0.beta"
 compute "google.golang.org/api/compute/v1"
)

type CloudAddressService interface {
 ReserveRegionAddress(address *compute.Address, region string) error
 GetRegionAddress(name string, region string) (*compute.Address, error)
 GetRegionAddressByIP(region, ipAddress string) (*compute.Address, error)
 DeleteRegionAddress(name, region string) error
 GetAlphaRegionAddress(name, region string) (*computealpha.Address, error)
 ReserveAlphaRegionAddress(addr *computealpha.Address, region string) error
 ReserveBetaRegionAddress(address *computebeta.Address, region string) error
 GetBetaRegionAddress(name string, region string) (*computebeta.Address, error)
 GetBetaRegionAddressByIP(region, ipAddress string) (*computebeta.Address, error)
 getNetworkTierFromAddress(name, region string) (string, error)
}
type CloudForwardingRuleService interface {
 GetRegionForwardingRule(name, region string) (*compute.ForwardingRule, error)
 CreateRegionForwardingRule(rule *compute.ForwardingRule, region string) error
 DeleteRegionForwardingRule(name, region string) error
 GetAlphaRegionForwardingRule(name, region string) (*computealpha.ForwardingRule, error)
 CreateAlphaRegionForwardingRule(rule *computealpha.ForwardingRule, region string) error
 getNetworkTierFromForwardingRule(name, region string) (string, error)
}
