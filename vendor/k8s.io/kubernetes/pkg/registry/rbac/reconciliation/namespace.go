package reconciliation

import (
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func tryEnsureNamespace(client corev1client.NamespaceInterface, namespace string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, getErr := client.Get(namespace, metav1.GetOptions{})
	if getErr == nil {
		return nil
	}
	if fatalGetErr := utilerrors.FilterOut(getErr, apierrors.IsNotFound, apierrors.IsForbidden); fatalGetErr != nil {
		return fatalGetErr
	}
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	_, createErr := client.Create(ns)
	return utilerrors.FilterOut(createErr, apierrors.IsAlreadyExists, apierrors.IsForbidden)
}
