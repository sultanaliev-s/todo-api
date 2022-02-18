package user

import "github.com/go-playground/validator"

func (r registrationRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("password", passwordValidation)
	return validate.Struct(r)
}

const specialCharacters string = `!@#$%^&*()_-+={[}]|\:;"'<,>.?/`

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return containsNumbersLettersLegals(password)
}

func containsNumbersLettersLegals(password string) bool {
	hasNums := false
	hasLetters := false
	for i := range password {
		if isLetter(password[i]) {
			hasLetters = true
		} else if isNumber(password[i]) {
			hasNums = true
		} else if !isSpecialCharacter(password[i]) {
			return false
		}
	}

	return hasLetters && hasNums
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func isSpecialCharacter(c byte) bool {
	for i := range specialCharacters {
		if c == specialCharacters[i] {
			return true
		}
	}
	return false
}
