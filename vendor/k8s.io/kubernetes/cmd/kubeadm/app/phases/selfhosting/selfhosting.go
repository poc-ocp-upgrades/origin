package selfhosting

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"
	clientscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	"os"
	"time"
)

const (
	selfHostingWaitTimeout          = 2 * time.Minute
	selfHostingFailureThreshold int = 5
)

func CreateSelfHostedControlPlane(manifestsDir, kubeConfigDir string, cfg *kubeadmapi.InitConfiguration, client clientset.Interface, waiter apiclient.Waiter, dryRun bool, certsInSecrets bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("creating self hosted control plane")
	waiter.SetTimeout(selfHostingWaitTimeout)
	klog.V(1).Infoln("getting mutators")
	mutators := GetMutatorsFromFeatureGates(certsInSecrets)
	if certsInSecrets {
		if err := uploadTLSSecrets(client, cfg.CertificatesDir); err != nil {
			return err
		}
		if err := uploadKubeConfigSecrets(client, kubeConfigDir); err != nil {
			return err
		}
	}
	for _, componentName := range kubeadmconstants.MasterComponents {
		start := time.Now()
		manifestPath := kubeadmconstants.GetStaticPodFilepath(componentName, manifestsDir)
		if _, err := os.Stat(manifestPath); err != nil {
			fmt.Printf("[self-hosted] The Static Pod for the component %q doesn't seem to be on the disk; trying the next one\n", componentName)
			continue
		}
		podSpec, err := loadPodSpecFromFile(manifestPath)
		if err != nil {
			return err
		}
		ds := BuildDaemonSet(componentName, podSpec, mutators)
		if err := apiclient.TryRunCommand(func() error {
			return apiclient.CreateOrUpdateDaemonSet(client, ds)
		}, selfHostingFailureThreshold); err != nil {
			return err
		}
		if err := waiter.WaitForPodsWithLabel(BuildSelfHostedComponentLabelQuery(componentName)); err != nil {
			return err
		}
		if !dryRun {
			if err := os.RemoveAll(manifestPath); err != nil {
				return errors.Wrapf(err, "unable to delete static pod manifest for %s ", componentName)
			}
		}
		staticPodName := fmt.Sprintf("%s-%s", componentName, cfg.NodeRegistration.Name)
		if err := waiter.WaitForPodToDisappear(staticPodName); err != nil {
			return err
		}
		if err := waiter.WaitForAPI(); err != nil {
			return err
		}
		fmt.Printf("[self-hosted] self-hosted %s ready after %f seconds\n", componentName, time.Since(start).Seconds())
	}
	return nil
}
func BuildDaemonSet(name string, podSpec *v1.PodSpec, mutators map[string][]PodSpecMutatorFunc) *apps.DaemonSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mutatePodSpec(mutators, name, podSpec)
	return &apps.DaemonSet{ObjectMeta: metav1.ObjectMeta{Name: kubeadmconstants.AddSelfHostedPrefix(name), Namespace: metav1.NamespaceSystem, Labels: BuildSelfhostedComponentLabels(name)}, Spec: apps.DaemonSetSpec{Selector: &metav1.LabelSelector{MatchLabels: BuildSelfhostedComponentLabels(name)}, Template: v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: BuildSelfhostedComponentLabels(name)}, Spec: *podSpec}, UpdateStrategy: apps.DaemonSetUpdateStrategy{Type: apps.RollingUpdateDaemonSetStrategyType}}}
}
func BuildSelfhostedComponentLabels(component string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return map[string]string{"k8s-app": kubeadmconstants.AddSelfHostedPrefix(component)}
}
func BuildSelfHostedComponentLabelQuery(componentName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("k8s-app=%s", kubeadmconstants.AddSelfHostedPrefix(componentName))
}
func loadPodSpecFromFile(filePath string) (*v1.PodSpec, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podDef, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file path %s", filePath)
	}
	if len(podDef) == 0 {
		return nil, errors.Errorf("file was empty: %s", filePath)
	}
	codec := clientscheme.Codecs.UniversalDecoder()
	pod := &v1.Pod{}
	if err = runtime.DecodeInto(codec, podDef, pod); err != nil {
		return nil, errors.Wrap(err, "failed decoding pod")
	}
	return &pod.Spec, nil
}
