package builds

import (
	"fmt"
	"strings"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/wait"
	exutil "github.com/openshift/origin/test/extended/util"
)

const (
	buildPrefixTS	= "ruby-sample-build-ts"
	buildPrefixTD	= "ruby-sample-build-td"
	buildPrefixTC	= "ruby-sample-build-tc"
)

func scrapeLogs(bldPrefix string, oc *exutil.CLI) {
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
	br, err := exutil.StartBuildAndWait(oc, bldPrefix)
	o.ExpectWithOffset(1, err).NotTo(o.HaveOccurred())
	out, err := br.Logs()
	o.Expect(err).NotTo(o.HaveOccurred())
	lines := strings.Split(out, "\n")
	found := false
	for _, line := range lines {
		if strings.Contains(line, "Pulling image") && strings.Contains(line, "ruby") {
			fmt.Fprintf(g.GinkgoWriter, "\n\nfound pull image line %s\n\n", line)
			found = true
			break
		}
	}
	if !found {
		fmt.Fprintf(g.GinkgoWriter, "\n\n build log dump on failed test:  %s\n\n", out)
		o.Expect(found).To(o.BeTrue())
	}
}
func checkPodFlag(bldPrefix string, oc *exutil.CLI) {
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
	_, err := exutil.StartBuildAndWait(oc, bldPrefix)
	o.ExpectWithOffset(1, err).NotTo(o.HaveOccurred())
	out, err := oc.Run("get").Args("pods", bldPrefix+"-1-build", "-o", "jsonpath='{.spec.containers[0].imagePullPolicy}'").Output()
	o.Expect(err).NotTo(o.HaveOccurred())
	o.Expect(out).To(o.Equal("'Always'"))
}

var _ = g.Describe("[Feature:Builds] forcePull should affect pulling builder images", func() {
	defer g.GinkgoRecover()
	var oc = exutil.NewCLI("forcepull", exutil.KubeConfigPath())
	g.Context("", func() {
		g.BeforeEach(func() {
			exutil.PreTestDump()
			g.By("granting system:build-strategy-custom")
			binding := fmt.Sprintf("custombuildaccess-%s", oc.Username())
			err := oc.AsAdmin().Run("create").Args("clusterrolebinding", binding, "--clusterrole", "system:build-strategy-custom", "--user", oc.Username()).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			err = wait.PollImmediate(time.Second, time.Minute, func() (bool, error) {
				err := oc.Run("auth").Args("can-i", "--all-namespaces", "--quiet", "create", "builds.build.openshift.io", "--subresource=custom").Execute()
				if err != nil {
					return false, nil
				}
				return true, nil
			})
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("waiting for openshift/ruby:latest ImageStreamTag")
			err = exutil.WaitForAnImageStreamTag(oc, "openshift", "ruby", "latest")
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("create application build configs for 3 strategies")
			apps := exutil.FixturePath("testdata", "forcepull-test.json")
			err = exutil.CreateResource(apps, oc)
			o.Expect(err).NotTo(o.HaveOccurred())
		})
		g.AfterEach(func() {
			binding := fmt.Sprintf("custombuildaccess-%s", oc.Username())
			err := oc.AsAdmin().Run("delete").Args("clusterrolebinding", binding).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			if g.CurrentGinkgoTestDescription().Failed {
				exutil.DumpPodStates(oc)
				exutil.DumpConfigMapStates(oc)
				exutil.DumpPodLogsStartingWith("", oc)
			}
		})
		g.It("ForcePull test case execution s2i", func() {
			g.Skip("TODO: force pull is moot until/unless we go back to sharing the image filesystem")
			g.By("when s2i force pull is true")
			scrapeLogs(buildPrefixTS, oc)
			scrapeLogs(buildPrefixTS, oc)
		})
		g.It("ForcePull test case execution docker", func() {
			g.Skip("TODO: force pull is moot until/unless we go back to sharing the image filesystem")
			g.By("docker when force pull is true")
			scrapeLogs(buildPrefixTD, oc)
			scrapeLogs(buildPrefixTD, oc)
		})
		g.It("ForcePull test case execution custom", func() {
			g.Skip("TODO: force pull is moot until/unless we go back to sharing the image filesystem")
			g.By("when custom force pull is true")
			checkPodFlag(buildPrefixTC, oc)
		})
	})
})
