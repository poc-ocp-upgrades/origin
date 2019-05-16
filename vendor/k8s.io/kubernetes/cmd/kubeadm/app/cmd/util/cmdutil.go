package util

import (
	goformat "fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func SubCmdRunE(name string) func(*cobra.Command, []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.Errorf("missing subcommand; %q is not meant to be run on its own", name)
		}
		return errors.Errorf("invalid subcommand: %q", args[0])
	}
}
func ValidateExactArgNumber(args []string, supportedArgs []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lenSupported := len(supportedArgs)
	validArgs := 0
	for _, arg := range args {
		if len(arg) > 0 {
			validArgs++
		}
		if validArgs > lenSupported {
			return errors.Errorf("too many arguments. Required arguments: %v", supportedArgs)
		}
	}
	if validArgs < lenSupported {
		return errors.Errorf("missing one or more required arguments. Required arguments: %v", supportedArgs)
	}
	return nil
}
func FindExistingKubeConfig(file string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if file != kubeadmconstants.GetAdminKubeConfigPath() {
		return file
	}
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.Precedence = append(rules.Precedence, kubeadmconstants.GetAdminKubeConfigPath())
	return rules.GetDefaultFilename()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
