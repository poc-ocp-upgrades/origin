package openstack

import (
	"context"
	"errors"
	"fmt"
	"github.com/gophercloud/gophercloud"
	volumeexpand "github.com/gophercloud/gophercloud/openstack/blockstorage/extensions/volumeactions"
	volumes_v1 "github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	volumes_v2 "github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	volumes_v3 "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	k8s_volume "k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type volumeService interface {
	createVolume(opts volumeCreateOpts) (string, string, error)
	getVolume(volumeID string) (Volume, error)
	deleteVolume(volumeName string) error
	expandVolume(volumeID string, newSize int) error
}
type VolumesV1 struct {
	blockstorage *gophercloud.ServiceClient
	opts         BlockStorageOpts
}
type VolumesV2 struct {
	blockstorage *gophercloud.ServiceClient
	opts         BlockStorageOpts
}
type VolumesV3 struct {
	blockstorage *gophercloud.ServiceClient
	opts         BlockStorageOpts
}
type Volume struct {
	AttachedServerID string
	AttachedDevice   string
	AvailabilityZone string
	ID               string
	Name             string
	Status           string
	Size             int
}
type volumeCreateOpts struct {
	Size         int
	Availability string
	Name         string
	VolumeType   string
	Metadata     map[string]string
}

var _ cloudprovider.PVLabeler = (*OpenStack)(nil)

const (
	volumeAvailableStatus = "available"
	volumeInUseStatus     = "in-use"
	volumeDeletedStatus   = "deleted"
	volumeErrorStatus     = "error"
	newtonMetadataVersion = "2016-06-30"
)

func (volumes *VolumesV1) createVolume(opts volumeCreateOpts) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumes_v1.CreateOpts{Name: opts.Name, Size: opts.Size, VolumeType: opts.VolumeType, AvailabilityZone: opts.Availability, Metadata: opts.Metadata}
	vol, err := volumes_v1.Create(volumes.blockstorage, createOpts).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("create_v1_volume", timeTaken, err)
	if err != nil {
		return "", "", err
	}
	return vol.ID, vol.AvailabilityZone, nil
}
func (volumes *VolumesV2) createVolume(opts volumeCreateOpts) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumes_v2.CreateOpts{Name: opts.Name, Size: opts.Size, VolumeType: opts.VolumeType, AvailabilityZone: opts.Availability, Metadata: opts.Metadata}
	vol, err := volumes_v2.Create(volumes.blockstorage, createOpts).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("create_v2_volume", timeTaken, err)
	if err != nil {
		return "", "", err
	}
	return vol.ID, vol.AvailabilityZone, nil
}
func (volumes *VolumesV3) createVolume(opts volumeCreateOpts) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumes_v3.CreateOpts{Name: opts.Name, Size: opts.Size, VolumeType: opts.VolumeType, AvailabilityZone: opts.Availability, Metadata: opts.Metadata}
	vol, err := volumes_v3.Create(volumes.blockstorage, createOpts).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("create_v3_volume", timeTaken, err)
	if err != nil {
		return "", "", err
	}
	return vol.ID, vol.AvailabilityZone, nil
}
func (volumes *VolumesV1) getVolume(volumeID string) (Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	volumeV1, err := volumes_v1.Get(volumes.blockstorage, volumeID).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("get_v1_volume", timeTaken, err)
	if err != nil {
		return Volume{}, fmt.Errorf("error occurred getting volume by ID: %s, err: %v", volumeID, err)
	}
	volume := Volume{AvailabilityZone: volumeV1.AvailabilityZone, ID: volumeV1.ID, Name: volumeV1.Name, Status: volumeV1.Status, Size: volumeV1.Size}
	if len(volumeV1.Attachments) > 0 && volumeV1.Attachments[0]["server_id"] != nil {
		volume.AttachedServerID = volumeV1.Attachments[0]["server_id"].(string)
		volume.AttachedDevice = volumeV1.Attachments[0]["device"].(string)
	}
	return volume, nil
}
func (volumes *VolumesV2) getVolume(volumeID string) (Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	volumeV2, err := volumes_v2.Get(volumes.blockstorage, volumeID).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("get_v2_volume", timeTaken, err)
	if err != nil {
		return Volume{}, fmt.Errorf("error occurred getting volume by ID: %s, err: %v", volumeID, err)
	}
	volume := Volume{AvailabilityZone: volumeV2.AvailabilityZone, ID: volumeV2.ID, Name: volumeV2.Name, Status: volumeV2.Status, Size: volumeV2.Size}
	if len(volumeV2.Attachments) > 0 {
		volume.AttachedServerID = volumeV2.Attachments[0].ServerID
		volume.AttachedDevice = volumeV2.Attachments[0].Device
	}
	return volume, nil
}
func (volumes *VolumesV3) getVolume(volumeID string) (Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	volumeV3, err := volumes_v3.Get(volumes.blockstorage, volumeID).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("get_v3_volume", timeTaken, err)
	if err != nil {
		return Volume{}, fmt.Errorf("error occurred getting volume by ID: %s, err: %v", volumeID, err)
	}
	volume := Volume{AvailabilityZone: volumeV3.AvailabilityZone, ID: volumeV3.ID, Name: volumeV3.Name, Status: volumeV3.Status, Size: volumeV3.Size}
	if len(volumeV3.Attachments) > 0 {
		volume.AttachedServerID = volumeV3.Attachments[0].ServerID
		volume.AttachedDevice = volumeV3.Attachments[0].Device
	}
	return volume, nil
}
func (volumes *VolumesV1) deleteVolume(volumeID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	err := volumes_v1.Delete(volumes.blockstorage, volumeID).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("delete_v1_volume", timeTaken, err)
	return err
}
func (volumes *VolumesV2) deleteVolume(volumeID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	err := volumes_v2.Delete(volumes.blockstorage, volumeID).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("delete_v2_volume", timeTaken, err)
	return err
}
func (volumes *VolumesV3) deleteVolume(volumeID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	err := volumes_v3.Delete(volumes.blockstorage, volumeID).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("delete_v3_volume", timeTaken, err)
	return err
}
func (volumes *VolumesV1) expandVolume(volumeID string, newSize int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumeexpand.ExtendSizeOpts{NewSize: newSize}
	err := volumeexpand.ExtendSize(volumes.blockstorage, volumeID, createOpts).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("expand_volume", timeTaken, err)
	return err
}
func (volumes *VolumesV2) expandVolume(volumeID string, newSize int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumeexpand.ExtendSizeOpts{NewSize: newSize}
	err := volumeexpand.ExtendSize(volumes.blockstorage, volumeID, createOpts).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("expand_volume", timeTaken, err)
	return err
}
func (volumes *VolumesV3) expandVolume(volumeID string, newSize int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	startTime := time.Now()
	createOpts := volumeexpand.ExtendSizeOpts{NewSize: newSize}
	err := volumeexpand.ExtendSize(volumes.blockstorage, volumeID, createOpts).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("expand_volume", timeTaken, err)
	return err
}
func (os *OpenStack) OperationPending(diskName string) (bool, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(diskName)
	if err != nil {
		return false, "", err
	}
	volumeStatus := volume.Status
	if volumeStatus == volumeErrorStatus {
		err = fmt.Errorf("status of volume %s is %s", diskName, volumeStatus)
		return false, volumeStatus, err
	}
	if volumeStatus == volumeAvailableStatus || volumeStatus == volumeInUseStatus || volumeStatus == volumeDeletedStatus {
		return false, volume.Status, nil
	}
	return true, volumeStatus, nil
}
func (os *OpenStack) AttachDisk(instanceID, volumeID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return "", err
	}
	cClient, err := os.NewComputeV2()
	if err != nil {
		return "", err
	}
	if volume.AttachedServerID != "" {
		if instanceID == volume.AttachedServerID {
			klog.V(4).Infof("Disk %s is already attached to instance %s", volumeID, instanceID)
			return volume.ID, nil
		}
		nodeName, err := os.GetNodeNameByID(volume.AttachedServerID)
		attachErr := fmt.Sprintf("disk %s path %s is attached to a different instance (%s)", volumeID, volume.AttachedDevice, volume.AttachedServerID)
		if err != nil {
			klog.Error(attachErr)
			return "", errors.New(attachErr)
		}
		devicePath := volume.AttachedDevice
		danglingErr := volumeutil.NewDanglingError(attachErr, nodeName, devicePath)
		klog.V(2).Infof("Found dangling volume %s attached to node %s", volumeID, nodeName)
		return "", danglingErr
	}
	startTime := time.Now()
	_, err = volumeattach.Create(cClient, instanceID, &volumeattach.CreateOpts{VolumeID: volume.ID}).Extract()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("attach_disk", timeTaken, err)
	if err != nil {
		return "", fmt.Errorf("failed to attach %s volume to %s compute: %v", volumeID, instanceID, err)
	}
	klog.V(2).Infof("Successfully attached %s volume to %s compute", volumeID, instanceID)
	return volume.ID, nil
}
func (os *OpenStack) DetachDisk(instanceID, volumeID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return err
	}
	if volume.Status == volumeAvailableStatus {
		klog.V(2).Infof("volume: %s has been detached from compute: %s ", volume.ID, instanceID)
		return nil
	}
	if volume.Status != volumeInUseStatus {
		return fmt.Errorf("can not detach volume %s, its status is %s", volume.Name, volume.Status)
	}
	cClient, err := os.NewComputeV2()
	if err != nil {
		return err
	}
	if volume.AttachedServerID != instanceID {
		return fmt.Errorf("disk: %s has no attachments or is not attached to compute: %s", volume.Name, instanceID)
	}
	startTime := time.Now()
	err = volumeattach.Delete(cClient, instanceID, volume.ID).ExtractErr()
	timeTaken := time.Since(startTime).Seconds()
	recordOpenstackOperationMetric("detach_disk", timeTaken, err)
	if err != nil {
		return fmt.Errorf("failed to delete volume %s from compute %s attached %v", volume.ID, instanceID, err)
	}
	klog.V(2).Infof("Successfully detached volume: %s from compute: %s", volume.ID, instanceID)
	return nil
}
func (os *OpenStack) ExpandVolume(volumeID string, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return oldSize, err
	}
	if volume.Status != volumeAvailableStatus {
		return oldSize, fmt.Errorf("volume status is not available")
	}
	volSizeGiB, err := volumeutil.RoundUpToGiBInt(newSize)
	if err != nil {
		return oldSize, err
	}
	newSizeQuant := resource.MustParse(fmt.Sprintf("%dGi", volSizeGiB))
	if volume.Size >= volSizeGiB {
		return newSizeQuant, nil
	}
	volumes, err := os.volumeService("")
	if err != nil {
		return oldSize, err
	}
	err = volumes.expandVolume(volumeID, volSizeGiB)
	if err != nil {
		return oldSize, err
	}
	return newSizeQuant, nil
}
func (os *OpenStack) getVolume(volumeID string) (Volume, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumes, err := os.volumeService("")
	if err != nil {
		return Volume{}, fmt.Errorf("unable to initialize cinder client for region: %s, err: %v", os.region, err)
	}
	return volumes.getVolume(volumeID)
}
func (os *OpenStack) CreateVolume(name string, size int, vtype, availability string, tags *map[string]string) (string, string, string, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volumes, err := os.volumeService("")
	if err != nil {
		return "", "", "", os.bsOpts.IgnoreVolumeAZ, fmt.Errorf("unable to initialize cinder client for region: %s, err: %v", os.region, err)
	}
	opts := volumeCreateOpts{Name: name, Size: size, VolumeType: vtype, Availability: availability}
	if tags != nil {
		opts.Metadata = *tags
	}
	volumeID, volumeAZ, err := volumes.createVolume(opts)
	if err != nil {
		return "", "", "", os.bsOpts.IgnoreVolumeAZ, fmt.Errorf("failed to create a %d GB volume: %v", size, err)
	}
	klog.Infof("Created volume %v in Availability Zone: %v Region: %v Ignore volume AZ: %v", volumeID, volumeAZ, os.region, os.bsOpts.IgnoreVolumeAZ)
	return volumeID, volumeAZ, os.region, os.bsOpts.IgnoreVolumeAZ, nil
}
func (os *OpenStack) GetDevicePathBySerialID(volumeID string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	candidateDeviceNodes := []string{fmt.Sprintf("virtio-%s", volumeID[:20]), fmt.Sprintf("scsi-0QEMU_QEMU_HARDDISK_%s", volumeID[:20]), fmt.Sprintf("wwn-0x%s", strings.Replace(volumeID, "-", "", -1))}
	files, _ := ioutil.ReadDir("/dev/disk/by-id/")
	for _, f := range files {
		for _, c := range candidateDeviceNodes {
			if c == f.Name() {
				klog.V(4).Infof("Found disk attached as %q; full devicepath: %s\n", f.Name(), path.Join("/dev/disk/by-id/", f.Name()))
				return path.Join("/dev/disk/by-id/", f.Name())
			}
		}
	}
	klog.V(4).Infof("Failed to find device for the volumeID: %q by serial ID", volumeID)
	return ""
}
func (os *OpenStack) getDevicePathFromInstanceMetadata(volumeID string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceMetadata, err := getMetadataFromMetadataService(newtonMetadataVersion)
	if err != nil {
		klog.V(4).Infof("Could not retrieve instance metadata. Error: %v", err)
		return ""
	}
	for _, device := range instanceMetadata.Devices {
		if device.Type == "disk" && device.Serial == volumeID {
			klog.V(4).Infof("Found disk metadata for volumeID %q. Bus: %q, Address: %q", volumeID, device.Bus, device.Address)
			diskPattern := fmt.Sprintf("/dev/disk/by-path/*-%s-%s", device.Bus, device.Address)
			diskPaths, err := filepath.Glob(diskPattern)
			if err != nil {
				klog.Errorf("could not retrieve disk path for volumeID: %q. Error filepath.Glob(%q): %v", volumeID, diskPattern, err)
				return ""
			}
			if len(diskPaths) == 1 {
				return diskPaths[0]
			}
			klog.Errorf("expecting to find one disk path for volumeID %q, found %d: %v", volumeID, len(diskPaths), diskPaths)
			return ""
		}
	}
	klog.V(4).Infof("Could not retrieve device metadata for volumeID: %q", volumeID)
	return ""
}
func (os *OpenStack) GetDevicePath(volumeID string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	devicePath := os.GetDevicePathBySerialID(volumeID)
	if devicePath == "" {
		devicePath = os.getDevicePathFromInstanceMetadata(volumeID)
	}
	if devicePath == "" {
		klog.Warningf("Failed to find device for the volumeID: %q", volumeID)
	}
	return devicePath
}
func (os *OpenStack) DeleteVolume(volumeID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	used, err := os.diskIsUsed(volumeID)
	if err != nil {
		return err
	}
	if used {
		msg := fmt.Sprintf("Cannot delete the volume %q, it's still attached to a node", volumeID)
		return k8s_volume.NewDeletedVolumeInUseError(msg)
	}
	volumes, err := os.volumeService("")
	if err != nil {
		return fmt.Errorf("unable to initialize cinder client for region: %s, err: %v", os.region, err)
	}
	err = volumes.deleteVolume(volumeID)
	return err
}
func (os *OpenStack) GetAttachmentDiskPath(instanceID, volumeID string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return "", err
	}
	if volume.Status != volumeInUseStatus {
		return "", fmt.Errorf("can not get device path of volume %s, its status is %s ", volume.Name, volume.Status)
	}
	if volume.AttachedServerID != "" {
		if instanceID == volume.AttachedServerID {
			return volume.AttachedDevice, nil
		}
		return "", fmt.Errorf("disk %q is attached to a different compute: %q, should be detached before proceeding", volumeID, volume.AttachedServerID)
	}
	return "", fmt.Errorf("volume %s has no ServerId", volumeID)
}
func (os *OpenStack) DiskIsAttached(instanceID, volumeID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if instanceID == "" {
		klog.Warningf("calling DiskIsAttached with empty instanceid: %s %s", instanceID, volumeID)
	}
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return false, err
	}
	return instanceID == volume.AttachedServerID, nil
}
func (os *OpenStack) DiskIsAttachedByName(nodeName types.NodeName, volumeID string) (bool, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cClient, err := os.NewComputeV2()
	if err != nil {
		return false, "", err
	}
	srv, err := getServerByName(cClient, nodeName)
	if err != nil {
		if err == ErrNotFound {
			return false, "", nil
		}
		return false, "", err
	}
	instanceID := "/" + srv.ID
	if ind := strings.LastIndex(instanceID, "/"); ind >= 0 {
		instanceID = instanceID[(ind + 1):]
	}
	attached, err := os.DiskIsAttached(instanceID, volumeID)
	return attached, instanceID, err
}
func (os *OpenStack) DisksAreAttached(instanceID string, volumeIDs []string) (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attached := make(map[string]bool)
	for _, volumeID := range volumeIDs {
		isAttached, err := os.DiskIsAttached(instanceID, volumeID)
		if err != nil && err != ErrNotFound {
			attached[volumeID] = true
			continue
		}
		attached[volumeID] = isAttached
	}
	return attached, nil
}
func (os *OpenStack) DisksAreAttachedByName(nodeName types.NodeName, volumeIDs []string) (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attached := make(map[string]bool)
	cClient, err := os.NewComputeV2()
	if err != nil {
		return attached, err
	}
	srv, err := getServerByName(cClient, nodeName)
	if err != nil {
		if err == ErrNotFound {
			for _, volumeID := range volumeIDs {
				attached[volumeID] = false
			}
			return attached, nil
		}
		return attached, err
	}
	instanceID := "/" + srv.ID
	if ind := strings.LastIndex(instanceID, "/"); ind >= 0 {
		instanceID = instanceID[(ind + 1):]
	}
	return os.DisksAreAttached(instanceID, volumeIDs)
}
func (os *OpenStack) diskIsUsed(volumeID string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	volume, err := os.getVolume(volumeID)
	if err != nil {
		return false, err
	}
	return volume.AttachedServerID != "", nil
}
func (os *OpenStack) ShouldTrustDevicePath() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return os.bsOpts.TrustDevicePath
}
func (os *OpenStack) GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.Cinder == nil {
		return nil, nil
	}
	if pv.Spec.Cinder.VolumeID == k8s_volume.ProvisionedVolumeName {
		return nil, nil
	}
	volume, err := os.getVolume(pv.Spec.Cinder.VolumeID)
	if err != nil {
		return nil, err
	}
	labels := make(map[string]string)
	labels[kubeletapis.LabelZoneFailureDomain] = volume.AvailabilityZone
	labels[kubeletapis.LabelZoneRegion] = os.region
	klog.V(4).Infof("The Volume %s has labels %v", pv.Spec.Cinder.VolumeID, labels)
	return labels, nil
}
func recordOpenstackOperationMetric(operation string, timeTaken float64, err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err != nil {
		openstackAPIRequestErrors.With(prometheus.Labels{"request": operation}).Inc()
	} else {
		openstackOperationsLatency.With(prometheus.Labels{"request": operation}).Observe(timeTaken)
	}
}
