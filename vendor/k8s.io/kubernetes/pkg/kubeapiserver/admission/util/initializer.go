package util

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/initialization"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func IsUpdatingInitializedObject(a admission.Attributes) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetOperation() != admission.Update {
		return false, nil
	}
	oldObj := a.GetOldObject()
	accessor, err := meta.Accessor(oldObj)
	if err != nil {
		return false, err
	}
	if initialization.IsInitialized(accessor.GetInitializers()) {
		return true, nil
	}
	return false, nil
}
func IsUpdatingUninitializedObject(a admission.Attributes) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetOperation() != admission.Update {
		return false, nil
	}
	oldObj := a.GetOldObject()
	accessor, err := meta.Accessor(oldObj)
	if err != nil {
		return false, err
	}
	if initialization.IsInitialized(accessor.GetInitializers()) {
		return false, nil
	}
	return true, nil
}
func IsInitializationCompletion(a admission.Attributes) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetOperation() != admission.Update {
		return false, nil
	}
	oldObj := a.GetOldObject()
	oldInitialized, err := initialization.IsObjectInitialized(oldObj)
	if err != nil {
		return false, err
	}
	if oldInitialized {
		return false, nil
	}
	newObj := a.GetObject()
	newInitialized, err := initialization.IsObjectInitialized(newObj)
	if err != nil {
		return false, err
	}
	return newInitialized, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
