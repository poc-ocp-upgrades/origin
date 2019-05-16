package validation

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"reflect"
)

type WrappingValidator struct {
	validate       *reflect.Value
	validateUpdate *reflect.Value
}

func (v *WrappingValidator) Validate(obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return callValidate(reflect.ValueOf(obj), *v.validate)
}
func (v *WrappingValidator) ValidateUpdate(obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if v.validateUpdate == nil {
		return field.ErrorList{field.Forbidden(field.NewPath("obj"), fmt.Sprintf("%v", obj))}
	}
	return callValidateUpdate(reflect.ValueOf(obj), reflect.ValueOf(old), *v.validateUpdate)
}
func NewValidationWrapper(validateFunction interface{}, validateUpdateFunction interface{}) (*WrappingValidator, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	validateFunctionValue := reflect.ValueOf(validateFunction)
	validateType := validateFunctionValue.Type()
	if err := verifyValidateFunctionSignature(validateType); err != nil {
		return nil, err
	}
	var validateUpdateFunctionValue *reflect.Value
	if validateUpdateFunction != nil {
		functionValue := reflect.ValueOf(validateUpdateFunction)
		validateUpdateType := functionValue.Type()
		if err := verifyValidateUpdateFunctionSignature(validateUpdateType); err != nil {
			return nil, err
		}
		validateUpdateFunctionValue = &functionValue
	}
	return &WrappingValidator{&validateFunctionValue, validateUpdateFunctionValue}, nil
}
func verifyValidateFunctionSignature(ft reflect.Type) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ft.Kind() != reflect.Func {
		return fmt.Errorf("expected func, got: %v", ft)
	}
	if ft.NumIn() != 1 {
		return fmt.Errorf("expected one 'in' param, got: %v", ft)
	}
	if ft.NumOut() != 1 {
		return fmt.Errorf("expected one 'out' param, got: %v", ft)
	}
	if ft.In(0).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'in' param 0, got: %v", ft)
	}
	errorType := reflect.TypeOf(&field.ErrorList{}).Elem()
	if ft.Out(0) != errorType {
		return fmt.Errorf("expected field.ErrorList return, got: %v", ft)
	}
	return nil
}
func verifyValidateUpdateFunctionSignature(ft reflect.Type) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ft.Kind() != reflect.Func {
		return fmt.Errorf("expected func, got: %v", ft)
	}
	if ft.NumIn() != 2 {
		return fmt.Errorf("expected two 'in' params, got: %v", ft)
	}
	if ft.NumOut() != 1 {
		return fmt.Errorf("expected one 'out' param, got: %v", ft)
	}
	if ft.In(0).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'in' param 0, got: %v", ft)
	}
	if ft.In(1).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'in' param 1, got: %v", ft)
	}
	errorType := reflect.TypeOf(&field.ErrorList{}).Elem()
	if ft.Out(0) != errorType {
		return fmt.Errorf("expected field.ErrorList return, got: %v", ft)
	}
	return nil
}
func callValidate(obj, validateMethod reflect.Value) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []reflect.Value{obj}
	ret := validateMethod.Call(args)[0].Interface()
	if ret == nil {
		return nil
	}
	return ret.(field.ErrorList)
}
func callValidateUpdate(obj, old, validateMethod reflect.Value) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	args := []reflect.Value{obj, old}
	ret := validateMethod.Call(args)[0].Interface()
	if ret == nil {
		return nil
	}
	return ret.(field.ErrorList)
}
