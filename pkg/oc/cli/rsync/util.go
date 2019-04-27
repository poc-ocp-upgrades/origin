package rsync

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"k8s.io/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	testRsyncCommand	= []string{"rsync", "--version"}
	testTarCommand		= []string{"tar", "--version"}
)

func executeWithLogging(e executor, cmd []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	w := &bytes.Buffer{}
	err := e.Execute(cmd, nil, w, w)
	klog.V(4).Infof("%s", w.String())
	klog.V(4).Infof("error: %v", err)
	return err
}
func isWindows() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return runtime.GOOS == "windows"
}
func hasLocalRsync() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := exec.LookPath("rsync")
	if err != nil {
		return false
	}
	return true
}
func isExitError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return false
	}
	_, exitErr := err.(*exec.ExitError)
	return exitErr
}
func checkRsync(e executor) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return executeWithLogging(e, testRsyncCommand)
}
func checkTar(e executor) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return executeWithLogging(e, testTarCommand)
}
func rsyncFlagsFromOptions(o *RsyncOptions) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := []string{}
	if o.Quiet {
		flags = append(flags, "-q")
	} else {
		flags = append(flags, "-v")
	}
	if o.Delete {
		flags = append(flags, "--delete")
	}
	if o.Compress {
		flags = append(flags, "-z")
	}
	if len(o.RsyncInclude) > 0 {
		for _, include := range o.RsyncInclude {
			flags = append(flags, fmt.Sprintf("--include=%s", include))
		}
	}
	if len(o.RsyncExclude) > 0 {
		for _, exclude := range o.RsyncExclude {
			flags = append(flags, fmt.Sprintf("--exclude=%s", exclude))
		}
	}
	if o.RsyncProgress {
		flags = append(flags, "--progress")
	}
	if o.RsyncNoPerms {
		flags = append(flags, "--no-perms")
	}
	return flags
}
func tarFlagsFromOptions(o *RsyncOptions) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := []string{}
	if !o.Quiet {
		flags = append(flags, "-v")
	}
	if len(o.RsyncInclude) > 0 {
		for _, include := range o.RsyncInclude {
			flags = append(flags, fmt.Sprintf("**/%s", include))
		}
		flags = append(flags, "*")
	}
	if len(o.RsyncExclude) > 0 {
		for _, exclude := range o.RsyncExclude {
			flags = append(flags, fmt.Sprintf("--exclude=%s", exclude))
		}
	}
	return flags
}
func rsyncSpecificFlags(o *RsyncOptions) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := []string{}
	if o.RsyncProgress {
		flags = append(flags, "--progress")
	}
	if o.RsyncNoPerms {
		flags = append(flags, "--no-perms")
	}
	if o.Compress {
		flags = append(flags, "-z")
	}
	return flags
}

type podAPIChecker struct {
	client		kubernetes.Interface
	namespace	string
	podName		string
}

func (p podAPIChecker) CheckPod() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := p.client.CoreV1().Pods(p.namespace).Get(p.podName, metav1.GetOptions{})
	return err
}
