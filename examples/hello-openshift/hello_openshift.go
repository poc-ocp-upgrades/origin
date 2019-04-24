package main

import (
	"fmt"
	"bytes"
	"runtime"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift!"
	}
	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")
}
func listenAndServe(port string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
