package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(user User) (User, error)
	GetByUsername(username string) (User, error)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return Service{repo: r}
}

func (s *Service) CreateUser(user User) (User, error) {
	hash, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	user.CreatedOn = time.Now().UTC()
	user, err := s.repo.Create(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Service) LoginUser(user User) (User, error) {
	foundUser, err := s.repo.GetByUsername(user.Username)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return User{}, err
	}

	return foundUser, nil
}
