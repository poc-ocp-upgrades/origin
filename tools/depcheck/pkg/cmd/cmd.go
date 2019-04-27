package cmd

import (
	"flag"
	"fmt"
	"io"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewCmdDepCheck(name string, out, errout io.Writer) *cobra.Command {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: fmt.Sprintf("%s (ARGUMENT) [OPTIONS]", name), Short: "Gather information about a dependency tree.", Long: "Modify or gather information about a dependency tree.", Example: fmt.Sprintf(pinImportsExample, name), RunE: func(c *cobra.Command, args []string) error {
		c.SetOutput(errout)
		return c.Help()
	}}
	cmd.AddCommand(NewCmdPinImports(name, out, errout))
	cmd.AddCommand(NewCmdTraceImports(name, out, errout))
	cmd.AddCommand(NewCmdAnalyzeImports(name, out, errout))
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.Set("logtostderr", "true")
	return cmd
}
