package v1

import (
	v1 "github.com/openshift/api/security/v1"
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
	scheme.AddTypeDefaultingFunc(&v1.PodSecurityPolicyReview{}, func(obj interface{}) {
		SetObjectDefaults_PodSecurityPolicyReview(obj.(*v1.PodSecurityPolicyReview))
	})
	scheme.AddTypeDefaultingFunc(&v1.PodSecurityPolicySelfSubjectReview{}, func(obj interface{}) {
		SetObjectDefaults_PodSecurityPolicySelfSubjectReview(obj.(*v1.PodSecurityPolicySelfSubjectReview))
	})
	scheme.AddTypeDefaultingFunc(&v1.PodSecurityPolicySubjectReview{}, func(obj interface{}) {
		SetObjectDefaults_PodSecurityPolicySubjectReview(obj.(*v1.PodSecurityPolicySubjectReview))
	})
	return nil
}
func SetObjectDefaults_PodSecurityPolicyReview(in *v1.PodSecurityPolicyReview) {
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
	for i := range in.Status.AllowedServiceAccounts {
		a := &in.Status.AllowedServiceAccounts[i]
		corev1.SetDefaults_PodSpec(&a.PodSecurityPolicySubjectReviewStatus.Template.Spec)
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Volumes {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Volumes[j]
			corev1.SetDefaults_Volume(b)
			if b.VolumeSource.HostPath != nil {
				corev1.SetDefaults_HostPathVolumeSource(b.VolumeSource.HostPath)
			}
			if b.VolumeSource.Secret != nil {
				corev1.SetDefaults_SecretVolumeSource(b.VolumeSource.Secret)
			}
			if b.VolumeSource.ISCSI != nil {
				corev1.SetDefaults_ISCSIVolumeSource(b.VolumeSource.ISCSI)
			}
			if b.VolumeSource.RBD != nil {
				corev1.SetDefaults_RBDVolumeSource(b.VolumeSource.RBD)
			}
			if b.VolumeSource.DownwardAPI != nil {
				corev1.SetDefaults_DownwardAPIVolumeSource(b.VolumeSource.DownwardAPI)
				for k := range b.VolumeSource.DownwardAPI.Items {
					c := &b.VolumeSource.DownwardAPI.Items[k]
					if c.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(c.FieldRef)
					}
				}
			}
			if b.VolumeSource.ConfigMap != nil {
				corev1.SetDefaults_ConfigMapVolumeSource(b.VolumeSource.ConfigMap)
			}
			if b.VolumeSource.AzureDisk != nil {
				corev1.SetDefaults_AzureDiskVolumeSource(b.VolumeSource.AzureDisk)
			}
			if b.VolumeSource.Projected != nil {
				corev1.SetDefaults_ProjectedVolumeSource(b.VolumeSource.Projected)
				for k := range b.VolumeSource.Projected.Sources {
					c := &b.VolumeSource.Projected.Sources[k]
					if c.DownwardAPI != nil {
						for l := range c.DownwardAPI.Items {
							d := &c.DownwardAPI.Items[l]
							if d.FieldRef != nil {
								corev1.SetDefaults_ObjectFieldSelector(d.FieldRef)
							}
						}
					}
					if c.ServiceAccountToken != nil {
						corev1.SetDefaults_ServiceAccountTokenProjection(c.ServiceAccountToken)
					}
				}
			}
			if b.VolumeSource.ScaleIO != nil {
				corev1.SetDefaults_ScaleIOVolumeSource(b.VolumeSource.ScaleIO)
			}
		}
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.InitContainers {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.InitContainers[j]
			corev1.SetDefaults_Container(b)
			for k := range b.Ports {
				c := &b.Ports[k]
				corev1.SetDefaults_ContainerPort(c)
			}
			for k := range b.Env {
				c := &b.Env[k]
				if c.ValueFrom != nil {
					if c.ValueFrom.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(c.ValueFrom.FieldRef)
					}
				}
			}
			corev1.SetDefaults_ResourceList(&b.Resources.Limits)
			corev1.SetDefaults_ResourceList(&b.Resources.Requests)
			if b.LivenessProbe != nil {
				corev1.SetDefaults_Probe(b.LivenessProbe)
				if b.LivenessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(b.LivenessProbe.Handler.HTTPGet)
				}
			}
			if b.ReadinessProbe != nil {
				corev1.SetDefaults_Probe(b.ReadinessProbe)
				if b.ReadinessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(b.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if b.Lifecycle != nil {
				if b.Lifecycle.PostStart != nil {
					if b.Lifecycle.PostStart.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(b.Lifecycle.PostStart.HTTPGet)
					}
				}
				if b.Lifecycle.PreStop != nil {
					if b.Lifecycle.PreStop.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(b.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
		for j := range a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Containers {
			b := &a.PodSecurityPolicySubjectReviewStatus.Template.Spec.Containers[j]
			corev1.SetDefaults_Container(b)
			for k := range b.Ports {
				c := &b.Ports[k]
				corev1.SetDefaults_ContainerPort(c)
			}
			for k := range b.Env {
				c := &b.Env[k]
				if c.ValueFrom != nil {
					if c.ValueFrom.FieldRef != nil {
						corev1.SetDefaults_ObjectFieldSelector(c.ValueFrom.FieldRef)
					}
				}
			}
			corev1.SetDefaults_ResourceList(&b.Resources.Limits)
			corev1.SetDefaults_ResourceList(&b.Resources.Requests)
			if b.LivenessProbe != nil {
				corev1.SetDefaults_Probe(b.LivenessProbe)
				if b.LivenessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(b.LivenessProbe.Handler.HTTPGet)
				}
			}
			if b.ReadinessProbe != nil {
				corev1.SetDefaults_Probe(b.ReadinessProbe)
				if b.ReadinessProbe.Handler.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(b.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if b.Lifecycle != nil {
				if b.Lifecycle.PostStart != nil {
					if b.Lifecycle.PostStart.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(b.Lifecycle.PostStart.HTTPGet)
					}
				}
				if b.Lifecycle.PreStop != nil {
					if b.Lifecycle.PreStop.HTTPGet != nil {
						corev1.SetDefaults_HTTPGetAction(b.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
	}
}
func SetObjectDefaults_PodSecurityPolicySelfSubjectReview(in *v1.PodSecurityPolicySelfSubjectReview) {
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
	corev1.SetDefaults_PodSpec(&in.Status.Template.Spec)
	for i := range in.Status.Template.Spec.Volumes {
		a := &in.Status.Template.Spec.Volumes[i]
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
	for i := range in.Status.Template.Spec.InitContainers {
		a := &in.Status.Template.Spec.InitContainers[i]
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
	for i := range in.Status.Template.Spec.Containers {
		a := &in.Status.Template.Spec.Containers[i]
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
func SetObjectDefaults_PodSecurityPolicySubjectReview(in *v1.PodSecurityPolicySubjectReview) {
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
	corev1.SetDefaults_PodSpec(&in.Status.Template.Spec)
	for i := range in.Status.Template.Spec.Volumes {
		a := &in.Status.Template.Spec.Volumes[i]
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
	for i := range in.Status.Template.Spec.InitContainers {
		a := &in.Status.Template.Spec.InitContainers[i]
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
	for i := range in.Status.Template.Spec.Containers {
		a := &in.Status.Template.Spec.Containers[i]
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
