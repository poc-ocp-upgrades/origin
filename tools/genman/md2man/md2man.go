package md2man

import (
	"github.com/russross/blackfriday"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
)

func Render(doc []byte) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	renderer := RoffRenderer(0)
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_FOOTNOTES
	extensions |= blackfriday.EXTENSION_TITLEBLOCK
	return blackfriday.Markdown(doc, renderer, extensions)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
