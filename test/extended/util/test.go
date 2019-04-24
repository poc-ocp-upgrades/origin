package util

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/ginkgo/types"
	"github.com/onsi/gomega"
	"k8s.io/klog"
	kapiv1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	kclientset "k8s.io/client-go/kubernetes"
	rbacv1client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	e2e "k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/testfiles"
	"k8s.io/kubernetes/test/e2e/generated"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
	"github.com/openshift/origin/pkg/oc/cli/admin/policy"
	securityclient "github.com/openshift/origin/pkg/security/generated/internalclientset"
	"github.com/openshift/origin/pkg/version"
	testutil "github.com/openshift/origin/test/util"
)

var (
	reportFileName	string
	syntheticSuite	string
	quiet		bool
)
var TestContext *e2e.TestContextType = &e2e.TestContext

func Init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.StringVar(&syntheticSuite, "suite", "", "DEPRECATED: Optional suite selector to filter which tests are run. Use focus.")
	e2e.ViperizeFlags()
	InitTest()
}
func InitStandardFlags() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.RegisterCommonFlags()
	e2e.RegisterClusterFlags()
	e2e.RegisterStorageFlags()
}
func InitTest() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	InitDefaultEnvironmentVariables()
	ginkgo.BeforeEach(checkSyntheticInput)
	TestContext.DeleteNamespace = os.Getenv("DELETE_NAMESPACE") != "false"
	TestContext.VerifyServiceAccount = true
	testfiles.AddFileSource(testfiles.BindataFileSource{Asset: generated.Asset, AssetNames: generated.AssetNames})
	TestContext.KubectlPath = "kubectl"
	TestContext.KubeConfig = KubeConfigPath()
	os.Setenv("KUBECONFIG", TestContext.KubeConfig)
	TestContext.NodeOSDistro = "custom"
	TestContext.MasterOSDistro = "custom"
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&clientcmd.ClientConfigLoadingRules{ExplicitPath: TestContext.KubeConfig}, &clientcmd.ConfigOverrides{})
	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		FatalErr(err)
	}
	TestContext.Host = cfg.Host
	reportFileName = os.Getenv("TEST_REPORT_FILE_NAME")
	if reportFileName == "" {
		reportFileName = "junit"
	}
	quiet = os.Getenv("TEST_OUTPUT_QUIET") == "true"
	TestContext.CreateTestingNS = createTestingNS
	klog.V(2).Infof("Extended test version %s", version.Get().String())
}
func ExecuteTest(t ginkgo.GinkgoTestingT, suite string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var r []ginkgo.Reporter
	if dir := os.Getenv("TEST_REPORT_DIR"); len(dir) > 0 {
		TestContext.ReportDir = dir
	}
	if TestContext.ReportDir != "" {
		if err := os.MkdirAll(TestContext.ReportDir, 0755); err != nil {
			klog.Errorf("Failed creating report directory: %v", err)
		}
		defer e2e.CoreDump(TestContext.ReportDir)
	}
	if config.GinkgoConfig.FocusString == "" && config.GinkgoConfig.SkipString == "" {
		config.GinkgoConfig.SkipString = "Skipped"
	}
	gomega.RegisterFailHandler(ginkgo.Fail)
	if TestContext.ReportDir != "" {
		r = append(r, reporters.NewJUnitReporter(path.Join(TestContext.ReportDir, fmt.Sprintf("%s_%02d.xml", reportFileName, config.GinkgoConfig.ParallelNode))))
	}
	AnnotateTestSuite()
	if quiet {
		r = append(r, NewSimpleReporter())
		ginkgo.RunSpecsWithCustomReporters(t, suite, r)
	} else {
		ginkgo.RunSpecsWithDefaultAndCustomReporters(t, suite, r)
	}
}
func AnnotateTestSuite() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var allLabels []string
	matches := make(map[string]*regexp.Regexp)
	stringMatches := make(map[string][]string)
	excludes := make(map[string]*regexp.Regexp)
	for label, items := range testMaps {
		sort.Strings(items)
		allLabels = append(allLabels, label)
		var remain []string
		for _, item := range items {
			re := regexp.MustCompile(item)
			if p, ok := re.LiteralPrefix(); ok {
				stringMatches[label] = append(stringMatches[label], p)
			} else {
				remain = append(remain, item)
			}
		}
		if len(remain) > 0 {
			matches[label] = regexp.MustCompile(strings.Join(remain, `|`))
		}
	}
	for label, items := range labelExcludes {
		sort.Strings(items)
		excludes[label] = regexp.MustCompile(strings.Join(items, `|`))
	}
	sort.Strings(allLabels)
	ginkgo.WalkTests(func(name string, node types.TestNode) {
		labels := ""
		for {
			count := 0
			for _, label := range allLabels {
				if strings.Contains(name, label) {
					continue
				}
				var hasLabel bool
				for _, segment := range stringMatches[label] {
					hasLabel = strings.Contains(name, segment)
					if hasLabel {
						break
					}
				}
				if !hasLabel {
					if re := matches[label]; re != nil {
						hasLabel = matches[label].MatchString(name)
					}
				}
				if hasLabel {
					if re, ok := excludes[label]; ok && re.MatchString(name) {
						continue
					}
					count++
					labels += " " + label
					name += " " + label
				}
			}
			if count == 0 {
				break
			}
		}
		if !excludedTestsFilter.MatchString(name) {
			isSerial := strings.Contains(name, "[Serial]")
			isConformance := strings.Contains(name, "[Conformance]")
			switch {
			case isSerial && isConformance:
				node.SetText(node.Text() + " [Suite:openshift/conformance/serial/minimal]")
			case isSerial:
				node.SetText(node.Text() + " [Suite:openshift/conformance/serial]")
			case isConformance:
				node.SetText(node.Text() + " [Suite:openshift/conformance/parallel/minimal]")
			default:
				node.SetText(node.Text() + " [Suite:openshift/conformance/parallel]")
			}
		}
		if strings.Contains(node.CodeLocation().FileName, "/origin/test/") && !strings.Contains(node.Text(), "[Suite:openshift") {
			node.SetText(node.Text() + " [Suite:openshift]")
		}
		if strings.Contains(node.CodeLocation().FileName, "/kubernetes/test/e2e/") {
			node.SetText(node.Text() + " [Suite:k8s]")
		}
		node.SetText(node.Text() + labels)
	})
}
func InitDefaultEnvironmentVariables() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ad := os.Getenv("ARTIFACT_DIR"); len(strings.TrimSpace(ad)) == 0 {
		os.Setenv("ARTIFACT_DIR", filepath.Join(os.TempDir(), "artifacts"))
	}
}
func isPackage(pkg string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(ginkgo.CurrentGinkgoTestDescription().FileName, pkg)
}
func isOriginTest() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isPackage("/origin/test/")
}
func isKubernetesE2ETest() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isPackage("/kubernetes/test/e2e/")
}
func testNameContains(name string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(ginkgo.CurrentGinkgoTestDescription().FullTestText, name)
}
func isOriginUpgradeTest() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isPackage("/origin/test/e2e/upgrade/")
}
func skipTestNamespaceCustomization() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (isPackage("/kubernetes/test/e2e/namespace.go") && (testNameContains("should always delete fast") || testNameContains("should delete fast enough")))
}
func createTestingNS(baseName string, c kclientset.Interface, labels map[string]string) (*kapiv1.Namespace, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns, err := e2e.CreateTestingNS(baseName, c, labels)
	if err != nil {
		return ns, err
	}
	klog.V(2).Infof("blah=%s", ginkgo.CurrentGinkgoTestDescription().FileName)
	if (isKubernetesE2ETest() && !skipTestNamespaceCustomization()) || isOriginUpgradeTest() {
		clientConfig, err := testutil.GetClusterAdminClientConfig(KubeConfigPath())
		if err != nil {
			return ns, err
		}
		securityClient, err := securityclient.NewForConfig(clientConfig)
		if err != nil {
			return ns, err
		}
		e2e.Logf("About to run a Kube e2e test, ensuring namespace is privileged")
		addE2EServiceAccountsToSCC(securityClient, []kapiv1.Namespace{*ns}, "privileged")
		addE2EServiceAccountsToSCC(securityClient, []kapiv1.Namespace{*ns}, "anyuid")
		addE2EServiceAccountsToSCC(securityClient, []kapiv1.Namespace{*ns}, "hostmount-anyuid")
		rbacClient, err := rbacv1client.NewForConfig(clientConfig)
		if err != nil {
			return ns, err
		}
		addRoleToE2EServiceAccounts(rbacClient, []kapiv1.Namespace{*ns}, bootstrappolicy.ViewRoleName)
		allowAllNodeScheduling(c, ns.Name)
	}
	return ns, err
}

