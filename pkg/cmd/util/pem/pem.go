package pem

import (
	"bytes"
	"encoding/pem"
	goformat "fmt"
	"io/ioutil"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func BlockFromFile(path string, blockType string) (*pem.Block, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, false, err
	}
	block, ok := BlockFromBytes(data, blockType)
	return block, ok, nil
}
func BlockFromBytes(data []byte, blockType string) (*pem.Block, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b := bytes.Buffer{}
	if err := pem.Encode(&b, block); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
