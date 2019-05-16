package controller

import (
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	serviceaccountcontrollers "github.com/openshift/origin/pkg/serviceaccounts/controllers"
	kapiv1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	sacontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
)

func RunServiceAccountController(ctx *ControllerContext) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(ctx.OpenshiftControllerConfig.ServiceAccount.ManagedNames) == 0 {
		klog.Infof("Skipped starting Service Account Manager, no managed names specified")
		return false, nil
	}
	options := sacontroller.DefaultServiceAccountsControllerOptions()
	options.ServiceAccounts = []kapiv1.ServiceAccount{}
	for _, saName := range ctx.OpenshiftControllerConfig.ServiceAccount.ManagedNames {
		if saName == "default" {
			continue
		}
		sa := kapiv1.ServiceAccount{}
		sa.Name = saName
		options.ServiceAccounts = append(options.ServiceAccounts, sa)
	}
	controller, err := sacontroller.NewServiceAccountsController(ctx.KubernetesInformers.Core().V1().ServiceAccounts(), ctx.KubernetesInformers.Core().V1().Namespaces(), ctx.ClientBuilder.ClientOrDie(bootstrappolicy.InfraServiceAccountControllerServiceAccountName), options)
	if err != nil {
		return true, nil
	}
	go controller.Run(3, ctx.Stop)
	return true, nil
}
func RunServiceAccountPullSecretsController(ctx *ControllerContext) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kc := ctx.ClientBuilder.ClientOrDie(bootstrappolicy.InfraServiceAccountPullSecretsControllerServiceAccountName)
	go serviceaccountcontrollers.NewDockercfgDeletedController(ctx.KubernetesInformers.Core().V1().Secrets(), kc, serviceaccountcontrollers.DockercfgDeletedControllerOptions{}).Run(ctx.Stop)
	go serviceaccountcontrollers.NewDockercfgTokenDeletedController(ctx.KubernetesInformers.Core().V1().Secrets(), kc, serviceaccountcontrollers.DockercfgTokenDeletedControllerOptions{}).Run(ctx.Stop)
	dockerURLsInitialized := make(chan struct{})
	dockercfgController := serviceaccountcontrollers.NewDockercfgController(ctx.KubernetesInformers.Core().V1().ServiceAccounts(), ctx.KubernetesInformers.Core().V1().Secrets(), kc, serviceaccountcontrollers.DockercfgControllerOptions{DockerURLsInitialized: dockerURLsInitialized})
	go dockercfgController.Run(5, ctx.Stop)
	dockerRegistryControllerOptions := serviceaccountcontrollers.DockerRegistryServiceControllerOptions{DockercfgController: dockercfgController, DockerURLsInitialized: dockerURLsInitialized, ClusterDNSSuffix: "cluster.local", AdditionalRegistryURLs: ctx.OpenshiftControllerConfig.DockerPullSecret.RegistryURLs}
	go serviceaccountcontrollers.NewDockerRegistryServiceController(ctx.KubernetesInformers.Core().V1().Secrets(), ctx.KubernetesInformers.Core().V1().Services(), kc, dockerRegistryControllerOptions).Run(10, ctx.Stop)
	return true, nil
}
