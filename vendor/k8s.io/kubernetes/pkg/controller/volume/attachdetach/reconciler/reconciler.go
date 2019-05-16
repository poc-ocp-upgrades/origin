package reconciler

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/metrics"
	"k8s.io/kubernetes/pkg/controller/volume/attachdetach/statusupdater"
	kevents "k8s.io/kubernetes/pkg/kubelet/events"
	"k8s.io/kubernetes/pkg/util/goroutinemap/exponentialbackoff"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util/operationexecutor"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

type Reconciler interface{ Run(stopCh <-chan struct{}) }

func NewReconciler(loopPeriod time.Duration, maxWaitForUnmountDuration time.Duration, syncDuration time.Duration, disableReconciliationSync bool, desiredStateOfWorld cache.DesiredStateOfWorld, actualStateOfWorld cache.ActualStateOfWorld, attacherDetacher operationexecutor.OperationExecutor, nodeStatusUpdater statusupdater.NodeStatusUpdater, recorder record.EventRecorder) Reconciler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &reconciler{loopPeriod: loopPeriod, maxWaitForUnmountDuration: maxWaitForUnmountDuration, syncDuration: syncDuration, disableReconciliationSync: disableReconciliationSync, desiredStateOfWorld: desiredStateOfWorld, actualStateOfWorld: actualStateOfWorld, attacherDetacher: attacherDetacher, nodeStatusUpdater: nodeStatusUpdater, timeOfLastSync: time.Now(), recorder: recorder}
}

type reconciler struct {
	loopPeriod                time.Duration
	maxWaitForUnmountDuration time.Duration
	syncDuration              time.Duration
	desiredStateOfWorld       cache.DesiredStateOfWorld
	actualStateOfWorld        cache.ActualStateOfWorld
	attacherDetacher          operationexecutor.OperationExecutor
	nodeStatusUpdater         statusupdater.NodeStatusUpdater
	timeOfLastSync            time.Time
	disableReconciliationSync bool
	recorder                  record.EventRecorder
}

