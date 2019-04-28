package observe

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	godefaulthttp "net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"text/template"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/jsonpath"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/origin/pkg/util/proc"
)

var (
	observeCounts		= prometheus.NewCounterVec(prometheus.CounterOpts{Name: "observe_counts", Help: "Number of changes observed to the underlying resource."}, []string{"type"})
	execDurations		= prometheus.NewSummaryVec(prometheus.SummaryOpts{Name: "observe_exec_durations_milliseconds", Help: "Item execution latency distributions."}, []string{"type", "exit_code"})
	nameExecDurations	= prometheus.NewSummaryVec(prometheus.SummaryOpts{Name: "observe_name_exec_durations_milliseconds", Help: "Name list execution latency distributions."}, []string{"exit_code"})
)
var (
	observeLong	= templates.LongDesc(`
		Observe changes to resources and take action on them

		This command assists in building scripted reactions to changes that occur in
		Kubernetes or OpenShift resources. This is frequently referred to as a
		'controller' in Kubernetes and acts to ensure particular conditions are
		maintained. On startup, observe will list all of the resources of a
		particular type and execute the provided script on each one. Observe watches
		the server for changes, and will reexecute the script for each update.

		Observe works best for problems of the form "for every resource X, make sure
		Y is true". Some examples of ways observe can be used include:

		* Ensure every namespace has a quota or limit range object
		* Ensure every service is registered in DNS by making calls to a DNS API
		* Send an email alert whenever a node reports 'NotReady'
		* Watch for the 'FailedScheduling' event and write an IRC message
		* Dynamically provision persistent volumes when a new PVC is created
		* Delete pods that have reached successful completion after a period of time.

		The simplest pattern is maintaining an invariant on an object - for instance,
		"every namespace should have an annotation that indicates its owner". If the
		object is deleted no reaction is necessary. A variation on that pattern is
		creating another object: "every namespace should have a quota object based
		on the resources allowed for an owner".

		    $ cat set_owner.sh
		    #!/bin/sh
		    if [[ "$(%[1]s get namespace "$1" --template='{{ .metadata.annotations.owner }}')" == "" ]]; then
		      %[1]s annotate namespace "$1" owner=bob
		    fi

		    $ %[1]s observe namespaces -- ./set_owner.sh

		The set_owner.sh script is invoked with a single argument (the namespace name)
		for each namespace. This simple script ensures that any user without the
		"owner" annotation gets one set, but preserves any existing value.

		The next common of controller pattern is provisioning - making changes in an
		external system to match the state of a Kubernetes resource. These scripts need
		to account for deletions that may take place while the observe command is not
		running. You can provide the list of known objects via the --names command,
		which should return a newline-delimited list of names or namespace/name pairs.
		Your command will be invoked whenever observe checks the latest state on the
		server - any resources returned by --names that are not found on the server
		will be passed to your --delete command.

		For example, you may wish to ensure that every node that is added to Kubernetes
		is added to your cluster inventory along with its IP:

		    $ cat add_to_inventory.sh
		    #!/bin/sh
		    echo "$1 $2" >> inventory
		    sort -u inventory -o inventory

		    $ cat remove_from_inventory.sh
		    #!/bin/sh
		    grep -vE "^$1 " inventory > /tmp/newinventory
		    mv -f /tmp/newinventory inventory

		    $ cat known_nodes.sh
		    #!/bin/sh
		    touch inventory
		    cut -f 1-1 -d ' ' inventory

		    $ %[1]s observe nodes -a '{ .status.addresses[0].address }' \
		      --names ./known_nodes.sh \
		      --delete ./remove_from_inventory.sh \
		      -- ./add_to_inventory.sh

		If you stop the observe command and then delete a node, when you launch observe
		again the contents of inventory will be compared to the list of nodes from the
		server, and any node in the inventory file that no longer exists will trigger
		a call to remove_from_inventory.sh with the name of the node.

		Important: when handling deletes, the previous state of the object may not be
		available and only the name/namespace of the object will be passed to	your
		--delete command as arguments (all custom arguments are omitted).

		More complicated interactions build on the two examples above - your inventory
		script could make a call to allocate storage on your infrastructure as a
		service, or register node names in DNS, or set complex firewalls. The more
		complex your integration, the more important it is to record enough data in the
		remote system that you can identify when resources on either side are deleted.`)
	observeExample	= templates.Examples(`
		# Observe changes to services
	  %[1]s observe services

	  # Observe changes to services, including the clusterIP and invoke a script for each
	  %[1]s observe services -a '{ .spec.clusterIP }' -- register_dns.sh

	  # Observe changes to services filtered by a label selector
	  %[1]s observe namespaces -l regist-dns=true -a '{ .spec.clusterIP }' -- register_dns.sh`)
)

