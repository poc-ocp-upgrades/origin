package util

import (
	"reflect"
	"testing"
	corev1 "k8s.io/api/core/v1"
	kmeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	appsv1 "github.com/openshift/api/apps/v1"
	appsapi "github.com/openshift/origin/pkg/apps/apis/apps"
	testtypes "github.com/openshift/origin/pkg/util/testing"
)

func TestAddConfigLabels(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var nilLabels map[string]string
	testCases := []struct {
		obj		runtime.Object
		addLabels	map[string]string
		err		bool
		expectedLabels	map[string]string
	}{{obj: &corev1.Pod{}, addLabels: nilLabels, err: false, expectedLabels: nilLabels}, {obj: &corev1.Pod{}, addLabels: map[string]string{}, err: false, expectedLabels: map[string]string{}}, {obj: &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "bar"}}}, addLabels: nilLabels, err: false, expectedLabels: map[string]string{"foo": "bar"}}, {obj: &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "bar"}}}, addLabels: map[string]string{}, err: false, expectedLabels: map[string]string{"foo": "bar"}}, {obj: &corev1.Pod{}, addLabels: map[string]string{"foo": "bar"}, err: false, expectedLabels: map[string]string{"foo": "bar"}}, {obj: &corev1.Service{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"baz": ""}}}, addLabels: map[string]string{"foo": "bar"}, err: false, expectedLabels: map[string]string{"foo": "bar", "baz": ""}}, {obj: &corev1.Service{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "same value"}}}, addLabels: map[string]string{"foo": "same value"}, err: false, expectedLabels: map[string]string{"foo": "same value"}}, {obj: &corev1.Service{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "first value"}}}, addLabels: map[string]string{"foo": "second value"}, err: false, expectedLabels: map[string]string{"foo": "second value"}}, {obj: &corev1.ReplicationController{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "same value"}}, Spec: corev1.ReplicationControllerSpec{Template: &corev1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{}}}}}, addLabels: map[string]string{"foo": "same value"}, err: false, expectedLabels: map[string]string{"foo": "same value"}}, {obj: &appsv1.DeploymentConfig{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "first value"}}, Spec: appsv1.DeploymentConfigSpec{Template: &corev1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"foo": "first value"}}}}}, addLabels: map[string]string{"bar": "second value"}, err: false, expectedLabels: map[string]string{"foo": "first value", "bar": "second value"}}, {obj: &testtypes.FakeLabelsResource{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"baz": ""}}}, addLabels: map[string]string{"foo": "bar"}, err: false, expectedLabels: map[string]string{"foo": "bar", "baz": ""}}}
	for i, test := range testCases {
		err := AddObjectLabels(test.obj, test.addLabels)
		if err != nil && !test.err {
			t.Errorf("Unexpected error while setting labels on testCase[%v]: %v.", i, err)
		} else if err == nil && test.err {
			t.Errorf("Unexpected non-error while setting labels on testCase[%v].", i)
		}
		accessor, err := kmeta.Accessor(test.obj)
		if err != nil {
			t.Error(err)
		}
		metaLabels := accessor.GetLabels()
		if e, a := test.expectedLabels, metaLabels; !reflect.DeepEqual(e, a) {
			t.Errorf("Unexpected labels on testCase[%v]. Expected: %#v, got: %#v.", i, e, a)
		}
		switch objType := test.obj.(type) {
		case *kapi.ReplicationController:
			if e, a := map[string]string{}, objType.Spec.Template.Labels; !reflect.DeepEqual(e, a) {
				t.Errorf("Unexpected labels on testCase[%v]. Expected: %#v, got: %#v.", i, e, a)
			}
		case *appsapi.DeploymentConfig:
			if e, a := test.expectedLabels, objType.Spec.Template.Labels; !reflect.DeepEqual(e, a) {
				t.Errorf("Unexpected labels on testCase[%v]. Expected: %#v, got: %#v.", i, e, a)
			}
		}
	}
}
func TestMergeInto(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var nilMap map[int]int
	testCases := []struct {
		dst		interface{}
		src		interface{}
		flags		int
		err		bool
		expected	interface{}
	}{{dst: nil, src: map[int]int{}, flags: 0, err: true, expected: nil}, {dst: map[int]int{}, src: nil, flags: 0, err: true, expected: map[int]int{}}, {dst: map[int]int{}, src: nilMap, flags: 0, err: false, expected: map[int]int{}}, {dst: nilMap, src: map[int]int{}, flags: 0, err: true, expected: nilMap}, {dst: &nilMap, src: map[int]int{}, flags: 0, err: true, expected: &nilMap}, {dst: map[int]int{}, src: map[int]int{}, flags: 0, err: false, expected: map[int]int{}}, {dst: map[int]byte{0: 0, 1: 1}, src: map[int]byte{2: 2, 3: 3}, flags: 0, err: false, expected: map[int]byte{0: 0, 1: 1, 2: 2, 3: 3}}, {dst: map[string]string{"foo": "bar"}, src: map[string]string{"foo": ""}, flags: 0, err: false, expected: map[string]string{"foo": "bar"}}, {dst: map[string]string{"foo": "bar"}, src: map[string]string{"foo": ""}, flags: OverwriteExistingDstKey, err: false, expected: map[string]string{"foo": ""}}, {dst: map[string]string{"foo": "bar"}, src: map[string]string{"foo": "bar"}, flags: ErrorOnExistingDstKey | OverwriteExistingDstKey, err: true, expected: map[string]string{"foo": "bar"}}, {dst: map[string]string{"foo": "bar"}, src: map[string]string{"foo": "bar"}, flags: ErrorOnDifferentDstKeyValue | OverwriteExistingDstKey, err: false, expected: map[string]string{"foo": "bar"}}, {dst: map[string]string{"foo": "bar"}, src: map[string]string{"foo": ""}, flags: ErrorOnDifferentDstKeyValue | OverwriteExistingDstKey, err: true, expected: map[string]string{"foo": "bar"}}}
	for i, test := range testCases {
		err := MergeInto(test.dst, test.src, test.flags)
		if err != nil && !test.err {
			t.Errorf("Unexpected error while merging maps on testCase[%v]: %v.", i, err)
		} else if err == nil && test.err {
			t.Errorf("Unexpected non-error while merging maps on testCase[%v].", i)
		}
		if !reflect.DeepEqual(test.dst, test.expected) {
			t.Errorf("Unexpected map on testCase[%v]. Expected: %#v, got: %#v.", i, test.expected, test.dst)
		}
	}
}
