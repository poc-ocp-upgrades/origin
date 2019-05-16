package openshiftadmission

import (
	goformat "fmt"
	configv1 "github.com/openshift/api/config/v1"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/apis/apiserver"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func convertOpenshiftAdmissionConfigToKubeAdmissionConfig(in map[string]configv1.AdmissionPluginConfig) (*apiserver.AdmissionConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := &apiserver.AdmissionConfiguration{}
	for _, pluginName := range sets.StringKeySet(in).List() {
		openshiftConfig := in[pluginName]
		kubeConfig := apiserver.AdmissionPluginConfiguration{Name: pluginName, Path: openshiftConfig.Location}
		kubeConfig.Configuration = &runtime.Unknown{Raw: openshiftConfig.Configuration.Raw}
		ret.Plugins = append(ret.Plugins, kubeConfig)
	}
	return ret, nil
}
func ToAdmissionConfigFile(pluginConfig map[string]configv1.AdmissionPluginConfig) (string, func(), error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cleanupFn := func() {
	}
	upstreamAdmissionConfig, err := convertOpenshiftAdmissionConfigToKubeAdmissionConfig(pluginConfig)
	if err != nil {
		return "", cleanupFn, err
	}
	configBytes, err := configapilatest.WriteYAML(upstreamAdmissionConfig)
	if err != nil {
		return "", cleanupFn, err
	}
	tempFile, err := ioutil.TempFile("", "master-config.yaml")
	if err != nil {
		return "", cleanupFn, err
	}
	cleanupFn = func() {
		os.Remove(tempFile.Name())
	}
	if _, err := tempFile.Write(configBytes); err != nil {
		return "", cleanupFn, err
	}
	tempFile.Close()
	return tempFile.Name(), cleanupFn, err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
