package fake

import (
 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 runtime "k8s.io/apimachinery/pkg/runtime"
 schema "k8s.io/apimachinery/pkg/runtime/schema"
 serializer "k8s.io/apimachinery/pkg/runtime/serializer"
 utilruntime "k8s.io/apimachinery/pkg/util/runtime"
 admissionregistrationinternalversion "k8s.io/kubernetes/pkg/apis/admissionregistration"
 appsinternalversion "k8s.io/kubernetes/pkg/apis/apps"
 auditregistrationinternalversion "k8s.io/kubernetes/pkg/apis/auditregistration"
 authenticationinternalversion "k8s.io/kubernetes/pkg/apis/authentication"
 authorizationinternalversion "k8s.io/kubernetes/pkg/apis/authorization"
 autoscalinginternalversion "k8s.io/kubernetes/pkg/apis/autoscaling"
 batchinternalversion "k8s.io/kubernetes/pkg/apis/batch"
 certificatesinternalversion "k8s.io/kubernetes/pkg/apis/certificates"
 coordinationinternalversion "k8s.io/kubernetes/pkg/apis/coordination"
 coreinternalversion "k8s.io/kubernetes/pkg/apis/core"
 eventsinternalversion "k8s.io/kubernetes/pkg/apis/events"
 extensionsinternalversion "k8s.io/kubernetes/pkg/apis/extensions"
 networkinginternalversion "k8s.io/kubernetes/pkg/apis/networking"
 policyinternalversion "k8s.io/kubernetes/pkg/apis/policy"
 rbacinternalversion "k8s.io/kubernetes/pkg/apis/rbac"
 schedulinginternalversion "k8s.io/kubernetes/pkg/apis/scheduling"
 settingsinternalversion "k8s.io/kubernetes/pkg/apis/settings"
 storageinternalversion "k8s.io/kubernetes/pkg/apis/storage"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)
var localSchemeBuilder = runtime.SchemeBuilder{admissionregistrationinternalversion.AddToScheme, coreinternalversion.AddToScheme, appsinternalversion.AddToScheme, auditregistrationinternalversion.AddToScheme, authenticationinternalversion.AddToScheme, authorizationinternalversion.AddToScheme, autoscalinginternalversion.AddToScheme, batchinternalversion.AddToScheme, certificatesinternalversion.AddToScheme, coordinationinternalversion.AddToScheme, eventsinternalversion.AddToScheme, extensionsinternalversion.AddToScheme, networkinginternalversion.AddToScheme, policyinternalversion.AddToScheme, rbacinternalversion.AddToScheme, schedulinginternalversion.AddToScheme, settingsinternalversion.AddToScheme, storageinternalversion.AddToScheme}
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
 utilruntime.Must(AddToScheme(scheme))
}
