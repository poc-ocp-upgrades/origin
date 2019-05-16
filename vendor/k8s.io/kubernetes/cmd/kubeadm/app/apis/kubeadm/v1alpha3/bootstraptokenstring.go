package v1alpha3

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type BootstrapTokenString struct {
	ID     string
	Secret string
}

func (bts BootstrapTokenString) MarshalJSON() ([]byte, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return []byte(fmt.Sprintf(`"%s"`, bts.String())), nil
}
func (bts *BootstrapTokenString) UnmarshalJSON(b []byte) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(b) == 0 {
		return nil
	}
	token := strings.Replace(string(b), `"`, ``, -1)
	newbts, err := NewBootstrapTokenString(token)
	if err != nil {
		return err
	}
	bts.ID = newbts.ID
	bts.Secret = newbts.Secret
	return nil
}
func (bts BootstrapTokenString) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(bts.ID) > 0 && len(bts.Secret) > 0 {
		return bootstraputil.TokenFromIDAndSecret(bts.ID, bts.Secret)
	}
	return ""
}
func NewBootstrapTokenString(token string) (*BootstrapTokenString, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	substrs := bootstraputil.BootstrapTokenRegexp.FindStringSubmatch(token)
	if len(substrs) != 3 {
		return nil, errors.Errorf("the bootstrap token %q was not of the form %q", token, bootstrapapi.BootstrapTokenPattern)
	}
	return &BootstrapTokenString{ID: substrs[1], Secret: substrs[2]}, nil
}
func NewBootstrapTokenStringFromIDAndSecret(id, secret string) (*BootstrapTokenString, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return NewBootstrapTokenString(bootstraputil.TokenFromIDAndSecret(id, secret))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
