package util

import (
	"fmt"
	buildv1 "github.com/openshift/api/build/v1"
	buildlister "github.com/openshift/client-go/build/listers/build/v1"
	"github.com/openshift/origin/pkg/api/apihelpers"
	"github.com/openshift/origin/pkg/build/buildapihelpers"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/credentialprovider"
	credentialprovidersecrets "k8s.io/kubernetes/pkg/credentialprovider/secrets"
	"net/url"
	"strings"
)

const (
	NoBuildLogsMessage        = "No logs are available."
	BuildWorkDirMount         = "/tmp/build"
	BuilderServiceAccountName = "builder"
	buildPodSuffix            = "build"
	BuildBlobsMetaCache       = "/var/lib/containers/cache"
	BuildBlobsContentCache    = "/var/cache/blobs"
)

type GeneratorFatalError struct{ Reason string }

func (e *GeneratorFatalError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("fatal error generating Build from BuildConfig: %s", e.Reason)
}
func IsFatalGeneratorError(err error) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, isFatal := err.(*GeneratorFatalError)
	return isFatal
}
func GetBuildPodName(build *buildv1.Build) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apihelpers.GetPodName(build.Name, buildPodSuffix)
}
func IsBuildComplete(build *buildv1.Build) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return IsTerminalPhase(build.Status.Phase)
}
func IsTerminalPhase(phase buildv1.BuildPhase) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch phase {
	case buildv1.BuildPhaseNew, buildv1.BuildPhasePending, buildv1.BuildPhaseRunning:
		return false
	}
	return true
}
func BuildNameForConfigVersion(name string, version int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s-%d", name, version)
}
func BuildConfigSelector(name string) labels.Selector {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return labels.Set{BuildConfigLabel: buildapihelpers.LabelValue(name)}.AsSelector()
}

type buildFilter func(*buildv1.Build) bool

