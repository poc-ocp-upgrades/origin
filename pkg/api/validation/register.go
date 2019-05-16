package validation

import (
	goformat "fmt"
	_ "github.com/openshift/origin/pkg/api/install"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	appsvalidation "github.com/openshift/origin/pkg/apps/apis/apps/validation"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	authorizationvalidation "github.com/openshift/origin/pkg/authorization/apis/authorization/validation"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	buildvalidation "github.com/openshift/origin/pkg/build/apis/build/validation"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imagevalidation "github.com/openshift/origin/pkg/image/apis/image/validation"
	networkapi "github.com/openshift/origin/pkg/network/apis/network"
	sdnvalidation "github.com/openshift/origin/pkg/network/apis/network/validation"
	oauthapi "github.com/openshift/origin/pkg/oauth/apis/oauth"
	oauthvalidation "github.com/openshift/origin/pkg/oauth/apis/oauth/validation"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectvalidation "github.com/openshift/origin/pkg/project/apis/project/validation"
	quotaapi "github.com/openshift/origin/pkg/quota/apis/quota"
	quotavalidation "github.com/openshift/origin/pkg/quota/apis/quota/validation"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
	routevalidation "github.com/openshift/origin/pkg/route/apis/route/validation"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityvalidation "github.com/openshift/origin/pkg/security/apis/security/validation"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
	templatevalidation "github.com/openshift/origin/pkg/template/apis/template/validation"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
	uservalidation "github.com/openshift/origin/pkg/user/apis/user/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerAll()
}
func registerAll() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	Validator.MustRegister(&authorizationapi.SelfSubjectRulesReview{}, true, authorizationvalidation.ValidateSelfSubjectRulesReview, nil)
	Validator.MustRegister(&authorizationapi.SubjectAccessReview{}, false, authorizationvalidation.ValidateSubjectAccessReview, nil)
	Validator.MustRegister(&authorizationapi.SubjectRulesReview{}, true, authorizationvalidation.ValidateSubjectRulesReview, nil)
	Validator.MustRegister(&authorizationapi.ResourceAccessReview{}, false, authorizationvalidation.ValidateResourceAccessReview, nil)
	Validator.MustRegister(&authorizationapi.LocalSubjectAccessReview{}, true, authorizationvalidation.ValidateLocalSubjectAccessReview, nil)
	Validator.MustRegister(&authorizationapi.LocalResourceAccessReview{}, true, authorizationvalidation.ValidateLocalResourceAccessReview, nil)
	Validator.MustRegister(&authorizationapi.Role{}, true, authorizationvalidation.ValidateLocalRole, authorizationvalidation.ValidateLocalRoleUpdate)
	Validator.MustRegister(&authorizationapi.RoleBinding{}, true, authorizationvalidation.ValidateLocalRoleBinding, authorizationvalidation.ValidateLocalRoleBindingUpdate)
	Validator.MustRegister(&authorizationapi.RoleBindingRestriction{}, true, authorizationvalidation.ValidateRoleBindingRestriction, authorizationvalidation.ValidateRoleBindingRestrictionUpdate)
	Validator.MustRegister(&authorizationapi.ClusterRole{}, false, authorizationvalidation.ValidateClusterRole, authorizationvalidation.ValidateClusterRoleUpdate)
	Validator.MustRegister(&authorizationapi.ClusterRoleBinding{}, false, authorizationvalidation.ValidateClusterRoleBinding, authorizationvalidation.ValidateClusterRoleBindingUpdate)
	Validator.MustRegister(&buildapi.Build{}, true, buildvalidation.ValidateBuild, buildvalidation.ValidateBuildUpdate)
	Validator.MustRegister(&buildapi.BuildConfig{}, true, buildvalidation.ValidateBuildConfig, buildvalidation.ValidateBuildConfigUpdate)
	Validator.MustRegister(&buildapi.BuildRequest{}, true, buildvalidation.ValidateBuildRequest, nil)
	Validator.MustRegister(&buildapi.BuildLogOptions{}, true, buildvalidation.ValidateBuildLogOptions, nil)
	Validator.MustRegister(&appsapi.DeploymentConfig{}, true, appsvalidation.ValidateDeploymentConfig, appsvalidation.ValidateDeploymentConfigUpdate)
	Validator.MustRegister(&appsapi.DeploymentConfigRollback{}, true, appsvalidation.ValidateDeploymentConfigRollback, nil)
	Validator.MustRegister(&appsapi.DeploymentLogOptions{}, true, appsvalidation.ValidateDeploymentLogOptions, nil)
	Validator.MustRegister(&appsapi.DeploymentRequest{}, true, appsvalidation.ValidateDeploymentRequest, nil)
	Validator.MustRegister(&imageapi.Image{}, false, imagevalidation.ValidateImage, imagevalidation.ValidateImageUpdate)
	Validator.MustRegister(&imageapi.ImageSignature{}, false, imagevalidation.ValidateImageSignature, imagevalidation.ValidateImageSignatureUpdate)
	Validator.MustRegister(&imageapi.ImageStream{}, true, imagevalidation.ValidateImageStream, imagevalidation.ValidateImageStreamUpdate)
	Validator.MustRegister(&imageapi.ImageStreamImport{}, true, imagevalidation.ValidateImageStreamImport, nil)
	Validator.MustRegister(&imageapi.ImageStreamMapping{}, true, imagevalidation.ValidateImageStreamMapping, nil)
	Validator.MustRegister(&imageapi.ImageStreamTag{}, true, imagevalidation.ValidateImageStreamTag, imagevalidation.ValidateImageStreamTagUpdate)
	Validator.MustRegister(&oauthapi.OAuthAccessToken{}, false, oauthvalidation.ValidateAccessToken, oauthvalidation.ValidateAccessTokenUpdate)
	Validator.MustRegister(&oauthapi.OAuthAuthorizeToken{}, false, oauthvalidation.ValidateAuthorizeToken, oauthvalidation.ValidateAuthorizeTokenUpdate)
	Validator.MustRegister(&oauthapi.OAuthClient{}, false, oauthvalidation.ValidateClient, oauthvalidation.ValidateClientUpdate)
	Validator.MustRegister(&oauthapi.OAuthClientAuthorization{}, false, oauthvalidation.ValidateClientAuthorization, oauthvalidation.ValidateClientAuthorizationUpdate)
	Validator.MustRegister(&oauthapi.OAuthRedirectReference{}, true, oauthvalidation.ValidateOAuthRedirectReference, nil)
	Validator.MustRegister(&projectapi.Project{}, false, projectvalidation.ValidateProject, projectvalidation.ValidateProjectUpdate)
	Validator.MustRegister(&projectapi.ProjectRequest{}, false, projectvalidation.ValidateProjectRequest, nil)
	Validator.MustRegister(&routeapi.Route{}, true, routevalidation.ValidateRoute, routevalidation.ValidateRouteUpdate)
	Validator.MustRegister(&networkapi.ClusterNetwork{}, false, sdnvalidation.ValidateClusterNetwork, sdnvalidation.ValidateClusterNetworkUpdate)
	Validator.MustRegister(&networkapi.HostSubnet{}, false, sdnvalidation.ValidateHostSubnet, sdnvalidation.ValidateHostSubnetUpdate)
	Validator.MustRegister(&networkapi.NetNamespace{}, false, sdnvalidation.ValidateNetNamespace, sdnvalidation.ValidateNetNamespaceUpdate)
	Validator.MustRegister(&networkapi.EgressNetworkPolicy{}, true, sdnvalidation.ValidateEgressNetworkPolicy, sdnvalidation.ValidateEgressNetworkPolicyUpdate)
	Validator.MustRegister(&templateapi.Template{}, true, templatevalidation.ValidateTemplate, templatevalidation.ValidateTemplateUpdate)
	Validator.MustRegister(&templateapi.TemplateInstance{}, true, templatevalidation.ValidateTemplateInstance, templatevalidation.ValidateTemplateInstanceUpdate)
	Validator.MustRegister(&templateapi.BrokerTemplateInstance{}, false, templatevalidation.ValidateBrokerTemplateInstance, templatevalidation.ValidateBrokerTemplateInstanceUpdate)
	Validator.MustRegister(&userapi.User{}, false, uservalidation.ValidateUser, uservalidation.ValidateUserUpdate)
	Validator.MustRegister(&userapi.Identity{}, false, uservalidation.ValidateIdentity, uservalidation.ValidateIdentityUpdate)
	Validator.MustRegister(&userapi.UserIdentityMapping{}, false, uservalidation.ValidateUserIdentityMapping, uservalidation.ValidateUserIdentityMappingUpdate)
	Validator.MustRegister(&userapi.Group{}, false, uservalidation.ValidateGroup, uservalidation.ValidateGroupUpdate)
	Validator.MustRegister(&securityapi.SecurityContextConstraints{}, false, securityvalidation.ValidateSecurityContextConstraints, securityvalidation.ValidateSecurityContextConstraintsUpdate)
	Validator.MustRegister(&securityapi.PodSecurityPolicySubjectReview{}, true, securityvalidation.ValidatePodSecurityPolicySubjectReview, nil)
	Validator.MustRegister(&securityapi.PodSecurityPolicySelfSubjectReview{}, true, securityvalidation.ValidatePodSecurityPolicySelfSubjectReview, nil)
	Validator.MustRegister(&securityapi.PodSecurityPolicyReview{}, true, securityvalidation.ValidatePodSecurityPolicyReview, nil)
	Validator.MustRegister(&quotaapi.ClusterResourceQuota{}, false, quotavalidation.ValidateClusterResourceQuota, quotavalidation.ValidateClusterResourceQuotaUpdate)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
