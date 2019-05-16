package v1

import (
	buildv1 "github.com/openshift/api/build/v1"
	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	buildinternalconversions "github.com/openshift/origin/pkg/build/apis/build/v1"
	"github.com/openshift/origin/pkg/cmd/server/apis/config"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	coreinternalconversions "k8s.io/kubernetes/pkg/apis/core"
)

var (
	LegacyGroupName             = ""
	LegacySchemeGroupVersion    = schema.GroupVersion{Group: LegacyGroupName, Version: "v1"}
	legacySchemeBuilder         = runtime.NewSchemeBuilder(legacyconfigv1.InstallLegacy, config.InstallLegacy, coreinternalconversions.AddToScheme, buildinternalconversions.Install, RegisterConversions, addConversionFuncs, addDefaultingFuncs)
	InstallLegacy               = legacySchemeBuilder.AddToScheme
	externalLegacySchemeBuilder = runtime.NewSchemeBuilder(legacyconfigv1.InstallLegacy, buildv1.Install)
	InstallLegacyExternal       = externalLegacySchemeBuilder.AddToScheme
	localSchemeBuilder          = runtime.NewSchemeBuilder()
)
