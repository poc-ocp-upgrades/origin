package integration

import (
	"math/rand"
	"net"
	"testing"
	"time"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	kinformers "k8s.io/client-go/informers"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	"github.com/openshift/origin/pkg/route/controller/ingressip"
	testserver "github.com/openshift/origin/test/util/server"
)

const sentinelName = "sentinel"

func TestIngressIPAllocation(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	masterConfig, err := testserver.DefaultMasterOptions()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)
	masterConfig.NetworkConfig.ExternalIPNetworkCIDRs = []string{"172.16.0.0/24"}
	masterConfig.NetworkConfig.IngressIPNetworkCIDR = "172.16.1.0/24"
	stopCh := make(chan struct{})
	defer close(stopCh)
	clusterAdminKubeConfig, err := testserver.StartConfiguredMasterWithOptions(masterConfig, stopCh)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	clientConfig, err := configapi.GetClientConfig(clusterAdminKubeConfig, &configapi.ClientConnectionOverrides{QPS: 20, Burst: 50})
	if err != nil {
		t.Fatal(err)
	}
	kc := kclientset.NewForConfigOrDie(clientConfig)
	stopChannel := make(chan struct{})
	defer close(stopChannel)
	received := make(chan bool)
	rand.Seed(time.Now().UTC().UnixNano())
	t.Log("start informer to watch for sentinel")
	_, informerController := cache.NewInformer(&cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		return kc.CoreV1().Services(metav1.NamespaceAll).List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		return kc.CoreV1().Services(metav1.NamespaceAll).Watch(options)
	}}, &v1.Service{}, time.Minute*10, cache.ResourceEventHandlerFuncs{UpdateFunc: func(old, cur interface{}) {
		service := cur.(*v1.Service)
		if service.Name == sentinelName && len(service.Spec.ExternalIPs) > 0 {
			received <- true
		}
	}})
	go informerController.Run(stopChannel)
	t.Log("start generating service events")
	go generateServiceEvents(t, kc)
	kubeInformers := kinformers.NewSharedInformerFactory(kc, 0)
	_, ipNet, err := net.ParseCIDR(masterConfig.NetworkConfig.IngressIPNetworkCIDR)
	c := ingressip.NewIngressIPController(kubeInformers.Core().V1().Services().Informer(), kc, ipNet, 10*time.Minute)
	kubeInformers.Start(stopChannel)
	go c.Run(stopChannel)
	t.Log("waiting for sentinel to be updated with external ip")
	select {
	case <-received:
	case <-time.After(time.Duration(90 * time.Second)):
		t.Fatal("took too long")
	}
	services, err := kc.CoreV1().Services(metav1.NamespaceDefault).List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	ips := sets.NewString()
	for _, s := range services.Items {
		typeLoadBalancer := s.Spec.Type == v1.ServiceTypeLoadBalancer
		hasAllocation := len(s.Status.LoadBalancer.Ingress) > 0
		switch {
		case !typeLoadBalancer && !hasAllocation:
			continue
		case !typeLoadBalancer && hasAllocation:
			t.Errorf("A service not of type load balancer has an ingress ip allocation")
			continue
		case typeLoadBalancer && !hasAllocation:
			t.Errorf("A service of type load balancer has not been allocated an ingress ip")
			continue
		}
		ingressIP := s.Status.LoadBalancer.Ingress[0].IP
		if ips.Has(ingressIP) {
			t.Errorf("One or more services have the same ingress ip")
			continue
		}
		ips.Insert(ingressIP)
		if len(s.Spec.ExternalIPs) == 0 || s.Spec.ExternalIPs[0] != ingressIP {
			t.Errorf("Service does not have the ingress ip as an external ip")
			continue
		}
	}
}

const (
	createOp	= iota
	updateOp
	deleteOp
)

func generateServiceEvents(t *testing.T, kc kclientset.Interface) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	maxMillisecondInterval := 25
	minServiceCount := 10
	maxOperations := minServiceCount + 30
	var services []*v1.Service
	for i := 0; i < maxOperations; {
		op := createOp
		if len(services) > minServiceCount {
			op = rand.Intn(deleteOp + 1)
		}
		switch op {
		case createOp:
			typeChoice := rand.Intn(2)
			typeLoadBalancer := false
			if typeChoice == 1 {
				typeLoadBalancer = true
			}
			s, err := createService(kc, "", typeLoadBalancer)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			services = append(services, s)
			t.Logf("Added service %s", s.Name)
		case updateOp:
			targetIndex := rand.Intn(len(services))
			name := services[targetIndex].Name
			s, err := kc.CoreV1().Services(metav1.NamespaceDefault).Get(name, metav1.GetOptions{})
			if err != nil {
				continue
			}
			if s.Spec.Type == v1.ServiceTypeLoadBalancer {
				s.Spec.Type = v1.ServiceTypeClusterIP
				s.Spec.Ports[0].NodePort = 0
			} else {
				s.Spec.Type = v1.ServiceTypeLoadBalancer
			}
			s, err = kc.CoreV1().Services(metav1.NamespaceDefault).Update(s)
			if err != nil {
				continue
			}
			t.Logf("Updated service %s", name)
		case deleteOp:
			targetIndex := rand.Intn(len(services))
			name := services[targetIndex].Name
			err := kc.CoreV1().Services(metav1.NamespaceDefault).Delete(name, nil)
			if err != nil {
				continue
			}
			services = append(services[:targetIndex], services[targetIndex+1:]...)
			t.Logf("Deleted service %s", name)
		}
		i++
		time.Sleep(time.Duration(rand.Intn(maxMillisecondInterval)) * time.Millisecond)
	}
	time.Sleep(time.Millisecond * 100)
	_, err := createService(kc, sentinelName, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
func createService(kc kclientset.Interface, name string, typeLoadBalancer bool) (*v1.Service, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceType := v1.ServiceTypeClusterIP
	if typeLoadBalancer {
		serviceType = v1.ServiceTypeLoadBalancer
	}
	service := &v1.Service{ObjectMeta: metav1.ObjectMeta{GenerateName: "service-", Name: name}, Spec: v1.ServiceSpec{Type: serviceType, Ports: []v1.ServicePort{{Protocol: "TCP", Port: 8080}}}}
	return kc.CoreV1().Services(metav1.NamespaceDefault).Create(service)
}
