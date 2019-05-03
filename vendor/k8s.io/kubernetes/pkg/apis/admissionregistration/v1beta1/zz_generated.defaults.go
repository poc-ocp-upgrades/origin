package v1beta1

import (
 v1beta1 "k8s.io/api/admissionregistration/v1beta1"
 runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 scheme.AddTypeDefaultingFunc(&v1beta1.MutatingWebhookConfiguration{}, func(obj interface{}) {
  SetObjectDefaults_MutatingWebhookConfiguration(obj.(*v1beta1.MutatingWebhookConfiguration))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.MutatingWebhookConfigurationList{}, func(obj interface{}) {
  SetObjectDefaults_MutatingWebhookConfigurationList(obj.(*v1beta1.MutatingWebhookConfigurationList))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.ValidatingWebhookConfiguration{}, func(obj interface{}) {
  SetObjectDefaults_ValidatingWebhookConfiguration(obj.(*v1beta1.ValidatingWebhookConfiguration))
 })
 scheme.AddTypeDefaultingFunc(&v1beta1.ValidatingWebhookConfigurationList{}, func(obj interface{}) {
  SetObjectDefaults_ValidatingWebhookConfigurationList(obj.(*v1beta1.ValidatingWebhookConfigurationList))
 })
 return nil
}
func SetObjectDefaults_MutatingWebhookConfiguration(in *v1beta1.MutatingWebhookConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Webhooks {
  a := &in.Webhooks[i]
  SetDefaults_Webhook(a)
 }
}
func SetObjectDefaults_MutatingWebhookConfigurationList(in *v1beta1.MutatingWebhookConfigurationList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_MutatingWebhookConfiguration(a)
 }
}
func SetObjectDefaults_ValidatingWebhookConfiguration(in *v1beta1.ValidatingWebhookConfiguration) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Webhooks {
  a := &in.Webhooks[i]
  SetDefaults_Webhook(a)
 }
}
func SetObjectDefaults_ValidatingWebhookConfigurationList(in *v1beta1.ValidatingWebhookConfigurationList) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for i := range in.Items {
  a := &in.Items[i]
  SetObjectDefaults_ValidatingWebhookConfiguration(a)
 }
}
