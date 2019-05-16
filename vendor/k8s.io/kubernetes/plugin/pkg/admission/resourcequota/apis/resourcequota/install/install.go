package install

import (
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	resourcequotaapi "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
	resourcequotav1alpha1 "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1"
	resourcequotav1beta1 "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1beta1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func Install(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(resourcequotaapi.AddToScheme(scheme))
	utilruntime.Must(resourcequotav1beta1.AddToScheme(scheme))
	utilruntime.Must(resourcequotav1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(resourcequotav1beta1.SchemeGroupVersion, resourcequotav1alpha1.SchemeGroupVersion))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
