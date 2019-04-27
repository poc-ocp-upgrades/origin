package server

import (
	"flag"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"k8s.io/apimachinery/pkg/util/wait"
)

func StartConfiguredMaster(masterConfig *configapi.MasterConfig) (string, error) {
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
	if v := flag.Lookup("test.v"); v == nil {
		panic("cannot be used outside of test code")
	}
	return StartConfiguredMasterWithOptions(masterConfig, wait.NeverStop)
}
func StartConfiguredMasterAPI(masterConfig *configapi.MasterConfig) (string, error) {
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
	if v := flag.Lookup("test.v"); v == nil {
		panic("cannot be used outside of test code")
	}
	if masterConfig.KubernetesMasterConfig.ControllerArguments == nil {
		masterConfig.KubernetesMasterConfig.ControllerArguments = map[string][]string{}
	}
	masterConfig.KubernetesMasterConfig.ControllerArguments["controllers"] = append(masterConfig.KubernetesMasterConfig.ControllerArguments["controllers"], "serviceaccount-token", "clusterrole-aggregation")
	return StartConfiguredMasterWithOptions(masterConfig, wait.NeverStop)
}
func StartTestMaster() (*configapi.MasterConfig, string, error) {
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
	master, err := DefaultMasterOptions()
	if err != nil {
		return nil, "", err
	}
	adminKubeConfigFile, err := StartConfiguredMaster(master)
	return master, adminKubeConfigFile, err
}
func StartTestMasterAPI() (*configapi.MasterConfig, string, error) {
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
	master, err := DefaultMasterOptions()
	if err != nil {
		return nil, "", err
	}
	adminKubeConfigFile, err := StartConfiguredMasterAPI(master)
	return master, adminKubeConfigFile, err
}
