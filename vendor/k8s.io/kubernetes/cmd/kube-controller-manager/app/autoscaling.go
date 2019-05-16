package app

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/scale"
	"k8s.io/kubernetes/pkg/controller/podautoscaler"
	"k8s.io/kubernetes/pkg/controller/podautoscaler/metrics"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/custom_metrics"
	"k8s.io/metrics/pkg/client/external_metrics"
	"net/http"
)

func startHPAController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "autoscaling", Version: "v1", Resource: "horizontalpodautoscalers"}] {
		return nil, false, nil
	}
	if ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerUseRESTClients {
		return startHPAControllerWithRESTClient(ctx)
	}
	return startHPAControllerWithLegacyClient(ctx)
}
func startHPAControllerWithRESTClient(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clientConfig := ctx.ClientBuilder.ConfigOrDie("horizontal-pod-autoscaler")
	hpaClient := ctx.ClientBuilder.ClientOrDie("horizontal-pod-autoscaler")
	apiVersionsGetter := custom_metrics.NewAvailableAPIsGetter(hpaClient.Discovery())
	go custom_metrics.PeriodicallyInvalidate(apiVersionsGetter, ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerSyncPeriod.Duration, ctx.Stop)
	metricsClient := metrics.NewRESTMetricsClient(resourceclient.NewForConfigOrDie(clientConfig), custom_metrics.NewForConfig(clientConfig, ctx.RESTMapper, apiVersionsGetter), external_metrics.NewForConfigOrDie(clientConfig))
	return startHPAControllerWithMetricsClient(ctx, metricsClient)
}
func startHPAControllerWithLegacyClient(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpaClient := ctx.ClientBuilder.ClientOrDie("horizontal-pod-autoscaler")
	metricsClient := metrics.NewHeapsterMetricsClient(hpaClient, metrics.DefaultHeapsterNamespace, metrics.DefaultHeapsterScheme, metrics.DefaultHeapsterService, metrics.DefaultHeapsterPort)
	return startHPAControllerWithMetricsClient(ctx, metricsClient)
}
func startHPAControllerWithMetricsClient(ctx ControllerContext, metricsClient metrics.MetricsClient) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hpaClient := ctx.ClientBuilder.ClientOrDie("horizontal-pod-autoscaler")
	hpaClientConfig := ctx.ClientBuilder.ConfigOrDie("horizontal-pod-autoscaler")
	scaleKindResolver := scale.NewDiscoveryScaleKindResolver(hpaClient.Discovery())
	scaleClient, err := scale.NewForConfig(hpaClientConfig, ctx.RESTMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	if err != nil {
		return nil, false, err
	}
	go podautoscaler.NewHorizontalController(hpaClient.CoreV1(), scaleClient, hpaClient.AutoscalingV1(), ctx.RESTMapper, metricsClient, ctx.InformerFactory.Autoscaling().V1().HorizontalPodAutoscalers(), ctx.InformerFactory.Core().V1().Pods(), ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerSyncPeriod.Duration, ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerDownscaleStabilizationWindow.Duration, ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerTolerance, ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerCPUInitializationPeriod.Duration, ctx.ComponentConfig.HPAController.HorizontalPodAutoscalerInitialReadinessDelay.Duration).Run(ctx.Stop)
	return nil, true, nil
}
