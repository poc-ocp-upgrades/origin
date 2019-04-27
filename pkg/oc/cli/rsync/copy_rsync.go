package rsync

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
)

type rsyncStrategy struct {
	Flags		[]string
	RshCommand	string
	LocalExecutor	executor
	RemoteExecutor	executor
	podChecker	podChecker
}

var rshExcludeFlags = sets.NewString("delete", "strategy", "quiet", "include", "exclude", "progress", "no-perms", "watch", "compress")

func DefaultRsyncRemoteShellToUse(cmd *cobra.Command) string {
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
	rshCmd := cmdutil.SiblingCommand(cmd, "rsh")
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if rshExcludeFlags.Has(flag.Name) {
			return
		}
		rshCmd = append(rshCmd, fmt.Sprintf("--%s=%s", flag.Name, flag.Value.String()))
	})
	return strings.Join(rsyncEscapeCommand(rshCmd), " ")
}
func NewRsyncStrategy(o *RsyncOptions) CopyStrategy {
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
	klog.V(4).Infof("Rsh command: %s", o.RshCmd)
	flags := []string{"--blocking-io"}
	flags = append(flags, rsyncDefaultFlags...)
	flags = append(flags, rsyncFlagsFromOptions(o)...)
	podName := o.Source.PodName
	if o.Source.Local() {
		podName = o.Destination.PodName
	}
	return &rsyncStrategy{Flags: flags, RshCommand: o.RshCmd, RemoteExecutor: newRemoteExecutor(o), LocalExecutor: newLocalExecutor(), podChecker: podAPIChecker{o.Client, o.Namespace, podName}}
}
func (r *rsyncStrategy) Copy(source, destination *PathSpec, out, errOut io.Writer) error {
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
	klog.V(3).Infof("Copying files with rsync")
	cmd := append([]string{"rsync"}, r.Flags...)
	cmd = append(cmd, "-e", r.RshCommand, source.RsyncPath(), destination.RsyncPath())
	errBuf := &bytes.Buffer{}
	err := r.LocalExecutor.Execute(cmd, nil, out, errBuf)
	if isExitError(err) {
		if podCheckErr := r.podChecker.CheckPod(); podCheckErr != nil {
			return podCheckErr
		}
		testRsyncErr := checkRsync(r.RemoteExecutor)
		if testRsyncErr != nil {
			return strategySetupError("rsync not available in container")
		}
	}
	io.Copy(errOut, errBuf)
	return err
}
func (r *rsyncStrategy) Validate() error {
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
	errs := []error{}
	if len(r.RshCommand) == 0 {
		errs = append(errs, errors.New("rsh command must be provided"))
	}
	if r.LocalExecutor == nil {
		errs = append(errs, errors.New("local executor must not be nil"))
	}
	if r.RemoteExecutor == nil {
		errs = append(errs, errors.New("remote executor must not be nil"))
	}
	if len(errs) > 0 {
		return kerrors.NewAggregate(errs)
	}
	return nil
}
func rsyncEscapeCommand(command []string) []string {
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
	var escapedCommand []string
	for _, val := range command {
		needsQuoted := strings.ContainsAny(val, `'" `)
		if needsQuoted {
			val = strings.Replace(val, `"`, `""`, -1)
			val = `"` + val + `"`
		}
		escapedCommand = append(escapedCommand, val)
	}
	return escapedCommand
}
func (r *rsyncStrategy) String() string {
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
	return "rsync"
}
