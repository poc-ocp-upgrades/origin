package db

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/openshift/origin/test/extended/util"
)

type PostgreSQL struct {
	podName		string
	masterPodName	string
}

func NewPostgreSQL(podName, masterPodName string) util.Database {
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
	if masterPodName == "" {
		masterPodName = podName
	}
	return &PostgreSQL{podName: podName, masterPodName: masterPodName}
}
func (m PostgreSQL) PodName() string {
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
	return m.podName
}
func (m PostgreSQL) IsReady(oc *util.CLI) (bool, error) {
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
	conf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return false, err
	}
	out, err := oc.Run("exec").Args(m.podName, "-c", conf.Container, "--", "bash", "-c", "psql postgresql://postgres@127.0.0.1 -x -c \"SELECT 1;\"").Output()
	if err != nil {
		switch err.(type) {
		case *util.ExitError, *exec.ExitError:
			return false, nil
		default:
			return false, err
		}
	}
	return strings.Contains(out, "-[ RECORD 1 ]\n?column? | 1"), nil
}
func (m PostgreSQL) Query(oc *util.CLI, query string) (string, error) {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return "", err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return "", err
	}
	return oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("psql postgres://%s:%s@127.0.0.1/%s -x -c \"%s\"", masterConf.Env["POSTGRESQL_USER"], masterConf.Env["POSTGRESQL_PASSWORD"], masterConf.Env["POSTGRESQL_DATABASE"], query)).Output()
}
func (m PostgreSQL) QueryPrivileged(oc *util.CLI, query string) (string, error) {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return "", err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return "", err
	}
	return oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("psql postgres://postgres:%s@127.0.0.1/%s -x -c \"%s\"", masterConf.Env["POSTGRESQL_ADMIN_PASSWORD"], masterConf.Env["POSTGRESQL_DATABASE"], query)).Output()
}
func (m PostgreSQL) TestRemoteLogin(oc *util.CLI, hostAddress string) error {
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
	container, err := firstContainerName(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.podName)
	if err != nil {
		return err
	}
	masterConf, err := getPodConfig(oc.KubeClient().CoreV1().Pods(oc.Namespace()), m.masterPodName)
	if err != nil {
		return err
	}
	err = oc.Run("exec").Args(m.podName, "-c", container, "--", "bash", "-c", fmt.Sprintf("psql postgres://%s:%s@%s/%s -x -c \"SELECT 1;\"", masterConf.Env["POSTGRESQL_USER"], masterConf.Env["POSTGRESQL_PASSWORD"], hostAddress, masterConf.Env["POSTGRESQL_DATABASE"])).Execute()
	return err
}
