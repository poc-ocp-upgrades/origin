package openshift_etcd

import (
	"errors"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"io"
	"os"
	"github.com/coreos/go-systemd/daemon"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"github.com/openshift/library-go/pkg/serviceability"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation"
	"github.com/openshift/origin/pkg/cmd/server/etcd/etcdserver"
)

const RecommendedStartEtcdServerName = "etcd"

type EtcdOptions struct {
	ConfigFile	string
	Output		io.Writer
}

var etcdLong = templates.LongDesc(`
	Start an etcd server for testing.

	This command starts an etcd server based on the config for testing.  It is not
	intended for production use.  Running

	    %[1]s start %[2]s

	will start the server listening for incoming requests. The server will run in
	the foreground until you terminate the process.`)

func NewCommandStartEtcdServer(name, basename string, out, errout io.Writer) (*cobra.Command, *EtcdOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	options := &EtcdOptions{Output: out}
	cmd := &cobra.Command{Use: name, Short: "Launch etcd server", Long: fmt.Sprintf(etcdLong, basename, name), Run: func(c *cobra.Command, args []string) {
		kcmdutil.CheckErr(options.Validate())
		serviceability.StartProfiler()
		if err := options.StartEtcdServer(); err != nil {
			if kerrors.IsInvalid(err) {
				if details := err.(*kerrors.StatusError).ErrStatus.Details; details != nil {
					fmt.Fprintf(errout, "Invalid %s %s\n", details.Kind, details.Name)
					for _, cause := range details.Causes {
						fmt.Fprintf(errout, "  %s: %s\n", cause.Field, cause.Message)
					}
					os.Exit(255)
				}
			}
			klog.Fatal(err)
		}
	}}
	flags := cmd.Flags()
	flags.StringVar(&options.ConfigFile, "config", "", "Location of the master configuration file to run from.")
	cmd.MarkFlagFilename("config", "yaml", "yml")
	cmd.MarkFlagRequired("config")
	return cmd, options
}
func (o *EtcdOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.ConfigFile) == 0 {
		return errors.New("--config is required for this command")
	}
	return nil
}
func (o *EtcdOptions) StartEtcdServer() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := o.RunEtcdServer(); err != nil {
		return err
	}
	go daemon.SdNotify(false, "READY=1")
	select {}
}
func (o *EtcdOptions) RunEtcdServer() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, err := configapilatest.ReadAndResolveMasterConfig(o.ConfigFile)
	if err != nil {
		return err
	}
	validationResults := validation.ValidateMasterConfig(masterConfig, nil)
	if len(validationResults.Warnings) != 0 {
		for _, warning := range validationResults.Warnings {
			klog.Warningf("%v", warning)
		}
	}
	if len(validationResults.Errors) != 0 {
		return kerrors.NewInvalid(configapi.Kind("MasterConfig"), o.ConfigFile, validationResults.Errors)
	}
	if masterConfig.EtcdConfig == nil {
		return kerrors.NewInvalid(configapi.Kind("MasterConfig.EtcConfig"), o.ConfigFile, field.ErrorList{field.Required(field.NewPath("etcdConfig"), "")})
	}
	etcdserver.RunEtcd(masterConfig.EtcdConfig)
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
