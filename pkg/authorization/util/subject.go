package util

import (
	goformat "fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func BuildRBACSubjects(users, groups []string) []rbacv1.Subject {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
