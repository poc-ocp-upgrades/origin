package v1

import (
	"github.com/openshift/api/oauth/v1"
	"github.com/openshift/origin/pkg/oauth/apis/oauth"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(oauth.Install, v1.Install, addFieldSelectorKeyConversions, RegisterDefaults)
	Install            = localSchemeBuilder.AddToScheme
)
