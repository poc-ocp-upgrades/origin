package v1

import (
	v1 "github.com/openshift/api/build/v1"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&v1.Build{}, func(obj interface{}) {
		SetObjectDefaults_Build(obj.(*v1.Build))
	})
	scheme.AddTypeDefaultingFunc(&v1.BuildConfig{}, func(obj interface{}) {
		SetObjectDefaults_BuildConfig(obj.(*v1.BuildConfig))
	})
	scheme.AddTypeDefaultingFunc(&v1.BuildConfigList{}, func(obj interface{}) {
		SetObjectDefaults_BuildConfigList(obj.(*v1.BuildConfigList))
	})
	scheme.AddTypeDefaultingFunc(&v1.BuildList{}, func(obj interface{}) {
		SetObjectDefaults_BuildList(obj.(*v1.BuildList))
	})
	scheme.AddTypeDefaultingFunc(&v1.BuildRequest{}, func(obj interface{}) {
		SetObjectDefaults_BuildRequest(obj.(*v1.BuildRequest))
	})
	return nil
}
func SetObjectDefaults_Build(in *v1.Build) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_BuildSource(&in.Spec.CommonSpec.Source)
	SetDefaults_BuildStrategy(&in.Spec.CommonSpec.Strategy)
	if in.Spec.CommonSpec.Strategy.DockerStrategy != nil {
		SetDefaults_DockerBuildStrategy(in.Spec.CommonSpec.Strategy.DockerStrategy)
		for i := range in.Spec.CommonSpec.Strategy.DockerStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.DockerStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
		for i := range in.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs {
			a := &in.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.SourceStrategy != nil {
		SetDefaults_SourceBuildStrategy(in.Spec.CommonSpec.Strategy.SourceStrategy)
		for i := range in.Spec.CommonSpec.Strategy.SourceStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.SourceStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.CustomStrategy != nil {
		SetDefaults_CustomBuildStrategy(in.Spec.CommonSpec.Strategy.CustomStrategy)
		for i := range in.Spec.CommonSpec.Strategy.CustomStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.CustomStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy != nil {
		for i := range in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	corev1.SetDefaults_ResourceList(&in.Spec.CommonSpec.Resources.Limits)
	corev1.SetDefaults_ResourceList(&in.Spec.CommonSpec.Resources.Requests)
}
func SetObjectDefaults_BuildConfig(in *v1.BuildConfig) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_BuildConfigSpec(&in.Spec)
	for i := range in.Spec.Triggers {
		a := &in.Spec.Triggers[i]
		SetDefaults_BuildTriggerPolicy(a)
	}
	SetDefaults_BuildSource(&in.Spec.CommonSpec.Source)
	SetDefaults_BuildStrategy(&in.Spec.CommonSpec.Strategy)
	if in.Spec.CommonSpec.Strategy.DockerStrategy != nil {
		SetDefaults_DockerBuildStrategy(in.Spec.CommonSpec.Strategy.DockerStrategy)
		for i := range in.Spec.CommonSpec.Strategy.DockerStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.DockerStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
		for i := range in.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs {
			a := &in.Spec.CommonSpec.Strategy.DockerStrategy.BuildArgs[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.SourceStrategy != nil {
		SetDefaults_SourceBuildStrategy(in.Spec.CommonSpec.Strategy.SourceStrategy)
		for i := range in.Spec.CommonSpec.Strategy.SourceStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.SourceStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.CustomStrategy != nil {
		SetDefaults_CustomBuildStrategy(in.Spec.CommonSpec.Strategy.CustomStrategy)
		for i := range in.Spec.CommonSpec.Strategy.CustomStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.CustomStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	if in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy != nil {
		for i := range in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy.Env {
			a := &in.Spec.CommonSpec.Strategy.JenkinsPipelineStrategy.Env[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
	corev1.SetDefaults_ResourceList(&in.Spec.CommonSpec.Resources.Limits)
	corev1.SetDefaults_ResourceList(&in.Spec.CommonSpec.Resources.Requests)
}
func SetObjectDefaults_BuildConfigList(in *v1.BuildConfigList) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_BuildConfig(a)
	}
}
func SetObjectDefaults_BuildList(in *v1.BuildList) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Build(a)
	}
}
func SetObjectDefaults_BuildRequest(in *v1.BuildRequest) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Env {
		a := &in.Env[i]
		if a.ValueFrom != nil {
			if a.ValueFrom.FieldRef != nil {
				corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
			}
		}
	}
	if in.DockerStrategyOptions != nil {
		for i := range in.DockerStrategyOptions.BuildArgs {
			a := &in.DockerStrategyOptions.BuildArgs[i]
			if a.ValueFrom != nil {
				if a.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(a.ValueFrom.FieldRef)
				}
			}
		}
	}
}
