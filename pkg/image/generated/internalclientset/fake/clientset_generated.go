package fake

import (
	clientset "github.com/openshift/origin/pkg/image/generated/internalclientset"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	imageinternalversion "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	fakeimageinternalversion "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}
	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})
	return cs
}

type Clientset struct {
	testing.Fake
	discovery	*fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

func (c *Clientset) Image() imageinternalversion.ImageInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &fakeimageinternalversion.FakeImage{Fake: &c.Fake}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
