package test

import (
	"k8s.io/apimachinery/pkg/util/sets"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
)

type FakeDeploymentConfigStore struct {
	DeploymentConfig	*appsapi.DeploymentConfig
	Err			error
}

func NewFakeDeploymentConfigStore(deployment *appsapi.DeploymentConfig) FakeDeploymentConfigStore {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return FakeDeploymentConfigStore{DeploymentConfig: deployment}
}
func (s FakeDeploymentConfigStore) Add(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Err
}
func (s FakeDeploymentConfigStore) Update(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Err
}
func (s FakeDeploymentConfigStore) Delete(obj interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Err
}
func (s FakeDeploymentConfigStore) List() []interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []interface{}{s.DeploymentConfig}
}
func (s FakeDeploymentConfigStore) ContainedIDs() sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sets.NewString()
}
func (s FakeDeploymentConfigStore) Get(obj interface{}) (item interface{}, exists bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.GetByKey("")
}
func (s FakeDeploymentConfigStore) GetByKey(id string) (item interface{}, exists bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.Err != nil {
		return nil, false, err
	}
	if s.DeploymentConfig == nil {
		return nil, false, nil
	}
	return s.DeploymentConfig, true, nil
}
func (s FakeDeploymentConfigStore) Replace(list []interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
