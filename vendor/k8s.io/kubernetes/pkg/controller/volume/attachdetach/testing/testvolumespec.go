package testing

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/fake"
	core "k8s.io/client-go/testing"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	"time"
	gotime "time"
)

const TestPluginName = "kubernetes.io/testPlugin"

func GetTestVolumeSpec(volumeName string, diskName v1.UniqueVolumeName) *volume.Spec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &volume.Spec{Volume: &v1.Volume{Name: volumeName, VolumeSource: v1.VolumeSource{GCEPersistentDisk: &v1.GCEPersistentDiskVolumeSource{PDName: string(diskName), FSType: "fake", ReadOnly: false}}}, PersistentVolume: &v1.PersistentVolume{Spec: v1.PersistentVolumeSpec{AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}}}}
}

var extraPods *v1.PodList

func CreateTestClient() *fake.Clientset {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fakeClient := &fake.Clientset{}
	extraPods = &v1.PodList{}
	fakeClient.AddReactor("list", "pods", func(action core.Action) (handled bool, ret runtime.Object, err error) {
		obj := &v1.PodList{}
		podNamePrefix := "mypod"
		namespace := "mynamespace"
		for i := 0; i < 5; i++ {
			podName := fmt.Sprintf("%s-%d", podNamePrefix, i)
			pod := v1.Pod{Status: v1.PodStatus{Phase: v1.PodRunning}, ObjectMeta: metav1.ObjectMeta{Name: podName, UID: types.UID(podName), Namespace: namespace, Labels: map[string]string{"name": podName}}, Spec: v1.PodSpec{Containers: []v1.Container{{Name: "containerName", Image: "containerImage", VolumeMounts: []v1.VolumeMount{{Name: "volumeMountName", ReadOnly: false, MountPath: "/mnt"}}}}, Volumes: []v1.Volume{{Name: "volumeName", VolumeSource: v1.VolumeSource{GCEPersistentDisk: &v1.GCEPersistentDiskVolumeSource{PDName: "pdName", FSType: "ext4", ReadOnly: false}}}}, NodeName: "mynode"}}
			obj.Items = append(obj.Items, pod)
		}
		for _, pod := range extraPods.Items {
			obj.Items = append(obj.Items, pod)
		}
		return true, obj, nil
	})
	fakeClient.AddReactor("create", "pods", func(action core.Action) (handled bool, ret runtime.Object, err error) {
		createAction := action.(core.CreateAction)
		pod := createAction.GetObject().(*v1.Pod)
		extraPods.Items = append(extraPods.Items, *pod)
		return true, createAction.GetObject(), nil
	})
	fakeClient.AddReactor("list", "nodes", func(action core.Action) (handled bool, ret runtime.Object, err error) {
		obj := &v1.NodeList{}
		nodeNamePrefix := "mynode"
		for i := 0; i < 5; i++ {
			var nodeName string
			if i != 0 {
				nodeName = fmt.Sprintf("%s-%d", nodeNamePrefix, i)
			} else {
				nodeName = nodeNamePrefix
			}
			node := v1.Node{ObjectMeta: metav1.ObjectMeta{Name: nodeName, Labels: map[string]string{"name": nodeName}, Annotations: map[string]string{util.ControllerManagedAttachAnnotation: "true"}}, Status: v1.NodeStatus{VolumesAttached: []v1.AttachedVolume{{Name: TestPluginName + "/lostVolumeName", DevicePath: "fake/path"}}}}
			obj.Items = append(obj.Items, node)
		}
		return true, obj, nil
	})
	fakeWatch := watch.NewFake()
	fakeClient.AddWatchReactor("*", core.DefaultWatchReactor(fakeWatch, nil))
	return fakeClient
}
func NewPod(uid, name string) *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Pod{ObjectMeta: metav1.ObjectMeta{UID: types.UID(uid), Name: name, Namespace: name}}
}
func NewPodWithVolume(podName, volumeName, nodeName string) *v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Pod{ObjectMeta: metav1.ObjectMeta{UID: types.UID(podName), Name: podName, Namespace: "mynamespace", Labels: map[string]string{"name": podName}}, Spec: v1.PodSpec{Containers: []v1.Container{{Name: "containerName", Image: "containerImage", VolumeMounts: []v1.VolumeMount{{Name: "volumeMountName", ReadOnly: false, MountPath: "/mnt"}}}}, Volumes: []v1.Volume{{Name: volumeName, VolumeSource: v1.VolumeSource{GCEPersistentDisk: &v1.GCEPersistentDiskVolumeSource{PDName: "pdName", FSType: "ext4", ReadOnly: false}}}}, NodeName: nodeName}}
}

