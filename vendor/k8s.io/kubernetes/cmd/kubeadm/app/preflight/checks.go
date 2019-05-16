package preflight

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/PuerkitoBio/purell"
	"github.com/blang/semver"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	versionutil "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	utilruntime "k8s.io/kubernetes/cmd/kubeadm/app/util/runtime"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/system"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"k8s.io/kubernetes/pkg/util/initsystem"
	ipvsutil "k8s.io/kubernetes/pkg/util/ipvs"
	kubeadmversion "k8s.io/kubernetes/pkg/version"
	utilsexec "k8s.io/utils/exec"
	"net"
	"net/http"
	"net/url"
	"os"
	goos "os"
	"path/filepath"
	"runtime"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

const (
	bridgenf                    = "/proc/sys/net/bridge/bridge-nf-call-iptables"
	bridgenf6                   = "/proc/sys/net/bridge/bridge-nf-call-ip6tables"
	ipv4Forward                 = "/proc/sys/net/ipv4/ip_forward"
	ipv6DefaultForwarding       = "/proc/sys/net/ipv6/conf/default/forwarding"
	externalEtcdRequestTimeout  = time.Duration(10 * time.Second)
	externalEtcdRequestRetries  = 3
	externalEtcdRequestInterval = time.Duration(5 * time.Second)
)

var (
	minExternalEtcdVersion = semver.MustParse(kubeadmconstants.MinExternalEtcdVersion)
)

type Error struct{ Msg string }

func (e *Error) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("[preflight] Some fatal errors occurred:\n%s%s", e.Msg, "[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`")
}
func (e *Error) Preflight() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return true
}

type Checker interface {
	Check() (warnings, errorList []error)
	Name() string
}
type ContainerRuntimeCheck struct{ runtime utilruntime.ContainerRuntime }

func (ContainerRuntimeCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "CRI"
}
func (crc ContainerRuntimeCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating the container runtime")
	if err := crc.runtime.IsRunning(); err != nil {
		errorList = append(errorList, err)
	}
	return warnings, errorList
}

type ServiceCheck struct {
	Service       string
	CheckIfActive bool
	Label         string
}

func (sc ServiceCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if sc.Label != "" {
		return sc.Label
	}
	return fmt.Sprintf("Service-%s", strings.Title(sc.Service))
}
func (sc ServiceCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating if the service is enabled and active")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}
	warnings = []error{}
	if !initSystem.ServiceExists(sc.Service) {
		warnings = append(warnings, errors.Errorf("%s service does not exist", sc.Service))
		return warnings, nil
	}
	if !initSystem.ServiceIsEnabled(sc.Service) {
		warnings = append(warnings, errors.Errorf("%s service is not enabled, please run 'systemctl enable %s.service'", sc.Service, sc.Service))
	}
	if sc.CheckIfActive && !initSystem.ServiceIsActive(sc.Service) {
		errorList = append(errorList, errors.Errorf("%s service is not active, please run 'systemctl start %s.service'", sc.Service, sc.Service))
	}
	return warnings, errorList
}

type FirewalldCheck struct{ ports []int }

func (FirewalldCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "Firewalld"
}
func (fc FirewalldCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating if the firewall is enabled and active")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}
	warnings = []error{}
	if !initSystem.ServiceExists("firewalld") {
		return nil, nil
	}
	if initSystem.ServiceIsActive("firewalld") {
		warnings = append(warnings, errors.Errorf("firewalld is active, please ensure ports %v are open or your cluster may not function correctly", fc.ports))
	}
	return warnings, errorList
}

type PortOpenCheck struct {
	port  int
	label string
}

func (poc PortOpenCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if poc.label != "" {
		return poc.label
	}
	return fmt.Sprintf("Port-%d", poc.port)
}
func (poc PortOpenCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating availability of port %d", poc.port)
	errorList = []error{}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", poc.port))
	if err != nil {
		errorList = append(errorList, errors.Errorf("Port %d is in use", poc.port))
	}
	if ln != nil {
		ln.Close()
	}
	return nil, errorList
}

type IsPrivilegedUserCheck struct{}

