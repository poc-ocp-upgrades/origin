package v1

import (
	"errors"
	buildv1 "github.com/openshift/api/build/v1"
	"k8s.io/client-go/rest"
	"net/url"
)

var ErrTriggerIsNotAWebHook = errors.New("the specified trigger is not a webhook")

type WebHookURLInterface interface {
	WebHookURL(name string, trigger *buildv1.BuildTriggerPolicy) (*url.URL, error)
}

func NewWebhookURLClient(c rest.Interface, ns string) WebHookURLInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &webhooks{client: c, ns: ns}
}

type webhooks struct {
	client rest.Interface
	ns     string
}

func (c *webhooks) WebHookURL(name string, trigger *buildv1.BuildTriggerPolicy) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hooks := c.client.Get().Namespace(c.ns).Resource("buildConfigs").Name(name).SubResource("webhooks")
	switch {
	case trigger.GenericWebHook != nil:
		return hooks.Suffix("<secret>", "generic").URL(), nil
	case trigger.GitHubWebHook != nil:
		return hooks.Suffix("<secret>", "github").URL(), nil
	case trigger.GitLabWebHook != nil:
		return hooks.Suffix("<secret>", "gitlab").URL(), nil
	case trigger.BitbucketWebHook != nil:
		return hooks.Suffix("<secret>", "bitbucket").URL(), nil
	default:
		return nil, ErrTriggerIsNotAWebHook
	}
}
