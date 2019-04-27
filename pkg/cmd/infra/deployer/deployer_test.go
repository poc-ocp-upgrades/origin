package deployer

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/origin/pkg/apps/strategy"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	appstest "github.com/openshift/origin/pkg/apps/util/test"
)

func TestDeployer_getDeploymentFail(t *testing.T) {
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
	deployer := &Deployer{strategyFor: func(config *appsv1.DeploymentConfig) (strategy.DeploymentStrategy, error) {
		t.Fatal("unexpected call")
		return nil, nil
	}, getDeployment: func(namespace, name string) (*corev1.ReplicationController, error) {
		return nil, fmt.Errorf("get error")
	}, getDeployments: func(namespace, configName string) (*corev1.ReplicationControllerList, error) {
		t.Fatal("unexpected call")
		return nil, nil
	}, scaler: &FakeScaler{}}
	err := deployer.Deploy("namespace", "name")
	if err == nil {
		t.Fatalf("expected an error")
	}
	t.Logf("got expected error: %v", err)
}
func TestDeployer_deployScenarios(t *testing.T) {
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
	mkd := func(version int64, status appsv1.DeploymentStatus, replicas int32, desired int32) *corev1.ReplicationController {
		deployment := mkdeployment(version, status)
		deployment.Spec.Replicas = &replicas
		if desired > 0 {
			deployment.Annotations[appsv1.DesiredReplicasAnnotation] = strconv.Itoa(int(desired))
		}
		return deployment
	}
	type scaleEvent struct {
		version	int64
		size	int32
	}
	scenarios := []struct {
		name		string
		deployments	[]*corev1.ReplicationController
		fromVersion	int64
		toVersion	int64
		scaleEvents	[]scaleEvent
	}{{"initial deployment", []*corev1.ReplicationController{mkd(1, appsv1.DeploymentStatusNew, 0, 3)}, 0, 1, []scaleEvent{}}, {"last deploy failed", []*corev1.ReplicationController{mkd(1, appsv1.DeploymentStatusComplete, 3, 0), mkd(2, appsv1.DeploymentStatusFailed, 1, 3), mkd(3, appsv1.DeploymentStatusNew, 0, 3)}, 1, 3, []scaleEvent{{2, 0}}}, {"sequential complete", []*corev1.ReplicationController{mkd(1, appsv1.DeploymentStatusComplete, 0, 0), mkd(2, appsv1.DeploymentStatusComplete, 3, 0), mkd(3, appsv1.DeploymentStatusNew, 0, 3)}, 2, 3, []scaleEvent{}}, {"sequential failure", []*corev1.ReplicationController{mkd(1, appsv1.DeploymentStatusFailed, 1, 3), mkd(2, appsv1.DeploymentStatusFailed, 1, 3), mkd(3, appsv1.DeploymentStatusNew, 0, 3)}, 0, 3, []scaleEvent{{1, 0}, {2, 0}}}, {"version mismatch", []*corev1.ReplicationController{mkd(1, appsv1.DeploymentStatusComplete, 0, 0), mkd(2, appsv1.DeploymentStatusNew, 3, 0), mkd(3, appsv1.DeploymentStatusComplete, 0, 3)}, 3, 2, []scaleEvent{}}}
	for _, s := range scenarios {
		t.Logf("executing scenario %s", s.name)
		findDeployment := func(version int64) *corev1.ReplicationController {
			for _, d := range s.deployments {
				if appsutil.DeploymentVersionFor(d) == version {
					return d
				}
			}
			return nil
		}
		var actualFrom, actualTo *corev1.ReplicationController
		to := findDeployment(s.toVersion)
		scaler := &FakeScaler{}
		deployer := &Deployer{out: &bytes.Buffer{}, errOut: &bytes.Buffer{}, strategyFor: func(config *appsv1.DeploymentConfig) (strategy.DeploymentStrategy, error) {
			return &testStrategy{deployFunc: func(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error {
				actualFrom = from
				actualTo = to
				return nil
			}}, nil
		}, getDeployment: func(namespace, name string) (*corev1.ReplicationController, error) {
			return to, nil
		}, getDeployments: func(namespace, configName string) (*corev1.ReplicationControllerList, error) {
			list := &corev1.ReplicationControllerList{}
			for _, d := range s.deployments {
				list.Items = append(list.Items, *d)
			}
			return list, nil
		}, scaler: scaler}
		err := deployer.Deploy(to.Namespace, to.Name)
		if s.toVersion < s.fromVersion {
			if err == nil {
				t.Fatalf("expected error when toVersion is older than newVersion")
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.fromVersion > 0 {
			if e, a := s.fromVersion, appsutil.DeploymentVersionFor(actualFrom); e != a {
				t.Fatalf("expected from.latestVersion %d, got %d", e, a)
			}
		}
		if e, a := s.toVersion, appsutil.DeploymentVersionFor(actualTo); e != a {
			t.Fatalf("expected to.latestVersion %d, got %d", e, a)
		}
		if e, a := len(s.scaleEvents), len(scaler.Events); e != a {
			t.Fatalf("expected %d scale events, got %d", e, a)
		}
		for _, expected := range s.scaleEvents {
			expectedTo := findDeployment(expected.version)
			expectedWasScaled := false
			for _, actual := range scaler.Events {
				if actual.Name != expectedTo.Name {
					continue
				}
				if e, a := uint(expected.size), actual.Size; e != a {
					t.Fatalf("expected version %d to be scaled to %d, got %d", expected.version, e, a)
				}
				expectedWasScaled = true
			}
			if !expectedWasScaled {
				t.Fatalf("expected version %d to be scaled to %d, but it wasn't scaled at all", expected.version, expected.size)
			}
		}
	}
}
func mkdeployment(version int64, status appsv1.DeploymentStatus) *corev1.ReplicationController {
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
	deployment, _ := appsutil.MakeDeployment(appstest.OkDeploymentConfig(version))
	deployment.Annotations[appsv1.DeploymentStatusAnnotation] = string(status)
	return deployment
}

type testStrategy struct {
	deployFunc func(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error
}

func (t *testStrategy) Deploy(from *corev1.ReplicationController, to *corev1.ReplicationController, desiredReplicas int) error {
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
	return t.deployFunc(from, to, desiredReplicas)
}

type FakeScaler struct{ Events []ScaleEvent }
type ScaleEvent struct {
	Name	string
	Size	uint
}

func (t *FakeScaler) Scale(namespace, name string, newSize uint, preconditions *kubectl.ScalePrecondition, retry, wait *kubectl.RetryParams, resource schema.GroupResource) error {
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
	t.Events = append(t.Events, ScaleEvent{name, newSize})
	return nil
}
func (t *FakeScaler) ScaleSimple(namespace, name string, preconditions *kubectl.ScalePrecondition, newSize uint, resource schema.GroupResource) (string, error) {
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
	return "", fmt.Errorf("unexpected call to ScaleSimple")
}
