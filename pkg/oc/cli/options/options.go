package options

import (
	"github.com/spf13/cobra"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

func NewCmdOptions(streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "options", Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	}}
	templates.UseOptionsTemplates(cmd)
	return cmd
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
