package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
	versionedinformers "k8s.io/client-go/informers"
	"k8s.io/kubernetes/pkg/kubeapiserver/authorizer"
	authzmodes "k8s.io/kubernetes/pkg/kubeapiserver/authorizer/modes"
	"strings"
	"time"
)

type BuiltInAuthorizationOptions struct {
	Modes                       []string
	PolicyFile                  string
	WebhookConfigFile           string
	WebhookCacheAuthorizedTTL   time.Duration
	WebhookCacheUnauthorizedTTL time.Duration
}

func NewBuiltInAuthorizationOptions() *BuiltInAuthorizationOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &BuiltInAuthorizationOptions{Modes: []string{authzmodes.ModeAlwaysAllow}, WebhookCacheAuthorizedTTL: 5 * time.Minute, WebhookCacheUnauthorizedTTL: 30 * time.Second}
}
func (s *BuiltInAuthorizationOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s == nil {
		return nil
	}
	allErrors := []error{}
	if len(s.Modes) == 0 {
		allErrors = append(allErrors, fmt.Errorf("at least one authorization-mode must be passed"))
	}
	allowedModes := sets.NewString(authzmodes.AuthorizationModeChoices...)
	modes := sets.NewString(s.Modes...)
	for _, mode := range s.Modes {
		if !allowedModes.Has(mode) {
			allErrors = append(allErrors, fmt.Errorf("authorization-mode %q is not a valid mode", mode))
		}
		if mode == authzmodes.ModeABAC {
			if s.PolicyFile == "" {
				allErrors = append(allErrors, fmt.Errorf("authorization-mode ABAC's authorization policy file not passed"))
			}
		}
		if mode == authzmodes.ModeWebhook {
			if s.WebhookConfigFile == "" {
				allErrors = append(allErrors, fmt.Errorf("authorization-mode Webhook's authorization config file not passed"))
			}
		}
	}
	if s.PolicyFile != "" && !modes.Has(authzmodes.ModeABAC) {
		allErrors = append(allErrors, fmt.Errorf("cannot specify --authorization-policy-file without mode ABAC"))
	}
	if s.WebhookConfigFile != "" && !modes.Has(authzmodes.ModeWebhook) {
		allErrors = append(allErrors, fmt.Errorf("cannot specify --authorization-webhook-config-file without mode Webhook"))
	}
	if len(s.Modes) != len(modes.List()) {
		allErrors = append(allErrors, fmt.Errorf("authorization-mode %q has mode specified more than once", s.Modes))
	}
	return allErrors
}
func (s *BuiltInAuthorizationOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringSliceVar(&s.Modes, "authorization-mode", s.Modes, ""+"Ordered list of plug-ins to do authorization on secure port. Comma-delimited list of: "+strings.Join(authzmodes.AuthorizationModeChoices, ",")+".")
	fs.StringVar(&s.PolicyFile, "authorization-policy-file", s.PolicyFile, ""+"File with authorization policy in csv format, used with --authorization-mode=ABAC, on the secure port.")
	fs.StringVar(&s.WebhookConfigFile, "authorization-webhook-config-file", s.WebhookConfigFile, ""+"File with webhook configuration in kubeconfig format, used with --authorization-mode=Webhook. "+"The API server will query the remote service to determine access on the API server's secure port.")
	fs.DurationVar(&s.WebhookCacheAuthorizedTTL, "authorization-webhook-cache-authorized-ttl", s.WebhookCacheAuthorizedTTL, "The duration to cache 'authorized' responses from the webhook authorizer.")
	fs.DurationVar(&s.WebhookCacheUnauthorizedTTL, "authorization-webhook-cache-unauthorized-ttl", s.WebhookCacheUnauthorizedTTL, "The duration to cache 'unauthorized' responses from the webhook authorizer.")
}
func (s *BuiltInAuthorizationOptions) ToAuthorizationConfig(versionedInformerFactory versionedinformers.SharedInformerFactory) authorizer.Config {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return authorizer.Config{AuthorizationModes: s.Modes, PolicyFile: s.PolicyFile, WebhookConfigFile: s.WebhookConfigFile, WebhookCacheAuthorizedTTL: s.WebhookCacheAuthorizedTTL, WebhookCacheUnauthorizedTTL: s.WebhookCacheUnauthorizedTTL, VersionedInformerFactory: versionedInformerFactory}
}