func BuildConfigBuilds(c buildlister.BuildLister, namespace, name string, filterFunc buildFilter) ([]*buildv1.Build, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	result, err := c.Builds(namespace).List(BuildConfigSelector(name))
	if err != nil {
		return nil, err
	}
	if filterFunc == nil {
		return result, nil
	}
	var filteredList []*buildv1.Build
	for _, b := range result {
		if filterFunc(b) {
			filteredList = append(filteredList, b)
		}
	}
	return filteredList, nil
}
func ConfigNameForBuild(build *buildv1.Build) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if build == nil {
		return ""
	}
	if build.Annotations != nil {
		if _, exists := build.Annotations[BuildConfigAnnotation]; exists {
			return build.Annotations[BuildConfigAnnotation]
		}
	}
	if _, exists := build.Labels[BuildConfigLabel]; exists {
		return build.Labels[BuildConfigLabel]
	}
	return build.Labels[BuildConfigLabelDeprecated]
}
func MergeTrustedEnvWithoutDuplicates(source []corev1.EnvVar, output *[]corev1.EnvVar, sourcePrecedence bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	MergeEnvWithoutDuplicates(source, output, sourcePrecedence, WhitelistEnvVarNames)
}
func MergeEnvWithoutDuplicates(source []corev1.EnvVar, output *[]corev1.EnvVar, sourcePrecedence bool, whitelist []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filteredSourceMap := make(map[string]corev1.EnvVar)
	for _, env := range source {
		allowed := false
		if len(whitelist) == 0 {
			allowed = true
		} else {
			for _, acceptable := range WhitelistEnvVarNames {
				if env.Name == acceptable {
					allowed = true
					break
				}
			}
		}
		if allowed {
			filteredSourceMap[env.Name] = env
		}
	}
	result := *output
	for i, env := range result {
		if v, found := filteredSourceMap[env.Name]; found {
			if sourcePrecedence {
				result[i].Value = v.Value
			}
			delete(filteredSourceMap, env.Name)
		}
	}
	for _, v := range source {
		if v, ok := filteredSourceMap[v.Name]; ok {
			result = append(result, v)
		}
	}
	*output = result
}
func GetBuildEnv(build *buildv1.Build) []corev1.EnvVar {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case build.Spec.Strategy.SourceStrategy != nil:
		return build.Spec.Strategy.SourceStrategy.Env
	case build.Spec.Strategy.DockerStrategy != nil:
		return build.Spec.Strategy.DockerStrategy.Env
	case build.Spec.Strategy.CustomStrategy != nil:
		return build.Spec.Strategy.CustomStrategy.Env
	case build.Spec.Strategy.JenkinsPipelineStrategy != nil:
		return build.Spec.Strategy.JenkinsPipelineStrategy.Env
	default:
		return nil
	}
}
func SetBuildEnv(build *buildv1.Build, env []corev1.EnvVar) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var oldEnv *[]corev1.EnvVar
	switch {
	case build.Spec.Strategy.SourceStrategy != nil:
		oldEnv = &build.Spec.Strategy.SourceStrategy.Env
	case build.Spec.Strategy.DockerStrategy != nil:
		oldEnv = &build.Spec.Strategy.DockerStrategy.Env
	case build.Spec.Strategy.CustomStrategy != nil:
		oldEnv = &build.Spec.Strategy.CustomStrategy.Env
	case build.Spec.Strategy.JenkinsPipelineStrategy != nil:
		oldEnv = &build.Spec.Strategy.JenkinsPipelineStrategy.Env
	default:
		return
	}
	*oldEnv = env
}
func UpdateBuildEnv(build *buildv1.Build, env []corev1.EnvVar) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buildEnv := GetBuildEnv(build)
	newEnv := []corev1.EnvVar{}
	for _, e := range buildEnv {
		exists := false
		for _, n := range env {
			if e.Name == n.Name {
				exists = true
				break
			}
		}
		if !exists {
			newEnv = append(newEnv, e)
		}
	}
	newEnv = append(newEnv, env...)
	SetBuildEnv(build, newEnv)
}
func FindDockerSecretAsReference(secrets []corev1.Secret, image string) *corev1.LocalObjectReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	emptyKeyring := credentialprovider.BasicDockerKeyring{}
	for _, secret := range secrets {
		secretList := []corev1.Secret{secret}
		keyring, err := credentialprovidersecrets.MakeDockerKeyring(secretList, &emptyKeyring)
		if err != nil {
			klog.V(2).Infof("Unable to make the Docker keyring for %s/%s secret: %v", secret.Name, secret.Namespace, err)
			continue
		}
		if _, found := keyring.Lookup(image); found {
			return &corev1.LocalObjectReference{Name: secret.Name}
		}
	}
	return nil
}
func FetchServiceAccountSecrets(secretStore v1lister.SecretLister, serviceAccountStore v1lister.ServiceAccountLister, namespace, serviceAccount string) ([]corev1.Secret, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var result []corev1.Secret
	sa, err := serviceAccountStore.ServiceAccounts(namespace).Get(serviceAccount)
	if err != nil {
		return result, fmt.Errorf("Error getting push/pull secrets for service account %s/%s: %v", namespace, serviceAccount, err)
	}
	for _, ref := range sa.Secrets {
		secret, err := secretStore.Secrets(namespace).Get(ref.Name)
		if err != nil {
			continue
		}
		result = append(result, *secret)
	}
	return result, nil
}
func UpdateCustomImageEnv(strategy *buildv1.CustomBuildStrategy, newImage string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if strategy.Env == nil {
		strategy.Env = make([]corev1.EnvVar, 1)
		strategy.Env[0] = corev1.EnvVar{Name: CustomBuildStrategyBaseImageKey, Value: newImage}
	} else {
		found := false
		for i := range strategy.Env {
			klog.V(4).Infof("Checking env variable %s %s", strategy.Env[i].Name, strategy.Env[i].Value)
			if strategy.Env[i].Name == CustomBuildStrategyBaseImageKey {
				found = true
				strategy.Env[i].Value = newImage
				klog.V(4).Infof("Updated env variable %s to %s", strategy.Env[i].Name, strategy.Env[i].Value)
				break
			}
		}
		if !found {
			strategy.Env = append(strategy.Env, corev1.EnvVar{Name: CustomBuildStrategyBaseImageKey, Value: newImage})
		}
	}
}
func ParseProxyURL(proxy string) (*url.URL, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	proxyURL, err := url.Parse(proxy)
	if err != nil || !strings.HasPrefix(proxyURL.Scheme, "http") {
		if proxyURL, err := url.Parse("http://" + proxy); err == nil {
			return proxyURL, nil
		}
	}
	return proxyURL, err
}
func GetInputReference(strategy buildv1.BuildStrategy) *corev1.ObjectReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case strategy.SourceStrategy != nil:
		return &strategy.SourceStrategy.From
	case strategy.DockerStrategy != nil:
		return strategy.DockerStrategy.From
	case strategy.CustomStrategy != nil:
		return &strategy.CustomStrategy.From
	default:
		return nil
	}
}
