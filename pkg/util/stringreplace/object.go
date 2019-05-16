package stringreplace

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"k8s.io/klog"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

func VisitObjectStrings(obj interface{}, visitor func(string) (string, bool)) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return visitValue(reflect.ValueOf(obj), visitor)
}
func visitValue(v reflect.Value, visitor func(string) (string, bool)) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v.IsNil() {
			return nil
		}
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		err := visitValue(v.Elem(), visitor)
		if err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		vt := v.Type().Elem()
		for i := 0; i < v.Len(); i++ {
			val, err := visitUnsettableValues(vt, v.Index(i), visitor)
			if err != nil {
				return err
			}
			v.Index(i).Set(val)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			err := visitValue(v.Field(i), visitor)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		vt := v.Type().Elem()
		for _, oldKey := range v.MapKeys() {
			newKey, err := visitUnsettableValues(oldKey.Type(), oldKey, visitor)
			if err != nil {
				return err
			}
			oldValue := v.MapIndex(oldKey)
			newValue, err := visitUnsettableValues(vt, oldValue, visitor)
			if err != nil {
				return err
			}
			v.SetMapIndex(oldKey, reflect.Value{})
			v.SetMapIndex(newKey, newValue)
		}
	case reflect.String:
		if !v.CanSet() {
			return fmt.Errorf("unable to set String value '%v'", v)
		}
		s, asString := visitor(v.String())
		if !asString {
			return fmt.Errorf("attempted to set String field to non-string value '%v'", s)
		}
		v.SetString(s)
	default:
		klog.V(5).Infof("Ignoring non-parameterizable field type '%s': %v", v.Kind(), v)
		return nil
	}
	return nil
}
func visitUnsettableValues(typeOf reflect.Type, original reflect.Value, visitor func(string) (string, bool)) (reflect.Value, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	val := reflect.New(typeOf).Elem()
	existing := original
	if existing.CanInterface() {
		existing = reflect.ValueOf(existing.Interface())
	}
	switch existing.Kind() {
	case reflect.String:
		s, asString := visitor(existing.String())
		if asString {
			val = reflect.ValueOf(s)
		} else {
			b := []byte(s)
			var data interface{}
			err := json.Unmarshal(b, &data)
			if err != nil {
				val = reflect.ValueOf(s)
			} else {
				val = reflect.ValueOf(data)
			}
		}
	default:
		if existing.IsValid() && existing.Kind() != reflect.Invalid {
			val.Set(existing)
		}
		visitValue(val, visitor)
	}
	return val, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
