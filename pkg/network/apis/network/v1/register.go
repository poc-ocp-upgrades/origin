package v1

import (
	"github.com/openshift/api/network/v1"
	"github.com/openshift/origin/pkg/network/apis/network"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(network.Install, v1.Install, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)