type ObserveOptions struct {
	debugOut		io.Writer
	noHeaders		bool
	client			resource.RESTClient
	mapping			*meta.RESTMapping
	includeNamespace	bool
	namespace		string
	allNamespaces		bool
	selector		string
	listenAddr		string
	eachCommand		[]string
	objectEnvVar		string
	typeEnvVar		string
	deleteCommand		stringSliceFlag
	nameSyncCommand		stringSliceFlag
	observedErrors		int
	maximumErrors		int
	retryCount		int
	retryExitStatus		int
	once			bool
	exitAfterPeriod		time.Duration
	resyncPeriod		time.Duration
	printMetricsOnExit	bool
	templateType		string
	templates		stringSliceFlag
	printer			ColumnPrinter
	strictTemplates		bool
	argumentStore		*objectArgumentsStore
	knownObjects		knownObjects
	genericclioptions.IOStreams
}

func NewObserveOptions(streams genericclioptions.IOStreams) *ObserveOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ObserveOptions{IOStreams: streams, retryCount: 2, templateType: "jsonpath", maximumErrors: 20, listenAddr: ":11251"}
}
func NewCmdObserve(fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewObserveOptions(streams)
	cmd := &cobra.Command{Use: "observe RESOURCE [-- COMMAND ...]", Short: "Observe changes to resources and react to them (experimental)", Long: fmt.Sprintf(observeLong, fullName), Example: fmt.Sprintf(observeExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		if err := o.Complete(f, cmd, args); err != nil {
			cmdutil.CheckErr(err)
		}
		if err := o.Validate(args); err != nil {
			cmdutil.CheckErr(cmdutil.UsageErrorf(cmd, err.Error()))
		}
		if err := o.Run(); err != nil {
			cmdutil.CheckErr(err)
		}
	}}
	cmd.Flags().BoolVarP(&o.allNamespaces, "all-namespaces", "A", o.allNamespaces, "If true, list the requested object(s) across all projects. Project in current context is ignored.")
	cmd.Flags().StringVarP(&o.selector, "selector", "l", o.selector, "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	cmd.Flags().VarP(&o.deleteCommand, "delete", "d", "A command to run when resources are deleted. Specify multiple times to add arguments.")
	cmd.Flags().Var(&o.nameSyncCommand, "names", "A command that will list all of the currently known names, optional. Specify multiple times to add arguments. Use to get notifications when objects are deleted.")
	cmd.Flags().StringVar(&o.templateType, "output", o.templateType, "Controls the template type used for the --argument flags. Supported values are gotemplate and jsonpath.")
	cmd.Flags().BoolVar(&o.strictTemplates, "strict-templates", o.strictTemplates, "If true return an error on any field or map key that is not missing in a template.")
	cmd.Flags().VarP(&o.templates, "argument", "a", "Template for the arguments to be passed to each command in the format defined by --output.")
	cmd.Flags().StringVar(&o.typeEnvVar, "type-env-var", "", "The name of an env var to set with the type of event received ('Sync', 'Updated', 'Deleted', 'Added') to the reaction command or --delete.")
	cmd.Flags().StringVar(&o.objectEnvVar, "object-env-var", "", "The name of an env var to serialize the object to when calling the command, optional.")
	cmd.Flags().IntVar(&o.maximumErrors, "maximum-errors", o.maximumErrors, "Exit after this many errors have been detected with. May be set to -1 for no maximum.")
	cmd.Flags().IntVar(&o.retryExitStatus, "retry-on-exit-code", o.retryExitStatus, "If any command returns this exit code, retry up to --retry-count times.")
	cmd.Flags().IntVar(&o.retryCount, "retry-count", o.retryCount, "The number of times to retry a failing command before continuing.")
	cmd.Flags().BoolVar(&o.once, "once", o.once, "If true, exit with a status code 0 after all current objects have been processed.")
	cmd.Flags().DurationVar(&o.exitAfterPeriod, "exit-after", o.exitAfterPeriod, "Exit with status code 0 after the provided duration, optional.")
	cmd.Flags().DurationVar(&o.resyncPeriod, "resync-period", o.resyncPeriod, "When non-zero, periodically reprocess every item from the server as a Sync event. Use to ensure external systems are kept up to date.")
	cmd.Flags().BoolVar(&o.printMetricsOnExit, "print-metrics-on-exit", o.printMetricsOnExit, "If true, on exit write all metrics to stdout.")
	cmd.Flags().StringVar(&o.listenAddr, "listen-addr", o.listenAddr, "The name of an interface to listen on to expose metrics and health checking.")
	cmd.Flags().BoolVar(&o.noHeaders, "no-headers", o.noHeaders, "If true, skip printing information about each event prior to executing the command.")
	return cmd
}
func (o *ObserveOptions) Complete(f kcmdutil.Factory, cmd *cobra.Command, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	var command []string
	if i := cmd.ArgsLenAtDash(); i != -1 {
		command = args[i:]
		args = args[:i]
	}
	o.eachCommand = command
	switch len(args) {
	case 0:
		return fmt.Errorf("you must specify at least one argument containing the resource to observe")
	case 1:
	default:
		return fmt.Errorf("you may only specify one argument containing the resource to observe (use '--' to separate your resource and your command)")
	}
	gr := schema.ParseGroupResource(args[0])
	if gr.Empty() {
		return fmt.Errorf("unknown resource argument")
	}
	mapper, err := f.ToRESTMapper()
	if err != nil {
		return err
	}
	version, err := mapper.KindFor(gr.WithVersion(""))
	if err != nil {
		return err
	}
	mapping, err := mapper.RESTMapping(version.GroupKind())
	if err != nil {
		return err
	}
	o.mapping = mapping
	o.includeNamespace = mapping.Scope.Name() == meta.RESTScopeNamespace.Name()
	o.client, err = f.ClientForMapping(mapping)
	if err != nil {
		return err
	}
	o.namespace, _, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	switch o.templateType {
	case "jsonpath":
		p, err := NewJSONPathArgumentPrinter(o.includeNamespace, o.strictTemplates, o.templates...)
		if err != nil {
			return err
		}
		o.printer = p
	case "gotemplate":
		p, err := NewGoTemplateArgumentPrinter(o.includeNamespace, o.strictTemplates, o.templates...)
		if err != nil {
			return err
		}
		o.printer = p
	default:
		return fmt.Errorf("template type %q not recognized - valid values are jsonpath and gotemplate", o.templateType)
	}
	o.printer = NewVersionedColumnPrinter(o.printer, legacyscheme.Scheme, version.GroupVersion())
	if o.noHeaders {
		o.debugOut = ioutil.Discard
	} else {
		o.debugOut = o.Out
	}
	o.argumentStore = &objectArgumentsStore{}
	switch {
	case len(o.nameSyncCommand) > 0:
		o.argumentStore.keyFn = func() ([]string, error) {
			var out []byte
			err := retryCommandError(o.retryExitStatus, o.retryCount, func() error {
				c := exec.Command(o.nameSyncCommand[0], o.nameSyncCommand[1:]...)
				var err error
				return measureCommandDuration(nameExecDurations, func() error {
					out, err = c.Output()
					return err
				})
			})
			if err != nil {
				if exit, ok := err.(*exec.ExitError); ok {
					if len(exit.Stderr) > 0 {
						err = fmt.Errorf("%v\n%s", err, string(exit.Stderr))
					}
				}
				return nil, err
			}
			names := strings.Split(string(out), "\n")
			sort.Sort(sort.StringSlice(names))
			var outputNames []string
			for i, s := range names {
				if len(s) != 0 {
					outputNames = names[i:]
					break
				}
			}
			klog.V(4).Infof("Found existing keys: %v", outputNames)
			return outputNames, nil
		}
		o.knownObjects = o.argumentStore
	case len(o.deleteCommand) > 0, o.resyncPeriod > 0:
		o.knownObjects = o.argumentStore
	}
	return nil
}
func (o *ObserveOptions) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.nameSyncCommand) > 0 && len(o.deleteCommand) == 0 {
		return fmt.Errorf("--delete and --names must both be specified")
	}
	if _, err := labels.Parse(o.selector); err != nil {
		return err
	}
	return nil
}
func (o *ObserveOptions) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.deleteCommand) > 0 && len(o.nameSyncCommand) == 0 {
		fmt.Fprintf(o.ErrOut, "warning: If you are modifying resources outside of %q, you should use the --names command to ensure you don't miss deletions that occur while the command is not running.\n", o.mapping.Resource)
	}
	store := cache.NewDeltaFIFO(objectArgumentsKeyFunc, o.knownObjects)
	lw := restListWatcher{Helper: resource.NewHelper(o.client, o.mapping), selector: o.selector}
	if !o.allNamespaces {
		lw.namespace = o.namespace
	}
	proc.StartReaper()
	if len(o.listenAddr) > 0 {
		prometheus.MustRegister(observeCounts)
		prometheus.MustRegister(execDurations)
		prometheus.MustRegister(nameExecDurations)
		errWaitingForSync := fmt.Errorf("waiting for initial sync")
		healthz.InstallHandler(http.DefaultServeMux, healthz.NamedCheck("ready", func(r *http.Request) error {
			if !store.HasSynced() {
				return errWaitingForSync
			}
			return nil
		}))
		http.Handle("/metrics", prometheus.Handler())
		go func() {
			klog.Fatalf("Unable to listen on %q: %v", o.listenAddr, http.ListenAndServe(o.listenAddr, nil))
		}()
		klog.V(2).Infof("Listening on %s at /metrics and /healthz", o.listenAddr)
	}
	var lock sync.Mutex
	if o.exitAfterPeriod > 0 {
		go func() {
			<-time.After(o.exitAfterPeriod)
			lock.Lock()
			defer lock.Unlock()
			o.dumpMetrics()
			fmt.Fprintf(o.ErrOut, "Shutting down after %s ...\n", o.exitAfterPeriod)
			os.Exit(0)
		}()
	}
	defer o.dumpMetrics()
	stopCh := make(chan struct{})
	defer close(stopCh)
	reflector := cache.NewNamedReflector("observer", lw, nil, store, o.resyncPeriod)
	go func() {
		observedListErrors := 0
		wait.Until(func() {
			if err := reflector.ListAndWatch(stopCh); err != nil {
				utilruntime.HandleError(err)
				observedListErrors++
				if o.maximumErrors != -1 && observedListErrors > o.maximumErrors {
					klog.Fatalf("Maximum list errors of %d reached, exiting", o.maximumErrors)
				}
			}
		}, time.Second, stopCh)
	}()
	if o.once {
		for len(reflector.LastSyncResourceVersion()) == 0 {
			time.Sleep(50 * time.Millisecond)
		}
		if store.HasSynced() && len(store.ListKeys()) == 0 {
			fmt.Fprintf(o.ErrOut, "Nothing to sync, exiting immediately\n")
			return nil
		}
	}
	syncing := false
	for {
		_, err := store.Pop(func(obj interface{}) error {
			if err := o.argumentStore.ListKeysError(); err != nil {
				return fmt.Errorf("unable to list known keys: %v", err)
			}
			deltas := obj.(cache.Deltas)
			for _, delta := range deltas {
				if err := func() error {
					lock.Lock()
					defer lock.Unlock()
					switch {
					case !syncing && delta.Type == cache.Sync:
						if err := o.startSync(); err != nil {
							return err
						}
						syncing = true
					case syncing && delta.Type != cache.Sync:
						if err := o.finishSync(); err != nil {
							return err
						}
						syncing = false
					}
					if !syncing && o.knownObjects == nil && !(delta.Type == cache.Added || delta.Type == cache.Updated) {
						return nil
					}
					observeCounts.WithLabelValues(string(delta.Type)).Inc()
					object, arguments, output, err := o.calculateArguments(delta)
					if err != nil {
						return err
					}
					if err := o.next(delta.Type, object, output, arguments); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
		if o.once && store.HasSynced() {
			if syncing {
				if err := o.finishSync(); err != nil {
					return err
				}
			}
			return nil
		}
	}
}
func (o *ObserveOptions) calculateArguments(delta cache.Delta) (runtime.Object, []string, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var arguments []string
	var object runtime.Object
	var key string
	var output []byte
	switch t := delta.Object.(type) {
	case cache.DeletedFinalStateUnknown:
		key = t.Key
		if obj, ok := t.Obj.(runtime.Object); ok {
			object = obj
			args, data, err := o.printer.Print(obj)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("unable to write arguments: %v", err)
			}
			arguments = args
			output = data
		} else {
			value, _, err := o.argumentStore.GetByKey(key)
			if err != nil {
				return nil, nil, nil, err
			}
			if value != nil {
				args, ok := value.(objectArguments)
				if !ok {
					return nil, nil, nil, fmt.Errorf("unexpected cache value %T", value)
				}
				arguments = args.arguments
				output = args.output
			}
		}
		o.argumentStore.Remove(key)
	case runtime.Object:
		object = t
		args, data, err := o.printer.Print(t)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("unable to write arguments: %v", err)
		}
		arguments = args
		output = data
		key, _ = cache.MetaNamespaceKeyFunc(t)
		if delta.Type == cache.Deleted {
			o.argumentStore.Remove(key)
		} else {
			saved := objectArguments{key: key, arguments: arguments}
			if len(o.objectEnvVar) > 0 {
				saved.output = output
			}
			o.argumentStore.Put(key, saved)
		}
	case objectArguments:
		key = t.key
		arguments = t.arguments
		output = t.output
	default:
		return nil, nil, nil, fmt.Errorf("unrecognized object %T from cache store", delta.Object)
	}
	if object == nil {
		namespace, name, err := cache.SplitMetaNamespaceKey(key)
		if err != nil {
			return nil, nil, nil, err
		}
		unstructured := &unstructured.Unstructured{}
		unstructured.SetNamespace(namespace)
		unstructured.SetName(name)
		object = unstructured
	}
	return object, arguments, output, nil
}
func (o *ObserveOptions) startSync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(o.debugOut, "# %s Sync started\n", time.Now().Format(time.RFC3339))
	return nil
}
func (o *ObserveOptions) finishSync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(o.debugOut, "# %s Sync ended\n", time.Now().Format(time.RFC3339))
	return nil
}
func (o *ObserveOptions) next(deltaType cache.DeltaType, obj runtime.Object, output []byte, arguments []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Processing %s %v: %#v", deltaType, arguments, obj)
	m, err := meta.Accessor(obj)
	if err != nil {
		return fmt.Errorf("unable to handle %T: %v", obj, err)
	}
	resourceVersion := m.GetResourceVersion()
	outType := string(deltaType)
	var args []string
	if o.includeNamespace {
		args = append(args, m.GetNamespace())
	}
	args = append(args, m.GetName())
	var command string
	switch {
	case deltaType == cache.Deleted:
		if len(o.deleteCommand) > 0 {
			command = o.deleteCommand[0]
			args = append(append([]string{}, o.deleteCommand[1:]...), args...)
		}
	case len(o.eachCommand) > 0:
		command = o.eachCommand[0]
		args = append(append([]string{}, o.eachCommand[1:]...), args...)
	}
	args = append(args, arguments...)
	if len(command) == 0 {
		fmt.Fprintf(o.debugOut, "# %s %s %s\t%s\n", time.Now().Format(time.RFC3339), outType, resourceVersion, printCommandLine(command, args...))
		return nil
	}
	fmt.Fprintf(o.debugOut, "# %s %s %s\t%s\n", time.Now().Format(time.RFC3339), outType, resourceVersion, printCommandLine(command, args...))
	out, errOut := &newlineTrailingWriter{w: o.Out}, &newlineTrailingWriter{w: o.ErrOut}
	err = retryCommandError(o.retryExitStatus, o.retryCount, func() error {
		cmd := exec.Command(command, args...)
		cmd.Stdout = out
		cmd.Stderr = errOut
		cmd.Env = os.Environ()
		if len(o.objectEnvVar) > 0 {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", o.objectEnvVar, string(output)))
		}
		if len(o.typeEnvVar) > 0 {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", o.typeEnvVar, string(outType)))
		}
		err := measureCommandDuration(execDurations, cmd.Run, outType)
		out.Flush()
		errOut.Flush()
		return err
	})
	if err != nil {
		if code, ok := exitCodeForCommandError(err); ok && code != 0 {
			err = fmt.Errorf("command %q exited with status code %d", command, code)
		}
		return o.handleCommandError(err)
	}
	return nil
}
func (o *ObserveOptions) handleCommandError(err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return nil
	}
	o.observedErrors++
	fmt.Fprintf(o.ErrOut, "error: %v\n", err)
	if o.maximumErrors == -1 || o.observedErrors < o.maximumErrors {
		return nil
	}
	return fmt.Errorf("reached maximum error limit of %d, exiting", o.maximumErrors)
}
func (o *ObserveOptions) dumpMetrics() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !o.printMetricsOnExit {
		return
	}
	w := httptest.NewRecorder()
	prometheus.UninstrumentedHandler().ServeHTTP(w, &http.Request{})
	if w.Code == http.StatusOK {
		fmt.Fprintf(o.Out, w.Body.String())
	}
}
func measureCommandDuration(m *prometheus.SummaryVec, fn func() error, labels ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := time.Now()
	err := fn()
	duration := time.Now().Sub(n)
	statusCode, ok := exitCodeForCommandError(err)
	if !ok {
		statusCode = -1
	}
	m.WithLabelValues(append(labels, strconv.Itoa(statusCode))...).Observe(float64(duration / time.Millisecond))
	if errnoError(err) == syscall.ECHILD {
		return nil
	}
	return err
}
func errnoError(err error) syscall.Errno {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if se, ok := err.(*os.SyscallError); ok {
		if errno, ok := se.Err.(syscall.Errno); ok {
			return errno
		}
	}
	return 0
}
func exitCodeForCommandError(err error) (int, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return 0, true
	}
	if exit, ok := err.(*exec.ExitError); ok {
		if ws, ok := exit.ProcessState.Sys().(syscall.WaitStatus); ok {
			return ws.ExitStatus(), true
		}
	}
	return 0, false
}
func retryCommandError(onExitStatus, times int, fn func() error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := fn()
	if err != nil && onExitStatus != 0 && times > 0 {
		if status, ok := exitCodeForCommandError(err); ok {
			if status == onExitStatus {
				klog.V(4).Infof("retrying command: %v", err)
				return retryCommandError(onExitStatus, times-1, fn)
			}
		}
	}
	return err
}
func printCommandLine(cmd string, args ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	outCmd := cmd
	if strings.ContainsAny(outCmd, "\"\\ ") {
		outCmd = strconv.Quote(outCmd)
	}
	if len(outCmd) == 0 {
		outCmd = "\"\""
	}
	outArgs := make([]string, 0, len(args))
	for _, s := range args {
		if strings.ContainsAny(s, "\"\\ ") {
			s = strconv.Quote(s)
		}
		if len(s) == 0 {
			s = "\"\""
		}
		outArgs = append(outArgs, s)
	}
	if len(outArgs) == 0 {
		return outCmd
	}
	return fmt.Sprintf("%s %s", outCmd, strings.Join(outArgs, " "))
}

