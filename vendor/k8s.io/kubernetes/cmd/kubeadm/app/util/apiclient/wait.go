package apiclient

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	"net/http"
	"time"
)

type Waiter interface {
	WaitForAPI() error
	WaitForPodsWithLabel(kvLabel string) error
	WaitForPodToDisappear(staticPodName string) error
	WaitForStaticPodSingleHash(nodeName string, component string) (string, error)
	WaitForStaticPodHashChange(nodeName, component, previousHash string) error
	WaitForStaticPodControlPlaneHashes(nodeName string) (map[string]string, error)
	WaitForHealthyKubelet(initalTimeout time.Duration, healthzEndpoint string) error
	WaitForKubeletAndFunc(f func() error) error
	SetTimeout(timeout time.Duration)
}
type KubeWaiter struct {
	client  clientset.Interface
	timeout time.Duration
	writer  io.Writer
}

func NewKubeWaiter(client clientset.Interface, timeout time.Duration, writer io.Writer) Waiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &KubeWaiter{client: client, timeout: timeout, writer: writer}
}
func (w *KubeWaiter) WaitForAPI() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	start := time.Now()
	return wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
		healthStatus := 0
		w.client.Discovery().RESTClient().Get().AbsPath("/healthz").Do().StatusCode(&healthStatus)
		if healthStatus != http.StatusOK {
			return false, nil
		}
		fmt.Printf("[apiclient] All control plane components are healthy after %f seconds\n", time.Since(start).Seconds())
		return true, nil
	})
}
func (w *KubeWaiter) WaitForPodsWithLabel(kvLabel string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lastKnownPodNumber := -1
	return wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
		listOpts := metav1.ListOptions{LabelSelector: kvLabel}
		pods, err := w.client.CoreV1().Pods(metav1.NamespaceSystem).List(listOpts)
		if err != nil {
			fmt.Fprintf(w.writer, "[apiclient] Error getting Pods with label selector %q [%v]\n", kvLabel, err)
			return false, nil
		}
		if lastKnownPodNumber != len(pods.Items) {
			fmt.Fprintf(w.writer, "[apiclient] Found %d Pods for label selector %s\n", len(pods.Items), kvLabel)
			lastKnownPodNumber = len(pods.Items)
		}
		if len(pods.Items) == 0 {
			return false, nil
		}
		for _, pod := range pods.Items {
			if pod.Status.Phase != v1.PodRunning {
				return false, nil
			}
		}
		return true, nil
	})
}
func (w *KubeWaiter) WaitForPodToDisappear(podName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
		_, err := w.client.CoreV1().Pods(metav1.NamespaceSystem).Get(podName, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			fmt.Printf("[apiclient] The old Pod %q is now removed (which is desired)\n", podName)
			return true, nil
		}
		return false, nil
	})
}
func (w *KubeWaiter) WaitForHealthyKubelet(initalTimeout time.Duration, healthzEndpoint string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	time.Sleep(initalTimeout)
	fmt.Printf("[kubelet-check] Initial timeout of %v passed.\n", initalTimeout)
	return TryRunCommand(func() error {
		client := &http.Client{Transport: netutil.SetOldTransportDefaults(&http.Transport{})}
		resp, err := client.Get(healthzEndpoint)
		if err != nil {
			fmt.Println("[kubelet-check] It seems like the kubelet isn't running or healthy.")
			fmt.Printf("[kubelet-check] The HTTP call equal to 'curl -sSL %s' failed with error: %v.\n", healthzEndpoint, err)
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("[kubelet-check] It seems like the kubelet isn't running or healthy.")
			fmt.Printf("[kubelet-check] The HTTP call equal to 'curl -sSL %s' returned HTTP code %d\n", healthzEndpoint, resp.StatusCode)
			return errors.New("the kubelet healthz endpoint is unhealthy")
		}
		return nil
	}, 5)
}
func (w *KubeWaiter) WaitForKubeletAndFunc(f func() error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errorChan := make(chan error)
	go func(errC chan error, waiter Waiter) {
		if err := waiter.WaitForHealthyKubelet(40*time.Second, fmt.Sprintf("http://localhost:%d/healthz", kubeadmconstants.KubeletHealthzPort)); err != nil {
			errC <- err
		}
	}(errorChan, w)
	go func(errC chan error, waiter Waiter) {
		errC <- f()
	}(errorChan, w)
	return <-errorChan
}
func (w *KubeWaiter) SetTimeout(timeout time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	w.timeout = timeout
}
func (w *KubeWaiter) WaitForStaticPodControlPlaneHashes(nodeName string) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	componentHash := ""
	var err error
	mirrorPodHashes := map[string]string{}
	for _, component := range constants.MasterComponents {
		err = wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
			componentHash, err = getStaticPodSingleHash(w.client, nodeName, component)
			if err != nil {
				return false, nil
			}
			return true, nil
		})
		if err != nil {
			return nil, err
		}
		mirrorPodHashes[component] = componentHash
	}
	return mirrorPodHashes, nil
}
func (w *KubeWaiter) WaitForStaticPodSingleHash(nodeName string, component string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	componentPodHash := ""
	var err error
	err = wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
		componentPodHash, err = getStaticPodSingleHash(w.client, nodeName, component)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	return componentPodHash, err
}
func (w *KubeWaiter) WaitForStaticPodHashChange(nodeName, component, previousHash string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return wait.PollImmediate(constants.APICallRetryInterval, w.timeout, func() (bool, error) {
		hash, err := getStaticPodSingleHash(w.client, nodeName, component)
		if err != nil {
			return false, nil
		}
		if hash == previousHash {
			return false, nil
		}
		return true, nil
	})
}
func getStaticPodSingleHash(client clientset.Interface, nodeName string, component string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	staticPodName := fmt.Sprintf("%s-%s", component, nodeName)
	staticPod, err := client.CoreV1().Pods(metav1.NamespaceSystem).Get(staticPodName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	staticPodHash := staticPod.Annotations[kubetypes.ConfigHashAnnotationKey]
	fmt.Printf("Static pod: %s hash: %s\n", staticPodName, staticPodHash)
	return staticPodHash, nil
}
func TryRunCommand(f func() error, failureThreshold int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backoff := wait.Backoff{Duration: 5 * time.Second, Factor: 2, Steps: failureThreshold}
	return wait.ExponentialBackoff(backoff, func() (bool, error) {
		err := f()
		if err != nil {
			return false, nil
		}
		return true, nil
	})
}
