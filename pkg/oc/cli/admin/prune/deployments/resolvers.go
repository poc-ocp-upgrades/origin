package deployments

import (
	"sort"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	appsv1 "github.com/openshift/api/apps/v1"
	appsutil "github.com/openshift/origin/pkg/apps/util"
)

type Resolver interface {
	Resolve() ([]*corev1.ReplicationController, error)
}
type mergeResolver struct{ resolvers []Resolver }

func (m *mergeResolver) Resolve() ([]*corev1.ReplicationController, error) {
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
	results := []*corev1.ReplicationController{}
	for _, resolver := range m.resolvers {
		items, err := resolver.Resolve()
		if err != nil {
			return nil, err
		}
		results = append(results, items...)
	}
	return results, nil
}
func NewOrphanDeploymentResolver(dataSet DataSet, deploymentStatusFilter []appsv1.DeploymentStatus) Resolver {
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
	filter := sets.NewString()
	for _, deploymentStatus := range deploymentStatusFilter {
		filter.Insert(string(deploymentStatus))
	}
	return &orphanDeploymentResolver{dataSet: dataSet, deploymentStatusFilter: filter}
}

type orphanDeploymentResolver struct {
	dataSet			DataSet
	deploymentStatusFilter	sets.String
}

func (o *orphanDeploymentResolver) Resolve() ([]*corev1.ReplicationController, error) {
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
	deployments, err := o.dataSet.ListDeployments()
	if err != nil {
		return nil, err
	}
	results := []*corev1.ReplicationController{}
	for _, deployment := range deployments {
		deploymentStatus := appsutil.DeploymentStatusFor(deployment)
		if !o.deploymentStatusFilter.Has(string(deploymentStatus)) {
			continue
		}
		_, exists, _ := o.dataSet.GetDeploymentConfig(deployment)
		if !exists {
			results = append(results, deployment)
		}
	}
	return results, nil
}

type perDeploymentConfigResolver struct {
	dataSet		DataSet
	keepComplete	int
	keepFailed	int
}

func NewPerDeploymentConfigResolver(dataSet DataSet, keepComplete int, keepFailed int) Resolver {
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
	return &perDeploymentConfigResolver{dataSet: dataSet, keepComplete: keepComplete, keepFailed: keepFailed}
}

type ByMostRecent []*corev1.ReplicationController

func (s ByMostRecent) Len() int {
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
	return len(s)
}
func (s ByMostRecent) Swap(i, j int) {
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
	s[i], s[j] = s[j], s[i]
}
func (s ByMostRecent) Less(i, j int) bool {
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
	return !s[i].CreationTimestamp.Before(&s[j].CreationTimestamp)
}
func (o *perDeploymentConfigResolver) Resolve() ([]*corev1.ReplicationController, error) {
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
	deploymentConfigs, err := o.dataSet.ListDeploymentConfigs()
	if err != nil {
		return nil, err
	}
	completeStates := sets.NewString(string(appsv1.DeploymentStatusComplete))
	failedStates := sets.NewString(string(appsv1.DeploymentStatusFailed))
	results := []*corev1.ReplicationController{}
	for _, deploymentConfig := range deploymentConfigs {
		deployments, err := o.dataSet.ListDeploymentsByDeploymentConfig(deploymentConfig)
		if err != nil {
			return nil, err
		}
		completeDeployments, failedDeployments := []*corev1.ReplicationController{}, []*corev1.ReplicationController{}
		for _, deployment := range deployments {
			status := appsutil.DeploymentStatusFor(deployment)
			if completeStates.Has(string(status)) {
				completeDeployments = append(completeDeployments, deployment)
			} else if failedStates.Has(string(status)) {
				failedDeployments = append(failedDeployments, deployment)
			}
		}
		sort.Sort(ByMostRecent(completeDeployments))
		sort.Sort(ByMostRecent(failedDeployments))
		if o.keepComplete >= 0 && o.keepComplete < len(completeDeployments) {
			results = append(results, completeDeployments[o.keepComplete:]...)
		}
		if o.keepFailed >= 0 && o.keepFailed < len(failedDeployments) {
			results = append(results, failedDeployments[o.keepFailed:]...)
		}
	}
	return results, nil
}
