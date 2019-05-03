package scheme

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 serializer "k8s.io/apimachinery/pkg/runtime/serializer"
 admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration/install"
 apps "k8s.io/kubernetes/pkg/apis/apps/install"
 auditregistration "k8s.io/kubernetes/pkg/apis/auditregistration/install"
 authentication "k8s.io/kubernetes/pkg/apis/authentication/install"
 authorization "k8s.io/kubernetes/pkg/apis/authorization/install"
 autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling/install"
 batch "k8s.io/kubernetes/pkg/apis/batch/install"
 certificates "k8s.io/kubernetes/pkg/apis/certificates/install"
 coordination "k8s.io/kubernetes/pkg/apis/coordination/install"
 core "k8s.io/kubernetes/pkg/apis/core/install"
 events "k8s.io/kubernetes/pkg/apis/events/install"
 extensions "k8s.io/kubernetes/pkg/apis/extensions/install"
 networking "k8s.io/kubernetes/pkg/apis/networking/install"
 policy "k8s.io/kubernetes/pkg/apis/policy/install"
 rbac "k8s.io/kubernetes/pkg/apis/rbac/install"
 scheduling "k8s.io/kubernetes/pkg/apis/scheduling/install"
 settings "k8s.io/kubernetes/pkg/apis/settings/install"
 storage "k8s.io/kubernetes/pkg/apis/storage/install"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
 Install(Scheme)
}
func Install(scheme *runtime.Scheme) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 admissionregistration.Install(scheme)
 core.Install(scheme)
 apps.Install(scheme)
 auditregistration.Install(scheme)
 authentication.Install(scheme)
 authorization.Install(scheme)
 autoscaling.Install(scheme)
 batch.Install(scheme)
 certificates.Install(scheme)
 coordination.Install(scheme)
 events.Install(scheme)
 extensions.Install(scheme)
 networking.Install(scheme)
 policy.Install(scheme)
 rbac.Install(scheme)
 scheduling.Install(scheme)
 settings.Install(scheme)
 storage.Install(scheme)
}
