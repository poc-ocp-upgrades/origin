package rest

import (
	"context"
	goformat "fmt"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type PassThroughStreamer struct {
	In          io.ReadCloser
	Flush       bool
	ContentType string
}

var _ rest.ResourceStreamer = &PassThroughStreamer{}

func (obj *PassThroughStreamer) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return schema.EmptyObjectKind
}
func (obj *PassThroughStreamer) DeepCopyObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	panic("passThroughStreamer does not implement DeepCopyObject")
}
func (s *PassThroughStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return s.In, s.Flush, s.ContentType, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
