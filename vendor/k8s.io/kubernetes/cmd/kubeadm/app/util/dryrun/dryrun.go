package dryrun

import (
	"fmt"
	goformat "fmt"
	"io"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type FileToPrint struct {
	RealPath  string
	PrintPath string
}

func NewFileToPrint(realPath, printPath string) FileToPrint {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return FileToPrint{RealPath: realPath, PrintPath: printPath}
}
func PrintDryRunFile(fileName, realDir, printDir string, w io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return PrintDryRunFiles([]FileToPrint{NewFileToPrint(filepath.Join(realDir, fileName), filepath.Join(printDir, fileName))}, w)
}
func PrintDryRunFiles(files []FileToPrint, w io.Writer) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	errs := []error{}
	for _, file := range files {
		if len(file.RealPath) == 0 {
			continue
		}
		fileBytes, err := ioutil.ReadFile(file.RealPath)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		outputFilePath := file.PrintPath
		if len(outputFilePath) == 0 {
			outputFilePath = file.RealPath
		}
		fmt.Fprintf(w, "[dryrun] Would write file %q with content:\n", outputFilePath)
		apiclient.PrintBytesWithLinePrefix(w, fileBytes, "\t")
	}
	return errors.NewAggregate(errs)
}

type Waiter struct{}

func NewWaiter() apiclient.Waiter {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &Waiter{}
}
func (w *Waiter) WaitForAPI() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Println("[dryrun] Would wait for the API Server's /healthz endpoint to return 'ok'")
	return nil
}
func (w *Waiter) WaitForPodsWithLabel(kvLabel string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[dryrun] Would wait for the Pods with the label %q in the %s namespace to become Running\n", kvLabel, metav1.NamespaceSystem)
	return nil
}
func (w *Waiter) WaitForPodToDisappear(podName string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[dryrun] Would wait for the %q Pod in the %s namespace to be deleted\n", podName, metav1.NamespaceSystem)
	return nil
}
func (w *Waiter) WaitForHealthyKubelet(_ time.Duration, healthzEndpoint string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fmt.Printf("[dryrun] Would make sure the kubelet %q endpoint is healthy\n", healthzEndpoint)
	return nil
}
func (w *Waiter) WaitForKubeletAndFunc(f func() error) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func (w *Waiter) SetTimeout(_ time.Duration) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
}
func (w *Waiter) WaitForStaticPodControlPlaneHashes(_ string) (map[string]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return map[string]string{constants.KubeAPIServer: "", constants.KubeControllerManager: "", constants.KubeScheduler: ""}, nil
}
func (w *Waiter) WaitForStaticPodSingleHash(_ string, _ string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "", nil
}
func (w *Waiter) WaitForStaticPodHashChange(_, _, _ string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
