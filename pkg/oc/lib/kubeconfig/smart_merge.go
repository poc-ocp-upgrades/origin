package kubeconfig

import (
	"reflect"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/openshift/origin/pkg/client/config"
)

func getClusterNicknameFromConfig(clientCfg *restclient.Config) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return config.GetClusterNicknameFromURL(clientCfg.Host)
}
func getUserNicknameFromConfig(clientCfg *restclient.Config) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userPartOfNick, err := getUserPartOfNickname(clientCfg)
	if err != nil {
		return "", err
	}
	clusterNick, err := getClusterNicknameFromConfig(clientCfg)
	if err != nil {
		return "", err
	}
	return userPartOfNick + "/" + clusterNick, nil
}
func getUserPartOfNickname(clientCfg *restclient.Config) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userClient, err := userv1typedclient.NewForConfig(clientCfg)
	if err != nil {
		return "", err
	}
	userInfo, err := userClient.Users().Get("~", metav1.GetOptions{})
	if kerrors.IsNotFound(err) || kerrors.IsForbidden(err) {
		switch {
		case len(clientCfg.BearerToken) > 0:
			userInfo.Name = clientCfg.BearerToken
		case len(clientCfg.Username) > 0:
			userInfo.Name = clientCfg.Username
		}
	} else if err != nil {
		return "", err
	}
	return userInfo.Name, nil
}
func getContextNicknameFromConfig(namespace string, clientCfg *restclient.Config) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userPartOfNick, err := getUserPartOfNickname(clientCfg)
	if err != nil {
		return "", err
	}
	clusterNick, err := getClusterNicknameFromConfig(clientCfg)
	if err != nil {
		return "", err
	}
	return namespace + "/" + clusterNick + "/" + userPartOfNick, nil
}
func CreateConfig(namespace string, clientCfg *restclient.Config) (*clientcmdapi.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterNick, err := getClusterNicknameFromConfig(clientCfg)
	if err != nil {
		return nil, err
	}
	userNick, err := getUserNicknameFromConfig(clientCfg)
	if err != nil {
		return nil, err
	}
	contextNick, err := getContextNicknameFromConfig(namespace, clientCfg)
	if err != nil {
		return nil, err
	}
	config := clientcmdapi.NewConfig()
	credentials := clientcmdapi.NewAuthInfo()
	credentials.Token = clientCfg.BearerToken
	credentials.ClientCertificate = clientCfg.TLSClientConfig.CertFile
	if len(credentials.ClientCertificate) == 0 {
		credentials.ClientCertificateData = clientCfg.TLSClientConfig.CertData
	}
	credentials.ClientKey = clientCfg.TLSClientConfig.KeyFile
	if len(credentials.ClientKey) == 0 {
		credentials.ClientKeyData = clientCfg.TLSClientConfig.KeyData
	}
	config.AuthInfos[userNick] = credentials
	cluster := clientcmdapi.NewCluster()
	cluster.Server = clientCfg.Host
	cluster.CertificateAuthority = clientCfg.CAFile
	if len(cluster.CertificateAuthority) == 0 {
		cluster.CertificateAuthorityData = clientCfg.CAData
	}
	cluster.InsecureSkipTLSVerify = clientCfg.Insecure
	config.Clusters[clusterNick] = cluster
	context := clientcmdapi.NewContext()
	context.Cluster = clusterNick
	context.AuthInfo = userNick
	context.Namespace = namespace
	config.Contexts[contextNick] = context
	config.CurrentContext = contextNick
	return config, nil
}
func MergeConfig(startingConfig, addition clientcmdapi.Config) (*clientcmdapi.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := startingConfig
	for requestedKey, value := range addition.Clusters {
		ret.Clusters[requestedKey] = value
	}
	for requestedKey, value := range addition.AuthInfos {
		ret.AuthInfos[requestedKey] = value
	}
	requestedContextNamesToActualContextNames := map[string]string{}
	for requestedKey, newContext := range addition.Contexts {
		actualContext := clientcmdapi.NewContext()
		actualContext.AuthInfo = newContext.AuthInfo
		actualContext.Cluster = newContext.Cluster
		actualContext.Namespace = newContext.Namespace
		actualContext.Extensions = newContext.Extensions
		if existingName := findExistingContextName(startingConfig, *actualContext); len(existingName) > 0 {
			requestedContextNamesToActualContextNames[requestedKey] = existingName
			continue
		}
		requestedContextNamesToActualContextNames[requestedKey] = requestedKey
		ret.Contexts[requestedKey] = actualContext
	}
	if len(addition.CurrentContext) > 0 {
		if newCurrentContext, exists := requestedContextNamesToActualContextNames[addition.CurrentContext]; exists {
			ret.CurrentContext = newCurrentContext
		} else {
			ret.CurrentContext = addition.CurrentContext
		}
	}
	return &ret, nil
}
func findExistingContextName(haystack clientcmdapi.Config, needle clientcmdapi.Context) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for key, context := range haystack.Contexts {
		context.LocationOfOrigin = ""
		if reflect.DeepEqual(context, needle) {
			return key
		}
	}
	return ""
}
