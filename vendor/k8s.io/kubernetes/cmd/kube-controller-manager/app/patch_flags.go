package app

import (
	"fmt"
	"io/ioutil"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
)

func getOpenShiftConfig(configFile string) (map[string]interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := kyaml.ToJSON(configBytes)
	if err != nil {
		return nil, err
	}
	config := map[string]interface{}{}
	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return nil, err
	}
	return config, nil
}
func applyOpenShiftConfigFlags(controllerManagerOptions *options.KubeControllerManagerOptions, openshiftConfig map[string]interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := applyOpenShiftConfigControllerArgs(controllerManagerOptions, openshiftConfig); err != nil {
		return err
	}
	if err := applyOpenShiftConfigDefaultProjectSelector(controllerManagerOptions, openshiftConfig); err != nil {
		return err
	}
	if err := applyOpenShiftConfigKubeDefaultProjectSelector(controllerManagerOptions, openshiftConfig); err != nil {
		return err
	}
	return nil
}
func applyOpenShiftConfigDefaultProjectSelector(controllerManagerOptions *options.KubeControllerManagerOptions, openshiftConfig map[string]interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projectConfig, ok := openshiftConfig["projectConfig"]
	if !ok {
		return nil
	}
	castProjectConfig := projectConfig.(map[string]interface{})
	defaultNodeSelector, ok := castProjectConfig["defaultNodeSelector"]
	if !ok {
		return nil
	}
	controllerManagerOptions.OpenShiftContext.OpenShiftDefaultProjectNodeSelector = defaultNodeSelector.(string)
	return nil
}
func applyOpenShiftConfigKubeDefaultProjectSelector(controllerManagerOptions *options.KubeControllerManagerOptions, openshiftConfig map[string]interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerManagerOptions.OpenShiftContext.KubeDefaultProjectNodeSelector = ""
	return nil
}
func applyOpenShiftConfigControllerArgs(controllerManagerOptions *options.KubeControllerManagerOptions, openshiftConfig map[string]interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var controllerArgs interface{}
	kubeMasterConfig, ok := openshiftConfig["kubernetesMasterConfig"]
	if !ok {
		controllerArgs, ok = openshiftConfig["extendedArguments"]
		if !ok || controllerArgs == nil {
			return nil
		}
	} else {
		castKubeMasterConfig := kubeMasterConfig.(map[string]interface{})
		controllerArgs, ok = castKubeMasterConfig["controllerArguments"]
		if !ok || controllerArgs == nil {
			controllerArgs, ok = openshiftConfig["extendedArguments"]
			if !ok || controllerArgs == nil {
				return nil
			}
		}
	}
	args := map[string][]string{}
	for key, value := range controllerArgs.(map[string]interface{}) {
		for _, arrayValue := range value.([]interface{}) {
			args[key] = append(args[key], arrayValue.(string))
		}
	}
	if err := applyFlags(args, controllerManagerOptions.Flags(KnownControllers(), ControllersDisabledByDefault.List())); len(err) > 0 {
		return kerrors.NewAggregate(err)
	}
	return nil
}
func applyFlags(args map[string][]string, flags apiserverflag.NamedFlagSets) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for key, value := range args {
		found := false
		for _, fs := range flags.FlagSets {
			if flag := fs.Lookup(key); flag != nil {
				for _, s := range value {
					if err := flag.Value.Set(s); err != nil {
						errs = append(errs, field.Invalid(field.NewPath(key), s, fmt.Sprintf("could not be set: %v", err)))
						break
					}
				}
				found = true
			}
		}
		if !found {
			errs = append(errs, field.Invalid(field.NewPath("flag"), key, "is not a valid flag"))
		}
	}
	return errs
}
