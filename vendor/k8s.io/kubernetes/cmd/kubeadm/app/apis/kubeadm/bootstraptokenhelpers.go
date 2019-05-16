package kubeadm

import (
	goformat "fmt"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	bootstrapapi "k8s.io/cluster-bootstrap/token/api"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	"time"
	gotime "time"
)

func (bt *BootstrapToken) ToSecret() *v1.Secret {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: bootstraputil.BootstrapTokenSecretName(bt.Token.ID), Namespace: metav1.NamespaceSystem}, Type: v1.SecretType(bootstrapapi.SecretTypeBootstrapToken), Data: encodeTokenSecretData(bt, time.Now())}
}
func encodeTokenSecretData(token *BootstrapToken, now time.Time) map[string][]byte {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data := map[string][]byte{bootstrapapi.BootstrapTokenIDKey: []byte(token.Token.ID), bootstrapapi.BootstrapTokenSecretKey: []byte(token.Token.Secret)}
	if len(token.Description) > 0 {
		data[bootstrapapi.BootstrapTokenDescriptionKey] = []byte(token.Description)
	}
	if token.Expires != nil {
		expirationString := token.Expires.Time.Format(time.RFC3339)
		data[bootstrapapi.BootstrapTokenExpirationKey] = []byte(expirationString)
	} else if token.TTL != nil && token.TTL.Duration > 0 {
		expirationString := now.Add(token.TTL.Duration).Format(time.RFC3339)
		data[bootstrapapi.BootstrapTokenExpirationKey] = []byte(expirationString)
	}
	for _, usage := range token.Usages {
		data[bootstrapapi.BootstrapTokenUsagePrefix+usage] = []byte("true")
	}
	if len(token.Groups) > 0 {
		data[bootstrapapi.BootstrapTokenExtraGroupsKey] = []byte(strings.Join(token.Groups, ","))
	}
	return data
}
func BootstrapTokenFromSecret(secret *v1.Secret) (*BootstrapToken, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	tokenID := getSecretString(secret, bootstrapapi.BootstrapTokenIDKey)
	if len(tokenID) == 0 {
		return nil, errors.Errorf("bootstrap Token Secret has no token-id data: %s", secret.Name)
	}
	if secret.Name != bootstraputil.BootstrapTokenSecretName(tokenID) {
		return nil, errors.Errorf("bootstrap token name is not of the form '%s(token-id)'. Actual: %q. Expected: %q", bootstrapapi.BootstrapTokenSecretPrefix, secret.Name, bootstraputil.BootstrapTokenSecretName(tokenID))
	}
	tokenSecret := getSecretString(secret, bootstrapapi.BootstrapTokenSecretKey)
	if len(tokenSecret) == 0 {
		return nil, errors.Errorf("bootstrap Token Secret has no token-secret data: %s", secret.Name)
	}
	bts, err := NewBootstrapTokenStringFromIDAndSecret(tokenID, tokenSecret)
	if err != nil {
		return nil, errors.Wrap(err, "bootstrap Token Secret is invalid and couldn't be parsed")
	}
	description := getSecretString(secret, bootstrapapi.BootstrapTokenDescriptionKey)
	secretExpiration := getSecretString(secret, bootstrapapi.BootstrapTokenExpirationKey)
	var expires *metav1.Time
	if len(secretExpiration) > 0 {
		expTime, err := time.Parse(time.RFC3339, secretExpiration)
		if err != nil {
			return nil, errors.Wrapf(err, "can't parse expiration time of bootstrap token %q", secret.Name)
		}
		expires = &metav1.Time{Time: expTime}
	}
	var usages []string
	for k, v := range secret.Data {
		if !strings.HasPrefix(k, bootstrapapi.BootstrapTokenUsagePrefix) {
			continue
		}
		if string(v) != "true" {
			continue
		}
		usages = append(usages, strings.TrimPrefix(k, bootstrapapi.BootstrapTokenUsagePrefix))
	}
	if usages != nil {
		sort.Strings(usages)
	}
	var groups []string
	groupsString := getSecretString(secret, bootstrapapi.BootstrapTokenExtraGroupsKey)
	g := strings.Split(groupsString, ",")
	if len(g) > 0 && len(g[0]) > 0 {
		groups = g
	}
	return &BootstrapToken{Token: bts, Description: description, Expires: expires, Usages: usages, Groups: groups}, nil
}
func getSecretString(secret *v1.Secret, key string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if secret.Data == nil {
		return ""
	}
	if val, ok := secret.Data[key]; ok {
		return string(val)
	}
	return ""
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
