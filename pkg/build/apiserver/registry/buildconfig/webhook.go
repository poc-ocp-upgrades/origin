package buildconfig

import (
	"context"
	"fmt"
	"github.com/openshift/api/build"
	buildv1 "github.com/openshift/api/build/v1"
	buildclienttyped "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	buildapi "github.com/openshift/origin/pkg/build/apis/build"
	buildv1helpers "github.com/openshift/origin/pkg/build/apis/build/v1"
	"github.com/openshift/origin/pkg/build/client"
	"github.com/openshift/origin/pkg/build/webhook"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	kubetypedclient "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/klog"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	"net/http"
	"strings"
)

var (
	webhookEncodingScheme       = runtime.NewScheme()
	webhookEncodingCodecFactory = serializer.NewCodecFactory(webhookEncodingScheme)
)

func init() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	utilruntime.Must(buildv1helpers.Install(webhookEncodingScheme))
	webhookEncodingCodecFactory = serializer.NewCodecFactory(webhookEncodingScheme)
}

type WebHook struct {
	groupVersion      schema.GroupVersion
	buildConfigClient buildclienttyped.BuildV1Interface
	secretsClient     kubetypedclient.SecretsGetter
	instantiator      client.BuildConfigInstantiator
	plugins           map[string]webhook.Plugin
}

func NewWebHookREST(buildConfigClient buildclienttyped.BuildV1Interface, secretsClient kubetypedclient.SecretsGetter, groupVersion schema.GroupVersion, plugins map[string]webhook.Plugin) *WebHook {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newWebHookREST(buildConfigClient, secretsClient, client.BuildConfigInstantiatorClient{Client: buildConfigClient}, groupVersion, plugins)
}
func newWebHookREST(buildConfigClient buildclienttyped.BuildV1Interface, secretsClient kubetypedclient.SecretsGetter, instantiator client.BuildConfigInstantiator, groupVersion schema.GroupVersion, plugins map[string]webhook.Plugin) *WebHook {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &WebHook{groupVersion: groupVersion, buildConfigClient: buildConfigClient, secretsClient: secretsClient, instantiator: instantiator, plugins: plugins}
}
func (h *WebHook) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &buildapi.Build{}
}
func (h *WebHook) Connect(ctx context.Context, name string, options runtime.Object, responder rest.Responder) (http.Handler, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &WebHookHandler{ctx: ctx, name: name, options: options.(*kapi.PodProxyOptions), responder: responder, groupVersion: h.groupVersion, plugins: h.plugins, buildConfigClient: h.buildConfigClient, secretsClient: h.secretsClient, instantiator: h.instantiator}, nil
}
func (h *WebHook) NewConnectOptions() (runtime.Object, bool, string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &kapi.PodProxyOptions{}, true, "path"
}
func (h *WebHook) ConnectMethods() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"POST"}
}

type WebHookHandler struct {
	ctx               context.Context
	name              string
	options           *kapi.PodProxyOptions
	responder         rest.Responder
	groupVersion      schema.GroupVersion
	plugins           map[string]webhook.Plugin
	buildConfigClient buildclienttyped.BuildV1Interface
	secretsClient     kubetypedclient.SecretsGetter
	instantiator      client.BuildConfigInstantiator
}

func (h *WebHookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := h.ProcessWebHook(w, r, h.ctx, h.name, h.options.Path); err != nil {
		h.responder.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (w *WebHookHandler) ProcessWebHook(writer http.ResponseWriter, req *http.Request, ctx context.Context, name, subpath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(strings.TrimPrefix(subpath, "/"), "/")
	if len(parts) != 2 {
		return errors.NewBadRequest(fmt.Sprintf("unexpected hook subpath %s", subpath))
	}
	secret, hookType := parts[0], parts[1]
	plugin, ok := w.plugins[hookType]
	if !ok {
		return errors.NewNotFound(build.Resource("buildconfighook"), hookType)
	}
	config, err := w.buildConfigClient.BuildConfigs(apirequest.NamespaceValue(ctx)).Get(name, metav1.GetOptions{})
	if err != nil {
		return errors.NewUnauthorized(fmt.Sprintf("the webhook %q for %q did not accept your secret", hookType, name))
	}
	triggers, err := plugin.GetTriggers(config)
	if err != nil {
		return errors.NewUnauthorized(fmt.Sprintf("the webhook %q for %q did not accept your secret", hookType, name))
	}
	klog.V(4).Infof("checking secret for %q webhook trigger of buildconfig %s/%s", hookType, config.Namespace, config.Name)
	trigger, err := webhook.CheckSecret(config.Namespace, secret, triggers, w.secretsClient)
	if err != nil {
		return errors.NewUnauthorized(fmt.Sprintf("the webhook %q for %q did not accept your secret", hookType, name))
	}
	revision, envvars, dockerStrategyOptions, proceed, err := plugin.Extract(config, trigger, req)
	if !proceed {
		switch err {
		case webhook.ErrSecretMismatch, webhook.ErrHookNotEnabled:
			return errors.NewUnauthorized(fmt.Sprintf("the webhook %q for %q did not accept your secret", hookType, name))
		case webhook.MethodNotSupported:
			return errors.NewMethodNotSupported(build.Resource("buildconfighook"), req.Method)
		}
		if _, ok := err.(*errors.StatusError); !ok && err != nil {
			return errors.NewInternalError(fmt.Errorf("hook failed: %v", err))
		}
		return err
	}
	warning := err
	buildTriggerCauses := webhook.GenerateBuildTriggerInfo(revision, hookType)
	request := &buildv1.BuildRequest{TriggeredBy: buildTriggerCauses, ObjectMeta: metav1.ObjectMeta{Name: name}, Revision: revision, Env: envvars, DockerStrategyOptions: dockerStrategyOptions}
	newBuild, err := w.instantiator.Instantiate(config.Namespace, request)
	if err != nil {
		return errors.NewInternalError(fmt.Errorf("could not generate a build: %v", err))
	}
	if newBuildEncoded, err := runtime.Encode(webhookEncodingCodecFactory.LegacyCodec(w.groupVersion), newBuild); err != nil {
		utilruntime.HandleError(err)
	} else {
		writer.Write(newBuildEncoded)
	}
	return warning
}
