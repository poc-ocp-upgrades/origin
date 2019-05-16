package compat

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	goos "os"
	"reflect"
	"regexp"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"testing"
	gotime "time"
)

func TestCompatibility(t *testing.T, version schema.GroupVersion, input []byte, validator func(obj runtime.Object) field.ErrorList, expectedKeys map[string]string, absentKeys []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	codec := legacyscheme.Codecs.LegacyCodec(version)
	obj, err := runtime.Decode(codec, input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	errs := validator(obj)
	if len(errs) != 0 {
		t.Fatalf("Unexpected validation errors: %v", errs)
	}
	output, err := runtime.Encode(codec, obj)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	generic := map[string]interface{}{}
	if err := json.Unmarshal(output, &generic); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	for k, expectedValue := range expectedKeys {
		keys := strings.Split(k, ".")
		if actualValue, ok, err := getJSONValue(generic, keys...); err != nil || !ok {
			t.Errorf("Unexpected error for %s: %v", k, err)
		} else if !reflect.DeepEqual(expectedValue, fmt.Sprintf("%v", actualValue)) {
			t.Errorf("Unexpected value for %v: expected %v, got %v", k, expectedValue, actualValue)
		}
	}
	for _, absentKey := range absentKeys {
		keys := strings.Split(absentKey, ".")
		actualValue, ok, err := getJSONValue(generic, keys...)
		if err == nil || ok {
			t.Errorf("Unexpected value found for key %s: %v", absentKey, actualValue)
		}
	}
	if t.Failed() {
		data, err := json.MarshalIndent(obj, "", "    ")
		if err != nil {
			t.Log(err)
		} else {
			t.Log(string(data))
		}
		t.Logf("2: Encoded value: %v", string(output))
	}
}
func getJSONValue(data map[string]interface{}, keys ...string) (interface{}, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
