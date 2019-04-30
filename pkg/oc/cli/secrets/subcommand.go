package secrets

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

const SecretsRecommendedName = "secrets"
const (
	SourceUsername		= "username"
	SourcePassword		= "password"
	SourceCertificate	= "ca.crt"
	SourcePrivateKey	= "ssh-privatekey"
	SourceGitConfig		= ".gitconfig"
)

var (
	secretsLong = templates.LongDesc(`
    Manage secrets in your project

    Secrets are used to store confidential information that should not be contained inside of an image.
    They are commonly used to hold things like keys for authentication to other internal systems like
    Docker registries.`)
)

func NewCmdSecrets(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmds := &cobra.Command{Use: name, Short: "Manage secrets", Long: secretsLong, Aliases: []string{"secret"}, Run: kcmdutil.DefaultSubCommandRun(streams.ErrOut)}
	cmds.AddCommand(NewCmdLinkSecret(LinkSecretRecommendedName, fullName+" "+LinkSecretRecommendedName, f, streams))
	cmds.AddCommand(NewCmdUnlinkSecret(UnlinkSecretRecommendedName, fullName+" "+UnlinkSecretRecommendedName, f, streams))
	return cmds
}
