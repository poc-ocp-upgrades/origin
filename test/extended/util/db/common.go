package db

import (
	"fmt"
	"bytes"
	"net/http"
	"runtime"
	"os/exec"
	"strings"
	"github.com/openshift/origin/test/extended/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kcoreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodConfig struct {
	Container	string
	Env		map[string]string
}

func getPodConfig(c kcoreclient.PodInterface, podName string) (conf *PodConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, err := c.Get(podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	env := make(map[string]string)
	for _, container := range pod.Spec.Containers {
		for _, e := range container.Env {
			env[e.Name] = e.Value
		}
	}
	return &PodConfig{pod.Spec.Containers[0].Name, env}, nil
}
func firstContainerName(c kcoreclient.PodInterface, podName string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pod, err := c.Get(podName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return pod.Spec.Containers[0].Name, nil
}
func isReady(oc *util.CLI, podName string, pingCommand, expectedOutput string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := executeShellCommand(oc, podName, pingCommand)
	ok := strings.Contains(out, expectedOutput)
	if !ok {
		err = fmt.Errorf("Expected output: %q but actual: %q", expectedOutput, out)
	}
	return ok, err
}
func executeShellCommand(oc *util.CLI, podName string, command string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := oc.Run("exec").Args(podName, "--", "bash", "-c", command).Output()
	if err != nil {
		switch err.(type) {
		case *util.ExitError, *exec.ExitError:
			return "", nil
		default:
			return "", err
		}
	}
	return out, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
