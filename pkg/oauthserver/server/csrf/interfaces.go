package csrf

import "net/http"

type CSRF interface {
	Generate(http.ResponseWriter, *http.Request) string
	Check(*http.Request, string) bool
}
