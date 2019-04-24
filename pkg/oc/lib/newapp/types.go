package newapp

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"strings"
	"k8s.io/klog"
)

type Tester interface {
	Has(dir string) (string, bool, error)
}
type Strategy int

const (
	StrategyUnspecified	Strategy	= iota
	StrategySource
	StrategyDocker
	StrategyPipeline
)

func (s Strategy) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch s {
	case StrategyUnspecified:
		return ""
	case StrategySource:
		return "source"
	case StrategyDocker:
		return "Docker"
	case StrategyPipeline:
		return "pipeline"
	}
	klog.Error("unknown strategy")
	return ""
}
func (s Strategy) Type() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "strategy"
}
func (s *Strategy) Set(str string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch strings.ToLower(str) {
	case "":
		*s = StrategyUnspecified
	case "docker":
		*s = StrategyDocker
	case "pipeline":
		*s = StrategyPipeline
	case "source":
		*s = StrategySource
	default:
		return fmt.Errorf("invalid strategy: %s. Must be 'docker', 'pipeline' or 'source'.", str)
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
