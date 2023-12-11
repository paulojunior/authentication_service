package service

import (
	contract "authentication/contract"
	"authentication/data"
	"errors"
)

type UserService struct {
	repository contract.UserRepository
}

func NewUserService(repo contract.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) FindByEmail(email string) (user data.User, err error) {
	if email == "" {
		return data.User{}, errors.New("invalid email")
	}

	user, err = s.repository.FindByEmail(email)
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}
