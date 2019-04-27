package auth

import (
	"fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	authv1client "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	oauthv1client "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	securityv1client "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

func reapForUser(userClient userv1client.UserV1Interface, authorizationClient authv1client.AuthorizationV1Interface, oauthClient oauthv1client.OauthV1Interface, securityClient securityv1client.SecurityContextConstraintsInterface, name string, out io.Writer) error {
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
	errors := []error{}
	removedSubject := corev1.ObjectReference{Kind: "User", Name: name}
	errors = append(errors, reapClusterBindings(removedSubject, authorizationClient, out)...)
	errors = append(errors, reapNamespacedBindings(removedSubject, authorizationClient, out)...)
	sccs, err := securityClient.List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, scc := range sccs.Items {
		retainedUsers := []string{}
		for _, user := range scc.Users {
			if user != name {
				retainedUsers = append(retainedUsers, user)
			}
		}
		if len(retainedUsers) != len(scc.Users) {
			updatedSCC := scc
			updatedSCC.Users = retainedUsers
			if _, err := securityClient.Update(&updatedSCC); err != nil && !kerrors.IsNotFound(err) {
				errors = append(errors, err)
			} else {
				fmt.Fprintf(out, "securitycontextconstraints.security.openshift.io/"+updatedSCC.Name+" updated\n")
			}
		}
	}
	groups, err := userClient.Groups().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, group := range groups.Items {
		retainedUsers := []string{}
		for _, user := range group.Users {
			if user != name {
				retainedUsers = append(retainedUsers, user)
			}
		}
		if len(retainedUsers) != len(group.Users) {
			updatedGroup := group
			updatedGroup.Users = retainedUsers
			if _, err := userClient.Groups().Update(&updatedGroup); err != nil && !kerrors.IsNotFound(err) {
				errors = append(errors, err)
			} else {
				fmt.Fprintf(out, "group.user.openshift.io/"+updatedGroup.Name+" updated\n")
			}
		}
	}
	authorizations, err := oauthClient.OAuthClientAuthorizations().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, authorization := range authorizations.Items {
		if authorization.UserName == name {
			if err := oauthClient.OAuthClientAuthorizations().Delete(authorization.Name, &metav1.DeleteOptions{}); err != nil && !kerrors.IsNotFound(err) {
				errors = append(errors, err)
			} else {
				fmt.Fprintf(out, "oauthclientauthorization.oauth.openshift.io/"+authorization.Name+" updated\n")
			}
		}
	}
	return utilerrors.NewAggregate(errors)
}
