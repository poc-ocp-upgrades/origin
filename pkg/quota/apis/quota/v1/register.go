package v1

import (
	"github.com/openshift/api/quota/v1"
	"github.com/openshift/origin/pkg/quota/apis/quota"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(quota.Install, v1.Install, AddConversionFuncs, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)
