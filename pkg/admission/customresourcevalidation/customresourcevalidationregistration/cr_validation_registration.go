package customresourcevalidationregistration

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/authentication"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/clusterresourcequota"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/config"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/console"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/features"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/image"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/oauth"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/project"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/rolebindingrestriction"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/scheduler"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation/securitycontextconstraints"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var AllCustomResourceValidators = []string{authentication.PluginName, features.PluginName, console.PluginName, image.PluginName, oauth.PluginName, project.PluginName, config.PluginName, scheduler.PluginName, clusterresourcequota.PluginName, securitycontextconstraints.PluginName, rolebindingrestriction.PluginName, securitycontextconstraints.DefaultingPluginName}

func RegisterCustomResourceValidation(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authentication.Register(plugins)
	features.Register(plugins)
	console.Register(plugins)
	image.Register(plugins)
	oauth.Register(plugins)
	project.Register(plugins)
	config.Register(plugins)
	scheduler.Register(plugins)
	clusterresourcequota.Register(plugins)
	securitycontextconstraints.Register(plugins)
	rolebindingrestriction.Register(plugins)
	securitycontextconstraints.RegisterDefaulting(plugins)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
