package securitycontext

import (
	goformat "fmt"
	api "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	gotime "time"
)

type PodSecurityContextAccessor interface {
	HostNetwork() bool
	HostPID() bool
	HostIPC() bool
	SELinuxOptions() *api.SELinuxOptions
	RunAsUser() *int64
	RunAsGroup() *int64
	RunAsNonRoot() *bool
	SupplementalGroups() []int64
	FSGroup() *int64
}
type PodSecurityContextMutator interface {
	PodSecurityContextAccessor
	SetHostNetwork(bool)
	SetHostPID(bool)
	SetHostIPC(bool)
	SetSELinuxOptions(*api.SELinuxOptions)
	SetRunAsUser(*int64)
	SetRunAsGroup(*int64)
	SetRunAsNonRoot(*bool)
	SetSupplementalGroups([]int64)
	SetFSGroup(*int64)
	PodSecurityContext() *api.PodSecurityContext
}

func NewPodSecurityContextAccessor(podSC *api.PodSecurityContext) PodSecurityContextAccessor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &podSecurityContextWrapper{podSC: podSC}
}
func NewPodSecurityContextMutator(podSC *api.PodSecurityContext) PodSecurityContextMutator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &podSecurityContextWrapper{podSC: podSC}
}

type podSecurityContextWrapper struct{ podSC *api.PodSecurityContext }

func (w *podSecurityContextWrapper) PodSecurityContext() *api.PodSecurityContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.podSC
}
func (w *podSecurityContextWrapper) ensurePodSC() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		w.podSC = &api.PodSecurityContext{}
	}
}
func (w *podSecurityContextWrapper) HostNetwork() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return false
	}
	return w.podSC.HostNetwork
}
func (w *podSecurityContextWrapper) SetHostNetwork(v bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == false {
		return
	}
	w.ensurePodSC()
	w.podSC.HostNetwork = v
}
func (w *podSecurityContextWrapper) HostPID() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return false
	}
	return w.podSC.HostPID
}
func (w *podSecurityContextWrapper) SetHostPID(v bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == false {
		return
	}
	w.ensurePodSC()
	w.podSC.HostPID = v
}
func (w *podSecurityContextWrapper) HostIPC() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return false
	}
	return w.podSC.HostIPC
}
func (w *podSecurityContextWrapper) SetHostIPC(v bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == false {
		return
	}
	w.ensurePodSC()
	w.podSC.HostIPC = v
}
func (w *podSecurityContextWrapper) SELinuxOptions() *api.SELinuxOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.SELinuxOptions
}
func (w *podSecurityContextWrapper) SetSELinuxOptions(v *api.SELinuxOptions) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == nil {
		return
	}
	w.ensurePodSC()
	w.podSC.SELinuxOptions = v
}
func (w *podSecurityContextWrapper) RunAsUser() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.RunAsUser
}
func (w *podSecurityContextWrapper) SetRunAsUser(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == nil {
		return
	}
	w.ensurePodSC()
	w.podSC.RunAsUser = v
}
func (w *podSecurityContextWrapper) RunAsGroup() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.RunAsGroup
}
func (w *podSecurityContextWrapper) SetRunAsGroup(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == nil {
		return
	}
	w.ensurePodSC()
	w.podSC.RunAsGroup = v
}
func (w *podSecurityContextWrapper) RunAsNonRoot() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.RunAsNonRoot
}
func (w *podSecurityContextWrapper) SetRunAsNonRoot(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == nil {
		return
	}
	w.ensurePodSC()
	w.podSC.RunAsNonRoot = v
}
func (w *podSecurityContextWrapper) SupplementalGroups() []int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.SupplementalGroups
}
func (w *podSecurityContextWrapper) SetSupplementalGroups(v []int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && len(v) == 0 {
		return
	}
	w.ensurePodSC()
	if len(v) == 0 && len(w.podSC.SupplementalGroups) == 0 {
		return
	}
	w.podSC.SupplementalGroups = v
}
func (w *podSecurityContextWrapper) FSGroup() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil {
		return nil
	}
	return w.podSC.FSGroup
}
func (w *podSecurityContextWrapper) SetFSGroup(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.podSC == nil && v == nil {
		return
	}
	w.ensurePodSC()
	w.podSC.FSGroup = v
}

