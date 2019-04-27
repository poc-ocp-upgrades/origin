package util

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"time"
	g "github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/wait"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

type Database interface {
	PodName() string
	IsReady(oc *CLI) (bool, error)
	Query(oc *CLI, query string) (string, error)
	QueryPrivileged(oc *CLI, query string) (string, error)
	TestRemoteLogin(oc *CLI, hostAddress string) error
}
type ReplicaSet interface {
	QueryPrimary(oc *CLI, query string) (string, error)
}

func WaitForQueryOutputSatisfies(oc *CLI, d Database, timeout time.Duration, admin bool, query string, predicate func(string) bool) error {
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
	err := wait.Poll(5*time.Second, timeout, func() (bool, error) {
		var (
			out	string
			err	error
		)
		if admin {
			out, err = d.QueryPrivileged(oc, query)
		} else {
			out, err = d.Query(oc, query)
		}
		fmt.Fprintf(g.GinkgoWriter, "Query %s result: %s\n", query, out)
		if _, ok := err.(*ExitError); ok {
			return false, nil
		}
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		if err != nil {
			e2e.Logf("failing immediately with error: %v, type=%v", err, reflect.TypeOf(err))
			return false, err
		}
		if predicate(out) {
			return true, nil
		}
		return false, nil
	})
	if err == wait.ErrWaitTimeout {
		return fmt.Errorf("timed out waiting for query: %q", query)
	}
	return err
}
func WaitForQueryOutputContains(oc *CLI, d Database, timeout time.Duration, admin bool, query, resultSubstr string) error {
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
	return WaitForQueryOutputSatisfies(oc, d, timeout, admin, query, func(resultOutput string) bool {
		return strings.Contains(resultOutput, resultSubstr)
	})
}
func WaitUntilUp(oc *CLI, d Database, timeout time.Duration) error {
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
	err := wait.Poll(2*time.Second, timeout, func() (bool, error) {
		return d.IsReady(oc)
	})
	if err == wait.ErrWaitTimeout {
		return fmt.Errorf("timed out waiting for pod %s get up", d.PodName())
	}
	return err
}
func WaitUntilAllHelpersAreUp(oc *CLI, helpers []Database) error {
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
	for _, m := range helpers {
		if err := WaitUntilUp(oc, m, 3*time.Minute); err != nil {
			return err
		}
	}
	return nil
}
