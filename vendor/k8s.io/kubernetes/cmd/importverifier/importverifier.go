package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	goformat "fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	goos "os"
	"os/exec"
	"path/filepath"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type Package struct {
	Dir          string   `yaml:",omitempty"`
	ImportPath   string   `yaml:",omitempty"`
	Imports      []string `yaml:",omitempty"`
	TestImports  []string `yaml:",omitempty"`
	XTestImports []string `yaml:",omitempty"`
}
type ImportRestriction struct {
	BaseDir         string   `yaml:"baseImportPath"`
	IgnoredSubTrees []string `yaml:"ignoredSubTrees,omitempty"`
	AllowedImports  []string `yaml:"allowedImports"`
	ExcludeTests    bool     `yaml:"excludeTests"`
}

func (i *ImportRestriction) ForbiddenImportsFor(pkg Package) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if restricted, err := i.isRestrictedDir(pkg.Dir); err != nil {
		return []string{}, err
	} else if !restricted {
		return []string{}, nil
	}
	return i.forbiddenImportsFor(pkg), nil
}
func (i *ImportRestriction) isRestrictedDir(dir string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if under, err := isPathUnder(i.BaseDir, dir); err != nil {
		return false, err
	} else if !under {
		return false, nil
	}
	for _, ignored := range i.IgnoredSubTrees {
		if under, err := isPathUnder(ignored, dir); err != nil {
			return false, err
		} else if under {
			return false, nil
		}
	}
	return true, nil
}
func isPathUnder(base, path string) (bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	absBase, err := filepath.Abs(base)
	if err != nil {
		return false, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	relPath, err := filepath.Rel(absBase, absPath)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(relPath, ".."), nil
}
func (i *ImportRestriction) forbiddenImportsFor(pkg Package) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	forbiddenImportSet := map[string]struct{}{}
	imports := pkg.Imports
	if !i.ExcludeTests {
		imports = append(imports, append(pkg.TestImports, pkg.XTestImports...)...)
	}
	for _, imp := range imports {
		path := extractVendorPath(imp)
		if i.isForbidden(path) {
			forbiddenImportSet[path] = struct{}{}
		}
	}
	var forbiddenImports []string
	for imp := range forbiddenImportSet {
		forbiddenImports = append(forbiddenImports, imp)
	}
	return forbiddenImports
}
func extractVendorPath(path string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	vendorPath := "/vendor/"
	if !strings.Contains(path, vendorPath) {
		return path
	}
	return path[strings.Index(path, vendorPath)+len(vendorPath):]
}
func (i *ImportRestriction) isForbidden(imp string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	importsBelowRoot := strings.HasPrefix(imp, rootPackage)
	importsBelowBase := strings.HasPrefix(imp, i.BaseDir)
	importsAllowed := false
	for _, allowed := range i.AllowedImports {
		exactlyImportsAllowed := imp == allowed
		importsBelowAllowed := strings.HasPrefix(imp, fmt.Sprintf("%s/", allowed))
		importsAllowed = importsAllowed || (importsBelowAllowed || exactlyImportsAllowed)
	}
	return importsBelowRoot && !importsBelowBase && !importsAllowed
}

var rootPackage string

func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s ROOT RESTRICTIONS.yaml", os.Args[0])
	}
	rootPackage = os.Args[1]
	configFile := os.Args[2]
	importRestrictions, err := loadImportRestrictions(configFile)
	if err != nil {
		log.Fatalf("Failed to load import restrictions: %v", err)
	}
	foundForbiddenImports := false
	for _, restriction := range importRestrictions {
		log.Printf("Inspecting imports under %s...\n", restriction.BaseDir)
		packages, err := resolvePackageTree(restriction.BaseDir)
		if err != nil {
			log.Fatalf("Failed to resolve package tree: %v", err)
		} else if len(packages) == 0 {
			log.Fatalf("Found no packages under tree %s", restriction.BaseDir)
		}
		log.Printf("- validating imports for %d packages in the tree", len(packages))
		restrictionViolated := false
		for _, pkg := range packages {
			if forbidden, err := restriction.ForbiddenImportsFor(pkg); err != nil {
				log.Fatalf("-- failed to validate imports: %v", err)
			} else if len(forbidden) != 0 {
				logForbiddenPackages(pkg.ImportPath, forbidden)
				restrictionViolated = true
			}
		}
		if restrictionViolated {
			foundForbiddenImports = true
			log.Println("- FAIL")
		} else {
			log.Println("- OK")
		}
	}
	if foundForbiddenImports {
		os.Exit(1)
	}
}
func loadImportRestrictions(configFile string) ([]ImportRestriction, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration from %s: %v", configFile, err)
	}
	var importRestrictions []ImportRestriction
	if err := yaml.Unmarshal(config, &importRestrictions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal from %s: %v", configFile, err)
	}
	return importRestrictions, nil
}
func resolvePackageTree(treeBase string) ([]Package, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := "go"
	args := []string{"list", "-json", fmt.Sprintf("%s...", treeBase)}
	stdout, err := exec.Command(cmd, args...).Output()
	if err != nil {
		var message string
		if ee, ok := err.(*exec.ExitError); ok {
			message = fmt.Sprintf("%v\n%v", ee, string(ee.Stderr))
		} else {
			message = fmt.Sprintf("%v", err)
		}
		return nil, fmt.Errorf("failed to run `%s %s`: %v", cmd, strings.Join(args, " "), message)
	}
	packages, err := decodePackages(bytes.NewReader(stdout))
	if err != nil {
		return nil, fmt.Errorf("failed to decode packages: %v", err)
	}
	return packages, nil
}
func decodePackages(r io.Reader) ([]Package, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var packages []Package
	decoder := json.NewDecoder(r)
	for decoder.More() {
		var pkg Package
		if err := decoder.Decode(&pkg); err != nil {
			return nil, fmt.Errorf("invalid package: %v", err)
		}
		packages = append(packages, pkg)
	}
	return packages, nil
}
func logForbiddenPackages(base string, forbidden []string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	log.Printf("-- found forbidden imports for %s:\n", base)
	for _, forbiddenPackage := range forbidden {
		log.Printf("--- %s\n", forbiddenPackage)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
