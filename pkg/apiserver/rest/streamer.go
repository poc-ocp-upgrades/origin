package rest

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

type PassThroughStreamer struct {
	In		io.ReadCloser
	Flush		bool
	ContentType	string
}

var _ rest.ResourceStreamer = &PassThroughStreamer{}

func (obj *PassThroughStreamer) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.EmptyObjectKind
}
func (obj *PassThroughStreamer) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("passThroughStreamer does not implement DeepCopyObject")
}
func (s *PassThroughStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.In, s.Flush, s.ContentType, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
