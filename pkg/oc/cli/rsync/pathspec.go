package rsync

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kvalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

type PathSpec struct {
	PodName	string
	Path	string
}

func (s *PathSpec) Local() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s.PodName) == 0
}
func (s *PathSpec) RsyncPath() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(s.PodName) > 0 {
		return fmt.Sprintf("%s:%s", s.PodName, s.Path)
	}
	if isWindows() {
		return convertWindowsPath(s.Path)
	}
	return s.Path
}
func (s *PathSpec) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.Local() {
		info, err := os.Stat(s.Path)
		if err != nil {
			return fmt.Errorf("invalid path %s: %v", s.Path, err)
		}
		if !info.IsDir() {
			return fmt.Errorf("path %s must point to a directory", s.Path)
		}
	}
	return nil
}
func isPathForPod(path string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(path, ":", 2)
	if len(parts) == 1 || (isWindows() && len(parts[0]) == 1) {
		return false
	}
	return true
}
func parsePathSpec(path string) (*PathSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(path, ":", 2)
	if !isPathForPod(path) {
		return &PathSpec{Path: path}, nil
	}
	if reasons := kvalidation.ValidatePodName(parts[0], false); len(reasons) != 0 {
		return nil, fmt.Errorf("invalid pod name %s: %s", parts[0], strings.Join(reasons, ", "))
	}
	return &PathSpec{PodName: parts[0], Path: parts[1]}, nil
}
func resolveResourceKindPath(f kcmdutil.Factory, path, namespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(path, ":", 2)
	if !isPathForPod(path) {
		return path, nil
	}
	podName := parts[0]
	if podSegs := strings.Split(podName, "/"); len(podSegs) > 1 {
		podName = podSegs[1]
	}
	r := f.NewBuilder().WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).NamespaceParam(namespace).SingleResourceType().ResourceNames("pods", podName).Do()
	if err := r.Err(); err != nil {
		return "", err
	}
	infos, err := r.Infos()
	if err != nil {
		return "", err
	}
	if len(infos) == 0 || infos[0].Mapping.Resource.GroupResource() != (schema.GroupResource{Resource: "pods"}) {
		return "", fmt.Errorf("error: expected resource to be of type pod, got %q", infos[0].Mapping.Resource)
	}
	return fmt.Sprintf("%s:%s", podName, parts[1]), nil
}
func convertWindowsPath(path string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(path, ":", 2)
	if len(parts) > 1 && len(parts[0]) == 1 {
		return fmt.Sprintf("/cygdrive/%s/%s", strings.ToLower(parts[0]), strings.TrimPrefix(filepath.ToSlash(parts[1]), "/"))
	}
	return filepath.ToSlash(path)
}
