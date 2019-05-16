package options

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func AddCertificateDirFlag(fs *pflag.FlagSet, certsDir *string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(certsDir, CertificatesDir, *certsDir, "The path where to save the certificates")
}
func AddCSRFlag(fs *pflag.FlagSet, csr *bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.BoolVar(csr, CSROnly, *csr, "Create CSRs instead of generating certificates")
}
func AddCSRDirFlag(fs *pflag.FlagSet, csrDir *string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(csrDir, CSRDir, *csrDir, "The path to output the CSRs and private keys to")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
