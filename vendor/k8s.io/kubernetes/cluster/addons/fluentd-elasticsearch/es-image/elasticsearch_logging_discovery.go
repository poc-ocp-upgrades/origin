package main

import (
	"flag"
	"fmt"
	goformat "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
	api "k8s.io/kubernetes/pkg/apis/core"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"time"
	gotime "time"
)

func buildConfigFromEnvs(masterURL, kubeconfigPath string) (*restclient.Config, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if kubeconfigPath == "" && masterURL == "" {
		kubeconfig, err := restclient.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return kubeconfig, nil
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath}, &clientcmd.ConfigOverrides{ClusterInfo: clientapi.Cluster{Server: masterURL}}).ClientConfig()
}
func flattenSubsets(subsets []api.EndpointSubset) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	ips := []string{}
	for _, ss := range subsets {
		for _, addr := range ss.Addresses {
			ips = append(ips, fmt.Sprintf(`"%s"`, addr.IP))
		}
	}
	return ips
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flag.Parse()
	klog.Info("Kubernetes Elasticsearch logging discovery")
	cc, err := buildConfigFromEnvs(os.Getenv("APISERVER_HOST"), os.Getenv("KUBE_CONFIG_FILE"))
	if err != nil {
		klog.Fatalf("Failed to make client: %v", err)
	}
	client, err := clientset.NewForConfig(cc)
	if err != nil {
		klog.Fatalf("Failed to make client: %v", err)
	}
	namespace := metav1.NamespaceSystem
	envNamespace := os.Getenv("NAMESPACE")
	if envNamespace != "" {
		if _, err := client.Core().Namespaces().Get(envNamespace, metav1.GetOptions{}); err != nil {
			klog.Fatalf("%s namespace doesn't exist: %v", envNamespace, err)
		}
		namespace = envNamespace
	}
	var elasticsearch *api.Service
	serviceName := os.Getenv("ELASTICSEARCH_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "elasticsearch-logging"
	}
	for t := time.Now(); time.Since(t) < 5*time.Minute; time.Sleep(10 * time.Second) {
		elasticsearch, err = client.Core().Services(namespace).Get(serviceName, metav1.GetOptions{})
		if err == nil {
			break
		}
	}
	if elasticsearch == nil {
		klog.Warningf("Failed to find the elasticsearch-logging service: %v", err)
		return
	}
	var endpoints *api.Endpoints
	addrs := []string{}
	count, _ := strconv.Atoi(os.Getenv("MINIMUM_MASTER_NODES"))
	for t := time.Now(); time.Since(t) < 5*time.Minute; time.Sleep(10 * time.Second) {
		endpoints, err = client.Core().Endpoints(namespace).Get(serviceName, metav1.GetOptions{})
		if err != nil {
			continue
		}
		addrs = flattenSubsets(endpoints.Subsets)
		klog.Infof("Found %s", addrs)
		if len(addrs) > 0 && len(addrs) >= count {
			break
		}
	}
	if err != nil {
		klog.Warningf("Error finding endpoints: %v", err)
		return
	}
	klog.Infof("Endpoints = %s", addrs)
	fmt.Printf("discovery.zen.ping.unicast.hosts: [%s]\n", strings.Join(addrs, ", "))
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
