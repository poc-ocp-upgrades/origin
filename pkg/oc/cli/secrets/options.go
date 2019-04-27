package secrets

import (
	"errors"
	"fmt"
	"os"
	"strings"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

type SecretOptions struct {
	TargetName	string
	SecretNames	[]string
	typeFlags	[]string
	Namespace	string
	BuilderFunc	func() *resource.Builder
	KubeClient	corev1client.CoreV1Interface
}

func (o *SecretOptions) Complete(f kcmdutil.Factory, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) < 2 {
		return errors.New("must have service account name and at least one secret name")
	}
	o.TargetName = args[0]
	o.SecretNames = args[1:]
	o.BuilderFunc = f.NewBuilder
	var err error
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	o.KubeClient, err = corev1client.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	o.Namespace, _, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	return nil
}
func (o SecretOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.TargetName) == 0 {
		return errors.New("service account name must be present")
	}
	if len(o.SecretNames) == 0 {
		return errors.New("secret name must be present")
	}
	if o.KubeClient == nil {
		return errors.New("KubeClient must be present")
	}
	for _, secretName := range o.SecretNames {
		if segs := strings.Split(secretName, "/"); len(segs) > 1 {
			if segs[0] != "secret" && segs[0] != "secrets" {
				return errors.New(fmt.Sprintf("expected resource of type secret, got %q", secretName))
			}
		}
	}
	return nil
}
func (o SecretOptions) GetServiceAccount() (*corev1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := o.BuilderFunc().WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).NamespaceParam(o.Namespace).ResourceNames("serviceaccounts", o.TargetName).SingleResourceType().Do()
	if r.Err() != nil {
		return nil, r.Err()
	}
	obj, err := r.Object()
	if err != nil {
		return nil, err
	}
	switch t := obj.(type) {
	case *corev1.ServiceAccount:
		return t, nil
	default:
		return nil, fmt.Errorf("unhandled object: %#v", t)
	}
}
func (o SecretOptions) GetSecretNames(secrets []*corev1.Secret) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := sets.String{}
	for _, secret := range secrets {
		names.Insert(parseSecretName(secret.Name))
	}
	return names
}
func parseSecretName(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	segs := strings.Split(name, "/")
	if len(segs) < 2 {
		return name
	}
	return segs[1]
}
func (o SecretOptions) GetMountSecretNames(serviceaccount *corev1.ServiceAccount) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := sets.String{}
	for _, secret := range serviceaccount.Secrets {
		names.Insert(secret.Name)
	}
	return names
}
func (o SecretOptions) GetPullSecretNames(serviceaccount *corev1.ServiceAccount) sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := sets.String{}
	for _, secret := range serviceaccount.ImagePullSecrets {
		names.Insert(secret.Name)
	}
	return names
}
func (o SecretOptions) GetSecrets(allowNonExisting bool) ([]*corev1.Secret, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	secrets := []*corev1.Secret{}
	hasNotFound := false
	for _, secretName := range o.SecretNames {
		r := o.BuilderFunc().WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).NamespaceParam(o.Namespace).ResourceNames("secrets", secretName).SingleResourceType().Do()
		if r.Err() != nil {
			return nil, false, r.Err()
		}
		obj, err := r.Object()
		if err != nil {
			if kerrors.IsNotFound(err) {
				fmt.Fprintf(os.Stderr, "secret %q not found\n", secretName)
				hasNotFound = true
				if allowNonExisting {
					obj = &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: secretName}}
				} else {
					continue
				}
			} else if err != nil {
				return nil, false, err
			}
		}
		switch t := obj.(type) {
		case *corev1.Secret:
			secrets = append(secrets, t)
		default:
			return nil, false, fmt.Errorf("unhandled object: %#v", t)
		}
	}
	if len(secrets) == 0 {
		return nil, false, errors.New("No valid secrets found")
	}
	return secrets, hasNotFound, nil
}
