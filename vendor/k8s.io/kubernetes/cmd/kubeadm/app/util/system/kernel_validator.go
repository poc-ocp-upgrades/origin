package system

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var _ Validator = &KernelValidator{}

type KernelValidator struct {
	kernelRelease string
	Reporter      Reporter
}

func (k *KernelValidator) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "kernel"
}

type kConfigOption string

const (
	builtIn           kConfigOption = "y"
	asModule          kConfigOption = "m"
	leftOut           kConfigOption = "n"
	validKConfigRegex               = "^CONFIG_[A-Z0-9_]+=[myn]"
	kConfigPrefix                   = "CONFIG_"
)

func (k *KernelValidator) Validate(spec SysSpec) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	helper := KernelValidatorHelperImpl{}
	release, err := helper.GetKernelReleaseVersion()
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed to get kernel release")
	}
	k.kernelRelease = release
	var errs []error
	errs = append(errs, k.validateKernelVersion(spec.KernelSpec))
	if len(spec.KernelSpec.Required) > 0 || len(spec.KernelSpec.Forbidden) > 0 || len(spec.KernelSpec.Optional) > 0 {
		errs = append(errs, k.validateKernelConfig(spec.KernelSpec))
	}
	return nil, errors.NewAggregate(errs)
}
func (k *KernelValidator) validateKernelVersion(kSpec KernelSpec) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versionRegexps := kSpec.Versions
	for _, versionRegexp := range versionRegexps {
		r := regexp.MustCompile(versionRegexp)
		if r.MatchString(k.kernelRelease) {
			k.Reporter.Report("KERNEL_VERSION", k.kernelRelease, good)
			return nil
		}
	}
	k.Reporter.Report("KERNEL_VERSION", k.kernelRelease, bad)
	return pkgerrors.Errorf("unsupported kernel release: %s", k.kernelRelease)
}
func (k *KernelValidator) validateKernelConfig(kSpec KernelSpec) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	allConfig, err := k.getKernelConfig()
	if err != nil {
		return pkgerrors.Wrap(err, "failed to parse kernel config")
	}
	return k.validateCachedKernelConfig(allConfig, kSpec)
}
func (k *KernelValidator) validateCachedKernelConfig(allConfig map[string]kConfigOption, kSpec KernelSpec) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	badConfigs := []string{}
	reportAndRecord := func(name, msg, desc string, result ValidationResultType) {
		if result == bad {
			badConfigs = append(badConfigs, name)
		}
		if result != good && desc != "" {
			msg = msg + " - " + desc
		}
		k.Reporter.Report(name, msg, result)
	}
	const (
		required = iota
		optional
		forbidden
	)
	validateOpt := func(config KernelConfig, expect int) {
		var found, missing ValidationResultType
		switch expect {
		case required:
			found, missing = good, bad
		case optional:
			found, missing = good, warn
		case forbidden:
			found, missing = bad, good
		}
		var name string
		var opt kConfigOption
		var ok bool
		for _, name = range append([]string{config.Name}, config.Aliases...) {
			name = kConfigPrefix + name
			if opt, ok = allConfig[name]; ok {
				break
			}
		}
		if !ok {
			reportAndRecord(name, "not set", config.Description, missing)
			return
		}
		switch opt {
		case builtIn:
			reportAndRecord(name, "enabled", config.Description, found)
		case asModule:
			reportAndRecord(name, "enabled (as module)", config.Description, found)
		case leftOut:
			reportAndRecord(name, "disabled", config.Description, missing)
		default:
			reportAndRecord(name, fmt.Sprintf("unknown option: %s", opt), config.Description, missing)
		}
	}
	for _, config := range kSpec.Required {
		validateOpt(config, required)
	}
	for _, config := range kSpec.Optional {
		validateOpt(config, optional)
	}
	for _, config := range kSpec.Forbidden {
		validateOpt(config, forbidden)
	}
	if len(badConfigs) > 0 {
		return pkgerrors.Errorf("unexpected kernel config: %s", strings.Join(badConfigs, " "))
	}
	return nil
}
func (k *KernelValidator) getKernelConfigReader() (io.Reader, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	possibePaths := []string{"/proc/config.gz", "/boot/config-" + k.kernelRelease, "/usr/src/linux-" + k.kernelRelease + "/.config", "/usr/src/linux/.config", "/usr/lib/modules/" + k.kernelRelease + "/config", "/usr/lib/ostree-boot/config-" + k.kernelRelease, "/usr/lib/kernel/config-" + k.kernelRelease, "/usr/src/linux-headers-" + k.kernelRelease + "/.config", "/lib/modules/" + k.kernelRelease + "/build/.config"}
	configsModule := "configs"
	modprobeCmd := "modprobe"
	loadModule := false
	for {
		for _, path := range possibePaths {
			_, err := os.Stat(path)
			if err != nil {
				continue
			}
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			var r io.Reader
			r = bytes.NewReader(b)
			if filepath.Ext(path) == ".gz" {
				r, err = gzip.NewReader(r)
				if err != nil {
					return nil, err
				}
			}
			return r, nil
		}
		if loadModule {
			break
		}
		output, err := exec.Command(modprobeCmd, configsModule).CombinedOutput()
		if err != nil {
			return nil, pkgerrors.Wrapf(err, "unable to load kernel module: %q, output: %q, err", configsModule, output)
		}
		defer exec.Command(modprobeCmd, "-r", configsModule).Run()
		loadModule = true
	}
	return nil, pkgerrors.Errorf("no config path in %v is available", possibePaths)
}
func (k *KernelValidator) getKernelConfig() (map[string]kConfigOption, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	r, err := k.getKernelConfigReader()
	if err != nil {
		return nil, err
	}
	return k.parseKernelConfig(r)
}
func (k *KernelValidator) parseKernelConfig(r io.Reader) (map[string]kConfigOption, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	config := map[string]kConfigOption{}
	regex := regexp.MustCompile(validKConfigRegex)
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}
		line := strings.TrimSpace(s.Text())
		if !regex.MatchString(line) {
			continue
		}
		fields := strings.Split(line, "=")
		if len(fields) != 2 {
			klog.Errorf("Unexpected fields number in config %q", line)
			continue
		}
		config[fields[0]] = kConfigOption(fields[1])
	}
	return config, nil
}
