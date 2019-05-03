package importer

import (
	"fmt"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/docker/distribution/manifest/schema1"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
	"github.com/docker/distribution/registry/api/errcode"
	v2 "github.com/docker/distribution/registry/api/v2"
	godigest "github.com/opencontainers/go-digest"
	"github.com/openshift/api/image"
	"github.com/openshift/origin/pkg/api/legacy"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/importer/dockerv1client"
	"github.com/openshift/origin/pkg/image/util"
	gocontext "golang.org/x/net/context"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/klog"
	"net/url"
	"runtime"
	"strings"
)

const ContextKeyV1RegistryClient = "v1-registry-client"

type Interface interface {
	Import(ctx gocontext.Context, isi *imageapi.ImageStreamImport, stream *imageapi.ImageStream) error
}
type RepositoryRetriever interface {
	Repository(ctx gocontext.Context, registry *url.URL, repoName string, insecure bool) (distribution.Repository, error)
}
type ImageStreamImporter struct {
	maximumTagsPerRepo      int
	retriever               RepositoryRetriever
	limiter                 flowcontrol.RateLimiter
	digestToRepositoryCache map[gocontext.Context]map[manifestKey]*imageapi.Image
	digestToLayerSizeCache  *ImageStreamLayerCache
}

func NewImageStreamImporter(retriever RepositoryRetriever, maximumTagsPerRepo int, limiter flowcontrol.RateLimiter, cache *ImageStreamLayerCache) *ImageStreamImporter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if limiter == nil {
		limiter = flowcontrol.NewFakeAlwaysRateLimiter()
	}
	if cache == nil {
		klog.V(5).Infof("the global layer cache is disabled")
	}
	return &ImageStreamImporter{maximumTagsPerRepo: maximumTagsPerRepo, retriever: retriever, limiter: limiter, digestToRepositoryCache: make(map[gocontext.Context]map[manifestKey]*imageapi.Image), digestToLayerSizeCache: cache}
}
func (i *ImageStreamImporter) Import(ctx gocontext.Context, isi *imageapi.ImageStreamImport, stream *imageapi.ImageStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if i.digestToLayerSizeCache == nil {
		cache, err := NewImageStreamLayerCache(DefaultImageStreamLayerCacheSize)
		if err != nil {
			return err
		}
		i.digestToLayerSizeCache = &cache
	}
	if _, ok := i.digestToRepositoryCache[ctx]; !ok {
		i.digestToRepositoryCache[ctx] = make(map[manifestKey]*imageapi.Image)
	}
	i.importImages(ctx, i.retriever, isi, stream, i.limiter)
	i.importFromRepository(ctx, i.retriever, isi, i.maximumTagsPerRepo, i.limiter)
	return nil
}
func (i *ImageStreamImporter) importImages(ctx gocontext.Context, retriever RepositoryRetriever, isi *imageapi.ImageStreamImport, stream *imageapi.ImageStream, limiter flowcontrol.RateLimiter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tags := make(map[manifestKey][]int)
	ids := make(map[manifestKey][]int)
	repositories := make(map[repositoryKey]*importRepository)
	cache := i.digestToRepositoryCache[ctx]
	isi.Status.Images = make([]imageapi.ImageImportStatus, len(isi.Spec.Images))
	for i := range isi.Spec.Images {
		spec := &isi.Spec.Images[i]
		from := spec.From
		if from.Kind != "DockerImage" {
			continue
		}
		var (
			err error
			ref imageapi.DockerImageReference
		)
		if from.Name != "*" {
			ref, err = imageapi.ParseDockerImageReference(from.Name)
			if err != nil {
				isi.Status.Images[i].Status = invalidStatus("", field.Invalid(field.NewPath("from", "name"), from.Name, fmt.Sprintf("invalid name: %v", err)))
				continue
			}
		} else {
			ref = imageapi.DockerImageReference{Name: from.Name}
		}
		defaultRef := ref.DockerClientDefaults()
		repoName := defaultRef.RepositoryName()
		registryURL := defaultRef.RegistryURL()
		key := repositoryKey{url: *registryURL, name: repoName}
		repo, ok := repositories[key]
		if !ok {
			repo = &importRepository{Ref: ref, Registry: &key.url, Name: key.name, Insecure: spec.ImportPolicy.Insecure}
			repositories[key] = repo
		}
		if len(defaultRef.ID) > 0 {
			id := manifestKey{repositoryKey: key}
			id.value = defaultRef.ID
			ids[id] = append(ids[id], i)
			if len(ids[id]) == 1 {
				repo.Digests = append(repo.Digests, importDigest{Name: defaultRef.ID, Image: cache[id]})
			}
		} else {
			var toName string
			if spec.To != nil {
				toName = spec.To.Name
			} else {
				toName = defaultRef.Tag
			}
			tagReference := stream.Spec.Tags[toName]
			preferArch := tagReference.Annotations[imageapi.ImporterPreferArchAnnotation]
			preferOS := tagReference.Annotations[imageapi.ImporterPreferOSAnnotation]
			tag := manifestKey{repositoryKey: key, preferArch: preferArch, preferOS: preferOS}
			tag.value = defaultRef.Tag
			tags[tag] = append(tags[tag], i)
			if len(tags[tag]) == 1 {
				repo.Tags = append(repo.Tags, importTag{Name: defaultRef.Tag, PreferArch: preferArch, PreferOS: preferOS, Image: cache[tag]})
			}
		}
	}
	for key, repo := range repositories {
		i.importRepositoryFromDocker(ctx, retriever, repo, limiter)
		for _, tag := range repo.Tags {
			j := manifestKey{repositoryKey: key, preferArch: tag.PreferArch, preferOS: tag.PreferOS}
			j.value = tag.Name
			if tag.Image != nil {
				cache[j] = tag.Image
			}
			for _, index := range tags[j] {
				if tag.Err != nil {
					setImageImportStatus(isi, index, tag.Name, tag.Err)
					continue
				}
				copied := *tag.Image
				image := &isi.Status.Images[index]
				ref := repo.Ref
				ref.Tag, ref.ID = tag.Name, copied.Name
				copied.DockerImageReference = ref.MostSpecific().Exact()
				image.Tag = tag.Name
				image.Image = &copied
				image.Status.Status = metav1.StatusSuccess
			}
		}
		for _, digest := range repo.Digests {
			j := manifestKey{repositoryKey: key}
			j.value = digest.Name
			if digest.Image != nil {
				cache[j] = digest.Image
			}
			for _, index := range ids[j] {
				if digest.Err != nil {
					setImageImportStatus(isi, index, "", digest.Err)
					continue
				}
				image := &isi.Status.Images[index]
				copied := *digest.Image
				ref := repo.Ref
				ref.Tag, ref.ID = "", copied.Name
				copied.DockerImageReference = ref.MostSpecific().Exact()
				image.Image = &copied
				image.Status.Status = metav1.StatusSuccess
			}
		}
	}
}
func (i *ImageStreamImporter) importFromRepository(ctx gocontext.Context, retriever RepositoryRetriever, isi *imageapi.ImageStreamImport, maximumTags int, limiter flowcontrol.RateLimiter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isi.Spec.Repository == nil {
		return
	}
	cache := i.digestToRepositoryCache[ctx]
	isi.Status.Repository = &imageapi.RepositoryImportStatus{}
	status := isi.Status.Repository
	spec := isi.Spec.Repository
	from := spec.From
	if from.Kind != "DockerImage" {
		return
	}
	var (
		err error
		ref imageapi.DockerImageReference
	)
	if from.Name != "*" {
		ref, err = imageapi.ParseDockerImageReference(from.Name)
		if err != nil {
			status.Status = invalidStatus("", field.Invalid(field.NewPath("from", "name"), from.Name, fmt.Sprintf("invalid name: %v", err)))
			return
		}
	} else {
		ref = imageapi.DockerImageReference{Name: from.Name}
	}
	defaultRef := ref.DockerClientDefaults()
	repoName := defaultRef.RepositoryName()
	registryURL := defaultRef.RegistryURL()
	key := repositoryKey{url: *registryURL, name: repoName}
	repo := &importRepository{Ref: ref, Registry: &key.url, Name: key.name, Insecure: spec.ImportPolicy.Insecure, MaximumTags: maximumTags}
	i.importRepositoryFromDocker(ctx, retriever, repo, limiter)
	if repo.Err != nil {
		status.Status = imageImportStatus(repo.Err, "", "repository")
		return
	}
	additional := []string{}
	tagKey := manifestKey{repositoryKey: key}
	for _, s := range repo.AdditionalTags {
		tagKey.value = s
		if image, ok := cache[tagKey]; ok {
			repo.Tags = append(repo.Tags, importTag{Name: s, Image: image})
		} else {
			additional = append(additional, s)
		}
	}
	status.AdditionalTags = additional
	failures := 0
	status.Status.Status = metav1.StatusSuccess
	status.Images = make([]imageapi.ImageImportStatus, len(repo.Tags))
	for i, tag := range repo.Tags {
		status.Images[i].Tag = tag.Name
		if tag.Err != nil {
			failures++
			status.Images[i].Status = imageImportStatus(tag.Err, "", "repository")
			continue
		}
		status.Images[i].Status.Status = metav1.StatusSuccess
		copied := *tag.Image
		ref.Tag, ref.ID = tag.Name, copied.Name
		copied.DockerImageReference = ref.MostSpecific().Exact()
		status.Images[i].Image = &copied
	}
	if failures > 0 {
		status.Status.Status = metav1.StatusFailure
		status.Status.Reason = metav1.StatusReason("ImportFailed")
		switch failures {
		case 1:
			status.Status.Message = "one of the images from this repository failed to import"
		default:
			status.Status.Message = fmt.Sprintf("%d of the images from this repository failed to import", failures)
		}
	}
}
func applyErrorToRepository(repository *importRepository, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	repository.Err = err
	for i := range repository.Tags {
		repository.Tags[i].Err = err
	}
	for i := range repository.Digests {
		repository.Digests[i].Err = err
	}
}
func formatRepositoryError(ref imageapi.DockerImageReference, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case isDockerError(err, v2.ErrorCodeManifestUnknown):
		err = kapierrors.NewNotFound(image.Resource("dockerimage"), ref.Exact())
	case isDockerError(err, errcode.ErrorCodeUnauthorized):
		err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q", ref.Exact()))
	case strings.HasSuffix(err.Error(), "no basic auth credentials"):
		err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q", ref.Exact()))
	case strings.HasSuffix(err.Error(), "incorrect username or password"):
		err = kapierrors.NewUnauthorized(fmt.Sprintf("incorrect username or password for image %q", ref.Exact()))
	}
	return err
}
func (isi *ImageStreamImporter) calculateImageSize(ctx gocontext.Context, bs distribution.BlobStore, image *imageapi.Image) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	blobSet := sets.NewString()
	size := int64(0)
	for i := range image.DockerImageLayers {
		layer := &image.DockerImageLayers[i]
		if blobSet.Has(layer.Name) {
			continue
		}
		blobSet.Insert(layer.Name)
		if layerSize, ok := isi.digestToLayerSizeCache.Get(layer.Name); ok {
			layerSize := layerSize.(int64)
			layer.LayerSize = layerSize
			size += layerSize
			continue
		}
		desc, err := bs.Stat(ctx, godigest.Digest(layer.Name))
		if err != nil {
			return err
		}
		isi.digestToLayerSizeCache.Add(layer.Name, desc.Size)
		layer.LayerSize = desc.Size
		size += desc.Size
	}
	if len(image.DockerImageConfig) > 0 && !blobSet.Has(image.DockerImageMetadata.ID) {
		blobSet.Insert(image.DockerImageMetadata.ID)
		size += int64(len(image.DockerImageConfig))
	}
	image.DockerImageMetadata.Size = size
	return nil
}
func manifestFromManifestList(ctx gocontext.Context, manifestList *manifestlist.DeserializedManifestList, ref imageapi.DockerImageReference, s distribution.ManifestService, preferArch, preferOS string) (distribution.Manifest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(manifestList.Manifests) == 0 {
		return nil, fmt.Errorf("no manifests in manifest list %s", ref.Exact())
	}
	if preferArch == "" {
		preferArch = runtime.GOARCH
	}
	if preferOS == "" {
		preferOS = runtime.GOOS
	}
	var manifestDigest godigest.Digest
	for _, manifestDescriptor := range manifestList.Manifests {
		if manifestDescriptor.Platform.Architecture == preferArch && manifestDescriptor.Platform.OS == preferOS {
			manifestDigest = manifestDescriptor.Digest
			break
		}
	}
	if manifestDigest == "" {
		klog.V(5).Infof("unable to find %s/%s manifest in manifest list %s, doing conservative fail by switching to the first one: %#+v", preferOS, preferArch, ref.Exact(), manifestList.Manifests[0])
		manifestDigest = manifestList.Manifests[0].Digest
	}
	manifest, err := s.Get(ctx, manifestDigest)
	if err != nil {
		klog.V(5).Infof("unable to get %s/%s manifest by digest %q for image %s: %#v", preferOS, preferArch, manifestDigest, ref.Exact(), err)
		return nil, formatRepositoryError(ref, err)
	}
	return manifest, err
}
func (isi *ImageStreamImporter) importManifest(ctx gocontext.Context, manifest distribution.Manifest, ref imageapi.DockerImageReference, d godigest.Digest, s distribution.ManifestService, b distribution.BlobStore, preferArch, preferOS string) (image *imageapi.Image, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if manifestList, ok := manifest.(*manifestlist.DeserializedManifestList); ok {
		manifest, err = manifestFromManifestList(ctx, manifestList, ref, s, preferArch, preferOS)
		if err != nil {
			return nil, err
		}
	}
	if signedManifest, isSchema1 := manifest.(*schema1.SignedManifest); isSchema1 {
		image, err = schema1ToImage(signedManifest, d)
	} else if deserializedManifest, isSchema2 := manifest.(*schema2.DeserializedManifest); isSchema2 {
		imageConfig, getImportConfigErr := b.Get(ctx, deserializedManifest.Config.Digest)
		if getImportConfigErr != nil {
			klog.V(5).Infof("unable to get image config by digest %q for image %s: %#v", d, ref.Exact(), getImportConfigErr)
			return image, formatRepositoryError(ref, getImportConfigErr)
		}
		image, err = schema2ToImage(deserializedManifest, imageConfig, d)
	} else {
		err = fmt.Errorf("unsupported image manifest type: %T", manifest)
		klog.V(5).Info(err)
	}
	if err != nil {
		return
	}
	if err := util.InternalImageWithMetadata(image); err != nil {
		return image, err
	}
	if image.DockerImageMetadata.Size == 0 {
		if err := isi.calculateImageSize(ctx, b, image); err != nil {
			return image, err
		}
	}
	return
}
func (isi *ImageStreamImporter) importRepositoryFromDocker(ctx gocontext.Context, retriever RepositoryRetriever, repository *importRepository, limiter flowcontrol.RateLimiter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(5).Infof("importing remote Docker repository registry=%s repository=%s insecure=%t", repository.Registry, repository.Name, repository.Insecure)
	repo, err := retriever.Repository(ctx, repository.Registry, repository.Name, repository.Insecure)
	if err != nil {
		klog.V(5).Infof("unable to access repository %#v: %#v", repository, err)
		switch {
		case err == reference.ErrReferenceInvalidFormat:
			err = field.Invalid(field.NewPath("from", "name"), repository.Name, "the provided repository name is not valid")
		case isDockerError(err, v2.ErrorCodeNameUnknown):
			err = kapierrors.NewNotFound(image.Resource("dockerimage"), repository.Ref.Exact())
		case isDockerError(err, errcode.ErrorCodeUnauthorized):
			err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q", repository.Ref.Exact()))
		case strings.Contains(err.Error(), "tls: oversized record received with length") && !repository.Insecure:
			err = kapierrors.NewBadRequest("this repository is HTTP only and requires the insecure flag to import")
		case strings.HasSuffix(err.Error(), "no basic auth credentials"):
			err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q and did not have credentials to the repository", repository.Ref.Exact()))
		case strings.HasSuffix(err.Error(), "does not support v2 API"):
			importRepositoryFromDockerV1(ctx, repository, limiter)
			return
		}
		applyErrorToRepository(repository, err)
		return
	}
	s, err := repo.Manifests(ctx)
	if err != nil {
		klog.V(5).Infof("unable to access manifests for repository %#v: %#v", repository, err)
		switch {
		case isDockerError(err, v2.ErrorCodeNameUnknown):
			err = kapierrors.NewNotFound(image.Resource("dockerimage"), repository.Ref.Exact())
		case isDockerError(err, errcode.ErrorCodeUnauthorized):
			err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q", repository.Ref.Exact()))
		case strings.HasSuffix(err.Error(), "no basic auth credentials"):
			err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q and did not have credentials to the repository", repository.Ref.Exact()))
		}
		applyErrorToRepository(repository, err)
		return
	}
	b := repo.Blobs(ctx)
	if count := repository.MaximumTags; count > 0 || count == -1 {
		tags, err := repo.Tags(ctx).All(ctx)
		if err != nil {
			klog.V(5).Infof("unable to access tags for repository %#v: %#v", repository, err)
			switch {
			case isDockerError(err, v2.ErrorCodeNameUnknown):
				err = kapierrors.NewNotFound(image.Resource("dockerimage"), repository.Ref.Exact())
			case isDockerError(err, errcode.ErrorCodeUnauthorized):
				err = kapierrors.NewUnauthorized(fmt.Sprintf("you may not have access to the Docker image %q", repository.Ref.Exact()))
			}
			repository.Err = err
			return
		}
		set := sets.NewString(tags...)
		if set.Has("") {
			set.Delete("")
			set.Insert(imageapi.DefaultImageTag)
		}
		tags = set.List()
		imageapi.PrioritizeTags(tags)
		for _, s := range tags {
			if count <= 0 && repository.MaximumTags != -1 {
				repository.AdditionalTags = append(repository.AdditionalTags, s)
				continue
			}
			count--
			repository.Tags = append(repository.Tags, importTag{Name: s})
		}
	}
	for i := range repository.Digests {
		importDigest := &repository.Digests[i]
		if importDigest.Err != nil || importDigest.Image != nil {
			continue
		}
		d, err := godigest.Parse(importDigest.Name)
		if err != nil {
			importDigest.Err = err
			continue
		}
		ref := repository.Ref
		ref.Tag = ""
		ref.ID = string(d)
		limiter.Accept()
		manifest, err := s.Get(ctx, d)
		if err != nil {
			klog.V(5).Infof("unable to get manifest by digest %q for image %s: %#v", d, ref.Exact(), err)
			importDigest.Err = formatRepositoryError(ref, err)
			continue
		}
		importDigest.Image, importDigest.Err = isi.importManifest(ctx, manifest, ref, d, s, b, "", "")
	}
	for i := range repository.Tags {
		importTag := &repository.Tags[i]
		if importTag.Err != nil || importTag.Image != nil {
			continue
		}
		ref := repository.Ref
		ref.Tag = importTag.Name
		ref.ID = ""
		limiter.Accept()
		manifest, err := s.Get(ctx, "", distribution.WithTag(importTag.Name))
		if err != nil {
			klog.V(5).Infof("unable to get manifest by tag %q for image %s: %#v", importTag.Name, ref.Exact(), err)
			importTag.Err = formatRepositoryError(ref, err)
			continue
		}
		importTag.Image, importTag.Err = isi.importManifest(ctx, manifest, ref, "", s, b, importTag.PreferArch, importTag.PreferOS)
	}
}
func importRepositoryFromDockerV1(ctx gocontext.Context, repository *importRepository, limiter flowcontrol.RateLimiter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value := ctx.Value(ContextKeyV1RegistryClient)
	if value == nil {
		err := kapierrors.NewForbidden(image.Resource(""), "", fmt.Errorf("registry %q does not support the v2 Registry API", repository.Registry.Host))
		err.ErrStatus.Reason = "NotV2Registry"
		applyErrorToRepository(repository, err)
		return
	}
	client, ok := value.(dockerv1client.Client)
	if !ok {
		err := kapierrors.NewForbidden(image.Resource(""), "", fmt.Errorf("registry %q does not support the v2 Registry API", repository.Registry.Host))
		err.ErrStatus.Reason = "NotV2Registry"
		return
	}
	conn, err := client.Connect(repository.Registry.Host, repository.Insecure)
	if err != nil {
		applyErrorToRepository(repository, err)
		return
	}
	if count := repository.MaximumTags; count > 0 || count == -1 {
		tagMap, err := conn.ImageTags(repository.Ref.Namespace, repository.Ref.Name)
		if err != nil {
			repository.Err = err
			return
		}
		tags := make([]string, 0, len(tagMap))
		for tag := range tagMap {
			tags = append(tags, tag)
		}
		set := sets.NewString(tags...)
		if set.Has("") {
			set.Delete("")
			set.Insert(imageapi.DefaultImageTag)
		}
		tags = set.List()
		imageapi.PrioritizeTags(tags)
		for _, s := range tags {
			if count <= 0 && repository.MaximumTags != -1 {
				repository.AdditionalTags = append(repository.AdditionalTags, s)
				continue
			}
			count--
			repository.Tags = append(repository.Tags, importTag{Name: s})
		}
	}
	for i := range repository.Digests {
		importDigest := &repository.Digests[i]
		if importDigest.Err != nil || importDigest.Image != nil {
			continue
		}
		limiter.Accept()
		image, err := conn.ImageByID(repository.Ref.Namespace, repository.Ref.Name, importDigest.Name)
		if err != nil {
			importDigest.Err = err
			continue
		}
		importDigest.Image, err = schema0ToImage(image)
		if err != nil {
			importDigest.Err = err
			continue
		}
	}
	for i := range repository.Tags {
		importTag := &repository.Tags[i]
		if importTag.Err != nil || importTag.Image != nil {
			continue
		}
		limiter.Accept()
		image, err := conn.ImageByTag(repository.Ref.Namespace, repository.Ref.Name, importTag.Name)
		if err != nil {
			importTag.Err = err
			continue
		}
		importTag.Image, err = schema0ToImage(image)
		if err != nil {
			importTag.Err = err
			continue
		}
	}
}

