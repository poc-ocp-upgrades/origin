package servicebroker

import (
	templateapiv1 "github.com/openshift/api/template/v1"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
	"github.com/openshift/origin/pkg/templateservicebroker/util"
	authorizationv1 "k8s.io/api/authorization/v1"
	kapiv1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog"
	"net/http"
	"reflect"
	"time"
)

func (b *Broker) ensureSecret(u user.Info, namespace string, brokerTemplateInstance *templateapiv1.BrokerTemplateInstance, instanceID string, preq *api.ProvisionRequest, template *templateapiv1.Template, didWork *bool) (*kapiv1.Secret, *api.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Template service broker: ensureSecret")
	blockOwnerDeletion := true
	secret := &kapiv1.Secret{ObjectMeta: metav1.ObjectMeta{Name: instanceID, OwnerReferences: []metav1.OwnerReference{{APIVersion: templateapiv1.SchemeGroupVersion.String(), Kind: "BrokerTemplateInstance", Name: brokerTemplateInstance.Name, UID: brokerTemplateInstance.UID, BlockOwnerDeletion: &blockOwnerDeletion}}}, Data: map[string][]byte{}}
	for k, v := range preq.Parameters {
		for _, param := range template.Parameters {
			if param.Name == k && ((len(v) == 0 && param.Generate == "") || len(v) != 0) {
				secret.Data[k] = []byte(v)
			}
		}
	}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "create", Group: kapiv1.GroupName, Resource: "secrets", Name: secret.Name}); err != nil {
		return nil, api.Forbidden(err)
	}
	createdSec, err := b.kc.CoreV1().Secrets(namespace).Create(secret)
	if err == nil {
		*didWork = true
		return createdSec, nil
	}
	if kerrors.IsAlreadyExists(err) {
		if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "get", Group: kapiv1.GroupName, Resource: "secrets", Name: secret.Name}); err != nil {
			return nil, api.Forbidden(err)
		}
		existingSec, err := b.kc.CoreV1().Secrets(namespace).Get(secret.Name, metav1.GetOptions{})
		if err == nil && reflect.DeepEqual(secret.Data, existingSec.Data) {
			return existingSec, nil
		}
		return nil, api.NewResponse(http.StatusConflict, api.ProvisionResponse{}, nil)
	}
	if kerrors.IsForbidden(err) {
		return nil, api.Forbidden(err)
	}
	return nil, api.InternalServerError(err)
}
func (b *Broker) ensureTemplateInstance(u user.Info, namespace string, brokerTemplateInstance *templateapiv1.BrokerTemplateInstance, instanceID string, template *templateapiv1.Template, secret *kapiv1.Secret, didWork *bool) (*templateapiv1.TemplateInstance, *api.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Template service broker: ensureTemplateInstance")
	extra := map[string]templateapiv1.ExtraValue{}
	for k, v := range u.GetExtra() {
		extra[k] = templateapiv1.ExtraValue(v)
	}
	blockOwnerDeletion := true
	templateInstance := &templateapiv1.TemplateInstance{ObjectMeta: metav1.ObjectMeta{Name: instanceID, Annotations: map[string]string{api.OpenServiceBrokerInstanceExternalID: instanceID}, OwnerReferences: []metav1.OwnerReference{{APIVersion: templateapiv1.SchemeGroupVersion.String(), Kind: "BrokerTemplateInstance", Name: brokerTemplateInstance.Name, UID: brokerTemplateInstance.UID, BlockOwnerDeletion: &blockOwnerDeletion}}}, Spec: templateapiv1.TemplateInstanceSpec{Template: *template, Secret: &kapiv1.LocalObjectReference{Name: secret.Name}, Requester: &templateapiv1.TemplateInstanceRequester{Username: u.GetName(), UID: u.GetUID(), Groups: u.GetGroups(), Extra: extra}}}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "create", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: instanceID}); err != nil {
		return nil, api.Forbidden(err)
	}
	createdTemplateInstance, err := b.templateclient.TemplateInstances(namespace).Create(templateInstance)
	if err == nil {
		*didWork = true
		return createdTemplateInstance, nil
	}
	if kerrors.IsAlreadyExists(err) {
		if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "get", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: templateInstance.Name}); err != nil {
			return nil, api.Forbidden(err)
		}
		existingTemplateInstance, err := b.templateclient.TemplateInstances(namespace).Get(templateInstance.Name, metav1.GetOptions{})
		if err == nil && reflect.DeepEqual(templateInstance.Spec, existingTemplateInstance.Spec) {
			return existingTemplateInstance, nil
		}
		return nil, api.NewResponse(http.StatusConflict, api.ProvisionResponse{}, nil)
	}
	if kerrors.IsForbidden(err) {
		return nil, api.Forbidden(err)
	}
	return nil, api.InternalServerError(err)
}
func (b *Broker) ensureBrokerTemplateInstanceUIDs(u user.Info, namespace string, brokerTemplateInstance *templateapiv1.BrokerTemplateInstance, secret *kapiv1.Secret, templateInstance *templateapiv1.TemplateInstance, didWork *bool) (*templateapiv1.BrokerTemplateInstance, *api.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Template service broker: ensureBrokerTemplateInstanceUIDs")
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "update", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: brokerTemplateInstance.Spec.TemplateInstance.Name}); err != nil {
		return nil, api.Forbidden(err)
	}
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		brokerTemplateInstance.Spec.Secret.UID = secret.UID
		brokerTemplateInstance.Spec.TemplateInstance.UID = templateInstance.UID
		newBrokerTemplateInstance, err := b.templateclient.BrokerTemplateInstances().Update(brokerTemplateInstance)
		switch {
		case err == nil:
			brokerTemplateInstance = newBrokerTemplateInstance
		case kerrors.IsConflict(err):
			var getErr error
			brokerTemplateInstance, getErr = b.templateclient.BrokerTemplateInstances().Get(brokerTemplateInstance.Name, metav1.GetOptions{})
			if getErr != nil {
				err = getErr
			}
		}
		return err
	})
	switch {
	case err == nil:
		*didWork = true
		return brokerTemplateInstance, nil
	case kerrors.IsConflict(err):
		return nil, api.NewResponse(http.StatusUnprocessableEntity, &api.ConcurrencyError, nil)
	}
	return nil, api.InternalServerError(err)
}
func (b *Broker) ensureBrokerTemplateInstance(u user.Info, namespace, instanceID string, didWork *bool) (*templateapiv1.BrokerTemplateInstance, *api.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Template service broker: ensureBrokerTemplateInstance")
	brokerTemplateInstance := &templateapiv1.BrokerTemplateInstance{ObjectMeta: metav1.ObjectMeta{Name: instanceID}, Spec: templateapiv1.BrokerTemplateInstanceSpec{TemplateInstance: kapiv1.ObjectReference{Kind: "TemplateInstance", Namespace: namespace, Name: instanceID}, Secret: kapiv1.ObjectReference{Kind: "Secret", Namespace: namespace, Name: instanceID}}}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "create", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: instanceID}); err != nil {
		return nil, api.Forbidden(err)
	}
	newBrokerTemplateInstance, err := b.templateclient.BrokerTemplateInstances().Create(brokerTemplateInstance)
	if err == nil {
		*didWork = true
		return newBrokerTemplateInstance, nil
	}
	if kerrors.IsAlreadyExists(err) {
		if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "get", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: instanceID}); err != nil {
			return nil, api.Forbidden(err)
		}
		existingBrokerTemplateInstance, err := b.templateclient.BrokerTemplateInstances().Get(brokerTemplateInstance.Name, metav1.GetOptions{})
		if err == nil && brokerTemplateInstance.Spec.TemplateInstance.Kind == existingBrokerTemplateInstance.Spec.TemplateInstance.Kind && brokerTemplateInstance.Spec.TemplateInstance.Namespace == existingBrokerTemplateInstance.Spec.TemplateInstance.Namespace && brokerTemplateInstance.Spec.TemplateInstance.Name == existingBrokerTemplateInstance.Spec.TemplateInstance.Name && brokerTemplateInstance.Spec.Secret.Kind == existingBrokerTemplateInstance.Spec.Secret.Kind && brokerTemplateInstance.Spec.Secret.Namespace == existingBrokerTemplateInstance.Spec.Secret.Namespace && brokerTemplateInstance.Spec.Secret.Name == existingBrokerTemplateInstance.Spec.Secret.Name {
			return existingBrokerTemplateInstance, nil
		}
		return nil, api.NewResponse(http.StatusConflict, api.ProvisionResponse{}, nil)
	}
	return nil, api.InternalServerError(err)
}
func (b *Broker) Provision(u user.Info, instanceID string, preq *api.ProvisionRequest) *api.Response {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("Template service broker: Provision: instanceID %s", instanceID)
	if errs := ValidateProvisionRequest(preq); len(errs) > 0 {
		return api.BadRequest(errs.ToAggregate())
	}
	namespace := preq.Context.Namespace
	template, err := b.lister.GetByUID(preq.ServiceID)
	if err != nil && !kerrors.IsNotFound(err) {
		return api.BadRequest(err)
	}
	if template == nil {
		klog.V(4).Infof("Template service broker: GetByUID didn't template %s", preq.ServiceID)
	out:
		for namespace := range b.templateNamespaces {
			templates, err := b.lister.Templates(namespace).List(labels.Everything())
			if err != nil {
				return api.InternalServerError(err)
			}
			for _, t := range templates {
				if string(t.UID) == preq.ServiceID {
					template = t
					break out
				}
			}
		}
	}
	if template == nil {
		klog.V(4).Infof("Template service broker: template %s not found", preq.ServiceID)
		return api.BadRequest(kerrors.NewNotFound(templateapiv1.Resource("templates"), preq.ServiceID))
	}
	if _, ok := b.templateNamespaces[template.Namespace]; !ok {
		return api.BadRequest(kerrors.NewNotFound(templateapiv1.Resource("templates"), preq.ServiceID))
	}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: template.Namespace, Verb: "get", Group: templateapiv1.GroupName, Resource: "templates", Name: template.Name}); err != nil {
		return api.Forbidden(err)
	}
	if err := util.Authorize(b.kc.AuthorizationV1().SubjectAccessReviews(), u, &authorizationv1.ResourceAttributes{Namespace: namespace, Verb: "create", Group: templateapiv1.GroupName, Resource: "templateinstances", Name: instanceID}); err != nil {
		return api.Forbidden(err)
	}
	didWork := false
	brokerTemplateInstance, resp := b.ensureBrokerTemplateInstance(u, namespace, instanceID, &didWork)
	if resp != nil {
		return resp
	}
	time.Sleep(b.gcCreateDelay)
	secret, resp := b.ensureSecret(u, namespace, brokerTemplateInstance, instanceID, preq, template, &didWork)
	if resp != nil {
		return resp
	}
	templateInstance, resp := b.ensureTemplateInstance(u, namespace, brokerTemplateInstance, instanceID, template, secret, &didWork)
	if resp != nil {
		return resp
	}
	_, resp = b.ensureBrokerTemplateInstanceUIDs(u, namespace, brokerTemplateInstance, secret, templateInstance, &didWork)
	if resp != nil {
		return resp
	}
	if didWork {
		return api.NewResponse(http.StatusAccepted, api.ProvisionResponse{Operation: api.OperationProvisioning}, nil)
	}
	return api.NewResponse(http.StatusOK, api.ProvisionResponse{Operation: api.OperationProvisioning}, nil)
}
