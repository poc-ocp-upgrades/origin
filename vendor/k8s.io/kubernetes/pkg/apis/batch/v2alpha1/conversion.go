package v2alpha1

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	for _, k := range []string{"Job", "JobTemplate", "CronJob"} {
		kind := k
		err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind(kind), func(label, value string) (string, string, error) {
			switch label {
			case "metadata.name", "metadata.namespace", "status.successful":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label %q not supported for %q", label, kind)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
