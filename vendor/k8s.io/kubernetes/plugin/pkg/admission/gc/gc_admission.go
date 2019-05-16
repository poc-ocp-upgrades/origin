package gc

import (
	"fmt"
	goformat "fmt"
	"io"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "OwnerReferencesPermissionEnforcement"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		whiteList := []whiteListItem{{groupResource: schema.GroupResource{Resource: "pods"}, subresource: "status"}}
		return &gcPermissionsEnforcement{Handler: admission.NewHandler(admission.Create, admission.Update), whiteList: whiteList}, nil
	})
}

type gcPermissionsEnforcement struct {
	*admission.Handler
	authorizer authorizer.Authorizer
	restMapper meta.RESTMapper
	whiteList  []whiteListItem
}

var _ admission.ValidationInterface = &gcPermissionsEnforcement{}

type whiteListItem struct {
	groupResource schema.GroupResource
	subresource   string
}

func (a *gcPermissionsEnforcement) isWhiteListed(groupResource schema.GroupResource, subresource string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, item := range a.whiteList {
		if item.groupResource == groupResource && item.subresource == subresource {
			return true
		}
	}
	return false
}
func (a *gcPermissionsEnforcement) Validate(attributes admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.isWhiteListed(attributes.GetResource().GroupResource(), attributes.GetSubresource()) {
		return nil
	}
	if !isChangingOwnerReference(attributes.GetObject(), attributes.GetOldObject()) {
		return nil
	}
	if attributes.GetOperation() != admission.Create {
		deleteAttributes := authorizer.AttributesRecord{User: attributes.GetUserInfo(), Verb: "delete", Namespace: attributes.GetNamespace(), APIGroup: attributes.GetResource().Group, APIVersion: attributes.GetResource().Version, Resource: attributes.GetResource().Resource, Subresource: attributes.GetSubresource(), Name: attributes.GetName(), ResourceRequest: true, Path: ""}
		decision, reason, err := a.authorizer.Authorize(deleteAttributes)
		if decision != authorizer.DecisionAllow {
			return admission.NewForbidden(attributes, fmt.Errorf("cannot set an ownerRef on a resource you can't delete: %v, %v", reason, err))
		}
	}
	newBlockingRefs := newBlockingOwnerDeletionRefs(attributes.GetObject(), attributes.GetOldObject())
	for _, ref := range newBlockingRefs {
		records, err := a.ownerRefToDeleteAttributeRecords(ref, attributes)
		if err != nil {
			return admission.NewForbidden(attributes, fmt.Errorf("cannot set blockOwnerDeletion in this case because cannot find RESTMapping for APIVersion %s Kind %s: %v", ref.APIVersion, ref.Kind, err))
		}
		for _, record := range records {
			decision, reason, err := a.authorizer.Authorize(record)
			if decision != authorizer.DecisionAllow {
				return admission.NewForbidden(attributes, fmt.Errorf("cannot set blockOwnerDeletion if an ownerReference refers to a resource you can't set finalizers on: %v, %v", reason, err))
			}
		}
	}
	return nil
}
func isChangingOwnerReference(newObj, oldObj runtime.Object) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newMeta, err := meta.Accessor(newObj)
	if err != nil {
		return false
	}
	if oldObj == nil {
		return len(newMeta.GetOwnerReferences()) > 0
	}
	oldMeta, err := meta.Accessor(oldObj)
	if err != nil {
		return false
	}
	oldOwners := oldMeta.GetOwnerReferences()
	newOwners := newMeta.GetOwnerReferences()
	if len(oldOwners) != len(newOwners) {
		return true
	}
	for i := range oldOwners {
		if !apiequality.Semantic.DeepEqual(oldOwners[i], newOwners[i]) {
			return true
		}
	}
	return false
}
func (a *gcPermissionsEnforcement) ownerRefToDeleteAttributeRecords(ref metav1.OwnerReference, attributes admission.Attributes) ([]authorizer.AttributesRecord, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var ret []authorizer.AttributesRecord
	groupVersion, err := schema.ParseGroupVersion(ref.APIVersion)
	if err != nil {
		return ret, err
	}
	mappings, err := a.restMapper.RESTMappings(schema.GroupKind{Group: groupVersion.Group, Kind: ref.Kind}, groupVersion.Version)
	if err != nil {
		return ret, err
	}
	for _, mapping := range mappings {
		ar := authorizer.AttributesRecord{User: attributes.GetUserInfo(), Verb: "update", APIGroup: mapping.Resource.Group, APIVersion: mapping.Resource.Version, Resource: mapping.Resource.Resource, Subresource: "finalizers", Name: ref.Name, ResourceRequest: true, Path: ""}
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			ar.Namespace = attributes.GetNamespace()
		}
		ret = append(ret, ar)
	}
	return ret, nil
}
func blockingOwnerRefs(refs []metav1.OwnerReference) []metav1.OwnerReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var ret []metav1.OwnerReference
	for _, ref := range refs {
		if ref.BlockOwnerDeletion != nil && *ref.BlockOwnerDeletion == true {
			ret = append(ret, ref)
		}
	}
	return ret
}
func indexByUID(refs []metav1.OwnerReference) map[types.UID]metav1.OwnerReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ret := make(map[types.UID]metav1.OwnerReference)
	for _, ref := range refs {
		ret[ref.UID] = ref
	}
	return ret
}
func newBlockingOwnerDeletionRefs(newObj, oldObj runtime.Object) []metav1.OwnerReference {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	newMeta, err := meta.Accessor(newObj)
	if err != nil {
		return nil
	}
	newRefs := newMeta.GetOwnerReferences()
	blockingNewRefs := blockingOwnerRefs(newRefs)
	if len(blockingNewRefs) == 0 {
		return nil
	}
	if oldObj == nil {
		return blockingNewRefs
	}
	oldMeta, err := meta.Accessor(oldObj)
	if err != nil {
		return blockingNewRefs
	}
	var ret []metav1.OwnerReference
	indexedOldRefs := indexByUID(oldMeta.GetOwnerReferences())
	for _, ref := range blockingNewRefs {
		oldRef, ok := indexedOldRefs[ref.UID]
		if !ok {
			ret = append(ret, ref)
			continue
		}
		wasNotBlocking := oldRef.BlockOwnerDeletion == nil || *oldRef.BlockOwnerDeletion == false
		if wasNotBlocking {
			ret = append(ret, ref)
		}
	}
	return ret
}
func (a *gcPermissionsEnforcement) SetAuthorizer(authorizer authorizer.Authorizer) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.authorizer = authorizer
}
func (a *gcPermissionsEnforcement) SetRESTMapper(restMapper meta.RESTMapper) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	a.restMapper = restMapper
}
func (a *gcPermissionsEnforcement) ValidateInitialization() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.authorizer == nil {
		return fmt.Errorf("missing authorizer")
	}
	if a.restMapper == nil {
		return fmt.Errorf("missing restMapper")
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
