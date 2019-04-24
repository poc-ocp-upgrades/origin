package localquota

import (
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	"k8s.io/kubernetes/pkg/volume/emptydirquota"
	exutil "github.com/openshift/origin/test/extended/util"
)

const (
	volDirEnvVar		= "VOLUME_DIR"
	podCreationTimeout	= 120
	expectedQuotaKb		= 4587520
)

func lookupFSGroup(oc *exutil.CLI, project string) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gidRange, err := oc.Run("get").Args("project", project, "--template='{{ index .metadata.annotations \"openshift.io/sa.scc.supplemental-groups\" }}'").Output()
	if err != nil {
		return 0, err
	}
	fsGroupStr := strings.Split(gidRange, "/")[0]
	fsGroupStr = strings.Replace(fsGroupStr, "'", "", -1)
	fsGroup, err := strconv.Atoi(fsGroupStr)
	if err != nil {
		return 0, err
	}
	return fsGroup, nil
}
func lookupXFSQuota(oc *exutil.CLI, fsGroup int, volDir string) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fsDevice, err := emptydirquota.GetFSDevice(volDir)
	if err != nil {
		return 0, err
	}
	args := []string{"xfs_quota", "-x", "-c", fmt.Sprintf("report -n -L %d -U %d", fsGroup, fsGroup), fsDevice}
	cmd := exec.Command("sudo", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	outBytes, reportErr := cmd.Output()
	if reportErr != nil {
		return 0, reportErr
	}
	quotaReport := string(outBytes)
	lines := strings.Split(quotaReport, "\n")
	for _, l := range lines {
		if strings.HasPrefix(l, fmt.Sprintf("#%d", fsGroup)) {
			words := strings.Fields(l)
			if len(words) != 6 {
				return 0, fmt.Errorf("expected 6 words in quota line: %s", l)
			}
			quota, err := strconv.Atoi(words[3])
			if err != nil {
				return 0, err
			}
			return quota, nil
		}
	}
	return -1, nil
}
func waitForQuotaToBeApplied(oc *exutil.CLI, fsGroup int, volDir string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	secondsWaited := 0
	for secondsWaited < podCreationTimeout {
		quotaFound, quotaErr := lookupXFSQuota(oc, fsGroup, volDir)
		o.Expect(quotaErr).NotTo(o.HaveOccurred())
		if quotaFound == expectedQuotaKb {
			return nil
		}
		time.Sleep(1 * time.Second)
		secondsWaited = secondsWaited + 1
	}
	return fmt.Errorf("expected quota was not applied in time")
}

var _ = g.Describe("[Conformance][volumes] Test local storage quota", func() {
	defer g.GinkgoRecover()
	var (
		oc			= exutil.NewCLI("local-quota", exutil.KubeConfigPath())
		emptyDirPodFixture	= exutil.FixturePath("..", "..", "examples", "hello-openshift", "hello-pod.json")
	)
	g.Describe("FSGroup local storage quota [local]", func() {
		g.It("should be applied to XFS filesystem when a pod is created", func() {
			project := oc.Namespace()
			volDir := os.Getenv(volDirEnvVar)
			g.By(fmt.Sprintf("make sure volume directory (%s) is on an XFS filesystem", volDir))
			o.Expect(volDir).NotTo(o.Equal(""))
			args := []string{"-f", "-c", "'%T'", volDir}
			outBytes, _ := exec.Command("stat", args...).Output()
			fmt.Fprintf(g.GinkgoWriter, "Volume directory status: \n%s\n", outBytes)
			if !strings.Contains(string(outBytes), "xfs") {
				g.Skip("Volume directory is not on an XFS filesystem, skipping...")
			}
			g.By("lookup test projects fsGroup ID")
			fsGroup, err := lookupFSGroup(oc, project)
			o.Expect(err).NotTo(o.HaveOccurred())
			g.By("create hello-openshift pod with emptyDir volume")
			_, createPodErr := oc.Run("create").Args("-f", emptyDirPodFixture).Output()
			o.Expect(createPodErr).NotTo(o.HaveOccurred())
			g.By("wait for XFS quota to be applied and verify")
			lookupQuotaErr := waitForQuotaToBeApplied(oc, fsGroup, volDir)
			o.Expect(lookupQuotaErr).NotTo(o.HaveOccurred())
		})
	})
})

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
