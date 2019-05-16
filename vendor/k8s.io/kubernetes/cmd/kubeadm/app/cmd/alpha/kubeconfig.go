package alpha

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	kubeconfigLongDesc = normalizer.LongDesc(`
	Kubeconfig file utilities.
	` + cmdutil.AlphaDisclaimer)
	userKubeconfigLongDesc = normalizer.LongDesc(`
	Outputs a kubeconfig file for an additional user.
	` + cmdutil.AlphaDisclaimer)
	userKubeconfigExample = normalizer.Examples(`
	# Outputs a kubeconfig file for an additional user named foo
	kubeadm alpha kubeconfig user --client-name=foo
	`)
)

func newCmdKubeConfigUtility(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "kubeconfig", Short: "Kubeconfig file utilities", Long: kubeconfigLongDesc}
	cmd.AddCommand(newCmdUserKubeConfig(out))
	return cmd
}
func newCmdUserKubeConfig(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cfg := &kubeadmapiv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(cfg)
	var token, clientName string
	var organizations []string
	cmd := &cobra.Command{Use: "user", Short: "Outputs a kubeconfig file for an additional user", Long: userKubeconfigLongDesc, Example: userKubeconfigExample, Run: func(cmd *cobra.Command, args []string) {
		if clientName == "" {
			kubeadmutil.CheckErr(errors.New("missing required argument --client-name"))
		}
		internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig("", cfg)
		kubeadmutil.CheckErr(err)
		if token != "" {
			kubeadmutil.CheckErr(kubeconfigphase.WriteKubeConfigWithToken(out, internalcfg, clientName, token))
			return
		}
		kubeadmutil.CheckErr(kubeconfigphase.WriteKubeConfigWithClientCert(out, internalcfg, clientName, organizations))
	}}
	cmd.Flags().StringVar(&cfg.CertificatesDir, "cert-dir", cfg.CertificatesDir, "The path where certificates are stored")
	cmd.Flags().StringVar(&cfg.LocalAPIEndpoint.AdvertiseAddress, "apiserver-advertise-address", cfg.LocalAPIEndpoint.AdvertiseAddress, "The IP address the API server is accessible on")
	cmd.Flags().Int32Var(&cfg.LocalAPIEndpoint.BindPort, "apiserver-bind-port", cfg.LocalAPIEndpoint.BindPort, "The port the API server is accessible on")
	cmd.Flags().StringVar(&token, "token", token, "The token that should be used as the authentication mechanism for this kubeconfig, instead of client certificates")
	cmd.Flags().StringVar(&clientName, "client-name", clientName, "The name of user. It will be used as the CN if client certificates are created")
	cmd.Flags().StringSliceVar(&organizations, "org", organizations, "The orgnizations of the client certificate. It will be used as the O if client certificates are created")
	return cmd
}
