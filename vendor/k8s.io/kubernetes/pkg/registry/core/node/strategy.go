package node

import (
 "context"
 "fmt"
 "net"
 "net/http"
 "net/url"
 "k8s.io/apimachinery/pkg/api/errors"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/fields"
 "k8s.io/apimachinery/pkg/labels"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apimachinery/pkg/types"
 utilnet "k8s.io/apimachinery/pkg/util/net"
 "k8s.io/apimachinery/pkg/util/validation/field"
 "k8s.io/apiserver/pkg/registry/generic"
 pkgstorage "k8s.io/apiserver/pkg/storage"
 "k8s.io/apiserver/pkg/storage/names"
 utilfeature "k8s.io/apiserver/pkg/util/feature"
 "k8s.io/kubernetes/pkg/api/legacyscheme"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/apis/core/validation"
 "k8s.io/kubernetes/pkg/features"
 "k8s.io/kubernetes/pkg/kubelet/client"
 proxyutil "k8s.io/kubernetes/pkg/proxy/util"
)

type nodeStrategy struct {
 runtime.ObjectTyper
 names.NameGenerator
}

var Strategy = nodeStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (nodeStrategy) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (nodeStrategy) AllowCreateOnUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (nodeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := obj.(*api.Node)
 if !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  node.Spec.ConfigSource = nil
 }
}
func (nodeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNode := obj.(*api.Node)
 oldNode := old.(*api.Node)
 newNode.Status = oldNode.Status
 if !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  newNode.Spec.ConfigSource = nil
  oldNode.Spec.ConfigSource = nil
 }
}
func (nodeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := obj.(*api.Node)
 return validation.ValidateNode(node)
}
func (nodeStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}
func (nodeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 errorList := validation.ValidateNode(obj.(*api.Node))
 return append(errorList, validation.ValidateNodeUpdate(obj.(*api.Node), old.(*api.Node))...)
}
func (nodeStrategy) AllowUnconditionalUpdate() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return true
}
func (ns nodeStrategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 n, ok := obj.(*api.Node)
 if !ok {
  return fmt.Errorf("unexpected object: %v", obj)
 }
 ns.PrepareForCreate(ctx, obj)
 if exact {
  return nil
 }
 n.Status = api.NodeStatus{}
 return nil
}

type nodeStatusStrategy struct{ nodeStrategy }

var StatusStrategy = nodeStatusStrategy{Strategy}

func (nodeStatusStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := obj.(*api.Node)
 if !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  node.Status.Config = nil
 }
}
func (nodeStatusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 newNode := obj.(*api.Node)
 oldNode := old.(*api.Node)
 newNode.Spec = oldNode.Spec
 if !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
  newNode.Status.Config = nil
  oldNode.Status.Config = nil
 }
}
func (nodeStatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return validation.ValidateNodeUpdate(obj.(*api.Node), old.(*api.Node))
}
func (nodeStatusStrategy) Canonicalize(obj runtime.Object) {
 _logClusterCodePath()
 defer _logClusterCodePath()
}

type ResourceGetter interface {
 Get(context.Context, string, *metav1.GetOptions) (runtime.Object, error)
}

func NodeToSelectableFields(node *api.Node) fields.Set {
 _logClusterCodePath()
 defer _logClusterCodePath()
 objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&node.ObjectMeta, false)
 specificFieldsSet := fields.Set{"spec.unschedulable": fmt.Sprint(node.Spec.Unschedulable)}
 return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 nodeObj, ok := obj.(*api.Node)
 if !ok {
  return nil, nil, false, fmt.Errorf("not a node")
 }
 return labels.Set(nodeObj.ObjectMeta.Labels), NodeToSelectableFields(nodeObj), nodeObj.Initializers != nil, nil
}
func MatchNode(label labels.Selector, field fields.Selector) pkgstorage.SelectionPredicate {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return pkgstorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs, IndexFields: []string{"metadata.name"}}
}
func NodeNameTriggerFunc(obj runtime.Object) []pkgstorage.MatchValue {
 _logClusterCodePath()
 defer _logClusterCodePath()
 node := obj.(*api.Node)
 result := pkgstorage.MatchValue{IndexName: "metadata.name", Value: node.ObjectMeta.Name}
 return []pkgstorage.MatchValue{result}
}
func ResourceLocation(getter ResourceGetter, connection client.ConnectionInfoGetter, proxyTransport http.RoundTripper, ctx context.Context, id string) (*url.URL, http.RoundTripper, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 schemeReq, name, portReq, valid := utilnet.SplitSchemeNamePort(id)
 if !valid {
  return nil, nil, errors.NewBadRequest(fmt.Sprintf("invalid node request %q", id))
 }
 info, err := connection.GetConnectionInfo(ctx, types.NodeName(name))
 if err != nil {
  return nil, nil, err
 }
 if portReq == "" || portReq == info.Port {
  return &url.URL{Scheme: info.Scheme, Host: net.JoinHostPort(info.Hostname, info.Port)}, info.Transport, nil
 }
 if err := proxyutil.IsProxyableHostname(ctx, &net.Resolver{}, info.Hostname); err != nil {
  return nil, nil, errors.NewBadRequest(err.Error())
 }
 return &url.URL{Scheme: schemeReq, Host: net.JoinHostPort(info.Hostname, portReq)}, proxyTransport, nil
}
