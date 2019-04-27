package db

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/openshift/origin/test/extended/util"
)

type MySQL struct {
	podName		string
	masterPodName	string
}

func NewMysql(podName, masterPodName string) util.Database {
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
	if masterPodName == "" {
		masterPodName = podName
	}
	return &MySQL{podName: podName, masterPodName: masterPodName}
}
func (m MySQL) PodName() string {
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
	return m.podName
}
func (m MySQL) IsReady(oc *util.CLI) (bool, error) {
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
	conf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return false, err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return false, err
	}
	out, err := oc.Run("exec").Args(m.podName, "-c", conf.Container, "--", "bash", "-c", fmt.Sprintf("mysqladmin -h localhost -u%s -p%s ping", masterConf.Env["MYSQL_USER"], masterConf.Env["MYSQL_PASSWORD"])).Output()
	if err != nil {
		switch err.(type) {
		case *util.ExitError, *exec.ExitError:
			return false, nil
		default:
			return false, err
		}
	}
	return strings.Contains(out, "mysqld is alive"), nil
}
func (m MySQL) Query(oc *util.CLI, query string) (string, error) {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return "", err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return "", err
	}
	return oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("mysql -h 127.0.0.1 -u%s -p%s -e \"%s\" %s", masterConf.Env["MYSQL_USER"], masterConf.Env["MYSQL_PASSWORD"], query, masterConf.Env["MYSQL_DATABASE"])).Output()
}
func (m MySQL) QueryPrivileged(oc *util.CLI, query string) (string, error) {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return "", err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return "", err
	}
	return oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("mysql -h 127.0.0.1 -uroot -e \"%s\" %s", query, masterConf.Env["MYSQL_DATABASE"])).Output()
}
func (m MySQL) TestRemoteLogin(oc *util.CLI, hostAddress string) error {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return err
	}
	err = oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("mysql -h %s -u%s -p%s -e \"SELECT 1;\" %s", hostAddress, masterConf.Env["MYSQL_USER"], masterConf.Env["MYSQL_PASSWORD"], masterConf.Env["MYSQL_DATABASE"])).Execute()
	return err
}
