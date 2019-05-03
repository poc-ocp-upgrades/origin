package meta

import (
 "fmt"
 "reflect"
 "strings"
)

func newArg(t reflect.Type) *arg {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := &arg{}
Loop:
 for {
  switch t.Kind() {
  case reflect.Ptr:
   ret.numPtr++
   t = t.Elem()
  default:
   ret.pkg = t.PkgPath()
   ret.typeName += t.Name()
   break Loop
  }
 }
 return ret
}

type arg struct {
 pkg, typeName string
 numPtr        int
}

func (a *arg) normalizedPkg() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 if a.pkg == "" {
  return ""
 }
 parts := strings.Split(a.pkg, "/")
 for i := 0; i < len(parts); i++ {
  if parts[i] == "vendor" {
   parts = parts[i+1:]
   break
  }
 }
 switch strings.Join(parts, "/") {
 case "google.golang.org/api/compute/v1":
  return "ga."
 case "google.golang.org/api/compute/v0.alpha":
  return "alpha."
 case "google.golang.org/api/compute/v0.beta":
  return "beta."
 default:
  panic(fmt.Errorf("unhandled package %q", a.pkg))
 }
}
func (a *arg) String() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var ret string
 for i := 0; i < a.numPtr; i++ {
  ret += "*"
 }
 ret += a.normalizedPkg()
 ret += a.typeName
 return ret
}
func newMethod(s *ServiceInfo, m reflect.Method) *Method {
 _logClusterCodePath()
 defer _logClusterCodePath()
 ret := &Method{ServiceInfo: s, m: m, kind: MethodOperation, ReturnType: ""}
 ret.init()
 return ret
}

type MethodKind int

const (
 MethodOperation MethodKind = iota
 MethodGet       MethodKind = iota
 MethodPaged     MethodKind = iota
)

type Method struct {
 *ServiceInfo
 m          reflect.Method
 kind       MethodKind
 ReturnType string
 ItemType   string
}

