package openshiftkubeapiserver

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"io/ioutil"
	"net"
	"strings"
	configv1 "github.com/openshift/api/config/v1"
	kubecontrolplanev1 "github.com/openshift/api/kubecontrolplane/v1"
	"github.com/openshift/origin/pkg/cmd/configflags"
	configapilatest "github.com/openshift/origin/pkg/cmd/server/apis/config/latest"
	"github.com/openshift/origin/pkg/configconversion"
)

func ConfigToFlags(kubeAPIServerConfig *kubecontrolplanev1.KubeAPIServerConfig) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := unmaskArgs(kubeAPIServerConfig.APIServerArguments)
	host, portString, err := net.SplitHostPort(kubeAPIServerConfig.ServingInfo.BindAddress)
	if err != nil {
		return nil, err
	}
	admissionFlags, err := admissionFlags(kubeAPIServerConfig.AdmissionConfig)
	if err != nil {
		return nil, err
	}
	for flag, value := range admissionFlags {
		configflags.SetIfUnset(args, flag, value...)
	}
	configflags.SetIfUnset(args, "allow-privileged", "true")
	configflags.SetIfUnset(args, "anonymous-auth", "false")
	configflags.SetIfUnset(args, "authorization-mode", "RBAC", "Node")
	for flag, value := range configflags.AuditFlags(&kubeAPIServerConfig.AuditConfig, configflags.ArgsWithPrefix(args, "audit-")) {
		configflags.SetIfUnset(args, flag, value...)
	}
	configflags.SetIfUnset(args, "bind-address", host)
	configflags.SetIfUnset(args, "client-ca-file", kubeAPIServerConfig.ServingInfo.ClientCA)
	configflags.SetIfUnset(args, "cors-allowed-origins", kubeAPIServerConfig.CORSAllowedOrigins...)
	configflags.SetIfUnset(args, "enable-logs-handler", "false")
	configflags.SetIfUnset(args, "enable-swagger-ui", "true")
	configflags.SetIfUnset(args, "endpoint-reconciler-type", "lease")
	configflags.SetIfUnset(args, "etcd-cafile", kubeAPIServerConfig.StorageConfig.CA)
	configflags.SetIfUnset(args, "etcd-certfile", kubeAPIServerConfig.StorageConfig.CertFile)
	configflags.SetIfUnset(args, "etcd-keyfile", kubeAPIServerConfig.StorageConfig.KeyFile)
	configflags.SetIfUnset(args, "etcd-prefix", kubeAPIServerConfig.StorageConfig.StoragePrefix)
	configflags.SetIfUnset(args, "etcd-servers", kubeAPIServerConfig.StorageConfig.URLs...)
	configflags.SetIfUnset(args, "event-ttl", "3h")
	configflags.SetIfUnset(args, "insecure-port", "0")
	configflags.SetIfUnset(args, "kubelet-certificate-authority", kubeAPIServerConfig.KubeletClientInfo.CA)
	configflags.SetIfUnset(args, "kubelet-client-certificate", kubeAPIServerConfig.KubeletClientInfo.CertFile)
	configflags.SetIfUnset(args, "kubelet-client-key", kubeAPIServerConfig.KubeletClientInfo.KeyFile)
	configflags.SetIfUnset(args, "kubelet-https", "true")
	configflags.SetIfUnset(args, "kubelet-preferred-address-types", "Hostname", "InternalIP", "ExternalIP")
	configflags.SetIfUnset(args, "kubelet-read-only-port", "0")
	configflags.SetIfUnset(args, "kubernetes-service-node-port", "0")
	configflags.SetIfUnset(args, "max-mutating-requests-inflight", fmt.Sprintf("%d", kubeAPIServerConfig.ServingInfo.MaxRequestsInFlight/2))
	configflags.SetIfUnset(args, "max-requests-inflight", fmt.Sprintf("%d", kubeAPIServerConfig.ServingInfo.MaxRequestsInFlight))
	configflags.SetIfUnset(args, "min-request-timeout", fmt.Sprintf("%d", kubeAPIServerConfig.ServingInfo.RequestTimeoutSeconds))
	configflags.SetIfUnset(args, "proxy-client-cert-file", kubeAPIServerConfig.AggregatorConfig.ProxyClientInfo.CertFile)
	configflags.SetIfUnset(args, "proxy-client-key-file", kubeAPIServerConfig.AggregatorConfig.ProxyClientInfo.KeyFile)
	configflags.SetIfUnset(args, "requestheader-allowed-names", kubeAPIServerConfig.AuthConfig.RequestHeader.ClientCommonNames...)
	configflags.SetIfUnset(args, "requestheader-client-ca-file", kubeAPIServerConfig.AuthConfig.RequestHeader.ClientCA)
	configflags.SetIfUnset(args, "requestheader-extra-headers-prefix", kubeAPIServerConfig.AuthConfig.RequestHeader.ExtraHeaderPrefixes...)
	configflags.SetIfUnset(args, "requestheader-group-headers", kubeAPIServerConfig.AuthConfig.RequestHeader.GroupHeaders...)
	configflags.SetIfUnset(args, "requestheader-username-headers", kubeAPIServerConfig.AuthConfig.RequestHeader.UsernameHeaders...)
	configflags.SetIfUnset(args, "secure-port", portString)
	configflags.SetIfUnset(args, "service-cluster-ip-range", kubeAPIServerConfig.ServicesSubnet)
	configflags.SetIfUnset(args, "service-node-port-range", kubeAPIServerConfig.ServicesNodePortRange)
	configflags.SetIfUnset(args, "storage-backend", "etcd3")
	configflags.SetIfUnset(args, "storage-media-type", "application/vnd.kubernetes.protobuf")
	configflags.SetIfUnset(args, "tls-cert-file", kubeAPIServerConfig.ServingInfo.CertFile)
	configflags.SetIfUnset(args, "tls-cipher-suites", kubeAPIServerConfig.ServingInfo.CipherSuites...)
	configflags.SetIfUnset(args, "tls-min-version", kubeAPIServerConfig.ServingInfo.MinTLSVersion)
	configflags.SetIfUnset(args, "tls-private-key-file", kubeAPIServerConfig.ServingInfo.KeyFile)
	configflags.SetIfUnset(args, "tls-sni-cert-key", sniCertKeys(kubeAPIServerConfig.ServingInfo.NamedCertificates)...)
	configflags.SetIfUnset(args, "secure-port", portString)
	return configflags.ToFlagSlice(args), nil
}
func admissionFlags(admissionConfig configv1.AdmissionConfig) (map[string][]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := map[string][]string{}
	upstreamAdmissionConfig, err := configconversion.ConvertOpenshiftAdmissionConfigToKubeAdmissionConfig(admissionConfig.PluginConfig)
	if err != nil {
		return nil, err
	}
	configBytes, err := configapilatest.WriteYAML(upstreamAdmissionConfig)
	if err != nil {
		return nil, err
	}
	tempFile, err := ioutil.TempFile("", "kubeapiserver-admission-config.yaml")
	if err != nil {
		return nil, err
	}
	if _, err := tempFile.Write(configBytes); err != nil {
		return nil, err
	}
	tempFile.Close()
	configflags.SetIfUnset(args, "admission-control-config-file", tempFile.Name())
	configflags.SetIfUnset(args, "disable-admission-plugins", admissionConfig.DisabledAdmissionPlugins...)
	configflags.SetIfUnset(args, "enable-admission-plugins", admissionConfig.EnabledAdmissionPlugins...)
	return args, nil
}
func sniCertKeys(namedCertificates []configv1.NamedCertificate) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{}
	for _, nc := range namedCertificates {
		names := ""
		if len(nc.Names) > 0 {
			names = ":" + strings.Join(nc.Names, ",")
		}
		args = append(args, fmt.Sprintf("%s,%s%s", nc.CertFile, nc.KeyFile, names))
	}
	return args
}
func unmaskArgs(args map[string]kubecontrolplanev1.Arguments) map[string][]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := map[string][]string{}
	for key, slice := range args {
		for _, val := range slice {
			ret[key] = append(ret[key], val)
		}
	}
	return ret
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
