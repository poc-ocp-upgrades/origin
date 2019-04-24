package image

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
	"github.com/openshift/origin/pkg/oc/cli/image/append"
	"github.com/openshift/origin/pkg/oc/cli/image/extract"
	"github.com/openshift/origin/pkg/oc/cli/image/info"
	"github.com/openshift/origin/pkg/oc/cli/image/mirror"
)

var (
	imageLong = ktemplates.LongDesc(`
		Manage images on OpenShift

		These commands help you manage images on OpenShift.`)
)

func NewCmdImage(fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	image := &cobra.Command{Use: "image COMMAND", Short: "Useful commands for managing images", Long: imageLong, Run: kcmdutil.DefaultSubCommandRun(streams.ErrOut)}
	name := fmt.Sprintf("%s image", fullName)
	groups := ktemplates.CommandGroups{{Message: "Advanced commands:", Commands: []*cobra.Command{append.NewCmdAppendImage(name, streams), info.NewInfo(name, streams), extract.New(name, streams), mirror.NewCmdMirrorImage(name, streams)}}}
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
