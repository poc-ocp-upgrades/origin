package podsecuritypolicy

import (
	"fmt"
	goformat "fmt"
	"io"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninit "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/client-go/informers"
	policylisters "k8s.io/client-go/listers/policy/v1beta1"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/policy"
	rbacregistry "k8s.io/kubernetes/pkg/registry/rbac"
	psp "k8s.io/kubernetes/pkg/security/podsecuritypolicy"
	psputil "k8s.io/kubernetes/pkg/security/podsecuritypolicy/util"
	"k8s.io/kubernetes/pkg/serviceaccount"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

const (
	PluginName = "PodSecurityPolicy"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		plugin := newPlugin(psp.NewSimpleStrategyFactory(), true)
		return plugin, nil
	})
}

type PodSecurityPolicyPlugin struct {
	*admission.Handler
	strategyFactory  psp.StrategyFactory
	failOnNoPolicies bool
	authz            authorizer.Authorizer
	lister           policylisters.PodSecurityPolicyLister
}

func (plugin *PodSecurityPolicyPlugin) SetAuthorizer(authz authorizer.Authorizer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.authz = authz
}
func (plugin *PodSecurityPolicyPlugin) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if plugin.authz == nil {
		return fmt.Errorf("%s requires an authorizer", PluginName)
	}
	if plugin.lister == nil {
		return fmt.Errorf("%s requires a lister", PluginName)
	}
	return nil
}

var _ admission.MutationInterface = &PodSecurityPolicyPlugin{}
var _ admission.ValidationInterface = &PodSecurityPolicyPlugin{}
var _ genericadmissioninit.WantsAuthorizer = &PodSecurityPolicyPlugin{}
var _ genericadmissioninit.WantsExternalKubeInformerFactory = &PodSecurityPolicyPlugin{}
var auditKeyPrefix = strings.ToLower(PluginName) + "." + policy.GroupName + ".k8s.io"

