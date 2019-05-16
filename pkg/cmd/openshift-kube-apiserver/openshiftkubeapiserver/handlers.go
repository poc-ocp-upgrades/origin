package openshiftkubeapiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/emicklei/go-restful"
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authenticationv1 "k8s.io/api/authentication/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	"net/http"
	"regexp"
)

type userAgentFilter struct {
	regex   *regexp.Regexp
	message string
	verbs   sets.String
}

func newUserAgentFilter(config kubecontrolplanev1.UserAgentMatchRule) (userAgentFilter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	regex, err := regexp.Compile(config.Regex)
	if err != nil {
		return userAgentFilter{}, err
	}
	userAgentFilter := userAgentFilter{regex: regex, verbs: sets.NewString(config.HTTPVerbs...)}
	return userAgentFilter, nil
}
func (f *userAgentFilter) matches(verb, userAgent string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(f.verbs) > 0 && !f.verbs.Has(verb) {
		return false
	}
	return f.regex.MatchString(userAgent)
}
func versionSkewFilter(handler http.Handler, userAgentMatchingConfig kubecontrolplanev1.UserAgentMatchingConfig) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filterConfig := userAgentMatchingConfig
	if len(filterConfig.RequiredClients) == 0 && len(filterConfig.DeniedClients) == 0 {
		return handler
	}
	defaultMessage := filterConfig.DefaultRejectionMessage
	if len(defaultMessage) == 0 {
		defaultMessage = "the cluster administrator has disabled access for this client, please upgrade or consult your administrator"
	}
	allowedFilters := []userAgentFilter{}
	deniedFilters := []userAgentFilter{}
	for _, config := range filterConfig.RequiredClients {
		userAgentFilter, err := newUserAgentFilter(config)
		if err != nil {
			klog.Errorf("Failure to compile User-Agent regex %v: %v", config.Regex, err)
			continue
		}
		allowedFilters = append(allowedFilters, userAgentFilter)
	}
	for _, config := range filterConfig.DeniedClients {
		userAgentFilter, err := newUserAgentFilter(config.UserAgentMatchRule)
		if err != nil {
			klog.Errorf("Failure to compile User-Agent regex %v: %v", config.Regex, err)
			continue
		}
		userAgentFilter.message = config.RejectionMessage
		if len(userAgentFilter.message) == 0 {
			userAgentFilter.message = defaultMessage
		}
		deniedFilters = append(deniedFilters, userAgentFilter)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if requestInfo, ok := apirequest.RequestInfoFrom(ctx); ok && requestInfo != nil && !requestInfo.IsResourceRequest {
			handler.ServeHTTP(w, req)
			return
		}
		userAgent := req.Header.Get("User-Agent")
		if len(allowedFilters) > 0 {
			foundMatch := false
			for _, filter := range allowedFilters {
				if filter.matches(req.Method, userAgent) {
					foundMatch = true
					break
				}
			}
			if !foundMatch {
				forbidden(defaultMessage, nil, w, req)
				return
			}
		}
		for _, filter := range deniedFilters {
			if filter.matches(req.Method, userAgent) {
				forbidden(filter.message, nil, w, req)
				return
			}
		}
		handler.ServeHTTP(w, req)
	})
}

const legacyImpersonateUserScopeHeader = "Impersonate-User-Scope"

func translateLegacyScopeImpersonation(handler http.Handler) http.Handler {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, scope := range req.Header[legacyImpersonateUserScopeHeader] {
			req.Header[authenticationv1.ImpersonateUserExtraHeaderPrefix+authorizationapi.ScopesKey] = append(req.Header[authenticationv1.ImpersonateUserExtraHeaderPrefix+authorizationapi.ScopesKey], scope)
		}
		handler.ServeHTTP(w, req)
	})
}
func forbidden(reason string, attributes kauthorizer.Attributes, w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resource := ""
	group := ""
	name := ""
	if attributes != nil {
		group = attributes.GetAPIGroup()
		resource = attributes.GetResource()
		name = attributes.GetName()
	}
	forbiddenError := kapierrors.NewForbidden(schema.GroupResource{Group: group, Resource: resource}, name, errors.New(""))
	forbiddenError.ErrStatus.Message = reason
	formatted := &bytes.Buffer{}
	output, err := runtime.Encode(legacyscheme.Codecs.LegacyCodec(coreapi.SchemeGroupVersion), &forbiddenError.ErrStatus)
	if err != nil {
		fmt.Fprintf(formatted, "%s", forbiddenError.Error())
	} else {
		json.Indent(formatted, output, "", "  ")
	}
	w.Header().Set("Content-Type", restful.MIME_JSON)
	w.WriteHeader(http.StatusForbidden)
	w.Write(formatted.Bytes())
}
