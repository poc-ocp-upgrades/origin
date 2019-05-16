package oauthclient

import (
	goformat "fmt"
	oauthapi "github.com/openshift/api/oauth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

type Getter interface {
	Get(name string, options metav1.GetOptions) (*oauthapi.OAuthClient, error)
}

func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
