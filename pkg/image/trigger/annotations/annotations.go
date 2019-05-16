package annotations

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	ometa "github.com/openshift/origin/pkg/api/imagereferencemutators"
	triggerapi "github.com/openshift/origin/pkg/image/apis/image/v1/trigger"
	"github.com/openshift/origin/pkg/image/trigger"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	goos "os"
	"reflect"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

func calculateAnnotationTriggers(m metav1.Object, prefix string) (string, string, []triggerapi.ObjectFieldTrigger, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var key, namespace string
	if namespace = m.GetNamespace(); len(namespace) > 0 {
		key = prefix + namespace + "/" + m.GetName()
	} else {
		key = prefix + m.GetName()
	}
	t, ok := m.GetAnnotations()[triggerapi.TriggerAnnotationKey]
	if !ok {
		return key, namespace, nil, nil
	}
	triggers := []triggerapi.ObjectFieldTrigger{}
	if err := json.Unmarshal([]byte(t), &triggers); err != nil {
		return key, namespace, nil, err
	}
	if hasDuplicateTriggers(triggers) {
		return key, namespace, nil, fmt.Errorf("duplicate triggers are not allowed")
	}
	return key, namespace, triggers, nil
}
func hasDuplicateTriggers(triggers []triggerapi.ObjectFieldTrigger) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for i := range triggers {
		for j := i + 1; j < len(triggers); j++ {
			if triggers[i].FieldPath == triggers[j].FieldPath {
				return true
			}
		}
	}
	return false
}
func parseContainerReference(path string) (init bool, selector string, remainder string, ok bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case strings.HasPrefix(path, "containers["):
		remainder = strings.TrimPrefix(path, "containers[")
	case strings.HasPrefix(path, "initContainers["):
		init = true
		remainder = strings.TrimPrefix(path, "initContainers[")
	default:
		return false, "", "", false
	}
	end := strings.Index(remainder, "]")
	if end == -1 {
		return false, "", "", false
	}
	selector = remainder[:end]
	remainder = remainder[end+1:]
	if len(remainder) > 0 && remainder[0] == '.' {
		remainder = remainder[1:]
	}
	return init, selector, remainder, true
}
func findContainerBySelector(spec ometa.PodSpecReferenceMutator, init bool, selector string) (ometa.ContainerMutator, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if i, err := strconv.Atoi(selector); err == nil {
		return spec.GetContainerByIndex(init, i)
	}
	if name := strings.TrimSuffix(strings.TrimPrefix(selector, "?(@.name==\""), "\")"); name != selector {
		return spec.GetContainerByName(name)
	}
	return nil, false
}
func ContainerForObjectFieldPath(obj runtime.Object, fieldPath string) (ometa.ContainerMutator, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	spec, err := ometa.GetPodSpecReferenceMutator(obj)
	if err != nil {
		return nil, fieldPath, err
	}
	specPath := spec.Path().String()
	containerPath := strings.TrimPrefix(fieldPath, specPath)
	if containerPath == fieldPath {
		return nil, fieldPath, fmt.Errorf("1 field path is not valid: %s", fieldPath)
	}
	containerPath = strings.TrimPrefix(containerPath, ".")
	init, selector, remainder, ok := parseContainerReference(containerPath)
	if !ok {
		return nil, fieldPath, fmt.Errorf("2 field path is not valid: %s", fieldPath)
	}
	container, ok := findContainerBySelector(spec, init, selector)
	if !ok {
		return nil, fieldPath, fmt.Errorf("no such container: %s", selector)
	}
	return container, remainder, nil
}
func UpdateObjectFromImages(obj runtime.Object, tagRetriever trigger.TagRetriever) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var updated runtime.Object
	m, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}
	spec, err := ometa.GetPodSpecReferenceMutator(obj)
	if err != nil {
		return nil, err
	}
	path := spec.Path()
	basePath := path.String() + "."
	_, _, triggers, err := calculateAnnotationTriggers(m, "/")
	if err != nil {
		return nil, err
	}
	klog.V(5).Infof("%T/%s has triggers: %#v", obj, m.GetName(), triggers)
	for _, trigger := range triggers {
		if trigger.Paused {
			continue
		}
		fieldPath := trigger.FieldPath
		if !strings.HasPrefix(trigger.FieldPath, basePath) {
			klog.V(5).Infof("%T/%s trigger %s did not match base path %s", obj, m.GetName(), trigger.FieldPath, basePath)
			continue
		}
		fieldPath = strings.TrimPrefix(fieldPath, basePath)
		namespace := trigger.From.Namespace
		if len(namespace) == 0 {
			namespace = m.GetNamespace()
		}
		ref, _, ok := tagRetriever.ImageStreamTag(namespace, trigger.From.Name)
		if !ok {
			klog.V(5).Infof("%T/%s detected no pending image on %s from %#v", obj, m.GetName(), trigger.FieldPath, trigger.From)
			continue
		}
		init, selector, remainder, ok := parseContainerReference(fieldPath)
		if !ok || remainder != "image" {
			return nil, fmt.Errorf("field path is not valid: %s", trigger.FieldPath)
		}
		container, ok := findContainerBySelector(spec, init, selector)
		if !ok {
			return nil, fmt.Errorf("no such container: %s", trigger.FieldPath)
		}
		if container.GetImage() != ref {
			if updated == nil {
				updated = obj.DeepCopyObject()
				spec, _ = ometa.GetPodSpecReferenceMutator(updated)
				container, _ = findContainerBySelector(spec, init, selector)
			}
			klog.V(5).Infof("%T/%s detected change on %s = %s", obj, m.GetName(), trigger.FieldPath, ref)
			container.SetImage(ref)
		}
	}
	return updated, nil
}
func ContainerImageChanged(oldObj, newObj runtime.Object, newTriggers []triggerapi.ObjectFieldTrigger) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, trigger := range newTriggers {
		if trigger.Paused {
			continue
		}
		newContainer, _, err := ContainerForObjectFieldPath(newObj, trigger.FieldPath)
		if err != nil {
			klog.V(5).Infof("%v", err)
			continue
		}
		oldContainer, _, err := ContainerForObjectFieldPath(oldObj, trigger.FieldPath)
		if err != nil {
			continue
		}
		if newContainer.GetImage() != oldContainer.GetImage() {
			return true
		}
	}
	return false
}

