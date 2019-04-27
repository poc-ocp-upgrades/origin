package fsnotification

import (
	"fmt"
	"os"
	"path/filepath"
	"k8s.io/klog"
	"github.com/fsnotify/fsnotify"
)

func AddRecursiveWatch(watcher *fsnotify.Watcher, path string) error {
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
	file, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error introspecting path %s: %v", path, err)
	}
	if !file.IsDir() {
		return nil
	}
	folders, err := getSubFolders(path)
	for _, v := range folders {
		klog.V(5).Infof("adding watch on path %s", v)
		err = watcher.Add(v)
		if err != nil {
			return fmt.Errorf("error adding watcher for path %s: %v", v, err)
		}
	}
	return nil
}
func getSubFolders(path string) (paths []string, err error) {
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
	err = filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths, err
}
