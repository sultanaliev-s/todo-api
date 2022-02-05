package user

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
