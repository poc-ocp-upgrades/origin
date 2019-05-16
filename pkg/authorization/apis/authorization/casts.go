package authorization

import (
	goformat "fmt"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ToRoleList(in *ClusterRoleList) *RoleList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := &RoleList{}
	for _, curr := range in.Items {
		ret.Items = append(ret.Items, *ToRole(&curr))
	}
	return ret
}
func ToRole(in *ClusterRole) *Role {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	ret := &Role{}
	ret.ObjectMeta = in.ObjectMeta
	ret.Rules = in.Rules
	return ret
}
func ToRoleBindingList(in *ClusterRoleBindingList) *RoleBindingList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := &RoleBindingList{}
	for _, curr := range in.Items {
		ret.Items = append(ret.Items, *ToRoleBinding(&curr))
	}
	return ret
}
func ToRoleBinding(in *ClusterRoleBinding) *RoleBinding {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if in == nil {
		return nil
	}
	ret := &RoleBinding{}
	ret.ObjectMeta = in.ObjectMeta
	ret.Subjects = in.Subjects
	ret.RoleRef = ToRoleRef(in.RoleRef)
	return ret
}
func ToRoleRef(in kapi.ObjectReference) kapi.ObjectReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := kapi.ObjectReference{}
	ret.Name = in.Name
	return ret
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
