package rsync

import (
	"testing"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var rshAllowedFlags = sets.NewString("container")

func TestRshExcludeFlags(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rsyncCmd := NewCmdRsync("rsync", "oc", nil, genericclioptions.NewTestIOStreamsDiscard())
	rsyncCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if !rshExcludeFlags.Has(flag.Name) && !rshAllowedFlags.Has(flag.Name) {
			t.Errorf("Unknown flag %s. Please add to rshExcludeFlags or to rshAllowedFlags", flag.Name)
		}
	})
}
func TestRsyncEscapeCommand(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stringsToEscape := []string{`thisshouldnotgetescapedorquoted`, `this should get quoted for spaces`, `this" should get escaped and quoted`, `"this should get escaped and quoted"`, `this\ should get quoted`, `this' should get quoted`}
	stringsShouldMatch := []string{`thisshouldnotgetescapedorquoted`, `"this should get quoted for spaces"`, `"this"" should get escaped and quoted"`, `"""this should get escaped and quoted"""`, `"this\ should get quoted"`, `"this' should get quoted"`}
	escapedStrings := rsyncEscapeCommand(stringsToEscape)
	for key, val := range escapedStrings {
		if val != stringsShouldMatch[key] {
			t.Errorf("%v did not match %v", val, stringsShouldMatch[key])
		}
	}
}
