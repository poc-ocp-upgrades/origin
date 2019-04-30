package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kapiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	batchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/kubernetes/pkg/apis/authorization"
	quota "k8s.io/kubernetes/pkg/quota/v1"
	e2e "k8s.io/kubernetes/test/e2e/framework"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	appsv1clienttyped "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	buildv1clienttyped "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	"github.com/openshift/library-go/pkg/git"
	"github.com/openshift/origin/pkg/api/apihelpers"
	appsutil "github.com/openshift/origin/pkg/apps/util"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	imagetypeclientset "github.com/openshift/origin/pkg/image/generated/internalclientset/typed/image/internalversion"
	"github.com/openshift/origin/test/extended/testdata"
)

const pvPrefix = "pv-"
const nfsPrefix = "nfs-"

func WaitForOpenShiftNamespaceImageStreams(oc *CLI) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	langs := []string{"ruby", "nodejs", "perl", "php", "python", "mysql", "postgresql", "mongodb", "jenkins"}
	scan := func() bool {
		for _, lang := range langs {
			e2e.Logf("Checking language %v \n", lang)
			is, err := oc.ImageClient().Image().ImageStreams("openshift").Get(lang, metav1.GetOptions{})
			if err != nil {
				e2e.Logf("ImageStream Error: %#v \n", err)
				return false
			}
			for tag := range is.Spec.Tags {
				e2e.Logf("Checking tag %v \n", tag)
				if _, ok := is.Status.Tags[tag]; !ok {
					e2e.Logf("Tag Error: %#v \n", ok)
					return false
				}
			}
		}
		return true
	}
	success := false
	for i := 0; i < 15; i++ {
		e2e.Logf("Running scan #%v \n", i)
		success = scan()
		if success {
			break
		}
		e2e.Logf("Sleeping for 10 seconds \n")
		time.Sleep(10 * time.Second)
	}
	if success {
		e2e.Logf("Success! \n")
		return nil
	}
	DumpImageStreams(oc)
	DumpSampleOperator(oc)
	return fmt.Errorf("Failed to import expected imagestreams")
}
func DumpImageStreams(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := oc.AsAdmin().Run("get").Args("is", "-n", "openshift", "-o", "yaml", "--config", KubeConfigPath()).Output()
	if err == nil {
		e2e.Logf("\n  imagestreams in openshift namespace: \n%s\n", out)
	} else {
		e2e.Logf("\n  error on getting imagestreams in openshift namespace: %+v\n%#v\n", err, out)
	}
	out, err = oc.AsAdmin().Run("get").Args("is", "-o", "yaml").Output()
	if err == nil {
		e2e.Logf("\n  imagestreams in dynamic test namespace: \n%s\n", out)
	} else {
		e2e.Logf("\n  error on getting imagestreams in dynamic test namespace: %+v\n%#v\n", err, out)
	}
	ids, err := ListImages()
	if err != nil {
		e2e.Logf("\n  got error on docker images %+v\n", err)
	} else {
		for _, id := range ids {
			e2e.Logf(" found local image %s\n", id)
		}
	}
}
func DumpSampleOperator(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := oc.AsAdmin().Run("get").Args("configs.samples.operator.openshift.io", "instance", "-n", "openshift-cluster-samples-operator", "-o", "yaml", "--config", KubeConfigPath()).Output()
	if err == nil {
		e2e.Logf("\n  samples operator CR: \n%s\n", out)
	} else {
		e2e.Logf("\n  error on getting samples operator CR: %+v\n%#v\n", err, out)
	}
	DumpPodLogsStartingWithInNamespace("cluster-samples-operator", "openshift-cluster-samples-operator", oc)
}
func DumpBuildLogs(bc string, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildOutput, err := oc.AsAdmin().Run("logs").Args("-f", "bc/"+bc, "--timestamps").Output()
	if err == nil {
		e2e.Logf("\n\n  build logs : %s\n\n", buildOutput)
	} else {
		e2e.Logf("\n\n  got error on build logs %+v\n\n", err)
	}
	ExamineDiskUsage()
	ExaminePodDiskUsage(oc)
}
func DumpBuilds(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buildOutput, err := oc.AsAdmin().Run("get").Args("builds", "-o", "yaml").Output()
	if err == nil {
		e2e.Logf("\n\n builds yaml:\n%s\n\n", buildOutput)
	} else {
		e2e.Logf("\n\n got error on build yaml dump: %#v\n\n", err)
	}
}
func GetDeploymentConfigPods(oc *CLI, dcName string, version int64) (*kapiv1.PodList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return oc.AdminKubeClient().CoreV1().Pods(oc.Namespace()).List(metav1.ListOptions{LabelSelector: ParseLabelsOrDie(fmt.Sprintf("%s=%s-%d", appsv1.DeployerPodForDeploymentLabel, dcName, version)).String()})
}
func GetApplicationPods(oc *CLI, dcName string) (*kapiv1.PodList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return oc.AdminKubeClient().CoreV1().Pods(oc.Namespace()).List(metav1.ListOptions{LabelSelector: ParseLabelsOrDie(fmt.Sprintf("deploymentconfig=%s", dcName)).String()})
}
func GetStatefulSetPods(oc *CLI, setName string) (*kapiv1.PodList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return oc.AdminKubeClient().CoreV1().Pods(oc.Namespace()).List(metav1.ListOptions{LabelSelector: ParseLabelsOrDie(fmt.Sprintf("name=%s", setName)).String()})
}
func DumpDeploymentLogs(dcName string, version int64, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Dumping deployment logs for deploymentconfig %q\n", dcName)
	pods, err := GetDeploymentConfigPods(oc, dcName, version)
	if err != nil {
		e2e.Logf("Unable to retrieve pods for deploymentconfig %q: %v\n", dcName, err)
		return
	}
	DumpPodLogs(pods.Items, oc)
}
func DumpApplicationPodLogs(dcName string, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Dumping application logs for deploymentconfig %q\n", dcName)
	pods, err := GetApplicationPods(oc, dcName)
	if err != nil {
		e2e.Logf("Unable to retrieve pods for deploymentconfig %q: %v\n", dcName, err)
		return
	}
	DumpPodLogs(pods.Items, oc)
}
func DumpPodStates(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Dumping pod state for namespace %s", oc.Namespace())
	out, err := oc.AsAdmin().Run("get").Args("pods", "-o", "yaml").Output()
	if err != nil {
		e2e.Logf("Error dumping pod states: %v", err)
		return
	}
	e2e.Logf(out)
}
func DumpPodStatesInNamespace(namespace string, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Dumping pod state for namespace %s", namespace)
	out, err := oc.AsAdmin().Run("get").Args("pods", "-n", namespace, "-o", "yaml").Output()
	if err != nil {
		e2e.Logf("Error dumping pod states: %v", err)
		return
	}
	e2e.Logf(out)
}
func DumpPodLogsStartingWith(prefix string, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podsToDump := []kapiv1.Pod{}
	podList, err := oc.AdminKubeClient().CoreV1().Pods(oc.Namespace()).List(metav1.ListOptions{})
	if err != nil {
		e2e.Logf("Error listing pods: %v", err)
		return
	}
	for _, pod := range podList.Items {
		if strings.HasPrefix(pod.Name, prefix) {
			podsToDump = append(podsToDump, pod)
		}
	}
	if len(podsToDump) > 0 {
		DumpPodLogs(podsToDump, oc)
	}
}
func DumpPodLogsStartingWithInNamespace(prefix, namespace string, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podsToDump := []kapiv1.Pod{}
	podList, err := oc.AdminKubeClient().CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		e2e.Logf("Error listing pods: %v", err)
		return
	}
	for _, pod := range podList.Items {
		if strings.HasPrefix(pod.Name, prefix) {
			podsToDump = append(podsToDump, pod)
		}
	}
	if len(podsToDump) > 0 {
		DumpPodLogs(podsToDump, oc)
	}
}
func DumpPodLogs(pods []kapiv1.Pod, oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, pod := range pods {
		descOutput, err := oc.AsAdmin().Run("describe").Args("pod/" + pod.Name).Output()
		if err == nil {
			e2e.Logf("Describing pod %q\n%s\n\n", pod.Name, descOutput)
		} else {
			e2e.Logf("Error retrieving description for pod %q: %v\n\n", pod.Name, err)
		}
		dumpContainer := func(container *kapiv1.Container) {
			depOutput, err := oc.AsAdmin().Run("logs").WithoutNamespace().Args("pod/"+pod.Name, "-c", container.Name, "-n", pod.Namespace).Output()
			if err == nil {
				e2e.Logf("Log for pod %q/%q\n---->\n%s\n<----end of log for %[1]q/%[2]q\n", pod.Name, container.Name, depOutput)
			} else {
				e2e.Logf("Error retrieving logs for pod %q/%q: %v\n\n", pod.Name, container.Name, err)
			}
		}
		for _, c := range pod.Spec.InitContainers {
			dumpContainer(&c)
		}
		for _, c := range pod.Spec.Containers {
			dumpContainer(&c)
		}
	}
}
func DumpPodsCommand(c kubernetes.Interface, ns string, selector labels.Selector, cmd string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podList, err := c.CoreV1().Pods(ns).List(metav1.ListOptions{LabelSelector: selector.String()})
	o.Expect(err).NotTo(o.HaveOccurred())
	values := make(map[string]string)
	for _, pod := range podList.Items {
		stdout, err := e2e.RunHostCmdWithRetries(pod.Namespace, pod.Name, cmd, e2e.StatefulSetPoll, e2e.StatefulPodTimeout)
		o.Expect(err).NotTo(o.HaveOccurred())
		values[pod.Name] = stdout
	}
	for name, stdout := range values {
		stdout = strings.TrimSuffix(stdout, "\n")
		e2e.Logf(name + ": " + strings.Join(strings.Split(stdout, "\n"), fmt.Sprintf("\n%s: ", name)))
	}
}
func DumpConfigMapStates(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Dumping configMap state for namespace %s", oc.Namespace())
	out, err := oc.AsAdmin().Run("get").Args("configmaps", "-o", "yaml").Output()
	if err != nil {
		e2e.Logf("Error dumping configMap states: %v", err)
		return
	}
	e2e.Logf(out)
}
func GetMasterThreadDump(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := oc.AsAdmin().Run("get").Args("--raw", "/debug/pprof/goroutine?debug=2").Output()
	if err == nil {
		e2e.Logf("\n\n Master thread stack dump:\n\n%s\n\n", string(out))
		return
	}
	e2e.Logf("\n\n got error on oc get --raw /debug/pprof/goroutine?godebug=2: %v\n\n", err)
}
func PreTestDump() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func ExamineDiskUsage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return
}
func ExaminePodDiskUsage(oc *CLI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return
}
func VarSubOnFile(srcFile string, destFile string, vars map[string]string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	srcData, err := ioutil.ReadFile(srcFile)
	if err == nil {
		srcString := string(srcData)
		for k, v := range vars {
			k = "${" + k + "}"
			srcString = strings.Replace(srcString, k, v, -1)
		}
		err = ioutil.WriteFile(destFile, []byte(srcString), 0644)
	}
	return err
}
func StartBuild(oc *CLI, args ...string) (stdout, stderr string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stdout, stderr, err = oc.Run("start-build").Args(args...).Outputs()
	e2e.Logf("\n\nstart-build output with args %v:\nError>%v\nStdOut>\n%s\nStdErr>\n%s\n\n", args, err, stdout, stderr)
	return stdout, stderr, err
}

