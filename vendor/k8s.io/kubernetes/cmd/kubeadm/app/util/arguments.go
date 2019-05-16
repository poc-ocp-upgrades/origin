package util

import (
	"fmt"
	goformat "fmt"
	"github.com/pkg/errors"
	goos "os"
	godefaultruntime "runtime"
	"sort"
	"strings"
	gotime "time"
)

func BuildArgumentListFromMap(baseArguments map[string]string, overrideArguments map[string]string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var command []string
	var keys []string
	for k := range overrideArguments {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := overrideArguments[k]
		command = append(command, fmt.Sprintf("--%s=%s", k, v))
	}
	keys = []string{}
	for k := range baseArguments {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := baseArguments[k]
		if _, overrideExists := overrideArguments[k]; !overrideExists {
			command = append(command, fmt.Sprintf("--%s=%s", k, v))
		}
	}
	return command
}
func ParseArgumentListToMap(arguments []string) map[string]string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	resultingMap := map[string]string{}
	for i, arg := range arguments {
		key, val, err := parseArgument(arg)
		if err != nil {
			if i != 0 {
				fmt.Printf("[kubeadm] WARNING: The component argument %q could not be parsed correctly. The argument must be of the form %q. Skipping...", arg, "--")
			}
			continue
		}
		resultingMap[key] = val
	}
	return resultingMap
}
func ReplaceArgument(command []string, argMutateFunc func(map[string]string) map[string]string) []string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	argMap := ParseArgumentListToMap(command)
	var newCommand []string
	if len(command) > 0 && !strings.HasPrefix(command[0], "--") {
		newCommand = append(newCommand, command[0])
	}
	newArgMap := argMutateFunc(argMap)
	newCommand = append(newCommand, BuildArgumentListFromMap(newArgMap, map[string]string{})...)
	return newCommand
}
func parseArgument(arg string) (string, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if !strings.HasPrefix(arg, "--") {
		return "", "", errors.New("the argument should start with '--'")
	}
	if !strings.Contains(arg, "=") {
		return "", "", errors.New("the argument should have a '=' between the flag and the value")
	}
	arg = strings.TrimPrefix(arg, "--")
	keyvalSlice := strings.SplitN(arg, "=", 2)
	if len(keyvalSlice) != 2 {
		return "", "", errors.New("the argument must have both a key and a value")
	}
	if len(keyvalSlice[0]) == 0 {
		return "", "", errors.New("the argument must have a key")
	}
	return keyvalSlice[0], keyvalSlice[1], nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
