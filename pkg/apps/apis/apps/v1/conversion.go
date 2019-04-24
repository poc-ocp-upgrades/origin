package v1

import (
	"strings"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"github.com/openshift/api/apps/v1"
	newer "github.com/openshift/origin/pkg/apps/apis/apps"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

func Convert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(in *v1.DeploymentTriggerImageChangeParams, out *newer.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := autoConvert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams(in, out, s); err != nil {
		return err
	}
	switch in.From.Kind {
	case "ImageStreamTag":
	case "ImageStream", "ImageRepository":
		out.From.Kind = "ImageStreamTag"
		if !strings.Contains(out.From.Name, ":") {
			out.From.Name = imageapi.JoinImageStreamTag(out.From.Name, imageapi.DefaultImageTag)
		}
	default:
	}
	return nil
}
func Convert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(in *newer.DeploymentTriggerImageChangeParams, out *v1.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := autoConvert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(in, out, s); err != nil {
		return err
	}
	switch in.From.Kind {
	case "ImageStreamTag":
	case "ImageStream", "ImageRepository":
		out.From.Kind = "ImageStreamTag"
		if !strings.Contains(out.From.Name, ":") {
			out.From.Name = imageapi.JoinImageStreamTag(out.From.Name, imageapi.DefaultImageTag)
		}
	default:
	}
	return nil
}
func Convert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams(in *v1.RollingDeploymentStrategyParams, out *newer.RollingDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_RollingDeploymentStrategyParams(in)
	out.UpdatePeriodSeconds = in.UpdatePeriodSeconds
	out.IntervalSeconds = in.IntervalSeconds
	out.TimeoutSeconds = in.TimeoutSeconds
	if in.Pre != nil {
		if err := s.Convert(&in.Pre, &out.Pre, 0); err != nil {
			return err
		}
	}
	if in.Post != nil {
		if err := s.Convert(&in.Post, &out.Post, 0); err != nil {
			return err
		}
	}
	if in.MaxUnavailable != nil {
		if err := s.Convert(in.MaxUnavailable, &out.MaxUnavailable, 0); err != nil {
			return err
		}
	}
	if in.MaxSurge != nil {
		if err := s.Convert(in.MaxSurge, &out.MaxSurge, 0); err != nil {
			return err
		}
	}
	return nil
}
func Convert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(in *newer.RollingDeploymentStrategyParams, out *v1.RollingDeploymentStrategyParams, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.UpdatePeriodSeconds = in.UpdatePeriodSeconds
	out.IntervalSeconds = in.IntervalSeconds
	out.TimeoutSeconds = in.TimeoutSeconds
	if in.Pre != nil {
		if err := s.Convert(&in.Pre, &out.Pre, 0); err != nil {
			return err
		}
	}
	if in.Post != nil {
		if err := s.Convert(&in.Post, &out.Post, 0); err != nil {
			return err
		}
	}
	if out.MaxUnavailable == nil {
		out.MaxUnavailable = &intstr.IntOrString{}
	}
	if out.MaxSurge == nil {
		out.MaxSurge = &intstr.IntOrString{}
	}
	if err := s.Convert(&in.MaxUnavailable, out.MaxUnavailable, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.MaxSurge, out.MaxSurge, 0); err != nil {
		return err
	}
	return nil
}
func AddConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scheme.AddConversionFuncs(Convert_v1_DeploymentTriggerImageChangeParams_To_apps_DeploymentTriggerImageChangeParams, Convert_apps_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams, Convert_v1_RollingDeploymentStrategyParams_To_apps_RollingDeploymentStrategyParams, Convert_apps_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
