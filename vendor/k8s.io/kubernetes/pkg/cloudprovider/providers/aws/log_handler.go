package aws

import (
 "github.com/aws/aws-sdk-go/aws/request"
 "k8s.io/klog"
)

func awsHandlerLogger(req *request.Request) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service, name := awsServiceAndName(req)
 klog.V(4).Infof("AWS request: %s %s", service, name)
}
func awsSendHandlerLogger(req *request.Request) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service, name := awsServiceAndName(req)
 klog.V(4).Infof("AWS API Send: %s %s %v %v", service, name, req.Operation, req.Params)
}
func awsValidateResponseHandlerLogger(req *request.Request) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service, name := awsServiceAndName(req)
 klog.V(4).Infof("AWS API ValidateResponse: %s %s %v %v %s", service, name, req.Operation, req.Params, req.HTTPResponse.Status)
}
func awsServiceAndName(req *request.Request) (string, string) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 service := req.ClientInfo.ServiceName
 name := "?"
 if req.Operation != nil {
  name = req.Operation.Name
 }
 return service, name
}
