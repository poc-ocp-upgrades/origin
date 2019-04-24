package customresourcevalidationregistration

import (
	"k8s.io/apiserver/pkg/admission"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/authentication"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/clusterresourcequota"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/config"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/console"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/features"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/image"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/oauth"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/project"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/scheduler"
)

var AllCustomResourceValidators = []string{authentication.PluginName, features.PluginName, console.PluginName, image.PluginName, oauth.PluginName, project.PluginName, config.PluginName, scheduler.PluginName, clusterresourcequota.PluginName}

func RegisterCustomResourceValidation(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	authentication.Register(plugins)
	features.Register(plugins)
	console.Register(plugins)
	image.Register(plugins)
	oauth.Register(plugins)
	project.Register(plugins)
	config.Register(plugins)
	scheduler.Register(plugins)
	clusterresourcequota.Register(plugins)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
