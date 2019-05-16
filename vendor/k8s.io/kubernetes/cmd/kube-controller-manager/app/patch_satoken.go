package app

import (
	"fmt"
	"io/ioutil"
	certutil "k8s.io/client-go/util/cert"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"path/filepath"
)

var applyOpenShiftServiceServingCertCA = func(in serviceaccountcontroller.TokensControllerOptions) serviceaccountcontroller.TokensControllerOptions {
	return in
}

func applyOpenShiftServiceServingCertCAFunc(openshiftConfigBase string, openshiftConfig map[string]interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serviceServingCertCAFilename := getServiceServingCertCAFilename(openshiftConfig)
	if len(serviceServingCertCAFilename) == 0 {
		return nil
	}
	resolvePath(&serviceServingCertCAFilename, openshiftConfigBase)
	serviceServingCA, err := ioutil.ReadFile(serviceServingCertCAFilename)
	if err != nil {
		return fmt.Errorf("error reading ca file for Service Serving Certificate Signer: %s: %v", serviceServingCertCAFilename, err)
	}
	if _, err := certutil.ParseCertsPEM(serviceServingCA); err != nil {
		return fmt.Errorf("error parsing ca file for Service Serving Certificate Signer: %s: %v", serviceServingCertCAFilename, err)
	}
	applyOpenShiftServiceServingCertCA = func(controllerOptions serviceaccountcontroller.TokensControllerOptions) serviceaccountcontroller.TokensControllerOptions {
		if len(serviceServingCA) == 0 {
			return controllerOptions
		}
		if len(controllerOptions.RootCA) > 0 {
			controllerOptions.ServiceServingCA = append(controllerOptions.ServiceServingCA, controllerOptions.RootCA...)
			controllerOptions.ServiceServingCA = append(controllerOptions.ServiceServingCA, []byte("\n")...)
		}
		controllerOptions.ServiceServingCA = append(controllerOptions.ServiceServingCA, serviceServingCA...)
		return controllerOptions
	}
	return nil
}
func getServiceServingCertCAFilename(config map[string]interface{}) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	controllerConfig, ok := config["controllerConfig"]
	if !ok {
		sscConfig, ok := config["serviceServingCert"]
		if !ok {
			return ""
		}
		sscConfigMap := sscConfig.(map[string]interface{})
		return sscConfigMap["certFile"].(string)
	}
	controllerConfigMap := controllerConfig.(map[string]interface{})
	sscConfig, ok := controllerConfigMap["serviceServingCert"]
	if !ok {
		return ""
	}
	sscConfigMap := sscConfig.(map[string]interface{})
	signerConfig, ok := sscConfigMap["signer"]
	if !ok {
		return ""
	}
	signerConfigMap := signerConfig.(map[string]interface{})
	return signerConfigMap["certFile"].(string)
}
func resolvePath(ref *string, base string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(*ref) > 0 {
		if !filepath.IsAbs(*ref) {
			*ref = filepath.Join(base, *ref)
		}
	}
	return nil
}
