package main

import (
	"bytes"
	"encoding/json"
	goflag "flag"
	"fmt"
	goformat "fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/spf13/pflag"
	"k8s.io/klog"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

var (
	listenAddress        string
	metricsPath          string
	etcdVersionScrapeURI string
	etcdMetricsScrapeURI string
	scrapeTimeout        time.Duration
)

func registerFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(&listenAddress, "listen-address", "localhost:9101", "Address to listen on for serving prometheus metrics")
	fs.StringVar(&metricsPath, "metrics-path", "/metrics", "Path under which prometheus metrics are to be served")
	fs.StringVar(&etcdVersionScrapeURI, "etcd-version-scrape-uri", "http://localhost:2379/version", "URI to scrape etcd version info")
	fs.StringVar(&etcdMetricsScrapeURI, "etcd-metrics-scrape-uri", "http://localhost:2379/metrics", "URI to scrape etcd metrics")
	fs.DurationVar(&scrapeTimeout, "scrape-timeout", 15*time.Second, "Timeout for trying to get stats from etcd")
}

const (
	namespace = "etcd"
)

var (
	customMetricRegistry = prometheus.NewRegistry()
	etcdVersion          = prometheus.NewGaugeVec(prometheus.GaugeOpts{Namespace: namespace, Name: "version_info", Help: "Etcd server's binary version"}, []string{"binary_version"})
	gatherer             = &monitorGatherer{exported: map[string]*exportedMetric{"etcd_grpc_requests_total": {rewriters: []rewriteFunc{func(mf *dto.MetricFamily) (*dto.MetricFamily, error) {
		mf = deepCopyMetricFamily(mf)
		renameLabels(mf, map[string]string{"grpc_method": "method", "grpc_service": "service"})
		return mf, nil
	}}}, "grpc_server_handled_total": {rewriters: []rewriteFunc{identity, func(mf *dto.MetricFamily) (*dto.MetricFamily, error) {
		mf = deepCopyMetricFamily(mf)
		renameMetric(mf, "etcd_grpc_requests_total")
		renameLabels(mf, map[string]string{"grpc_method": "method", "grpc_service": "service"})
		filterMetricsByLabels(mf, map[string]string{"grpc_type": "unary"})
		groupCounterMetricsByLabels(mf, map[string]bool{"grpc_type": true, "grpc_code": true})
		return mf, nil
	}}}, "etcd_grpc_unary_requests_duration_seconds": {rewriters: []rewriteFunc{func(mf *dto.MetricFamily) (*dto.MetricFamily, error) {
		mf = deepCopyMetricFamily(mf)
		renameMetric(mf, "grpc_server_handling_seconds")
		tpeName := "grpc_type"
		tpeVal := "unary"
		for _, m := range mf.Metric {
			m.Label = append(m.Label, &dto.LabelPair{Name: &tpeName, Value: &tpeVal})
		}
		return mf, nil
	}}}, "grpc_server_handling_seconds": {}}}
)

type monitorGatherer struct{ exported map[string]*exportedMetric }
type exportedMetric struct{ rewriters []rewriteFunc }
type rewriteFunc func(mf *dto.MetricFamily) (*dto.MetricFamily, error)

func (m *monitorGatherer) Gather() ([]*dto.MetricFamily, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdMetrics, err := scrapeMetrics()
	if err != nil {
		return nil, err
	}
	exported, err := m.rewriteExportedMetrics(etcdMetrics)
	if err != nil {
		return nil, err
	}
	custom, err := customMetricRegistry.Gather()
	if err != nil {
		return nil, err
	}
	result := make([]*dto.MetricFamily, 0, len(exported)+len(custom))
	result = append(result, exported...)
	result = append(result, custom...)
	return result, nil
}
func (m *monitorGatherer) rewriteExportedMetrics(metrics map[string]*dto.MetricFamily) ([]*dto.MetricFamily, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	results := make([]*dto.MetricFamily, 0, len(metrics))
	for n, mf := range metrics {
		if e, ok := m.exported[n]; ok {
			if e.rewriters == nil {
				results = append(results, mf)
			} else {
				for _, rewriter := range e.rewriters {
					new, err := rewriter(mf)
					if err != nil {
						return nil, err
					}
					results = append(results, new)
				}
			}
		} else {
			results = append(results, mf)
		}
	}
	return results, nil
}

type EtcdVersion struct {
	BinaryVersion  string `json:"etcdserver"`
	ClusterVersion string `json:"etcdcluster"`
}

