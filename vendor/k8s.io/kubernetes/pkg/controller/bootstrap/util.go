package bootstrap

import (
	"k8s.io/api/core/v1"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	"k8s.io/klog"
	"regexp"
	"time"
)

var namePattern = `^` + regexp.QuoteMeta(bootstrapapi.BootstrapTokenSecretPrefix) + `([a-z0-9]{6})$`
var nameRegExp = regexp.MustCompile(namePattern)

func getSecretString(secret *v1.Secret, key string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, ok := secret.Data[key]
	if !ok {
		return ""
	}
	return string(data)
}
func parseSecretName(name string) (secretID string, ok bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r := nameRegExp.FindStringSubmatch(name)
	if r == nil {
		return "", false
	}
	return r[1], true
}
func validateSecretForSigning(secret *v1.Secret) (tokenID, tokenSecret string, ok bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nameTokenID, ok := parseSecretName(secret.Name)
	if !ok {
		klog.V(3).Infof("Invalid secret name: %s. Must be of form %s<secret-id>.", secret.Name, bootstrapapi.BootstrapTokenSecretPrefix)
		return "", "", false
	}
	tokenID = getSecretString(secret, bootstrapapi.BootstrapTokenIDKey)
	if len(tokenID) == 0 {
		klog.V(3).Infof("No %s key in %s/%s Secret", bootstrapapi.BootstrapTokenIDKey, secret.Namespace, secret.Name)
		return "", "", false
	}
	if nameTokenID != tokenID {
		klog.V(3).Infof("Token ID (%s) doesn't match secret name: %s", tokenID, nameTokenID)
		return "", "", false
	}
	tokenSecret = getSecretString(secret, bootstrapapi.BootstrapTokenSecretKey)
	if len(tokenSecret) == 0 {
		klog.V(3).Infof("No %s key in %s/%s Secret", bootstrapapi.BootstrapTokenSecretKey, secret.Namespace, secret.Name)
		return "", "", false
	}
	if isSecretExpired(secret) {
		return "", "", false
	}
	okToSign := getSecretString(secret, bootstrapapi.BootstrapTokenUsageSigningKey)
	if okToSign != "true" {
		return "", "", false
	}
	return tokenID, tokenSecret, true
}
func isSecretExpired(secret *v1.Secret) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	expiration := getSecretString(secret, bootstrapapi.BootstrapTokenExpirationKey)
	if len(expiration) > 0 {
		expTime, err2 := time.Parse(time.RFC3339, expiration)
		if err2 != nil {
			klog.V(3).Infof("Unparseable expiration time (%s) in %s/%s Secret: %v. Treating as expired.", expiration, secret.Namespace, secret.Name, err2)
			return true
		}
		if time.Now().After(expTime) {
			klog.V(3).Infof("Expired bootstrap token in %s/%s Secret: %v", secret.Namespace, secret.Name, expiration)
			return true
		}
	}
	return false
}
