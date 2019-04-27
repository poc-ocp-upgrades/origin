package app

import (
	"fmt"
	"sort"
	"strings"
	"github.com/openshift/origin/pkg/oc/lib/newapp"
	"k8s.io/apimachinery/pkg/util/errors"
)

func IsComponentReference(s string) error {
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
	if len(s) == 0 {
		return fmt.Errorf("empty string provided to component reference check")
	}
	all := strings.Split(s, "+")
	_, _, _, err := componentWithSource(all[0])
	return err
}
func componentWithSource(s string) (component, repo string, builder bool, err error) {
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
	if strings.Contains(s, "~") {
		segs := strings.SplitN(s, "~", 2)
		if len(segs) == 2 {
			builder = true
			switch {
			case len(segs[0]) == 0:
				err = fmt.Errorf("when using '[image]~[code]' form for %q, you must specify a image name", s)
				return
			case len(segs[1]) == 0:
				component = segs[0]
			default:
				component = segs[0]
				repo = segs[1]
			}
		}
	} else {
		component = s
	}
	return
}

type ComponentReference interface {
	Input() *ComponentInput
	Resolve() error
	Search() error
	NeedsSource() bool
}
type ComponentReferences []ComponentReference

func (r ComponentReferences) filter(filterFunc func(ref ComponentReference) bool) ComponentReferences {
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
	refs := ComponentReferences{}
	for _, ref := range r {
		if filterFunc(ref) {
			refs = append(refs, ref)
		}
	}
	return refs
}
func (r ComponentReferences) HasSource() bool {
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
	return len(r.filter(func(ref ComponentReference) bool {
		return ref.Input().Uses != nil
	})) > 0
}
func (r ComponentReferences) NeedsSource() (refs ComponentReferences) {
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
	return r.filter(func(ref ComponentReference) bool {
		return ref.NeedsSource()
	})
}
func (r ComponentReferences) UseSource() (refs ComponentReferences) {
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
	return r.filter(func(ref ComponentReference) bool {
		return ref.Input().Uses != nil
	})
}
func (r ComponentReferences) ImageComponentRefs() (refs ComponentReferences) {
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
	return r.filter(func(ref ComponentReference) bool {
		if ref.Input().ScratchImage {
			return true
		}
		return ref.Input() != nil && ref.Input().ResolvedMatch != nil && ref.Input().ResolvedMatch.IsImage()
	})
}
func (r ComponentReferences) TemplateComponentRefs() (refs ComponentReferences) {
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
	return r.filter(func(ref ComponentReference) bool {
		return ref.Input() != nil && ref.Input().ResolvedMatch != nil && ref.Input().ResolvedMatch.IsTemplate()
	})
}
func (r ComponentReferences) InstallableComponentRefs() (refs ComponentReferences) {
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
	return r.filter(func(ref ComponentReference) bool {
		return ref.Input() != nil && ref.Input().ResolvedMatch != nil && ref.Input().ResolvedMatch.GeneratorInput.Job
	})
}
func (r ComponentReferences) String() string {
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
	return r.HumanString(",")
}
func (r ComponentReferences) HumanString(separator string) string {
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
	components := []string{}
	for _, ref := range r {
		components = append(components, ref.Input().Value)
	}
	return strings.Join(components, separator)
}

type GroupedComponentReferences ComponentReferences

