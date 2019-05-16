package main

import (
	"errors"
	goflag "flag"
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/server"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	cloudcontrollermanager "k8s.io/kubernetes/cmd/cloud-controller-manager/app"
	kubeapiserver "k8s.io/kubernetes/cmd/kube-apiserver/app"
	kubecontrollermanager "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kubeproxy "k8s.io/kubernetes/cmd/kube-proxy/app"
	kubescheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus"
	kubectl "k8s.io/kubernetes/pkg/kubectl/cmd"
	_ "k8s.io/kubernetes/pkg/version/prometheus"
	"math/rand"
	"os"
	goos "os"
	"path"
	"path/filepath"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rand.Seed(time.Now().UnixNano())
	hyperkubeCommand, allCommandFns := NewHyperKubeCommand(server.SetupSignalHandler())
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	logs.InitLogs()
	defer logs.FlushLogs()
	basename := filepath.Base(os.Args[0])
	if err := commandFor(basename, hyperkubeCommand, allCommandFns).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
func commandFor(basename string, defaultCommand *cobra.Command, commands []func() *cobra.Command) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, commandFn := range commands {
		command := commandFn()
		if command.Name() == basename {
			return command
		}
		for _, alias := range command.Aliases {
			if alias == basename {
				return command
			}
		}
	}
	return defaultCommand
}
func NewHyperKubeCommand(stopCh <-chan struct{}) (*cobra.Command, []func() *cobra.Command) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiserver := func() *cobra.Command {
		ret := kubeapiserver.NewAPIServerCommand(stopCh)
		ret.Aliases = []string{"apiserver"}
		return ret
	}
	controller := func() *cobra.Command {
		ret := kubecontrollermanager.NewControllerManagerCommand(stopCh)
		ret.Aliases = []string{"controller-manager"}
		return ret
	}
	proxy := func() *cobra.Command {
		ret := kubeproxy.NewProxyCommand()
		ret.Aliases = []string{"proxy"}
		return ret
	}
	scheduler := func() *cobra.Command {
		ret := kubescheduler.NewSchedulerCommand(stopCh)
		ret.Aliases = []string{"scheduler"}
		return ret
	}
	kubectlCmd := func() *cobra.Command {
		return kubectl.NewDefaultKubectlCommand()
	}
	kubelet := func() *cobra.Command {
		return kubelet.NewKubeletCommand(stopCh)
	}
	cloudController := func() *cobra.Command {
		return cloudcontrollermanager.NewCloudControllerManagerCommand()
	}
	commandFns := []func() *cobra.Command{apiserver, controller, proxy, scheduler, kubectlCmd, kubelet, cloudController}
	makeSymlinksFlag := false
	cmd := &cobra.Command{Use: "hyperkube", Short: "Request a new project", Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 || !makeSymlinksFlag {
			cmd.Help()
			os.Exit(1)
		}
		if err := makeSymlinks(os.Args[0], commandFns); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		}
	}}
	cmd.Flags().BoolVar(&makeSymlinksFlag, "make-symlinks", makeSymlinksFlag, "create a symlink for each server in current directory")
	cmd.Flags().MarkHidden("make-symlinks")
	for i := range commandFns {
		cmd.AddCommand(commandFns[i]())
	}
	return cmd, commandFns
}
func makeSymlinks(targetName string, commandFns []func() *cobra.Command) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	var errs bool
	for _, commandFn := range commandFns {
		command := commandFn()
		link := path.Join(wd, command.Name())
		err := os.Symlink(targetName, link)
		if err != nil {
			errs = true
			fmt.Println(err)
		}
	}
	if errs {
		return errors.New("Error creating one or more symlinks.")
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
