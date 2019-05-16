package restrictedendpoints

import (
	"fmt"
	goformat "fmt"
	configlatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/network/admission/apis/restrictedendpoints"
	"io"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/klog"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"net"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

const RestrictedEndpointsPluginName = "network.openshift.io/RestrictedEndpointsAdmission"

func RegisterRestrictedEndpoints(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(RestrictedEndpointsPluginName, func(config io.Reader) (admission.Interface, error) {
		pluginConfig, err := readConfig(config)
		if err != nil {
			return nil, err
		}
		if pluginConfig == nil {
			klog.Infof("Admission plugin %q is not configured so it will be disabled.", RestrictedEndpointsPluginName)
			return nil, nil
		}
		restrictedNetworks, err := ParseSimpleCIDRRules(pluginConfig.RestrictedCIDRs)
		if err != nil {
			return nil, err
		}
		return NewRestrictedEndpointsAdmission(restrictedNetworks), nil
	})
}
func readConfig(reader io.Reader) (*restrictedendpoints.RestrictedEndpointsAdmissionConfig, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if reader == nil || reflect.ValueOf(reader).IsNil() {
		return nil, nil
	}
	obj, err := configlatest.ReadYAML(reader)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, nil
	}
	config, ok := obj.(*restrictedendpoints.RestrictedEndpointsAdmissionConfig)
	if !ok {
		return nil, fmt.Errorf("unexpected config object: %#v", obj)
	}
	return config, nil
}

type restrictedEndpointsAdmission struct {
	*admission.Handler
	authorizer         authorizer.Authorizer
	restrictedNetworks []*net.IPNet
}

var _ = initializer.WantsAuthorizer(&restrictedEndpointsAdmission{})
var _ = admission.ValidationInterface(&restrictedEndpointsAdmission{})

func ParseSimpleCIDRRules(rules []string) (networks []*net.IPNet, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, s := range rules {
		_, cidr, err := net.ParseCIDR(s)
		if err != nil {
			return nil, err
		}
		networks = append(networks, cidr)
	}
	return networks, nil
}
func NewRestrictedEndpointsAdmission(restrictedNetworks []*net.IPNet) *restrictedEndpointsAdmission {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &restrictedEndpointsAdmission{Handler: admission.NewHandler(admission.Create, admission.Update), restrictedNetworks: restrictedNetworks}
}
func (r *restrictedEndpointsAdmission) SetAuthorizer(a authorizer.Authorizer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r.authorizer = a
}
func (r *restrictedEndpointsAdmission) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if r.authorizer == nil {
		return fmt.Errorf("missing authorizer")
	}
	return nil
}
func (r *restrictedEndpointsAdmission) findRestrictedIP(ep *kapi.Endpoints) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, subset := range ep.Subsets {
		for _, addr := range subset.Addresses {
			ip := net.ParseIP(addr.IP)
			if ip == nil {
				continue
			}
			for _, net := range r.restrictedNetworks {
				if net.Contains(ip) {
					return addr.IP
				}
			}
		}
	}
	return ""
}
func (r *restrictedEndpointsAdmission) checkAccess(attr admission.Attributes) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authzAttr := authorizer.AttributesRecord{User: attr.GetUserInfo(), Verb: "create", Namespace: attr.GetNamespace(), Resource: "endpoints", Subresource: "restricted", APIGroup: kapi.GroupName, Name: attr.GetName(), ResourceRequest: true}
	authorized, _, err := r.authorizer.Authorize(authzAttr)
	return authorized == authorizer.DecisionAllow, err
}
func (r *restrictedEndpointsAdmission) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != kapi.Resource("endpoints") {
		return nil
	}
	ep, ok := a.GetObject().(*kapi.Endpoints)
	if !ok {
		return nil
	}
	old, ok := a.GetOldObject().(*kapi.Endpoints)
	if ok && reflect.DeepEqual(ep.Subsets, old.Subsets) {
		return nil
	}
	restrictedIP := r.findRestrictedIP(ep)
	if restrictedIP == "" {
		return nil
	}
	allow, err := r.checkAccess(a)
	if err != nil {
		return err
	}
	if !allow {
		return admission.NewForbidden(a, fmt.Errorf("endpoint address %s is not allowed", restrictedIP))
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
