package settings

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type PodPreset struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec PodPresetSpec
}
type PodPresetSpec struct {
	Selector     metav1.LabelSelector
	Env          []api.EnvVar
	EnvFrom      []api.EnvFromSource
	Volumes      []api.Volume
	VolumeMounts []api.VolumeMount
}
type PodPresetList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PodPreset
}
