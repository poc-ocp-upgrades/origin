package originpolymorphichelpers

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/kubectl/describe"
	odescribe "github.com/openshift/origin/pkg/oc/lib/describe"
)

func NewDescriberFn(delegate describe.DescriberFunc) describe.DescriberFunc {
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
	return func(restClientGetter genericclioptions.RESTClientGetter, mapping *meta.RESTMapping) (describe.Describer, error) {
		clientConfig, err := restClientGetter.ToRESTConfig()
		if err != nil {
			return nil, fmt.Errorf("unable to create client config %s: %v", mapping.GroupVersionKind.Kind, err)
		}
		kubeClient, err := kubernetes.NewForConfig(clientConfig)
		if err != nil {
			return nil, fmt.Errorf("unable to create client %s: %v", mapping.GroupVersionKind.Kind, err)
		}
		describer, ok := odescribe.DescriberFor(mapping.GroupVersionKind.GroupKind(), clientConfig, kubeClient, clientConfig.Host)
		if ok {
			return describer, nil
		}
		return delegate(restClientGetter, mapping)
	}
}
