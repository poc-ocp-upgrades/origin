package gce

import (
	"context"
	"encoding/json"
	"fmt"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce/cloud/meta"
	"k8s.io/kubernetes/pkg/features"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	"k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	"net/http"
	"strings"
)

type DiskType string

const (
	DiskTypeSSD                      = "pd-ssd"
	DiskTypeStandard                 = "pd-standard"
	diskTypeDefault                  = DiskTypeStandard
	diskTypeURITemplateSingleZone    = "%s/zones/%s/diskTypes/%s"
	diskTypeURITemplateRegional      = "%s/regions/%s/diskTypes/%s"
	diskTypePersistent               = "PERSISTENT"
	diskSourceURITemplateSingleZone  = "%s/zones/%s/disks/%s"
	diskSourceURITemplateRegional    = "%s/regions/%s/disks/%s"
	replicaZoneURITemplateSingleZone = "%s/zones/%s"
)

type diskServiceManager interface {
	CreateDiskOnCloudProvider(name string, sizeGb int64, tagsStr string, diskType string, zone string) error
	CreateRegionalDiskOnCloudProvider(name string, sizeGb int64, tagsStr string, diskType string, zones sets.String) error
	DeleteDiskOnCloudProvider(zone string, disk string) error
	DeleteRegionalDiskOnCloudProvider(diskName string) error
	AttachDiskOnCloudProvider(disk *Disk, readWrite string, instanceZone string, instanceName string) error
	DetachDiskOnCloudProvider(instanceZone string, instanceName string, devicePath string) error
	ResizeDiskOnCloudProvider(disk *Disk, sizeGb int64, zone string) error
	RegionalResizeDiskOnCloudProvider(disk *Disk, sizeGb int64) error
	GetDiskFromCloudProvider(zone string, diskName string) (*Disk, error)
	GetRegionalDiskFromCloudProvider(diskName string) (*Disk, error)
}
type gceServiceManager struct{ gce *Cloud }

var _ diskServiceManager = &gceServiceManager{}

