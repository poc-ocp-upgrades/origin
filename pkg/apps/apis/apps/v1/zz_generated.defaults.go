package v1

import (
	v1 "github.com/openshift/api/apps/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.DeploymentConfig{}, func(obj interface{}) {
		SetObjectDefaults_DeploymentConfig(obj.(*v1.DeploymentConfig))
	})
	scheme.AddTypeDefaultingFunc(&v1.DeploymentConfigList{}, func(obj interface{}) {
		SetObjectDefaults_DeploymentConfigList(obj.(*v1.DeploymentConfigList))
	})
	return nil
}
func SetObjectDefaults_DeploymentConfig(in *v1.DeploymentConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_DeploymentConfig(in)
	SetDefaults_DeploymentConfigSpec(&in.Spec)
	SetDefaults_DeploymentStrategy(&in.Spec.Strategy)
	if in.Spec.Strategy.CustomParams != nil {
		for i := range in.Spec.Strategy.CustomParams.Environment {
			a := &in.Spec.Strategy.CustomParams.Environment[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.Strategy.RecreateParams != nil {
		SetDefaults_RecreateDeploymentStrategyParams(in.Spec.Strategy.RecreateParams)
		if in.Spec.Strategy.RecreateParams.Pre != nil {
			if in.Spec.Strategy.RecreateParams.Pre.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Pre.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Pre.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RecreateParams.Mid != nil {
			if in.Spec.Strategy.RecreateParams.Mid.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Mid.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Mid.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RecreateParams.Post != nil {
			if in.Spec.Strategy.RecreateParams.Post.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RecreateParams.Post.ExecNewPod.Env {
					a := &in.Spec.Strategy.RecreateParams.Post.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
	}
	if in.Spec.Strategy.RollingParams != nil {
		SetDefaults_RollingDeploymentStrategyParams(in.Spec.Strategy.RollingParams)
		if in.Spec.Strategy.RollingParams.Pre != nil {
			if in.Spec.Strategy.RollingParams.Pre.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RollingParams.Pre.ExecNewPod.Env {
					a := &in.Spec.Strategy.RollingParams.Pre.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
		if in.Spec.Strategy.RollingParams.Post != nil {
			if in.Spec.Strategy.RollingParams.Post.ExecNewPod != nil {
				for i := range in.Spec.Strategy.RollingParams.Post.ExecNewPod.Env {
					a := &in.Spec.Strategy.RollingParams.Post.ExecNewPod.Env[i]
					if a.ValueFrom != nil {
						if a.ValueFrom.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
						}
					}
				}
			}
		}
	}
	corev1.SetDefaults_ResourceList(&in.Spec.Strategy.Resources.Limits)
	corev1.SetDefaults_ResourceList(&in.Spec.Strategy.Resources.Requests)
	if in.Spec.Template != nil {
		corev1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
		for i := range in.Spec.Template.Spec.Volumes {
			a := &in.Spec.Template.Spec.Volumes[i]
			corev1.SetDefaults_Volume(a)
			if a.VolumeSource.HostPath != nil {
				corev1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
			}
			if a.VolumeSource.Secret != nil {
				corev1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
			}
			if a.VolumeSource.ISCSI != nil {
				corev1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
			}
			if a.VolumeSource.RBD != nil {
				corev1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
			}
			if a.VolumeSource.DownwardAPI != nil {
				corev1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
				for j := range a.VolumeSource.DownwardAPI.Items {
					b := &a.VolumeSource.DownwardAPI.Items[j]
					if b.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(b.FieldRef)
					}
				}
			}
			if a.VolumeSource.ConfigMap != nil {
				corev1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
			}
			if a.VolumeSource.AzureDisk != nil {
				corev1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
			}
			if a.VolumeSource.Projected != nil {
				corev1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
				for j := range a.VolumeSource.Projected.Sources {
					b := &a.VolumeSource.Projected.Sources[j]
					if b.DownwardAPI != nil {
						for k := range b.DownwardAPI.Items {
							c := &b.DownwardAPI.Items[k]
							if c.FieldRef != nil {
								corev1.SetDefaults_ObjectFieldSelector(c.FieldRef)
							}
						}
					}
					if b.ServiceAccountToken != nil {
						corev1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
					}
				}
			}
			if a.VolumeSource.ScaleIO != nil {
				corev1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
			}
		}
		for i := range in.Spec.Template.Spec.InitContainers {
			a := &in.Spec.Template.Spec.InitContainers[i]
			corev1.SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				corev1.SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			corev1.SetDefaults_ResourceList(&a.Resources.Limits)
			corev1.SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				corev1.SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				corev1.SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.Containers {
			a := &in.Spec.Template.Spec.Containers[i]
			corev1.SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				corev1.SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			corev1.SetDefaults_ResourceList(&a.Resources.Limits)
			corev1.SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				corev1.SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				corev1.SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
	}
}
func SetObjectDefaults_DeploymentConfigList(in *v1.DeploymentConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_DeploymentConfig(a)
	}
}
