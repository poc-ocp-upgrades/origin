package migrate

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"time"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/client-go/discovery"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

type MigrateVisitFunc func(info *resource.Info) (Reporter, error)
type MigrateActionFunc func(info *resource.Info, reporter Reporter) error
type MigrateFilterFunc func(info *resource.Info) (bool, error)
type Reporter interface{ Changed() bool }
type ReporterBool bool

func (r ReporterBool) Changed() bool {
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
	return bool(r)
}
func AlwaysRequiresMigration(_ *resource.Info) (Reporter, error) {
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
	return ReporterBool(true), nil
}
func timeStampNow() string {
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
	return time.Now().Format("0102 15:04:05.000000")
}

type flusher interface{ Flush() error }
type syncer interface{ Sync() error }

var _ io.Writer = &syncedWriter{}

type syncedWriter struct {
	lock	sync.Mutex
	writer	io.Writer
}

func (w *syncedWriter) Write(p []byte) (int, error) {
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
	w.lock.Lock()
	n, err := w.write(p)
	w.lock.Unlock()
	return n, err
}
func (w *syncedWriter) write(p []byte) (int, error) {
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
	n, err := w.writer.Write(p)
	if f, ok := w.writer.(flusher); ok {
		f.Flush()
	}
	if s, ok := w.writer.(syncer); ok {
		s.Sync()
	}
	return n, err
}

type ResourceOptions struct {
	PrintFlags		*genericclioptions.PrintFlags
	Printer			printers.ResourcePrinter
	Unstructured		bool
	AllNamespaces		bool
	Include			[]string
	Filenames		[]string
	Confirm			bool
	Output			string
	FromKey			string
	ToKey			string
	OverlappingResources	[]sets.String
	DefaultExcludes		[]schema.GroupResource
	Builder			*resource.Builder
	SaveFn			MigrateActionFunc
	PrintFn			MigrateActionFunc
	FilterFn		MigrateFilterFunc
	DryRun			bool
	Summarize		bool
	Workers			int
	SyncOut			bool
	genericclioptions.IOStreams
}

