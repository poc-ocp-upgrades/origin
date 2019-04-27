package basicauthpassword

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
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
	expectedSubject := "12345"
	expectedName := "My Name"
	expectedEmail := "mylogin@example.com"
	expectedPreferredUsername := "myusername"
	data := fmt.Sprintf(`
	{
		"sub":"%s",
		"name": "%s",
		"email": "%s",
		"preferred_username": "%s",
		"additional_field": "should be ignored"
	}`, expectedSubject, expectedName, expectedEmail, expectedPreferredUsername)
	user := &RemoteUserData{}
	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user.Subject != expectedSubject {
		t.Errorf("Expected %s, got %s", expectedSubject, user.Subject)
	}
	if user.Name != expectedName {
		t.Errorf("Expected %s, got %s", expectedName, user.Name)
	}
	if user.Email != expectedEmail {
		t.Errorf("Expected %s, got %s", expectedEmail, user.Email)
	}
	if user.PreferredUsername != expectedPreferredUsername {
		t.Errorf("Expected %s, got %s", expectedPreferredUsername, user.PreferredUsername)
	}
}
