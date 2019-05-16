package audit

import (
	goformat "fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/apis/audit/install"
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"
	"k8s.io/kubernetes/cmd/kubeadm/app/util"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func CreateDefaultAuditLogPolicy(policyFile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	policy := auditv1.Policy{TypeMeta: metav1.TypeMeta{APIVersion: auditv1.SchemeGroupVersion.String(), Kind: "Policy"}, Rules: []auditv1.PolicyRule{{Level: auditv1.LevelMetadata}}}
	return writePolicyToDisk(policyFile, &policy)
}
func writePolicyToDisk(policyFile string, policy *auditv1.Policy) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := os.MkdirAll(filepath.Dir(policyFile), 0700); err != nil {
		return errors.Wrapf(err, "failed to create directory %q: ", filepath.Dir(policyFile))
	}
	scheme := runtime.NewScheme()
	install.Install(scheme)
	codecs := serializer.NewCodecFactory(scheme)
	serialized, err := util.MarshalToYamlForCodecs(policy, auditv1.SchemeGroupVersion, codecs)
	if err != nil {
		return errors.Wrap(err, "failed to marshal audit policy to YAML")
	}
	if err := ioutil.WriteFile(policyFile, serialized, 0600); err != nil {
		return errors.Wrapf(err, "failed to write audit policy to %v: ", policyFile)
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
