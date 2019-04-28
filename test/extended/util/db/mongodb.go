package db

import (
	"errors"
	"fmt"
	"github.com/openshift/origin/test/extended/util"
)

type MongoDB struct{ podName string }

func NewMongoDB(podName string) util.Database {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MongoDB{podName: podName}
}
func (m MongoDB) PodName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.podName
}
func (m MongoDB) IsReady(oc *util.CLI) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isReady(oc, m.podName, `mongo --quiet --eval '{"ping", 1}'`, "1")
}
func (m MongoDB) Query(oc *util.CLI, query string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return executeShellCommand(oc, m.podName, fmt.Sprintf(`mongo --quiet "$MONGODB_DATABASE" --username "$MONGODB_USER" --password "$MONGODB_PASSWORD" --eval '%s'`, query))
}
func (m MongoDB) QueryPrivileged(oc *util.CLI, query string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", errors.New("not implemented")
}
func (m MongoDB) TestRemoteLogin(oc *util.CLI, hostAddress string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return errors.New("not implemented")
}
func (m MongoDB) QueryPrimary(oc *util.CLI, query string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return executeShellCommand(oc, m.podName, fmt.Sprintf(`mongo --quiet "$MONGODB_DATABASE" --username "$MONGODB_USER" --password "$MONGODB_PASSWORD" --host "$MONGODB_REPLICA_NAME/localhost" --eval '%s'`, query))
}