func (IsPrivilegedUserCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "IsPrivilegedUser"
}

type DirAvailableCheck struct {
	Path  string
	Label string
}

func (dac DirAvailableCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dac.Label != "" {
		return dac.Label
	}
	return fmt.Sprintf("DirAvailable-%s", strings.Replace(dac.Path, "/", "-", -1))
}
func (dac DirAvailableCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating the existence and emptiness of directory %s", dac.Path)
	errorList = []error{}
	if _, err := os.Stat(dac.Path); os.IsNotExist(err) {
		return nil, nil
	}
	f, err := os.Open(dac.Path)
	if err != nil {
		errorList = append(errorList, errors.Wrapf(err, "unable to check if %s is empty", dac.Path))
		return nil, errorList
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err != io.EOF {
		errorList = append(errorList, errors.Errorf("%s is not empty", dac.Path))
	}
	return nil, errorList
}

type FileAvailableCheck struct {
	Path  string
	Label string
}

func (fac FileAvailableCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if fac.Label != "" {
		return fac.Label
	}
	return fmt.Sprintf("FileAvailable-%s", strings.Replace(fac.Path, "/", "-", -1))
}
func (fac FileAvailableCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating the existence of file %s", fac.Path)
	errorList = []error{}
	if _, err := os.Stat(fac.Path); err == nil {
		errorList = append(errorList, errors.Errorf("%s already exists", fac.Path))
	}
	return nil, errorList
}

type FileExistingCheck struct {
	Path  string
	Label string
}

func (fac FileExistingCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if fac.Label != "" {
		return fac.Label
	}
	return fmt.Sprintf("FileExisting-%s", strings.Replace(fac.Path, "/", "-", -1))
}
func (fac FileExistingCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating the existence of file %s", fac.Path)
	errorList = []error{}
	if _, err := os.Stat(fac.Path); err != nil {
		errorList = append(errorList, errors.Errorf("%s doesn't exist", fac.Path))
	}
	return nil, errorList
}

type FileContentCheck struct {
	Path    string
	Content []byte
	Label   string
}

func (fcc FileContentCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if fcc.Label != "" {
		return fcc.Label
	}
	return fmt.Sprintf("FileContent-%s", strings.Replace(fcc.Path, "/", "-", -1))
}
func (fcc FileContentCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating the contents of file %s", fcc.Path)
	f, err := os.Open(fcc.Path)
	if err != nil {
		return nil, []error{errors.Errorf("%s does not exist", fcc.Path)}
	}
	lr := io.LimitReader(f, int64(len(fcc.Content)))
	defer f.Close()
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, lr)
	if err != nil {
		return nil, []error{errors.Errorf("%s could not be read", fcc.Path)}
	}
	if !bytes.Equal(buf.Bytes(), fcc.Content) {
		return nil, []error{errors.Errorf("%s contents are not set to %s", fcc.Path, fcc.Content)}
	}
	return nil, []error{}
}

type InPathCheck struct {
	executable string
	mandatory  bool
	exec       utilsexec.Interface
	label      string
	suggestion string
}

func (ipc InPathCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if ipc.label != "" {
		return ipc.label
	}
	return fmt.Sprintf("FileExisting-%s", strings.Replace(ipc.executable, "/", "-", -1))
}
func (ipc InPathCheck) Check() (warnings, errs []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating the presence of executable %s", ipc.executable)
	_, err := ipc.exec.LookPath(ipc.executable)
	if err != nil {
		if ipc.mandatory {
			return nil, []error{errors.Errorf("%s not found in system path", ipc.executable)}
		}
		warningMessage := fmt.Sprintf("%s not found in system path", ipc.executable)
		if ipc.suggestion != "" {
			warningMessage += fmt.Sprintf("\nSuggestion: %s", ipc.suggestion)
		}
		return []error{errors.New(warningMessage)}, nil
	}
	return nil, nil
}

type HostnameCheck struct{ nodeName string }

