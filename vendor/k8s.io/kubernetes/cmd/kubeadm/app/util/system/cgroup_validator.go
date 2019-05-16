package system

import (
	"bufio"
	goformat "fmt"
	"github.com/pkg/errors"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var _ Validator = &CgroupsValidator{}

type CgroupsValidator struct{ Reporter Reporter }

func (c *CgroupsValidator) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "cgroups"
}

const (
	cgroupsConfigPrefix = "CGROUPS_"
)

func (c *CgroupsValidator) Validate(spec SysSpec) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	subsystems, err := c.getCgroupSubsystems()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cgroup subsystems")
	}
	return nil, c.validateCgroupSubsystems(spec.Cgroups, subsystems)
}
func (c *CgroupsValidator) validateCgroupSubsystems(cgroupSpec, subsystems []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	missing := []string{}
	for _, cgroup := range cgroupSpec {
		found := false
		for _, subsystem := range subsystems {
			if cgroup == subsystem {
				found = true
				break
			}
		}
		item := cgroupsConfigPrefix + strings.ToUpper(cgroup)
		if found {
			c.Reporter.Report(item, "enabled", good)
		} else {
			c.Reporter.Report(item, "missing", bad)
			missing = append(missing, cgroup)
		}
	}
	if len(missing) > 0 {
		return errors.Errorf("missing cgroups: %s", strings.Join(missing, " "))
	}
	return nil
}
func (c *CgroupsValidator) getCgroupSubsystems() ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f, err := os.Open("/proc/cgroups")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	subsystems := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return nil, err
		}
		text := s.Text()
		if text[0] != '#' {
			parts := strings.Fields(text)
			if len(parts) >= 4 && parts[3] != "0" {
				subsystems = append(subsystems, parts[0])
			}
		}
	}
	return subsystems, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
