package label

import (
	"bytes"
	"fmt"
	goformat "fmt"
	"io"
	"k8s.io/apiserver/pkg/admission"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/aws"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/azure"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	"k8s.io/kubernetes/pkg/features"
	kubeapiserveradmission "k8s.io/kubernetes/pkg/kubeapiserver/admission"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
	vol "k8s.io/kubernetes/pkg/volume"
	volumeutil "k8s.io/kubernetes/pkg/volume/util"
	goos "os"
	godefaultruntime "runtime"
	"sync"
	gotime "time"
)

const (
	PluginName = "PersistentVolumeLabel"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		persistentVolumeLabelAdmission := newPersistentVolumeLabel()
		return persistentVolumeLabelAdmission, nil
	})
}

var _ = admission.Interface(&persistentVolumeLabel{})

type persistentVolumeLabel struct {
	*admission.Handler
	mutex            sync.Mutex
	ebsVolumes       aws.Volumes
	cloudConfig      []byte
	gceCloudProvider *gce.Cloud
	azureProvider    *azure.Cloud
}

var _ admission.MutationInterface = &persistentVolumeLabel{}
var _ kubeapiserveradmission.WantsCloudConfig = &persistentVolumeLabel{}

func newPersistentVolumeLabel() *persistentVolumeLabel {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.Warning("PersistentVolumeLabel admission controller is deprecated. " + "Please remove this controller from your configuration files and scripts.")
	return &persistentVolumeLabel{Handler: admission.NewHandler(admission.Create)}
}
func (l *persistentVolumeLabel) SetCloudConfig(cloudConfig []byte) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.cloudConfig = cloudConfig
}
func nodeSelectorRequirementKeysExistInNodeSelectorTerms(reqs []api.NodeSelectorRequirement, terms []api.NodeSelectorTerm) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, req := range reqs {
		for _, term := range terms {
			for _, r := range term.MatchExpressions {
				if r.Key == req.Key {
					return true
				}
			}
		}
	}
	return false
}
func (l *persistentVolumeLabel) Admit(a admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.GetResource().GroupResource() != api.Resource("persistentvolumes") {
		return nil
	}
	obj := a.GetObject()
	if obj == nil {
		return nil
	}
	volume, ok := obj.(*api.PersistentVolume)
	if !ok {
		return nil
	}
	var volumeLabels map[string]string
	if volume.Spec.AWSElasticBlockStore != nil {
		labels, err := l.findAWSEBSLabels(volume)
		if err != nil {
			return admission.NewForbidden(a, fmt.Errorf("error querying AWS EBS volume %s: %v", volume.Spec.AWSElasticBlockStore.VolumeID, err))
		}
		volumeLabels = labels
	}
	if volume.Spec.GCEPersistentDisk != nil {
		labels, err := l.findGCEPDLabels(volume)
		if err != nil {
			return admission.NewForbidden(a, fmt.Errorf("error querying GCE PD volume %s: %v", volume.Spec.GCEPersistentDisk.PDName, err))
		}
		volumeLabels = labels
	}
	if volume.Spec.AzureDisk != nil {
		labels, err := l.findAzureDiskLabels(volume)
		if err != nil {
			return admission.NewForbidden(a, fmt.Errorf("error querying AzureDisk volume %s: %v", volume.Spec.AzureDisk.DiskName, err))
		}
		volumeLabels = labels
	}
	requirements := make([]api.NodeSelectorRequirement, 0)
	if len(volumeLabels) != 0 {
		if volume.Labels == nil {
			volume.Labels = make(map[string]string)
		}
		for k, v := range volumeLabels {
			volume.Labels[k] = v
			var values []string
			if k == kubeletapis.LabelZoneFailureDomain {
				zones, err := volumeutil.LabelZonesToSet(v)
				if err != nil {
					return admission.NewForbidden(a, fmt.Errorf("failed to convert label string for Zone: %s to a Set", v))
				}
				values = zones.UnsortedList()
			} else {
				values = []string{v}
			}
			requirements = append(requirements, api.NodeSelectorRequirement{Key: k, Operator: api.NodeSelectorOpIn, Values: values})
		}
		if utilfeature.DefaultFeatureGate.Enabled(features.VolumeScheduling) {
			if volume.Spec.NodeAffinity == nil {
				volume.Spec.NodeAffinity = new(api.VolumeNodeAffinity)
			}
			if volume.Spec.NodeAffinity.Required == nil {
				volume.Spec.NodeAffinity.Required = new(api.NodeSelector)
			}
			if len(volume.Spec.NodeAffinity.Required.NodeSelectorTerms) == 0 {
				volume.Spec.NodeAffinity.Required.NodeSelectorTerms = make([]api.NodeSelectorTerm, 1)
			}
			if nodeSelectorRequirementKeysExistInNodeSelectorTerms(requirements, volume.Spec.NodeAffinity.Required.NodeSelectorTerms) {
				klog.V(4).Infof("NodeSelectorRequirements for cloud labels %v conflict with existing NodeAffinity %v. Skipping addition of NodeSelectorRequirements for cloud labels.", requirements, volume.Spec.NodeAffinity)
			} else {
				for _, req := range requirements {
					for i := range volume.Spec.NodeAffinity.Required.NodeSelectorTerms {
						volume.Spec.NodeAffinity.Required.NodeSelectorTerms[i].MatchExpressions = append(volume.Spec.NodeAffinity.Required.NodeSelectorTerms[i].MatchExpressions, req)
					}
				}
			}
		}
	}
	return nil
}
func (l *persistentVolumeLabel) findAWSEBSLabels(volume *api.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volume.Spec.AWSElasticBlockStore.VolumeID == vol.ProvisionedVolumeName {
		return nil, nil
	}
	ebsVolumes, err := l.getEBSVolumes()
	if err != nil {
		return nil, err
	}
	if ebsVolumes == nil {
		return nil, fmt.Errorf("unable to build AWS cloud provider for EBS")
	}
	spec := aws.KubernetesVolumeID(volume.Spec.AWSElasticBlockStore.VolumeID)
	labels, err := ebsVolumes.GetVolumeLabels(spec)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
func (l *persistentVolumeLabel) getEBSVolumes() (aws.Volumes, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.ebsVolumes == nil {
		var cloudConfigReader io.Reader
		if len(l.cloudConfig) > 0 {
			cloudConfigReader = bytes.NewReader(l.cloudConfig)
		}
		cloudProvider, err := cloudprovider.GetCloudProvider("aws", cloudConfigReader)
		if err != nil || cloudProvider == nil {
			return nil, err
		}
		awsCloudProvider, ok := cloudProvider.(*aws.Cloud)
		if !ok {
			return nil, fmt.Errorf("error retrieving AWS cloud provider")
		}
		l.ebsVolumes = awsCloudProvider
	}
	return l.ebsVolumes, nil
}
func (l *persistentVolumeLabel) findGCEPDLabels(volume *api.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volume.Spec.GCEPersistentDisk.PDName == vol.ProvisionedVolumeName {
		return nil, nil
	}
	provider, err := l.getGCECloudProvider()
	if err != nil {
		return nil, err
	}
	if provider == nil {
		return nil, fmt.Errorf("unable to build GCE cloud provider for PD")
	}
	zone := volume.Labels[kubeletapis.LabelZoneFailureDomain]
	labels, err := provider.GetAutoLabelsForPD(volume.Spec.GCEPersistentDisk.PDName, zone)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
func (l *persistentVolumeLabel) getGCECloudProvider() (*gce.Cloud, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.gceCloudProvider == nil {
		var cloudConfigReader io.Reader
		if len(l.cloudConfig) > 0 {
			cloudConfigReader = bytes.NewReader(l.cloudConfig)
		}
		cloudProvider, err := cloudprovider.GetCloudProvider("gce", cloudConfigReader)
		if err != nil || cloudProvider == nil {
			return nil, err
		}
		gceCloudProvider, ok := cloudProvider.(*gce.Cloud)
		if !ok {
			return nil, fmt.Errorf("error retrieving GCE cloud provider")
		}
		l.gceCloudProvider = gceCloudProvider
	}
	return l.gceCloudProvider, nil
}
func (l *persistentVolumeLabel) getAzureCloudProvider() (*azure.Cloud, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.azureProvider == nil {
		var cloudConfigReader io.Reader
		if len(l.cloudConfig) > 0 {
			cloudConfigReader = bytes.NewReader(l.cloudConfig)
		}
		cloudProvider, err := cloudprovider.GetCloudProvider("azure", cloudConfigReader)
		if err != nil || cloudProvider == nil {
			return nil, err
		}
		azureProvider, ok := cloudProvider.(*azure.Cloud)
		if !ok {
			return nil, fmt.Errorf("error retrieving Azure cloud provider")
		}
		l.azureProvider = azureProvider
	}
	return l.azureProvider, nil
}
func (l *persistentVolumeLabel) findAzureDiskLabels(volume *api.PersistentVolume) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if volume.Spec.AzureDisk.DiskName == vol.ProvisionedVolumeName {
		return nil, nil
	}
	provider, err := l.getAzureCloudProvider()
	if err != nil {
		return nil, err
	}
	if provider == nil {
		return nil, fmt.Errorf("unable to build Azure cloud provider for AzureDisk")
	}
	return provider.GetAzureDiskLabels(volume.Spec.AzureDisk.DataDiskURI)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