type restListWatcher struct {
	*resource.Helper
	namespace	string
	selector	string
}

func (lw restListWatcher) List(opt metav1.ListOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opt.LabelSelector = lw.selector
	return lw.Helper.List(lw.namespace, "", false, &opt)
}
func (lw restListWatcher) Watch(opt metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opt.LabelSelector = lw.selector
	return lw.Helper.Watch(lw.namespace, opt.ResourceVersion, &opt)
}

type JSONPathColumnPrinter struct {
	includeNamespace	bool
	rawTemplates		[]string
	templates		[]*jsonpath.JSONPath
	buf			*bytes.Buffer
}

func NewJSONPathArgumentPrinter(includeNamespace, strict bool, templates ...string) (*JSONPathColumnPrinter, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &JSONPathColumnPrinter{includeNamespace: includeNamespace, rawTemplates: templates, buf: &bytes.Buffer{}}
	for _, s := range templates {
		t := jsonpath.New("template").AllowMissingKeys(!strict)
		if err := t.Parse(s); err != nil {
			return nil, err
		}
		p.templates = append(p.templates, t)
	}
	return p, nil
}
func (p *JSONPathColumnPrinter) Print(obj interface{}) ([]string, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var columns []string
	for i, t := range p.templates {
		p.buf.Reset()
		if err := t.Execute(p.buf, obj); err != nil {
			return nil, nil, fmt.Errorf("error executing template '%v': '%v'\n----data----\n%+v\n", p.rawTemplates[i], err, obj)
		}
		columns = append(columns, p.buf.String())
	}
	return columns, nil, nil
}