type TestPlugin struct {
	ErrorEncountered  bool
	attachedVolumeMap map[string][]string
	detachedVolumeMap map[string][]string
	pluginLock        *sync.RWMutex
}

func (plugin *TestPlugin) Init(host volume.VolumeHost) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (plugin *TestPlugin) GetPluginName() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return TestPluginName
}
func (plugin *TestPlugin) GetVolumeName(spec *volume.Spec) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.Lock()
	defer plugin.pluginLock.Unlock()
	if spec == nil {
		klog.Errorf("GetVolumeName called with nil volume spec")
		plugin.ErrorEncountered = true
	}
	return spec.Name(), nil
}
func (plugin *TestPlugin) CanSupport(spec *volume.Spec) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.Lock()
	defer plugin.pluginLock.Unlock()
	if spec == nil {
		klog.Errorf("CanSupport called with nil volume spec")
		plugin.ErrorEncountered = true
	}
	return true
}
func (plugin *TestPlugin) RequiresRemount() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (plugin *TestPlugin) NewMounter(spec *volume.Spec, podRef *v1.Pod, opts volume.VolumeOptions) (volume.Mounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.Lock()
	defer plugin.pluginLock.Unlock()
	if spec == nil {
		klog.Errorf("NewMounter called with nil volume spec")
		plugin.ErrorEncountered = true
	}
	return nil, nil
}
func (plugin *TestPlugin) NewUnmounter(name string, podUID types.UID) (volume.Unmounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (plugin *TestPlugin) ConstructVolumeSpec(volumeName, mountPath string) (*volume.Spec, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fakeVolume := &v1.Volume{Name: volumeName, VolumeSource: v1.VolumeSource{GCEPersistentDisk: &v1.GCEPersistentDiskVolumeSource{PDName: "pdName", FSType: "ext4", ReadOnly: false}}}
	return volume.NewSpecFromVolume(fakeVolume), nil
}
func (plugin *TestPlugin) NewAttacher() (volume.Attacher, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attacher := testPluginAttacher{ErrorEncountered: &plugin.ErrorEncountered, attachedVolumeMap: plugin.attachedVolumeMap, pluginLock: plugin.pluginLock}
	return &attacher, nil
}
func (plugin *TestPlugin) NewDeviceMounter() (volume.DeviceMounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return plugin.NewAttacher()
}
func (plugin *TestPlugin) NewDetacher() (volume.Detacher, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	detacher := testPluginDetacher{detachedVolumeMap: plugin.detachedVolumeMap, pluginLock: plugin.pluginLock}
	return &detacher, nil
}
func (plugin *TestPlugin) NewDeviceUnmounter() (volume.DeviceUnmounter, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return plugin.NewDetacher()
}
func (plugin *TestPlugin) GetDeviceMountRefs(deviceMountPath string) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{}, nil
}
func (plugin *TestPlugin) SupportsMountOption() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (plugin *TestPlugin) SupportsBulkVolumeVerification() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (plugin *TestPlugin) GetErrorEncountered() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.RLock()
	defer plugin.pluginLock.RUnlock()
	return plugin.ErrorEncountered
}
func (plugin *TestPlugin) GetAttachedVolumes() map[string][]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.RLock()
	defer plugin.pluginLock.RUnlock()
	ret := make(map[string][]string)
	for nodeName, volumeList := range plugin.attachedVolumeMap {
		ret[nodeName] = make([]string, len(volumeList))
		copy(ret[nodeName], volumeList)
	}
	return ret
}
func (plugin *TestPlugin) GetDetachedVolumes() map[string][]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugin.pluginLock.RLock()
	defer plugin.pluginLock.RUnlock()
	ret := make(map[string][]string)
	for nodeName, volumeList := range plugin.detachedVolumeMap {
		ret[nodeName] = make([]string, len(volumeList))
		copy(ret[nodeName], volumeList)
	}
	return ret
}
func CreateTestPlugin() []volume.VolumePlugin {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attachedVolumes := make(map[string][]string)
	detachedVolumes := make(map[string][]string)
	return []volume.VolumePlugin{&TestPlugin{ErrorEncountered: false, attachedVolumeMap: attachedVolumes, detachedVolumeMap: detachedVolumes, pluginLock: &sync.RWMutex{}}}
}