type importTag struct {
	Name       string
	PreferArch string
	PreferOS   string
	Image      *imageapi.Image
	Err        error
}
type importDigest struct {
	Name  string
	Image *imageapi.Image
	Err   error
}
type importRepository struct {
	Ref            imageapi.DockerImageReference
	Registry       *url.URL
	Name           string
	Insecure       bool
	Tags           []importTag
	Digests        []importDigest
	MaximumTags    int
	AdditionalTags []string
	Err            error
}
type repositoryKey struct {
	url  url.URL
	name string
}
type manifestKey struct {
	repositoryKey
	value      string
	preferArch string
	preferOS   string
}

func imageImportStatus(err error, kind, position string) metav1.Status {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch t := err.(type) {
	case kapierrors.APIStatus:
		return t.Status()
	case *field.Error:
		return kapierrors.NewInvalid(image.Kind(kind), position, field.ErrorList{t}).ErrStatus
	default:
		return kapierrors.NewInternalError(err).ErrStatus
	}
}
func setImageImportStatus(images *imageapi.ImageStreamImport, i int, tag string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	images.Status.Images[i].Tag = tag
	images.Status.Images[i].Status = imageImportStatus(err, "", "")
}
func invalidStatus(position string, errs ...*field.Error) metav1.Status {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kapierrors.NewInvalid(legacy.Kind(""), position, errs).ErrStatus
}