var (
	testMaps		= map[string][]string{"[Local]": {`\[Feature:ImagePrune\]`}, "[Disabled:Alpha]": {`\[Feature:Initializers\]`, `\[Feature:PodPreemption\]`, `\[Feature:RunAsGroup\]`, `\[NodeAlphaFeature:VolumeSubpathEnvExpansion\]`, `AdmissionWebhook`, `\[NodeAlphaFeature:NodeLease\]`, `\[Feature:TTLAfterFinished\]`, `\[Feature:GPUDevicePlugin\]`}, "[Disabled:Unimplemented]": {`\[Feature:Networking-IPv6\]`, `Monitoring`, `Cluster level logging`, `Kibana`, `Ubernetes`, `kube-ui`, `Kubernetes Dashboard`, `\[Feature:ServiceLoadBalancer\]`, `PersistentVolumes-local`, `\[Feature:RuntimeClass\]`, `\[Feature:CustomResourceWebhookConversion\]`, `NetworkPolicy between server and client should allow egress access on one named port`, `should proxy to cadvisor`}, "[Disabled:SpecialConfig]": {`\[Feature:ImageQuota\]`, `\[Feature:Audit\]`, `\[Feature:LocalStorageCapacityIsolation\]`, `kube-dns-autoscaler`, `should check if Kubernetes master services is included in cluster-info`, `DNS configMap`, `vsphere`, `Cinder`, `Ceph RBD`, `GlusterFS`, `Horizontal pod autoscaling`, `authentication: OpenLDAP`, `NodeProblemDetector`, `Advanced Audit should audit API calls`, `Metadata Concealment`, `Firewall rule should have correct firewall rules for e2e cluster`}, "[Disabled:Broken]": {`\[Feature:BlockVolume\]`, `\[Feature:Example\]`, `mount an API token into pods`, `ServiceAccounts should ensure a single API token exists`, `should test kube-proxy`, `unchanging, static URL paths for kubernetes api services`, "PersistentVolumes NFS when invoking the Recycle reclaim policy", `should propagate mounts to the host`, `Simple pod should handle in-cluster config`, `Services should be able to up and down services`, `Network should set TCP CLOSE_WAIT timeout`, `should allow ingress access on one named port`, `should answer endpoint and wildcard queries for the cluster`, `\[NodeFeature:Sysctls\]`, `validates that there is no conflict between pods with same hostPort but different hostIP and protocol`, `Pod should perfer to scheduled to nodes pod can tolerate`, `Services should be able to create a functioning NodePort service`, `SSH`, `SELinux relabeling`, `Volumes CephFS`, `should support inline execution and attach`, `should idle the service and DeploymentConfig properly`, `\[Feature:Volumes\]`, `\[Driver: csi-hostpath`, `\[Driver: nfs\] \[Testpattern: Pre-provisioned PV \(block volmode\)\] volumeMode should fail to create pod by failing to mount volume`, `\[Driver: aws\] \[Testpattern: Dynamic PV \(block volmode\)\] volumeMode should create sc, pod, pv, and pvc, read/write to the pv, and delete all created resources`, `\[Feature:NodeAuthenticator\]`, `PreemptionExecutionPath`, `\[Volume type: blockfswithoutformat\]`, `CSI Volumes CSI attach test using HostPath driver`, `CSI Volumes CSI plugin test using CSI driver: hostPath`, `Volume metrics should create volume metrics in Volume Manager`}, "[Slow]": {`\[sig-scalability\]`, `should create and stop a working application`, `\[Feature:PerformanceDNS\]`, `should ensure that critical pod is scheduled in case there is no resources available`, "Pod should avoid to schedule to node that have avoidPod annotation", "Pod should be schedule to node that satisify the PodAffinity", "Pod should be prefer scheduled to node that satisify the NodeAffinity", "Pod should be schedule to node that don't match the PodAntiAffinity terms", `validates that there exists conflict between pods with same hostPort and protocol but one using 0\.0\.0\.0 hostIP`}, "[Flaky]": {`Job should run a job to completion when tasks sometimes fail and are not locally restarted`, `openshift mongodb replication creating from a template`, `should use be able to process many pods and reuse local volumes`}, "[Serial]": {`\[Disruptive\]`, `\[Feature:Performance\]`, `\[Feature:ManualPerformance\]`, `\[Feature:HighDensityPerformance\]`, `Service endpoints latency`, `Clean up pods on node`, `should allow starting 95 pods per node`, `DynamicProvisioner should test that deleting a claim before the volume is provisioned deletes the volume`, `Should be able to support the 1\.7 Sample API Server using the current Aggregator`}, "[Suite:openshift/scalability]": {}}
	labelExcludes		= map[string][]string{}
	excludedTests		= []string{`\[Disabled:`, `\[Disruptive\]`, `\[Skipped\]`, `\[Slow\]`, `\[Flaky\]`, `\[local\]`, `\[Local\]`}
	excludedTestsFilter	= regexp.MustCompile(strings.Join(excludedTests, `|`))
)

