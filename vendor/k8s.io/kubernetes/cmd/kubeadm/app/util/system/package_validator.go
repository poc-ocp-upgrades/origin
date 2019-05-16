package system

import (
	"fmt"
	"github.com/blang/semver"
	pkgerrors "github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
	"os/exec"
	"strings"
)

const semVerDotsCount int = 2

type packageManager interface {
	getPackageVersion(packageName string) (string, error)
}

func newPackageManager() (packageManager, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if m, ok := newDPKG(); ok {
		return m, nil
	}
	return nil, pkgerrors.New("failed to find package manager")
}

type dpkg struct{}

func newDPKG() (packageManager, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	_, err := exec.LookPath("dpkg-query")
	if err != nil {
		return nil, false
	}
	return dpkg{}, true
}
func (_ dpkg) getPackageVersion(packageName string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	output, err := exec.Command("dpkg-query", "--show", "--showformat='${Version}'", packageName).Output()
	if err != nil {
		return "", pkgerrors.Wrap(err, "dpkg-query failed")
	}
	version := extractUpstreamVersion(string(output))
	if version == "" {
		return "", pkgerrors.New("no version information")
	}
	return version, nil
}

type packageValidator struct {
	reporter      Reporter
	kernelRelease string
	osDistro      string
}

func (self *packageValidator) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "package"
}
func (self *packageValidator) Validate(spec SysSpec) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(spec.PackageSpecs) == 0 {
		return nil, nil
	}
	var err error
	if self.kernelRelease, err = getKernelRelease(); err != nil {
		return nil, err
	}
	if self.osDistro, err = getOSDistro(); err != nil {
		return nil, err
	}
	manager, err := newPackageManager()
	if err != nil {
		return nil, err
	}
	specs := applyPackageSpecOverride(spec.PackageSpecs, spec.PackageSpecOverrides, self.osDistro)
	return self.validate(specs, manager)
}
func (self *packageValidator) validate(packageSpecs []PackageSpec, manager packageManager) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for _, spec := range packageSpecs {
		packageName := resolvePackageName(spec.Name, self.kernelRelease)
		nameWithVerRange := fmt.Sprintf("%s (%s)", packageName, spec.VersionRange)
		version, err := manager.getPackageVersion(packageName)
		if err != nil {
			klog.V(1).Infof("Failed to get the version for the package %q: %s\n", packageName, err)
			errs = append(errs, err)
			self.reporter.Report(nameWithVerRange, "not installed", bad)
			continue
		}
		if spec.VersionRange == "" {
			self.reporter.Report(packageName, version, good)
			continue
		}
		sv, err := semver.Make(toSemVer(version))
		if err != nil {
			klog.Errorf("Failed to convert %q to semantic version: %s\n", version, err)
			errs = append(errs, err)
			self.reporter.Report(nameWithVerRange, "internal error", bad)
			continue
		}
		versionRange := semver.MustParseRange(toSemVerRange(spec.VersionRange))
		if versionRange(sv) {
			self.reporter.Report(nameWithVerRange, version, good)
		} else {
			errs = append(errs, pkgerrors.Errorf("package \"%s %s\" does not meet the spec \"%s (%s)\"", packageName, sv, packageName, spec.VersionRange))
			self.reporter.Report(nameWithVerRange, version, bad)
		}
	}
	return nil, errors.NewAggregate(errs)
}
func getKernelRelease() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	output, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "", pkgerrors.Wrap(err, "failed to get kernel release")
	}
	return strings.TrimSpace(string(output)), nil
}
func getOSDistro() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f := "/etc/lsb-release"
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return "", pkgerrors.Wrapf(err, "failed to read %q", f)
	}
	content := string(b)
	switch {
	case strings.Contains(content, "Ubuntu"):
		return "ubuntu", nil
	case strings.Contains(content, "Chrome OS"):
		return "cos", nil
	case strings.Contains(content, "CoreOS"):
		return "coreos", nil
	default:
		return "", pkgerrors.Errorf("failed to get OS distro: %s", content)
	}
}
func resolvePackageName(packageName string, kernelRelease string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	packageName = strings.Replace(packageName, "${KERNEL_RELEASE}", kernelRelease, -1)
	return packageName
}
func applyPackageSpecOverride(packageSpecs []PackageSpec, overrides []PackageSpecOverride, osDistro string) []PackageSpec {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var override *PackageSpecOverride
	for _, o := range overrides {
		if o.OSDistro == osDistro {
			override = &o
			break
		}
	}
	if override == nil {
		return packageSpecs
	}
	var out []PackageSpec
	subtractions := make(map[string]bool)
	for _, spec := range override.Subtractions {
		subtractions[spec.Name] = true
	}
	for _, spec := range packageSpecs {
		if _, ok := subtractions[spec.Name]; !ok {
			out = append(out, spec)
		}
	}
	return append(out, override.Additions...)
}
func extractUpstreamVersion(version string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	version = strings.Trim(version, " '")
	if i := strings.Index(version, ":"); i != -1 {
		version = version[i+1:]
	}
	if i := strings.Index(version, "-"); i != -1 {
		version = version[:i]
	}
	return version
}
func toSemVerRange(input string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var output []string
	fields := strings.Fields(input)
	for _, f := range fields {
		numDots, hasDigits := 0, false
		for _, c := range f {
			switch {
			case c == '.':
				numDots++
			case c >= '0' && c <= '9':
				hasDigits = true
			}
		}
		if hasDigits && numDots < semVerDotsCount {
			f = strings.TrimRight(f, " ")
			f += ".x"
		}
		output = append(output, f)
	}
	return strings.Join(output, " ")
}
func toSemVer(version string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if i := strings.IndexFunc(version, func(c rune) bool {
		if (c < '0' || c > '9') && c != '.' {
			return true
		}
		return false
	}); i != -1 {
		version = version[:i]
	}
	version = strings.TrimRight(version, ".")
	if version == "" {
		return ""
	}
	numDots := strings.Count(version, ".")
	switch {
	case numDots < semVerDotsCount:
		version += strings.Repeat(".0", semVerDotsCount-numDots)
	case numDots > semVerDotsCount:
		for numDots != semVerDotsCount {
			if i := strings.LastIndex(version, "."); i != -1 {
				version = version[:i]
				numDots--
			}
		}
	}
	var subs []string
	for _, s := range strings.Split(version, ".") {
		s := strings.TrimLeft(s, "0")
		if s == "" {
			s = "0"
		}
		subs = append(subs, s)
	}
	return strings.Join(subs, ".")
}
