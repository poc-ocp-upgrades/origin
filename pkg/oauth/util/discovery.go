package util

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/RangelReale/osin"
	osinv1 "github.com/openshift/api/osin/v1"
	"github.com/openshift/origin/pkg/authorization/authorizer/scope"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/oauth/apis/oauth/validation"
	"github.com/openshift/origin/pkg/oauth/urls"
	"io/ioutil"
	"k8s.io/klog"
	"net/url"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type OauthAuthorizationServerMetadata struct {
	Issuer                        string                    `json:"issuer"`
	AuthorizationEndpoint         string                    `json:"authorization_endpoint"`
	TokenEndpoint                 string                    `json:"token_endpoint"`
	ScopesSupported               []string                  `json:"scopes_supported"`
	ResponseTypesSupported        osin.AllowedAuthorizeType `json:"response_types_supported"`
	GrantTypesSupported           osin.AllowedAccessType    `json:"grant_types_supported"`
	CodeChallengeMethodsSupported []string                  `json:"code_challenge_methods_supported"`
}

func getOauthMetadata(masterPublicURL string) OauthAuthorizationServerMetadata {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return OauthAuthorizationServerMetadata{Issuer: masterPublicURL, AuthorizationEndpoint: urls.OpenShiftOAuthAuthorizeURL(masterPublicURL), TokenEndpoint: urls.OpenShiftOAuthTokenURL(masterPublicURL), ScopesSupported: scope.DefaultSupportedScopes(), ResponseTypesSupported: osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}, GrantTypesSupported: osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.AccessRequestType("implicit")}, CodeChallengeMethodsSupported: validation.CodeChallengeMethodsSupported}
}
func validateURL(urlString string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	urlObj, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("%q is an invalid URL: %v", urlString, err)
	}
	if len(urlObj.Scheme) == 0 {
		return fmt.Errorf("must contain a valid scheme")
	}
	if len(urlObj.Host) == 0 {
		return fmt.Errorf("must contain a valid host")
	}
	return nil
}
func LoadOAuthMetadataFile(metadataFile string) ([]byte, *OauthAuthorizationServerMetadata, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read External OAuth Metadata file: %v", err)
	}
	oauthMetadata := &OauthAuthorizationServerMetadata{}
	if err := json.Unmarshal(data, oauthMetadata); err != nil {
		return nil, nil, fmt.Errorf("unable to decode External OAuth Metadata file: %v", err)
	}
	if err := validateURL(oauthMetadata.Issuer); err != nil {
		return nil, nil, fmt.Errorf("error validating External OAuth Metadata Issuer field: %v", err)
	}
	if err := validateURL(oauthMetadata.AuthorizationEndpoint); err != nil {
		return nil, nil, fmt.Errorf("error validating External OAuth Metadata AuthorizationEndpoint field: %v", err)
	}
	if err := validateURL(oauthMetadata.TokenEndpoint); err != nil {
		return nil, nil, fmt.Errorf("error validating External OAuth Metadata TokenEndpoint field: %v", err)
	}
	return data, oauthMetadata, nil
}
func PrepOauthMetadata(oauthConfig *osinv1.OAuthConfig, oauthMetadataFile string) ([]byte, *OauthAuthorizationServerMetadata, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(oauthMetadataFile) > 0 {
		return LoadOAuthMetadataFile(oauthMetadataFile)
	}
	if oauthConfig != nil && len(oauthConfig.MasterPublicURL) != 0 {
		metadataStruct := getOauthMetadata(oauthConfig.MasterPublicURL)
		metadata, err := json.MarshalIndent(metadataStruct, "", "  ")
		if err != nil {
			klog.Errorf("Unable to initialize OAuth authorization server metadata route: %v", err)
			return nil, nil, err
		}
		return metadata, &metadataStruct, nil
	}
	return nil, nil, nil
}
func DeprecatedPrepOauthMetadata(oauthConfig *configapi.OAuthConfig, oauthMetadataFile string) ([]byte, *OauthAuthorizationServerMetadata, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if oauthConfig != nil {
		metadataStruct := getOauthMetadata(oauthConfig.MasterPublicURL)
		metadata, err := json.MarshalIndent(metadataStruct, "", "  ")
		if err != nil {
			klog.Errorf("Unable to initialize OAuth authorization server metadata route: %v", err)
			return nil, nil, err
		}
		return metadata, &metadataStruct, nil
	}
	if len(oauthMetadataFile) > 0 {
		return LoadOAuthMetadataFile(oauthMetadataFile)
	}
	return nil, nil, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