type ContainerSecurityContextAccessor interface {
	Capabilities() *api.Capabilities
	Privileged() *bool
	ProcMount() api.ProcMountType
	SELinuxOptions() *api.SELinuxOptions
	RunAsUser() *int64
	RunAsGroup() *int64
	RunAsNonRoot() *bool
	ReadOnlyRootFilesystem() *bool
	AllowPrivilegeEscalation() *bool
}
type ContainerSecurityContextMutator interface {
	ContainerSecurityContextAccessor
	ContainerSecurityContext() *api.SecurityContext
	SetCapabilities(*api.Capabilities)
	SetPrivileged(*bool)
	SetSELinuxOptions(*api.SELinuxOptions)
	SetRunAsUser(*int64)
	SetRunAsGroup(*int64)
	SetRunAsNonRoot(*bool)
	SetReadOnlyRootFilesystem(*bool)
	SetAllowPrivilegeEscalation(*bool)
}

func NewContainerSecurityContextAccessor(containerSC *api.SecurityContext) ContainerSecurityContextAccessor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &containerSecurityContextWrapper{containerSC: containerSC}
}
func NewContainerSecurityContextMutator(containerSC *api.SecurityContext) ContainerSecurityContextMutator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &containerSecurityContextWrapper{containerSC: containerSC}
}

type containerSecurityContextWrapper struct{ containerSC *api.SecurityContext }

func (w *containerSecurityContextWrapper) ContainerSecurityContext() *api.SecurityContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC
}
func (w *containerSecurityContextWrapper) ensureContainerSC() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		w.containerSC = &api.SecurityContext{}
	}
}
func (w *containerSecurityContextWrapper) Capabilities() *api.Capabilities {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.Capabilities
}
func (w *containerSecurityContextWrapper) SetCapabilities(v *api.Capabilities) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.Capabilities = v
}
func (w *containerSecurityContextWrapper) Privileged() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.Privileged
}
func (w *containerSecurityContextWrapper) SetPrivileged(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.Privileged = v
}
func (w *containerSecurityContextWrapper) ProcMount() api.ProcMountType {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return api.DefaultProcMount
	}
	if w.containerSC.ProcMount == nil {
		return api.DefaultProcMount
	}
	return *w.containerSC.ProcMount
}
func (w *containerSecurityContextWrapper) SELinuxOptions() *api.SELinuxOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.SELinuxOptions
}
func (w *containerSecurityContextWrapper) SetSELinuxOptions(v *api.SELinuxOptions) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.SELinuxOptions = v
}
func (w *containerSecurityContextWrapper) RunAsUser() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.RunAsUser
}
func (w *containerSecurityContextWrapper) SetRunAsUser(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.RunAsUser = v
}
func (w *containerSecurityContextWrapper) RunAsGroup() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.RunAsGroup
}
func (w *containerSecurityContextWrapper) SetRunAsGroup(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.RunAsGroup = v
}
func (w *containerSecurityContextWrapper) RunAsNonRoot() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.RunAsNonRoot
}
func (w *containerSecurityContextWrapper) SetRunAsNonRoot(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.RunAsNonRoot = v
}
func (w *containerSecurityContextWrapper) ReadOnlyRootFilesystem() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.ReadOnlyRootFilesystem
}
func (w *containerSecurityContextWrapper) SetReadOnlyRootFilesystem(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.ReadOnlyRootFilesystem = v
}
func (w *containerSecurityContextWrapper) AllowPrivilegeEscalation() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil {
		return nil
	}
	return w.containerSC.AllowPrivilegeEscalation
}
func (w *containerSecurityContextWrapper) SetAllowPrivilegeEscalation(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if w.containerSC == nil && v == nil {
		return
	}
	w.ensureContainerSC()
	w.containerSC.AllowPrivilegeEscalation = v
}
func NewEffectiveContainerSecurityContextAccessor(podSC PodSecurityContextAccessor, containerSC ContainerSecurityContextMutator) ContainerSecurityContextAccessor {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &effectiveContainerSecurityContextWrapper{podSC: podSC, containerSC: containerSC}
}
func NewEffectiveContainerSecurityContextMutator(podSC PodSecurityContextAccessor, containerSC ContainerSecurityContextMutator) ContainerSecurityContextMutator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &effectiveContainerSecurityContextWrapper{podSC: podSC, containerSC: containerSC}
}

