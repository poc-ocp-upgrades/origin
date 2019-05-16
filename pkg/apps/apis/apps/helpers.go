package apps

import (
	goformat "fmt"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func DeploymentToPodLogOptions(opts *DeploymentLogOptions) *kapi.PodLogOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &kapi.PodLogOptions{Container: opts.Container, Follow: opts.Follow, SinceSeconds: opts.SinceSeconds, SinceTime: opts.SinceTime, Timestamps: opts.Timestamps, TailLines: opts.TailLines, LimitBytes: opts.LimitBytes}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
