package componentconfigs

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	AddToScheme(Scheme)
}
func AddToScheme(scheme *runtime.Scheme) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(Known.AddToScheme(scheme))
}
