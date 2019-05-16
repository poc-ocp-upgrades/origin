package util

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/apis/storage"
	"k8s.io/kubernetes/pkg/features"
)

func DropDisabledAlphaFields(class *storage.StorageClass) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
		class.VolumeBindingMode = nil
		class.AllowedTopologies = nil
	}
}
