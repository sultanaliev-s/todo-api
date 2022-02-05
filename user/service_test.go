package user

import (
	"testing"
)

func TestBCreateUser(t *testing.T) {
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
