package app

import (
	"fmt"
	"k8s.io/kubernetes/pkg/controller/bootstrap"
	"net/http"
)

func startBootstrapSignerController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bsc, err := bootstrap.NewBootstrapSigner(ctx.ClientBuilder.ClientOrDie("bootstrap-signer"), ctx.InformerFactory.Core().V1().Secrets(), ctx.InformerFactory.Core().V1().ConfigMaps(), bootstrap.DefaultBootstrapSignerOptions())
	if err != nil {
		return nil, true, fmt.Errorf("error creating BootstrapSigner controller: %v", err)
	}
	go bsc.Run(ctx.Stop)
	return nil, true, nil
}
func startTokenCleanerController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tcc, err := bootstrap.NewTokenCleaner(ctx.ClientBuilder.ClientOrDie("token-cleaner"), ctx.InformerFactory.Core().V1().Secrets(), bootstrap.DefaultTokenCleanerOptions())
	if err != nil {
		return nil, true, fmt.Errorf("error creating TokenCleaner controller: %v", err)
	}
	go tcc.Run(ctx.Stop)
	return nil, true, nil
}
