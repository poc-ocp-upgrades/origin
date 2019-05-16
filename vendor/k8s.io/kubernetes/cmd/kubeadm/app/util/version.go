package util

import (
	"errors"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"io/ioutil"
	netutil "k8s.io/apimachinery/pkg/util/net"
	versionutil "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	pkgversion "k8s.io/kubernetes/pkg/version"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	getReleaseVersionTimeout = time.Duration(10 * time.Second)
)

var (
	kubeReleaseBucketURL  = "https://dl.k8s.io"
	kubeReleaseRegex      = regexp.MustCompile(`^v?(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)([-0-9a-zA-Z_\.+]*)?$`)
	kubeReleaseLabelRegex = regexp.MustCompile(`^[[:lower:]]+(-[-\w_\.]+)?$`)
	kubeBucketPrefixes    = regexp.MustCompile(`^((release|ci|ci-cross)/)?([-\w_\.+]+)$`)
)

func KubernetesReleaseVersion(version string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ver := normalizedBuildVersion(version)
	if len(ver) != 0 {
		return ver, nil
	}
	bucketURL, versionLabel, err := splitVersion(version)
	if err != nil {
		return "", err
	}
	ver = normalizedBuildVersion(versionLabel)
	if len(ver) != 0 {
		return ver, nil
	}
	if kubeReleaseLabelRegex.MatchString(versionLabel) {
		var clientVersion string
		clientVersion, _ = kubeadmVersion(pkgversion.Get().String())
		url := fmt.Sprintf("%s/%s.txt", bucketURL, versionLabel)
		body, err := fetchFromURL(url, getReleaseVersionTimeout)
		if err != nil {
			if body != "" {
				return "", err
			}
			klog.Infof("could not fetch a Kubernetes version from the internet: %v", err)
			klog.Infof("falling back to the local client version: %s", clientVersion)
			return KubernetesReleaseVersion(clientVersion)
		}
		body, err = validateStableVersion(body, clientVersion)
		if err != nil {
			return "", err
		}
		return KubernetesReleaseVersion(body)
	}
	return "", pkgerrors.Errorf("version %q doesn't match patterns for neither semantic version nor labels (stable, latest, ...)", version)
}
func KubernetesVersionToImageTag(version string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allowed := regexp.MustCompile(`[^-a-zA-Z0-9_\.]`)
	return allowed.ReplaceAllString(version, "_")
}
func KubernetesIsCIVersion(version string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subs := kubeBucketPrefixes.FindAllStringSubmatch(version, 1)
	if len(subs) == 1 && len(subs[0]) == 4 && strings.HasPrefix(subs[0][2], "ci") {
		return true
	}
	return false
}
func normalizedBuildVersion(version string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeReleaseRegex.MatchString(version) {
		if strings.HasPrefix(version, "v") {
			return version
		}
		return "v" + version
	}
	return ""
}
func splitVersion(version string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var urlSuffix string
	subs := kubeBucketPrefixes.FindAllStringSubmatch(version, 1)
	if len(subs) != 1 || len(subs[0]) != 4 {
		return "", "", pkgerrors.Errorf("invalid version %q", version)
	}
	switch {
	case strings.HasPrefix(subs[0][2], "ci"):
		urlSuffix = subs[0][2]
	default:
		urlSuffix = "release"
	}
	url := fmt.Sprintf("%s/%s", kubeReleaseBucketURL, urlSuffix)
	return url, subs[0][3], nil
}
func fetchFromURL(url string, timeout time.Duration) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(2).Infof("fetching Kubernetes version from URL: %s", url)
	client := &http.Client{Timeout: timeout, Transport: netutil.SetOldTransportDefaults(&http.Transport{})}
	resp, err := client.Get(url)
	if err != nil {
		return "", pkgerrors.Errorf("unable to get URL %q: %s", url, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", pkgerrors.Errorf("unable to read content of URL %q: %s", url, err.Error())
	}
	bodyString := strings.TrimSpace(string(body))
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("unable to fetch file. URL: %q, status: %v", url, resp.Status)
		return bodyString, errors.New(msg)
	}
	return bodyString, nil
}
func kubeadmVersion(info string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v, err := versionutil.ParseSemantic(info)
	if err != nil {
		return "", pkgerrors.Wrap(err, "kubeadm version error")
	}
	pre := v.PreRelease()
	patch := v.Patch()
	if len(pre) > 0 {
		if patch > 0 {
			patch = patch - 1
			pre = ""
		} else {
			split := strings.Split(pre, ".")
			if len(split) > 2 {
				pre = split[0] + "." + split[1]
			} else if len(split) < 2 {
				pre = split[0] + ".0"
			}
			pre = "-" + pre
		}
	}
	vStr := fmt.Sprintf("v%d.%d.%d%s", v.Major(), v.Minor(), patch, pre)
	return vStr, nil
}
func validateStableVersion(remoteVersion, clientVersion string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if clientVersion == "" {
		klog.Infof("could not obtain client version; using remote version: %s", remoteVersion)
		return remoteVersion, nil
	}
	verRemote, err := versionutil.ParseGeneric(remoteVersion)
	if err != nil {
		return "", pkgerrors.Wrap(err, "remote version error")
	}
	verClient, err := versionutil.ParseGeneric(clientVersion)
	if err != nil {
		return "", pkgerrors.Wrap(err, "client version error")
	}
	if verClient.Major() < verRemote.Major() || (verClient.Major() == verRemote.Major()) && verClient.Minor() < verRemote.Minor() {
		estimatedRelease := fmt.Sprintf("stable-%d.%d", verClient.Major(), verClient.Minor())
		klog.Infof("remote version is much newer: %s; falling back to: %s", remoteVersion, estimatedRelease)
		return estimatedRelease, nil
	}
	return remoteVersion, nil
}
