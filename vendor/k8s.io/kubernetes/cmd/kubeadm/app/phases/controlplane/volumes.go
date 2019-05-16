package controlplane

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	staticpodutil "k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	"os"
	"path/filepath"
	"strings"
)

const (
	caCertsVolumeName       = "ca-certs"
	caCertsVolumePath       = "/etc/ssl/certs"
	flexvolumeDirVolumeName = "flexvolume-dir"
	flexvolumeDirVolumePath = "/usr/libexec/kubernetes/kubelet-plugins/volume/exec"
)

var caCertsExtraVolumePaths = []string{"/etc/pki", "/usr/share/ca-certificates", "/usr/local/share/ca-certificates", "/etc/ca-certificates"}

func getHostPathVolumesForTheControlPlane(cfg *kubeadmapi.InitConfiguration) controlPlaneHostPathMounts {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	hostPathDirectoryOrCreate := v1.HostPathDirectoryOrCreate
	hostPathFileOrCreate := v1.HostPathFileOrCreate
	mounts := newControlPlaneHostPathMounts()
	mounts.NewHostPathMount(kubeadmconstants.KubeAPIServer, kubeadmconstants.KubeCertificatesVolumeName, cfg.CertificatesDir, cfg.CertificatesDir, true, &hostPathDirectoryOrCreate)
	mounts.NewHostPathMount(kubeadmconstants.KubeAPIServer, caCertsVolumeName, caCertsVolumePath, caCertsVolumePath, true, &hostPathDirectoryOrCreate)
	if cfg.Etcd.External != nil {
		etcdVols, etcdVolMounts := getEtcdCertVolumes(cfg.Etcd.External, cfg.CertificatesDir)
		mounts.AddHostPathMounts(kubeadmconstants.KubeAPIServer, etcdVols, etcdVolMounts)
	}
	mounts.NewHostPathMount(kubeadmconstants.KubeControllerManager, kubeadmconstants.KubeCertificatesVolumeName, cfg.CertificatesDir, cfg.CertificatesDir, true, &hostPathDirectoryOrCreate)
	mounts.NewHostPathMount(kubeadmconstants.KubeControllerManager, caCertsVolumeName, caCertsVolumePath, caCertsVolumePath, true, &hostPathDirectoryOrCreate)
	controllerManagerKubeConfigFile := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ControllerManagerKubeConfigFileName)
	mounts.NewHostPathMount(kubeadmconstants.KubeControllerManager, kubeadmconstants.KubeConfigVolumeName, controllerManagerKubeConfigFile, controllerManagerKubeConfigFile, true, &hostPathFileOrCreate)
	if stat, err := os.Stat(flexvolumeDirVolumePath); err == nil && stat.IsDir() {
		mounts.NewHostPathMount(kubeadmconstants.KubeControllerManager, flexvolumeDirVolumeName, flexvolumeDirVolumePath, flexvolumeDirVolumePath, false, &hostPathDirectoryOrCreate)
	}
	schedulerKubeConfigFile := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.SchedulerKubeConfigFileName)
	mounts.NewHostPathMount(kubeadmconstants.KubeScheduler, kubeadmconstants.KubeConfigVolumeName, schedulerKubeConfigFile, schedulerKubeConfigFile, true, &hostPathFileOrCreate)
	for _, caCertsExtraVolumePath := range caCertsExtraVolumePaths {
		if isExtraVolumeMountNeeded(caCertsExtraVolumePath) {
			caCertsExtraVolumeName := strings.Replace(caCertsExtraVolumePath, "/", "-", -1)[1:]
			mounts.NewHostPathMount(kubeadmconstants.KubeAPIServer, caCertsExtraVolumeName, caCertsExtraVolumePath, caCertsExtraVolumePath, true, &hostPathDirectoryOrCreate)
			mounts.NewHostPathMount(kubeadmconstants.KubeControllerManager, caCertsExtraVolumeName, caCertsExtraVolumePath, caCertsExtraVolumePath, true, &hostPathDirectoryOrCreate)
		}
	}
	mounts.AddExtraHostPathMounts(kubeadmconstants.KubeAPIServer, cfg.APIServer.ExtraVolumes)
	mounts.AddExtraHostPathMounts(kubeadmconstants.KubeControllerManager, cfg.ControllerManager.ExtraVolumes)
	mounts.AddExtraHostPathMounts(kubeadmconstants.KubeScheduler, cfg.Scheduler.ExtraVolumes)
	return mounts
}

