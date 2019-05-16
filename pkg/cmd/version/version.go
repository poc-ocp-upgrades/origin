package version

import (
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/apimachinery/pkg/version"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func NewCmdVersion(fullName string, versionInfo version.Info, out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "version", Short: "Display version", Long: "Display version", Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(out, "%s %v\n", fullName, versionInfo)
	}}
	return cmd
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
