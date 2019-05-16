package phases

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
	"k8s.io/kubernetes/pkg/util/normalizer"
	"strings"
)

var (
	saKeyLongDesc = fmt.Sprintf(normalizer.LongDesc(`
		Generates the private key for signing service account tokens along with its public key, and saves them into
		%s and %s files.
		If both files already exist, kubeadm skips the generation step and existing files will be used.
		`+cmdutil.AlphaDisclaimer), kubeadmconstants.ServiceAccountPrivateKeyName, kubeadmconstants.ServiceAccountPublicKeyName)
	genericLongDesc = normalizer.LongDesc(`
		Generates the %[1]s, and saves them into %[2]s.cert and %[2]s.key files.%[3]s

		If both files already exist, kubeadm skips the generation step and existing files will be used.
		` + cmdutil.AlphaDisclaimer)
)
var (
	csrOnly bool
	csrDir  string
)

type certsData interface {
	Cfg() *kubeadmapi.InitConfiguration
	ExternalCA() bool
	CertificateDir() string
	CertificateWriteDir() string
}

func NewCertsPhase() workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return workflow.Phase{Name: "certs", Short: "Certificate generation", Phases: newCertSubPhases(), Run: runCerts, Long: cmdutil.MacroCommandLongDescription}
}
func localFlags() *pflag.FlagSet {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	set := pflag.NewFlagSet("csr", pflag.ExitOnError)
	options.AddCSRFlag(set, &csrOnly)
	options.AddCSRDirFlag(set, &csrDir)
	return set
}
func newCertSubPhases() []workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subPhases := []workflow.Phase{}
	allPhase := workflow.Phase{Name: "all", Short: "Generates all certificates", InheritFlags: getCertPhaseFlags("all"), RunAllSiblings: true}
	subPhases = append(subPhases, allPhase)
	certTree, _ := certsphase.GetDefaultCertList().AsMap().CertTree()
	for ca, certList := range certTree {
		caPhase := newCertSubPhase(ca, runCAPhase(ca))
		subPhases = append(subPhases, caPhase)
		for _, cert := range certList {
			certPhase := newCertSubPhase(cert, runCertPhase(cert, ca))
			certPhase.LocalFlags = localFlags()
			subPhases = append(subPhases, certPhase)
		}
	}
	saPhase := workflow.Phase{Name: "sa", Short: "Generates a private key for signing service account tokens along with its public key", Long: saKeyLongDesc, Run: runCertsSa, InheritFlags: []string{options.CertificatesDir}}
	subPhases = append(subPhases, saPhase)
	return subPhases
}
func newCertSubPhase(certSpec *certsphase.KubeadmCert, run func(c workflow.RunData) error) workflow.Phase {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	phase := workflow.Phase{Name: certSpec.Name, Short: fmt.Sprintf("Generates the %s", certSpec.LongName), Long: fmt.Sprintf(genericLongDesc, certSpec.LongName, certSpec.BaseName, getSANDescription(certSpec)), Run: run, InheritFlags: getCertPhaseFlags(certSpec.Name)}
	return phase
}
func getCertPhaseFlags(name string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags := []string{options.CertificatesDir, options.CfgPath, options.CSROnly, options.CSRDir}
	if name == "all" || name == "apiserver" {
		flags = append(flags, options.APIServerAdvertiseAddress, options.APIServerCertSANs, options.NetworkingDNSDomain, options.NetworkingServiceSubnet)
	}
	return flags
}
func getSANDescription(certSpec *certsphase.KubeadmCert) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	defaultConfig := &kubeadmapiv1beta1.InitConfiguration{LocalAPIEndpoint: kubeadmapiv1beta1.APIEndpoint{AdvertiseAddress: "127.0.0.1"}}
	defaultInternalConfig := &kubeadmapi.InitConfiguration{}
	kubeadmscheme.Scheme.Default(defaultConfig)
	err := kubeadmscheme.Scheme.Convert(defaultConfig, defaultInternalConfig, nil)
	kubeadmutil.CheckErr(err)
	certConfig, err := certSpec.GetConfig(defaultInternalConfig)
	kubeadmutil.CheckErr(err)
	if len(certConfig.AltNames.DNSNames) == 0 && len(certConfig.AltNames.IPs) == 0 {
		return ""
	}
	sans := []string{}
	for _, dnsName := range certConfig.AltNames.DNSNames {
		if dnsName != "" {
			sans = append(sans, dnsName)
		}
	}
	for _, ip := range certConfig.AltNames.IPs {
		sans = append(sans, ip.String())
	}
	return fmt.Sprintf("\n\nDefault SANs are %s", strings.Join(sans, ", "))
}
func runCertsSa(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(certsData)
	if !ok {
		return errors.New("certs phase invoked with an invalid data struct")
	}
	if data.ExternalCA() {
		fmt.Printf("[certs] External CA mode: Using existing sa keys\n")
		return nil
	}
	cfg := data.Cfg()
	cfg.CertificatesDir = data.CertificateWriteDir()
	defer func() {
		cfg.CertificatesDir = data.CertificateDir()
	}()
	return certsphase.CreateServiceAccountKeyAndPublicKeyFiles(cfg)
}
func runCerts(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := c.(certsData)
	if !ok {
		return errors.New("certs phase invoked with an invalid data struct")
	}
	fmt.Printf("[certs] Using certificateDir folder %q\n", data.CertificateWriteDir())
	return nil
}
func runCAPhase(ca *certsphase.KubeadmCert) func(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(c workflow.RunData) error {
		data, ok := c.(certsData)
		if !ok {
			return errors.New("certs phase invoked with an invalid data struct")
		}
		if _, err := pkiutil.TryLoadCertFromDisk(data.CertificateDir(), ca.BaseName); err == nil {
			if _, err := pkiutil.TryLoadKeyFromDisk(data.CertificateDir(), ca.BaseName); err == nil {
				fmt.Printf("[certs] Using existing %s certificate authority\n", ca.BaseName)
				return nil
			}
			fmt.Printf("[certs] Using existing %s keyless certificate authority", ca.BaseName)
			return nil
		}
		if data.Cfg().Etcd.External != nil && ca.Name == "etcd-ca" {
			fmt.Printf("[certs] External etcd mode: Skipping %s certificate authority generation\n", ca.BaseName)
			return nil
		}
		cfg := data.Cfg()
		cfg.CertificatesDir = data.CertificateWriteDir()
		defer func() {
			cfg.CertificatesDir = data.CertificateDir()
		}()
		return certsphase.CreateCACertAndKeyFiles(ca, cfg)
	}
}
func runCertPhase(cert *certsphase.KubeadmCert, caCert *certsphase.KubeadmCert) func(c workflow.RunData) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(c workflow.RunData) error {
		data, ok := c.(certsData)
		if !ok {
			return errors.New("certs phase invoked with an invalid data struct")
		}
		if certData, _, err := pkiutil.TryLoadCertAndKeyFromDisk(data.CertificateDir(), cert.BaseName); err == nil {
			caCertData, err := pkiutil.TryLoadCertFromDisk(data.CertificateDir(), caCert.BaseName)
			if err != nil {
				return errors.Wrapf(err, "couldn't load CA certificate %s", caCert.Name)
			}
			if err := certData.CheckSignatureFrom(caCertData); err != nil {
				return errors.Wrapf(err, "[certs] certificate %s not signed by CA certificate %s", cert.BaseName, caCert.BaseName)
			}
			fmt.Printf("[certs] Using existing %s certificate and key on disk\n", cert.BaseName)
			return nil
		}
		if csrOnly {
			fmt.Printf("[certs] Generating CSR for %s instead of certificate\n", cert.BaseName)
			if csrDir == "" {
				csrDir = data.CertificateWriteDir()
			}
			return certsphase.CreateCSR(cert, data.Cfg(), csrDir)
		}
		if data.Cfg().Etcd.External != nil && cert.CAName == "etcd-ca" {
			fmt.Printf("[certs] External etcd mode: Skipping %s certificate authority generation\n", cert.BaseName)
			return nil
		}
		cfg := data.Cfg()
		cfg.CertificatesDir = data.CertificateWriteDir()
		defer func() {
			cfg.CertificatesDir = data.CertificateDir()
		}()
		return certsphase.CreateCertAndKeyFilesWithCA(cert, caCert, cfg)
	}
}
