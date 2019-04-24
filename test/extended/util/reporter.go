package util

import (
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters/stenographer"
	"github.com/onsi/ginkgo/types"
)

const maxDescriptionLength = 100

type SimpleReporter struct {
	stenographer	stenographer.Stenographer
	Output		io.Writer
}

func NewSimpleReporter() *SimpleReporter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SimpleReporter{Output: os.Stdout, stenographer: stenographer.New(!config.DefaultReporterConfig.NoColor, false)}
}
func (r *SimpleReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(r.Output, "=== SUITE %s (%d total specs, %d will run):\n", summary.SuiteDescription, summary.NumberOfTotalSpecs, summary.NumberOfSpecsThatWillBeRun)
}
func (r *SimpleReporter) BeforeSuiteDidRun(*types.SetupSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *SimpleReporter) SpecWillRun(spec *types.SpecSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.printRunLine(spec)
}
func (r *SimpleReporter) SpecDidComplete(spec *types.SpecSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.handleSpecFailure(spec)
	r.printStatusLine(spec)
}
func (r *SimpleReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *SimpleReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *SimpleReporter) handleSpecFailure(spec *types.SpecSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch spec.State {
	case types.SpecStateFailed:
		r.stenographer.AnnounceSpecFailed(spec, true, false)
	case types.SpecStatePanicked:
		r.stenographer.AnnounceSpecPanicked(spec, true, false)
	case types.SpecStateTimedOut:
		r.stenographer.AnnounceSpecTimedOut(spec, true, false)
	}
}
func (r *SimpleReporter) printStatusLine(spec *types.SpecSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runTime := ""
	if runTime = fmt.Sprintf(" (%v)", spec.RunTime); runTime == " (0)" {
		runTime = ""
	}
	fmt.Fprintf(r.Output, "%4s%-16s %s%s\n", " ", stateToString(spec.State), specDescription(spec), runTime)
}
func (r *SimpleReporter) printRunLine(spec *types.SpecSummary) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(r.Output, "=== RUN %s:\n", trimLocation(spec.ComponentCodeLocations[1]))
}
func specDescription(spec *types.SpecSummary) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name := ""
	for _, t := range spec.ComponentTexts[1:len(spec.ComponentTexts)] {
		name += strings.TrimSpace(t) + " "
	}
	if len(name) == 0 {
		name = fmt.Sprintf("FIXME: Spec without valid name (%s)", spec.ComponentTexts)
	}
	return short(strings.TrimSpace(name))
}
func short(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runes := []rune(s)
	if len(runes) > maxDescriptionLength {
		return string(runes[:maxDescriptionLength]) + " ..."
	}
	return s
}
func bold(v string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "\033[1m" + v + "\033[0m"
}
func red(v string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "\033[31m" + v + "\033[0m"
}
func magenta(v string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "\033[35m" + v + "\033[0m"
}
func stateToString(s types.SpecState) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch s {
	case types.SpecStatePassed:
		return bold("ok")
	case types.SpecStateSkipped:
		return magenta("skip")
	case types.SpecStateFailed:
		return red("fail")
	case types.SpecStateTimedOut:
		return red("timed")
	case types.SpecStatePanicked:
		return red("panic")
	case types.SpecStatePending:
		return magenta("pending")
	default:
		return bold(fmt.Sprintf("%v", s))
	}
}
func trimLocation(l types.CodeLocation) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delimiter := "/openshift/origin/"
	return fmt.Sprintf("%q", l.FileName[strings.LastIndex(l.FileName, delimiter)+len(delimiter):])
}
