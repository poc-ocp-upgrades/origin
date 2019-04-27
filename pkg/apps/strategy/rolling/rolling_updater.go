package rolling

import (
	"fmt"
	"io"
	"strconv"
	"time"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	scaleclient "k8s.io/client-go/scale"
	"k8s.io/client-go/util/integer"
	"k8s.io/client-go/util/retry"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	deploymentutil "k8s.io/kubernetes/pkg/controller/deployment/util"
	"k8s.io/kubernetes/pkg/kubectl"
)

func ControllerHasDesiredReplicas(rcClient coreclient.ReplicationControllersGetter, controller *api.ReplicationController) wait.ConditionFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	desiredGeneration := controller.Generation
	return func() (bool, error) {
		ctrl, err := rcClient.ReplicationControllers(controller.Namespace).Get(controller.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return ctrl.Status.ObservedGeneration >= desiredGeneration && ctrl.Status.Replicas == *ctrl.Spec.Replicas, nil
	}
}

const (
	kubectlAnnotationPrefix		= "kubectl.kubernetes.io/"
	sourceIdAnnotation		= kubectlAnnotationPrefix + "update-source-id"
	desiredReplicasAnnotation	= kubectlAnnotationPrefix + "desired-replicas"
	originalReplicasAnnotation	= kubectlAnnotationPrefix + "original-replicas"
)

type RollingUpdaterConfig struct {
	Out		io.Writer
	OldRc		*api.ReplicationController
	NewRc		*api.ReplicationController
	UpdatePeriod	time.Duration
	Interval	time.Duration
	Timeout		time.Duration
	MinReadySeconds	int32
	CleanupPolicy	RollingUpdaterCleanupPolicy
	MaxUnavailable	intstr.IntOrString
	MaxSurge	intstr.IntOrString
	OnProgress	func(oldRc, newRc *api.ReplicationController, percentage int) error
}
type RollingUpdaterCleanupPolicy string

const (
	DeleteRollingUpdateCleanupPolicy	RollingUpdaterCleanupPolicy	= "Delete"
	PreserveRollingUpdateCleanupPolicy	RollingUpdaterCleanupPolicy	= "Preserve"
	RenameRollingUpdateCleanupPolicy	RollingUpdaterCleanupPolicy	= "Rename"
)

type RollingUpdater struct {
	rcClient			coreclient.ReplicationControllersGetter
	podClient			coreclient.PodsGetter
	scaleClient			scaleclient.ScalesGetter
	ns				string
	scaleAndWait			func(rc *api.ReplicationController, retry *kubectl.RetryParams, wait *kubectl.RetryParams) (*api.ReplicationController, error)
	getOrCreateTargetController	func(controller *api.ReplicationController, sourceId string) (*api.ReplicationController, bool, error)
	cleanup				func(oldRc, newRc *api.ReplicationController, config *RollingUpdaterConfig) error
	getReadyPods			func(oldRc, newRc *api.ReplicationController, minReadySeconds int32) (int32, int32, error)
	nowFn				func() metav1.Time
}

