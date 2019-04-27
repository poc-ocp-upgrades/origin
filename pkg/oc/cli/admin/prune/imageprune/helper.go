package imageprune

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"github.com/docker/distribution/registry/api/errcode"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	kmeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	ref "k8s.io/client-go/tools/reference"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	imagev1 "github.com/openshift/api/image/v1"
	"github.com/openshift/library-go/pkg/image/reference"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/util/netutils"
)

type imgByAge []*imagev1.Image

func (ba imgByAge) Len() int {
	_logClusterCodePath()
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
	return len(ba)
}
func (ba imgByAge) Swap(i, j int) {
	_logClusterCodePath()
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
	ba[i], ba[j] = ba[j], ba[i]
}
func (ba imgByAge) Less(i, j int) bool {
	_logClusterCodePath()
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
	return ba[i].CreationTimestamp.After(ba[j].CreationTimestamp.Time)
}

type isByAge []imagev1.ImageStream

func (ba isByAge) Len() int {
	_logClusterCodePath()
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
	return len(ba)
}
func (ba isByAge) Swap(i, j int) {
	_logClusterCodePath()
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
	ba[i], ba[j] = ba[j], ba[i]
}
func (ba isByAge) Less(i, j int) bool {
	_logClusterCodePath()
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
	return ba[i].CreationTimestamp.After(ba[j].CreationTimestamp.Time)
}
func DetermineRegistryHost(images *imagev1.ImageList, imageStreams *imagev1.ImageStreamList) (string, error) {
	_logClusterCodePath()
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
	var pullSpec string
	var managedImages []*imagev1.Image
	for i := range images.Items {
		image := &images.Items[i]
		if image.Annotations[imageapi.ManagedByOpenShiftAnnotation] != "true" {
			continue
		}
		managedImages = append(managedImages, image)
	}
	sort.Sort(imgByAge(managedImages))
	if len(managedImages) > 0 {
		pullSpec = managedImages[0].DockerImageReference
	} else {
		sort.Sort(isByAge(imageStreams.Items))
		for _, is := range imageStreams.Items {
			if len(is.Status.DockerImageRepository) == 0 {
				continue
			}
			pullSpec = is.Status.DockerImageRepository
		}
	}
	if len(pullSpec) == 0 {
		return "", fmt.Errorf("no managed image found")
	}
	ref, err := reference.Parse(pullSpec)
	if err != nil {
		return "", fmt.Errorf("unable to parse %q: %v", pullSpec, err)
	}
	if len(ref.Registry) == 0 {
		return "", fmt.Errorf("%s does not include a registry", pullSpec)
	}
	return ref.Registry, nil
}

type RegistryPinger interface {
	Ping(registry string) (*url.URL, error)
}
type DefaultRegistryPinger struct {
	Client		*http.Client
	Insecure	bool
}

func (drp *DefaultRegistryPinger) Ping(registry string) (*url.URL, error) {
	_logClusterCodePath()
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
	var (
		registryURL	*url.URL
		err		error
	)
pathLoop:
	for _, path := range []string{"/", "/healthz"} {
		registryURL, err = TryProtocolsWithRegistryURL(registry, drp.Insecure, func(u url.URL) error {
			u.Path = path
			healthResponse, err := drp.Client.Get(u.String())
			if err != nil {
				return err
			}
			defer healthResponse.Body.Close()
			if healthResponse.StatusCode != http.StatusOK {
				return &retryPath{err: fmt.Errorf("unexpected status: %s", healthResponse.Status)}
			}
			return nil
		})
		switch t := err.(type) {
		case *retryPath:
			err = t.err
			continue pathLoop
		case kerrors.Aggregate:
			for _, err := range t.Errors() {
				if _, ok := err.(*retryPath); ok {
					continue pathLoop
				}
			}
		}
		break
	}
	return registryURL, err
}

type DryRunRegistryPinger struct{}

