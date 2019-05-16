package history

import (
	"bytes"
	"fmt"
	goformat "fmt"
	"hash/fnv"
	apps "k8s.io/api/apps/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	clientset "k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"
	hashutil "k8s.io/kubernetes/pkg/util/hash"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strconv"
	gotime "time"
)

const ControllerRevisionHashLabel = "controller.kubernetes.io/hash"

func ControllerRevisionName(prefix string, hash string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(prefix) > 223 {
		prefix = prefix[:223]
	}
	return fmt.Sprintf("%s-%s", prefix, hash)
}
func NewControllerRevision(parent metav1.Object, parentKind schema.GroupVersionKind, templateLabels map[string]string, data runtime.RawExtension, revision int64, collisionCount *int32) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	labelMap := make(map[string]string)
	for k, v := range templateLabels {
		labelMap[k] = v
	}
	blockOwnerDeletion := true
	isController := true
	cr := &apps.ControllerRevision{ObjectMeta: metav1.ObjectMeta{Labels: labelMap, OwnerReferences: []metav1.OwnerReference{{APIVersion: parentKind.GroupVersion().String(), Kind: parentKind.Kind, Name: parent.GetName(), UID: parent.GetUID(), BlockOwnerDeletion: &blockOwnerDeletion, Controller: &isController}}}, Data: data, Revision: revision}
	hash := HashControllerRevision(cr, collisionCount)
	cr.Name = ControllerRevisionName(parent.GetName(), hash)
	cr.Labels[ControllerRevisionHashLabel] = hash
	return cr, nil
}
func HashControllerRevision(revision *apps.ControllerRevision, probe *int32) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hf := fnv.New32()
	if len(revision.Data.Raw) > 0 {
		hf.Write(revision.Data.Raw)
	}
	if revision.Data.Object != nil {
		hashutil.DeepHashObject(hf, revision.Data.Object)
	}
	if probe != nil {
		hf.Write([]byte(strconv.FormatInt(int64(*probe), 10)))
	}
	return rand.SafeEncodeString(fmt.Sprint(hf.Sum32()))
}
func SortControllerRevisions(revisions []*apps.ControllerRevision) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	sort.Sort(byRevision(revisions))
}
func EqualRevision(lhs *apps.ControllerRevision, rhs *apps.ControllerRevision) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var lhsHash, rhsHash *uint32
	if lhs == nil || rhs == nil {
		return lhs == rhs
	}
	if hs, found := lhs.Labels[ControllerRevisionHashLabel]; found {
		hash, err := strconv.ParseInt(hs, 10, 32)
		if err == nil {
			lhsHash = new(uint32)
			*lhsHash = uint32(hash)
		}
	}
	if hs, found := rhs.Labels[ControllerRevisionHashLabel]; found {
		hash, err := strconv.ParseInt(hs, 10, 32)
		if err == nil {
			rhsHash = new(uint32)
			*rhsHash = uint32(hash)
		}
	}
	if lhsHash != nil && rhsHash != nil && *lhsHash != *rhsHash {
		return false
	}
	return bytes.Equal(lhs.Data.Raw, rhs.Data.Raw) && apiequality.Semantic.DeepEqual(lhs.Data.Object, rhs.Data.Object)
}
func FindEqualRevisions(revisions []*apps.ControllerRevision, needle *apps.ControllerRevision) []*apps.ControllerRevision {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var eq []*apps.ControllerRevision
	for i := range revisions {
		if EqualRevision(revisions[i], needle) {
			eq = append(eq, revisions[i])
		}
	}
	return eq
}

type byRevision []*apps.ControllerRevision

func (br byRevision) Len() int {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return len(br)
}
func (br byRevision) Less(i, j int) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return br[i].Revision < br[j].Revision
}
func (br byRevision) Swap(i, j int) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	br[i], br[j] = br[j], br[i]
}

