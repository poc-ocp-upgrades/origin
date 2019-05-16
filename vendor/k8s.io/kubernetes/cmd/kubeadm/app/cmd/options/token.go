package options

import (
	"fmt"
	"github.com/spf13/pflag"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	kubeadmapiv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"strings"
)

func NewBootstrapTokenOptions() *BootstrapTokenOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bto := &BootstrapTokenOptions{&kubeadmapiv1beta1.BootstrapToken{}, ""}
	kubeadmapiv1beta1.SetDefaults_BootstrapToken(bto.BootstrapToken)
	return bto
}

type BootstrapTokenOptions struct {
	*kubeadmapiv1beta1.BootstrapToken
	TokenStr string
}

func (bto *BootstrapTokenOptions) AddTokenFlag(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(&bto.TokenStr, "token", "", "The token to use for establishing bidirectional trust between nodes and masters. The format is [a-z0-9]{6}\\.[a-z0-9]{16} - e.g. abcdef.0123456789abcdef")
}
func (bto *BootstrapTokenOptions) AddTTLFlag(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	bto.AddTTLFlagWithName(fs, "token-ttl")
}
func (bto *BootstrapTokenOptions) AddTTLFlagWithName(fs *pflag.FlagSet, flagName string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.DurationVar(&bto.TTL.Duration, flagName, bto.TTL.Duration, "The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire")
}
func (bto *BootstrapTokenOptions) AddUsagesFlag(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringSliceVar(&bto.Usages, "usages", bto.Usages, fmt.Sprintf("Describes the ways in which this token can be used. You can pass --usages multiple times or provide a comma separated list of options. Valid options: [%s]", strings.Join(kubeadmconstants.DefaultTokenUsages, ",")))
}
func (bto *BootstrapTokenOptions) AddGroupsFlag(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringSliceVar(&bto.Groups, "groups", bto.Groups, fmt.Sprintf("Extra groups that this token will authenticate as when used for authentication. Must match %q", bootstrapapi.BootstrapGroupPattern))
}
func (bto *BootstrapTokenOptions) AddDescriptionFlag(fs *pflag.FlagSet) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fs.StringVar(&bto.Description, "description", bto.Description, "A human friendly description of how this token is used.")
}
func (bto *BootstrapTokenOptions) ApplyTo(cfg *kubeadmapiv1beta1.InitConfiguration) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(bto.TokenStr) > 0 {
		var err error
		bto.Token, err = kubeadmapiv1beta1.NewBootstrapTokenString(bto.TokenStr)
		if err != nil {
			return err
		}
	}
	cfg.BootstrapTokens = []kubeadmapiv1beta1.BootstrapToken{*bto.BootstrapToken}
	return nil
}
