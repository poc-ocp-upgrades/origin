package nodeidentifier

import (
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/user"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func NewDefaultNodeIdentifier() NodeIdentifier {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return defaultNodeIdentifier{}
}

type defaultNodeIdentifier struct{}

const nodeUserNamePrefix = "system:node:"

func (defaultNodeIdentifier) NodeIdentity(u user.Info) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if u == nil {
		return "", false
	}
	userName := u.GetName()
	if !strings.HasPrefix(userName, nodeUserNamePrefix) {
		return "", false
	}
	isNode := false
	for _, g := range u.GetGroups() {
		if g == user.NodesGroup {
			isNode = true
			break
		}
	}
	if !isNode {
		return "", false
	}
	nodeName := strings.TrimPrefix(userName, nodeUserNamePrefix)
	return nodeName, true
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
