package admin

import (
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"github.com/openshift/origin/pkg/cmd/util/term"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	pemutil "github.com/openshift/origin/pkg/cmd/util/pem"
)

const DecryptCommandName = "decrypt"

type DecryptOptions struct {
	EncryptedFile	string
	EncryptedData	[]byte
	EncryptedReader	io.Reader
	DecryptedFile	string
	DecryptedWriter	io.Writer
	KeyFile		string
}

var decryptExample = templates.Examples(`
	# Decrypt an encrypted file to a cleartext file:
	%[1]s --key=secret.key --in=secret.encrypted --out=secret.decrypted

	# Decrypt from stdin to stdout:
	%[1]s --key=secret.key < secret2.encrypted > secret2.decrypted`)

func NewDecryptOptions(streams genericclioptions.IOStreams) *DecryptOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DecryptOptions{EncryptedReader: streams.In, DecryptedWriter: streams.Out}
}
func NewCommandDecrypt(commandName string, fullName, encryptFullName string, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := NewDecryptOptions(streams)
	cmd := &cobra.Command{Use: commandName, Short: fmt.Sprintf("Decrypt data encrypted with %q", encryptFullName), Example: fmt.Sprintf(decryptExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Validate(args))
		kcmdutil.CheckErr(o.Decrypt())
	}}
	cmd.Flags().StringVar(&o.EncryptedFile, "in", o.EncryptedFile, fmt.Sprintf("File containing encrypted data, in the format written by %q.", encryptFullName))
	cmd.Flags().StringVar(&o.DecryptedFile, "out", o.DecryptedFile, "File to write the decrypted data to. Written to stdout if omitted.")
	cmd.Flags().StringVar(&o.KeyFile, "key", o.KeyFile, fmt.Sprintf("The file to read the decrypting key from. Must be a PEM file in the format written by %q.", encryptFullName))
	cmd.MarkFlagFilename("in")
	cmd.MarkFlagFilename("out")
	cmd.MarkFlagFilename("key")
	return cmd
}
func (o *DecryptOptions) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		return errors.New("no arguments are supported")
	}
	if len(o.EncryptedFile) == 0 && len(o.EncryptedData) == 0 && (o.EncryptedReader == nil || term.IsTerminalReader(o.EncryptedReader)) {
		return errors.New("no input data specified")
	}
	if len(o.EncryptedFile) > 0 && len(o.EncryptedData) > 0 {
		return errors.New("cannot specify both an input file and data")
	}
	if len(o.KeyFile) == 0 {
		return errors.New("no key specified")
	}
	return nil
}
func (o *DecryptOptions) Decrypt() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var data []byte
	switch {
	case len(o.EncryptedFile) > 0:
		if d, err := ioutil.ReadFile(o.EncryptedFile); err != nil {
			return err
		} else {
			data = d
		}
	case len(o.EncryptedData) > 0:
		data = o.EncryptedData
	case o.EncryptedReader != nil && !term.IsTerminalReader(o.EncryptedReader):
		if d, err := ioutil.ReadAll(o.EncryptedReader); err != nil {
			return err
		} else {
			data = d
		}
	}
	if len(data) == 0 {
		return fmt.Errorf("no input data specified")
	}
	dataBlock, ok := pemutil.BlockFromBytes(data, configapi.StringSourceEncryptedBlockType)
	if !ok {
		return fmt.Errorf("input does not contain a valid PEM block of type %q", configapi.StringSourceEncryptedBlockType)
	}
	keyBlock, ok, err := pemutil.BlockFromFile(o.KeyFile, configapi.StringSourceKeyBlockType)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("%s does not contain a valid PEM block of type %q", o.KeyFile, configapi.StringSourceKeyBlockType)
	}
	if len(keyBlock.Bytes) == 0 {
		return fmt.Errorf("%s does not contain a key", o.KeyFile)
	}
	password := keyBlock.Bytes
	plaintext, err := x509.DecryptPEMBlock(dataBlock, password)
	if err != nil {
		return err
	}
	switch {
	case len(o.DecryptedFile) > 0:
		if err := ioutil.WriteFile(o.DecryptedFile, plaintext, os.FileMode(0600)); err != nil {
			return err
		}
	case o.DecryptedWriter != nil:
		fmt.Fprint(o.DecryptedWriter, string(plaintext))
		if term.IsTerminalWriter(o.DecryptedWriter) {
			fmt.Fprintln(o.DecryptedWriter)
		}
	}
	return nil
}
