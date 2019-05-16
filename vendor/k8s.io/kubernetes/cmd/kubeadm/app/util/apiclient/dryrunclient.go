package apiclient

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientset "k8s.io/client-go/kubernetes"
	fakeclientset "k8s.io/client-go/kubernetes/fake"
	core "k8s.io/client-go/testing"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"strings"
)

type DryRunGetter interface {
	HandleGetAction(core.GetAction) (bool, runtime.Object, error)
	HandleListAction(core.ListAction) (bool, runtime.Object, error)
}
type MarshalFunc func(runtime.Object, schema.GroupVersion) ([]byte, error)

func DefaultMarshalFunc(obj runtime.Object, gv schema.GroupVersion) ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return kubeadmutil.MarshalToYaml(obj, gv)
}

type DryRunClientOptions struct {
	Writer          io.Writer
	Getter          DryRunGetter
	PrependReactors []core.Reactor
	AppendReactors  []core.Reactor
	MarshalFunc     MarshalFunc
	PrintGETAndLIST bool
}

func GetDefaultDryRunClientOptions(drg DryRunGetter, w io.Writer) DryRunClientOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return DryRunClientOptions{Writer: w, Getter: drg, PrependReactors: []core.Reactor{}, AppendReactors: []core.Reactor{}, MarshalFunc: DefaultMarshalFunc, PrintGETAndLIST: false}
}

type actionWithName interface {
	core.Action
	GetName() string
}
type actionWithObject interface {
	core.Action
	GetObject() runtime.Object
}

func NewDryRunClient(drg DryRunGetter, w io.Writer) clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewDryRunClientWithOpts(GetDefaultDryRunClientOptions(drg, w))
}
func NewDryRunClientWithOpts(opts DryRunClientOptions) clientset.Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	client := fakeclientset.NewSimpleClientset()
	defaultReactorChain := []core.Reactor{&core.SimpleReactor{Verb: "*", Resource: "*", Reaction: func(action core.Action) (bool, runtime.Object, error) {
		logDryRunAction(action, opts.Writer, opts.MarshalFunc)
		return false, nil, nil
	}}, &core.SimpleReactor{Verb: "get", Resource: "*", Reaction: func(action core.Action) (bool, runtime.Object, error) {
		getAction, ok := action.(core.GetAction)
		if !ok {
			return true, nil, errors.New("can't cast get reactor event action object to GetAction interface")
		}
		handled, obj, err := opts.Getter.HandleGetAction(getAction)
		if opts.PrintGETAndLIST {
			objBytes, err := opts.MarshalFunc(obj, action.GetResource().GroupVersion())
			if err == nil {
				fmt.Println("[dryrun] Returning faked GET response:")
				PrintBytesWithLinePrefix(opts.Writer, objBytes, "\t")
			}
		}
		return handled, obj, err
	}}, &core.SimpleReactor{Verb: "list", Resource: "*", Reaction: func(action core.Action) (bool, runtime.Object, error) {
		listAction, ok := action.(core.ListAction)
		if !ok {
			return true, nil, errors.New("can't cast list reactor event action object to ListAction interface")
		}
		handled, objs, err := opts.Getter.HandleListAction(listAction)
		if opts.PrintGETAndLIST {
			objBytes, err := opts.MarshalFunc(objs, action.GetResource().GroupVersion())
			if err == nil {
				fmt.Println("[dryrun] Returning faked LIST response:")
				PrintBytesWithLinePrefix(opts.Writer, objBytes, "\t")
			}
		}
		return handled, objs, err
	}}, &core.SimpleReactor{Verb: "create", Resource: "*", Reaction: successfulModificationReactorFunc}, &core.SimpleReactor{Verb: "update", Resource: "*", Reaction: successfulModificationReactorFunc}, &core.SimpleReactor{Verb: "delete", Resource: "*", Reaction: successfulModificationReactorFunc}, &core.SimpleReactor{Verb: "delete-collection", Resource: "*", Reaction: successfulModificationReactorFunc}, &core.SimpleReactor{Verb: "patch", Resource: "*", Reaction: successfulModificationReactorFunc}}
	fullReactorChain := append(opts.PrependReactors, defaultReactorChain...)
	fullReactorChain = append(fullReactorChain, opts.AppendReactors...)
	client.Fake.ReactionChain = append(fullReactorChain, client.Fake.ReactionChain...)
	return client
}
func successfulModificationReactorFunc(action core.Action) (bool, runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objAction, ok := action.(actionWithObject)
	if ok {
		return true, objAction.GetObject(), nil
	}
	return true, nil, nil
}
func logDryRunAction(action core.Action, w io.Writer, marshalFunc MarshalFunc) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	group := action.GetResource().Group
	if len(group) == 0 {
		group = "core"
	}
	fmt.Fprintf(w, "[dryrun] Would perform action %s on resource %q in API group \"%s/%s\"\n", strings.ToUpper(action.GetVerb()), action.GetResource().Resource, group, action.GetResource().Version)
	namedAction, ok := action.(actionWithName)
	if ok {
		fmt.Fprintf(w, "[dryrun] Resource name: %q\n", namedAction.GetName())
	}
	objAction, ok := action.(actionWithObject)
	if ok && objAction.GetObject() != nil {
		objBytes, err := marshalFunc(objAction.GetObject(), action.GetResource().GroupVersion())
		if err == nil {
			fmt.Println("[dryrun] Attached object:")
			PrintBytesWithLinePrefix(w, objBytes, "\t")
		}
	}
	patchAction, ok := action.(core.PatchAction)
	if ok {
		fmt.Fprintf(w, "[dryrun] Attached patch:\n\t%s\n", strings.Replace(string(patchAction.GetPatch()), `\"`, `"`, -1))
	}
}
func PrintBytesWithLinePrefix(w io.Writer, objBytes []byte, linePrefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	scanner := bufio.NewScanner(bytes.NewReader(objBytes))
	for scanner.Scan() {
		fmt.Fprintf(w, "%s%s\n", linePrefix, scanner.Text())
	}
}
