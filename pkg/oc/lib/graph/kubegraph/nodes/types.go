package nodes

import (
	"fmt"
	"reflect"
	kappsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	osgraph "github.com/openshift/origin/pkg/oc/lib/graph/genericgraph"
)

var (
	ServiceNodeKind				= reflect.TypeOf(corev1.Service{}).Name()
	PodNodeKind				= reflect.TypeOf(corev1.Pod{}).Name()
	PodSpecNodeKind				= reflect.TypeOf(corev1.PodSpec{}).Name()
	PodTemplateSpecNodeKind			= reflect.TypeOf(corev1.PodTemplateSpec{}).Name()
	ReplicationControllerNodeKind		= reflect.TypeOf(corev1.ReplicationController{}).Name()
	ReplicationControllerSpecNodeKind	= reflect.TypeOf(corev1.ReplicationControllerSpec{}).Name()
	ServiceAccountNodeKind			= reflect.TypeOf(corev1.ServiceAccount{}).Name()
	SecretNodeKind				= reflect.TypeOf(corev1.Secret{}).Name()
	PersistentVolumeClaimNodeKind		= reflect.TypeOf(corev1.PersistentVolumeClaim{}).Name()
	JobNodeKind				= reflect.TypeOf(batchv1.Job{}).Name()
	JobSpecNodeKind				= reflect.TypeOf(batchv1.JobSpec{}).Name()
	HorizontalPodAutoscalerNodeKind		= reflect.TypeOf(autoscalingv1.HorizontalPodAutoscaler{}).Name()
	StatefulSetNodeKind			= reflect.TypeOf(kappsv1.StatefulSet{}).Name()
	StatefulSetSpecNodeKind			= reflect.TypeOf(kappsv1.StatefulSetSpec{}).Name()
	DeploymentNodeKind			= reflect.TypeOf(kappsv1.Deployment{}).Name()
	DeploymentSpecNodeKind			= reflect.TypeOf(kappsv1.DeploymentSpec{}).Name()
	ReplicaSetNodeKind			= reflect.TypeOf(kappsv1.ReplicaSet{}).Name()
	ReplicaSetSpecNodeKind			= reflect.TypeOf(kappsv1.ReplicaSetSpec{}).Name()
	DaemonSetNodeKind			= reflect.TypeOf(kappsv1.DaemonSet{}).Name()
)

func ServiceNodeName(o *corev1.Service) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(ServiceNodeKind, o)
}

type ServiceNode struct {
	osgraph.Node
	*corev1.Service
	IsFound	bool
}

func (n ServiceNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Service
}
func (n ServiceNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(ServiceNodeName(n.Service))
}
func (*ServiceNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ServiceNodeKind
}
func (n ServiceNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func PodNodeName(o *corev1.Pod) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(PodNodeKind, o)
}

type PodNode struct {
	osgraph.Node
	*corev1.Pod
}

func (n PodNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Pod
}
func (n PodNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(PodNodeName(n.Pod))
}
func (n PodNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodNodeName(n.Pod)
}
func (*PodNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodNodeKind
}
func PodSpecNodeName(o *corev1.PodSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", PodSpecNodeKind, ownerName))
}

type PodSpecNode struct {
	osgraph.Node
	*corev1.PodSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n PodSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.PodSpec
}
func (n PodSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n PodSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodSpecNodeName(n.PodSpec, n.OwnerName)
}
func (*PodSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodSpecNodeKind
}
func ReplicaSetNodeName(o *kappsv1.ReplicaSet) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(ReplicaSetNodeKind, o)
}

type ReplicaSetNode struct {
	osgraph.Node
	ReplicaSet	*kappsv1.ReplicaSet
	IsFound		bool
}

func (n ReplicaSetNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n ReplicaSetNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ReplicaSet
}
func (n ReplicaSetNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(ReplicaSetNodeName(n.ReplicaSet))
}
func (n ReplicaSetNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicaSetNodeName(n.ReplicaSet)
}
func (*ReplicaSetNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicaSetNodeKind
}
func ReplicaSetSpecNodeName(o *kappsv1.ReplicaSetSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", ReplicaSetSpecNodeKind, ownerName))
}

