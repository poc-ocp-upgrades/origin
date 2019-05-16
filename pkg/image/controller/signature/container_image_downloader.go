package signature

import (
	"context"
	"crypto/sha256"
	"fmt"
	goformat "fmt"
	"github.com/containers/image/docker"
	imagev1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type containerImageSignatureDownloader struct {
	ctx     context.Context
	timeout time.Duration
}

func NewContainerImageSignatureDownloader(ctx context.Context, timeout time.Duration) SignatureDownloader {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &containerImageSignatureDownloader{ctx: ctx, timeout: timeout}
}

type GetSignaturesError struct{ error }

func (s *containerImageSignatureDownloader) DownloadImageSignatures(image *imagev1.Image) ([]imagev1.ImageSignature, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	reference, err := docker.ParseReference("//" + image.DockerImageReference)
	if err != nil {
		return nil, err
	}
	source, err := reference.NewImageSource(nil, nil)
	if err != nil {
		klog.V(4).Infof("Failed to get %q: %v", image.DockerImageReference, err)
		return []imagev1.ImageSignature{}, nil
	}
	defer source.Close()
	ctx, cancel := context.WithTimeout(s.ctx, s.timeout)
	defer cancel()
	signatures, err := source.GetSignatures(ctx)
	if err != nil {
		klog.V(4).Infof("Failed to get signatures for %v due to: %v", source.Reference(), err)
		return []imagev1.ImageSignature{}, GetSignaturesError{err}
	}
	ret := []imagev1.ImageSignature{}
	for _, blob := range signatures {
		sig := imagev1.ImageSignature{Type: imageapi.ImageSignatureTypeAtomicImageV1}
		sig.Name = imageapi.JoinImageStreamImage(image.Name, fmt.Sprintf("%x", sha256.Sum256(blob)))
		sig.Content = blob
		sig.CreationTimestamp = metav1.Now()
		ret = append(ret, sig)
	}
	return ret, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
