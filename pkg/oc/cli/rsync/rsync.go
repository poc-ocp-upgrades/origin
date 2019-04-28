package rsync

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/origin/pkg/util/fsnotification"
)

const (
	RsyncRecommendedName	= "rsync"
	noRsyncUnixWarning	= "WARNING: rsync command not found in path. Please use your package manager to install it.\n"
	noRsyncWindowsWarning	= "WARNING: rsync command not found in path. Download cwRsync for Windows and add it to your PATH.\n"
)

var (
	rsyncLong	= templates.LongDesc(`
		Copy local files to or from a pod container

		This command will copy local files to or from a remote container.
		It only copies the changed files using the rsync command from your OS.
		To ensure optimum performance, install rsync locally. In UNIX systems,
		use your package manager. In Windows, install cwRsync from
		https://www.itefix.net/cwrsync.

		If no container is specified, the first container of the pod is used
		for the copy.

		The following flags are passed to rsync by default:
		--archive --no-owner --no-group --omit-dir-times --numeric-ids
		`)
	rsyncExample	= templates.Examples(`
	  # Synchronize a local directory with a pod directory
	  %[1]s ./local/dir/ POD:/remote/dir

	  # Synchronize a pod directory with a local directory
	  %[1]s POD:/remote/dir/ ./local/dir`)
	rsyncDefaultFlags	= []string{"--archive", "--no-owner", "--no-group", "--omit-dir-times", "--numeric-ids"}
)

type CopyStrategy interface {
	Copy(source, destination *PathSpec, out, errOut io.Writer) error
	Validate() error
	String() string
}
type executor interface {
	Execute(command []string, in io.Reader, out, err io.Writer) error
}
type forwarder interface {
	ForwardPorts(ports []string, stopChan <-chan struct{}) error
}
type podChecker interface{ CheckPod() error }
type RsyncOptions struct {
	Namespace		string
	ContainerName		string
	Source			*PathSpec
	Destination		*PathSpec
	Strategy		CopyStrategy
	StrategyName		string
	Quiet			bool
	Delete			bool
	Watch			bool
	Compress		bool
	SuggestedCmdUsage	string
	RshCmd			string
	RsyncInclude		[]string
	RsyncExclude		[]string
	RsyncProgress		bool
	RsyncNoPerms		bool
	Config			*rest.Config
	Client			kubernetes.Interface
	genericclioptions.IOStreams
}