func NewResourceOptions(streams genericclioptions.IOStreams) *ResourceOptions {
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
	return &ResourceOptions{PrintFlags: genericclioptions.NewPrintFlags("migrated").WithTypeSetter(scheme.Scheme), IOStreams: streams, AllNamespaces: true}
}
func (o *ResourceOptions) WithIncludes(include []string) *ResourceOptions {
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
	o.Include = include
	return o
}
func (o *ResourceOptions) WithExcludes(defaultExcludes []schema.GroupResource) *ResourceOptions {
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
	o.DefaultExcludes = defaultExcludes
	return o
}
func (o *ResourceOptions) WithOverlappingResources(resources []sets.String) *ResourceOptions {
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
	o.OverlappingResources = resources
	return o
}
func (o *ResourceOptions) WithUnstructured() *ResourceOptions {
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
	o.Unstructured = true
	return o
}
func (o *ResourceOptions) WithAllNamespaces() *ResourceOptions {
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
	o.AllNamespaces = true
	return o
}
func (o *ResourceOptions) Bind(c *cobra.Command) {
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
	c.Flags().StringSliceVar(&o.Include, "include", o.Include, "Resource types to migrate. Passing --filename will override this flag.")
	c.Flags().BoolVarP(&o.AllNamespaces, "all-namespaces", "A", o.AllNamespaces, "Migrate objects in all namespaces. Defaults to true.")
	c.Flags().BoolVar(&o.Confirm, "confirm", o.Confirm, "If true, all requested objects will be migrated. Defaults to false.")
	c.Flags().StringVar(&o.FromKey, "from-key", o.FromKey, "If specified, only migrate items with a key (namespace/name or name) greater than or equal to this value")
	c.Flags().StringVar(&o.ToKey, "to-key", o.ToKey, "If specified, only migrate items with a key (namespace/name or name) less than this value")
	o.PrintFlags.AddFlags(c)
	usage := "Filename, directory, or URL to docker-compose.yml file to use"
	kcmdutil.AddJsonFilenameFlag(c.Flags(), &o.Filenames, usage)
}
func (o *ResourceOptions) Complete(f kcmdutil.Factory, c *cobra.Command) error {
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
	o.Output = kcmdutil.GetFlagString(c, "output")
	if o.DryRun {
		o.PrintFlags.Complete("%s (dry run)")
	}
	var err error
	o.Printer, err = o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	switch {
	case len(o.Output) > 0:
		first := true
		o.PrintFn = func(info *resource.Info, _ Reporter) error {
			if o.Output == "yaml" && !first {
				fmt.Fprintln(o.Out, "---")
			}
			first = false
			return o.Printer.PrintObj(info.Object, o.Out)
		}
		o.DryRun = true
	case o.Confirm:
		o.DryRun = false
	default:
		o.DryRun = true
	}
	namespace, explicitNamespace, err := f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	allNamespaces := !explicitNamespace && o.AllNamespaces
	if len(o.FromKey) > 0 || len(o.ToKey) > 0 {
		o.FilterFn = func(info *resource.Info) (bool, error) {
			var key string
			if info.Mapping.Scope.Name() == meta.RESTScopeNameNamespace {
				key = info.Namespace + "/" + info.Name
			} else {
				if !allNamespaces {
					return false, nil
				}
				key = info.Name
			}
			if len(o.FromKey) > 0 && o.FromKey > key {
				return false, nil
			}
			if len(o.ToKey) > 0 && o.ToKey <= key {
				return false, nil
			}
			return true, nil
		}
	}
	discoveryClient, err := f.ToDiscoveryClient()
	if err != nil {
		return err
	}
	discoveryClient.Invalidate()
	_, _ = discoveryClient.ServerResources()
	mapper, err := f.ToRESTMapper()
	if err != nil {
		return err
	}
	resourceNames := sets.NewString()
	for i, s := range o.Include {
		if resourceNames.Has(s) {
			continue
		}
		if s != "*" {
			resourceNames.Insert(s)
			continue
		}
		exclude := sets.NewString()
		for _, gr := range o.DefaultExcludes {
			if len(o.OverlappingResources) > 0 {
				for _, others := range o.OverlappingResources {
					if !others.Has(gr.String()) {
						continue
					}
					exclude.Insert(others.List()...)
					break
				}
			}
			exclude.Insert(gr.String())
		}
		candidate := sets.NewString()
		all, err := FindAllCanonicalResources(discoveryClient, mapper)
		if err != nil {
			return fmt.Errorf("could not calculate the list of available resources: %v", err)
		}
		for _, gr := range all {
			if resourceNames.Has(gr.Resource) || resourceNames.Has(gr.String()) || exclude.Has(gr.String()) {
				continue
			}
			candidate.Insert(gr.String())
		}
		candidate.Delete(exclude.List()...)
		include := candidate
		if len(o.OverlappingResources) > 0 {
			include = sets.NewString()
			for _, k := range candidate.List() {
				reduce := k
				for _, others := range o.OverlappingResources {
					if !others.Has(k) {
						continue
					}
					reduce = others.List()[0]
					break
				}
				include.Insert(reduce)
			}
		}
		klog.V(4).Infof("Found the following resources from the server: %v", include.List())
		last := o.Include[i+1:]
		o.Include = append([]string{}, o.Include[:i]...)
		o.Include = append(o.Include, include.List()...)
		o.Include = append(o.Include, last...)
		break
	}
	if o.Workers == 0 {
		o.Workers = 1
	}
	if len(o.Output) > 0 && o.Workers > 1 {
		o.SyncOut = true
	}
	if o.SyncOut {
		o.Out = &syncedWriter{writer: o.Out}
		o.ErrOut = &syncedWriter{writer: o.ErrOut}
	}
	o.Builder = f.NewBuilder().AllNamespaces(allNamespaces).FilenameParam(false, &resource.FilenameOptions{Recursive: false, Filenames: o.Filenames}).ContinueOnError().DefaultNamespace().RequireObject(true).SelectAllParam(true).Flatten().RequestChunksOf(500)
	if o.Unstructured {
		o.Builder.Unstructured()
	} else {
		o.Builder.WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...)
	}
	if !allNamespaces {
		o.Builder.NamespaceParam(namespace)
	}
	if len(o.Filenames) == 0 {
		o.Builder.ResourceTypes(o.Include...)
	}
	return nil
}
func (o *ResourceOptions) Validate() error {
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
	if len(o.Filenames) == 0 && len(o.Include) == 0 {
		return fmt.Errorf("you must specify at least one resource or resource type to migrate with --include or --filenames")
	}
	if o.Workers < 1 {
		return fmt.Errorf("invalid value %d for workers, must be at least 1", o.Workers)
	}
	return nil
}
func (o *ResourceOptions) Visitor() *ResourceVisitor {
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
	return &ResourceVisitor{Out: o.Out, Builder: &resourceBuilder{builder: o.Builder}, SaveFn: o.SaveFn, PrintFn: o.PrintFn, FilterFn: o.FilterFn, DryRun: o.DryRun, Workers: o.Workers}
}

type Builder interface {
	Visitor(fns ...resource.ErrMatchFunc) (resource.Visitor, error)
}
type resourceBuilder struct{ builder *resource.Builder }

func (r *resourceBuilder) Visitor(fns ...resource.ErrMatchFunc) (resource.Visitor, error) {
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
	result := r.builder.Do().IgnoreErrors(fns...)
	return result, result.Err()
}

