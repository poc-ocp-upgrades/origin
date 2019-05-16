package upgrade

import (
	"fmt"
	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	"time"
)

const (
	prepullPrefix = "upgrade-prepull-"
)

type Prepuller interface {
	CreateFunc(string) error
	WaitFunc(string)
	DeleteFunc(string) error
}
type DaemonSetPrepuller struct {
	client clientset.Interface
	cfg    *kubeadmapi.ClusterConfiguration
	waiter apiclient.Waiter
}

func NewDaemonSetPrepuller(client clientset.Interface, waiter apiclient.Waiter, cfg *kubeadmapi.ClusterConfiguration) *DaemonSetPrepuller {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &DaemonSetPrepuller{client: client, cfg: cfg, waiter: waiter}
}
func (d *DaemonSetPrepuller) CreateFunc(component string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var image string
	if component == constants.Etcd {
		image = images.GetEtcdImage(d.cfg)
	} else {
		image = images.GetKubernetesImage(component, d.cfg)
	}
	ds := buildPrePullDaemonSet(component, image)
	if err := apiclient.CreateOrUpdateDaemonSet(d.client, ds); err != nil {
		return errors.Wrapf(err, "unable to create a DaemonSet for prepulling the component %q", component)
	}
	return nil
}
func (d *DaemonSetPrepuller) WaitFunc(component string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[upgrade/prepull] Prepulling image for component %s.\n", component)
	d.waiter.WaitForPodsWithLabel("k8s-app=upgrade-prepull-" + component)
}
func (d *DaemonSetPrepuller) DeleteFunc(component string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dsName := addPrepullPrefix(component)
	if err := apiclient.DeleteDaemonSetForeground(d.client, metav1.NamespaceSystem, dsName); err != nil {
		return errors.Wrapf(err, "unable to cleanup the DaemonSet used for prepulling %s", component)
	}
	fmt.Printf("[upgrade/prepull] Prepulled image for component %s.\n", component)
	return nil
}
func PrepullImagesInParallel(kubePrepuller Prepuller, timeout time.Duration, componentsToPrepull []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[upgrade/prepull] Will prepull images for components %v\n", componentsToPrepull)
	timeoutChan := time.After(timeout)
	for _, component := range componentsToPrepull {
		if err := kubePrepuller.CreateFunc(component); err != nil {
			return err
		}
	}
	prePulledChan := make(chan string, len(componentsToPrepull))
	for _, component := range componentsToPrepull {
		go func(c string) {
			kubePrepuller.WaitFunc(c)
			prePulledChan <- c
		}(component)
	}
	if err := waitForItemsFromChan(timeoutChan, prePulledChan, len(componentsToPrepull), kubePrepuller.DeleteFunc); err != nil {
		return err
	}
	fmt.Println("[upgrade/prepull] Successfully prepulled the images for all the control plane components")
	return nil
}
func waitForItemsFromChan(timeoutChan <-chan time.Time, stringChan chan string, n int, cleanupFunc func(string) error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	i := 0
	for {
		select {
		case <-timeoutChan:
			return errors.New("The prepull operation timed out")
		case result := <-stringChan:
			i++
			if err := cleanupFunc(result); err != nil {
				return err
			}
			if i == n {
				return nil
			}
		}
	}
}
func addPrepullPrefix(component string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s%s", prepullPrefix, component)
}
func buildPrePullDaemonSet(component, image string) *apps.DaemonSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var gracePeriodSecs int64
	return &apps.DaemonSet{ObjectMeta: metav1.ObjectMeta{Name: addPrepullPrefix(component), Namespace: metav1.NamespaceSystem}, Spec: apps.DaemonSetSpec{Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"k8s-app": addPrepullPrefix(component)}}, Template: v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"k8s-app": addPrepullPrefix(component)}}, Spec: v1.PodSpec{Containers: []v1.Container{{Name: component, Image: image, Command: []string{"/bin/sleep", "3600"}}}, NodeSelector: map[string]string{constants.LabelNodeRoleMaster: ""}, Tolerations: []v1.Toleration{constants.MasterToleration}, TerminationGracePeriodSeconds: &gracePeriodSecs}}}}
}
