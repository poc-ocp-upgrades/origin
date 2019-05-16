package apiclient

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	core "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/clientcmd"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type ClientBackedDryRunGetter struct {
	client        clientset.Interface
	dynamicClient dynamic.Interface
}

var _ DryRunGetter = &ClientBackedDryRunGetter{}

func NewClientBackedDryRunGetter(config *rest.Config) (*ClientBackedDryRunGetter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &ClientBackedDryRunGetter{client: client, dynamicClient: dynamicClient}, nil
}
func NewClientBackedDryRunGetterFromKubeconfig(file string) (*ClientBackedDryRunGetter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := clientcmd.LoadFromFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load kubeconfig")
	}
	clientConfig, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create API client configuration from kubeconfig")
	}
	return NewClientBackedDryRunGetter(clientConfig)
}
func (clg *ClientBackedDryRunGetter) HandleGetAction(action core.GetAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	unstructuredObj, err := clg.dynamicClient.Resource(action.GetResource()).Namespace(action.GetNamespace()).Get(action.GetName(), metav1.GetOptions{})
	printIfNotExists(err)
	if err != nil {
		return true, nil, err
	}
	newObj, err := decodeUnstructuredIntoAPIObject(action, unstructuredObj)
	if err != nil {
		fmt.Printf("error after decode: %v %v\n", unstructuredObj, err)
		return true, nil, err
	}
	return true, newObj, err
}
func (clg *ClientBackedDryRunGetter) HandleListAction(action core.ListAction) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	listOpts := metav1.ListOptions{LabelSelector: action.GetListRestrictions().Labels.String(), FieldSelector: action.GetListRestrictions().Fields.String()}
	unstructuredList, err := clg.dynamicClient.Resource(action.GetResource()).Namespace(action.GetNamespace()).List(listOpts)
	if err != nil {
		return true, nil, err
	}
	newObj, err := decodeUnstructuredIntoAPIObject(action, unstructuredList)
	if err != nil {
		fmt.Printf("error after decode: %v %v\n", unstructuredList, err)
		return true, nil, err
	}
	return true, newObj, err
}
func (clg *ClientBackedDryRunGetter) Client() clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return clg.client
}
func decodeUnstructuredIntoAPIObject(action core.Action, unstructuredObj runtime.Unstructured) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objBytes, err := json.Marshal(unstructuredObj)
	if err != nil {
		return nil, err
	}
	newObj, err := runtime.Decode(clientsetscheme.Codecs.UniversalDecoder(action.GetResource().GroupVersion()), objBytes)
	if err != nil {
		return nil, err
	}
	return newObj, nil
}
func printIfNotExists(err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if apierrors.IsNotFound(err) {
		fmt.Println("[dryrun] The GET request didn't yield any result, the API Server returned a NotFound error.")
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
