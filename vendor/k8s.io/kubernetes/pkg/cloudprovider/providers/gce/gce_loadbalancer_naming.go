package gce

import (
 "crypto/sha1"
 "encoding/hex"
 "fmt"
 "strings"
 "k8s.io/api/core/v1"
 "k8s.io/apimachinery/pkg/types"
 "k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
)

func makeInstanceGroupName(clusterID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 prefix := "k8s-ig"
 if clusterID == "" {
  return prefix
 }
 return fmt.Sprintf("%s--%s", prefix, clusterID)
}
func makeBackendServiceName(loadBalancerName, clusterID string, shared bool, scheme cloud.LbScheme, protocol v1.Protocol, svcAffinity v1.ServiceAffinity) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if shared {
  hash := sha1.New()
  hash.Write([]byte(string(svcAffinity)))
  hashed := hex.EncodeToString(hash.Sum(nil))
  hashed = hashed[:16]
  return fmt.Sprintf("k8s-%s-%s-%s-nmv1-%s", clusterID, strings.ToLower(string(scheme)), strings.ToLower(string(protocol)), hashed)
 }
 return loadBalancerName
}
func makeHealthCheckName(loadBalancerName, clusterID string, shared bool) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if shared {
  return fmt.Sprintf("k8s-%s-node", clusterID)
 }
 return loadBalancerName
}
func makeHealthCheckFirewallNameFromHC(healthCheckName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return healthCheckName + "-hc"
}
func makeHealthCheckFirewallName(loadBalancerName, clusterID string, shared bool) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if shared {
  return fmt.Sprintf("k8s-%s-node-hc", clusterID)
 }
 return loadBalancerName + "-hc"
}
func makeBackendServiceDescription(nm types.NamespacedName, shared bool) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if shared {
  return ""
 }
 return fmt.Sprintf(`{"kubernetes.io/service-name":"%s"}`, nm.String())
}
func makeServiceDescription(serviceName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf(`{"kubernetes.io/service-name":"%s"}`, serviceName)
}
func MakeNodesHealthCheckName(clusterID string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("k8s-%v-node", clusterID)
}
func makeHealthCheckDescription(serviceName string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf(`{"kubernetes.io/service-name":"%s"}`, serviceName)
}
func MakeHealthCheckFirewallName(clusterID, hcName string, isNodesHealthCheck bool) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if isNodesHealthCheck {
  return MakeNodesHealthCheckName(clusterID) + "-http-hc"
 }
 return "k8s-" + hcName + "-http-hc"
}
func MakeFirewallName(name string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf("k8s-fw-%s", name)
}
func makeFirewallDescription(serviceName, ipAddress string) string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fmt.Sprintf(`{"kubernetes.io/service-name":"%s", "kubernetes.io/service-ip":"%s"}`, serviceName, ipAddress)
}
