package app

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/klog"
	kubeoptions "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	"k8s.io/kubernetes/pkg/controller/certificates/approver"
	"k8s.io/kubernetes/pkg/controller/certificates/cleaner"
	"k8s.io/kubernetes/pkg/controller/certificates/rootcacertpublisher"
	"k8s.io/kubernetes/pkg/controller/certificates/signer"
	"k8s.io/kubernetes/pkg/features"
	"net/http"
	"os"
)

func startCSRSigningController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1beta1", Resource: "certificatesigningrequests"}] {
		return nil, false, nil
	}
	if ctx.ComponentConfig.CSRSigningController.ClusterSigningCertFile == "" || ctx.ComponentConfig.CSRSigningController.ClusterSigningKeyFile == "" {
		return nil, false, nil
	}
	var keyFileExists, keyUsesDefault, certFileExists, certUsesDefault bool
	_, err := os.Stat(ctx.ComponentConfig.CSRSigningController.ClusterSigningCertFile)
	certFileExists = !os.IsNotExist(err)
	certUsesDefault = (ctx.ComponentConfig.CSRSigningController.ClusterSigningCertFile == kubeoptions.DefaultClusterSigningCertFile)
	_, err = os.Stat(ctx.ComponentConfig.CSRSigningController.ClusterSigningKeyFile)
	keyFileExists = !os.IsNotExist(err)
	keyUsesDefault = (ctx.ComponentConfig.CSRSigningController.ClusterSigningKeyFile == kubeoptions.DefaultClusterSigningKeyFile)
	switch {
	case (keyFileExists && keyUsesDefault) || (certFileExists && certUsesDefault):
		klog.Warningf("You might be using flag defaulting for --cluster-signing-cert-file and" + " --cluster-signing-key-file. These defaults are deprecated and will be removed" + " in a subsequent release. Please pass these options explicitly.")
	case (!keyFileExists && keyUsesDefault) && (!certFileExists && certUsesDefault):
		return nil, false, nil
	default:
	}
	c := ctx.ClientBuilder.ClientOrDie("certificate-controller")
	signer, err := signer.NewCSRSigningController(c, ctx.InformerFactory.Certificates().V1beta1().CertificateSigningRequests(), ctx.ComponentConfig.CSRSigningController.ClusterSigningCertFile, ctx.ComponentConfig.CSRSigningController.ClusterSigningKeyFile, ctx.ComponentConfig.CSRSigningController.ClusterSigningDuration.Duration)
	if err != nil {
		return nil, false, fmt.Errorf("failed to start certificate controller: %v", err)
	}
	go signer.Run(1, ctx.Stop)
	return nil, true, nil
}
func startCSRApprovingController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1beta1", Resource: "certificatesigningrequests"}] {
		return nil, false, nil
	}
	approver := approver.NewCSRApprovingController(ctx.ClientBuilder.ClientOrDie("certificate-controller"), ctx.InformerFactory.Certificates().V1beta1().CertificateSigningRequests())
	go approver.Run(1, ctx.Stop)
	return nil, true, nil
}
func startCSRCleanerController(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cleaner := cleaner.NewCSRCleanerController(ctx.ClientBuilder.ClientOrDie("certificate-controller").CertificatesV1beta1().CertificateSigningRequests(), ctx.InformerFactory.Certificates().V1beta1().CertificateSigningRequests())
	go cleaner.Run(1, ctx.Stop)
	return nil, true, nil
}
func startRootCACertPublisher(ctx ControllerContext) (http.Handler, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.BoundServiceAccountTokenVolume) {
		return nil, false, nil
	}
	var (
		rootCA []byte
		err    error
	)
	if ctx.ComponentConfig.SAController.RootCAFile != "" {
		if rootCA, err = readCA(ctx.ComponentConfig.SAController.RootCAFile); err != nil {
			return nil, true, fmt.Errorf("error parsing root-ca-file at %s: %v", ctx.ComponentConfig.SAController.RootCAFile, err)
		}
	} else {
		rootCA = ctx.ClientBuilder.ConfigOrDie("root-ca-cert-publisher").CAData
	}
	sac, err := rootcacertpublisher.NewPublisher(ctx.InformerFactory.Core().V1().ConfigMaps(), ctx.InformerFactory.Core().V1().Namespaces(), ctx.ClientBuilder.ClientOrDie("root-ca-cert-publisher"), rootCA)
	if err != nil {
		return nil, true, fmt.Errorf("error creating root CA certificate publisher: %v", err)
	}
	go sac.Run(1, ctx.Stop)
	return nil, true, nil
}
