package routes

import (
	goformat "fmt"
	"github.com/emicklei/go-restful"
	"net/http"
	goos "os"
	"path"
	godefaultruntime "runtime"
	gotime "time"
)

type Logs struct{}

func (l Logs) Install(c *restful.Container) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ws := new(restful.WebService)
	ws.Path("/logs")
	ws.Doc("get log files")
	ws.Route(ws.GET("/{logpath:*}").To(logFileHandler).Param(ws.PathParameter("logpath", "path to the log").DataType("string")))
	ws.Route(ws.GET("/").To(logFileListHandler))
	c.Add(ws)
}
func logFileHandler(req *restful.Request, resp *restful.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	logdir := "/var/log"
	actual := path.Join(logdir, req.PathParameter("logpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}
func logFileListHandler(req *restful.Request, resp *restful.Response) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	logdir := "/var/log"
	http.ServeFile(resp.ResponseWriter, req.Request, logdir)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
