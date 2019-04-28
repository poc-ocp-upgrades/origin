package rsync

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
)

type copyStrategies []CopyStrategy

var _ CopyStrategy = copyStrategies{}

func (ss copyStrategies) Copy(source, destination *PathSpec, out, errOut io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	foundStrategy := false
	for _, s := range ss {
		errBuf := &bytes.Buffer{}
		err = s.Copy(source, destination, out, errBuf)
		if _, isSetupError := err.(strategySetupError); isSetupError {
			klog.V(4).Infof("Error output:\n%s", errBuf.String())
			fmt.Fprintf(errOut, "WARNING: cannot use %s: %v\n", s.String(), err.Error())
			continue
		}
		io.Copy(errOut, errBuf)
		foundStrategy = true
		break
	}
	if !foundStrategy {
		err = strategySetupError("No available strategies to copy.")
	}
	return err
}
func (ss copyStrategies) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs []error
	for _, s := range ss {
		err := s.Validate()
		if err != nil {
			errs = append(errs, fmt.Errorf("invalid %v strategy: %v", s, err))
		}
	}
	return errors.NewAggregate(errs)
}
func (ss copyStrategies) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := []string{}
	for _, s := range ss {
		names = append(names, s.String())
	}
	return strings.Join(names, ",")
}

type strategySetupError string

func (e strategySetupError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(e)
}
