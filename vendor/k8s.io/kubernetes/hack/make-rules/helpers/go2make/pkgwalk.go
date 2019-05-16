package main

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"sort"
)

type VisitFunc func(importPath string, absPath string) error

var ErrSkipPkg = fmt.Errorf("package skipped")

func WalkPkg(pkgName string, visit VisitFunc) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pkg, err := findPackage(pkgName)
	if err != nil {
		return err
	}
	if err := visit(pkg.ImportPath, pkg.Dir); err == ErrSkipPkg {
		return nil
	} else if err != nil {
		return err
	}
	infos, err := readDirInfos(pkg.Dir)
	if err != nil {
		return err
	}
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		name := info.Name()
		if name[0] == '_' || (len(name) > 1 && name[0] == '.') || name == "testdata" {
			continue
		}
		err := WalkPkg(pkgName+"/"+name, visit)
		if err != nil {
			return err
		}
	}
	return nil
}
func findPackage(pkgName string) (*build.Package, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	debug("find", pkgName)
	pkg, err := build.Import(pkgName, getwd(), build.FindOnly)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}
func readDirInfos(dirPath string) ([]os.FileInfo, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	names, err := readDirNames(dirPath)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	infos := make([]os.FileInfo, 0, len(names))
	for _, n := range names {
		info, err := os.Stat(path.Join(dirPath, n))
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	return infos, nil
}
func readDirNames(dirPath string) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	d, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	return names, nil
}
