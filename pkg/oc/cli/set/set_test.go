package set

import (
	"testing"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdtesting "k8s.io/kubernetes/pkg/kubectl/cmd/testing"
)

func TestLocalAndDryRunFlags(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tf := kcmdtesting.NewTestFactory().WithNamespace("test")
	defer tf.Cleanup()
	setCmd := NewCmdSet("", tf, genericclioptions.NewTestIOStreamsDiscard())
	ensureLocalAndDryRunFlagsOnChildren(t, setCmd, "")
}
func ensureLocalAndDryRunFlagsOnChildren(t *testing.T, c *cobra.Command, prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, cmd := range c.Commands() {
		name := prefix + cmd.Name()
		if localFlag := cmd.Flag("local"); localFlag == nil {
			t.Errorf("Command %s does not implement the --local flag", name)
		}
		if dryRunFlag := cmd.Flag("dry-run"); dryRunFlag == nil {
			t.Errorf("Command %s does not implement the --dry-run flag", name)
		}
		ensureLocalAndDryRunFlagsOnChildren(t, cmd, name+".")
	}
}
