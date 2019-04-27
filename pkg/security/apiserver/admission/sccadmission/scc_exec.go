package sccadmission

import (
	"fmt"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/client-go/kubernetes"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	coreapiv1conversions "k8s.io/kubernetes/pkg/apis/core/v1"
	securityv1informers "github.com/openshift/client-go/security/informers/externalversions"
	oadmission "github.com/openshift/origin/pkg/cmd/server/admission"
)

func RegisterSCCExecRestrictions(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register("security.openshift.io/SCCExecRestrictions", func(config io.Reader) (admission.Interface, error) {
		execAdmitter := NewSCCExecRestrictions()
		return execAdmitter, nil
	})
}

var (
	_	= initializer.WantsAuthorizer(&sccExecRestrictions{})
	_	= initializer.WantsExternalKubeClientSet(&sccExecRestrictions{})
	_	= oadmission.WantsSecurityInformer(&sccExecRestrictions{})
	_	= admission.ValidationInterface(&sccExecRestrictions{})
)

type sccExecRestrictions struct {
	*admission.Handler
	constraintAdmission	*constraint
	client			kubernetes.Interface
}

func (d *sccExecRestrictions) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.GetOperation() != admission.Connect {
		return nil
	}
	if a.GetResource().GroupResource() != coreapi.Resource("pods") {
		return nil
	}
	if a.GetSubresource() != "attach" && a.GetSubresource() != "exec" {
		return nil
	}
	pod, err := d.client.CoreV1().Pods(a.GetNamespace()).Get(a.GetName(), metav1.GetOptions{})
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	internalPod := &coreapi.Pod{}
	if err := coreapiv1conversions.Convert_v1_Pod_To_core_Pod(pod, internalPod, nil); err != nil {
		return admission.NewForbidden(a, err)
	}
	createAttributes := admission.NewAttributesRecord(internalPod, nil, coreapi.Kind("Pod").WithVersion(""), a.GetNamespace(), a.GetName(), a.GetResource(), "", admission.Create, false, a.GetUserInfo())
	if err := d.constraintAdmission.Admit(createAttributes); err != nil {
		return admission.NewForbidden(a, fmt.Errorf("%s operation is not allowed because the pod's security context exceeds your permissions: %v", a.GetSubresource(), err))
	}
	return nil
}
func NewSCCExecRestrictions() *sccExecRestrictions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &sccExecRestrictions{Handler: admission.NewHandler(admission.Connect), constraintAdmission: NewConstraint()}
}
func (d *sccExecRestrictions) SetExternalKubeClientSet(c kubernetes.Interface) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.client = c
	d.constraintAdmission.SetExternalKubeClientSet(c)
}
func (d *sccExecRestrictions) SetSecurityInformers(informers securityv1informers.SharedInformerFactory) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.constraintAdmission.SetSecurityInformers(informers)
}
func (d *sccExecRestrictions) SetAuthorizer(authorizer authorizer.Authorizer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.constraintAdmission.SetAuthorizer(authorizer)
}
func (d *sccExecRestrictions) ValidateInitialization() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d.constraintAdmission.ValidateInitialization()
}
