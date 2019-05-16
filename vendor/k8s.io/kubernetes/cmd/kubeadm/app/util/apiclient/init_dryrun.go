package apiclient

import (
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	core "k8s.io/client-go/testing"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"net"
	"strings"
)

type InitDryRunGetter struct {
	masterName    string
	serviceSubnet string
}

var _ DryRunGetter = &InitDryRunGetter{}

func NewInitDryRunGetter(masterName string, serviceSubnet string) *InitDryRunGetter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &InitDryRunGetter{masterName: masterName, serviceSubnet: serviceSubnet}
}
func (idr *InitDryRunGetter) HandleGetAction(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	funcs := []func(core.GetAction) (bool, runtime.Object, error){idr.handleKubernetesService, idr.handleGetNode, idr.handleSystemNodesClusterRoleBinding, idr.handleGetBootstrapToken}
	for _, f := range funcs {
		handled, obj, err := f(action)
		if handled {
			return handled, obj, err
		}
	}
	return false, nil, nil
}
func (idr *InitDryRunGetter) HandleListAction(action core.ListAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false, nil, nil
}
func (idr *InitDryRunGetter) handleKubernetesService(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if action.GetName() != "kubernetes" || action.GetNamespace() != metav1.NamespaceDefault || action.GetResource().Resource != "services" {
		return false, nil, nil
	}
	_, svcSubnet, err := net.ParseCIDR(idr.serviceSubnet)
	if err != nil {
		return true, nil, errors.Wrapf(err, "error parsing CIDR %q", idr.serviceSubnet)
	}
	internalAPIServerVirtualIP, err := ipallocator.GetIndexedIP(svcSubnet, 1)
	if err != nil {
		return true, nil, errors.Wrapf(err, "unable to get first IP address from the given CIDR (%s)", svcSubnet.String())
	}
	return true, &v1.Service{ObjectMeta: metav1.ObjectMeta{Name: "kubernetes", Namespace: metav1.NamespaceDefault, Labels: map[string]string{"component": "apiserver", "provider": "kubernetes"}}, Spec: v1.ServiceSpec{ClusterIP: internalAPIServerVirtualIP.String(), Ports: []v1.ServicePort{{Name: "https", Port: 443, TargetPort: intstr.FromInt(6443)}}}}, nil
}
func (idr *InitDryRunGetter) handleGetNode(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if action.GetName() != idr.masterName || action.GetResource().Resource != "nodes" {
		return false, nil, nil
	}
	return true, &v1.Node{ObjectMeta: metav1.ObjectMeta{Name: idr.masterName, Labels: map[string]string{"kubernetes.io/hostname": idr.masterName}, Annotations: map[string]string{}}}, nil
}
func (idr *InitDryRunGetter) handleSystemNodesClusterRoleBinding(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if action.GetName() != constants.NodesClusterRoleBinding || action.GetResource().Resource != "clusterrolebindings" {
		return false, nil, nil
	}
	return true, nil, apierrors.NewNotFound(action.GetResource().GroupResource(), "clusterrolebinding not found")
}
func (idr *InitDryRunGetter) handleGetBootstrapToken(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(action.GetName(), "bootstrap-token-") || action.GetNamespace() != metav1.NamespaceSystem || action.GetResource().Resource != "secrets" {
		return false, nil, nil
	}
	return true, nil, apierrors.NewNotFound(action.GetResource().GroupResource(), "secret not found")
}
