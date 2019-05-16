package options

import (
	goformat "fmt"
	"github.com/openshift/origin/pkg/cmd/util/variable"
	"github.com/spf13/pflag"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ImageFormatArgs struct{ ImageTemplate variable.ImageTemplate }

func BindImageFormatArgs(args *ImageFormatArgs, flags *pflag.FlagSet, prefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags.StringVar(&args.ImageTemplate.Format, "images", args.ImageTemplate.Format, "When fetching images used by the cluster for important components, use this format on both master and nodes. The latest release will be used by default.")
	flags.BoolVar(&args.ImageTemplate.Latest, "latest-images", args.ImageTemplate.Latest, "If true, attempt to use the latest images for the cluster instead of the latest release.")
}
func NewDefaultImageFormatArgs() *ImageFormatArgs {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := &ImageFormatArgs{ImageTemplate: variable.NewDefaultImageTemplate()}
	return config
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
