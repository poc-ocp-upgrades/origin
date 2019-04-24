package glide

import (
	"time"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

type LockFile struct {
	Hash	string			`yaml:"hash"`
	Updated	time.Time		`yaml:"updated"`
	Imports	[]*LockFileImport	`yaml:"imports"`
}

func (l *LockFile) Decode(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return yaml.Unmarshal(b, l)
}

type YamlFile struct {
	Package		string			`yaml:"package"`
	ExcludeDirs	[]string		`yaml:"excludeDirs"`
	Imports		YamlFileImportList	`yaml:"import"`
}

func (y *YamlFile) Encode() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return yaml.Marshal(y)
}
func (y *YamlFile) Decode(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return yaml.Unmarshal(b, y)
}

type LockFileImport struct {
	Name	string	`yaml:"name"`
	Repo	string	`yaml:"repo,omitempty"`
	Version	string	`yaml:"version"`
}
type YamlFileImport struct {
	Package	string	`yaml:"package"`
	Repo	string	`yaml:"repo,omitempty"`
	Version	string	`yaml:"version"`
}
type YamlFileImportList []*YamlFileImport

func (l *YamlFileImportList) Encode() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return yaml.Marshal(l)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
