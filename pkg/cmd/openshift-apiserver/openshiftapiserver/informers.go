package openshiftapiserver

import (
	"time"
	authorizationv1client "github.com/openshift/client-go/authorization/clientset/versioned"
	authorizationv1informer "github.com/openshift/client-go/authorization/informers/externalversions"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned"
	imagev1informer "github.com/openshift/client-go/image/informers/externalversions"
	oauthv1client "github.com/openshift/client-go/oauth/clientset/versioned"
	oauthv1informer "github.com/openshift/client-go/oauth/informers/externalversions"
	quotaclient "github.com/openshift/client-go/quota/clientset/versioned"
	quotainformer "github.com/openshift/client-go/quota/informers/externalversions"
	routev1client "github.com/openshift/client-go/route/clientset/versioned"
	routev1informer "github.com/openshift/client-go/route/informers/externalversions"
	securityv1client "github.com/openshift/client-go/security/clientset/versioned"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions"
	userv1client "github.com/openshift/client-go/user/clientset/versioned"
	userv1informer "github.com/openshift/client-go/user/informers/externalversions"
	kexternalinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/rest"
)

type InformerHolder struct {
	kubernetesInformers	kexternalinformers.SharedInformerFactory
	authorizationInformers	authorizationv1informer.SharedInformerFactory
	imageInformers		imagev1informer.SharedInformerFactory
	oauthInformers		oauthv1informer.SharedInformerFactory
	quotaInformers		quotainformer.SharedInformerFactory
	routeInformers		routev1informer.SharedInformerFactory
	securityInformers	securityv1informer.SharedInformerFactory
	userInformers		userv1informer.SharedInformerFactory
}

func NewInformers(kubeInformers kexternalinformers.SharedInformerFactory, kubeClientConfig *rest.Config, loopbackClientConfig *rest.Config) (*InformerHolder, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	authorizationClient, err := authorizationv1client.NewForConfig(nonProtobufConfig(kubeClientConfig))
	if err != nil {
		return nil, err
	}
	imageClient, err := imagev1client.NewForConfig(loopbackClientConfig)
	if err != nil {
		return nil, err
	}
	oauthClient, err := oauthv1client.NewForConfig(loopbackClientConfig)
	if err != nil {
		return nil, err
	}
	quotaClient, err := quotaclient.NewForConfig(nonProtobufConfig(kubeClientConfig))
	if err != nil {
		return nil, err
	}
	routerClient, err := routev1client.NewForConfig(loopbackClientConfig)
	if err != nil {
		return nil, err
	}
	securityClient, err := securityv1client.NewForConfig(nonProtobufConfig(kubeClientConfig))
	if err != nil {
		return nil, err
	}
	userClient, err := userv1client.NewForConfig(loopbackClientConfig)
	if err != nil {
		return nil, err
	}
	const defaultInformerResyncPeriod = 10 * time.Minute
	return &InformerHolder{kubernetesInformers: kubeInformers, authorizationInformers: authorizationv1informer.NewSharedInformerFactory(authorizationClient, defaultInformerResyncPeriod), imageInformers: imagev1informer.NewSharedInformerFactory(imageClient, defaultInformerResyncPeriod), oauthInformers: oauthv1informer.NewSharedInformerFactory(oauthClient, defaultInformerResyncPeriod), quotaInformers: quotainformer.NewSharedInformerFactory(quotaClient, defaultInformerResyncPeriod), routeInformers: routev1informer.NewSharedInformerFactory(routerClient, defaultInformerResyncPeriod), securityInformers: securityv1informer.NewSharedInformerFactory(securityClient, defaultInformerResyncPeriod), userInformers: userv1informer.NewSharedInformerFactory(userClient, defaultInformerResyncPeriod)}, nil
}
func nonProtobufConfig(inConfig *rest.Config) *rest.Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	npConfig := rest.CopyConfig(inConfig)
	npConfig.ContentConfig.AcceptContentTypes = "application/json"
	npConfig.ContentConfig.ContentType = "application/json"
	return npConfig
}
func (i *InformerHolder) GetKubernetesInformers() kexternalinformers.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.kubernetesInformers
}
func (i *InformerHolder) GetOpenshiftAuthorizationInformers() authorizationv1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.authorizationInformers
}
func (i *InformerHolder) GetOpenshiftImageInformers() imagev1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.imageInformers
}
func (i *InformerHolder) GetOpenshiftOauthInformers() oauthv1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.oauthInformers
}
func (i *InformerHolder) GetOpenshiftQuotaInformers() quotainformer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.quotaInformers
}
func (i *InformerHolder) GetOpenshiftRouteInformers() routev1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.routeInformers
}
func (i *InformerHolder) GetOpenshiftSecurityInformers() securityv1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.securityInformers
}
func (i *InformerHolder) GetOpenshiftUserInformers() userv1informer.SharedInformerFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.userInformers
}
func (i *InformerHolder) Start(stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i.kubernetesInformers.Start(stopCh)
	i.authorizationInformers.Start(stopCh)
	i.imageInformers.Start(stopCh)
	i.oauthInformers.Start(stopCh)
	i.quotaInformers.Start(stopCh)
	i.routeInformers.Start(stopCh)
	i.securityInformers.Start(stopCh)
	i.userInformers.Start(stopCh)
}
