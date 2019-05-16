package upgrade

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"strings"
)

const (
	MaximumAllowedMinorVersionUpgradeSkew   = 1
	MaximumAllowedMinorVersionDowngradeSkew = 1
	MaximumAllowedMinorVersionKubeletSkew   = 1
)

type VersionSkewPolicyErrors struct {
	Mandatory []error
	Skippable []error
}

func EnforceVersionPolicies(versionGetter VersionGetter, newK8sVersionStr string, newK8sVersion *version.Version, allowExperimentalUpgrades, allowRCUpgrades bool) *VersionSkewPolicyErrors {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	skewErrors := &VersionSkewPolicyErrors{Mandatory: []error{}, Skippable: []error{}}
	clusterVersionStr, clusterVersion, err := versionGetter.ClusterVersion()
	if err != nil {
		skewErrors.Mandatory = append(skewErrors.Mandatory, errors.Wrap(err, "Unable to fetch cluster version"))
		return skewErrors
	}
	kubeadmVersionStr, kubeadmVersion, err := versionGetter.KubeadmVersion()
	if err != nil {
		skewErrors.Mandatory = append(skewErrors.Mandatory, errors.Wrap(err, "Unable to fetch kubeadm version"))
		return skewErrors
	}
	kubeletVersions, err := versionGetter.KubeletVersions()
	if err != nil {
		skewErrors.Skippable = append(skewErrors.Skippable, errors.Wrap(err, "Unable to fetch kubelet version"))
	}
	if constants.MinimumControlPlaneVersion.AtLeast(newK8sVersion) {
		skewErrors.Mandatory = append(skewErrors.Mandatory, errors.Errorf("Specified version to upgrade to %q is equal to or lower than the minimum supported version %q. Please specify a higher version to upgrade to", newK8sVersionStr, clusterVersionStr))
	}
	if newK8sVersion.Minor() > clusterVersion.Minor()+MaximumAllowedMinorVersionUpgradeSkew {
		tooLargeUpgradeSkewErr := errors.Errorf("Specified version to upgrade to %q is too high; kubeadm can upgrade only %d minor version at a time", newK8sVersionStr, MaximumAllowedMinorVersionUpgradeSkew)
		if len(newK8sVersion.PreRelease()) == 0 {
			skewErrors.Mandatory = append(skewErrors.Mandatory, tooLargeUpgradeSkewErr)
		} else {
			skewErrors.Skippable = append(skewErrors.Skippable, tooLargeUpgradeSkewErr)
		}
	}
	if newK8sVersion.Minor() < clusterVersion.Minor()-MaximumAllowedMinorVersionDowngradeSkew {
		tooLargeDowngradeSkewErr := errors.Errorf("Specified version to downgrade to %q is too low; kubeadm can downgrade only %d minor version at a time", newK8sVersionStr, MaximumAllowedMinorVersionDowngradeSkew)
		if len(newK8sVersion.PreRelease()) == 0 {
			skewErrors.Mandatory = append(skewErrors.Mandatory, tooLargeDowngradeSkewErr)
		} else {
			skewErrors.Skippable = append(skewErrors.Skippable, tooLargeDowngradeSkewErr)
		}
	}
	if kubeadmVersion.LessThan(newK8sVersion) {
		if newK8sVersion.Minor() > kubeadmVersion.Minor() {
			tooLargeKubeadmSkew := errors.Errorf("Specified version to upgrade to %q is at least one minor release higher than the kubeadm minor release (%d > %d). Such an upgrade is not supported", newK8sVersionStr, newK8sVersion.Minor(), kubeadmVersion.Minor())
			if len(newK8sVersion.PreRelease()) == 0 {
				skewErrors.Mandatory = append(skewErrors.Mandatory, tooLargeKubeadmSkew)
			} else {
				skewErrors.Skippable = append(skewErrors.Skippable, tooLargeKubeadmSkew)
			}
		} else {
			skewErrors.Skippable = append(skewErrors.Skippable, errors.Errorf("Specified version to upgrade to %q is higher than the kubeadm version %q. Upgrade kubeadm first using the tool you used to install kubeadm", newK8sVersionStr, kubeadmVersionStr))
		}
	}
	if kubeadmVersion.Major() > newK8sVersion.Major() || kubeadmVersion.Minor() > newK8sVersion.Minor() {
		skewErrors.Skippable = append(skewErrors.Skippable, errors.Errorf("Kubeadm version %s can only be used to upgrade to Kubernetes version %d.%d", kubeadmVersionStr, kubeadmVersion.Major(), kubeadmVersion.Minor()))
	}
	if err = detectUnstableVersionError(newK8sVersion, newK8sVersionStr, allowExperimentalUpgrades, allowRCUpgrades); err != nil {
		skewErrors.Skippable = append(skewErrors.Skippable, err)
	}
	if kubeletVersions != nil {
		if err = detectTooOldKubelets(newK8sVersion, kubeletVersions); err != nil {
			skewErrors.Skippable = append(skewErrors.Skippable, err)
		}
	}
	if len(skewErrors.Skippable) == 0 && len(skewErrors.Mandatory) == 0 {
		return nil
	}
	return skewErrors
}
func detectUnstableVersionError(newK8sVersion *version.Version, newK8sVersionStr string, allowExperimentalUpgrades, allowRCUpgrades bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(newK8sVersion.PreRelease()) == 0 {
		return nil
	}
	if allowExperimentalUpgrades {
		return nil
	}
	if strings.HasPrefix(newK8sVersion.PreRelease(), "rc") && allowRCUpgrades {
		return nil
	}
	return errors.Errorf("Specified version to upgrade to %q is an unstable version and such upgrades weren't allowed via setting the --allow-*-upgrades flags", newK8sVersionStr)
}
func detectTooOldKubelets(newK8sVersion *version.Version, kubeletVersions map[string]uint16) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tooOldKubeletVersions := []string{}
	for versionStr := range kubeletVersions {
		kubeletVersion, err := version.ParseSemantic(versionStr)
		if err != nil {
			return errors.Errorf("couldn't parse kubelet version %s", versionStr)
		}
		if newK8sVersion.Minor() > kubeletVersion.Minor()+MaximumAllowedMinorVersionKubeletSkew {
			tooOldKubeletVersions = append(tooOldKubeletVersions, versionStr)
		}
	}
	if len(tooOldKubeletVersions) == 0 {
		return nil
	}
	return errors.Errorf("There are kubelets in this cluster that are too old that have these versions %v", tooOldKubeletVersions)
}
