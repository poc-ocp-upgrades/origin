package abac

import (
	"bufio"
	"fmt"
	goformat "fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/apis/abac"
	_ "k8s.io/kubernetes/pkg/apis/abac/latest"
	"k8s.io/kubernetes/pkg/apis/abac/v0"
	"os"
	goos "os"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

type policyLoadError struct {
	path string
	line int
	data []byte
	err  error
}

func (p policyLoadError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if p.line >= 0 {
		return fmt.Sprintf("error reading policy file %s, line %d: %s: %v", p.path, p.line, string(p.data), p.err)
	}
	return fmt.Sprintf("error reading policy file %s: %v", p.path, p.err)
}

type policyList []*abac.Policy

func NewFromFile(path string) (policyList, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pl := make(policyList, 0)
	decoder := abac.Codecs.UniversalDecoder()
	i := 0
	unversionedLines := 0
	for scanner.Scan() {
		i++
		p := &abac.Policy{}
		b := scanner.Bytes()
		trimmed := strings.TrimSpace(string(b))
		if len(trimmed) == 0 || strings.HasPrefix(trimmed, "#") {
			continue
		}
		decodedObj, _, err := decoder.Decode(b, nil, nil)
		if err != nil {
			if !(runtime.IsMissingVersion(err) || runtime.IsMissingKind(err) || runtime.IsNotRegisteredError(err)) {
				return nil, policyLoadError{path, i, b, err}
			}
			unversionedLines++
			oldPolicy := &v0.Policy{}
			if err := runtime.DecodeInto(decoder, b, oldPolicy); err != nil {
				return nil, policyLoadError{path, i, b, err}
			}
			if err := abac.Scheme.Convert(oldPolicy, p, nil); err != nil {
				return nil, policyLoadError{path, i, b, err}
			}
			pl = append(pl, p)
			continue
		}
		decodedPolicy, ok := decodedObj.(*abac.Policy)
		if !ok {
			return nil, policyLoadError{path, i, b, fmt.Errorf("unrecognized object: %#v", decodedObj)}
		}
		pl = append(pl, decodedPolicy)
	}
	if unversionedLines > 0 {
		klog.Warningf("Policy file %s contained unversioned rules. See docs/admin/authorization.md#abac-mode for ABAC file format details.", path)
	}
	if err := scanner.Err(); err != nil {
		return nil, policyLoadError{path, -1, nil, err}
	}
	return pl, nil
}
func matches(p abac.Policy, a authorizer.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if subjectMatches(p, a.GetUser()) {
		if verbMatches(p, a) {
			if resourceMatches(p, a) {
				return true
			}
			if nonResourceMatches(p, a) {
				return true
			}
		}
	}
	return false
}
func subjectMatches(p abac.Policy, user user.Info) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	matched := false
	if user == nil {
		return false
	}
	username := user.GetName()
	groups := user.GetGroups()
	if len(p.Spec.User) > 0 {
		if p.Spec.User == "*" {
			matched = true
		} else {
			matched = p.Spec.User == username
			if !matched {
				return false
			}
		}
	}
	if len(p.Spec.Group) > 0 {
		if p.Spec.Group == "*" {
			matched = true
		} else {
			matched = false
			for _, group := range groups {
				if p.Spec.Group == group {
					matched = true
				}
			}
			if !matched {
				return false
			}
		}
	}
	return matched
}
func verbMatches(p abac.Policy, a authorizer.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.IsReadOnly() {
		return true
	}
	if !p.Spec.Readonly {
		return true
	}
	return false
}
func nonResourceMatches(p abac.Policy, a authorizer.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !a.IsResourceRequest() {
		if p.Spec.NonResourcePath == "*" {
			return true
		}
		if p.Spec.NonResourcePath == a.GetPath() {
			return true
		}
		if strings.HasSuffix(p.Spec.NonResourcePath, "*") && strings.HasPrefix(a.GetPath(), strings.TrimRight(p.Spec.NonResourcePath, "*")) {
			return true
		}
	}
	return false
}
func resourceMatches(p abac.Policy, a authorizer.Attributes) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if a.IsResourceRequest() {
		if p.Spec.Namespace == "*" || p.Spec.Namespace == a.GetNamespace() {
			if p.Spec.Resource == "*" || p.Spec.Resource == a.GetResource() {
				if p.Spec.APIGroup == "*" || p.Spec.APIGroup == a.GetAPIGroup() {
					return true
				}
			}
		}
	}
	return false
}
func (pl policyList) Authorize(a authorizer.Attributes) (authorizer.Decision, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	for _, p := range pl {
		if matches(*p, a) {
			return authorizer.DecisionAllow, "", nil
		}
	}
	return authorizer.DecisionNoOpinion, "No policy matched.", nil
}
func (pl policyList) RulesFor(user user.Info, namespace string) ([]authorizer.ResourceRuleInfo, []authorizer.NonResourceRuleInfo, bool, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var (
		resourceRules    []authorizer.ResourceRuleInfo
		nonResourceRules []authorizer.NonResourceRuleInfo
	)
	for _, p := range pl {
		if subjectMatches(*p, user) {
			if p.Spec.Namespace == "*" || p.Spec.Namespace == namespace {
				if len(p.Spec.Resource) > 0 {
					r := authorizer.DefaultResourceRuleInfo{Verbs: getVerbs(p.Spec.Readonly), APIGroups: []string{p.Spec.APIGroup}, Resources: []string{p.Spec.Resource}}
					var resourceRule authorizer.ResourceRuleInfo = &r
					resourceRules = append(resourceRules, resourceRule)
				}
				if len(p.Spec.NonResourcePath) > 0 {
					r := authorizer.DefaultNonResourceRuleInfo{Verbs: getVerbs(p.Spec.Readonly), NonResourceURLs: []string{p.Spec.NonResourcePath}}
					var nonResourceRule authorizer.NonResourceRuleInfo = &r
					nonResourceRules = append(nonResourceRules, nonResourceRule)
				}
			}
		}
	}
	return resourceRules, nonResourceRules, false, nil
}
func getVerbs(isReadOnly bool) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if isReadOnly {
		return []string{"get", "list", "watch"}
	}
	return []string{"*"}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
