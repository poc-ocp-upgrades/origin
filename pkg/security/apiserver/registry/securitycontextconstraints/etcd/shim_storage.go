package etcd

import (
	"k8s.io/apiserver/pkg/registry/rest"
)

func AddSCC(sccStorage *REST) func(restStorage map[string]rest.Storage) map[string]rest.Storage {
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
	return func(restStorage map[string]rest.Storage) map[string]rest.Storage {
		restStorage["securityContextConstraints"] = sccStorage
		return restStorage
	}
}
