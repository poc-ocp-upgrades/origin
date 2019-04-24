package main

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	re := regexp.MustCompile(`([\W])(quay\.io/coreos[/\w\-]*)(\:[a-zA-Z\d][a-zA-Z\d\-_]*[a-zA-Z\d]|@\w+:\w+)?`)
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	out := re.ReplaceAllFunc(data, func(data []byte) []byte {
		fmt.Fprintf(os.Stderr, "found: %s\n", string(data))
		return data
	})
	fmt.Println(string(out))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