func getVersion(lastSeenBinaryVersion *string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", etcdVersionScrapeURI, nil)
	if err != nil {
		return fmt.Errorf("Failed to create GET request for etcd version: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to receive GET response for etcd version: %v", err)
	}
	defer resp.Body.Close()
	var version EtcdVersion
	if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
		return fmt.Errorf("Failed to decode etcd version JSON: %v", err)
	}
	if *lastSeenBinaryVersion == version.BinaryVersion {
		return nil
	}
	if *lastSeenBinaryVersion != "" {
		deleted := etcdVersion.Delete(prometheus.Labels{"binary_version": *lastSeenBinaryVersion})
		if !deleted {
			return fmt.Errorf("Failed to delete previous version's metric")
		}
	}
	etcdVersion.With(prometheus.Labels{"binary_version": version.BinaryVersion}).Set(0)
	*lastSeenBinaryVersion = version.BinaryVersion
	return nil
}
func getVersionPeriodically(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	lastSeenBinaryVersion := ""
	for {
		if err := getVersion(&lastSeenBinaryVersion); err != nil {
			klog.Errorf("Failed to fetch etcd version: %v", err)
		}
		select {
		case <-stopCh:
			break
		case <-time.After(scrapeTimeout):
		}
	}
}
func scrapeMetrics() (map[string]*dto.MetricFamily, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest("GET", etcdMetricsScrapeURI, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create GET request for etcd metrics: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to receive GET response for etcd metrics: %v", err)
	}
	defer resp.Body.Close()
	var textParser expfmt.TextParser
	return textParser.TextToMetricFamilies(resp.Body)
}
func renameMetric(mf *dto.MetricFamily, name string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mf.Name = &name
}
func renameLabels(mf *dto.MetricFamily, nameMapping map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, m := range mf.Metric {
		for _, lbl := range m.Label {
			if alias, ok := nameMapping[*lbl.Name]; ok {
				lbl.Name = &alias
			}
		}
	}
}
func filterMetricsByLabels(mf *dto.MetricFamily, labelValues map[string]string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buf := mf.Metric[:0]
	for _, m := range mf.Metric {
		shouldRemove := false
		for _, lbl := range m.Label {
			if val, ok := labelValues[*lbl.Name]; ok && val != *lbl.Value {
				shouldRemove = true
				break
			}
		}
		if !shouldRemove {
			buf = append(buf, m)
		}
	}
	mf.Metric = buf
}
func groupCounterMetricsByLabels(mf *dto.MetricFamily, names map[string]bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buf := mf.Metric[:0]
	deleteLabels(mf, names)
	byLabels := map[string]*dto.Metric{}
	for _, m := range mf.Metric {
		if metric, ok := byLabels[labelsKey(m.Label)]; ok {
			metric.Counter.Value = proto.Float64(*metric.Counter.Value + *m.Counter.Value)
		} else {
			byLabels[labelsKey(m.Label)] = m
			buf = append(buf, m)
		}
	}
	mf.Metric = buf
}
func labelsKey(lbls []*dto.LabelPair) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var buf bytes.Buffer
	for i, lbl := range lbls {
		buf.WriteString(lbl.String())
		if i < len(lbls)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
func deleteLabels(mf *dto.MetricFamily, names map[string]bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, m := range mf.Metric {
		buf := m.Label[:0]
		for _, lbl := range m.Label {
			shouldRemove := names[*lbl.Name]
			if !shouldRemove {
				buf = append(buf, lbl)
			}
		}
		m.Label = buf
	}
}
func identity(mf *dto.MetricFamily) (*dto.MetricFamily, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return mf, nil
}
func deepCopyMetricFamily(mf *dto.MetricFamily) *dto.MetricFamily {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := &dto.MetricFamily{}
	r.Name = mf.Name
	r.Help = mf.Help
	r.Type = mf.Type
	r.Metric = make([]*dto.Metric, len(mf.Metric))
	for i, m := range mf.Metric {
		r.Metric[i] = deepCopyMetric(m)
	}
	return r
}
func deepCopyMetric(m *dto.Metric) *dto.Metric {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := &dto.Metric{}
	r.Label = make([]*dto.LabelPair, len(m.Label))
	for i, lp := range m.Label {
		r.Label[i] = deepCopyLabelPair(lp)
	}
	r.Gauge = m.Gauge
	r.Counter = m.Counter
	r.Summary = m.Summary
	r.Untyped = m.Untyped
	r.Histogram = m.Histogram
	r.TimestampMs = m.TimestampMs
	return r
}
func deepCopyLabelPair(lp *dto.LabelPair) *dto.LabelPair {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := &dto.LabelPair{}
	r.Name = lp.Name
	r.Value = lp.Value
	return r
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	registerFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	customMetricRegistry.MustRegister(etcdVersion)
	customMetricRegistry.Unregister(prometheus.NewGoCollector())
	stopCh := make(chan struct{})
	defer close(stopCh)
	go getVersionPeriodically(stopCh)
	klog.Infof("Listening on: %v", listenAddress)
	http.Handle(metricsPath, promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}))
	klog.Errorf("Stopped listening/serving metrics: %v", http.ListenAndServe(listenAddress, nil))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