func checkSyntheticInput() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkSuiteSkips()
}
func checkSuiteSkips() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case isOriginTest():
		if strings.Contains(config.GinkgoConfig.SkipString, "Synthetic Origin") {
			ginkgo.Skip("skipping all openshift/origin tests")
		}
	case isKubernetesE2ETest():
		if strings.Contains(config.GinkgoConfig.SkipString, "Synthetic Kubernetes") {
			ginkgo.Skip("skipping all k8s.io/kubernetes tests")
		}
	}
}

var longRetry = wait.Backoff{Steps: 100}

func allowAllNodeScheduling(c kclientset.Interface, namespace string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := retry.RetryOnConflict(longRetry, func() error {
		ns, err := c.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if ns.Annotations == nil {
			ns.Annotations = make(map[string]string)
		}
		ns.Annotations["openshift.io/node-selector"] = ""
		_, err = c.CoreV1().Namespaces().Update(ns)
		return err
	})
	if err != nil {
		FatalErr(err)
	}
}
func addE2EServiceAccountsToSCC(securityClient securityclient.Interface, namespaces []kapiv1.Namespace, sccName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := retry.RetryOnConflict(longRetry, func() error {
		scc, err := securityClient.Security().SecurityContextConstraints().Get(sccName, metav1.GetOptions{})
		if err != nil {
			if apierrs.IsNotFound(err) {
				return nil
			}
			return err
		}
		for _, ns := range namespaces {
			if strings.HasPrefix(ns.Name, "e2e-") {
				scc.Groups = append(scc.Groups, fmt.Sprintf("system:serviceaccounts:%s", ns.Name))
			}
		}
		if _, err := securityClient.Security().SecurityContextConstraints().Update(scc); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		FatalErr(err)
	}
}
func addRoleToE2EServiceAccounts(rbacClient rbacv1client.RbacV1Interface, namespaces []kapiv1.Namespace, roleName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := retry.RetryOnConflict(longRetry, func() error {
		for _, ns := range namespaces {
			if strings.HasPrefix(ns.Name, "e2e-") && ns.Status.Phase != kapiv1.NamespaceTerminating {
				sa := fmt.Sprintf("system:serviceaccount:%s:default", ns.Name)
				addRole := &policy.RoleModificationOptions{RoleBindingNamespace: ns.Name, RoleKind: "ClusterRole", RoleName: roleName, RbacClient: rbacClient, Users: []string{sa}, PrintFlags: genericclioptions.NewPrintFlags(""), ToPrinter: func(string) (printers.ResourcePrinter, error) {
					return printers.NewDiscardingPrinter(), nil
				}}
				if err := addRole.AddRole(); err != nil {
					e2e.Logf("Warning: Failed to add role to e2e service account: %v", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		FatalErr(err)
	}
}
