package main

import (
	"fmt"
	goformat "fmt"
	"io"
	"io/ioutil"
	"k8s.io/klog"
	"os"
	goos "os"
	"os/exec"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type DataDirectory struct {
	path        string
	versionFile *VersionFile
}

func OpenOrCreateDataDirectory(path string) (*DataDirectory, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	exists, err := exists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		klog.Infof("data directory '%s' does not exist, creating it", path)
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return nil, fmt.Errorf("failed to create data directory %s: %v", path, err)
		}
	}
	versionFile := &VersionFile{path: filepath.Join(path, versionFilename)}
	return &DataDirectory{path, versionFile}, nil
}
func (d *DataDirectory) Initialize(target *EtcdVersionPair) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	isEmpty, err := d.IsEmpty()
	if err != nil {
		return err
	}
	if isEmpty {
		klog.Infof("data directory '%s' is empty, writing target version '%s' to version.txt", d.path, target)
		err = d.versionFile.Write(target)
		if err != nil {
			return fmt.Errorf("failed to write version.txt to '%s': %v", d.path, err)
		}
		return nil
	}
	return nil
}
func (d *DataDirectory) Backup() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	backupDir := fmt.Sprintf("%s.bak", d.path)
	err := os.RemoveAll(backupDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(backupDir, 0777)
	if err != nil {
		return err
	}
	err = exec.Command("cp", "-r", d.path, backupDir).Run()
	if err != nil {
		return err
	}
	return nil
}
func (d *DataDirectory) IsEmpty() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dir, err := os.Open(d.path)
	if err != nil {
		return false, fmt.Errorf("failed to open data directory %s: %v", d.path, err)
	}
	defer dir.Close()
	_, err = dir.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
func (d *DataDirectory) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return d.path
}

type VersionFile struct{ path string }

func (v *VersionFile) Exists() (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return exists(v.path)
}
func (v *VersionFile) Read() (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data, err := ioutil.ReadFile(v.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read version file %s: %v", v.path, err)
	}
	txt := strings.TrimSpace(string(data))
	vp, err := ParseEtcdVersionPair(txt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse etcd '<version>/<storage-version>' string from version.txt file contents '%s': %v", txt, err)
	}
	return vp, nil
}
func (v *VersionFile) Write(vp *EtcdVersionPair) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	data := []byte(fmt.Sprintf("%s/%s", vp.version, vp.storageVersion))
	return ioutil.WriteFile(v.path, data, 0666)
}
func exists(path string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
