package util

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"os"
	"strings"
)

func SiblingCommand(cmd *cobra.Command, name string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := cmd.Parent()
	command := []string{}
	for c != nil {
		klog.V(5).Infof("Found parent command: %s", c.Name())
		command = append([]string{c.Name()}, command...)
		c = c.Parent()
	}
	klog.V(4).Infof("Setting root command to: %s", os.Args[0])
	command[0] = os.Args[0]
	command = append(command, name)
	klog.V(4).Infof("The sibling command is: %s", strings.Join(command, " "))
	return command
}