func (rc *reconciler) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	wait.Until(rc.reconciliationLoopFunc(), rc.loopPeriod, stopCh)
}
func (rc *reconciler) reconciliationLoopFunc() func() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func() {
		rc.reconcile()
		if rc.disableReconciliationSync {
			klog.V(5).Info("Skipping reconciling attached volumes still attached since it is disabled via the command line.")
		} else if rc.syncDuration < time.Second {
			klog.V(5).Info("Skipping reconciling attached volumes still attached since it is set to less than one second via the command line.")
		} else if time.Since(rc.timeOfLastSync) > rc.syncDuration {
			klog.V(5).Info("Starting reconciling attached volumes still attached")
			rc.sync()
		}
	}
}
func (rc *reconciler) sync() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer rc.updateSyncTime()
	rc.syncStates()
}
func (rc *reconciler) updateSyncTime() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	rc.timeOfLastSync = time.Now()
}
func (rc *reconciler) syncStates() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumesPerNode := rc.actualStateOfWorld.GetAttachedVolumesPerNode()
	rc.attacherDetacher.VerifyVolumesAreAttached(volumesPerNode, rc.actualStateOfWorld)
}
func (rc *reconciler) isMultiAttachForbidden(volumeSpec *volume.Spec) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volumeSpec.Volume != nil {
		if volumeSpec.Volume.AzureDisk != nil || volumeSpec.Volume.Cinder != nil {
			return true
		}
	}
	if volumeSpec.PersistentVolume != nil {
		if len(volumeSpec.PersistentVolume.Spec.AccessModes) == 0 {
			return false
		}
		for _, accessMode := range volumeSpec.PersistentVolume.Spec.AccessModes {
			if accessMode == v1.ReadWriteMany || accessMode == v1.ReadOnlyMany {
				return false
			}
		}
		return true
	}
	return false
}
func (rc *reconciler) reconcile() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, attachedVolume := range rc.actualStateOfWorld.GetAttachedVolumes() {
		if !rc.desiredStateOfWorld.VolumeExists(attachedVolume.VolumeName, attachedVolume.NodeName) {
			if rc.attacherDetacher.IsOperationPending(attachedVolume.VolumeName, "") {
				klog.V(10).Infof("Operation for volume %q is already running. Can't start detach for %q", attachedVolume.VolumeName, attachedVolume.NodeName)
				continue
			}
			elapsedTime, err := rc.actualStateOfWorld.SetDetachRequestTime(attachedVolume.VolumeName, attachedVolume.NodeName)
			if err != nil {
				klog.Errorf("Cannot trigger detach because it fails to set detach request time with error %v", err)
				continue
			}
			timeout := elapsedTime > rc.maxWaitForUnmountDuration
			if attachedVolume.MountedByNode && !timeout {
				klog.V(12).Infof(attachedVolume.GenerateMsgDetailed("Cannot detach volume because it is still mounted", ""))
				continue
			}
			err = rc.actualStateOfWorld.RemoveVolumeFromReportAsAttached(attachedVolume.VolumeName, attachedVolume.NodeName)
			if err != nil {
				klog.V(5).Infof("RemoveVolumeFromReportAsAttached failed while removing volume %q from node %q with: %v", attachedVolume.VolumeName, attachedVolume.NodeName, err)
			}
			err = rc.nodeStatusUpdater.UpdateNodeStatuses()
			if err != nil {
				klog.Errorf(attachedVolume.GenerateErrorDetailed("UpdateNodeStatuses failed while attempting to report volume as attached", err).Error())
				continue
			}
			klog.V(5).Infof(attachedVolume.GenerateMsgDetailed("Starting attacherDetacher.DetachVolume", ""))
			verifySafeToDetach := !timeout
			err = rc.attacherDetacher.DetachVolume(attachedVolume.AttachedVolume, verifySafeToDetach, rc.actualStateOfWorld)
			if err == nil {
				if !timeout {
					klog.Infof(attachedVolume.GenerateMsgDetailed("attacherDetacher.DetachVolume started", ""))
				} else {
					metrics.RecordForcedDetachMetric()
					klog.Warningf(attachedVolume.GenerateMsgDetailed("attacherDetacher.DetachVolume started", fmt.Sprintf("This volume is not safe to detach, but maxWaitForUnmountDuration %v expired, force detaching", rc.maxWaitForUnmountDuration)))
				}
			}
			if err != nil && !exponentialbackoff.IsExponentialBackoff(err) {
				klog.Errorf(attachedVolume.GenerateErrorDetailed("attacherDetacher.DetachVolume failed to start", err).Error())
			}
		}
	}
	rc.attachDesiredVolumes()
	err := rc.nodeStatusUpdater.UpdateNodeStatuses()
	if err != nil {
		klog.Warningf("UpdateNodeStatuses failed with: %v", err)
	}
}
func (rc *reconciler) attachDesiredVolumes() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, volumeToAttach := range rc.desiredStateOfWorld.GetVolumesToAttach() {
		if rc.actualStateOfWorld.VolumeNodeExists(volumeToAttach.VolumeName, volumeToAttach.NodeName) {
			if klog.V(5) {
				klog.Infof(volumeToAttach.GenerateMsgDetailed("Volume attached--touching", ""))
			}
			rc.actualStateOfWorld.ResetDetachRequestTime(volumeToAttach.VolumeName, volumeToAttach.NodeName)
			continue
		}
		if rc.attacherDetacher.IsOperationPending(volumeToAttach.VolumeName, "") {
			if klog.V(10) {
				klog.Infof("Operation for volume %q is already running. Can't start attach for %q", volumeToAttach.VolumeName, volumeToAttach.NodeName)
			}
			continue
		}
		if rc.isMultiAttachForbidden(volumeToAttach.VolumeSpec) {
			nodes := rc.actualStateOfWorld.GetNodesForVolume(volumeToAttach.VolumeName)
			if len(nodes) > 0 {
				if !volumeToAttach.MultiAttachErrorReported {
					rc.reportMultiAttachError(volumeToAttach, nodes)
					rc.desiredStateOfWorld.SetMultiAttachError(volumeToAttach.VolumeName, volumeToAttach.NodeName)
				}
				continue
			}
		}
		if klog.V(5) {
			klog.Infof(volumeToAttach.GenerateMsgDetailed("Starting attacherDetacher.AttachVolume", ""))
		}
		err := rc.attacherDetacher.AttachVolume(volumeToAttach.VolumeToAttach, rc.actualStateOfWorld)
		if err == nil {
			klog.Infof(volumeToAttach.GenerateMsgDetailed("attacherDetacher.AttachVolume started", ""))
		}
		if err != nil && !exponentialbackoff.IsExponentialBackoff(err) {
			klog.Errorf(volumeToAttach.GenerateErrorDetailed("attacherDetacher.AttachVolume failed to start", err).Error())
		}
	}
}
func (rc *reconciler) reportMultiAttachError(volumeToAttach cache.VolumeToAttach, nodes []types.NodeName) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	otherNodes := []types.NodeName{}
	otherNodesStr := []string{}
	for _, node := range nodes {
		if node != volumeToAttach.NodeName {
			otherNodes = append(otherNodes, node)
			otherNodesStr = append(otherNodesStr, string(node))
		}
	}
	pods := rc.desiredStateOfWorld.GetVolumePodsOnNodes(otherNodes, volumeToAttach.VolumeName)
	if len(pods) == 0 {
		simpleMsg, _ := volumeToAttach.GenerateMsg("Multi-Attach error", "Volume is already exclusively attached to one node and can't be attached to another")
		for _, pod := range volumeToAttach.ScheduledPods {
			rc.recorder.Eventf(pod, v1.EventTypeWarning, kevents.FailedAttachVolume, simpleMsg)
		}
		nodeList := strings.Join(otherNodesStr, ", ")
		detailedMsg := volumeToAttach.GenerateMsgDetailed("Multi-Attach error", fmt.Sprintf("Volume is already exclusively attached to node %s and can't be attached to another", nodeList))
		klog.Warningf(detailedMsg)
		return
	}
	for _, scheduledPod := range volumeToAttach.ScheduledPods {
		localPodNames := []string{}
		otherPods := 0
		for _, pod := range pods {
			if pod.Namespace == scheduledPod.Namespace {
				localPodNames = append(localPodNames, pod.Name)
			} else {
				otherPods++
			}
		}
		var msg string
		if len(localPodNames) > 0 {
			msg = fmt.Sprintf("Volume is already used by pod(s) %s", strings.Join(localPodNames, ", "))
			if otherPods > 0 {
				msg = fmt.Sprintf("%s and %d pod(s) in different namespaces", msg, otherPods)
			}
		} else {
			msg = fmt.Sprintf("Volume is already used by %d pod(s) in different namespaces", otherPods)
		}
		simpleMsg, _ := volumeToAttach.GenerateMsg("Multi-Attach error", msg)
		rc.recorder.Eventf(scheduledPod, v1.EventTypeWarning, kevents.FailedAttachVolume, simpleMsg)
	}
	podNames := []string{}
	for _, pod := range pods {
		podNames = append(podNames, pod.Namespace+"/"+pod.Name)
	}
	detailedMsg := volumeToAttach.GenerateMsgDetailed("Multi-Attach error", fmt.Sprintf("Volume is already used by pods %s on node %s", strings.Join(podNames, ", "), strings.Join(otherNodesStr, ", ")))
	klog.Warningf(detailedMsg)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
