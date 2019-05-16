package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/duration"
	clientset "k8s.io/client-go/kubernetes"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/options"
	phaseutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases"
	cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	tokenphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func NewCmdToken(out io.Writer, errW io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	kubeConfigFile := kubeadmconstants.GetAdminKubeConfigPath()
	var dryRun bool
	tokenCmd := &cobra.Command{Use: "token", Short: "Manage bootstrap tokens.", Long: dedent.Dedent(`
			This command manages bootstrap tokens. It is optional and needed only for advanced use cases.

			In short, bootstrap tokens are used for establishing bidirectional trust between a client and a server.
			A bootstrap token can be used when a client (for example a node that is about to join the cluster) needs
			to trust the server it is talking to. Then a bootstrap token with the "signing" usage can be used.
			bootstrap tokens can also function as a way to allow short-lived authentication to the API Server
			(the token serves as a way for the API Server to trust the client), for example for doing the TLS Bootstrap.

			What is a bootstrap token more exactly?
			 - It is a Secret in the kube-system namespace of type "bootstrap.kubernetes.io/token".
			 - A bootstrap token must be of the form "[a-z0-9]{6}.[a-z0-9]{16}". The former part is the public token ID,
			   while the latter is the Token Secret and it must be kept private at all circumstances!
			 - The name of the Secret must be named "bootstrap-token-(token-id)".

			You can read more about bootstrap tokens here:
			  https://kubernetes.io/docs/admin/bootstrap-tokens/
		`), RunE: cmdutil.SubCmdRunE("token")}
	options.AddKubeConfigFlag(tokenCmd.PersistentFlags(), &kubeConfigFile)
	tokenCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", dryRun, "Whether to enable dry-run mode or not")
	cfg := &kubeadmapiv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(cfg)
	var cfgPath string
	var printJoinCommand bool
	bto := options.NewBootstrapTokenOptions()
	createCmd := &cobra.Command{Use: "create [token]", DisableFlagsInUseLine: true, Short: "Create bootstrap tokens on the server.", Long: dedent.Dedent(`
			This command will create a bootstrap token for you.
			You can specify the usages for this token, the "time to live" and an optional human friendly description.

			The [token] is the actual token to write.
			This should be a securely generated random token of the form "[a-z0-9]{6}.[a-z0-9]{16}".
			If no [token] is given, kubeadm will generate a random token instead.
		`), Run: func(tokenCmd *cobra.Command, args []string) {
		if len(args) > 0 {
			bto.TokenStr = args[0]
		}
		klog.V(1).Infoln("[token] validating mixed arguments")
		err := validation.ValidateMixedArguments(tokenCmd.Flags())
		kubeadmutil.CheckErr(err)
		err = bto.ApplyTo(cfg)
		kubeadmutil.CheckErr(err)
		klog.V(1).Infoln("[token] getting Clientsets from kubeconfig file")
		kubeConfigFile = cmdutil.FindExistingKubeConfig(kubeConfigFile)
		client, err := getClientset(kubeConfigFile, dryRun)
		kubeadmutil.CheckErr(err)
		err = RunCreateToken(out, client, cfgPath, cfg, printJoinCommand, kubeConfigFile)
		kubeadmutil.CheckErr(err)
	}}
	createCmd.Flags().StringVar(&cfgPath, "config", cfgPath, "Path to kubeadm config file (WARNING: Usage of a configuration file is experimental)")
	createCmd.Flags().BoolVar(&printJoinCommand, "print-join-command", false, "Instead of printing only the token, print the full 'kubeadm join' flag needed to join the cluster using the token.")
	bto.AddTTLFlagWithName(createCmd.Flags(), "ttl")
	bto.AddUsagesFlag(createCmd.Flags())
	bto.AddGroupsFlag(createCmd.Flags())
	bto.AddDescriptionFlag(createCmd.Flags())
	tokenCmd.AddCommand(createCmd)
	tokenCmd.AddCommand(NewCmdTokenGenerate(out))
	listCmd := &cobra.Command{Use: "list", Short: "List bootstrap tokens on the server.", Long: dedent.Dedent(`
			This command will list all bootstrap tokens for you.
		`), Run: func(tokenCmd *cobra.Command, args []string) {
		kubeConfigFile = cmdutil.FindExistingKubeConfig(kubeConfigFile)
		client, err := getClientset(kubeConfigFile, dryRun)
		kubeadmutil.CheckErr(err)
		err = RunListTokens(out, errW, client)
		kubeadmutil.CheckErr(err)
	}}
	tokenCmd.AddCommand(listCmd)
	deleteCmd := &cobra.Command{Use: "delete [token-value]", DisableFlagsInUseLine: true, Short: "Delete bootstrap tokens on the server.", Long: dedent.Dedent(`
			This command will delete a given bootstrap token for you.

			The [token-value] is the full Token of the form "[a-z0-9]{6}.[a-z0-9]{16}" or the
			Token ID of the form "[a-z0-9]{6}" to delete.
		`), Run: func(tokenCmd *cobra.Command, args []string) {
		if len(args) < 1 {
			kubeadmutil.CheckErr(errors.Errorf("missing subcommand; 'token delete' is missing token of form %q", bootstrapapi.BootstrapTokenIDPattern))
		}
		kubeConfigFile = cmdutil.FindExistingKubeConfig(kubeConfigFile)
		client, err := getClientset(kubeConfigFile, dryRun)
		kubeadmutil.CheckErr(err)
		err = RunDeleteToken(out, client, args[0])
		kubeadmutil.CheckErr(err)
	}}
	tokenCmd.AddCommand(deleteCmd)
	return tokenCmd
}
func NewCmdTokenGenerate(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &cobra.Command{Use: "generate", Short: "Generate and print a bootstrap token, but do not create it on the server.", Long: dedent.Dedent(`
			This command will print out a randomly-generated bootstrap token that can be used with
			the "init" and "join" commands.

			You don't have to use this command in order to generate a token. You can do so
			yourself as long as it is in the format "[a-z0-9]{6}.[a-z0-9]{16}". This
			command is provided for convenience to generate tokens in the given format.

			You can also use "kubeadm init" without specifying a token and it will
			generate and print one for you.
		`), Run: func(cmd *cobra.Command, args []string) {
		err := RunGenerateToken(out)
		kubeadmutil.CheckErr(err)
	}}
}
func RunCreateToken(out io.Writer, client clientset.Interface, cfgPath string, cfg *kubeadmapiv1beta1.InitConfiguration, printJoinCommand bool, kubeConfigFile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	phaseutil.SetKubernetesVersion(cfg)
	klog.V(1).Infoln("[token] loading configurations")
	internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(cfgPath, cfg)
	if err != nil {
		return err
	}
	klog.V(1).Infoln("[token] creating token")
	if err := tokenphase.CreateNewTokens(client, internalcfg.BootstrapTokens); err != nil {
		return err
	}
	if printJoinCommand {
		joinCommand, err := cmdutil.GetJoinCommand(kubeConfigFile, internalcfg.BootstrapTokens[0].Token.String(), false)
		if err != nil {
			return errors.Wrap(err, "failed to get join command")
		}
		fmt.Fprintln(out, joinCommand)
	} else {
		fmt.Fprintln(out, internalcfg.BootstrapTokens[0].Token.String())
	}
	return nil
}
func RunGenerateToken(out io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("[token] generating random token")
	token, err := bootstraputil.GenerateBootstrapToken()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, token)
	return nil
}
func RunListTokens(out io.Writer, errW io.Writer, client clientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("[token] preparing selector for bootstrap token")
	tokenSelector := fields.SelectorFromSet(map[string]string{"type": string(bootstrapapi.SecretTypeBootstrapToken)})
	listOptions := metav1.ListOptions{FieldSelector: tokenSelector.String()}
	klog.V(1).Infoln("[token] retrieving list of bootstrap tokens")
	secrets, err := client.CoreV1().Secrets(metav1.NamespaceSystem).List(listOptions)
	if err != nil {
		return errors.Wrap(err, "failed to list bootstrap tokens")
	}
	w := tabwriter.NewWriter(out, 10, 4, 3, ' ', 0)
	fmt.Fprintln(w, "TOKEN\tTTL\tEXPIRES\tUSAGES\tDESCRIPTION\tEXTRA GROUPS")
	for _, secret := range secrets.Items {
		token, err := kubeadmapi.BootstrapTokenFromSecret(&secret)
		if err != nil {
			fmt.Fprintf(errW, "%v", err)
			continue
		}
		humanFriendlyTokenOutput := humanReadableBootstrapToken(token)
		fmt.Fprintln(w, humanFriendlyTokenOutput)
	}
	w.Flush()
	return nil
}
func RunDeleteToken(out io.Writer, client clientset.Interface, tokenIDOrToken string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokenID := tokenIDOrToken
	klog.V(1).Infoln("[token] parsing token ID")
	if !bootstraputil.IsValidBootstrapTokenID(tokenIDOrToken) {
		bts, err := kubeadmapiv1beta1.NewBootstrapTokenString(tokenIDOrToken)
		if err != nil {
			return errors.Errorf("given token or token id %q didn't match pattern %q or %q", tokenIDOrToken, bootstrapapi.BootstrapTokenIDPattern, bootstrapapi.BootstrapTokenIDPattern)
		}
		tokenID = bts.ID
	}
	tokenSecretName := bootstraputil.BootstrapTokenSecretName(tokenID)
	klog.V(1).Infoln("[token] deleting token")
	if err := client.CoreV1().Secrets(metav1.NamespaceSystem).Delete(tokenSecretName, nil); err != nil {
		return errors.Wrap(err, "failed to delete bootstrap token")
	}
	fmt.Fprintf(out, "bootstrap token with id %q deleted\n", tokenID)
	return nil
}
func humanReadableBootstrapToken(token *kubeadmapi.BootstrapToken) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	description := token.Description
	if len(description) == 0 {
		description = "<none>"
	}
	ttl := "<forever>"
	expires := "<never>"
	if token.Expires != nil {
		ttl = duration.ShortHumanDuration(token.Expires.Sub(time.Now()))
		expires = token.Expires.Format(time.RFC3339)
	}
	usagesString := strings.Join(token.Usages, ",")
	if len(usagesString) == 0 {
		usagesString = "<none>"
	}
	groupsString := strings.Join(token.Groups, ",")
	if len(groupsString) == 0 {
		groupsString = "<none>"
	}
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s", token.Token.String(), ttl, expires, usagesString, description, groupsString)
}
func getClientset(file string, dryRun bool) (clientset.Interface, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if dryRun {
		dryRunGetter, err := apiclient.NewClientBackedDryRunGetterFromKubeconfig(file)
		if err != nil {
			return nil, err
		}
		return apiclient.NewDryRunClient(dryRunGetter, os.Stdout), nil
	}
	return kubeconfigutil.ClientSetFromFile(file)
}
