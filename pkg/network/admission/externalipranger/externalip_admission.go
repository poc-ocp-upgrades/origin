package externalipranger

import (
	"fmt"
	goformat "fmt"
	configlatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/network/admission/apis/externalipranger"
	"io"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
	admission "k8s.io/apiserver/pkg/admission"
	"k8s.io/klog"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"net"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const ExternalIPPluginName = "network.openshift.io/ExternalIPRanger"

func RegisterExternalIP(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register("network.openshift.io/ExternalIPRanger", func(config io.Reader) (admission.Interface, error) {
		pluginConfig, err := readConfig(config)
		if err != nil {
			return nil, err
		}
		if pluginConfig == nil {
			klog.Infof("Admission plugin %q is not configured so it will be disabled.", ExternalIPPluginName)
			return nil, nil
		}
		reject, admit, err := ParseRejectAdmitCIDRRules(pluginConfig.ExternalIPNetworkCIDRs)
		if err != nil {
			return nil, err
		}
		return NewExternalIPRanger(reject, admit, pluginConfig.AllowIngressIP), nil
	})
}
func readConfig(reader io.Reader) (*externalipranger.ExternalIPRangerAdmissionConfig, error) {
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
	config, ok := obj.(*externalipranger.ExternalIPRangerAdmissionConfig)
	if !ok {
		return nil, fmt.Errorf("unexpected config object: %#v", obj)
	}
	return config, nil
}

type externalIPRanger struct {
	*admission.Handler
	reject         []*net.IPNet
	admit          []*net.IPNet
	allowIngressIP bool
}

var _ admission.Interface = &externalIPRanger{}
var _ admission.ValidationInterface = &externalIPRanger{}

func ParseRejectAdmitCIDRRules(rules []string) (reject, admit []*net.IPNet, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, s := range rules {
		negate := false
		if strings.HasPrefix(s, "!") {
			negate = true
			s = s[1:]
		}
		_, cidr, err := net.ParseCIDR(s)
		if err != nil {
			return nil, nil, err
		}
		if negate {
			reject = append(reject, cidr)
		} else {
			admit = append(admit, cidr)
		}
	}
	return reject, admit, nil
}
func NewExternalIPRanger(reject, admit []*net.IPNet, allowIngressIP bool) *externalIPRanger {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &externalIPRanger{Handler: admission.NewHandler(admission.Create, admission.Update), reject: reject, admit: admit, allowIngressIP: allowIngressIP}
}

type NetworkSlice []*net.IPNet

func (s NetworkSlice) Contains(ip net.IP) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, cidr := range s {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
func (r *externalIPRanger) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != kapi.Resource("services") {
		return nil
	}
	svc, ok := a.GetObject().(*kapi.Service)
	if !ok {
		return nil
	}
	ingressIP := ""
	retrieveIngressIP := a.GetOperation() == admission.Update && r.allowIngressIP && svc.Spec.Type == kapi.ServiceTypeLoadBalancer
	if retrieveIngressIP {
		old, ok := a.GetOldObject().(*kapi.Service)
		ipPresent := ok && old != nil && len(old.Status.LoadBalancer.Ingress) > 0
		if ipPresent {
			ingressIP = old.Status.LoadBalancer.Ingress[0].IP
		}
	}
	var errs field.ErrorList
	switch {
	case len(svc.Spec.ExternalIPs) > 0 && len(r.admit) == 0:
		onlyIngressIP := len(svc.Spec.ExternalIPs) == 1 && svc.Spec.ExternalIPs[0] == ingressIP
		if !onlyIngressIP {
			errs = append(errs, field.Forbidden(field.NewPath("spec", "externalIPs"), "externalIPs have been disabled"))
		}
	case len(svc.Spec.ExternalIPs) > 0 && len(r.admit) > 0:
		for i, s := range svc.Spec.ExternalIPs {
			ip := net.ParseIP(s)
			if ip == nil {
				errs = append(errs, field.Forbidden(field.NewPath("spec", "externalIPs").Index(i), "externalIPs must be a valid address"))
				continue
			}
			notIngressIP := s != ingressIP
			if (NetworkSlice(r.reject).Contains(ip) || !NetworkSlice(r.admit).Contains(ip)) && notIngressIP {
				errs = append(errs, field.Forbidden(field.NewPath("spec", "externalIPs").Index(i), "externalIP is not allowed"))
				continue
			}
		}
	}
	if len(errs) > 0 {
		return apierrs.NewInvalid(a.GetKind().GroupKind(), a.GetName(), errs)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
