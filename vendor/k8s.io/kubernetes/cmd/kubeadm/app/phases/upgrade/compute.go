package upgrade

import (
	"fmt"
	goformat "fmt"
	versionutil "k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Upgrade struct {
	Description string
	Before      ClusterState
	After       ClusterState
}

func (u *Upgrade) CanUpgradeKubelets() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(u.Before.KubeletVersions) > 1 {
		return true
	}
	if len(u.Before.KubeletVersions) == 0 {
		return false
	}
	_, sameVersionFound := u.Before.KubeletVersions[u.After.KubeVersion]
	return !sameVersionFound
}
func (u *Upgrade) CanUpgradeEtcd() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return u.Before.EtcdVersion != u.After.EtcdVersion
}

type ClusterState struct {
	KubeVersion     string
	DNSType         kubeadmapi.DNSAddOnType
	DNSVersion      string
	KubeadmVersion  string
	KubeletVersions map[string]uint16
	EtcdVersion     string
}

func GetAvailableUpgrades(versionGetterImpl VersionGetter, experimentalUpgradesAllowed, rcUpgradesAllowed bool, etcdClient etcdutil.ClusterInterrogator, dnsType kubeadmapi.DNSAddOnType, client clientset.Interface) ([]Upgrade, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[upgrade] Fetching available versions to upgrade to")
	upgrades := []Upgrade{}
	clusterVersionStr, clusterVersion, err := versionGetterImpl.ClusterVersion()
	if err != nil {
		return upgrades, err
	}
	kubeadmVersionStr, kubeadmVersion, err := versionGetterImpl.KubeadmVersion()
	if err != nil {
		return upgrades, err
	}
	stableVersionStr, stableVersion, err := versionGetterImpl.VersionFromCILabel("stable", "stable version")
	if err != nil {
		fmt.Printf("[upgrade/versions] WARNING: %v\n", err)
		fmt.Println("[upgrade/versions] WARNING: Falling back to current kubeadm version as latest stable version")
		stableVersionStr, stableVersion = kubeadmVersionStr, kubeadmVersion
	}
	kubeletVersions, err := versionGetterImpl.KubeletVersions()
	if err != nil {
		return upgrades, err
	}
	etcdVersion, err := etcdClient.GetVersion()
	if err != nil {
		return upgrades, err
	}
	currentDNSType, dnsVersion, err := dns.DeployedDNSAddon(client)
	if err != nil {
		return nil, err
	}
	beforeState := ClusterState{KubeVersion: clusterVersionStr, DNSType: currentDNSType, DNSVersion: dnsVersion, KubeadmVersion: kubeadmVersionStr, KubeletVersions: kubeletVersions, EtcdVersion: etcdVersion}
	canDoMinorUpgrade := clusterVersion.LessThan(stableVersion)
	if patchVersionBranchExists(clusterVersion, stableVersion) {
		currentBranch := getBranchFromVersion(clusterVersionStr)
		versionLabel := fmt.Sprintf("stable-%s", currentBranch)
		description := fmt.Sprintf("version in the v%s series", currentBranch)
		patchVersionStr, patchVersion, err := versionGetterImpl.VersionFromCILabel(versionLabel, description)
		if err != nil {
			fmt.Printf("[upgrade/versions] WARNING: %v\n", err)
		} else {
			canDoMinorUpgrade = minorUpgradePossibleWithPatchRelease(stableVersion, patchVersion)
			if patchUpgradePossible(clusterVersion, patchVersion) {
				newKubeadmVer := patchVersionStr
				if kubeadmVersion.AtLeast(patchVersion) {
					newKubeadmVer = kubeadmVersionStr
				}
				upgrades = append(upgrades, Upgrade{Description: description, Before: beforeState, After: ClusterState{KubeVersion: patchVersionStr, DNSType: dnsType, DNSVersion: kubeadmconstants.GetDNSVersion(dnsType), KubeadmVersion: newKubeadmVer, EtcdVersion: getSuggestedEtcdVersion(patchVersionStr)}})
			}
		}
	}
	if canDoMinorUpgrade {
		upgrades = append(upgrades, Upgrade{Description: "stable version", Before: beforeState, After: ClusterState{KubeVersion: stableVersionStr, DNSType: dnsType, DNSVersion: kubeadmconstants.GetDNSVersion(dnsType), KubeadmVersion: stableVersionStr, EtcdVersion: getSuggestedEtcdVersion(stableVersionStr)}})
	}
	if experimentalUpgradesAllowed || rcUpgradesAllowed {
		latestVersionStr, latestVersion, err := versionGetterImpl.VersionFromCILabel("latest", "experimental version")
		if err != nil {
			return upgrades, err
		}
		minorUnstable := latestVersion.Components()[1]
		previousBranch := fmt.Sprintf("latest-1.%d", minorUnstable-1)
		previousBranchLatestVersionStr, previousBranchLatestVersion, err := versionGetterImpl.VersionFromCILabel(previousBranch, "")
		if err != nil {
			return upgrades, err
		}
		if rcUpgradesAllowed && rcUpgradePossible(clusterVersion, previousBranchLatestVersion) {
			upgrades = append(upgrades, Upgrade{Description: "release candidate version", Before: beforeState, After: ClusterState{KubeVersion: previousBranchLatestVersionStr, DNSType: dnsType, DNSVersion: kubeadmconstants.GetDNSVersion(dnsType), KubeadmVersion: previousBranchLatestVersionStr, EtcdVersion: getSuggestedEtcdVersion(previousBranchLatestVersionStr)}})
		}
		if experimentalUpgradesAllowed && clusterVersion.LessThan(latestVersion) {
			unstableKubeVersion := latestVersionStr
			unstableKubeDNSVersion := kubeadmconstants.GetDNSVersion(dnsType)
			if latestVersion.PreRelease() == "alpha.0" {
				unstableKubeVersion = previousBranchLatestVersionStr
				unstableKubeDNSVersion = kubeadmconstants.GetDNSVersion(dnsType)
			}
			upgrades = append(upgrades, Upgrade{Description: "experimental version", Before: beforeState, After: ClusterState{KubeVersion: unstableKubeVersion, DNSType: dnsType, DNSVersion: unstableKubeDNSVersion, KubeadmVersion: unstableKubeVersion, EtcdVersion: getSuggestedEtcdVersion(unstableKubeVersion)}})
		}
	}
	fmt.Println("")
	return upgrades, nil
}
func getBranchFromVersion(version string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v := versionutil.MustParseGeneric(version)
	return fmt.Sprintf("%d.%d", v.Major(), v.Minor())
}
func patchVersionBranchExists(clusterVersion, stableVersion *versionutil.Version) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return stableVersion.AtLeast(clusterVersion)
}
func patchUpgradePossible(clusterVersion, patchVersion *versionutil.Version) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return clusterVersion.LessThan(patchVersion)
}
func rcUpgradePossible(clusterVersion, previousBranchLatestVersion *versionutil.Version) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return strings.HasPrefix(previousBranchLatestVersion.PreRelease(), "rc") && clusterVersion.LessThan(previousBranchLatestVersion)
}
func minorUpgradePossibleWithPatchRelease(stableVersion, patchVersion *versionutil.Version) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return patchVersion.LessThan(stableVersion)
}
func getSuggestedEtcdVersion(kubernetesVersion string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdVersion, err := kubeadmconstants.EtcdSupportedVersion(kubernetesVersion)
	if err != nil {
		fmt.Printf("[upgrade/versions] WARNING: No recommended etcd for requested Kubernetes version (%s)\n", kubernetesVersion)
		return "N/A"
	}
	return etcdVersion.String()
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