type ReplicaSetSpecNode struct {
	osgraph.Node
	ReplicaSetSpec	*kappsv1.ReplicaSetSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n ReplicaSetSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ReplicaSetSpec
}
func (n ReplicaSetSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n ReplicaSetSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicaSetSpecNodeName(n.ReplicaSetSpec, n.OwnerName)
}
func (*ReplicaSetSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicaSetSpecNodeKind
}
func ReplicationControllerNodeName(o *corev1.ReplicationController) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(ReplicationControllerNodeKind, o)
}

type ReplicationControllerNode struct {
	osgraph.Node
	ReplicationController	*corev1.ReplicationController
	IsFound			bool
}

func (n ReplicationControllerNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n ReplicationControllerNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ReplicationController
}
func (n ReplicationControllerNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(ReplicationControllerNodeName(n.ReplicationController))
}
func (n ReplicationControllerNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicationControllerNodeName(n.ReplicationController)
}
func (*ReplicationControllerNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicationControllerNodeKind
}
func ReplicationControllerSpecNodeName(o *corev1.ReplicationControllerSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", ReplicationControllerSpecNodeKind, ownerName))
}

type ReplicationControllerSpecNode struct {
	osgraph.Node
	ReplicationControllerSpec	*corev1.ReplicationControllerSpec
	Namespace			string
	OwnerName			osgraph.UniqueName
}

func (n ReplicationControllerSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ReplicationControllerSpec
}
func (n ReplicationControllerSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n ReplicationControllerSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicationControllerSpecNodeName(n.ReplicationControllerSpec, n.OwnerName)
}
func (*ReplicationControllerSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ReplicationControllerSpecNodeKind
}
func PodTemplateSpecNodeName(o *corev1.PodTemplateSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", PodTemplateSpecNodeKind, ownerName))
}

type PodTemplateSpecNode struct {
	osgraph.Node
	*corev1.PodTemplateSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n PodTemplateSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.PodTemplateSpec
}
func (n PodTemplateSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n PodTemplateSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodTemplateSpecNodeName(n.PodTemplateSpec, n.OwnerName)
}
func (*PodTemplateSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PodTemplateSpecNodeKind
}
func ServiceAccountNodeName(o *corev1.ServiceAccount) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(ServiceAccountNodeKind, o)
}

type ServiceAccountNode struct {
	osgraph.Node
	*corev1.ServiceAccount
	IsFound	bool
}

func (n ServiceAccountNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n ServiceAccountNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ServiceAccount
}
func (n ServiceAccountNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(ServiceAccountNodeName(n.ServiceAccount))
}
func (*ServiceAccountNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ServiceAccountNodeKind
}
func SecretNodeName(o *corev1.Secret) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(SecretNodeKind, o)
}

type SecretNode struct {
	osgraph.Node
	*corev1.Secret
	IsFound	bool
}

func (n SecretNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n SecretNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Secret
}
func (n SecretNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(SecretNodeName(n.Secret))
}
func (*SecretNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SecretNodeKind
}
func PersistentVolumeClaimNodeName(o *corev1.PersistentVolumeClaim) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(PersistentVolumeClaimNodeKind, o)
}

type PersistentVolumeClaimNode struct {
	osgraph.Node
	PersistentVolumeClaim	*corev1.PersistentVolumeClaim
	IsFound			bool
}

func (n PersistentVolumeClaimNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n PersistentVolumeClaimNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.PersistentVolumeClaim
}
func (n PersistentVolumeClaimNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (*PersistentVolumeClaimNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PersistentVolumeClaimNodeKind
}
func (n PersistentVolumeClaimNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PersistentVolumeClaimNodeName(n.PersistentVolumeClaim)
}
func JobNodeName(o *batchv1.Job) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(JobNodeKind, o)
}

type JobNode struct {
	osgraph.Node
	Job	*batchv1.Job
	IsFound	bool
}

