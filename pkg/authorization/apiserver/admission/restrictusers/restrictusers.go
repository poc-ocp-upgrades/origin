package restrictusers

import (
	"errors"
	"fmt"
	goformat "fmt"
	userapi "github.com/openshift/api/user/v1"
	authorizationtypedclient "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	userclient "github.com/openshift/client-go/user/clientset/versioned"
	userinformer "github.com/openshift/client-go/user/informers/externalversions"
	oadmission "github.com/openshift/origin/pkg/cmd/server/admission"
	usercache "github.com/openshift/origin/pkg/user/cache"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/apis/rbac"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register("authorization.openshift.io/RestrictSubjectBindings", func(config io.Reader) (admission.Interface, error) {
		return NewRestrictUsersAdmission()
	})
}

type GroupCache interface {
	GroupsFor(string) ([]*userapi.Group, error)
}
type restrictUsersAdmission struct {
	*admission.Handler
	roleBindingRestrictionsGetter authorizationtypedclient.RoleBindingRestrictionsGetter
	userClient                    userclient.Interface
	kubeClient                    kubernetes.Interface
	groupCache                    GroupCache
}

var _ = oadmission.WantsRESTClientConfig(&restrictUsersAdmission{})
var _ = oadmission.WantsUserInformer(&restrictUsersAdmission{})
var _ = initializer.WantsExternalKubeClientSet(&restrictUsersAdmission{})
var _ = admission.ValidationInterface(&restrictUsersAdmission{})

func NewRestrictUsersAdmission() (admission.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &restrictUsersAdmission{Handler: admission.NewHandler(admission.Create, admission.Update)}, nil
}
func (q *restrictUsersAdmission) SetExternalKubeClientSet(c kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.kubeClient = c
}
func (q *restrictUsersAdmission) SetRESTClientConfig(restClientConfig rest.Config) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	jsonClientConfig := rest.CopyConfig(&restClientConfig)
	jsonClientConfig.ContentConfig.AcceptContentTypes = "application/json"
	jsonClientConfig.ContentConfig.ContentType = "application/json"
	q.roleBindingRestrictionsGetter, err = authorizationtypedclient.NewForConfig(jsonClientConfig)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	q.userClient, err = userclient.NewForConfig(&restClientConfig)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
}
func (q *restrictUsersAdmission) SetUserInformer(userInformers userinformer.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	q.groupCache = usercache.NewGroupCache(userInformers.User().V1().Groups())
}
func subjectsDelta(elementsToIgnore, elements []rbac.Subject) []rbac.Subject {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result := []rbac.Subject{}
	for _, el := range elements {
		keep := true
		for _, skipEl := range elementsToIgnore {
			if el == skipEl {
				keep = false
				break
			}
		}
		if keep {
			result = append(result, el)
		}
	}
	return result
}
func (q *restrictUsersAdmission) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != rbac.Resource("rolebindings") {
		return nil
	}
	if len(a.GetSubresource()) != 0 {
		return nil
	}
	ns := a.GetNamespace()
	if len(ns) == 0 {
		return nil
	}
	var oldSubjects []rbac.Subject
	obj, oldObj := a.GetObject(), a.GetOldObject()
	rolebinding, ok := obj.(*rbac.RoleBinding)
	if !ok {
		return admission.NewForbidden(a, fmt.Errorf("wrong object type for new rolebinding: %T", obj))
	}
	if len(rolebinding.Subjects) == 0 {
		klog.V(4).Infof("No new subjects; admitting")
		return nil
	}
	if oldObj != nil {
		oldrolebinding, ok := oldObj.(*rbac.RoleBinding)
		if !ok {
			return admission.NewForbidden(a, fmt.Errorf("wrong object type for old rolebinding: %T", oldObj))
		}
		oldSubjects = oldrolebinding.Subjects
	}
	klog.V(4).Infof("Handling rolebinding %s/%s", rolebinding.Namespace, rolebinding.Name)
	newSubjects := subjectsDelta(oldSubjects, rolebinding.Subjects)
	if len(newSubjects) == 0 {
		klog.V(4).Infof("No new subjects; admitting")
		return nil
	}
	roleBindingRestrictionList, err := q.roleBindingRestrictionsGetter.RoleBindingRestrictions(ns).List(metav1.ListOptions{})
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if len(roleBindingRestrictionList.Items) == 0 {
		klog.V(4).Infof("No rolebinding restrictions specified; admitting")
		return nil
	}
	checkers := []SubjectChecker{}
	for _, rbr := range roleBindingRestrictionList.Items {
		checker, err := NewSubjectChecker(&rbr.Spec)
		if err != nil {
			return admission.NewForbidden(a, err)
		}
		checkers = append(checkers, checker)
	}
	roleBindingRestrictionContext, err := newRoleBindingRestrictionContext(ns, q.kubeClient, q.userClient.UserV1(), q.groupCache)
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	checker := NewUnionSubjectChecker(checkers)
	errs := []error{}
	for _, subject := range newSubjects {
		allowed, err := checker.Allowed(subject, roleBindingRestrictionContext)
		if err != nil {
			errs = append(errs, err)
		}
		if !allowed {
			errs = append(errs, fmt.Errorf("rolebindings to %s %q are not allowed in project %q", subject.Kind, subject.Name, ns))
		}
	}
	if len(errs) != 0 {
		return admission.NewForbidden(a, kerrors.NewAggregate(errs))
	}
	klog.V(4).Infof("All new subjects are allowed; admitting")
	return nil
}
func (q *restrictUsersAdmission) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if q.kubeClient == nil {
		return errors.New("RestrictUsersAdmission plugin requires a Kubernetes client")
	}
	if q.roleBindingRestrictionsGetter == nil {
		return errors.New("RestrictUsersAdmission plugin requires an OpenShift client")
	}
	if q.userClient == nil {
		return errors.New("RestrictUsersAdmission plugin requires an OpenShift user client")
	}
	if q.groupCache == nil {
		return errors.New("RestrictUsersAdmission plugin requires a group cache")
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
