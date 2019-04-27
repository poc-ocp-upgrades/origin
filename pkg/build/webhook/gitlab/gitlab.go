package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	buildv1 "github.com/openshift/api/build/v1"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	"github.com/openshift/origin/pkg/build/webhook"
)

type WebHookPlugin struct{}

func New() *WebHookPlugin {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &WebHookPlugin{}
}

type commit struct {
	ID	string				`json:"id,omitempty"`
	Author	buildv1.SourceControlUser	`json:"author,omitempty"`
	Message	string				`json:"message,omitempty"`
}
type pushEvent struct {
	Ref	string		`json:"ref,omitempty"`
	After	string		`json:"after,omitempty"`
	Commits	[]commit	`json:"commits,omitempty"`
}

func (p *WebHookPlugin) Extract(buildCfg *buildv1.BuildConfig, trigger *buildv1.WebHookTrigger, req *http.Request) (revision *buildv1.SourceRevision, envvars []corev1.EnvVar, dockerStrategyOptions *buildv1.DockerStrategyOptions, proceed bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(4).Infof("Verifying build request for BuildConfig %s/%s", buildCfg.Namespace, buildCfg.Name)
	if err = verifyRequest(req); err != nil {
		return revision, envvars, dockerStrategyOptions, proceed, err
	}
	method := getEvent(req.Header)
	if method != "Push Hook" {
		return revision, envvars, dockerStrategyOptions, proceed, errors.NewBadRequest(fmt.Sprintf("Unknown X-Gitlab-Event %s", method))
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return revision, envvars, dockerStrategyOptions, proceed, errors.NewBadRequest(err.Error())
	}
	var event pushEvent
	if err = json.Unmarshal(body, &event); err != nil {
		return revision, envvars, dockerStrategyOptions, proceed, errors.NewBadRequest(err.Error())
	}
	if !webhook.GitRefMatches(event.Ref, webhook.DefaultConfigRef, &buildCfg.Spec.Source) {
		klog.V(2).Infof("Skipping build for BuildConfig %s/%s.  Branch reference from '%s' does not match configuration", buildCfg.Namespace, buildCfg.Name, event)
		return revision, envvars, dockerStrategyOptions, proceed, err
	}
	lastCommit := event.Commits[len(event.Commits)-1]
	revision = &buildv1.SourceRevision{Git: &buildv1.GitSourceRevision{Commit: lastCommit.ID, Author: lastCommit.Author, Committer: lastCommit.Author, Message: lastCommit.Message}}
	return revision, envvars, dockerStrategyOptions, true, err
}
func (p *WebHookPlugin) GetTriggers(buildConfig *buildv1.BuildConfig) ([]*buildv1.WebHookTrigger, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	triggers := buildapihelpers.FindTriggerPolicy(buildv1.GitLabWebHookBuildTriggerType, buildConfig)
	webhookTriggers := []*buildv1.WebHookTrigger{}
	for _, trigger := range triggers {
		if trigger.GitLabWebHook != nil {
			webhookTriggers = append(webhookTriggers, trigger.GitLabWebHook)
		}
	}
	if len(webhookTriggers) == 0 {
		return nil, webhook.ErrHookNotEnabled
	}
	return webhookTriggers, nil
}
func verifyRequest(req *http.Request) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if method := req.Method; method != "POST" {
		return webhook.MethodNotSupported
	}
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return errors.NewBadRequest(fmt.Sprintf("non-parseable Content-Type %s (%s)", contentType, err))
	}
	if mediaType != "application/json" {
		return errors.NewBadRequest(fmt.Sprintf("unsupported Content-Type %s", contentType))
	}
	if len(getEvent(req.Header)) == 0 {
		return errors.NewBadRequest("missing X-Gitlab-Event")
	}
	return nil
}
func getEvent(header http.Header) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return header.Get("X-Gitlab-Event")
}
