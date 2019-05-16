package controlplane

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	certphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	staticpodutil "k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	authzmodes "k8s.io/kubernetes/pkg/kubeapiserver/authorizer/modes"
	"net"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

func CreateInitStaticPodManifestFiles(manifestDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("[control-plane] creating static Pod files")
	return CreateStaticPodFiles(manifestDir, cfg, kubeadmconstants.KubeAPIServer, kubeadmconstants.KubeControllerManager, kubeadmconstants.KubeScheduler)
}
func GetStaticPodSpecs(cfg *kubeadmapi.InitConfiguration, k8sVersion *version.Version) map[string]v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mounts := getHostPathVolumesForTheControlPlane(cfg)
	staticPodSpecs := map[string]v1.Pod{kubeadmconstants.KubeAPIServer: staticpodutil.ComponentPod(v1.Container{Name: kubeadmconstants.KubeAPIServer, Image: images.GetKubernetesImage(kubeadmconstants.KubeAPIServer, &cfg.ClusterConfiguration), ImagePullPolicy: v1.PullIfNotPresent, Command: getAPIServerCommand(cfg), VolumeMounts: staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeAPIServer)), LivenessProbe: staticpodutil.ComponentProbe(cfg, kubeadmconstants.KubeAPIServer, int(cfg.LocalAPIEndpoint.BindPort), "/healthz", v1.URISchemeHTTPS), Resources: staticpodutil.ComponentResources("250m"), Env: getProxyEnvVars()}, mounts.GetVolumes(kubeadmconstants.KubeAPIServer)), kubeadmconstants.KubeControllerManager: staticpodutil.ComponentPod(v1.Container{Name: kubeadmconstants.KubeControllerManager, Image: images.GetKubernetesImage(kubeadmconstants.KubeControllerManager, &cfg.ClusterConfiguration), ImagePullPolicy: v1.PullIfNotPresent, Command: getControllerManagerCommand(cfg, k8sVersion), VolumeMounts: staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeControllerManager)), LivenessProbe: staticpodutil.ComponentProbe(cfg, kubeadmconstants.KubeControllerManager, 10252, "/healthz", v1.URISchemeHTTP), Resources: staticpodutil.ComponentResources("200m"), Env: getProxyEnvVars()}, mounts.GetVolumes(kubeadmconstants.KubeControllerManager)), kubeadmconstants.KubeScheduler: staticpodutil.ComponentPod(v1.Container{Name: kubeadmconstants.KubeScheduler, Image: images.GetKubernetesImage(kubeadmconstants.KubeScheduler, &cfg.ClusterConfiguration), ImagePullPolicy: v1.PullIfNotPresent, Command: getSchedulerCommand(cfg), VolumeMounts: staticpodutil.VolumeMountMapToSlice(mounts.GetVolumeMounts(kubeadmconstants.KubeScheduler)), LivenessProbe: staticpodutil.ComponentProbe(cfg, kubeadmconstants.KubeScheduler, 10251, "/healthz", v1.URISchemeHTTP), Resources: staticpodutil.ComponentResources("100m"), Env: getProxyEnvVars()}, mounts.GetVolumes(kubeadmconstants.KubeScheduler))}
	return staticPodSpecs
}
func CreateStaticPodFiles(manifestDir string, cfg *kubeadmapi.InitConfiguration, componentNames ...string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	k8sVersion, err := version.ParseSemantic(cfg.KubernetesVersion)
	if err != nil {
		return err
	}
	klog.V(1).Infoln("[control-plane] getting StaticPodSpecs")
	specs := GetStaticPodSpecs(cfg, k8sVersion)
	for _, componentName := range componentNames {
		spec, exists := specs[componentName]
		if !exists {
			return errors.Errorf("couldn't retrive StaticPodSpec for %q", componentName)
		}
		if err := staticpodutil.WriteStaticPodToDisk(componentName, manifestDir, spec); err != nil {
			return errors.Wrapf(err, "failed to create static pod manifest file for %q", componentName)
		}
		klog.V(1).Infof("[control-plane] wrote static Pod manifest for component %q to %q\n", componentName, kubeadmconstants.GetStaticPodFilepath(componentName, manifestDir))
	}
	return nil
}
func getAPIServerCommand(cfg *kubeadmapi.InitConfiguration) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultArguments := map[string]string{"advertise-address": cfg.LocalAPIEndpoint.AdvertiseAddress, "insecure-port": "0", "enable-admission-plugins": "NodeRestriction", "service-cluster-ip-range": cfg.Networking.ServiceSubnet, "service-account-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.ServiceAccountPublicKeyName), "client-ca-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName), "tls-cert-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerCertName), "tls-private-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKeyName), "kubelet-client-certificate": filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKubeletClientCertName), "kubelet-client-key": filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerKubeletClientKeyName), "enable-bootstrap-token-auth": "true", "secure-port": fmt.Sprintf("%d", cfg.LocalAPIEndpoint.BindPort), "allow-privileged": "true", "kubelet-preferred-address-types": "InternalIP,ExternalIP,Hostname", "requestheader-username-headers": "X-Remote-User", "requestheader-group-headers": "X-Remote-Group", "requestheader-extra-headers-prefix": "X-Remote-Extra-", "requestheader-client-ca-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertName), "requestheader-allowed-names": "front-proxy-client", "proxy-client-cert-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyClientCertName), "proxy-client-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyClientKeyName)}
	command := []string{"kube-apiserver"}
	if cfg.Etcd.External != nil {
		defaultArguments["etcd-servers"] = strings.Join(cfg.Etcd.External.Endpoints, ",")
		if cfg.Etcd.External.CAFile != "" {
			defaultArguments["etcd-cafile"] = cfg.Etcd.External.CAFile
		}
		if cfg.Etcd.External.CertFile != "" && cfg.Etcd.External.KeyFile != "" {
			defaultArguments["etcd-certfile"] = cfg.Etcd.External.CertFile
			defaultArguments["etcd-keyfile"] = cfg.Etcd.External.KeyFile
		}
	} else {
		defaultArguments["etcd-servers"] = fmt.Sprintf("https://127.0.0.1:%d", kubeadmconstants.EtcdListenClientPort)
		defaultArguments["etcd-cafile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdCACertName)
		defaultArguments["etcd-certfile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerEtcdClientCertName)
		defaultArguments["etcd-keyfile"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.APIServerEtcdClientKeyName)
		if cfg.Etcd.Local != nil {
			if value, ok := cfg.Etcd.Local.ExtraArgs["advertise-client-urls"]; ok {
				defaultArguments["etcd-servers"] = value
			}
		}
	}
	if cfg.APIServer.ExtraArgs == nil {
		cfg.APIServer.ExtraArgs = map[string]string{}
	}
	cfg.APIServer.ExtraArgs["authorization-mode"] = getAuthzModes(cfg.APIServer.ExtraArgs["authorization-mode"])
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.APIServer.ExtraArgs)...)
	return command
}
func getAuthzModes(authzModeExtraArgs string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	modes := []string{authzmodes.ModeNode, authzmodes.ModeRBAC}
	if strings.Contains(authzModeExtraArgs, authzmodes.ModeABAC) {
		modes = append(modes, authzmodes.ModeABAC)
	}
	if strings.Contains(authzModeExtraArgs, authzmodes.ModeWebhook) {
		modes = append(modes, authzmodes.ModeWebhook)
	}
	return strings.Join(modes, ",")
}
func calcNodeCidrSize(podSubnet string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	maskSize := "24"
	if ip, podCidr, err := net.ParseCIDR(podSubnet); err == nil {
		if ip.To4() == nil {
			var nodeCidrSize int
			podNetSize, totalBits := podCidr.Mask.Size()
			switch {
			case podNetSize == 112:
				nodeCidrSize = 120
			case podNetSize < 112:
				nodeCidrSize = totalBits - ((totalBits-podNetSize-1)/8-1)*8
			default:
				nodeCidrSize = podNetSize
			}
			maskSize = strconv.Itoa(nodeCidrSize)
		}
	}
	return maskSize
}
func getControllerManagerCommand(cfg *kubeadmapi.InitConfiguration, k8sVersion *version.Version) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultArguments := map[string]string{"address": "127.0.0.1", "leader-elect": "true", "kubeconfig": filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ControllerManagerKubeConfigFileName), "root-ca-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName), "service-account-private-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.ServiceAccountPrivateKeyName), "cluster-signing-cert-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName), "cluster-signing-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.CAKeyName), "use-service-account-credentials": "true", "controllers": "*,bootstrapsigner,tokencleaner"}
	if k8sVersion.Major() >= 1 && k8sVersion.Minor() >= 12 {
		defaultArguments["authentication-kubeconfig"] = filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ControllerManagerKubeConfigFileName)
		defaultArguments["authorization-kubeconfig"] = filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ControllerManagerKubeConfigFileName)
		defaultArguments["client-ca-file"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.CACertName)
		defaultArguments["requestheader-client-ca-file"] = filepath.Join(cfg.CertificatesDir, kubeadmconstants.FrontProxyCACertName)
	}
	if res, _ := certphase.UsingExternalCA(cfg); res {
		defaultArguments["cluster-signing-key-file"] = ""
		defaultArguments["cluster-signing-cert-file"] = ""
	}
	if cfg.Networking.PodSubnet != "" {
		maskSize := calcNodeCidrSize(cfg.Networking.PodSubnet)
		defaultArguments["allocate-node-cidrs"] = "true"
		defaultArguments["cluster-cidr"] = cfg.Networking.PodSubnet
		defaultArguments["node-cidr-mask-size"] = maskSize
	}
	command := []string{"kube-controller-manager"}
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.ControllerManager.ExtraArgs)...)
	return command
}
func getSchedulerCommand(cfg *kubeadmapi.InitConfiguration) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultArguments := map[string]string{"address": "127.0.0.1", "leader-elect": "true", "kubeconfig": filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.SchedulerKubeConfigFileName)}
	command := []string{"kube-scheduler"}
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.Scheduler.ExtraArgs)...)
	return command
}
func getProxyEnvVars() []v1.EnvVar {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	envs := []v1.EnvVar{}
	for _, env := range os.Environ() {
		pos := strings.Index(env, "=")
		if pos == -1 {
			continue
		}
		name := env[:pos]
		value := env[pos+1:]
		if strings.HasSuffix(strings.ToLower(name), "_proxy") && value != "" {
			envVar := v1.EnvVar{Name: name, Value: value}
			envs = append(envs, envVar)
		}
	}
	return envs
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
