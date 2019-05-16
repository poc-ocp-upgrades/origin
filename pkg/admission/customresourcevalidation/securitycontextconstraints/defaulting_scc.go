package securitycontextconstraints

import (
	"bytes"
	goformat "fmt"
	securityv1 "github.com/openshift/api/security/v1"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	securityapiv1 "github.com/openshift/origin/pkg/security/apis/security/v1"
	"io"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const DefaultingPluginName = "security.openshift.io/DefaultSecurityContextConstraints"

func RegisterDefaulting(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(DefaultingPluginName, func(config io.Reader) (admission.Interface, error) {
		return NewDefaulter(), nil
	})
}

type defaultSCC struct {
	*admission.Handler
	scheme       *runtime.Scheme
	codecFactory runtimeserializer.CodecFactory
}

var _ admission.MutationInterface = &defaultSCC{}

func NewDefaulter() admission.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scheme := runtime.NewScheme()
	codecFactory := runtimeserializer.NewCodecFactory(scheme)
	utilruntime.Must(securityv1.Install(scheme))
	utilruntime.Must(securityapiv1.Install(scheme))
	return &defaultSCC{Handler: admission.NewHandler(admission.Create, admission.Update), scheme: scheme, codecFactory: codecFactory}
}
func (a *defaultSCC) Admit(attributes admission.Attributes) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.shouldIgnore(attributes) {
		return nil
	}
	unstructuredOrig, ok := attributes.GetObject().(*unstructured.Unstructured)
	if !ok {
		return nil
	}
	buf := &bytes.Buffer{}
	if err := unstructured.UnstructuredJSONScheme.Encode(unstructuredOrig, buf); err != nil {
		return err
	}
	uncastObj, err := runtime.Decode(a.codecFactory.LegacyCodec(securityv1.GroupVersion), buf.Bytes())
	if err != nil {
		return err
	}
	internalSCC := uncastObj.(*securityapi.SecurityContextConstraints)
	outSCCExternal := &securityv1.SecurityContextConstraints{}
	if err := a.scheme.Convert(internalSCC, outSCCExternal, nil); err != nil {
		return apierrors.NewForbidden(attributes.GetResource().GroupResource(), attributes.GetName(), err)
	}
	defaultedBytes, err := runtime.Encode(a.codecFactory.LegacyCodec(securityv1.GroupVersion), outSCCExternal)
	if err := a.scheme.Convert(internalSCC, outSCCExternal, nil); err != nil {
		return apierrors.NewForbidden(attributes.GetResource().GroupResource(), attributes.GetName(), err)
	}
	outUnstructured := &unstructured.Unstructured{}
	if _, _, err := unstructured.UnstructuredJSONScheme.Decode(defaultedBytes, nil, outUnstructured); err != nil {
		return err
	}
	unstructuredOrig.Object = outUnstructured.Object
	return nil
}
func (a *defaultSCC) shouldIgnore(attributes admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attributes.GetResource().GroupResource() != (schema.GroupResource{Group: "security.openshift.io", Resource: "securitycontextconstraints"}) {
		return true
	}
	if len(attributes.GetSubresource()) > 0 {
		return true
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
