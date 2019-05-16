package main

import (
	"fmt"
	"k8s.io/klog"
	"os"
	"os/exec"
	"strings"
	"time"
)

type EtcdMigrateServer struct {
	cfg    *EtcdMigrateCfg
	client EtcdMigrateClient
	cmd    *exec.Cmd
}

func NewEtcdMigrateServer(cfg *EtcdMigrateCfg, client EtcdMigrateClient) *EtcdMigrateServer {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &EtcdMigrateServer{cfg: cfg, client: client}
}
func (r *EtcdMigrateServer) Start(version *EtcdVersion) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	etcdCmd := exec.Command(fmt.Sprintf("%s/etcd-%s", r.cfg.binPath, version), "--name", r.cfg.name, "--initial-cluster", r.cfg.initialCluster, "--debug", "--data-dir", r.cfg.dataDirectory, "--listen-client-urls", fmt.Sprintf("http://127.0.0.1:%d", r.cfg.port), "--advertise-client-urls", fmt.Sprintf("http://127.0.0.1:%d", r.cfg.port), "--listen-peer-urls", r.cfg.peerListenUrls, "--initial-advertise-peer-urls", r.cfg.peerAdvertiseUrls)
	if r.cfg.etcdServerArgs != "" {
		extraArgs := strings.Fields(r.cfg.etcdServerArgs)
		etcdCmd.Args = append(etcdCmd.Args, extraArgs...)
	}
	fmt.Printf("Starting server %s: %+v\n", r.cfg.name, etcdCmd.Args)
	etcdCmd.Stdout = os.Stdout
	etcdCmd.Stderr = os.Stderr
	err := etcdCmd.Start()
	if err != nil {
		return err
	}
	interval := time.NewTicker(time.Millisecond * 500)
	defer interval.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Minute * 2)
		done <- true
	}()
	for {
		select {
		case <-interval.C:
			err := r.client.SetEtcdVersionKeyValue(version)
			if err != nil {
				klog.Infof("Still waiting for etcd to start, current error: %v", err)
			} else {
				klog.Infof("Etcd on port %d is up.", r.cfg.port)
				r.cmd = etcdCmd
				return nil
			}
		case <-done:
			err = etcdCmd.Process.Kill()
			if err != nil {
				return fmt.Errorf("error killing etcd: %v", err)
			}
			return fmt.Errorf("Timed out waiting for etcd on port %d", r.cfg.port)
		}
	}
}
func (r *EtcdMigrateServer) Stop() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if r.cmd == nil {
		return fmt.Errorf("cannot stop EtcdMigrateServer that has not been started")
	}
	err := r.cmd.Process.Signal(os.Interrupt)
	if err != nil {
		return fmt.Errorf("error sending SIGINT to etcd for graceful shutdown: %v", err)
	}
	gracefulWait := time.Minute * 2
	stopped := make(chan bool)
	timedout := make(chan bool)
	go func() {
		time.Sleep(gracefulWait)
		timedout <- true
	}()
	go func() {
		select {
		case <-stopped:
			return
		case <-timedout:
			klog.Infof("etcd server has not terminated gracefully after %s, killing it.", gracefulWait)
			r.cmd.Process.Kill()
			return
		}
	}()
	err = r.cmd.Wait()
	stopped <- true
	if exiterr, ok := err.(*exec.ExitError); ok {
		klog.Infof("etcd server stopped (signal: %s)", exiterr.Error())
	} else if err != nil {
		return fmt.Errorf("error waiting for etcd to stop: %v", err)
	}
	klog.Infof("Stopped etcd server %s", r.cfg.name)
	return nil
}
