package openshiftapiserver

import (
	"strconv"
	"k8s.io/apimachinery/pkg/runtime/schema"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/server/options"
	apiserverstorage "k8s.io/apiserver/pkg/server/storage"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/origin/pkg/cmd/openshift-apiserver/openshiftapiserver/configprocessing"
)

func NewRESTOptionsGetter(startingFlags map[string][]string, storageConfig configv1.EtcdStorageConfig) (genericregistry.RESTOptionsGetter, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	targetRAMMB := 0
	if targetRamString := startingFlags["target-ram-mb"]; len(targetRamString) == 1 {
		targetRAMMB, err = strconv.Atoi(targetRamString[0])
		if err != nil {
			return nil, err
		}
	}
	etcdOptions, err := configprocessing.GetEtcdOptions(startingFlags, storageConfig, newHeuristicWatchCacheSizes(targetRAMMB))
	storageFactory := apiserverstorage.NewDefaultStorageFactory(etcdOptions.StorageConfig, etcdOptions.DefaultStorageMediaType, legacyscheme.Codecs, apiserverstorage.NewDefaultResourceEncodingConfig(legacyscheme.Scheme), &serverstorage.ResourceConfig{}, specialDefaultResourcePrefixes)
	restOptionsGetter := &options.StorageFactoryRestOptionsFactory{Options: *etcdOptions, StorageFactory: storageFactory}
	return restOptionsGetter, nil
}
func newHeuristicWatchCacheSizes(expectedRAMCapacityMB int) map[schema.GroupResource]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterSize := expectedRAMCapacityMB / 60
	watchCacheSizes := make(map[schema.GroupResource]int)
	watchCacheSizes[schema.GroupResource{Group: "network.openshift.io", Resource: "hostsubnets"}] = maxInt(5*clusterSize, 100)
	watchCacheSizes[schema.GroupResource{Group: "network.openshift.io", Resource: "netnamespaces"}] = maxInt(5*clusterSize, 100)
	watchCacheSizes[schema.GroupResource{Group: "network.openshift.io", Resource: "egressnetworkpolicies"}] = maxInt(10*clusterSize, 100)
	return watchCacheSizes
}
func maxInt(a, b int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a > b {
		return a
	}
	return b
}

var specialDefaultResourcePrefixes = map[schema.GroupResource]string{{Resource: "clusterpolicies"}: "authorization/cluster/policies", {Resource: "clusterpolicies", Group: "authorization.openshift.io"}: "authorization/cluster/policies", {Resource: "clusterpolicybindings"}: "authorization/cluster/policybindings", {Resource: "clusterpolicybindings", Group: "authorization.openshift.io"}: "authorization/cluster/policybindings", {Resource: "policies"}: "authorization/local/policies", {Resource: "policies", Group: "authorization.openshift.io"}: "authorization/local/policies", {Resource: "policybindings"}: "authorization/local/policybindings", {Resource: "policybindings", Group: "authorization.openshift.io"}: "authorization/local/policybindings", {Resource: "oauthaccesstokens"}: "oauth/accesstokens", {Resource: "oauthaccesstokens", Group: "oauth.openshift.io"}: "oauth/accesstokens", {Resource: "oauthauthorizetokens"}: "oauth/authorizetokens", {Resource: "oauthauthorizetokens", Group: "oauth.openshift.io"}: "oauth/authorizetokens", {Resource: "oauthclients"}: "oauth/clients", {Resource: "oauthclients", Group: "oauth.openshift.io"}: "oauth/clients", {Resource: "oauthclientauthorizations"}: "oauth/clientauthorizations", {Resource: "oauthclientauthorizations", Group: "oauth.openshift.io"}: "oauth/clientauthorizations", {Resource: "identities"}: "useridentities", {Resource: "identities", Group: "user.openshift.io"}: "useridentities", {Resource: "clusternetworks"}: "registry/sdnnetworks", {Resource: "clusternetworks", Group: "network.openshift.io"}: "registry/sdnnetworks", {Resource: "egressnetworkpolicies"}: "registry/egressnetworkpolicy", {Resource: "egressnetworkpolicies", Group: "network.openshift.io"}: "registry/egressnetworkpolicy", {Resource: "hostsubnets"}: "registry/sdnsubnets", {Resource: "hostsubnets", Group: "network.openshift.io"}: "registry/sdnsubnets", {Resource: "netnamespaces"}: "registry/sdnnetnamespaces", {Resource: "netnamespaces", Group: "network.openshift.io"}: "registry/sdnnetnamespaces"}
