package podsecuritypolicyselfsubjectreview

import (
	"context"
	"fmt"
	goformat "fmt"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityvalidation "github.com/openshift/origin/pkg/security/apis/security/validation"
	podsecuritypolicysubjectreview "github.com/openshift/origin/pkg/security/apiserver/registry/podsecuritypolicysubjectreview"
	scc "github.com/openshift/origin/pkg/security/apiserver/securitycontextconstraints"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
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
	return &securityapi.PodSecurityPolicySelfSubjectReview{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pspssr, ok := obj.(*securityapi.PodSecurityPolicySelfSubjectReview)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("not a PodSecurityPolicySelfSubjectReview: %#v", obj))
	}
	if errs := securityvalidation.ValidatePodSecurityPolicySelfSubjectReview(pspssr); len(errs) > 0 {
		return nil, apierrors.NewInvalid(coreapi.Kind("PodSecurityPolicySelfSubjectReview"), "", errs)
	}
	userInfo, ok := apirequest.UserFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("no user data associated with context"))
	}
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, apierrors.NewBadRequest("namespace parameter required.")
	}
	users := []user.Info{userInfo}
	saName := pspssr.Spec.Template.Spec.ServiceAccountName
	if len(saName) > 0 {
		users = append(users, serviceaccount.UserInfo(ns, saName, ""))
	}
	matchedConstraints, err := r.sccMatcher.FindApplicableSCCs(ns, users...)
	if err != nil {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("unable to find SecurityContextConstraints: %v", err))
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
		filled, err := podsecuritypolicysubjectreview.FillPodSecurityPolicySubjectReviewStatus(&pspssr.Status, provider, pspssr.Spec.Template.Spec, constraint)
		if err != nil {
			klog.Errorf("unable to fill PodSecurityPolicySelfSubjectReview from constraint %v", err)
			continue
		}
		if filled {
			return pspssr, nil
		}
	}
	return pspssr, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
