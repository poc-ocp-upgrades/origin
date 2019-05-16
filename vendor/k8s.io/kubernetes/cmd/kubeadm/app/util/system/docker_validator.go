package system

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"regexp"
)

var _ Validator = &DockerValidator{}

type DockerValidator struct{ Reporter Reporter }

func (d *DockerValidator) Name() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return "docker"
}

const (
	dockerConfigPrefix           = "DOCKER_"
	latestValidatedDockerVersion = "18.06"
)

func (d *DockerValidator) Validate(spec SysSpec) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if spec.RuntimeSpec.DockerSpec == nil {
		return nil, nil
	}
	c, err := client.NewClient(dockerEndpoint, "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create docker client")
	}
	info, err := c.Info(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get docker info")
	}
	return d.validateDockerInfo(spec.RuntimeSpec.DockerSpec, info)
}
func (d *DockerValidator) validateDockerInfo(spec *DockerSpec, info types.Info) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matched := false
	for _, v := range spec.Version {
		r := regexp.MustCompile(v)
		if r.MatchString(info.ServerVersion) {
			d.Reporter.Report(dockerConfigPrefix+"VERSION", info.ServerVersion, good)
			matched = true
		}
	}
	if !matched {
		ver := `\d{2}\.\d+\.\d+(?:-[a-z]{2})?`
		r := regexp.MustCompile(ver)
		if r.MatchString(info.ServerVersion) {
			d.Reporter.Report(dockerConfigPrefix+"VERSION", info.ServerVersion, good)
			w := errors.Errorf("this Docker version is not on the list of validated versions: %s. Latest validated version: %s", info.ServerVersion, latestValidatedDockerVersion)
			return w, nil
		}
		d.Reporter.Report(dockerConfigPrefix+"VERSION", info.ServerVersion, bad)
		return nil, errors.Errorf("unsupported docker version: %s", info.ServerVersion)
	}
	item := dockerConfigPrefix + "GRAPH_DRIVER"
	for _, gd := range spec.GraphDriver {
		if info.Driver == gd {
			d.Reporter.Report(item, info.Driver, good)
			return nil, nil
		}
	}
	d.Reporter.Report(item, info.Driver, bad)
	return nil, errors.Errorf("unsupported graph driver: %s", info.Driver)
}
