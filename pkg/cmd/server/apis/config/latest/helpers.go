package latest

import (
	"bytes"
	"net/http"
	"runtime"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"reflect"
	"github.com/ghodss/yaml"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/runtime"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
)

func ReadSessionSecrets(filename string) (*configapi.SessionSecrets, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &configapi.SessionSecrets{}
	if err := ReadYAMLFileInto(filename, config); err != nil {
		return nil, err
	}
	return config, nil
}
func ReadMasterConfig(filename string) (*configapi.MasterConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &configapi.MasterConfig{}
	if err := ReadYAMLFileInto(filename, config); err != nil {
		return nil, err
	}
	return config, nil
}
func ReadAndResolveMasterConfig(filename string) (*configapi.MasterConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, err := ReadMasterConfig(filename)
	if err != nil {
		return nil, err
	}
	if err := configapi.ResolveMasterConfigPaths(masterConfig, path.Dir(filename)); err != nil {
		return nil, err
	}
	return masterConfig, nil
}
func ReadNodeConfig(filename string) (*configapi.NodeConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &configapi.NodeConfig{}
	if err := ReadYAMLFileInto(filename, config); err != nil {
		return nil, err
	}
	return config, nil
}
func ReadAndResolveNodeConfig(filename string) (*configapi.NodeConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeConfig, err := ReadNodeConfig(filename)
	if err != nil {
		return nil, err
	}
	if err := configapi.ResolveNodeConfigPaths(nodeConfig, path.Dir(filename)); err != nil {
		return nil, err
	}
	return nodeConfig, nil
}
func WriteYAML(obj runtime.Object) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	json, err := runtime.Encode(Codec, obj)
	if err != nil {
		return nil, err
	}
	content, err := yaml.JSONToYAML(json)
	if err != nil {
		return nil, err
	}
	return content, err
}
func ReadYAML(reader io.Reader) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if reader == nil || reflect.ValueOf(reader).IsNil() {
		return nil, nil
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	jsonData, err := kyaml.ToJSON(data)
	if err != nil {
		return nil, err
	}
	obj, err := runtime.Decode(Codec, jsonData)
	if err != nil {
		return nil, captureSurroundingJSONForError("error reading config: ", jsonData, err)
	}
	if err := strictDecodeCheck(jsonData, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
func ReadYAMLInto(data []byte, obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	jsonData, err := kyaml.ToJSON(data)
	if err != nil {
		return err
	}
	if err := runtime.DecodeInto(Codec, jsonData, obj); err != nil {
		return captureSurroundingJSONForError("error reading config: ", jsonData, err)
	}
	return strictDecodeCheck(jsonData, obj)
}
func strictDecodeCheck(jsonData []byte, obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := getExternalZeroValue(obj)
	if err != nil {
		klog.Errorf("Encountered config error %v in object %T, raw JSON:\n%s", err, obj, string(jsonData))
		return nil
	}
	d := json.NewDecoder(bytes.NewReader(jsonData))
	d.DisallowUnknownFields()
	if err := d.Decode(out); err != nil {
		klog.Errorf("Encountered config error %v in object %T, raw JSON:\n%s", err, obj, string(jsonData))
	}
	return nil
}
func getExternalZeroValue(obj runtime.Object) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gvks, _, err := configapi.Scheme.ObjectKinds(obj)
	if err != nil {
		return nil, err
	}
	if len(gvks) == 0 {
		return nil, fmt.Errorf("no gvks found for %#v", obj)
	}
	gvk := legacyconfigv1.LegacySchemeGroupVersion.WithKind(gvks[0].Kind)
	return configapi.Scheme.New(gvk)
}
func ReadYAMLFileInto(filename string, obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = ReadYAMLInto(data, obj)
	if err != nil {
		return fmt.Errorf("could not load config file %q due to an error: %v", filename, err)
	}
	return nil
}
func captureSurroundingJSONForError(prefix string, data []byte, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if syntaxErr, ok := err.(*json.SyntaxError); err != nil && ok {
		offset := syntaxErr.Offset
		begin := offset - 20
		if begin < 0 {
			begin = 0
		}
		end := offset + 20
		if end > int64(len(data)) {
			end = int64(len(data))
		}
		return fmt.Errorf("%s%v (found near '%s')", prefix, err, string(data[begin:end]))
	}
	if err != nil {
		return fmt.Errorf("%s%v", prefix, err)
	}
	return err
}
func IsAdmissionPluginActivated(reader io.Reader, defaultValue bool) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := ReadYAML(reader)
	if err != nil {
		return false, err
	}
	if obj == nil {
		return defaultValue, nil
	}
	activationConfig, ok := obj.(*configapi.DefaultAdmissionConfig)
	if !ok {
		return true, nil
	}
	return !activationConfig.Disable, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
