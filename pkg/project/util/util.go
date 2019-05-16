package util

import (
	"fmt"
	goformat "fmt"
	oapi "github.com/openshift/origin/pkg/api"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	apistorage "k8s.io/apiserver/pkg/storage"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ConvertNamespaceFromExternal(namespace *corev1.Namespace) *projectapi.Project {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	internalFinalizers := []kapi.FinalizerName{}
	for _, externalFinalizer := range namespace.Spec.Finalizers {
		internalFinalizers = append(internalFinalizers, kapi.FinalizerName(externalFinalizer))
	}
	return &projectapi.Project{ObjectMeta: namespace.ObjectMeta, Spec: projectapi.ProjectSpec{Finalizers: internalFinalizers}, Status: projectapi.ProjectStatus{Phase: kapi.NamespacePhase(namespace.Status.Phase)}}
}
func ConvertProjectToExternal(project *projectapi.Project) *corev1.Namespace {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	externalFinalizers := []corev1.FinalizerName{}
	for _, internalFinalizer := range project.Spec.Finalizers {
		externalFinalizers = append(externalFinalizers, corev1.FinalizerName(internalFinalizer))
	}
	namespace := &corev1.Namespace{ObjectMeta: project.ObjectMeta, Spec: corev1.NamespaceSpec{Finalizers: externalFinalizers}, Status: corev1.NamespaceStatus{Phase: corev1.NamespacePhase(project.Status.Phase)}}
	if namespace.Annotations == nil {
		namespace.Annotations = map[string]string{}
	}
	namespace.Annotations[oapi.OpenShiftDisplayName] = project.Annotations[oapi.OpenShiftDisplayName]
	return namespace
}
func ConvertNamespaceList(namespaceList *corev1.NamespaceList) *projectapi.ProjectList {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projects := &projectapi.ProjectList{}
	for _, n := range namespaceList.Items {
		projects.Items = append(projects.Items, *ConvertNamespaceFromExternal(&n))
	}
	return projects
}
func getAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	projectObj, ok := obj.(*projectapi.Project)
	if !ok {
		return nil, nil, false, fmt.Errorf("not a project")
	}
	return labels.Set(projectObj.Labels), projectToSelectableFields(projectObj), projectObj.Initializers != nil, nil
}
func MatchProject(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}
func projectToSelectableFields(projectObj *projectapi.Project) fields.Set {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&projectObj.ObjectMeta, false)
	specificFieldsSet := fields.Set{"status.phase": string(projectObj.Status.Phase)}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
