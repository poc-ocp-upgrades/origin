package v1

import (
	"github.com/openshift/api/route/v1"
	"github.com/openshift/origin/pkg/route/apis/route"
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(route.Install, v1.Install, corev1conversions.AddToScheme, addFieldSelectorKeyConversions, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)
