package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/openshift/origin/pkg/api/apihelpers"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/klog"
)

var invalidNameCharactersRegexp = regexp.MustCompile("[^-a-z0-9]")

type UniqueNameGenerator interface {
	Generate(NameSuggester) (string, error)
}

func NewUniqueNameGenerator(name string) UniqueNameGenerator {
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
	return &uniqueNameGenerator{name, map[string]int{}}
}

type uniqueNameGenerator struct {
	originalName	string
	names		map[string]int
}

func (ung *uniqueNameGenerator) Generate(suggester NameSuggester) (string, error) {
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
	name := ung.originalName
	if len(name) == 0 {
		var ok bool
		name, ok = suggester.SuggestName()
		if !ok {
			return "", ErrNameRequired
		}
	}
	return ung.ensureValidName(name)
}
func (ung *uniqueNameGenerator) ensureValidName(name string) (string, error) {
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
	names := ung.names
	if len(name) < 2 {
		return "", fmt.Errorf("invalid name: %s", name)
	}
	if !IsParameterizableValue(name) {
		name = strings.ToLower(name)
		name = invalidNameCharactersRegexp.ReplaceAllString(name, "")
		name = strings.TrimLeft(name, "-")
		if len(name) > kvalidation.DNS1123SubdomainMaxLength {
			klog.V(4).Infof("Trimming %s to maximum allowable length (%d)\n", name, kvalidation.DNS1123SubdomainMaxLength)
			name = name[:kvalidation.DNS1123SubdomainMaxLength]
		}
	}
	count, existing := names[name]
	if !existing {
		names[name] = 0
		return name, nil
	}
	count++
	names[name] = count
	newName := apihelpers.GetName(name, strconv.Itoa(count), kvalidation.DNS1123SubdomainMaxLength)
	return newName, nil
}
