package controller

import (
	"crypto/tls"
	"crypto/x509"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/library-go/pkg/crypto"
	cmdutil "github.com/openshift/origin/pkg/cmd/util"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	apifilters "k8s.io/apiserver/pkg/endpoints/filters"
	apiserver "k8s.io/apiserver/pkg/server"
	apiserverfilters "k8s.io/apiserver/pkg/server/filters"
	"k8s.io/apiserver/pkg/server/healthz"
	genericmux "k8s.io/apiserver/pkg/server/mux"
	genericroutes "k8s.io/apiserver/pkg/server/routes"
	authzwebhook "k8s.io/apiserver/plugin/pkg/authorizer/webhook"
	clientgoclientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"net/http"
	"time"
)

func RunControllerServer(servingInfo configv1.HTTPServingInfo, kubeExternal clientgoclientset.Interface) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	clientCAs, err := getClientCertCAPool(servingInfo)
	if err != nil {
		return err
	}
	mux := genericmux.NewPathRecorderMux("master-healthz")
	healthz.InstallHandler(mux, healthz.PingHealthz, healthz.LogHealthz)
	initReadinessCheckRoute(mux, "/healthz/ready", func() bool {
		return true
	})
	genericroutes.Profiling{}.Install(mux)
	genericroutes.MetricsWithReset{}.Install(mux)
	tokenReview := kubeExternal.AuthenticationV1beta1().TokenReviews()
	authn, err := newRemoteAuthenticator(tokenReview, clientCAs, 5*time.Minute)
	if err != nil {
		return err
	}
	sarClient := kubeExternal.AuthorizationV1beta1().SubjectAccessReviews()
	remoteAuthz, err := authzwebhook.NewFromInterface(sarClient, 5*time.Minute, 5*time.Minute)
	if err != nil {
		return err
	}
	requestInfoResolver := apiserver.NewRequestInfoResolver(&apiserver.Config{})
	authz := newBypassAuthorizer(remoteAuthz, "/healthz", "/healthz/ready")
	handler := apifilters.WithAuthorization(mux, authz, legacyscheme.Codecs)
	handler = apifilters.WithAuthentication(handler, authn, apifilters.Unauthorized(legacyscheme.Codecs, false), nil)
	handler = apiserverfilters.WithPanicRecovery(handler)
	handler = apifilters.WithRequestInfo(handler, requestInfoResolver)
	return serveControllers(servingInfo, handler)
}
func initReadinessCheckRoute(mux *genericmux.PathRecorderMux, path string, readyFunc func() bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if readyFunc() {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
}
func serveControllers(servingInfo configv1.HTTPServingInfo, handler http.Handler) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	timeout := servingInfo.RequestTimeoutSeconds
	if timeout == -1 {
		timeout = 0
	}
	server := &http.Server{Addr: servingInfo.BindAddress, Handler: handler, ReadTimeout: time.Duration(timeout) * time.Second, WriteTimeout: time.Duration(timeout) * time.Second, MaxHeaderBytes: 1 << 20}
	clientCAs, err := getClientCertCAPool(servingInfo)
	if err != nil {
		return err
	}
	go utilwait.Forever(func() {
		klog.Infof("Started health checks at %s", servingInfo.BindAddress)
		extraCerts, err := getNamedCertificateMap(servingInfo.NamedCertificates)
		if err != nil {
			klog.Fatal(err)
		}
		server.TLSConfig = crypto.SecureTLSConfig(&tls.Config{ClientAuth: tls.RequestClientCert, ClientCAs: clientCAs, GetCertificate: cmdutil.GetCertificateFunc(extraCerts), MinVersion: crypto.TLSVersionOrDie(servingInfo.MinTLSVersion), CipherSuites: crypto.CipherSuitesOrDie(servingInfo.CipherSuites)})
		klog.Fatal(cmdutil.ListenAndServeTLS(server, servingInfo.BindNetwork, servingInfo.CertFile, servingInfo.KeyFile))
	}, 0)
	return nil
}
func getClientCertCAPool(servingInfo configv1.HTTPServingInfo) (*x509.CertPool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	roots := x509.NewCertPool()
	certs, err := cmdutil.CertificatesFromFile(servingInfo.ClientCA)
	if err != nil {
		return nil, err
	}
	for _, root := range certs {
		roots.AddCert(root)
	}
	return roots, nil
}
func getNamedCertificateMap(namedCertificates []configv1.NamedCertificate) (map[string]*tls.Certificate, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(namedCertificates) == 0 {
		return nil, nil
	}
	namedCerts := map[string]*tls.Certificate{}
	for _, namedCertificate := range namedCertificates {
		cert, err := tls.LoadX509KeyPair(namedCertificate.CertFile, namedCertificate.KeyFile)
		if err != nil {
			return nil, err
		}
		for _, name := range namedCertificate.Names {
			namedCerts[name] = &cert
		}
	}
	return namedCerts, nil
}
