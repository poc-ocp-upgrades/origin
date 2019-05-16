package upgrade

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/upgrade"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type planFlags struct {
	*applyPlanFlags
	newK8sVersionStr string
}

func NewCmdPlan(apf *applyPlanFlags) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := &planFlags{applyPlanFlags: apf}
	cmd := &cobra.Command{Use: "plan [version] [flags]", Short: "Check which versions are available to upgrade to and validate whether your current cluster is upgradeable. To skip the internet check, pass in the optional [version] parameter.", Run: func(_ *cobra.Command, args []string) {
		var err error
		flags.ignorePreflightErrorsSet, err = validation.ValidateIgnorePreflightErrors(flags.ignorePreflightErrors)
		kubeadmutil.CheckErr(err)
		err = runPreflightChecks(flags.ignorePreflightErrorsSet)
		kubeadmutil.CheckErr(err)
		if flags.cfgPath != "" {
			klog.V(1).Infof("fetching configuration from file %s", flags.cfgPath)
			cfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(flags.cfgPath, &kubeadmapiv1beta1.InitConfiguration{})
			kubeadmutil.CheckErr(err)
			if cfg.KubernetesVersion != "" {
				flags.newK8sVersionStr = cfg.KubernetesVersion
			}
		}
		if len(args) == 1 {
			flags.newK8sVersionStr = args[0]
		}
		err = RunPlan(flags)
		kubeadmutil.CheckErr(err)
	}}
	addApplyPlanFlags(cmd.Flags(), flags.applyPlanFlags)
	return cmd
}
func RunPlan(flags *planFlags) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("[upgrade/plan] verifying health of cluster")
	klog.V(1).Infof("[upgrade/plan] retrieving configuration from cluster")
	upgradeVars, err := enforceRequirements(flags.applyPlanFlags, false, flags.newK8sVersionStr)
	if err != nil {
		return err
	}
	var etcdClient etcdutil.ClusterInterrogator
	isExternalEtcd := upgradeVars.cfg.Etcd.External != nil
	if isExternalEtcd {
		client, err := etcdutil.New(upgradeVars.cfg.Etcd.External.Endpoints, upgradeVars.cfg.Etcd.External.CAFile, upgradeVars.cfg.Etcd.External.CertFile, upgradeVars.cfg.Etcd.External.KeyFile)
		if err != nil {
			return err
		}
		etcdClient = client
	} else {
		client, err := etcdutil.NewFromCluster(upgradeVars.client, upgradeVars.cfg.CertificatesDir)
		if err != nil {
			return err
		}
		etcdClient = client
	}
	klog.V(1).Infof("[upgrade/plan] computing upgrade possibilities")
	availUpgrades, err := upgrade.GetAvailableUpgrades(upgradeVars.versionGetter, flags.allowExperimentalUpgrades, flags.allowRCUpgrades, etcdClient, upgradeVars.cfg.DNS.Type, upgradeVars.client)
	if err != nil {
		return errors.Wrap(err, "[upgrade/versions] FATAL")
	}
	printAvailableUpgrades(availUpgrades, os.Stdout, isExternalEtcd)
	return nil
}
func printAvailableUpgrades(upgrades []upgrade.Upgrade, w io.Writer, isExternalEtcd bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(upgrades) == 0 {
		fmt.Fprintln(w, "Awesome, you're up-to-date! Enjoy!")
		return
	}
	tabw := tabwriter.NewWriter(w, 10, 4, 3, ' ', 0)
	for _, upgrade := range upgrades {
		newK8sVersion, err := version.ParseSemantic(upgrade.After.KubeVersion)
		if err != nil {
			fmt.Fprintf(w, "Unable to parse normalized version %q as a semantic version\n", upgrade.After.KubeVersion)
			continue
		}
		UnstableVersionFlag := ""
		if len(newK8sVersion.PreRelease()) != 0 {
			if strings.HasPrefix(newK8sVersion.PreRelease(), "rc") {
				UnstableVersionFlag = " --allow-release-candidate-upgrades"
			} else {
				UnstableVersionFlag = " --allow-experimental-upgrades"
			}
		}
		if isExternalEtcd && upgrade.CanUpgradeEtcd() {
			fmt.Fprintln(w, "External components that should be upgraded manually before you upgrade the control plane with 'kubeadm upgrade apply':")
			fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
			fmt.Fprintf(tabw, "Etcd\t%s\t%s\n", upgrade.Before.EtcdVersion, upgrade.After.EtcdVersion)
			tabw.Flush()
			fmt.Fprintln(w, "")
		}
		if upgrade.CanUpgradeKubelets() {
			fmt.Fprintln(w, "Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':")
			fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
			firstPrinted := false
			for _, oldVersion := range sortedSliceFromStringIntMap(upgrade.Before.KubeletVersions) {
				nodeCount := upgrade.Before.KubeletVersions[oldVersion]
				if !firstPrinted {
					fmt.Fprintf(tabw, "Kubelet\t%d x %s\t%s\n", nodeCount, oldVersion, upgrade.After.KubeVersion)
					firstPrinted = true
					continue
				}
				fmt.Fprintf(tabw, "\t%d x %s\t%s\n", nodeCount, oldVersion, upgrade.After.KubeVersion)
			}
			tabw.Flush()
			fmt.Fprintln(w, "")
		}
		fmt.Fprintf(w, "Upgrade to the latest %s:\n", upgrade.Description)
		fmt.Fprintln(w, "")
		fmt.Fprintln(tabw, "COMPONENT\tCURRENT\tAVAILABLE")
		fmt.Fprintf(tabw, "API Server\t%s\t%s\n", upgrade.Before.KubeVersion, upgrade.After.KubeVersion)
		fmt.Fprintf(tabw, "Controller Manager\t%s\t%s\n", upgrade.Before.KubeVersion, upgrade.After.KubeVersion)
		fmt.Fprintf(tabw, "Scheduler\t%s\t%s\n", upgrade.Before.KubeVersion, upgrade.After.KubeVersion)
		fmt.Fprintf(tabw, "Kube Proxy\t%s\t%s\n", upgrade.Before.KubeVersion, upgrade.After.KubeVersion)
		printCoreDNS, printKubeDNS := false, false
		coreDNSBeforeVersion, coreDNSAfterVersion, kubeDNSBeforeVersion, kubeDNSAfterVersion := "", "", "", ""
		switch upgrade.Before.DNSType {
		case kubeadmapi.CoreDNS:
			printCoreDNS = true
			coreDNSBeforeVersion = upgrade.Before.DNSVersion
		case kubeadmapi.KubeDNS:
			printKubeDNS = true
			kubeDNSBeforeVersion = upgrade.Before.DNSVersion
		}
		switch upgrade.After.DNSType {
		case kubeadmapi.CoreDNS:
			printCoreDNS = true
			coreDNSAfterVersion = upgrade.After.DNSVersion
		case kubeadmapi.KubeDNS:
			printKubeDNS = true
			kubeDNSAfterVersion = upgrade.After.DNSVersion
		}
		if printCoreDNS {
			fmt.Fprintf(tabw, "CoreDNS\t%s\t%s\n", coreDNSBeforeVersion, coreDNSAfterVersion)
		}
		if printKubeDNS {
			fmt.Fprintf(tabw, "Kube DNS\t%s\t%s\n", kubeDNSBeforeVersion, kubeDNSAfterVersion)
		}
		if !isExternalEtcd {
			fmt.Fprintf(tabw, "Etcd\t%s\t%s\n", upgrade.Before.EtcdVersion, upgrade.After.EtcdVersion)
		}
		tabw.Flush()
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "You can now apply the upgrade by executing the following command:")
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "\tkubeadm upgrade apply %s%s\n", upgrade.After.KubeVersion, UnstableVersionFlag)
		fmt.Fprintln(w, "")
		if upgrade.Before.KubeadmVersion != upgrade.After.KubeadmVersion {
			fmt.Fprintf(w, "Note: Before you can perform this upgrade, you have to update kubeadm to %s.\n", upgrade.After.KubeadmVersion)
			fmt.Fprintln(w, "")
		}
		fmt.Fprintln(w, "_____________________________________________________________________")
		fmt.Fprintln(w, "")
	}
}
func sortedSliceFromStringIntMap(strMap map[string]uint16) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	strSlice := []string{}
	for k := range strMap {
		strSlice = append(strSlice, k)
	}
	sort.Strings(strSlice)
	return strSlice
}
