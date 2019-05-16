package system

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/errors"
)

type Validator interface {
	Name() string
	Validate(SysSpec) (error, error)
}
type Reporter interface {
	Report(string, string, ValidationResultType) error
}

func Validate(spec SysSpec, validators []Validator) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	var warns []error
	for _, v := range validators {
		fmt.Printf("Validating %s...\n", v.Name())
		warn, err := v.Validate(spec)
		errs = append(errs, err)
		warns = append(warns, warn)
	}
	return errors.NewAggregate(warns), errors.NewAggregate(errs)
}
func ValidateSpec(spec SysSpec, runtime string) (error, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var osValidators = []Validator{&OSValidator{Reporter: DefaultReporter}, &KernelValidator{Reporter: DefaultReporter}, &CgroupsValidator{Reporter: DefaultReporter}, &packageValidator{reporter: DefaultReporter}}
	var dockerValidators = []Validator{&DockerValidator{Reporter: DefaultReporter}}
	validators := osValidators
	switch runtime {
	case "docker":
		validators = append(validators, dockerValidators...)
	}
	return Validate(spec, validators)
}
