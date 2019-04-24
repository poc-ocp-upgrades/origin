package importer

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"github.com/spf13/cobra"
	"github.com/openshift/origin/pkg/oc/cli/importer/appjson"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

var (
	importLong = templates.LongDesc(`
		Import outside applications into OpenShift

		These commands assist in bringing existing applications into OpenShift.`)
)

func NewCmdImport(fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "import COMMAND", Short: "Commands that import applications", Long: importLong, Run: kcmdutil.DefaultSubCommandRun(streams.ErrOut)}
	name := fmt.Sprintf("%s import", fullName)
	cmd.AddCommand(appjson.NewCmdAppJSON(name, f, streams))
	return cmd
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
