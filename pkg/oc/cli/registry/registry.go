package registry

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	ktemplates "k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/origin/pkg/cmd/templates"
	"github.com/openshift/origin/pkg/oc/cli/registry/info"
	"github.com/openshift/origin/pkg/oc/cli/registry/login"
)

var (
	imageLong = ktemplates.LongDesc(`
		Manage the integrated registry on OpenShift

		These commands help you work with an integrated OpenShift registry.`)
)

func NewCmd(fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	image := &cobra.Command{Use: "registry COMMAND", Short: "Commands for working with the registry", Long: imageLong, Run: kcmdutil.DefaultSubCommandRun(streams.ErrOut)}
	name := fmt.Sprintf("%s registry", fullName)
	groups := ktemplates.CommandGroups{{Message: "Advanced commands:", Commands: []*cobra.Command{info.NewRegistryInfoCmd(name, f, streams), login.NewRegistryLoginCmd(name, f, streams)}}}
	groups.Add(image)
	templates.ActsAsRootCommand(image, []string{"options"}, groups...)
	return image
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
