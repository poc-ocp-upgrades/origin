package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfig "k8s.io/kubernetes/pkg/controller/apis/config"
)

type PersistentVolumeBinderControllerOptions struct {
	PVClaimBinderSyncPeriod metav1.Duration
	VolumeConfiguration     kubectrlmgrconfig.VolumeConfiguration
}

func (o *PersistentVolumeBinderControllerOptions) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return
	}
	fs.DurationVar(&o.PVClaimBinderSyncPeriod.Duration, "pvclaimbinder-sync-period", o.PVClaimBinderSyncPeriod.Duration, "The period for syncing persistent volumes and persistent volume claims")
	fs.StringVar(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.PodTemplateFilePathNFS, "pv-recycler-pod-template-filepath-nfs", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.PodTemplateFilePathNFS, "The file path to a pod definition used as a template for NFS persistent volume recycling")
	fs.Int32Var(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.MinimumTimeoutNFS, "pv-recycler-minimum-timeout-nfs", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.MinimumTimeoutNFS, "The minimum ActiveDeadlineSeconds to use for an NFS Recycler pod")
	fs.Int32Var(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.IncrementTimeoutNFS, "pv-recycler-increment-timeout-nfs", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.IncrementTimeoutNFS, "the increment of time added per Gi to ActiveDeadlineSeconds for an NFS scrubber pod")
	fs.StringVar(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.PodTemplateFilePathHostPath, "pv-recycler-pod-template-filepath-hostpath", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.PodTemplateFilePathHostPath, "The file path to a pod definition used as a template for HostPath persistent volume recycling. This is for development and testing only and will not work in a multi-node cluster.")
	fs.Int32Var(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.MinimumTimeoutHostPath, "pv-recycler-minimum-timeout-hostpath", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.MinimumTimeoutHostPath, "The minimum ActiveDeadlineSeconds to use for a HostPath Recycler pod.  This is for development and testing only and will not work in a multi-node cluster.")
	fs.Int32Var(&o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.IncrementTimeoutHostPath, "pv-recycler-timeout-increment-hostpath", o.VolumeConfiguration.PersistentVolumeRecyclerConfiguration.IncrementTimeoutHostPath, "the increment of time added per Gi to ActiveDeadlineSeconds for a HostPath scrubber pod.  This is for development and testing only and will not work in a multi-node cluster.")
	fs.BoolVar(&o.VolumeConfiguration.EnableHostPathProvisioning, "enable-hostpath-provisioner", o.VolumeConfiguration.EnableHostPathProvisioning, "Enable HostPath PV provisioning when running without a cloud provider. This allows testing and development of provisioning features.  HostPath provisioning is not supported in any way, won't work in a multi-node cluster, and should not be used for anything other than testing or development.")
	fs.BoolVar(&o.VolumeConfiguration.EnableDynamicProvisioning, "enable-dynamic-provisioning", o.VolumeConfiguration.EnableDynamicProvisioning, "Enable dynamic provisioning for environments that support it.")
	fs.StringVar(&o.VolumeConfiguration.FlexVolumePluginDir, "flex-volume-plugin-dir", o.VolumeConfiguration.FlexVolumePluginDir, "Full path of the directory in which the flex volume plugin should search for additional third party volume plugins.")
}
func (o *PersistentVolumeBinderControllerOptions) ApplyTo(cfg *kubectrlmgrconfig.PersistentVolumeBinderControllerConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	cfg.PVClaimBinderSyncPeriod = o.PVClaimBinderSyncPeriod
	cfg.VolumeConfiguration = o.VolumeConfiguration
	return nil
}
func (o *PersistentVolumeBinderControllerOptions) Validate() []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o == nil {
		return nil
	}
	errs := []error{}
	return errs
}
