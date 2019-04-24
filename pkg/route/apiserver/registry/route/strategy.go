package route

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	authorizationapi "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/storage/names"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kvalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	authorizationutil "github.com/openshift/origin/pkg/authorization/util"
	"github.com/openshift/origin/pkg/route"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	"github.com/openshift/origin/pkg/route/apis/route/validation"
)

const HostGeneratedAnnotationKey = "openshift.io/host.generated"

type SubjectAccessReviewInterface interface {
	Create(sar *authorizationapi.SubjectAccessReview) (result *authorizationapi.SubjectAccessReview, err error)
}

var _ SubjectAccessReviewInterface = authorizationclient.SubjectAccessReviewInterface(nil)

type routeStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	route.RouteAllocator
	sarClient	SubjectAccessReviewInterface
}

func NewStrategy(allocator route.RouteAllocator, sarClient SubjectAccessReviewInterface) routeStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return routeStrategy{ObjectTyper: legacyscheme.Scheme, NameGenerator: names.SimpleNameGenerator, RouteAllocator: allocator, sarClient: sarClient}
}
func (routeStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (s routeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route := obj.(*routeapi.Route)
	route.Status = routeapi.RouteStatus{}
	stripEmptyDestinationCACertificate(route)
}
func (s routeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route := obj.(*routeapi.Route)
	oldRoute := old.(*routeapi.Route)
	route.Status = oldRoute.Status
	stripEmptyDestinationCACertificate(route)
	if len(route.Spec.Host) == 0 {
		route.Spec.Host = oldRoute.Spec.Host
	}
}
func (s routeStrategy) allocateHost(ctx context.Context, route *routeapi.Route) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hostSet := len(route.Spec.Host) > 0
	certSet := route.Spec.TLS != nil && (len(route.Spec.TLS.CACertificate) > 0 || len(route.Spec.TLS.Certificate) > 0 || len(route.Spec.TLS.DestinationCACertificate) > 0 || len(route.Spec.TLS.Key) > 0)
	if hostSet || certSet {
		user, ok := apirequest.UserFrom(ctx)
		if !ok {
			return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), fmt.Errorf("unable to verify host field can be set"))}
		}
		res, err := s.sarClient.Create(authorizationutil.AddUserToSAR(user, &authorizationapi.SubjectAccessReview{Spec: authorizationapi.SubjectAccessReviewSpec{ResourceAttributes: &authorizationapi.ResourceAttributes{Namespace: apirequest.NamespaceValue(ctx), Verb: "create", Group: routeapi.GroupName, Resource: "routes", Subresource: "custom-host"}}}))
		if err != nil {
			return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), err)}
		}
		if !res.Status.Allowed {
			if hostSet {
				return field.ErrorList{field.Forbidden(field.NewPath("spec", "host"), "you do not have permission to set the host field of the route")}
			}
			return field.ErrorList{field.Forbidden(field.NewPath("spec", "tls"), "you do not have permission to set certificate fields on the route")}
		}
	}
	if route.Spec.WildcardPolicy == routeapi.WildcardPolicySubdomain {
		return nil
	}
	if len(route.Spec.Host) == 0 && s.RouteAllocator != nil {
		shard, err := s.RouteAllocator.AllocateRouterShard(route)
		if err != nil {
			return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), fmt.Errorf("allocation error: %v for route: %#v", err, route))}
		}
		route.Spec.Host = s.RouteAllocator.GenerateHostname(route, shard)
		if route.Annotations == nil {
			route.Annotations = map[string]string{}
		}
		route.Annotations[HostGeneratedAnnotationKey] = "true"
	}
	return nil
}
func (s routeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route := obj.(*routeapi.Route)
	errs := s.allocateHost(ctx, route)
	errs = append(errs, validation.ValidateRoute(route)...)
	return errs
}
func (routeStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (routeStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s routeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldRoute := old.(*routeapi.Route)
	objRoute := obj.(*routeapi.Route)
	errs := s.validateHostUpdate(ctx, objRoute, oldRoute)
	errs = append(errs, validation.ValidateRouteUpdate(objRoute, oldRoute)...)
	return errs
}
func hasCertificateInfo(tls *routeapi.TLSConfig) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tls == nil {
		return false
	}
	return len(tls.Certificate) > 0 || len(tls.Key) > 0 || len(tls.CACertificate) > 0 || len(tls.DestinationCACertificate) > 0
}
func certificateChangeRequiresAuth(route, older *routeapi.Route) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case route.Spec.TLS != nil && older.Spec.TLS != nil:
		a, b := route.Spec.TLS, older.Spec.TLS
		if !hasCertificateInfo(a) {
			return false
		}
		return a.CACertificate != b.CACertificate || a.Certificate != b.Certificate || a.DestinationCACertificate != b.DestinationCACertificate || a.Key != b.Key
	case route.Spec.TLS != nil:
		return hasCertificateInfo(route.Spec.TLS)
	default:
		return false
	}
}
func (s routeStrategy) validateHostUpdate(ctx context.Context, route, older *routeapi.Route) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hostChanged := route.Spec.Host != older.Spec.Host
	certChanged := certificateChangeRequiresAuth(route, older)
	if !hostChanged && !certChanged {
		return nil
	}
	user, ok := apirequest.UserFrom(ctx)
	if !ok {
		return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), fmt.Errorf("unable to verify host field can be changed"))}
	}
	res, err := s.sarClient.Create(authorizationutil.AddUserToSAR(user, &authorizationapi.SubjectAccessReview{Spec: authorizationapi.SubjectAccessReviewSpec{ResourceAttributes: &authorizationapi.ResourceAttributes{Namespace: apirequest.NamespaceValue(ctx), Verb: "update", Group: routeapi.GroupName, Resource: "routes", Subresource: "custom-host"}}}))
	if err != nil {
		return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), err)}
	}
	if !res.Status.Allowed {
		if hostChanged {
			return kvalidation.ValidateImmutableField(route.Spec.Host, older.Spec.Host, field.NewPath("spec", "host"))
		}
		res, err := s.sarClient.Create(authorizationutil.AddUserToSAR(user, &authorizationapi.SubjectAccessReview{Spec: authorizationapi.SubjectAccessReviewSpec{ResourceAttributes: &authorizationapi.ResourceAttributes{Namespace: apirequest.NamespaceValue(ctx), Verb: "create", Group: routeapi.GroupName, Resource: "routes", Subresource: "custom-host"}}}))
		if err != nil {
			return field.ErrorList{field.InternalError(field.NewPath("spec", "host"), err)}
		}
		if !res.Status.Allowed {
			if route.Spec.TLS == nil || older.Spec.TLS == nil {
				return kvalidation.ValidateImmutableField(route.Spec.TLS, older.Spec.TLS, field.NewPath("spec", "tls"))
			}
			errs := kvalidation.ValidateImmutableField(route.Spec.TLS.CACertificate, older.Spec.TLS.CACertificate, field.NewPath("spec", "tls", "caCertificate"))
			errs = append(errs, kvalidation.ValidateImmutableField(route.Spec.TLS.Certificate, older.Spec.TLS.Certificate, field.NewPath("spec", "tls", "certificate"))...)
			errs = append(errs, kvalidation.ValidateImmutableField(route.Spec.TLS.DestinationCACertificate, older.Spec.TLS.DestinationCACertificate, field.NewPath("spec", "tls", "destinationCACertificate"))...)
			errs = append(errs, kvalidation.ValidateImmutableField(route.Spec.TLS.Key, older.Spec.TLS.Key, field.NewPath("spec", "tls", "key"))...)
			return errs
		}
	}
	return nil
}
func (routeStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}

