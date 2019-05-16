package importer

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/image/registryclient"
	kapiv1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/credentialprovider"
	credentialprovidersecrets "k8s.io/kubernetes/pkg/credentialprovider/secrets"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"sync"
	gotime "time"
)

var (
	emptyKeyring = &credentialprovider.BasicDockerKeyring{}
)

func NewCredentialsForSecrets(secrets []kapiv1.Secret) *SecretCredentialStore {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &SecretCredentialStore{secrets: secrets, RefreshTokenStore: registryclient.NewRefreshTokenStore()}
}
func NewLazyCredentialsForSecrets(secretsFn func() ([]kapiv1.Secret, error)) *SecretCredentialStore {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &SecretCredentialStore{secretsFn: secretsFn, RefreshTokenStore: registryclient.NewRefreshTokenStore()}
}

type SecretCredentialStore struct {
	lock      sync.Mutex
	secrets   []kapiv1.Secret
	secretsFn func() ([]kapiv1.Secret, error)
	err       error
	keyring   credentialprovider.DockerKeyring
	registryclient.RefreshTokenStore
}

func (s *SecretCredentialStore) Basic(url *url.URL) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return basicCredentialsFromKeyring(s.init(), url)
}
func (s *SecretCredentialStore) Err() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.err
}
func (s *SecretCredentialStore) init() credentialprovider.DockerKeyring {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.keyring != nil {
		return s.keyring
	}
	if s.secrets == nil {
		if s.secretsFn != nil {
			s.secrets, s.err = s.secretsFn()
		}
	}
	keyring, err := credentialprovidersecrets.MakeDockerKeyring(s.secrets, emptyKeyring)
	if err != nil {
		klog.V(5).Infof("Loading keyring failed for credential store: %v", err)
		s.err = err
		keyring = emptyKeyring
	}
	s.keyring = keyring
	return keyring
}
func basicCredentialsFromKeyring(keyring credentialprovider.DockerKeyring, target *url.URL) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var value string
	if len(target.Scheme) == 0 || target.Scheme == "https" {
		value = target.Host + target.Path
	} else {
		if !strings.Contains(target.Host, ":") {
			value = target.Host + ":80" + target.Path
		} else {
			value = target.Host + target.Path
		}
	}
	pathWithSlash := target.Path + "/"
	if strings.HasPrefix(pathWithSlash, "/v1/") || strings.HasPrefix(pathWithSlash, "/v2/") {
		value = target.Host + target.Path[3:]
	}
	configs, found := keyring.Lookup(value)
	if !found || len(configs) == 0 {
		if value == "auth.docker.io/token" {
			klog.V(5).Infof("Being asked for %s (%s), trying %s for legacy behavior", target, value, "index.docker.io/v1")
			return basicCredentialsFromKeyring(keyring, &url.URL{Host: "index.docker.io", Path: "/v1"})
		}
		if value == "index.docker.io" {
			klog.V(5).Infof("Being asked for %s (%s), trying %s for legacy behavior", target, value, "docker.io")
			return basicCredentialsFromKeyring(keyring, &url.URL{Host: "docker.io"})
		}
		if (strings.HasSuffix(target.Host, ":443") && target.Scheme == "https") || (strings.HasSuffix(target.Host, ":80") && target.Scheme == "http") {
			host := strings.SplitN(target.Host, ":", 2)[0]
			klog.V(5).Infof("Being asked for %s (%s), trying %s without port", target, value, host)
			return basicCredentialsFromKeyring(keyring, &url.URL{Scheme: target.Scheme, Host: host, Path: target.Path})
		}
		klog.V(5).Infof("Unable to find a secret to match %s (%s)", target, value)
		return "", ""
	}
	klog.V(5).Infof("Found secret to match %s (%s): %s", target, value, configs[0].ServerAddress)
	return configs[0].Username, configs[0].Password
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
