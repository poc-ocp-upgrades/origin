package apihelpers

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func LegacyMetaV1FieldSelectorConversionWithName(label, value string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch label {
	case "name":
		return "metadata.name", value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
