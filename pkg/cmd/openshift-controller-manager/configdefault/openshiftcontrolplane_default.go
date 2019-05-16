package configdefault

import (
	goformat "fmt"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	"github.com/openshift/library-go/pkg/config/configdefaults"
	leaderelectionconverter "github.com/openshift/library-go/pkg/config/leaderelection"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func SetRecommendedOpenShiftControllerConfigDefaults(config *openshiftcontrolplanev1.OpenShiftControllerManagerConfig) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	configdefaults.SetRecommendedHTTPServingInfoDefaults(config.ServingInfo)
	configdefaults.SetRecommendedKubeClientConfigDefaults(&config.KubeClientConfig)
	config.LeaderElection = leaderelectionconverter.LeaderElectionDefaulting(config.LeaderElection, "kube-system", "openshift-master-controllers")
	configdefaults.DefaultStringSlice(&config.Controllers, []string{"*"})
	configdefaults.DefaultString(&config.Network.ServiceNetworkCIDR, "10.0.0.0/24")
	if config.ImageImport.MaxScheduledImageImportsPerMinute == 0 {
		config.ImageImport.MaxScheduledImageImportsPerMinute = 60
	}
	if config.ImageImport.ScheduledImageImportMinimumIntervalSeconds == 0 {
		config.ImageImport.ScheduledImageImportMinimumIntervalSeconds = 15 * 60
	}
	configdefaults.DefaultString(&config.SecurityAllocator.UIDAllocatorRange, "1000000000-1999999999/10000")
	configdefaults.DefaultString(&config.SecurityAllocator.MCSAllocatorRange, "s0:/2")
	if config.SecurityAllocator.MCSLabelsPerProject == 0 {
		config.SecurityAllocator.MCSLabelsPerProject = 5
	}
	if config.ResourceQuota.MinResyncPeriod.Duration == 0 {
		config.ResourceQuota.MinResyncPeriod.Duration = 5 * time.Minute
	}
	if config.ResourceQuota.SyncPeriod.Duration == 0 {
		config.ResourceQuota.SyncPeriod.Duration = 12 * time.Hour
	}
	if config.ResourceQuota.ConcurrentSyncs == 0 {
		config.ResourceQuota.ConcurrentSyncs = 5
	}
	if config.ImageImport.MaxScheduledImageImportsPerMinute == 0 {
		config.ImageImport.MaxScheduledImageImportsPerMinute = 60
	}
	if config.ImageImport.ScheduledImageImportMinimumIntervalSeconds == 0 {
		config.ImageImport.ScheduledImageImportMinimumIntervalSeconds = 15 * 60
	}
	configdefaults.DefaultStringSlice(&config.ServiceAccount.ManagedNames, []string{"builder", "deployer"})
	configdefaults.DefaultString(&config.Deployer.ImageTemplateFormat.Format, "quay.io/openshift/origin-${component}:${version}")
	configdefaults.DefaultString(&config.Build.ImageTemplateFormat.Format, "quay.io/openshift/origin-${component}:${version}")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
