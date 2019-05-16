package buildlog

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

type PipeStreamer struct {
	In          *io.PipeWriter
	Out         *io.PipeReader
	Flush       bool
	ContentType string
}

var _ rest.ResourceStreamer = &PipeStreamer{}

func (obj *PipeStreamer) GetObjectKind() schema.ObjectKind {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return schema.EmptyObjectKind
}
func (obj *PipeStreamer) DeepCopyObject() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	panic("buildlog.PipeStreamer does not implement DeepCopyObject")
}
func (s *PipeStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flush = s.Flush
	stream = s.Out
	return
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
