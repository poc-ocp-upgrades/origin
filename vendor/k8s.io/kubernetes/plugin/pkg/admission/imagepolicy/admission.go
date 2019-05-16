package imagepolicy

import (
	"encoding/json"
	"errors"
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/api/imagepolicy/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/cache"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/util/webhook"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
	_ "k8s.io/kubernetes/pkg/apis/imagepolicy/install"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const PluginName = "ImagePolicyWebhook"

var AuditKeyPrefix = strings.ToLower(PluginName) + ".image-policy.k8s.io/"

const (
	ImagePolicyFailedOpenKeySuffix    string = "failed-open"
	ImagePolicyAuditRequiredKeySuffix string = "audit-required"
)

var (
	groupVersions = []schema.GroupVersion{v1alpha1.SchemeGroupVersion}
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		newImagePolicyWebhook, err := NewImagePolicyWebhook(config)
		if err != nil {
			return nil, err
		}
		return newImagePolicyWebhook, nil
	})
}

type Plugin struct {
	*admission.Handler
	webhook       *webhook.GenericWebhook
	responseCache *cache.LRUExpireCache
	allowTTL      time.Duration
	denyTTL       time.Duration
	retryBackoff  time.Duration
	defaultAllow  bool
}

var _ admission.ValidationInterface = &Plugin{}

func (a *Plugin) statusTTL(status v1alpha1.ImageReviewStatus) time.Duration {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if status.Allowed {
		return a.allowTTL
	}
	return a.denyTTL
}
func (a *Plugin) filterAnnotations(allAnnotations map[string]string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	annotations := make(map[string]string)
	for k, v := range allAnnotations {
		if strings.Contains(k, ".image-policy.k8s.io/") {
			annotations[k] = v
		}
	}
	return annotations
}
func (a *Plugin) webhookError(pod *api.Pod, attributes admission.Attributes, err error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err != nil {
		klog.V(2).Infof("error contacting webhook backend: %s", err)
		if a.defaultAllow {
			attributes.AddAnnotation(AuditKeyPrefix+ImagePolicyFailedOpenKeySuffix, "true")
			annotations := pod.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations[api.ImagePolicyFailedOpenKey] = "true"
			pod.ObjectMeta.SetAnnotations(annotations)
			klog.V(2).Infof("resource allowed in spite of webhook backend failure")
			return nil
		}
		klog.V(2).Infof("resource not allowed due to webhook backend failure ")
		return admission.NewForbidden(attributes, err)
	}
	return nil
}
func (a *Plugin) Validate(attributes admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attributes.GetSubresource() != "" || attributes.GetResource().GroupResource() != api.Resource("pods") {
		return nil
	}
	pod, ok := attributes.GetObject().(*api.Pod)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Pod but was unable to be converted")
	}
	var imageReviewContainerSpecs []v1alpha1.ImageReviewContainerSpec
	containers := make([]api.Container, 0, len(pod.Spec.Containers)+len(pod.Spec.InitContainers))
	containers = append(containers, pod.Spec.Containers...)
	containers = append(containers, pod.Spec.InitContainers...)
	for _, c := range containers {
		imageReviewContainerSpecs = append(imageReviewContainerSpecs, v1alpha1.ImageReviewContainerSpec{Image: c.Image})
	}
	imageReview := v1alpha1.ImageReview{Spec: v1alpha1.ImageReviewSpec{Containers: imageReviewContainerSpecs, Annotations: a.filterAnnotations(pod.Annotations), Namespace: attributes.GetNamespace()}}
	if err := a.admitPod(pod, attributes, &imageReview); err != nil {
		return admission.NewForbidden(attributes, err)
	}
	return nil
}
func (a *Plugin) admitPod(pod *api.Pod, attributes admission.Attributes, review *v1alpha1.ImageReview) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cacheKey, err := json.Marshal(review.Spec)
	if err != nil {
		return err
	}
	if entry, ok := a.responseCache.Get(string(cacheKey)); ok {
		review.Status = entry.(v1alpha1.ImageReviewStatus)
	} else {
		result := a.webhook.WithExponentialBackoff(func() rest.Result {
			return a.webhook.RestClient.Post().Body(review).Do()
		})
		if err := result.Error(); err != nil {
			return a.webhookError(pod, attributes, err)
		}
		var statusCode int
		if result.StatusCode(&statusCode); statusCode < 200 || statusCode >= 300 {
			return a.webhookError(pod, attributes, fmt.Errorf("Error contacting webhook: %d", statusCode))
		}
		if err := result.Into(review); err != nil {
			return a.webhookError(pod, attributes, err)
		}
		a.responseCache.Add(string(cacheKey), review.Status, a.statusTTL(review.Status))
	}
	for k, v := range review.Status.AuditAnnotations {
		if err := attributes.AddAnnotation(AuditKeyPrefix+k, v); err != nil {
			klog.Warningf("failed to set admission audit annotation %s to %s: %v", AuditKeyPrefix+k, v, err)
		}
	}
	if !review.Status.Allowed {
		if len(review.Status.Reason) > 0 {
			return fmt.Errorf("image policy webhook backend denied one or more images: %s", review.Status.Reason)
		}
		return errors.New("one or more images rejected by webhook backend")
	}
	return nil
}
func NewImagePolicyWebhook(configFile io.Reader) (*Plugin, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if configFile == nil {
		return nil, fmt.Errorf("no config specified")
	}
	var config AdmissionConfig
	d := yaml.NewYAMLOrJSONDecoder(configFile, 4096)
	err := d.Decode(&config)
	if err != nil {
		return nil, err
	}
	whConfig := config.ImagePolicyWebhook
	if err := normalizeWebhookConfig(&whConfig); err != nil {
		return nil, err
	}
	gw, err := webhook.NewGenericWebhook(legacyscheme.Scheme, legacyscheme.Codecs, whConfig.KubeConfigFile, groupVersions, whConfig.RetryBackoff)
	if err != nil {
		return nil, err
	}
	return &Plugin{Handler: admission.NewHandler(admission.Create, admission.Update), webhook: gw, responseCache: cache.NewLRUExpireCache(1024), allowTTL: whConfig.AllowTTL, denyTTL: whConfig.DenyTTL, defaultAllow: whConfig.DefaultAllow}, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
