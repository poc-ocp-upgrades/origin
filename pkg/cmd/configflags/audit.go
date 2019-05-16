package configflags

import (
	goformat "fmt"
	configv1 "github.com/openshift/api/config/v1"
	"io/ioutil"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"strconv"
	gotime "time"
)

const defaultAuditPolicyFilePath = "openshift.local.audit/policy.yaml"

func AuditFlags(c *configv1.AuditConfig, args map[string][]string) map[string][]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !c.Enabled {
		return args
	}
	auditPolicyFilePath := c.PolicyFile
	if len(c.PolicyConfiguration.Raw) > 0 && string(c.PolicyConfiguration.Raw) != "null" {
		if len(auditPolicyFilePath) == 0 {
			auditPolicyFilePath = defaultAuditPolicyFilePath
		}
		if err := os.MkdirAll(filepath.Dir(auditPolicyFilePath), 0755); err != nil {
			utilruntime.HandleError(err)
		}
		if err := ioutil.WriteFile(auditPolicyFilePath, c.PolicyConfiguration.Raw, 0644); err != nil {
			utilruntime.HandleError(err)
		}
	}
	SetIfUnset(args, "audit-log-maxbackup", strconv.Itoa(int(c.MaximumRetainedFiles)))
	SetIfUnset(args, "audit-log-maxsize", strconv.Itoa(int(c.MaximumFileSizeMegabytes)))
	SetIfUnset(args, "audit-log-maxage", strconv.Itoa(int(c.MaximumFileRetentionDays)))
	auditFilePath := c.AuditFilePath
	if len(auditFilePath) == 0 {
		auditFilePath = "-"
	}
	SetIfUnset(args, "audit-log-path", auditFilePath)
	if len(auditPolicyFilePath) > 0 {
		SetIfUnset(args, "audit-policy-file", auditPolicyFilePath)
	}
	if len(c.LogFormat) > 0 {
		SetIfUnset(args, "audit-log-format", string(c.LogFormat))
	}
	if len(c.WebHookMode) > 0 {
		SetIfUnset(args, "audit-webhook-mode", string(c.WebHookMode))
	}
	SetIfUnset(args, "audit-webhook-config-file", string(c.WebHookKubeConfig))
	return args
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
