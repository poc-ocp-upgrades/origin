package printers

import (
 "io"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourcePrinter interface {
 PrintObj(runtime.Object, io.Writer) error
}
type ResourcePrinterFunc func(runtime.Object, io.Writer) error

func (fn ResourcePrinterFunc) PrintObj(obj runtime.Object, w io.Writer) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return fn(obj, w)
}

type PrintOptions struct {
 OutputFormatType     string
 OutputFormatArgument string
 NoHeaders            bool
 WithNamespace        bool
 WithKind             bool
 Wide                 bool
 ShowAll              bool
 ShowLabels           bool
 AbsoluteTimestamps   bool
 Kind                 schema.GroupKind
 ColumnLabels         []string
 SortBy               string
 AllowMissingKeys     bool
}
