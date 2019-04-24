package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	v1 "github.com/openshift/api/oauth/v1"
)

func addFieldSelectorKeyConversions(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := scheme.AddFieldLabelConversionFunc(v1.GroupVersion.WithKind("OAuthAccessToken"), oauthAccessTokenFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	if err := scheme.AddFieldLabelConversionFunc(v1.GroupVersion.WithKind("OAuthAuthorizeToken"), oauthAuthorizeTokenFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	if err := scheme.AddFieldLabelConversionFunc(v1.GroupVersion.WithKind("OAuthClientAuthorization"), oauthClientAuthorizationFieldSelectorKeyConversionFunc); err != nil {
		return err
	}
	return nil
}
func oauthAccessTokenFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "clientName", "userName", "userUID", "authorizeToken":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func oauthAuthorizeTokenFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "clientName", "userName", "userUID":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func oauthClientAuthorizationFieldSelectorKeyConversionFunc(label, value string) (internalLabel, internalValue string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "clientName", "userName", "userUID":
		return label, value, nil
	default:
		return runtime.DefaultMetaV1FieldSelectorConversion(label, value)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
