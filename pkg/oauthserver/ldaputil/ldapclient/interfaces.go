package ldapclient

import (
	goformat "fmt"
	"gopkg.in/ldap.v2"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Config interface {
	Connect() (client ldap.Client, err error)
	GetBindCredentials() (bindDN, bindPassword string)
	Host() string
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
