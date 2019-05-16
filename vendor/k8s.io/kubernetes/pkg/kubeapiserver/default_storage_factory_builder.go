package kubeapiserver

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	serveroptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/server/options/encryptionconfig"
	"k8s.io/apiserver/pkg/server/resourceconfig"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/admissionregistration"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/batch"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/events"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/networking"
	"k8s.io/kubernetes/pkg/apis/policy"
	apisstorage "k8s.io/kubernetes/pkg/apis/storage"
	kubeapiserveroptions "k8s.io/kubernetes/pkg/kubeapiserver/options"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var SpecialDefaultResourcePrefixes = map[schema.GroupResource]string{{Group: "", Resource: "replicationcontrollers"}: "controllers", {Group: "", Resource: "endpoints"}: "services/endpoints", {Group: "", Resource: "nodes"}: "minions", {Group: "", Resource: "services"}: "services/specs", {Group: "extensions", Resource: "ingresses"}: "ingress", {Group: "extensions", Resource: "podsecuritypolicies"}: "podsecuritypolicy", {Group: "policy", Resource: "podsecuritypolicies"}: "podsecuritypolicy"}

func NewStorageFactoryConfig() *StorageFactoryConfig {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &StorageFactoryConfig{Serializer: legacyscheme.Codecs, DefaultResourceEncoding: serverstorage.NewDefaultResourceEncodingConfig(legacyscheme.Scheme), ResourceEncodingOverrides: []schema.GroupVersionResource{batch.Resource("cronjobs").WithVersion("v1beta1"), apisstorage.Resource("volumeattachments").WithVersion("v1beta1"), admissionregistration.Resource("initializerconfigurations").WithVersion("v1alpha1")}}
}

type StorageFactoryConfig struct {
	StorageConfig                    storagebackend.Config
	ApiResourceConfig                *serverstorage.ResourceConfig
	DefaultResourceEncoding          *serverstorage.DefaultResourceEncodingConfig
	DefaultStorageMediaType          string
	Serializer                       runtime.StorageSerializer
	StorageEncodingOverrides         map[string]schema.GroupVersion
	ResourceEncodingOverrides        []schema.GroupVersionResource
	EtcdServersOverrides             []string
	EncryptionProviderConfigFilepath string
}

func (c *StorageFactoryConfig) Complete(etcdOptions *serveroptions.EtcdOptions, serializationOptions *kubeapiserveroptions.StorageSerializationOptions) (*completedStorageFactoryConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	storageGroupsToEncodingVersion, err := serializationOptions.StorageGroupsToEncodingVersion()
	if err != nil {
		return nil, fmt.Errorf("error generating storage version map: %s", err)
	}
	c.StorageEncodingOverrides = storageGroupsToEncodingVersion
	c.StorageConfig = etcdOptions.StorageConfig
	c.DefaultStorageMediaType = etcdOptions.DefaultStorageMediaType
	c.EtcdServersOverrides = etcdOptions.EtcdServersOverrides
	c.EncryptionProviderConfigFilepath = etcdOptions.EncryptionProviderConfigFilepath
	return &completedStorageFactoryConfig{c}, nil
}

type completedStorageFactoryConfig struct{ *StorageFactoryConfig }

func (c *completedStorageFactoryConfig) New() (*serverstorage.DefaultStorageFactory, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resourceEncodingConfig := resourceconfig.MergeGroupEncodingConfigs(c.DefaultResourceEncoding, c.StorageEncodingOverrides)
	resourceEncodingConfig = resourceconfig.MergeResourceEncodingConfigs(resourceEncodingConfig, c.ResourceEncodingOverrides)
	storageFactory := serverstorage.NewDefaultStorageFactory(c.StorageConfig, c.DefaultStorageMediaType, c.Serializer, resourceEncodingConfig, c.ApiResourceConfig, SpecialDefaultResourcePrefixes)
	storageFactory.AddCohabitatingResources(networking.Resource("networkpolicies"), extensions.Resource("networkpolicies"))
	storageFactory.AddCohabitatingResources(apps.Resource("deployments"), extensions.Resource("deployments"))
	storageFactory.AddCohabitatingResources(apps.Resource("daemonsets"), extensions.Resource("daemonsets"))
	storageFactory.AddCohabitatingResources(apps.Resource("replicasets"), extensions.Resource("replicasets"))
	storageFactory.AddCohabitatingResources(api.Resource("events"), events.Resource("events"))
	storageFactory.AddCohabitatingResources(policy.Resource("podsecuritypolicies"), extensions.Resource("podsecuritypolicies"))
	for _, override := range c.EtcdServersOverrides {
		tokens := strings.Split(override, "#")
		apiresource := strings.Split(tokens[0], "/")
		group := apiresource[0]
		resource := apiresource[1]
		groupResource := schema.GroupResource{Group: group, Resource: resource}
		servers := strings.Split(tokens[1], ";")
		storageFactory.SetEtcdLocation(groupResource, servers)
	}
	if len(c.EncryptionProviderConfigFilepath) != 0 {
		transformerOverrides, err := encryptionconfig.GetTransformerOverrides(c.EncryptionProviderConfigFilepath)
		if err != nil {
			return nil, err
		}
		for groupResource, transformer := range transformerOverrides {
			storageFactory.SetTransformer(groupResource, transformer)
		}
	}
	return storageFactory, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
