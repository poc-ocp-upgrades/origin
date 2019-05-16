package imagepolicy

import (
	"fmt"
	"k8s.io/klog"
	"time"
)

const (
	defaultRetryBackoff = time.Duration(500) * time.Millisecond
	minRetryBackoff     = time.Duration(1)
	maxRetryBackoff     = time.Duration(5) * time.Minute
	defaultAllowTTL     = time.Duration(5) * time.Minute
	defaultDenyTTL      = time.Duration(30) * time.Second
	minAllowTTL         = time.Duration(1) * time.Second
	maxAllowTTL         = time.Duration(30) * time.Minute
	minDenyTTL          = time.Duration(1) * time.Second
	maxDenyTTL          = time.Duration(30) * time.Minute
	useDefault          = time.Duration(0)
	disableTTL          = time.Duration(-1)
)

type imagePolicyWebhookConfig struct {
	KubeConfigFile string        `json:"kubeConfigFile"`
	AllowTTL       time.Duration `json:"allowTTL"`
	DenyTTL        time.Duration `json:"denyTTL"`
	RetryBackoff   time.Duration `json:"retryBackoff"`
	DefaultAllow   bool          `json:"defaultAllow"`
}
type AdmissionConfig struct {
	ImagePolicyWebhook imagePolicyWebhookConfig `json:"imagePolicy"`
}

func normalizeWebhookConfig(config *imagePolicyWebhookConfig) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config.RetryBackoff, err = normalizeConfigDuration("backoff", time.Millisecond, config.RetryBackoff, minRetryBackoff, maxRetryBackoff, defaultRetryBackoff)
	if err != nil {
		return err
	}
	config.AllowTTL, err = normalizeConfigDuration("allow cache", time.Second, config.AllowTTL, minAllowTTL, maxAllowTTL, defaultAllowTTL)
	if err != nil {
		return err
	}
	config.DenyTTL, err = normalizeConfigDuration("deny cache", time.Second, config.DenyTTL, minDenyTTL, maxDenyTTL, defaultDenyTTL)
	if err != nil {
		return err
	}
	return nil
}
func normalizeConfigDuration(name string, scale, value, min, max, defaultValue time.Duration) (time.Duration, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if value == disableTTL {
		klog.V(2).Infof("image policy webhook %s disabled", name)
		return time.Duration(0), nil
	}
	if value == useDefault {
		klog.V(2).Infof("image policy webhook %s using default value", name)
		return defaultValue, nil
	}
	value *= scale
	if value < min || value > max {
		return value, fmt.Errorf("valid value is between %v and %v, got %v", min, max, value)
	}
	return value, nil
}
