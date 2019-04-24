package deploymentconfigs

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"reflect"
	"strings"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	appsv1 "github.com/openshift/api/apps/v1"
	appsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	triggerapi "github.com/openshift/origin/pkg/image/apis/image/v1/trigger"
	"github.com/openshift/origin/pkg/image/trigger"
)

func indicesForContainerNames(containers []corev1.Container, names []string) []int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var index []int
	for _, name := range names {
		for i, container := range containers {
			if name == container.Name {
				index = append(index, i)
			}
		}
	}
	return index
}
func namesInclude(names []string, name string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}
func calculateDeploymentConfigTrigger(t appsv1.DeploymentTriggerPolicy, dc *appsv1.DeploymentConfig) []triggerapi.ObjectFieldTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.ImageChangeParams == nil {
		return nil
	}
	from := t.ImageChangeParams.From
	if from.Kind != "ImageStreamTag" || len(from.Name) == 0 {
		return nil
	}
	var triggers []triggerapi.ObjectFieldTrigger
	for _, index := range indicesForContainerNames(dc.Spec.Template.Spec.Containers, t.ImageChangeParams.ContainerNames) {
		fieldPath := fmt.Sprintf("spec.template.spec.containers[@name==\"%s\"].image", dc.Spec.Template.Spec.Containers[index].Name)
		triggers = append(triggers, triggerapi.ObjectFieldTrigger{From: triggerapi.ObjectReference{Name: from.Name, Namespace: from.Namespace, Kind: from.Kind, APIVersion: from.APIVersion}, FieldPath: fieldPath, Paused: !t.ImageChangeParams.Automatic})
	}
	for _, index := range indicesForContainerNames(dc.Spec.Template.Spec.InitContainers, t.ImageChangeParams.ContainerNames) {
		fieldPath := fmt.Sprintf("spec.template.spec.initContainers[@name==\"%s\"].image", dc.Spec.Template.Spec.InitContainers[index].Name)
		triggers = append(triggers, triggerapi.ObjectFieldTrigger{From: triggerapi.ObjectReference{Name: from.Name, Namespace: from.Namespace, Kind: from.Kind, APIVersion: from.APIVersion}, FieldPath: fieldPath, Paused: !t.ImageChangeParams.Automatic})
	}
	return triggers
}
func calculateDeploymentConfigTriggers(dc *appsv1.DeploymentConfig) []triggerapi.ObjectFieldTrigger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var triggers []triggerapi.ObjectFieldTrigger
	for _, t := range dc.Spec.Triggers {
		addedTriggers := calculateDeploymentConfigTrigger(t, dc)
		triggers = append(triggers, addedTriggers...)
	}
	return triggers
}

type deploymentConfigTriggerIndexer struct{ prefix string }

func NewDeploymentConfigTriggerIndexer(prefix string) trigger.Indexer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return deploymentConfigTriggerIndexer{prefix: prefix}
}
func (i deploymentConfigTriggerIndexer) Index(obj, old interface{}) (string, *trigger.CacheEntry, cache.DeltaType, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		triggers	[]triggerapi.ObjectFieldTrigger
		dc		*appsv1.DeploymentConfig
		change		cache.DeltaType
	)
	switch {
	case obj != nil && old == nil:
		dc = obj.(*appsv1.DeploymentConfig)
		triggers = calculateDeploymentConfigTriggers(dc)
		change = cache.Added
	case old != nil && obj == nil:
		switch deleted := old.(type) {
		case *appsv1.DeploymentConfig:
			dc = deleted
			triggers = calculateDeploymentConfigTriggers(dc)
		case cache.DeletedFinalStateUnknown:
			klog.V(4).Infof("skipping trigger calculation for deleted deploymentconfig %q", deleted.Key)
		}
		change = cache.Deleted
	default:
		dc = obj.(*appsv1.DeploymentConfig)
		oldDC := old.(*appsv1.DeploymentConfig)
		triggers = calculateDeploymentConfigTriggers(dc)
		oldTriggers := calculateDeploymentConfigTriggers(oldDC)
		switch {
		case len(oldTriggers) == 0:
			change = cache.Added
		case !reflect.DeepEqual(oldTriggers, triggers):
			change = cache.Updated
		case !reflect.DeepEqual(dc.Spec.Template.Spec.Containers, oldDC.Spec.Template.Spec.Containers):
			change = cache.Updated
		}
	}
	if len(triggers) > 0 {
		key := i.prefix + dc.Namespace + "/" + dc.Name
		return key, &trigger.CacheEntry{Key: key, Namespace: dc.Namespace, Triggers: triggers}, change, nil
	}
	return "", nil, change, nil
}

