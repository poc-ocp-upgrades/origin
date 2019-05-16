package fsnotification

import (
	"fmt"
	goformat "fmt"
	"github.com/fsnotify/fsnotify"
	"k8s.io/klog"
	"os"
	goos "os"
	"path/filepath"
	godefaultruntime "runtime"
	gotime "time"
)

func AddRecursiveWatch(watcher *fsnotify.Watcher, path string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
