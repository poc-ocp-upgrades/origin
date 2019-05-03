package config

import (
	"crypto/x509"
	"fmt"
	pemutil "github.com/openshift/origin/pkg/cmd/util/pem"
	"io/ioutil"
	"os"
)

func GetStringSourceFileReferences(s *StringSource) []*string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == nil {
		return nil
	}
	return []*string{&s.File, &s.KeyFile}
}
func ResolveStringValue(s StringSource) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var value string
	switch {
	case len(s.Value) > 0:
		value = s.Value
	case len(s.Env) > 0:
		value = os.Getenv(s.Env)
	case len(s.File) > 0:
		data, err := ioutil.ReadFile(s.File)
		if err != nil {
			return "", err
		}
		value = string(data)
	default:
		value = ""
	}
	if len(s.KeyFile) == 0 {
		return value, nil
	}
	keyData, err := ioutil.ReadFile(s.KeyFile)
	if err != nil {
		return "", err
	}
	secretBlock, ok := pemutil.BlockFromBytes([]byte(value), StringSourceEncryptedBlockType)
	if !ok {
		return "", fmt.Errorf("no valid PEM block of type %q found in data", StringSourceEncryptedBlockType)
	}
	keyBlock, ok := pemutil.BlockFromBytes(keyData, StringSourceKeyBlockType)
	if !ok {
		return "", fmt.Errorf("no valid PEM block of type %q found in key", StringSourceKeyBlockType)
	}
	data, err := x509.DecryptPEMBlock(secretBlock, keyBlock.Bytes)
	return string(data), err
}