func NewRollingUpdater(namespace string, rcClient coreclient.ReplicationControllersGetter, podClient coreclient.PodsGetter, sc scaleclient.ScalesGetter) *RollingUpdater {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	updater := &RollingUpdater{rcClient: rcClient, podClient: podClient, scaleClient: sc, ns: namespace}
	updater.scaleAndWait = updater.scaleAndWaitWithScaler
	updater.getOrCreateTargetController = updater.getOrCreateTargetControllerWithClient
	updater.getReadyPods = updater.readyPods
	updater.cleanup = updater.cleanupWithClients
	updater.nowFn = metav1.Now
	return updater
}
func (r *RollingUpdater) Update(config *RollingUpdaterConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := config.Out
	oldRc := config.OldRc
	if oldRc.Spec.Replicas == nil {
		one := int32(1)
		oldRc.Spec.Replicas = &one
	}
	scaleRetryParams := kubectl.NewRetryParams(config.Interval, config.Timeout)
	sourceId := fmt.Sprintf("%s:%s", oldRc.Name, oldRc.UID)
	newRc, existed, err := r.getOrCreateTargetController(config.NewRc, sourceId)
	if err != nil {
		return err
	}
	if newRc.Spec.Replicas == nil {
		one := int32(1)
		newRc.Spec.Replicas = &one
	}
	if existed {
		fmt.Fprintf(out, "Continuing update with existing controller %s.\n", newRc.Name)
	} else {
		fmt.Fprintf(out, "Created %s\n", newRc.Name)
	}
	desiredAnnotation, err := strconv.Atoi(newRc.Annotations[desiredReplicasAnnotation])
	if err != nil {
		return fmt.Errorf("unable to parse annotation for %s: %s=%s", newRc.Name, desiredReplicasAnnotation, newRc.Annotations[desiredReplicasAnnotation])
	}
	desired := int32(desiredAnnotation)
	_, hasOriginalAnnotation := oldRc.Annotations[originalReplicasAnnotation]
	if !hasOriginalAnnotation {
		existing, err := r.rcClient.ReplicationControllers(oldRc.Namespace).Get(oldRc.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if existing.Spec.Replicas != nil {
			originReplicas := strconv.Itoa(int(*existing.Spec.Replicas))
			applyUpdate := func(rc *api.ReplicationController) {
				if rc.Annotations == nil {
					rc.Annotations = map[string]string{}
				}
				rc.Annotations[originalReplicasAnnotation] = originReplicas
			}
			if oldRc, err = updateRcWithRetries(r.rcClient, existing.Namespace, existing, applyUpdate); err != nil {
				return err
			}
		}
	}
	maxSurge, maxUnavailable, err := deploymentutil.ResolveFenceposts(&config.MaxSurge, &config.MaxUnavailable, desired)
	if err != nil {
		return err
	}
	if desired > 0 && maxUnavailable == 0 && maxSurge == 0 {
		return fmt.Errorf("one of maxSurge or maxUnavailable must be specified")
	}
	minAvailable := int32(integer.IntMax(0, int(desired-maxUnavailable)))
	if desired == 0 {
		maxUnavailable = *oldRc.Spec.Replicas
		minAvailable = 0
	}
	fmt.Fprintf(out, "Scaling up %s from %d to %d, scaling down %s from %d to 0 (keep %d pods available, don't exceed %d pods)\n", newRc.Name, *newRc.Spec.Replicas, desired, oldRc.Name, *oldRc.Spec.Replicas, minAvailable, desired+maxSurge)
	goal := desired - *newRc.Spec.Replicas
	if goal < 0 {
		goal = -goal
	}
	progress := func(complete bool) error {
		if config.OnProgress == nil {
			return nil
		}
		progress := desired - *newRc.Spec.Replicas
		if progress < 0 {
			progress = -progress
		}
		percentage := 100
		if !complete && goal > 0 {
			percentage = int((goal - progress) * 100 / goal)
		}
		return config.OnProgress(oldRc, newRc, percentage)
	}
	progressDeadline := time.Now().UnixNano() + config.Timeout.Nanoseconds()
	for *newRc.Spec.Replicas != desired || *oldRc.Spec.Replicas != 0 {
		newReplicas := *newRc.Spec.Replicas
		oldReplicas := *oldRc.Spec.Replicas
		scaledRc, err := r.scaleUp(newRc, oldRc, desired, maxSurge, maxUnavailable, scaleRetryParams, config)
		if err != nil {
			return err
		}
		newRc = scaledRc
		if err := progress(false); err != nil {
			return err
		}
		time.Sleep(config.UpdatePeriod)
		scaledRc, err = r.scaleDown(newRc, oldRc, desired, minAvailable, maxUnavailable, maxSurge, config)
		if err != nil {
			return err
		}
		oldRc = scaledRc
		if err := progress(false); err != nil {
			return err
		}
		progressMade := (*newRc.Spec.Replicas != newReplicas) || (*oldRc.Spec.Replicas != oldReplicas)
		if progressMade {
			progressDeadline = time.Now().UnixNano() + config.Timeout.Nanoseconds()
		} else if time.Now().UnixNano() > progressDeadline {
			return fmt.Errorf("timed out waiting for any update progress to be made")
		}
	}
	if err := progress(true); err != nil {
		return err
	}
	return r.cleanup(oldRc, newRc, config)
}
func (r *RollingUpdater) scaleUp(newRc, oldRc *api.ReplicationController, desired, maxSurge, maxUnavailable int32, scaleRetryParams *kubectl.RetryParams, config *RollingUpdaterConfig) (*api.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if *newRc.Spec.Replicas == desired {
		return newRc, nil
	}
	increment := (desired + maxSurge) - (*oldRc.Spec.Replicas + *newRc.Spec.Replicas)
	if *oldRc.Spec.Replicas == 0 {
		increment = desired - *newRc.Spec.Replicas
	}
	if increment <= 0 {
		return newRc, nil
	}
	*newRc.Spec.Replicas += increment
	if *newRc.Spec.Replicas > desired {
		*newRc.Spec.Replicas = desired
	}
	fmt.Fprintf(config.Out, "Scaling %s up to %d\n", newRc.Name, *newRc.Spec.Replicas)
	scaledRc, err := r.scaleAndWait(newRc, scaleRetryParams, scaleRetryParams)
	if err != nil {
		return nil, err
	}
	return scaledRc, nil
}
func (r *RollingUpdater) scaleDown(newRc, oldRc *api.ReplicationController, desired, minAvailable, maxUnavailable, maxSurge int32, config *RollingUpdaterConfig) (*api.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if *oldRc.Spec.Replicas == 0 {
		return oldRc, nil
	}
	_, newAvailable, err := r.getReadyPods(oldRc, newRc, config.MinReadySeconds)
	if err != nil {
		return nil, err
	}
	allPods := *oldRc.Spec.Replicas + *newRc.Spec.Replicas
	newUnavailable := *newRc.Spec.Replicas - newAvailable
	decrement := allPods - minAvailable - newUnavailable
	if decrement <= 0 {
		return oldRc, nil
	}
	*oldRc.Spec.Replicas -= decrement
	if *oldRc.Spec.Replicas < 0 {
		*oldRc.Spec.Replicas = 0
	}
	if *newRc.Spec.Replicas == desired && newAvailable == desired {
		*oldRc.Spec.Replicas = 0
	}
	fmt.Fprintf(config.Out, "Scaling %s down to %d\n", oldRc.Name, *oldRc.Spec.Replicas)
	retryWait := &kubectl.RetryParams{Interval: config.Interval, Timeout: config.Timeout}
	scaledRc, err := r.scaleAndWait(oldRc, retryWait, retryWait)
	if err != nil {
		return nil, err
	}
	return scaledRc, nil
}
func (r *RollingUpdater) scaleAndWaitWithScaler(rc *api.ReplicationController, retry, wait *kubectl.RetryParams) (*api.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scaler := kubectl.NewScaler(r.scaleClient)
	if err := scaler.Scale(rc.Namespace, rc.Name, uint(*rc.Spec.Replicas), &kubectl.ScalePrecondition{Size: -1}, retry, wait, schema.GroupResource{Resource: "replicationcontrollers"}); err != nil {
		return nil, err
	}
	return r.rcClient.ReplicationControllers(rc.Namespace).Get(rc.Name, metav1.GetOptions{})
}
func (r *RollingUpdater) readyPods(oldRc, newRc *api.ReplicationController, minReadySeconds int32) (int32, int32, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	controllers := []*api.ReplicationController{oldRc, newRc}
	oldReady := int32(0)
	newReady := int32(0)
	if r.nowFn == nil {
		r.nowFn = metav1.Now
	}
	for i := range controllers {
		controller := controllers[i]
		selector := labels.Set(controller.Spec.Selector).AsSelector()
		options := metav1.ListOptions{LabelSelector: selector.String()}
		pods, err := r.podClient.Pods(controller.Namespace).List(options)
		if err != nil {
			return 0, 0, err
		}
		for _, pod := range pods.Items {
			if pod.DeletionTimestamp != nil {
				continue
			}
			if !podutil.IsPodAvailable(&pod, minReadySeconds, r.nowFn()) {
				continue
			}
			switch controller.Name {
			case oldRc.Name:
				oldReady++
			case newRc.Name:
				newReady++
			}
		}
	}
	return oldReady, newReady, nil
}
func (r *RollingUpdater) getOrCreateTargetControllerWithClient(controller *api.ReplicationController, sourceId string) (*api.ReplicationController, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	existingRc, err := r.existingController(controller)
	if err != nil {
		if !errors.IsNotFound(err) {
			return nil, false, err
		}
		if controller.Spec.Replicas == nil || *controller.Spec.Replicas <= 0 {
			return nil, false, fmt.Errorf("Invalid controller spec for %s; required: > 0 replicas, actual: %d\n", controller.Name, controller.Spec.Replicas)
		}
		if controller.Annotations == nil {
			controller.Annotations = map[string]string{}
		}
		controller.Annotations[desiredReplicasAnnotation] = fmt.Sprintf("%d", controller.Spec.Replicas)
		controller.Annotations[sourceIdAnnotation] = sourceId
		zero := int32(0)
		controller.Spec.Replicas = &zero
		newRc, err := r.rcClient.ReplicationControllers(r.ns).Create(controller)
		return newRc, false, err
	}
	annotations := existingRc.Annotations
	source := annotations[sourceIdAnnotation]
	_, ok := annotations[desiredReplicasAnnotation]
	if source != sourceId || !ok {
		return nil, false, fmt.Errorf("missing/unexpected annotations for controller %s, expected %s : %s", controller.Name, sourceId, annotations)
	}
	return existingRc, true, nil
}
func (r *RollingUpdater) existingController(controller *api.ReplicationController) (*api.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(controller.Name) == 0 && len(controller.GenerateName) > 0 {
		return nil, errors.NewNotFound(api.Resource("replicationcontrollers"), controller.Name)
	}
	return r.rcClient.ReplicationControllers(controller.Namespace).Get(controller.Name, metav1.GetOptions{})
}
func (r *RollingUpdater) cleanupWithClients(oldRc, newRc *api.ReplicationController, config *RollingUpdaterConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	newRc, err = r.rcClient.ReplicationControllers(r.ns).Get(newRc.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	applyUpdate := func(rc *api.ReplicationController) {
		delete(rc.Annotations, sourceIdAnnotation)
		delete(rc.Annotations, desiredReplicasAnnotation)
	}
	if newRc, err = updateRcWithRetries(r.rcClient, r.ns, newRc, applyUpdate); err != nil {
		return err
	}
	if err = wait.Poll(config.Interval, config.Timeout, ControllerHasDesiredReplicas(r.rcClient, newRc)); err != nil {
		return err
	}
	newRc, err = r.rcClient.ReplicationControllers(r.ns).Get(newRc.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	switch config.CleanupPolicy {
	case DeleteRollingUpdateCleanupPolicy:
		fmt.Fprintf(config.Out, "Update succeeded. Deleting %s\n", oldRc.Name)
		return r.rcClient.ReplicationControllers(r.ns).Delete(oldRc.Name, nil)
	case RenameRollingUpdateCleanupPolicy:
		fmt.Fprintf(config.Out, "Update succeeded. Deleting old controller: %s\n", oldRc.Name)
		if err := r.rcClient.ReplicationControllers(r.ns).Delete(oldRc.Name, nil); err != nil {
			return err
		}
		fmt.Fprintf(config.Out, "Renaming %s to %s\n", newRc.Name, oldRc.Name)
		return Rename(r.rcClient, newRc, oldRc.Name)
	case PreserveRollingUpdateCleanupPolicy:
		return nil
	default:
		return nil
	}
}
func Rename(c coreclient.ReplicationControllersGetter, rc *api.ReplicationController, newName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldName := rc.Name
	rc.Name = newName
	rc.ResourceVersion = ""
	propagation := metav1.DeletePropagationOrphan
	err := c.ReplicationControllers(rc.Namespace).Delete(oldName, &metav1.DeleteOptions{PropagationPolicy: &propagation})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	err = wait.Poll(5*time.Second, 60*time.Second, func() (bool, error) {
		_, err := c.ReplicationControllers(rc.Namespace).Get(oldName, metav1.GetOptions{})
		if err == nil {
			return false, nil
		} else if errors.IsNotFound(err) {
			return true, nil
		} else {
			return false, err
		}
	})
	if err != nil {
		return err
	}
	_, err = c.ReplicationControllers(rc.Namespace).Create(rc)
	return err
}

type NewControllerConfig struct {
	Namespace		string
	OldName, NewName	string
	Image			string
	Container		string
	DeploymentKey		string
	PullPolicy		api.PullPolicy
}
type updateRcFunc func(controller *api.ReplicationController)

func updateRcWithRetries(rcClient coreclient.ReplicationControllersGetter, namespace string, rc *api.ReplicationController, applyUpdate updateRcFunc) (*api.ReplicationController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldRc := rc.DeepCopy()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (e error) {
		applyUpdate(rc)
		if rc, e = rcClient.ReplicationControllers(namespace).Update(rc); e == nil {
			return
		}
		updateErr := e
		if rc, e = rcClient.ReplicationControllers(namespace).Get(oldRc.Name, metav1.GetOptions{}); e != nil {
			rc = oldRc
		}
		return updateErr
	})
	return rc, err
}
