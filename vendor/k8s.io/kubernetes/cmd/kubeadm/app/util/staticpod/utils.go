package staticpod

import (
	"bytes"
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	"net"
	"net/url"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

const (
	kubeControllerManagerAddressArg = "address"
	kubeSchedulerAddressArg         = "address"
	etcdListenClientURLsArg         = "listen-client-urls"
)

func ComponentPod(container v1.Container, volumes map[string]v1.Volume) v1.Pod {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.Pod{TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Pod"}, ObjectMeta: metav1.ObjectMeta{Name: container.Name, Namespace: metav1.NamespaceSystem, Annotations: map[string]string{kubetypes.CriticalPodAnnotationKey: ""}, Labels: map[string]string{"component": container.Name, "tier": "control-plane"}}, Spec: v1.PodSpec{Containers: []v1.Container{container}, PriorityClassName: "system-cluster-critical", HostNetwork: true, Volumes: VolumeMapToSlice(volumes)}}
}
func ComponentResources(cpu string) v1.ResourceRequirements {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceName(v1.ResourceCPU): resource.MustParse(cpu)}}
}
func ComponentProbe(cfg *kubeadmapi.InitConfiguration, componentName string, port int, path string, scheme v1.URIScheme) *v1.Probe {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Probe{Handler: v1.Handler{HTTPGet: &v1.HTTPGetAction{Host: GetProbeAddress(cfg, componentName), Path: path, Port: intstr.FromInt(port), Scheme: scheme}}, InitialDelaySeconds: 15, TimeoutSeconds: 15, FailureThreshold: 8}
}
func EtcdProbe(cfg *kubeadmapi.InitConfiguration, componentName string, port int, certsDir string, CACertName string, CertName string, KeyName string) *v1.Probe {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tlsFlags := fmt.Sprintf("--cacert=%[1]s/%[2]s --cert=%[1]s/%[3]s --key=%[1]s/%[4]s", certsDir, CACertName, CertName, KeyName)
	cmd := fmt.Sprintf("ETCDCTL_API=3 etcdctl --endpoints=https://[%s]:%d %s get foo", GetProbeAddress(cfg, componentName), port, tlsFlags)
	return &v1.Probe{Handler: v1.Handler{Exec: &v1.ExecAction{Command: []string{"/bin/sh", "-ec", cmd}}}, InitialDelaySeconds: 15, TimeoutSeconds: 15, FailureThreshold: 8}
}
func NewVolume(name, path string, pathType *v1.HostPathType) v1.Volume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.Volume{Name: name, VolumeSource: v1.VolumeSource{HostPath: &v1.HostPathVolumeSource{Path: path, Type: pathType}}}
}
func NewVolumeMount(name, path string, readOnly bool) v1.VolumeMount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v1.VolumeMount{Name: name, MountPath: path, ReadOnly: readOnly}
}
func VolumeMapToSlice(volumes map[string]v1.Volume) []v1.Volume {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v := make([]v1.Volume, 0, len(volumes))
	for _, vol := range volumes {
		v = append(v, vol)
	}
	sort.Slice(v, func(i, j int) bool {
		return strings.Compare(v[i].Name, v[j].Name) == -1
	})
	return v
}
func VolumeMountMapToSlice(volumeMounts map[string]v1.VolumeMount) []v1.VolumeMount {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v := make([]v1.VolumeMount, 0, len(volumeMounts))
	for _, volMount := range volumeMounts {
		v = append(v, volMount)
	}
	sort.Slice(v, func(i, j int) bool {
		return strings.Compare(v[i].Name, v[j].Name) == -1
	})
	return v
}
func GetExtraParameters(overrides map[string]string, defaults map[string]string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var command []string
	for k, v := range overrides {
		if len(v) > 0 {
			command = append(command, fmt.Sprintf("--%s=%s", k, v))
		}
	}
	for k, v := range defaults {
		if _, overrideExists := overrides[k]; !overrideExists {
			command = append(command, fmt.Sprintf("--%s=%s", k, v))
		}
	}
	return command
}
func WriteStaticPodToDisk(componentName, manifestDir string, pod v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := os.MkdirAll(manifestDir, 0700); err != nil {
		return errors.Wrapf(err, "failed to create directory %q", manifestDir)
	}
	serialized, err := util.MarshalToYaml(&pod, v1.SchemeGroupVersion)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal manifest for %q to YAML", componentName)
	}
	filename := kubeadmconstants.GetStaticPodFilepath(componentName, manifestDir)
	if err := ioutil.WriteFile(filename, serialized, 0600); err != nil {
		return errors.Wrapf(err, "failed to write static pod manifest file for %q (%q)", componentName, filename)
	}
	return nil
}
func ReadStaticPodFromDisk(manifestPath string) (*v1.Pod, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	buf, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return &v1.Pod{}, errors.Wrapf(err, "failed to read manifest for %q", manifestPath)
	}
	obj, err := util.UnmarshalFromYaml(buf, v1.SchemeGroupVersion)
	if err != nil {
		return &v1.Pod{}, errors.Errorf("failed to unmarshal manifest for %q from YAML: %v", manifestPath, err)
	}
	pod := obj.(*v1.Pod)
	return pod, nil
}
func GetProbeAddress(cfg *kubeadmapi.InitConfiguration, componentName string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch {
	case componentName == kubeadmconstants.KubeAPIServer:
		if cfg.LocalAPIEndpoint.AdvertiseAddress != "" {
			return cfg.LocalAPIEndpoint.AdvertiseAddress
		}
	case componentName == kubeadmconstants.KubeControllerManager:
		if addr, exists := cfg.ControllerManager.ExtraArgs[kubeControllerManagerAddressArg]; exists {
			return addr
		}
	case componentName == kubeadmconstants.KubeScheduler:
		if addr, exists := cfg.Scheduler.ExtraArgs[kubeSchedulerAddressArg]; exists {
			return addr
		}
	case componentName == kubeadmconstants.Etcd:
		if cfg.Etcd.Local != nil && cfg.Etcd.Local.ExtraArgs != nil {
			if arg, exists := cfg.Etcd.Local.ExtraArgs[etcdListenClientURLsArg]; exists {
				if strings.ContainsAny(arg, ",") {
					arg = strings.Split(arg, ",")[0]
				}
				parsedURL, err := url.Parse(arg)
				if err != nil || parsedURL.Hostname() == "" {
					break
				}
				if ip := net.ParseIP(parsedURL.Hostname()); ip != nil {
					if ip.Equal(net.IPv4zero) {
						return "127.0.0.1"
					}
					if ip.Equal(net.IPv6zero) {
						return net.IPv6loopback.String()
					}
					return ip.String()
				}
				addrs, err := net.LookupIP(parsedURL.Hostname())
				if err != nil {
					break
				}
				var ip net.IP
				for _, addr := range addrs {
					if addr.To4() != nil {
						ip = addr
						break
					}
					if addr.To16() != nil && ip == nil {
						ip = addr
					}
				}
				return ip.String()
			}
		}
	}
	return "127.0.0.1"
}
func ManifestFilesAreEqual(path1, path2 string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	content1, err := ioutil.ReadFile(path1)
	if err != nil {
		return false, err
	}
	content2, err := ioutil.ReadFile(path2)
	if err != nil {
		return false, err
	}
	return bytes.Equal(content1, content2), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
