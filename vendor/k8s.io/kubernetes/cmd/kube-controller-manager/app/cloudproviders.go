package app

import (
	"fmt"
	"k8s.io/client-go/informers"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
)

func createCloudProvider(cloudProvider string, externalCloudVolumePlugin string, cloudConfigFile string, allowUntaggedCloud bool, sharedInformers informers.SharedInformerFactory) (cloudprovider.Interface, ControllerLoopMode, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cloud cloudprovider.Interface
	var loopMode ControllerLoopMode
	var err error
	if cloudprovider.IsExternal(cloudProvider) {
		loopMode = ExternalLoops
		if externalCloudVolumePlugin == "" {
			return nil, loopMode, nil
		}
		cloud, err = cloudprovider.InitCloudProvider(externalCloudVolumePlugin, cloudConfigFile)
	} else {
		loopMode = IncludeCloudLoops
		cloud, err = cloudprovider.InitCloudProvider(cloudProvider, cloudConfigFile)
	}
	if err != nil {
		return nil, loopMode, fmt.Errorf("cloud provider could not be initialized: %v", err)
	}
	if cloud != nil && cloud.HasClusterID() == false {
		if allowUntaggedCloud == true {
			klog.Warning("detected a cluster without a ClusterID.  A ClusterID will be required in the future.  Please tag your cluster to avoid any future issues")
		} else {
			return nil, loopMode, fmt.Errorf("no ClusterID Found.  A ClusterID is required for the cloud provider to function properly.  This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}
	if informerUserCloud, ok := cloud.(cloudprovider.InformerUser); ok {
		informerUserCloud.SetInformers(sharedInformers)
	}
	return cloud, loopMode, err
}
