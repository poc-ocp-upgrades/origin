package gce

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
)

type LoadBalancerType string

const (
	ServiceAnnotationLoadBalancerType                           = "cloud.google.com/load-balancer-type"
	LBTypeInternal                             LoadBalancerType = "Internal"
	deprecatedTypeInternalLowerCase            LoadBalancerType = "internal"
	ServiceAnnotationILBBackendShare                            = "alpha.cloud.google.com/load-balancer-backend-share"
	deprecatedServiceAnnotationILBBackendShare                  = "cloud.google.com/load-balancer-backend-share"
	NetworkTierAnnotationKey                                    = "cloud.google.com/network-tier"
	NetworkTierAnnotationStandard                               = cloud.NetworkTierStandard
	NetworkTierAnnotationPremium                                = cloud.NetworkTierPremium
)

func GetLoadBalancerAnnotationType(service *v1.Service) (LoadBalancerType, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v := LoadBalancerType("")
	if service.Spec.Type != v1.ServiceTypeLoadBalancer {
		return v, false
	}
	l, ok := service.Annotations[ServiceAnnotationLoadBalancerType]
	v = LoadBalancerType(l)
	if !ok {
		return v, false
	}
	switch v {
	case LBTypeInternal, deprecatedTypeInternalLowerCase:
		return LBTypeInternal, true
	default:
		return v, false
	}
}
func GetLoadBalancerAnnotationBackendShare(service *v1.Service) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if l, exists := service.Annotations[ServiceAnnotationILBBackendShare]; exists && l == "true" {
		return true
	}
	if l, exists := service.Annotations[deprecatedServiceAnnotationILBBackendShare]; exists && l == "true" {
		klog.Warningf("Annotation %q is deprecated and replaced with an alpha-specific key: %q", deprecatedServiceAnnotationILBBackendShare, ServiceAnnotationILBBackendShare)
		return true
	}
	return false
}
func GetServiceNetworkTier(service *v1.Service) (cloud.NetworkTier, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l, ok := service.Annotations[NetworkTierAnnotationKey]
	if !ok {
		return cloud.NetworkTierDefault, nil
	}
	v := cloud.NetworkTier(l)
	switch v {
	case cloud.NetworkTierStandard:
		fallthrough
	case cloud.NetworkTierPremium:
		return v, nil
	default:
		return cloud.NetworkTierDefault, fmt.Errorf("unsupported network tier: %q", v)
	}
}
