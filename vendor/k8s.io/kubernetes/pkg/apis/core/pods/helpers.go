package pods

import (
	"fmt"
	goformat "fmt"
	"k8s.io/kubernetes/pkg/fieldpath"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ConvertDownwardAPIFieldLabel(version, label, value string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if version != "v1" {
		return "", "", fmt.Errorf("unsupported pod version: %s", version)
	}
	if path, _, ok := fieldpath.SplitMaybeSubscriptedPath(label); ok {
		switch path {
		case "metadata.annotations", "metadata.labels":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label does not support subscript: %s", label)
		}
	}
	switch label {
	case "metadata.annotations", "metadata.labels", "metadata.name", "metadata.namespace", "metadata.uid", "spec.nodeName", "spec.restartPolicy", "spec.serviceAccountName", "spec.schedulerName", "status.phase", "status.hostIP", "status.podIP":
		return label, value, nil
	case "spec.host":
		return "spec.nodeName", value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
