package user

import (
	"testing"
)

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