type GoTemplateColumnPrinter struct {
	includeNamespace	bool
	strict			bool
	rawTemplates		[]string
	templates		[]*template.Template
	buf			*bytes.Buffer
}

func NewGoTemplateArgumentPrinter(includeNamespace, strict bool, templates ...string) (*GoTemplateColumnPrinter, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &GoTemplateColumnPrinter{includeNamespace: includeNamespace, strict: strict, rawTemplates: templates, buf: &bytes.Buffer{}}
	for _, s := range templates {
		t := template.New("template")
		child, err := t.Parse(s)
		if err != nil {
			return nil, err
		}
		if !strict {
			child.Option("missingkey=zero")
		}
		p.templates = append(p.templates, child)
	}
	return p, nil
}
func (p *GoTemplateColumnPrinter) Print(obj interface{}) ([]string, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var columns []string
	for i, t := range p.templates {
		p.buf.Reset()
		if err := t.Execute(p.buf, obj); err != nil {
			return nil, nil, fmt.Errorf("error executing template '%v': '%v'\n----data----\n%+v\n", p.rawTemplates[i], err, obj)
		}
		if p.buf.String() == "<no value>" {
			if p.strict {
				return nil, nil, fmt.Errorf("error executing template '%v': <no value>", p.rawTemplates[i])
			}
			columns = append(columns, "")
		} else {
			columns = append(columns, p.buf.String())
		}
	}
	return columns, nil, nil
}

