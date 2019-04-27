package graph

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"github.com/gonum/graph/simple"
)

type GraphOptions struct {
	Packages	*PackageList
	Roots		[]string
	Excludes	[]string
	Filters		[]string
	BuildTags	[]string
}

func (o *GraphOptions) Complete() error {
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
	return nil
}
func (o *GraphOptions) Validate() error {
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
	if o.Packages == nil || len(o.Packages.Packages) == 0 {
		return errors.New("a list of Go Packages is required in order to build the graph")
	}
	return nil
}
func (o *GraphOptions) BuildGraph() (*MutableDirectedGraph, error) {
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
	g := NewMutableDirectedGraph(o.Roots)
	filteredPackages := []Package{}
	for _, pkg := range o.Packages.Packages {
		if isExcludedPath(pkg.ImportPath, o.Excludes) {
			continue
		}
		if !isValidPackagePath(pkg.ImportPath) {
			continue
		}
		n := &Node{Id: g.NewNodeID(), UniqueName: pkg.ImportPath}
		err := g.AddNode(n)
		if err != nil {
			return nil, err
		}
		filteredPackages = append(filteredPackages, pkg)
	}
	for _, pkg := range filteredPackages {
		from, exists := g.NodeByName(pkg.ImportPath)
		if !exists {
			return nil, fmt.Errorf("expected node for package %q was not found in graph", pkg.ImportPath)
		}
		for _, dependency := range append(pkg.Imports, pkg.TestImports...) {
			if isExcludedPath(dependency, o.Excludes) {
				continue
			}
			if !isValidPackagePath(dependency) {
				continue
			}
			to, exists := g.NodeByName(dependency)
			if !exists {
				klog.V(1).Infof("Skipping unvisited (missing) dependency %q, which is imported by package %q", dependency, pkg.ImportPath)
				continue
			}
			if g.HasEdgeFromTo(from, to) {
				continue
			}
			g.SetEdge(simple.Edge{F: from, T: to})
		}
	}
	if len(o.Filters) > 0 {
		var err error
		g, err = FilterPackages(g, o.Filters)
		if err != nil {
			return nil, err
		}
	}
	g.PruneOrphans()
	return g, nil
}

type GraphFlags struct {
	Filter		string
	Exclude		string
	Openshift	bool
	RepoImportPath	string
	Entrypoints	[]string
	BuildTags	[]string
}

func (o *GraphFlags) calculateRoots(excludes []string) ([]string, error) {
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
	packages, err := getPackageMetadata(o.Entrypoints, excludes, o.BuildTags)
	if err != nil {
		return nil, err
	}
	roots := []string{}
	for _, pkg := range packages.Packages {
		roots = append(roots, pkg.ImportPath)
	}
	return roots, nil
}
func (o *GraphFlags) ToOptions(out, errout io.Writer) (*GraphOptions, error) {
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
	opts := &GraphOptions{BuildTags: o.BuildTags}
	if len(o.RepoImportPath) == 0 {
		return nil, errors.New("the go-import path for the repository must be specified via --root")
	}
	if len(o.Entrypoints) == 0 {
		return nil, errors.New("at least one entrypoint path must be provided")
	}
	if o.Openshift && (len(o.Exclude) > 0 || len(o.Filter) > 0) {
		return nil, errors.New("--exclude or --filter are mutually exclusive with --openshift")
	}
	o.Entrypoints = ensureEntrypointPrefix(o.Entrypoints, o.RepoImportPath)
	if o.Openshift {
		opts.Excludes = getOpenShiftExcludes()
		filters, err := getOpenShiftFilters()
		if err != nil {
			return nil, err
		}
		opts.Filters = filters
	} else {
		if len(o.Exclude) > 0 {
			f, err := ioutil.ReadFile(o.Exclude)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(f, &opts.Excludes)
			if err != nil {
				return nil, err
			}
		}
		if len(o.Filter) > 0 {
			f, err := ioutil.ReadFile(o.Filter)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(f, &opts.Filters)
			if err != nil {
				return nil, err
			}
		}
	}
	packages, err := getPackageMetadata(ensureVendorEntrypoint(o.Entrypoints, o.RepoImportPath), opts.Excludes, opts.BuildTags)
	if err != nil {
		return nil, err
	}
	opts.Packages = packages
	roots, err := o.calculateRoots(opts.Excludes)
	if err != nil {
		return nil, err
	}
	opts.Roots = roots
	return opts, nil
}
func (o *GraphFlags) AddFlags(cmd *cobra.Command) {
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
	cmd.Flags().BoolVar(&o.Openshift, "openshift", o.Openshift, "generate and use OpenShift-specific lists of excluded packages and filters.")
	cmd.Flags().StringVar(&o.RepoImportPath, "root", o.RepoImportPath, "Go import-path of repository to analyze (e.g. github.com/openshift/origin)")
	cmd.Flags().StringSliceVar(&o.Entrypoints, "entry", o.Entrypoints, "filepaths for packages within the specified --root relative to the repo's import path (e.g. ./cmd/...). Paths ending in an ellipsis (...) are traversed recursively.")
	cmd.Flags().StringSliceVar(&o.BuildTags, "tag", o.BuildTags, "list of build tags (not matching target system) of files to include during Go package traversal.")
	cmd.Flags().StringVarP(&o.Exclude, "exclude", "e", "", "optional path to json file containing a list of import-paths of packages within the specified repository to recursively exclude.")
	cmd.Flags().StringVarP(&o.Filter, "filter", "c", "", "optional path to json file containing a list of import-paths of packages to collapse sub-packages into.")
}
func isExcludedPath(path string, excludes []string) bool {
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
	for _, exclude := range excludes {
		if strings.HasPrefix(path, exclude) {
			return true
		}
	}
	return false
}
func ensureEntrypointPrefix(entrypoints []string, prefix string) []string {
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
	for idx, entry := range entrypoints {
		if strings.HasPrefix(entry, prefix) {
			continue
		}
		entrypoints[idx] = path.Join(prefix, entry)
	}
	return entrypoints
}
func ensureVendorEntrypoint(entrypoints []string, prefix string) []string {
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
	hasVendor := false
	for _, entry := range entrypoints {
		if strings.HasSuffix(path.Clean(entry), "/vendor") {
			hasVendor = true
			break
		}
	}
	if !hasVendor {
		vendor := ensureEntrypointPrefix([]string{"vendor"}, prefix)
		vendor = ensureEntrypointEllipsis(vendor)
		entrypoints = append(entrypoints, vendor[0])
	}
	return entrypoints
}
func ensureEntrypointEllipsis(entypoints []string) []string {
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
	parsedRoots := []string{}
	for _, entry := range entypoints {
		if strings.HasSuffix(entry, "...") {
			parsedRoots = append(parsedRoots, entry)
			continue
		}
		slash := ""
		if string(entry[len(entry)-1]) != "/" {
			slash = "/"
		}
		entry = strings.Join([]string{entry, slash, "..."}, "")
		parsedRoots = append(parsedRoots, entry)
	}
	return parsedRoots
}
