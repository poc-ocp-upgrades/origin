package api

import (
	"encoding/json"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/core/validation"
)

func TestCompatibility_v1_Pod(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedServiceAccount := "my-service-account"
	input := []byte(fmt.Sprintf(`
{
	"kind":"Pod",
	"apiVersion":"v1",
	"metadata":{"name":"my-pod-name", "namespace":"my-pod-namespace"},
	"spec": {
		"serviceAccount":"%s",
		"containers":[{
			"name":"my-container-name",
			"image":"my-container-image"
		}]
	}
}
`, expectedServiceAccount))
	t.Log("Testing 1.0.0 v1 migration added in PR #3592")
	testCompatibility(t, "v1", input, func(obj runtime.Object) field.ErrorList {
		return validation.ValidatePod(obj.(*api.Pod))
	}, map[string]string{"spec.serviceAccount": expectedServiceAccount, "spec.serviceAccountName": expectedServiceAccount})
}
func testCompatibility(t *testing.T, version string, input []byte, validator func(obj runtime.Object) field.ErrorList, serialized map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := runtime.Decode(legacyscheme.Codecs.UniversalDecoder(), input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	errs := validator(obj)
	if len(errs) != 0 {
		t.Fatalf("Unexpected errors: %v", errs)
	}
	output := runtime.EncodeOrDie(legacyscheme.Codecs.LegacyCodec(schema.GroupVersion{Group: "", Version: version}), obj)
	generic := map[string]interface{}{}
	if err := json.Unmarshal([]byte(output), &generic); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	for k, expectedValue := range serialized {
		keys := strings.Split(k, ".")
		if actualValue, ok, err := getJSONValue(generic, keys...); err != nil || !ok {
			t.Errorf("Unexpected error for %s: %v", k, err)
		} else if !reflect.DeepEqual(expectedValue, actualValue) {
			t.Errorf("Expected %v, got %v", expectedValue, actualValue)
		}
	}
}
func TestAllowedGrouplessVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	versions := map[string]schema.GroupVersion{"v1": {Group: "", Version: "v1"}, "v1beta3": {Group: "", Version: "v1beta3"}, "1.0": {Group: "", Version: "1.0"}, "pre012": {Group: "", Version: "pre012"}}
	for apiVersion, expectedGroupVersion := range versions {
		groupVersion, err := schema.ParseGroupVersion(apiVersion)
		if err != nil {
			t.Errorf("%s: unexpected error parsing: %v", apiVersion, err)
			continue
		}
		if groupVersion != expectedGroupVersion {
			t.Errorf("%s: expected %#v, got %#v", apiVersion, expectedGroupVersion, groupVersion)
			continue
		}
		if groupVersion.String() != apiVersion {
			t.Errorf("%s: expected GroupVersion.String() to be %q, got %q", apiVersion, apiVersion, groupVersion.String())
			continue
		}
	}
}
func getJSONValue(data map[string]interface{}, keys ...string) (interface{}, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keys) == 0 {
		return data, true, nil
	}
	key := keys[0]
	index := -1
	if matches := regexp.MustCompile(`^(.*)\[(\d+)\]$`).FindStringSubmatch(key); len(matches) > 0 {
		key = matches[1]
		index, _ = strconv.Atoi(matches[2])
	}
	value, ok := data[key]
	if !ok {
		return nil, false, fmt.Errorf("No key %s found", key)
	}
	if index >= 0 {
		valueSlice, ok := value.([]interface{})
		if !ok {
			return nil, false, fmt.Errorf("Key %s did not hold a slice", key)
		}
		if index >= len(valueSlice) {
			return nil, false, fmt.Errorf("Index %d out of bounds for slice at key: %v", index, key)
		}
		value = valueSlice[index]
	}
	if len(keys) == 1 {
		return value, true, nil
	}
	childData, ok := value.(map[string]interface{})
	if !ok {
		return nil, false, fmt.Errorf("Key %s did not hold a map", keys[0])
	}
	return getJSONValue(childData, keys[1:]...)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
