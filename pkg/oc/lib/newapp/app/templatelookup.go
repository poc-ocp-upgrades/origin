package app

import (
	"encoding/json"
	"fmt"
	"strings"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	templatev1 "github.com/openshift/api/template/v1"
	templatev1typedclient "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
)

type TemplateSearcher struct {
	Client			templatev1typedclient.TemplatesGetter
	Namespaces		[]string
	StopOnExactMatch	bool
}

func (r TemplateSearcher) Type() string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "templates loaded in accessible projects"
}
func (r TemplateSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches := ComponentMatches{}
	var errs []error
	for _, term := range terms {
		ref, err := parseTemplateReference(term)
		if err != nil {
			klog.V(2).Infof("template references must be of the form [<namespace>/]<name>, term %q did not qualify", term)
			continue
		}
		if term == "__template_fail" {
			errs = append(errs, fmt.Errorf("unable to find the specified template: %s", term))
			continue
		}
		namespaces := r.Namespaces
		if ref.hasNamespace() {
			namespaces = []string{ref.Namespace}
		}
		checkedNamespaces := sets.NewString()
		for _, namespace := range namespaces {
			if checkedNamespaces.Has(namespace) {
				continue
			}
			checkedNamespaces.Insert(namespace)
			templates, err := r.Client.Templates(namespace).List(metav1.ListOptions{})
			if err != nil {
				if errors.IsNotFound(err) || errors.IsForbidden(err) {
					continue
				}
				errs = append(errs, err)
				continue
			}
			exact := false
			for i := range templates.Items {
				template := &templates.Items[i]
				klog.V(4).Infof("checking namespace %s for template %s", namespace, ref.Name)
				if score, scored := templateScorer(*template, ref.Name); scored {
					if score == 0.0 {
						exact = true
					}
					klog.V(4).Infof("Adding template %q in project %q with score %f", template.Name, template.Namespace, score)
					fullName := fmt.Sprintf("%s/%s", template.Namespace, template.Name)
					matches = append(matches, &ComponentMatch{Value: term, Argument: fmt.Sprintf("--template=%q", fullName), Name: fullName, Description: fmt.Sprintf("Template %q in project %q", template.Name, template.Namespace), Score: score, Template: template})
				}
			}
			if exact && precise {
				break
			}
		}
	}
	return matches, errs
}
func IsPossibleTemplateFile(value string) (bool, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isFile(value)
}

type TemplateFileSearcher struct {
	Builder		*resource.Builder
	Namespace	string
}

func (r *TemplateFileSearcher) Type() string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "template files"
}
func (r *TemplateFileSearcher) Search(precise bool, terms ...string) (ComponentMatches, []error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches := ComponentMatches{}
	var errs []error
	for _, term := range terms {
		if term == "__templatefile_fail" {
			errs = append(errs, fmt.Errorf("unable to find the specified template file: %s", term))
			continue
		}
		var isSingleItemImplied bool
		obj, err := r.Builder.WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).NamespaceParam(r.Namespace).RequireNamespace().FilenameParam(false, &resource.FilenameOptions{Recursive: false, Filenames: terms}).Do().IntoSingleItemImplied(&isSingleItemImplied).Object()
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "does not exist") && strings.Contains(err.Error(), "the path"):
				continue
			case strings.Contains(err.Error(), "not a directory") && strings.Contains(err.Error(), "the path"):
				continue
			default:
				if syntaxErr, ok := err.(*json.SyntaxError); ok {
					err = fmt.Errorf("at offset %d: %v", syntaxErr.Offset, err)
				}
				errs = append(errs, fmt.Errorf("unable to load template file %q: %v", term, err))
				continue
			}
		}
		if list, isList := obj.(*corev1.List); isList && !isSingleItemImplied {
			if len(list.Items) == 1 {
				obj = list.Items[0].Object
				isSingleItemImplied = true
			}
		}
		if !isSingleItemImplied {
			errs = append(errs, fmt.Errorf("there is more than one object in %q", term))
			continue
		}
		template, ok := obj.(*templatev1.Template)
		if !ok {
			errs = append(errs, fmt.Errorf("object in %q is not a template", term))
			continue
		}
		matches = append(matches, &ComponentMatch{Value: term, Argument: fmt.Sprintf("--file=%q", template.Name), Name: template.Name, Description: fmt.Sprintf("Template file %s", term), Score: 0, Template: template})
	}
	return matches, errs
}

type templateReference struct {
	Namespace	string
	Name		string
}

func parseTemplateReference(s string) (templateReference, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ref templateReference
	parts := strings.Split(s, "/")
	switch len(parts) {
	case 2:
		ref.Namespace = parts[0]
		ref.Name = parts[1]
		break
	case 1:
		ref.Name = parts[0]
		break
	default:
		return ref, fmt.Errorf("the template reference must be either the template name or namespace and template name separated by slashes")
	}
	return ref, nil
}
func (r templateReference) hasNamespace() bool {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(r.Namespace) > 0
}
func (r templateReference) String() string {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.hasNamespace() {
		return fmt.Sprintf("%s/%s", r.Namespace, r.Name)
	}
	return r.Name
}
