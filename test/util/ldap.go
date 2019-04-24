package util

import "github.com/vjeantet/ldapserver"

type testLDAPServer struct {
	Passwords	map[string]string
	BindRequests	[]ldapserver.BindRequest
	SearchRequests	[]ldapserver.SearchRequest
	SearchResults	[]*ldapserver.SearchResultEntry
	server		*ldapserver.Server
}

func NewTestLDAPServer() *testLDAPServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &testLDAPServer{}
	routes := ldapserver.NewRouteMux()
	routes.Bind(t.handleBind)
	routes.Search(t.handleSearch)
	t.server = ldapserver.NewServer()
	t.server.Handle(routes)
	return t
}
func (t *testLDAPServer) Start(address string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go t.server.ListenAndServe(address)
}
func (t *testLDAPServer) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.server.Stop()
}
func (t *testLDAPServer) ResetRequests() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.BindRequests = []ldapserver.BindRequest{}
	t.SearchRequests = []ldapserver.SearchRequest{}
}
func (t *testLDAPServer) AddSearchResult(dn string, attributes map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := ldapserver.NewSearchResultEntry()
	e.SetDn(dn)
	for k, v := range attributes {
		e.AddAttribute(ldapserver.AttributeDescription(k), ldapserver.AttributeValue(v))
	}
	t.SearchResults = append(t.SearchResults, e)
}
func (t *testLDAPServer) SetPassword(dn string, password string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.Passwords == nil {
		t.Passwords = map[string]string{}
	}
	t.Passwords[dn] = password
}
func (t *testLDAPServer) handleBind(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := m.GetBindRequest()
	t.BindRequests = append(t.BindRequests, r)
	dn := string(r.GetLogin())
	password := string(r.GetPassword())
	if len(dn) == 0 || len(password) == 0 {
		w.Write(ldapserver.NewBindResponse(ldapserver.LDAPResultUnwillingToPerform))
		return
	}
	expectedPassword, ok := t.Passwords[dn]
	if !ok || expectedPassword != password {
		w.Write(ldapserver.NewBindResponse(ldapserver.LDAPResultInvalidCredentials))
		return
	}
	w.Write(ldapserver.NewBindResponse(ldapserver.LDAPResultSuccess))
}
func (t *testLDAPServer) handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := m.GetSearchRequest()
	t.SearchRequests = append(t.SearchRequests, r)
	for _, entry := range t.SearchResults {
		w.Write(entry)
	}
	w.Write(ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess))
}
