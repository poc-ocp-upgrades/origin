package trigger

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/image/apis/image/v1/trigger"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
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

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
