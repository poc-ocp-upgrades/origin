package alpha

import (
	goformat "fmt"
	"github.com/spf13/cobra"
	"io"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewCmdAlpha(in io.Reader, out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "alpha", Short: "Kubeadm experimental sub-commands"}
	cmd.AddCommand(newCmdCertsUtility())
	cmd.AddCommand(newCmdKubeletUtility())
	cmd.AddCommand(newCmdKubeConfigUtility(out))
	cmd.AddCommand(newCmdPreFlightUtility())
	cmd.AddCommand(NewCmdSelfhosting(in))
	cmd.AddCommand(newCmdPhase(out))
	return cmd
}
func newCmdPhase(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "phase", Short: "Invoke subsets of kubeadm functions separately for a manual install", Long: cmdutil.MacroCommandLongDescription}
	return cmd
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
