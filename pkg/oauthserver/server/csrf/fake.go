package csrf

import "net/http"

type FakeCSRF struct{ Token string }

func (c *FakeCSRF) Generate(w http.ResponseWriter, req *http.Request) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.Token
}
func (c *FakeCSRF) Check(req *http.Request, value string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.Token == value
}
