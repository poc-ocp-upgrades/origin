package testutil

import (
	"encoding/json"
	"errors"
	"fmt"
	goformat "fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/fake"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	ref "k8s.io/client-go/tools/reference"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	api "k8s.io/kubernetes/pkg/apis/core"
	utilnode "k8s.io/kubernetes/pkg/util/node"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"sync"
	"testing"
	"time"
	gotime "time"
)

var (
	keyFunc = cache.DeletionHandlingMetaNamespaceKeyFunc
)

type FakeNodeHandler struct {
	*fake.Clientset
	CreateHook          func(*FakeNodeHandler, *v1.Node) bool
	Existing            []*v1.Node
	CreatedNodes        []*v1.Node
	DeletedNodes        []*v1.Node
	UpdatedNodes        []*v1.Node
	UpdatedNodeStatuses []*v1.Node
	RequestCount        int
	lock                sync.Mutex
	DeleteWaitChan      chan struct{}
	PatchWaitChan       chan struct{}
}
type FakeLegacyHandler struct {
	v1core.CoreV1Interface
	n *FakeNodeHandler
}

func (m *FakeNodeHandler) GetUpdatedNodesCopy() []*v1.Node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer m.lock.Unlock()
	updatedNodesCopy := make([]*v1.Node, len(m.UpdatedNodes), len(m.UpdatedNodes))
	for i, ptr := range m.UpdatedNodes {
		updatedNodesCopy[i] = ptr
	}
	return updatedNodesCopy
}
func (m *FakeNodeHandler) Core() v1core.CoreV1Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FakeLegacyHandler{m.Clientset.Core(), m}
}
func (m *FakeNodeHandler) CoreV1() v1core.CoreV1Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FakeLegacyHandler{m.Clientset.CoreV1(), m}
}
func (m *FakeLegacyHandler) Nodes() v1core.NodeInterface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return m.n
}
func (m *FakeNodeHandler) Create(node *v1.Node) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		m.lock.Unlock()
	}()
	for _, n := range m.Existing {
		if n.Name == node.Name {
			return nil, apierrors.NewAlreadyExists(api.Resource("nodes"), node.Name)
		}
	}
	if m.CreateHook == nil || m.CreateHook(m, node) {
		nodeCopy := *node
		m.CreatedNodes = append(m.CreatedNodes, &nodeCopy)
		return node, nil
	}
	return nil, errors.New("create error")
}
func (m *FakeNodeHandler) Get(name string, opts metav1.GetOptions) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		m.lock.Unlock()
	}()
	for i := range m.UpdatedNodes {
		if m.UpdatedNodes[i].Name == name {
			nodeCopy := *m.UpdatedNodes[i]
			return &nodeCopy, nil
		}
	}
	for i := range m.Existing {
		if m.Existing[i].Name == name {
			nodeCopy := *m.Existing[i]
			return &nodeCopy, nil
		}
	}
	return nil, nil
}
func (m *FakeNodeHandler) List(opts metav1.ListOptions) (*v1.NodeList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		m.lock.Unlock()
	}()
	var nodes []*v1.Node
	for i := 0; i < len(m.UpdatedNodes); i++ {
		if !contains(m.UpdatedNodes[i], m.DeletedNodes) {
			nodes = append(nodes, m.UpdatedNodes[i])
		}
	}
	for i := 0; i < len(m.Existing); i++ {
		if !contains(m.Existing[i], m.DeletedNodes) && !contains(m.Existing[i], nodes) {
			nodes = append(nodes, m.Existing[i])
		}
	}
	for i := 0; i < len(m.CreatedNodes); i++ {
		if !contains(m.CreatedNodes[i], m.DeletedNodes) && !contains(m.CreatedNodes[i], nodes) {
			nodes = append(nodes, m.CreatedNodes[i])
		}
	}
	nodeList := &v1.NodeList{}
	for _, node := range nodes {
		nodeList.Items = append(nodeList.Items, *node)
	}
	return nodeList, nil
}
func (m *FakeNodeHandler) Delete(id string, opt *metav1.DeleteOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		if m.DeleteWaitChan != nil {
			m.DeleteWaitChan <- struct{}{}
		}
		m.lock.Unlock()
	}()
	m.DeletedNodes = append(m.DeletedNodes, NewNode(id))
	return nil
}
func (m *FakeNodeHandler) DeleteCollection(opt *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (m *FakeNodeHandler) Update(node *v1.Node) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		m.lock.Unlock()
	}()
	nodeCopy := *node
	for i, updateNode := range m.UpdatedNodes {
		if updateNode.Name == nodeCopy.Name {
			m.UpdatedNodes[i] = &nodeCopy
			return node, nil
		}
	}
	m.UpdatedNodes = append(m.UpdatedNodes, &nodeCopy)
	return node, nil
}
func (m *FakeNodeHandler) UpdateStatus(node *v1.Node) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		m.lock.Unlock()
	}()
	var origNodeCopy v1.Node
	found := false
	for i := range m.Existing {
		if m.Existing[i].Name == node.Name {
			origNodeCopy = *m.Existing[i]
			found = true
			break
		}
	}
	updatedNodeIndex := -1
	for i := range m.UpdatedNodes {
		if m.UpdatedNodes[i].Name == node.Name {
			origNodeCopy = *m.UpdatedNodes[i]
			updatedNodeIndex = i
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("Not found node %v", node)
	}
	origNodeCopy.Status = node.Status
	if updatedNodeIndex < 0 {
		m.UpdatedNodes = append(m.UpdatedNodes, &origNodeCopy)
	} else {
		m.UpdatedNodes[updatedNodeIndex] = &origNodeCopy
	}
	nodeCopy := *node
	m.UpdatedNodeStatuses = append(m.UpdatedNodeStatuses, &nodeCopy)
	return node, nil
}
func (m *FakeNodeHandler) PatchStatus(nodeName string, data []byte) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.RequestCount++
	return m.Patch(nodeName, types.StrategicMergePatchType, data, "status")
}
func (m *FakeNodeHandler) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return watch.NewFake(), nil
}
func (m *FakeNodeHandler) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Node, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	m.lock.Lock()
	defer func() {
		m.RequestCount++
		if m.PatchWaitChan != nil {
			m.PatchWaitChan <- struct{}{}
		}
		m.lock.Unlock()
	}()
	var nodeCopy v1.Node
	for i := range m.Existing {
		if m.Existing[i].Name == name {
			nodeCopy = *m.Existing[i]
		}
	}
	updatedNodeIndex := -1
	for i := range m.UpdatedNodes {
		if m.UpdatedNodes[i].Name == name {
			nodeCopy = *m.UpdatedNodes[i]
			updatedNodeIndex = i
		}
	}
	originalObjJS, err := json.Marshal(nodeCopy)
	if err != nil {
		klog.Errorf("Failed to marshal %v", nodeCopy)
		return nil, nil
	}
	var originalNode v1.Node
	if err = json.Unmarshal(originalObjJS, &originalNode); err != nil {
		klog.Errorf("Failed to unmarshal original object: %v", err)
		return nil, nil
	}
	var patchedObjJS []byte
	switch pt {
	case types.JSONPatchType:
		patchObj, err := jsonpatch.DecodePatch(data)
		if err != nil {
			klog.Error(err.Error())
			return nil, nil
		}
		if patchedObjJS, err = patchObj.Apply(originalObjJS); err != nil {
			klog.Error(err.Error())
			return nil, nil
		}
	case types.MergePatchType:
		if patchedObjJS, err = jsonpatch.MergePatch(originalObjJS, data); err != nil {
			klog.Error(err.Error())
			return nil, nil
		}
	case types.StrategicMergePatchType:
		if patchedObjJS, err = strategicpatch.StrategicMergePatch(originalObjJS, data, originalNode); err != nil {
			klog.Error(err.Error())
			return nil, nil
		}
	default:
		klog.Errorf("unknown Content-Type header for patch: %v", pt)
		return nil, nil
	}
	var updatedNode v1.Node
	if err = json.Unmarshal(patchedObjJS, &updatedNode); err != nil {
		klog.Errorf("Failed to unmarshal patched object: %v", err)
		return nil, nil
	}
	if updatedNodeIndex < 0 {
		m.UpdatedNodes = append(m.UpdatedNodes, &updatedNode)
	} else {
		m.UpdatedNodes[updatedNodeIndex] = &updatedNode
	}
	return &updatedNode, nil
}

