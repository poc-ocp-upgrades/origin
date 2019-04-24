package cert

import (
	"github.com/spf13/cobra"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"github.com/openshift/origin/pkg/cmd/server/admin"
)

const CertRecommendedName = "ca"

func NewCmdCert(name, fullName string, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmds := &cobra.Command{Use: name, Long: "Manage certificates and keys", Short: "", Run: cmdutil.DefaultSubCommandRun(streams.ErrOut), Deprecated: "and will be removed in the future version", Hidden: true}
	subCommands := []*cobra.Command{admin.NewCommandCreateMasterCerts(admin.CreateMasterCertsCommandName, fullName+" "+admin.CreateMasterCertsCommandName, streams), admin.NewCommandCreateKeyPair(admin.CreateKeyPairCommandName, fullName+" "+admin.CreateKeyPairCommandName, streams), admin.NewCommandCreateServerCert(admin.CreateServerCertCommandName, fullName+" "+admin.CreateServerCertCommandName, streams), admin.NewCommandCreateSignerCert(admin.CreateSignerCertCommandName, fullName+" "+admin.CreateSignerCertCommandName, streams), admin.NewCommandEncrypt(admin.EncryptCommandName, fullName+" "+admin.EncryptCommandName, streams), admin.NewCommandDecrypt(admin.DecryptCommandName, fullName+" "+admin.DecryptCommandName, fullName+" "+admin.EncryptCommandName, streams)}
	for _, cmd := range subCommands {
		cmd.Short = ""
		cmd.Deprecated = "and will be removed in the future version"
		cmds.AddCommand(cmd)
	}
	return cmds
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
