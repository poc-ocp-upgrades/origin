package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type CloudControllerManagerConfiguration struct {
	metav1.TypeMeta
	Generic                   kubectrlmgrconfig.GenericControllerManagerConfiguration
	KubeCloudShared           kubectrlmgrconfig.KubeCloudSharedConfiguration
	ServiceController         kubectrlmgrconfig.ServiceControllerConfiguration
	NodeStatusUpdateFrequency metav1.Duration
}
