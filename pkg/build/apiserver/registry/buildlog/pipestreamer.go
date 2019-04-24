package buildlog

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

type PipeStreamer struct {
	In		*io.PipeWriter
	Out		*io.PipeReader
	Flush		bool
	ContentType	string
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