var buildPathPattern = regexp.MustCompile(`^build\.build\.openshift\.io/([\w\-\._]+)$`)

type LogDumperFunc func(oc *CLI, br *BuildResult) (string, error)

func NewBuildResult(oc *CLI, build *buildv1.Build) *BuildResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &BuildResult{Oc: oc, BuildName: build.Name, BuildPath: "builds/" + build.Name}
}

type BuildResult struct {
	BuildPath		string
	BuildName		string
	StartBuildStdErr	string
	StartBuildStdOut	string
	StartBuildErr		error
	BuildConfigName		string
	Build			*buildv1.Build
	BuildAttempt		bool
	BuildSuccess		bool
	BuildFailure		bool
	BuildCancelled		bool
	BuildTimeout		bool
	LogDumper		LogDumperFunc
	Oc			*CLI
}

func (t *BuildResult) DumpLogs() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("\n\n*****************************************\n")
	e2e.Logf("Dumping Build Result: %#v\n", *t)
	if t == nil {
		e2e.Logf("No build result available!\n\n")
		return
	}
	desc, err := t.Oc.Run("describe").Args(t.BuildPath).Output()
	e2e.Logf("\n** Build Description:\n")
	if err != nil {
		e2e.Logf("Error during description retrieval: %+v\n", err)
	} else {
		e2e.Logf("%s\n", desc)
	}
	e2e.Logf("\n** Build Logs:\n")
	buildOuput, err := t.Logs()
	if err != nil {
		e2e.Logf("Error during log retrieval: %+v\n", err)
	} else {
		e2e.Logf("%s\n", buildOuput)
	}
	e2e.Logf("\n\n")
	t.dumpRegistryLogs()
}
func (t *BuildResult) dumpRegistryLogs() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buildStarted *time.Time
	oc := t.Oc
	e2e.Logf("\n** Registry Logs:\n")
	if t.Build != nil && !t.Build.CreationTimestamp.IsZero() {
		buildStarted = &t.Build.CreationTimestamp.Time
	} else {
		proj, err := oc.ProjectClient().Project().Projects().Get(oc.Namespace(), metav1.GetOptions{})
		if err != nil {
			e2e.Logf("Failed to get project %s: %v\n", oc.Namespace(), err)
		} else {
			buildStarted = &proj.CreationTimestamp.Time
		}
	}
	if buildStarted == nil {
		e2e.Logf("Could not determine test' start time\n\n\n")
		return
	}
	since := time.Now().Sub(*buildStarted)
	savedNamespace := t.Oc.Namespace()
	oadm := t.Oc.AsAdmin().SetNamespace("default")
	out, err := oadm.Run("logs").Args("dc/docker-registry", "--since="+since.String()).Output()
	if err != nil {
		e2e.Logf("Error during log retrieval: %+v\n", err)
	} else {
		e2e.Logf("%s\n", out)
	}
	oadm = t.Oc.AsAdmin().SetNamespace("openshift-image-registry")
	out, err = oadm.Run("logs").Args("deployment/image-registry", "--since="+since.String()).Output()
	if err != nil {
		e2e.Logf("Error during log retrieval: %+v\n", err)
	} else {
		e2e.Logf("%s\n", out)
	}
	t.Oc.SetNamespace(savedNamespace)
	e2e.Logf("\n\n")
}
func (t *BuildResult) Logs() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t == nil || t.BuildPath == "" {
		return "", fmt.Errorf("Not enough information to retrieve logs for %#v", *t)
	}
	if t.LogDumper != nil {
		return t.LogDumper(t.Oc, t)
	}
	buildOuput, err := t.Oc.Run("logs").Args("-f", t.BuildPath, "--timestamps").Output()
	if err != nil {
		return "", fmt.Errorf("Error retrieving logs for %#v: %v", *t, err)
	}
	return buildOuput, nil
}
func (t *BuildResult) LogsNoTimestamp() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t == nil || t.BuildPath == "" {
		return "", fmt.Errorf("Not enough information to retrieve logs for %#v", *t)
	}
	if t.LogDumper != nil {
		return t.LogDumper(t.Oc, t)
	}
	buildOuput, err := t.Oc.Run("logs").Args("-f", t.BuildPath).Output()
	if err != nil {
		return "", fmt.Errorf("Error retrieving logs for %#v: %v", *t, err)
	}
	return buildOuput, nil
}
func (t *BuildResult) AssertSuccess() *BuildResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !t.BuildSuccess {
		t.DumpLogs()
	}
	o.ExpectWithOffset(1, t.BuildSuccess).To(o.BeTrue())
	return t
}
func (t *BuildResult) AssertFailure() *BuildResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !t.BuildFailure {
		t.DumpLogs()
	}
	o.ExpectWithOffset(1, t.BuildFailure).To(o.BeTrue())
	return t
}
func StartBuildResult(oc *CLI, args ...string) (result *BuildResult, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args = append(args, "-o=name")
	stdout, stderr, err := StartBuild(oc, args...)
	buildPath := strings.TrimSpace(strings.Split(stdout, "\n")[0])
	result = &BuildResult{Build: nil, BuildPath: buildPath, StartBuildStdOut: stdout, StartBuildStdErr: stderr, StartBuildErr: nil, BuildAttempt: false, BuildSuccess: false, BuildFailure: false, BuildCancelled: false, BuildTimeout: false, Oc: oc}
	result.StartBuildErr = err
	matches := buildPathPattern.FindStringSubmatch(buildPath)
	if len(matches) != 2 {
		return result, fmt.Errorf("Build path output did not match expected format 'build/name' : %q", buildPath)
	}
	result.BuildName = matches[1]
	return result, nil
}
func StartBuildAndWait(oc *CLI, args ...string) (result *BuildResult, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result, err = StartBuildResult(oc, args...)
	if err != nil {
		return result, err
	}
	return result, WaitForBuildResult(oc.BuildClient().BuildV1().Builds(oc.Namespace()), result)
}
func WaitForBuildResult(c buildv1clienttyped.BuildInterface, result *BuildResult) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Waiting for %s to complete\n", result.BuildName)
	err := WaitForABuild(c, result.BuildName, func(b *buildv1.Build) bool {
		result.Build = b
		result.BuildSuccess = CheckBuildSuccess(b)
		return result.BuildSuccess
	}, func(b *buildv1.Build) bool {
		result.Build = b
		result.BuildFailure = CheckBuildFailed(b)
		return result.BuildFailure
	}, func(b *buildv1.Build) bool {
		result.Build = b
		result.BuildCancelled = CheckBuildCancelled(b)
		return result.BuildCancelled
	})
	if result.Build == nil {
		return fmt.Errorf("Severe error waiting for build: %v", err)
	}
	result.BuildAttempt = true
	result.BuildTimeout = !(result.BuildFailure || result.BuildSuccess || result.BuildCancelled)
	e2e.Logf("Done waiting for %s: %#v\n with error: %v\n", result.BuildName, *result, err)
	return nil
}
func WaitForABuild(c buildv1clienttyped.BuildInterface, name string, isOK, isFailed, isCanceled func(*buildv1.Build) bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isOK == nil {
		isOK = CheckBuildSuccess
	}
	if isFailed == nil {
		isFailed = CheckBuildFailed
	}
	if isCanceled == nil {
		isCanceled = CheckBuildCancelled
	}
	err := wait.Poll(1*time.Second, 2*time.Minute, func() (bool, error) {
		if _, err := c.Get(name, metav1.GetOptions{}); err != nil {
			return false, nil
		}
		return true, nil
	})
	if err == wait.ErrWaitTimeout {
		return fmt.Errorf("Timed out waiting for build %q to be created", name)
	}
	if err != nil {
		return err
	}
	err = wait.Poll(5*time.Second, 10*time.Minute, func() (bool, error) {
		list, err := c.List(metav1.ListOptions{FieldSelector: fields.Set{"metadata.name": name}.AsSelector().String()})
		if err != nil {
			e2e.Logf("error listing builds: %v", err)
			return false, err
		}
		for i := range list.Items {
			if name == list.Items[i].Name && (isOK(&list.Items[i]) || isCanceled(&list.Items[i])) {
				return true, nil
			}
			if name != list.Items[i].Name {
				return false, fmt.Errorf("While listing builds named %s, found unexpected build %#v", name, list.Items[i])
			}
			if isFailed(&list.Items[i]) {
				return false, fmt.Errorf("The build %q status is %q", name, list.Items[i].Status.Phase)
			}
		}
		return false, nil
	})
	if err != nil {
		e2e.Logf("WaitForABuild returning with error: %v", err)
	}
	if err == wait.ErrWaitTimeout {
		return fmt.Errorf("Timed out waiting for build %q to complete", name)
	}
	return err
}
func CheckBuildSuccess(b *buildv1.Build) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.Status.Phase == buildv1.BuildPhaseComplete
}
func CheckBuildFailed(b *buildv1.Build) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.Status.Phase == buildv1.BuildPhaseFailed || b.Status.Phase == buildv1.BuildPhaseError
}
func CheckBuildCancelled(b *buildv1.Build) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.Status.Phase == buildv1.BuildPhaseCancelled
}
func WaitForServiceAccount(c corev1client.ServiceAccountInterface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waitFn := func() (bool, error) {
		sc, err := c.Get(name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) || errors.IsForbidden(err) {
				return false, nil
			}
			return false, err
		}
		for _, s := range sc.Secrets {
			if strings.Contains(s.Name, "dockercfg") {
				return true, nil
			}
		}
		return false, nil
	}
	return wait.Poll(time.Duration(100*time.Millisecond), 3*time.Minute, waitFn)
}
func WaitForAnImageStream(client imagetypeclientset.ImageStreamInterface, name string, isOK, isFailed func(*imageapi.ImageStream) bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		list, err := client.List(metav1.ListOptions{FieldSelector: fields.Set{"metadata.name": name}.AsSelector().String()})
		if err != nil {
			return err
		}
		for i := range list.Items {
			if isOK(&list.Items[i]) {
				return nil
			}
			if isFailed(&list.Items[i]) {
				return fmt.Errorf("The image stream %q status is %q", name, list.Items[i].Annotations[imageapi.DockerImageRepositoryCheckAnnotation])
			}
		}
		rv := list.ResourceVersion
		w, err := client.Watch(metav1.ListOptions{FieldSelector: fields.Set{"metadata.name": name}.AsSelector().String(), ResourceVersion: rv})
		if err != nil {
			return err
		}
		defer w.Stop()
		for {
			val, ok := <-w.ResultChan()
			if !ok {
				break
			}
			if e, ok := val.Object.(*imageapi.ImageStream); ok {
				if isOK(e) {
					return nil
				}
				if isFailed(e) {
					return fmt.Errorf("The image stream %q status is %q", name, e.Annotations[imageapi.DockerImageRepositoryCheckAnnotation])
				}
			}
		}
	}
}
func WaitForAnImageStreamTag(oc *CLI, namespace, name, tag string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return TimedWaitForAnImageStreamTag(oc, namespace, name, tag, time.Second*300)
}
func TimedWaitForAnImageStreamTag(oc *CLI, namespace, name, tag string, waitTimeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.By(fmt.Sprintf("waiting for an is importer to import a tag %s into a stream %s", tag, name))
	start := time.Now()
	c := make(chan error)
	go func() {
		err := WaitForAnImageStream(oc.ImageClient().Image().ImageStreams(namespace), name, func(is *imageapi.ImageStream) bool {
			if history, exists := is.Status.Tags[tag]; !exists || len(history.Items) == 0 {
				return false
			}
			return true
		}, func(is *imageapi.ImageStream) bool {
			return time.Now().After(start.Add(waitTimeout))
		})
		c <- err
	}()
	select {
	case e := <-c:
		return e
	case <-time.After(waitTimeout):
		return fmt.Errorf("timed out while waiting of an image stream tag %s/%s:%s", namespace, name, tag)
	}
}
func CheckImageStreamLatestTagPopulated(i *imageapi.ImageStream) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := i.Status.Tags["latest"]
	return ok
}
func CheckImageStreamTagNotFound(i *imageapi.ImageStream) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(i.Annotations[imageapi.DockerImageRepositoryCheckAnnotation], "not") || strings.Contains(i.Annotations[imageapi.DockerImageRepositoryCheckAnnotation], "error")
}
func WaitForDeploymentConfig(kc kubernetes.Interface, dcClient appsv1clienttyped.DeploymentConfigsGetter, namespace, name string, version int64, enforceNotProgressing bool, cli *CLI) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("waiting for deploymentconfig %s/%s to be available with version %d\n", namespace, name, version)
	var dc *appsv1.DeploymentConfig
	start := time.Now()
	err := wait.Poll(time.Second, 15*time.Minute, func() (done bool, err error) {
		dc, err = dcClient.DeploymentConfigs(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if dc.Status.LatestVersion < version {
			return false, nil
		}
		var progressing, available *appsv1.DeploymentCondition
		for i, condition := range dc.Status.Conditions {
			switch condition.Type {
			case appsv1.DeploymentProgressing:
				progressing = &dc.Status.Conditions[i]
			case appsv1.DeploymentAvailable:
				available = &dc.Status.Conditions[i]
			}
		}
		if enforceNotProgressing {
			if progressing != nil && progressing.Status == corev1.ConditionFalse {
				return false, fmt.Errorf("not progressing")
			}
		}
		if progressing != nil && progressing.Status == corev1.ConditionTrue && progressing.Reason == appsutil.NewRcAvailableReason && available != nil && available.Status == corev1.ConditionTrue {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		e2e.Logf("got error %q when waiting for deploymentconfig %s/%s to be available with version %d\n", err, namespace, name, version)
		cli.Run("get").Args("dc", dc.Name, "-o", "yaml").Execute()
		DumpDeploymentLogs(name, version, cli)
		DumpApplicationPodLogs(name, cli)
		return err
	}
	requirement, err := labels.NewRequirement(appsutil.DeploymentLabel, selection.Equals, []string{appsutil.LatestDeploymentNameForConfigAndVersion(dc.Name, dc.Status.LatestVersion)})
	if err != nil {
		return err
	}
	podnames, err := GetPodNamesByFilter(kc.CoreV1().Pods(namespace), labels.NewSelector().Add(*requirement), func(kapiv1.Pod) bool {
		return true
	})
	if err != nil {
		return err
	}
	e2e.Logf("deploymentconfig %s/%s available after %s\npods: %s\n", namespace, name, time.Now().Sub(start), strings.Join(podnames, ", "))
	return nil
}
func isUsageSynced(received, expected corev1.ResourceList, expectedIsUpperLimit bool) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceNames := quota.ResourceNames(expected)
	masked := quota.Mask(received, resourceNames)
	if len(masked) != len(expected) {
		return false
	}
	if expectedIsUpperLimit {
		if le, _ := quota.LessThanOrEqual(masked, expected); !le {
			return false
		}
	} else {
		if le, _ := quota.LessThanOrEqual(expected, masked); !le {
			return false
		}
	}
	return true
}
func WaitForResourceQuotaSync(client corev1client.ResourceQuotaInterface, name string, expectedUsage corev1.ResourceList, expectedIsUpperLimit bool, timeout time.Duration) (corev1.ResourceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	endTime := startTime.Add(timeout)
	expectedResourceNames := quota.ResourceNames(expectedUsage)
	list, err := client.List(metav1.ListOptions{FieldSelector: fields.Set{"metadata.name": name}.AsSelector().String()})
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		used := quota.Mask(list.Items[i].Status.Used, expectedResourceNames)
		if isUsageSynced(used, expectedUsage, expectedIsUpperLimit) {
			return used, nil
		}
	}
	rv := list.ResourceVersion
	w, err := client.Watch(metav1.ListOptions{FieldSelector: fields.Set{"metadata.name": name}.AsSelector().String(), ResourceVersion: rv})
	if err != nil {
		return nil, err
	}
	defer w.Stop()
	for time.Now().Before(endTime) {
		select {
		case val, ok := <-w.ResultChan():
			if !ok {
				continue
			}
			if rq, ok := val.Object.(*corev1.ResourceQuota); ok {
				used := quota.Mask(rq.Status.Used, expectedResourceNames)
				if isUsageSynced(used, expectedUsage, expectedIsUpperLimit) {
					return used, nil
				}
			}
		case <-time.After(endTime.Sub(time.Now())):
			return nil, wait.ErrWaitTimeout
		}
	}
	return nil, wait.ErrWaitTimeout
}
func GetPodNamesByFilter(c corev1client.PodInterface, label labels.Selector, predicate func(kapiv1.Pod) bool) (podNames []string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	podList, err := c.List(metav1.ListOptions{LabelSelector: label.String()})
	if err != nil {
		return nil, err
	}
	for _, pod := range podList.Items {
		if predicate(pod) {
			podNames = append(podNames, pod.Name)
		}
	}
	return podNames, nil
}
func WaitForAJob(c batchv1client.JobInterface, name string, timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(1*time.Second, timeout, func() (bool, error) {
		j, e := c.Get(name, metav1.GetOptions{})
		if e != nil {
			return true, e
		}
		for _, c := range j.Status.Conditions {
			if (c.Type == batchv1.JobComplete || c.Type == batchv1.JobFailed) && c.Status == kapiv1.ConditionTrue {
				return true, nil
			}
		}
		return false, nil
	})
}
func WaitForPods(c corev1client.PodInterface, label labels.Selector, predicate func(kapiv1.Pod) bool, count int, timeout time.Duration) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var podNames []string
	err := wait.Poll(1*time.Second, timeout, func() (bool, error) {
		p, e := GetPodNamesByFilter(c, label, predicate)
		if e != nil {
			return true, e
		}
		if len(p) != count {
			return false, nil
		}
		podNames = p
		return true, nil
	})
	return podNames, err
}
func CheckPodIsRunning(pod kapiv1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pod.Status.Phase == kapiv1.PodRunning
}
func CheckPodIsSucceeded(pod kapiv1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pod.Status.Phase == kapiv1.PodSucceeded
}
func CheckPodIsReady(pod kapiv1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod.Status.Phase != kapiv1.PodRunning {
		return false
	}
	for _, cond := range pod.Status.Conditions {
		if cond.Type != kapiv1.PodReady {
			continue
		}
		return cond.Status == kapiv1.ConditionTrue
	}
	return false
}
func CheckPodNoOp(pod kapiv1.Pod) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func WaitUntilPodIsGone(c corev1client.PodInterface, podName string, timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(1*time.Second, timeout, func() (bool, error) {
		_, err := c.Get(podName, metav1.GetOptions{})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return true, nil
			}
			return true, err
		}
		return false, nil
	})
}
func GetDockerImageReference(c imagetypeclientset.ImageStreamInterface, name, tag string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	imageStream, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	isTag, ok := imageStream.Status.Tags[tag]
	if !ok {
		return "", fmt.Errorf("ImageStream %q does not have tag %q", name, tag)
	}
	if len(isTag.Items) == 0 {
		return "", fmt.Errorf("ImageStreamTag %q is empty", tag)
	}
	return isTag.Items[0].DockerImageReference, nil
}
func GetPodForContainer(container kapiv1.Container) *kapiv1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name := apihelpers.GetPodName("test-pod", string(uuid.NewUUID()))
	return &kapiv1.Pod{TypeMeta: metav1.TypeMeta{Kind: "Pod", APIVersion: "v1"}, ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{"name": name}}, Spec: kapiv1.PodSpec{Containers: []kapiv1.Container{container}, RestartPolicy: kapiv1.RestartPolicyNever}}
}
func KubeConfigPath() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return os.Getenv("KUBECONFIG")
}
func ArtifactDirPath() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := os.Getenv("ARTIFACT_DIR")
	o.Expect(path).NotTo(o.BeNil())
	o.Expect(path).NotTo(o.BeEmpty())
	return path
}
func ArtifactPath(elem ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return filepath.Join(append([]string{ArtifactDirPath()}, elem...)...)
}

