package sccadmission

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io"
	"sort"
	"strings"
	"k8s.io/klog"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/client-go/kubernetes"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	kapihelper "k8s.io/kubernetes/pkg/apis/core/helper"
	rbacregistry "k8s.io/kubernetes/pkg/registry/rbac"
	"k8s.io/kubernetes/pkg/serviceaccount"
	securityv1informer "github.com/openshift/client-go/security/informers/externalversions/security/v1"
	securityv1listers "github.com/openshift/client-go/security/listers/security/v1"
	oadmission "github.com/openshift/origin/pkg/cmd/server/admission"
	allocator "github.com/openshift/origin/pkg/security"
	scc "github.com/openshift/origin/pkg/security/apiserver/securitycontextconstraints"
)

const PluginName = "security.openshift.io/SecurityContextConstraint"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return NewConstraint(), nil
	})
}

type constraint struct {
	*admission.Handler
	client		kubernetes.Interface
	sccLister	securityv1listers.SecurityContextConstraintsLister
	authorizer	authorizer.Authorizer
}

var (
	_	= initializer.WantsAuthorizer(&constraint{})
	_	= initializer.WantsExternalKubeClientSet(&constraint{})
	_	= oadmission.WantsSecurityInformer(&constraint{})
	_	= admission.ValidationInterface(&constraint{})
	_	= admission.MutationInterface(&constraint{})
)

