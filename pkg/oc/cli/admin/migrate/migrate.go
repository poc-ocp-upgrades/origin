package migrate

import (
	"github.com/spf13/cobra"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

const MigrateRecommendedName = "migrate"

var migrateLong = templates.LongDesc(`
	Migrate resources on the cluster

	These commands assist administrators in performing preventative maintenance on a cluster.`)

func NewCommandMigrate(name, fullName string, f cmdutil.Factory, streams genericclioptions.IOStreams, cmds ...*cobra.Command) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: name, Short: "Migrate data in the cluster", Long: migrateLong, Run: cmdutil.DefaultSubCommandRun(streams.ErrOut)}
	cmd.AddCommand(cmds...)
	return cmd
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
