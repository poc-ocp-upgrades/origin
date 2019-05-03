package graph

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	baseRepoRegex = regexp.MustCompile("[a-zA-Z0-9]+\\.([a-z0-9])+\\/.+")
)

type PackageError struct {
	ImportStack []string
	Pos         string
	Err         string
}

func (e *PackageError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Err
}

type Package struct {
	Dir         string
	ImportPath  string
	Imports     []string
	TestImports []string
	Error       *PackageError
}
type PackageList struct{ Packages []Package }

func (p *PackageList) Add(pkg Package) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.Packages = append(p.Packages, pkg)
}
func getPackageMetadata(entrypoints, ignoredPaths, buildTags []string) (*PackageList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{"list", "-e", "--json"}
	if len(buildTags) > 0 {
		args = append(args, append([]string{"--tags"}, buildTags...)...)
	}
	golist := exec.Command("go", append(args, entrypoints...)...)
	r, w := io.Pipe()
	golist.Stdout = w
	golist.Stderr = os.Stderr
	defer r.Close()
	done := make(chan bool)
	pkgs := &PackageList{}
	go func(list *PackageList) {
		decoder := json.NewDecoder(r)
		for {
			var pkg Package
			err := decoder.Decode(&pkg)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			if pkg.Error != nil {
				if containsPrefix(pkg.ImportPath, ignoredPaths) {
					fmt.Fprintf(os.Stderr, "warning: error encountered on excluded path %s: %v\n", pkg.ImportPath, pkg.Error)
					continue
				}
				fmt.Fprintf(os.Stderr, "error: %v\n", pkg.Error)
				golist.Process.Kill()
				close(done)
				return
			}
			list.Add(pkg)
		}
		close(done)
	}(pkgs)
	if err := golist.Run(); err != nil {
		w.Close()
		return nil, err
	}
	w.Close()
	<-done
	return pkgs, nil
}
func containsPrefix(needle string, prefixes []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, prefix := range prefixes {
		if strings.HasPrefix(needle, prefix) {
			return true
		}
	}
	return false
}
func isValidPackagePath(path string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return baseRepoRegex.Match([]byte(path))
}
