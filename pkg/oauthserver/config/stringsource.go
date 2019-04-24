package config

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"k8s.io/apimachinery/pkg/util/sets"
	configv1 "github.com/openshift/api/config/v1"
	osinv1 "github.com/openshift/api/osin/v1"
)

const (
	StringSourceEncryptedBlockType	= "ENCRYPTED STRING"
	StringSourceKeyBlockType	= "ENCRYPTING KEY"
)

var ValidGrantHandlerTypes = sets.NewString(string(osinv1.GrantHandlerAuto), string(osinv1.GrantHandlerPrompt), string(osinv1.GrantHandlerDeny))

func ResolveStringValue(s configv1.StringSource) (string, error) {
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
	secretBlock, ok := blockFromBytes([]byte(value), StringSourceEncryptedBlockType)
	if !ok {
		return "", fmt.Errorf("no valid PEM block of type %q found in data", StringSourceEncryptedBlockType)
	}
	keyBlock, ok := blockFromBytes(keyData, StringSourceKeyBlockType)
	if !ok {
		return "", fmt.Errorf("no valid PEM block of type %q found in key", StringSourceKeyBlockType)
	}
	data, err := x509.DecryptPEMBlock(secretBlock, keyBlock.Bytes)
	return string(data), err
}
func blockFromBytes(data []byte, blockType string) (*pem.Block, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		block, remaining := pem.Decode(data)
		if block == nil {
			return nil, false
		}
		if block.Type == blockType {
			return block, true
		}
		data = remaining
	}
}
