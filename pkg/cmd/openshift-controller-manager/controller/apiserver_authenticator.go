package controller

import (
	"crypto/x509"
	goformat "fmt"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/group"
	"k8s.io/apiserver/pkg/authentication/request/anonymous"
	"k8s.io/apiserver/pkg/authentication/request/bearertoken"
	"k8s.io/apiserver/pkg/authentication/request/union"
	x509request "k8s.io/apiserver/pkg/authentication/request/x509"
	"k8s.io/apiserver/pkg/authentication/token/cache"
	webhooktoken "k8s.io/apiserver/plugin/pkg/authenticator/token/webhook"
	authenticationclient "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

func newRemoteAuthenticator(tokenReview authenticationclient.TokenReviewInterface, clientCAs *x509.CertPool, cacheTTL time.Duration) (authenticator.Request, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	authenticators := []authenticator.Request{}
	tokenAuthenticator, err := webhooktoken.NewFromInterface(tokenReview, nil)
	if err != nil {
		return nil, err
	}
	cachingTokenAuth := cache.New(tokenAuthenticator, false, cacheTTL, cacheTTL)
	authenticators = append(authenticators, bearertoken.New(cachingTokenAuth))
	if clientCAs != nil {
		opts := x509request.DefaultVerifyOptions()
		opts.Roots = clientCAs
		certauth := x509request.New(opts, x509request.CommonNameUserConversion)
		authenticators = append(authenticators, certauth)
	}
	return union.NewFailOnError(group.NewAuthenticatedGroupAdder(union.New(authenticators...)), anonymous.NewAuthenticator()), nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
