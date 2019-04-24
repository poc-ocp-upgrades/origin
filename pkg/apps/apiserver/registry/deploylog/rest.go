package deploylog

import (
	"context"
	"fmt"
	"sort"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	genericrest "k8s.io/apiserver/pkg/registry/generic/rest"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/controller"
	"github.com/openshift/api/apps"
	appsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	apiserverrest "github.com/openshift/origin/pkg/apiserver/rest"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	"github.com/openshift/origin/pkg/apps/apis/apps/validation"
	appsutil "github.com/openshift/origin/pkg/apps/util"
)

const (
	defaultTimeout	= 60 * time.Second
	defaultInterval	= 1 * time.Second
)

type REST struct {
	dcClient	appsclient.DeploymentConfigsGetter
	rcClient	corev1client.ReplicationControllersGetter
	podClient	corev1client.PodsGetter
	timeout		time.Duration
	interval	time.Duration
	getLogsFn	func(podNamespace, podName string, logOpts *corev1.PodLogOptions) (runtime.Object, error)
}

var _ = rest.GetterWithOptions(&REST{})

func NewREST(dcClient appsclient.DeploymentConfigsGetter, client kubernetes.Interface) *REST {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &REST{dcClient: dcClient, rcClient: client.CoreV1(), podClient: client.CoreV1(), timeout: defaultTimeout, interval: defaultInterval}
	r.getLogsFn = r.getLogs
	return r
}
func (r *REST) NewGetOptions() (runtime.Object, bool, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentLogOptions{}, false, ""
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &appsapi.DeploymentLog{}
}
func (r *REST) Get(ctx context.Context, name string, opts runtime.Object) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("namespace parameter required.")
	}
	deployLogOpts, ok := opts.(*appsapi.DeploymentLogOptions)
	if !ok {
		return nil, apierrors.NewBadRequest("did not get an expected options.")
	}
	if errs := validation.ValidateDeploymentLogOptions(deployLogOpts); len(errs) > 0 {
		return nil, apierrors.NewInvalid(apps.Kind("DeploymentLogOptions"), "", errs)
	}
	config, err := r.dcClient.DeploymentConfigs(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, apierrors.NewNotFound(apps.Resource("deploymentconfig"), name)
	}
	desiredVersion := config.Status.LatestVersion
	if desiredVersion == 0 {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("no deployment exists for deploymentConfig %q", config.Name))
	}
	switch {
	case deployLogOpts.Version == nil:
		if deployLogOpts.Previous {
			desiredVersion--
			if desiredVersion < 1 {
				return nil, apierrors.NewBadRequest(fmt.Sprintf("no previous deployment exists for deploymentConfig %q", config.Name))
			}
		}
	case *deployLogOpts.Version <= 0 || *deployLogOpts.Version > config.Status.LatestVersion:
		return nil, apierrors.NewBadRequest(fmt.Sprintf("invalid version for deploymentConfig %q: %d", config.Name, *deployLogOpts.Version))
	default:
		desiredVersion = *deployLogOpts.Version
	}
	targetName := appsutil.DeploymentNameForConfigVersion(config.Name, desiredVersion)
	target, err := r.waitForExistingDeployment(namespace, targetName)
	if err != nil {
		return nil, err
	}
	podName := appsutil.DeployerPodNameForDeployment(target.Name)
	labelForDeployment := fmt.Sprintf("%s/%s", target.Namespace, target.Name)
	status := appsutil.DeploymentStatusFor(target)
	switch status {
	case appsv1.DeploymentStatusNew, appsv1.DeploymentStatusPending:
		if deployLogOpts.NoWait {
			klog.V(4).Infof("Deployment %s is in %s state. No logs to retrieve yet.", labelForDeployment, status)
			return &genericrest.LocationStreamer{}, nil
		}
		klog.V(4).Infof("Deployment %s is in %s state, waiting for it to start...", labelForDeployment, status)
		if err := WaitForRunningDeployerPod(r.podClient, target, r.timeout); err != nil {
			return nil, apierrors.NewBadRequest(fmt.Sprintf("failed to run deployer pod %s: %v", podName, err))
		}
		latest, err := WaitForRunningDeployment(r.rcClient, target, r.timeout)
		if err == wait.ErrWaitTimeout {
			return nil, apierrors.NewServerTimeout(kapi.Resource("ReplicationController"), "get", 2)
		}
		if err != nil {
			return nil, apierrors.NewBadRequest(fmt.Sprintf("unable to wait for deployment %s to run: %v", labelForDeployment, err))
		}
		if appsutil.IsCompleteDeployment(latest) {
			podName, err = r.returnApplicationPodName(target)
			if err != nil {
				return nil, err
			}
		}
	case appsv1.DeploymentStatusComplete:
		podName, err = r.returnApplicationPodName(target)
		if err != nil {
			return nil, err
		}
	}
	logOpts := DeploymentToPodLogOptions(deployLogOpts)
	return r.getLogsFn(namespace, podName, logOpts)
}
func (r *REST) getLogs(podNamespace, podName string, logOpts *corev1.PodLogOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logRequest := r.podClient.Pods(podNamespace).GetLogs(podName, logOpts)
	readerCloser, err := logRequest.Stream()
	if err != nil {
		return nil, err
	}
	return &apiserverrest.PassThroughStreamer{In: readerCloser, Flush: logOpts.Follow, ContentType: "text/plain"}, nil
}
func (r *REST) waitForExistingDeployment(namespace, name string) (*corev1.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		target	*corev1.ReplicationController
		err	error
	)
	condition := func() (bool, error) {
		target, err = r.rcClient.ReplicationControllers(namespace).Get(name, metav1.GetOptions{})
		switch {
		case apierrors.IsNotFound(err):
			return false, nil
		case err != nil:
			return false, err
		}
		return true, nil
	}
	err = wait.PollImmediate(r.interval, r.timeout, condition)
	if err == wait.ErrWaitTimeout {
		err = apierrors.NewNotFound(kapi.Resource("replicationcontrollers"), name)
	}
	return target, err
}
func (r *REST) returnApplicationPodName(target *corev1.ReplicationController) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	selector := labels.SelectorFromValidatedSet(labels.Set(target.Spec.Selector))
	sortBy := func(pods []*corev1.Pod) sort.Interface {
		return controller.ByLogging(pods)
	}
	firstPod, _, err := GetFirstPod(r.podClient, target.Namespace, selector.String(), r.timeout, sortBy)
	if err != nil {
		return "", apierrors.NewInternalError(err)
	}
	return firstPod.Name, nil
}
func GetFirstPod(client corev1client.PodsGetter, namespace string, selector string, timeout time.Duration, sortBy func([]*corev1.Pod) sort.Interface) (*corev1.Pod, int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lw := &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		options.LabelSelector = selector
		return client.Pods(namespace).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		options.LabelSelector = selector
		return client.Pods(namespace).Watch(options)
	}}
	var initialPods []*corev1.Pod
	preconditionFunc := func(store cache.Store) (bool, error) {
		items := store.List()
		if len(items) > 0 {
			for _, item := range items {
				pod, ok := item.(*corev1.Pod)
				if !ok {
					return true, fmt.Errorf("unexpected store item type: %#v", item)
				}
				initialPods = append(initialPods, pod)
			}
			sort.Sort(sortBy(initialPods))
			return true, nil
		}
		return false, nil
	}
	ctx, cancel := watchtools.ContextWithOptionalTimeout(context.Background(), timeout)
	defer cancel()
	event, err := watchtools.UntilWithSync(ctx, lw, &corev1.Pod{}, preconditionFunc, func(event watch.Event) (bool, error) {
		switch event.Type {
		case watch.Added, watch.Modified:
			return true, nil
		case watch.Deleted:
			return true, fmt.Errorf("pod got deleted %#v", event.Object)
		case watch.Error:
			return true, fmt.Errorf("unexpected error %#v", event.Object)
		default:
			return true, fmt.Errorf("unexpected event type: %T", event.Type)
		}
	})
	if err != nil {
		return nil, 0, err
	}
	if len(initialPods) > 0 {
		return initialPods[0], len(initialPods), nil
	}
	pod, ok := event.Object.(*corev1.Pod)
	if !ok {
		return nil, 0, fmt.Errorf("%#v is not a pod event", event)
	}
	return pod, 1, nil
}
func WaitForRunningDeployerPod(podClient corev1client.PodsGetter, rc *corev1.ReplicationController, timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podName := appsutil.DeployerPodNameForDeployment(rc.Name)
	canGetLogs := func(p *corev1.Pod) bool {
		return corev1.PodSucceeded == p.Status.Phase || corev1.PodFailed == p.Status.Phase || corev1.PodRunning == p.Status.Phase
	}
	fieldSelector := fields.OneTermEqualSelector("metadata.name", podName).String()
	lw := &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		options.FieldSelector = fieldSelector
		return podClient.Pods(rc.Namespace).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		options.FieldSelector = fieldSelector
		return podClient.Pods(rc.Namespace).Watch(options)
	}}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := watchtools.UntilWithSync(ctx, lw, &corev1.Pod{}, nil, func(e watch.Event) (bool, error) {
		switch e.Type {
		case watch.Added, watch.Modified:
			newPod, ok := e.Object.(*corev1.Pod)
			if !ok {
				return true, fmt.Errorf("unknown event object %#v", e.Object)
			}
			return canGetLogs(newPod), nil
		case watch.Deleted:
			return true, fmt.Errorf("pod got deleted %#v", e.Object)
		case watch.Error:
			return true, fmt.Errorf("encountered error while watching for pod: %v", e.Object)
		default:
			return true, fmt.Errorf("unexpected event type: %T", e.Type)
		}
	})
	return err
}
func DeploymentToPodLogOptions(opts *appsapi.DeploymentLogOptions) *corev1.PodLogOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.PodLogOptions{Container: opts.Container, Follow: opts.Follow, SinceSeconds: opts.SinceSeconds, SinceTime: opts.SinceTime, Timestamps: opts.Timestamps, TailLines: opts.TailLines, LimitBytes: opts.LimitBytes}
}
