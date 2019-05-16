package envresolve

import (
	"fmt"
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/api/v1/resource"
	"k8s.io/kubernetes/pkg/fieldpath"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ResourceStore struct {
	SecretStore    map[string]*corev1.Secret
	ConfigMapStore map[string]*corev1.ConfigMap
}

func NewResourceStore() *ResourceStore {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ResourceStore{SecretStore: make(map[string]*corev1.Secret), ConfigMapStore: make(map[string]*corev1.ConfigMap)}
}
func getSecretRefValue(client kubernetes.Interface, namespace string, store *ResourceStore, secretSelector *corev1.SecretKeySelector) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secret, ok := store.SecretStore[secretSelector.Name]
	if !ok {
		var err error
		secret, err = client.CoreV1().Secrets(namespace).Get(secretSelector.Name, metav1.GetOptions{})
		if err != nil {
			return "", err
		}
		store.SecretStore[secretSelector.Name] = secret
	}
	if data, ok := secret.Data[secretSelector.Key]; ok {
		return string(data), nil
	}
	return "", fmt.Errorf("key %s not found in secret %s", secretSelector.Key, secretSelector.Name)
}
func getConfigMapRefValue(client kubernetes.Interface, namespace string, store *ResourceStore, configMapSelector *corev1.ConfigMapKeySelector) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configMap, ok := store.ConfigMapStore[configMapSelector.Name]
	if !ok {
		var err error
		configMap, err = client.CoreV1().ConfigMaps(namespace).Get(configMapSelector.Name, metav1.GetOptions{})
		if err != nil {
			return "", err
		}
		store.ConfigMapStore[configMapSelector.Name] = configMap
	}
	if data, ok := configMap.Data[configMapSelector.Key]; ok {
		return string(data), nil
	}
	return "", fmt.Errorf("key %s not found in config map %s", configMapSelector.Key, configMapSelector.Name)
}
func getFieldRef(obj runtime.Object, from *corev1.EnvVarSource) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fieldpath.ExtractFieldPathAsString(obj, from.FieldRef.FieldPath)
}
func getResourceFieldRef(from *corev1.EnvVarSource, c *corev1.Container) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return resource.ExtractContainerResourceValue(from.ResourceFieldRef, c)
}
func GetEnvVarRefValue(kc kubernetes.Interface, ns string, store *ResourceStore, from *corev1.EnvVarSource, obj runtime.Object, c *corev1.Container) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if from.SecretKeyRef != nil {
		return getSecretRefValue(kc, ns, store, from.SecretKeyRef)
	}
	if from.ConfigMapKeyRef != nil {
		return getConfigMapRefValue(kc, ns, store, from.ConfigMapKeyRef)
	}
	if from.FieldRef != nil {
		return getFieldRef(obj, from)
	}
	if from.ResourceFieldRef != nil {
		return getResourceFieldRef(from, c)
	}
	return "", fmt.Errorf("invalid valueFrom")
}
func GetEnvVarRefString(from *corev1.EnvVarSource) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if from.ConfigMapKeyRef != nil {
		return fmt.Sprintf("configmap %s, key %s", from.ConfigMapKeyRef.Name, from.ConfigMapKeyRef.Key)
	}
	if from.SecretKeyRef != nil {
		return fmt.Sprintf("secret %s, key %s", from.SecretKeyRef.Name, from.SecretKeyRef.Key)
	}
	if from.FieldRef != nil {
		return fmt.Sprintf("field path %s", from.FieldRef.FieldPath)
	}
	if from.ResourceFieldRef != nil {
		containerPrefix := ""
		if from.ResourceFieldRef.ContainerName != "" {
			containerPrefix = fmt.Sprintf("%s/", from.ResourceFieldRef.ContainerName)
		}
		return fmt.Sprintf("resource field %s%s", containerPrefix, from.ResourceFieldRef.Resource)
	}
	return "invalid valueFrom"
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
