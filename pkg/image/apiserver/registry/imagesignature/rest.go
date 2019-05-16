package imagesignature

import (
	"context"
	"errors"
	goformat "fmt"
	imagegroup "github.com/openshift/api/image"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imageclient "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct{ imageClient imageclient.ImagesGetter }

var _ rest.Creater = &REST{}
var _ rest.GracefulDeleter = &REST{}
var _ rest.Scoper = &REST{}

func NewREST(imageClient imageclient.ImagesGetter) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{imageClient: imageClient}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &imageapi.ImageSignature{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	signature := obj.(*imageapi.ImageSignature)
	if err := rest.BeforeCreate(Strategy, ctx, obj); err != nil {
		return nil, err
	}
	if createValidation != nil {
		if err := createValidation(obj.DeepCopyObject()); err != nil {
			return nil, err
		}
	}
	imageName, _, err := imageapi.SplitImageSignatureName(signature.Name)
	if err != nil {
		return nil, kapierrors.NewBadRequest(err.Error())
	}
	image, err := r.imageClient.Images().Get(imageName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if byName, byContent := imageapi.IndexOfImageSignatureByName(image.Signatures, signature.Name), imageapi.IndexOfImageSignature(image.Signatures, signature.Type, signature.Content); byName >= 0 || byContent >= 0 {
		return nil, kapierrors.NewAlreadyExists(imagegroup.Resource("imageSignatures"), signature.Name)
	}
	image.Signatures = append(image.Signatures, *signature)
	image, err = r.imageClient.Images().Update(image)
	if err != nil {
		return nil, err
	}
	byName := imageapi.IndexOfImageSignatureByName(image.Signatures, signature.Name)
	if byName < 0 {
		return nil, kapierrors.NewInternalError(errors.New("failed to store given signature"))
	}
	return &image.Signatures[byName], nil
}
func (r *REST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	imageName, _, err := imageapi.SplitImageSignatureName(name)
	if err != nil {
		return nil, false, kapierrors.NewBadRequest("ImageSignatures must be accessed with <imageName>@<signatureName>")
	}
	image, err := r.imageClient.Images().Get(imageName, metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	index := imageapi.IndexOfImageSignatureByName(image.Signatures, name)
	if index < 0 {
		return nil, false, kapierrors.NewNotFound(imagegroup.Resource("imageSignatures"), name)
	}
	size := len(image.Signatures)
	copy(image.Signatures[index:size-1], image.Signatures[index+1:size])
	image.Signatures = image.Signatures[0 : size-1]
	if _, err := r.imageClient.Images().Update(image); err != nil {
		return nil, false, err
	}
	return &metav1.Status{Status: metav1.StatusSuccess}, true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
