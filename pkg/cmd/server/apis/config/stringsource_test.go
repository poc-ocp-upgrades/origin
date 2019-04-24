package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	encryptedData	= []byte(`-----BEGIN ENCRYPTED STRING-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,8db8d0db237a56c8ea8d3c5c21acc5b6

9JUj8ezx/3MDgiiWbnx9QA==
-----END ENCRYPTED STRING-----`)
	encryptingKey	= []byte(`-----BEGIN ENCRYPTING KEY-----
f3zQfReuhwI1BvBNglhZdjgSocKKqABwyGafHJcdORw=
-----END ENCRYPTING KEY-----`)
)

func TestStringSource(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	os.Setenv("TestStringSource_present_env", "envvalue")
	os.Setenv("TestStringSource_encrypted_env", string(encryptedData))
	emptyFile, err := ioutil.TempFile("", "empty_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(emptyFile.Name())
	fooFile, err := ioutil.TempFile("", "foo_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(emptyFile.Name())
	if err := ioutil.WriteFile(fooFile.Name(), []byte(`filevalue`), os.FileMode(0755)); err != nil {
		t.Fatal(err)
	}
	encryptedFile, err := ioutil.TempFile("", "encrypted_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(encryptedFile.Name())
	if err := ioutil.WriteFile(encryptedFile.Name(), encryptedData, os.FileMode(0755)); err != nil {
		t.Fatal(err)
	}
	validKeyFile, err := ioutil.TempFile("", "valid_key_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(validKeyFile.Name())
	if err := ioutil.WriteFile(validKeyFile.Name(), encryptingKey, os.FileMode(0755)); err != nil {
		t.Fatal(err)
	}
	invalidKeyFile, err := ioutil.TempFile("", "invalid_key_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(invalidKeyFile.Name())
	if err := ioutil.WriteFile(invalidKeyFile.Name(), []byte(`invalid key value`), os.FileMode(0600)); err != nil {
		t.Fatal(err)
	}
	testcases := map[string]struct {
		StringSource	StringSource
		ExpectedValue	string
		ExpectedError	string
	}{"empty": {StringSource: StringSource{}, ExpectedValue: "", ExpectedError: ""}, "value": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Value: "foo"}}, ExpectedValue: "foo", ExpectedError: ""}, "env empty": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Env: "empty_env"}}, ExpectedValue: "", ExpectedError: ""}, "env present": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Env: "TestStringSource_present_env"}}, ExpectedValue: "envvalue", ExpectedError: ""}, "file missing": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: "missing_file"}}, ExpectedValue: "", ExpectedError: "missing_file: no such file"}, "file empty": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: emptyFile.Name()}}, ExpectedValue: "", ExpectedError: ""}, "file present": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: fooFile.Name()}}, ExpectedValue: "filevalue", ExpectedError: ""}, "encrypted env with missing key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Env: "TestStringSource_encrypted_env", KeyFile: "missing_key"}}, ExpectedValue: "", ExpectedError: "missing_key: no such file"}, "encrypted env with invalid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Env: "TestStringSource_encrypted_env", KeyFile: invalidKeyFile.Name()}}, ExpectedValue: "", ExpectedError: "no valid PEM block"}, "encrypted env with valid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Env: "TestStringSource_encrypted_env", KeyFile: validKeyFile.Name()}}, ExpectedValue: "encryptedvalue", ExpectedError: ""}, "encrypted value with missing key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Value: string(encryptedData), KeyFile: "missing_key"}}, ExpectedValue: "", ExpectedError: "missing_key: no such file"}, "encrypted value with invalid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Value: string(encryptedData), KeyFile: invalidKeyFile.Name()}}, ExpectedValue: "", ExpectedError: "no valid PEM block"}, "encrypted value with valid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{Value: string(encryptedData), KeyFile: validKeyFile.Name()}}, ExpectedValue: "encryptedvalue", ExpectedError: ""}, "missing encrypted file with valid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: "missing_file", KeyFile: validKeyFile.Name()}}, ExpectedValue: "", ExpectedError: "missing_file: no such file"}, "encrypted file with missing key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: encryptedFile.Name(), KeyFile: "missing_key"}}, ExpectedValue: "", ExpectedError: "missing_key: no such file"}, "encrypted file with invalid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: encryptedFile.Name(), KeyFile: invalidKeyFile.Name()}}, ExpectedValue: "", ExpectedError: "no valid PEM block"}, "encrypted file with valid key": {StringSource: StringSource{StringSourceSpec: StringSourceSpec{File: encryptedFile.Name(), KeyFile: validKeyFile.Name()}}, ExpectedValue: "encryptedvalue", ExpectedError: ""}}
	for k, tc := range testcases {
		value, err := ResolveStringValue(tc.StringSource)
		if len(tc.ExpectedError) > 0 && (err == nil || !strings.Contains(err.Error(), tc.ExpectedError)) {
			t.Errorf("%s: expected error containing %q, got %q", k, tc.ExpectedError, err.Error())
		}
		if len(tc.ExpectedError) == 0 && err != nil {
			t.Errorf("%s: got unexpected error: %v", k, err)
		}
		if tc.ExpectedValue != value {
			t.Errorf("%s: Expected value=%q, got %q", k, tc.ExpectedValue, value)
		}
	}
}