var (
	fixtureDirLock	sync.Once
	fixtureDir	string
)

func FixturePath(elem ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case len(elem) == 0:
		panic("must specify path")
	case len(elem) > 3 && elem[0] == ".." && elem[1] == ".." && elem[2] == "examples":
		elem = elem[2:]
	case len(elem) > 3 && elem[0] == ".." && elem[1] == ".." && elem[2] == "install":
		elem = elem[2:]
	case len(elem) > 3 && elem[0] == ".." && elem[1] == "integration":
		elem = append([]string{"test"}, elem[1:]...)
	case elem[0] == "testdata":
		elem = append([]string{"test", "extended"}, elem...)
	default:
		panic(fmt.Sprintf("Fixtures must be in test/extended/testdata or examples not %s", path.Join(elem...)))
	}
	fixtureDirLock.Do(func() {
		dir, err := ioutil.TempDir("", "fixture-testdata-dir")
		if err != nil {
			panic(err)
		}
		fixtureDir = dir
	})
	relativePath := path.Join(elem...)
	fullPath := path.Join(fixtureDir, relativePath)
	if err := testdata.RestoreAsset(fixtureDir, relativePath); err != nil {
		if err := testdata.RestoreAssets(fixtureDir, relativePath); err != nil {
			panic(err)
		}
		if err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
			if err := os.Chmod(path, 0640); err != nil {
				return err
			}
			if stat, err := os.Lstat(path); err == nil && stat.IsDir() {
				return os.Chmod(path, 0755)
			}
			return nil
		}); err != nil {
			panic(err)
		}
	} else {
		if err := os.Chmod(fullPath, 0640); err != nil {
			panic(err)
		}
	}
	p, err := filepath.Abs(fullPath)
	if err != nil {
		panic(err)
	}
	return p
}
func FetchURL(oc *CLI, url string, retryTimeout time.Duration) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := oc.KubeFramework().Namespace.Name
	execPodName := CreateExecPodOrFail(oc.AdminKubeClient().CoreV1(), ns, string(uuid.NewUUID()))
	defer func() {
		oc.AdminKubeClient().CoreV1().Pods(ns).Delete(execPodName, metav1.NewDeleteOptions(1))
	}()
	execPod, err := oc.AdminKubeClient().CoreV1().Pods(ns).Get(execPodName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	var response string
	waitFn := func() (bool, error) {
		e2e.Logf("Waiting up to %v to wget %s", retryTimeout, url)
		cmd := fmt.Sprintf("curl -vvv %s", url)
		response, err = e2e.RunHostCmd(execPod.Namespace, execPod.Name, cmd)
		if err != nil {
			e2e.Logf("got err: %v, retry until timeout", err)
			return false, nil
		}
		if strings.TrimSpace(response) == "" {
			e2e.Logf("got empty stdout, retry until timeout")
			return false, nil
		}
		return true, nil
	}
	pollErr := wait.Poll(time.Duration(1*time.Second), retryTimeout, waitFn)
	if pollErr == wait.ErrWaitTimeout {
		return "", fmt.Errorf("Timed out while fetching url %q", url)
	}
	if pollErr != nil {
		return "", pollErr
	}
	return response, nil
}
func ParseLabelsOrDie(str string) labels.Selector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret, err := labels.Parse(str)
	if err != nil {
		panic(fmt.Sprintf("cannot parse '%v': %v", str, err))
	}
	return ret
}
func GetEndpointAddress(oc *CLI, name string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := e2e.WaitForEndpoint(oc.KubeFramework().ClientSet, oc.Namespace(), name)
	if err != nil {
		return "", err
	}
	endpoint, err := oc.KubeClient().CoreV1().Endpoints(oc.Namespace()).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", endpoint.Subsets[0].Addresses[0].IP, endpoint.Subsets[0].Ports[0].Port), nil
}
func CreateExecPodOrFail(client corev1client.CoreV1Interface, ns, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e2e.Logf("Creating new exec pod")
	execPod := e2e.NewHostExecPodSpec(ns, name)
	created, err := client.Pods(ns).Create(execPod)
	o.Expect(err).NotTo(o.HaveOccurred())
	err = wait.PollImmediate(e2e.Poll, 5*time.Minute, func() (bool, error) {
		retrievedPod, err := client.Pods(execPod.Namespace).Get(created.Name, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}
		return retrievedPod.Status.Phase == kapiv1.PodRunning, nil
	})
	o.Expect(err).NotTo(o.HaveOccurred())
	return created.Name
}
func CheckForBuildEvent(client corev1client.CoreV1Interface, build *buildv1.Build, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var expectedEvent *kapiv1.Event
	err := wait.PollImmediate(e2e.Poll, 1*time.Minute, func() (bool, error) {
		events, err := client.Events(build.Namespace).Search(legacyscheme.Scheme, build)
		if err != nil {
			return false, err
		}
		for _, event := range events.Items {
			e2e.Logf("Found event %#v", event)
			if reason == event.Reason {
				expectedEvent = &event
				return true, nil
			}
		}
		return false, nil
	})
	o.ExpectWithOffset(1, err).NotTo(o.HaveOccurred(), "Should be able to get events from the build")
	o.ExpectWithOffset(1, expectedEvent).NotTo(o.BeNil(), "Did not find a %q event on build %s/%s", reason, build.Namespace, build.Name)
	o.ExpectWithOffset(1, expectedEvent.Message).To(o.Equal(fmt.Sprintf(message, build.Namespace, build.Name)))
}

