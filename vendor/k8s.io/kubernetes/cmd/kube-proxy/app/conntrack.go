package app

import (
	"errors"
	goformat "fmt"
	"io/ioutil"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/util/mount"
	"k8s.io/kubernetes/pkg/util/sysctl"
	goos "os"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	gotime "time"
)

type Conntracker interface {
	SetMax(max int) error
	SetTCPEstablishedTimeout(seconds int) error
	SetTCPCloseWaitTimeout(seconds int) error
}
type realConntracker struct{}

var readOnlySysFSError = errors.New("readOnlySysFS")

func (rct realConntracker) SetMax(max int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := rct.setIntSysCtl("nf_conntrack_max", max); err != nil {
		return err
	}
	klog.Infof("Setting nf_conntrack_max to %d", max)
	hashsize, err := readIntStringFile("/sys/module/nf_conntrack/parameters/hashsize")
	if err != nil {
		return err
	}
	if hashsize >= (max / 4) {
		return nil
	}
	writable, err := isSysFSWritable()
	if err != nil {
		return err
	}
	if !writable {
		return readOnlySysFSError
	}
	klog.Infof("Setting conntrack hashsize to %d", max/4)
	return writeIntStringFile("/sys/module/nf_conntrack/parameters/hashsize", max/4)
}
func (rct realConntracker) SetTCPEstablishedTimeout(seconds int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rct.setIntSysCtl("nf_conntrack_tcp_timeout_established", seconds)
}
func (rct realConntracker) SetTCPCloseWaitTimeout(seconds int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return rct.setIntSysCtl("nf_conntrack_tcp_timeout_close_wait", seconds)
}
func (realConntracker) setIntSysCtl(name string, value int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	entry := "net/netfilter/" + name
	sys := sysctl.New()
	if val, _ := sys.GetSysctl(entry); val != value {
		klog.Infof("Set sysctl '%v' to %v", entry, value)
		if err := sys.SetSysctl(entry, value); err != nil {
			return err
		}
	}
	return nil
}
func isSysFSWritable() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	const permWritable = "rw"
	const sysfsDevice = "sysfs"
	m := mount.New("")
	mountPoints, err := m.List()
	if err != nil {
		klog.Errorf("failed to list mount points: %v", err)
		return false, err
	}
	for _, mountPoint := range mountPoints {
		if mountPoint.Type != sysfsDevice {
			continue
		}
		if len(mountPoint.Opts) > 0 && mountPoint.Opts[0] == permWritable {
			return true, nil
		}
		klog.Errorf("sysfs is not writable: %+v (mount options are %v)", mountPoint, mountPoint.Opts)
		return false, readOnlySysFSError
	}
	return false, errors.New("No sysfs mounted")
}
func readIntStringFile(filename string) (int, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(strings.TrimSpace(string(b)))
}
func writeIntStringFile(filename string, value int) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return ioutil.WriteFile(filename, []byte(strconv.Itoa(value)), 0640)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
