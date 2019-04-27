package validation

import (
	"testing"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/validation/field"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
)

func TestFailingKubeletArgs(t *testing.T) {
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
	args := configapi.ExtendedArguments{}
	args["port"] = []string{"invalid-value"}
	args["missing-key"] = []string{"value"}
	errs := ValidateKubeletExtendedArguments(args, nil)
	if len(errs) != 2 {
		t.Fatalf("expected 2 errors, not %v", errs)
	}
	var (
		portErr		*field.Error
		missingErr	*field.Error
	)
	for _, err := range errs {
		switch err.Field {
		case "port":
			portErr = err
		case "flag":
			missingErr = err
		}
	}
	if portErr == nil {
		t.Fatalf("missing port")
	}
	if missingErr == nil {
		t.Fatalf("missing missing-key")
	}
	if e, a := "port", portErr.Field; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := "invalid-value", portErr.BadValue.(string); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := `could not be set: strconv.ParseInt: parsing "invalid-value": invalid syntax`, portErr.Detail; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := "flag", missingErr.Field; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := "missing-key", missingErr.BadValue.(string); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := `is not a valid flag`, missingErr.Detail; e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
}
func TestInvalidProjectEmptyDirQuota(t *testing.T) {
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
	negQuota := resource.MustParse("-1000Mi")
	nodeCfg := configapi.NodeConfig{VolumeConfig: configapi.NodeVolumeConfig{LocalQuota: configapi.LocalQuota{PerFSGroup: &negQuota}}}
	errs := ValidateNodeConfig(&nodeCfg, nil)
	var emptyDirQuotaError *field.Error
	for _, err := range errs.Errors {
		t.Logf("Found error: %s", err.Field)
		if err.Field == "volumeConfig.localQuota.perFSGroup" {
			emptyDirQuotaError = err
		}
	}
	if emptyDirQuotaError == nil {
		t.Fatalf("expected volumeConfig.localQuota.perFSGroup error but got none")
	}
	if emptyDirQuotaError.Type != field.ErrorTypeInvalid {
		t.Errorf("unexpected error for negative volumeConfig.localQuota.perFSGroup: %s", emptyDirQuotaError.Detail)
	}
}
