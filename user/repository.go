package user

import "errors"

type UserRepo struct {
	repo []*User
}

func NewRepo() UserRepo {
	return UserRepo{}
}

func (ur *UserRepo) Create(user User) (User, error) {
	id := 1
	if len(ur.repo) > 0 {
		id = ur.repo[len(ur.repo)-1].ID + 1
	}
	user.ID = id
	ur.repo = append(ur.repo, &user)
	return user, nil
}

var ErrUserNotFound = errUserNotFound()

func errUserNotFound() error {
	return errors.New("user not found")
}

func (ur *UserRepo) GetByUsername(username string) (User, error) {
	for i := range ur.repo {
		if ur.repo[i].Username == username {
			return *ur.repo[i], nil
		}
	}
	return User{}, ErrUserNotFound
}
