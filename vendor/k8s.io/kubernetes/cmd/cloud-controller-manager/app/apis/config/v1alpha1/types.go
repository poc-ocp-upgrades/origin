package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfigv1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
)

type CloudControllerManagerConfiguration struct {
	metav1.TypeMeta           `json:",inline"`
	Generic                   kubectrlmgrconfigv1alpha1.GenericControllerManagerConfiguration
	KubeCloudShared           kubectrlmgrconfigv1alpha1.KubeCloudSharedConfiguration
	ServiceController         kubectrlmgrconfigv1alpha1.ServiceControllerConfiguration
	NodeStatusUpdateFrequency metav1.Duration
}