func (HostnameCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "Hostname"
}
func (hc HostnameCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("checking whether the given node name is reachable using net.LookupHost")
	errorList = []error{}
	warnings = []error{}
	addr, err := net.LookupHost(hc.nodeName)
	if addr == nil {
		warnings = append(warnings, errors.Errorf("hostname \"%s\" could not be reached", hc.nodeName))
	}
	if err != nil {
		warnings = append(warnings, errors.Wrapf(err, "hostname \"%s\"", hc.nodeName))
	}
	return warnings, errorList
}

type HTTPProxyCheck struct {
	Proto string
	Host  string
}

func (hst HTTPProxyCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "HTTPProxy"
}
func (hst HTTPProxyCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infof("validating if the connectivity type is via proxy or direct")
	u := (&url.URL{Scheme: hst.Proto, Host: hst.Host}).String()
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, []error{err}
	}
	proxy, err := netutil.SetOldTransportDefaults(&http.Transport{}).Proxy(req)
	if err != nil {
		return nil, []error{err}
	}
	if proxy != nil {
		return []error{errors.Errorf("Connection to %q uses proxy %q. If that is not intended, adjust your proxy settings", u, proxy)}, nil
	}
	return nil, nil
}

type HTTPProxyCIDRCheck struct {
	Proto string
	CIDR  string
}

func (HTTPProxyCIDRCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "HTTPProxyCIDR"
}
func (subnet HTTPProxyCIDRCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating http connectivity to first IP address in the CIDR")
	if len(subnet.CIDR) == 0 {
		return nil, nil
	}
	_, cidr, err := net.ParseCIDR(subnet.CIDR)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "error parsing CIDR %q", subnet.CIDR)}
	}
	testIP, err := ipallocator.GetIndexedIP(cidr, 1)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "unable to get first IP address from the given CIDR (%s)", cidr.String())}
	}
	testIPstring := testIP.String()
	if len(testIP) == net.IPv6len {
		testIPstring = fmt.Sprintf("[%s]:1234", testIP)
	}
	url := fmt.Sprintf("%s://%s/", subnet.Proto, testIPstring)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, []error{err}
	}
	proxy, err := netutil.SetOldTransportDefaults(&http.Transport{}).Proxy(req)
	if err != nil {
		return nil, []error{err}
	}
	if proxy != nil {
		return []error{errors.Errorf("connection to %q uses proxy %q. This may lead to malfunctional cluster setup. Make sure that Pod and Services IP ranges specified correctly as exceptions in proxy configuration", subnet.CIDR, proxy)}, nil
	}
	return nil, nil
}

type SystemVerificationCheck struct{ IsDocker bool }

func (SystemVerificationCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "SystemVerification"
}
func (sysver SystemVerificationCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("running all checks")
	bufw := bufio.NewWriterSize(os.Stdout, 1*1024*1024)
	reporter := &system.StreamReporter{WriteStream: bufw}
	var errs []error
	var warns []error
	var validators = []system.Validator{&system.KernelValidator{Reporter: reporter}}
	if sysver.IsDocker {
		validators = append(validators, &system.DockerValidator{Reporter: reporter})
	}
	if runtime.GOOS == "linux" {
		validators = append(validators, &system.OSValidator{Reporter: reporter}, &system.CgroupsValidator{Reporter: reporter})
	}
	for _, v := range validators {
		warn, err := v.Validate(system.DefaultSysSpec)
		if err != nil {
			errs = append(errs, err)
		}
		if warn != nil {
			warns = append(warns, warn)
		}
	}
	if len(errs) != 0 {
		fmt.Println("[preflight] The system verification failed. Printing the output from the verification:")
		bufw.Flush()
		return warns, errs
	}
	return warns, nil
}

type KubernetesVersionCheck struct {
	KubeadmVersion    string
	KubernetesVersion string
}