type FakeRecorder struct {
	sync.Mutex
	source v1.EventSource
	Events []*v1.Event
	clock  clock.Clock
}

func (f *FakeRecorder) Event(obj runtime.Object, eventtype, reason, message string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.generateEvent(obj, metav1.Now(), eventtype, reason, message)
}
func (f *FakeRecorder) Eventf(obj runtime.Object, eventtype, reason, messageFmt string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.Event(obj, eventtype, reason, fmt.Sprintf(messageFmt, args...))
}
func (f *FakeRecorder) PastEventf(obj runtime.Object, timestamp metav1.Time, eventtype, reason, messageFmt string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (f *FakeRecorder) AnnotatedEventf(obj runtime.Object, annotations map[string]string, eventtype, reason, messageFmt string, args ...interface{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.Eventf(obj, eventtype, reason, messageFmt, args)
}
func (f *FakeRecorder) generateEvent(obj runtime.Object, timestamp metav1.Time, eventtype, reason, message string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.Lock()
	defer f.Unlock()
	ref, err := ref.GetReference(legacyscheme.Scheme, obj)
	if err != nil {
		klog.Errorf("Encountered error while getting reference: %v", err)
		return
	}
	event := f.makeEvent(ref, eventtype, reason, message)
	event.Source = f.source
	if f.Events != nil {
		f.Events = append(f.Events, event)
	}
}
func (f *FakeRecorder) makeEvent(ref *v1.ObjectReference, eventtype, reason, message string) *v1.Event {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	t := metav1.Time{Time: f.clock.Now()}
	namespace := ref.Namespace
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}
	clientref := v1.ObjectReference{Kind: ref.Kind, Namespace: ref.Namespace, Name: ref.Name, UID: ref.UID, APIVersion: ref.APIVersion, ResourceVersion: ref.ResourceVersion, FieldPath: ref.FieldPath}
	return &v1.Event{ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("%v.%x", ref.Name, t.UnixNano()), Namespace: namespace}, InvolvedObject: clientref, Reason: reason, Message: message, FirstTimestamp: t, LastTimestamp: t, Count: 1, Type: eventtype}
}
func NewFakeRecorder() *FakeRecorder {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &FakeRecorder{source: v1.EventSource{Component: "nodeControllerTest"}, Events: []*v1.Event{}, clock: clock.NewFakeClock(time.Now())}
}
func NewNode(name string) *v1.Node {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Node{ObjectMeta: metav1.ObjectMeta{Name: name}, Status: v1.NodeStatus{Capacity: v1.ResourceList{v1.ResourceName(v1.ResourceCPU): resource.MustParse("10"), v1.ResourceName(v1.ResourceMemory): resource.MustParse("10G")}}}
}
func NewPod(name, host string) *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: name}, Spec: v1.PodSpec{NodeName: host}, Status: v1.PodStatus{Conditions: []v1.PodCondition{{Type: v1.PodReady, Status: v1.ConditionTrue}}}}
	return pod
}
func contains(node *v1.Node, nodes []*v1.Node) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := 0; i < len(nodes); i++ {
		if node.Name == nodes[i].Name {
			return true
		}
	}
	return false
}
func GetZones(nodeHandler *FakeNodeHandler) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodes, _ := nodeHandler.List(metav1.ListOptions{})
	zones := sets.NewString()
	for _, node := range nodes.Items {
		zones.Insert(utilnode.GetZoneKey(&node))
	}
	return zones.List()
}
func CreateZoneID(region, zone string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return region + ":\x00:" + zone
}
func GetKey(obj interface{}, t *testing.T) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
	if ok {
		obj = tombstone.Obj
	}
	val := reflect.ValueOf(obj).Elem()
	name := val.FieldByName("Name").String()
	kind := val.FieldByName("Kind").String()
	if len(name) == 0 || len(kind) == 0 {
		t.Errorf("Unexpected object %v", obj)
	}
	key, err := keyFunc(obj)
	if err != nil {
		t.Errorf("Unexpected error getting key for %v %v: %v", kind, name, err)
		return ""
	}
	return key
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
