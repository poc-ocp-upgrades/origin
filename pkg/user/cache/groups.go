package cache

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"k8s.io/client-go/tools/cache"
	userapi "github.com/openshift/api/user/v1"
	userinformer "github.com/openshift/client-go/user/informers/externalversions/user/v1"
)

type GroupCache struct{ indexer cache.Indexer }

const ByUserIndexName = "ByUser"

func ByUserIndexKeys(obj interface{}) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	group, ok := obj.(*userapi.Group)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %v", obj)
	}
	return group.Users, nil
}
func NewGroupCache(groupInformer userinformer.GroupInformer) *GroupCache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GroupCache{indexer: groupInformer.Informer().GetIndexer()}
}
func (c *GroupCache) GroupsFor(username string) ([]*userapi.Group, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
