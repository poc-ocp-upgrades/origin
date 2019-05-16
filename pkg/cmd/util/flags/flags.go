package flags

import (
	"fmt"
	goformat "fmt"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apiserverflag "k8s.io/apiserver/pkg/util/flag"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

func apply(args map[string][]string, flags apiserverflag.NamedFlagSets, ignoreMissing bool) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var errs []error
	for key, value := range args {
		found := false
		for _, fs := range flags.FlagSets {
			if flag := fs.Lookup(key); flag != nil {
				found = true
				for _, s := range value {
					if err := flag.Value.Set(s); err != nil {
						errs = append(errs, field.Invalid(field.NewPath(key), s, fmt.Sprintf("could not be set: %v", err)))
						break
					}
				}
			}
		}
		if !found && !ignoreMissing {
			errs = append(errs, field.Invalid(field.NewPath("flag"), key, "is not a valid flag"))
		}
	}
	return errs
}
func Resolve(args map[string][]string, flagSet apiserverflag.NamedFlagSets) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apply(args, flagSet, false)
}
func ResolveIgnoreMissing(args map[string][]string, flagSet apiserverflag.NamedFlagSets) []error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return apply(args, flagSet, true)
}

type ComponentFlag struct {
	enabled    string
	disabled   string
	enabledSet func() bool
	calculated sets.String
	allowed    sets.String
	mappings   map[string][]string
}

func NewComponentFlag(mappings map[string][]string, allowed ...string) *ComponentFlag {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	set := sets.NewString(allowed...)
	return &ComponentFlag{allowed: set, mappings: mappings, enabled: strings.Join(set.List(), ","), enabledSet: func() bool {
		return false
	}}
}
func (f *ComponentFlag) DefaultEnable(components ...string) *ComponentFlag {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.enabled = strings.Join(f.allowed.Union(sets.NewString(components...)).List(), ",")
	return f
}
func (f *ComponentFlag) DefaultDisable(components ...string) *ComponentFlag {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.enabled = strings.Join(f.allowed.Difference(sets.NewString(components...)).List(), ",")
	return f
}
func (f *ComponentFlag) Disable(components ...string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	f.Calculated().Delete(components...)
}
func (f *ComponentFlag) Enabled(name string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return f.Calculated().Has(name)
}
func (f *ComponentFlag) Calculated() sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if f.calculated == nil {
		f.calculated = f.Expand(f.enabled).Difference(f.Expand(f.disabled)).Intersection(f.allowed)
	}
	return f.calculated
}
func (f *ComponentFlag) Validate() (sets.String, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	enabled := f.Expand(f.enabled)
	disabled := f.Expand(f.disabled)
	if diff := enabled.Difference(f.allowed); enabled.Len() > 0 && diff.Len() > 0 {
		return nil, fmt.Errorf("the following components are not recognized: %s", strings.Join(diff.List(), ", "))
	}
	if diff := disabled.Difference(f.allowed); disabled.Len() > 0 && diff.Len() > 0 {
		return nil, fmt.Errorf("the following components are not recognized: %s", strings.Join(diff.List(), ", "))
	}
	if inter := enabled.Intersection(disabled); f.enabledSet() && inter.Len() > 0 {
		return nil, fmt.Errorf("the following components can't be both disabled and enabled: %s", strings.Join(inter.List(), ", "))
	}
	return enabled.Difference(disabled), nil
}
func (f *ComponentFlag) Expand(value string) sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(value) == 0 {
		return sets.NewString()
	}
	items := strings.Split(value, ",")
	set := sets.NewString()
	for _, s := range items {
		if mapped, ok := f.mappings[s]; ok {
			set.Insert(mapped...)
		} else {
			set.Insert(s)
		}
	}
	return set
}
func (f *ComponentFlag) Allowed() sets.String {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return sets.NewString(f.allowed.List()...)
}
func (f *ComponentFlag) Mappings() map[string][]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	copied := make(map[string][]string)
	for k, v := range f.mappings {
		copiedV := make([]string, len(v))
		copy(copiedV, v)
		copied[k] = copiedV
	}
	return copied
}
func (f *ComponentFlag) Bind(flags *pflag.FlagSet, flagFormat, messagePrefix string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flags.StringVar(&f.enabled, fmt.Sprintf(flagFormat, "enable"), f.enabled, messagePrefix+" enable")
	flags.StringVar(&f.disabled, fmt.Sprintf(flagFormat, "disable"), f.disabled, messagePrefix+" disable")
	f.enabledSet = func() bool {
		return flags.Lookup(fmt.Sprintf(flagFormat, "enable")).Changed
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