func (m GroupedComponentReferences) Len() int {
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
	return len(m)
}
func (m GroupedComponentReferences) Swap(i, j int) {
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
	m[i], m[j] = m[j], m[i]
}
func (m GroupedComponentReferences) Less(i, j int) bool {
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
	return m[i].Input().GroupID < m[j].Input().GroupID
}
func (r ComponentReferences) Group() (refs []ComponentReferences) {
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
	sorted := make(GroupedComponentReferences, len(r))
	copy(sorted, r)
	sort.Sort(sorted)
	groupID := -1
	for _, ref := range sorted {
		if ref.Input().GroupID != groupID {
			refs = append(refs, ComponentReferences{})
		}
		groupID = ref.Input().GroupID
		refs[len(refs)-1] = append(refs[len(refs)-1], ref)
	}
	return
}
func (components ComponentReferences) Resolve() error {
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
	for _, ref := range components {
		if err := ref.Resolve(); err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return errors.NewAggregate(errs)
}
func (components ComponentReferences) Search() error {
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
	for _, ref := range components {
		if err := ref.Search(); err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return errors.NewAggregate(errs)
}

type GeneratorJobReference struct {
	Ref	ComponentReference
	Input	GeneratorInput
	Err	error
}
type ReferenceBuilder struct {
	refs	ComponentReferences
	repos	SourceRepositories
	errs	[]error
	groupID	int
}

func (r *ReferenceBuilder) AddComponents(inputs []string, fn func(*ComponentInput) ComponentReference) ComponentReferences {
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
	refs := ComponentReferences{}
	for _, s := range inputs {
		for _, s := range strings.Split(s, "+") {
			input, repo, err := NewComponentInput(s)
			if err != nil {
				r.errs = append(r.errs, err)
				continue
			}
			input.GroupID = r.groupID
			ref := fn(input)
			if len(repo) != 0 {
				repository, ok := r.AddSourceRepository(repo, newapp.StrategySource)
				if !ok {
					continue
				}
				input.Use(repository)
				repository.UsedBy(ref)
			}
			refs = append(refs, ref)
		}
		r.groupID++
	}
	r.refs = append(r.refs, refs...)
	return refs
}
func (r *ReferenceBuilder) AddGroups(inputs []string) {
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
	for _, s := range inputs {
		groups := strings.Split(s, "+")
		if len(groups) == 1 {
			r.errs = append(r.errs, fmt.Errorf("group %q only contains a single name", s))
			continue
		}
		to := -1
		for _, group := range groups {
			var match ComponentReference
			for _, ref := range r.refs {
				if group == ref.Input().Value {
					match = ref
					break
				}
			}
			if match == nil {
				r.errs = append(r.errs, fmt.Errorf("the name %q from the group definition is not in use, and can't be used", group))
				break
			}
			if to == -1 {
				to = match.Input().GroupID
			} else {
				match.Input().GroupID = to
			}
		}
	}
}
func (r *ReferenceBuilder) AddSourceRepository(input string, strategy newapp.Strategy) (*SourceRepository, bool) {
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
	for _, existing := range r.repos {
		if input == existing.location {
			return existing, true
		}
	}
	source, err := NewSourceRepository(input, strategy)
	if err != nil {
		r.errs = append(r.errs, err)
		return nil, false
	}
	r.repos = append(r.repos, source)
	return source, true
}
func (r *ReferenceBuilder) AddExistingSourceRepository(source *SourceRepository) {
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
	r.repos = append(r.repos, source)
}
func (r *ReferenceBuilder) Result() (ComponentReferences, SourceRepositories, []error) {
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
	return r.refs, r.repos, r.errs
}
func NewComponentInput(input string) (*ComponentInput, string, error) {
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
	component, repo, builder, err := componentWithSource(input)
	if err != nil {
		return nil, "", err
	}
	return &ComponentInput{From: input, Argument: input, Value: component, ExpectToBuild: builder}, repo, nil
}

type ComponentInput struct {
	GroupID		int
	From		string
	Argument	string
	Value		string
	ExpectToBuild	bool
	ScratchImage	bool
	Uses		*SourceRepository
	ResolvedMatch	*ComponentMatch
	SearchMatches	ComponentMatches
	Resolver
	Searcher
}

func (i *ComponentInput) Input() *ComponentInput {
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
	return i
}
func (i *ComponentInput) NeedsSource() bool {
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
	return i.ExpectToBuild && i.Uses == nil
}
func (i *ComponentInput) Resolve() error {
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
	if i.Resolver == nil {
		return ErrNoMatch{Value: i.Value, Qualifier: "no resolver defined"}
	}
	match, err := i.Resolver.Resolve(i.Value)
	if err != nil {
		return err
	}
	i.Value = match.Value
	i.Argument = match.Argument
	i.ResolvedMatch = match
	return nil
}
func (i *ComponentInput) Search() error {
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
	if i.Searcher == nil {
		return ErrNoMatch{Value: i.Value, Qualifier: "no searcher defined"}
	}
	matches, err := i.Searcher.Search(false, i.Value)
	if matches != nil {
		i.SearchMatches = matches
	}
	return errors.NewAggregate(err)
}
func (i *ComponentInput) String() string {
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
	return i.Value
}
func (i *ComponentInput) Use(repo *SourceRepository) {
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
	i.Uses = repo
}
