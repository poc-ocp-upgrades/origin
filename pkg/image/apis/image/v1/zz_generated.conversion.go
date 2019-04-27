package v1

import (
	unsafe "unsafe"
	v1 "github.com/openshift/api/image/v1"
	image "github.com/openshift/origin/pkg/image/apis/image"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
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
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
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
	if err := s.AddGeneratedConversionFunc((*v1.Image)(nil), (*image.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Image_To_image_Image(a.(*v1.Image), b.(*image.Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.Image)(nil), (*v1.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_Image_To_v1_Image(a.(*image.Image), b.(*v1.Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageBlobReferences)(nil), (*image.ImageBlobReferences)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageBlobReferences_To_image_ImageBlobReferences(a.(*v1.ImageBlobReferences), b.(*image.ImageBlobReferences), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageBlobReferences)(nil), (*v1.ImageBlobReferences)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageBlobReferences_To_v1_ImageBlobReferences(a.(*image.ImageBlobReferences), b.(*v1.ImageBlobReferences), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageImportSpec)(nil), (*image.ImageImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageImportSpec_To_image_ImageImportSpec(a.(*v1.ImageImportSpec), b.(*image.ImageImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageImportSpec)(nil), (*v1.ImageImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageImportSpec_To_v1_ImageImportSpec(a.(*image.ImageImportSpec), b.(*v1.ImageImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageImportStatus)(nil), (*image.ImageImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageImportStatus_To_image_ImageImportStatus(a.(*v1.ImageImportStatus), b.(*image.ImageImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageImportStatus)(nil), (*v1.ImageImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageImportStatus_To_v1_ImageImportStatus(a.(*image.ImageImportStatus), b.(*v1.ImageImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageLayer)(nil), (*image.ImageLayer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageLayer_To_image_ImageLayer(a.(*v1.ImageLayer), b.(*image.ImageLayer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageLayer)(nil), (*v1.ImageLayer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageLayer_To_v1_ImageLayer(a.(*image.ImageLayer), b.(*v1.ImageLayer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageLayerData)(nil), (*image.ImageLayerData)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageLayerData_To_image_ImageLayerData(a.(*v1.ImageLayerData), b.(*image.ImageLayerData), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageLayerData)(nil), (*v1.ImageLayerData)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageLayerData_To_v1_ImageLayerData(a.(*image.ImageLayerData), b.(*v1.ImageLayerData), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageList)(nil), (*image.ImageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageList_To_image_ImageList(a.(*v1.ImageList), b.(*image.ImageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageList)(nil), (*v1.ImageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageList_To_v1_ImageList(a.(*image.ImageList), b.(*v1.ImageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageLookupPolicy)(nil), (*image.ImageLookupPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(a.(*v1.ImageLookupPolicy), b.(*image.ImageLookupPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageLookupPolicy)(nil), (*v1.ImageLookupPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(a.(*image.ImageLookupPolicy), b.(*v1.ImageLookupPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageSignature)(nil), (*image.ImageSignature)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageSignature_To_image_ImageSignature(a.(*v1.ImageSignature), b.(*image.ImageSignature), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageSignature)(nil), (*v1.ImageSignature)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageSignature_To_v1_ImageSignature(a.(*image.ImageSignature), b.(*v1.ImageSignature), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStream)(nil), (*image.ImageStream)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStream_To_image_ImageStream(a.(*v1.ImageStream), b.(*image.ImageStream), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStream)(nil), (*v1.ImageStream)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStream_To_v1_ImageStream(a.(*image.ImageStream), b.(*v1.ImageStream), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamImage)(nil), (*image.ImageStreamImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamImage_To_image_ImageStreamImage(a.(*v1.ImageStreamImage), b.(*image.ImageStreamImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamImage)(nil), (*v1.ImageStreamImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamImage_To_v1_ImageStreamImage(a.(*image.ImageStreamImage), b.(*v1.ImageStreamImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamImport)(nil), (*image.ImageStreamImport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamImport_To_image_ImageStreamImport(a.(*v1.ImageStreamImport), b.(*image.ImageStreamImport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamImport)(nil), (*v1.ImageStreamImport)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamImport_To_v1_ImageStreamImport(a.(*image.ImageStreamImport), b.(*v1.ImageStreamImport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamImportSpec)(nil), (*image.ImageStreamImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamImportSpec_To_image_ImageStreamImportSpec(a.(*v1.ImageStreamImportSpec), b.(*image.ImageStreamImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamImportSpec)(nil), (*v1.ImageStreamImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamImportSpec_To_v1_ImageStreamImportSpec(a.(*image.ImageStreamImportSpec), b.(*v1.ImageStreamImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamImportStatus)(nil), (*image.ImageStreamImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamImportStatus_To_image_ImageStreamImportStatus(a.(*v1.ImageStreamImportStatus), b.(*image.ImageStreamImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamImportStatus)(nil), (*v1.ImageStreamImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamImportStatus_To_v1_ImageStreamImportStatus(a.(*image.ImageStreamImportStatus), b.(*v1.ImageStreamImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamLayers)(nil), (*image.ImageStreamLayers)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamLayers_To_image_ImageStreamLayers(a.(*v1.ImageStreamLayers), b.(*image.ImageStreamLayers), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamLayers)(nil), (*v1.ImageStreamLayers)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamLayers_To_v1_ImageStreamLayers(a.(*image.ImageStreamLayers), b.(*v1.ImageStreamLayers), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamList)(nil), (*image.ImageStreamList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamList_To_image_ImageStreamList(a.(*v1.ImageStreamList), b.(*image.ImageStreamList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamList)(nil), (*v1.ImageStreamList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamList_To_v1_ImageStreamList(a.(*image.ImageStreamList), b.(*v1.ImageStreamList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamMapping)(nil), (*image.ImageStreamMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamMapping_To_image_ImageStreamMapping(a.(*v1.ImageStreamMapping), b.(*image.ImageStreamMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamMapping)(nil), (*v1.ImageStreamMapping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamMapping_To_v1_ImageStreamMapping(a.(*image.ImageStreamMapping), b.(*v1.ImageStreamMapping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamSpec)(nil), (*image.ImageStreamSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamSpec_To_image_ImageStreamSpec(a.(*v1.ImageStreamSpec), b.(*image.ImageStreamSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamSpec)(nil), (*v1.ImageStreamSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamSpec_To_v1_ImageStreamSpec(a.(*image.ImageStreamSpec), b.(*v1.ImageStreamSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamStatus)(nil), (*image.ImageStreamStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamStatus_To_image_ImageStreamStatus(a.(*v1.ImageStreamStatus), b.(*image.ImageStreamStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamStatus)(nil), (*v1.ImageStreamStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamStatus_To_v1_ImageStreamStatus(a.(*image.ImageStreamStatus), b.(*v1.ImageStreamStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamTag)(nil), (*image.ImageStreamTag)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamTag_To_image_ImageStreamTag(a.(*v1.ImageStreamTag), b.(*image.ImageStreamTag), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamTag)(nil), (*v1.ImageStreamTag)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamTag_To_v1_ImageStreamTag(a.(*image.ImageStreamTag), b.(*v1.ImageStreamTag), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ImageStreamTagList)(nil), (*image.ImageStreamTagList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamTagList_To_image_ImageStreamTagList(a.(*v1.ImageStreamTagList), b.(*image.ImageStreamTagList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.ImageStreamTagList)(nil), (*v1.ImageStreamTagList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamTagList_To_v1_ImageStreamTagList(a.(*image.ImageStreamTagList), b.(*v1.ImageStreamTagList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RepositoryImportSpec)(nil), (*image.RepositoryImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositoryImportSpec_To_image_RepositoryImportSpec(a.(*v1.RepositoryImportSpec), b.(*image.RepositoryImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.RepositoryImportSpec)(nil), (*v1.RepositoryImportSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_RepositoryImportSpec_To_v1_RepositoryImportSpec(a.(*image.RepositoryImportSpec), b.(*v1.RepositoryImportSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.RepositoryImportStatus)(nil), (*image.RepositoryImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositoryImportStatus_To_image_RepositoryImportStatus(a.(*v1.RepositoryImportStatus), b.(*image.RepositoryImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.RepositoryImportStatus)(nil), (*v1.RepositoryImportStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_RepositoryImportStatus_To_v1_RepositoryImportStatus(a.(*image.RepositoryImportStatus), b.(*v1.RepositoryImportStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SignatureCondition)(nil), (*image.SignatureCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SignatureCondition_To_image_SignatureCondition(a.(*v1.SignatureCondition), b.(*image.SignatureCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.SignatureCondition)(nil), (*v1.SignatureCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_SignatureCondition_To_v1_SignatureCondition(a.(*image.SignatureCondition), b.(*v1.SignatureCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SignatureGenericEntity)(nil), (*image.SignatureGenericEntity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(a.(*v1.SignatureGenericEntity), b.(*image.SignatureGenericEntity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.SignatureGenericEntity)(nil), (*v1.SignatureGenericEntity)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(a.(*image.SignatureGenericEntity), b.(*v1.SignatureGenericEntity), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SignatureIssuer)(nil), (*image.SignatureIssuer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SignatureIssuer_To_image_SignatureIssuer(a.(*v1.SignatureIssuer), b.(*image.SignatureIssuer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.SignatureIssuer)(nil), (*v1.SignatureIssuer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_SignatureIssuer_To_v1_SignatureIssuer(a.(*image.SignatureIssuer), b.(*v1.SignatureIssuer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.SignatureSubject)(nil), (*image.SignatureSubject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SignatureSubject_To_image_SignatureSubject(a.(*v1.SignatureSubject), b.(*image.SignatureSubject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.SignatureSubject)(nil), (*v1.SignatureSubject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_SignatureSubject_To_v1_SignatureSubject(a.(*image.SignatureSubject), b.(*v1.SignatureSubject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagEvent)(nil), (*image.TagEvent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagEvent_To_image_TagEvent(a.(*v1.TagEvent), b.(*image.TagEvent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.TagEvent)(nil), (*v1.TagEvent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_TagEvent_To_v1_TagEvent(a.(*image.TagEvent), b.(*v1.TagEvent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagEventCondition)(nil), (*image.TagEventCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagEventCondition_To_image_TagEventCondition(a.(*v1.TagEventCondition), b.(*image.TagEventCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.TagEventCondition)(nil), (*v1.TagEventCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_TagEventCondition_To_v1_TagEventCondition(a.(*image.TagEventCondition), b.(*v1.TagEventCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagImportPolicy)(nil), (*image.TagImportPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagImportPolicy_To_image_TagImportPolicy(a.(*v1.TagImportPolicy), b.(*image.TagImportPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.TagImportPolicy)(nil), (*v1.TagImportPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_TagImportPolicy_To_v1_TagImportPolicy(a.(*image.TagImportPolicy), b.(*v1.TagImportPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagReference)(nil), (*image.TagReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagReference_To_image_TagReference(a.(*v1.TagReference), b.(*image.TagReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.TagReference)(nil), (*v1.TagReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_TagReference_To_v1_TagReference(a.(*image.TagReference), b.(*v1.TagReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.TagReferencePolicy)(nil), (*image.TagReferencePolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TagReferencePolicy_To_image_TagReferencePolicy(a.(*v1.TagReferencePolicy), b.(*image.TagReferencePolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*image.TagReferencePolicy)(nil), (*v1.TagReferencePolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_TagReferencePolicy_To_v1_TagReferencePolicy(a.(*image.TagReferencePolicy), b.(*v1.TagReferencePolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*image.ImageStreamSpec)(nil), (*v1.ImageStreamSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamSpec_To_v1_ImageStreamSpec(a.(*image.ImageStreamSpec), b.(*v1.ImageStreamSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*image.ImageStreamStatus)(nil), (*v1.ImageStreamStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_ImageStreamStatus_To_v1_ImageStreamStatus(a.(*image.ImageStreamStatus), b.(*v1.ImageStreamStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*image.Image)(nil), (*v1.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_image_Image_To_v1_Image(a.(*image.Image), b.(*v1.Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ImageStreamSpec)(nil), (*image.ImageStreamSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamSpec_To_image_ImageStreamSpec(a.(*v1.ImageStreamSpec), b.(*image.ImageStreamSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.ImageStreamStatus)(nil), (*image.ImageStreamStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ImageStreamStatus_To_image_ImageStreamStatus(a.(*v1.ImageStreamStatus), b.(*image.ImageStreamStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.Image)(nil), (*image.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Image_To_image_Image(a.(*v1.Image), b.(*image.Image), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1_Image_To_image_Image(in *v1.Image, out *image.Image, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.DockerImageReference = in.DockerImageReference
	if err := s.Convert(&in.DockerImageMetadata, &out.DockerImageMetadata, 0); err != nil {
		return err
	}
	out.DockerImageMetadataVersion = in.DockerImageMetadataVersion
	out.DockerImageManifest = in.DockerImageManifest
	out.DockerImageLayers = *(*[]image.ImageLayer)(unsafe.Pointer(&in.DockerImageLayers))
	out.Signatures = *(*[]image.ImageSignature)(unsafe.Pointer(&in.Signatures))
	out.DockerImageSignatures = *(*[][]byte)(unsafe.Pointer(&in.DockerImageSignatures))
	out.DockerImageManifestMediaType = in.DockerImageManifestMediaType
	out.DockerImageConfig = in.DockerImageConfig
	return nil
}
func autoConvert_image_Image_To_v1_Image(in *image.Image, out *v1.Image, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.DockerImageReference = in.DockerImageReference
	if err := s.Convert(&in.DockerImageMetadata, &out.DockerImageMetadata, 0); err != nil {
		return err
	}
	out.DockerImageMetadataVersion = in.DockerImageMetadataVersion
	out.DockerImageManifest = in.DockerImageManifest
	out.DockerImageLayers = *(*[]v1.ImageLayer)(unsafe.Pointer(&in.DockerImageLayers))
	out.Signatures = *(*[]v1.ImageSignature)(unsafe.Pointer(&in.Signatures))
	out.DockerImageSignatures = *(*[][]byte)(unsafe.Pointer(&in.DockerImageSignatures))
	out.DockerImageManifestMediaType = in.DockerImageManifestMediaType
	out.DockerImageConfig = in.DockerImageConfig
	return nil
}
func autoConvert_v1_ImageBlobReferences_To_image_ImageBlobReferences(in *v1.ImageBlobReferences, out *image.ImageBlobReferences, s conversion.Scope) error {
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
	out.ImageMissing = in.ImageMissing
	out.Layers = *(*[]string)(unsafe.Pointer(&in.Layers))
	out.Config = (*string)(unsafe.Pointer(in.Config))
	return nil
}
func Convert_v1_ImageBlobReferences_To_image_ImageBlobReferences(in *v1.ImageBlobReferences, out *image.ImageBlobReferences, s conversion.Scope) error {
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
	return autoConvert_v1_ImageBlobReferences_To_image_ImageBlobReferences(in, out, s)
}
func autoConvert_image_ImageBlobReferences_To_v1_ImageBlobReferences(in *image.ImageBlobReferences, out *v1.ImageBlobReferences, s conversion.Scope) error {
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
	out.ImageMissing = in.ImageMissing
	out.Layers = *(*[]string)(unsafe.Pointer(&in.Layers))
	out.Config = (*string)(unsafe.Pointer(in.Config))
	return nil
}
func Convert_image_ImageBlobReferences_To_v1_ImageBlobReferences(in *image.ImageBlobReferences, out *v1.ImageBlobReferences, s conversion.Scope) error {
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
	return autoConvert_image_ImageBlobReferences_To_v1_ImageBlobReferences(in, out, s)
}
func autoConvert_v1_ImageImportSpec_To_image_ImageImportSpec(in *v1.ImageImportSpec, out *image.ImageImportSpec, s conversion.Scope) error {
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
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(core.LocalObjectReference)
		if err := corev1.Convert_v1_LocalObjectReference_To_core_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if err := Convert_v1_TagImportPolicy_To_image_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_v1_TagReferencePolicy_To_image_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}
func Convert_v1_ImageImportSpec_To_image_ImageImportSpec(in *v1.ImageImportSpec, out *image.ImageImportSpec, s conversion.Scope) error {
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
	return autoConvert_v1_ImageImportSpec_To_image_ImageImportSpec(in, out, s)
}
func autoConvert_image_ImageImportSpec_To_v1_ImageImportSpec(in *image.ImageImportSpec, out *v1.ImageImportSpec, s conversion.Scope) error {
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
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(apicorev1.LocalObjectReference)
		if err := corev1.Convert_core_LocalObjectReference_To_v1_LocalObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.To = nil
	}
	if err := Convert_image_TagImportPolicy_To_v1_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_image_TagReferencePolicy_To_v1_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}
func Convert_image_ImageImportSpec_To_v1_ImageImportSpec(in *image.ImageImportSpec, out *v1.ImageImportSpec, s conversion.Scope) error {
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
	return autoConvert_image_ImageImportSpec_To_v1_ImageImportSpec(in, out, s)
}
func autoConvert_v1_ImageImportStatus_To_image_ImageImportStatus(in *v1.ImageImportStatus, out *image.ImageImportStatus, s conversion.Scope) error {
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
	out.Status = in.Status
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(image.Image)
		if err := Convert_v1_Image_To_image_Image(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Image = nil
	}
	out.Tag = in.Tag
	return nil
}
func Convert_v1_ImageImportStatus_To_image_ImageImportStatus(in *v1.ImageImportStatus, out *image.ImageImportStatus, s conversion.Scope) error {
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
	return autoConvert_v1_ImageImportStatus_To_image_ImageImportStatus(in, out, s)
}
func autoConvert_image_ImageImportStatus_To_v1_ImageImportStatus(in *image.ImageImportStatus, out *v1.ImageImportStatus, s conversion.Scope) error {
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
	out.Tag = in.Tag
	out.Status = in.Status
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(v1.Image)
		if err := Convert_image_Image_To_v1_Image(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Image = nil
	}
	return nil
}
func Convert_image_ImageImportStatus_To_v1_ImageImportStatus(in *image.ImageImportStatus, out *v1.ImageImportStatus, s conversion.Scope) error {
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
	return autoConvert_image_ImageImportStatus_To_v1_ImageImportStatus(in, out, s)
}
func autoConvert_v1_ImageLayer_To_image_ImageLayer(in *v1.ImageLayer, out *image.ImageLayer, s conversion.Scope) error {
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
	out.Name = in.Name
	out.LayerSize = in.LayerSize
	out.MediaType = in.MediaType
	return nil
}
func Convert_v1_ImageLayer_To_image_ImageLayer(in *v1.ImageLayer, out *image.ImageLayer, s conversion.Scope) error {
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
	return autoConvert_v1_ImageLayer_To_image_ImageLayer(in, out, s)
}
func autoConvert_image_ImageLayer_To_v1_ImageLayer(in *image.ImageLayer, out *v1.ImageLayer, s conversion.Scope) error {
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
	out.Name = in.Name
	out.LayerSize = in.LayerSize
	out.MediaType = in.MediaType
	return nil
}
func Convert_image_ImageLayer_To_v1_ImageLayer(in *image.ImageLayer, out *v1.ImageLayer, s conversion.Scope) error {
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
	return autoConvert_image_ImageLayer_To_v1_ImageLayer(in, out, s)
}
func autoConvert_v1_ImageLayerData_To_image_ImageLayerData(in *v1.ImageLayerData, out *image.ImageLayerData, s conversion.Scope) error {
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
	out.LayerSize = (*int64)(unsafe.Pointer(in.LayerSize))
	out.MediaType = in.MediaType
	return nil
}
func Convert_v1_ImageLayerData_To_image_ImageLayerData(in *v1.ImageLayerData, out *image.ImageLayerData, s conversion.Scope) error {
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
	return autoConvert_v1_ImageLayerData_To_image_ImageLayerData(in, out, s)
}
func autoConvert_image_ImageLayerData_To_v1_ImageLayerData(in *image.ImageLayerData, out *v1.ImageLayerData, s conversion.Scope) error {
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
	out.LayerSize = (*int64)(unsafe.Pointer(in.LayerSize))
	out.MediaType = in.MediaType
	return nil
}
func Convert_image_ImageLayerData_To_v1_ImageLayerData(in *image.ImageLayerData, out *v1.ImageLayerData, s conversion.Scope) error {
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
	return autoConvert_image_ImageLayerData_To_v1_ImageLayerData(in, out, s)
}
func autoConvert_v1_ImageList_To_image_ImageList(in *v1.ImageList, out *image.ImageList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]image.Image, len(*in))
		for i := range *in {
			if err := Convert_v1_Image_To_image_Image(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ImageList_To_image_ImageList(in *v1.ImageList, out *image.ImageList, s conversion.Scope) error {
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
	return autoConvert_v1_ImageList_To_image_ImageList(in, out, s)
}
func autoConvert_image_ImageList_To_v1_ImageList(in *image.ImageList, out *v1.ImageList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.Image, len(*in))
		for i := range *in {
			if err := Convert_image_Image_To_v1_Image(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_image_ImageList_To_v1_ImageList(in *image.ImageList, out *v1.ImageList, s conversion.Scope) error {
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
	return autoConvert_image_ImageList_To_v1_ImageList(in, out, s)
}
func autoConvert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(in *v1.ImageLookupPolicy, out *image.ImageLookupPolicy, s conversion.Scope) error {
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
	out.Local = in.Local
	return nil
}
func Convert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(in *v1.ImageLookupPolicy, out *image.ImageLookupPolicy, s conversion.Scope) error {
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
	return autoConvert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(in, out, s)
}
func autoConvert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(in *image.ImageLookupPolicy, out *v1.ImageLookupPolicy, s conversion.Scope) error {
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
	out.Local = in.Local
	return nil
}
func Convert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(in *image.ImageLookupPolicy, out *v1.ImageLookupPolicy, s conversion.Scope) error {
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
	return autoConvert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(in, out, s)
}
func autoConvert_v1_ImageSignature_To_image_ImageSignature(in *v1.ImageSignature, out *image.ImageSignature, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.Type = in.Type
	out.Content = *(*[]byte)(unsafe.Pointer(&in.Content))
	out.Conditions = *(*[]image.SignatureCondition)(unsafe.Pointer(&in.Conditions))
	out.ImageIdentity = in.ImageIdentity
	out.SignedClaims = *(*map[string]string)(unsafe.Pointer(&in.SignedClaims))
	out.Created = (*metav1.Time)(unsafe.Pointer(in.Created))
	out.IssuedBy = (*image.SignatureIssuer)(unsafe.Pointer(in.IssuedBy))
	out.IssuedTo = (*image.SignatureSubject)(unsafe.Pointer(in.IssuedTo))
	return nil
}
func Convert_v1_ImageSignature_To_image_ImageSignature(in *v1.ImageSignature, out *image.ImageSignature, s conversion.Scope) error {
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
	return autoConvert_v1_ImageSignature_To_image_ImageSignature(in, out, s)
}
func autoConvert_image_ImageSignature_To_v1_ImageSignature(in *image.ImageSignature, out *v1.ImageSignature, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.Type = in.Type
	out.Content = *(*[]byte)(unsafe.Pointer(&in.Content))
	out.Conditions = *(*[]v1.SignatureCondition)(unsafe.Pointer(&in.Conditions))
	out.ImageIdentity = in.ImageIdentity
	out.SignedClaims = *(*map[string]string)(unsafe.Pointer(&in.SignedClaims))
	out.Created = (*metav1.Time)(unsafe.Pointer(in.Created))
	out.IssuedBy = (*v1.SignatureIssuer)(unsafe.Pointer(in.IssuedBy))
	out.IssuedTo = (*v1.SignatureSubject)(unsafe.Pointer(in.IssuedTo))
	return nil
}
func Convert_image_ImageSignature_To_v1_ImageSignature(in *image.ImageSignature, out *v1.ImageSignature, s conversion.Scope) error {
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
	return autoConvert_image_ImageSignature_To_v1_ImageSignature(in, out, s)
}
func autoConvert_v1_ImageStream_To_image_ImageStream(in *v1.ImageStream, out *image.ImageStream, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ImageStreamSpec_To_image_ImageStreamSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ImageStreamStatus_To_image_ImageStreamStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ImageStream_To_image_ImageStream(in *v1.ImageStream, out *image.ImageStream, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStream_To_image_ImageStream(in, out, s)
}
func autoConvert_image_ImageStream_To_v1_ImageStream(in *image.ImageStream, out *v1.ImageStream, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_image_ImageStreamSpec_To_v1_ImageStreamSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_image_ImageStreamStatus_To_v1_ImageStreamStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_ImageStream_To_v1_ImageStream(in *image.ImageStream, out *v1.ImageStream, s conversion.Scope) error {
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
	return autoConvert_image_ImageStream_To_v1_ImageStream(in, out, s)
}
func autoConvert_v1_ImageStreamImage_To_image_ImageStreamImage(in *v1.ImageStreamImage, out *image.ImageStreamImage, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_Image_To_image_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ImageStreamImage_To_image_ImageStreamImage(in *v1.ImageStreamImage, out *image.ImageStreamImage, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamImage_To_image_ImageStreamImage(in, out, s)
}
func autoConvert_image_ImageStreamImage_To_v1_ImageStreamImage(in *image.ImageStreamImage, out *v1.ImageStreamImage, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_image_Image_To_v1_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_ImageStreamImage_To_v1_ImageStreamImage(in *image.ImageStreamImage, out *v1.ImageStreamImage, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamImage_To_v1_ImageStreamImage(in, out, s)
}
func autoConvert_v1_ImageStreamImport_To_image_ImageStreamImport(in *v1.ImageStreamImport, out *image.ImageStreamImport, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ImageStreamImportSpec_To_image_ImageStreamImportSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ImageStreamImportStatus_To_image_ImageStreamImportStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ImageStreamImport_To_image_ImageStreamImport(in *v1.ImageStreamImport, out *image.ImageStreamImport, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamImport_To_image_ImageStreamImport(in, out, s)
}
func autoConvert_image_ImageStreamImport_To_v1_ImageStreamImport(in *image.ImageStreamImport, out *v1.ImageStreamImport, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_image_ImageStreamImportSpec_To_v1_ImageStreamImportSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_image_ImageStreamImportStatus_To_v1_ImageStreamImportStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_ImageStreamImport_To_v1_ImageStreamImport(in *image.ImageStreamImport, out *v1.ImageStreamImport, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamImport_To_v1_ImageStreamImport(in, out, s)
}
func autoConvert_v1_ImageStreamImportSpec_To_image_ImageStreamImportSpec(in *v1.ImageStreamImportSpec, out *image.ImageStreamImportSpec, s conversion.Scope) error {
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
	out.Import = in.Import
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(image.RepositoryImportSpec)
		if err := Convert_v1_RepositoryImportSpec_To_image_RepositoryImportSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]image.ImageImportSpec, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageImportSpec_To_image_ImageImportSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}
func Convert_v1_ImageStreamImportSpec_To_image_ImageStreamImportSpec(in *v1.ImageStreamImportSpec, out *image.ImageStreamImportSpec, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamImportSpec_To_image_ImageStreamImportSpec(in, out, s)
}
func autoConvert_image_ImageStreamImportSpec_To_v1_ImageStreamImportSpec(in *image.ImageStreamImportSpec, out *v1.ImageStreamImportSpec, s conversion.Scope) error {
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
	out.Import = in.Import
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(v1.RepositoryImportSpec)
		if err := Convert_image_RepositoryImportSpec_To_v1_RepositoryImportSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]v1.ImageImportSpec, len(*in))
		for i := range *in {
			if err := Convert_image_ImageImportSpec_To_v1_ImageImportSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}
func Convert_image_ImageStreamImportSpec_To_v1_ImageStreamImportSpec(in *image.ImageStreamImportSpec, out *v1.ImageStreamImportSpec, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamImportSpec_To_v1_ImageStreamImportSpec(in, out, s)
}
func autoConvert_v1_ImageStreamImportStatus_To_image_ImageStreamImportStatus(in *v1.ImageStreamImportStatus, out *image.ImageStreamImportStatus, s conversion.Scope) error {
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
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(image.ImageStream)
		if err := Convert_v1_ImageStream_To_image_ImageStream(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Import = nil
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(image.RepositoryImportStatus)
		if err := Convert_v1_RepositoryImportStatus_To_image_RepositoryImportStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]image.ImageImportStatus, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageImportStatus_To_image_ImageImportStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}
func Convert_v1_ImageStreamImportStatus_To_image_ImageStreamImportStatus(in *v1.ImageStreamImportStatus, out *image.ImageStreamImportStatus, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamImportStatus_To_image_ImageStreamImportStatus(in, out, s)
}
func autoConvert_image_ImageStreamImportStatus_To_v1_ImageStreamImportStatus(in *image.ImageStreamImportStatus, out *v1.ImageStreamImportStatus, s conversion.Scope) error {
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
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(v1.ImageStream)
		if err := Convert_image_ImageStream_To_v1_ImageStream(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Import = nil
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(v1.RepositoryImportStatus)
		if err := Convert_image_RepositoryImportStatus_To_v1_RepositoryImportStatus(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]v1.ImageImportStatus, len(*in))
		for i := range *in {
			if err := Convert_image_ImageImportStatus_To_v1_ImageImportStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}
func Convert_image_ImageStreamImportStatus_To_v1_ImageStreamImportStatus(in *image.ImageStreamImportStatus, out *v1.ImageStreamImportStatus, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamImportStatus_To_v1_ImageStreamImportStatus(in, out, s)
}
func autoConvert_v1_ImageStreamLayers_To_image_ImageStreamLayers(in *v1.ImageStreamLayers, out *image.ImageStreamLayers, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.Blobs = *(*map[string]image.ImageLayerData)(unsafe.Pointer(&in.Blobs))
	out.Images = *(*map[string]image.ImageBlobReferences)(unsafe.Pointer(&in.Images))
	return nil
}
func Convert_v1_ImageStreamLayers_To_image_ImageStreamLayers(in *v1.ImageStreamLayers, out *image.ImageStreamLayers, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamLayers_To_image_ImageStreamLayers(in, out, s)
}
func autoConvert_image_ImageStreamLayers_To_v1_ImageStreamLayers(in *image.ImageStreamLayers, out *v1.ImageStreamLayers, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	out.Blobs = *(*map[string]v1.ImageLayerData)(unsafe.Pointer(&in.Blobs))
	out.Images = *(*map[string]v1.ImageBlobReferences)(unsafe.Pointer(&in.Images))
	return nil
}
func Convert_image_ImageStreamLayers_To_v1_ImageStreamLayers(in *image.ImageStreamLayers, out *v1.ImageStreamLayers, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamLayers_To_v1_ImageStreamLayers(in, out, s)
}
func autoConvert_v1_ImageStreamList_To_image_ImageStreamList(in *v1.ImageStreamList, out *image.ImageStreamList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]image.ImageStream, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageStream_To_image_ImageStream(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ImageStreamList_To_image_ImageStreamList(in *v1.ImageStreamList, out *image.ImageStreamList, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamList_To_image_ImageStreamList(in, out, s)
}
func autoConvert_image_ImageStreamList_To_v1_ImageStreamList(in *image.ImageStreamList, out *v1.ImageStreamList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ImageStream, len(*in))
		for i := range *in {
			if err := Convert_image_ImageStream_To_v1_ImageStream(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_image_ImageStreamList_To_v1_ImageStreamList(in *image.ImageStreamList, out *v1.ImageStreamList, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamList_To_v1_ImageStreamList(in, out, s)
}
func autoConvert_v1_ImageStreamMapping_To_image_ImageStreamMapping(in *v1.ImageStreamMapping, out *image.ImageStreamMapping, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_Image_To_image_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Tag = in.Tag
	return nil
}
func Convert_v1_ImageStreamMapping_To_image_ImageStreamMapping(in *v1.ImageStreamMapping, out *image.ImageStreamMapping, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamMapping_To_image_ImageStreamMapping(in, out, s)
}
func autoConvert_image_ImageStreamMapping_To_v1_ImageStreamMapping(in *image.ImageStreamMapping, out *v1.ImageStreamMapping, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_image_Image_To_v1_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Tag = in.Tag
	return nil
}
func Convert_image_ImageStreamMapping_To_v1_ImageStreamMapping(in *image.ImageStreamMapping, out *v1.ImageStreamMapping, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamMapping_To_v1_ImageStreamMapping(in, out, s)
}
func autoConvert_v1_ImageStreamSpec_To_image_ImageStreamSpec(in *v1.ImageStreamSpec, out *image.ImageStreamSpec, s conversion.Scope) error {
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
	if err := Convert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(&in.LookupPolicy, &out.LookupPolicy, s); err != nil {
		return err
	}
	out.DockerImageRepository = in.DockerImageRepository
	return nil
}
func autoConvert_image_ImageStreamSpec_To_v1_ImageStreamSpec(in *image.ImageStreamSpec, out *v1.ImageStreamSpec, s conversion.Scope) error {
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
	if err := Convert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(&in.LookupPolicy, &out.LookupPolicy, s); err != nil {
		return err
	}
	out.DockerImageRepository = in.DockerImageRepository
	return nil
}
func autoConvert_v1_ImageStreamStatus_To_image_ImageStreamStatus(in *v1.ImageStreamStatus, out *image.ImageStreamStatus, s conversion.Scope) error {
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
	out.DockerImageRepository = in.DockerImageRepository
	out.PublicDockerImageRepository = in.PublicDockerImageRepository
	return nil
}
func autoConvert_image_ImageStreamStatus_To_v1_ImageStreamStatus(in *image.ImageStreamStatus, out *v1.ImageStreamStatus, s conversion.Scope) error {
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
	out.DockerImageRepository = in.DockerImageRepository
	out.PublicDockerImageRepository = in.PublicDockerImageRepository
	return nil
}
func autoConvert_v1_ImageStreamTag_To_image_ImageStreamTag(in *v1.ImageStreamTag, out *image.ImageStreamTag, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(image.TagReference)
		if err := Convert_v1_TagReference_To_image_TagReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Tag = nil
	}
	out.Generation = in.Generation
	if err := Convert_v1_ImageLookupPolicy_To_image_ImageLookupPolicy(&in.LookupPolicy, &out.LookupPolicy, s); err != nil {
		return err
	}
	out.Conditions = *(*[]image.TagEventCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_v1_Image_To_image_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_ImageStreamTag_To_image_ImageStreamTag(in *v1.ImageStreamTag, out *image.ImageStreamTag, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamTag_To_image_ImageStreamTag(in, out, s)
}
func autoConvert_image_ImageStreamTag_To_v1_ImageStreamTag(in *image.ImageStreamTag, out *v1.ImageStreamTag, s conversion.Scope) error {
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
	out.ObjectMeta = in.ObjectMeta
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(v1.TagReference)
		if err := Convert_image_TagReference_To_v1_TagReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Tag = nil
	}
	out.Generation = in.Generation
	out.Conditions = *(*[]v1.TagEventCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_image_ImageLookupPolicy_To_v1_ImageLookupPolicy(&in.LookupPolicy, &out.LookupPolicy, s); err != nil {
		return err
	}
	if err := Convert_image_Image_To_v1_Image(&in.Image, &out.Image, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_ImageStreamTag_To_v1_ImageStreamTag(in *image.ImageStreamTag, out *v1.ImageStreamTag, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamTag_To_v1_ImageStreamTag(in, out, s)
}
func autoConvert_v1_ImageStreamTagList_To_image_ImageStreamTagList(in *v1.ImageStreamTagList, out *image.ImageStreamTagList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]image.ImageStreamTag, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageStreamTag_To_image_ImageStreamTag(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_v1_ImageStreamTagList_To_image_ImageStreamTagList(in *v1.ImageStreamTagList, out *image.ImageStreamTagList, s conversion.Scope) error {
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
	return autoConvert_v1_ImageStreamTagList_To_image_ImageStreamTagList(in, out, s)
}
func autoConvert_image_ImageStreamTagList_To_v1_ImageStreamTagList(in *image.ImageStreamTagList, out *v1.ImageStreamTagList, s conversion.Scope) error {
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ImageStreamTag, len(*in))
		for i := range *in {
			if err := Convert_image_ImageStreamTag_To_v1_ImageStreamTag(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
func Convert_image_ImageStreamTagList_To_v1_ImageStreamTagList(in *image.ImageStreamTagList, out *v1.ImageStreamTagList, s conversion.Scope) error {
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
	return autoConvert_image_ImageStreamTagList_To_v1_ImageStreamTagList(in, out, s)
}
func autoConvert_v1_RepositoryImportSpec_To_image_RepositoryImportSpec(in *v1.RepositoryImportSpec, out *image.RepositoryImportSpec, s conversion.Scope) error {
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
	if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := Convert_v1_TagImportPolicy_To_image_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_v1_TagReferencePolicy_To_image_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}
func Convert_v1_RepositoryImportSpec_To_image_RepositoryImportSpec(in *v1.RepositoryImportSpec, out *image.RepositoryImportSpec, s conversion.Scope) error {
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
	return autoConvert_v1_RepositoryImportSpec_To_image_RepositoryImportSpec(in, out, s)
}
func autoConvert_image_RepositoryImportSpec_To_v1_RepositoryImportSpec(in *image.RepositoryImportSpec, out *v1.RepositoryImportSpec, s conversion.Scope) error {
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
	if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := Convert_image_TagImportPolicy_To_v1_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_image_TagReferencePolicy_To_v1_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}
func Convert_image_RepositoryImportSpec_To_v1_RepositoryImportSpec(in *image.RepositoryImportSpec, out *v1.RepositoryImportSpec, s conversion.Scope) error {
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
	return autoConvert_image_RepositoryImportSpec_To_v1_RepositoryImportSpec(in, out, s)
}
func autoConvert_v1_RepositoryImportStatus_To_image_RepositoryImportStatus(in *v1.RepositoryImportStatus, out *image.RepositoryImportStatus, s conversion.Scope) error {
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
	out.Status = in.Status
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]image.ImageImportStatus, len(*in))
		for i := range *in {
			if err := Convert_v1_ImageImportStatus_To_image_ImageImportStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.AdditionalTags = *(*[]string)(unsafe.Pointer(&in.AdditionalTags))
	return nil
}
func Convert_v1_RepositoryImportStatus_To_image_RepositoryImportStatus(in *v1.RepositoryImportStatus, out *image.RepositoryImportStatus, s conversion.Scope) error {
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
	return autoConvert_v1_RepositoryImportStatus_To_image_RepositoryImportStatus(in, out, s)
}
func autoConvert_image_RepositoryImportStatus_To_v1_RepositoryImportStatus(in *image.RepositoryImportStatus, out *v1.RepositoryImportStatus, s conversion.Scope) error {
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
	out.Status = in.Status
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]v1.ImageImportStatus, len(*in))
		for i := range *in {
			if err := Convert_image_ImageImportStatus_To_v1_ImageImportStatus(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.AdditionalTags = *(*[]string)(unsafe.Pointer(&in.AdditionalTags))
	return nil
}
func Convert_image_RepositoryImportStatus_To_v1_RepositoryImportStatus(in *image.RepositoryImportStatus, out *v1.RepositoryImportStatus, s conversion.Scope) error {
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
	return autoConvert_image_RepositoryImportStatus_To_v1_RepositoryImportStatus(in, out, s)
}
func autoConvert_v1_SignatureCondition_To_image_SignatureCondition(in *v1.SignatureCondition, out *image.SignatureCondition, s conversion.Scope) error {
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
	out.Type = image.SignatureConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_v1_SignatureCondition_To_image_SignatureCondition(in *v1.SignatureCondition, out *image.SignatureCondition, s conversion.Scope) error {
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
	return autoConvert_v1_SignatureCondition_To_image_SignatureCondition(in, out, s)
}
func autoConvert_image_SignatureCondition_To_v1_SignatureCondition(in *image.SignatureCondition, out *v1.SignatureCondition, s conversion.Scope) error {
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
	out.Type = v1.SignatureConditionType(in.Type)
	out.Status = apicorev1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_image_SignatureCondition_To_v1_SignatureCondition(in *image.SignatureCondition, out *v1.SignatureCondition, s conversion.Scope) error {
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
	return autoConvert_image_SignatureCondition_To_v1_SignatureCondition(in, out, s)
}
func autoConvert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(in *v1.SignatureGenericEntity, out *image.SignatureGenericEntity, s conversion.Scope) error {
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
	out.Organization = in.Organization
	out.CommonName = in.CommonName
	return nil
}
func Convert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(in *v1.SignatureGenericEntity, out *image.SignatureGenericEntity, s conversion.Scope) error {
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
	return autoConvert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(in, out, s)
}
func autoConvert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(in *image.SignatureGenericEntity, out *v1.SignatureGenericEntity, s conversion.Scope) error {
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
	out.Organization = in.Organization
	out.CommonName = in.CommonName
	return nil
}
func Convert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(in *image.SignatureGenericEntity, out *v1.SignatureGenericEntity, s conversion.Scope) error {
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
	return autoConvert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(in, out, s)
}
func autoConvert_v1_SignatureIssuer_To_image_SignatureIssuer(in *v1.SignatureIssuer, out *image.SignatureIssuer, s conversion.Scope) error {
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
	if err := Convert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(&in.SignatureGenericEntity, &out.SignatureGenericEntity, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_SignatureIssuer_To_image_SignatureIssuer(in *v1.SignatureIssuer, out *image.SignatureIssuer, s conversion.Scope) error {
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
	return autoConvert_v1_SignatureIssuer_To_image_SignatureIssuer(in, out, s)
}
func autoConvert_image_SignatureIssuer_To_v1_SignatureIssuer(in *image.SignatureIssuer, out *v1.SignatureIssuer, s conversion.Scope) error {
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
	if err := Convert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(&in.SignatureGenericEntity, &out.SignatureGenericEntity, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_SignatureIssuer_To_v1_SignatureIssuer(in *image.SignatureIssuer, out *v1.SignatureIssuer, s conversion.Scope) error {
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
	return autoConvert_image_SignatureIssuer_To_v1_SignatureIssuer(in, out, s)
}
func autoConvert_v1_SignatureSubject_To_image_SignatureSubject(in *v1.SignatureSubject, out *image.SignatureSubject, s conversion.Scope) error {
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
	if err := Convert_v1_SignatureGenericEntity_To_image_SignatureGenericEntity(&in.SignatureGenericEntity, &out.SignatureGenericEntity, s); err != nil {
		return err
	}
	out.PublicKeyID = in.PublicKeyID
	return nil
}
func Convert_v1_SignatureSubject_To_image_SignatureSubject(in *v1.SignatureSubject, out *image.SignatureSubject, s conversion.Scope) error {
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
	return autoConvert_v1_SignatureSubject_To_image_SignatureSubject(in, out, s)
}
func autoConvert_image_SignatureSubject_To_v1_SignatureSubject(in *image.SignatureSubject, out *v1.SignatureSubject, s conversion.Scope) error {
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
	if err := Convert_image_SignatureGenericEntity_To_v1_SignatureGenericEntity(&in.SignatureGenericEntity, &out.SignatureGenericEntity, s); err != nil {
		return err
	}
	out.PublicKeyID = in.PublicKeyID
	return nil
}
func Convert_image_SignatureSubject_To_v1_SignatureSubject(in *image.SignatureSubject, out *v1.SignatureSubject, s conversion.Scope) error {
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
	return autoConvert_image_SignatureSubject_To_v1_SignatureSubject(in, out, s)
}
func autoConvert_v1_TagEvent_To_image_TagEvent(in *v1.TagEvent, out *image.TagEvent, s conversion.Scope) error {
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
	out.Created = in.Created
	out.DockerImageReference = in.DockerImageReference
	out.Image = in.Image
	out.Generation = in.Generation
	return nil
}
func Convert_v1_TagEvent_To_image_TagEvent(in *v1.TagEvent, out *image.TagEvent, s conversion.Scope) error {
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
	return autoConvert_v1_TagEvent_To_image_TagEvent(in, out, s)
}
func autoConvert_image_TagEvent_To_v1_TagEvent(in *image.TagEvent, out *v1.TagEvent, s conversion.Scope) error {
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
	out.Created = in.Created
	out.DockerImageReference = in.DockerImageReference
	out.Image = in.Image
	out.Generation = in.Generation
	return nil
}
func Convert_image_TagEvent_To_v1_TagEvent(in *image.TagEvent, out *v1.TagEvent, s conversion.Scope) error {
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
	return autoConvert_image_TagEvent_To_v1_TagEvent(in, out, s)
}
func autoConvert_v1_TagEventCondition_To_image_TagEventCondition(in *v1.TagEventCondition, out *image.TagEventCondition, s conversion.Scope) error {
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
	out.Type = image.TagEventConditionType(in.Type)
	out.Status = core.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	out.Generation = in.Generation
	return nil
}
func Convert_v1_TagEventCondition_To_image_TagEventCondition(in *v1.TagEventCondition, out *image.TagEventCondition, s conversion.Scope) error {
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
	return autoConvert_v1_TagEventCondition_To_image_TagEventCondition(in, out, s)
}
func autoConvert_image_TagEventCondition_To_v1_TagEventCondition(in *image.TagEventCondition, out *v1.TagEventCondition, s conversion.Scope) error {
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
	out.Type = v1.TagEventConditionType(in.Type)
	out.Status = apicorev1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	out.Generation = in.Generation
	return nil
}
func Convert_image_TagEventCondition_To_v1_TagEventCondition(in *image.TagEventCondition, out *v1.TagEventCondition, s conversion.Scope) error {
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
	return autoConvert_image_TagEventCondition_To_v1_TagEventCondition(in, out, s)
}
func autoConvert_v1_TagImportPolicy_To_image_TagImportPolicy(in *v1.TagImportPolicy, out *image.TagImportPolicy, s conversion.Scope) error {
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
	out.Insecure = in.Insecure
	out.Scheduled = in.Scheduled
	return nil
}
func Convert_v1_TagImportPolicy_To_image_TagImportPolicy(in *v1.TagImportPolicy, out *image.TagImportPolicy, s conversion.Scope) error {
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
	return autoConvert_v1_TagImportPolicy_To_image_TagImportPolicy(in, out, s)
}
func autoConvert_image_TagImportPolicy_To_v1_TagImportPolicy(in *image.TagImportPolicy, out *v1.TagImportPolicy, s conversion.Scope) error {
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
	out.Insecure = in.Insecure
	out.Scheduled = in.Scheduled
	return nil
}
func Convert_image_TagImportPolicy_To_v1_TagImportPolicy(in *image.TagImportPolicy, out *v1.TagImportPolicy, s conversion.Scope) error {
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
	return autoConvert_image_TagImportPolicy_To_v1_TagImportPolicy(in, out, s)
}
func autoConvert_v1_TagReference_To_image_TagReference(in *v1.TagReference, out *image.TagReference, s conversion.Scope) error {
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
	out.Name = in.Name
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(core.ObjectReference)
		if err := corev1.Convert_v1_ObjectReference_To_core_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Reference = in.Reference
	out.Generation = (*int64)(unsafe.Pointer(in.Generation))
	if err := Convert_v1_TagImportPolicy_To_image_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_v1_TagReferencePolicy_To_image_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1_TagReference_To_image_TagReference(in *v1.TagReference, out *image.TagReference, s conversion.Scope) error {
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
	return autoConvert_v1_TagReference_To_image_TagReference(in, out, s)
}
func autoConvert_image_TagReference_To_v1_TagReference(in *image.TagReference, out *v1.TagReference, s conversion.Scope) error {
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
	out.Name = in.Name
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(apicorev1.ObjectReference)
		if err := corev1.Convert_core_ObjectReference_To_v1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	out.Reference = in.Reference
	out.Generation = (*int64)(unsafe.Pointer(in.Generation))
	if err := Convert_image_TagImportPolicy_To_v1_TagImportPolicy(&in.ImportPolicy, &out.ImportPolicy, s); err != nil {
		return err
	}
	if err := Convert_image_TagReferencePolicy_To_v1_TagReferencePolicy(&in.ReferencePolicy, &out.ReferencePolicy, s); err != nil {
		return err
	}
	return nil
}
func Convert_image_TagReference_To_v1_TagReference(in *image.TagReference, out *v1.TagReference, s conversion.Scope) error {
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
	return autoConvert_image_TagReference_To_v1_TagReference(in, out, s)
}
func autoConvert_v1_TagReferencePolicy_To_image_TagReferencePolicy(in *v1.TagReferencePolicy, out *image.TagReferencePolicy, s conversion.Scope) error {
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
	out.Type = image.TagReferencePolicyType(in.Type)
	return nil
}
func Convert_v1_TagReferencePolicy_To_image_TagReferencePolicy(in *v1.TagReferencePolicy, out *image.TagReferencePolicy, s conversion.Scope) error {
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
	return autoConvert_v1_TagReferencePolicy_To_image_TagReferencePolicy(in, out, s)
}
func autoConvert_image_TagReferencePolicy_To_v1_TagReferencePolicy(in *image.TagReferencePolicy, out *v1.TagReferencePolicy, s conversion.Scope) error {
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
	out.Type = v1.TagReferencePolicyType(in.Type)
	return nil
}
func Convert_image_TagReferencePolicy_To_v1_TagReferencePolicy(in *image.TagReferencePolicy, out *v1.TagReferencePolicy, s conversion.Scope) error {
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
	return autoConvert_image_TagReferencePolicy_To_v1_TagReferencePolicy(in, out, s)
}
