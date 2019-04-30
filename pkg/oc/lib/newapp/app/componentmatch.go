package app

import (
	dockerv10 "github.com/openshift/api/image/docker10"
	imagev1 "github.com/openshift/api/image/v1"
	templatev1 "github.com/openshift/api/template/v1"
)

type ComponentMatch struct {
	Value		string
	Argument	string
	Name		string
	Description	string
	Score		float32
	Insecure	bool
	LocalOnly	bool
	NoTagsFound	bool
	Virtual		bool
	DockerImage	*dockerv10.DockerImage
	ImageStream	*imagev1.ImageStream
	ImageTag	string
	Template	*templatev1.Template
	Builder		bool
	GeneratorInput	GeneratorInput
	Meta		map[string]string
}

func (m *ComponentMatch) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Argument
}
func (m *ComponentMatch) IsImage() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Template == nil
}
func (m *ComponentMatch) IsTemplate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Template != nil
}
func (m *ComponentMatch) Exact() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Score == 0.0
}

type ComponentMatches []*ComponentMatch

func (m ComponentMatches) Exact() ComponentMatches {
	_logClusterCodePath()
	defer _logClusterCodePath()
	exact := ComponentMatches{}
	for _, match := range m {
		if match.Exact() {
			exact = append(exact, match)
		}
	}
	return exact
}
func (m ComponentMatches) Inexact() ComponentMatches {
	_logClusterCodePath()
	defer _logClusterCodePath()
	inexact := ComponentMatches{}
	for _, match := range m {
		if !match.Exact() {
			inexact = append(inexact, match)
		}
	}
	return inexact
}

type ScoredComponentMatches []*ComponentMatch

func (m ScoredComponentMatches) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m ScoredComponentMatches) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m ScoredComponentMatches) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m[i].Score < m[j].Score
}
func (m ScoredComponentMatches) Exact() []*ComponentMatch {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := []*ComponentMatch{}
	for _, match := range m {
		if match.Score == 0.0 {
			out = append(out, match)
		}
	}
	return out
}
