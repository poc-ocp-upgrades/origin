package kubelet

import (
	"fmt"
	"k8s.io/kubernetes/pkg/util/initsystem"
)

func TryStartKubelet() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		fmt.Println("[kubelet-start] no supported init system detected, won't make sure the kubelet is running properly.")
		return
	}
	if !initSystem.ServiceExists("kubelet") {
		fmt.Println("[kubelet-start] couldn't detect a kubelet service, can't make sure the kubelet is running properly.")
	}
	fmt.Println("[kubelet-start] Activating the kubelet service")
	if err := initSystem.ServiceRestart("kubelet"); err != nil {
		fmt.Printf("[kubelet-start] WARNING: unable to start the kubelet service: [%v]\n", err)
		fmt.Printf("[kubelet-start] please ensure kubelet is reloaded and running manually.\n")
	}
}
func TryStopKubelet() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		fmt.Println("[kubelet-start] no supported init system detected, won't make sure the kubelet not running for a short period of time while setting up configuration for it.")
		return
	}
	if !initSystem.ServiceExists("kubelet") {
		fmt.Println("[kubelet-start] couldn't detect a kubelet service, can't make sure the kubelet not running for a short period of time while setting up configuration for it.")
	}
	if err := initSystem.ServiceStop("kubelet"); err != nil {
		fmt.Printf("[kubelet-start] WARNING: unable to stop the kubelet service momentarily: [%v]\n", err)
	}
}
