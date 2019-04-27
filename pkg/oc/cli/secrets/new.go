package secrets

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions/printers"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

const NewSecretRecommendedCommandName = "new"

var (
	newLong	= templates.LongDesc(`
    Create a new secret based on a file or directory

    Key files can be specified using their file path, in which case a default name will be given to them, or optionally
    with a name and file path, in which case the given name will be used. Specifying a directory will create a secret
    using with all valid keys in that directory.`)
	newExample	= templates.Examples(`
    # Create a new secret named my-secret with a key named ssh-privatekey
    %[1]s my-secret ~/.ssh/ssh-privatekey

    # Create a new secret named my-secret with keys named ssh-privatekey and ssh-publickey instead of the names of the keys on disk
    %[1]s my-secret ssh-privatekey=~/.ssh/id_rsa ssh-publickey=~/.ssh/id_rsa.pub

    # Create a new secret named my-secret with keys for each file in the folder "bar"
    %[1]s my-secret path/to/bar

    # Create a new .dockercfg secret named my-secret
    %[1]s my-secret path/to/.dockercfg

    # Create a new .docker/config.json secret named my-secret
    %[1]s my-secret .dockerconfigjson=path/to/.docker/config.json`)
)

type CreateSecretOptions struct {
	PrintFlags		*genericclioptions.PrintFlags
	Printer			printers.ResourcePrinter
	Name			string
	SecretTypeName		string
	Sources			[]string
	SecretsInterface	corev1client.SecretInterface
	Quiet			bool
	AllowUnknownTypes	bool
	genericclioptions.IOStreams
}

func NewCreateSecretOptions(streams genericclioptions.IOStreams) *CreateSecretOptions {
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
	return &CreateSecretOptions{PrintFlags: genericclioptions.NewPrintFlags("created").WithTypeSetter(scheme.Scheme), IOStreams: streams}
}
func NewCmdCreateSecret(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
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
	o := NewCreateSecretOptions(streams)
	cmd := &cobra.Command{Use: fmt.Sprintf("%s NAME [KEY=]SOURCE ...", name), Short: "Create a new secret based on a key file or on files within a directory", Long: newLong, Example: fmt.Sprintf(newExample, fullName), Deprecated: "use oc create secret", Hidden: true, Run: func(c *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Complete(args, f))
		kcmdutil.CheckErr(o.Validate())
		kcmdutil.CheckErr(o.Run())
	}}
	cmd.Flags().BoolVarP(&o.Quiet, "quiet", "q", o.Quiet, "If true, suppress warnings")
	cmd.Flags().BoolVar(&o.AllowUnknownTypes, "confirm", o.AllowUnknownTypes, "If true, allow unknown secret types.")
	cmd.Flags().StringVar(&o.SecretTypeName, "type", "", "The type of secret")
	o.PrintFlags.AddFlags(cmd)
	return cmd
}
func (o *CreateSecretOptions) Complete(args []string, f kcmdutil.Factory) error {
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
	if len(args) > 0 {
		o.Name = args[0]
	}
	if len(args) > 1 {
		o.Sources = append(o.Sources, args[1:]...)
	}
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	kubeClient, err := corev1client.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	namespace, _, err := f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	o.SecretsInterface = kubeClient.Secrets(namespace)
	o.Printer, err = o.PrintFlags.ToPrinter()
	if err != nil {
		return err
	}
	return nil
}
func (o *CreateSecretOptions) Validate() error {
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
	if len(o.Name) == 0 {
		return errors.New("secret name is required")
	}
	if len(o.Sources) == 0 {
		return errors.New("at least one source file or directory must be specified")
	}
	if !o.AllowUnknownTypes {
		switch o.SecretTypeName {
		case string(corev1.SecretTypeOpaque), "":
		default:
			found := false
			for _, secretType := range KnownSecretTypes {
				if o.SecretTypeName == string(secretType.Type) {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("unknown secret type %q; use --confirm to use it anyway", o.SecretTypeName)
			}
		}
	}
	return nil
}
func (o *CreateSecretOptions) Run() error {
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
	secret, err := o.BundleSecret()
	if err != nil {
		return err
	}
	if o.PrintFlags.OutputFormat != nil && len(*o.PrintFlags.OutputFormat) == 0 {
		persistedSecret, err := o.SecretsInterface.Create(secret)
		if err != nil {
			return err
		}
		return o.Printer.PrintObj(persistedSecret, o.Out)
	}
	return o.Printer.PrintObj(secret, o.Out)
}
func (o *CreateSecretOptions) BundleSecret() (*corev1.Secret, error) {
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
	secretData := make(map[string][]byte)
	for _, source := range o.Sources {
		keyName, filePath, err := parseSource(source)
		if err != nil {
			return nil, err
		}
		info, err := os.Stat(filePath)
		if err != nil {
			switch err := err.(type) {
			case *os.PathError:
				return nil, fmt.Errorf("error reading %s: %v", filePath, err.Err)
			default:
				return nil, fmt.Errorf("error reading %s: %v", filePath, err)
			}
		}
		if info.IsDir() {
			if strings.Contains(source, "=") {
				return nil, errors.New("Cannot give a key name for a directory path.")
			}
			fileList, err := ioutil.ReadDir(filePath)
			if err != nil {
				return nil, fmt.Errorf("error listing files in %s: %v", filePath, err)
			}
			for _, item := range fileList {
				itemPath := path.Join(filePath, item.Name())
				if !item.Mode().IsRegular() {
					if o.ErrOut != nil && o.Quiet != true {
						fmt.Fprintf(o.ErrOut, "Skipping resource %s\n", itemPath)
					}
				} else {
					keyName = item.Name()
					err = addKeyToSecret(keyName, itemPath, secretData)
					if err != nil {
						return nil, err
					}
				}
			}
		} else {
			err = addKeyToSecret(keyName, filePath, secretData)
			if err != nil {
				return nil, err
			}
		}
	}
	if len(secretData) == 0 {
		return nil, errors.New("No files selected")
	}
	secretType := corev1.SecretType(o.SecretTypeName)
	if len(o.SecretTypeName) == 0 {
		secretType = corev1.SecretTypeOpaque
		for _, knownSecretType := range KnownSecretTypes {
			if knownSecretType.Matches(secretData) {
				secretType = knownSecretType.Type
			}
		}
	}
	secret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: o.Name}, Type: secretType, Data: secretData}
	return secret, nil
}
func addKeyToSecret(keyName, filePath string, secretData map[string][]byte) error {
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
	if errors := kvalidation.IsConfigMapKey(keyName); len(errors) > 0 {
		return fmt.Errorf("%v is not a valid key name for a secret: %s", keyName, strings.Join(errors, ", "))
	}
	if _, entryExists := secretData[keyName]; entryExists {
		return fmt.Errorf("cannot add key %s from path %s, another key by that name already exists: %v.", keyName, filePath, secretData)
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	secretData[keyName] = data
	return nil
}
func parseSource(source string) (keyName, filePath string, err error) {
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
	numSeparators := strings.Count(source, "=")
	switch {
	case numSeparators == 0:
		return path.Base(source), source, nil
	case numSeparators == 1 && strings.HasPrefix(source, "="):
		return "", "", fmt.Errorf("key name for file path %v missing.", strings.TrimPrefix(source, "="))
	case numSeparators == 1 && strings.HasSuffix(source, "="):
		return "", "", fmt.Errorf("file path for key name %v missing.", strings.TrimSuffix(source, "="))
	case numSeparators > 1:
		return "", "", errors.New("Key names or file paths cannot contain '='.")
	default:
		components := strings.Split(source, "=")
		return components[0], components[1], nil
	}
}