type effectiveContainerSecurityContextWrapper struct {
	podSC       PodSecurityContextAccessor
	containerSC ContainerSecurityContextMutator
}

func (w *effectiveContainerSecurityContextWrapper) ContainerSecurityContext() *api.SecurityContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.ContainerSecurityContext()
}
func (w *effectiveContainerSecurityContextWrapper) Capabilities() *api.Capabilities {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.Capabilities()
}
func (w *effectiveContainerSecurityContextWrapper) SetCapabilities(v *api.Capabilities) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.Capabilities(), v) {
		w.containerSC.SetCapabilities(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) Privileged() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.Privileged()
}
func (w *effectiveContainerSecurityContextWrapper) SetPrivileged(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.Privileged(), v) {
		w.containerSC.SetPrivileged(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) ProcMount() api.ProcMountType {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.ProcMount()
}
func (w *effectiveContainerSecurityContextWrapper) SELinuxOptions() *api.SELinuxOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if v := w.containerSC.SELinuxOptions(); v != nil {
		return v
	}
	return w.podSC.SELinuxOptions()
}
func (w *effectiveContainerSecurityContextWrapper) SetSELinuxOptions(v *api.SELinuxOptions) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.SELinuxOptions(), v) {
		w.containerSC.SetSELinuxOptions(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) RunAsUser() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if v := w.containerSC.RunAsUser(); v != nil {
		return v
	}
	return w.podSC.RunAsUser()
}
func (w *effectiveContainerSecurityContextWrapper) SetRunAsUser(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.RunAsUser(), v) {
		w.containerSC.SetRunAsUser(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) RunAsGroup() *int64 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if v := w.containerSC.RunAsGroup(); v != nil {
		return v
	}
	return w.podSC.RunAsGroup()
}
func (w *effectiveContainerSecurityContextWrapper) SetRunAsGroup(v *int64) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.RunAsGroup(), v) {
		w.containerSC.SetRunAsGroup(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) RunAsNonRoot() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if v := w.containerSC.RunAsNonRoot(); v != nil {
		return v
	}
	return w.podSC.RunAsNonRoot()
}
func (w *effectiveContainerSecurityContextWrapper) SetRunAsNonRoot(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.RunAsNonRoot(), v) {
		w.containerSC.SetRunAsNonRoot(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) ReadOnlyRootFilesystem() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.ReadOnlyRootFilesystem()
}
func (w *effectiveContainerSecurityContextWrapper) SetReadOnlyRootFilesystem(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.ReadOnlyRootFilesystem(), v) {
		w.containerSC.SetReadOnlyRootFilesystem(v)
	}
}
func (w *effectiveContainerSecurityContextWrapper) AllowPrivilegeEscalation() *bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return w.containerSC.AllowPrivilegeEscalation()
}
func (w *effectiveContainerSecurityContextWrapper) SetAllowPrivilegeEscalation(v *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !reflect.DeepEqual(w.AllowPrivilegeEscalation(), v) {
		w.containerSC.SetAllowPrivilegeEscalation(v)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
