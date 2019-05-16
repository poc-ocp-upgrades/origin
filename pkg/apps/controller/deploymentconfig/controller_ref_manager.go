package deploymentconfig

import (
	"fmt"
	goformat "fmt"
	"k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	klabels "k8s.io/apimachinery/pkg/labels"
	kschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kutilerrors "k8s.io/apimachinery/pkg/util/errors"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	kcontroller "k8s.io/kubernetes/pkg/controller"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type RCControlInterface interface {
	PatchReplicationController(namespace, name string, data []byte) error
}
type RealRCControl struct {
	KubeClient kclientset.Interface
	Recorder   record.EventRecorder
}

var _ RCControlInterface = &RealRCControl{}

func (r RealRCControl) PatchReplicationController(namespace, name string, data []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := r.KubeClient.CoreV1().ReplicationControllers(namespace).Patch(name, types.StrategicMergePatchType, data)
	return err
}

type RCControllerRefManager struct {
	kcontroller.BaseControllerRefManager
	controllerKind kschema.GroupVersionKind
	rcControl      RCControlInterface
}

func NewRCControllerRefManager(rcControl RCControlInterface, controller kmetav1.Object, selector klabels.Selector, controllerKind kschema.GroupVersionKind, canAdopt func() error) *RCControllerRefManager {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &RCControllerRefManager{BaseControllerRefManager: kcontroller.BaseControllerRefManager{Controller: controller, Selector: selector, CanAdoptFunc: canAdopt}, controllerKind: controllerKind, rcControl: rcControl}
}
func (m *RCControllerRefManager) ClaimReplicationController(rc *v1.ReplicationController) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	match := func(obj kmetav1.Object) bool {
		return m.Selector.Matches(klabels.Set(obj.GetLabels()))
	}
	adopt := func(obj kmetav1.Object) error {
		return m.AdoptReplicationController(obj.(*v1.ReplicationController))
	}
	release := func(obj kmetav1.Object) error {
		return m.ReleaseReplicationController(obj.(*v1.ReplicationController))
	}
	return m.ClaimObject(rc, match, adopt, release)
}
func (m *RCControllerRefManager) ClaimReplicationControllers(rcs []*v1.ReplicationController) ([]*v1.ReplicationController, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var claimed []*v1.ReplicationController
	var errlist []error
	for _, rc := range rcs {
		ok, err := m.ClaimReplicationController(rc)
		if err != nil {
			errlist = append(errlist, err)
			continue
		}
		if ok {
			claimed = append(claimed, rc)
		}
	}
	return claimed, kutilerrors.NewAggregate(errlist)
}
func (m *RCControllerRefManager) AdoptReplicationController(rs *v1.ReplicationController) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := m.CanAdopt(); err != nil {
		return fmt.Errorf("can't adopt ReplicationController %s/%s (%s): %v", rs.Namespace, rs.Name, rs.UID, err)
	}
	addControllerPatch := fmt.Sprintf(`{"metadata":{
			"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],
			"uid":"%s"
			}
		}`, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName(), m.Controller.GetUID(), rs.UID)
	return m.rcControl.PatchReplicationController(rs.Namespace, rs.Name, []byte(addControllerPatch))
}
func (m *RCControllerRefManager) ReleaseReplicationController(rc *v1.ReplicationController) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(4).Infof("patching ReplicationController %s/%s to remove its controllerRef to %s/%s:%s", rc.Namespace, rc.Name, m.controllerKind.GroupVersion(), m.controllerKind.Kind, m.Controller.GetName())
	deleteOwnerRefPatch := fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, m.Controller.GetUID(), rc.UID)
	err := m.rcControl.PatchReplicationController(rc.Namespace, rc.Name, []byte(deleteOwnerRefPatch))
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil
		}
		if kerrors.IsInvalid(err) {
			return nil
		}
	}
	return err
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
