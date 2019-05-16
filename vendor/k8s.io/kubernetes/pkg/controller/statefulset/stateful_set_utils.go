package statefulset

import (
	"bytes"
	"encoding/json"
	"fmt"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
	podutil "k8s.io/kubernetes/pkg/api/v1/pod"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/history"
	"regexp"
	"strconv"
)

const maxUpdateRetries = 10

var updateConflictError = fmt.Errorf("aborting update after %d attempts", maxUpdateRetries)
var patchCodec = scheme.Codecs.LegacyCodec(apps.SchemeGroupVersion)

type overlappingStatefulSets []*apps.StatefulSet

func (o overlappingStatefulSets) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(o)
}
func (o overlappingStatefulSets) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o[i], o[j] = o[j], o[i]
}
func (o overlappingStatefulSets) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o[i].CreationTimestamp.Equal(&o[j].CreationTimestamp) {
		return o[i].Name < o[j].Name
	}
	return o[i].CreationTimestamp.Before(&o[j].CreationTimestamp)
}

var statefulPodRegex = regexp.MustCompile("(.*)-([0-9]+)$")

func getParentNameAndOrdinal(pod *v1.Pod) (string, int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parent := ""
	ordinal := -1
	subMatches := statefulPodRegex.FindStringSubmatch(pod.Name)
	if len(subMatches) < 3 {
		return parent, ordinal
	}
	parent = subMatches[1]
	if i, err := strconv.ParseInt(subMatches[2], 10, 32); err == nil {
		ordinal = int(i)
	}
	return parent, ordinal
}
func getParentName(pod *v1.Pod) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parent, _ := getParentNameAndOrdinal(pod)
	return parent
}
func getOrdinal(pod *v1.Pod) int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, ordinal := getParentNameAndOrdinal(pod)
	return ordinal
}
func getPodName(set *apps.StatefulSet, ordinal int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s-%d", set.Name, ordinal)
}
func getPersistentVolumeClaimName(set *apps.StatefulSet, claim *v1.PersistentVolumeClaim, ordinal int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s-%s-%d", claim.Name, set.Name, ordinal)
}
func isMemberOf(set *apps.StatefulSet, pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return getParentName(pod) == set.Name
}
func identityMatches(set *apps.StatefulSet, pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parent, ordinal := getParentNameAndOrdinal(pod)
	return ordinal >= 0 && set.Name == parent && pod.Name == getPodName(set, ordinal) && pod.Namespace == set.Namespace && pod.Labels[apps.StatefulSetPodNameLabel] == pod.Name
}
func storageMatches(set *apps.StatefulSet, pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ordinal := getOrdinal(pod)
	if ordinal < 0 {
		return false
	}
	volumes := make(map[string]v1.Volume, len(pod.Spec.Volumes))
	for _, volume := range pod.Spec.Volumes {
		volumes[volume.Name] = volume
	}
	for _, claim := range set.Spec.VolumeClaimTemplates {
		volume, found := volumes[claim.Name]
		if !found || volume.VolumeSource.PersistentVolumeClaim == nil || volume.VolumeSource.PersistentVolumeClaim.ClaimName != getPersistentVolumeClaimName(set, &claim, ordinal) {
			return false
		}
	}
	return true
}
func getPersistentVolumeClaims(set *apps.StatefulSet, pod *v1.Pod) map[string]v1.PersistentVolumeClaim {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ordinal := getOrdinal(pod)
	templates := set.Spec.VolumeClaimTemplates
	claims := make(map[string]v1.PersistentVolumeClaim, len(templates))
	for i := range templates {
		claim := templates[i]
		claim.Name = getPersistentVolumeClaimName(set, &claim, ordinal)
		claim.Namespace = set.Namespace
		claim.Labels = set.Spec.Selector.MatchLabels
		claims[templates[i].Name] = claim
	}
	return claims
}
func updateStorage(set *apps.StatefulSet, pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	currentVolumes := pod.Spec.Volumes
	claims := getPersistentVolumeClaims(set, pod)
	newVolumes := make([]v1.Volume, 0, len(claims))
	for name, claim := range claims {
		newVolumes = append(newVolumes, v1.Volume{Name: name, VolumeSource: v1.VolumeSource{PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: claim.Name, ReadOnly: false}}})
	}
	for i := range currentVolumes {
		if _, ok := claims[currentVolumes[i].Name]; !ok {
			newVolumes = append(newVolumes, currentVolumes[i])
		}
	}
	pod.Spec.Volumes = newVolumes
}
func initIdentity(set *apps.StatefulSet, pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	updateIdentity(set, pod)
	pod.Spec.Hostname = pod.Name
	pod.Spec.Subdomain = set.Spec.ServiceName
}
func updateIdentity(set *apps.StatefulSet, pod *v1.Pod) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod.Name = getPodName(set, getOrdinal(pod))
	pod.Namespace = set.Namespace
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}
	pod.Labels[apps.StatefulSetPodNameLabel] = pod.Name
}
func isRunningAndReady(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod.Status.Phase == v1.PodRunning && podutil.IsPodReady(pod)
}
func isCreated(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod.Status.Phase != ""
}
func isFailed(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod.Status.Phase == v1.PodFailed
}
func isTerminating(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return pod.DeletionTimestamp != nil
}
func isHealthy(pod *v1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return isRunningAndReady(pod) && !isTerminating(pod)
}
func allowsBurst(set *apps.StatefulSet) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return set.Spec.PodManagementPolicy == apps.ParallelPodManagement
}
func setPodRevision(pod *v1.Pod, revision string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}
	pod.Labels[apps.StatefulSetRevisionLabel] = revision
}
func getPodRevision(pod *v1.Pod) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod.Labels == nil {
		return ""
	}
	return pod.Labels[apps.StatefulSetRevisionLabel]
}
func newStatefulSetPod(set *apps.StatefulSet, ordinal int) *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod, _ := controller.GetPodFromTemplate(&set.Spec.Template, set, metav1.NewControllerRef(set, controllerKind))
	pod.Name = getPodName(set, ordinal)
	initIdentity(set, pod)
	updateStorage(set, pod)
	return pod
}
func newVersionedStatefulSetPod(currentSet, updateSet *apps.StatefulSet, currentRevision, updateRevision string, ordinal int) *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if currentSet.Spec.UpdateStrategy.Type == apps.RollingUpdateStatefulSetStrategyType && (currentSet.Spec.UpdateStrategy.RollingUpdate == nil && ordinal < int(currentSet.Status.CurrentReplicas)) || (currentSet.Spec.UpdateStrategy.RollingUpdate != nil && ordinal < int(*currentSet.Spec.UpdateStrategy.RollingUpdate.Partition)) {
		pod := newStatefulSetPod(currentSet, ordinal)
		setPodRevision(pod, currentRevision)
		return pod
	}
	pod := newStatefulSetPod(updateSet, ordinal)
	setPodRevision(pod, updateRevision)
	return pod
}
func Match(ss *apps.StatefulSet, history *apps.ControllerRevision) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	patch, err := getPatch(ss)
	if err != nil {
		return false, err
	}
	return bytes.Equal(patch, history.Data.Raw), nil
}
func getPatch(set *apps.StatefulSet) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	str, err := runtime.Encode(patchCodec, set)
	if err != nil {
		return nil, err
	}
	var raw map[string]interface{}
	json.Unmarshal([]byte(str), &raw)
	objCopy := make(map[string]interface{})
	specCopy := make(map[string]interface{})
	spec := raw["spec"].(map[string]interface{})
	template := spec["template"].(map[string]interface{})
	specCopy["template"] = template
	template["$patch"] = "replace"
	objCopy["spec"] = specCopy
	patch, err := json.Marshal(objCopy)
	return patch, err
}
func newRevision(set *apps.StatefulSet, revision int64, collisionCount *int32) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	patch, err := getPatch(set)
	if err != nil {
		return nil, err
	}
	cr, err := history.NewControllerRevision(set, controllerKind, set.Spec.Template.Labels, runtime.RawExtension{Raw: patch}, revision, collisionCount)
	if err != nil {
		return nil, err
	}
	if cr.ObjectMeta.Annotations == nil {
		cr.ObjectMeta.Annotations = make(map[string]string)
	}
	for key, value := range set.Annotations {
		cr.ObjectMeta.Annotations[key] = value
	}
	return cr, nil
}
func ApplyRevision(set *apps.StatefulSet, revision *apps.ControllerRevision) (*apps.StatefulSet, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clone := set.DeepCopy()
	patched, err := strategicpatch.StrategicMergePatch([]byte(runtime.EncodeOrDie(patchCodec, clone)), revision.Data.Raw, clone)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(patched, clone)
	if err != nil {
		return nil, err
	}
	return clone, nil
}
func nextRevision(revisions []*apps.ControllerRevision) int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	count := len(revisions)
	if count <= 0 {
		return 1
	}
	return revisions[count-1].Revision + 1
}
func inconsistentStatus(set *apps.StatefulSet, status *apps.StatefulSetStatus) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return status.ObservedGeneration > set.Status.ObservedGeneration || status.Replicas != set.Status.Replicas || status.CurrentReplicas != set.Status.CurrentReplicas || status.ReadyReplicas != set.Status.ReadyReplicas || status.UpdatedReplicas != set.Status.UpdatedReplicas || status.CurrentRevision != set.Status.CurrentRevision || status.UpdateRevision != set.Status.UpdateRevision
}
func completeRollingUpdate(set *apps.StatefulSet, status *apps.StatefulSetStatus) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if set.Spec.UpdateStrategy.Type == apps.RollingUpdateStatefulSetStrategyType && status.UpdatedReplicas == status.Replicas && status.ReadyReplicas == status.Replicas {
		status.CurrentReplicas = status.UpdatedReplicas
		status.CurrentRevision = status.UpdateRevision
	}
}

type ascendingOrdinal []*v1.Pod

func (ao ascendingOrdinal) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(ao)
}
func (ao ascendingOrdinal) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ao[i], ao[j] = ao[j], ao[i]
}
func (ao ascendingOrdinal) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return getOrdinal(ao[i]) < getOrdinal(ao[j])
}
