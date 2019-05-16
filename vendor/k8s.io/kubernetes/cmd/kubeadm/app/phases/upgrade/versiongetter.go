package upgrade

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	versionutil "k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/version"
)

type VersionGetter interface {
	ClusterVersion() (string, *versionutil.Version, error)
	KubeadmVersion() (string, *versionutil.Version, error)
	VersionFromCILabel(string, string) (string, *versionutil.Version, error)
	KubeletVersions() (map[string]uint16, error)
}
type KubeVersionGetter struct {
	client clientset.Interface
	w      io.Writer
}

func NewKubeVersionGetter(client clientset.Interface, writer io.Writer) VersionGetter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &KubeVersionGetter{client: client, w: writer}
}
func (g *KubeVersionGetter) ClusterVersion() (string, *versionutil.Version, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clusterVersionInfo, err := g.client.Discovery().ServerVersion()
	if err != nil {
		return "", nil, errors.Wrap(err, "Couldn't fetch cluster version from the API Server")
	}
	fmt.Fprintf(g.w, "[upgrade/versions] Cluster version: %s\n", clusterVersionInfo.String())
	clusterVersion, err := versionutil.ParseSemantic(clusterVersionInfo.String())
	if err != nil {
		return "", nil, errors.Wrap(err, "Couldn't parse cluster version")
	}
	return clusterVersionInfo.String(), clusterVersion, nil
}
func (g *KubeVersionGetter) KubeadmVersion() (string, *versionutil.Version, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeadmVersionInfo := version.Get()
	fmt.Fprintf(g.w, "[upgrade/versions] kubeadm version: %s\n", kubeadmVersionInfo.String())
	kubeadmVersion, err := versionutil.ParseSemantic(kubeadmVersionInfo.String())
	if err != nil {
		return "", nil, errors.Wrap(err, "Couldn't parse kubeadm version")
	}
	return kubeadmVersionInfo.String(), kubeadmVersion, nil
}
func (g *KubeVersionGetter) VersionFromCILabel(ciVersionLabel, description string) (string, *versionutil.Version, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionStr, err := kubeadmutil.KubernetesReleaseVersion(ciVersionLabel)
	if err != nil {
		return "", nil, errors.Wrapf(err, "Couldn't fetch latest %s from the internet", description)
	}
	if description != "" {
		fmt.Fprintf(g.w, "[upgrade/versions] Latest %s: %s\n", description, versionStr)
	}
	ver, err := versionutil.ParseSemantic(versionStr)
	if err != nil {
		return "", nil, errors.Wrapf(err, "Couldn't parse latest %s", description)
	}
	return versionStr, ver, nil
}
func (g *KubeVersionGetter) KubeletVersions() (map[string]uint16, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nodes, err := g.client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, errors.New("couldn't list all nodes in cluster")
	}
	return computeKubeletVersions(nodes.Items), nil
}
func computeKubeletVersions(nodes []v1.Node) map[string]uint16 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeletVersions := map[string]uint16{}
	for _, node := range nodes {
		kver := node.Status.NodeInfo.KubeletVersion
		if _, found := kubeletVersions[kver]; !found {
			kubeletVersions[kver] = 1
			continue
		}
		kubeletVersions[kver]++
	}
	return kubeletVersions
}

type OfflineVersionGetter struct {
	VersionGetter
	version string
}

func NewOfflineVersionGetter(versionGetter VersionGetter, version string) VersionGetter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &OfflineVersionGetter{VersionGetter: versionGetter, version: version}
}
func (o *OfflineVersionGetter) VersionFromCILabel(ciVersionLabel, description string) (string, *versionutil.Version, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if o.version == "" {
		return o.VersionGetter.VersionFromCILabel(ciVersionLabel, description)
	}
	ver, err := versionutil.ParseSemantic(o.version)
	if err != nil {
		return "", nil, errors.Wrapf(err, "Couldn't parse version %s", description)
	}
	return o.version, ver, nil
}