type ResourceVisitor struct {
	Out		io.Writer
	Builder		Builder
	SaveFn		MigrateActionFunc
	PrintFn		MigrateActionFunc
	FilterFn	MigrateFilterFunc
	DryRun		bool
	Workers		int
}

func (o *ResourceVisitor) Visit(fn MigrateVisitFunc) error {
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
	dryRun := o.DryRun
	summarize := true
	actionFn := o.SaveFn
	switch {
	case o.PrintFn != nil:
		actionFn = o.PrintFn
		dryRun = true
		summarize = false
	case dryRun:
		actionFn = nil
	}
	out := o.Out
	visitor, err := o.Builder.Visitor(errors.IsMethodNotSupported, errors.IsNotFound)
	if err != nil {
		return err
	}
	work := make(chan workData, 10*o.Workers)
	results := make(chan resultData, 10*o.Workers)
	t := &migrateTracker{out: out, dryRun: dryRun, resourcesWithErrors: sets.NewString(), results: results}
	workersWG := sync.WaitGroup{}
	for w := 0; w < o.Workers; w++ {
		workersWG.Add(1)
		go func() {
			defer workersWG.Done()
			worker := &migrateWorker{retries: 10, work: work, results: results, migrateFn: fn, actionFn: actionFn, filterFn: o.FilterFn}
			worker.run()
		}()
	}
	consumerWG := sync.WaitGroup{}
	consumerWG.Add(1)
	go func() {
		defer consumerWG.Done()
		t.run()
	}()
	err = visitor.Visit(func(info *resource.Info, err error) error {
		work <- workData{info: info, err: err}
		return nil
	})
	close(work)
	workersWG.Wait()
	close(results)
	consumerWG.Wait()
	if summarize {
		if dryRun {
			fmt.Fprintf(out, "summary (dry run): total=%d errors=%d ignored=%d unchanged=%d migrated=%d\n", t.found, t.errors, t.ignored, t.unchanged, t.found-t.errors-t.unchanged-t.ignored)
		} else {
			fmt.Fprintf(out, "summary: total=%d errors=%d ignored=%d unchanged=%d migrated=%d\n", t.found, t.errors, t.ignored, t.unchanged, t.found-t.errors-t.unchanged-t.ignored)
		}
	}
	if t.resourcesWithErrors.Len() > 0 {
		fmt.Fprintf(out, "info: to rerun only failing resources, add --include=%s\n", strings.Join(t.resourcesWithErrors.List(), ","))
	}
	switch {
	case err != nil:
		fmt.Fprintf(out, "error: exited without processing all resources: %v\n", err)
		err = kcmdutil.ErrExit
	case t.errors > 0:
		fmt.Fprintf(out, "error: %d resources failed to migrate\n", t.errors)
		err = kcmdutil.ErrExit
	}
	return err
}

var ErrUnchanged = fmt.Errorf("migration was not necessary")
var ErrRecalculate = fmt.Errorf("recalculate migration")

type MigrateError error
type ErrRetriable struct{ MigrateError }

func (ErrRetriable) Temporary() bool {
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
	return true
}

type ErrNotRetriable struct{ MigrateError }

func (ErrNotRetriable) Temporary() bool {
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
	return false
}

type TemporaryError interface {
	error
	Temporary() bool
}
type attemptResult int

const (
	attemptResultSuccess	attemptResult	= iota
	attemptResultError
	attemptResultUnchanged
	attemptResultIgnore
)

type workData struct {
	info	*resource.Info
	err	error
}
type resultData struct {
	found	bool
	retry	bool
	result	attemptResult
	data	workData
}
type migrateTracker struct {
	out					io.Writer
	dryRun					bool
	found, ignored, unchanged, errors	int
	resourcesWithErrors			sets.String
	results					<-chan resultData
}

func (t *migrateTracker) report(prefix string, info *resource.Info, err error) {
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
	ns := info.Namespace
	if len(ns) > 0 {
		ns = " -n " + ns
	}
	groupResource := info.Mapping.Resource.GroupResource()
	groupResourceStr := (&groupResource).String()
	if err != nil {
		fmt.Fprintf(t.out, "E%s %-10s%s %s/%s: %v\n", timeStampNow(), prefix, ns, groupResourceStr, info.Name, err)
	} else {
		fmt.Fprintf(t.out, "I%s %-10s%s %s/%s\n", timeStampNow(), prefix, ns, groupResourceStr, info.Name)
	}
}
func (t *migrateTracker) run() {
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
	for r := range t.results {
		if r.found {
			t.found++
		}
		if r.retry {
			t.report("retry:", r.data.info, r.data.err)
			continue
		}
		switch r.result {
		case attemptResultError:
			t.report("error:", r.data.info, r.data.err)
			t.errors++
			groupResource := r.data.info.Mapping.Resource.GroupResource()
			t.resourcesWithErrors.Insert((&groupResource).String())
		case attemptResultIgnore:
			t.ignored++
			if klog.V(2) {
				t.report("ignored:", r.data.info, nil)
			}
		case attemptResultUnchanged:
			t.unchanged++
			if klog.V(2) {
				t.report("unchanged:", r.data.info, nil)
			}
		case attemptResultSuccess:
			if klog.V(1) {
				if t.dryRun {
					t.report("migrated (dry run):", r.data.info, nil)
				} else {
					t.report("migrated:", r.data.info, nil)
				}
			}
		}
	}
}

