package filter

import (
	"errors"
	"fmt"
	goformat "fmt"
	"k8s.io/klog"
	goos "os"
	"reflect"
	"regexp"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var (
	None *F
)

func Regexp(fieldName, v string) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndRegexp(fieldName, v)
}
func NotRegexp(fieldName, v string) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndNotRegexp(fieldName, v)
}
func EqualInt(fieldName string, v int) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndEqualInt(fieldName, v)
}
func NotEqualInt(fieldName string, v int) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndNotEqualInt(fieldName, v)
}
func EqualBool(fieldName string, v bool) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndEqualBool(fieldName, v)
}
func NotEqualBool(fieldName string, v bool) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return (&F{}).AndNotEqualBool(fieldName, v)
}

type F struct{ predicates []filterPredicate }

func (fl *F) And(rest *F) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, rest.predicates...)
	return fl
}
func (fl *F) AndRegexp(fieldName, v string) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: equals, s: &v})
	return fl
}
func (fl *F) AndNotRegexp(fieldName, v string) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: notEquals, s: &v})
	return fl
}
func (fl *F) AndEqualInt(fieldName string, v int) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: equals, i: &v})
	return fl
}
func (fl *F) AndNotEqualInt(fieldName string, v int) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: notEquals, i: &v})
	return fl
}
func (fl *F) AndEqualBool(fieldName string, v bool) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: equals, b: &v})
	return fl
}
func (fl *F) AndNotEqualBool(fieldName string, v bool) *F {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	fl.predicates = append(fl.predicates, filterPredicate{fieldName: fieldName, op: notEquals, b: &v})
	return fl
}
func (fl *F) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(fl.predicates) == 1 {
		return fl.predicates[0].String()
	}
	var pl []string
	for _, p := range fl.predicates {
		pl = append(pl, "("+p.String()+")")
	}
	return strings.Join(pl, " ")
}
func (fl *F) Match(obj interface{}) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if fl == nil {
		return true
	}
	for _, p := range fl.predicates {
		if !p.match(obj) {
			return false
		}
	}
	return true
}

type filterOp int

const (
	equals    filterOp = iota
	notEquals filterOp = iota
)

type filterPredicate struct {
	fieldName string
	op        filterOp
	s         *string
	i         *int
	b         *bool
}

func (fp *filterPredicate) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var op string
	switch fp.op {
	case equals:
		op = "eq"
	case notEquals:
		op = "ne"
	default:
		op = "invalidOp"
	}
	var value string
	switch {
	case fp.s != nil:
		value = *fp.s
	case fp.i != nil:
		value = fmt.Sprintf("%d", *fp.i)
	case fp.b != nil:
		value = fmt.Sprintf("%t", *fp.b)
	default:
		value = "invalidValue"
	}
	return fmt.Sprintf("%s %s %s", fp.fieldName, op, value)
}
func (fp *filterPredicate) match(o interface{}) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v, err := extractValue(fp.fieldName, o)
	klog.V(6).Infof("extractValue(%q, %#v) = %v, %v", fp.fieldName, o, v, err)
	if err != nil {
		return false
	}
	var match bool
	switch x := v.(type) {
	case string:
		if fp.s == nil {
			return false
		}
		re, err := regexp.Compile(*fp.s)
		if err != nil {
			klog.Errorf("Match regexp %q is invalid: %v", *fp.s, err)
			return false
		}
		match = re.Match([]byte(x))
	case int:
		if fp.i == nil {
			return false
		}
		match = x == *fp.i
	case bool:
		if fp.b == nil {
			return false
		}
		match = x == *fp.b
	}
	switch fp.op {
	case equals:
		return match
	case notEquals:
		return !match
	}
	return false
}
func snakeToCamelCase(s string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(s, "_")
	var ret string
	for _, x := range parts {
		ret += strings.Title(x)
	}
	return ret
}
func extractValue(path string, o interface{}) (interface{}, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(path, ".")
	for _, f := range parts {
		v := reflect.ValueOf(o)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, errors.New("field is nil")
			}
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot get field from non-struct (%T)", o)
		}
		v = v.FieldByName(snakeToCamelCase(f))
		if !v.IsValid() {
			return nil, fmt.Errorf("cannot get field %q as it is not a valid field in %T", f, o)
		}
		if !v.CanInterface() {
			return nil, fmt.Errorf("cannot get field %q in obj of type %T", f, o)
		}
		o = v.Interface()
	}
	switch o.(type) {
	case string, int, bool:
		return o, nil
	}
	return nil, fmt.Errorf("unhandled object of type %T", o)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
