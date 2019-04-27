package errors

import (
	"fmt"
	"runtime"
)

const (
	KubeConfigFileSolutionWindows	= `
Make sure that the value of the --config flag passed contains a valid path:
   --config=c:\path\to\valid\file
`
	KubeConfigFileSolutionUnix	= `
Make sure that the value of the --config flag passed contains a valid path:
   --config=/path/to/valid/file
`
	KubeConfigSolutionUnix	= `
You can unset the KUBECONFIG variable to use the default location for it:
   unset KUBECONFIG

Or you can set its value to a file that can be written to:
   export KUBECONFIG=/path/to/file
`
	KubeConfigSolutionWindows	= `
You can clear the KUBECONFIG variable to use the default location for it:
   set KUBECONFIG=

Or you can set its value to a file that can be written to:
   set KUBECONFIG=c:\path\to\file`
)

func ErrKubeConfigNotWriteable(file string, isExplicitFile bool, err error) error {
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
	return NewError("KUBECONFIG is set to a file that cannot be created or modified: %s", file).WithCause(err).WithSolution(kubeConfigSolution(isExplicitFile))
}
func kubeConfigSolution(isExplicitFile bool) string {
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
	switch runtime.GOOS {
	case "windows":
		if isExplicitFile {
			return KubeConfigFileSolutionWindows
		}
		return KubeConfigSolutionWindows
	default:
		if isExplicitFile {
			return KubeConfigFileSolutionUnix
		}
		return KubeConfigSolutionUnix
	}
}
func NoProjectsExistMessage(canRequestProjects bool, commandName string) string {
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
	if !canRequestProjects {
		return fmt.Sprintf("You don't have any projects. Contact your system administrator to request a project.\n")
	}
	return fmt.Sprintf(`You don't have any projects. You can try to create a new project, by running

    %s new-project <projectname>

`, commandName)
}
