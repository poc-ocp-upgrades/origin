package secrets

import (
	"testing"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestValidateSSHAuth(t *testing.T) {
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
	tests := []struct {
		testName	string
		args		[]string
		options		func(genericclioptions.IOStreams) *CreateSSHAuthSecretOptions
		expErr		bool
	}{{testName: "validArgs", args: []string{"testSecret"}, options: func(streams genericclioptions.IOStreams) *CreateSSHAuthSecretOptions {
		o := NewCreateSSHAuthSecretOptions(streams)
		o.SecretName = "testSecret"
		o.PrivateKeyPath = "./bsFixtures/valid/ssh-privatekey"
		return o
	}, expErr: false}, {testName: "validArgsWithCertificate", args: []string{"testSecret"}, options: func(streams genericclioptions.IOStreams) *CreateSSHAuthSecretOptions {
		o := NewCreateSSHAuthSecretOptions(streams)
		o.SecretName = "testSecret"
		o.PrivateKeyPath = "./bsFixtures/valid/ssh-privatekey"
		o.CertificatePath = "./bsFixtures/valid/ca.crt"
		return o
	}, expErr: false}, {testName: "noName", args: []string{}, options: func(streams genericclioptions.IOStreams) *CreateSSHAuthSecretOptions {
		o := NewCreateSSHAuthSecretOptions(streams)
		o.SecretName = "testSecret"
		o.PrivateKeyPath = "./bsFixtures/valid/ssh-privatekey"
		o.CertificatePath = "./bsFixtures/valid/ca.crt"
		return o
	}, expErr: true}, {testName: "noParams", args: []string{"testSecret"}, options: func(streams genericclioptions.IOStreams) *CreateSSHAuthSecretOptions {
		o := NewCreateSSHAuthSecretOptions(streams)
		o.SecretName = "testSecret"
		return o
	}, expErr: true}}
	for _, test := range tests {
		options := test.options(genericclioptions.NewTestIOStreamsDiscard())
		err := options.Validate(test.args)
		if test.expErr {
			if err == nil {
				t.Errorf("%s: unexpected error: %v", test.testName, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: unexpected error: %v", test.testName, err)
		}
		secret, err := options.NewSSHAuthSecret()
		if err != nil {
			t.Errorf("%s: unexpected error: %v", test.testName, err)
		}
		if secret.Type != corev1.SecretTypeSSHAuth {
			t.Errorf("%s: unexpected secret.Type: %v", test.testName, secret.Type)
		}
	}
}
