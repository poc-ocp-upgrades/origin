package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	authorizationv1 "k8s.io/api/authorization/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta/testrestmapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	clientgotesting "k8s.io/client-go/testing"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/utils/clock"
	templatev1 "github.com/openshift/api/template/v1"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	batchv1.AddToScheme(legacyscheme.Scheme)
}

type roundtripper func(*http.Request) (*http.Response, error)

func (rt roundtripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rt(r)
}

type fakeClock struct {
	clock.RealClock
	now	time.Time
}

func (f *fakeClock) Now() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.now
}
func TestControllerCheckReadiness(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClock := &fakeClock{now: time.Unix(0, 0)}
	job := batchv1.Job{TypeMeta: metav1.TypeMeta{APIVersion: "batch/v1", Kind: "Job"}, ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{WaitForReadyAnnotation: "true"}}}
	fakerestconfig := &rest.Config{WrapTransport: func(http.RoundTripper) http.RoundTripper {
		return roundtripper(func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(job)
			if err != nil {
				panic(err)
			}
			return &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewBuffer(b))}, nil
		})
	}}
	client, err := dynamic.NewForConfig(fakerestconfig)
	if err != nil {
		t.Fatal(err)
	}
	fakeclientset := &fake.Clientset{}
	sarClient := fake.NewSimpleClientset()
	c := &TemplateInstanceController{dynamicRestMapper: testrestmapper.TestOnlyStaticRESTMapper(legacyscheme.Scheme, legacyscheme.Scheme.PrioritizedVersionsAllGroups()...), sarClient: sarClient.AuthorizationV1(), kc: fakeclientset, clock: fakeClock, dynamicClient: client}
	sarClient.PrependReactor("create", "subjectaccessreviews", func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &authorizationv1.SubjectAccessReview{Status: authorizationv1.SubjectAccessReviewStatus{Allowed: true}}, nil
	})
	templateInstance := &templatev1.TemplateInstance{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: metav1.Time{Time: fakeClock.now}}, Spec: templatev1.TemplateInstanceSpec{Requester: &templatev1.TemplateInstanceRequester{}}, Status: templatev1.TemplateInstanceStatus{Objects: []templatev1.TemplateInstanceObject{{Ref: corev1.ObjectReference{APIVersion: "batch/v1", Kind: "Job", Namespace: "namespace", Name: "name"}}}}}
	ready, err := c.checkReadiness(templateInstance)
	if ready || err != nil {
		t.Error(ready, err)
	}
	fakeClock.now = fakeClock.now.Add(readinessTimeout + 1)
	ready, err = c.checkReadiness(templateInstance)
	if ready || err == nil || err != TimeoutErr {
		t.Error(ready, err)
	}
	fakeClock.now = time.Unix(0, 0)
	job.Status.CompletionTime = &metav1.Time{Time: fakeClock.now}
	ready, err = c.checkReadiness(templateInstance)
	if !ready || err != nil {
		t.Error(ready, err)
	}
	job.Status.Failed = 1
	ready, err = c.checkReadiness(templateInstance)
	if ready || err == nil || err.Error() != "readiness failed on Job namespace/name" {
		t.Error(ready, err)
	}
}