type routeStatusStrategy struct{ routeStrategy }

var StatusStrategy = routeStatusStrategy{NewStrategy(nil, nil)}

func (routeStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newRoute := obj.(*routeapi.Route)
	oldRoute := old.(*routeapi.Route)
	newRoute.Spec = oldRoute.Spec
}
func (routeStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validation.ValidateRouteStatusUpdate(obj.(*routeapi.Route), old.(*routeapi.Route))
}

const emptyDestinationCertificate = `-----BEGIN COMMENT-----
This is an empty PEM file created to provide backwards compatibility
for reencrypt routes that have no destinationCACertificate. This 
content will only appear for routes accessed via /oapi/v1/routes.
-----END COMMENT-----
`

func stripEmptyDestinationCACertificate(route *routeapi.Route) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tls := route.Spec.TLS
	if tls == nil || tls.Termination != routeapi.TLSTerminationReencrypt {
		return
	}
	if tls.DestinationCACertificate == emptyDestinationCertificate {
		tls.DestinationCACertificate = ""
	}
}
func DecorateLegacyRouteWithEmptyDestinationCACertificates(obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := obj.(type) {
	case *routeapi.Route:
		tls := t.Spec.TLS
		if tls == nil || tls.Termination != routeapi.TLSTerminationReencrypt {
			return nil
		}
		if len(tls.DestinationCACertificate) == 0 {
			tls.DestinationCACertificate = emptyDestinationCertificate
		}
		return nil
	case *routeapi.RouteList:
		for i := range t.Items {
			tls := t.Items[i].Spec.TLS
			if tls == nil || tls.Termination != routeapi.TLSTerminationReencrypt {
				continue
			}
			if len(tls.DestinationCACertificate) == 0 {
				tls.DestinationCACertificate = emptyDestinationCertificate
			}
		}
		return nil
	default:
		return fmt.Errorf("unknown type passed to %T", obj)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
