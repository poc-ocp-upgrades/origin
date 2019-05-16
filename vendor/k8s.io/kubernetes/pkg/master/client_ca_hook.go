package master

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	genericapiserver "k8s.io/apiserver/pkg/server"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type ClientCARegistrationHook struct {
	ClientCA                         []byte
	RequestHeaderUsernameHeaders     []string
	RequestHeaderGroupHeaders        []string
	RequestHeaderExtraHeaderPrefixes []string
	RequestHeaderCA                  []byte
	RequestHeaderAllowedNames        []string
}

func (h ClientCARegistrationHook) PostStartHook(hookContext genericapiserver.PostStartHookContext) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := wait.Poll(1*time.Second, 30*time.Second, func() (done bool, err error) {
		client, err := corev1client.NewForConfig(hookContext.LoopbackClientConfig)
		if err != nil {
			utilruntime.HandleError(err)
			return false, nil
		}
		return h.tryToWriteClientCAs(client)
	})
	if err != nil {
		return fmt.Errorf("unable to initialize client CA configmap: %v", err)
	}
	return nil
}
func (h ClientCARegistrationHook) tryToWriteClientCAs(client corev1client.CoreV1Interface) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := createNamespaceIfNeeded(client, metav1.NamespaceSystem); err != nil {
		utilruntime.HandleError(err)
		return false, nil
	}
	data := map[string]string{}
	if len(h.ClientCA) > 0 {
		data["client-ca-file"] = string(h.ClientCA)
	}
	if len(h.RequestHeaderCA) > 0 {
		var err error
		data["requestheader-username-headers"], err = jsonSerializeStringSlice(h.RequestHeaderUsernameHeaders)
		if err != nil {
			return false, err
		}
		data["requestheader-group-headers"], err = jsonSerializeStringSlice(h.RequestHeaderGroupHeaders)
		if err != nil {
			return false, err
		}
		data["requestheader-extra-headers-prefix"], err = jsonSerializeStringSlice(h.RequestHeaderExtraHeaderPrefixes)
		if err != nil {
			return false, err
		}
		data["requestheader-client-ca-file"] = string(h.RequestHeaderCA)
		data["requestheader-allowed-names"], err = jsonSerializeStringSlice(h.RequestHeaderAllowedNames)
		if err != nil {
			return false, err
		}
	}
	if err := writeConfigMap(client, "extension-apiserver-authentication", data); err != nil {
		utilruntime.HandleError(err)
		return false, nil
	}
	return true, nil
}
func jsonSerializeStringSlice(in []string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	out, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(out), err
}
func writeConfigMap(client corev1client.ConfigMapsGetter, name string, data map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	existing, err := client.ConfigMaps(metav1.NamespaceSystem).Get(name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		_, err := client.ConfigMaps(metav1.NamespaceSystem).Create(&corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: metav1.NamespaceSystem, Name: name}, Data: data})
		return err
	}
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(existing.Data, data) {
		existing.Data = data
		_, err = client.ConfigMaps(metav1.NamespaceSystem).Update(existing)
	}
	return err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