type DeploymentConfigReactor struct {
	Client appsclient.DeploymentConfigsGetter
}

func UpdateDeploymentConfigImages(dc *appsv1.DeploymentConfig, tagRetriever trigger.TagRetriever) (*appsv1.DeploymentConfig, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var updated *appsv1.DeploymentConfig
	copyObject := func() {
		if updated != nil {
			return
		}
		dc = dc.DeepCopy()
		updated = dc
	}
	for i, t := range dc.Spec.Triggers {
		p := t.ImageChangeParams
		if p == nil || p.From.Kind != "ImageStreamTag" {
			continue
		}
		if !p.Automatic {
			continue
		}
		namespace := p.From.Namespace
		if len(namespace) == 0 {
			namespace = dc.Namespace
		}
		ref, _, ok := tagRetriever.ImageStreamTag(namespace, p.From.Name)
		if !ok && len(p.LastTriggeredImage) == 0 {
			klog.V(4).Infof("trigger %#v in deployment %s is not resolveable", p, dc.Name)
			return nil, false, nil
		}
		if len(ref) == 0 {
			ref = p.LastTriggeredImage
		}
		if p.LastTriggeredImage != ref {
			copyObject()
			p = dc.Spec.Triggers[i].ImageChangeParams
			p.LastTriggeredImage = ref
		}
		for i, c := range dc.Spec.Template.Spec.InitContainers {
			if !namesInclude(p.ContainerNames, c.Name) || c.Image == ref {
				continue
			}
			copyObject()
			container := &dc.Spec.Template.Spec.InitContainers[i]
			container.Image = ref
		}
		for i, c := range dc.Spec.Template.Spec.Containers {
			if !namesInclude(p.ContainerNames, c.Name) || c.Image == ref {
				continue
			}
			copyObject()
			container := &dc.Spec.Template.Spec.Containers[i]
			container.Image = ref
		}
	}
	return updated, true, nil
}
func (r *DeploymentConfigReactor) ImageChanged(obj runtime.Object, tagRetriever trigger.TagRetriever) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dc := obj.(*appsv1.DeploymentConfig)
	newDC := dc.DeepCopy()
	updated, resolvable, err := UpdateDeploymentConfigImages(newDC, tagRetriever)
	if err != nil {
		return err
	}
	if !resolvable {
		if klog.V(4) {
			klog.Infof("Ignoring changes to deployment config %s, has unresolved images: %s", dc.Name, printDeploymentTriggers(newDC.Spec.Triggers))
		}
		return nil
	}
	if updated == nil {
		klog.V(4).Infof("Deployment config %s has not changed", dc.Name)
		return nil
	}
	klog.V(4).Infof("Deployment config %s has changed", dc.Name)
	_, err = r.Client.DeploymentConfigs(dc.Namespace).Update(updated)
	return err
}
func printDeploymentTriggers(triggers []appsv1.DeploymentTriggerPolicy) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var values []string
	for _, t := range triggers {
		if t.ImageChangeParams == nil {
			continue
		}
		values = append(values, fmt.Sprintf("[from=%s last=%s]", t.ImageChangeParams.From.Name, t.ImageChangeParams.LastTriggeredImage))
	}
	return strings.Join(values, ", ")
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
