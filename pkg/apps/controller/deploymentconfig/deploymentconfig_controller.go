package deploymentconfig

import (
	"fmt"
	"reflect"
	"strings"
	"k8s.io/klog"
	"k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kcoreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	kcorelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	kcontroller "k8s.io/kubernetes/pkg/controller"
	appsv1 "github.com/openshift/api/apps/v1"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	appsv1lister "github.com/openshift/client-go/apps/listers/apps/v1"
	appsutil "github.com/openshift/origin/pkg/apps/util"
)

const (
	maxRetryCount = 15
)

type fatalError string

func (e fatalError) Error() string {
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
	return fmt.Sprintf("fatal error handling deployment config: %s", string(e))
}

type DeploymentConfigController struct {
	appsClient	appsv1client.DeploymentConfigsGetter
	kubeClient	kcoreclient.ReplicationControllersGetter
	queue		workqueue.RateLimitingInterface
	dcIndex		cache.Indexer
	dcLister	appsv1lister.DeploymentConfigLister
	dcStoreSynced	func() bool
	rcLister	kcorelisters.ReplicationControllerLister
	rcListerSynced	func() bool
	rcControl	RCControlInterface
	codec		runtime.Codec
	recorder	record.EventRecorder
}

func (c *DeploymentConfigController) Handle(config *appsv1.DeploymentConfig) error {
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
	klog.V(5).Infof("Reconciling %s/%s", config.Namespace, config.Name)
	if appsutil.IsInitialDeployment(config) && !appsutil.HasTrigger(config) {
		return c.updateStatus(config, []*v1.ReplicationController{}, true)
	}
	rcList, err := c.rcLister.ReplicationControllers(config.Namespace).List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error while deploymentConfigController listing replication controllers: %v", err)
	}
	canAdoptFunc := kcontroller.RecheckDeletionTimestamp(func() (metav1.Object, error) {
		fresh, err := c.appsClient.DeploymentConfigs(config.Namespace).Get(config.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if fresh.UID != config.UID {
			return nil, fmt.Errorf("original DeploymentConfig %s/%s is gone: got uid %s, wanted %s", config.Namespace, config.Name, fresh.UID, config.UID)
		}
		return fresh, nil
	})
	cm := NewRCControllerRefManager(c.rcControl, config, appsutil.ConfigSelector(config.Name), appsv1.GroupVersion.WithKind("DeploymentConfig"), canAdoptFunc)
	existingDeployments, err := cm.ClaimReplicationControllers(rcList)
	if err != nil {
		return fmt.Errorf("error while deploymentConfigController claiming replication controllers: %v", err)
	}
	if config.DeletionTimestamp != nil {
		return c.updateStatus(config, existingDeployments, true)
	}
	if config.Spec.Paused {
		if err := c.cleanupOldDeployments(existingDeployments, config); err != nil {
			c.recorder.Eventf(config, v1.EventTypeWarning, "DeploymentCleanupFailed", "Couldn't clean up deployments: %v", err)
		}
		return c.updateStatus(config, existingDeployments, true)
	}
	latestExists, latestDeployment := appsutil.LatestDeploymentInfo(config, existingDeployments)
	if !latestExists {
		if err := c.cancelRunningRollouts(config, existingDeployments, cm); err != nil {
			return err
		}
	}
	for i, container := range config.Spec.Template.Spec.Containers {
		if len(strings.TrimSpace(container.Image)) == 0 {
			klog.V(4).Infof("Postponing rollout #%d for DeploymentConfig %s/%s because of invalid or unresolved image for container #%d (name=%s)", config.Status.LatestVersion, config.Namespace, config.Name, i, container.Name)
			return c.updateStatus(config, existingDeployments, true)
		}
	}
	configCopy := config.DeepCopy()
	shouldTrigger, shouldSkip, err := triggerActivated(configCopy, latestExists, latestDeployment)
	if err != nil {
		return fmt.Errorf("triggerActivated failed: %v", err)
	}
	if shouldSkip {
		return c.updateStatus(configCopy, existingDeployments, true)
	}
	if shouldTrigger {
		configCopy.Status.LatestVersion++
		_, err := c.appsClient.DeploymentConfigs(configCopy.Namespace).UpdateStatus(configCopy)
		return err
	}
	if latestExists {
		if !appsutil.IsTerminatedDeployment(latestDeployment) {
			return c.updateStatus(config, existingDeployments, false)
		}
		return c.reconcileDeployments(existingDeployments, config, cm)
	}
	deployment, err := appsutil.MakeDeployment(config)
	if err != nil {
		return fatalError(fmt.Sprintf("couldn't make deployment from (potentially invalid) deployment config %s: %v", appsutil.LabelForDeploymentConfig(config), err))
	}
	created, err := c.kubeClient.ReplicationControllers(config.Namespace).Create(deployment)
	if err != nil {
		if kapierrors.IsAlreadyExists(err) {
			rc, err := c.rcLister.ReplicationControllers(deployment.Namespace).Get(deployment.Name)
			if err != nil {
				return fmt.Errorf("error while deploymentConfigController getting the replication controller %s/%s: %v", deployment.Namespace, deployment.Name, err)
			}
			isOurs, err := cm.ClaimReplicationController(rc)
			if err != nil {
				return fmt.Errorf("error while deploymentConfigController claiming the replication controller: %v", err)
			}
			if isOurs {
				return c.updateStatus(config, existingDeployments, true)
			} else {
				err = fmt.Errorf("replication controller %s already exists and deployment config is not allowed to claim it", deployment.Name)
				c.recorder.Eventf(config, v1.EventTypeWarning, "DeploymentCreationFailed", "Couldn't deploy version %d: %v", config.Status.LatestVersion, err)
				return c.updateStatus(config, existingDeployments, true)
			}
		}
		c.recorder.Eventf(config, v1.EventTypeWarning, "DeploymentCreationFailed", "Couldn't deploy version %d: %s", config.Status.LatestVersion, err)
		cond := appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionFalse, appsutil.FailedRcCreateReason, err.Error())
		_ = c.updateStatus(config, existingDeployments, true, *cond)
		return fmt.Errorf("couldn't create deployment for deployment config %s: %v", appsutil.LabelForDeploymentConfig(config), err)
	}
	msg := fmt.Sprintf("Created new replication controller %q for version %d", created.Name, config.Status.LatestVersion)
	c.recorder.Eventf(config, v1.EventTypeNormal, "DeploymentCreated", msg)
	existingDeployments = append(existingDeployments, created)
	if err := c.cleanupOldDeployments(existingDeployments, config); err != nil {
		c.recorder.Eventf(config, v1.EventTypeWarning, "DeploymentCleanupFailed", "Couldn't clean up deployments: %v", err)
	}
	cond := appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionTrue, appsutil.NewReplicationControllerReason, msg)
	return c.updateStatus(config, existingDeployments, true, *cond)
}
func (c *DeploymentConfigController) reconcileDeployments(existingDeployments []*v1.ReplicationController, config *appsv1.DeploymentConfig, cm *RCControllerRefManager) error {
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
	activeDeployment := appsutil.ActiveDeployment(existingDeployments)
	var updatedDeployments []*v1.ReplicationController
	for i := range existingDeployments {
		deployment := existingDeployments[i]
		toAppend := deployment
		isActiveDeployment := activeDeployment != nil && deployment.Name == activeDeployment.Name
		oldReplicaCount := deployment.Spec.Replicas
		if oldReplicaCount == nil {
			zero := int32(0)
			oldReplicaCount = &zero
		}
		newReplicaCount := int32(0)
		if isActiveDeployment {
			newReplicaCount = config.Spec.Replicas
		}
		if config.Spec.Test {
			klog.V(4).Infof("Deployment config %q is test and deployment %q will be scaled down", appsutil.LabelForDeploymentConfig(config), appsutil.LabelForDeployment(deployment))
			newReplicaCount = 0
		}
		var copied *v1.ReplicationController
		if newReplicaCount != *oldReplicaCount {
			if err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
				rc, err := c.rcLister.ReplicationControllers(deployment.Namespace).Get(deployment.Name)
				if err != nil {
					return err
				}
				isOurs, err := cm.ClaimReplicationController(rc)
				if err != nil {
					return fmt.Errorf("error while deploymentConfigController claiming the replication controller %s/%s: %v", rc.Namespace, rc.Name, err)
				}
				if !isOurs {
					return fmt.Errorf("deployment config %s/%s (%v) no longer owns replication controller %s/%s (%v)", config.Namespace, config.Name, config.UID, deployment.Namespace, deployment.Name, deployment.UID)
				}
				copied = rc.DeepCopy()
				copied.Spec.Replicas = &newReplicaCount
				copied, err = c.kubeClient.ReplicationControllers(copied.Namespace).Update(copied)
				return err
			}); err != nil {
				c.recorder.Eventf(config, v1.EventTypeWarning, "ReplicationControllerScaleFailed", "Failed to scale replication controler %q from %d to %d: %v", deployment.Name, *oldReplicaCount, newReplicaCount, err)
				return err
			}
			c.recorder.Eventf(config, v1.EventTypeNormal, "ReplicationControllerScaled", "Scaled replication controller %q from %d to %d", copied.Name, *oldReplicaCount, newReplicaCount)
			toAppend = copied
		}
		updatedDeployments = append(updatedDeployments, toAppend)
	}
	if err := c.cleanupOldDeployments(updatedDeployments, config); err != nil {
		c.recorder.Eventf(config, v1.EventTypeWarning, "ReplicationControllerCleanupFailed", "Couldn't clean up replication controllers: %v", err)
	}
	return c.updateStatus(config, updatedDeployments, true)
}
func (c *DeploymentConfigController) updateStatus(config *appsv1.DeploymentConfig, deployments []*v1.ReplicationController, updateObservedGeneration bool, additional ...appsv1.DeploymentCondition) error {
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
	newStatus := calculateStatus(config, deployments, updateObservedGeneration, additional...)
	if reflect.DeepEqual(newStatus, config.Status) {
		return nil
	}
	copied := config.DeepCopy()
	copied.Status = newStatus
	if _, err := c.appsClient.DeploymentConfigs(copied.Namespace).UpdateStatus(copied); err != nil {
		return err
	}
	klog.V(4).Infof(fmt.Sprintf("Updated status for DeploymentConfig: %s, ", appsutil.LabelForDeploymentConfig(config)) + fmt.Sprintf("replicas %d->%d (need %d), ", config.Status.Replicas, newStatus.Replicas, config.Spec.Replicas) + fmt.Sprintf("readyReplicas %d->%d, ", config.Status.ReadyReplicas, newStatus.ReadyReplicas) + fmt.Sprintf("availableReplicas %d->%d, ", config.Status.AvailableReplicas, newStatus.AvailableReplicas) + fmt.Sprintf("unavailableReplicas %d->%d, ", config.Status.UnavailableReplicas, newStatus.UnavailableReplicas) + fmt.Sprintf("sequence No: %v->%v", config.Status.ObservedGeneration, newStatus.ObservedGeneration))
	return nil
}
func (c *DeploymentConfigController) cancelRunningRollouts(config *appsv1.DeploymentConfig, existingDeployments []*v1.ReplicationController, cm *RCControllerRefManager) error {
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
	awaitingCancellations := false
	for i := range existingDeployments {
		deployment := existingDeployments[i]
		if appsutil.IsTerminatedDeployment(deployment) {
			continue
		}
		awaitingCancellations = true
		if appsutil.IsDeploymentCancelled(deployment) {
			continue
		}
		var updatedDeployment *v1.ReplicationController
		err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			rc, err := c.rcLister.ReplicationControllers(deployment.Namespace).Get(deployment.Name)
			if kapierrors.IsNotFound(err) {
				return nil
			}
			if err != nil {
				return err
			}
			isOurs, err := cm.ClaimReplicationController(rc)
			if err != nil {
				return fmt.Errorf("error while deploymentConfigController claiming the replication controller %s/%s: %v", rc.Namespace, rc.Name, err)
			}
			if !isOurs {
				return nil
			}
			copied := rc.DeepCopy()
			appsutil.SetCancelledByNewerDeployment(copied)
			updatedDeployment, err = c.kubeClient.ReplicationControllers(copied.Namespace).Update(copied)
			return err
		})
		if err != nil {
			c.recorder.Eventf(config, v1.EventTypeWarning, "DeploymentCancellationFailed", "Failed to cancel deployment %q superceded by version %d: %s", deployment.Name, config.Status.LatestVersion, err)
			return err
		}
		if updatedDeployment != nil {
			existingDeployments[i] = updatedDeployment
			c.recorder.Eventf(config, v1.EventTypeNormal, "DeploymentCancelled", "Cancelled deployment %q superceded by version %d", deployment.Name, config.Status.LatestVersion)
		}
	}
	if awaitingCancellations {
		c.recorder.Eventf(config, v1.EventTypeNormal, "DeploymentAwaitingCancellation", "Deployment of version %d awaiting cancellation of older running deployments", config.Status.LatestVersion)
		return fmt.Errorf("found previous inflight deployment for %s - requeuing", appsutil.LabelForDeploymentConfig(config))
	}
	return nil
}
func calculateStatus(config *appsv1.DeploymentConfig, rcs []*v1.ReplicationController, updateObservedGeneration bool, additional ...appsv1.DeploymentCondition) appsv1.DeploymentConfigStatus {
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
	latestReplicas := int32(0)
	latestExists, latestRC := appsutil.LatestDeploymentInfo(config, rcs)
	if !latestExists {
		latestRC = nil
	} else {
		latestReplicas = appsutil.GetStatusReplicaCountForDeployments([]*v1.ReplicationController{latestRC})
	}
	available := appsutil.GetAvailableReplicaCountForReplicationControllers(rcs)
	total := appsutil.GetReplicaCountForDeployments(rcs)
	unavailableReplicas := total - available
	if unavailableReplicas < 0 {
		unavailableReplicas = 0
	}
	generation := config.Status.ObservedGeneration
	if updateObservedGeneration {
		generation = config.Generation
	}
	status := appsv1.DeploymentConfigStatus{LatestVersion: config.Status.LatestVersion, Details: config.Status.Details, ObservedGeneration: generation, Replicas: appsutil.GetStatusReplicaCountForDeployments(rcs), UpdatedReplicas: latestReplicas, AvailableReplicas: available, ReadyReplicas: appsutil.GetReadyReplicaCountForReplicationControllers(rcs), UnavailableReplicas: unavailableReplicas, Conditions: config.Status.Conditions}
	updateConditions(config, &status, latestRC)
	for _, cond := range additional {
		appsutil.SetDeploymentCondition(&status, cond)
	}
	return status
}
func updateConditions(config *appsv1.DeploymentConfig, newStatus *appsv1.DeploymentConfigStatus, latestRC *v1.ReplicationController) {
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
	if newStatus.AvailableReplicas >= config.Spec.Replicas-appsutil.MaxUnavailable(config) && newStatus.AvailableReplicas > 0 {
		minAvailability := appsutil.NewDeploymentCondition(appsv1.DeploymentAvailable, v1.ConditionTrue, "", "Deployment config has minimum availability.")
		appsutil.SetDeploymentCondition(newStatus, *minAvailability)
	} else {
		noMinAvailability := appsutil.NewDeploymentCondition(appsv1.DeploymentAvailable, v1.ConditionFalse, "", "Deployment config does not have minimum availability.")
		appsutil.SetDeploymentCondition(newStatus, *noMinAvailability)
	}
	if latestRC != nil {
		switch appsutil.DeploymentStatusFor(latestRC) {
		case appsv1.DeploymentStatusPending:
			msg := fmt.Sprintf("replication controller %q is waiting for pod %q to run", latestRC.Name, appsutil.DeployerPodNameForDeployment(latestRC.Name))
			condition := appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionUnknown, "", msg)
			appsutil.SetDeploymentCondition(newStatus, *condition)
		case appsv1.DeploymentStatusRunning:
			if appsutil.IsProgressing(config, newStatus) {
				appsutil.RemoveDeploymentCondition(newStatus, appsv1.DeploymentProgressing)
				msg := fmt.Sprintf("replication controller %q is progressing", latestRC.Name)
				condition := appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionTrue, string(appsv1.ReplicationControllerUpdatedReason), msg)
				appsutil.SetDeploymentCondition(newStatus, *condition)
			}
		case appsv1.DeploymentStatusFailed:
			var condition *appsv1.DeploymentCondition
			if appsutil.IsDeploymentCancelled(latestRC) {
				msg := fmt.Sprintf("rollout of replication controller %q was cancelled", latestRC.Name)
				condition = appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionFalse, appsutil.CancelledRolloutReason, msg)
			} else {
				msg := fmt.Sprintf("replication controller %q has failed progressing", latestRC.Name)
				condition = appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionFalse, appsutil.TimedOutReason, msg)
			}
			appsutil.SetDeploymentCondition(newStatus, *condition)
		case appsv1.DeploymentStatusComplete:
			msg := fmt.Sprintf("replication controller %q successfully rolled out", latestRC.Name)
			condition := appsutil.NewDeploymentCondition(appsv1.DeploymentProgressing, v1.ConditionTrue, appsutil.NewRcAvailableReason, msg)
			appsutil.SetDeploymentCondition(newStatus, *condition)
		}
	}
}
func (c *DeploymentConfigController) handleErr(err error, key interface{}) {
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
	if err == nil {
		c.queue.Forget(key)
		return
	}
	if _, isFatal := err.(fatalError); isFatal {
		utilruntime.HandleError(err)
		c.queue.Forget(key)
		return
	}
	verbosity := klog.Level(2)
	if c.queue.NumRequeues(key) < maxRetryCount {
		if kapierrors.IsConflict(err) {
			verbosity = klog.Level(4)
		}
		klog.V(verbosity).Infof("Error syncing deployment config %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	klog.V(2).Infof("Dropping deployment config %q out of the queue: %v", key, err)
	c.queue.Forget(key)
}
func (c *DeploymentConfigController) cleanupOldDeployments(existingDeployments []*v1.ReplicationController, deploymentConfig *appsv1.DeploymentConfig) error {
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
	if deploymentConfig.Spec.RevisionHistoryLimit == nil {
		return nil
	}
	prunableDeployments := appsutil.DeploymentsForCleanup(deploymentConfig, existingDeployments)
	if len(prunableDeployments) <= int(*deploymentConfig.Spec.RevisionHistoryLimit) {
		return nil
	}
	deletionErrors := []error{}
	for i := 0; i < (len(prunableDeployments) - int(*deploymentConfig.Spec.RevisionHistoryLimit)); i++ {
		deployment := prunableDeployments[i]
		if *deployment.Spec.Replicas != 0 {
			continue
		}
		policy := metav1.DeletePropagationBackground
		err := c.kubeClient.ReplicationControllers(deployment.Namespace).Delete(deployment.Name, &metav1.DeleteOptions{PropagationPolicy: &policy})
		if err != nil && !kapierrors.IsNotFound(err) {
			deletionErrors = append(deletionErrors, err)
		}
	}
	return kutilerrors.NewAggregate(deletionErrors)
}
func triggerActivated(config *appsv1.DeploymentConfig, latestExists bool, latestDeployment *v1.ReplicationController) (bool, bool, error) {
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
	if config.Spec.Paused {
		return false, false, nil
	}
	imageTrigger := appsutil.HasImageChangeTrigger(config)
	configTrigger := appsutil.HasChangeTrigger(config)
	hasTrigger := imageTrigger || configTrigger
	if !hasTrigger {
		return false, false, nil
	}
	if appsutil.IsInitialDeployment(config) {
		hasAvailableImages := appsutil.HasLastTriggeredImage(config)
		if imageTrigger {
			if hasAvailableImages {
				klog.V(4).Infof("Rolling out initial deployment for %s/%s as it now have images available", config.Namespace, config.Name)
				appsutil.RecordConfigChangeCause(config)
				return true, false, nil
			}
			klog.V(4).Infof("Rolling out initial deployment for %s/%s deferred until its images are ready", config.Namespace, config.Name)
			return false, true, nil
		}
		if configTrigger {
			klog.V(4).Infof("Rolling out initial deployment for %s/%s", config.Namespace, config.Name)
			appsutil.RecordConfigChangeCause(config)
			return true, false, nil
		}
		return false, false, nil
	}
	if !latestExists {
		return false, false, nil
	}
	if latestDeployment == nil {
		return false, false, nil
	}
	if imageTrigger {
		if ok, imageNames := appsutil.HasUpdatedImages(config, latestDeployment); ok {
			klog.V(4).Infof("Rolling out #%d deployment for %s/%s caused by image changes (%s)", config.Status.LatestVersion+1, config.Namespace, config.Name, strings.Join(imageNames, ","))
			appsutil.RecordImageChangeCauses(config, imageNames)
			return true, false, nil
		}
	}
	if configTrigger {
		isLatest, changes, err := appsutil.HasLatestPodTemplate(config, latestDeployment)
		if err != nil {
			return false, false, fmt.Errorf("error while checking for latest pod template in replication controller: %v", err)
		}
		if !isLatest {
			klog.V(4).Infof("Rolling out #%d deployment for %s/%s caused by config change, diff: %s", config.Status.LatestVersion+1, config.Namespace, config.Name, changes)
			appsutil.RecordConfigChangeCause(config)
			return true, false, nil
		}
	}
	return false, false, nil
}