type podExecutor struct {
	client	*CLI
	podName	string
}

func NewPodExecutor(oc *CLI, name, image string) (*podExecutor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := oc.Run("run").Args(name, "--labels", "name="+name, "--image", image, "--restart", "Never", "--command", "--", "/bin/bash", "-c", "sleep infinity").Output()
	if err != nil {
		return nil, fmt.Errorf("error: %v\n(%s)", err, out)
	}
	_, err = WaitForPods(oc.KubeClient().CoreV1().Pods(oc.Namespace()), ParseLabelsOrDie("name="+name), CheckPodIsReady, 1, 3*time.Minute)
	if err != nil {
		return nil, err
	}
	return &podExecutor{client: oc, podName: name}, nil
}
func (r *podExecutor) Exec(script string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var out string
	waitErr := wait.PollImmediate(1*time.Second, 3*time.Minute, func() (bool, error) {
		var err error
		out, err = r.client.Run("exec").Args(r.podName, "--", "/bin/bash", "-c", script).Output()
		return true, err
	})
	return out, waitErr
}
func (r *podExecutor) CopyFromHost(local, remote string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := r.client.Run("cp").Args(local, fmt.Sprintf("%s:%s", r.podName, remote)).Output()
	return err
}

type GitRepo struct {
	baseTempDir	string
	upstream	git.Repository
	upstreamPath	string
	repo		git.Repository
	RepoPath	string
}

