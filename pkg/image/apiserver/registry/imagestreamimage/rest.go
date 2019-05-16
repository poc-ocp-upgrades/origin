package imagestreamimage

import (
	"context"
	goformat "fmt"
	imagegroup "github.com/openshift/api/image"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apiserver/registry/image"
	"github.com/openshift/origin/pkg/image/apiserver/registry/imagestream"
	"github.com/openshift/origin/pkg/image/util"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type REST struct {
	imageRegistry       image.Registry
	imageStreamRegistry imagestream.Registry
	rest.TableConvertor
}

var _ rest.Getter = &REST{}
var _ rest.ShortNamesProvider = &REST{}
var _ rest.Scoper = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"isimage"}
}
func NewREST(imageRegistry image.Registry, imageStreamRegistry imagestream.Registry) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{imageRegistry, imageStreamRegistry, printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)}}
}
func (r *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &imageapi.ImageStreamImage{}
}
func (s *REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}
func parseNameAndID(input string) (name string, id string, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, id, err = imageapi.ParseImageStreamImageName(input)
	if err != nil {
		err = errors.NewBadRequest("ImageStreamImages must be retrieved with <name>@<id>")
	}
	return
}
func (r *REST) Get(ctx context.Context, id string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	name, imageID, err := parseNameAndID(id)
	if err != nil {
		return nil, err
	}
	repo, err := r.imageStreamRegistry.GetImageStream(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if repo.Status.Tags == nil {
		return nil, errors.NewNotFound(imagegroup.Resource("imagestreamimage"), id)
	}
	event, err := imageapi.ResolveImageID(repo, imageID)
	if err != nil {
		return nil, err
	}
	imageName := event.Image
	image, err := r.imageRegistry.GetImage(ctx, imageName, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if err := util.InternalImageWithMetadata(image); err != nil {
		return nil, err
	}
	image.DockerImageManifest = ""
	image.DockerImageConfig = ""
	isi := imageapi.ImageStreamImage{ObjectMeta: metav1.ObjectMeta{Namespace: apirequest.NamespaceValue(ctx), Name: imageapi.JoinImageStreamImage(name, imageID), CreationTimestamp: image.ObjectMeta.CreationTimestamp, Annotations: repo.Annotations}, Image: *image}
	return &isi, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