type Interface interface {
	ListControllerRevisions(parent metav1.Object, selector labels.Selector) ([]*apps.ControllerRevision, error)
	CreateControllerRevision(parent metav1.Object, revision *apps.ControllerRevision, collisionCount *int32) (*apps.ControllerRevision, error)
	DeleteControllerRevision(revision *apps.ControllerRevision) error
	UpdateControllerRevision(revision *apps.ControllerRevision, newRevision int64) (*apps.ControllerRevision, error)
	AdoptControllerRevision(parent metav1.Object, parentKind schema.GroupVersionKind, revision *apps.ControllerRevision) (*apps.ControllerRevision, error)
	ReleaseControllerRevision(parent metav1.Object, revision *apps.ControllerRevision) (*apps.ControllerRevision, error)
}

func NewHistory(client clientset.Interface, lister appslisters.ControllerRevisionLister) Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &realHistory{client, lister}
}
func NewFakeHistory(informer appsinformers.ControllerRevisionInformer) Interface {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &fakeHistory{informer.Informer().GetIndexer(), informer.Lister()}
}

type realHistory struct {
	client clientset.Interface
	lister appslisters.ControllerRevisionLister
}

func (rh *realHistory) ListControllerRevisions(parent metav1.Object, selector labels.Selector) ([]*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	history, err := rh.lister.ControllerRevisions(parent.GetNamespace()).List(selector)
	if err != nil {
		return nil, err
	}
	var owned []*apps.ControllerRevision
	for i := range history {
		ref := metav1.GetControllerOf(history[i])
		if ref == nil || ref.UID == parent.GetUID() {
			owned = append(owned, history[i])
		}
	}
	return owned, err
}
func (rh *realHistory) CreateControllerRevision(parent metav1.Object, revision *apps.ControllerRevision, collisionCount *int32) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if collisionCount == nil {
		return nil, fmt.Errorf("collisionCount should not be nil")
	}
	clone := revision.DeepCopy()
	for {
		hash := HashControllerRevision(revision, collisionCount)
		clone.Name = ControllerRevisionName(parent.GetName(), hash)
		ns := parent.GetNamespace()
		created, err := rh.client.AppsV1().ControllerRevisions(ns).Create(clone)
		if errors.IsAlreadyExists(err) {
			exists, err := rh.client.AppsV1().ControllerRevisions(ns).Get(clone.Name, metav1.GetOptions{})
			if err != nil {
				return nil, err
			}
			if bytes.Equal(exists.Data.Raw, clone.Data.Raw) {
				return exists, nil
			}
			*collisionCount++
			continue
		}
		return created, err
	}
}
func (rh *realHistory) UpdateControllerRevision(revision *apps.ControllerRevision, newRevision int64) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clone := revision.DeepCopy()
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		if clone.Revision == newRevision {
			return nil
		}
		clone.Revision = newRevision
		updated, updateErr := rh.client.AppsV1().ControllerRevisions(clone.Namespace).Update(clone)
		if updateErr == nil {
			return nil
		}
		if updated != nil {
			clone = updated
		}
		if updated, err := rh.lister.ControllerRevisions(clone.Namespace).Get(clone.Name); err == nil {
			clone = updated.DeepCopy()
		}
		return updateErr
	})
	return clone, err
}
func (rh *realHistory) DeleteControllerRevision(revision *apps.ControllerRevision) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rh.client.AppsV1().ControllerRevisions(revision.Namespace).Delete(revision.Name, nil)
}
func (rh *realHistory) AdoptControllerRevision(parent metav1.Object, parentKind schema.GroupVersionKind, revision *apps.ControllerRevision) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if owner := metav1.GetControllerOf(revision); owner != nil {
		return nil, fmt.Errorf("attempt to adopt revision owned by %v", owner)
	}
	return rh.client.AppsV1().ControllerRevisions(parent.GetNamespace()).Patch(revision.GetName(), types.StrategicMergePatchType, []byte(fmt.Sprintf(`{"metadata":{"ownerReferences":[{"apiVersion":"%s","kind":"%s","name":"%s","uid":"%s","controller":true,"blockOwnerDeletion":true}],"uid":"%s"}}`, parentKind.GroupVersion().String(), parentKind.Kind, parent.GetName(), parent.GetUID(), revision.UID)))
}
func (rh *realHistory) ReleaseControllerRevision(parent metav1.Object, revision *apps.ControllerRevision) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	released, err := rh.client.AppsV1().ControllerRevisions(revision.GetNamespace()).Patch(revision.GetName(), types.StrategicMergePatchType, []byte(fmt.Sprintf(`{"metadata":{"ownerReferences":[{"$patch":"delete","uid":"%s"}],"uid":"%s"}}`, parent.GetUID(), revision.UID)))
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		if errors.IsInvalid(err) {
			return nil, nil
		}
	}
	return released, err
}