func (r GitRepo) AddAndCommit(file, content string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir := filepath.Dir(file)
	if err := os.MkdirAll(filepath.Join(r.RepoPath, dir), 0777); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(r.RepoPath, file), []byte(content), 0666); err != nil {
		return err
	}
	if err := r.repo.Add(r.RepoPath, file); err != nil {
		return err
	}
	if err := r.repo.Commit(r.RepoPath, "added file "+file); err != nil {
		return err
	}
	return nil
}
func (r GitRepo) Remove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.baseTempDir != "" {
		os.RemoveAll(r.baseTempDir)
	}
}
func NewGitRepo(repoName string) (GitRepo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDir, err := ioutil.TempDir(os.TempDir(), repoName)
	if err != nil {
		return GitRepo{}, err
	}
	repoPath := filepath.Join(testDir, repoName)
	upstreamPath := repoPath + `.git`
	upstream := git.NewRepository()
	if err = upstream.Init(upstreamPath, true); err != nil {
		return GitRepo{baseTempDir: testDir}, err
	}
	repo := git.NewRepository()
	if err = repo.Clone(repoPath, upstreamPath); err != nil {
		return GitRepo{baseTempDir: testDir}, err
	}
	return GitRepo{testDir, upstream, upstreamPath, repo, repoPath}, nil
}
func WaitForUserBeAuthorized(oc *CLI, user, verb, resource string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sar := &authorization.SubjectAccessReview{Spec: authorization.SubjectAccessReviewSpec{ResourceAttributes: &authorization.ResourceAttributes{Namespace: oc.Namespace(), Verb: verb, Resource: resource}, User: user}}
	return wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		resp, err := oc.InternalAdminKubeClient().Authorization().SubjectAccessReviews().Create(sar)
		if err == nil && resp != nil && resp.Status.Allowed {
			return true, nil
		}
		return false, err
	})
}
func GetRouterPodTemplate(oc *CLI) (*corev1.PodTemplateSpec, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	appsclient := oc.AdminAppsClient().AppsV1()
	k8sappsclient := oc.AdminKubeClient().AppsV1()
	for _, ns := range []string{"default", "openshift-ingress", "tectonic-ingress"} {
		dc, err := appsclient.DeploymentConfigs(ns).Get("router", metav1.GetOptions{})
		if err == nil {
			return dc.Spec.Template, ns, nil
		}
		if !errors.IsNotFound(err) {
			return nil, "", err
		}
		deploy, err := k8sappsclient.Deployments(ns).Get("router", metav1.GetOptions{})
		if err == nil {
			return &deploy.Spec.Template, ns, nil
		}
		if !errors.IsNotFound(err) {
			return nil, "", err
		}
		deploy, err = k8sappsclient.Deployments(ns).Get("router-default", metav1.GetOptions{})
		if err == nil {
			return &deploy.Spec.Template, ns, nil
		}
		if !errors.IsNotFound(err) {
			return nil, "", err
		}
	}
	return nil, "", errors.NewNotFound(schema.GroupResource{Group: "apps.openshift.io", Resource: "deploymentconfigs"}, "router")
}
func FindImageFormatString(oc *CLI) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	template, _, err := GetRouterPodTemplate(oc)
	if err == nil {
		if strings.Contains(template.Spec.Containers[0].Image, "haproxy-router") {
			return strings.Replace(template.Spec.Containers[0].Image, "haproxy-router", "${component}", -1), true
		}
	}
	return "openshift/origin-${component}:latest", false
}
func FindCLIImage(oc *CLI) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	is, err := oc.AdminImageClient().ImageV1().ImageStreams("openshift").Get("cli", metav1.GetOptions{})
	if err == nil {
		for _, tag := range is.Spec.Tags {
			if tag.Name == "latest" && tag.From != nil && tag.From.Kind == "DockerImage" {
				return tag.From.Name, true
			}
		}
	}
	format, ok := FindImageFormatString(oc)
	return strings.Replace(format, "${component}", "cli", -1), ok
}
func FindRouterImage(oc *CLI) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	format, ok := FindImageFormatString(oc)
	return strings.Replace(format, "${component}", "haproxy-router", -1), ok
}
func IsClusterOperated(oc *CLI) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configclient := oc.AdminConfigClient().ConfigV1()
	o, err := configclient.Images().Get("cluster", metav1.GetOptions{})
	if o == nil || err != nil {
		e2e.Logf("Could not find image config object, assuming non-4.0 installed cluster: %v", err)
		return false
	}
	return true
}
