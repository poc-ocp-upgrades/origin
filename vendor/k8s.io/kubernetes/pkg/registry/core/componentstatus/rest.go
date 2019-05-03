package componentstatus

import (
 "context"
 "fmt"
 "sync"
 metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
 metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 "k8s.io/apimachinery/pkg/runtime"
 "k8s.io/apiserver/pkg/registry/rest"
 api "k8s.io/kubernetes/pkg/apis/core"
 "k8s.io/kubernetes/pkg/probe"
)

type REST struct{ GetServersToValidate func() map[string]*Server }

func NewStorage(serverRetriever func() map[string]*Server) *REST {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &REST{GetServersToValidate: serverRetriever}
}
func (*REST) NamespaceScoped() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return false
}
func (rs *REST) New() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.ComponentStatus{}
}
func (rs *REST) NewList() runtime.Object {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &api.ComponentStatusList{}
}
func (rs *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
 servers := rs.GetServersToValidate()
 if server, ok := servers[name]; !ok {
  return nil, fmt.Errorf("Component not found: %s", name)
 } else {
  return rs.getComponentStatus(name, server), nil
 }
}
func ToConditionStatus(s probe.Result) api.ConditionStatus {
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
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
 _logClusterCodePath()
 defer _logClusterCodePath()
 return []string{"cs"}
}
