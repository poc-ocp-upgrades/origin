package v1

import (
	"github.com/openshift/api/template/v1"
	"github.com/openshift/origin/pkg/template/apis/template"
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(template.Install, v1.Install, corev1conversions.AddToScheme, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)
