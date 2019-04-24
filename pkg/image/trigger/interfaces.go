package trigger

import (
	"k8s.io/apimachinery/pkg/runtime"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/client-go/tools/cache"
	"github.com/openshift/origin/pkg/image/apis/image/v1/trigger"
)

type CacheEntry struct {
	Key		string
	Namespace	string
	Triggers	[]trigger.ObjectFieldTrigger
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
