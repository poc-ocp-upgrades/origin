package create

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

type CreateSubcommandOptions struct {
	genericclioptions.IOStreams
	PrintFlags		*genericclioptions.PrintFlags
	Name			string
	DryRun			bool
	Namespace		string
	EnforceNamespace	bool
	Printer			printers.ResourcePrinter
}

func NewCreateSubcommandOptions(ioStreams genericclioptions.IOStreams) *CreateSubcommandOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &CreateSubcommandOptions{PrintFlags: genericclioptions.NewPrintFlags("created").WithTypeSetter(scheme.Scheme), IOStreams: ioStreams}
}
func (o *CreateSubcommandOptions) Complete(f genericclioptions.RESTClientGetter, cmd *cobra.Command, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	name, err := NameFromCommandArgs(cmd, args)
	if err != nil {
		return err
	}
	o.Name = name
	o.Namespace, o.EnforceNamespace, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	o.DryRun = cmdutil.GetDryRunFlag(cmd)
	if o.DryRun {
		o.PrintFlags.Complete("%s (dry run)")
	}
	o.Printer, err = o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	return nil
}
func NameFromCommandArgs(cmd *cobra.Command, args []string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	argsLen := cmd.ArgsLenAtDash()
	if argsLen == -1 {
		argsLen = len(args)
	}
	if argsLen != 1 {
		return "", cmdutil.UsageErrorf(cmd, "exactly one NAME is required, got %d", argsLen)
	}
	return args[0], nil
}
