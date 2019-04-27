package deployments

import (
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	appsv1 "github.com/openshift/api/apps/v1"
)

type Pruner interface {
	Prune(deleter DeploymentDeleter) error
}
type DeploymentDeleter interface {
	DeleteDeployment(deployment *corev1.ReplicationController) error
}
type pruner struct{ resolver Resolver }

var _ Pruner = &pruner{}

type PrunerOptions struct {
	KeepYoungerThan		time.Duration
	Orphans			bool
	KeepComplete		int
	KeepFailed		int
	DeploymentConfigs	[]*appsv1.DeploymentConfig
	Deployments		[]*corev1.ReplicationController
}

func NewPruner(options PrunerOptions) Pruner {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(1).Infof("Creating deployment pruner with keepYoungerThan=%v, orphans=%v, keepComplete=%v, keepFailed=%v", options.KeepYoungerThan, options.Orphans, options.KeepComplete, options.KeepFailed)
	filter := &andFilter{filterPredicates: []FilterPredicate{FilterDeploymentsPredicate, FilterZeroReplicaSize, NewFilterBeforePredicate(options.KeepYoungerThan)}}
	deployments := filter.Filter(options.Deployments)
	dataSet := NewDataSet(options.DeploymentConfigs, deployments)
	resolvers := []Resolver{}
	if options.Orphans {
		inactiveDeploymentStatus := []appsv1.DeploymentStatus{appsv1.DeploymentStatusComplete, appsv1.DeploymentStatusFailed}
		resolvers = append(resolvers, NewOrphanDeploymentResolver(dataSet, inactiveDeploymentStatus))
	}
	resolvers = append(resolvers, NewPerDeploymentConfigResolver(dataSet, options.KeepComplete, options.KeepFailed))
	return &pruner{resolver: &mergeResolver{resolvers: resolvers}}
}
func (p *pruner) Prune(deleter DeploymentDeleter) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	deployments, err := p.resolver.Resolve()
	if err != nil {
		return err
	}
	for _, deployment := range deployments {
		if err := deleter.DeleteDeployment(deployment); err != nil {
			return err
		}
	}
	return nil
}

type deploymentDeleter struct {
	deployments corev1client.ReplicationControllersGetter
}

var _ DeploymentDeleter = &deploymentDeleter{}

func NewDeploymentDeleter(deployments corev1client.ReplicationControllersGetter) DeploymentDeleter {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &deploymentDeleter{deployments: deployments}
}
func (p *deploymentDeleter) DeleteDeployment(deployment *corev1.ReplicationController) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Deleting deployment %q", deployment.Name)
	policy := metav1.DeletePropagationBackground
	return p.deployments.ReplicationControllers(deployment.Namespace).Delete(deployment.Name, &metav1.DeleteOptions{PropagationPolicy: &policy})
}
