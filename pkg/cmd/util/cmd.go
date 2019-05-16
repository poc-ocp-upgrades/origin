package util

import (
	"errors"
	"fmt"
	goformat "fmt"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	goos "os"
	"path/filepath"
	"regexp"
	godefaultruntime "runtime"
	"strings"
	gotime "time"
)

var commaSepVarsPattern = regexp.MustCompile(".*=.*,.*=.*")

func ReplaceCommandName(from, to string, c *cobra.Command) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	c.Example = strings.Replace(c.Example, from, to, -1)
	c.Long = strings.Replace(c.Long, from, to, -1)
	for _, sub := range c.Commands() {
		ReplaceCommandName(from, to, sub)
	}
	return c
}
func GetDisplayFilename(filename string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if absName, err := filepath.Abs(filename); err == nil {
		return absName
	}
	return filename
}
func ResolveResource(defaultResource schema.GroupResource, resourceString string, mapper meta.RESTMapper) (schema.GroupResource, string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if mapper == nil {
		return schema.GroupResource{}, "", errors.New("mapper cannot be nil")
	}
	var name string
	parts := strings.Split(resourceString, "/")
	switch len(parts) {
	case 1:
		name = parts[0]
	case 2:
		name = parts[1]
		groupResource := schema.ParseGroupResource(parts[0])
		groupResource.Resource = strings.ToLower(groupResource.Resource)
		gvr, err := mapper.ResourceFor(groupResource.WithVersion(""))
		if err != nil {
			return schema.GroupResource{}, "", err
		}
		return gvr.GroupResource(), name, nil
	default:
		return schema.GroupResource{}, "", fmt.Errorf("invalid resource format: %s", resourceString)
	}
	return defaultResource, name, nil
}
func WarnAboutCommaSeparation(errout io.Writer, values []string, flag string) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errout == nil {
		return
	}
	for _, value := range values {
		if commaSepVarsPattern.MatchString(value) {
			fmt.Fprintf(errout, "warning: %s no longer accepts comma-separated lists of values. %q will be treated as a single key-value pair.\n", flag, value)
		}
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
