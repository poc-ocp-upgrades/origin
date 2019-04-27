package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"github.com/joho/godotenv"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation"
)

type Environment map[string]string

func ParseEnvironmentAllowEmpty(vals ...string) Environment {
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
	env := make(Environment)
	for _, s := range vals {
		if i := strings.Index(s, "="); i == -1 {
			env[s] = ""
		} else {
			env[s[:i]] = s[i+1:]
		}
	}
	return env
}
func ParseEnvironment(vals ...string) (Environment, []string, []error) {
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
	errs := []error{}
	duplicates := []string{}
	env := make(Environment)
	for _, s := range vals {
		p := strings.SplitN(s, "=", 2)
		if err := validation.IsEnvVarName(p[0]); len(err) != 0 || len(p) != 2 {
			errs = append(errs, fmt.Errorf("invalid parameter assignment in %q, %v", s, err))
			continue
		}
		key, val := p[0], p[1]
		if _, exists := env[key]; exists {
			duplicates = append(duplicates, key)
			continue
		}
		env[key] = val
	}
	return env, duplicates, errs
}
func NewEnvironment(envs ...map[string]string) Environment {
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
	if len(envs) == 1 {
		return envs[0]
	}
	out := make(Environment)
	out.Add(envs...)
	return out
}
func (e Environment) Add(envs ...map[string]string) {
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
	for _, env := range envs {
		for k, v := range env {
			e[k] = v
		}
	}
}
func (e Environment) AddIfNotPresent(more Environment) []string {
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
	duplicates := []string{}
	for k, v := range more {
		if _, exists := e[k]; exists {
			duplicates = append(duplicates, k)
		} else {
			e[k] = v
		}
	}
	return duplicates
}
func (e Environment) List() []corev1.EnvVar {
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
	env := []corev1.EnvVar{}
	for k, v := range e {
		env = append(env, corev1.EnvVar{Name: k, Value: v})
	}
	sort.Sort(sortedEnvVar(env))
	return env
}

type sortedEnvVar []corev1.EnvVar

func (m sortedEnvVar) Len() int {
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
	return len(m)
}
func (m sortedEnvVar) Swap(i, j int) {
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
	m[i], m[j] = m[j], m[i]
}
func (m sortedEnvVar) Less(i, j int) bool {
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
	return m[i].Name < m[j].Name
}
func JoinEnvironment(a, b []corev1.EnvVar) (out []corev1.EnvVar) {
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
	out = a
	for i := range b {
		exists := false
		for j := range a {
			if a[j].Name == b[i].Name {
				exists = true
				break
			}
		}
		if exists {
			continue
		}
		out = append(out, b[i])
	}
	return out
}
func LoadEnvironmentFile(filename string, stdin io.Reader) (Environment, error) {
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
	errorFilename := filename
	if filename == "-" && stdin != nil {
		temp, err := ioutil.TempFile("", "origin-env-stdin")
		if err != nil {
			return nil, fmt.Errorf("Cannot create temporary file: %s", err)
		}
		filename = temp.Name()
		errorFilename = "stdin"
		defer os.Remove(filename)
		if _, err = io.Copy(temp, stdin); err != nil {
			return nil, fmt.Errorf("Cannot write to temporary file %q: %s", filename, err)
		}
		temp.Close()
	}
	if info, err := os.Stat(filename); err == nil && info.IsDir() {
		return nil, fmt.Errorf("Cannot read variables from %q: is a directory", filename)
	} else if err != nil {
		return nil, fmt.Errorf("Cannot stat %q: %s", filename, err)
	}
	env, err := godotenv.Read(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot read variables from file %q: %s", errorFilename, err)
	}
	for k, v := range env {
		if err := validation.IsEnvVarName(k); len(err) != 0 {
			return nil, fmt.Errorf("invalid parameter assignment in %s=%s", k, v)
		}
	}
	return env, nil
}
func ParseAndCombineEnvironment(envs []string, filenames []string, stdin io.Reader, dupfn func(string, string) error) (Environment, error) {
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
	vars, duplicates, errs := ParseEnvironment(envs...)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	for _, s := range duplicates {
		if err := dupfn(s, ""); err != nil {
			return nil, err
		}
	}
	for _, fname := range filenames {
		fileVars, err := LoadEnvironmentFile(fname, stdin)
		if err != nil {
			return nil, err
		}
		duplicates = vars.AddIfNotPresent(fileVars)
		for _, s := range duplicates {
			if err := dupfn(s, fname); err != nil {
				return nil, err
			}
		}
	}
	return vars, nil
}
