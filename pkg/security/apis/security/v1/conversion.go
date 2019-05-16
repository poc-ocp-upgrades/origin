package v1

import (
	goformat "fmt"
	"github.com/openshift/api/security/v1"
	securityapi "github.com/openshift/origin/pkg/security/apis/security"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func AddConversionFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := scheme.AddConversionFuncs(Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints, Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints)
	if err != nil {
		return err
	}
	return nil
}
func Convert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(in *v1.SecurityContextConstraints, out *securityapi.SecurityContextConstraints, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return autoConvert_v1_SecurityContextConstraints_To_security_SecurityContextConstraints(in, out, s)
}
func Convert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(in *securityapi.SecurityContextConstraints, out *v1.SecurityContextConstraints, s conversion.Scope) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := autoConvert_security_SecurityContextConstraints_To_v1_SecurityContextConstraints(in, out, s); err != nil {
		return err
	}
	if in.Volumes != nil {
		for _, v := range in.Volumes {
			switch v {
			case securityapi.FSTypeHostPath, securityapi.FSTypeAll:
				out.AllowHostDirVolumePlugin = true
			}
		}
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