type fakeHistory struct {
	indexer cache.Indexer
	lister  appslisters.ControllerRevisionLister
}

func (fh *fakeHistory) ListControllerRevisions(parent metav1.Object, selector labels.Selector) ([]*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	history, err := fh.lister.ControllerRevisions(parent.GetNamespace()).List(selector)
	if err != nil {
		return nil, err
	}
	var owned []*apps.ControllerRevision
	for i := range history {
		ref := metav1.GetControllerOf(history[i])
		if ref == nil || ref.UID == parent.GetUID() {
			owned = append(owned, history[i])
		}
	}
	return owned, err
}
func (fh *fakeHistory) addRevision(revision *apps.ControllerRevision) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(revision)
	if err != nil {
		return nil, err
	}
	obj, found, err := fh.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if found {
		foundRevision := obj.(*apps.ControllerRevision)
		return foundRevision, errors.NewAlreadyExists(apps.Resource("controllerrevision"), revision.Name)
	}
	return revision, fh.indexer.Update(revision)
}
func (fh *fakeHistory) CreateControllerRevision(parent metav1.Object, revision *apps.ControllerRevision, collisionCount *int32) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if collisionCount == nil {
		return nil, fmt.Errorf("collisionCount should not be nil")
	}
	clone := revision.DeepCopy()
	clone.Namespace = parent.GetNamespace()
	for {
		hash := HashControllerRevision(revision, collisionCount)
		clone.Name = ControllerRevisionName(parent.GetName(), hash)
		created, err := fh.addRevision(clone)
		if errors.IsAlreadyExists(err) {
			*collisionCount++
			continue
		}
		return created, err
	}
}
func (fh *fakeHistory) DeleteControllerRevision(revision *apps.ControllerRevision) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(revision)
	if err != nil {
		return err
	}
	obj, found, err := fh.indexer.GetByKey(key)
	if err != nil {
		return err
	}
	if !found {
		return errors.NewNotFound(apps.Resource("controllerrevisions"), revision.Name)
	}
	return fh.indexer.Delete(obj)
}
func (fh *fakeHistory) UpdateControllerRevision(revision *apps.ControllerRevision, newRevision int64) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clone := revision.DeepCopy()
	clone.Revision = newRevision
	return clone, fh.indexer.Update(clone)
}
func (fh *fakeHistory) AdoptControllerRevision(parent metav1.Object, parentKind schema.GroupVersionKind, revision *apps.ControllerRevision) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	blockOwnerDeletion := true
	isController := true
	if owner := metav1.GetControllerOf(revision); owner != nil {
		return nil, fmt.Errorf("attempt to adopt revision owned by %v", owner)
	}
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(revision)
	if err != nil {
		return nil, err
	}
	_, found, err := fh.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.NewNotFound(apps.Resource("controllerrevisions"), revision.Name)
	}
	clone := revision.DeepCopy()
	clone.OwnerReferences = append(clone.OwnerReferences, metav1.OwnerReference{APIVersion: parentKind.GroupVersion().String(), Kind: parentKind.Kind, Name: parent.GetName(), UID: parent.GetUID(), BlockOwnerDeletion: &blockOwnerDeletion, Controller: &isController})
	return clone, fh.indexer.Update(clone)
}
func (fh *fakeHistory) ReleaseControllerRevision(parent metav1.Object, revision *apps.ControllerRevision) (*apps.ControllerRevision, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(revision)
	if err != nil {
		return nil, err
	}
	_, found, err := fh.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	clone := revision.DeepCopy()
	refs := clone.OwnerReferences
	clone.OwnerReferences = nil
	for i := range refs {
		if refs[i].UID != parent.GetUID() {
			clone.OwnerReferences = append(clone.OwnerReferences, refs[i])
		}
	}
	return clone, fh.indexer.Update(clone)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
