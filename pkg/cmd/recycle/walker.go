package recycle

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

type walkFunc func(path string, info os.FileInfo) error
type walker struct {
	walkFn		walkFunc
	fsuid		int64
	lstat		func(path string) (os.FileInfo, error)
	getuid		func(info os.FileInfo) (int64, error)
	setfsuid	func(uid int) error
	readDirNames	func(dirname string) ([]string, error)
}
type walkError struct {
	path		string
	info		os.FileInfo
	operation	string
	err		error
}

func (w walkError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var mode interface{} = "unknown"
	if w.info != nil {
		mode = w.info.Mode()
	}
	return fmt.Sprintf("%s (%s), %s: %s", w.path, mode, w.operation, w.err)
}
func makeWalkError(path string, info os.FileInfo, err error, operation string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, isWalkError := err.(walkError); isWalkError {
		return err
	}
	return walkError{path, info, operation, err}
}
func newWalker(walkFn walkFunc) *walker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &walker{walkFn: walkFn, fsuid: int64(os.Getuid()), lstat: os.Lstat, getuid: getuid, setfsuid: setfsuid, readDirNames: readDirNames}
}
func (w *walker) Walk(root string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	info, err := w.lstat(root)
	if err != nil {
		return makeWalkError(root, info, err, "lstat root dir")
	}
	err = w.becomeOwner(info)
	if err != nil {
		return makeWalkError(root, info, err, "becoming root dir owner")
	}
	return w.walk(root, info)
}
func (w *walker) walk(path string, info os.FileInfo) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	if info.IsDir() {
		previousFSuid := w.fsuid
		err = w.becomeOwner(info)
		if err != nil {
			return makeWalkError(path, info, err, "becoming dir owner")
		}
		names, err := w.readDirNames(path)
		if err != nil {
			return makeWalkError(path, info, err, fmt.Sprintf("reading dir names as %d", w.fsuid))
		}
		for _, name := range names {
			filename := filepath.Join(path, name)
			fileInfo, err := w.lstat(filename)
			if err != nil {
				return makeWalkError(path, info, err, fmt.Sprintf("lstat child as %d", w.fsuid))
			}
			err = w.walk(filename, fileInfo)
			if err != nil {
				return err
			}
		}
		err = w.becomeUid(previousFSuid)
		if err != nil {
			return makeWalkError(path, info, err, "returning to previous uid")
		}
	}
	err = w.walkFn(path, info)
	if err != nil {
		return makeWalkError(path, info, err, "calling walkFn")
	}
	return nil
}
func (w *walker) becomeOwner(info os.FileInfo) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uid, err := w.getuid(info)
	if err != nil {
		return err
	}
	return w.becomeUid(uid)
}
func (w *walker) becomeUid(uid int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if w.fsuid == uid {
		return nil
	}
	if err := w.setfsuid(int(uid)); err != nil {
		return err
	}
	w.fsuid = uid
	return nil
}
func readDirNames(dirname string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}
