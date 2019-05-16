package v1alpha1

import (
	"fmt"
	goformat "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	kubeproxyconfigv1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
	"k8s.io/kubernetes/pkg/kubelet/qos"
	"k8s.io/kubernetes/pkg/master/ports"
	"k8s.io/utils/pointer"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	"time"
	gotime "time"
)

func addDefaultingFuncs(scheme *kruntime.Scheme) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return RegisterDefaults(scheme)
}
func SetDefaults_KubeProxyConfiguration(obj *kubeproxyconfigv1alpha1.KubeProxyConfiguration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(obj.BindAddress) == 0 {
		obj.BindAddress = "0.0.0.0"
	}
	if obj.HealthzBindAddress == "" {
		obj.HealthzBindAddress = fmt.Sprintf("0.0.0.0:%v", ports.ProxyHealthzPort)
	} else if !strings.Contains(obj.HealthzBindAddress, ":") {
		obj.HealthzBindAddress += fmt.Sprintf(":%v", ports.ProxyHealthzPort)
	}
	if obj.MetricsBindAddress == "" {
		obj.MetricsBindAddress = fmt.Sprintf("127.0.0.1:%v", ports.ProxyStatusPort)
	} else if !strings.Contains(obj.MetricsBindAddress, ":") {
		obj.MetricsBindAddress += fmt.Sprintf(":%v", ports.ProxyStatusPort)
	}
	if obj.OOMScoreAdj == nil {
		temp := int32(qos.KubeProxyOOMScoreAdj)
		obj.OOMScoreAdj = &temp
	}
	if obj.ResourceContainer == "" {
		obj.ResourceContainer = "/kube-proxy"
	}
	if obj.IPTables.SyncPeriod.Duration == 0 {
		obj.IPTables.SyncPeriod = metav1.Duration{Duration: 30 * time.Second}
	}
	if obj.IPVS.SyncPeriod.Duration == 0 {
		obj.IPVS.SyncPeriod = metav1.Duration{Duration: 30 * time.Second}
	}
	zero := metav1.Duration{}
	if obj.UDPIdleTimeout == zero {
		obj.UDPIdleTimeout = metav1.Duration{Duration: 250 * time.Millisecond}
	}
	if obj.Conntrack.Max == nil {
		if obj.Conntrack.MaxPerCore == nil {
			obj.Conntrack.MaxPerCore = pointer.Int32Ptr(32 * 1024)
		}
		if obj.Conntrack.Min == nil {
			obj.Conntrack.Min = pointer.Int32Ptr(128 * 1024)
		}
	}
	if obj.IPTables.MasqueradeBit == nil {
		temp := int32(14)
		obj.IPTables.MasqueradeBit = &temp
	}
	if obj.Conntrack.TCPEstablishedTimeout == nil {
		obj.Conntrack.TCPEstablishedTimeout = &metav1.Duration{Duration: 24 * time.Hour}
	}
	if obj.Conntrack.TCPCloseWaitTimeout == nil {
		obj.Conntrack.TCPCloseWaitTimeout = &metav1.Duration{Duration: 1 * time.Hour}
	}
	if obj.ConfigSyncPeriod.Duration == 0 {
		obj.ConfigSyncPeriod.Duration = 15 * time.Minute
	}
	if len(obj.ClientConnection.ContentType) == 0 {
		obj.ClientConnection.ContentType = "application/vnd.kubernetes.protobuf"
	}
	if obj.ClientConnection.QPS == 0.0 {
		obj.ClientConnection.QPS = 5.0
	}
	if obj.ClientConnection.Burst == 0 {
		obj.ClientConnection.Burst = 10
	}
	if obj.FeatureGates == nil {
		obj.FeatureGates = make(map[string]bool)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