func (KubernetesVersionCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "KubernetesVersion"
}
func (kubever KubernetesVersionCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating Kubernetes and kubeadm version")
	if strings.HasPrefix(kubever.KubeadmVersion, "v0.0.0") {
		return nil, nil
	}
	kadmVersion, err := versionutil.ParseSemantic(kubever.KubeadmVersion)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "couldn't parse kubeadm version %q", kubever.KubeadmVersion)}
	}
	k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "couldn't parse Kubernetes version %q", kubever.KubernetesVersion)}
	}
	firstUnsupportedVersion := versionutil.MustParseSemantic(fmt.Sprintf("%d.%d.%s", kadmVersion.Major(), kadmVersion.Minor()+1, "0-0"))
	if k8sVersion.AtLeast(firstUnsupportedVersion) {
		return []error{errors.Errorf("Kubernetes version is greater than kubeadm version. Please consider to upgrade kubeadm. Kubernetes version: %s. Kubeadm version: %d.%d.x", k8sVersion, kadmVersion.Components()[0], kadmVersion.Components()[1])}, nil
	}
	return nil, nil
}

type KubeletVersionCheck struct {
	KubernetesVersion string
	exec              utilsexec.Interface
}

func (KubeletVersionCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "KubeletVersion"
}
func (kubever KubeletVersionCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating kubelet version")
	kubeletVersion, err := GetKubeletVersion(kubever.exec)
	if err != nil {
		return nil, []error{errors.Wrap(err, "couldn't get kubelet version")}
	}
	if kubeletVersion.LessThan(kubeadmconstants.MinimumKubeletVersion) {
		return nil, []error{errors.Errorf("Kubelet version %q is lower than kubadm can support. Please upgrade kubelet", kubeletVersion)}
	}
	if kubever.KubernetesVersion != "" {
		k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
		if err != nil {
			return nil, []error{errors.Wrapf(err, "couldn't parse Kubernetes version %q", kubever.KubernetesVersion)}
		}
		if kubeletVersion.Major() > k8sVersion.Major() || kubeletVersion.Minor() > k8sVersion.Minor() {
			return nil, []error{errors.Errorf("the kubelet version is higher than the control plane version. This is not a supported version skew and may lead to a malfunctional cluster. Kubelet version: %q Control plane version: %q", kubeletVersion, k8sVersion)}
		}
	}
	return nil, nil
}

type SwapCheck struct{}

func (SwapCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "Swap"
}
func (swc SwapCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating whether swap is enabled or not")
	f, err := os.Open("/proc/swaps")
	if err != nil {
		return nil, nil
	}
	defer f.Close()
	var buf []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, []error{errors.Wrap(err, "error parsing /proc/swaps")}
	}
	if len(buf) > 1 {
		return nil, []error{errors.New("running with swap on is not supported. Please disable swap")}
	}
	return nil, nil
}

type etcdVersionResponse struct {
	Etcdserver  string `json:"etcdserver"`
	Etcdcluster string `json:"etcdcluster"`
}
type ExternalEtcdVersionCheck struct{ Etcd kubeadmapi.Etcd }

