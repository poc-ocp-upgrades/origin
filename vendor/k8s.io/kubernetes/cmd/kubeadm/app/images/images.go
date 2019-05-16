package images

import (
	"fmt"
	goformat "fmt"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func GetGenericImage(prefix, image, tag string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s:%s", prefix, image, tag)
}
func GetKubernetesImage(image string, cfg *kubeadmapi.ClusterConfiguration) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.UseHyperKubeImage {
		image = constants.HyperKube
	}
	repoPrefix := cfg.GetControlPlaneImageRepository()
	kubernetesImageTag := kubeadmutil.KubernetesVersionToImageTag(cfg.KubernetesVersion)
	return GetGenericImage(repoPrefix, image, kubernetesImageTag)
}
func GetDNSImage(cfg *kubeadmapi.ClusterConfiguration, imageName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dnsImageRepository := cfg.ImageRepository
	if cfg.DNS.ImageRepository != "" {
		dnsImageRepository = cfg.DNS.ImageRepository
	}
	dnsImageTag := constants.GetDNSVersion(cfg.DNS.Type)
	if cfg.DNS.ImageTag != "" {
		dnsImageTag = cfg.DNS.ImageTag
	}
	return GetGenericImage(dnsImageRepository, imageName, dnsImageTag)
}
func GetEtcdImage(cfg *kubeadmapi.ClusterConfiguration) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdImageRepository := cfg.ImageRepository
	if cfg.Etcd.Local != nil && cfg.Etcd.Local.ImageRepository != "" {
		etcdImageRepository = cfg.Etcd.Local.ImageRepository
	}
	etcdImageTag := constants.DefaultEtcdVersion
	etcdVersion, err := constants.EtcdSupportedVersion(cfg.KubernetesVersion)
	if err == nil {
		etcdImageTag = etcdVersion.String()
	}
	if cfg.Etcd.Local != nil && cfg.Etcd.Local.ImageTag != "" {
		etcdImageTag = cfg.Etcd.Local.ImageTag
	}
	return GetGenericImage(etcdImageRepository, constants.Etcd, etcdImageTag)
}
func GetPauseImage(cfg *kubeadmapi.ClusterConfiguration) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return GetGenericImage(cfg.ImageRepository, "pause", constants.PauseVersion)
}
func GetAllImages(cfg *kubeadmapi.ClusterConfiguration) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	imgs := []string{}
	if cfg.UseHyperKubeImage {
		imgs = append(imgs, GetKubernetesImage(constants.HyperKube, cfg))
	} else {
		imgs = append(imgs, GetKubernetesImage(constants.KubeAPIServer, cfg))
		imgs = append(imgs, GetKubernetesImage(constants.KubeControllerManager, cfg))
		imgs = append(imgs, GetKubernetesImage(constants.KubeScheduler, cfg))
		imgs = append(imgs, GetKubernetesImage(constants.KubeProxy, cfg))
	}
	imgs = append(imgs, GetPauseImage(cfg))
	if cfg.Etcd.Local != nil {
		imgs = append(imgs, GetEtcdImage(cfg))
	}
	if cfg.DNS.Type == kubeadmapi.CoreDNS {
		imgs = append(imgs, GetDNSImage(cfg, constants.CoreDNSImageName))
	} else {
		imgs = append(imgs, GetDNSImage(cfg, constants.KubeDNSKubeDNSImageName))
		imgs = append(imgs, GetDNSImage(cfg, constants.KubeDNSSidecarImageName))
		imgs = append(imgs, GetDNSImage(cfg, constants.KubeDNSDnsMasqNannyImageName))
	}
	return imgs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
