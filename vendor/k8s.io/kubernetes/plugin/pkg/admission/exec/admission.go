package exec

import (
	"fmt"
	goformat "fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/kubernetes"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const (
	DenyEscalatingExec   = "DenyEscalatingExec"
	DenyExecOnPrivileged = "DenyExecOnPrivileged"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(DenyEscalatingExec, func(config io.Reader) (admission.Interface, error) {
		return NewDenyEscalatingExec(), nil
	})
	plugins.Register(DenyExecOnPrivileged, func(config io.Reader) (admission.Interface, error) {
		return NewDenyExecOnPrivileged(), nil
	})
}

type DenyExec struct {
	*admission.Handler
	client      kubernetes.Interface
	hostNetwork bool
	hostIPC     bool
	hostPID     bool
	privileged  bool
}

var _ admission.ValidationInterface = &DenyExec{}
var _ = genericadmissioninitializer.WantsExternalKubeClientSet(&DenyExec{})

func NewDenyExecOnPrivileged() *DenyExec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &DenyExec{Handler: admission.NewHandler(admission.Connect), hostNetwork: false, hostIPC: false, hostPID: false, privileged: true}
}
func NewDenyEscalatingExec() *DenyExec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &DenyExec{Handler: admission.NewHandler(admission.Connect), hostNetwork: true, hostIPC: true, hostPID: true, privileged: true}
}
func (d *DenyExec) SetExternalKubeClientSet(client kubernetes.Interface) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	d.client = client
}
func (d *DenyExec) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if d.client == nil {
		return fmt.Errorf("missing client")
	}
	return nil
}
func (d *DenyExec) Validate(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	path := a.GetResource().Resource
	if subresource := a.GetSubresource(); subresource != "" {
		path = path + "/" + subresource
	}
	if path != "pods/exec" && path != "pods/attach" {
		return nil
	}
	pod, err := d.client.CoreV1().Pods(a.GetNamespace()).Get(a.GetName(), metav1.GetOptions{})
	if err != nil {
		return admission.NewForbidden(a, err)
	}
	if d.hostNetwork && pod.Spec.HostNetwork {
		return admission.NewForbidden(a, fmt.Errorf("cannot exec into or attach to a container using host network"))
	}
	if d.hostPID && pod.Spec.HostPID {
		return admission.NewForbidden(a, fmt.Errorf("cannot exec into or attach to a container using host pid"))
	}
	if d.hostIPC && pod.Spec.HostIPC {
		return admission.NewForbidden(a, fmt.Errorf("cannot exec into or attach to a container using host ipc"))
	}
	if d.privileged && isPrivileged(pod) {
		return admission.NewForbidden(a, fmt.Errorf("cannot exec into or attach to a privileged container"))
	}
	return nil
}
func isPrivileged(pod *corev1.Pod) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, c := range pod.Spec.InitContainers {
		if c.SecurityContext == nil || c.SecurityContext.Privileged == nil {
			continue
		}
		if *c.SecurityContext.Privileged {
			return true
		}
	}
	for _, c := range pod.Spec.Containers {
		if c.SecurityContext == nil || c.SecurityContext.Privileged == nil {
			continue
		}
		if *c.SecurityContext.Privileged {
			return true
		}
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