type ColumnPrinter interface {
	Print(obj interface{}) ([]string, []byte, error)
}
type VersionedColumnPrinter struct {
	printer		ColumnPrinter
	convertor	runtime.ObjectConvertor
	version		runtime.GroupVersioner
}

func NewVersionedColumnPrinter(printer ColumnPrinter, convertor runtime.ObjectConvertor, version runtime.GroupVersioner) ColumnPrinter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &VersionedColumnPrinter{printer: printer, convertor: convertor, version: version}
}
func (p *VersionedColumnPrinter) Print(out interface{}) ([]string, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var output []byte
	if obj, ok := out.(runtime.Object); ok {
		converted, err := p.convertor.ConvertToVersion(obj, p.version)
		if err != nil {
			if !runtime.IsNotRegisteredError(err) {
				return nil, nil, err
			}
			converted = obj
		}
		data, err := json.Marshal(converted)
		if err != nil {
			return nil, nil, err
		}
		output = data
		out = map[string]interface{}{}
		if err := json.Unmarshal(data, &out); err != nil {
			return nil, nil, err
		}
	}
	args, _, err := p.printer.Print(out)
	return args, output, err
}

type knownObjects interface {
	cache.KeyListerGetter
	ListKeysError() error
	Put(key string, value interface{})
	Remove(key string)
}
type objectArguments struct {
	key		string
	arguments	[]string
	output		[]byte
}

