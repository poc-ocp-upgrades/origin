package apidocs

import (
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"github.com/go-openapi/spec"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func RefType(s *spec.Schema) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.TrimPrefix(s.Ref.String(), "#/definitions/")
}
func FriendlyTypeName(s *spec.Schema) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	refType := RefType(s)
	if refType == "" {
		return s.Type[0]
	}
	parts := strings.Split(refType, ".")
	return strings.Join(parts[len(parts)-2:], ".")
}
func EscapeMediaTypes(mediatypes []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rv := make([]string, len(mediatypes))
	for i, mediatype := range mediatypes {
		rv[i] = mediatype
		if mediatype == "*/*" {
			rv[i] = `\*/*`
		}
	}
	return rv
}
func GroupVersionKinds(s spec.Schema) []schema.GroupVersionKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := s.Extensions["x-kubernetes-group-version-kind"]
	if e == nil {
		return nil
	}
	gvks := make([]schema.GroupVersionKind, 0, len(e.([]interface{})))
	for _, gvk := range e.([]interface{}) {
		gvk := gvk.(map[string]interface{})
		gvks = append(gvks, schema.GroupVersionKind{Group: gvk["group"].(string), Version: gvk["version"].(string), Kind: gvk["kind"].(string)})
	}
	return gvks
}

var opNames = []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch"}

func Operations(path spec.PathItem) map[string]*spec.Operation {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ops := make(map[string]*spec.Operation, len(opNames))
	v := reflect.ValueOf(path)
	for _, opName := range opNames {
		op := v.FieldByName(opName).Interface().(*spec.Operation)
		if op != nil {
			ops[opName] = op
		}
	}
	return ops
}

var envStyleRegexp = regexp.MustCompile(`\{[^}]+\}`)

func EnvStyle(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return envStyleRegexp.ReplaceAllStringFunc(s, func(s string) string {
		return "$" + strings.ToUpper(s[1:len(s)-1])
	})
}

var alreadyPluralSuffixes = []string{"versions", "constraints", "endpoints"}

func Pluralise(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := strings.ToLower(s)
	for _, ss := range alreadyPluralSuffixes {
		if strings.HasSuffix(l, ss) {
			return s
		}
	}
	if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	}
	return s + "s"
}
func SortedKeys(m interface{}, t reflect.Type) interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("wrong type")
	}
	s := reflect.MakeSlice(reflect.SliceOf(v.Type().Key()), v.Len(), v.Len())
	for i, k := range v.MapKeys() {
		s.Index(i).Set(k)
	}
	s = s.Convert(t)
	sort.Sort(s.Interface().(sort.Interface))
	return s.Interface()
}
func ToUpper(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
func ReverseStringSlice(s []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := make([]string, len(s))
	for i := range s {
		r[len(r)-1-i] = s[i]
	}
	return r
}
