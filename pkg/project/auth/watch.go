package auth

import (
	"errors"
	"sync"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/authentication/user"
	kstorage "k8s.io/apiserver/pkg/storage"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	projectcache "github.com/openshift/origin/pkg/project/cache"
	projectutil "github.com/openshift/origin/pkg/project/util"
)

type CacheWatcher interface {
	GroupMembershipChanged(namespaceName string, users, groups sets.String)
}
type WatchableCache interface {
	RemoveWatcher(CacheWatcher)
	List(userInfo user.Info, selector labels.Selector) (*corev1.NamespaceList, error)
}
type userProjectWatcher struct {
	user			user.Info
	visibleNamespaces	sets.String
	cacheIncoming		chan watch.Event
	cacheError		chan error
	outgoing		chan watch.Event
	userStop		chan struct{}
	stopLock		sync.Mutex
	emit			func(watch.Event)
	projectCache		*projectcache.ProjectCache
	authCache		WatchableCache
	initialProjects		[]corev1.Namespace
	knownProjects		map[string]string
}

var (
	watchChannelHWM kstorage.HighWaterMark
)

func NewUserProjectWatcher(user user.Info, visibleNamespaces sets.String, projectCache *projectcache.ProjectCache, authCache WatchableCache, includeAllExistingProjects bool, predicate kstorage.SelectionPredicate) *userProjectWatcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaces, _ := authCache.List(user, labels.Everything())
	knownProjects := map[string]string{}
	for _, namespace := range namespaces.Items {
		knownProjects[namespace.Name] = namespace.ResourceVersion
	}
	initialProjects := []corev1.Namespace{}
	if includeAllExistingProjects {
		initialProjects = append(initialProjects, namespaces.Items...)
	}
	w := &userProjectWatcher{user: user, visibleNamespaces: visibleNamespaces, cacheIncoming: make(chan watch.Event, 1000), cacheError: make(chan error, 1), outgoing: make(chan watch.Event), userStop: make(chan struct{}), projectCache: projectCache, authCache: authCache, initialProjects: initialProjects, knownProjects: knownProjects}
	w.emit = func(e watch.Event) {
		if project, ok := e.Object.(*projectapi.Project); ok {
			if matches, err := predicate.Matches(project); err != nil || !matches {
				return
			}
		}
		select {
		case w.outgoing <- e:
		case <-w.userStop:
		}
	}
	return w
}
func (w *userProjectWatcher) GroupMembershipChanged(namespaceName string, users, groups sets.String) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !w.visibleNamespaces.Has("*") && !w.visibleNamespaces.Has(namespaceName) {
		return
	}
	hasAccess := users.Has(w.user.GetName()) || groups.HasAny(w.user.GetGroups()...)
	_, known := w.knownProjects[namespaceName]
	switch {
	case !hasAccess && known:
		delete(w.knownProjects, namespaceName)
		select {
		case w.cacheIncoming <- watch.Event{Type: watch.Deleted, Object: projectutil.ConvertNamespaceFromExternal(&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespaceName}})}:
		default:
			w.authCache.RemoveWatcher(w)
			w.cacheError <- errors.New("delete notification timeout")
		}
	case hasAccess:
		namespace, err := w.projectCache.GetNamespace(namespaceName)
		if err != nil {
			utilruntime.HandleError(err)
			return
		}
		event := watch.Event{Type: watch.Added, Object: projectutil.ConvertNamespaceFromExternal(namespace)}
		if lastResourceVersion, known := w.knownProjects[namespaceName]; known {
			event.Type = watch.Modified
			if lastResourceVersion == namespace.ResourceVersion {
				return
			}
		}
		w.knownProjects[namespaceName] = namespace.ResourceVersion
		select {
		case w.cacheIncoming <- event:
		default:
			w.authCache.RemoveWatcher(w)
			w.cacheError <- errors.New("add notification timeout")
		}
	}
}
func (w *userProjectWatcher) Watch() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(w.outgoing)
	defer func() {
		w.authCache.RemoveWatcher(w)
	}()
	defer utilruntime.HandleCrash()
	for i := range w.initialProjects {
		select {
		case err := <-w.cacheError:
			w.emit(makeErrorEvent(err))
			return
		default:
		}
		w.emit(watch.Event{Type: watch.Added, Object: projectutil.ConvertNamespaceFromExternal(&w.initialProjects[i])})
	}
	for {
		select {
		case err := <-w.cacheError:
			w.emit(makeErrorEvent(err))
			return
		case <-w.userStop:
			return
		case event := <-w.cacheIncoming:
			if curLen := int64(len(w.cacheIncoming)); watchChannelHWM.Update(curLen) {
				klog.V(2).Infof("watch: %v objects queued in project cache watching channel.", curLen)
			}
			w.emit(event)
		}
	}
}
func makeErrorEvent(err error) watch.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return watch.Event{Type: watch.Error, Object: &metav1.Status{Status: metav1.StatusFailure, Message: err.Error()}}
}
func (w *userProjectWatcher) ResultChan() <-chan watch.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.outgoing
}
func (w *userProjectWatcher) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.stopLock.Lock()
	defer w.stopLock.Unlock()
	select {
	case <-w.userStop:
		return
	default:
	}
	close(w.userStop)
}
