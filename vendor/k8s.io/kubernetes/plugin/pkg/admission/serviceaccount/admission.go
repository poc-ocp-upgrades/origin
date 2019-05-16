package serviceaccount

import (
	"fmt"
	goformat "fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/apiserver/pkg/storage/names"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	podutil "k8s.io/kubernetes/pkg/api/pod"
	api "k8s.io/kubernetes/pkg/apis/core"
	kubefeatures "k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/kubeapiserver/admission/util"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"math/rand"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"time"
	gotime "time"
)

const (
	DefaultServiceAccountName         = "default"
	EnforceMountableSecretsAnnotation = "kubernetes.io/enforce-mountable-secrets"
	ServiceAccountVolumeName          = "kube-api-access"
	DefaultAPITokenMountPath          = "/var/run/secrets/kubernetes.io/serviceaccount"
	PluginName                        = "ServiceAccount"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		serviceAccountAdmission := NewServiceAccount()
		return serviceAccountAdmission, nil
	})
}

var _ = admission.Interface(&serviceAccount{})

type serviceAccount struct {
	*admission.Handler
	LimitSecretReferences    bool
	RequireAPIToken          bool
	MountServiceAccountToken bool
	client                   kubernetes.Interface
	serviceAccountLister     corev1listers.ServiceAccountLister
	secretLister             corev1listers.SecretLister
	generateName             func(string) string
	featureGate              utilfeature.FeatureGate
}

var _ admission.MutationInterface = &serviceAccount{}
var _ admission.ValidationInterface = &serviceAccount{}
var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&serviceAccount{})
var _ = genericadmissioninitializer.WantsExternalKubeInformerFactory(&serviceAccount{})

