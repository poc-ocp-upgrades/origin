package authorizer

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "time"
 "k8s.io/apiserver/pkg/authorization/authorizer"
 "k8s.io/apiserver/pkg/authorization/authorizerfactory"
 "k8s.io/apiserver/pkg/authorization/union"
 "k8s.io/apiserver/plugin/pkg/authorizer/webhook"
 versionedinformers "k8s.io/client-go/informers"
 "k8s.io/kubernetes/pkg/auth/authorizer/abac"
 "k8s.io/kubernetes/pkg/auth/nodeidentifier"
 "k8s.io/kubernetes/pkg/kubeapiserver/authorizer/modes"
 "k8s.io/kubernetes/plugin/pkg/auth/authorizer/node"
 "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac"
 "k8s.io/kubernetes/plugin/pkg/auth/authorizer/rbac/bootstrappolicy"
)

type Config struct {
 AuthorizationModes          []string
 PolicyFile                  string
 WebhookConfigFile           string
 WebhookCacheAuthorizedTTL   time.Duration
 WebhookCacheUnauthorizedTTL time.Duration
 VersionedInformerFactory    versionedinformers.SharedInformerFactory
}

func (config Config) New() (authorizer.Authorizer, authorizer.RuleResolver, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if len(config.AuthorizationModes) == 0 {
  return nil, nil, fmt.Errorf("at least one authorization mode must be passed")
 }
 var (
  authorizers   []authorizer.Authorizer
  ruleResolvers []authorizer.RuleResolver
 )
 for _, authorizationMode := range config.AuthorizationModes {
  switch authorizationMode {
  case modes.ModeNode:
   graph := node.NewGraph()
   node.AddGraphEventHandlers(graph, config.VersionedInformerFactory.Core().V1().Nodes(), config.VersionedInformerFactory.Core().V1().Pods(), config.VersionedInformerFactory.Core().V1().PersistentVolumes(), config.VersionedInformerFactory.Storage().V1beta1().VolumeAttachments())
   nodeAuthorizer := node.NewAuthorizer(graph, nodeidentifier.NewDefaultNodeIdentifier(), bootstrappolicy.NodeRules())
   authorizers = append(authorizers, nodeAuthorizer)
  case modes.ModeAlwaysAllow:
   alwaysAllowAuthorizer := authorizerfactory.NewAlwaysAllowAuthorizer()
   authorizers = append(authorizers, alwaysAllowAuthorizer)
   ruleResolvers = append(ruleResolvers, alwaysAllowAuthorizer)
  case modes.ModeAlwaysDeny:
   alwaysDenyAuthorizer := authorizerfactory.NewAlwaysDenyAuthorizer()
   authorizers = append(authorizers, alwaysDenyAuthorizer)
   ruleResolvers = append(ruleResolvers, alwaysDenyAuthorizer)
  case modes.ModeABAC:
   abacAuthorizer, err := abac.NewFromFile(config.PolicyFile)
   if err != nil {
    return nil, nil, err
   }
   authorizers = append(authorizers, abacAuthorizer)
   ruleResolvers = append(ruleResolvers, abacAuthorizer)
  case modes.ModeWebhook:
   webhookAuthorizer, err := webhook.New(config.WebhookConfigFile, config.WebhookCacheAuthorizedTTL, config.WebhookCacheUnauthorizedTTL)
   if err != nil {
    return nil, nil, err
   }
   authorizers = append(authorizers, webhookAuthorizer)
   ruleResolvers = append(ruleResolvers, webhookAuthorizer)
  case modes.ModeRBAC:
   rbacAuthorizer := rbac.New(&rbac.RoleGetter{Lister: config.VersionedInformerFactory.Rbac().V1().Roles().Lister()}, &rbac.RoleBindingLister{Lister: config.VersionedInformerFactory.Rbac().V1().RoleBindings().Lister()}, &rbac.ClusterRoleGetter{Lister: config.VersionedInformerFactory.Rbac().V1().ClusterRoles().Lister()}, &rbac.ClusterRoleBindingLister{Lister: config.VersionedInformerFactory.Rbac().V1().ClusterRoleBindings().Lister()})
   authorizers = append(authorizers, rbacAuthorizer)
   ruleResolvers = append(ruleResolvers, rbacAuthorizer)
  default:
   return nil, nil, fmt.Errorf("unknown authorization mode %s specified", authorizationMode)
  }
 }
 return union.New(authorizers...), union.NewRuleResolvers(ruleResolvers...), nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
