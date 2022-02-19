package user

import (
	"reflect"
	"strings"
	"testing"
)

func TestRegistrationRequestUsernameHasNecessaryValidationTags(t *testing.T) {
	expectedArgs := []string{"required", "gt=3", "lte=30", "alphanum"}
	r := registrationRequest{}
	rt := reflect.TypeOf(r)
	usernameField, ok := rt.FieldByName("Username")
	if !ok {
		t.Error("expected to have a 'Username' field")
	}
	validate, ok := usernameField.Tag.Lookup("validate")
	if !ok {
		t.Error("expected to have a 'validate' tag")
	}
	args := strings.Split(validate, ",")
	for i := range expectedArgs {
		matched := false
		for j := range args {
			if expectedArgs[i] == args[j] {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("expected to have validate argument %v", expectedArgs[i])
		}
	}
}

func TestRegistrationRequestPasswordHasNecessaryValidationTags(t *testing.T) {
	expectedArgs := []string{"required", "gt=8", "lte=128", "password"}
	r := registrationRequest{}
	rt := reflect.TypeOf(r)
	passwordField, ok := rt.FieldByName("Password")
	if !ok {
		t.Error("expected to have a 'Password' field")
	}
	validate, ok := passwordField.Tag.Lookup("validate")
	if !ok {
		t.Error("expected to have a 'validate' tag")
	}
	args := strings.Split(validate, ",")
	for i := range expectedArgs {
		matched := false
		for j := range args {
			if expectedArgs[i] == args[j] {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("expected to have validate argument %v", expectedArgs[i])
		}
	}
}

func TestLoginRequestUsernameHasNecessaryValidationTags(t *testing.T) {
	expectedArgs := []string{"required"}
	r := loginRequest{}
	rt := reflect.TypeOf(r)
	usernameField, ok := rt.FieldByName("Username")
	if !ok {
		t.Error("expected to have a 'Username' field")
	}
	validate, ok := usernameField.Tag.Lookup("validate")
	if !ok {
		t.Error("expected to have a 'validate' tag")
	}
	args := strings.Split(validate, ",")
	for i := range expectedArgs {
		matched := false
		for j := range args {
			if expectedArgs[i] == args[j] {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("expected to have validate argument %v", expectedArgs[i])
		}
	}
}

func TestLoginRequestPasswordHasNecessaryValidationTags(t *testing.T) {
	expectedArgs := []string{"required"}
	r := loginRequest{}
	rt := reflect.TypeOf(r)
	passwordField, ok := rt.FieldByName("Password")
	if !ok {
		t.Error("expected to have a 'Password' field")
	}
	validate, ok := passwordField.Tag.Lookup("validate")
	if !ok {
		t.Error("expected to have a 'validate' tag")
	}
	args := strings.Split(validate, ",")
	for i := range expectedArgs {
		matched := false
		for j := range args {
			if expectedArgs[i] == args[j] {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("expected to have validate argument %v", expectedArgs[i])
		}
	}
}
