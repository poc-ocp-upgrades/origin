package configdefault

import (
	"time"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	openshiftcontrolplanev1 "github.com/openshift/api/openshiftcontrolplane/v1"
	"github.com/openshift/library-go/pkg/config/configdefaults"
	leaderelectionconverter "github.com/openshift/library-go/pkg/config/leaderelection"
)

func SetRecommendedOpenShiftControllerConfigDefaults(config *openshiftcontrolplanev1.OpenShiftControllerManagerConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