type migrateWorker struct {
	retries		int
	work		<-chan workData
	results		chan<- resultData
	migrateFn	MigrateVisitFunc
	actionFn	MigrateActionFunc
	filterFn	MigrateFilterFunc
}

func (t *migrateWorker) run() {
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
	for data := range t.work {
		if data.err == nil && t.filterFn != nil {
			ok, err := t.filterFn(data.info)
			if err != nil {
				t.results <- resultData{found: true, result: attemptResultError, data: workData{info: data.info, err: err}}
				continue
			}
			if !ok {
				t.results <- resultData{found: true, result: attemptResultIgnore, data: data}
				continue
			}
		}
		if data.err != nil {
			t.results <- resultData{result: attemptResultError, data: data}
			continue
		}
		result, err := t.try(data.info, t.retries)
		t.results <- resultData{found: true, result: result, data: workData{info: data.info, err: err}}
	}
}
func (t *migrateWorker) try(info *resource.Info, retries int) (attemptResult, error) {
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
	reporter, err := t.migrateFn(info)
	if err != nil {
		return attemptResultError, err
	}
	if reporter == nil {
		return attemptResultIgnore, nil
	}
	if !reporter.Changed() {
		return attemptResultUnchanged, nil
	}
	if t.actionFn != nil {
		if err := t.actionFn(info, reporter); err != nil {
			if err == ErrUnchanged {
				return attemptResultUnchanged, nil
			}
			if canRetry(err) {
				if retries > 0 {
					if bool(klog.V(1)) && err != ErrRecalculate {
						t.results <- resultData{retry: true, data: workData{info: info, err: err}}
					}
					result, err := t.try(info, retries-1)
					switch result {
					case attemptResultUnchanged, attemptResultIgnore:
						result = attemptResultSuccess
					}
					return result, err
				}
			}
			return attemptResultError, err
		}
	}
	return attemptResultSuccess, nil
}
func canRetry(err error) bool {
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
	if temp, ok := err.(TemporaryError); ok && temp.Temporary() {
		return true
	}
	return err == ErrRecalculate
}
func DefaultRetriable(info *resource.Info, err error) error {
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
	switch {
	case err == nil:
		return nil
	case errors.IsNotFound(err):
		return ErrUnchanged
	case errors.IsMethodNotSupported(err):
		return ErrNotRetriable{err}
	case errors.IsConflict(err):
		if refreshErr := info.Get(); refreshErr != nil {
			if errors.IsNotFound(refreshErr) {
				return ErrUnchanged
			}
			return ErrNotRetriable{err}
		}
		return ErrRetriable{err}
	case errors.IsServerTimeout(err):
		return ErrRetriable{err}
	default:
		return err
	}
}
func FindAllCanonicalResources(d discovery.ServerResourcesInterface, m meta.RESTMapper) ([]schema.GroupResource, error) {
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
	set := make(map[schema.GroupResource]struct{})
	all, err := d.ServerResources()
	if err != nil {
		return nil, err
	}
	for _, serverResource := range all {
		gv, err := schema.ParseGroupVersion(serverResource.GroupVersion)
		if err != nil {
			continue
		}
		for _, r := range serverResource.APIResources {
			if strings.Contains(r.Name, "/") {
				continue
			}
			if !sets.NewString(r.Verbs...).HasAll("list", "update") {
				continue
			}
			if mapping, err := m.RESTMapping(schema.GroupKind{Group: gv.Group, Kind: r.Kind}, gv.Version); err == nil {
				if _, err := m.KindsFor(mapping.Resource); err == nil {
					set[mapping.Resource.GroupResource()] = struct{}{}
				}
			}
		}
	}
	var groupResources []schema.GroupResource
	for k := range set {
		groupResources = append(groupResources, k)
	}
	sort.Sort(groupResourcesByName(groupResources))
	return groupResources, nil
}

type groupResourcesByName []schema.GroupResource

func (g groupResourcesByName) Len() int {
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
	return len(g)
}
func (g groupResourcesByName) Less(i, j int) bool {
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
	if g[i].Resource < g[j].Resource {
		return true
	}
	if g[i].Resource > g[j].Resource {
		return false
	}
	return g[i].Group < g[j].Group
}
func (g groupResourcesByName) Swap(i, j int) {
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
	g[i], g[j] = g[j], g[i]
}
