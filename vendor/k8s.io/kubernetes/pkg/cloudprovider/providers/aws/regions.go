package aws

import (
 "sync"
 "k8s.io/klog"
 "k8s.io/apimachinery/pkg/util/sets"
 awscredentialprovider "k8s.io/kubernetes/pkg/credentialprovider/aws"
)

var wellKnownRegions = [...]string{"ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "cn-north-1", "cn-northwest-1", "us-gov-west-1"}
var awsRegionsMutex sync.Mutex
var awsRegions sets.String

func recognizeRegion(region string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 awsRegionsMutex.Lock()
 defer awsRegionsMutex.Unlock()
 if awsRegions == nil {
  awsRegions = sets.NewString()
 }
 if awsRegions.Has(region) {
  klog.V(6).Infof("found AWS region %q again - ignoring", region)
  return
 }
 klog.V(4).Infof("found AWS region %q", region)
 awscredentialprovider.RegisterCredentialsProvider(region)
 awsRegions.Insert(region)
}
func recognizeWellKnownRegions() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, region := range wellKnownRegions {
  recognizeRegion(region)
 }
}
func isRegionValid(region string) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 awsRegionsMutex.Lock()
 defer awsRegionsMutex.Unlock()
 return awsRegions.Has(region)
}
