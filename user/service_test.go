package user

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	repository := NewRepo()
	service := NewService(&repository)
	userIn := User{
		ID:       0,
		Username: "user",
		Password: "Password123",
		Image:    nil,
	}

	userOut, err := service.CreateUser(userIn)

	if err != nil {
		t.Error(err)
	}
	if userOut.ID == userIn.ID {
		t.Errorf("user's ID was not set. User ID = %d", userOut.ID)
	}
	if userOut.Password == userIn.Password {
		t.Errorf("user's password was not hashed")
	}
	if userOut.CreatedOn.IsZero() {
		t.Errorf("user's createdOn field was not set")
	}
}

func TestLoginUser_NotFound(t *testing.T) {
	repository := NewRepo()
	service := NewService(&repository)
	userIn := User{
		ID:       0,
		Username: "user",
		Password: "Password123",
		Image:    nil,
	}
	_, err := service.CreateUser(userIn)
	if err != nil {
		t.Error(err)
	}
	user := User{
		Username: "Random",
		Password: "Password123",
	}

	_, err = service.LoginUser(user)
	if err != ErrUserNotFound {
		t.Errorf("user should not be found")
	}
}

func TestLoginUser_Found(t *testing.T) {
	repository := NewRepo()
	service := NewService(&repository)
	userIn := User{
		ID:       0,
		Username: "user",
		Password: "Password123",
		Image:    nil,
	}
	_, err := service.CreateUser(userIn)
	if err != nil {
		t.Error(err)
	}
	user := User{
		Username: "user",
		Password: "Password123",
	}

	user, err = service.LoginUser(user)
	if err != nil {
		t.Errorf("user should be found")
	}
	if user.Username != userIn.Username {
		t.Errorf("found wrong user")
	}
}

func TestLoginUser_WrongPassword(t *testing.T) {
	repository := NewRepo()
	service := NewService(&repository)
	userIn := User{
		ID:       0,
		Username: "user",
		Password: "Password123",
		Image:    nil,
	}
	_, err := service.CreateUser(userIn)
	if err != nil {
		t.Error(err)
	}
	user := User{
		Username: "user",
		Password: "WrongPassword",
	}

	user, err = service.LoginUser(user)
	if err == nil {
		t.Errorf("wrong password should cause error")
	}
}

func TestLoginUser_RightPassword(t *testing.T) {
	repository := NewRepo()
	service := NewService(&repository)
	userIn := User{
		ID:       0,
		Username: "user",
		Password: "Password123",
		Image:    nil,
	}
	_, err := service.CreateUser(userIn)
	if err != nil {
		t.Error(err)
	}
	user := User{
		Username: "user",
		Password: "Password123",
	}

	user, err = service.LoginUser(user)
	if err != nil {
		t.Errorf("valid credentials should be accepted")
	}
}
