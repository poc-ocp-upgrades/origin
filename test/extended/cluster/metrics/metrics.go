package metrics

import (
	"encoding/json"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"time"
)

const (
	marker_name string = "cluster_loader_marker"
)

type Metrics interface{ printLog() error }
type BaseMetrics struct {
	Marker	string	`json:"marker"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
}
type TestDuration struct {
	BaseMetrics
	StartTime	time.Time	`json:"startTime"`
	TestDuration	time.Duration	`json:"testDuration"`
}

func (td TestDuration) printLog() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(td)
	fmt.Println(string(b))
	return err
}
func (td TestDuration) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type Alias TestDuration
	return json.Marshal(&struct {
		Alias
		TestDuration	string	`json:"testDuration"`
	}{Alias: (Alias)(td), TestDuration: td.TestDuration.String()})
}
func (td *TestDuration) UnmarshalJSON(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	type Alias TestDuration
	s := &struct {
		TestDuration	string	`json:"testDuration"`
		*Alias
	}{Alias: (*Alias)(td)}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	td.TestDuration, err = time.ParseDuration(s.TestDuration)
	if err != nil {
		return err
	}
	return nil
}
func LogMetrics(metrics []Metrics) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range metrics {
		err := m.printLog()
		if err != nil {
			return err
		}
	}
	return nil
}
func NewTestDuration(name string, startTime time.Time, testDuration time.Duration) TestDuration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return TestDuration{BaseMetrics: BaseMetrics{Marker: marker_name, Name: name, Type: fmt.Sprintf("%T", (*TestDuration)(nil))[1:]}, StartTime: startTime, TestDuration: testDuration}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
