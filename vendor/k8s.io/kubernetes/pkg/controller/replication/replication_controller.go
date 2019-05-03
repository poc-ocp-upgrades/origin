package replication

import (
 "k8s.io/api/core/v1"
 coreinformers "k8s.io/client-go/informers/core/v1"
 clientset "k8s.io/client-go/kubernetes"
 "k8s.io/client-go/kubernetes/scheme"
 v1core "k8s.io/client-go/kubernetes/typed/core/v1"
 "k8s.io/client-go/tools/record"
 "k8s.io/klog"
 "k8s.io/kubernetes/pkg/controller"
 "k8s.io/kubernetes/pkg/controller/replicaset"
)

const (
 BurstReplicas = replicaset.BurstReplicas
)

type ReplicationManager struct {
 replicaset.ReplicaSetController
}

func NewReplicationManager(podInformer coreinformers.PodInformer, rcInformer coreinformers.ReplicationControllerInformer, kubeClient clientset.Interface, burstReplicas int) *ReplicationManager {
 _logClusterCodePath()
 defer _logClusterCodePath()
 eventBroadcaster := record.NewBroadcaster()
 eventBroadcaster.StartLogging(klog.Infof)
 eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
 return &ReplicationManager{*replicaset.NewBaseController(informerAdapter{rcInformer}, podInformer, clientsetAdapter{kubeClient}, burstReplicas, v1.SchemeGroupVersion.WithKind("ReplicationController"), "replication_controller", "replicationmanager", podControlAdapter{controller.RealPodControl{KubeClient: kubeClient, Recorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "replication-controller"})}})}
}
