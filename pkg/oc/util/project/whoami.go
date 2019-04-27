package project

import (
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	userv1 "github.com/openshift/api/user/v1"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

func WhoAmI(clientConfig *restclient.Config) (*userv1.User, error) {
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
	client, err := userv1typedclient.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	me, err := client.Users().Get("~", metav1.GetOptions{})
	if kerrors.IsNotFound(err) || kerrors.IsForbidden(err) {
		switch {
		case len(clientConfig.BearerToken) > 0:
			return &userv1.User{ObjectMeta: metav1.ObjectMeta{Name: clientConfig.BearerToken}}, nil
		case len(clientConfig.Username) > 0:
			return &userv1.User{ObjectMeta: metav1.ObjectMeta{Name: clientConfig.Username}}, nil
		}
	}
	if err != nil {
		return nil, err
	}
	return me, nil
}
