package options

import (
	"errors"
	goformat "fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	kubeschedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	kubeschedulerscheme "k8s.io/kubernetes/pkg/scheduler/apis/config/scheme"
	kubeschedulerconfigv1alpha1 "k8s.io/kubernetes/pkg/scheduler/apis/config/v1alpha1"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func loadConfigFromFile(file string) (*kubeschedulerconfig.KubeSchedulerConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return loadConfig(data)
}
func loadConfig(data []byte) (*kubeschedulerconfig.KubeSchedulerConfiguration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configObj := &kubeschedulerconfig.KubeSchedulerConfiguration{}
	if err := runtime.DecodeInto(kubeschedulerscheme.Codecs.UniversalDecoder(), data, configObj); err != nil {
		return nil, err
	}
	return configObj, nil
}
func WriteConfigFile(fileName string, cfg *kubeschedulerconfig.KubeSchedulerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
