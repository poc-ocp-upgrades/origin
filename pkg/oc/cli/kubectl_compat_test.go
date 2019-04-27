package cli

import (
	"bytes"
	"io/ioutil"
	"testing"
	"k8s.io/apimachinery/pkg/util/sets"
	kcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
)

var MissingCommands = sets.NewString("namespace", "rolling-update", "cordon", "drain", "uncordon", "taint", "top", "certificate", "run-container", "alpha")
var WhitelistedCommands = sets.NewString()

func TestKubectlCompatibility(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	oc := NewOcCommand("oc", "oc", &bytes.Buffer{}, ioutil.Discard, ioutil.Discard)
	kubectl := kcmd.NewKubectlCommand(nil, ioutil.Discard, ioutil.Discard)
kubectlLoop:
	for _, kubecmd := range kubectl.Commands() {
		for _, occmd := range oc.Commands() {
			if kubecmd.Name() == occmd.Name() {
				if MissingCommands.Has(kubecmd.Name()) {
					t.Errorf("%s was supposed to be missing", kubecmd.Name())
					continue
				}
				if WhitelistedCommands.Has(kubecmd.Name()) {
					t.Errorf("%s was supposed to be whitelisted", kubecmd.Name())
					continue
				}
				continue kubectlLoop
			}
		}
		if MissingCommands.Has(kubecmd.Name()) || WhitelistedCommands.Has(kubecmd.Name()) {
			continue
		}
		t.Errorf("missing %q,", kubecmd.Name())
	}
}
func TestValidateDisabled(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	oc := NewOcCommand("oc", "oc", &bytes.Buffer{}, ioutil.Discard, ioutil.Discard)
	kubectl := kcmd.NewKubectlCommand(nil, ioutil.Discard, ioutil.Discard)
	for _, kubecmd := range kubectl.Commands() {
		for _, occmd := range oc.Commands() {
			if kubecmd.Name() == occmd.Name() {
				ocValidateFlag := occmd.Flags().Lookup("validate")
				if ocValidateFlag == nil {
					continue
				}
				if ocValidateFlag.Value.String() != "false" {
					t.Errorf("%s --validate is not defaulting to false", occmd.Name())
				}
			}
		}
	}
}