type controlPlaneHostPathMounts struct {
	volumes      map[string]map[string]v1.Volume
	volumeMounts map[string]map[string]v1.VolumeMount
}

func newControlPlaneHostPathMounts() controlPlaneHostPathMounts {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return controlPlaneHostPathMounts{volumes: map[string]map[string]v1.Volume{}, volumeMounts: map[string]map[string]v1.VolumeMount{}}
}
func (c *controlPlaneHostPathMounts) NewHostPathMount(component, mountName, hostPath, containerPath string, readOnly bool, hostPathType *v1.HostPathType) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vol := staticpodutil.NewVolume(mountName, hostPath, hostPathType)
	c.addComponentVolume(component, vol)
	volMount := staticpodutil.NewVolumeMount(mountName, containerPath, readOnly)
	c.addComponentVolumeMount(component, volMount)
}
func (c *controlPlaneHostPathMounts) AddHostPathMounts(component string, vols []v1.Volume, volMounts []v1.VolumeMount) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, v := range vols {
		c.addComponentVolume(component, v)
	}
	for _, v := range volMounts {
		c.addComponentVolumeMount(component, v)
	}
}
func (c *controlPlaneHostPathMounts) AddExtraHostPathMounts(component string, extraVols []kubeadmapi.HostPathMount) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, extraVol := range extraVols {
		fmt.Printf("[controlplane] Adding extra host path mount %q to %q\n", extraVol.Name, component)
		hostPathType := extraVol.PathType
		c.NewHostPathMount(component, extraVol.Name, extraVol.HostPath, extraVol.MountPath, extraVol.ReadOnly, &hostPathType)
	}
}
func (c *controlPlaneHostPathMounts) GetVolumes(component string) map[string]v1.Volume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.volumes[component]
}
func (c *controlPlaneHostPathMounts) GetVolumeMounts(component string) map[string]v1.VolumeMount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.volumeMounts[component]
}
func (c *controlPlaneHostPathMounts) addComponentVolume(component string, vol v1.Volume) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, ok := c.volumes[component]; !ok {
		c.volumes[component] = map[string]v1.Volume{}
	}
	c.volumes[component][vol.Name] = vol
}
func (c *controlPlaneHostPathMounts) addComponentVolumeMount(component string, volMount v1.VolumeMount) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, ok := c.volumeMounts[component]; !ok {
		c.volumeMounts[component] = map[string]v1.VolumeMount{}
	}
	c.volumeMounts[component][volMount.Name] = volMount
}
func getEtcdCertVolumes(etcdCfg *kubeadmapi.ExternalEtcd, k8sCertificatesDir string) ([]v1.Volume, []v1.VolumeMount) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	certPaths := []string{etcdCfg.CAFile, etcdCfg.CertFile, etcdCfg.KeyFile}
	certDirs := sets.NewString()
	for _, certPath := range certPaths {
		certDir := filepath.Dir(certPath)
		extraVolumePath := false
		for _, caCertsExtraVolumePath := range caCertsExtraVolumePaths {
			if strings.HasPrefix(certDir, caCertsExtraVolumePath) {
				extraVolumePath = true
				break
			}
		}
		if certDir == "." || extraVolumePath || strings.HasPrefix(certDir, caCertsVolumePath) || strings.HasPrefix(certDir, k8sCertificatesDir) {
			continue
		}
		alreadyExists := false
		for _, existingCertDir := range certDirs.List() {
			if strings.HasPrefix(existingCertDir, certDir) {
				certDirs.Delete(existingCertDir)
			} else if strings.HasPrefix(certDir, existingCertDir) {
				alreadyExists = true
			}
		}
		if alreadyExists {
			continue
		}
		certDirs.Insert(certDir)
	}
	volumes := []v1.Volume{}
	volumeMounts := []v1.VolumeMount{}
	pathType := v1.HostPathDirectoryOrCreate
	for i, certDir := range certDirs.List() {
		name := fmt.Sprintf("etcd-certs-%d", i)
		volumes = append(volumes, staticpodutil.NewVolume(name, certDir, &pathType))
		volumeMounts = append(volumeMounts, staticpodutil.NewVolumeMount(name, certDir, true))
	}
	return volumes, volumeMounts
}
func isExtraVolumeMountNeeded(caCertsExtraVolumePath string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := os.Stat(caCertsExtraVolumePath); err == nil {
		return true
	}
	return false
}
