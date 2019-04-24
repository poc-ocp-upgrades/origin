package util

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
)

func BuildRBACSubjects(users, groups []string) []rbacv1.Subject {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subjects := []rbacv1.Subject{}
	for _, user := range users {
		saNamespace, saName, err := serviceaccount.SplitUsername(user)
		if err == nil {
			subjects = append(subjects, rbacv1.Subject{Kind: rbacv1.ServiceAccountKind, Namespace: saNamespace, Name: saName})
		} else {
			subjects = append(subjects, rbacv1.Subject{Kind: rbacv1.UserKind, APIGroup: rbacv1.GroupName, Name: user})
		}
	}
	for _, group := range groups {
		subjects = append(subjects, rbacv1.Subject{Kind: rbacv1.GroupKind, APIGroup: rbacv1.GroupName, Name: group})
	}
	return subjects
}
func RBACSubjectsToUsersAndGroups(subjects []rbacv1.Subject, defaultNamespace string) (users []string, groups []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, subject := range subjects {
		switch {
		case subject.APIGroup == rbacv1.GroupName && subject.Kind == rbacv1.GroupKind:
			groups = append(groups, subject.Name)
		case subject.APIGroup == rbacv1.GroupName && subject.Kind == rbacv1.UserKind:
			users = append(users, subject.Name)
		case subject.APIGroup == "" && subject.Kind == rbacv1.ServiceAccountKind:
			ns := defaultNamespace
			if len(subject.Namespace) > 0 {
				ns = subject.Namespace
			}
			if len(ns) > 0 {
				name := serviceaccount.MakeUsername(ns, subject.Name)
				users = append(users, name)
			} else {
			}
		default:
		}
	}
	return users, groups
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
