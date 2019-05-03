package gce

import (
 "fmt"
)

const (
 AlphaFeatureNetworkTiers = "NetworkTiers"
)

type AlphaFeatureGate struct{ features map[string]bool }

func (af *AlphaFeatureGate) Enabled(key string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return af.features[key]
}
func NewAlphaFeatureGate(features []string) *AlphaFeatureGate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 featureMap := make(map[string]bool)
 for _, name := range features {
  featureMap[name] = true
 }
 return &AlphaFeatureGate{featureMap}
}
func (g *Cloud) alphaFeatureEnabled(feature string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if !g.AlphaFeatureGate.Enabled(feature) {
  return fmt.Errorf("alpha feature %q is not enabled", feature)
 }
 return nil
}
