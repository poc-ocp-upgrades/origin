package main

import (
	"fmt"
	goformat "fmt"
	"log"
	"net/http"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

var (
	version  string
	subtitle string
	color    string
)

const htmlContent = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Deployment Demonstration</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    HTML{height:100%%;}
    BODY{font-family:Helvetica,Arial;display:flex;display:-webkit-flex;align-items:center;justify-content:center;-webkit-align-items:center;-webkit-box-align:center;-webkit-justify-content:center;height:100%%;}
    .box{background:%[3]s;color:white;text-align:center;border-radius:10px;display:inline-block;}
    H1{font-size:10em;line-height:1.5em;margin:0 0.5em;}
    H2{margin-top:0;}
  </style>
</head>
<body>
<div class="box"><h1>%[1]s</h1><h2>%[2]s</h2></div>
</body>
</html>`

func deploymentHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Fprintf(w, htmlContent, version, subtitle, color)
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Fprintln(w, "ok")
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	version = "v1"
	if len(os.Args) > 1 {
		version = os.Args[1]
	}
	subtitle = os.Getenv("SUBTITLE")
	color = os.Getenv("COLOR")
	if len(color) == 0 {
		color = "#303030"
	}
	http.HandleFunc("/", deploymentHandler)
	http.HandleFunc("/_healthz", healthHandler)
	log.Printf("Listening on :8080 at %s ...", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
