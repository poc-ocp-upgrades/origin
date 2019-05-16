package etcd

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	etcdutil "k8s.io/kubernetes/cmd/kubeadm/app/util/etcd"
	staticpodutil "k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const (
	etcdVolumeName           = "etcd-data"
	certsVolumeName          = "etcd-certs"
	etcdHealthyCheckInterval = 5 * time.Second
	etcdHealthyCheckRetries  = 8
)

func CreateLocalEtcdStaticPodManifestFile(manifestDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.ClusterConfiguration.Etcd.External != nil {
		return errors.New("etcd static pod manifest cannot be generated for cluster using external etcd")
	}
	emptyInitialCluster := []etcdutil.Member{}
	spec := GetEtcdPodSpec(cfg, emptyInitialCluster)
	if err := staticpodutil.WriteStaticPodToDisk(kubeadmconstants.Etcd, manifestDir, spec); err != nil {
		return err
	}
	klog.V(1).Infof("[etcd] wrote Static Pod manifest for a local etcd member to %q\n", kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.Etcd, manifestDir))
	return nil
}
func CheckLocalEtcdClusterStatus(client clientset.Interface, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[etcd] Checking etcd cluster health")
	klog.V(1).Info("creating etcd client that connects to etcd pods")
	etcdClient, err := etcdutil.NewFromCluster(client, cfg.CertificatesDir)
	if err != nil {
		return err
	}
	_, err = etcdClient.GetClusterStatus()
	if err != nil {
		return errors.Wrap(err, "etcd cluster is not healthy")
	}
	return nil
}
func CreateStackedEtcdStaticPodManifestFile(client clientset.Interface, manifestDir string, cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Info("creating etcd client that connects to etcd pods")
	etcdClient, err := etcdutil.NewFromCluster(client, cfg.CertificatesDir)
	if err != nil {
		return err
	}
	etcdPeerAddress := etcdutil.GetPeerURL(cfg)
	klog.V(1).Infof("Adding etcd member: %s", etcdPeerAddress)
	initialCluster, err := etcdClient.AddMember(cfg.NodeRegistration.Name, etcdPeerAddress)
	if err != nil {
		return err
	}
	fmt.Println("[etcd] Announced new etcd member joining to the existing etcd cluster")
	klog.V(1).Infof("Updated etcd member list: %v", initialCluster)
	klog.V(1).Info("Creating local etcd static pod manifest file")
	spec := GetEtcdPodSpec(cfg, initialCluster)
	if err := staticpodutil.WriteStaticPodToDisk(kubeadmconstants.Etcd, manifestDir, spec); err != nil {
		return err
	}
	fmt.Printf("[etcd] Wrote Static Pod manifest for a local etcd member to %q\n", kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.Etcd, manifestDir))
	fmt.Printf("[etcd] Waiting for the new etcd member to join the cluster. This can take up to %v\n", etcdHealthyCheckInterval*etcdHealthyCheckRetries)
	noDelay := 0 * time.Second
	if _, err := etcdClient.WaitForClusterAvailable(noDelay, etcdHealthyCheckRetries, etcdHealthyCheckInterval); err != nil {
		return err
	}
	return nil
}
func GetEtcdPodSpec(cfg *kubeadmapi.InitConfiguration, initialCluster []etcdutil.Member) v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pathType := v1.HostPathDirectoryOrCreate
	etcdMounts := map[string]v1.Volume{etcdVolumeName: staticpodutil.NewVolume(etcdVolumeName, cfg.Etcd.Local.DataDir, &pathType), certsVolumeName: staticpodutil.NewVolume(certsVolumeName, cfg.CertificatesDir+"/etcd", &pathType)}
	return staticpodutil.ComponentPod(v1.Container{Name: kubeadmconstants.Etcd, Command: getEtcdCommand(cfg, initialCluster), Image: images.GetEtcdImage(&cfg.ClusterConfiguration), ImagePullPolicy: v1.PullIfNotPresent, VolumeMounts: []v1.VolumeMount{staticpodutil.NewVolumeMount(etcdVolumeName, cfg.Etcd.Local.DataDir, false), staticpodutil.NewVolumeMount(certsVolumeName, cfg.CertificatesDir+"/etcd", false)}, LivenessProbe: staticpodutil.EtcdProbe(cfg, kubeadmconstants.Etcd, kubeadmconstants.EtcdListenClientPort, cfg.CertificatesDir, kubeadmconstants.EtcdCACertName, kubeadmconstants.EtcdHealthcheckClientCertName, kubeadmconstants.EtcdHealthcheckClientKeyName)}, etcdMounts)
}
func getEtcdCommand(cfg *kubeadmapi.InitConfiguration, initialCluster []etcdutil.Member) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultArguments := map[string]string{"name": cfg.GetNodeName(), "listen-client-urls": fmt.Sprintf("%s,%s", etcdutil.GetClientURLByIP("127.0.0.1"), etcdutil.GetClientURL(cfg)), "advertise-client-urls": etcdutil.GetClientURL(cfg), "listen-peer-urls": etcdutil.GetPeerURL(cfg), "initial-advertise-peer-urls": etcdutil.GetPeerURL(cfg), "data-dir": cfg.Etcd.Local.DataDir, "cert-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdServerCertName), "key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdServerKeyName), "trusted-ca-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdCACertName), "client-cert-auth": "true", "peer-cert-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdPeerCertName), "peer-key-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdPeerKeyName), "peer-trusted-ca-file": filepath.Join(cfg.CertificatesDir, kubeadmconstants.EtcdCACertName), "peer-client-cert-auth": "true", "snapshot-count": "10000"}
	if len(initialCluster) == 0 {
		defaultArguments["initial-cluster"] = fmt.Sprintf("%s=%s", cfg.GetNodeName(), etcdutil.GetPeerURL(cfg))
	} else {
		endpoints := []string{}
		for _, member := range initialCluster {
			endpoints = append(endpoints, fmt.Sprintf("%s=%s", member.Name, member.PeerURL))
		}
		defaultArguments["initial-cluster"] = strings.Join(endpoints, ",")
		defaultArguments["initial-cluster-state"] = "existing"
	}
	command := []string{"etcd"}
	command = append(command, kubeadmutil.BuildArgumentListFromMap(defaultArguments, cfg.Etcd.Local.ExtraArgs)...)
	return command
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
