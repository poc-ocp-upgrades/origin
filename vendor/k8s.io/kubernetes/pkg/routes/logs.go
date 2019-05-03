package routes

import (
 "net/http"
 "path"
 "github.com/emicklei/go-restful"
)

type Logs struct{}

func (l Logs) Install(c *restful.Container) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ws := new(restful.WebService)
 ws.Path("/logs")
 ws.Doc("get log files")
 ws.Route(ws.GET("/{logpath:*}").To(logFileHandler).Param(ws.PathParameter("logpath", "path to the log").DataType("string")))
 ws.Route(ws.GET("/").To(logFileListHandler))
 c.Add(ws)
}
func logFileHandler(req *restful.Request, resp *restful.Response) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 logdir := "/var/log"
 actual := path.Join(logdir, req.PathParameter("logpath"))
 http.ServeFile(resp.ResponseWriter, req.Request, actual)
}
func logFileListHandler(req *restful.Request, resp *restful.Response) {
 _logClusterCodePath()
 defer _logClusterCodePath()
 logdir := "/var/log"
 http.ServeFile(resp.ResponseWriter, req.Request, logdir)
}
