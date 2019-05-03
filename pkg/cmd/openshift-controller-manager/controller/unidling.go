package controller

import (
	appsclient "github.com/openshift/client-go/apps/clientset/versioned"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	unidlingcontroller "github.com/openshift/origin/pkg/unidling/controller"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/scale"
	"time"
)

func RunUnidlingController(ctx *ControllerContext) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resyncPeriod := 2 * time.Hour
	clientConfig := ctx.ClientBuilder.ConfigOrDie(bootstrappolicy.InfraUnidlingControllerServiceAccountName)
	appsClient, err := appsclient.NewForConfig(clientConfig)
	if err != nil {
		return false, err
	}
	scaleKindResolver := scale.NewDiscoveryScaleKindResolver(appsClient.Discovery())
	scaleClient, err := scale.NewForConfig(clientConfig, ctx.RestMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	if err != nil {
		return false, err
	}
	coreClient := ctx.ClientBuilder.ClientOrDie(bootstrappolicy.InfraUnidlingControllerServiceAccountName).CoreV1()
	controller := unidlingcontroller.NewUnidlingController(scaleClient, ctx.RestMapper, coreClient, coreClient, appsClient.AppsV1(), coreClient, resyncPeriod)
	go controller.Run(ctx.Stop)
	return true, nil
}
