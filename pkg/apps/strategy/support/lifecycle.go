package support

import (
	"context"
	"fmt"
	appsv1 "github.com/openshift/api/apps/v1"
	imageapiv1 "github.com/openshift/api/image/v1"
	imageclienttyped "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	strategyutil "github.com/openshift/origin/pkg/apps/strategy/util"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	"github.com/openshift/origin/pkg/util"
	"io"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"strings"
	"sync"
	"time"
)

const (
	hookContainerName      = "lifecycle"
	deploymentPodTypeLabel = "openshift.io/deployer-pod.type"
	deploymentAnnotation   = "openshift.io/deployment.name"
)

type HookExecutor interface {
	Execute(hook *appsv1.LifecycleHook, rc *corev1.ReplicationController, suffix, label string) error
}

var _ HookExecutor = &hookExecutor{}

type hookExecutor struct {
	pods       corev1client.PodsGetter
	tags       imageclienttyped.ImageStreamTagsGetter
	out        io.Writer
	events     corev1client.EventsGetter
	getPodLogs func(*corev1.Pod) (io.ReadCloser, error)
}

func NewHookExecutor(kubeClient kubernetes.Interface, imageClient imageclienttyped.ImageStreamTagsGetter, out io.Writer) HookExecutor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	executor := &hookExecutor{tags: imageClient, pods: kubeClient.CoreV1(), events: kubeClient.CoreV1(), out: out}
	executor.getPodLogs = func(pod *corev1.Pod) (io.ReadCloser, error) {
		opts := &corev1.PodLogOptions{Container: hookContainerName, Follow: true, Timestamps: false}
		return executor.pods.Pods(pod.Namespace).GetLogs(pod.Name, opts).Stream()
	}
	return executor
}
func (e *hookExecutor) Execute(hook *appsv1.LifecycleHook, rc *corev1.ReplicationController, suffix, label string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	switch {
	case len(hook.TagImages) > 0:
		tagEventMessages := []string{}
		for _, t := range hook.TagImages {
			image, ok := findContainerImage(rc, t.ContainerName)
			if ok {
				tagEventMessages = append(tagEventMessages, fmt.Sprintf("image %q as %q", image, t.To.Name))
			}
		}
		strategyutil.RecordConfigEvent(e.events, rc, kapi.EventTypeNormal, "Started", fmt.Sprintf("Running %s-hook (TagImages) %s for rc %s/%s", label, strings.Join(tagEventMessages, ","), rc.Namespace, rc.Name))
		err = e.tagImages(hook, rc, suffix, label)
	case hook.ExecNewPod != nil:
		strategyutil.RecordConfigEvent(e.events, rc, kapi.EventTypeNormal, "Started", fmt.Sprintf("Running %s-hook (%q) for rc %s/%s", label, strings.Join(hook.ExecNewPod.Command, " "), rc.Namespace, rc.Name))
		err = e.executeExecNewPod(hook, rc, suffix, label)
	}
	if err == nil {
		strategyutil.RecordConfigEvent(e.events, rc, kapi.EventTypeNormal, "Completed", fmt.Sprintf("The %s-hook for rc %s/%s completed successfully", label, rc.Namespace, rc.Name))
		return nil
	}
	switch hook.FailurePolicy {
	case appsv1.LifecycleHookFailurePolicyAbort, appsv1.LifecycleHookFailurePolicyRetry:
		strategyutil.RecordConfigEvent(e.events, rc, kapi.EventTypeWarning, "Failed", fmt.Sprintf("The %s-hook failed: %v, aborting rollout of %s/%s", label, err, rc.Namespace, rc.Name))
		return fmt.Errorf("the %s hook failed: %v, aborting rollout of %s/%s", label, err, rc.Namespace, rc.Name)
	case appsv1.LifecycleHookFailurePolicyIgnore:
		strategyutil.RecordConfigEvent(e.events, rc, kapi.EventTypeWarning, "Failed", fmt.Sprintf("The %s-hook failed: %v (ignore), rollout of %s/%s will continue", label, err, rc.Namespace, rc.Name))
		return nil
	default:
		return err
	}
}
func findContainerImage(rc *corev1.ReplicationController, containerName string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if rc.Spec.Template == nil {
		return "", false
	}
	for _, container := range rc.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			return container.Image, true
		}
	}
	return "", false
}
func (e *hookExecutor) tagImages(hook *appsv1.LifecycleHook, rc *corev1.ReplicationController, suffix, label string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for _, action := range hook.TagImages {
		value, ok := findContainerImage(rc, action.ContainerName)
		if !ok {
			errs = append(errs, fmt.Errorf("unable to find image for container %q, container could not be found", action.ContainerName))
			continue
		}
		namespace := action.To.Namespace
		if len(namespace) == 0 {
			namespace = rc.Namespace
		}
		if _, err := e.tags.ImageStreamTags(namespace).Update(&imageapiv1.ImageStreamTag{ObjectMeta: metav1.ObjectMeta{Name: action.To.Name, Namespace: namespace}, Tag: &imageapiv1.TagReference{From: &corev1.ObjectReference{Kind: "DockerImage", Name: value}}}); err != nil {
			errs = append(errs, err)
			continue
		}
		fmt.Fprintf(e.out, "--> %s: Tagged %q into %s/%s\n", label, value, action.To.Namespace, action.To.Name)
	}
	return utilerrors.NewAggregate(errs)
}
func (e *hookExecutor) executeExecNewPod(hook *appsv1.LifecycleHook, rc *corev1.ReplicationController, suffix, label string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := appsutil.DecodeDeploymentConfig(rc)
	if err != nil {
		return err
	}
	deployerPod, err := e.pods.Pods(rc.Namespace).Get(appsutil.DeployerPodNameForDeployment(rc.Name), metav1.GetOptions{})
	if err != nil {
		return err
	}
	var startTime time.Time
	if deployerPod.Status.StartTime != nil {
		startTime = deployerPod.Status.StartTime.Time
	} else {
		startTime = time.Now()
	}
	podSpec, err := createHookPodManifest(hook, rc, &config.Spec.Strategy, suffix, startTime)
	if err != nil {
		return err
	}
	completed, created := false, false
	pod, err := e.pods.Pods(rc.Namespace).Create(podSpec)
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return fmt.Errorf("couldn't create lifecycle pod for %s: %v", rc.Name, err)
		}
		completed = true
		pod = podSpec
		pod.Namespace = rc.Namespace
	} else {
		created = true
		fmt.Fprintf(e.out, "--> %s: Running hook pod ...\n", label)
	}
	var updatedPod *corev1.Pod
	restarts := int32(0)
	alreadyRead := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	listWatcher := &cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		options.FieldSelector = fields.OneTermEqualSelector("metadata.name", pod.Name).String()
		return e.pods.Pods(pod.Namespace).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		options.FieldSelector = fields.OneTermEqualSelector("metadata.name", pod.Name).String()
		return e.pods.Pods(pod.Namespace).Watch(options)
	}}
	preconditionFunc := func(store cache.Store) (bool, error) {
		_, exists, err := store.Get(&metav1.ObjectMeta{Namespace: pod.Namespace, Name: pod.Name})
		if err != nil {
			return true, err
		}
		if !exists {
			return true, apierrors.NewNotFound(corev1.Resource("pods"), pod.Name)
		}
		return false, nil
	}
	_, err = watchtools.UntilWithSync(context.TODO(), listWatcher, &corev1.Pod{}, preconditionFunc, func(event watch.Event) (bool, error) {
		switch event.Type {
		case watch.Error:
			return false, apierrors.FromObject(event.Object)
		case watch.Added, watch.Modified:
			updatedPod = event.Object.(*corev1.Pod)
		case watch.Deleted:
			err := fmt.Errorf("%s: pod/%s[%s] unexpectedly deleted", label, pod.Name, pod.Namespace)
			fmt.Fprintf(e.out, "%v\n", err)
			return false, err
		}
		switch updatedPod.Status.Phase {
		case corev1.PodRunning:
			completed = false
			canRetry, restartCount := canRetryReading(updatedPod, restarts)
			if alreadyRead && !canRetry {
				break
			}
			if restarts != restartCount {
				wg.Add(1)
				restarts = restartCount
				fmt.Fprintf(e.out, "--> %s: Retrying hook pod (retry #%d)\n", label, restartCount)
			}
			alreadyRead = true
			go e.readPodLogs(pod, wg)
		case corev1.PodSucceeded, corev1.PodFailed:
			if completed {
				if updatedPod.Status.Phase == corev1.PodSucceeded {
					fmt.Fprintf(e.out, "--> %s: Hook pod already succeeded\n", label)
				}
				wg.Done()
				return true, nil
			}
			if !created {
				fmt.Fprintf(e.out, "--> %s: Hook pod is already running ...\n", label)
			}
			if !alreadyRead {
				go e.readPodLogs(pod, wg)
			}
			return true, nil
		default:
			completed = false
		}
		return false, nil
	})
	if err != nil {
		return err
	}
	wg.Wait()
	if updatedPod.Status.Phase == corev1.PodFailed {
		fmt.Fprintf(e.out, "--> %s: Failed\n", label)
		return fmt.Errorf(updatedPod.Status.Message)
	}
	if !completed {
		fmt.Fprintf(e.out, "--> %s: Success\n", label)
	}
	return nil
}
func (e *hookExecutor) readPodLogs(pod *corev1.Pod, wg *sync.WaitGroup) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer wg.Done()
	logStream, err := e.getPodLogs(pod)
	if err != nil || logStream == nil {
		fmt.Fprintf(e.out, "warning: Unable to retrieve hook logs from %s: %v\n", pod.Name, err)
		return
	}
	defer logStream.Close()
	if _, err := io.Copy(e.out, logStream); err != nil {
		fmt.Fprintf(e.out, "\nwarning: Unable to read all logs from %s, continuing: %v\n", pod.Name, err)
	}
}
func createHookPodManifest(hook *appsv1.LifecycleHook, rc *corev1.ReplicationController, strategy *appsv1.DeploymentStrategy, hookType string, startTime time.Time) (*corev1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	exec := hook.ExecNewPod
	var baseContainer *corev1.Container
	for _, container := range rc.Spec.Template.Spec.Containers {
		if container.Name == exec.ContainerName {
			baseContainer = &container
			break
		}
	}
	if baseContainer == nil {
		return nil, fmt.Errorf("no container named '%s' found in rc template", exec.ContainerName)
	}
	envMap := map[string]corev1.EnvVar{}
	mergedEnv := []corev1.EnvVar{}
	for _, env := range baseContainer.Env {
		envMap[env.Name] = env
	}
	for _, env := range exec.Env {
		envMap[env.Name] = env
	}
	for k, v := range envMap {
		mergedEnv = append(mergedEnv, corev1.EnvVar{Name: k, Value: v.Value, ValueFrom: v.ValueFrom})
	}
	mergedEnv = append(mergedEnv, corev1.EnvVar{Name: "OPENSHIFT_DEPLOYMENT_NAME", Value: rc.Name})
	mergedEnv = append(mergedEnv, corev1.EnvVar{Name: "OPENSHIFT_DEPLOYMENT_NAMESPACE", Value: rc.Namespace})
	defaultActiveDeadline := appsutil.MaxDeploymentDurationSeconds
	if strategy.ActiveDeadlineSeconds != nil {
		defaultActiveDeadline = *(strategy.ActiveDeadlineSeconds)
	}
	maxDeploymentDurationSeconds := defaultActiveDeadline - int64(time.Since(startTime).Seconds())
	restartPolicy := corev1.RestartPolicyNever
	if hook.FailurePolicy == appsv1.LifecycleHookFailurePolicyRetry {
		restartPolicy = corev1.RestartPolicyOnFailure
	}
	volumes := []corev1.Volume{}
	volumeNames := sets.NewString()
	for _, volume := range rc.Spec.Template.Spec.Volumes {
		for _, name := range exec.Volumes {
			if volume.Name == name {
				volumes = append(volumes, volume)
				volumeNames.Insert(volume.Name)
			}
		}
	}
	volumeMounts := []corev1.VolumeMount{}
	for _, mount := range baseContainer.VolumeMounts {
		if volumeNames.Has(mount.Name) {
			volumeMounts = append(volumeMounts, corev1.VolumeMount{Name: mount.Name, ReadOnly: mount.ReadOnly, MountPath: mount.MountPath, SubPath: mount.SubPath})
		}
	}
	imagePullSecrets := []corev1.LocalObjectReference{}
	for _, pullSecret := range rc.Spec.Template.Spec.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, corev1.LocalObjectReference{Name: pullSecret.Name})
	}
	gracePeriod := int64(10)
	podSecurityContextCopy := rc.Spec.Template.Spec.SecurityContext.DeepCopy()
	securityContextCopy := baseContainer.SecurityContext.DeepCopy()
	pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: apihelpers.GetPodName(rc.Name, hookType), Namespace: rc.Namespace, Annotations: map[string]string{deploymentAnnotation: rc.Name}, Labels: map[string]string{appsv1.DeployerPodForDeploymentLabel: rc.Name, deploymentPodTypeLabel: hookType}}, Spec: corev1.PodSpec{Containers: []corev1.Container{{Name: hookContainerName, Image: baseContainer.Image, ImagePullPolicy: baseContainer.ImagePullPolicy, Command: exec.Command, WorkingDir: baseContainer.WorkingDir, Env: mergedEnv, Resources: baseContainer.Resources, VolumeMounts: volumeMounts, SecurityContext: securityContextCopy}}, SecurityContext: podSecurityContextCopy, Volumes: volumes, ActiveDeadlineSeconds: &maxDeploymentDurationSeconds, NodeSelector: rc.Spec.Template.Spec.NodeSelector, RestartPolicy: restartPolicy, ImagePullSecrets: imagePullSecrets, TerminationGracePeriodSeconds: &gracePeriod}}
	util.MergeInto(pod.Labels, strategy.Labels, 0)
	util.MergeInto(pod.Annotations, strategy.Annotations, 0)
	return pod, nil
}
func canRetryReading(pod *corev1.Pod, restarts int32) (bool, int32) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(pod.Status.ContainerStatuses) == 0 {
		return false, int32(0)
	}
	restartCount := pod.Status.ContainerStatuses[0].RestartCount
	return pod.Spec.RestartPolicy == corev1.RestartPolicyOnFailure && restartCount > restarts, restartCount
}