func (m *Method) IsOperation() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return m.kind == MethodOperation
}
func (m *Method) IsPaged() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return m.kind == MethodPaged
}
func (m *Method) IsGet() bool {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return m.kind == MethodGet
}
func (m *Method) argsSkip() int {
 _logClusterCodePath()
 defer _logClusterCodePath()
 switch m.keyType {
 case Zonal:
  return 4
 case Regional:
  return 4
 case Global:
  return 3
 }
 panic(fmt.Errorf("invalid KeyType %v", m.keyType))
}
func (m *Method) args(skip int, nameArgs bool, prefix []string) []string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var args []*arg
 fType := m.m.Func.Type()
 for i := 0; i < fType.NumIn(); i++ {
  t := fType.In(i)
  args = append(args, newArg(t))
 }
 var a []string
 for i := skip; i < fType.NumIn(); i++ {
  if nameArgs {
   a = append(a, fmt.Sprintf("arg%d %s", i-skip, args[i]))
  } else {
   a = append(a, args[i].String())
  }
 }
 return append(prefix, a...)
}
func (m *Method) init() {
 _logClusterCodePath()
 defer _logClusterCodePath()
 fType := m.m.Func.Type()
 if fType.NumIn() < m.argsSkip() {
  err := fmt.Errorf("method %q.%q, arity = %d which is less than required (< %d)", m.Service, m.Name(), fType.NumIn(), m.argsSkip())
  panic(err)
 }
 for i := 1; i < m.argsSkip(); i++ {
  if fType.In(i).Kind() != reflect.String {
   panic(fmt.Errorf("method %q.%q: skipped args can only be strings", m.Service, m.Name()))
  }
 }
 if fType.NumOut() != 1 || fType.Out(0).Kind() != reflect.Ptr || !strings.HasSuffix(fType.Out(0).Elem().Name(), "Call") {
  panic(fmt.Errorf("method %q.%q: generator only supports methods returning an *xxxCall object", m.Service, m.Name()))
 }
 returnType := fType.Out(0)
 returnTypeName := fType.Out(0).Elem().Name()
 doMethod, ok := returnType.MethodByName("Do")
 if !ok {
  panic(fmt.Errorf("method %q.%q: return type %q does not have a Do() method", m.Service, m.Name(), returnTypeName))
 }
 _, hasPages := returnType.MethodByName("Pages")
 switch doMethod.Func.Type().NumOut() {
 case 2:
  out0 := doMethod.Func.Type().Out(0)
  if out0.Kind() != reflect.Ptr {
   panic(fmt.Errorf("method %q.%q: return type %q of Do() = S, _; S must be pointer type (%v)", m.Service, m.Name(), returnTypeName, out0))
  }
  m.ReturnType = out0.Elem().Name()
  switch {
  case out0.Elem().Name() == "Operation":
   m.kind = MethodOperation
  case hasPages:
   m.kind = MethodPaged
   listType := out0.Elem()
   itemsField, ok := listType.FieldByName("Items")
   if !ok {
    panic(fmt.Errorf("method %q.%q: paged return type %q does not have a .Items field", m.Service, m.Name(), listType.Name()))
   }
   itemsType := itemsField.Type
   if itemsType.Kind() != reflect.Slice && itemsType.Elem().Kind() != reflect.Ptr {
    panic(fmt.Errorf("method %q.%q: paged return type %q.Items is not an array of pointers", m.Service, m.Name(), listType.Name()))
   }
   m.ItemType = itemsType.Elem().Elem().Name()
  default:
   m.kind = MethodGet
  }
  if doMethod.Func.Type().Out(1).Name() != "error" {
   panic(fmt.Errorf("method %q.%q: return type %q of Do() = S, T; T must be 'error'", m.Service, m.Name(), returnTypeName))
  }
  break
 default:
  panic(fmt.Errorf("method %q.%q: %q Do() return type is not handled by the generator", m.Service, m.Name(), returnTypeName))
 }
}
func (m *Method) Name() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return m.m.Name
}
func (m *Method) CallArgs() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 var args []string
 for i := m.argsSkip(); i < m.m.Func.Type().NumIn(); i++ {
  args = append(args, fmt.Sprintf("arg%d", i-m.argsSkip()))
 }
 if len(args) == 0 {
  return ""
 }
 return fmt.Sprintf(", %s", strings.Join(args, ", "))
}
func (m *Method) MockHookName() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 return m.m.Name + "Hook"
}
func (m *Method) MockHook() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 args := m.args(m.argsSkip(), false, []string{"context.Context", "*meta.Key"})
 if m.kind == MethodPaged {
  args = append(args, "*filter.F")
 }
 args = append(args, fmt.Sprintf("*%s", m.MockWrapType()))
 switch m.kind {
 case MethodOperation:
  return fmt.Sprintf("%v func(%v) error", m.MockHookName(), strings.Join(args, ", "))
 case MethodGet:
  return fmt.Sprintf("%v func(%v) (*%v.%v, error)", m.MockHookName(), strings.Join(args, ", "), m.Version(), m.ReturnType)
 case MethodPaged:
  return fmt.Sprintf("%v func(%v) ([]*%v.%v, error)", m.MockHookName(), strings.Join(args, ", "), m.Version(), m.ItemType)
 default:
  panic(fmt.Errorf("invalid method kind: %v", m.kind))
 }
}
func (m *Method) FcnArgs() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 args := m.args(m.argsSkip(), true, []string{"ctx context.Context", "key *meta.Key"})
 if m.kind == MethodPaged {
  args = append(args, "fl *filter.F")
 }
 switch m.kind {
 case MethodOperation:
  return fmt.Sprintf("%v(%v) error", m.m.Name, strings.Join(args, ", "))
 case MethodGet:
  return fmt.Sprintf("%v(%v) (*%v.%v, error)", m.m.Name, strings.Join(args, ", "), m.Version(), m.ReturnType)
 case MethodPaged:
  return fmt.Sprintf("%v(%v) ([]*%v.%v, error)", m.m.Name, strings.Join(args, ", "), m.Version(), m.ItemType)
 default:
  panic(fmt.Errorf("invalid method kind: %v", m.kind))
 }
}
func (m *Method) InterfaceFunc() string {
 _logClusterCodePath()
 defer _logClusterCodePath()
 args := []string{"context.Context", "*meta.Key"}
 args = m.args(m.argsSkip(), false, args)
 if m.kind == MethodPaged {
  args = append(args, "*filter.F")
 }
 switch m.kind {
 case MethodOperation:
  return fmt.Sprintf("%v(%v) error", m.m.Name, strings.Join(args, ", "))
 case MethodGet:
  return fmt.Sprintf("%v(%v) (*%v.%v, error)", m.m.Name, strings.Join(args, ", "), m.Version(), m.ReturnType)
 case MethodPaged:
  return fmt.Sprintf("%v(%v) ([]*%v.%v, error)", m.m.Name, strings.Join(args, ", "), m.Version(), m.ItemType)
 default:
  panic(fmt.Errorf("invalid method kind: %v", m.kind))
 }
}
