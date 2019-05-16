package variable

import (
	"fmt"
	goformat "fmt"
	"k8s.io/klog"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type ImageTemplate struct {
	Format    string
	Latest    bool
	EnvFormat string
}

var (
	DefaultImagePrefix string
	defaultImageFormat = DefaultImagePrefix + "-${component}:${version}"
)

const defaultImageEnvFormat = "OPENSHIFT_%s_IMAGE"

func NewDefaultImageTemplate() ImageTemplate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ImageTemplate{Format: defaultImageFormat, Latest: false, EnvFormat: defaultImageEnvFormat}
}
func (t *ImageTemplate) ExpandOrDie(component string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value, err := t.Expand(component)
	if err != nil {
		klog.Fatalf("Unable to find an image for %q due to an error processing the format: %v", component, err)
	}
	return value
}
func (t *ImageTemplate) Expand(component string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	template := t.Format
	if len(t.EnvFormat) > 0 {
		if s, ok := t.imageComponentEnvExpander(component); ok {
			template = s
		}
	}
	value, err := ExpandStrict(template, func(key string) (string, bool) {
		switch key {
		case "component":
			return component, true
		case "version":
			if t.Latest {
				return "latest", true
			}
		}
		return "", false
	}, Versions)
	return value, err
}
func (t *ImageTemplate) imageComponentEnvExpander(key string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	s := strings.Replace(strings.ToUpper(key), "-", "_", -1)
	val := os.Getenv(fmt.Sprintf(t.EnvFormat, s))
	if len(val) == 0 {
		return "", false
	}
	return val, true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
