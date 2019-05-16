package eventratelimit

import (
	goformat "fmt"
	"io"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/client-go/util/flowcontrol"
	api "k8s.io/kubernetes/pkg/apis/core"
	eventratelimitapi "k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit"
	"k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/validation"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

const PluginName = "EventRateLimit"

func Register(plugins *admission.Plugins) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	plugins.Register(PluginName, func(config io.Reader) (admission.Interface, error) {
		configuration, err := LoadConfiguration(config)
		if err != nil {
			return nil, err
		}
		if configuration != nil {
			if errs := validation.ValidateConfiguration(configuration); len(errs) != 0 {
				return nil, errs.ToAggregate()
			}
		}
		return newEventRateLimit(configuration, realClock{})
	})
}

type Plugin struct {
	*admission.Handler
	limitEnforcers []*limitEnforcer
}

var _ admission.ValidationInterface = &Plugin{}

func newEventRateLimit(config *eventratelimitapi.Configuration, clock flowcontrol.Clock) (*Plugin, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	limitEnforcers := make([]*limitEnforcer, 0, len(config.Limits))
	for _, limitConfig := range config.Limits {
		enforcer, err := newLimitEnforcer(limitConfig, clock)
		if err != nil {
			return nil, err
		}
		limitEnforcers = append(limitEnforcers, enforcer)
	}
	eventRateLimitAdmission := &Plugin{Handler: admission.NewHandler(admission.Create, admission.Update), limitEnforcers: limitEnforcers}
	return eventRateLimitAdmission, nil
}
func (a *Plugin) Validate(attr admission.Attributes) (err error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if attr.GetKind().GroupKind() != api.Kind("Event") {
		return nil
	}
	if attr.IsDryRun() {
		return nil
	}
	var errors []error
	for _, enforcer := range a.limitEnforcers {
		if err := enforcer.accept(attr); err != nil {
			errors = append(errors, err)
		}
	}
	if aggregatedErr := utilerrors.NewAggregate(errors); aggregatedErr != nil {
		return apierrors.NewTooManyRequestsError(aggregatedErr.Error())
	}
	return nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
