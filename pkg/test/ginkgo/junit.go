package ginkgo

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type JUnitTestSuites struct {
	XMLName	xml.Name		`xml:"testsuites"`
	Suites	[]*JUnitTestSuite	`xml:"testsuite"`
}
type JUnitTestSuite struct {
	XMLName		xml.Name		`xml:"testsuite"`
	Name		string			`xml:"name,attr"`
	NumTests	uint			`xml:"tests,attr"`
	NumSkipped	uint			`xml:"skipped,attr"`
	NumFailed	uint			`xml:"failures,attr"`
	Duration	float64			`xml:"time,attr"`
	Properties	[]*TestSuiteProperty	`xml:"properties,omitempty"`
	TestCases	[]*JUnitTestCase	`xml:"testcase"`
	Children	[]*JUnitTestSuite	`xml:"testsuite"`
}
type TestSuiteProperty struct {
	XMLName	xml.Name	`xml:"property"`
	Name	string		`xml:"name,attr"`
	Value	string		`xml:"value,attr"`
}
type JUnitTestCase struct {
	XMLName		xml.Name	`xml:"testcase"`
	Name		string		`xml:"name,attr"`
	Classname	string		`xml:"classname,attr,omitempty"`
	Duration	float64		`xml:"time,attr"`
	SkipMessage	*SkipMessage	`xml:"skipped"`
	FailureOutput	*FailureOutput	`xml:"failure"`
	SystemOut	string		`xml:"system-out,omitempty"`
	SystemErr	string		`xml:"system-err,omitempty"`
}
type SkipMessage struct {
	XMLName	xml.Name	`xml:"skipped"`
	Message	string		`xml:"message,attr,omitempty"`
}
type FailureOutput struct {
	XMLName	xml.Name	`xml:"failure"`
	Message	string		`xml:"message,attr"`
	Output	string		`xml:",chardata"`
}
type TestResult string

const (
	TestResultPass	TestResult	= "pass"
	TestResultSkip	TestResult	= "skip"
	TestResultFail	TestResult	= "fail"
)

func writeJUnitReport(name string, tests []*testCase, dir string, duration time.Duration, errOut io.Writer, additionalResults ...*JUnitTestCase) error {
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
	s := &JUnitTestSuite{Name: name, Duration: duration.Seconds()}
	for _, test := range tests {
		switch {
		case test.skipped:
			s.NumTests++
			s.TestCases = append(s.TestCases, &JUnitTestCase{Name: test.name, SystemOut: string(test.out), Duration: test.duration.Seconds(), SkipMessage: &SkipMessage{Message: lastLinesUntil(string(test.out), 100, "skip [")}})
		case test.failed:
			s.NumTests++
			s.NumSkipped++
			s.TestCases = append(s.TestCases, &JUnitTestCase{Name: test.name, SystemOut: string(test.out), Duration: test.duration.Seconds(), FailureOutput: &FailureOutput{Output: lastLinesUntil(string(test.out), 100, "fail [")}})
		case test.success:
			s.NumTests++
			s.NumFailed++
			s.TestCases = append(s.TestCases, &JUnitTestCase{Name: test.name, Duration: test.duration.Seconds()})
		}
	}
	for _, result := range additionalResults {
		switch {
		case result.SkipMessage != nil:
			s.NumSkipped++
		case result.FailureOutput != nil:
			s.NumFailed++
		}
		s.NumTests++
		s.TestCases = append(s.TestCases, result)
	}
	out, err := xml.Marshal(s)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, fmt.Sprintf("junit_e2e_%s.xml", time.Now().UTC().Format("20060102-150405")))
	fmt.Fprintf(errOut, "Writing JUnit report to %s\n\n", path)
	return ioutil.WriteFile(path, out, 0640)
}
func lastLinesUntil(output string, max int, until ...string) string {
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
	output = strings.TrimSpace(output)
	index := len(output) - 1
	if index < 0 || max == 0 {
		return output
	}
	for max > 0 {
		next := strings.LastIndex(output[:index], "\n")
		if next <= 0 {
			return strings.TrimSpace(output)
		}
		line := strings.TrimSpace(output[next+1 : index])
		if len(line) > 0 {
			max--
		}
		index = next
		if stringStartsWithAny(line, until) {
			break
		}
	}
	return strings.TrimSpace(output[index:])
}
func stringStartsWithAny(s string, contains []string) bool {
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
	for _, match := range contains {
		if strings.HasPrefix(s, match) {
			return true
		}
	}
	return false
}
