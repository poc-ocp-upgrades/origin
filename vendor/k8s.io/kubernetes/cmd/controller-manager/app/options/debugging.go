package options

import (
	"github.com/spf13/pflag"
	apiserverconfig "k8s.io/apiserver/pkg/apis/config"
)

type DebuggingOptions struct {
	EnableProfiling           bool
	EnableContentionProfiling bool
}

func (o *DebuggingOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.BoolVar(&o.EnableProfiling, "profiling", o.EnableProfiling, "Enable profiling via web interface host:port/debug/pprof/")
	fs.BoolVar(&o.EnableContentionProfiling, "contention-profiling", o.EnableContentionProfiling, "Enable lock contention profiling, if profiling is enabled")
}
func (o *DebuggingOptions) ApplyTo(cfg *apiserverconfig.DebuggingConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.EnableProfiling = o.EnableProfiling
	cfg.EnableContentionProfiling = o.EnableContentionProfiling
	return nil
}
func (o *DebuggingOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
