package apparmor

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"k8s.io/api/core/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	utilfile "k8s.io/kubernetes/pkg/util/file"
	"os"
	"path"
	"strings"
)

var isDisabledBuild bool

type Validator interface {
	Validate(pod *v1.Pod) error
	ValidateHost() error
}

func NewValidator(runtime string) Validator {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := validateHost(runtime); err != nil {
		return &validator{validateHostErr: err}
	}
	appArmorFS, err := getAppArmorFS()
	if err != nil {
		return &validator{validateHostErr: fmt.Errorf("error finding AppArmor FS: %v", err)}
	}
	return &validator{appArmorFS: appArmorFS}
}

type validator struct {
	validateHostErr error
	appArmorFS      string
}

func (v *validator) Validate(pod *v1.Pod) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !isRequired(pod) {
		return nil
	}
	if v.ValidateHost() != nil {
		return v.validateHostErr
	}
	loadedProfiles, err := v.getLoadedProfiles()
	if err != nil {
		return fmt.Errorf("could not read loaded profiles: %v", err)
	}
	for _, container := range pod.Spec.InitContainers {
		if err := validateProfile(GetProfileName(pod, container.Name), loadedProfiles); err != nil {
			return err
		}
	}
	for _, container := range pod.Spec.Containers {
		if err := validateProfile(GetProfileName(pod, container.Name), loadedProfiles); err != nil {
			return err
		}
	}
	return nil
}
func (v *validator) ValidateHost() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v.validateHostErr
}
func validateHost(runtime string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !utilfeature.DefaultFeatureGate.Enabled(features.AppArmor) {
		return errors.New("AppArmor disabled by feature-gate")
	}
	if isDisabledBuild {
		return errors.New("Binary not compiled for linux")
	}
	if !IsAppArmorEnabled() {
		return errors.New("AppArmor is not enabled on the host")
	}
	if runtime != kubetypes.DockerContainerRuntime && runtime != kubetypes.RemoteContainerRuntime {
		return fmt.Errorf("AppArmor is only enabled for 'docker' and 'remote' runtimes. Found: %q.", runtime)
	}
	return nil
}
func validateProfile(profile string, loadedProfiles map[string]bool) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if err := ValidateProfileFormat(profile); err != nil {
		return err
	}
	if strings.HasPrefix(profile, ProfileNamePrefix) {
		profileName := strings.TrimPrefix(profile, ProfileNamePrefix)
		if !loadedProfiles[profileName] {
			return fmt.Errorf("profile %q is not loaded", profileName)
		}
	}
	return nil
}
func ValidateProfileFormat(profile string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if profile == "" || profile == ProfileRuntimeDefault || profile == ProfileNameUnconfined {
		return nil
	}
	if !strings.HasPrefix(profile, ProfileNamePrefix) {
		return fmt.Errorf("invalid AppArmor profile name: %q", profile)
	}
	return nil
}
func (v *validator) getLoadedProfiles() (map[string]bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	profilesPath := path.Join(v.appArmorFS, "profiles")
	profilesFile, err := os.Open(profilesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", profilesPath, err)
	}
	defer profilesFile.Close()
	profiles := map[string]bool{}
	scanner := bufio.NewScanner(profilesFile)
	for scanner.Scan() {
		profileName := parseProfileName(scanner.Text())
		if profileName == "" {
			continue
		}
		profiles[profileName] = true
	}
	return profiles, nil
}
func parseProfileName(profileLine string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	modeIndex := strings.IndexRune(profileLine, '(')
	if modeIndex < 0 {
		return ""
	}
	return strings.TrimSpace(profileLine[:modeIndex])
}
func getAppArmorFS() (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	mountsFile, err := os.Open("/proc/mounts")
	if err != nil {
		return "", fmt.Errorf("could not open /proc/mounts: %v", err)
	}
	defer mountsFile.Close()
	scanner := bufio.NewScanner(mountsFile)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}
		if fields[2] == "securityfs" {
			appArmorFS := path.Join(fields[1], "apparmor")
			if ok, err := utilfile.FileExists(appArmorFS); !ok {
				msg := fmt.Sprintf("path %s does not exist", appArmorFS)
				if err != nil {
					return "", fmt.Errorf("%s: %v", msg, err)
				} else {
					return "", errors.New(msg)
				}
			} else {
				return appArmorFS, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning mounts: %v", err)
	}
	return "", errors.New("securityfs not found")
}
func IsAppArmorEnabled() bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if _, err := os.Stat("/sys/kernel/security/apparmor"); err == nil && os.Getenv("container") == "" {
		if _, err = os.Stat("/sbin/apparmor_parser"); err == nil {
			buf, err := ioutil.ReadFile("/sys/module/apparmor/parameters/enabled")
			return err == nil && len(buf) > 1 && buf[0] == 'Y'
		}
	}
	return false
}