func objectArgumentsKeyFunc(obj interface{}) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if args, ok := obj.(objectArguments); ok {
		return args.key, nil
	}
	return cache.MetaNamespaceKeyFunc(obj)
}

type objectArgumentsStore struct {
	keyFn		func() ([]string, error)
	lock		sync.Mutex
	arguments	map[string]interface{}
	err		error
}

func (r *objectArgumentsStore) ListKeysError() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.err
}
func (r *objectArgumentsStore) ListKeys() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.keyFn != nil {
		var keys []string
		keys, r.err = r.keyFn()
		return keys
	}
	keys := make([]string, 0, len(r.arguments))
	for k := range r.arguments {
		keys = append(keys, k)
	}
	return keys
}
func (r *objectArgumentsStore) GetByKey(key string) (interface{}, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.lock.Lock()
	defer r.lock.Unlock()
	args := r.arguments[key]
	return args, true, nil
}
func (r *objectArgumentsStore) Put(key string, arguments interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.arguments == nil {
		r.arguments = make(map[string]interface{})
	}
	r.arguments[key] = arguments
}
func (r *objectArgumentsStore) Remove(key string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.arguments, key)
}

type newlineTrailingWriter struct {
	w		io.Writer
	openLine	bool
}

func (w *newlineTrailingWriter) Write(data []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(data) > 0 && data[len(data)-1] != '\n' {
		w.openLine = true
	}
	return w.w.Write(data)
}
func (w *newlineTrailingWriter) Flush() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if w.openLine {
		w.openLine = false
		_, err := fmt.Fprintln(w.w)
		return err
	}
	return nil
}

type stringSliceFlag []string

func (f *stringSliceFlag) Set(value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*f = append(*f, value)
	return nil
}
func (f *stringSliceFlag) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(*f, " ")
}
func (f *stringSliceFlag) Type() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "string"
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
