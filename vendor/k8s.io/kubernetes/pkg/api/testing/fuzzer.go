package testing

import (
	"fmt"
	fuzz "github.com/google/gofuzz"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	apitesting "k8s.io/apimachinery/pkg/api/apitesting"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	genericfuzzer "k8s.io/apimachinery/pkg/apis/meta/fuzzer"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	admissionregistrationfuzzer "k8s.io/kubernetes/pkg/apis/admissionregistration/fuzzer"
	"k8s.io/kubernetes/pkg/apis/apps"
	appsfuzzer "k8s.io/kubernetes/pkg/apis/apps/fuzzer"
	auditregistrationfuzzer "k8s.io/kubernetes/pkg/apis/auditregistration/fuzzer"
	autoscalingfuzzer "k8s.io/kubernetes/pkg/apis/autoscaling/fuzzer"
	batchfuzzer "k8s.io/kubernetes/pkg/apis/batch/fuzzer"
	certificatesfuzzer "k8s.io/kubernetes/pkg/apis/certificates/fuzzer"
	api "k8s.io/kubernetes/pkg/apis/core"
	corefuzzer "k8s.io/kubernetes/pkg/apis/core/fuzzer"
	extensionsfuzzer "k8s.io/kubernetes/pkg/apis/extensions/fuzzer"
	networkingfuzzer "k8s.io/kubernetes/pkg/apis/networking/fuzzer"
	policyfuzzer "k8s.io/kubernetes/pkg/apis/policy/fuzzer"
	rbacfuzzer "k8s.io/kubernetes/pkg/apis/rbac/fuzzer"
	storagefuzzer "k8s.io/kubernetes/pkg/apis/storage/fuzzer"
)

func overrideGenericFuncs(codecs runtimeserializer.CodecFactory) []interface{} {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []interface{}{func(j *runtime.Object, c fuzz.Continue) {
		if true {
			*j = &runtime.Unknown{Raw: []byte(`{"apiVersion":"unknown.group/unknown","kind":"Something","someKey":"someValue"}`), ContentType: runtime.ContentTypeJSON}
		} else {
			types := []runtime.Object{&api.Pod{}, &api.ReplicationController{}}
			t := types[c.Rand.Intn(len(types))]
			c.Fuzz(t)
			*j = t
		}
	}, func(r *runtime.RawExtension, c fuzz.Continue) {
		types := []runtime.Object{&api.Pod{}, &apps.Deployment{}, &api.Service{}}
		obj := types[c.Rand.Intn(len(types))]
		c.Fuzz(obj)
		var codec runtime.Codec
		switch obj.(type) {
		case *apps.Deployment:
			codec = apitesting.TestCodec(codecs, appsv1.SchemeGroupVersion)
		default:
			codec = apitesting.TestCodec(codecs, v1.SchemeGroupVersion)
		}
		bytes, err := runtime.Encode(codec, obj)
		if err != nil {
			panic(fmt.Sprintf("Failed to encode object: %v", err))
		}
		r.Raw = bytes
	}}
}

var FuzzerFuncs = fuzzer.MergeFuzzerFuncs(genericfuzzer.Funcs, overrideGenericFuncs, corefuzzer.Funcs, extensionsfuzzer.Funcs, appsfuzzer.Funcs, batchfuzzer.Funcs, autoscalingfuzzer.Funcs, rbacfuzzer.Funcs, policyfuzzer.Funcs, certificatesfuzzer.Funcs, admissionregistrationfuzzer.Funcs, auditregistrationfuzzer.Funcs, storagefuzzer.Funcs, networkingfuzzer.Funcs)
