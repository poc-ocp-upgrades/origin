package dockercredentials

import (
	goformat "fmt"
	"github.com/docker/distribution/registry/client/auth"
	"github.com/openshift/origin/pkg/image/registryclient"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/credentialprovider"
	"net/url"
	"os"
	goos "os"
	"path/filepath"
	"runtime"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var (
	emptyKeyring = &credentialprovider.BasicDockerKeyring{}
)

func NewLocal() auth.CredentialStore {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	keyring := &credentialprovider.BasicDockerKeyring{}
	keyring.Add(defaultClientDockerConfig())
	return &keyringCredentialStore{DockerKeyring: keyring, RefreshTokenStore: registryclient.NewRefreshTokenStore()}
}
func NewFromFile(path string) (auth.CredentialStore, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg, err := credentialprovider.ReadSpecificDockerConfigJsonFile(path)
	if err != nil {
		return nil, err
	}
	keyring := &credentialprovider.BasicDockerKeyring{}
	keyring.Add(cfg)
	return &keyringCredentialStore{DockerKeyring: keyring, RefreshTokenStore: registryclient.NewRefreshTokenStore()}, nil
}

type keyringCredentialStore struct {
	credentialprovider.DockerKeyring
	registryclient.RefreshTokenStore
}

func (s *keyringCredentialStore) Basic(url *url.URL) (string, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return BasicFromKeyring(s.DockerKeyring, url)
}
func BasicFromKeyring(keyring credentialprovider.DockerKeyring, target *url.URL) (string, string) {
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
			return BasicFromKeyring(keyring, &url.URL{Host: "index.docker.io", Path: "/v1"})
		}
		if value == "index.docker.io" {
			klog.V(5).Infof("Being asked for %s (%s), trying %s for legacy behavior", target, value, "docker.io")
			return BasicFromKeyring(keyring, &url.URL{Host: "docker.io"})
		}
		if (strings.HasSuffix(target.Host, ":443") && target.Scheme == "https") || (strings.HasSuffix(target.Host, ":80") && target.Scheme == "http") {
			host := strings.SplitN(target.Host, ":", 2)[0]
			klog.V(5).Infof("Being asked for %s (%s), trying %s without port", target, value, host)
			return BasicFromKeyring(keyring, &url.URL{Scheme: target.Scheme, Host: host, Path: target.Path})
		}
		klog.V(5).Infof("Unable to find a secret to match %s (%s)", target, value)
		return "", ""
	}
	klog.V(5).Infof("Found secret to match %s (%s): %s", target, value, configs[0].ServerAddress)
	return configs[0].Username, configs[0].Password
}
func defaultClientDockerConfig() credentialprovider.DockerConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg, err := credentialprovider.ReadDockerConfigJSONFile(defaultPathsForCredentials()); err == nil {
		return cfg
	}
	if cfg, err := credentialprovider.ReadDockercfgFile(defaultPathsForLegacyCredentials()); err == nil {
		return cfg
	}
	return credentialprovider.DockerConfig{}
}
func defaultPathsForCredentials() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if runtime.GOOS == "windows" {
		return []string{filepath.Join(os.Getenv("USERPROFILE"), ".docker")}
	}
	return []string{filepath.Join(os.Getenv("HOME"), ".docker")}
}
func defaultPathsForLegacyCredentials() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if runtime.GOOS == "windows" {
		return []string{os.Getenv("USERPROFILE")}
	}
	return []string{os.Getenv("HOME")}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
