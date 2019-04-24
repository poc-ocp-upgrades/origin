package controller

import (
	"fmt"
	"sync"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	imagev1 "github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	metrics "github.com/openshift/origin/pkg/image/metrics/prometheus"
	imageutil "github.com/openshift/origin/pkg/image/util"
)

const reasonUnknown = "Unknown"
const reasonInvalidImageReference = "InvalidImageReference"

type ImportMetricCounter struct {
	counterMutex		sync.Mutex
	importSuccessCounts	metrics.ImportSuccessCounts
	importErrorCounts	metrics.ImportErrorCounts
}

func NewImportMetricCounter() *ImportMetricCounter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ImportMetricCounter{importSuccessCounts: make(metrics.ImportSuccessCounts), importErrorCounts: make(metrics.ImportErrorCounts)}
}
func (c *ImportMetricCounter) Increment(isi *imagev1.ImageStreamImport, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isi == nil {
		if err == nil {
			return
		}
		c.counterMutex.Lock()
		defer c.counterMutex.Unlock()
		info := defaultErrorInfoReason(&metrics.ImportErrorInfo{}, err)
		c.importErrorCounts[*info]++
		return
	}
	c.countRepositoryImport(isi, err)
	if len(isi.Status.Images) == 0 {
		return
	}
	c.counterMutex.Lock()
	defer c.counterMutex.Unlock()
	enumerateIsImportStatuses(isi, func(info *metrics.ImportErrorInfo) {
		if len(info.Reason) == 0 {
			c.importSuccessCounts[info.Registry]++
		} else {
			c.importErrorCounts[*defaultErrorInfoReason(info, err)]++
		}
	})
}
func (c *ImportMetricCounter) countRepositoryImport(isi *imagev1.ImageStreamImport, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errInfo := getIsImportRepositoryInfo(isi)
	if errInfo == nil {
		return
	}
	c.counterMutex.Lock()
	defer c.counterMutex.Unlock()
	if len(errInfo.Reason) == 0 {
		c.importSuccessCounts[errInfo.Registry]++
	} else {
		c.importErrorCounts[*defaultErrorInfoReason(errInfo, err)]++
	}
}
func (c *ImportMetricCounter) Collect() (metrics.ImportSuccessCounts, metrics.ImportErrorCounts, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.counterMutex.Lock()
	defer c.counterMutex.Unlock()
	success := metrics.ImportSuccessCounts{}
	for registry, count := range c.importSuccessCounts {
		success[registry] = count
	}
	failures := metrics.ImportErrorCounts{}
	for info, count := range c.importErrorCounts {
		failures[info] = count
	}
	return success, failures, nil
}
func getIsImportRepositoryInfo(isi *imagev1.ImageStreamImport) *metrics.ImportErrorInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isi.Status.Repository == nil || isi.Spec.Repository == nil {
		return nil
	}
	ref := isi.Spec.Repository.From
	if ref.Kind != "DockerImage" {
		return nil
	}
	imgRef, err := imageapi.ParseDockerImageReference(ref.Name)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to parse isi.spec.repository.from.name %q: %v", ref.Name, err))
		return nil
	}
	info := mkImportInfo(imgRef.DockerClientDefaults().Registry, &isi.Status.Repository.Status)
	return &info
}
func enumerateIsImportStatuses(isi *imagev1.ImageStreamImport, cb func(*metrics.ImportErrorInfo)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(isi.Status.Images) == 0 {
		return
	}
	for i, status := range isi.Status.Images {
		var registry string
		imgRef, err := getImageDockerReferenceForImage(isi, i)
		if err != nil {
			utilruntime.HandleError(err)
		} else {
			if imgRef == nil {
				continue
			}
			imageutil.SetDockerClientDefaults(imgRef)
			registry = imgRef.Registry
		}
		info := mkImportInfo(registry, &status.Status)
		if err != nil {
			info.Reason = reasonInvalidImageReference
		}
		cb(&info)
	}
}
func getImageDockerReferenceForImage(isi *imagev1.ImageStreamImport, index int) (*imagev1.DockerImageReference, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		imgRef	imagev1.DockerImageReference
		err	error
	)
	if index >= 0 && index < len(isi.Spec.Images) {
		imgSpec := &isi.Spec.Images[index]
		if imgSpec.From.Kind == "DockerImage" {
			imgRef, err = imageutil.ParseDockerImageReference(imgSpec.From.Name)
			if err == nil {
				return &imgRef, nil
			}
			err = fmt.Errorf("failed to parse isi.spec.images[%d].from.name %q: %v", index, imgSpec.From.Name, err)
		}
	}
	if index < 0 || index >= len(isi.Status.Images) {
		return nil, err
	}
	img := isi.Status.Images[index].Image
	if img == nil {
		return nil, err
	}
	imgRef, err = imageutil.ParseDockerImageReference(img.DockerImageReference)
	if err != nil {
		return nil, fmt.Errorf("failed to parse isi.status.images[%d].image.dockerImageReference %q: %v", index, img.DockerImageReference, err)
	}
	return &imgRef, nil
}
func mkImportInfo(registry string, status *metav1.Status) metrics.ImportErrorInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var reason string
	if status.Status != metav1.StatusSuccess {
		reason = string(status.Reason)
		if len(reason) == 0 {
			reason = reasonUnknown
		}
	}
	return metrics.ImportErrorInfo{Registry: registry, Reason: reason}
}
func defaultErrorInfoReason(info *metrics.ImportErrorInfo, err error) *metrics.ImportErrorInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(info.Reason) == 0 && err != nil {
		info.Reason = string(apierrs.ReasonForError(err))
		if len(info.Reason) == 0 {
			info.Reason = reasonUnknown
		}
	}
	return info
}