type annotationTriggerIndexer struct{ prefix string }

func NewAnnotationTriggerIndexer(prefix string) trigger.Indexer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return annotationTriggerIndexer{prefix: prefix}
}
func (i annotationTriggerIndexer) Index(obj, old interface{}) (string, *trigger.CacheEntry, cache.DeltaType, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		triggers  []triggerapi.ObjectFieldTrigger
		key       string
		namespace string
		change    cache.DeltaType
	)
	switch {
	case obj != nil && old == nil:
		m, err := meta.Accessor(obj)
		if err != nil {
			return "", nil, change, err
		}
		key, namespace, triggers, err = calculateAnnotationTriggers(m, i.prefix)
		if err != nil {
			return "", nil, change, err
		}
		change = cache.Added
	case old != nil && obj == nil:
		m, err := meta.Accessor(old)
		if err != nil {
			return "", nil, change, err
		}
		key, namespace, triggers, err = calculateAnnotationTriggers(m, i.prefix)
		if err != nil {
			return "", nil, change, err
		}
		change = cache.Deleted
	default:
		m, err := meta.Accessor(obj)
		if err != nil {
			return "", nil, change, err
		}
		key, namespace, triggers, err = calculateAnnotationTriggers(m, i.prefix)
		if err != nil {
			return "", nil, change, err
		}
		oldM, err := meta.Accessor(old)
		if err != nil {
			return "", nil, change, err
		}
		_, _, oldTriggers, err := calculateAnnotationTriggers(oldM, i.prefix)
		if err != nil {
			return "", nil, change, err
		}
		switch {
		case len(oldTriggers) == 0:
			change = cache.Added
		case !reflect.DeepEqual(oldTriggers, triggers):
			change = cache.Updated
		case ContainerImageChanged(old.(runtime.Object), obj.(runtime.Object), triggers):
			change = cache.Updated
		}
	}
	if len(triggers) > 0 {
		return key, &trigger.CacheEntry{Key: key, Namespace: namespace, Triggers: triggers}, change, nil
	}
	return "", nil, change, nil
}

type AnnotationUpdater interface {
	Update(obj runtime.Object) error
}
type AnnotationReactor struct{ Updater AnnotationUpdater }

func (r *AnnotationReactor) ImageChanged(obj runtime.Object, tagRetriever trigger.TagRetriever) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	changed, err := UpdateObjectFromImages(obj, tagRetriever)
	if err != nil {
		return err
	}
	if changed != nil {
		return r.Updater.Update(changed)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