func (manager *gceServiceManager) CreateDiskOnCloudProvider(name string, sizeGb int64, tagsStr string, diskType string, zone string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	diskTypeURI, err := manager.getDiskTypeURI(manager.gce.region, singleZone{zone}, diskType)
	if err != nil {
		return err
	}
	diskToCreateV1 := &compute.Disk{Name: name, SizeGb: sizeGb, Description: tagsStr, Type: diskTypeURI}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.Disks().Insert(ctx, meta.ZonalKey(name, zone), diskToCreateV1)
}
func (manager *gceServiceManager) CreateRegionalDiskOnCloudProvider(name string, sizeGb int64, tagsStr string, diskType string, replicaZones sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		return fmt.Errorf("the regional PD feature is only available with the %s Kubernetes feature gate enabled", features.GCERegionalPersistentDisk)
	}
	diskTypeURI, err := manager.getDiskTypeURI(manager.gce.region, multiZone{replicaZones}, diskType)
	if err != nil {
		return err
	}
	fullyQualifiedReplicaZones := []string{}
	for _, replicaZone := range replicaZones.UnsortedList() {
		fullyQualifiedReplicaZones = append(fullyQualifiedReplicaZones, manager.getReplicaZoneURI(replicaZone))
	}
	diskToCreate := &compute.Disk{Name: name, SizeGb: sizeGb, Description: tagsStr, Type: diskTypeURI, ReplicaZones: fullyQualifiedReplicaZones}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.RegionDisks().Insert(ctx, meta.RegionalKey(name, manager.gce.region), diskToCreate)
}
func (manager *gceServiceManager) AttachDiskOnCloudProvider(disk *Disk, readWrite string, instanceZone string, instanceName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	source, err := manager.getDiskSourceURI(disk)
	if err != nil {
		return err
	}
	attachedDiskV1 := &compute.AttachedDisk{DeviceName: disk.Name, Kind: disk.Kind, Mode: readWrite, Source: source, Type: diskTypePersistent}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.Instances().AttachDisk(ctx, meta.ZonalKey(instanceName, instanceZone), attachedDiskV1)
}
func (manager *gceServiceManager) DetachDiskOnCloudProvider(instanceZone string, instanceName string, devicePath string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.Instances().DetachDisk(ctx, meta.ZonalKey(instanceName, instanceZone), devicePath)
}
func (manager *gceServiceManager) GetDiskFromCloudProvider(zone string, diskName string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if zone == "" {
		return nil, fmt.Errorf("can not fetch disk %q, zone is empty", diskName)
	}
	if diskName == "" {
		return nil, fmt.Errorf("can not fetch disk, zone is specified (%q), but disk name is empty", zone)
	}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	diskStable, err := manager.gce.c.Disks().Get(ctx, meta.ZonalKey(diskName, zone))
	if err != nil {
		return nil, err
	}
	zoneInfo := singleZone{strings.TrimSpace(lastComponent(diskStable.Zone))}
	if zoneInfo.zone == "" {
		zoneInfo = singleZone{zone}
	}
	region, err := manager.getRegionFromZone(zoneInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to extract region from zone for %q/%q err=%v", zone, diskName, err)
	}
	return &Disk{ZoneInfo: zoneInfo, Region: region, Name: diskStable.Name, Kind: diskStable.Kind, Type: diskStable.Type, SizeGb: diskStable.SizeGb}, nil
}
func (manager *gceServiceManager) GetRegionalDiskFromCloudProvider(diskName string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		return nil, fmt.Errorf("the regional PD feature is only available with the %s Kubernetes feature gate enabled", features.GCERegionalPersistentDisk)
	}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	diskBeta, err := manager.gce.c.RegionDisks().Get(ctx, meta.RegionalKey(diskName, manager.gce.region))
	if err != nil {
		return nil, err
	}
	zones := sets.NewString()
	for _, zoneURI := range diskBeta.ReplicaZones {
		zones.Insert(lastComponent(zoneURI))
	}
	return &Disk{ZoneInfo: multiZone{zones}, Region: lastComponent(diskBeta.Region), Name: diskBeta.Name, Kind: diskBeta.Kind, Type: diskBeta.Type, SizeGb: diskBeta.SizeGb}, nil
}
func (manager *gceServiceManager) DeleteDiskOnCloudProvider(zone string, diskName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.Disks().Delete(ctx, meta.ZonalKey(diskName, zone))
}
func (manager *gceServiceManager) DeleteRegionalDiskOnCloudProvider(diskName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		return fmt.Errorf("the regional PD feature is only available with the %s Kubernetes feature gate enabled", features.GCERegionalPersistentDisk)
	}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.RegionDisks().Delete(ctx, meta.RegionalKey(diskName, manager.gce.region))
}
func (manager *gceServiceManager) getDiskSourceURI(disk *Disk) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getProjectsAPIEndpoint := manager.getProjectsAPIEndpoint()
	switch zoneInfo := disk.ZoneInfo.(type) {
	case singleZone:
		if zoneInfo.zone == "" || disk.Region == "" {
			return "", fmt.Errorf("PD does not have zone/region information: %#v", disk)
		}
		return getProjectsAPIEndpoint + fmt.Sprintf(diskSourceURITemplateSingleZone, manager.gce.projectID, zoneInfo.zone, disk.Name), nil
	case multiZone:
		if zoneInfo.replicaZones == nil || zoneInfo.replicaZones.Len() <= 0 {
			return "", fmt.Errorf("PD is regional but does not have any replicaZones specified: %v", disk)
		}
		return getProjectsAPIEndpoint + fmt.Sprintf(diskSourceURITemplateRegional, manager.gce.projectID, disk.Region, disk.Name), nil
	case nil:
		return "", fmt.Errorf("PD did not have ZoneInfo: %v", disk)
	default:
		return "", fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
	}
}
func (manager *gceServiceManager) getDiskTypeURI(diskRegion string, diskZoneInfo zoneType, diskType string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	getProjectsAPIEndpoint := manager.getProjectsAPIEndpoint()
	switch zoneInfo := diskZoneInfo.(type) {
	case singleZone:
		if zoneInfo.zone == "" {
			return "", fmt.Errorf("zone is empty: %v", zoneInfo)
		}
		return getProjectsAPIEndpoint + fmt.Sprintf(diskTypeURITemplateSingleZone, manager.gce.projectID, zoneInfo.zone, diskType), nil
	case multiZone:
		if zoneInfo.replicaZones == nil || zoneInfo.replicaZones.Len() <= 0 {
			return "", fmt.Errorf("zoneInfo is regional but does not have any replicaZones specified: %v", zoneInfo)
		}
		return getProjectsAPIEndpoint + fmt.Sprintf(diskTypeURITemplateRegional, manager.gce.projectID, diskRegion, diskType), nil
	case nil:
		return "", fmt.Errorf("zoneInfo nil")
	default:
		return "", fmt.Errorf("zoneInfo has unexpected type %T", zoneInfo)
	}
}
func (manager *gceServiceManager) getReplicaZoneURI(zone string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return manager.getProjectsAPIEndpoint() + fmt.Sprintf(replicaZoneURITemplateSingleZone, manager.gce.projectID, zone)
}
func (manager *gceServiceManager) getRegionFromZone(zoneInfo zoneType) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var zone string
	switch zoneInfo := zoneInfo.(type) {
	case singleZone:
		if zoneInfo.zone == "" {
			return "", fmt.Errorf("PD is single zone, but zone is not specified: %#v", zoneInfo)
		}
		zone = zoneInfo.zone
	case multiZone:
		if zoneInfo.replicaZones == nil || zoneInfo.replicaZones.Len() <= 0 {
			return "", fmt.Errorf("PD is regional but does not have any replicaZones specified: %v", zoneInfo)
		}
		zone = zoneInfo.replicaZones.UnsortedList()[0]
	case nil:
		return "", fmt.Errorf("zoneInfo is nil")
	default:
		return "", fmt.Errorf("zoneInfo has unexpected type %T", zoneInfo)
	}
	region, err := GetGCERegion(zone)
	if err != nil {
		klog.Warningf("failed to parse GCE region from zone %q: %v", zone, err)
		region = manager.gce.region
	}
	return region, nil
}
func (manager *gceServiceManager) ResizeDiskOnCloudProvider(disk *Disk, sizeGb int64, zone string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resizeServiceRequest := &compute.DisksResizeRequest{SizeGb: sizeGb}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.Disks().Resize(ctx, meta.ZonalKey(disk.Name, zone), resizeServiceRequest)
}
func (manager *gceServiceManager) RegionalResizeDiskOnCloudProvider(disk *Disk, sizeGb int64) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		return fmt.Errorf("the regional PD feature is only available with the %s Kubernetes feature gate enabled", features.GCERegionalPersistentDisk)
	}
	resizeServiceRequest := &compute.RegionDisksResizeRequest{SizeGb: sizeGb}
	ctx, cancel := cloud.ContextWithCallTimeout()
	defer cancel()
	return manager.gce.c.RegionDisks().Resize(ctx, meta.RegionalKey(disk.Name, disk.Region), resizeServiceRequest)
}

