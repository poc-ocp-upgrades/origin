package image

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
	"github.com/openshift/origin/pkg/image/apis/image/reference"
)

type ImageList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]Image
}

const (
	ManagedByOpenShiftAnnotation					= "openshift.io/image.managed"
	DockerImageRepositoryCheckAnnotation				= "openshift.io/image.dockerRepositoryCheck"
	InsecureRepositoryAnnotation					= "openshift.io/image.insecureRepository"
	ExcludeImageSecretAnnotation					= "openshift.io/image.excludeSecret"
	DockerImageLayersOrderAnnotation				= "image.openshift.io/dockerLayersOrder"
	DockerImageLayersOrderAscending					= "ascending"
	DockerImageLayersOrderDescending				= "descending"
	ImporterPreferArchAnnotation					= "importer.image.openshift.io/prefer-arch"
	ImporterPreferOSAnnotation					= "importer.image.openshift.io/prefer-os"
	ImageManifestBlobStoredAnnotation				= "image.openshift.io/manifestBlobStored"
	DefaultImageTag							= "latest"
	ResourceImageStreams			core.ResourceName	= "openshift.io/imagestreams"
	ResourceImageStreamImages		core.ResourceName	= "openshift.io/images"
	ResourceImageStreamTags			core.ResourceName	= "openshift.io/image-tags"
	LimitTypeImage				core.LimitType		= "openshift.io/Image"
	LimitTypeImageStream			core.LimitType		= "openshift.io/ImageStream"
)

type Image struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	DockerImageReference		string
	DockerImageMetadata		DockerImage
	DockerImageMetadataVersion	string
	DockerImageManifest		string
	DockerImageLayers		[]ImageLayer
	Signatures			[]ImageSignature
	DockerImageSignatures		[][]byte
	DockerImageManifestMediaType	string
	DockerImageConfig		string
}
type ImageLayer struct {
	Name		string
	LayerSize	int64
	MediaType	string
}

const (
	ImageSignatureTypeAtomicImageV1 string = "AtomicImageV1"
)

type ImageSignature struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Type		string
	Content		[]byte
	Conditions	[]SignatureCondition
	ImageIdentity	string
	SignedClaims	map[string]string
	Created		*metav1.Time
	IssuedBy	*SignatureIssuer
	IssuedTo	*SignatureSubject
}

const (
	SignatureTrusted	= "Trusted"
	SignatureForImage	= "ForImage"
	SignatureExpired	= "Expired"
	SignatureRevoked	= "Revoked"
)

type SignatureConditionType string
type SignatureCondition struct {
	Type			SignatureConditionType
	Status			core.ConditionStatus
	LastProbeTime		metav1.Time
	LastTransitionTime	metav1.Time
	Reason			string
	Message			string
}
type SignatureGenericEntity struct {
	Organization	string
	CommonName	string
}
type SignatureIssuer struct{ SignatureGenericEntity }
type SignatureSubject struct {
	SignatureGenericEntity
	PublicKeyID	string
}
type ImageStreamList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ImageStream
}
type ImageStream struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ImageStreamSpec
	Status	ImageStreamStatus
}
type ImageStreamSpec struct {
	LookupPolicy		ImageLookupPolicy
	DockerImageRepository	string
	Tags			map[string]TagReference
}
type ImageLookupPolicy struct{ Local bool }
type TagReference struct {
	Name		string
	Annotations	map[string]string
	From		*core.ObjectReference
	Reference	bool
	Generation	*int64
	ImportPolicy	TagImportPolicy
	ReferencePolicy	TagReferencePolicy
}
type TagImportPolicy struct {
	Insecure	bool
	Scheduled	bool
}
type TagReferencePolicyType string

const (
	SourceTagReferencePolicy	TagReferencePolicyType	= "Source"
	LocalTagReferencePolicy		TagReferencePolicyType	= "Local"
)

type TagReferencePolicy struct{ Type TagReferencePolicyType }
type ImageStreamStatus struct {
	DockerImageRepository		string
	PublicDockerImageRepository	string
	Tags				map[string]TagEventList
}
type TagEventList struct {
	Items		[]TagEvent
	Conditions	[]TagEventCondition
}
type TagEvent struct {
	Created			metav1.Time
	DockerImageReference	string
	Image			string
	Generation		int64
}
type TagEventConditionType string

const (
	ImportSuccess TagEventConditionType = "ImportSuccess"
)

type TagEventCondition struct {
	Type			TagEventConditionType
	Status			core.ConditionStatus
	LastTransitionTime	metav1.Time
	Reason			string
	Message			string
	Generation		int64
}
type ImageStreamMapping struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	DockerImageRepository	string
	Image			Image
	Tag			string
}
type ImageStreamTag struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Tag		*TagReference
	Generation	int64
	Conditions	[]TagEventCondition
	LookupPolicy	ImageLookupPolicy
	Image		Image
}
type ImageStreamTagList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ImageStreamTag
}
type ImageStreamImage struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Image	Image
}
type DockerImageReference = reference.DockerImageReference
type ImageStreamLayers struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Blobs	map[string]ImageLayerData
	Images	map[string]ImageBlobReferences
}
type ImageBlobReferences struct {
	ImageMissing	bool
	Layers		[]string
	Config		*string
}
type ImageLayerData struct {
	LayerSize	*int64
	MediaType	string
}
type ImageStreamImport struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ImageStreamImportSpec
	Status	ImageStreamImportStatus
}
type ImageStreamImportSpec struct {
	Import		bool
	Repository	*RepositoryImportSpec
	Images		[]ImageImportSpec
}
type ImageStreamImportStatus struct {
	Import		*ImageStream
	Repository	*RepositoryImportStatus
	Images		[]ImageImportStatus
}
type RepositoryImportSpec struct {
	From		core.ObjectReference
	ImportPolicy	TagImportPolicy
	ReferencePolicy	TagReferencePolicy
	IncludeManifest	bool
}
type RepositoryImportStatus struct {
	Status		metav1.Status
	Images		[]ImageImportStatus
	AdditionalTags	[]string
}
type ImageImportSpec struct {
	From		core.ObjectReference
	To		*core.LocalObjectReference
	ImportPolicy	TagImportPolicy
	ReferencePolicy	TagReferencePolicy
	IncludeManifest	bool
}
type ImageImportStatus struct {
	Tag	string
	Status	metav1.Status
	Image	*Image
}
