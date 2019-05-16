package options

import (
	"github.com/openshift/origin/pkg/cmd/flagtypes"
	"github.com/spf13/pflag"
)

type ListenArg struct{ ListenAddr flagtypes.Addr }

func BindListenArg(args *ListenArg, flags *pflag.FlagSet, prefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags.Var(&args.ListenAddr, prefix+"listen", "The address to listen for connections on (scheme://host:port).")
}
func NewDefaultListenArg() *ListenArg {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := &ListenArg{ListenAddr: flagtypes.Addr{Value: "0.0.0.0:8443", DefaultScheme: "https", DefaultPort: 8443, AllowPrefix: true}.Default()}
	return config
}
func (l *ListenArg) UseTLS() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return l.ListenAddr.URL.Scheme == "https"
}
