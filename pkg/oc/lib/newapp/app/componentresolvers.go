package app

import (
	"sort"
	"strings"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/util/errors"
)

type Resolver interface {
	Resolve(value string) (*ComponentMatch, error)
}
type Searcher interface {
	Search(precise bool, terms ...string) (ComponentMatches, []error)
	Type() string
}
type WeightedResolver struct {
	Searcher
	Weight	float32
}
type PerfectMatchWeightedResolver []WeightedResolver

func (r PerfectMatchWeightedResolver) Resolve(value string) (*ComponentMatch, error) {
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
	var errs []error
	types := []string{}
	candidates := ScoredComponentMatches{}
	var group MultiSimpleSearcher
	var groupWeight float32 = 0.0
	for i, resolver := range r {
		if len(group) == 0 || resolver.Weight == groupWeight {
			group = append(group, resolver.Searcher)
			groupWeight = resolver.Weight
			if i != len(r)-1 && r[i+1].Weight == groupWeight {
				continue
			}
		}
		matches, err := group.Search(true, value)
		if err != nil {
			klog.V(5).Infof("Error from resolver: %v\n", err)
			errs = append(errs, err...)
		}
		types = append(types, group.Type())
		sort.Sort(ScoredComponentMatches(matches))
		if len(matches) > 0 && matches[0].Score == 0.0 && (len(matches) == 1 || matches[1].Score != 0.0) {
			return matches[0], errors.NewAggregate(errs)
		}
		for _, m := range matches {
			if groupWeight != 0.0 {
				m.Score = groupWeight * m.Score
			}
			candidates = append(candidates, m)
		}
		group = nil
	}
	switch len(candidates) {
	case 0:
		return nil, ErrNoMatch{Value: value, Errs: errs, Type: strings.Join(types, ", ")}
	case 1:
		if candidates[0].Score != 0.0 {
			if candidates[0].NoTagsFound {
				return nil, ErrNoTagsFound{Value: value, Match: candidates[0], Errs: errs}
			}
			return nil, ErrPartialMatch{Value: value, Match: candidates[0], Errs: errs}
		}
		return candidates[0], errors.NewAggregate(errs)
	default:
		sort.Sort(candidates)
		if candidates[0].Score < candidates[1].Score {
			if candidates[0].Score != 0.0 {
				return nil, ErrPartialMatch{Value: value, Match: candidates[0], Errs: errs}
			}
			return candidates[0], errors.NewAggregate(errs)
		}
		return nil, ErrMultipleMatches{Value: value, Matches: candidates, Errs: errs}
	}
}

type FirstMatchResolver struct{ Searcher Searcher }

func (r FirstMatchResolver) Resolve(value string) (*ComponentMatch, error) {
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
	matches, err := r.Searcher.Search(true, value)
	if len(matches) == 0 {
		return nil, ErrNoMatch{Value: value, Errs: err, Type: r.Searcher.Type()}
	}
	return matches[0], errors.NewAggregate(err)
}

type HighestScoreResolver struct{ Searcher Searcher }

func (r HighestScoreResolver) Resolve(value string) (*ComponentMatch, error) {
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
	matches, err := r.Searcher.Search(true, value)
	if len(matches) == 0 {
		return nil, ErrNoMatch{Value: value, Errs: err, Type: r.Searcher.Type()}
	}
	sort.Sort(ScoredComponentMatches(matches))
	return matches[0], errors.NewAggregate(err)
}

type HighestUniqueScoreResolver struct{ Searcher Searcher }

func (r HighestUniqueScoreResolver) Resolve(value string) (*ComponentMatch, error) {
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
	matches, err := r.Searcher.Search(true, value)
	sort.Sort(ScoredComponentMatches(matches))
	switch len(matches) {
	case 0:
		return nil, ErrNoMatch{Value: value, Errs: err, Type: r.Searcher.Type()}
	case 1:
		return matches[0], errors.NewAggregate(err)
	default:
		if matches[0].Score == matches[1].Score {
			equal := ComponentMatches{}
			for _, m := range matches {
				if m.Score != matches[0].Score {
					break
				}
				equal = append(equal, m)
			}
			return nil, ErrMultipleMatches{Value: value, Matches: equal, Errs: err}
		}
		return matches[0], errors.NewAggregate(err)
	}
}

type UniqueExactOrInexactMatchResolver struct{ Searcher Searcher }

func (r UniqueExactOrInexactMatchResolver) Resolve(value string) (*ComponentMatch, error) {
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
	matches, err := r.Searcher.Search(true, value)
	sort.Sort(ScoredComponentMatches(matches))
	exact := matches.Exact()
	switch len(exact) {
	case 0:
		inexact := matches.Inexact()
		switch len(inexact) {
		case 0:
			return nil, ErrNoMatch{Value: value, Errs: err, Type: r.Searcher.Type()}
		case 1:
			return inexact[0], errors.NewAggregate(err)
		default:
			return nil, ErrMultipleMatches{Value: value, Matches: matches, Errs: err}
		}
	case 1:
		return exact[0], errors.NewAggregate(err)
	default:
		return nil, ErrMultipleMatches{Value: value, Matches: matches, Errs: err}
	}
}

type PipelineResolver struct{}

func (r PipelineResolver) Resolve(value string) (*ComponentMatch, error) {
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
	return &ComponentMatch{Value: value, LocalOnly: true}, nil
}

type MultiSimpleSearcher []Searcher

func (s MultiSimpleSearcher) Type() string {
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
	t := []string{}
	for _, searcher := range s {
		t = append(t, searcher.Type())
	}
	return strings.Join(t, ", ")
}
func (s MultiSimpleSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
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
	var errs []error
	componentMatches := ComponentMatches{}
	for _, searcher := range s {
		matches, err := searcher.Search(precise, terms...)
		if len(err) > 0 {
			errs = append(errs, err...)
			continue
		}
		componentMatches = append(componentMatches, matches...)
	}
	sort.Sort(ScoredComponentMatches(componentMatches))
	return componentMatches, errs
}

type WeightedSearcher struct {
	Searcher
	Weight	float32
}
type MultiWeightedSearcher []WeightedSearcher

func (s MultiWeightedSearcher) Type() string {
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
	t := []string{}
	for _, searcher := range s {
		t = append(t, searcher.Type())
	}
	return strings.Join(t, ", ")
}
func (s MultiWeightedSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
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
	componentMatches := ComponentMatches{}
	var errs []error
	for _, searcher := range s {
		matches, err := searcher.Search(precise, terms...)
		if len(err) > 0 {
			errs = append(errs, err...)
			continue
		}
		for _, match := range matches {
			match.Score += searcher.Weight
			componentMatches = append(componentMatches, match)
		}
	}
	sort.Sort(ScoredComponentMatches(componentMatches))
	return componentMatches, errs
}