func NewServiceAccount() *serviceAccount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &serviceAccount{Handler: admission.NewHandler(admission.Create, admission.Update), LimitSecretReferences: false, MountServiceAccountToken: true, RequireAPIToken: true, generateName: names.SimpleNameGenerator.GenerateName, featureGate: utilfeature.DefaultFeatureGate}
}
func (a *serviceAccount) SetExternalKubeClientSet(cl kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.client = cl
}
func (a *serviceAccount) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serviceAccountInformer := f.Core().V1().ServiceAccounts()
	a.serviceAccountLister = serviceAccountInformer.Lister()
	secretInformer := f.Core().V1().Secrets()
	a.secretLister = secretInformer.Lister()
	a.SetReadyFunc(func() bool {
		return serviceAccountInformer.Informer().HasSynced() && secretInformer.Informer().HasSynced()
	})
}
func (a *serviceAccount) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.client == nil {
		return fmt.Errorf("missing client")
	}
	if a.secretLister == nil {
		return fmt.Errorf("missing secretLister")
	}
	if a.serviceAccountLister == nil {
		return fmt.Errorf("missing serviceAccountLister")
	}
	return nil
}
func (s *serviceAccount) Admit(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	updateInitialized, err := util.IsUpdatingInitializedObject(a)
	if err != nil {
		return err
	}
	if updateInitialized {
		return nil
	}
	pod := a.GetObject().(*api.Pod)
	if _, isMirrorPod := pod.Annotations[api.MirrorPodAnnotationKey]; isMirrorPod {
		return s.Validate(a)
	}
	if len(pod.Spec.ServiceAccountName) == 0 {
		pod.Spec.ServiceAccountName = DefaultServiceAccountName
	}
	serviceAccount, err := s.getServiceAccount(a.GetNamespace(), pod.Spec.ServiceAccountName)
	if err != nil {
		return admission.NewForbidden(a, fmt.Errorf("error looking up service account %s/%s: %v", a.GetNamespace(), pod.Spec.ServiceAccountName, err))
	}
	if s.MountServiceAccountToken && shouldAutomount(serviceAccount, pod) {
		if err := s.mountServiceAccountToken(serviceAccount, pod); err != nil {
			if _, ok := err.(errors.APIStatus); ok {
				return err
			}
			return admission.NewForbidden(a, err)
		}
	}
	if len(pod.Spec.ImagePullSecrets) == 0 {
		pod.Spec.ImagePullSecrets = make([]api.LocalObjectReference, len(serviceAccount.ImagePullSecrets))
		for i := 0; i < len(serviceAccount.ImagePullSecrets); i++ {
			pod.Spec.ImagePullSecrets[i].Name = serviceAccount.ImagePullSecrets[i].Name
		}
	}
	return s.Validate(a)
}
func (s *serviceAccount) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if shouldIgnore(a) {
		return nil
	}
	updateInitialized, err := util.IsUpdatingInitializedObject(a)
	if err != nil {
		return err
	}
	if updateInitialized {
		return nil
	}
	pod := a.GetObject().(*api.Pod)
	if _, isMirrorPod := pod.Annotations[api.MirrorPodAnnotationKey]; isMirrorPod {
		if len(pod.Spec.ServiceAccountName) != 0 {
			return admission.NewForbidden(a, fmt.Errorf("a mirror pod may not reference service accounts"))
		}
		hasSecrets := false
		podutil.VisitPodSecretNames(pod, func(name string) bool {
			hasSecrets = true
			return false
		})
		if hasSecrets {
			return admission.NewForbidden(a, fmt.Errorf("a mirror pod may not reference secrets"))
		}
		for _, v := range pod.Spec.Volumes {
			if proj := v.Projected; proj != nil {
				for _, projSource := range proj.Sources {
					if projSource.ServiceAccountToken != nil {
						return admission.NewForbidden(a, fmt.Errorf("a mirror pod may not use ServiceAccountToken volume projections"))
					}
				}
			}
		}
		return nil
	}
	serviceAccount, err := s.getServiceAccount(a.GetNamespace(), pod.Spec.ServiceAccountName)
	if err != nil {
		return admission.NewForbidden(a, fmt.Errorf("error looking up service account %s/%s: %v", a.GetNamespace(), pod.Spec.ServiceAccountName, err))
	}
	if s.enforceMountableSecrets(serviceAccount) {
		if err := s.limitSecretReferences(serviceAccount, pod); err != nil {
			return admission.NewForbidden(a, err)
		}
	}
	return nil
}
func shouldIgnore(a admission.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != api.Resource("pods") {
		return true
	}
	obj := a.GetObject()
	if obj == nil {
		return true
	}
	_, ok := obj.(*api.Pod)
	if !ok {
		return true
	}
	return false
}
func shouldAutomount(sa *corev1.ServiceAccount, pod *api.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pod.Spec.AutomountServiceAccountToken != nil {
		return *pod.Spec.AutomountServiceAccountToken
	}
	if sa.AutomountServiceAccountToken != nil {
		return *sa.AutomountServiceAccountToken
	}
	return true
}
func (s *serviceAccount) enforceMountableSecrets(serviceAccount *corev1.ServiceAccount) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.LimitSecretReferences {
		return true
	}
	if value, ok := serviceAccount.Annotations[EnforceMountableSecretsAnnotation]; ok {
		enforceMountableSecretCheck, _ := strconv.ParseBool(value)
		return enforceMountableSecretCheck
	}
	return false
}
func (s *serviceAccount) getServiceAccount(namespace string, name string) (*corev1.ServiceAccount, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serviceAccount, err := s.serviceAccountLister.ServiceAccounts(namespace).Get(name)
	if err == nil {
		return serviceAccount, nil
	}
	if !errors.IsNotFound(err) {
		return nil, err
	}
	numAttempts := 1
	if name == DefaultServiceAccountName {
		numAttempts = 10
	}
	retryInterval := time.Duration(rand.Int63n(100)+int64(100)) * time.Millisecond
	for i := 0; i < numAttempts; i++ {
		if i != 0 {
			time.Sleep(retryInterval)
		}
		serviceAccount, err := s.client.Core().ServiceAccounts(namespace).Get(name, metav1.GetOptions{})
		if err == nil {
			return serviceAccount, nil
		}
		if !errors.IsNotFound(err) {
			return nil, err
		}
	}
	return nil, errors.NewNotFound(api.Resource("serviceaccount"), name)
}
func (s *serviceAccount) getReferencedServiceAccountToken(serviceAccount *corev1.ServiceAccount) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(serviceAccount.Secrets) == 0 {
		return "", nil
	}
	tokens, err := s.getServiceAccountTokens(serviceAccount)
	if err != nil {
		return "", err
	}
	accountTokens := sets.NewString()
	for _, token := range tokens {
		accountTokens.Insert(token.Name)
	}
	for _, secret := range serviceAccount.Secrets {
		if accountTokens.Has(secret.Name) {
			return secret.Name, nil
		}
	}
	return "", nil
}
func (s *serviceAccount) getServiceAccountTokens(serviceAccount *corev1.ServiceAccount) ([]*corev1.Secret, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	secrets, err := s.secretLister.Secrets(serviceAccount.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	tokens := []*corev1.Secret{}
	for _, secret := range secrets {
		if secret.Type != corev1.SecretTypeServiceAccountToken {
			continue
		}
		if serviceaccount.IsServiceAccountToken(secret, serviceAccount) {
			tokens = append(tokens, secret)
		}
	}
	return tokens, nil
}
func (s *serviceAccount) limitSecretReferences(serviceAccount *corev1.ServiceAccount, pod *api.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mountableSecrets := sets.NewString()
	for _, s := range serviceAccount.Secrets {
		mountableSecrets.Insert(s.Name)
	}
	for _, volume := range pod.Spec.Volumes {
		source := volume.VolumeSource
		if source.Secret == nil {
			continue
		}
		secretName := source.Secret.SecretName
		if !mountableSecrets.Has(secretName) {
			return fmt.Errorf("volume with secret.secretName=\"%s\" is not allowed because service account %s does not reference that secret", secretName, serviceAccount.Name)
		}
	}
	for _, container := range pod.Spec.InitContainers {
		for _, env := range container.Env {
			if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
				if !mountableSecrets.Has(env.ValueFrom.SecretKeyRef.Name) {
					return fmt.Errorf("init container %s with envVar %s referencing secret.secretName=\"%s\" is not allowed because service account %s does not reference that secret", container.Name, env.Name, env.ValueFrom.SecretKeyRef.Name, serviceAccount.Name)
				}
			}
		}
	}
	for _, container := range pod.Spec.Containers {
		for _, env := range container.Env {
			if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
				if !mountableSecrets.Has(env.ValueFrom.SecretKeyRef.Name) {
					return fmt.Errorf("container %s with envVar %s referencing secret.secretName=\"%s\" is not allowed because service account %s does not reference that secret", container.Name, env.Name, env.ValueFrom.SecretKeyRef.Name, serviceAccount.Name)
				}
			}
		}
	}
	pullSecrets := sets.NewString()
	for _, s := range serviceAccount.ImagePullSecrets {
		pullSecrets.Insert(s.Name)
	}
	for i, pullSecretRef := range pod.Spec.ImagePullSecrets {
		if !pullSecrets.Has(pullSecretRef.Name) {
			return fmt.Errorf(`imagePullSecrets[%d].name="%s" is not allowed because service account %s does not reference that imagePullSecret`, i, pullSecretRef.Name, serviceAccount.Name)
		}
	}
	return nil
}
func (s *serviceAccount) mountServiceAccountToken(serviceAccount *corev1.ServiceAccount, pod *api.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	serviceAccountToken, err := s.getReferencedServiceAccountToken(serviceAccount)
	if err != nil {
		return fmt.Errorf("Error looking up service account token for %s/%s: %v", serviceAccount.Namespace, serviceAccount.Name, err)
	}
	if len(serviceAccountToken) == 0 {
		if s.RequireAPIToken {
			err := errors.NewServerTimeout(schema.GroupResource{Resource: "serviceaccounts"}, "create pod", 1)
			err.ErrStatus.Message = fmt.Sprintf("No API token found for service account %q, retry after the token is automatically created and added to the service account", serviceAccount.Name)
			return err
		}
		return nil
	}
	tokenVolumeName := ""
	hasTokenVolume := false
	allVolumeNames := sets.NewString()
	for _, volume := range pod.Spec.Volumes {
		allVolumeNames.Insert(volume.Name)
		if (!s.featureGate.Enabled(kubefeatures.BoundServiceAccountTokenVolume) && volume.Secret != nil && volume.Secret.SecretName == serviceAccountToken) || (s.featureGate.Enabled(kubefeatures.BoundServiceAccountTokenVolume) && strings.HasPrefix(volume.Name, ServiceAccountVolumeName+"-")) {
			tokenVolumeName = volume.Name
			hasTokenVolume = true
			break
		}
	}
	if len(tokenVolumeName) == 0 {
		if s.featureGate.Enabled(kubefeatures.BoundServiceAccountTokenVolume) {
			tokenVolumeName = s.generateName(ServiceAccountVolumeName + "-")
		} else {
			tokenVolumeName = serviceAccountToken
			if allVolumeNames.Has(tokenVolumeName) {
				tokenVolumeName = s.generateName(fmt.Sprintf("%s-", serviceAccountToken))
			}
		}
	}
	volumeMount := api.VolumeMount{Name: tokenVolumeName, ReadOnly: true, MountPath: DefaultAPITokenMountPath}
	needsTokenVolume := false
	for i, container := range pod.Spec.InitContainers {
		existingContainerMount := false
		for _, volumeMount := range container.VolumeMounts {
			if volumeMount.MountPath == DefaultAPITokenMountPath {
				existingContainerMount = true
				break
			}
		}
		if !existingContainerMount {
			pod.Spec.InitContainers[i].VolumeMounts = append(pod.Spec.InitContainers[i].VolumeMounts, volumeMount)
			needsTokenVolume = true
		}
	}
	for i, container := range pod.Spec.Containers {
		existingContainerMount := false
		for _, volumeMount := range container.VolumeMounts {
			if volumeMount.MountPath == DefaultAPITokenMountPath {
				existingContainerMount = true
				break
			}
		}
		if !existingContainerMount {
			pod.Spec.Containers[i].VolumeMounts = append(pod.Spec.Containers[i].VolumeMounts, volumeMount)
			needsTokenVolume = true
		}
	}
	if !hasTokenVolume && needsTokenVolume {
		pod.Spec.Volumes = append(pod.Spec.Volumes, s.createVolume(tokenVolumeName, serviceAccountToken))
	}
	return nil
}
func (s *serviceAccount) createVolume(tokenVolumeName, secretName string) api.Volume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.featureGate.Enabled(kubefeatures.BoundServiceAccountTokenVolume) {
		return api.Volume{Name: tokenVolumeName, VolumeSource: api.VolumeSource{Projected: &api.ProjectedVolumeSource{Sources: []api.VolumeProjection{{ServiceAccountToken: &api.ServiceAccountTokenProjection{Path: "token", ExpirationSeconds: 60 * 60}}, {ConfigMap: &api.ConfigMapProjection{LocalObjectReference: api.LocalObjectReference{Name: "kube-root-ca.crt"}, Items: []api.KeyToPath{{Key: "ca.crt", Path: "ca.crt"}}}}, {DownwardAPI: &api.DownwardAPIProjection{Items: []api.DownwardAPIVolumeFile{{Path: "namespace", FieldRef: &api.ObjectFieldSelector{APIVersion: "v1", FieldPath: "metadata.namespace"}}}}}}}}}
	}
	return api.Volume{Name: tokenVolumeName, VolumeSource: api.VolumeSource{Secret: &api.SecretVolumeSource{SecretName: secretName}}}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