func (n JobNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n JobNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Job
}
func (n JobNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(JobNodeName(n.Job))
}
func (n JobNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return JobNodeName(n.Job)
}
func (*JobNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return JobNodeKind
}
func JobSpecNodeName(o *batchv1.JobSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", JobSpecNodeKind, ownerName))
}

type JobSpecNode struct {
	osgraph.Node
	JobSpec		*batchv1.JobSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n JobSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.JobSpec
}
func (n JobSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n JobSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return JobSpecNodeName(n.JobSpec, n.OwnerName)
}
func (*JobSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return JobSpecNodeKind
}
func HorizontalPodAutoscalerNodeName(o *autoscalingv1.HorizontalPodAutoscaler) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(HorizontalPodAutoscalerNodeKind, o)
}

type HorizontalPodAutoscalerNode struct {
	osgraph.Node
	HorizontalPodAutoscaler	*autoscalingv1.HorizontalPodAutoscaler
}

func (n HorizontalPodAutoscalerNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.HorizontalPodAutoscaler
}
func (n HorizontalPodAutoscalerNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (*HorizontalPodAutoscalerNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return HorizontalPodAutoscalerNodeKind
}
func (n HorizontalPodAutoscalerNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return HorizontalPodAutoscalerNodeName(n.HorizontalPodAutoscaler)
}
func DeploymentNodeName(o *kappsv1.Deployment) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(DeploymentNodeKind, o)
}

type DeploymentNode struct {
	osgraph.Node
	Deployment	*kappsv1.Deployment
	IsFound		bool
}

func (n DeploymentNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n DeploymentNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Deployment
}
func (n DeploymentNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n DeploymentNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DeploymentNodeName(n.Deployment)
}
func (*DeploymentNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DeploymentNodeKind
}
func DeploymentSpecNodeName(o *kappsv1.DeploymentSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", DeploymentSpecNodeKind, ownerName))
}

type DeploymentSpecNode struct {
	osgraph.Node
	DeploymentSpec	*kappsv1.DeploymentSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n DeploymentSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.DeploymentSpec
}
func (n DeploymentSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n DeploymentSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DeploymentSpecNodeName(n.DeploymentSpec, n.OwnerName)
}
func (*DeploymentSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DeploymentSpecNodeKind
}
func StatefulSetNodeName(o *kappsv1.StatefulSet) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(StatefulSetNodeKind, o)
}

type StatefulSetNode struct {
	osgraph.Node
	StatefulSet	*kappsv1.StatefulSet
	IsFound		bool
}

func (n StatefulSetNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n StatefulSetNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.StatefulSet
}
func (n StatefulSetNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n StatefulSetNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return StatefulSetNodeName(n.StatefulSet)
}
func (*StatefulSetNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return StatefulSetNodeKind
}
func StatefulSetSpecNodeName(o *kappsv1.StatefulSetSpec, ownerName osgraph.UniqueName) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.UniqueName(fmt.Sprintf("%s|%v", StatefulSetSpecNodeKind, ownerName))
}

type StatefulSetSpecNode struct {
	osgraph.Node
	StatefulSetSpec	*kappsv1.StatefulSetSpec
	Namespace	string
	OwnerName	osgraph.UniqueName
}

func (n StatefulSetSpecNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.StatefulSetSpec
}
func (n StatefulSetSpecNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(n.UniqueName())
}
func (n StatefulSetSpecNode) UniqueName() osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return StatefulSetSpecNodeName(n.StatefulSetSpec, n.OwnerName)
}
func (*StatefulSetSpecNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return StatefulSetSpecNodeKind
}
func DaemonSetNodeName(o *kappsv1.DaemonSet) osgraph.UniqueName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return osgraph.GetUniqueRuntimeObjectNodeName(DaemonSetNodeKind, o)
}

type DaemonSetNode struct {
	osgraph.Node
	DaemonSet	*kappsv1.DaemonSet
	IsFound		bool
}

func (n DaemonSetNode) Found() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.IsFound
}
func (n DaemonSetNode) Object() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.DaemonSet
}
func (n DaemonSetNode) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(DaemonSetNodeName(n.DaemonSet))
}
func (*DaemonSetNode) Kind() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DaemonSetNodeKind
}