type Disks interface {
	AttachDisk(diskName string, nodeName types.NodeName, readOnly bool, regional bool) error
	DetachDisk(devicePath string, nodeName types.NodeName) error
	DiskIsAttached(diskName string, nodeName types.NodeName) (bool, error)
	DisksAreAttached(diskNames []string, nodeName types.NodeName) (map[string]bool, error)
	CreateDisk(name string, diskType string, zone string, sizeGb int64, tags map[string]string) error
	CreateRegionalDisk(name string, diskType string, replicaZones sets.String, sizeGb int64, tags map[string]string) error
	DeleteDisk(diskToDelete string) error
	ResizeDisk(diskToResize string, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error)
	GetAutoLabelsForPD(name string, zone string) (map[string]string, error)
}

var _ Disks = (*Cloud)(nil)
var _ cloudprovider.PVLabeler = (*Cloud)(nil)

type Disk struct {
	ZoneInfo zoneType
	Region   string
	Name     string
	Kind     string
	Type     string
	SizeGb   int64
}
type zoneType interface{ isZoneType() }
type multiZone struct{ replicaZones sets.String }
type singleZone struct{ zone string }

func (m multiZone) isZoneType() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (s singleZone) isZoneType() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func newDiskMetricContextZonal(request, region, zone string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("disk", request, region, zone, computeV1Version)
}
func newDiskMetricContextRegional(request, region string) *metricContext {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return newGenericMetricContext("disk", request, region, unusedMetricLabel, computeV1Version)
}
func (g *Cloud) GetLabelsForVolume(ctx context.Context, pv *v1.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if pv.Spec.GCEPersistentDisk.PDName == volume.ProvisionedVolumeName {
		return nil, nil
	}
	zone := pv.Labels[kubeletapis.LabelZoneFailureDomain]
	labels, err := g.GetAutoLabelsForPD(pv.Spec.GCEPersistentDisk.PDName, zone)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
func (g *Cloud) AttachDisk(diskName string, nodeName types.NodeName, readOnly bool, regional bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceName := mapNodeNameToInstanceName(nodeName)
	instance, err := g.getInstanceByName(instanceName)
	if err != nil {
		return fmt.Errorf("error getting instance %q", instanceName)
	}
	var disk *Disk
	var mc *metricContext
	if regional && utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		disk, err = g.getRegionalDiskByName(diskName)
		if err != nil {
			return err
		}
		mc = newDiskMetricContextRegional("attach", g.region)
	} else {
		disk, err = g.getDiskByName(diskName, instance.Zone)
		if err != nil {
			return err
		}
		mc = newDiskMetricContextZonal("attach", g.region, instance.Zone)
	}
	readWrite := "READ_WRITE"
	if readOnly {
		readWrite = "READ_ONLY"
	}
	return mc.Observe(g.manager.AttachDiskOnCloudProvider(disk, readWrite, instance.Zone, instance.Name))
}
func (g *Cloud) DetachDisk(devicePath string, nodeName types.NodeName) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceName := mapNodeNameToInstanceName(nodeName)
	inst, err := g.getInstanceByName(instanceName)
	if err != nil {
		if err == cloudprovider.InstanceNotFound {
			klog.Warningf("Instance %q does not exist. DetachDisk will assume PD %q is not attached to it.", instanceName, devicePath)
			return nil
		}
		return fmt.Errorf("error getting instance %q", instanceName)
	}
	mc := newDiskMetricContextZonal("detach", g.region, inst.Zone)
	return mc.Observe(g.manager.DetachDiskOnCloudProvider(inst.Zone, inst.Name, devicePath))
}
func (g *Cloud) DiskIsAttached(diskName string, nodeName types.NodeName) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	instanceName := mapNodeNameToInstanceName(nodeName)
	instance, err := g.getInstanceByName(instanceName)
	if err != nil {
		if err == cloudprovider.InstanceNotFound {
			klog.Warningf("Instance %q does not exist. DiskIsAttached will assume PD %q is not attached to it.", instanceName, diskName)
			return false, nil
		}
		return false, err
	}
	for _, disk := range instance.Disks {
		if disk.DeviceName == diskName {
			return true, nil
		}
	}
	return false, nil
}
func (g *Cloud) DisksAreAttached(diskNames []string, nodeName types.NodeName) (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	attached := make(map[string]bool)
	for _, diskName := range diskNames {
		attached[diskName] = false
	}
	instanceName := mapNodeNameToInstanceName(nodeName)
	instance, err := g.getInstanceByName(instanceName)
	if err != nil {
		if err == cloudprovider.InstanceNotFound {
			klog.Warningf("Instance %q does not exist. DisksAreAttached will assume PD %v are not attached to it.", instanceName, diskNames)
			return attached, nil
		}
		return attached, err
	}
	for _, instanceDisk := range instance.Disks {
		for _, diskName := range diskNames {
			if instanceDisk.DeviceName == diskName {
				attached[diskName] = true
			}
		}
	}
	return attached, nil
}
func (g *Cloud) CreateDisk(name string, diskType string, zone string, sizeGb int64, tags map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	curZones, err := g.GetAllCurrentZones()
	if err != nil {
		return err
	}
	if !curZones.Has(zone) {
		return fmt.Errorf("kubernetes does not have a node in zone %q", zone)
	}
	tagsStr, err := g.encodeDiskTags(tags)
	if err != nil {
		return err
	}
	diskType, err = getDiskType(diskType)
	if err != nil {
		return err
	}
	mc := newDiskMetricContextZonal("create", g.region, zone)
	err = g.manager.CreateDiskOnCloudProvider(name, sizeGb, tagsStr, diskType, zone)
	mc.Observe(err)
	if isGCEError(err, "alreadyExists") {
		klog.Warningf("GCE PD %q already exists, reusing", name)
		return nil
	}
	return err
}
func (g *Cloud) CreateRegionalDisk(name string, diskType string, replicaZones sets.String, sizeGb int64, tags map[string]string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	curZones, err := g.GetAllCurrentZones()
	if err != nil {
		return err
	}
	if !curZones.IsSuperset(replicaZones) {
		return fmt.Errorf("kubernetes does not have nodes in specified zones: %q. Zones that contain nodes: %q", replicaZones.Difference(curZones), curZones)
	}
	tagsStr, err := g.encodeDiskTags(tags)
	if err != nil {
		return err
	}
	diskType, err = getDiskType(diskType)
	if err != nil {
		return err
	}
	mc := newDiskMetricContextRegional("create", g.region)
	err = g.manager.CreateRegionalDiskOnCloudProvider(name, sizeGb, tagsStr, diskType, replicaZones)
	mc.Observe(err)
	if isGCEError(err, "alreadyExists") {
		klog.Warningf("GCE PD %q already exists, reusing", name)
		return nil
	}
	return err
}
func getDiskType(diskType string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch diskType {
	case DiskTypeSSD, DiskTypeStandard:
		return diskType, nil
	case "":
		return diskTypeDefault, nil
	default:
		return "", fmt.Errorf("invalid GCE disk type %q", diskType)
	}
}
func (g *Cloud) DeleteDisk(diskToDelete string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	err := g.doDeleteDisk(diskToDelete)
	if isGCEError(err, "resourceInUseByAnotherResource") {
		return volume.NewDeletedVolumeInUseError(err.Error())
	}
	if err == cloudprovider.DiskNotFound {
		return nil
	}
	return err
}
func (g *Cloud) ResizeDisk(diskToResize string, oldSize resource.Quantity, newSize resource.Quantity) (resource.Quantity, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := g.GetDiskByNameUnknownZone(diskToResize)
	if err != nil {
		return oldSize, err
	}
	requestGIB := volumeutil.RoundUpToGiB(newSize)
	newSizeQuant := resource.MustParse(fmt.Sprintf("%dGi", requestGIB))
	if disk.SizeGb >= requestGIB {
		return newSizeQuant, nil
	}
	var mc *metricContext
	switch zoneInfo := disk.ZoneInfo.(type) {
	case singleZone:
		mc = newDiskMetricContextZonal("resize", disk.Region, zoneInfo.zone)
		err := g.manager.ResizeDiskOnCloudProvider(disk, requestGIB, zoneInfo.zone)
		if err != nil {
			return oldSize, mc.Observe(err)
		}
		return newSizeQuant, mc.Observe(err)
	case multiZone:
		if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
			return oldSize, fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
		}
		mc = newDiskMetricContextRegional("resize", disk.Region)
		err := g.manager.RegionalResizeDiskOnCloudProvider(disk, requestGIB)
		if err != nil {
			return oldSize, mc.Observe(err)
		}
		return newSizeQuant, mc.Observe(err)
	case nil:
		return oldSize, fmt.Errorf("PD has nil ZoneInfo: %v", disk)
	default:
		return oldSize, fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
	}
}
func (g *Cloud) GetAutoLabelsForPD(name string, zone string) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var disk *Disk
	var err error
	if zone == "" {
		disk, err = g.GetDiskByNameUnknownZone(name)
		if err != nil {
			return nil, err
		}
	} else {
		if utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
			zoneSet, err := volumeutil.LabelZonesToSet(zone)
			if err != nil {
				klog.Warningf("Failed to parse zone field: %q. Will use raw field.", zone)
			}
			if len(zoneSet) > 1 {
				disk, err = g.getRegionalDiskByName(name)
				if err != nil {
					return nil, err
				}
			} else {
				disk, err = g.getDiskByName(name, zone)
				if err != nil {
					return nil, err
				}
			}
		} else {
			disk, err = g.getDiskByName(name, zone)
			if err != nil {
				return nil, err
			}
		}
	}
	labels := make(map[string]string)
	switch zoneInfo := disk.ZoneInfo.(type) {
	case singleZone:
		if zoneInfo.zone == "" || disk.Region == "" {
			return nil, fmt.Errorf("PD did not have zone/region information: %v", disk)
		}
		labels[kubeletapis.LabelZoneFailureDomain] = zoneInfo.zone
		labels[kubeletapis.LabelZoneRegion] = disk.Region
	case multiZone:
		if zoneInfo.replicaZones == nil || zoneInfo.replicaZones.Len() <= 0 {
			return nil, fmt.Errorf("PD is regional but does not have any replicaZones specified: %v", disk)
		}
		labels[kubeletapis.LabelZoneFailureDomain] = volumeutil.ZonesSetToLabelValue(zoneInfo.replicaZones)
		labels[kubeletapis.LabelZoneRegion] = disk.Region
	case nil:
		return nil, fmt.Errorf("PD did not have ZoneInfo: %v", disk)
	default:
		return nil, fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
	}
	return labels, nil
}
func (g *Cloud) findDiskByName(diskName string, zone string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mc := newDiskMetricContextZonal("get", g.region, zone)
	disk, err := g.manager.GetDiskFromCloudProvider(zone, diskName)
	if err == nil {
		return disk, mc.Observe(nil)
	}
	if !isHTTPErrorCode(err, http.StatusNotFound) {
		return nil, mc.Observe(err)
	}
	return nil, mc.Observe(nil)
}
func (g *Cloud) getDiskByName(diskName string, zone string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := g.findDiskByName(diskName, zone)
	if disk == nil && err == nil {
		return nil, fmt.Errorf("GCE persistent disk not found: diskName=%q zone=%q", diskName, zone)
	}
	return disk, err
}
func (g *Cloud) findRegionalDiskByName(diskName string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mc := newDiskMetricContextRegional("get", g.region)
	disk, err := g.manager.GetRegionalDiskFromCloudProvider(diskName)
	if err == nil {
		return disk, mc.Observe(nil)
	}
	if !isHTTPErrorCode(err, http.StatusNotFound) {
		return nil, mc.Observe(err)
	}
	return nil, mc.Observe(nil)
}
func (g *Cloud) getRegionalDiskByName(diskName string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := g.findRegionalDiskByName(diskName)
	if disk == nil && err == nil {
		return nil, fmt.Errorf("GCE regional persistent disk not found: diskName=%q", diskName)
	}
	return disk, err
}
func (g *Cloud) GetDiskByNameUnknownZone(diskName string) (*Disk, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
		regionalDisk, err := g.getRegionalDiskByName(diskName)
		if err == nil {
			return regionalDisk, err
		}
	}
	var found *Disk
	for _, zone := range g.managedZones {
		disk, err := g.findDiskByName(diskName, zone)
		if err != nil {
			return nil, err
		}
		if disk == nil {
			continue
		}
		if found != nil {
			switch zoneInfo := disk.ZoneInfo.(type) {
			case multiZone:
				if zoneInfo.replicaZones.Has(zone) {
					klog.Warningf("GCE PD name (%q) was found in multiple zones (%q), but ok because it is a RegionalDisk.", diskName, zoneInfo.replicaZones)
					continue
				}
				return nil, fmt.Errorf("GCE PD name was found in multiple zones: %q", diskName)
			default:
				return nil, fmt.Errorf("GCE PD name was found in multiple zones: %q", diskName)
			}
		}
		found = disk
	}
	if found != nil {
		return found, nil
	}
	klog.Warningf("GCE persistent disk %q not found in managed zones (%s)", diskName, strings.Join(g.managedZones, ","))
	return nil, cloudprovider.DiskNotFound
}
func (g *Cloud) encodeDiskTags(tags map[string]string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(tags) == 0 {
		return "", nil
	}
	enc, err := json.Marshal(tags)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
func (g *Cloud) doDeleteDisk(diskToDelete string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	disk, err := g.GetDiskByNameUnknownZone(diskToDelete)
	if err != nil {
		return err
	}
	var mc *metricContext
	switch zoneInfo := disk.ZoneInfo.(type) {
	case singleZone:
		mc = newDiskMetricContextZonal("delete", disk.Region, zoneInfo.zone)
		return mc.Observe(g.manager.DeleteDiskOnCloudProvider(zoneInfo.zone, disk.Name))
	case multiZone:
		if !utilfeature.DefaultFeatureGate.Enabled(features.GCERegionalPersistentDisk) {
			return fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
		}
		mc = newDiskMetricContextRegional("delete", disk.Region)
		return mc.Observe(g.manager.DeleteRegionalDiskOnCloudProvider(disk.Name))
	case nil:
		return fmt.Errorf("PD has nil ZoneInfo: %v", disk)
	default:
		return fmt.Errorf("disk.ZoneInfo has unexpected type %T", zoneInfo)
	}
}
func isGCEError(err error, reason string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	apiErr, ok := err.(*googleapi.Error)
	if !ok {
		return false
	}
	for _, e := range apiErr.Errors {
		if e.Reason == reason {
			return true
		}
	}
	return false
}