func (ExternalEtcdVersionCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "ExternalEtcdVersion"
}
func (evc ExternalEtcdVersionCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("validating the external etcd version")
	if evc.Etcd.External.Endpoints == nil {
		return nil, nil
	}
	var config *tls.Config
	var err error
	if config, err = evc.configRootCAs(config); err != nil {
		errorList = append(errorList, err)
		return nil, errorList
	}
	if config, err = evc.configCertAndKey(config); err != nil {
		errorList = append(errorList, err)
		return nil, errorList
	}
	client := evc.getHTTPClient(config)
	for _, endpoint := range evc.Etcd.External.Endpoints {
		if _, err := url.Parse(endpoint); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to parse external etcd endpoint %s", endpoint))
			continue
		}
		resp := etcdVersionResponse{}
		var err error
		versionURL := fmt.Sprintf("%s/%s", endpoint, "version")
		if tmpVersionURL, err := purell.NormalizeURLString(versionURL, purell.FlagRemoveDuplicateSlashes); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to normalize external etcd version url %s", versionURL))
			continue
		} else {
			versionURL = tmpVersionURL
		}
		if err = getEtcdVersionResponse(client, versionURL, &resp); err != nil {
			errorList = append(errorList, err)
			continue
		}
		etcdVersion, err := semver.Parse(resp.Etcdserver)
		if err != nil {
			errorList = append(errorList, errors.Wrapf(err, "couldn't parse external etcd version %q", resp.Etcdserver))
			continue
		}
		if etcdVersion.LT(minExternalEtcdVersion) {
			errorList = append(errorList, errors.Errorf("this version of kubeadm only supports external etcd version >= %s. Current version: %s", kubeadmconstants.MinExternalEtcdVersion, resp.Etcdserver))
			continue
		}
	}
	return nil, errorList
}
func (evc ExternalEtcdVersionCheck) configRootCAs(config *tls.Config) (*tls.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var CACertPool *x509.CertPool
	if evc.Etcd.External.CAFile != "" {
		CACert, err := ioutil.ReadFile(evc.Etcd.External.CAFile)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load external etcd's server certificate %s", evc.Etcd.External.CAFile)
		}
		CACertPool = x509.NewCertPool()
		CACertPool.AppendCertsFromPEM(CACert)
	}
	if CACertPool != nil {
		if config == nil {
			config = &tls.Config{}
		}
		config.RootCAs = CACertPool
	}
	return config, nil
}
func (evc ExternalEtcdVersionCheck) configCertAndKey(config *tls.Config) (*tls.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var cert tls.Certificate
	if evc.Etcd.External.CertFile != "" && evc.Etcd.External.KeyFile != "" {
		var err error
		cert, err = tls.LoadX509KeyPair(evc.Etcd.External.CertFile, evc.Etcd.External.KeyFile)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load external etcd's certificate and key pair %s, %s", evc.Etcd.External.CertFile, evc.Etcd.External.KeyFile)
		}
		if config == nil {
			config = &tls.Config{}
		}
		config.Certificates = []tls.Certificate{cert}
	}
	return config, nil
}
func (evc ExternalEtcdVersionCheck) getHTTPClient(config *tls.Config) *http.Client {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if config != nil {
		transport := netutil.SetOldTransportDefaults(&http.Transport{TLSClientConfig: config})
		return &http.Client{Transport: transport, Timeout: externalEtcdRequestTimeout}
	}
	return &http.Client{Timeout: externalEtcdRequestTimeout, Transport: netutil.SetOldTransportDefaults(&http.Transport{})}
}
func getEtcdVersionResponse(client *http.Client, url string, target interface{}) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	loopCount := externalEtcdRequestRetries + 1
	var err error
	var stopRetry bool
	for loopCount > 0 {
		if loopCount <= externalEtcdRequestRetries {
			time.Sleep(externalEtcdRequestInterval)
		}
		stopRetry, err = func() (stopRetry bool, err error) {
			r, err := client.Get(url)
			if err != nil {
				loopCount--
				return false, err
			}
			defer r.Body.Close()
			if r != nil && r.StatusCode >= 500 && r.StatusCode <= 599 {
				loopCount--
				return false, errors.Errorf("server responded with non-successful status: %s", r.Status)
			}
			return true, json.NewDecoder(r.Body).Decode(target)
		}()
		if stopRetry {
			break
		}
	}
	return err
}

type ImagePullCheck struct {
	runtime   utilruntime.ContainerRuntime
	imageList []string
}

func (ImagePullCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "ImagePull"
}
func (ipc ImagePullCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, image := range ipc.imageList {
		ret, err := ipc.runtime.ImageExists(image)
		if ret && err == nil {
			klog.V(1).Infof("image exists: %s", image)
			continue
		}
		if err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to check if image %s exists", image))
		}
		klog.V(1).Infof("pulling %s", image)
		if err := ipc.runtime.PullImage(image); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to pull image %s", image))
		}
	}
	return warnings, errorList
}

type NumCPUCheck struct{ NumCPU int }

