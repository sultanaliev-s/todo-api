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

func TestContainsNumbersLettersLegals_Fail(t *testing.T) {
	strs := []string{"InvalidPassword",
		"0123456789",
		"ПриветМир",
		"AlmostValid)НоНет"}
	expected := false
	res := make([]bool, len(strs))

	for i := range strs {
		res[i] = containsNumbersLettersLegals(strs[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf(
				"%v is not valid. Expected %v, got %v",
				strs[i], expected, res[i])
		}
	}
}

func TestContainsNumbersLettersLegals_Ok(t *testing.T) {
	strs := []string{"ValidPassword1",
		"0123456789A",
		"Validity123%$#",
		"propagandaisusele55"}
	expected := true
	res := make([]bool, len(strs))

	for i := range strs {
		res[i] = containsNumbersLettersLegals(strs[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf(
				"%v is valid. Expected %v, got %v",
				strs[i], expected, res[i])
		}
	}
}

func TestIsLetter_NotLetters(t *testing.T) {
	chars := []byte{'0', '1', '9', '@', '%', '&', '#'}
	expected := false
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isLetter(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf(
				"%c is not a letter. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}

func TestIsLetter_Letters(t *testing.T) {
	chars := []byte{'a', 'm', 'z', 'A', 'N', 'Z', 'Y'}
	expected := true
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isLetter(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf("%c is a letter. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}

func TestIsNumber_NotNumbers(t *testing.T) {
	chars := []byte{'a', 'z', 'A', 'Z', '%', '&', '#'}
	expected := false
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isNumber(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf(
				"%c is not a number. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}

func TestIsNumber_Numbers(t *testing.T) {
	chars := []byte{'0', '1', '2', '5', '6', '7', '9'}
	expected := true
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isNumber(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf("%c is a number. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}

func TestIsSpecialCharacter_NotSpecialCharacter(t *testing.T) {
	chars := []byte("abcdefghzABCDEFGHMKDNDSZ123456789")
	expected := false
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isSpecialCharacter(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf("%c is not a number. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}

func TestIsSpecialCharacter_SpecialCharacters(t *testing.T) {
	chars := []byte(`!@#$%^&*()_-+={[}]|\:;"'<,>.?/`)
	expected := true
	res := make([]bool, len(chars))

	for i := range chars {
		res[i] = isSpecialCharacter(chars[i])
	}

	for i := range res {
		if res[i] != expected {
			t.Errorf("%c is a special character. Expected %v, got %v",
				chars[i], expected, res[i])
		}
	}
}
