package cache

import (
	"fmt"
	goformat "fmt"
	userapi "github.com/openshift/api/user/v1"
	userinformer "github.com/openshift/client-go/user/informers/externalversions/user/v1"
	"k8s.io/client-go/tools/cache"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type GroupCache struct{ indexer cache.Indexer }

const ByUserIndexName = "ByUser"

func ByUserIndexKeys(obj interface{}) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	group, ok := obj.(*userapi.Group)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %v", obj)
	}
	return group.Users, nil
}
func NewGroupCache(groupInformer userinformer.GroupInformer) *GroupCache {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &GroupCache{indexer: groupInformer.Informer().GetIndexer()}
}
func (c *GroupCache) GroupsFor(username string) ([]*userapi.Group, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	objs, err := c.indexer.ByIndex(ByUserIndexName, username)
	if err != nil {
		return nil, err
	}
	groups := make([]*userapi.Group, len(objs))
	for i := range objs {
		groups[i] = objs[i].(*userapi.Group)
	}
	return groups, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
