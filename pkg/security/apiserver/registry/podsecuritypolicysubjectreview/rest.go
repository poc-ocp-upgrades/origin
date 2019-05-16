package podsecuritypolicysubjectreview

import (
	"context"
	"fmt"
	goformat "fmt"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityvalidation "github.com/openshift/origin/pkg/security/apis/security/validation"
	scc "github.com/openshift/origin/pkg/security/apiserver/securitycontextconstraints"
	corev1 "k8s.io/api/core/v1"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	kapiref "k8s.io/kubernetes/pkg/api/ref"
	coreapi "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/serviceaccount"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct {
	sccMatcher scc.SCCMatcher
	client     kubernetes.Interface
}

var _ rest.Creater = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(m scc.SCCMatcher, c kubernetes.Interface) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{sccMatcher: m, client: c}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &securityapi.PodSecurityPolicySubjectReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pspsr, ok := obj.(*securityapi.PodSecurityPolicySubjectReview)
	if !ok {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("not a PodSecurityPolicySubjectReview: %#v", obj))
	}
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, kapierrors.NewBadRequest("namespace parameter required.")
	}
	if errs := securityvalidation.ValidatePodSecurityPolicySubjectReview(pspsr); len(errs) > 0 {
		return nil, kapierrors.NewInvalid(coreapi.Kind("PodSecurityPolicySubjectReview"), "", errs)
	}
	var users []user.Info
	specUser := &user.DefaultInfo{Name: pspsr.Spec.User, Groups: pspsr.Spec.Groups}
	if len(specUser.Name) > 0 || len(specUser.Groups) > 0 {
		users = append(users, specUser)
	}
	saName := pspsr.Spec.Template.Spec.ServiceAccountName
	if len(saName) > 0 {
		users = append(users, serviceaccount.UserInfo(ns, saName, ""))
	}
	matchedConstraints, err := r.sccMatcher.FindApplicableSCCs(ns, users...)
	if err != nil {
		return nil, kapierrors.NewBadRequest(fmt.Sprintf("unable to find SecurityContextConstraints: %v", err))
	}
	var namespace *corev1.Namespace
	for _, constraint := range matchedConstraints {
		var (
			provider scc.SecurityContextConstraintsProvider
			err      error
		)
		if provider, namespace, err = scc.CreateProviderFromConstraint(ns, namespace, constraint, r.client); err != nil {
			klog.Errorf("Unable to create provider for constraint: %v", err)
			continue
		}
		filled, err := FillPodSecurityPolicySubjectReviewStatus(&pspsr.Status, provider, pspsr.Spec.Template.Spec, constraint)
		if err != nil {
			klog.Errorf("unable to fill PodSecurityPolicySubjectReviewStatus from constraint %v", err)
			continue
		}
		if filled {
			return pspsr, nil
		}
	}
	return pspsr, nil
}
func FillPodSecurityPolicySubjectReviewStatus(s *securityapi.PodSecurityPolicySubjectReviewStatus, provider scc.SecurityContextConstraintsProvider, spec coreapi.PodSpec, constraint *securityapi.SecurityContextConstraints) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := &coreapi.Pod{Spec: spec}
	if errs := scc.AssignSecurityContext(provider, pod, field.NewPath(fmt.Sprintf("provider %s: ", provider.GetSCCName()))); len(errs) > 0 {
		klog.Errorf("unable to assign SecurityContextConstraints provider: %v", errs)
		s.Reason = "CantAssignSecurityContextConstraintProvider"
		return false, fmt.Errorf("unable to assign SecurityContextConstraints provider: %v", errs.ToAggregate())
	}
	ref, err := kapiref.GetReference(legacyscheme.Scheme, constraint)
	if err != nil {
		s.Reason = "CantObtainReference"
		return false, fmt.Errorf("unable to get SecurityContextConstraints reference: %v", err)
	}
	s.AllowedBy = ref
	if len(spec.ServiceAccountName) > 0 {
		s.Template.Spec = pod.Spec
	}
	return true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
