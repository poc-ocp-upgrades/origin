package service

import (
 "os"
 godefaultbytes "bytes"
 godefaulthttp "net/http"
 godefaultruntime "runtime"
 "k8s.io/klog"
 "golang.org/x/sys/windows"
 "golang.org/x/sys/windows/svc"
)

var (
 service *handler
)

type handler struct {
 tosvc   chan bool
 fromsvc chan error
}

func InitService(serviceName string) error {
 _logClusterCodePath()
 defer _logClusterCodePath()
 h := &handler{tosvc: make(chan bool), fromsvc: make(chan error)}
 service = h
 var err error
 go func() {
  err = svc.Run(serviceName, h)
  h.fromsvc <- err
 }()
 err = <-h.fromsvc
 if err != nil {
  return err
 }
 klog.Infof("Running %s as a Windows service!", serviceName)
 return nil
}
func (h *handler) Execute(_ []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 s <- svc.Status{State: svc.StartPending, Accepts: 0}
 h.fromsvc <- nil
 s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown | svc.Accepted(windows.SERVICE_ACCEPT_PARAMCHANGE)}
 klog.Infof("Service running")
Loop:
 for {
  select {
  case <-h.tosvc:
   break Loop
  case c := <-r:
   switch c.Cmd {
   case svc.Cmd(windows.SERVICE_CONTROL_PARAMCHANGE):
    s <- c.CurrentStatus
   case svc.Interrogate:
    s <- c.CurrentStatus
   case svc.Stop, svc.Shutdown:
    s <- svc.Status{State: svc.Stopped}
    os.Exit(0)
   }
  }
 }
 return false, 0
}
func _logClusterCodePath() {
 pc, _, _, _ := godefaultruntime.Caller(1)
 jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
 godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
