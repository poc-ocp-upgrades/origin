package componentstatus

import (
	"context"
	"fmt"
	goformat "fmt"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/probe"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

type REST struct{ GetServersToValidate func() map[string]*Server }

func NewStorage(serverRetriever func() map[string]*Server) *REST {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &REST{GetServersToValidate: serverRetriever}
}
func (*REST) NamespaceScoped() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return false
}
func (rs *REST) New() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &api.ComponentStatus{}
}
func (rs *REST) NewList() runtime.Object {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &api.ComponentStatusList{}
}
func (rs *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	servers := rs.GetServersToValidate()
	wait := sync.WaitGroup{}
	wait.Add(len(servers))
	statuses := make(chan api.ComponentStatus, len(servers))
	for k, v := range servers {
		go func(name string, server *Server) {
			defer wait.Done()
			status := rs.getComponentStatus(name, server)
			statuses <- *status
		}(k, v)
	}
	wait.Wait()
	close(statuses)
	reply := []api.ComponentStatus{}
	for status := range statuses {
		reply = append(reply, status)
	}
	return &api.ComponentStatusList{Items: reply}, nil
}
func (rs *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	servers := rs.GetServersToValidate()
	if server, ok := servers[name]; !ok {
		return nil, fmt.Errorf("Component not found: %s", name)
	} else {
		return rs.getComponentStatus(name, server), nil
	}
}
func ToConditionStatus(s probe.Result) api.ConditionStatus {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch s {
	case probe.Success:
		return api.ConditionTrue
	case probe.Failure:
		return api.ConditionFalse
	default:
		return api.ConditionUnknown
	}
}
func (rs *REST) getComponentStatus(name string, server *Server) *api.ComponentStatus {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	status, msg, err := server.DoServerCheck()
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	c := &api.ComponentCondition{Type: api.ComponentHealthy, Status: ToConditionStatus(status), Message: msg, Error: errorMsg}
	retVal := &api.ComponentStatus{Conditions: []api.ComponentCondition{*c}}
	retVal.Name = name
	return retVal
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []string{"cs"}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
