package cache

import (
	"fmt"
	goformat "fmt"
	projectapi "github.com/openshift/origin/pkg/project/apis/project"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func NewProjectCache(namespaces cache.SharedIndexInformer, client corev1client.NamespaceInterface, defaultNodeSelector string) *ProjectCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := namespaces.GetIndexer().AddIndexers(cache.Indexers{"requester": indexNamespaceByRequester}); err != nil {
		panic(err)
	}
	return &ProjectCache{Client: client, Store: namespaces.GetIndexer(), HasSynced: namespaces.GetController().HasSynced, DefaultNodeSelector: defaultNodeSelector}
}

type ProjectCache struct {
	Client              corev1client.NamespaceInterface
	Store               cache.Indexer
	HasSynced           cache.InformerSynced
	DefaultNodeSelector string
}

func (p *ProjectCache) GetNamespace(name string) (*corev1.Namespace, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	key := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	namespaceObj, exists, err := p.Store.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		time.Sleep(50 * time.Millisecond)
		namespaceObj, exists, err = p.Store.Get(key)
		if err != nil {
			return nil, err
		}
		if exists {
			klog.V(4).Infof("found %s in cache after waiting", name)
		}
	}
	var namespace *corev1.Namespace
	if exists {
		namespace = namespaceObj.(*corev1.Namespace)
	} else {
		namespace, err = p.Client.Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("namespace %s does not exist", name)
		}
		klog.V(4).Infof("found %s via storage lookup", name)
	}
	return namespace, nil
}
func (c *ProjectCache) Run(stopCh <-chan struct{}) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defer runtime.HandleCrash()
	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		return
	}
	<-stopCh
}
func (c *ProjectCache) Running() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.Store != nil
}
func NewFake(c corev1client.NamespaceInterface, store cache.Indexer, defaultNodeSelector string) *ProjectCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ProjectCache{Client: c, Store: store, DefaultNodeSelector: defaultNodeSelector}
}
func NewCacheStore(keyFn cache.KeyFunc) cache.Indexer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return cache.NewIndexer(keyFn, cache.Indexers{"requester": indexNamespaceByRequester})
}
func indexNamespaceByRequester(obj interface{}) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	requester := obj.(*corev1.Namespace).Annotations[projectapi.ProjectRequester]
	return []string{requester}, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
