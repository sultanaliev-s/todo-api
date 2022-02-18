package user

import (
	"testing"
)

func TestRegistrationRequest_RequiredUsernameAndPassword(t *testing.T) {
	req := registrationRequest{}
	err := req.Validate()
	if err == nil {
		t.Error("username and password should be required")
	}
}

func TestRegistrationRequest_RequiredUsername(t *testing.T) {
	req := registrationRequest{Password: "ValidPassword12345"}
	err := req.Validate()
	if err == nil {
		t.Error("username should be required")
	}
}

func TestRegistrationRequest_RequiredPassword(t *testing.T) {
	req := registrationRequest{Username: "ValidUsername12345"}
	err := req.Validate()
	if err == nil {
		t.Error("password should be required")
	}
}

func TestRegistrationRequest_UsernameNotAlphanumeric(t *testing.T) {
	req := registrationRequest{
		Username: "InvalidUsername12345!#",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err == nil {
		t.Error("username should be validated to be alphanumeric")
	}
}

func TestRegistrationRequest_UsernameAlphanumeric(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid username (%v) should be accepted %v", req.Username, err)
	}
}

func TestRegistrationRequest_UsernameLessMinLength(t *testing.T) {
	req := registrationRequest{
		Username: "Inv",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err == nil {
		t.Error("username should be validated to be longer than 3")
	}
}

func TestRegistrationRequest_UsernameMinLength(t *testing.T) {
	req := registrationRequest{
		Username: "Valid",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid username (%v) should be accepted %v", req.Username, err)
	}
}

func TestRegistrationRequest_UsernameGreaterMaxLength(t *testing.T) {
	req := registrationRequest{
		Username: "InvalidUsername012345678901234567809123456789",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err == nil {
		t.Error("username should be validated to be shorter than 31")
	}
}

func TestRegistrationRequest_UsernameMaxLength(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername0123456789",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid username (%v) should be accepted %v", req.Username, err)
	}
}

func TestRegistrationRequest_PasswordNotSecure(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "Invalid"}

	err := req.Validate()
	if err == nil {
		t.Error("password should be checked to contain " +
			"at least one letter and at least one number")
	}
}

func TestRegistrationRequest_PasswordSecure(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "ValidPassword12345!@"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid password (%v) should be accepted %v", req.Username, err)
	}
}

func TestRegistrationRequest_PasswordLessMinLength(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "Invalid"}

	err := req.Validate()
	if err == nil {
		t.Error("password should be validated to be longer than 8")
	}
}

func TestRegistrationRequest_PasswordMinLength(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "ValidPassword12345"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid password (%v) should be accepted %v", req.Username, err)
	}
}

func TestRegistrationRequest_PasswordGreaterMaxLength(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "TooLongPassword0123456789012345678901234567890123456789" +
			"012345678901234567890123456789012345678901234567890123456789" +
			"012345678901234567890123456789012345678901234567890123456789"}

	err := req.Validate()
	if err == nil {
		t.Error("password should be validated to be shorter than 129")
	}
}

func TestRegistrationRequest_PasswordMaxLength(t *testing.T) {
	req := registrationRequest{
		Username: "ValidUsername12345",
		Password: "ValidPassword0123456789012345678901234567890123456789" +
			"012345678901234567890123456789012345678901234567890123456789"}

	err := req.Validate()
	if err != nil {
		t.Errorf("valid password (%v) should be accepted %v", req.Username, err)
	}
}
