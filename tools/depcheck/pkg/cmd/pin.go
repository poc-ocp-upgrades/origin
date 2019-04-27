package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/openshift/origin/tools/depcheck/glide"
)

var pinImportsExample = `# output the contents of a glide.yaml file with new
# imports pinned to a SHA from a given glide.lock file
%[1]s pin --glide.lock=/path/to/glide.lock --glide.yaml=/path/to/glide.yaml
`

type PinImportsOpts struct {
	LockFile			*glide.LockFile
	YamlFile			*glide.YamlFile
	ExistingGlideYamlContent	[]byte
	Out				io.Writer
	ErrOut				io.Writer
}
type PinImportsFlags struct {
	LockFileName	string
	YamlFileName	string
}

func (o *PinImportsFlags) ToOptions(out, errout io.Writer) (*PinImportsOpts, error) {
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
	if len(o.LockFileName) == 0 || len(o.YamlFileName) == 0 {
		return nil, fmt.Errorf("both a --glide.lock path and a --glide.yaml path must be specified")
	}
	yamlFileBytes, err := ioutil.ReadFile(o.YamlFileName)
	if err != nil {
		return nil, err
	}
	lockFileBytes, err := ioutil.ReadFile(o.LockFileName)
	if err != nil {
		return nil, err
	}
	if len(yamlFileBytes) == 0 {
		return nil, fmt.Errorf("--glide.yaml path contained an empty file")
	}
	if len(lockFileBytes) == 0 {
		return nil, fmt.Errorf("--glide.lock path contained an empty file")
	}
	lockfile := &glide.LockFile{}
	yamlfile := &glide.YamlFile{}
	if err := lockfile.Decode(lockFileBytes); err != nil {
		return nil, err
	}
	if err := yamlfile.Decode(yamlFileBytes); err != nil {
		return nil, err
	}
	return &PinImportsOpts{LockFile: lockfile, YamlFile: yamlfile, ExistingGlideYamlContent: yamlFileBytes, Out: out, ErrOut: errout}, nil
}
func NewCmdPinImports(parent string, out, errout io.Writer) *cobra.Command {
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
	flags := &PinImportsFlags{}
	cmd := &cobra.Command{Use: "pin --glide.lock=FOO.lock --glide.yaml=FOO.yaml", Short: "Outputs a glide.yaml with all unspecified repos pinned to use SHAs specified in a glide.lock file.", Long: `Outputs a glide.yaml with all unspecified repos pinned to use SHAs specified in a glide.lock file.
Any existing dependencies on a glide.yaml file containing a "repo" field are ignored.`, Example: fmt.Sprintf(pinImportsExample, parent), RunE: func(c *cobra.Command, args []string) error {
		opts, err := flags.ToOptions(out, errout)
		if err != nil {
			return err
		}
		if err := opts.Validate(); err != nil {
			return err
		}
		return opts.Run()
	}}
	cmd.Flags().StringVar(&flags.LockFileName, "glide.lock", "glide.lock", "Path to glide.lock file used to indicate which dependencies should be updated or pinned to a specific revision in the glide.yaml file")
	cmd.Flags().StringVar(&flags.YamlFileName, "glide.yaml", "glide.yaml", "Path to glide.yaml file used to pin dependencies from the given --glide.lock file")
	return cmd
}
func (o *PinImportsOpts) Validate() error {
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
	if o.YamlFile == nil || o.LockFile == nil {
		return fmt.Errorf("both a glide.yaml and a glide.lock file are required")
	}
	if o.Out == nil || o.ErrOut == nil {
		return fmt.Errorf("no output or error output buffer given")
	}
	return nil
}
func (o *PinImportsOpts) Run() error {
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
	missing, warnings, err := glide.MissingImports(o.LockFile, o.YamlFile)
	if err != nil {
		return fmt.Errorf("error: unable to calculate missing imports: %v\n", err)
	}
	for _, warning := range warnings {
		fmt.Fprintf(o.ErrOut, "%s\n", warning)
	}
	missingBytes, err := missing.Encode()
	if err != nil {
		return err
	}
	o.Out.Write(o.ExistingGlideYamlContent)
	fmt.Fprintln(o.Out, "\n\n# generated by tools/depcheck")
	o.Out.Write(missingBytes)
	return nil
}
