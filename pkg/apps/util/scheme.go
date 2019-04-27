package util

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsv1 "github.com/openshift/api/apps/v1"
	"github.com/openshift/origin/pkg/api/legacy"
	appsv1helpers "github.com/openshift/origin/pkg/apps/apis/apps/v1"
)

var (
	annotationDecodingScheme	= runtime.NewScheme()
	annotationDecoder		runtime.Decoder
	annotationEncodingScheme	= runtime.NewScheme()
	annotationEncoder		runtime.Encoder
)

func init() {
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
	legacy.InstallInternalLegacyApps(annotationDecodingScheme)
	utilruntime.Must(appsv1helpers.Install(annotationDecodingScheme))
	annotationDecoderCodecFactory := serializer.NewCodecFactory(annotationDecodingScheme)
	annotationDecoder = annotationDecoderCodecFactory.UniversalDecoder(appsv1.GroupVersion)
	utilruntime.Must(appsv1helpers.Install(annotationEncodingScheme))
	annotationEncoderCodecFactory := serializer.NewCodecFactory(annotationEncodingScheme)
	annotationEncoder = annotationEncoderCodecFactory.LegacyCodec(appsv1.GroupVersion)
}