type testPluginAttacher struct {
	ErrorEncountered  *bool
	attachedVolumeMap map[string][]string
	pluginLock        *sync.RWMutex
}

func (attacher *testPluginAttacher) Attach(spec *volume.Spec, nodeName types.NodeName) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attacher.pluginLock.Lock()
	defer attacher.pluginLock.Unlock()
	if spec == nil {
		*attacher.ErrorEncountered = true
		klog.Errorf("Attach called with nil volume spec")
		return "", fmt.Errorf("Attach called with nil volume spec")
	}
	attacher.attachedVolumeMap[string(nodeName)] = append(attacher.attachedVolumeMap[string(nodeName)], spec.Name())
	return spec.Name(), nil
}
func (attacher *testPluginAttacher) VolumesAreAttached(specs []*volume.Spec, nodeName types.NodeName) (map[*volume.Spec]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil, nil
}
func (attacher *testPluginAttacher) WaitForAttach(spec *volume.Spec, devicePath string, pod *v1.Pod, timeout time.Duration) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attacher.pluginLock.Lock()
	defer attacher.pluginLock.Unlock()
	if spec == nil {
		*attacher.ErrorEncountered = true
		klog.Errorf("WaitForAttach called with nil volume spec")
		return "", fmt.Errorf("WaitForAttach called with nil volume spec")
	}
	fakePath := fmt.Sprintf("%s/%s", devicePath, spec.Name())
	return fakePath, nil
}
func (attacher *testPluginAttacher) GetDeviceMountPath(spec *volume.Spec) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attacher.pluginLock.Lock()
	defer attacher.pluginLock.Unlock()
	if spec == nil {
		*attacher.ErrorEncountered = true
		klog.Errorf("GetDeviceMountPath called with nil volume spec")
		return "", fmt.Errorf("GetDeviceMountPath called with nil volume spec")
	}
	return "", nil
}
func (attacher *testPluginAttacher) MountDevice(spec *volume.Spec, devicePath string, deviceMountPath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attacher.pluginLock.Lock()
	defer attacher.pluginLock.Unlock()
	if spec == nil {
		*attacher.ErrorEncountered = true
		klog.Errorf("MountDevice called with nil volume spec")
		return fmt.Errorf("MountDevice called with nil volume spec")
	}
	return nil
}

type testPluginDetacher struct {
	detachedVolumeMap map[string][]string
	pluginLock        *sync.RWMutex
}

func (detacher *testPluginDetacher) Detach(volumeName string, nodeName types.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	detacher.pluginLock.Lock()
	defer detacher.pluginLock.Unlock()
	detacher.detachedVolumeMap[string(nodeName)] = append(detacher.detachedVolumeMap[string(nodeName)], volumeName)
	return nil
}
func (detacher *testPluginDetacher) UnmountDevice(deviceMountPath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