func (NumCPUCheck) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "NumCPU"
}
func (ncc NumCPUCheck) Check() (warnings, errorList []error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	numCPU := runtime.NumCPU()
	if numCPU < ncc.NumCPU {
		errorList = append(errorList, errors.Errorf("the number of available CPUs %d is less than the required %d", numCPU, ncc.NumCPU))
	}
	return warnings, errorList
}
func RunInitMasterChecks(execer utilsexec.Interface, cfg *kubeadmapi.InitConfiguration, ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
		return err
	}
	manifestsDir := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)
	checks := []Checker{NumCPUCheck{NumCPU: kubeadmconstants.MasterNumCPU}, KubernetesVersionCheck{KubernetesVersion: cfg.KubernetesVersion, KubeadmVersion: kubeadmversion.Get().GitVersion}, FirewalldCheck{ports: []int{int(cfg.LocalAPIEndpoint.BindPort), 10250}}, PortOpenCheck{port: int(cfg.LocalAPIEndpoint.BindPort)}, PortOpenCheck{port: 10251}, PortOpenCheck{port: 10252}, FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeAPIServer, manifestsDir)}, FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeControllerManager, manifestsDir)}, FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeScheduler, manifestsDir)}, FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.Etcd, manifestsDir)}, HTTPProxyCheck{Proto: "https", Host: cfg.LocalAPIEndpoint.AdvertiseAddress}, HTTPProxyCIDRCheck{Proto: "https", CIDR: cfg.Networking.ServiceSubnet}, HTTPProxyCIDRCheck{Proto: "https", CIDR: cfg.Networking.PodSubnet}}
	checks = addCommonChecks(execer, cfg, checks)
	if cfg.ComponentConfigs.KubeProxy != nil && cfg.ComponentConfigs.KubeProxy.Mode == ipvsutil.IPVSProxyMode {
		checks = append(checks, ipvsutil.RequiredIPVSKernelModulesAvailableCheck{Executor: execer})
	}
	if cfg.Etcd.Local != nil {
		checks = append(checks, PortOpenCheck{port: kubeadmconstants.EtcdListenClientPort}, PortOpenCheck{port: kubeadmconstants.EtcdListenPeerPort}, DirAvailableCheck{Path: cfg.Etcd.Local.DataDir})
	}
	if cfg.Etcd.External != nil {
		if cfg.Etcd.External.CAFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.CAFile, Label: "ExternalEtcdClientCertificates"})
		}
		if cfg.Etcd.External.CertFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.CertFile, Label: "ExternalEtcdClientCertificates"})
		}
		if cfg.Etcd.External.KeyFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.KeyFile, Label: "ExternalEtcdClientCertificates"})
		}
		checks = append(checks, ExternalEtcdVersionCheck{Etcd: cfg.Etcd})
	}
	if ip := net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress); ip != nil {
		if ip.To4() == nil && ip.To16() != nil {
			checks = append(checks, FileContentCheck{Path: bridgenf6, Content: []byte{'1'}}, FileContentCheck{Path: ipv6DefaultForwarding, Content: []byte{'1'}})
		}
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}
func RunJoinNodeChecks(execer utilsexec.Interface, cfg *kubeadmapi.JoinConfiguration, ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
		return err
	}
	checks := []Checker{DirAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)}, FileAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.KubeletKubeConfigFileName)}, FileAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.KubeletBootstrapKubeConfigFileName)}}
	checks = addCommonChecks(execer, cfg, checks)
	if cfg.ControlPlane == nil {
		checks = append(checks, FileAvailableCheck{Path: cfg.CACertPath})
	}
	addIPv6Checks := false
	if cfg.Discovery.BootstrapToken != nil {
		ipstr, _, err := net.SplitHostPort(cfg.Discovery.BootstrapToken.APIServerEndpoint)
		if err == nil {
			checks = append(checks, HTTPProxyCheck{Proto: "https", Host: ipstr})
			if !addIPv6Checks {
				if ip := net.ParseIP(ipstr); ip != nil {
					if ip.To4() == nil && ip.To16() != nil {
						addIPv6Checks = true
					}
				}
			}
		}
	}
	if addIPv6Checks {
		checks = append(checks, FileContentCheck{Path: bridgenf6, Content: []byte{'1'}}, FileContentCheck{Path: ipv6DefaultForwarding, Content: []byte{'1'}})
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}
func RunOptionalJoinNodeChecks(execer utilsexec.Interface, initCfg *kubeadmapi.InitConfiguration, ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	checks := []Checker{}
	if initCfg.ComponentConfigs.KubeProxy != nil && initCfg.ComponentConfigs.KubeProxy.Mode == ipvsutil.IPVSProxyMode {
		checks = append(checks, ipvsutil.RequiredIPVSKernelModulesAvailableCheck{Executor: execer})
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}
func addCommonChecks(execer utilsexec.Interface, cfg kubeadmapi.CommonConfiguration, checks []Checker) []Checker {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	containerRuntime, err := utilruntime.NewContainerRuntime(execer, cfg.GetCRISocket())
	isDocker := false
	if err != nil {
		fmt.Printf("[preflight] WARNING: Couldn't create the interface used for talking to the container runtime: %v\n", err)
	} else {
		checks = append(checks, ContainerRuntimeCheck{runtime: containerRuntime})
		if containerRuntime.IsDocker() {
			isDocker = true
			checks = append(checks, ServiceCheck{Service: "docker", CheckIfActive: true})
		}
	}
	if runtime.GOOS == "linux" {
		if !isDocker {
			checks = append(checks, InPathCheck{executable: "crictl", mandatory: true, exec: execer})
		}
		checks = append(checks, FileContentCheck{Path: bridgenf, Content: []byte{'1'}}, FileContentCheck{Path: ipv4Forward, Content: []byte{'1'}}, SwapCheck{}, InPathCheck{executable: "ip", mandatory: true, exec: execer}, InPathCheck{executable: "iptables", mandatory: true, exec: execer}, InPathCheck{executable: "mount", mandatory: true, exec: execer}, InPathCheck{executable: "nsenter", mandatory: true, exec: execer}, InPathCheck{executable: "ebtables", mandatory: false, exec: execer}, InPathCheck{executable: "ethtool", mandatory: false, exec: execer}, InPathCheck{executable: "socat", mandatory: false, exec: execer}, InPathCheck{executable: "tc", mandatory: false, exec: execer}, InPathCheck{executable: "touch", mandatory: false, exec: execer})
	}
	checks = append(checks, SystemVerificationCheck{IsDocker: isDocker}, HostnameCheck{nodeName: cfg.GetNodeName()}, KubeletVersionCheck{KubernetesVersion: cfg.GetKubernetesVersion(), exec: execer}, ServiceCheck{Service: "kubelet", CheckIfActive: false}, PortOpenCheck{port: 10250})
	return checks
}
func RunRootCheckOnly(ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	checks := []Checker{IsPrivilegedUserCheck{}}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}
func RunPullImagesCheck(execer utilsexec.Interface, cfg *kubeadmapi.InitConfiguration, ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	containerRuntime, err := utilruntime.NewContainerRuntime(utilsexec.New(), cfg.GetCRISocket())
	if err != nil {
		return err
	}
	checks := []Checker{ImagePullCheck{runtime: containerRuntime, imageList: images.GetAllImages(&cfg.ClusterConfiguration)}}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}
func RunChecks(checks []Checker, ww io.Writer, ignorePreflightErrors sets.String) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	type checkErrors struct {
		Name   string
		Errors []error
	}
	found := []checkErrors{}
	for _, c := range checks {
		name := c.Name()
		warnings, errs := c.Check()
		if setHasItemOrAll(ignorePreflightErrors, name) {
			warnings = append(warnings, errs...)
			errs = []error{}
		}
		for _, w := range warnings {
			io.WriteString(ww, fmt.Sprintf("\t[WARNING %s]: %v\n", name, w))
		}
		if len(errs) > 0 {
			found = append(found, checkErrors{Name: name, Errors: errs})
		}
	}
	if len(found) > 0 {
		var errs bytes.Buffer
		for _, c := range found {
			for _, i := range c.Errors {
				errs.WriteString(fmt.Sprintf("\t[ERROR %s]: %v\n", c.Name, i.Error()))
			}
		}
		return &Error{Msg: errs.String()}
	}
	return nil
}
func setHasItemOrAll(s sets.String, item string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if s.Has("all") || s.Has(strings.ToLower(item)) {
		return true
	}
	return false
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
