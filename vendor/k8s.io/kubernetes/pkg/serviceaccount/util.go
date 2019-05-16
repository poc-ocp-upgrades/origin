package serviceaccount

import (
	"k8s.io/api/core/v1"
	apiserverserviceaccount "k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/apiserver/pkg/authentication/user"
)

const (
	PodNameKey = "authentication.kubernetes.io/pod-name"
	PodUIDKey  = "authentication.kubernetes.io/pod-uid"
)

func UserInfo(namespace, name, uid string) user.Info {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&ServiceAccountInfo{Name: name, Namespace: namespace, UID: uid}).UserInfo()
}

type ServiceAccountInfo struct {
	Name, Namespace, UID string
	PodName, PodUID      string
}

func (sa *ServiceAccountInfo) UserInfo() user.Info {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	info := &user.DefaultInfo{Name: apiserverserviceaccount.MakeUsername(sa.Namespace, sa.Name), UID: sa.UID, Groups: apiserverserviceaccount.MakeGroupNames(sa.Namespace)}
	if sa.PodName != "" && sa.PodUID != "" {
		info.Extra = map[string][]string{PodNameKey: {sa.PodName}, PodUIDKey: {sa.PodUID}}
	}
	return info
}
func IsServiceAccountToken(secret *v1.Secret, sa *v1.ServiceAccount) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if secret.Type != v1.SecretTypeServiceAccountToken {
		return false
	}
	name := secret.Annotations[v1.ServiceAccountNameKey]
	uid := secret.Annotations[v1.ServiceAccountUIDKey]
	if name != sa.Name {
		return false
	}
	if len(uid) > 0 && uid != string(sa.UID) {
		return false
	}
	return true
}
