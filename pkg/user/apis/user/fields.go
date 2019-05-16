package user

import (
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/fields"
	runtime "k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func IdentityFieldSelector(obj runtime.Object, fieldSet fields.Set) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	identity, ok := obj.(*Identity)
	if !ok {
		return fmt.Errorf("%T not an Identity", obj)
	}
	fieldSet["providerName"] = identity.ProviderName
	fieldSet["providerUserName"] = identity.ProviderUserName
	fieldSet["user.name"] = identity.User.Name
	fieldSet["user.uid"] = string(identity.User.UID)
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
