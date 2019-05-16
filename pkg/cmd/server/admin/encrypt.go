package admin

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	configapi "github.com/openshift/origin/pkg/cmd/server/apis/config"
	pemutil "github.com/openshift/origin/pkg/cmd/util/pem"
	"github.com/openshift/origin/pkg/cmd/util/term"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
	"os"
	"unicode"
	"unicode/utf8"
)

const EncryptCommandName = "encrypt"

type EncryptOptions struct {
	CleartextFile   string
	CleartextData   []byte
	CleartextReader io.Reader
	EncryptedFile   string
	EncryptedWriter io.Writer
	KeyFile         string
	GenKeyFile      string
	PromptWriter    io.Writer
}

var encryptExample = templates.Examples(`
	# Encrypt the content of secret.txt with a generated key:
	%[1]s --genkey=secret.key --in=secret.txt --out=secret.encrypted

	# Encrypt the content of secret2.txt with an existing key:
	%[1]s --key=secret.key < secret2.txt > secret2.encrypted`)

func NewEncryptOptions(streams genericclioptions.IOStreams) *EncryptOptions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &EncryptOptions{CleartextReader: streams.In, EncryptedWriter: streams.Out, PromptWriter: streams.ErrOut}
}
func NewCommandEncrypt(commandName string, fullName string, streams genericclioptions.IOStreams) *cobra.Command {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	o := NewEncryptOptions(streams)
	cmd := &cobra.Command{Use: commandName, Short: "Encrypt data with AES-256-CBC encryption", Example: fmt.Sprintf(encryptExample, fullName), Run: func(cmd *cobra.Command, args []string) {
		kcmdutil.CheckErr(o.Validate(args))
		kcmdutil.CheckErr(o.Encrypt())
	}}
	cmd.Flags().StringVar(&o.CleartextFile, "in", o.CleartextFile, "File containing the data to encrypt. Read from stdin if omitted.")
	cmd.Flags().StringVar(&o.EncryptedFile, "out", o.EncryptedFile, "File to write the encrypted data to. Written to stdout if omitted.")
	cmd.Flags().StringVar(&o.KeyFile, "key", o.KeyFile, "File containing the encrypting key from in the format written by --genkey.")
	cmd.Flags().StringVar(&o.GenKeyFile, "genkey", o.GenKeyFile, "File to write a randomly generated key to.")
	cmd.MarkFlagFilename("in")
	cmd.MarkFlagFilename("out")
	cmd.MarkFlagFilename("key")
	cmd.MarkFlagFilename("genkey")
	return cmd
}
func (o *EncryptOptions) Validate(args []string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(args) != 0 {
		return errors.New("no arguments are supported")
	}
	if len(o.CleartextFile) == 0 && len(o.CleartextData) == 0 && o.CleartextReader == nil {
		return errors.New("an input file, data, or reader is required")
	}
	if len(o.CleartextFile) > 0 && len(o.CleartextData) > 0 {
		return errors.New("cannot specify both an input file and data")
	}
	if len(o.EncryptedFile) == 0 && o.EncryptedWriter == nil {
		return errors.New("an output file or writer is required")
	}
	if len(o.GenKeyFile) > 0 && len(o.KeyFile) > 0 {
		return errors.New("either --genkey or --key may be specified, not both")
	}
	if len(o.GenKeyFile) == 0 && len(o.KeyFile) == 0 {
		return errors.New("--genkey or --key is required")
	}
	return nil
}
func (o *EncryptOptions) Encrypt() error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var data []byte
	var warnWhitespace = true
	switch {
	case len(o.CleartextFile) > 0:
		if d, err := ioutil.ReadFile(o.CleartextFile); err != nil {
			return err
		} else {
			data = d
		}
	case len(o.CleartextData) > 0:
		warnWhitespace = false
		data = o.CleartextData
	case o.CleartextReader != nil && term.IsTerminalReader(o.CleartextReader) && o.PromptWriter != nil:
		data = []byte(term.PromptForString(o.CleartextReader, o.PromptWriter, "Data to encrypt: "))
	case o.CleartextReader != nil:
		if d, err := ioutil.ReadAll(o.CleartextReader); err != nil {
			return err
		} else {
			data = d
		}
	}
	if warnWhitespace && (o.PromptWriter != nil) && (len(data) > 0) {
		r1, _ := utf8.DecodeRune(data)
		r2, _ := utf8.DecodeLastRune(data)
		if unicode.IsSpace(r1) || unicode.IsSpace(r2) {
			fmt.Fprintln(o.PromptWriter, "Warning: Data includes leading or trailing whitespace, which will be included in the encrypted value")
		}
	}
	var key []byte
	switch {
	case len(o.KeyFile) > 0:
		if block, ok, err := pemutil.BlockFromFile(o.KeyFile, configapi.StringSourceKeyBlockType); err != nil {
			return err
		} else if !ok {
			return fmt.Errorf("%s does not contain a valid PEM block of type %q", o.KeyFile, configapi.StringSourceKeyBlockType)
		} else if len(block.Bytes) == 0 {
			return fmt.Errorf("%s does not contain a key", o.KeyFile)
		} else {
			key = block.Bytes
		}
	case len(o.GenKeyFile) > 0:
		key = make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			return err
		}
	}
	if len(key) == 0 {
		return errors.New("--genkey or --key is required")
	}
	dataBlock, err := x509.EncryptPEMBlock(rand.Reader, configapi.StringSourceEncryptedBlockType, data, key, x509.PEMCipherAES256)
	if err != nil {
		return err
	}
	if len(o.EncryptedFile) > 0 {
		if err := pemutil.BlockToFile(o.EncryptedFile, dataBlock, os.FileMode(0644)); err != nil {
			return err
		}
	} else if o.EncryptedWriter != nil {
		encryptedBytes, err := pemutil.BlockToBytes(dataBlock)
		if err != nil {
			return err
		}
		n, err := o.EncryptedWriter.Write(encryptedBytes)
		if err != nil {
			return err
		}
		if n != len(encryptedBytes) {
			return fmt.Errorf("could not completely write encrypted data")
		}
	}
	if len(o.GenKeyFile) > 0 {
		keyBlock := &pem.Block{Bytes: key, Type: configapi.StringSourceKeyBlockType}
		if err := pemutil.BlockToFile(o.GenKeyFile, keyBlock, os.FileMode(0600)); err != nil {
			return err
		}
	}
	return nil
}
