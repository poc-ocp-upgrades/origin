package openshift_controller_manager

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/coreos/go-systemd/daemon"
	configv1 "github.com/openshift/api/config/v1"
	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	"github.com/openshift/library-go/pkg/config/helpers"
	"github.com/openshift/library-go/pkg/serviceability"
	"github.com/openshift/origin/pkg/cmd/openshift-controller-manager/configdefault"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/cmd/server/apis/config/validation"
	"github.com/openshift/origin/pkg/configconversion"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"os"
	goos "os"
	"path"
	godefaultruntime "runtime"
	gotime "time"
)

const RecommendedStartControllerManagerName = "openshift-controller-manager"

type OpenShiftControllerManager struct {
	ConfigFilePath string
	Output         io.Writer
}

var longDescription = templates.LongDesc(`
	Start the OpenShift controllers`)

func NewOpenShiftControllerManagerCommand(name, basename string, out, errout io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	options := &OpenShiftControllerManager{Output: out}
	cmd := &cobra.Command{Use: name, Short: "Start the OpenShift controllers", Long: longDescription, Run: func(c *cobra.Command, args []string) {
		rest.CommandNameOverride = name
		kcmdutil.CheckErr(options.Validate())
		serviceability.StartProfiler()
		if err := options.StartControllerManager(); err != nil {
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
	flags.StringVar(&options.ConfigFilePath, "config", options.ConfigFilePath, "Location of the master configuration file to run from.")
	cmd.MarkFlagFilename("config", "yaml", "yml")
	cmd.MarkFlagRequired("config")
	return cmd
}
func (o *OpenShiftControllerManager) Validate() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(o.ConfigFilePath) == 0 {
		return errors.New("--config is required for this command")
	}
	return nil
}
func (o *OpenShiftControllerManager) StartControllerManager() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := o.RunControllerManager(); err != nil {
		return err
	}
	go daemon.SdNotify(false, "READY=1")
	select {}
}
func (o *OpenShiftControllerManager) RunControllerManager() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configContent, err := ioutil.ReadFile(o.ConfigFilePath)
	if err != nil {
		return err
	}
	scheme := runtime.NewScheme()
	utilruntime.Must(openshiftcontrolplanev1.Install(scheme))
	codecs := serializer.NewCodecFactory(scheme)
	obj, err := runtime.Decode(codecs.UniversalDecoder(openshiftcontrolplanev1.GroupVersion, configv1.GroupVersion), configContent)
	switch {
	case runtime.IsMissingVersion(err):
	case runtime.IsMissingKind(err):
	case runtime.IsNotRegisteredError(err):
	case err != nil:
		return err
	case err == nil:
		absoluteConfigFile, err := api.MakeAbs(o.ConfigFilePath, "")
		if err != nil {
			return err
		}
		configFileLocation := path.Dir(absoluteConfigFile)
		config := obj.(*openshiftcontrolplanev1.OpenShiftControllerManagerConfig)
		if config.ServingInfo == nil {
			config.ServingInfo = &configv1.HTTPServingInfo{}
		}
		if err := helpers.ResolvePaths(configconversion.GetOpenShiftControllerConfigFileReferences(config), configFileLocation); err != nil {
			return err
		}
		configdefault.SetRecommendedOpenShiftControllerConfigDefaults(config)
		clientConfig, err := helpers.GetKubeClientConfig(config.KubeClientConfig)
		if err != nil {
			return err
		}
		return RunOpenShiftControllerManager(config, clientConfig)
	}
	masterConfig, err := configapilatest.ReadAndResolveMasterConfig(o.ConfigFilePath)
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
		return kerrors.NewInvalid(configapi.Kind("MasterConfig"), "master-config.yaml", validationResults.Errors)
	}
	externalMasterConfig, err := configapi.Scheme.ConvertToVersion(masterConfig, legacyconfigv1.LegacySchemeGroupVersion)
	if err != nil {
		return err
	}
	config := ConvertMasterConfigToOpenshiftControllerConfig(externalMasterConfig.(*legacyconfigv1.MasterConfig))
	clientConfig, err := helpers.GetKubeClientConfig(config.KubeClientConfig)
	if err != nil {
		return err
	}
	return RunOpenShiftControllerManager(config, clientConfig)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