func NewRsyncOptions(streams genericclioptions.IOStreams) *RsyncOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RsyncOptions{IOStreams: streams}
}
func NewCmdRsync(name, parent string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewRsyncOptions(streams)
	cmd := &cobra.Command{Use: fmt.Sprintf("%s SOURCE DESTINATION", name), Short: "Copy files between local filesystem and a pod", Long: rsyncLong, Example: fmt.Sprintf(rsyncExample, parent+" "+name), Run: func(c *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(f, c, args))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.RunRsync())
	}}
	cmd.Flags().StringVarP(&o.ContainerName, "container", "c", "", "Container within the pod")
	cmd.Flags().StringVar(&o.StrategyName, "strategy", "", "Specify which strategy to use for copy: rsync, rsync-daemon, or tar")
	cmd.Flags().BoolVarP(&o.Quiet, "quiet", "q", false, "Suppress non-error messages")
	cmd.Flags().BoolVar(&o.Delete, "delete", false, "If true, delete files not present in source")
	cmd.Flags().StringSliceVar(&o.RsyncExclude, "exclude", nil, "If true, exclude files matching specified pattern")
	cmd.Flags().StringSliceVar(&o.RsyncInclude, "include", nil, "If true, include files matching specified pattern")
	cmd.Flags().BoolVar(&o.RsyncProgress, "progress", false, "If true, show progress during transfer")
	cmd.Flags().BoolVar(&o.RsyncNoPerms, "no-perms", false, "If true, do not transfer permissions")
	cmd.Flags().BoolVarP(&o.Watch, "watch", "w", false, "Watch directory for changes and resync automatically")
	cmd.Flags().BoolVar(&o.Compress, "compress", false, "compress file data during the transfer")
	return cmd
}
func warnNoRsync(out io.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isWindows() {
		fmt.Fprintf(out, noRsyncWindowsWarning)
		return
	}
	fmt.Fprintf(out, noRsyncUnixWarning)
}
func (o *RsyncOptions) GetCopyStrategy(name string) (CopyStrategy, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch name {
	case "":
		return NewDefaultCopyStrategy(o), nil
	case "rsync":
		return NewRsyncStrategy(o), nil
	case "rsync-daemon":
		return NewRsyncDaemonStrategy(o), nil
	case "tar":
		return NewTarStrategy(o), nil
	default:
		return nil, fmt.Errorf("unknown strategy: %s", name)
	}
}
func (o *RsyncOptions) Complete(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch n := len(args); {
	case n == 0:
		cmd.Help()
		fallthrough
	case n < 2:
		return kcmdutil.UsageErrorf(cmd, "SOURCE_DIR and POD:DESTINATION_DIR are required arguments")
	case n > 2:
		return kcmdutil.UsageErrorf(cmd, "only SOURCE_DIR and POD:DESTINATION_DIR should be specified as arguments")
	}
	var err error
	if o.Config, err = f.ToRESTConfig(); err != nil {
		return err
	}
	if o.Client, err = kubernetes.NewForConfig(o.Config); err != nil {
		return err
	}
	namespace, _, err := f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	o.Namespace = namespace
	parsedSourcePath, err := resolveResourceKindPath(f, args[0], namespace)
	if err != nil {
		return err
	}
	parsedDestPath, err := resolveResourceKindPath(f, args[1], namespace)
	if err != nil {
		return err
	}
	o.Source, err = parsePathSpec(parsedSourcePath)
	if err != nil {
		return err
	}
	o.Destination, err = parsePathSpec(parsedDestPath)
	if err != nil {
		return err
	}
	fullCmdName := ""
	cmdParent := cmd.Parent()
	if cmdParent != nil {
		fullCmdName = cmdParent.CommandPath()
	}
	if len(fullCmdName) > 0 && kcmdutil.IsSiblingCommandExists(cmd, "describe") {
		o.SuggestedCmdUsage = fmt.Sprintf("Use '%s describe pod/%s -n %s' to see all of the containers in this pod.", fullCmdName, o.PodName(), o.Namespace)
	}
	o.RshCmd = DefaultRsyncRemoteShellToUse(cmd)
	o.Strategy, err = o.GetCopyStrategy(o.StrategyName)
	if err != nil {
		return err
	}
	return nil
}
func (o *RsyncOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.Out == nil || o.ErrOut == nil {
		return errors.New("output and error streams must be specified")
	}
	if o.Source == nil || o.Destination == nil {
		return errors.New("source and destination must be specified")
	}
	if err := o.Source.Validate(); err != nil {
		return err
	}
	if err := o.Destination.Validate(); err != nil {
		return err
	}
	if o.Source.Local() == o.Destination.Local() {
		return errors.New("rsync is only valid between a local directory and a pod directory; " + "specify a pod directory as [PODNAME]:[DIR]")
	}
	if o.Destination.Local() && o.Watch {
		return errors.New("\"--watch\" can only be used with a local source directory")
	}
	if err := o.Strategy.Validate(); err != nil {
		return err
	}
	return nil
}
func (o *RsyncOptions) RunRsync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := o.Strategy.Copy(o.Source, o.Destination, o.Out, o.ErrOut); err != nil {
		return err
	}
	if !o.Watch {
		return nil
	}
	return o.WatchAndSync()
}
func (o *RsyncOptions) WatchAndSync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		changeLock	sync.Mutex
		dirty		bool
		lastChange	time.Time
		watchError	error
	)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("error setting up filesystem watcher: %v", err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				changeLock.Lock()
				klog.V(5).Infof("filesystem watch event: %s", event)
				lastChange = time.Now()
				dirty = true
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					if e := watcher.Remove(event.Name); e != nil {
						klog.V(5).Infof("error removing watch for %s: %v", event.Name, e)
					}
				} else {
					if e := fsnotification.AddRecursiveWatch(watcher, event.Name); e != nil && watchError == nil {
						watchError = e
					}
				}
				changeLock.Unlock()
			case err := <-watcher.Errors:
				changeLock.Lock()
				watchError = fmt.Errorf("error watching filesystem for changes: %v", err)
				changeLock.Unlock()
			}
		}
	}()
	err = fsnotification.AddRecursiveWatch(watcher, o.Source.Path)
	if err != nil {
		return fmt.Errorf("error watching source path %s: %v", o.Source.Path, err)
	}
	delay := 2 * time.Second
	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		changeLock.Lock()
		if watchError != nil {
			return watchError
		}
		if dirty && time.Now().After(lastChange.Add(delay)) {
			klog.V(1).Info("Synchronizing filesystem changes...")
			err = o.Strategy.Copy(o.Source, o.Destination, o.Out, o.ErrOut)
			if err != nil {
				return err
			}
			klog.V(1).Info("Done.")
			dirty = false
		}
		changeLock.Unlock()
		<-ticker.C
	}
}
func (o *RsyncOptions) PodName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.Source.PodName) > 0 {
		return o.Source.PodName
	}
	return o.Destination.PodName
}
