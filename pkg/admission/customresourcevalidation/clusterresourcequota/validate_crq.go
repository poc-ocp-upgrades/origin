package clusterresourcequota

import (
	"fmt"
	goformat "fmt"
	quotav1 "github.com/openshift/api/quota/v1"
	"github.com/openshift/origin/pkg/admission/customresourcevalidation"
	quotavalidation "github.com/openshift/origin/pkg/admission/customresourcevalidation/clusterresourcequota/validation"
	"io"
	"k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "quota.openshift.io/ValidateClusterResourceQuota"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		return customresourcevalidation.NewValidator(map[schema.GroupResource]bool{{Group: quotav1.GroupName, Resource: "clusterresourcequotas"}: true}, map[schema.GroupVersionKind]customresourcevalidation.ObjectValidator{quotav1.GroupVersion.WithKind("ClusterResourceQuota"): clusterResourceQuotaV1{}})
	})
}
func toClusterResourceQuota(uncastObj runtime.Object) (*quotav1.ClusterResourceQuota, field.ErrorList) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if uncastObj == nil {
		return nil, nil
	}
	allErrs := field.ErrorList{}
	obj, ok := uncastObj.(*quotav1.ClusterResourceQuota)
	if !ok {
		return nil, append(allErrs, field.NotSupported(field.NewPath("kind"), fmt.Sprintf("%T", uncastObj), []string{"ClusterResourceQuota"}), field.NotSupported(field.NewPath("apiVersion"), fmt.Sprintf("%T", uncastObj), []string{quotav1.GroupVersion.String()}))
	}
	return obj, nil
}

type clusterResourceQuotaV1 struct{}

func (clusterResourceQuotaV1) ValidateCreate(obj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterResourceQuotaObj, errs := toClusterResourceQuota(obj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&clusterResourceQuotaObj.ObjectMeta, false, validation.NameIsDNSSubdomain, field.NewPath("metadata"))...)
	errs = append(errs, quotavalidation.ValidateClusterResourceQuota(clusterResourceQuotaObj)...)
	return errs
}
func (clusterResourceQuotaV1) ValidateUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterResourceQuotaObj, errs := toClusterResourceQuota(obj)
	if len(errs) > 0 {
		return errs
	}
	clusterResourceQuotaOldObj, errs := toClusterResourceQuota(oldObj)
	if len(errs) > 0 {
		return errs
	}
	errs = append(errs, validation.ValidateObjectMeta(&clusterResourceQuotaObj.ObjectMeta, false, validation.NameIsDNSSubdomain, field.NewPath("metadata"))...)
	errs = append(errs, quotavalidation.ValidateClusterResourceQuotaUpdate(clusterResourceQuotaObj, clusterResourceQuotaOldObj)...)
	return errs
}
func (c clusterResourceQuotaV1) ValidateStatusUpdate(obj runtime.Object, oldObj runtime.Object) field.ErrorList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.ValidateUpdate(obj, oldObj)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
