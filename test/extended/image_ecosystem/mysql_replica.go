package image_ecosystem

import (
	"fmt"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	"github.com/openshift/api/template"
	exutil "github.com/openshift/origin/test/extended/util"
	"github.com/openshift/origin/test/extended/util/db"
	testutil "github.com/openshift/origin/test/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kcoreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

type testCase struct {
	Version		string
	TemplatePath	string
	TemplateName	string
	SkipReplication	bool
}

var (
	testCases	= []testCase{{"5.7", "https://raw.githubusercontent.com/sclorg/mysql-container/master/examples/replica/mysql_replica.json", "mysql-replication-example", false}}
	helperTemplate	= exutil.FixturePath("..", "..", "examples", "db-templates", "mysql-ephemeral-template.json")
	helperName	= "mysql-helper"
)

func CreateMySQLReplicationHelpers(c kcoreclient.PodInterface, masterDeployment, slaveDeployment, helperDeployment string, slaveCount int) (exutil.Database, []exutil.Database, exutil.Database) {
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
	podNames, err := exutil.WaitForPods(c, exutil.ParseLabelsOrDie(fmt.Sprintf("deployment=%s", masterDeployment)), exutil.CheckPodIsRunning, 1, 4*time.Minute)
	o.Expect(err).NotTo(o.HaveOccurred())
	masterPod := podNames[0]
	slavePods, err := exutil.WaitForPods(c, exutil.ParseLabelsOrDie(fmt.Sprintf("deployment=%s", slaveDeployment)), exutil.CheckPodIsRunning, slaveCount, 6*time.Minute)
	o.Expect(err).NotTo(o.HaveOccurred())
	master := db.NewMysql(masterPod, "")
	slaves := make([]exutil.Database, len(slavePods))
	for i := range slavePods {
		slave := db.NewMysql(slavePods[i], masterPod)
		slaves[i] = slave
	}
	helperNames, err := exutil.WaitForPods(c, exutil.ParseLabelsOrDie(fmt.Sprintf("deployment=%s", helperDeployment)), exutil.CheckPodIsRunning, 1, 4*time.Minute)
	o.Expect(err).NotTo(o.HaveOccurred())
	helper := db.NewMysql(helperNames[0], masterPod)
	return master, slaves, helper
}
func replicationTestFactory(oc *exutil.CLI, tc testCase, cleanup func()) func() {
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
	return func() {
		defer cleanup()
		err := testutil.WaitForPolicyUpdate(oc.KubeClient().AuthorizationV1(), oc.Namespace(), "create", template.Resource("templates"), true)
		o.Expect(err).NotTo(o.HaveOccurred())
		exutil.WaitForOpenShiftNamespaceImageStreams(oc)
		err = oc.Run("create").Args("-f", tc.TemplatePath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("new-app").Args("--template", tc.TemplateName).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("new-app").Args("-f", helperTemplate, "-p", fmt.Sprintf("MYSQL_VERSION=%s", tc.Version), "-p", fmt.Sprintf("DATABASE_SERVICE_NAME=%s", helperName)).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		g.By("waiting for the deployment to complete")
		err = exutil.WaitForDeploymentConfig(oc.KubeClient(), oc.AppsClient().AppsV1(), oc.Namespace(), helperName, 1, true, oc)
		o.Expect(err).NotTo(o.HaveOccurred())
		g.By("waiting for an endpoint")
		err = e2e.WaitForEndpoint(oc.KubeFramework().ClientSet, oc.Namespace(), helperName)
		o.Expect(err).NotTo(o.HaveOccurred())
		tableCounter := 0
		assertReplicationIsWorking := func(masterDeployment, slaveDeployment string, slaveCount int) (exutil.Database, []exutil.Database, exutil.Database) {
			tableCounter++
			table := fmt.Sprintf("table_%0.2d", tableCounter)
			g.By("creating replication helpers")
			master, slaves, helper := CreateMySQLReplicationHelpers(oc.KubeClient().CoreV1().Pods(oc.Namespace()), masterDeployment, slaveDeployment, fmt.Sprintf("%s-1", helperName), slaveCount)
			o.Expect(exutil.WaitUntilAllHelpersAreUp(oc, []exutil.Database{master})).NotTo(o.HaveOccurred())
			o.Expect(exutil.WaitUntilAllHelpersAreUp(oc, slaves)).NotTo(o.HaveOccurred())
			g.By("wait for mysql-master endpoint")
			err = e2e.WaitForEndpoint(oc.KubeFramework().ClientSet, oc.Namespace(), "mysql-master")
			o.Expect(err).NotTo(o.HaveOccurred())
			err := helper.TestRemoteLogin(oc, "mysql-master")
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("create new table")
			_, err = master.Query(oc, fmt.Sprintf("CREATE TABLE %s (col1 VARCHAR(20), col2 VARCHAR(20));", table))
			o.Expect(err).NotTo(o.HaveOccurred())
			_, err = master.Query(oc, fmt.Sprintf("INSERT INTO %s (col1, col2) VALUES ('val1', 'val2');", table))
			o.Expect(err).NotTo(o.HaveOccurred())
			err = exutil.WaitForQueryOutputContains(oc, master, 10*time.Second, false, fmt.Sprintf("SELECT * FROM %s\\G;", table), "col1: val1\ncol2: val2")
			o.Expect(err).NotTo(o.HaveOccurred())
			for _, slave := range slaves {
				err = exutil.WaitForQueryOutputContains(oc, slave, 90*time.Second, false, fmt.Sprintf("SELECT * FROM %s\\G;", table), "col1: val1\ncol2: val2")
				o.Expect(err).NotTo(o.HaveOccurred())
			}
			return master, slaves, helper
		}
		g.By("after initial deployment")
		master, _, _ := assertReplicationIsWorking("mysql-master-1", "mysql-slave-1", 1)
		if tc.SkipReplication {
			return
		}
		g.By("after master is restarted by changing the Deployment Config")
		err = oc.Run("set", "env").Args("dc", "mysql-master", "MYSQL_ROOT_PASSWORD=newpass").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = exutil.WaitUntilPodIsGone(oc.KubeClient().CoreV1().Pods(oc.Namespace()), master.PodName(), 2*time.Minute)
		if err != nil {
			e2e.Logf("Checking if pod %s still exists", master.PodName())
			oc.Run("get").Args("pod", master.PodName(), "-o", "yaml").Execute()
		}
		master, _, _ = assertReplicationIsWorking("mysql-master-2", "mysql-slave-1", 1)
		g.By("after master is restarted by deleting the pod")
		err = oc.Run("delete").Args("pod", "-l", "deployment=mysql-master-2").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = exutil.WaitUntilPodIsGone(oc.KubeClient().CoreV1().Pods(oc.Namespace()), master.PodName(), 2*time.Minute)
		if err != nil {
			e2e.Logf("Checking if pod %s still exists", master.PodName())
			oc.Run("get").Args("pod", master.PodName(), "-o", "yaml").Execute()
		}
		o.Expect(err).NotTo(o.HaveOccurred())
		assertReplicationIsWorking("mysql-master-2", "mysql-slave-1", 1)
		_, slaves, _ := assertReplicationIsWorking("mysql-master-2", "mysql-slave-1", 1)
		g.By("after slave is restarted by deleting the pod")
		err = oc.Run("delete").Args("pod", "-l", "deployment=mysql-slave-1").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = exutil.WaitUntilPodIsGone(oc.KubeClient().CoreV1().Pods(oc.Namespace()), slaves[0].PodName(), 2*time.Minute)
		if err != nil {
			e2e.Logf("Checking if pod %s still exists", slaves[0].PodName())
			oc.Run("get").Args("pod", slaves[0].PodName(), "-o", "yaml").Execute()
		}
		o.Expect(err).NotTo(o.HaveOccurred())
		assertReplicationIsWorking("mysql-master-2", "mysql-slave-1", 1)
		pods, err := oc.KubeClient().CoreV1().Pods(oc.Namespace()).List(metav1.ListOptions{LabelSelector: exutil.ParseLabelsOrDie("deployment=mysql-slave-1").String()})
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(len(pods.Items)).To(o.Equal(1))
	}
}
