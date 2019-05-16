package app

import (
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/controller"
)

func createPVRecyclerSA(openshiftConfig string, clientBuilder controller.ControllerClientBuilder) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(openshiftConfig) == 0 {
		return nil
	}
	coreClient, err := clientBuilder.Client("pv-recycler-controller-creator")
	if err != nil {
		return err
	}
	if _, err := coreClient.CoreV1().Namespaces().Create(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "openshift-infra"}}); err != nil {
	}
	_, err = coreClient.CoreV1().ServiceAccounts("openshift-infra").Create(&v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: "openshift-infra", Name: "pv-recycler-controller"}})
	if apierrors.IsAlreadyExists(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
