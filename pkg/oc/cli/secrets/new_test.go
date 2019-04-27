package secrets

import (
	"os"
	"testing"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kapi "k8s.io/kubernetes/pkg/apis/core"
)

func TestValidate(t *testing.T) {
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
		options		func(genericclioptions.IOStreams) *CreateSecretOptions
		expErr		bool
	}{{testName: "validArgs", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/www.google.com"}
		return o
	}}, {testName: "noName", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Sources = []string{"./bsFixtures/www.google.com"}
		return o
	}, expErr: true}, {testName: "noFilesPassed", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		return o
	}, expErr: true}}
	for _, test := range tests {
		options := test.options(genericclioptions.NewTestIOStreamsDiscard())
		err := options.Validate()
		if err != nil && !test.expErr {
			t.Errorf("%s: unexpected error: %v", test.testName, err)
		}
	}
}
func TestCreateSecret(t *testing.T) {
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
	os.Symlink(".", "./bsFixtures/dir/symbolic")
	defer os.Remove("./bsFixtures/dir/symbolic")
	tests := []struct {
		testName		string
		options			func(genericclioptions.IOStreams) *CreateSecretOptions
		expErr			bool
		errStreamContent	string
	}{{testName: "validSources", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/www.google.com", "./bsFixtures/dirNoSubdir"}
		return o
	}}, {testName: "allowsMixedCaseAndDash", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/invalid/invalid-DNS"}
		return o
	}}, {testName: "failsWithUnderscore", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/invalid/not\\valid"}
		return o
	}, expErr: true}, {testName: "leadingDotsAllowed", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/leadingdot/.dockercfg"}
		return o
	}}, {testName: "filesSameName", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/www.google.com", "./bsFixtures/multiple/www.google.com"}
		return o
	}, expErr: true}, {testName: "testQuietTrue", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/dir"}
		o.Quiet = true
		return o
	}}, {testName: "testQuietFalse", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/dir"}
		return o
	}, errStreamContent: "Skipping resource bsFixtures/dir/symbolic\n"}, {testName: "testNamedKeys", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".googlename=./bsFixtures/www.google.com"}
		return o
	}, expErr: false}, {testName: "testNamedDir", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".somename=./bsFixtures/dirNoSubdir"}
		return o
	}, expErr: true}, {testName: "testUnnamedDir", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"./bsFixtures/dirContainsMany"}
		return o
	}, expErr: false}, {testName: "testMalformedName", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".google=name=./bsFixtures/www.google.com"}
		return o
	}, expErr: true}, {testName: "testMissingName", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{"=./bsFixtures/www.google.com"}
		return o
	}, expErr: true}, {testName: "testMissingPath", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".somename="}
		return o
	}, expErr: true}, {testName: "testNamesAvoidCollision", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".googlename=./bsFixtures/www.google.com", ".othergooglename=./bsFixtures/multiple/www.google.com"}
		return o
	}, expErr: false}, {testName: "testNameCollision", options: func(streams genericclioptions.IOStreams) *CreateSecretOptions {
		o := NewCreateSecretOptions(streams)
		o.Name = "testSecret"
		o.Sources = []string{".googlename=./bsFixtures/www.google.com", ".googlename=./bsFixtures/multiple/www.google.com"}
		return o
	}, expErr: true}}
	for _, test := range tests {
		streams, _, _, errStream := genericclioptions.NewTestIOStreams()
		options := test.options(streams)
		err := options.Validate()
		if err != nil {
			t.Errorf("unexpected error")
		}
		_, err = options.BundleSecret()
		if err != nil && !test.expErr {
			t.Errorf("%s: unexpected error: %s", test.testName, err)
		}
		if err == nil && test.expErr {
			t.Errorf("%s: missing expected error", test.testName)
		}
		if string(errStream.Bytes()) != test.errStreamContent {
			t.Errorf("%s: expected %s, got %v", test.testName, test.errStreamContent, string(errStream.Bytes()))
		}
	}
}
func TestSecretTypeSpecified(t *testing.T) {
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
	options := CreateSecretOptions{Name: "any", SecretTypeName: string(kapi.SecretTypeDockercfg), Sources: []string{"./bsFixtures/www.google.com"}, IOStreams: genericclioptions.NewTestIOStreamsDiscard()}
	secret, err := options.BundleSecret()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if secret.Type != corev1.SecretTypeDockercfg {
		t.Errorf("expected %v, got %v", kapi.SecretTypeDockercfg, secret.Type)
	}
}
func TestSecretTypeDiscovered(t *testing.T) {
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
	options := CreateSecretOptions{Name: "any", Sources: []string{"./bsFixtures/leadingdot/.dockercfg"}, IOStreams: genericclioptions.NewTestIOStreamsDiscard()}
	secret, err := options.BundleSecret()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if secret.Type != corev1.SecretTypeDockercfg {
		t.Errorf("expected %v, got %v", kapi.SecretTypeDockercfg, secret.Type)
	}
}
