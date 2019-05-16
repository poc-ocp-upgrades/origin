package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"k8s.io/klog"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/version"
	"sigs.k8s.io/yaml"
)

type Version struct {
	ClientVersion *apimachineryversion.Info `json:"clientVersion"`
}

func NewCmdVersion(out io.Writer) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := &cobra.Command{Use: "version", Short: "Print the version of kubeadm", Run: func(cmd *cobra.Command, args []string) {
		err := RunVersion(out, cmd)
		kubeadmutil.CheckErr(err)
	}}
	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}
func RunVersion(out io.Writer, cmd *cobra.Command) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	klog.V(1).Infoln("[version] retrieving version info")
	clientVersion := version.Get()
	v := Version{ClientVersion: &clientVersion}
	const flag = "output"
	of, err := cmd.Flags().GetString(flag)
	if err != nil {
		klog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	switch of {
	case "":
		fmt.Fprintf(out, "kubeadm version: %#v\n", v.ClientVersion)
	case "short":
		fmt.Fprintf(out, "%s\n", v.ClientVersion.GitVersion)
	case "yaml":
		y, err := yaml.Marshal(&v)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	case "json":
		y, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	default:
		return errors.Errorf("invalid output format: %s", of)
	}
	return nil
}
