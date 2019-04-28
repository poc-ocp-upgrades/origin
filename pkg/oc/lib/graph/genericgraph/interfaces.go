package genericgraph

import (
	"github.com/gonum/graph"
)

type Marker struct {
	Node		graph.Node
	RelatedNodes	[]graph.Node
	Severity	Severity
	Key		string
	Message		string
	Suggestion	Suggestion
}
type Severity string

const (
	InfoSeverity	Severity	= "info"
	WarningSeverity	Severity	= "warning"
	ErrorSeverity	Severity	= "error"
)

type Markers []Marker
type MarkerScanner func(g Graph, f Namer) []Marker

func (m Markers) BySeverity(severity Severity) []Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []Marker{}
	for i := range m {
		if m[i].Severity == severity {
			ret = append(ret, m[i])
		}
	}
	return ret
}
func (m Markers) FilterByNamespace(namespace string) Markers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	filtered := Markers{}
	for i := range m {
		markerNodes := []graph.Node{}
		markerNodes = append(markerNodes, m[i].Node)
		markerNodes = append(markerNodes, m[i].RelatedNodes...)
		hasCrossNamespaceLink := false
		for _, node := range markerNodes {
			if IsFromDifferentNamespace(namespace, node) {
				hasCrossNamespaceLink = true
				break
			}
		}
		if !hasCrossNamespaceLink {
			filtered = append(filtered, m[i])
		}
	}
	return filtered
}

type BySeverity []Marker

func (m BySeverity) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m BySeverity) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m BySeverity) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lhs := m[i]
	rhs := m[j]
	switch lhs.Severity {
	case ErrorSeverity:
		switch rhs.Severity {
		case ErrorSeverity:
			return false
		}
	case WarningSeverity:
		switch rhs.Severity {
		case ErrorSeverity, WarningSeverity:
			return false
		}
	case InfoSeverity:
		switch rhs.Severity {
		case ErrorSeverity, WarningSeverity, InfoSeverity:
			return false
		}
	}
	return true
}

type ByNodeID []Marker

func (m ByNodeID) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m ByNodeID) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m ByNodeID) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m[i].Node == nil {
		return true
	}
	if m[j].Node == nil {
		return false
	}
	return m[i].Node.ID() < m[j].Node.ID()
}

type ByKey []Marker

func (m ByKey) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m)
}
func (m ByKey) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m[i], m[j] = m[j], m[i]
}
func (m ByKey) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m[i].Key < m[j].Key
}

type Suggestion string

func (s Suggestion) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(s)
}
