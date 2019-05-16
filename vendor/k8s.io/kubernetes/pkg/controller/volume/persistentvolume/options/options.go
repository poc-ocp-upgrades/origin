package options

import (
	goformat "fmt"
	"github.com/spf13/pflag"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type VolumeConfigFlags struct {
	PersistentVolumeRecyclerMaximumRetry                int
	PersistentVolumeRecyclerMinimumTimeoutNFS           int
	PersistentVolumeRecyclerPodTemplateFilePathNFS      string
	PersistentVolumeRecyclerIncrementTimeoutNFS         int
	PersistentVolumeRecyclerPodTemplateFilePathHostPath string
	PersistentVolumeRecyclerMinimumTimeoutHostPath      int
	PersistentVolumeRecyclerIncrementTimeoutHostPath    int
	EnableHostPathProvisioning                          bool
	EnableDynamicProvisioning                           bool
}
type PersistentVolumeControllerOptions struct {
	PVClaimBinderSyncPeriod time.Duration
	VolumeConfigFlags       VolumeConfigFlags
}

func NewPersistentVolumeControllerOptions() PersistentVolumeControllerOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return PersistentVolumeControllerOptions{PVClaimBinderSyncPeriod: 15 * time.Second, VolumeConfigFlags: VolumeConfigFlags{PersistentVolumeRecyclerMaximumRetry: 3, PersistentVolumeRecyclerMinimumTimeoutNFS: 300, PersistentVolumeRecyclerIncrementTimeoutNFS: 30, PersistentVolumeRecyclerMinimumTimeoutHostPath: 60, PersistentVolumeRecyclerIncrementTimeoutHostPath: 30, EnableHostPathProvisioning: false, EnableDynamicProvisioning: true}}
}
func (o *PersistentVolumeControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.DurationVar(&o.PVClaimBinderSyncPeriod, "pvclaimbinder-sync-period", o.PVClaimBinderSyncPeriod, "The period for syncing persistent volumes and persistent volume claims")
	fs.StringVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerPodTemplateFilePathNFS, "pv-recycler-pod-template-filepath-nfs", o.VolumeConfigFlags.PersistentVolumeRecyclerPodTemplateFilePathNFS, "The file path to a pod definition used as a template for NFS persistent volume recycling")
	fs.IntVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerMinimumTimeoutNFS, "pv-recycler-minimum-timeout-nfs", o.VolumeConfigFlags.PersistentVolumeRecyclerMinimumTimeoutNFS, "The minimum ActiveDeadlineSeconds to use for an NFS Recycler pod")
	fs.IntVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerIncrementTimeoutNFS, "pv-recycler-increment-timeout-nfs", o.VolumeConfigFlags.PersistentVolumeRecyclerIncrementTimeoutNFS, "the increment of time added per Gi to ActiveDeadlineSeconds for an NFS scrubber pod")
	fs.StringVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerPodTemplateFilePathHostPath, "pv-recycler-pod-template-filepath-hostpath", o.VolumeConfigFlags.PersistentVolumeRecyclerPodTemplateFilePathHostPath, "The file path to a pod definition used as a template for HostPath persistent volume recycling. "+"This is for development and testing only and will not work in a multi-node cluster.")
	fs.IntVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerMinimumTimeoutHostPath, "pv-recycler-minimum-timeout-hostpath", o.VolumeConfigFlags.PersistentVolumeRecyclerMinimumTimeoutHostPath, "The minimum ActiveDeadlineSeconds to use for a HostPath Recycler pod. This is for development and testing only and will not work in a multi-node cluster.")
	fs.IntVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerIncrementTimeoutHostPath, "pv-recycler-timeout-increment-hostpath", o.VolumeConfigFlags.PersistentVolumeRecyclerIncrementTimeoutHostPath, "the increment of time added per Gi to ActiveDeadlineSeconds for a HostPath scrubber pod. "+"This is for development and testing only and will not work in a multi-node cluster.")
	fs.IntVar(&o.VolumeConfigFlags.PersistentVolumeRecyclerMaximumRetry, "pv-recycler-maximum-retry", o.VolumeConfigFlags.PersistentVolumeRecyclerMaximumRetry, "Maximum number of attempts to recycle or delete a persistent volume")
	fs.BoolVar(&o.VolumeConfigFlags.EnableHostPathProvisioning, "enable-hostpath-provisioner", o.VolumeConfigFlags.EnableHostPathProvisioning, "Enable HostPath PV provisioning when running without a cloud provider. This allows testing and development of provisioning features. "+"HostPath provisioning is not supported in any way, won't work in a multi-node cluster, and should not be used for anything other than testing or development.")
	fs.BoolVar(&o.VolumeConfigFlags.EnableDynamicProvisioning, "enable-dynamic-provisioning", o.VolumeConfigFlags.EnableDynamicProvisioning, "Enable dynamic provisioning for environments that support it.")
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
