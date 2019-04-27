package image_ecosystem

import (
	"fmt"
	"time"
	g "github.com/onsi/ginkgo"
	exutil "github.com/openshift/origin/test/extended/util"
	dbutil "github.com/openshift/origin/test/extended/util/db"
)

func readRecordFromPod(oc *exutil.CLI, podName string) error {
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
	findCmd := "rs.slaveOk(); printjson(db.test.find({}, {_id: 0}).toArray())"
	fmt.Fprintf(g.GinkgoWriter, "DEBUG: reading record from the pod %v\n", podName)
	mongoPod := dbutil.NewMongoDB(podName)
	return exutil.WaitForQueryOutputContains(oc, mongoPod, 1*time.Minute, false, findCmd, `{ "status" : "passed" }`)
}
