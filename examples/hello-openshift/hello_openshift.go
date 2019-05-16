package main

import (
	"fmt"
	goformat "fmt"
	"net/http"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift!"
	}
	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")
}
func listenAndServe(port string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)
	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)
	select {}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
