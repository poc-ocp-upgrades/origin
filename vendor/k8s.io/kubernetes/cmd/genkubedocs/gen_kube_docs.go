package main

import (
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/server"
	ccmapp "k8s.io/kubernetes/cmd/cloud-controller-manager/app"
	"k8s.io/kubernetes/cmd/genutils"
	apiservapp "k8s.io/kubernetes/cmd/kube-apiserver/app"
	cmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	proxyapp "k8s.io/kubernetes/cmd/kube-proxy/app"
	schapp "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubeadmapp "k8s.io/kubernetes/cmd/kubeadm/app/cmd"
	kubeletapp "k8s.io/kubernetes/cmd/kubelet/app"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	path := ""
	module := ""
	if len(os.Args) == 3 {
		path = os.Args[1]
		module = os.Args[2]
	} else {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory] [module] \n", os.Args[0])
		os.Exit(1)
	}
	outDir, err := genutils.OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}
	switch module {
	case "kube-apiserver":
		apiserver := apiservapp.NewAPIServerCommand(server.SetupSignalHandler())
		doc.GenMarkdownTree(apiserver, outDir)
	case "kube-controller-manager":
		controllermanager := cmapp.NewControllerManagerCommand()
		doc.GenMarkdownTree(controllermanager, outDir)
	case "cloud-controller-manager":
		cloudcontrollermanager := ccmapp.NewCloudControllerManagerCommand()
		doc.GenMarkdownTree(cloudcontrollermanager, outDir)
	case "kube-proxy":
		proxy := proxyapp.NewProxyCommand()
		doc.GenMarkdownTree(proxy, outDir)
	case "kube-scheduler":
		scheduler := schapp.NewSchedulerCommand()
		doc.GenMarkdownTree(scheduler, outDir)
	case "kubelet":
		kubelet := kubeletapp.NewKubeletCommand(server.SetupSignalHandler())
		doc.GenMarkdownTree(kubelet, outDir)
	case "kubeadm":
		pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
		kubeadm := kubeadmapp.NewKubeadmCommand(os.Stdin, os.Stdout, os.Stderr)
		doc.GenMarkdownTree(kubeadm, outDir)
		MarkdownPostProcessing(kubeadm, outDir, cleanupForInclude)
	default:
		fmt.Fprintf(os.Stderr, "Module %s is not supported", module)
		os.Exit(1)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
