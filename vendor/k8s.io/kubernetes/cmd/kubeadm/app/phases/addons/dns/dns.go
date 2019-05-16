package dns

import (
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/mholt/caddy/caddyfile"
	"github.com/pkg/errors"
	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kuberuntime "k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

const (
	KubeDNSServiceAccountName  = "kube-dns"
	kubeDNSStubDomain          = "stubDomains"
	kubeDNSUpstreamNameservers = "upstreamNameservers"
	kubeDNSFederation          = "federations"
)

func DeployedDNSAddon(client clientset.Interface) (kubeadmapi.DNSAddOnType, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	deploymentsClient := client.AppsV1().Deployments(metav1.NamespaceSystem)
	deployments, err := deploymentsClient.List(metav1.ListOptions{LabelSelector: "k8s-app=kube-dns"})
	if err != nil {
		return "", "", errors.Wrap(err, "couldn't retrieve DNS addon deployments")
	}
	switch len(deployments.Items) {
	case 0:
		return "", "", nil
	case 1:
		addonName := deployments.Items[0].Name
		addonType := kubeadmapi.CoreDNS
		if addonName == kubeadmconstants.KubeDNSDeploymentName {
			addonType = kubeadmapi.KubeDNS
		}
		addonImage := deployments.Items[0].Spec.Template.Spec.Containers[0].Image
		addonImageParts := strings.Split(addonImage, ":")
		addonVersion := addonImageParts[len(addonImageParts)-1]
		return addonType, addonVersion, nil
	default:
		return "", "", errors.Errorf("multiple DNS addon deployments found: %v", deployments.Items)
	}
}
func EnsureDNSAddon(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if cfg.DNS.Type == kubeadmapi.CoreDNS {
		return coreDNSAddon(cfg, client)
	}
	return kubeDNSAddon(cfg, client)
}
func kubeDNSAddon(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := CreateServiceAccount(client); err != nil {
		return err
	}
	dnsip, err := kubeadmconstants.GetDNSIP(cfg.Networking.ServiceSubnet)
	if err != nil {
		return err
	}
	var dnsBindAddr, dnsProbeAddr string
	if dnsip.To4() == nil {
		dnsBindAddr = "::1"
		dnsProbeAddr = "[" + dnsBindAddr + "]"
	} else {
		dnsBindAddr = "127.0.0.1"
		dnsProbeAddr = dnsBindAddr
	}
	dnsDeploymentBytes, err := kubeadmutil.ParseTemplate(KubeDNSDeployment, struct{ DeploymentName, KubeDNSImage, DNSMasqImage, SidecarImage, DNSBindAddr, DNSProbeAddr, DNSDomain, MasterTaintKey string }{DeploymentName: kubeadmconstants.KubeDNSDeploymentName, KubeDNSImage: images.GetDNSImage(&cfg.ClusterConfiguration, kubeadmconstants.KubeDNSKubeDNSImageName), DNSMasqImage: images.GetDNSImage(&cfg.ClusterConfiguration, kubeadmconstants.KubeDNSDnsMasqNannyImageName), SidecarImage: images.GetDNSImage(&cfg.ClusterConfiguration, kubeadmconstants.KubeDNSSidecarImageName), DNSBindAddr: dnsBindAddr, DNSProbeAddr: dnsProbeAddr, DNSDomain: cfg.Networking.DNSDomain, MasterTaintKey: kubeadmconstants.LabelNodeRoleMaster})
	if err != nil {
		return errors.Wrap(err, "error when parsing kube-dns deployment template")
	}
	dnsServiceBytes, err := kubeadmutil.ParseTemplate(KubeDNSService, struct{ DNSIP string }{DNSIP: dnsip.String()})
	if err != nil {
		return errors.Wrap(err, "error when parsing kube-proxy configmap template")
	}
	if err := createKubeDNSAddon(dnsDeploymentBytes, dnsServiceBytes, client); err != nil {
		return err
	}
	fmt.Println("[addons] Applied essential addon: kube-dns")
	return nil
}
func CreateServiceAccount(client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apiclient.CreateOrUpdateServiceAccount(client, &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: KubeDNSServiceAccountName, Namespace: metav1.NamespaceSystem}})
}
func createKubeDNSAddon(deploymentBytes, serviceBytes []byte, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubednsDeployment := &apps.Deployment{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), deploymentBytes, kubednsDeployment); err != nil {
		return errors.Wrap(err, "unable to decode kube-dns deployment")
	}
	if err := apiclient.CreateOrUpdateDeployment(client, kubednsDeployment); err != nil {
		return err
	}
	kubednsService := &v1.Service{}
	return createDNSService(kubednsService, serviceBytes, client)
}
func coreDNSAddon(cfg *kubeadmapi.InitConfiguration, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	coreDNSDeploymentBytes, err := kubeadmutil.ParseTemplate(CoreDNSDeployment, struct{ DeploymentName, Image, MasterTaintKey string }{DeploymentName: kubeadmconstants.CoreDNSDeploymentName, Image: images.GetDNSImage(&cfg.ClusterConfiguration, kubeadmconstants.CoreDNSImageName), MasterTaintKey: kubeadmconstants.LabelNodeRoleMaster})
	if err != nil {
		return errors.Wrap(err, "error when parsing CoreDNS deployment template")
	}
	kubeDNSConfigMap, err := client.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(kubeadmconstants.KubeDNSConfigMap, metav1.GetOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	stubDomain, err := translateStubDomainOfKubeDNSToProxyCoreDNS(kubeDNSStubDomain, kubeDNSConfigMap)
	if err != nil {
		return err
	}
	upstreamNameserver, err := translateUpstreamNameServerOfKubeDNSToUpstreamProxyCoreDNS(kubeDNSUpstreamNameservers, kubeDNSConfigMap)
	if err != nil {
		return err
	}
	coreDNSDomain := cfg.Networking.DNSDomain
	federations, err := translateFederationsofKubeDNSToCoreDNS(kubeDNSFederation, coreDNSDomain, kubeDNSConfigMap)
	if err != nil {
		return err
	}
	coreDNSConfigMapBytes, err := kubeadmutil.ParseTemplate(CoreDNSConfigMap, struct{ DNSDomain, UpstreamNameserver, Federation, StubDomain string }{DNSDomain: coreDNSDomain, UpstreamNameserver: upstreamNameserver, Federation: federations, StubDomain: stubDomain})
	if err != nil {
		return errors.Wrap(err, "error when parsing CoreDNS configMap template")
	}
	dnsip, err := kubeadmconstants.GetDNSIP(cfg.Networking.ServiceSubnet)
	if err != nil {
		return err
	}
	coreDNSServiceBytes, err := kubeadmutil.ParseTemplate(KubeDNSService, struct{ DNSIP string }{DNSIP: dnsip.String()})
	if err != nil {
		return errors.Wrap(err, "error when parsing CoreDNS service template")
	}
	if err := createCoreDNSAddon(coreDNSDeploymentBytes, coreDNSServiceBytes, coreDNSConfigMapBytes, client); err != nil {
		return err
	}
	fmt.Println("[addons] Applied essential addon: CoreDNS")
	return nil
}
func createCoreDNSAddon(deploymentBytes, serviceBytes, configBytes []byte, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	coreDNSConfigMap := &v1.ConfigMap{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), configBytes, coreDNSConfigMap); err != nil {
		return errors.Wrap(err, "unable to decode CoreDNS configmap")
	}
	if err := apiclient.CreateOrRetainConfigMap(client, coreDNSConfigMap, kubeadmconstants.CoreDNSConfigMap); err != nil {
		return err
	}
	coreDNSClusterRoles := &rbac.ClusterRole{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), []byte(CoreDNSClusterRole), coreDNSClusterRoles); err != nil {
		return errors.Wrap(err, "unable to decode CoreDNS clusterroles")
	}
	if err := apiclient.CreateOrUpdateClusterRole(client, coreDNSClusterRoles); err != nil {
		return err
	}
	coreDNSClusterRolesBinding := &rbac.ClusterRoleBinding{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), []byte(CoreDNSClusterRoleBinding), coreDNSClusterRolesBinding); err != nil {
		return errors.Wrap(err, "unable to decode CoreDNS clusterrolebindings")
	}
	if err := apiclient.CreateOrUpdateClusterRoleBinding(client, coreDNSClusterRolesBinding); err != nil {
		return err
	}
	coreDNSServiceAccount := &v1.ServiceAccount{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), []byte(CoreDNSServiceAccount), coreDNSServiceAccount); err != nil {
		return errors.Wrap(err, "unable to decode CoreDNS serviceaccount")
	}
	if err := apiclient.CreateOrUpdateServiceAccount(client, coreDNSServiceAccount); err != nil {
		return err
	}
	coreDNSDeployment := &apps.Deployment{}
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), deploymentBytes, coreDNSDeployment); err != nil {
		return errors.Wrap(err, "unable to decode CoreDNS deployment")
	}
	if err := apiclient.CreateOrUpdateDeployment(client, coreDNSDeployment); err != nil {
		return err
	}
	coreDNSService := &v1.Service{}
	return createDNSService(coreDNSService, serviceBytes, client)
}
func createDNSService(dnsService *v1.Service, serviceBytes []byte, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := kuberuntime.DecodeInto(clientsetscheme.Codecs.UniversalDecoder(), serviceBytes, dnsService); err != nil {
		return errors.Wrap(err, "unable to decode the DNS service")
	}
	if _, err := client.CoreV1().Services(metav1.NamespaceSystem).Create(dnsService); err != nil {
		if !apierrors.IsAlreadyExists(err) && !apierrors.IsInvalid(err) {
			return errors.Wrap(err, "unable to create a new DNS service")
		}
		if _, err := client.CoreV1().Services(metav1.NamespaceSystem).Update(dnsService); err != nil {
			return errors.Wrap(err, "unable to create/update the DNS service")
		}
	}
	return nil
}
func translateStubDomainOfKubeDNSToProxyCoreDNS(dataField string, kubeDNSConfigMap *v1.ConfigMap) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeDNSConfigMap == nil {
		return "", nil
	}
	if proxy, ok := kubeDNSConfigMap.Data[dataField]; ok {
		stubDomainData := make(map[string][]string)
		err := json.Unmarshal([]byte(proxy), &stubDomainData)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse JSON from 'kube-dns ConfigMap")
		}
		var proxyStanza []interface{}
		for domain, proxyIP := range stubDomainData {
			pStanza := map[string]interface{}{}
			pStanza["keys"] = []string{domain + ":53"}
			pStanza["body"] = [][]string{{"errors"}, {"cache", "30"}, {"loop"}, append([]string{"proxy", "."}, proxyIP...)}
			proxyStanza = append(proxyStanza, pStanza)
		}
		stanzasBytes, err := json.Marshal(proxyStanza)
		if err != nil {
			return "", err
		}
		corefileStanza, err := caddyfile.FromJSON(stanzasBytes)
		if err != nil {
			return "", err
		}
		return prepCorefileFormat(string(corefileStanza), 4), nil
	}
	return "", nil
}
func translateUpstreamNameServerOfKubeDNSToUpstreamProxyCoreDNS(dataField string, kubeDNSConfigMap *v1.ConfigMap) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeDNSConfigMap == nil {
		return "", nil
	}
	if upstreamValues, ok := kubeDNSConfigMap.Data[dataField]; ok {
		var upstreamProxyIP []string
		err := json.Unmarshal([]byte(upstreamValues), &upstreamProxyIP)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse JSON from 'kube-dns ConfigMap")
		}
		coreDNSProxyStanzaList := strings.Join(upstreamProxyIP, " ")
		return coreDNSProxyStanzaList, nil
	}
	return "/etc/resolv.conf", nil
}
func translateFederationsofKubeDNSToCoreDNS(dataField, coreDNSDomain string, kubeDNSConfigMap *v1.ConfigMap) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeDNSConfigMap == nil {
		return "", nil
	}
	if federation, ok := kubeDNSConfigMap.Data[dataField]; ok {
		var (
			federationStanza []interface{}
			body             [][]string
		)
		federationData := make(map[string]string)
		err := json.Unmarshal([]byte(federation), &federationData)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse JSON from kube-dns ConfigMap")
		}
		fStanza := map[string]interface{}{}
		for name, domain := range federationData {
			body = append(body, []string{name, domain})
		}
		federationStanza = append(federationStanza, fStanza)
		fStanza["keys"] = []string{"federation " + coreDNSDomain}
		fStanza["body"] = body
		stanzasBytes, err := json.Marshal(federationStanza)
		if err != nil {
			return "", err
		}
		corefileStanza, err := caddyfile.FromJSON(stanzasBytes)
		if err != nil {
			return "", err
		}
		return prepCorefileFormat(string(corefileStanza), 8), nil
	}
	return "", nil
}
func prepCorefileFormat(s string, indentation int) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := []string{}
	for _, line := range strings.Split(s, "\n") {
		indented := strings.Repeat(" ", indentation) + line
		r = append(r, indented)
	}
	corefile := strings.Join(r, "\n")
	return "\n" + strings.Replace(corefile, "\t", "   ", -1)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
