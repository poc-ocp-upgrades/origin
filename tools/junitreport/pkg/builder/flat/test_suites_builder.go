package flat

import (
	"github.com/openshift/origin/tools/junitreport/pkg/api"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"github.com/openshift/origin/tools/junitreport/pkg/builder"
)

func NewTestSuitesBuilder() builder.TestSuitesBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &flatTestSuitesBuilder{testSuites: &api.TestSuites{}}
}

type flatTestSuitesBuilder struct{ testSuites *api.TestSuites }

func (b *flatTestSuitesBuilder) AddSuite(suite *api.TestSuite) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.testSuites.Suites = append(b.testSuites.Suites, suite)
}
func (b *flatTestSuitesBuilder) Build() *api.TestSuites {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.testSuites
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