func newPlugin(strategyFactory psp.StrategyFactory, failOnNoPolicies bool) *PodSecurityPolicyPlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &PodSecurityPolicyPlugin{Handler: admission.NewHandler(admission.Create, admission.Update), strategyFactory: strategyFactory, failOnNoPolicies: failOnNoPolicies}
}
func (a *PodSecurityPolicyPlugin) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	podSecurityPolicyInformer := f.Policy().V1beta1().PodSecurityPolicies()
	a.lister = podSecurityPolicyInformer.Lister()
	a.SetReadyFunc(podSecurityPolicyInformer.Informer().HasSynced)
}
func (c *PodSecurityPolicyPlugin) Admit(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ignore, err := shouldIgnore(a); err != nil {
		return err
	} else if ignore {
		return nil
	}
	if a.GetOperation() != admission.Create {
		return nil
	}
	pod := a.GetObject().(*api.Pod)
	allowedPod, pspName, validationErrs, err := c.computeSecurityContext(a, pod, true, "")
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if allowedPod != nil {
		*pod = *allowedPod
		klog.V(4).Infof("pod %s (generate: %s) in namespace %s validated against provider %s", pod.Name, pod.GenerateName, a.GetNamespace(), pspName)
		if pod.ObjectMeta.Annotations == nil {
			pod.ObjectMeta.Annotations = map[string]string{}
		}
		pod.ObjectMeta.Annotations[psputil.ValidatedPSPAnnotation] = pspName
		key := auditKeyPrefix + "/" + "admit-policy"
		if err := a.AddAnnotation(key, pspName); err != nil {
			klog.Warningf("failed to set admission audit annotation %s to %s: %v", key, pspName, err)
		}
		return nil
	}
	klog.V(4).Infof("unable to validate pod %s (generate: %s) in namespace %s against any pod security policy: %v", pod.Name, pod.GenerateName, a.GetNamespace(), validationErrs)
	return admission.NewForbidden(a, fmt.Errorf("unable to validate against any pod security policy: %v", validationErrs))
}
func (c *PodSecurityPolicyPlugin) Validate(a admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ignore, err := shouldIgnore(a); err != nil {
		return err
	} else if ignore {
		return nil
	}
	pod := a.GetObject().(*api.Pod)
	allowedPod, pspName, validationErrs, err := c.computeSecurityContext(a, pod, false, pod.ObjectMeta.Annotations[psputil.ValidatedPSPAnnotation])
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if apiequality.Semantic.DeepEqual(pod, allowedPod) {
		key := auditKeyPrefix + "/" + "validate-policy"
		if err := a.AddAnnotation(key, pspName); err != nil {
			klog.Warningf("failed to set admission audit annotation %s to %s: %v", key, pspName, err)
		}
		return nil
	}
	klog.V(4).Infof("unable to validate pod %s (generate: %s) in namespace %s against any pod security policy: %v", pod.Name, pod.GenerateName, a.GetNamespace(), validationErrs)
	return admission.NewForbidden(a, fmt.Errorf("unable to validate against any pod security policy: %v", validationErrs))
}
func shouldIgnore(a admission.Attributes) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != api.Resource("pods") {
		return true, nil
	}
	if len(a.GetSubresource()) != 0 {
		return true, nil
	}
	if _, ok := a.GetObject().(*api.Pod); !ok {
		return false, admission.NewForbidden(a, fmt.Errorf("unexpected type %T", a.GetObject()))
	}
	if a.GetOperation() == admission.Update && rbacregistry.IsOnlyMutatingGCFields(a.GetObject(), a.GetOldObject(), apiequality.Semantic) {
		return true, nil
	}
	return false, nil
}
func (c *PodSecurityPolicyPlugin) computeSecurityContext(a admission.Attributes, pod *api.Pod, specMutationAllowed bool, validatedPSPHint string) (*api.Pod, string, field.ErrorList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("getting pod security policies for pod %s (generate: %s)", pod.Name, pod.GenerateName)
	var saInfo user.Info
	if len(pod.Spec.ServiceAccountName) > 0 {
		saInfo = serviceaccount.UserInfo(a.GetNamespace(), pod.Spec.ServiceAccountName, "")
	}
	policies, err := c.lister.List(labels.Everything())
	if err != nil {
		return nil, "", nil, err
	}
	if len(policies) == 0 && !c.failOnNoPolicies {
		return pod, "", nil, nil
	}
	sort.SliceStable(policies, func(i, j int) bool {
		if !specMutationAllowed {
			if policies[i].Name == validatedPSPHint {
				return true
			}
			if policies[j].Name == validatedPSPHint {
				return false
			}
		}
		return strings.Compare(policies[i].Name, policies[j].Name) < 0
	})
	providers, errs := c.createProvidersFromPolicies(policies, pod.Namespace)
	for _, err := range errs {
		klog.V(4).Infof("provider creation error: %v", err)
	}
	if len(providers) == 0 {
		return nil, "", nil, fmt.Errorf("no providers available to validate pod request")
	}
	var (
		allowedMutatedPod   *api.Pod
		allowingMutatingPSP string
		validationErrs      = map[string]field.ErrorList{}
	)
	for _, provider := range providers {
		podCopy := pod.DeepCopy()
		if errs := assignSecurityContext(provider, podCopy); len(errs) > 0 {
			validationErrs[provider.GetPSPName()] = errs
			continue
		}
		mutated := !apiequality.Semantic.DeepEqual(pod, podCopy)
		if mutated && !specMutationAllowed {
			continue
		}
		if !isAuthorizedForPolicy(a.GetUserInfo(), saInfo, a.GetNamespace(), provider.GetPSPName(), c.authz) {
			continue
		}
		switch {
		case !mutated:
			return podCopy, provider.GetPSPName(), nil, nil
		case specMutationAllowed && allowedMutatedPod == nil:
			allowedMutatedPod = podCopy
			allowingMutatingPSP = provider.GetPSPName()
		}
	}
	if allowedMutatedPod != nil {
		return allowedMutatedPod, allowingMutatingPSP, nil, nil
	}
	aggregate := field.ErrorList{}
	for psp, errs := range validationErrs {
		if isAuthorizedForPolicy(a.GetUserInfo(), saInfo, a.GetNamespace(), psp, c.authz) {
			aggregate = append(aggregate, errs...)
		}
	}
	return nil, "", aggregate, nil
}
func assignSecurityContext(provider psp.Provider, pod *api.Pod) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := field.ErrorList{}
	err := provider.DefaultPodSecurityContext(pod)
	if err != nil {
		errs = append(errs, field.Invalid(field.NewPath("spec", "securityContext"), pod.Spec.SecurityContext, err.Error()))
	}
	errs = append(errs, provider.ValidatePod(pod)...)
	for i := range pod.Spec.InitContainers {
		err := provider.DefaultContainerSecurityContext(pod, &pod.Spec.InitContainers[i])
		if err != nil {
			errs = append(errs, field.Invalid(field.NewPath("spec", "initContainers").Index(i).Child("securityContext"), "", err.Error()))
			continue
		}
		errs = append(errs, provider.ValidateContainer(pod, &pod.Spec.InitContainers[i], field.NewPath("spec", "initContainers").Index(i))...)
	}
	for i := range pod.Spec.Containers {
		err := provider.DefaultContainerSecurityContext(pod, &pod.Spec.Containers[i])
		if err != nil {
			errs = append(errs, field.Invalid(field.NewPath("spec", "containers").Index(i).Child("securityContext"), "", err.Error()))
			continue
		}
		errs = append(errs, provider.ValidateContainer(pod, &pod.Spec.Containers[i], field.NewPath("spec", "containers").Index(i))...)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
func (c *PodSecurityPolicyPlugin) createProvidersFromPolicies(psps []*policyv1beta1.PodSecurityPolicy, namespace string) ([]psp.Provider, []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		providers []psp.Provider
		errs      []error
	)
	for _, constraint := range psps {
		provider, err := psp.NewSimpleProvider(constraint, namespace, c.strategyFactory)
		if err != nil {
			errs = append(errs, fmt.Errorf("error creating provider for PSP %s: %v", constraint.Name, err))
			continue
		}
		providers = append(providers, provider)
	}
	return providers, errs
}
func isAuthorizedForPolicy(user, sa user.Info, namespace, policyName string, authz authorizer.Authorizer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return authorizedForPolicy(sa, namespace, policyName, authz) || authorizedForPolicy(user, namespace, policyName, authz)
}
func authorizedForPolicy(info user.Info, namespace string, policyName string, authz authorizer.Authorizer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return authorizedForPolicyInAPIGroup(info, namespace, policyName, policy.GroupName, authz) || authorizedForPolicyInAPIGroup(info, namespace, policyName, extensions.GroupName, authz)
}
func authorizedForPolicyInAPIGroup(info user.Info, namespace, policyName, apiGroupName string, authz authorizer.Authorizer) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if info == nil {
		return false
	}
	attr := buildAttributes(info, namespace, policyName, apiGroupName)
	decision, reason, err := authz.Authorize(attr)
	if err != nil {
		klog.V(5).Infof("cannot authorize for policy: %v,%v", reason, err)
	}
	return (decision == authorizer.DecisionAllow)
}
func buildAttributes(info user.Info, namespace, policyName, apiGroupName string) authorizer.Attributes {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attr := authorizer.AttributesRecord{User: info, Verb: "use", Namespace: namespace, Name: policyName, APIGroup: apiGroupName, Resource: "podsecuritypolicies", ResourceRequest: true}
	return attr
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