func (*DryRunRegistryPinger) Ping(registry string) (*url.URL, error) {
	_logClusterCodePath()
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
	return url.Parse("https://" + registry)
}
func TryProtocolsWithRegistryURL(registry string, allowInsecure bool, action func(registryURL url.URL) error) (*url.URL, error) {
	_logClusterCodePath()
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
	var errs []error
	if !strings.Contains(registry, "://") {
		registry = "unset://" + registry
	}
	url, err := url.Parse(registry)
	if err != nil {
		return nil, err
	}
	var protos []string
	switch {
	case len(url.Scheme) > 0 && url.Scheme != "unset":
		protos = []string{url.Scheme}
	case allowInsecure || netutils.IsPrivateAddress(registry):
		protos = []string{"https", "http"}
	default:
		protos = []string{"https"}
	}
	registry = url.Host
	for _, proto := range protos {
		klog.V(4).Infof("Trying protocol %s for the registry URL %s", proto, registry)
		url.Scheme = proto
		err := action(*url)
		if err == nil {
			return url, nil
		}
		if err != nil {
			klog.V(4).Infof("Error with %s for %s: %v", proto, registry, err)
		}
		if _, ok := err.(*errcode.Errors); ok {
			return url, err
		}
		errs = append(errs, err)
		if proto == "https" && strings.Contains(err.Error(), "server gave HTTP response to HTTPS client") && !allowInsecure {
			errs = append(errs, fmt.Errorf("\n* Append --force-insecure if you really want to prune the registry using insecure connection."))
		} else if proto == "http" && strings.Contains(err.Error(), "malformed HTTP response") {
			errs = append(errs, fmt.Errorf("\n* Are you trying to connect to a TLS-enabled registry without TLS?"))
		}
	}
	return nil, kerrors.NewAggregate(errs)
}

type retryPath struct{ err error }

func (rp *retryPath) Error() string {
	_logClusterCodePath()
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
	return rp.err.Error()
}

type ErrBadReference struct {
	kind		string
	namespace	string
	name		string
	targetKind	string
	reference	string
	reason		string
}

func newErrBadReferenceToImage(reference string, obj *corev1.ObjectReference, reason string) error {
	_logClusterCodePath()
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
	kind := "<UnknownType>"
	namespace := ""
	name := "<unknown-name>"
	if obj != nil {
		kind = obj.Kind
		namespace = obj.Namespace
		name = obj.Name
	}
	return &ErrBadReference{kind: kind, namespace: namespace, name: name, reference: reference, reason: reason}
}
func newErrBadReferenceTo(targetKind, reference string, obj *corev1.ObjectReference, reason string) error {
	_logClusterCodePath()
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
	return &ErrBadReference{kind: obj.Kind, namespace: obj.Namespace, name: obj.Name, targetKind: targetKind, reference: reference, reason: reason}
}
func (e *ErrBadReference) Error() string {
	_logClusterCodePath()
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
	return e.String()
}
func (e *ErrBadReference) String() string {
	_logClusterCodePath()
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
	name := e.name
	if len(e.namespace) > 0 {
		name = e.namespace + "/" + name
	}
	targetKind := "docker image"
	if len(e.targetKind) > 0 {
		targetKind = e.targetKind
	}
	return fmt.Sprintf("%s[%s]: invalid %s reference %q: %s", e.kind, name, targetKind, e.reference, e.reason)
}
func getName(obj runtime.Object) string {
	_logClusterCodePath()
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
	accessor, err := kmeta.Accessor(obj)
	if err != nil {
		klog.V(4).Infof("Error getting accessor for %#v", obj)
		return "<unknown>"
	}
	ns := accessor.GetNamespace()
	if len(ns) == 0 {
		return accessor.GetName()
	}
	return fmt.Sprintf("%s/%s", ns, accessor.GetName())
}
func getKindName(obj *corev1.ObjectReference) string {
	_logClusterCodePath()
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
	if obj == nil {
		return "unknown object"
	}
	name := obj.Name
	if len(obj.Namespace) > 0 {
		name = obj.Namespace + "/" + name
	}
	return fmt.Sprintf("%s[%s]", obj.Kind, name)
}
func getRef(obj runtime.Object) *corev1.ObjectReference {
	_logClusterCodePath()
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
	ref, err := ref.GetReference(scheme.Scheme, obj)
	if err != nil {
		klog.Errorf("failed to get reference to object %T: %v", obj, err)
		return nil
	}
	return ref
}