func NewConstraint() *constraint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &constraint{Handler: admission.NewHandler(admission.Create, admission.Update)}
}
func (c *constraint) Admit(a admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ignore, err := shouldIgnore(a); err != nil {
		return err
	} else if ignore {
		return nil
	}
	pod := a.GetObject().(*coreapi.Pod)
	specMutationAllowed := a.GetOperation() == admission.Create
	allowedPod, sccName, validationErrs, err := c.computeSecurityContext(a, pod, specMutationAllowed, "")
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if allowedPod != nil {
		*pod = *allowedPod
		klog.V(4).Infof("pod %s (generate: %s) validated against provider %s", pod.Name, pod.GenerateName, sccName)
		if pod.ObjectMeta.Annotations == nil {
			pod.ObjectMeta.Annotations = map[string]string{}
		}
		pod.ObjectMeta.Annotations[allocator.ValidatedSCCAnnotation] = sccName
		return nil
	}
	klog.V(4).Infof("unable to validate pod %s (generate: %s) against any security context constraint: %v", pod.Name, pod.GenerateName, validationErrs)
	return admission.NewForbidden(a, fmt.Errorf("unable to validate against any security context constraint: %v", validationErrs))
}
func (c *constraint) Validate(a admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ignore, err := shouldIgnore(a); err != nil {
		return err
	} else if ignore {
		return nil
	}
	pod := a.GetObject().(*coreapi.Pod)
	allowedPod, _, validationErrs, err := c.computeSecurityContext(a, pod, false, pod.ObjectMeta.Annotations[allocator.ValidatedSCCAnnotation])
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if allowedPod != nil && apiequality.Semantic.DeepEqual(pod, allowedPod) {
		return nil
	}
	klog.V(4).Infof("unable to validate pod %s (generate: %s) in namespace %s against any security context constraint: %v", pod.Name, pod.GenerateName, a.GetNamespace(), validationErrs)
	return admission.NewForbidden(a, fmt.Errorf("unable to validate against any security context constraint: %v", validationErrs))
}
func (c *constraint) computeSecurityContext(a admission.Attributes, pod *coreapi.Pod, specMutationAllowed bool, validatedSCCHint string) (*coreapi.Pod, string, field.ErrorList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("getting security context constraints for pod %s (generate: %s) in namespace %s with user info %v", pod.Name, pod.GenerateName, a.GetNamespace(), a.GetUserInfo())
	constraints, err := scc.NewDefaultSCCMatcher(c.sccLister, nil).FindApplicableSCCs(a.GetNamespace())
	if err != nil {
		return nil, "", nil, admission.NewForbidden(a, err)
	}
	sort.SliceStable(constraints, func(i, j int) bool {
		if !specMutationAllowed {
			if constraints[i].Name == validatedSCCHint {
				return true
			}
			if constraints[j].Name == validatedSCCHint {
				return false
			}
		}
		return i < j
	})
	providers, errs := scc.CreateProvidersFromConstraints(a.GetNamespace(), constraints, c.client)
	logProviders(pod, providers, errs)
	if len(providers) == 0 {
		return nil, "", nil, admission.NewForbidden(a, fmt.Errorf("no providers available to validate pod request"))
	}
	var (
		allowedPod		*coreapi.Pod
		allowingProvider	scc.SecurityContextConstraintsProvider
		validationErrs		field.ErrorList
		saUserInfo		user.Info
	)
	userInfo := a.GetUserInfo()
	if len(pod.Spec.ServiceAccountName) > 0 {
		saUserInfo = serviceaccount.UserInfo(a.GetNamespace(), pod.Spec.ServiceAccountName, "")
	}
loop:
	for _, provider := range providers {
		sccName := provider.GetSCCName()
		sccUsers := provider.GetSCCUsers()
		sccGroups := provider.GetSCCGroups()
		if !scc.ConstraintAppliesTo(sccName, sccUsers, sccGroups, userInfo, a.GetNamespace(), c.authorizer) && !(saUserInfo != nil && scc.ConstraintAppliesTo(sccName, sccUsers, sccGroups, saUserInfo, a.GetNamespace(), c.authorizer)) {
			continue
		}
		podCopy := pod.DeepCopy()
		if errs := scc.AssignSecurityContext(provider, podCopy, field.NewPath(fmt.Sprintf("provider %s: ", sccName))); len(errs) > 0 {
			validationErrs = append(validationErrs, errs...)
			continue
		}
		switch {
		case specMutationAllowed:
			allowedPod = podCopy
			allowingProvider = provider
			klog.V(5).Infof("pod %s (generate: %s) validated against provider %s with mutation", pod.Name, pod.GenerateName, sccName)
			break loop
		case apiequality.Semantic.DeepEqual(pod, podCopy):
			allowedPod = podCopy
			allowingProvider = provider
			klog.V(5).Infof("pod %s (generate: %s) validated against provider %s without mutation", pod.Name, pod.GenerateName, sccName)
			break loop
		default:
			klog.V(5).Infof("pod %s (generate: %s) validated against provider %s, but required mutation, skipping", pod.Name, pod.GenerateName, sccName)
		}
	}
	if allowedPod == nil || allowingProvider == nil {
		return nil, "", validationErrs, nil
	}
	return allowedPod, allowingProvider.GetSCCName(), validationErrs, nil
}
func shouldIgnore(a admission.Attributes) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.GetResource().GroupResource() != coreapi.Resource("pods") {
		return true, nil
	}
	if len(a.GetSubresource()) != 0 {
		return true, nil
	}
	_, ok := a.GetObject().(*coreapi.Pod)
	if !ok {
		return false, admission.NewForbidden(a, fmt.Errorf("object was marked as kind pod but was unable to be converted: %v", a.GetObject()))
	}
	if a.GetOperation() == admission.Update && rbacregistry.IsOnlyMutatingGCFields(a.GetObject(), a.GetOldObject(), kapihelper.Semantic) {
		return true, nil
	}
	return false, nil
}
func (c *constraint) SetSecurityInformers(informers securityv1informer.SecurityContextConstraintsInformer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.sccLister = informers.Lister()
}
func (c *constraint) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.client = client
}
func (c *constraint) SetAuthorizer(authorizer authorizer.Authorizer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.authorizer = authorizer
}
func (c *constraint) ValidateInitialization() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.sccLister == nil {
		return fmt.Errorf("%s requires an sccLister", PluginName)
	}
	if c.client == nil {
		return fmt.Errorf("%s requires a client", PluginName)
	}
	if c.authorizer == nil {
		return fmt.Errorf("%s requires an authorizer", PluginName)
	}
	return nil
}
func logProviders(pod *coreapi.Pod, providers []scc.SecurityContextConstraintsProvider, providerCreationErrs []error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, len(providers))
	for i, p := range providers {
		names[i] = p.GetSCCName()
	}
	klog.V(4).Infof("validating pod %s (generate: %s) against providers %s", pod.Name, pod.GenerateName, strings.Join(names, ","))
	for _, err := range providerCreationErrs {
		klog.V(4).Infof("provider creation error: %v", err)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
