package network

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

const PodNetworkCommandName = "pod-network"

var (
	podNetworkLong = templates.LongDesc(`
		Manage pod network in the cluster

		This command provides common pod network operations for administrators.`)
)

func NewCmdPodNetwork(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
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
	cmds := &cobra.Command{Use: name, Short: "Manage pod network", Long: podNetworkLong, Run: kcmdutil.DefaultSubCommandRun(streams.ErrOut)}
	cmds.AddCommand(NewCmdJoinProjectsNetwork(JoinProjectsNetworkCommandName, fullName+" "+JoinProjectsNetworkCommandName, f, streams))
	cmds.AddCommand(NewCmdMakeGlobalProjectsNetwork(MakeGlobalProjectsNetworkCommandName, fullName+" "+MakeGlobalProjectsNetworkCommandName, f, streams))
	cmds.AddCommand(NewCmdIsolateProjectsNetwork(IsolateProjectsNetworkCommandName, fullName+" "+IsolateProjectsNetworkCommandName, f, streams))
	return cmds
}
