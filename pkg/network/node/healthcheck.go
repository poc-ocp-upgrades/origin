package node

import (
	"fmt"
	"time"
	"k8s.io/klog"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	"github.com/openshift/origin/pkg/util/ovs/ovsclient"
)

const (
	ovsDialTimeout		= 5 * time.Second
	ovsHealthcheckInterval	= 30 * time.Second
	ovsRecoveryTimeout	= 10 * time.Second
	ovsDialDefaultNetwork	= "unix"
	ovsDialDefaultAddress	= "/var/run/openvswitch/db.sock"
)

func waitForOVS(network, addr string) error {
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
	return utilwait.PollImmediate(time.Second, time.Minute, func() (bool, error) {
		c, err := ovsclient.DialTimeout(network, addr, ovsDialTimeout)
		if err != nil {
			klog.V(2).Infof("waiting for OVS to start: %v", err)
			return false, nil
		}
		defer c.Close()
		if err := c.Ping(); err != nil {
			klog.V(2).Infof("waiting for OVS to start, ping failed: %v", err)
			return false, nil
		}
		return true, nil
	})
}
func runOVSHealthCheck(network, addr string, healthFn func() error) {
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
	go utilwait.Until(func() {
		c, err := ovsclient.DialTimeout(network, addr, ovsDialTimeout)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("SDN healthcheck unable to connect to OVS server: %v", err))
			return
		}
		defer c.Close()
		err = c.WaitForDisconnect()
		utilruntime.HandleError(fmt.Errorf("SDN healthcheck disconnected from OVS server: %v", err))
		err = utilwait.PollImmediate(100*time.Millisecond, ovsRecoveryTimeout, func() (bool, error) {
			c, err := ovsclient.DialTimeout(network, addr, ovsDialTimeout)
			if err != nil {
				klog.V(2).Infof("SDN healthcheck unable to reconnect to OVS server: %v", err)
				return false, nil
			}
			defer c.Close()
			if err := c.Ping(); err != nil {
				klog.V(2).Infof("SDN healthcheck unable to ping OVS server: %v", err)
				return false, nil
			}
			if err := healthFn(); err != nil {
				return false, fmt.Errorf("OVS health check failed: %v", err)
			}
			return true, nil
		})
		if err != nil {
			klog.Fatalf("SDN healthcheck detected unhealthy OVS server, restarting: %v", err)
		}
	}, ovsDialTimeout, utilwait.NeverStop)
	go utilwait.Until(func() {
		c, err := ovsclient.DialTimeout(network, addr, ovsDialTimeout)
		if err != nil {
			klog.V(2).Infof("SDN healthcheck unable to reconnect to OVS server: %v", err)
			return
		}
		defer c.Close()
		if err := c.Ping(); err != nil {
			klog.V(2).Infof("SDN healthcheck unable to ping OVS server: %v", err)
			return
		}
		if err := healthFn(); err != nil {
			klog.Fatalf("SDN healthcheck detected unhealthy OVS server, restarting: %v", err)
		}
		klog.V(4).Infof("SDN healthcheck succeeded")
	}, ovsHealthcheckInterval, utilwait.NeverStop)
}
